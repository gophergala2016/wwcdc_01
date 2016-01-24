package db

import "github.com/gophergala2016/wwcdc_01/models"

func AddReview(review *models.Review) error {
	query := `insert into gophergala_reviews 
    (user_id, learning_resource_id, description, userfulness,
     course_cost, course_length, course_size, course_interaction, 
     finished, description, review_date) 
    values($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`
	_, err := db.Exec(query, review.UserID, review.LearningResourceID,
		review.Description, review.Usefulness)
	if err != nil {
		return err
	}
	return nil
}
