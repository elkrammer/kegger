<template>
  <div class="box">
    <h1 class="title">Invite Settings</h1>

    <div v-for="setting in settings" :key="setting.name">

      <div class="columns is-multiline">
        <div class="column is-two-thirds">
          <b-field :label="setting.description">
            <b-input :value="setting.value" v-model="setting.value"></b-input>
          </b-field>

          <b-field>

            <b-upload v-if="setting.name == 'invite_image_en'" v-model="invite_image_en">
              <a class="button is-info is-rounded">
                <b-icon icon="upload"></b-icon>
                <span>Change</span>
              </a>
            </b-upload>

            <b-upload v-if="setting.name == 'invite_image_es'" v-model="invite_image_es">
              <a class="button is-info is-rounded">
                <b-icon icon="upload"></b-icon>
                <span>Change</span>
              </a>
            </b-upload>

            <b-upload v-if="setting.name == 'signature_image'" v-model="signature_image">
              <a class="button is-info is-rounded">
                <b-icon icon="upload"></b-icon>
                <span>Change</span>
              </a>
            </b-upload>

            <span class="file-name" v-if="setting.name == 'invite_image_es' && invite_image_es">
              {{ invite_image_es.name }}
            </span>

            <span class="file-name" v-if="setting.name == 'invite_image_en' && invite_image_en">
              {{ invite_image_en.name }}
            </span>

            <span class="file-name" v-if="setting.name == 'signature_image' && signature_image">
              {{ signature_image.name }}
            </span>

          </b-field>
        </div>

        <div class="column">
          <figure class="image">
            <img v-if="setting.name == 'invite_image_es'" :src="es_Invite_url" class="is-shady " @click="modal(es_Invite_url)">
            <img v-if="setting.name == 'invite_image_en'" :src="en_Invite_url" class="is-shady " @click="modal(en_Invite_url)">
            <img v-if="setting.name == 'signature_image'" :src="signature_image_url" class="is-shady " @click="modal(signature_image_url)">
          </figure>
        </div>

        <b-modal :active.sync="isModalActive">
          <p class="image is-4by5">
          <img :src="selected">
          </p>
        </b-modal>

      </div>
    </div>

    <br><br>
    <b-button is-rounded :disabled="!saveEnabled" @click="saveChanges" class="is-success">
      Save
    </b-button>

  </div>
</template>

<script>
import { mapGetters } from "vuex";
export default {
  name: "invite_settings",
  data() {
    return {
      isModalActive: false,
      invite_image_en: null,
      invite_image_es: null,
      signature_image: null,
      selected: null,
    }
  },
  computed: {
    ...mapGetters({
      settings: "settings/inviteSettings",
      appSettings: "settings/appSettings",
    }),
    es_Invite_url() {
      const img = this.settings.find(s => s.name === 'invite_image_es').value;
      const url = this.appSettings.find(s => s.name === 'kegger_api_url').value;
      return url + img;
    },
    en_Invite_url() {
      const img = this.settings.find(s => s.name === 'invite_image_en').value;
      const url = this.appSettings.find(s => s.name === 'kegger_api_url').value;
      return url + img;
    },
    signature_image_url() {
      const img = this.settings.find(s => s.name === 'signature_image').value;
      const url = this.appSettings.find(s => s.name === 'kegger_api_url').value;
      return url + img;
    },
    saveEnabled() {
      if (this.invite_image_en || this.invite_image_es || this.signature_image) {
        return true;
      }
      return false;
    },
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
    async saveChanges() {
      try {

        if (this.invite_image_en) {
          let data = new FormData();
          data.append('file', this.invite_image_en);
          data.append('name', 'invite_image_en')
          await this.$store.dispatch("settings/uploadInviteImage", data);
        }

        if (this.invite_image_es) {
          let data = new FormData();
          data.append('file', this.invite_image_es);
          data.append('name', 'invite_image_es')
          await this.$store.dispatch("settings/uploadInviteImage", data);
        }

        if (this.signature_image) {
          let data = new FormData();
          data.append('file', this.signature_image);
          data.append('name', 'signature_image')
          await this.$store.dispatch("settings/uploadSignatureImage", data);
        }

        this.getInviteSettings();

        const msg = `Invite Images updated!`
        this.$buefy.toast.open({
          message: msg,
          type: 'is-success',
          position: 'is-bottom',
          duration: 3000,
        })
        return;
      } catch (error) {
        console.log(error);
      }
    },
    modal(img) {
      this.selected = img;
      this.isModalActive = true;
    }
  },
  created() {
    if (!this.settings) {
      this.getInviteSettings();
    }
  },
}
</script>
