package db

import "github.com/gophergala2016/wwcdc_01/models"

func AddScreenshot(screenshot *models.Screenshot) error {
	query := `insert into gophergala_reviews 
    (user_id, learning_resource_id, name, picture_url)
    values($1, $2, $3, $4)`
	_, err := db.Exec(query, screenshot.UserID, screenshot.LearningResourceID,
		screenshot.Name, screenshot.Url)
	if err != nil {
		return err
	}
	return nil
}
