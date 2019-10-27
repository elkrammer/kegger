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
      <b-navbar-item href="#">
        Home
      </b-navbar-item>
      <b-navbar-item href="#">
        Parties
      </b-navbar-item>
    </template>
  </b-navbar>
</template>

<script>
  export default {
    name: "navbar",
  };
</script>
