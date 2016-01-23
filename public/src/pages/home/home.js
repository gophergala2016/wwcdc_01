import Component from 'can/component/';
import Map from 'can/map/';
import 'can/map/define/';
import './home.less!';
import template from './home.stache!';

export const ViewModel = Map.extend({
  define: {
    message: {
      value: 'This is the pages-home component'
    }
  }
});

export default Component.extend({
  tag: 'pages-home',
  viewModel: ViewModel,
  template
});