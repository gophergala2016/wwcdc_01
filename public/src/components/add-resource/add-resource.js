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
    }
  },
  send(event) {
    event.preventDefault();
    let resource = this.attr('learningResource').attr();
    console.log(resource)
    new LearningResources(resource).save().then(() => {
      console.log('hellooo')
      this.resetValues()
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