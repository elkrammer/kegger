// this store is used to manage invites
import Vue from "vue";
import Vuex from "vuex";
import axios from "@/common/axios";
import { setAuthorizationHeader } from "@/common/utilities";

Vue.use(Vuex);

const SEND_INVITE = "SEND_INVITE";

const invite = {
  namespaced: true,
  state: {
    invite: [],
  },
  mutations: {
    SEND_INVITE(state, data) {
      state.invite = data;
    },
  },
  actions: {
    async sendInvite({ commit, rootGetters }, guest_id) {
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
        commit(SEND_INVITE, response.data);
        return response.data;
      } catch (error) {
        return Promise.reject(error);
      }
    },
  }
};

export default invite;
