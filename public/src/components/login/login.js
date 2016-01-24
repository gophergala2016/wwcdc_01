import Component from 'can/component/';
import Map from 'can/map/';
import 'can/map/define/';
import './login.less!';
import template from './login.stache!';
//import LearningResources from 'gophergala/models/learning-resources';

export const ViewModel = Map.extend({
  define:{
    userLogin: {
      value: {},
      Type: LearningResources
    }
  },
/*  send(event) {
    event.preventDefault();
    let resource = this.attr('learningResource').attr();
    console.log(resource)
    new LearningResources(resource).save().then(() => {
      console.log('hellooo')
      this.resetValues()
    });
  }, */
  resetValues() {
    this.attr('userLogin', {});
  }
});

export default Component.extend({
  tag: 'login',
  viewModel: ViewModel,
  template
});
