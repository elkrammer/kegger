<template>
  <section class="section">
    <div class="container">

      <div class="columns">

        <div class="column">
          <h1 class="title">Invites</h1>
          <b-field>
            <div style="margin-bottom: 5px" class="buttons">
              <b-button :disabled="selected === null" type="is-success">View Invite</b-button>
              <b-button v-if="selected !== null" type="is-info">Send Invite</b-button>
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
        default-sort="Name"
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
            <div>{{ props.row.invitation_sent | dateParse('YYYY.MM.DD') | dateFormat('MMM DD YYYY') }}</div>
          </div>
        </b-table-column>

        <b-table-column label="Invite Opened" sortable>
          <div v-if="props.row.invitation_opened">
            <div>{{ props.row.invitation_opened | dateParse('YYYY.MM.DD') | dateFormat('MMM DD YYYY') }}</div>
          </div>
        </b-table-column>

        </template>
      </b-table>

    </div>
  </section>
</template>

<script>
  import { mapGetters } from "vuex";
  export default {
    name: "invites",
    data() {
      return {
        guests: [],
        selected: null,
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
      mapData() {
        for (var i=0; i < this.parties.length; i++) {
          for (var j=0; j < this.parties[i].Guests.length; j++) {
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
