<template>
  <section>

    <div class="modal-card">
      <header class="modal-card-head">
        <p class="modal-card-title">Create User</p>
        <b-icon icon="user-plus" class="is-pulled-right" size="is-large"></b-icon>
      </header>
      <section class="modal-card-body">


        <form v-on:submit.prevent>

          <b-field label="Name">
            <b-input placeholder="Name" icon="id-card" v-model="user.name"></b-input>
          </b-field>

          <b-field label="Email">
            <b-input placeholder="Email" icon="envelope" v-model="user.email"></b-input>
          </b-field>

          <b-field label="Password">
            <b-input type="password" icon="lock" placeholder="Password" v-model="user.password"></b-input>
          </b-field>

        </form>
      </section>
      <footer class="modal-card-foot">
        <button class="button is-success" @click="createUser" :disabled="$v.$invalid">Create</button>
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
        submitted: false,
      }
    },
    computed: {
      ...mapGetters({
        users: "user/users",
      })
    },
    methods: {
      async createUser() {
        try {

            this.submitted = true;
            await this.$store.dispatch("users/createUser", this.user);
            const msg = `User ${this.user.name} created!`
            this.$buefy.toast.open({
                message: msg,
                type: 'is-success',
                position: 'is-bottom',
                duration: 3000,
            })
            this.$parent.close();
        } catch (error) {
            const msg = `There was an error creating user ${this.user.name}`
            this.$buefy.toast.open({
                message: msg,
                type: 'is-danger',
                position: 'is-bottom',
                duration: 3000,
            })
            console.log(error);
        }
      },
    },
  }
</script>
