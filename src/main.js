import Vue from 'vue'
import Vuelidate from 'vuelidate';
import Buefy from 'buefy'
import VueFilterDateFormat from 'vue-filter-date-format';
import VueFilterDateParse from 'vue-filter-date-parse';

import App from './App.vue'
import router from './router';
import store from './store';

Vue.use(Buefy, { defaultIconPack: 'fas' })
Vue.use(Vuelidate);
Vue.use(VueFilterDateParse);
Vue.use(VueFilterDateFormat);

Vue.config.productionTip = false
Vue.prototype.$log = console.log;

new Vue({
  store,
  router,
  render: h => h(App),
}).$mount('#app')
