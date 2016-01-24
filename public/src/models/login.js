import can from 'can';
import superMap from 'can-connect/can/super-map/';
import tag from 'can-connect/can/tag/';
import 'can/map/define/define';

export const userLogin = can.Map.extend({
  define: {
    email: {
      type: 'string'
    },
    password: {
      type: 'string'
    }
  }
});

userLogin.List = can.List.extend({
  Map: userLogin
}, {});

export const loginConnection = superMap({
  url: '/api/user-auth',
  idProp: 'id',
  Map: userLogin,
  List: userLogin.List,
  name: 'login'
});

tag('login-model', loginConnection);

export default userLogin;
