import Component from 'can/component/';
import Map from 'can/map/';
import 'can/map/define/';
import './resources-list.less!';
import template from './resources-list.stache!';
import LearningResources from 'gophergala/models/learning-resources';
import route from "can/route/";

export const ViewModel = Map.extend({
  define: {
  },
  changeRoute (resource, a, event) {
    event.preventDefault();
    route.attr({page: 'resource', slug: resource.attr('id')})
  }
});

export default Component.extend({
  tag: 'resources-list',
  viewModel: ViewModel,
  template
});