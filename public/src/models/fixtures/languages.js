import fixture from 'can-fixture';

const store = fixture.store([{
  id: 0,
  description: 'First item'
}, {
  id: 1,
  description: 'Second item'
}]);

fixture({
  'GET api/languages': store.findAll,
  'GET api/languages/{id}': store.findOne,
  'POST api/languages': store.create,
  'PUT api/languages/{id}': store.update,
  'DELETE api/languages/{id}': store.destroy
});

export default store;
