<template>
  <div class="box">
    <h1 class="title">Environment variables</h1>

    <div v-for="(value, name) in settings" :key="name">
      <b-field :label="name" type="is-info">
        <b-input :value="value" rounded disabled />
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
        VUE_APP_API_SERVER: config.API_SERVER,
        VUE_APP_EMAIL_FROM_NAME: config.EMAIL_FROM_NAME,
        VUE_APP_EMAIL_FROM: config.EMAIL_FROM,
        VUE_APP_EMAIL_DEBUG: config.EMAIL_DEBUG,
        VUE_APP_EMAIL_TO: config.EMAIL_TO
      },
    };
  },
}
</script>
