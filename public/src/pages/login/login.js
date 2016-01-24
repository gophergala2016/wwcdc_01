import Component from 'can/component/';
import Map from 'can/map/';
import 'can/map/define/';
import './login.less!';
import template from './login.stache!';

export const ViewModel = Map.extend({
  define: {
    message: {
      value: 'This is the pages-login component'
    }
  }
});

export default Component.extend({
  tag: 'pages-login',
  viewModel: ViewModel,
  template
});
