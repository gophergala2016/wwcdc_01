import Component from 'can/component/';
import Map from 'can/map/';
import 'can/map/define/';
import './resource-list-item.less!';
import template from './resource-list-item.stache!';

export const ViewModel = Map.extend({
  define: {
    message: {
      value: 'This is the resource-list-item component'
    }
  }
});

export default Component.extend({
  tag: 'resource-list-item',
  viewModel: ViewModel,
  template
});