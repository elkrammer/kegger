// this store sets parameters for the current logged in user
// eg: 'user': {'user': "admin@admin.com", 'acccessToken': "...", 'refreshToken': "..." }

import Vue from "vue";
import Vuex from "vuex";
import jwtDecode from "jwt-decode";
import axios from "@/common/axios";
import store from "@/store/index";
import router from "@/router";

Vue.use(Vuex);

let authTokenRequest;

function getAuthToken() {
    // if our currently stored token expires soon
    if (store.getters['user/tokenExpiration'] - 240 <= (Date.now() / 1000).toFixed(0)) {
        // if not already requesting a token
        if (!authTokenRequest) {
            authTokenRequest = store.dispatch("user/refreshUserTokens");
            authTokenRequest.then(() => {
                authTokenRequest = null; // reset authTokenRequest
            }).catch(err => {
                console.log("Error in getAuthToken: " + err);
            });
            return authTokenRequest;
        }
    }
}

async function logoutOfApplication() {
    await store.dispatch("user/userLogout")
    router.replace({ name: "login" });
}

// axios interceptor that will handle expired tokens
axios.interceptors.response.use(undefined, async error => {
    // prevent endless redirects
    if (error.request !== undefined) {
        if (error.request.responseURL.includes('login')) {
            return Promise.reject(error)
        }
    }

    if (error.response && error.response.status === 401 &&
        error.response.data.message === "invalid or expired jwt" &&
        store.getters['user/retries'] < 3) {
        try {
            let response = await getAuthToken();

            if (!response) {
                store.dispatch("user/increaseRetry");
                return Promise.reject(error.config);
            }

            await store.dispatch("user/setUserAndTokens", {
                accessToken: response.data.access_token,
                refreshToken: response.data.refresh_token
            });

            error.config.headers["Authorization"] = "Bearer " + store.getters["user/accessToken"];
            error.config.baseURL = ""; // baseURL needs to be zeroud out to prevent tripping over the baseURL set in main axios instance
            return axios.request(error.config);
            // return axios(error.config);
        } catch (error) {
            return
        }
    } else if (error.response && error.response.status === 401 && error.response.data.message === "invalid or expired jwt" && store.getters['user/retries'] > 2) {
        logoutOfApplication();
        return Promise.reject(error);
    }

    // if someone gets here we don't want to log them out since it's more of a general error
    return Promise.reject(error);
});

const SET_USER = "SET_USER";
const STORE_ACCESS_TOKEN = "STORE_ACCESS_TOKEN";
const STORE_REFRESH_TOKEN = "STORE_REFRESH_TOKEN";
const LOGOUT_USER = "LOGOUT_USER";
const INCREASE_RETRY = 'INCREASE_RETRY';

const user = {
    namespaced: true,
    state: {
        user: null,
        accessToken: null,
        refreshToken: null,
        tokenExpiration: null,
        retries: 0,
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
        LOGOUT_USER(state) {
            state.user = null;
            state.accessToken = null;
            state.refreshToken = null;
            state.tokenExpiration = null;
            state.retries = 0;
            localStorage.removeItem("accessToken");
            localStorage.removeItem("refreshToken");
        },
        INCREASE_RETRY(state) {
            state.retries += 1
        }
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
        retries(state) {
            return state.retries;
        },
    },
    actions: {
        async setUserAndTokens({ commit }, data) {
            try {
                let decoded = jwtDecode(data.accessToken);
                commit(SET_USER, decoded.username);
                commit(STORE_ACCESS_TOKEN, data.accessToken);
                commit(STORE_REFRESH_TOKEN, data.refreshToken);
                this.tokenExpiration = decoded.exp;
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
        increaseRetry({ commit }) {
            commit(INCREASE_RETRY);
        },
    }
};

export default user;
