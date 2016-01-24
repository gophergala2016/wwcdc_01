package db

import (
	"io/ioutil"

	"github.com/go-contrib/uuid"
	"github.com/gophergala2016/wwcdc_01/restapi/operations"
)

func AddScreenshot(screenshot *operations.AddScreenshotParams) error {
	query := `insert into pictures
    (learning_resource_id, picture_url)
    values($1, $2)`
	buf, err := ioutil.ReadAll(screenshot.Screenshot.Data)
	if err != nil {
		return err
	}
	u1 := uuid.NewV4()
	ioutil.WriteFile("./public/screenshots/"+u1.String()+".png", buf, 0644)
	url := "https://boolmeover.com/screenshots/" + u1.String() + ".png"
	_, err = db.Exec(query, screenshot.ID, url)
	if err != nil {
		return err
	}
	return nil
}
