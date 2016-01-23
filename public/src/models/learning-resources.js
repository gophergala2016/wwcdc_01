import can from 'can';
import superMap from 'can-connect/can/super-map/';
import tag from 'can-connect/can/tag/';
import 'can/map/define/define';

export const LearningResources = can.Map.extend({
  define: {
    name: {
      type: 'string'
    },
    url: {
      type: 'string'
    },
    type: {
      type: 'string'
    },
    languages: {
      value: []
    }
  }
});

LearningResources.List = can.List.extend({
  Map: LearningResources
}, {});

export const learningResourcesConnection = superMap({
  url: '/api/learning-resources',
  idProp: 'id',
  Map: LearningResources,
  List: LearningResources.List,
  name: 'learning-resources'
});

tag('learning-resources-model', learningResourcesConnection);

export default LearningResources;
