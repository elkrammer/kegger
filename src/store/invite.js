// this store is used to manage invites
import Vue from "vue";
import Vuex from "vuex";
import axios from "@/common/axios";
import { setAuthorizationHeader } from "@/common/utilities";

Vue.use(Vuex);

const GET_INVITE = "GET_INVITE";

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
          from_name: "Wingding",
          from_email: "wingding@wingding.com",
          to: [
            {
              name: "Mauricio",
              email: "mauricioc.bahamonde@gmail.com"
            },
          ],
          subject: "Test message from wingding",
          message: "<strong>test from wingding.</strong>"
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
  }
};

export default invite;
