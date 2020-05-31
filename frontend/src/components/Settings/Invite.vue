<template>
  <div class="box">
    <h1 class="title">Invite Settings</h1>

    <div v-for="setting in settings" :key="setting.name">

      <b-field :label="setting.description">
        <b-input :value="setting.value" v-model="setting.value"></b-input>
      </b-field>

      <b-collapse v-if="setting.name == 'invite_background'" :open="false" aria-id="changeBackground">
        <button
          class="button is-success"
          slot="trigger"
          aria-controls="contentIdForA11y1">
          Change Image
        </button>

        <br/>

        <b-field class="file">

          <b-upload v-model="file">
            <a class="button is-rounded">
              <b-icon icon="upload"></b-icon>
              <span>Click to upload</span>
            </a>
          </b-upload>
          <span class="file-name" v-if="file">
            {{ file.name }}
          </span>
          <b-button v-if="file" is-rounded @click="uploadInviteBgImage" class="is-info">
            Save
          </b-button>

        </b-field>

      </b-collapse>

    </div>

  </div>
</template>

<script>
import { mapGetters } from "vuex";
export default {
  name: "invite_settings",
  data() {
    return {
      file: null,
    }
  },
  computed: {
    ...mapGetters({
      settings: "settings/inviteSettings",
    }),
  },
  methods: {
    async getInviteSettings() {
      try {
        const response = await this.$store.dispatch("settings/getSettings");
        return response.data;
      } catch (error) {
        console.log(error);
      }
    },
    async uploadInviteBgImage() {
      try {
        let data = new FormData();
        data.append('file', this.file);
        const response = await this.$store.dispatch("settings/uploadInviteBgImage", data);
        const msg = `Invite Background image updated!`
        this.$buefy.toast.open({
          message: msg,
          type: 'is-success',
          position: 'is-bottom',
          duration: 3000,
        })
        this.settings.find(s => s.name === 'invite_background').value = './uploads/' + this.file.name;
        return response.data;
      } catch (error) {
        console.log(error);
      }
    },
  },
  created() {
    if (!this.settings) {
      this.getInviteSettings();
    }
  },
}
</script>
