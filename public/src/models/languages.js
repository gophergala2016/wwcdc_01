import can from 'can';
import superMap from 'can-connect/can/super-map/';
import tag from 'can-connect/can/tag/';
import 'can/map/define/define';

export const Languages = can.Map.extend({
  define: {}
});

Languages.List = can.List.extend({
  Map: Languages
}, {});

export const languagesConnection = superMap({
  url: 'api/languages',
  idProp: 'id',
  Map: Languages,
  List: Languages.List,
  name: 'languages'
});

tag('languages-model', languagesConnection);

export default Languages;
