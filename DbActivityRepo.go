package interfaces

import (
	"fmt"
	"log"

	Domain "github.com/Psinobious/Maser-DC/Domain"
)

func (repo *DbActivityRepo) Store(activity Domain.Activity) error {
	statement := fmt.Sprintf(`INSERT INTO public."Activity" ("ActivityID","Title","ActivityType") 
							VALUES('%s','%s','%s');`,
		activity.ActivityID, activity.Title, activity.ActivityType)
	repo.mutex.Lock()
	err := repo.dbHandler.Execute(statement)
	repo.mutex.Unlock()

	if err != nil {
		fmt.Println(err)
	}
	return err
}
func (repo *DbActivityRepo) Delete(ActivityID string) error {
	statement := fmt.Sprintf(`DELETE FROM public."Activity" WHERE "ActivityID"= '%s';`)
	repo.mutex.Lock()
	err := repo.dbHandler.Execute(statement)
	repo.mutex.Unlock()

	if err != nil {
		fmt.Println(err)
	}
	return err
}
func (repo *DbActivityRepo) FindById(ActivityID string) (Domain.Activity, error) {
	var activity Domain.Activity

	statement := fmt.Sprintf(`SELECT "ActivityID", "Title", "ActivityType"
							  FROM public."Activity" WHERE "ActivityID" = '%s';`, ActivityID)

	repo.mutex.RLock()
	row, err := repo.dbHandler.Query(statement)
	repo.mutex.RUnlock()

	if err := row.Scan(&activity.ActivityID, &activity.Title, &activity.ActivityType); err != nil {
		log.Fatal(err)
	}
	return activity, err
}
func (repo *DbActivityRepo) Update(activity Domain.Activity) error {
	statement := fmt.Sprintf(`UPDATE public."Activity" SET "Title"='%s',"ActivityType"='%s'
							WHERE "ActivityID"='%s'`,
		activity.Title, activity.ActivityType, activity.ActivityID)
	repo.mutex.Lock()
	err := repo.dbHandler.Execute(statement)
	repo.mutex.Unlock()

	if err != nil {
		fmt.Println(err)
	}
	return err
}
