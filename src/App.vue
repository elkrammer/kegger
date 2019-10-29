<template>
  <div id="app">
    <NavBar v-if="showNavBar" />
    <router-view/>
  </div>
</template>

<script>
import NavBar from "@/components/Layout/NavBar";

export default {
  name: "app",
  components: {
    NavBar,
  },
  computed: {
    showNavBar() {
      return this.$route.name !== 'login';
    }
  }
};
</script>

<style lang="scss">
@import "~bulma/sass/utilities/_all";
// bulma customizations go here

// import Bulma and Buefy styles
@import "~bulma";
@import "~buefy/src/scss/buefy";
</style>
