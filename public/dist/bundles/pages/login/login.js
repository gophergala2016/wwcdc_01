/*gophergala@0.0.0#components/login/login.stache!can@2.3.10#view/stache/system*/
define("gophergala@0.0.0#components/login/login.stache!can@2.3.10#view/stache/system",["module","can/view/stache/stache","can/view/stache/mustache_core"],function(t,e,a){var r=e([{tokenType:"start",args:["form",!1]},{tokenType:"attrStart",args:["($submit)"]},{tokenType:"attrValue",args:["send(%event)"]},{tokenType:"attrEnd",args:["($submit)"]},{tokenType:"end",args:["form",!1]},{tokenType:"chars",args:["\n  "]},{tokenType:"start",args:["div",!1]},{tokenType:"attrStart",args:["class"]},{tokenType:"attrValue",args:["form-group"]},{tokenType:"attrEnd",args:["class"]},{tokenType:"end",args:["div",!1]},{tokenType:"chars",args:["\n    "]},{tokenType:"start",args:["label",!1]},{tokenType:"attrStart",args:["for"]},{tokenType:"attrValue",args:["email"]},{tokenType:"attrEnd",args:["for"]},{tokenType:"end",args:["label",!1]},{tokenType:"chars",args:["E-mail"]},{tokenType:"close",args:["label"]},{tokenType:"chars",args:["\n    "]},{tokenType:"start",args:["input",!0]},{tokenType:"attrStart",args:["id"]},{tokenType:"attrValue",args:["email"]},{tokenType:"attrEnd",args:["id"]},{tokenType:"attrStart",args:["type"]},{tokenType:"attrValue",args:["text"]},{tokenType:"attrEnd",args:["type"]},{tokenType:"attrStart",args:["class"]},{tokenType:"attrValue",args:["form-control"]},{tokenType:"attrEnd",args:["class"]},{tokenType:"attrStart",args:["placeholder"]},{tokenType:"attrValue",args:["E-mail"]},{tokenType:"attrEnd",args:["placeholder"]},{tokenType:"attrStart",args:["{($value)}"]},{tokenType:"attrValue",args:["login.email"]},{tokenType:"attrEnd",args:["{($value)}"]},{tokenType:"end",args:["input",!0]},{tokenType:"chars",args:["\n  "]},{tokenType:"close",args:["div"]},{tokenType:"chars",args:["\n  "]},{tokenType:"start",args:["div",!1]},{tokenType:"attrStart",args:["class"]},{tokenType:"attrValue",args:["form-group"]},{tokenType:"attrEnd",args:["class"]},{tokenType:"end",args:["div",!1]},{tokenType:"chars",args:["\n    "]},{tokenType:"start",args:["label",!1]},{tokenType:"attrStart",args:["for"]},{tokenType:"attrValue",args:["password"]},{tokenType:"attrEnd",args:["for"]},{tokenType:"end",args:["label",!1]},{tokenType:"chars",args:["Password"]},{tokenType:"close",args:["label"]},{tokenType:"chars",args:["\n    "]},{tokenType:"start",args:["input",!0]},{tokenType:"attrStart",args:["id"]},{tokenType:"attrValue",args:["password"]},{tokenType:"attrEnd",args:["id"]},{tokenType:"attrStart",args:["type"]},{tokenType:"attrValue",args:["text"]},{tokenType:"attrEnd",args:["type"]},{tokenType:"attrStart",args:["class"]},{tokenType:"attrValue",args:["form-control"]},{tokenType:"attrEnd",args:["class"]},{tokenType:"attrStart",args:["placeholder"]},{tokenType:"attrValue",args:["Password"]},{tokenType:"attrEnd",args:["placeholder"]},{tokenType:"attrStart",args:["{($value)}"]},{tokenType:"attrValue",args:["login.password"]},{tokenType:"attrEnd",args:["{($value)}"]},{tokenType:"end",args:["input",!0]},{tokenType:"chars",args:["\n  "]},{tokenType:"close",args:["div"]},{tokenType:"chars",args:["\n  "]},{tokenType:"start",args:["div",!1]},{tokenType:"attrStart",args:["class"]},{tokenType:"attrValue",args:["form-group"]},{tokenType:"attrEnd",args:["class"]},{tokenType:"end",args:["div",!1]},{tokenType:"chars",args:["\n    "]},{tokenType:"start",args:["input",!0]},{tokenType:"attrStart",args:["type"]},{tokenType:"attrValue",args:["submit"]},{tokenType:"attrEnd",args:["type"]},{tokenType:"attrStart",args:["class"]},{tokenType:"attrValue",args:["btn btn-primary btn-block"]},{tokenType:"attrEnd",args:["class"]},{tokenType:"attrStart",args:["value"]},{tokenType:"attrValue",args:["Register"]},{tokenType:"attrEnd",args:["value"]},{tokenType:"end",args:["input",!0]},{tokenType:"chars",args:["\n  "]},{tokenType:"close",args:["div"]},{tokenType:"chars",args:["\n  "]},{tokenType:"start",args:["div",!1]},{tokenType:"attrStart",args:["class"]},{tokenType:"attrValue",args:["form-group"]},{tokenType:"attrEnd",args:["class"]},{tokenType:"end",args:["div",!1]},{tokenType:"chars",args:["\n    "]},{tokenType:"start",args:["input",!0]},{tokenType:"attrStart",args:["type"]},{tokenType:"attrValue",args:["submit"]},{tokenType:"attrEnd",args:["type"]},{tokenType:"attrStart",args:["class"]},{tokenType:"attrValue",args:["btn btn-primary btn-block"]},{tokenType:"attrEnd",args:["class"]},{tokenType:"attrStart",args:["value"]},{tokenType:"attrValue",args:["Login"]},{tokenType:"attrEnd",args:["value"]},{tokenType:"end",args:["input",!0]},{tokenType:"chars",args:["\n  "]},{tokenType:"close",args:["div"]},{tokenType:"chars",args:["\n"]},{tokenType:"close",args:["form"]},{tokenType:"chars",args:["\n"]},{tokenType:"done",args:[]}]);return function(e,s,n){var o={module:t};return s instanceof a.Options||(s=new a.Options(s||{})),r(e,s.add(o),n)}});
/*gophergala@0.0.0#models/login*/
define("gophergala@0.0.0#models/login",["exports","can","can-connect/can/super-map/","can-connect/can/tag/","can/map/define/define"],function(e,n,t,a,i){"use strict";function d(e){return e&&e.__esModule?e:{"default":e}}Object.defineProperty(e,"__esModule",{value:!0});var o=d(n),l=d(t),u=d(a),r=o["default"].Map.extend({define:{email:{type:"string"},password:{type:"string"}}});e.userLogin=r,r.List=o["default"].List.extend({Map:r},{});var s=l["default"]({url:"/api/user-auth",idProp:"id",Map:r,List:r.List,name:"login"});e.loginConnection=s,u["default"]("login-model",s),e["default"]=r});
/*gophergala@0.0.0#components/login/login*/
define("gophergala@0.0.0#components/login/login",["exports","can/component/","can/map/","can/map/define/","./login.less!","./login.stache!","gophergala/models/login"],function(e,n,t,o,l,a,i){"use strict";function s(e){return e&&e.__esModule?e:{"default":e}}Object.defineProperty(e,"__esModule",{value:!0});var u=s(n),d=s(t),r=s(a),g=s(i),f=d["default"].extend({define:{message:{value:"This is the login component",Type:g["default"]}},send:function(e){var n=this;e.preventDefault();var t=this.attr("userLogin").attr();console.log(t),new g["default"](t).save().then(function(){console.log("login hellooo"),n.resetValues()})},resetValues:function(){this.attr("userLogin",{})}});e.ViewModel=f,e["default"]=u["default"].extend({tag:"login",viewModel:f,template:r["default"]})});
/*gophergala@0.0.0#pages/login/login.stache!can@2.3.10#view/stache/system*/
define("gophergala@0.0.0#pages/login/login.stache!can@2.3.10#view/stache/system",["module","can/view/stache/stache","can/view/stache/mustache_core","gophergala/components/login/"],function(e,n,a){var t=n([{tokenType:"start",args:["can-import",!0]},{tokenType:"attrStart",args:["from"]},{tokenType:"attrValue",args:["gophergala/components/login/"]},{tokenType:"attrEnd",args:["from"]},{tokenType:"end",args:["can-import",!0]},{tokenType:"chars",args:["\n"]},{tokenType:"start",args:["h2",!1]},{tokenType:"end",args:["h2",!1]},{tokenType:"chars",args:["Login or Register!"]},{tokenType:"close",args:["h2"]},{tokenType:"chars",args:["\n"]},{tokenType:"start",args:["login",!0]},{tokenType:"end",args:["login",!0]},{tokenType:"chars",args:["\n"]},{tokenType:"done",args:[]}]);return function(n,o,r){var s={module:e};return o instanceof a.Options||(o=new a.Options(o||{})),t(n,o.add(s),r)}});
/*gophergala@0.0.0#pages/login/login*/
define("gophergala@0.0.0#pages/login/login",["exports","can/component/","can/map/","can/map/define/","./login.less!","./login.stache!"],function(e,n,a,t,l,o){"use strict";function i(e){return e&&e.__esModule?e:{"default":e}}Object.defineProperty(e,"__esModule",{value:!0});var d=i(n),s=i(a),g=i(o),u=s["default"].extend({define:{message:{value:"This is the pages-login component"}}});e.ViewModel=u,e["default"]=d["default"].extend({tag:"pages-login",viewModel:u,template:g["default"]})});