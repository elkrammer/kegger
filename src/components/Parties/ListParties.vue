<template>
  <section class="section">
    <div class="container">
      <h1 class="title">Parties</h1>

      <section>
        <div style="margin-bottom: 30px">
          <b-button @click="modalActive = true" type="is-success">Create New Party</b-button>
        </div>

        <b-modal
          :active.sync="modalActive"
          has-modal-card
          trap-focus
          aria-role="dialog"
          aria-modal>
          <CreateParty/>
        </b-modal>

      </section>

      <b-table
        :data="parties"
        ref="table"
        hoverable
        detailed
        show-detail-icon
        custom-detail-row
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
          <tr v-for="guest in props.row.Guests" :key="guest.index">
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

  export default {
    name: 'list_parties',
    components: { CreateParty },
    data() {
      return {
        modalActive: false,
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
      }
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
