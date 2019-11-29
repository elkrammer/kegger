import Vue from "vue";
import Vuex from "vuex";

Vue.use(Vuex);

import user from "@/store/user";
import users from "@/store/users";
import party from "@/store/party";
import guest from "@/store/guest";
import settings from "@/store/settings";

const store = new Vuex.Store({
  modules: {
    user,
    users,
    party,
    guest,
    settings,
  }
});

export default store;
