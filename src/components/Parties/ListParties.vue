<template>
  <section class="section">
    <div class="container">
      <h1 class="title">Parties</h1>

      <section>
        <div style="margin-bottom: 30px" class="buttons">
          <b-button @click="createActive = true" type="is-success">Create New Party</b-button>
          <b-button @click="editActive = true" v-if="selected !== null" type="is-info">Edit Party</b-button>
          <b-button @click="deleteActive = true" v-if="selected !== null" type="is-danger">Delete Party</b-button>
          <b-button @click="unSelect" v-if="selected !== null" style="margin-left: 100px;" type="is-warning">Unselect</b-button>
        </div>

        <b-modal
          :active.sync="createActive"
          has-modal-card
          trap-focus
          aria-role="dialog"
          aria-modal>
          <CreateParty/>
        </b-modal>

        <b-modal
          :active.sync="editActive"
          has-modal-card
          trap-focus
          aria-role="dialog"
          aria-modal>
          <EditParty v-if="selected !== null" :party_id="selected.ID"/>
        </b-modal>

        <b-modal
          :active.sync="deleteActive"
          has-modal-card
          trap-focus
          aria-role="dialog"
          aria-modal>
          <DeleteParty v-if="selected !== null" :party_id="selected.ID" :party_name="selected.Name"/>
        </b-modal>

      </section>

      <b-table
        :data="parties"
        ref="table"
        hoverable
        detailed
        show-detail-icon
        custom-detail-row
        :selected.sync="selected"
        default-sort="Name"
        sort-icon="chevron-up"
        sort-icon-size="is-small"
        detail-key="ID"
        >
        <template slot-scope="props">

          <b-table-column field="Name" label="Party Name" sortable>
            {{ props.row.Name }}
          </b-table-column>

        <b-table-column field="HostName" label="Host" sortable>
          {{ props.row.HostName }}
        </b-table-column>

        <b-table-column field="IsAttending" label="Attending" sortable>

          <div class="has-text-success" v-if="props.row.IsAttending === true">
            <span class="icon-text">Yep</span>
            <b-icon pack="fas" icon="thumbs-up">
            </b-icon>
          </div>
          <div class="has-text-danger" v-else>
            <span class="icon-text">Nay</span>
            <b-icon pack="fas" icon="thumbs-down">
            </b-icon>
          </div>

        </b-table-column>

        <b-table-column field="InvitationSent" label="Invite Sent" sortable>
          <div v-if="props.row.InvitationSent !== null">
            {{ props.row.InvitationSent }}
          </div>
          <div v-else>
            Not Sent
          </div>
        </b-table-column>

        <b-table-column field="InvitationOpened" label="Invite Opened" sortable>
          <div v-if="props.row.InvitationOpened !== null">
            <div>{{ props.row.InvitationOpened | dateParse('YYYY.MM.DD') | dateFormat('MMM DD YYYY') }}</div>
          </div>
          <div v-else>
            Not Opened
          </div>
        </b-table-column>

        <b-table-column field="Comments" label="Comments" sortable>
          {{ props.row.comments }}
        </b-table-column>

        </template>

        <template slot="detail" slot-scope="props">
          <tr v-for="guest in props.row.Guests" :key="guest.id">
            <td></td>
            <td style="padding-left: 30px;">{{ guest.first_name }} {{ guest.last_name }}</td>
            <td></td>
            <td>
              <!-- TODO: Fix me in model! :) -->
              <span class="icon-text">Nay</span>
              <b-icon icon="thumbs-down"></b-icon>
            </td>
            <td>
              <b-icon pack="fas"
                      :icon="guest.isAttending === true ? 'thumbs-up' : 'thumbs-down'">
              </b-icon>
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
  import DeleteParty from "@/components/Parties/DeleteParty.vue";

  export default {
    name: 'list_parties',
    components: { CreateParty, EditParty, DeleteParty },
    data() {
      return {
        createActive: false,
        editActive: false,
        deleteActive: false,
        selected: null,
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
      })
    },
    created() {
      this.loadParties();
    },
  };
</script>

<style lang="scss">

.icon-text {
  margin-right: 10px;
}

</style>
