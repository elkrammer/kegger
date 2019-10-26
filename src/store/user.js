import Vue from "vue";
import Vuex from "vuex";

import axios from "@/common/axios";

Vue.use(Vuex);

// axios interceptor that will handle expired tokens
axios.interceptors.response.use(undefined, async error => {
  if (error.response.status === 401 && error.response.data.message === "TOKEN_EXPIRED" && !error.config.__isRetryRequest) {
    try {
      console.log("token expired")
    } catch (error) {
      return Promise.reject(error);
    }
  }
});

const user = {
  namespaced: true,
  state: {
    user: null,
    accessToken: null,
    refreshToken: null
  },
  mutations: {

  },
  getters: {

  },
  actions: {

  },
}

export default user;
