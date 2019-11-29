<template>
  <div class="box">
    <h1 class="title">Event Settings</h1>

    <div v-for="setting in settings" :key="setting.name">
      <ul>
        <li>
          {{ setting.description }} : {{ setting.value }}
        </li>
      </ul>
    </div>

  </div>
</template>

<script>
  import { mapGetters } from "vuex";

  export default {
    name: "event_settings",
    computed: {
      ...mapGetters({
        settings: "settings/settings",
      })
    },
    methods: {
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
    }
  }
</script>
