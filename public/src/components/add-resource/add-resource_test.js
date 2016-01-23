import QUnit from 'steal-qunit';
import { ViewModel } from './add-resource';

// ViewModel unit tests
QUnit.module('gophergala/components/add-resource');

QUnit.test('Has message', function(){
  var vm = new ViewModel();
  QUnit.equal(vm.attr('message'), 'This is the add-resource component');
});
