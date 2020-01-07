// this store is used to manage application settings
import Vue from "vue";
import Vuex from "vuex";
import axios from "@/common/axios";
import { setAuthorizationHeader } from "@/common/utilities";

Vue.use(Vuex);

const GET_SETTINGS = "GET_SETTINGS";
const UPDATE_SETTINGS = "UPDATE_SETTINGS";

const settings = {
  namespaced: true,
  state: {
    settings: [],
  },
  mutations: {
    GET_SETTINGS(state, data) {
      state.settings = data;
    },
    UPDATE_SETTINGS(state, data) {
      Vue.set(state.settings, data);
    },
  },
  getters: {
    settings(state) {
      return state.settings;
    },
  },
  actions: {
    async getSettings({ commit, rootGetters }) {
      try {
        setAuthorizationHeader(rootGetters["user/accessToken"]);
        const response = await axios.get("/api/settings");
        commit(GET_SETTINGS, response.data);
        return response.data;
      } catch (error) {
        return Promise.reject(error);
      }
    },
    async updateSettings({ commit, rootGetters }, data) {
      try {
        setAuthorizationHeader(rootGetters["user/accessToken"]);
        const response = await axios.put("/api/settings", data);
        commit(UPDATE_SETTINGS, response.data);
        return response.data;
      } catch (error) {
        return Promise.reject(error);
      }
    },
  }
};

export default settings;
