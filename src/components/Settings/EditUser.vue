<template>
  <section>

    <div class="modal-card">
      <header class="modal-card-head">
        <p class="modal-card-title">User {{ user.name }}</p>
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
            <b-input type="password" icon="lock" placeholder="Password" v-model="password"></b-input>
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
          this.$parent.close();
        } catch (error) {
          console.log(error);
        }
      },
    },
    created() {
      this.getUser();
    },
  }
</script>
