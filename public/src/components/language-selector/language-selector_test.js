import QUnit from 'steal-qunit';
import { ViewModel } from './language-selector';

// ViewModel unit tests
QUnit.module('gophergala/components/language-selector');

QUnit.test('Has message', function(){
  var vm = new ViewModel();
  QUnit.equal(vm.attr('message'), 'This is the language-selector component');
});
