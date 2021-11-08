package validate

import (
	"fmt"
	"sort"

	"join_table/pkg/model"

	"github.com/google/go-cmp/cmp"
	"gorm.io/gorm"
)

func All(db *gorm.DB) error {
	models := []interface{}{model.Kiosk{}, model.Event{}, model.KioskEvent{}, model.Checkin{}}
	for i := range models {
		err := AssertModelMigration(db, models[i])
		if err != nil {
			return err
		}
	}

	return nil
}

func AssertModelMigration(db *gorm.DB, model interface{}) error {
	var existing []InformationSchemaResult

	db.Raw(`SELECT * FROM information_schema.columns WHERE table_schema = ?`, "checkin").Scan(&existing)
	tx := db.Begin()

	err := tx.AutoMigrate(model)
	if err != nil {
		return err
	}

	var postMigration []InformationSchemaResult
	err = tx.Raw(`SELECT * FROM information_schema.columns WHERE table_schema = ?`, "checkin").Scan(&postMigration).Error
	if err != nil {
		return err
	}
	err = tx.Rollback().Error
	if err != nil {
		return err
	}
	sort.Slice(existing, tableCmp(existing))
	sort.Slice(postMigration, tableCmp(postMigration))
	diff := cmp.Diff(existing, postMigration)
	if diff != "" {
		return fmt.Errorf("existing does not match post migration, db changes detected %s", diff)
	}
	return nil
}

func tableCmp(results []InformationSchemaResult) func(a int, b int) bool {
	return func(a, b int) bool {
		if results[a].TableName == results[b].TableName {
			return results[a].ColumnName < results[b].ColumnName
		}
		return results[a].TableName < results[b].TableName
	}
}

type InformationSchemaResult struct {
	gorm.Model
	TableName   string
	TableSchema string
	ColumnName  string
}
