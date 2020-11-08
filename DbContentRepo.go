package interfaces

import (
	"fmt"
	"log"

	Domain "github.com/Psinobious/Maser-DC/Domain"
)

func (repo *DbContentRepo) FindById(contentID string) (Domain.Content, error) {
	var content Domain.Content

	statement := fmt.Sprintf(`SELECT "ContentID", "ActivityID", "Title", "Group", "Reference"
							  FROM public."Content" WHERE "ContentID" = '%s';`, contentID)

	repo.mutex.RLock()
	row, err := repo.dbHandler.Query(statement)
	repo.mutex.RUnlock()

	if err := row.Scan(&content.ContentID, &content.ActivityID, &content.Title, &content.Group, &content.Reference); err != nil {
		log.Fatal(err)
	}
	return content, err
}
func (repo *DbContentRepo) FincByActivityID(activityID string) ([]Domain.Content, error) {
	var content Domain.Content
	var list []Domain.Content

	statement := fmt.Sprintf(`SELECT "ContentID", "ActivityID", "Title", "Group", "Reference"
							  FROM public."Content" WHERE "ActivityID" = '%s';`, activityID)

	repo.mutex.RLock()
	row, err := repo.dbHandler.Query(statement)
	repo.mutex.RUnlock()

	for row.Next() {
		if err := row.Scan(&content.ContentID, &content.ActivityID, &content.Title, &content.Group, &content.Reference); err != nil {
			log.Fatal(err)
			list = append(list, content)
		}
	}
	if err != nil {
		fmt.Println(err)
	}
	return list, err
}
func (repo *DbContentRepo) Store(content Domain.Content) error {
	statement := fmt.Sprintf(`INSERT INTO public."Content" ("ContentID", "ActivityID", "Reference", "Group", "Title",) 
														VALUES('%s','%s','%s','%s','%s','%s');`,
		content.ContentID, content.ActivityID, content.Reference, content.Group, content.Title)

	repo.mutex.Lock()
	err := repo.dbHandler.Execute(statement)
	repo.mutex.Unlock()

	if err != nil {
		fmt.Println(err)
	}
	return err
}
func (repo *DbContentRepo) Remove(contentID string) error {
	statement := fmt.Sprintf(`DELETE FROM public."Content" WHERE "ContentID"= '%s';`, contentID)
	repo.mutex.Lock()
	err := repo.dbHandler.Execute(statement)
	repo.mutex.Unlock()

	if err != nil {
		fmt.Println(err)
	}
	return err
}
func (repo *DbContentRepo) Update(content Domain.Content) error {
	statement := fmt.Sprintf(`UPDATE public."Content" SET "ActivityID"='%s', "Title" ='%s', "Group"='%s', "Reference"='%s' WHERE "ContentID"= '%s';`,
		content.ActivityID, content.Title, content.Group, content.Reference, content.ContentID)
	repo.mutex.Lock()
	err := repo.dbHandler.Execute(statement)
	repo.mutex.Unlock()

	if err != nil {
		fmt.Println(err)
	}
	return err
}
