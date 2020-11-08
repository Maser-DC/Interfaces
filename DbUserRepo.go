package interfaces

import (
	"database/sql"
	"fmt"
	"log"

	Domain "github.com/Psinobious/Maser-DC/Domain"
)

func (repo *DbUserRepo) Store(client *Domain.Client) error {
	statement := fmt.Sprintf(`INSERT INTO public."User" ("userid", "password", "email", "first_name", "last_name") VALUES ('%s','%s','%s','%s','%s');`,
		client.ClientID, client.Password, client.Email, client.FirstName, client.LastName)

	repo.mutex.Lock()
	err := repo.dbHandler.Execute(statement)
	repo.mutex.Unlock()

	if err != nil {
		fmt.Println(err)
	}
	return err
}
func (repo *DbUserRepo) FindById(userID string) (Domain.Client, error) {
	var user Domain.Client
	statement := fmt.Sprintf(`SELECT * FROM public."User" WHERE "userid" = '%s';`, userID)

	repo.mutex.RLock()
	row, err := repo.dbHandler.Query(statement)
	repo.mutex.RUnlock()

	row.Next()
	err = row.Scan(&user)
	if err != nil && err == sql.ErrNoRows {
		log.Fatal("No Results Found")
	}

	return user, err
}
func (repo *DbUserRepo) CheckIfExist(userID string) bool {
	var user Domain.Client
	statement := fmt.Sprintf(`SELECT "userid" FROM public."User" WHERE "userid" = '%s';`, userID)

	repo.mutex.RLock()
	row, err := repo.dbHandler.Query(statement)
	repo.mutex.RUnlock()

	row.Next()
	err = row.Scan(&user)
	if err != nil && err == sql.ErrNoRows {
		return false
	}
	return true
}
func (repo *DbUserRepo) Delete(userID string) error {
	statement := fmt.Sprintf(`DELETE FROM public."User" WHERE "userid" = '%s';`, userID)
	repo.mutex.Lock()
	err := repo.dbHandler.Execute(statement)
	repo.mutex.Unlock()

	return err
}
func (repo *DbUserRepo) Update(client Domain.Client) error {
	statement := fmt.Sprint(`UPDATE public."User" SET "email" = '%s', "first_name" = '%s', "last_name" = '%s' WHERE "userid" = '%s';`,
		client.Email, client.FirstName, client.LastName, client.ClientID)
	repo.mutex.Lock()
	err := repo.dbHandler.Execute(statement)
	repo.mutex.Unlock()
	return err
}
