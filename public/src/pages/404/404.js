import Component from 'can/component/';
import Map from 'can/map/';
import 'can/map/define/';
import './404.less!';
import template from './404.stache!';

export const ViewModel = Map.extend({
  define: {
    message: {
      value: 'This is the pages-404 component'
    }
  }
});

export default Component.extend({
  tag: 'pages-404',
  viewModel: ViewModel,
  template
});