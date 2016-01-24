import QUnit from 'steal-qunit';
import Languages from './languages';

QUnit.module('models/languages');

QUnit.test('getList', function(){
  stop();
  Languages.getList().then(function(items) {
    QUnit.equal(items.length, 2);
    QUnit.equal(items.attr('0.description'), 'First item');
    start();
  });
});
