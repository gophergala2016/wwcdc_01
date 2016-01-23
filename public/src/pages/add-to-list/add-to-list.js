import Component from 'can/component/';
import Map from 'can/map/';
import 'can/map/define/';
import './add-to-list.less!';
import template from './add-to-list.stache!';

export const ViewModel = Map.extend({
  define: {
    message: {
      value: 'This is the pages-add-to-list component'
    }
  }
});

export default Component.extend({
  tag: 'pages-add-to-list',
  viewModel: ViewModel,
  template
});