package migrations

import (
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type Migration interface {
	Name() string
	Migrate(db *gorm.DB) error
	CheckCondition(db *gorm.DB) bool
}

type MigrationLock struct {
	gorm.Model
	Locked bool
}

type MigrationEntry struct {
	gorm.Model
	ID      uint   `gorm:"primaryKey"`
	Name    string `gorm:"index"`
	Type    string
	Checked bool
}

func InitMigrationTable(db *gorm.DB) error {
	var err error

	// Init Migration Lock table
	if !db.Migrator().HasTable(&MigrationLock{}) {
		err = db.Migrator().CreateTable(&MigrationLock{})
		log.Info("Created 'migration_lock' table")
		if err != nil {
			return err
		}

		// Init Migration Lock entry
		db.Create(&MigrationLock{
			Locked: false,
		})
	}

	// Init Migration Entry table
	if !db.Migrator().HasTable(&MigrationEntry{}) {
		err = db.Migrator().CreateTable(&MigrationEntry{})
		log.Info("Created 'migration_entry' table")
		if err != nil {
			return err
		}
	}

	return nil
}

func getAllMigrations() []Migration {
	var migrations = make([]Migration, 0)

	migrations = append(migrations, Migration_20220724155200{})

	return migrations
}

func DoMigration(db *gorm.DB) error {
	locker := &MigrationLock{}
	db.First(&locker)

	if locker.Locked {
		log.Fatal("Migration stopped due to being locked by other migration")
	}

	locker.Locked = true
	db.Save(locker)

	defer func() {
		locker.Locked = false
		db.Save(locker)
	}()

	err := db.Transaction(func(tx *gorm.DB) error {

		migrations := getAllMigrations()

		for _, migration := range migrations {
			name := migration.Name()

			migrationEntry := &MigrationEntry{}
			result := db.First(migrationEntry).Where("name = ?", name)

			if result.RowsAffected == 0 || migration.CheckCondition(tx) {
				err := migration.Migrate(tx)
				if err != nil {
					log.Warnf("Encounter an error, rolling back. Error details: %s", err.Error())
					return err
				}
				db.Create(&MigrationEntry{
					Name: name,
				})
			} else {
				log.Infof("Skipping %s...", name)
			}
		}

		log.Info("Complete migration")
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
