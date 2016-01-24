import fixture from 'can-fixture';

const store = fixture.store([{
  "id":9,
  "name":"Geek for Geeks",
  "type":"online",
  "url":"http://www.geeksforgeeks.org/fundamentals-of-algorithms/",
  "languages": ["go", "js"]
}, {
  "id":8,
  "name":"Another One",
  "type":"online",
  "url":"http://www.geeksforgeeks.org/fundamentals-of-algorithms/",
  "languages": ["python", "ruby", "erlang"]
}]);

fixture({
  'GET /api/learning-resources': store.findAll,
  'GET /api/learning-resources/{id}': store.findOne,
  'POST /api/learning-resources': store.create,
  'PUT /api/learning-resources/{id}': store.update,
  'DELETE /api/learning-resources/{id}': store.destroy
});

export default store;
