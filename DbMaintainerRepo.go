package interfaces

import (
	"database/sql"
	"fmt"
	"log"

	Domain "github.com/Psinobious/Maser-DC/Domain"
)

func (repo *DbMaintainerRepo) FindUsers(ActivityID string) ([]Domain.Maintainer, error) {
	var list []Domain.Maintainer
	var Maintainer Domain.Maintainer

	statement := fmt.Sprintf(`SELECT "UserID", "ActivityID", "RoleID"
							FROM public."Maintainer' WHERE "ActivityID"='%s';`)
	repo.mutex.RLock()
	row, err := repo.dbHandler.Query(statement)
	for row.Next() {
		if err := row.Scan(&Maintainer.UserID, &Maintainer.ActivityID, &Maintainer.RoleID); err != nil {
			log.Fatal(err)
		}
		list = append(list, Maintainer)
	}
	repo.mutex.RUnlock()
	if err != nil {
		log.Fatal(err)
	}
	return list, err
}
func (repo *DbMaintainerRepo) FindActivities(UserID string) ([]Domain.Maintainer, error) {
	var list []Domain.Maintainer
	var Maintainer Domain.Maintainer

	statement := fmt.Sprintf(`SELECT "UserID", "ActivityID", "RoleID"
							FROM public."Maintainer' WHERE "UserID"='%s';`)
	repo.mutex.RLock()
	row, err := repo.dbHandler.Query(statement)
	for row.Next() {
		if err := row.Scan(&Maintainer.UserID, &Maintainer.ActivityID, &Maintainer.RoleID); err != nil {
			log.Fatal(err)
		}
		list = append(list, Maintainer)
	}
	repo.mutex.RUnlock()
	if err != nil {
		log.Fatal(err)
	}
	return list, err
}
func (repo *DbMaintainerRepo) FindById(UserID string, ActivityID string) (Domain.Maintainer, error) {
	var maintainer Domain.Maintainer
	statement := fmt.Sprintf(`SELECT "UserID", "ActivityID", "RoleID" FROM public."Maintainer" 
							 WHERE "UserID"='%s' AND "ActivityID"='%s'`,
		UserID, ActivityID)
	repo.mutex.RLock()
	row, err := repo.dbHandler.Query(statement)
	repo.mutex.RUnlock()
	if err != nil {
		fmt.Println(err)
	}
	row.Next()
	err = row.Scan(&maintainer.UserID, &maintainer.ActivityID, &maintainer.RoleID)
	if err != nil && err == sql.ErrNoRows {
		log.Fatal("No Results Found")
	}
	return maintainer, err
}
func (repo *DbMaintainerRepo) Store(maintainer Domain.Maintainer) error {
	statement := fmt.Sprintf(`INSERT INTO public."Maintainer"("UserID","ActivityID","RoleID") 
							VALUES('%s','%s','%s')`,
		maintainer.UserID, maintainer.ActivityID, maintainer.RoleID)
	repo.mutex.Lock()
	err := repo.dbHandler.Execute(statement)
	repo.mutex.Unlock()

	if err != nil {
		fmt.Println(err)
	}
	return err
}
func (repo *DbMaintainerRepo) Remove(UserID string, ActivityID string) error {
	statement := fmt.Sprintf(`DELETE FROM public."Maintainer" WHERE "UserID"='%s' AND "ActivityID"='%s'`,
		UserID, ActivityID)
	repo.mutex.Lock()
	err := repo.dbHandler.Execute(statement)
	repo.mutex.Unlock()

	return err
}
func (repo *DbMaintainerRepo) Purge(UserID string) error {
	statement := fmt.Sprintf(`DELETE FROM public."Maintainer" WHERE "UserID"='%s'`,
		UserID)
	repo.mutex.Lock()
	err := repo.dbHandler.Execute(statement)
	repo.mutex.Unlock()

	return err
}
func (repo *DbMaintainerRepo) ChangeRole(maintainer Domain.Maintainer) error {
	statement := fmt.Sprintf(`UPDATE public."Maintainer" SET "RoleID" ='%s' WHERE "UserID"= '%s' AND "ActivityID"='%s';`,
		maintainer.RoleID, maintainer.UserID, maintainer.ActivityID)
	repo.mutex.Lock()
	err := repo.dbHandler.Execute(statement)
	repo.mutex.Unlock()

	if err != nil {
		fmt.Println(err)
	}
	return err
}
