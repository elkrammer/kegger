<template>
  <section class="hero is-primary is-fullheight has-background-info">
    <div class="hero-body">
      <div class="container">
        <div class="columns is-centered">
          <div class="column is-5-tablet is-4-desktop is-3-widescreen">
            <h1 class="title">Wingdings - RSVP Manager</h1>
            <form class="box" v-on:submit.prevent autocomplete="off">

              <div class="form-group"
                   v-bind:class="{ 'form-group--error': $v.credentials.email.$error,
                                   'form-group--success': !$v.credentials.email.$invalid && $v.credentials.email.$dirty}"
                   />

                <div class="field">
                  <label for="" class="label">Email</label>
                  <div class="control has-icons-left">
                    <input type="email"
                           placeholder="e.g. john@doe.com"
                           class="input"
                           v-model.trim="credentials.email"
                           required>
                    <span class="icon is-small is-left">
                      <i class="fa fa-envelope"></i>
                    </span>
                  </div>
                </div>

                <div
                  class="form-group__message"
                  v-if="$v.credentials.email.$error"
                  >
                  Please enter your email.
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

                  <div class="form-group__message" v-if="$v.credentials.password.$error">
                    Please enter your password.
                  </div>

                  <div class="field">
                    <label for="" class="checkbox">
                      <input type="checkbox">
                      Remember me
                    </label>
                  </div>
                  <div class="field">
                    <button v-on:click="submit()" class="button is-success">
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
  import { required, email } from "vuelidate/lib/validators";
  export default {
    name: "login",
    data() {
      return {
        credentials: {
          email: "",
          password: ""
        },
        pending: false
      };
    },
    methods: {
      async submit() {

        if (this.$v.$invalid) {
          this.$v.$touch();
          return;
        }

        this.pending = true;

        const credentials = {
          email: this.credentials.email,
          password: this.credentials.password
        };

        try {
          await this.$store.dispatch("user/userLogin", credentials);
          this.credentials.email= "";
          this.credentials.password = "";
          this.$v.$reset();
          this.$router.push({ name: "home" });
        } catch (error) {
          console.log("Invalid login");
        } finally {
          this.pending = false;
        }
      }
    },
    validations: {
      credentials: {
        email: { required, email },
        password: { required }
      }
    }
  };
</script>
