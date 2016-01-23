import QUnit from 'steal-qunit';
import { ViewModel } from './resources-list';

// ViewModel unit tests
QUnit.module('gophergala/components/resources-list');

QUnit.test('Has message', function(){
  var vm = new ViewModel();
  QUnit.equal(vm.attr('message'), 'This is the resources-list component');
});
