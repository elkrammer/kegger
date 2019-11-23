<template>
  <section>

    <div class="modal-card">
      <header class="modal-card-head">
        <p class="modal-card-title">Invite for {{ guest.first_name }} {{ guest.last_name }}</p>
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
    name: "view_invite",
    props: ['guest_id'],
    data() {
      return {
        guest: [],
      }
    },
    computed: {
      ...mapGetters({
        guests: "guest/guests",
      })
    },
    methods: {
      async getGuest() {
        try {
          const response = await this.$store.dispatch("guest/getGuest", this.guest_id);
          this.guest= response;
        } catch (error) {
          console.log(error);
        }
      },
    },
    created() {
      this.getGuest();
    },
  }
</script>
