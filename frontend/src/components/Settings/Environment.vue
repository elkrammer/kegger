<template>
  <div class="box">
    <h1 class="title">Environment variables</h1>

    <div v-for="(value, name) in settings" :key="name">
      <b-field :label="name" type="is-info">
        <b-input :value="value" rounded></b-input>
      </b-field>
    </div>

  </div>
</template>

<script>

export default {
  name: "environment_settings",
  data() {
    return {
      settings: {
        VUE_APP_API_SERVER: process.env.VUE_APP_API_SERVER,
        VUE_APP_EMAIL_FROM_NAME: process.env.VUE_APP_EMAIL_FROM_NAME,
        VUE_APP_EMAIL_FROM: process.env.VUE_APP_EMAIL_FROM,
        VUE_APP_EMAIL_DEBUG: process.env.VUE_APP_EMAIL_DEBUG,
        VUE_APP_EMAIL_TO: process.env.VUE_APP_EMAIL_TO
      },
    };
  },
}
</script>
