import QUnit from 'steal-qunit';
import { ViewModel } from './reviews-list-item';

// ViewModel unit tests
QUnit.module('gophergala/components/reviews-list-item');

QUnit.test('Has message', function(){
  var vm = new ViewModel();
  QUnit.equal(vm.attr('message'), 'This is the reviews-list-item component');
});
