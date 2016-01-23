import Component from 'can/component/';
import Map from 'can/map/';
import 'can/map/define/';
import './resources-list.less!';
import template from './resources-list.stache!';

export const ViewModel = Map.extend({
  define: {
    message: {
      value: 'This is the resources-list component'
    }
  }
});

export default Component.extend({
  tag: 'resources-list',
  viewModel: ViewModel,
  template
});