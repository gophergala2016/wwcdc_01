import fixture from 'can-fixture';

const store = fixture.store([ {
    "id": 0,
    "user_id": 0,
    "learning_resource_id": 9,
    "usefulness": 0,
    "finished": true,
    "cost": 0,
    "course_size": 0,
    "course_length": 0,
    "course_interaction": 0,
    "description": "string",
    "review_date": "2016-01-24"
  },  {
    "id": 0,
    "user_id": 0,
    "learning_resource_id": 9,
    "usefulness": 0,
    "finished": true,
    "cost": 0,
    "course_size": 0,
    "course_length": 0,
    "course_interaction": 0,
    "description": "string",
    "review_date": "2016-01-24"
  }]);

fixture({
  'GET /api/learning-resources/{id}/reviews': store.findAll,
  'GET /api/learning-resources/{id}/reviews/{id}': store.findOne,
  'POST /api/learning-resources/{id}/reviews': store.create,
  'PUT /api/learning-resources/{id}/reviews/{id}': store.update,
  'DELETE /api/learning-resources/{id}/reviews/{id}': store.destroy
});

export default store;
