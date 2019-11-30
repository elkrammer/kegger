<template>
  <section>

    <div class="modal-card">
      <header class="modal-card-head">
        <p class="modal-card-title">Invite for {{ guest.first_name }} {{ guest.last_name }}</p>
      </header>
      <section class="modal-card-body">
        <p>
        Hello {{ guest.first_name }}!<br />
        You've been invited to {{ eventName }}<br />
        to be celebrated on {{ eventDate }}.
        <br />
        <br />
        <span v-if="guest.plus_one">
          You are allowed to bring a guest.
        </span>

        </p>
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
        settings: "settings/settings",
      }),
      eventName() {
        return this.settings.find(s => s.name === 'event_name').value;
      },
      eventDate() {
        return this.settings.find(s => s.name === 'event_date').value;
      }

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
