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
        const disabled = ["login", "invitation"];
        const showNavBar = !disabled.includes(this.$route.name);
        return showNavBar;
      }
    }
  };
</script>

<style lang="scss">
@import "variables";
@import "~bulma";
@import "~buefy/src/scss/buefy";
</style>
