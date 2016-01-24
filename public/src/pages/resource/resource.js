import Component from 'can/component/';
import Map from 'can/map/';
import 'can/map/define/';
import './resource.less!';
import template from './resource.stache!';

export const ViewModel = Map.extend({
  define: {
    message: {
      value: 'This is the pages-resource component'
    }
  }
});

export default Component.extend({
  tag: 'pages-resource',
  viewModel: ViewModel,
  template
});