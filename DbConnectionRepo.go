package interfaces

import (
	"fmt"
	"log"

	Domain "github.com/Psinobious/Maser-DC/Domain"
)

func (repo *DbConnectionRepo) Store(connection Domain.Connection) error {
	statement := fmt.Sprintf(`INSERT INTO public."Connection" ("ClientID", "ActivityID", "Notifications") 
							VALUES('%s','%s','%b');`,
		connection.ClientID, connection.ActivityID, connection.Notifications)
	repo.mutex.Lock()
	err := repo.dbHandler.Execute(statement)
	repo.mutex.Unlock()

	if err != nil {
		fmt.Println(err)
	}
	return err
}
func (repo *DbConnectionRepo) Delete(ClientID string, ActivityID string) error {
	statement := fmt.Sprintf(`DELETE FROM public."Connection" WHERE "ClientID"= '%s' AND "ActivityID"='%s';`,
		ClientID, ActivityID)
	repo.mutex.Lock()
	err := repo.dbHandler.Execute(statement)
	repo.mutex.Unlock()

	if err != nil {
		fmt.Println(err)
	}
	return err
}
func (repo *DbConnectionRepo) Purge(ClientID string) error {
	statement := fmt.Sprintf(`DELETE FROM public."Connection" WHERE "ClientID"= '%s';`,
		ClientID)
	repo.mutex.Lock()
	err := repo.dbHandler.Execute(statement)
	repo.mutex.Unlock()

	if err != nil {
		fmt.Println(err)
	}
	return err
}
func (repo *DbConnectionRepo) Update(connection Domain.Connection) error {
	statement := fmt.Sprintf(`UPDATE public."Connection" SET "ClientID"='%s', "ActivityID" ='%s', "Notifications"='%s', "ConnectionList"='%s' WHERE "ContentID"= '%s';`,
		connection.ClientID, connection.ActivityID)
	repo.mutex.Lock()
	err := repo.dbHandler.Execute(statement)
	repo.mutex.Unlock()

	if err != nil {
		fmt.Println(err)
	}
	return err
}
func (repo *DbConnectionRepo) FindUsersById(ActivityID string) ([]Domain.Connection, error) {
	var connection Domain.Connection
	var list []Domain.Connection

	statement := fmt.Sprintf(`SELECT "ClientID", "ActivityID", "Notifications"
							  FROM public."Connection" WHERE "ActivityID" = '%s';`, ActivityID)

	repo.mutex.RLock()
	row, err := repo.dbHandler.Query(statement)

	for row.Next() {
		if err := row.Scan(&connection.ClientID, &connection.ActivityID, &connection.Notifications); err != nil {
			log.Fatal(err)
		}
		list = append(list, connection)
	}
	repo.mutex.RUnlock()
	return list, err
}
func (repo *DbConnectionRepo) FindActivitiesById(UserID string) ([]Domain.Connection, error) {
	var connection Domain.Connection
	var list []Domain.Connection

	statement := fmt.Sprintf(`SELECT "ClientID", "ActivityID", "Notifications"
							  FROM public."Connection" WHERE "UserID" = '%s';`, UserID)

	repo.mutex.RLock()
	row, err := repo.dbHandler.Query(statement)

	for row.Next() {
		if err := row.Scan(&connection.ClientID, &connection.ActivityID, &connection.Notifications); err != nil {
			log.Fatal(err)
		}
		list = append(list, connection)
	}
	repo.mutex.RUnlock()
	return list, err
}
func (repo *DbConnectionRepo) FindById(ClientID string, ActivityID string) (Domain.Connection, error) {
	var connection Domain.Connection

	statement := fmt.Sprintf(`SELECT "ClientID", "ActivityID", "Notifications"
							  FROM public."Connection" WHERE "UserID" = '%s' AND "ActivityID"='%s';`, ClientID)

	repo.mutex.RLock()
	row, err := repo.dbHandler.Query(statement)
	repo.mutex.RUnlock()

	if err := row.Scan(&connection.ClientID, &connection.ActivityID, &connection.Notifications); err != nil {
		log.Fatal(err)
	}
	return connection, err
}
