package db

import (
	"fmt"
	"log"

	"github.com/gophergala2016/wwcdc_01/models"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

var db = sqlx.MustConnect("pgx", "postgres://gophergala:gophergala1@localhost:5432/gophergala?sslmode=disable")

func FindLearningResources(t string) ([]*models.LearningResource, error) {
	log.Println("querying learning resources on", t)
	query := ""
	ret := []*models.LearningResource{}
	tx := db.MustBegin()
	var err error
	if t == "all" {
		query = `select lr.name as name, lr.url as url, lr.id as id, t.name as type
             from gophergala_learning_resources lr, gophergala_types t 
             where lr.type_id = t.id`
		err = tx.Select(&ret, query)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		query = `select lr.name as name, lr.url as url, lr.id as id, t.name as type
             from gophergala_learning_resources lr, gophergala_types t 
             where t.name=$1 and lr.type_id = t.id`
		err = tx.Select(&ret, query, t)
		if err != nil {
			fmt.Println(err)
		}
	}
	log.Println("first query len:", len(ret), ret)
	log.Println("starting second query set..")
	query = `select l.name as name
            from gophergala_languages l, 
                 gophergala_learning_resource_languages lrl
            where l.id = lrl.language_id and lrl.learning_resource_id = $1`
	for _, x := range ret {
		err = tx.Select(&(x.Languages), query, x.ID)
		if err != nil {
			fmt.Println(err)
		}
	}
	log.Println("after second query:", ret)
	for _, x := range ret {
		fmt.Println(*x)
	}
	tx.Commit()
	return ret, nil
}

func FindLearningResourceByID(id int64) (*models.LearningResource, error) {
	query := `select lr.name as name, lr.url as url, lr.id as id, t.name as type
             from gophergala_learning_resources lr, gophergala_types t 
             where lr.id=$1 and lr.type_id = t.id`
	ret := models.LearningResource{}
	err := db.Get(&ret, query, id)
	if err != nil {
		fmt.Println(err)
	}
	query = `select l.name as name
            from gophergala_languages l, 
                 gophergala_learning_resource_languages lrl
            where l.id = lrl.language_id and lrl.learning_resource_id = $1`
	err = db.Select(&ret.Languages, query, id)
	return &ret, nil
}

func CreateLearningResource(lr *models.LearningResource) error {
	var typeId int64
	var languageIds []int64
	tx := db.MustBegin()
	query := "select id from gophergala_types where name=$1;"
	err := tx.Get(&typeId, query, lr.Type)
	if err != nil {
		return err
	}
	query = "select id from gophergala_languages where name in (?);"
	query, args, err := sqlx.In(query, lr.Languages)
	if err != nil {
		return err
	}
	query = db.Rebind(query)
	err = tx.Select(&languageIds, query, args...)
	if err != nil {
		return err
	}
	query = "insert into gophergala_learning_resources (name, url, type_id) values($1, $2, $3);"
	_, err = tx.Exec(query, lr.Name, lr.URL, typeId)
	if err != nil {
		return err
	}
	// hacky ... sometimes causes weird dups if you put the same one in twice
	query = "select id from gophergala_learning_resources where name=$1;"
	var lastId int64
	err = tx.Get(&lastId, query, lr.Name)
	if err != nil {
		return err
	}
	for _, languageId := range languageIds {
		query = "insert into gophergala_learning_resource_languages (learning_resource_id, language_id) values($1, $2);"
		_, err = tx.Exec(query, lastId, languageId)
		if err != nil {
			return err
		}
	}
	err = tx.Commit()
	return err
}
