import Vue from "vue";
import Vuex from "vuex";

Vue.use(Vuex);
import axios from "@/common/axios";
import { setAuthorizationHeader } from "@/common/utilities";

const GET_PARTIES = "GET_PARTIES";
const ADD_PARTY = "ADD_PARTY";

const party = {
  namespaced: true,
  state: {
    parties: []
  },
  mutations: {
    GET_PARTIES(state, data) {
      state.parties = data;
    },
    ADD_PARTY(state, data) {
      state.parties.push(data);
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
    async createParty({ commit, rootGetters }, data) {
      try {
        setAuthorizationHeader(rootGetters["user/accessToken"]);
        const response = await axios.post("/api/party", {
          name: data.name,
          HostId: data.host_id,
          Comments: data.comments,
          Guests: data.guests,
        });
        commit(ADD_PARTY, response.data);
        return response.data;
      } catch (error) {
        return Promise.reject(error);
      }
    },
  },
}

export default party;
