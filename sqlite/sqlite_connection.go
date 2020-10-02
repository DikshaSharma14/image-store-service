package sim_sqlite

import (
	. "image-store-service/logger"
	"image-store-service/model"
	"path/filepath"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var SqliteInstance *gorm.DB

// Get the SqliteInstance of SQLite database. The database is served by a file named offline.db. If the database is already
// created then the existing database SqliteInstance is returned else a new database SqliteInstance is created. Return error in case
// of any error occurs
func init() {
	if SqliteInstance == nil {
		Logger.Info("Initializing SQLite Instance")

		db, err := gorm.Open("sqlite3", filepath.Join(".", "image-store.db"))
		if err != nil {
			Logger.Critical("Failed to initialize SQLite Instance", err)
		}
		db.LogMode(true)
		db.SingularTable(true)
		if !db.HasTable(&model.Album{}) {
			db.CreateTable(&model.Album{})
		}
		if !db.HasTable(&model.Image{}) {
			db.CreateTable(&model.Image{})
			//db.Model(&model.Image{}).AddForeignKey("album_id", "album(id)", "CASCADE", "CASCADE")
		}
		SqliteInstance = db
		Logger.Info("Initialized SQLite Instance")
	}
}
