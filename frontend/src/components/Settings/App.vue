<template>
  <div class="box">
    <h1 class="title">App Settings</h1>

    <div v-for="setting in settings" :key="setting.name">

      <b-field :label="setting.description">
        <b-input :value="setting.value" v-model="setting.value"></b-input>
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
  name: "app_settings",
  data() {
    return {
    }
  },
  computed: {
    ...mapGetters({
      settings: "settings/appSettings",
    }),
  },
  methods: {
    async getAppSettings() {
      try {
        const response = await this.$store.dispatch("settings/getSettings");
        return response.data;
      } catch (error) {
        console.log(error);
      }
    },
    async saveSettings() {
      try {
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
    if (!this.settings) {
      this.getAppSettings();
    }
  },
}
</script>
