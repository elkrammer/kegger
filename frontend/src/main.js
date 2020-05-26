import Vue from 'vue'
import Buefy from 'buefy'
import VueMoment from 'vue-moment'
import moment from 'moment-timezone'

import App from './App.vue'
import router from './router';
import store from './store';

Vue.use(Buefy, { defaultIconPack: 'fas' })
Vue.use(VueMoment, {
    moment,
})

Vue.config.productionTip = false
Vue.prototype.$log = console.log;

new Vue({
  store,
  router,
  render: h => h(App),
}).$mount('#app')
