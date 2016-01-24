import Component from 'can/component/';
import Map from 'can/map/';
import 'can/map/define/';
import './resource-list-item.less!';
import template from './resource-list-item.stache!';
import LearningResources from 'gophergala/models/learning-resources';

export const ViewModel = Map.extend({
  define: {
    resource: {
      value: {},
      Type: LearningResources
    },
    languages: {
      get(){
        return this.resource.attr('languages').join(', ') || '';
      }
    }
  }
});

export default Component.extend({
  tag: 'resource-list-item',
  viewModel: ViewModel,
  template
});