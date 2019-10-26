<template>
  <section class="hero is-primary is-fullheight has-background-info">
    <div class="hero-body">
      <div class="container">
        <div class="columns is-centered">
          <div class="column is-5-tablet is-4-desktop is-3-widescreen">
            <h1 class="title">Wingdings - RSVP Manager</h1>
            <form class="box" v-on:submit.prevent autocomplete="off">

              <div class="field">
                <label for="" class="label">Email</label>
                <div class="control has-icons-left">
                  <input type="email"
                         placeholder="e.g. john@doe.com"
                         class="input"
                         v-model.trim="credentials.username"
                         required>
                  <span class="icon is-small is-left">
                    <i class="fa fa-envelope"></i>
                  </span>
                </div>
              </div>
              <div class="field">
                <label for="" class="label">Password</label>
                <div class="control has-icons-left">
                  <input type="password"
                         placeholder="*******"
                         class="input"
                         v-model.trim="credentials.password"
                         required>
                  <span class="icon is-small is-left">
                    <i class="fa fa-lock"></i>
                  </span>
                </div>
              </div>
              <div class="field">
                <label for="" class="checkbox">
                  <input type="checkbox">
                  Remember me
                </label>
              </div>
              <div class="field">
                <button v-on:click="login()" class="button is-success">
                  Login
                </button>
              </div>
            </form>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script>
  export default {
    name: "login",
    data() {
      return {
        credentials: {
          username: "",
          password: ""
        },
        pending: false
      }
    }
  }
</script>
