import fixture from 'can-fixture';

const store = fixture.store([{
  id: 0,
  description: 'First item'
}, {
  id: 1,
  description: 'Second item'
}]);

fixture({
  'GET /api/reviews': store.findAll,
  'GET /api/reviews/{id}': store.findOne,
  'POST /api/reviews': store.create,
  'PUT /api/reviews/{id}': store.update,
  'DELETE /api/reviews/{id}': store.destroy
});

export default store;
