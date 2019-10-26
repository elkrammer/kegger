import Vue from "vue";
import Vuex from "vuex";

Vue.use(Vuex);

const user = {
  state: {
    user: null,
    accessToken: null,
    refreshToken: null
  },
}

export default user;
