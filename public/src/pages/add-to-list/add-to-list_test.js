import QUnit from 'steal-qunit';
import { ViewModel } from './add-to-list';

// ViewModel unit tests
QUnit.module('gophergala/pages/add-to-list');

QUnit.test('Has message', function(){
  var vm = new ViewModel();
  QUnit.equal(vm.attr('message'), 'This is the pages-add-to-list component');
});
