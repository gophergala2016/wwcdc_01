import can from 'can';
import 'can/map/define/define';

import Model from "can-connect/can/model/";

let Languages = Model.extend({
  findOne: "/api/languages"
},{});

export default Languages;
