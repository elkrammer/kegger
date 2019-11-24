// this store is used to get the admin user list, create new admins, etc

import Vue from "vue";
import Vuex from "vuex";
import axios from "@/common/axios";
import { setAuthorizationHeader } from "@/common/utilities";

Vue.use(Vuex);

const GET_USERS = "GET_USERS";
const GET_USER = "GET_USER";
const ADD_USER = "ADD_USER";
const DELETE_USER = "DELETE_USER";

const users = {
  namespaced: true,
  state: {
    users: [],
  },
  mutations: {
    GET_USERS(state, data) {
      state.users = data;
    },
    GET_USER(state, data) {
      state.user = data;
    },
    ADD_USER(state, data) {
      state.users.push(data);
    },
    DELETE_USER(state, index) {
      const itemId = state.users.find(user => user.id === index);
      state.users.splice(state.users.indexOf(itemId), 1);
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
    async getUser({ commit, rootGetters }, id) {
      try {
        setAuthorizationHeader(rootGetters["user/accessToken"]);
        const response = await axios.get("/api/user/" + id);
        commit(GET_USER, response.data);
        return response.data;
      } catch (error) {
        return Promise.reject(error);
      }
    },
    async createUser({ commit, rootGetters }, data) {
      try {
        setAuthorizationHeader(rootGetters["user/accessToken"]);
        const response = await axios.post("/api/user", {
          name: data.name,
          email: data.email,
          password: data.password,
        });
        commit(ADD_USER, response.data);
        return response.data;
      } catch (error) {
        return Promise.reject(error);
      }
    },
    async deleteUser({ commit, rootGetters }, id) {
      try {
        setAuthorizationHeader(rootGetters["user/accessToken"]);
        const response = await axios.delete("/api/user/" + id);
        commit(DELETE_USER, id);
        return response.data;
      } catch (error) {
        return Promise.reject(error);
      }
    },
  }
};

export default users;
