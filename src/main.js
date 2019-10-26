import Vue from 'vue'
import Vuelidate from "vuelidate";
import Buefy from 'buefy'
import App from './App.vue'
import router from "./router";
import store from "./store";
import 'buefy/dist/buefy.css'

Vue.use(Buefy)
Vue.use(Vuelidate);

Vue.config.productionTip = false

new Vue({
  store,
  router,
  render: h => h(App),
}).$mount('#app')
