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
            v-model="party.name"
            required />
        </b-field>

        <b-field label="Host">
          <b-select placeholder="Select Host" v-model="party.host_id">
            <option
              v-for="user in users"
              :value="user.id"
              :key="user.id"
              :v-bind:value="user.id">
            {{ user.name }}
            </option>
          </b-select>
        </b-field>

        <b-field label="Comments">
          <b-input maxlength="200" rows="2" v-model="party.comments" type="textarea" />
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

  // TODO: do proper form validations
  export default {
    name: 'edit_party',
    props: ['party_id'],
    data() {
      return {
        party: [],
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
            const msg = `Changes saved!`
            this.$buefy.toast.open({
                message: msg,
                type: 'is-success',
                position: 'is-bottom',
                duration: 3000,
            })
            this.$parent.close();
        } catch (error) {
            const msg = `There was an error editing party: ${error}`
            this.$buefy.toast.open({
                message: msg,
                type: 'is-danger',
                position: 'is-bottom',
                duration: 3000,
            })
            console.log(error);
        }
      }
    },
  }
</script>
