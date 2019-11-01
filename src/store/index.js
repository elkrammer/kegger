import Vue from "vue";
import Vuex from "vuex";

Vue.use(Vuex);

import user from "@/store/user";
import party from "@/store/party";

const store = new Vuex.Store({
  modules: {
    user,
    party,
  }
});

export default store;
