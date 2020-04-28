package migrations

import (
	errs2 "github.com/breathbath/go_utils/utils/errs"
	baseDb "github.com/breathbath/go_utils/utils/sqlDb"
	log "github.com/sirupsen/logrus"
	"strings"
)

type Registry struct {
	items []Item
	dbg *baseDb.DbGateway
}

func NewRegistry(dbg *baseDb.DbGateway) *Registry {
	return &Registry{
		items: []Item{},
		dbg: dbg,
	}
}

func (mr *Registry) migrateItem(item Item) error {
	if len(item.GetQueries()) == 0 {
		log.Info("Nothing to migrate")
		return nil
	}

	log.Infof("Starting migration %s", item.GetId())
	tx, err := mr.dbg.Begin()
	if err != nil {
		return err
	}

	log.Infof("Will execute %d queries", len(item.GetQueries()))
	for _, q := range item.GetQueries() {
		log.Info("Executing\n'%s'", q.sql)
		_, err := tx.Exec(q.sql, q.args...)
		errs := errs2.NewErrorContainer()
		if err != nil {
			errs.AddError(err)
			err = tx.Rollback()
			errs.AddError(err)

			return errs.Result(" ")
		}
	}

	return tx.Commit()
}

func (mr *Registry) migrationWasExecuted(uid string) (bool, error) {
	var id string
	return mr.dbg.ScanScalarByQuery(&id, "SELECT uid FROM migrations WHERE uid=?", uid)
}

func (mr *Registry) Execute() error {
	migratedCount := 0
	for _, mgrItem := range mr.items {
		migrationWasExecuted, err := mr.migrationWasExecuted(mgrItem.GetId())
		if err != nil {
			if strings.Contains(err.Error(), `Table 'visits.migrations' doesn't exist`) {
				log.Info("Table 'migrations' doesn't exist will create a new one")
				_, err = mr.dbg.Exec(
					`CREATE TABLE migrations (
  uid varchar(155) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  PRIMARY KEY (uid),
  UNIQUE KEY uid (uid)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;`,
				)
				if err != nil {
					return err
				}
			} else {
				return err
			}
		}
		if migrationWasExecuted {
			continue
		}

		err = mr.migrateItem(mgrItem)
		if err != nil {
			return err
		}

		migratedCount++
		_, err = mr.dbg.Exec(
			"INSERT IGNORE INTO migrations (uid) VALUES (?)",
			mgrItem.GetId(),
		)
		if err != nil {
			return err
		}

		log.Info("Will execute migration post query callback")
		err = mgrItem.GetPostQueryCallback()()
		if err != nil {
			return err
		}
		log.Infof("Finished %s", mgrItem.GetId())
	}

	if migratedCount == 0 {
		log.Info("No migrations to execute")
	} else {
		log.Infof("Executed %d migrations", migratedCount)
	}

	return nil
}

func (mr *Registry) RegisterMigration(item Item) {
	mr.items = append(mr.items, item)
}
