// this store sets parameters for the current logged in user
// eg: 'user': {'user': "admin@admin.com", 'accessToken': "...", 'refreshToken': "..." }

import Vue from "vue";
import Vuex from "vuex";

import jwtDecode from "jwt-decode";

import axios from "@/common/axios";
import router from "@/router";

Vue.use(Vuex);

const SET_TOKENS = "SET_TOKENS";
const FETCH_USER = "FETCH_USER";
const LOGOUT_USER = "LOGOUT_USER";

const user = {
    namespaced: true,
    state: {
        email: null,
        isLoggedIn: false,
        accessToken: null,
        refreshToken: null,
    },
    mutations: {
        SET_TOKENS(state, token) {
          state.accessToken = token.access_token;
          state.refreshToken = token.refresh_token;
          localStorage.setItem("accessToken", token.access_token);
          localStorage.setItem("refreshToken", token.refresh_token);
        },
        FETCH_USER(state, user) {
          let userInfo = jwtDecode(user);
          state.email = userInfo.username;
          state.accessToken = localStorage.getItem("accessToken");
          state.refreshToken = localStorage.getItem("refreshToken")
          state.isLoggedIn = true;
        },
        LOGOUT_USER(state) {
          state.email = null;
          state.accessToken = null;
          state.refreshToken = null;
          localStorage.removeItem("accessToken");
          localStorage.removeItem("refreshToken");
        },
    },
    getters: {
        user(state) {
            return state.email;
        },
        accessToken(state) {
            return state.accessToken;
        },
    },
    actions: {
        async userLogin({ commit }, credentials) {
            try {
                const response = await axios.post("/login", {
                    email: credentials.email,
                    password: credentials.password
                });
                commit(SET_TOKENS, response.data)
                router.push({ name: "list_parties" }).catch(err => {console.log(err)});
            } catch (error) {
                return Promise.reject(error);
            }
        },
        async validateToken({ commit, dispatch }, token) {
          try {
            const response = await axios.post("/token", { access_token: token, token_type: 'bearer' }); 
            commit(FETCH_USER, response.data.access_token)
            return response;
          } catch(error) {
            // attempt to refresh token
            dispatch('refreshToken');
          }
        },
        async refreshToken({ commit }) {
          try {
            let token = localStorage.getItem("refreshToken");
            const response = await axios.post("/token", { access_token: token, token_type: 'bearer' }); 
            commit(SET_TOKENS, response.data)
            return response;
          } catch(error) {
            commit(LOGOUT_USER);
            return Promise.reject(error);
          }
        },
        async userLogout({ commit }) {
            try {
                commit(LOGOUT_USER);
                router.push({ name: "login" }).catch(err => {console.log(err)});
            } catch (error) {
                return Promise.reject(error);
            }
        },
    }
};

export default user;
