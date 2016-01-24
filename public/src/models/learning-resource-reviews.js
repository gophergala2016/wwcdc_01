import can from 'can';
import superMap from 'can-connect/can/super-map/';
import tag from 'can-connect/can/tag/';
import 'can/map/define/define';

export const LearningResourceReviews = can.Map.extend({
  define: {}
});

LearningResourceReviews.List = can.List.extend({
  Map: LearningResourceReviews
}, {});

export const learningResourceReviewsConnection = superMap({
  url: '/api/learning-resources/{id}/reviews',
  idProp: 'id',
  Map: LearningResourceReviews,
  List: LearningResourceReviews.List,
  name: 'learningResourceReviews'
});

tag('learningResourceReviews-model', learningResourceReviewsConnection);

export default LearningResourceReviews;
