import Vue from "vue";
import Vuex from "vuex";

Vue.use(Vuex);
import axios from "@/common/axios";
import { setAuthorizationHeader } from "@/common/utilities";

const GET_PARTIES = "GET_PARTIES";
const ADD_PARTY_TO_STACK = "ADD_PARTY_TO_STACK";

const party = {
  namespaced: true,
  state: {
    parties: []
  },
  mutations: {
    GET_PARTIES(state, data) {
      state.parties = data;
    },
    ADD_PARTY_TO_STACK(state, party) {
      state.parties.unshift(party);
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
    async createParty({ rootGetters }, data) {
      try {
        setAuthorizationHeader(rootGetters["user/accessToken"]);
        const response = await axios.post("/api/party", {
          name: data.name,
          HostId: data.host_id,
          Comments: data.comments,
          Guests: data.guests,
        });
        return response.data;
      } catch (error) {
        return Promise.reject(error);
      }
    },

    // only mutations
    async addPartyToStack({ commit }, party) {
      try {
        commit(ADD_PARTY_TO_STACK, party);
      } catch (error) {
        return Promise.reject(error);
      }
    }
  },
}

export default party;
