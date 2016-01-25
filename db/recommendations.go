package db

import (
  "github.com/gophergala2016/wwcdc_01/models"

  _ "gopkg.in/cq.v1"
)

func FindRecommendationsForLearningResource(id int64) (*models.Recommendations, error) {
  recommendations := models.Recommendations{}

  return &recommendations, nil
}
