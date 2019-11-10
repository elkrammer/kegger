<template>
  <form v-on:submit.prevent action="">
    <div class="modal-card">
      <header class="modal-card-head">
        <p class="modal-card-title">Create New Party</p>
      </header>
      <section class="modal-card-body">
        <b-field label="Party Name">
          <b-input
            placeholder="Party Name"
            v-model="formProps.name"
            required>
          </b-input>
        </b-field>

        <b-field label="Host">
          <b-select placeholder="Select Host" v-model="formProps.host_id">
            <option
              v-for="user in users"
              :value="user.id"
              :key="user.id"
              :v-bind:value="user.id">
            {{ user.name }}
            </option>
          </b-select>
        </b-field>

        <div class="guests">
          <b-field label="Guests"></b-field>
          <b-field grouped group-multiline>

            <b-field label="First Name">
              <b-input
                placeholder="First Name"
                v-model="guest.first_name"
                required>
              </b-input>
            </b-field>

            <b-field label="Last Name">
              <b-input
                placeholder="Last Name"
                v-model="guest.last_name"
                required>
              </b-input>
            </b-field>

            <b-field label="Email">
              <b-input
                placeholder="Email"
                v-model="guest.email"
                required>
              </b-input>
            </b-field>

            <b-field>
              <p class="control">
              <button style="margin-top: 32px" class="button is-success" @click="addGuest()">Add Guest</button>
              </p>
            </b-field>

          </b-field>

          <div v-if="formProps.guests.length >0">
            <ul>
              <li v-for="(guest, index) in formProps.guests" :key="index">
                <b-icon icon="user"></b-icon>
                {{ guest.first_name }} {{ guest.last_name }}
              </li>
            </ul>
          </div>
        </div>

        <b-field label="Comments">
          <b-input maxlength="200" rows="2" v-model="formProps.comments" type="textarea"></b-input>
        </b-field>

      </section>
      <footer class="modal-card-foot">
        <button class="button is-success" @click="createParty()">Create</button>
        <button class="button" @click="$parent.close()">Close</button>
      </footer>
    </div>
  </form>
</template>

<script>
  import { mapGetters } from "vuex";
  export default {
    name: 'create_party',
    data() {
      return {
        formProps: {
          name: '',
          host_id: null,
          comments: '',
          guests: [],
        },
        guest: {
          first_name: '',
          last_name: '',
          email: '',
        },
      }
    },
    methods: {
      async getHosts() {
        try {
          const response = await this.$store.dispatch("users/getUsers");
          return response.data;
        } catch (error) {
          console.log(error);
        }
      },
      async createParty() {
        try {
          await this.$store.dispatch("party/createParty", this.formProps);
          this.$parent.close();
        } catch (error) {
          console.log(error);
        }
      },
      addGuest() {
        this.formProps.guests.push({...this.guest});
        this.guest.first_name = '';
        this.guest.last_name= '';
        this.guest.email = '';
      }
    },
    computed: {
      ...mapGetters({
        users: "users/users",
      }),
    },
    created() {
      this.getHosts();
    },
  }
</script>

<style lang="scss">
.guests {
  background-color: #fffaff;
}
</style>
