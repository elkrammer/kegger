import Vue from "vue";
import Vuex from "vuex";
import axios from "@/common/axios";
import { setAuthorizationHeader } from "@/common/auth";

Vue.use(Vuex);

const GET_GUEST = "GET_GUEST";
const GET_GUESTS = "GET_GUESTS";
const EDIT_GUESTS = "EDIT_GUESTS";
const DELETE_GUEST = "DELETE_GUEST";

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
    EDIT_GUESTS(state, data) {
      Vue.set(state.guests, data);
    },
    DELETE_GUEST(state, index) {
      const itemId = state.guests.find(guest => guest.ID === index);
      state.guests.splice(state.guests.indexOf(itemId), 1);
    },
  },
  getters: {
    guests(state) {
      return state.guests;
    }
  },
  actions: {
    async getGuest({ commit, rootGetters }, id) {
      try {
        setAuthorizationHeader(rootGetters["user/accessToken"]);
        const response = await axios.get("/api/guest/" + id);
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
  },
};

export default guest;
