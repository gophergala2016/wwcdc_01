import QUnit from 'steal-qunit';
import { ViewModel } from './login';

// ViewModel unit tests
QUnit.module('gophergala/pages/login');

QUnit.test('Has message', function(){
  var vm = new ViewModel();
  QUnit.equal(vm.attr('message'), 'This is the pages-login component');
});
