// this store is used to get the admin user list, create new admins, etc

import Vue from "vue";
import Vuex from "vuex";
import axios from "@/common/axios";
import { setAuthorizationHeader } from "@/common/utilities";

Vue.use(Vuex);

const GET_USERS = "GET_USERS";

const users = {
  namespaced: true,
  state: {
    users: [],
  },
  mutations: {
    GET_USERS(state, data) {
      state.users= data;
    },
  },
  getters: {
    users(state) {
      return state.users;
    },
  },
  actions: {
    async getUsers({ commit, rootGetters }) {
      try {
        setAuthorizationHeader(rootGetters["user/accessToken"]);
        const response = await axios.get("/api/user");
        commit(GET_USERS, response.data);
        return response.data;
      } catch (error) {
        return Promise.reject(error);
      }
    },
  }
};

export default users;
