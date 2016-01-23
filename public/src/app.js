import AppMap from "can-ssr/app-map";
import route from "can/route/";
import 'can/map/define/';
import 'can/route/pushstate/';

//!steal-remove-start
import 'gophergala/models/fixtures/';
//!steal-remove-end

const AppViewModel = AppMap.extend({
  define: {
    message: {
      value: 'Hello World!',
      serialize: false
    },
    title: {
      value: 'gophergala',
      serialize: false
    }
  }
});

route('/:page', { page: 'home' })

export default AppViewModel;
