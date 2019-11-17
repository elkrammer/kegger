import Vue from "vue";
import Vuex from "vuex";
import axios from "@/common/axios";
import { setAuthorizationHeader } from "@/common/utilities";

Vue.use(Vuex);

const GET_GUEST = "GET_GUEST";
const GET_GUESTS = "GET_GUESTS";
const ADD_GUEST = "ADD_GUEST";
const EDIT_GUESTS = "EDIT_GUESTS";
const DELETE_GUEST = "DELETE_GUEST";
const ADD_EMPTY_GUEST = "ADD_EMPTY_GUEST"

const guest = {
  namespaced: true,
  state: {
    guests: [],
  },
  mutations: {
    GET_GUEST(state, data) {
      state.guest = data;
    },
    GET_GUESTS(state, data) {
      state.guests = data;
    },
    ADD_GUEST(state, data) {
      state.guests.push(data);
    },
    EDIT_GUESTS(state, data) {
      Vue.set(state.guests, data);
    },
    DELETE_GUEST(state, index) {
      const itemId = state.guests.find(guest => guest.ID === index);
      state.guests.splice(state.guests.indexOf(itemId), 1);
    },
    ADD_EMPTY_GUEST(state, party_refer) {
      const guest = {
        first_name:  '',
        last_name:    '',
        email:  '',
        party_refer: party_refer,
      };
      state.guests.unshift(guest);
    }
  },
  getters: {
    guests(state) {
      return state.guests;
    }
  },
  actions: {
    async getGuest({ commit, rootGetters }, index) {
      try {
        setAuthorizationHeader(rootGetters["user/accessToken"]);
        const response = await axios.get("/api/guest/" + index);
        commit(GET_GUEST, response.data);
        return response.data;
      } catch (error) {
        return Promise.reject(error);
      }
    },
    async getGuests({ commit, rootGetters }, party_id) {
      try {
        setAuthorizationHeader(rootGetters["user/accessToken"]);
        const response = await axios.get("/api/guests/" + party_id);
        commit(GET_GUESTS, response.data);
        return response.data;
      } catch (error) {
        return Promise.reject(error);
      }
    },
    async createGuest({ commit, rootGetters }, data) {
      try {
        setAuthorizationHeader(rootGetters["user/accessToken"]);
        const response = await axios.post("/api/guest", {
          name: data.name,
          host_id: data.host_id,
          comments: data.comments,
          Guests: data.guests,
        });
        commit(ADD_GUEST, response.data);
        return response.data;
      } catch (error) {
        return Promise.reject(error);
      }
    },
    async editGuests({ commit, rootGetters }, data) {
      try{
        setAuthorizationHeader(rootGetters["user/accessToken"]);
        const response = await axios.put("/api/guest/" + data[0].party_refer, data);
        commit(EDIT_GUESTS, response.data);
        return response.data;
      } catch(error) {
        return Promise.reject(error);
      }
    },
    async deleteGuest({ commit, rootGetters }, index) {
      try {
        setAuthorizationHeader(rootGetters["user/accessToken"]);
        const response = await axios.delete("/api/guest/" + index);
        commit(DELETE_GUEST, index);
        return response.data;
      } catch(error) {
        return Promise.reject(error);
      }
    },
    addEmptyGuest({ commit }, party_id) {
      commit(ADD_EMPTY_GUEST, party_id);
    },
  },
};

export default guest;
