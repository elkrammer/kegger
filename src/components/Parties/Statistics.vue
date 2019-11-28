<template>
  <div class="column is-two-fifths" v-if="this.parties.length > 0">
    <div class="box totalsbox">
      <nav class="level">
        <div class="level-item has-text-centered">
          <div>
            <p class="heading">Parties</p>
            <p class="title">{{ this.parties.length }}</p>
          </div>
        </div>

        <div class="level-item has-text-centered">
          <div>
            <p class="heading">Guests</p>
            <p class="title">{{ this.totalGuests }}</p>
          </div>
        </div>

        <div class="level-item has-text-centered">
          <div>
            <p class="heading">Attending</p>
            <p class="title">{{ this.guestsAttending }}</p>
          </div>
        </div>

        <div class="level-item has-text-centered">
          <div>
            <p class="heading">Pending</p>
            <p class="title">{{ this.guestsPending }}</p>
          </div>
        </div>

        <div class="level-item has-text-centered">
          <div>
            <p class="heading">Not Attending</p>
            <p class="title">{{ this.guestsNotAttending }}</p>
          </div>
        </div>
      </nav>
    </div>
  </div>

</template>

<script>
  import { mapGetters } from "vuex";

  export default {
    name: 'statistics',
    data() {
      return {
        guestsAttending: 0,
        guestsNotAttending: 0,
        guestsPending: 0,
        totalGuests: 0,
      };
    },
    methods: {
      statistics() {

       // reset stats
        this.guestsAttending = 0;
        this.guestsNotAttending = 0;
        this.guestsPending = 0;
        this.totalGuests = 0;

        for (var i = 0; i < this.parties.length; i++) { // parties loop
          let guestSize = this.parties[i].Guests.length;
          for (var j = 0; j < guestSize; j++) { // guests Loop
            // guest has plus_one
            if (this.parties[i].Guests[j].plus_one) {
              this.totalGuests += 2
            } else {
              this.totalGuests += 1
            }
            if (this.parties[i].Guests[j].is_attending) { // guest is attending
              this.guestsAttending += 1
            } else if (!this.parties[i].Guests[j].is_attending && this.parties[i].Guests[j].invitation_opened === null) { // guest hasn't seen invite
              this.guestsPending += 1
            } else if (!this.parties[i].Guests[j].is_attending && this.parties[i].Guests[j].invitation_opened !== null) { // guest saw the invite!
              this.guestsNotAttending += 1
            }
          }
        }
      },
    },
    computed: {
      ...mapGetters({
        parties: "party/parties"
      }),
    },
    watch: {
      parties() {
        this.statistics();
      },
    },
  };

</script>

<style lang="scss">
@import "@/variables";

.totalsbox {
  background-color: #f1ffe7;
}

</style>
