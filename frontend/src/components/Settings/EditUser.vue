<template>
  <section>

    <div class="modal-card">
      <header class="modal-card-head">
        <p class="modal-card-title">User {{ user.name }}</p>
      </header>
      <section class="modal-card-body">

        <form v-on:submit.prevent>

          <b-field label="Name">
            <b-input placeholder="Name" icon="id-card" v-model="user.name" />
          </b-field>

          <b-field label="Email">
            <b-input placeholder="Email" icon="envelope" v-model="user.email" />
          </b-field>

          <b-field label="Password">
            <b-input type="password" icon="lock" placeholder="Password" v-model="password" />
          </b-field>

        </form>

      </section>
      <footer class="modal-card-foot">
        <button class="button is-success" @click="editUser">Save</button>
        <button class="button" @click="$parent.close()">Close</button>
      </footer>
    </div>

  </section>
</template>

<script>
  import { mapGetters } from "vuex";

  export default {
    name: "edit_user",
    props: ['user_id'],
    data() {
      return {
        user: [],
        password: '',
      }
    },
    computed: {
      ...mapGetters({
        users: "user/users",
      })
    },
    created() {
      this.getUser();
    },
    methods: {
      async getUser() {
        try {
          const response = await this.$store.dispatch("users/getUser", this.user_id);
          this.user = response;
        } catch (error) {
          console.log(error);
        }
      },
      async editUser() {
        if (this.password !== '') {
          this.user.password = this.password;
        }
        try {
            const response = await this.$store.dispatch("users/editUser", this.user);
            this.user = response;
            const msg = `User ${this.user.name} updated!`
            this.$buefy.toast.open({
                message: msg,
                type: 'is-success',
                position: 'is-bottom',
                duration: 3000,
            })
            this.$parent.close();
        } catch (error) {
            const msg = `There was an error editing user ${this.user.name}`
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
