import Component from 'can/component/';
import Map from 'can/map/';
import 'can/map/define/';
import './create-review.less!';
import template from './create-review.stache!';
import Reviews from 'gophergala/models/reviews';

export const ViewModel = Map.extend({
  define:{
    review: {
      value: {},
      Type: Reviews
    },
    saving: {
      value: false,
      type:'boolean'
    }
  },
  send(event) {
    event.preventDefault();
    let resource = this.attr('review').attr();
    resource.learing_resource_id = this.attr('resouceId');
    this.attr('saving', true);
    new Reviews(resource).save().then(() => {
      this.resetValues();
      this.attr('created', Math.random()*100)
      this.attr('saving', false);
    });
  },
  resetValues() {
    this.attr('review', {});
  }
});

export default Component.extend({
  tag: 'create-review',
  viewModel: ViewModel,
  template
});