import QUnit from 'steal-qunit';
import { ViewModel } from './resource-detail';

// ViewModel unit tests
QUnit.module('gophergala/components/resource-detail');

QUnit.test('Has message', function(){
  var vm = new ViewModel();
  QUnit.equal(vm.attr('message'), 'This is the resource-detail component');
});
