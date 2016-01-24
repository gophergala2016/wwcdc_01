import fixture from 'can-fixture';

const store = fixture.store([{
  id: 0,
  description: 'First item'
}, {
  id: 1,
  description: 'Second item'
}]);

fixture({
  'GET /api/learning-resources/{id}/reviews': store.findAll,
  'GET /api/learning-resources/{id}/reviews/{id}': store.findOne,
  'POST /api/learning-resources/{id}/reviews': store.create,
  'PUT /api/learning-resources/{id}/reviews/{id}': store.update,
  'DELETE /api/learning-resources/{id}/reviews/{id}': store.destroy
});

export default store;
