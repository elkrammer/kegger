<template>
  <b-navbar>

    <template slot="brand">
      <b-navbar-item tag="router-link" :to="{ path: '/' }">
        <img
          src="@/assets/logo-small.png"
          alt="WingDing Logo"
          >
      </b-navbar-item>
    </template>

    <template slot="start">
      <b-navbar-item>
        <router-link active-class="active" to="/home">Home</router-link>
      </b-navbar-item>
      <b-navbar-item>
        <router-link class="nav-link" active-class="active" to="/parties">Parties</router-link>
      </b-navbar-item>
    </template>

  </b-navbar>

</template>

<script>
  export default {
    name: "navbar",
  };
</script>
