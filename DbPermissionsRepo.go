package interfaces

import (
	"database/sql"
	"fmt"
	"log"
)

func (repo *DbPermissionsRepo) AddPermission(RoleID string, UserID string) error {
	statement := fmt.Sprintf(`INSERT INTO public."Permissions" (RoleID, UserID) VALUES (%s,%s);`, RoleID, UserID)
	repo.mutex.Lock()
	err := repo.dbHandler.Execute(statement)
	repo.mutex.Unlock()
	if err != nil {
		log.Println(err)
	}
	return err
}
func (repo *DbPermissionsRepo) CheckPermission(userID string, roleID string) (bool, error) {
	statement := fmt.Sprintf(`SELECT "RoleID" FROM public."Permissions" WHERE "userid" = '%s' AND "roleid" = '%s'`, userID)

	repo.mutex.RLock()
	err := repo.dbHandler.Execute(statement)
	repo.mutex.RUnlock()

	if err != nil && err == sql.ErrNoRows {
		return false, err
	}
	return true, err
}
func (repo *DbPermissionsRepo) Permissions(userID string) ([]string, error) {
	statement := fmt.Sprintf(`SELECT "RoleID" FROM public."Permissions" WHERE "userid" = '%s'`, userID)
	var list []string
	var permission string

	repo.mutex.RLock()
	row, err := repo.dbHandler.Query(statement)
	for row.Next() {
		if err := row.Scan(&permission); err != nil {
			log.Fatal(err)
		}
		list = append(list, permission)
	}
	repo.mutex.RUnlock()
	if err != nil {
		log.Println(err)
	}
	return list, err
}
func (repo *DbPermissionsRepo) DeletePermission(RoleID string, UserID string) error {
	statement := fmt.Sprintf(`REMOVE FROM public."Permissions" WHERE "RoleID" = '%s' AND "UserID" = '%s';`, RoleID, UserID)
	repo.mutex.Lock()
	err := repo.dbHandler.Execute(statement)
	repo.mutex.Unlock()
	if err != nil {
		log.Println(err)
	}
	return err
}
func (repo *DbPermissionsRepo) Purge(userID string) error {
	statement := fmt.Sprintf(`REMOVE FROM public."Permissions" WHERE "UserID" = '%s';`, userID)

	repo.mutex.Lock()
	err := repo.dbHandler.Execute(statement)
	repo.mutex.Unlock()

	if err != nil {
		log.Println(err)
	}
	return err
}
