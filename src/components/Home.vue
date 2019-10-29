<template>
  <div class="hello">
    <h1>{{ msg }}</h1>
    <p>
    you are now logged in!
    <button type="button" value="Logout" @click="logout">Logout</button>
    </p>
  </div>
</template>

<script>
  export default {
    name: 'Home',
    props: {
      msg: String
    },
    methods: {
      async logout() {
        await this.$store.dispatch("user/userLogout");
        this.$router.push({ name: "login" });
        document.location.href = "/";
      }
    }

  }

</script>


