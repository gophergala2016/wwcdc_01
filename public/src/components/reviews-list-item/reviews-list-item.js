import Component from 'can/component/';
import Map from 'can/map/';
import 'can/map/define/';
import './reviews-list-item.less!';
import template from './reviews-list-item.stache!';

export const ViewModel = Map.extend({
  define: {
    cost: {
      type: 'number',
      get(setVal) {
        if (setVal) {
          return '$'+ setVal.toFixed(2);
        }
        return 'no cost';
      }
    }
  }
});

export default Component.extend({
  tag: 'reviews-list-item',
  viewModel: ViewModel,
  template
});