package repository

import "sync"

var (
	once          sync.Once
	dbEngine      *db.Engine
	dbrepoFactory *dbRepoFactory
)

func Init(c *config.Database) (*gorm.DB, error) {
	dbEngine = db.New(c)
	return dbEngine.Run()
}
