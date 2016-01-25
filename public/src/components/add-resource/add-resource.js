import Component from 'can/component/';
import Map from 'can/map/';
import 'can/map/define/';
import './add-resource.less!';
import template from './add-resource.stache!';
import LearningResources from 'gophergala/models/learning-resources';

export const ViewModel = Map.extend({
  define:{
    learningResource: {
      value: {},
      Type: LearningResources
    },
    saving: {
      value: false,
      type:'boolean'
    }
  },
  send(event) {
    event.preventDefault();
    let resource = this.attr('learningResource').attr();
    this.attr('saving', true);
    new LearningResources(resource).save().then(() => {
      this.resetValues();
      this.attr('saving', false);
    });
  },
  resetValues() {
    this.attr('learningResource', {});
  }
});


export default Component.extend({
  tag: 'add-resource',
  viewModel: ViewModel,
  template
});