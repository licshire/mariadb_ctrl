package upgrader

import (
	"errors"
	"regexp"
	"time"

	"code.cloudfoundry.org/lager"
	"github.com/cloudfoundry/mariadb_ctrl/config"
	"github.com/cloudfoundry/mariadb_ctrl/mariadb_helper"
	"github.com/cloudfoundry/mariadb_ctrl/os_helper"
)

//go:generate counterfeiter . Upgrader
type Upgrader interface {
	Upgrade() error
	NeedsUpgrade() (bool, error)
}

type upgrader struct {
	osHelper      os_helper.OsHelper
	config        config.Upgrader
	logger        lager.Logger
	mariadbHelper mariadb_helper.DBHelper
}

var (
	DBReachablePollingAttempts = 30
	DBReachablePollingDelay    = 10 * time.Second
)

func NewUpgrader(
	osHelper os_helper.OsHelper,
	config config.Upgrader,
	logger lager.Logger,
	mariadbHelper mariadb_helper.DBHelper) Upgrader {

	return upgrader{
		osHelper:      osHelper,
		config:        config,
		logger:        logger,
		mariadbHelper: mariadbHelper,
	}
}

func (u upgrader) Upgrade() (err error) {
	u.startStandaloneDatabaseSynchronously()

	u.logger.Info("Performing upgrade")
	output, upgrade_err := u.mariadbHelper.Upgrade()

	if upgrade_err != nil {
		acceptableErrorsCompiled, _ := regexp.Compile(
			"already upgraded|Unknown command|WSREP has not yet prepared node")

		if acceptableErrorsCompiled.MatchString(output) {
			u.logger.Info(
				"output string matches acceptable errors - continuing startup.",
				lager.Data{"upgradeErr": upgrade_err, "upgradeOutput": output},
			)
		} else {
			u.logger.Info(
				"output string does not match acceptable errors - aborting startup.",
				lager.Data{"upgradeErr": upgrade_err, "upgradeOutput": output},
			)
			err = upgrade_err
		}
	} else {
		u.logger.Info(
			"Upgrade applied successfully",
			lager.Data{"upgradeOutput": output},
		)
	}

	if err != nil {
		return
	}

	u.stopStandaloneDatabaseSynchronously()

	return
}

func (u upgrader) startStandaloneDatabaseSynchronously() {
	u.mariadbHelper.StartMysqldInStandAlone()

	for tries := 0; tries < DBReachablePollingAttempts; tries++ {
		if u.mariadbHelper.IsDatabaseReachable() {
			return
		}

		u.osHelper.Sleep(DBReachablePollingDelay)
	}

	u.logger.Fatal("Database is not reachable after 30 tries.", errors.New("Database is not reachable after 30 tries."))
}

func (u upgrader) stopStandaloneDatabaseSynchronously() {
	u.mariadbHelper.StopMysqld()
}

func (u upgrader) NeedsUpgrade() (bool, error) {
	if !u.osHelper.FileExists(u.config.LastUpgradedVersionFile) {
		u.logger.Info(
			"Upgrade required",
			lager.Data{
				"reason":                  "Last Upgraded version file does not exist in data dir",
				"lastUpgradedVersionFile": u.config.LastUpgradedVersionFile,
			})
		return true, nil
	}

	if !u.osHelper.FileExists(u.config.PackageVersionFile) {
		u.logger.Info(
			"Cannot determine whether upgrade is required.",
			lager.Data{
				"reason":             "Package version file does not exist",
				"packageVersionFile": u.config.PackageVersionFile,
			})
		return false, errors.New("MariaDB package is invalid because it is missing the version file.")
	}

	existing_version, err := u.osHelper.ReadFile(u.config.LastUpgradedVersionFile)
	if err != nil {
		u.logger.Info(
			"Cannot determine whether upgrade is required.",
			lager.Data{
				"reason":                  "Error reading last upgraded version file",
				"lastUpgradedVersionFile": u.config.LastUpgradedVersionFile,
				"err": err,
			})
		return false, errors.New("Could not read last upgraded version file in the data dir.")
	}

	package_version, err := u.osHelper.ReadFile(u.config.PackageVersionFile)
	if err != nil {
		u.logger.Info(
			"Cannot determine whether upgrade is required.",
			lager.Data{
				"reason":             "Error reading package version file",
				"packageVersionFile": u.config.PackageVersionFile,
				"err":                err,
			})
		return false, errors.New("MariaDB package is invalid because the version file is not readable.")
	}

	if existing_version != package_version {
		u.logger.Info("Need to upgrade to latest version.")
		return true, nil
	}
	u.logger.Info("Already upgraded to latest version, starting normally.")
	return false, nil
}
