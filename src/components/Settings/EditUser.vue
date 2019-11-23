<template>
  <section>

    <div class="modal-card">
      <header class="modal-card-head">
        <p class="modal-card-title">User {{ user.name }}</p>
      </header>
      <section class="modal-card-body">
        <b-icon icon="dizzy" size="is-large"></b-icon>
      </section>
      <footer class="modal-card-foot">
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
    },
    created() {
      this.getUser();
    },
  }
</script>
