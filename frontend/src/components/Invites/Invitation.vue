<template>

  <section
    class="invite-background"
    :style="{ 'background-image': inviteBackground }"
    >
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

      <article class="message">
        <div class="message-header">
          <p>RSVP</p>
        </div>
        <div class="message-body">
          <b-switch type="is-active" size="is-large" v-model="attending">
            Are you attending?
          </b-switch>
        </div>
      </article>

    </div>

    <div v-else>
      <h1>Invalid Invitation ID</h1>
    </div>

  </section>
</template>

<script>
import { mapGetters } from "vuex";

export default {
  name: "invitation",
  props: ['id'],
  data() {
    return {
      attending: null,
    }
  },
  computed: {
    ...mapGetters({
      invite: "invite/invite",
    }),
    inviteBackground() {
      const path = process.env.VUE_APP_API_SERVER + this.invite.invite_background;
      return 'url("' + path + '")';
    }
  },
  methods: {
    async getInvite() {
      try {
        await this.$store.dispatch("invite/getInvite", this.id);
        // update invitation_opened if this is the first time it's being opened
        if (this.invite && this.invite.guest.invitation_opened == null) {
          this.setInviteOpened();
        }
        // sync local attending state with the invite
        this.attending = this.invite.guest.is_attending;
      } catch (error) {
        console.log(error);
      }
    },
    async setInviteOpened() {
      try {
        const data = {
          action: "opened",
          invitation_id: this.id
        };
        await this.$store.dispatch("invite/updateInvite", data);
      } catch (error) {
        console.log(error);
      }
    },
    async setAttending() {
      try {
        const data = {
          action: "attending",
          invitation_id: this.id,
          is_attending: this.attending
        };
        await this.$store.dispatch("invite/updateInvite", data);
      } catch (error) {
        console.log(error);
      }
    },
  },
  watch: {
    attending: function() {
      this.setAttending();
    }
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

.message {
  text-align: center;
  margin: auto;
  margin-top: 30px;
  width: 30%;
}

.invite-background {
  background-color: #3D9970;
  background-size: cover;
  background-repeat: no-repeat;
  background-position: center;
  background-color: #3D9970;
}

.invite-main {
  height: 1000px;
}

.title-box {
  padding-top: 50px;
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
