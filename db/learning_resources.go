package db

import (
	"github.com/gophergala2016/wwcdc_01/models"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

var db = sqlx.MustConnect("pgx", "postgres://gophergala:gophergala1@localhost:5432/gophergala?sslmode=disable")

func CreateLearningResource(lr *models.LearningResource) error {
	var typeId int64
	var languageIds []int64
	tx := db.MustBegin()
	query := "select id from gophergala_types where name=$1;"
	err := tx.Get(&typeId, query, lr.Type)
	if err != nil {
		return err
	}
	query = "SELECT id FROM gophergala_languages WHERE name IN (?);"
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
