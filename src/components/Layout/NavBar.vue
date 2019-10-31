<template>
  <b-navbar shadow>

    <template slot="brand">
      <b-navbar-item to="/">
        <img
          src="@/assets/logo-small.png"
          alt="WingDing Logo"
          >
      </b-navbar-item>
    </template>

    <template slot="start">
      <b-navbar-item tag="router-link" :to="{ name: 'home' }">
        <b-icon pack="fas" icon="home"></b-icon>
        <span class="item-text">Home</span>
      </b-navbar-item>
      <b-navbar-item tag="router-link" :to="{ name: 'listparties' }">
        <b-icon pack="fas" icon="users"></b-icon>
        <span class="item-text">Parties</span>
      </b-navbar-item>
    </template>

    <template slot="end">
      <b-navbar-item>
        <div class="buttons">
          <b-button @click="logout" type="is-link">Log out</b-button>
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

<style lang="scss">
@import "@/variables";
.navbar-start .navbar-item:first-of-type {
  margin-left: 30px;
}

.navbar-item span.item-text {
  margin-left: 8px;
}
</style>
