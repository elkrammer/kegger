<template>
  <div class="box">
    <h1 class="title">Event Settings</h1>

    <div v-for="setting in settings" :key="setting.name">
      <b-field :label="setting.description">
        <b-input v-if="setting.description != 'Event Date'" :value="setting.value" v-model="setting.value"></b-input>
        <b-datetimepicker
          placeholder="Click to select..."
          icon="calendar-alt"
          v-model="eventDate"
          v-if="setting.description === 'Event Date'"
          >
        </b-datetimepicker>
      </b-field>
    </div>

    <div class="buttons" style="margin-top: 30px;">
      <button class="button is-success" @click="saveSettings">Save</button>
    </div>

  </div>
</template>

<script>
import { mapGetters } from "vuex";

export default {
  name: "event_settings",
  data() {
    return {
      eventDate: new Date(),
    };
  },
  computed: {
    ...mapGetters({
      settings: "settings/eventSettings",
    }),
  },
  methods: {
    async getEventSettings() {
      try {
        const response = await this.$store.dispatch("settings/getSettings");
        return response.data;
      } catch (error) {
        console.log(error);
      }
    },
    async saveSettings() {
      try {
        // update our state with the new selected date
        this.settings.find(s => s.name === 'event_date').value = this.eventDate.toISOString();
        const response = await this.$store.dispatch("settings/updateSettings", this.settings);
        const msg = `Changes saved!`
        this.$buefy.toast.open({
          message: msg,
          type: 'is-success',
          position: 'is-bottom',
          duration: 3000,
        })
        return response.data;
      } catch (error) {
        const msg = `There was an error saving event settings: ${error}`
        this.$buefy.toast.open({
          message: msg,
          type: 'is-danger',
          position: 'is-bottom',
          duration: 3000,
        })
        console.log(error);
      }
    }
  },
  created() {
    this.getEventSettings();
    if (this.settings.length > 0) {
      let date = this.settings.find(s => s.name === 'event_date');
      this.eventDate = new Date(date.value);
    }
  }
}
</script>
