import Vue from "vue";
import Vuex from "vuex";

Vue.use(Vuex);
import axios from "@/common/axios";
import { setAuthorizationHeader } from "@/common/utilities";

const GET_PARTIES = "GET_PARTIES";

const party = {
  namespaced: true,
  state: {
    parties: []
  },
  mutations: {
    GET_PARTIES(state, data) {
      state.parties = data;
    },
  },
  getters: {
    parties(state) {
      return state.parties;
    }
  },
  actions: {
    async getParties({ commit, rootGetters }) {
      try {
        setAuthorizationHeader(rootGetters["user/accessToken"]);
        const response = await axios.get("/api/party");
        commit(GET_PARTIES, response.data);
        return response.data;
      } catch (error) {
        return Promise.reject(error);
      }
    },
  },
}

export default party;
