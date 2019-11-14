<template>
  <form v-on:submit.prevent>
    <div class="modal-card" style="width: auto">
      <header class="modal-card-head">
        <p class="modal-card-title">Edit Party</p>
      </header>
      <section class="modal-card-body">
        <b-field label="Party Name">
          <b-input
            placeholder="Party Name"
            v-model="party.Name"
            required>
          </b-input>
        </b-field>

        <b-field label="Host">
          <b-select placeholder="Select Host" v-model="party.HostId">
            <option
              v-for="user in users"
              :value="user.id"
              :key="user.id"
              :v-bind:value="user.id">
            {{ user.name }}
            </option>
          </b-select>
        </b-field>

      </section>
      <footer class="modal-card-foot">
        <button class="button is-success" @click="editParty">Edit</button>
        <button class="button" @click="$parent.close()">Close</button>
      </footer>
    </div>
  </form>
</template>

<script>
  import { mapGetters } from "vuex";
  export default {
    name: 'edit_party',
    props: ['party_id'],
    data() {
      return {
        party: [],
      }
    },
    methods: {
      async getParty() {
        try {
          const response = await this.$store.dispatch("party/getParty", this.party_id);
          this.party = response;
        } catch (error) {
          console.log(error);
        }
      },
      async getHosts() {
        try {
          const response = await this.$store.dispatch("users/getUsers");
          return response.data;
        } catch (error) {
          console.log(error);
        }
      },
      async editParty() {
        try {
          await this.$store.dispatch("party/editParty", this.party);
          this.$parent.close();
        } catch (error) {
          console.log(error);
        }
      }
    },
    computed: {
      ...mapGetters({
        parties: "party/parties",
        users: "users/users"
      })
    },
    created() {
      this.getParty();
      this.getHosts();
    },
  }
</script>
