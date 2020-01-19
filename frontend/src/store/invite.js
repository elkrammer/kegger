// this store is used to manage invites
import Vue from "vue";
import Vuex from "vuex";
import axios from "@/common/axios";
import { setAuthorizationHeader } from "@/common/utilities";

Vue.use(Vuex);

const GET_INVITE = "GET_INVITE";
const UPDATE_INVITE = "UPDATE_INVITE";

const invite = {
  namespaced: true,
  state: {
    invite: [],
  },
  mutations: {
    GET_INVITE(state, data) {
      state.invite = data;
    },
  },
  getters: {
    invite(state) {
      return state.invite;
    },
  },
  actions: {
    async sendInvite({ rootGetters }, guest_id) {
      try {
        setAuthorizationHeader(rootGetters["user/accessToken"]);
        const response = await axios.post("/api/sendinvite/" + guest_id, {
          from_name: process.env.VUE_APP_EMAIL_FROM_NAME,
          from_email: process.env.VUE_APP_EMAIL_FROM,
          to: [
            {
              name: process.env.VUE_APP_EMAIL_TO_NAME, // DEBUGGING
              email: process.env.VUE_APP_EMAIL_TO // DEBUGGING
            },
          ],
          subject: process.env.VUE_APP_EMAIL_SUBJECT,
        });
        return response.data;
      } catch (error) {
        return Promise.reject(error);
      }
    },
    async getInvite({ commit }, id) {
      try {
        const response = await axios.get("/invite/" + id);
        commit(GET_INVITE, response.data);
        return response.data;
      } catch (error) {
        return Promise.reject(error);
      }
    },
    async updateInvite({ commit }, data) {
      try {
        if (data.action === "opened") {
          var date = new Date();
          const response = await axios.put("/invite", {
            invitation_id: data.invitation_id,
            invitation_opened: date.toISOString()
          });
          commit(UPDATE_INVITE, response.data);
        } else if (data.action === "attending") {
          const response = await axios.put("/invite", {
            invitation_id: data.invitation_id,
            is_attending: data.is_attending
          });
          commit(UPDATE_INVITE, response.data);
        }
      } catch (error) {
        return Promise.reject(error);
      }
    },
  }
};

export default invite;
