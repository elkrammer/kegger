// this store is used to manage application settings
import Vue from "vue";
import Vuex from "vuex";
import axios from "@/common/axios";
import { setAuthorizationHeader } from "@/common/utilities";

Vue.use(Vuex);

const GET_SETTINGS = "GET_SETTINGS";
const UPDATE_SETTINGS = "UPDATE_SETTINGS";
const UPLOAD_INVITE_BG_IMAGE = "UPDATE_SETTINGS";

const settings = {
    namespaced: true,
    state: {
        settings: [],
    },
    mutations: {
        GET_SETTINGS(state, data) {
            state.settings = data;
        },
        UPDATE_SETTINGS(state, data) {
            Vue.set(state.settings, data);
        },
    },
    getters: {
        eventSettings(state) {
            const event = [
                'event_name',
                'event_date',
                'event_location',
                'groom_name',
                'bride_name',
                'dress_code',
            ];
            return state.settings.filter(obj => {
                return event.includes(obj.name)
            })
        },
        inviteSettings(state) {
            const invite = ['invite_background'];
            return state.settings.filter(obj => {
                return invite.includes(obj.name)
            })
        },
    },
    actions: {
        async getSettings({ commit, rootGetters }) {
            try {
                setAuthorizationHeader(rootGetters["user/accessToken"]);
                const response = await axios.get("/api/settings");
                commit(GET_SETTINGS, response.data);
                return response.data;
            } catch (error) {
                return Promise.reject(error);
            }
        },
        async updateSettings({ commit, rootGetters }, data) {
            try {
                setAuthorizationHeader(rootGetters["user/accessToken"]);
                const response = await axios.put("/api/settings", data);
                commit(UPDATE_SETTINGS, response.data);
                return response.data;
            } catch (error) {
                return Promise.reject(error);
            }
        },
        async uploadInviteBgImage({ commit, rootGetters }, data) {
            try {
                setAuthorizationHeader(rootGetters["user/accessToken"]);
                const response = await axios.post("/api/settings/uploadInviteBg", data);
                commit(UPLOAD_INVITE_BG_IMAGE, response.data);
                return response.data;
            } catch(error) {
                return Promise.reject(error);
            }
        },
    }
};

export default settings;
