import QUnit from 'steal-qunit';
import Reviews from './reviews';

QUnit.module('models/reviews');

QUnit.test('getList', function(){
  stop();
  Reviews.getList().then(function(items) {
    QUnit.equal(items.length, 2);
    QUnit.equal(items.attr('0.description'), 'First item');
    start();
  });
});
