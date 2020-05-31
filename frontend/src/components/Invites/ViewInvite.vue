<template>
  <section>
    <div class="modal-card" style="width: 900px;">
      <header class="modal-card-head">
        <p class="modal-card-title">Invite for {{ guest.first_name }} {{ guest.last_name }}</p>
      </header>
      <section
        class="modal-card-body invite-background"
        :style="{ 'background-image': inviteBackground }"
        >
        <div class="invite-main" v-if="eventSettings.length > 0">

          <div class="title-box">
            <span>Save the Date!</span>
          </div>

          <div class="event-date">
            <span>{{ eventDate | moment('timezone', 'America/Toronto', 'MMM do YYYY H:mm') }}</span>
          </div>

          <div class="invite">
            <p>
            Hello {{ guest.first_name }}!<br />
            You've been invited to {{ eventName }}.<br /><br />
            </p>
            <span v-if="guest.plus_one">
              You are allowed to bring a guest. <br /><br />
            </span>

            <span>
              This event will take place at {{ eventLocation }}. <br /><br />
              The dress code for this event is {{ dressCode }}. <br />
            </span>
          </div>

        </div>

      </section>
      <footer class="modal-card-foot">
        <button class="button" @click="$parent.close()">Close</button>
      </footer>
    </div>

  </section>
</template>

<script>
import { mapGetters } from "vuex";

export default {
  name: "view_invite",
  props: ['guest_id'],
  data() {
    return {
      guest: [],
    }
  },
  computed: {
    ...mapGetters({
      guests: "guest/guests",
      eventSettings: "settings/eventSettings",
      inviteSettings: "settings/inviteSettings"
    }),
    eventName() {
      return this.eventSettings.find(s => s.name === 'event_name').value;
    },
    eventDate() {
      return this.eventSettings.find(s => s.name === 'event_date').value;
    },
    eventLocation() {
      return this.eventSettings.find(s => s.name === 'event_location').value;
    },
    dressCode() {
      return this.eventSettings.find(s => s.name === 'dress_code').value;
    },
    inviteBackground() {
      const path = process.env.VUE_APP_API_SERVER + this.inviteSettings.find(s => s.name === 'invite_background').value;
      return 'url("' + path + '")';
    }
  },
  methods: {
    async getGuest() {
      try {
        const response = await this.$store.dispatch("guest/getGuest", this.guest_id);
        this.guest= response;
      } catch (error) {
        console.log(error);
      }
    },
    async getSettings() {
      try {
        const response = await this.$store.dispatch("settings/getSettings");
        return response.data;
      } catch (error) {
        console.log(error);
      }
    },
  },
  created() {
    this.getSettings();
    this.getGuest();
  },
}
</script>


<style lang="scss" scoped>
@import "@/variables";
@import url('https://fonts.googleapis.com/css?family=Roboto+Condensed');
@import url('https://db.onlinewebfonts.com/c/19106915702fcec8936aa3e6cf1019aa?family=ConeriaScriptW01-Medium');

.invite-background {
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
