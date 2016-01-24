import Component from 'can/component/';
import Map from 'can/map/';
import 'can/map/define/';
import './login.less!';
import template from './login.stache!';
import userLogin from 'gophergala/models/login';

export const ViewModel = Map.extend({
  define:{
    message: {
  //    value: {},
      value: 'This is the login component',
      Type: userLogin
    }
  },
 /* send(event) {
    event.preventDefault();
    let user = this.attr('userLogin').attr();
    console.log(user)
    new userLogin(user).save().then(() => {
      console.log('login hellooo')
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
