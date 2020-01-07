<template>
  <form v-on:submit.prevent>
    <div class="modal-card" style="width: auto">
      <header class="modal-card-head">
        <p class="modal-card-title">Delete User</p>
      </header>
      <section class="modal-card-body">
        <p>Are you sure that you want to delete User {{ user_name }}?</p>
      </section>
      <footer class="modal-card-foot">
        <button class="button is-danger" @click="deleteUser">Delete</button>
        <button class="button" @click="$parent.close()">Close</button>
      </footer>
    </div>
  </form>
</template>

<script>
  import { mapGetters } from "vuex";
  export default {
    name: 'delete_user',
    props: ['user_id', 'user_name'],
    methods: {
      async deleteUser() {
        try {
            await this.$store.dispatch("users/deleteUser", this.user_id);
            const msg = `User ${this.user.name} deleted!`
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
    computed: {
      ...mapGetters({
        users: "users/users",
      })
    },
  }
</script>
