import QUnit from 'steal-qunit';
import { ViewModel } from './create-review';

// ViewModel unit tests
QUnit.module('gophergala/components/create-review');

QUnit.test('Has message', function(){
  var vm = new ViewModel();
  QUnit.equal(vm.attr('message'), 'This is the create-review component');
});
