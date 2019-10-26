import Vue from "vue";
import Vuex from "vuex";
import jwtDecode from "jwt-decode";
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

const SET_USER = "SET_USER";
const STORE_ACCESS_TOKEN = "STORE_ACCESS_TOKEN";
const STORE_REFRESH_TOKEN = "STORE_REFRESH_TOKEN";

const user = {
  namespaced: true,
  state: {
    user: null,
    accessToken: null,
    refreshToken: null
  },
  mutations: {
    SET_USER(state, data) {
      state.user = data;
    },
    STORE_ACCESS_TOKEN(state, accessToken) {
      state.accessToken = accessToken
      localStorage.setItem("accessToken", accessToken);
    },
    STORE_REFRESH_TOKEN(state, refreshToken) {
      state.refreshToken = refreshToken
      localStorage.setItem("refreshToken", refreshToken);
    },
  },
  getters: {

  },
  actions: {
    async setUserAndTokens({ commit }, data) {
      try {
        let decoded = jwtDecode(data.accessToken);
        commit(SET_USER, decoded.data);
        commit(STORE_ACCESS_TOKEN, data.accessToken);
        commit(STORE_REFRESH_TOKEN, data.refreshToken);
      } catch (error) {
        return Promise.reject(error.response ? error.response : error);
      }
    },
    async userLogin({ dispatch}, credentials) {
      try {
        const response = await axios.post("/login", {
          email: credentials.email,
          password: credentials.password
        });
        return await dispatch("setUserAndTokens", {
          accessToken: response.data.accessToken,
          refreshToken: response.data.refreshToken
        });
      } catch (error) {
        return Promise.reject(error.response ? error.response : error);
      }
    },
  }
}

export default user;
