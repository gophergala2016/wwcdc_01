import Component from 'can/component/';
import Map from 'can/map/';
import 'can/map/define/';
import './resource-detail.less!';
import template from './resource-detail.stache!';

export const ViewModel = Map.extend({
  define: {
    message: {
      value: 'This is the resource-detail component'
    }
  }
});

export default Component.extend({
  tag: 'resource-detail',
  viewModel: ViewModel,
  template
});