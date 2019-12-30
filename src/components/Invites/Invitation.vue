<template>
  <div class="invite-main" v-if="Object.entries(this.invite).length > 0">

    <div class="title-box">
      <span>Save the Date!</span>
    </div>

    <div class="event-date">
      <span>{{ invite.event_date | moment('timezone', 'America/Toronto', 'MMM do YYYY H:mm') }}</span>
    </div>

    <div class="invite">
      <p>
        Hello {{ this.invite.guest.first_name }}!<br />
        You've been invited to {{ this.invite.event_name }}.<br /><br />
      </p>
      <span v-if="this.invite.guest.plus_one">
        You are allowed to bring a guest. <br /><br />
      </span>

      <span>
        This event will take place at {{ invite.event_location }}. <br /><br />
        The dress code for this event is {{ invite.dress_code }}. <br />
      </span>
    </div>

  </div>

  <div v-else>
    <h1>Invalid Invitation ID</h1>
  </div>

</template>

<script>
import { mapGetters } from "vuex";

export default {
  name: "invitation",
  props: ['id'],
  data() {
    return {

    }
  },
  computed: {
    ...mapGetters({
      invite: "invite/invite",
    }),
  },
  methods: {
    async getInvite() {
      try {
        await this.$store.dispatch("invite/getInvite", this.id);
      } catch (error) {
        console.log(error);
      }
    },
  },
  created() {
    this.getInvite();
  },
}
</script>

<style lang="scss" scoped>
@import "@/variables";
@import url('https://fonts.googleapis.com/css?family=Roboto+Condensed');
@import url('https://db.onlinewebfonts.com/c/19106915702fcec8936aa3e6cf1019aa?family=ConeriaScriptW01-Medium');

.invite-background {
  background-image: url("~@/assets/invite_bg.jpg");
  background-size: cover;
  background-repeat: no-repeat;
  background-position: center;
}

.invite-main {
  height: 4000px;
}

.title-box {
  font-family: 'ConeriaScriptW01-Medium', sans-serif;
  font-size: 6.0em;
  color: white;
  text-align: center;
}

.event-date {
  font-family: 'ConeriaScriptW01-Medium', sans-serif;
  font-size: 3.2em;
  color: yellow;
  text-align: center;
}

.invite {
  font-family: 'Roboto Condensed', sans-serif;
  font-size: 2.2em;
  margin: auto;
  margin-top: 50px;
  color: white;
  background-color: rgba(0,0,0,0.6);
  width: 80%;
}

</style>
