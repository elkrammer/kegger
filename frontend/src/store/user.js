// this store sets parameters for the current logged in user
// eg: 'user': {'user': "admin@admin.com", 'acccessToken': "...", 'refreshToken': "..." }

import Vue from "vue";
import Vuex from "vuex";
import jwtDecode from "jwt-decode";
import axios from "@/common/axios";
import store from "@/store/index";

Vue.use(Vuex);

let isAlreadyFetchingAccessToken = false;
let subscribers = [];

function onAccessTokenFetched(access_token) {
    subscribers = subscribers.filter(callback => callback(access_token))
}

function addSubscriber(callback) {
    subscribers.push(callback)
}

axios.interceptors.response.use(function (response) {
    return response
}, function (error) {
    const { config, response: { status } } = error
    const originalRequest = config

    if (status === 401 && error.response.data.message === "invalid or expired jwt") {
        if (!isAlreadyFetchingAccessToken) {
            isAlreadyFetchingAccessToken = true
            store.dispatch("user/refreshUserTokens").then((access_token = localStorage.getItem('accessToken')) => {
                isAlreadyFetchingAccessToken = false
                onAccessTokenFetched(access_token)
            })
        }

        const retryOriginalRequest = new Promise((resolve) => {
            addSubscriber(access_token => {
                originalRequest.headers.Authorization = 'Bearer ' + access_token
                resolve(axios(originalRequest))
            })
        })
        return retryOriginalRequest
    }
    return Promise.reject(error)
})

const SET_USER = "SET_USER";
const STORE_ACCESS_TOKEN = "STORE_ACCESS_TOKEN";
const STORE_REFRESH_TOKEN = "STORE_REFRESH_TOKEN";
const STORE_TOKEN_EXPIRATION = "STORE_TOKEN_EXPIRATION";
const LOGOUT_USER = "LOGOUT_USER";

const user = {
    namespaced: true,
    state: {
        user: null,
        accessToken: null,
        refreshToken: null,
        tokenExpiration: null,
    },
    mutations: {
        SET_USER(state, data) {
            state.user = data;
        },
        STORE_ACCESS_TOKEN(state, accessToken) {
            state.accessToken = accessToken
            localStorage.setItem("accessToken", accessToken);
        },
        STORE_REFRESH_TOKEN(state, refreshToken) {
            state.refreshToken = refreshToken
            localStorage.setItem("refreshToken", refreshToken);
        },
        STORE_TOKEN_EXPIRATION(state, tokenExpiration) {
            state.tokenExpiration = tokenExpiration
            localStorage.setItem("tokenExpiration", tokenExpiration);
        },
        LOGOUT_USER(state) {
            state.user = null;
            state.accessToken = null;
            state.refreshToken = null;
            state.tokenExpiration = null;
            localStorage.removeItem("accessToken");
            localStorage.removeItem("refreshToken");
        },
    },
    getters: {
        user(state) {
            return state.user;
        },
        accessToken(state) {
            return state.accessToken;
        },
        refreshToken(state) {
            return state.refreshToken;
        },
        tokenExpiration(state) {
            return state.tokenExpiration;
        },
    },
    actions: {
        async setUserAndTokens({ commit }, data) {
            try {
                let decoded = jwtDecode(data.accessToken);
                commit(SET_USER, decoded.username);
                commit(STORE_ACCESS_TOKEN, data.accessToken);
                commit(STORE_REFRESH_TOKEN, data.refreshToken);
                commit(STORE_TOKEN_EXPIRATION, decoded.exp);
            } catch (error) {
                return Promise.reject(error);
            }
        },
        async userLogin({ dispatch }, credentials) {
            try {
                const response = await axios.post("/login", {
                    email: credentials.email,
                    password: credentials.password
                });
                return await dispatch("setUserAndTokens", {
                    accessToken: response.data.access_token,
                    refreshToken: response.data.refresh_token
                });
            } catch (error) {
                return Promise.reject(error);
            }
        },
        async refreshUserTokens({ dispatch }) {
            try {
                const token = this.refreshToken || localStorage.getItem('refreshToken');
                const response = await axios.post("/token", {
                    refresh_token: token
                });
                return await dispatch("setUserAndTokens", {
                    accessToken: response.data.access_token,
                    refreshToken: response.data.refresh_token
                });
            } catch (error) {
                return Promise.reject(error);
            }
        },
        async userLogout({ commit }) {
            try {
                commit(LOGOUT_USER);
            } catch (error) {
                return Promise.reject(error);
            }
        },
    }
};

export default user;
