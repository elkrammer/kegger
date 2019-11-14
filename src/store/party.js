import Vue from "vue";
import Vuex from "vuex";

Vue.use(Vuex);
import axios from "@/common/axios";
import { setAuthorizationHeader } from "@/common/utilities";

const GET_PARTY = "GET_PARTY";
const GET_PARTIES = "GET_PARTIES";
const ADD_PARTY = "ADD_PARTY";
const EDIT_PARTY = "EDIT_PARTY";
const DELETE_PARTY = "DELETE_PARTY";

const party = {
  namespaced: true,
  state: {
    parties: [],
  },
  mutations: {
    GET_PARTY(state, data) {
      state.party = data;
    },
    GET_PARTIES(state, data) {
      state.parties = data;
    },
    ADD_PARTY(state, data) {
      state.parties.push(data);
    },
    EDIT_PARTY(state, data) {
      const itemId = state.parties.findIndex(party => party.ID === data.ID);
      Vue.set(state.parties, itemId, data);
    },
    DELETE_PARTY(state, index) {
      const itemId = state.parties.find(party => party.ID === index);
      state.parties.splice(state.parties.indexOf(itemId), 1);
    },
  },
  getters: {
    parties(state) {
      return state.parties;
    }
  },
  actions: {
    async getParty({ commit, rootGetters }, index) {
      try {
        setAuthorizationHeader(rootGetters["user/accessToken"]);
        const response = await axios.get("/api/party/" + index);
        commit(GET_PARTY, response.data);
        return response.data;
      } catch (error) {
        return Promise.reject(error);
      }
    },
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
    async editParty({ commit, rootGetters }, data) {
      try{
        setAuthorizationHeader(rootGetters["user/accessToken"]);
        const response = await axios.put("/api/party/" + data.ID, data);
        commit(EDIT_PARTY, response.data);
        return response.data;
      } catch(error) {
        return Promise.reject(error);
      }
    },
    async deleteParty({ commit, rootGetters }, index) {
      try {
        setAuthorizationHeader(rootGetters["user/accessToken"]);
        const response = await axios.delete("/api/party/" + index);
        commit(DELETE_PARTY, index);
        return response.data;
      } catch(error) {
        return Promise.reject(error);
      }
    },
  },
};

export default party;
