<template>
  <b-navbar>

    <template slot="brand">
      <b-navbar-item tag="router-link" :to="{ name: 'home' }">
        <img
          src="@/assets/logo-small.png"
          alt="WingDing Logo"
          >
      </b-navbar-item>
    </template>

    <template slot="start">
      <b-navbar-item tag="router-link" :to="{ name: 'home' }">Home</b-navbar-item>
      <b-navbar-item tag="router-link" :to="{ name: 'parties' }">Parties</b-navbar-item>
    </template>

    <template slot="end">
      <b-navbar-item>
        <div class="buttons">
          <b-button @click="logout" type="is-primary">Log out</b-button>
        </div>
      </b-navbar-item>
    </template>

  </b-navbar>

</template>

<script>
  export default {
    name: "navbar",
    methods: {
      async logout() {
        await this.$store.dispatch("user/userLogout");
        this.$router.push({ name: "login" });
        document.location.href = "/";
      }
    }
  };
</script>
