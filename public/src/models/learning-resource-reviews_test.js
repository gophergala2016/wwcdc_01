import QUnit from 'steal-qunit';
import LearningResourceReviews from './learning-resource-reviews';

QUnit.module('models/learning-resource-reviews');

QUnit.test('getList', function(){
  stop();
  LearningResourceReviews.getList().then(function(items) {
    QUnit.equal(items.length, 2);
    QUnit.equal(items.attr('0.description'), 'First item');
    start();
  });
});
