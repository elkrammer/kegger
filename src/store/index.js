import Vue from "vue";
import Vuex from "vuex";

Vue.use(Vuex);

import user from "@/store/user";
import users from "@/store/users";
import party from "@/store/party";

const store = new Vuex.Store({
  modules: {
    user,
    users,
    party,
  }
});

export default store;
