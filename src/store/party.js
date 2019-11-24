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

      // for every party get a status of whether Guests are attending or not...
      let partySize = state.parties.length;
      let isPartyAttending = ''; // possible values: "true", "false", "partial"

      for (var i = 0; i < partySize; i++) {
        let attending = 0; // total guests attending
        let guestSize = state.parties[i].Guests.length;
        for (var j = 0; j < guestSize; j++) {
          if (state.parties[i].Guests[j].is_attending) {
            attending += 1
          }
        }

        if (attending === guestSize) {
          isPartyAttending = "true"
        }

        else if (attending > 0 && attending < guestSize) {
          isPartyAttending = "partial"
        }

        else {
          isPartyAttending = "false"
        }

        // add isPartyAttending property to every party
        Vue.set(state.parties[i], 'isPartyAttending', isPartyAttending)
      }
    },
    ADD_PARTY(state, data) {
      Vue.set(data, 'isPartyAttending', "false")
      state.parties.push(data);
    },
    EDIT_PARTY(state, data) {
      const itemId = state.parties.findIndex(party => party.id === data.id);
      Vue.set(state.parties, itemId, data);
    },
    DELETE_PARTY(state, index) {
      const itemId = state.parties.find(party => party.id === index);
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
          host_id: data.host_id,
          comments: data.comments,
          Guests: data.guests,
        });
        commit(ADD_PARTY, response.data);
        return response.data;
      } catch (error) {
        return Promise.reject(error);
      }
    },
    async editParty({ commit, rootGetters }, data) {
      try {
        setAuthorizationHeader(rootGetters["user/accessToken"]);
        const response = await axios.put("/api/party/" + data.id, data);
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
