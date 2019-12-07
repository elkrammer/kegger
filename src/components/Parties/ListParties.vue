<template>
  <section class="section">
    <div class="container">

      <section>
        <b-modal
          :active.sync="createPartyActive"
          has-modal-card
          trap-focus
          aria-role="dialog"
          aria-modal>
          <CreateParty/>
        </b-modal>

        <b-modal
          :active.sync="editPartyActive"
          has-modal-card
          trap-focus
          aria-role="dialog"
          aria-modal>
          <EditParty v-if="selected !== null" :party_id="selected.id"/>
        </b-modal>

        <b-modal
          :active.sync="editGuestsActive"
          has-modal-card
          trap-focus
          aria-role="dialog"
          aria-modal>
          <EditGuests v-if="selected !== null" :party_id="selected.id"/>
        </b-modal>

        <b-modal
          :active.sync="deletePartyActive"
          has-modal-card
          trap-focus
          aria-role="dialog"
          aria-modal>
          <DeleteParty v-if="selected !== null" :party_id="selected.id" :party_name="selected.name"/>
        </b-modal>
      </section>

      <div class="columns">

        <div class="column">
          <h1 class="title">Parties</h1>
          <div style="margin-bottom: 5px" class="buttons">
            <b-button @click="createPartyActive = true" type="is-success">Create New Party</b-button>
            <b-button @click="editPartyActive = true" v-if="selected !== null" type="is-info">Edit Party</b-button>
            <b-button @click="editGuestsActive = true" v-if="selected !== null" class="is-dark">Edit Guests</b-button>
            <b-button @click="deletePartyActive = true" v-if="selected !== null" type="is-danger">Delete Party</b-button>
          </div>
        </div>

        <Statistics/>

      </div>

      <div class="columns">
        <div class="column is-half">
          <b-field>
            <b-input placeholder="Party Name" v-model="search" icon="search"></b-input>
            <p class="control">
            <b-button @click="search = ''" v-if="search !== ''" type="is-warning">Clear</b-button>
            <b-button @click="unSelect" v-if="selected !== null" style="margin-left: 50px;" type="is-warning">Unselect</b-button>
            </p>
          </b-field>
        </div>
      </div>

      <b-table
        id="parties-table"
        :data="filteredPartyList"
        ref="table"
        hoverable
        detailed
        paginated
        show-detail-icon
        custom-detail-row
        per-page="10"
        :selected.sync="selected"
        default-sort="Name"
        default-sort-direction="asc"
        sort-icon="chevron-up"
        sort-icon-size="is-small"
        detail-key="id"
        >
        <template slot-scope="props">

          <b-table-column field="Name" label="Party Name" width="200" sortable>
            {{ props.row.name }}
          </b-table-column>

        <b-table-column field="HostName" label="Host" width="200" sortable>
          {{ props.row.host_name }}
        </b-table-column>

        <b-table-column field="isPartyAttending" label="Attending" width="100" visible sortable>
          <div class="has-text-success" v-if="props.row.isPartyAttending === 'true'">
            <span class="icon-text">Yep</span>
            <b-icon pack="fas" icon="thumbs-up">
            </b-icon>
          </div>

          <div class="has-text-warning" v-if="props.row.isPartyAttending === 'partial'">
            <span class="icon-text">Partial</span>
            <b-icon pack="fas" icon="exclamation-triangle">
            </b-icon>
          </div>

          <div class="has-text-danger" v-if="props.row.isPartyAttending === 'false'">
            <span class="icon-text">Nay</span>
            <b-icon pack="fas" icon="thumbs-down">
            </b-icon>
          </div>
        </b-table-column>


        <b-table-column field="plusOne" label="Plus One" width="100">
          &#10240;
        </b-table-column>


        <b-table-column field="Comments" label="Comments" width="200" sortable>
          {{ props.row.comments }}
        </b-table-column>

        </template>

        <template slot="detail" slot-scope="props">
          <tr v-for="guest in props.row.Guests" :key=" 'A' + guest.id">
            <td></td>
            <td style="padding-left: 30px;">{{ guest.first_name }} {{ guest.last_name }}</td>
            <td></td>
            <td>
              <div class="has-text-success" v-if="guest.is_attending === true">
                <span class="icon-text">Yep</span>
                <b-icon pack="fas" icon="thumbs-up">
                </b-icon>
              </div>
              <div class="has-text-danger" v-else>
                <span class="icon-text">Nay</span>
                <b-icon pack="fas" icon="thumbs-down">
                </b-icon>
              </div>
            </td>

            <td>
              <div class="has-text-success" v-if="guest.plus_one === true">
                <span>Yes</span>
              </div>
              <div v-else>
                <span>No</span>
              </div>
            </td>

            <td></td>
          </tr>
        </template>

      </b-table>

    </div>

  </section>
</template>

<script>
  import { mapGetters } from "vuex";
  import CreateParty from "@/components/Parties/CreateParty.vue";
  import EditParty from "@/components/Parties/EditParty.vue";
  import EditGuests from "@/components/Guests/EditGuests.vue";
  import DeleteParty from "@/components/Parties/DeleteParty.vue";
  import Statistics from "@/components/Parties/Statistics.vue";

  export default {
    name: 'list_parties',
    components: { CreateParty, EditParty, DeleteParty, EditGuests, Statistics },
    data() {
      return {
        createPartyActive: false,
        editPartyActive: false,
        deletePartyActive: false,
        editGuestsActive: false,
        selected: null,
        isPartyAttending: null,
        search: '',
        statistics: {
          guestsAttending: 0,
          guestsNotAttending: 0,
          totalGuests: 0,
          guestsPending: 0,
        },
      };
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
      toggle(row) {
        this.$refs.table.toggleDetails(row)
      },
      unSelect() {
        this.selected = null;
      },
    },
    computed: {
      ...mapGetters({
        parties: "party/parties"
      }),
      filteredPartyList() {
        return this.parties.filter(party => {
          return party.name.toLowerCase().includes(this.search.toLowerCase())
        })
      },
    },
    created() {
      this.loadParties();
    },
  };
</script>

<style lang="scss">
@import "@/variables";

.icon-text {
  margin-right: 10px;
}

#parties-table td {
    word-break: break-all;
    vertical-align: middle;
}

</style>
