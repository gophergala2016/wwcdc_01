import Component from 'can/component/';
import Map from 'can/map/';
import 'can/map/define/';
import './reviews-list.less!';
import template from './reviews-list.stache!';

export const ViewModel = Map.extend({
  define: {
    message: {
      value: 'This is the reviews-list component'
    },
    created: {
      value: 0,
    }
  }
});

export default Component.extend({
  tag: 'reviews-list',
  viewModel: ViewModel,
  template
});