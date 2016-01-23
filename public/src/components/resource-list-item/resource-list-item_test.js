import QUnit from 'steal-qunit';
import { ViewModel } from './resource-list-item';

// ViewModel unit tests
QUnit.module('gophergala/components/resource-list-item');

QUnit.test('Has message', function(){
  var vm = new ViewModel();
  QUnit.equal(vm.attr('message'), 'This is the resource-list-item component');
});
