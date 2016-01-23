import F from 'funcunit';
import QUnit from 'steal-qunit';

F.attach(QUnit);

QUnit.module('gophergala functional smoke test', {
  beforeEach() {
    F.open('../development.html');
  }
});

QUnit.test('gophergala main page shows up', function() {
  F('title').text('gophergala', 'Title is set');
});
