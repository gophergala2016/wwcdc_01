import Component from 'can/component/';
import Map from 'can/map/';
import 'can/map/define/';
import './create-review.less!';
import template from './create-review.stache!';

export const ViewModel = Map.extend({
  define: {
    message: {
      value: 'This is the create-review component'
    }
  }
});

export default Component.extend({
  tag: 'create-review',
  viewModel: ViewModel,
  template
});