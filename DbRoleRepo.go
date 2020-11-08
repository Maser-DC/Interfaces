package interfaces

import (
	"database/sql"
	"fmt"
	"log"

	Domain "github.com/Psinobious/Maser-DC/Domain"
)

func (repo *DbRoleRepo) Store(role Domain.Role) error {
	statement := fmt.Sprintf(`INSERT INTO public."Role" ("RoleID", "Title", "Description", "RoleType") VALUES ('%s','%s','%s', '%s');`,
		role.RoleID, role.Title, role.Description, role.RoleType)

	repo.mutex.Lock()
	err := repo.dbHandler.Execute(statement)
	repo.mutex.Unlock()

	if err != nil {
		fmt.Println(err)
	}
	return err
}
func (repo *DbRoleRepo) FindById(roleID string) (Domain.Role, error) {
	var role Domain.Role

	statement := fmt.Sprintf(`SELECT * FROM public."Role" WHERE "RoleID" = '%s';`, roleID)
	repo.mutex.RLock()
	row, err := repo.dbHandler.Query(statement)
	repo.mutex.RUnlock()
	if err != nil {
		fmt.Println(err)
	}
	row.Next()
	err = row.Scan(&role)
	if err != nil && err == sql.ErrNoRows {
		log.Fatal("No Results Found")
	}
	return role, err
}
func (repo *DbRoleRepo) FindByRoleTypes(RoleType string) ([]Domain.Role, error) {
	var list []Domain.Role
	var role Domain.Role

	statement := fmt.Sprintf(`SELECT "RoleID", "Title", "Description", "RoleType" FROM public."Role" WHERE "RoleType" = '%s';`, RoleType)

	repo.mutex.RLock()
	row, err := repo.dbHandler.Query(statement)
	for row.Next() {
		if err := row.Scan(&role.RoleID, &role.Title, &role.Description, &role.RoleType); err != nil {
			log.Fatal(err)
		}
		list = append(list, role)
	}
	repo.mutex.RUnlock()

	if err != nil {
		log.Fatal(err)
	}

	return list, err
}
func (repo *DbRoleRepo) RemoveRole(roleID string) error {
	statement := fmt.Sprintf(`DELETE FROM public."Role" WHERE "RoleID" = '%s';`, roleID)
	repo.mutex.Lock()
	err := repo.dbHandler.Execute(statement)
	repo.mutex.Unlock()

	if err != nil {
		fmt.Println(err)
	}
	return err
}
func (repo *DbRoleRepo) UpdateRole(role Domain.Role) error {
	statement := fmt.Sprintf(`UPDATE public."Role" SET "RoleID" = '%s', "Title" ='%s', "Description"='%s', "RoleType"='s' WHERE "RoleID"='%s';`,
		role.Title, role.Description, role.RoleType, role.RoleID)

	repo.mutex.Lock()
	err := repo.dbHandler.Execute(statement)
	repo.mutex.Unlock()

	if err != nil {
		fmt.Println(err)
	}
	return err
}
