import can from 'can';
import superMap from 'can-connect/can/super-map/';
import tag from 'can-connect/can/tag/';
import 'can/map/define/define';

export const Reviews = can.Map.extend({
  define: {
    "learning_resource_id": {
      value:0
    },
    "usefulness": {
      value:0
    },
    "finished":{
      value: false
    },
    "cost": {
      value:0
    },
    "course_size": {
      value:0
    },
    "course_length": {
      value:0
    },
    "course_interaction": {
      value:0
    },
    "description": {
      value:''
    },
  }
});

Reviews.List = can.List.extend({
  Map: Reviews
}, {});

export const reviewsConnection = superMap({
  url: '/api/reviews',
  idProp: 'id',
  Map: Reviews,
  List: Reviews.List,
  name: 'reviews'
});

tag('reviews-model', reviewsConnection);

export default Reviews;
