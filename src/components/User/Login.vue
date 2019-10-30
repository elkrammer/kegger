<template>
  <section class="hero is-primary is-fullheight has-background-info">
    <div class="hero-body">
      <div class="container">
        <div class="columns is-centered">
          <div class="column is-5-tablet is-4-desktop is-3-widescreen">
            <img src="@/assets/logo.png" alt="WingDing Logo" />

            <b-message v-if="errors.length" :auto-close=true :duration="8000" type="is-danger">
              <ul>
                <li v-for="(error, index) in errors" :key="index">{{ error }}</li>
              </ul>
            </b-message>

            <form class="box" v-on:submit.prevent autocomplete="off">

              <b-field label="Email">
                <div class="control has-icons-left">
                  <b-input type="email"
                           placeholder="e.g. john@doe.com"
                           v-model.trim="email"
                           @input="delayTouch($v.email)"
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
                         required>
                  </b-input>
                  <span class="icon is-small is-left">
                    <i class="fa fa-lock"></i>
                  </span>
                </div>
              </b-field>

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

  const touchMap = new WeakMap()

  export default {
    name: "login",
    data() {
      return {
        email: "",
        password: "",
        submitted: false,
        errors: [],
      };
    },
    methods: {
      async submit() {
        if (this.$v.$invalid) {
          this.$v.$touch();
          return;
        }

        this.submitted = true;
        const credentials = {
          email: this.email,
          password: this.password
        };

        try {
          await this.$store.dispatch("user/userLogin", credentials);
          this.email= "";
          this.password = "";
          this.$v.$reset();
          this.$router.push({ name: "home" });
        } catch (error) {
          this.errors.push("Invalid Login Credentials");
        } finally {
          this.submitted = false;
        }
      },
      delayTouch($v) {
        $v.$reset()
        if (touchMap.has($v)) {
          clearTimeout(touchMap.get($v));
        }
        touchMap.set($v, setTimeout($v.$touch, 1000))
      }
    },
    validations: {
      email: { required, email },
      password: { required },
    }
  };
</script>
