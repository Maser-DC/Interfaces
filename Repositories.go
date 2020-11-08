package interfaces

import "sync"

type DBHandler interface {
	Execute(statement string) error
	Query(statement string) (Row, error)
}
type Row interface {
	Scan(dest ...interface{}) error
	Next() bool
}
type DbRepo struct {
	dbHandlers map[string]DBHandler
	dbHandler  DBHandler
	mutex      sync.RWMutex
}

type DbUserRepo DbRepo
type DbPermissionsRepo DbRepo
type DbRoleRepo DbRepo

type DbActivityRepo DbRepo
type DbConnectionRepo DbRepo
type DbContentRepo DbRepo
type DbMaintainerRepo DbRepo

func NewDbUserRepo(dbHandlers map[string]DBHandler) *DbUserRepo {
	dbUserRepo := new(DbUserRepo)
	dbUserRepo.dbHandlers = dbHandlers
	dbUserRepo.dbHandler = dbHandlers["DbUserRepo"]
	return dbUserRepo
}
func NewDbPermissionsRepo(dbHandlers map[string]DBHandler) *DbPermissionsRepo {
	dbPermissionsRepo := new(DbPermissionsRepo)
	dbPermissionsRepo.dbHandlers = dbHandlers
	dbPermissionsRepo.dbHandler = dbHandlers["DbPermissionsRepo"]
	return dbPermissionsRepo
}
func NewDbRoleRepo(dbHandlers map[string]DBHandler) *DbRoleRepo {
	dbRoleRepo := new(DbRoleRepo)
	dbRoleRepo.dbHandlers = dbHandlers
	dbRoleRepo.dbHandler = dbHandlers["DbRoleRepo"]
	return dbRoleRepo
}
func NewDbActivityRepo(dbHandlers map[string]DBHandler) *DbActivityRepo {
	dbActivityRepo := new(DbActivityRepo)
	dbActivityRepo.dbHandlers = dbHandlers
	dbActivityRepo.dbHandler = dbHandlers["DbActivityRepo"]
	return dbActivityRepo
}
func NewDbMaintainerRepo(dbHandlers map[string]DBHandler) *DbMaintainerRepo {
	dbMaintainerRepo := new(DbMaintainerRepo)
	dbMaintainerRepo.dbHandlers = dbHandlers
	dbMaintainerRepo.dbHandler = dbHandlers["DbMaintainerRepo"]
	return dbMaintainerRepo
}
func NewDbContentRepo(dbHandlers map[string]DBHandler) *DbContentRepo {
	dbContentRepo := new(DbContentRepo)
	dbContentRepo.dbHandlers = dbHandlers
	dbContentRepo.dbHandler = dbHandlers["DbContentRepo"]
	return dbContentRepo
}
func NewDbConnectionRepo(dbHandlers map[string]DBHandler) *DbConnectionRepo {
	dbConnectionRepo := new(DbConnectionRepo)
	dbConnectionRepo.dbHandlers = dbHandlers
	dbConnectionRepo.dbHandler = dbHandlers["DbConnectionRepo"]
	return dbConnectionRepo
}
