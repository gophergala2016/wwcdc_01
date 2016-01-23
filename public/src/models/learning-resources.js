import can from 'can';
import superMap from 'can-connect/can/super-map/';
import tag from 'can-connect/can/tag/';
import 'can/map/define/define';

export const Learning-resources = can.Map.extend({
  define: {}
});

Learning-resources.List = can.List.extend({
  Map: Learning-resources
}, {});

export const learning-resourcesConnection = superMap({
  url: '/api/learning-resources',
  idProp: 'id',
  Map: Learning-resources,
  List: Learning-resources.List,
  name: 'learning-resources'
});

tag('learning-resources-model', learning-resourcesConnection);

export default Learning-resources;
