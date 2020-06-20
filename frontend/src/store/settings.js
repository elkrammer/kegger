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
            const invite = ['invite_image_en', 'invite_image_es'];
            return state.settings.filter(obj => {
                return invite.includes(obj.name)
            })
        },
        appSettings(state) {
            const app = ['wedding_website_url', 'kegger_frontend_url', 'kegger_api_url']
            return state.settings.filter(obj => {
                return app.includes(obj.name)
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
        async uploadInviteImage({ commit, rootGetters }, data) {
            try {
                setAuthorizationHeader(rootGetters["user/accessToken"]);
                const response = await axios.post("/api/settings/uploadInviteImg", data);
                commit(UPLOAD_INVITE_BG_IMAGE, response.data);
                return response.data;
            } catch(error) {
                return Promise.reject(error);
            }
        },
    }
};

export default settings;
