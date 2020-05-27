<template>
  <section class="hero is-primary is-fullheight has-background-black-ter">
    <div class="hero-body">
      <div class="container">
        <div class="columns is-centered">
          <div class="column is-5-tablet is-4-desktop is-3-widescreen">
            <img src="@/assets/logo.png" alt="Kegger Logo" />

            <b-message v-if="loginError" auto-close :duration="5000" type="is-danger">
              {{ loginError }}
            </b-message>

            <form ref="form" class="box" @submit.prevent="submit" autocomplete="off">

              <b-field label="Email">
                <div class="control has-icons-left">
                  <b-input type="email"
                           placeholder="e.g. john@doe.com"
                           v-model.trim="email"
                           required>
                  </b-input>
                  <span class="icon is-small is-left">
                    <i class="fa fa-envelope"></i>
                  </span>
                </div>
              </b-field>

              <b-field label="Password">
                <div class="control has-icons-left">
                  <b-input type="password"
                           placeholder="*******"
                           v-model.trim="password"
                           minlength="4"
                           required>
                  </b-input>
                  <span class="icon is-small is-left">
                    <i class="fa fa-lock"></i>
                  </span>
                </div>
              </b-field>

              <div class="field">
                <label for="checkbox" class="checkbox">
                  <input type="checkbox">
                  Remember me
                </label>
              </div>
              <div class="field">
                <button v-on:click="submit()" :disabled="!isFormValid()" class="button is-success">
                  Login
                </button>
              </div>
            </form>
          </div>
        </div>

        <br><br><br><br><br><br><br><br>
        <div class="content has-text-centered has-text-primary">
          <p>
          <a href="https://github.com/elkrammer/kegger">Kegger</a> - RSVP Manager<br />
          © 2020 <a href="https://github.com/elkrammer">Mauricio Bahamonde</a><br /><br />
          The source code is licensed under <a href="http://opensource.org/licenses/mit-license.php">MIT</a>.
          </p>
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
      email: "",
      password: "",
      loginError: "",
      submitted: false,
    };
  },
  methods: {
    async submit() {
      const credentials = {
        email: this.email,
        password: this.password
      };

      try {
        await this.$store.dispatch("user/userLogin", credentials);
        this.submitted = true;
        this.$router.push({ name: "list_parties" });
      } catch (error) {
        this.loginError = "Invalid Login Credentials";
      }

      // reset vars
      this.loginError = "";
      this.email= "";
      this.password = "";
    },
    isFormValid() {
      if (this.$refs.form && this.$refs.form.checkValidity()) {
        return true;
      }
      return false;
    }
  },
};
</script>
