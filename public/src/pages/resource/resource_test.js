import QUnit from 'steal-qunit';
import { ViewModel } from './resource';

// ViewModel unit tests
QUnit.module('gophergala/pages/resource');

QUnit.test('Has message', function(){
  var vm = new ViewModel();
  QUnit.equal(vm.attr('message'), 'This is the pages-resource component');
});
