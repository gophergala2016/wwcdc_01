import QUnit from 'steal-qunit';
import { ViewModel } from './404';

// ViewModel unit tests
QUnit.module('gophergala/pages/404');

QUnit.test('Has message', function(){
  var vm = new ViewModel();
  QUnit.equal(vm.attr('message'), 'This is the pages-404 component');
});
