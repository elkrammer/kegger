<template>
  <section class="section">
    <div class="container">

      <section>
        <b-modal
          :active.sync="viewInviteActive"
          has-modal-card
          trap-focus
          aria-role="dialog"
          aria-modal>
          <ViewInvite v-if="selected !== null" :guest_id="selected.id" />
        </b-modal>

      </section>

      <div class="columns">
        <div class="column">
          <h1 class="title">Invites</h1>
          <b-field>
            <div style="margin-bottom: 5px" class="buttons">
              <b-button @click="viewInviteActive = true" :disabled="selected === null" type="is-success">View Invite</b-button>
              <b-button v-if="selected !== null && !selected.invitation_sent" @click="sendInvite" type="is-info">Send Invite</b-button>
              <b-button v-if="selected !== null" @click="selected = null" style="margin-left: 50px;" type="is-warning">Unselect</b-button>
            </div>
          </b-field>
        </div>
      </div>

      <b-table
        :data="this.guests"
        ref="table"
        hoverable
        paginated
        default-sort="name"
        per-page="10"
        default-sort-direction="asc"
        :selected.sync="selected"
        sort-icon="chevron-up"
        sort-icon-size="is-small"
        >
        <template slot-scope="props">

          <b-table-column field="name" label="Guest Name" searchable sortable>
            {{ props.row.name }}
          </b-table-column>

        <b-table-column field="party_name" label="Party Name" searchable sortable>
          {{ props.row.party_name }}
        </b-table-column>

        <b-table-column label="Invite Sent" sortable>
          <div v-if="props.row.invitation_sent">
            <div>{{ props.row.invitation_sent | moment('timezone', 'America/Toronto', 'M/D/YY H:mm') }}</div>
          </div>
        </b-table-column>

        <b-table-column label="Invite Opened" sortable>
          <div v-if="props.row.invitation_opened">
            <div>{{ props.row.invitation_opened | moment('timezone', 'America/Toronto', 'M/D/YY H:mm') }}</div>
          </div>
        </b-table-column>

        </template>
      </b-table>

    </div>
  </section>
</template>

<script>
import { mapGetters } from "vuex";
import ViewInvite from "@/components/Invites/ViewInvite.vue";

export default {
  name: "invites",
  components: { ViewInvite },
  data() {
    return {
      guests: [],
      selected: null,
      viewInviteActive: false,
    }
  },
  methods: {
    async loadParties() {
      try {
        const response = await this.$store.dispatch("party/getParties");
        return response.data;
      } catch (error) {
        console.log(error);
      }
    },
    async sendInvite() {
      try {
        const response = await this.$store.dispatch("invite/sendInvite", {
          id: this.selected.id,
          email: this.selected.email,
          name: this.selected.name
        });
        const msg = `Invitation to ${this.selected.name} successfully sent`
        this.$buefy.toast.open({
          message: msg,
          type: 'is-success',
          position: 'is-bottom',
          duration: 3000,
        })
        var d = new Date();
        var now = d.toISOString();
        const itemId = this.guests.findIndex(guest => guest.id === this.selected.id);
        this.guests[itemId].invitation_sent = now;
        return response.data;
      } catch (error) {
        const msg = `There was an error sending the invitation to ${this.selected.name}`
        this.$buefy.toast.open({
          message: msg,
          type: 'is-danger',
          position: 'is-bottom',
          duration: 3000,
        })
      }
    },
    // this function re-maps the data obtained from the loadParties API call into the this.guests array like so:
    // {"id":1,"party_refer":1,"invitation_id":"1UfD6v1YuzssdgJKxWSYRJZjlsS","invitation_sent":null,"invitation_opened":null,"email":"foo@bar.com","plus_one":false,"is_attending":false,"party_name":"foo","name":"foo"}
    mapData() {
      for (var i = 0; i < this.parties.length; i++) {
        for (var j = 0; j < this.parties[i].Guests.length; j++) {
          let guest = this.parties[i].Guests[j];
          guest.party_name = this.parties[i].name;
          guest.name = guest.first_name + ' ' + guest.last_name;
          delete guest.first_name;
          delete guest.last_name;
          this.guests.push(guest);
        }
      }
    },
  },
  computed: {
    ...mapGetters({
      parties: "party/parties"
    }),
  },
  created() {
    this.loadParties();
    this.mapData();
  }
}
</script>

<style lang="scss">
@import "@/variables";
</style>
