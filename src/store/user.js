import Vue from "vue";
import Vuex from "vuex";
import jwtDecode from "jwt-decode";
import axios from "@/common/axios";
import store from "@/store/index";
import router from "@/router";
import { setAuthorizationHeader } from "@/common/utilities";

Vue.use(Vuex);

let authTokenRequest;
async function getAuthToken() {
  if (!authTokenRequest) {
    authTokenRequest = store.dispatch("user/refreshUserTokens");
    authTokenRequest.then(() => {
      authTokenRequest = null;
    }).catch(() => {
      authTokenRequest = null;
    });
  }
}

async function logoutOfApplication() {
  await store.dispatch("user/userLogout")
  router.replace({ name: "login" });
}

// axios interceptor that will handle expired tokens
axios.interceptors.response.use(undefined, async error => {
  if (error.response.status === 401 && error.response.data.message === "TOKEN_EXPIRED" && !error.config.__isRetryRequest) {
    try {
      let response = await getAuthToken();
      await store.dispatch("user/setUserAndTokens", {
        accessToken: response.data.access_token,
        refreshToken: response.data.refresh_token
      });
      error.config.headers["Authorization"] = "Bearer " + store.getters["user/accessToken"];
      error.config.__isRetryRequest = true;
      error.config.baseURL = ""; // baseURL needs to be zeroud out to prevent tripping over the baseURL set in main axios instance
      return axios(error.config);
    } catch (error) {
      logoutOfApplication();
      return Promise.reject(error);
    }
  }

  // handle user that isn't correctly logged in
  if (error.response.status === 401 && error.response.data.message === "AUTHENTICATION_ERROR") {
    logoutOfApplication();
    return Promise.reject(error);
  }

  // if someone gets here we don't want to log them out since it's more of a general error
  return Promise.reject(error);
});

const SET_USER = "SET_USER";
const STORE_ACCESS_TOKEN = "STORE_ACCESS_TOKEN";
const STORE_REFRESH_TOKEN = "STORE_REFRESH_TOKEN";
const LOGOUT_USER = "LOGOUT_USER";

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
    LOGOUT_USER(state) {
      state.user = null;
      state.accessToken = null;
      state.refreshToken = null;
      localStorage.removeItem("accessToken");
      localStorage.removeItem("refreshToken");
    }
  },
  getters: {
    user(state) {
      return state.user;
    },
    accessToken(state) {
      return state.accessToken;
    },
    refreshToken(state) {
      return state.refreshToken;
    }
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
          accessToken: response.data.access_token,
          refreshToken: response.data.refresh_token
        });
      } catch (error) {
        return Promise.reject(error.response ? error.response : error);
      }
    },
    async refreshUserTokens({ getters, rootGetters }) {
      try {
        setAuthorizationHeader(rootGetters["user/accessToken"]);
        return await axios.post("/token", {
          username: getters.user.username,
          refreshToken: getters.refreshToken
        });
      } catch (error) {
        return Promise.reject(error.response ? error.response : error);
      }
    },
    async userLogout({ commit }) {
      try {
        commit(LOGOUT_USER);
      } catch (error) {
        return Promise.reject(error.response ? error.response : error);
      }
    }
  }
}

export default user;
