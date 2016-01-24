import Component from 'can/component/';
import Map from 'can/map/';
import 'can/map/define/';
import './reviews-list-item.less!';
import template from './reviews-list-item.stache!';

export const ViewModel = Map.extend({
  define: {
    message: {
      value: 'This is the reviews-list-item component'
    }
  }
});

export default Component.extend({
  tag: 'reviews-list-item',
  viewModel: ViewModel,
  template
});