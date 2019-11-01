<template>
  <section class="section">
    <div class="container">
      <h1 class="title">Parties</h1>

      <b-table
        :data="parties"
        ref="table"
        hoverable
        detailed
        show-detail-icon
        custom-detail-row
        detail-key="Id"
        >
        <template slot-scope="props">

          <b-table-column field="name" label="Party Name" sortable>
            {{ props.row.name }}
          </b-table-column>

        <b-table-column field="invitationOpened" label="Invite Seen?" sortable>
          <b-icon pack="fas"
                  :icon="props.row.invitationOpened === true ? 'thumbs-up' : 'thumbs-down'">
          </b-icon>
        </b-table-column>

        <b-table-column field="isAttending" label="Attending" sortable>
          <b-icon pack="fas"
                  :icon="props.row.isAttending === true ? 'thumbs-up' : 'thumbs-down'">
          </b-icon>
        </b-table-column>

        <b-table-column field="comments" label="Comments" sortable>
          {{ props.row.comments }}
        </b-table-column>

        </template>

        <template slot="detail" slot-scope="props">
          <tr v-for="guest in props.row.Guests" :key="guest.Id">
            <td></td>
            <td style="padding-left: 30px;">{{ guest.firstName }} {{ guest.lastName }}</td>
            <td>
              <!-- TODO: Fix me in model! :) -->
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

  export default {
    name: 'list_parties',
    data() {
      return {
        showDetailIcon: true,
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
    }
  };
</script>

<style lang="scss">
</style>
