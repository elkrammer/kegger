<template>
  <section>

    <div class="modal-card">
      <header class="modal-card-head">
        <p class="modal-card-title">Create User</p>
        <b-icon icon="user-plus" class="is-pulled-right" size="is-large"></b-icon>
      </header>
      <section class="modal-card-body">

        <form ref="form" @submit.prevent="createUser">

          <b-field label="Name">
            <b-input placeholder="Name" icon="id-card" v-model="user.name" required></b-input>
          </b-field>

          <b-field label="Email">
            <b-input placeholder="Email" type="email" icon="envelope" v-model="user.email" required></b-input>
          </b-field>

          <b-field label="Password">
            <b-input placeholder="Password" type="password" icon="lock" v-model="user.password" required></b-input>
          </b-field>

        </form>
      </section>
      <footer class="modal-card-foot">
        <button class="button is-success" :disabled="!isFormValid()" @click="createUser">Create</button>
        <button class="button" @click="$parent.close()">Close</button>
      </footer>
    </div>

  </section>
</template>

<script>
import { mapGetters } from "vuex";

export default {
  name: "create_user",
  data() {
    return {
      user: {
        name: '',
        email: '',
        password: '',
      },
    }
  },
  computed: {
    ...mapGetters({
      users: "user/users",
    })
  },
  methods: {
    async createUser() {
      if (!this.isFormValid) {
        return;
      }
      try {
        await this.$store.dispatch("users/createUser", this.user);
        this.$parent.close();
        const msg = `User ${this.user.name} created!`
        this.$buefy.toast.open({
          message: msg,
          type: 'is-success',
          position: 'is-bottom',
          duration: 3000,
        })
      } catch (error) {
        console.log(error);
        const msg = `There was an error creating user ${this.user.name}`
        this.$buefy.toast.open({
          message: msg,
          type: 'is-danger',
          position: 'is-bottom',
          duration: 3000,
        })
      }
    },
    isFormValid() {
      if (this.$refs.form && this.$refs.form.checkValidity()) {
        return true;
      }
      return false;
    }
  },
}
</script>
