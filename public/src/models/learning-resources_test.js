import QUnit from 'steal-qunit';
import Learning-resources from './learning-resources';

QUnit.module('models/learning-resources');

QUnit.test('getList', function(){
  stop();
  Learning-resources.getList().then(function(items) {
    QUnit.equal(items.length, 2);
    QUnit.equal(items.attr('0.description'), 'First item');
    start();
  });
});
