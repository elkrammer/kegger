<template>
  <form v-on:submit.prevent>
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
                v-model="formProps.guest.first_name"
                required>
              </b-input>
            </b-field>

            <b-field label="Last Name">
              <b-input
                placeholder="Last Name"
                v-model="formProps.guest.last_name"
                required>
              </b-input>
            </b-field>

            <b-field label="Email">
              <b-input
                placeholder="Email"
                v-model.trim="formProps.guest.email"
                required>
              </b-input>
            </b-field>

            <b-field>
              <p class="control">
              <button style="margin-top: 32px"
                      class="button is-success"
                      @click="addGuest()">Add Guest</button>
              <button style="margin-top: 32px; margin-left: 15px;"
                      class="button is-warning"
                      @click="setTBDGuest()">TBD</button>
              </p>
            </b-field>

          </b-field>

          <div v-if="formProps.guests.length >0">
            <ul>
              <li class="guest" v-for="(guest, index) in formProps.guests" :key="index">
                <b-icon icon="user"></b-icon>
                {{ guest.first_name }} {{ guest.last_name }} - {{ guest.email }}
                <b-button @click="deleteGuest(index)" type="is-danger" icon-right="trash-alt" size="is-small" />
              </li>
            </ul>
          </div>
        </div>

        <b-field label="Comments">
          <b-input maxlength="200" rows="2" v-model="formProps.comments" type="textarea"></b-input>
        </b-field>

      </section>
      <footer class="modal-card-foot">
        <button class="button is-success" :disabled="$v.$invalid" @click="createParty()">Create</button>
        <button class="button" @click="$parent.close()">Close</button>
      </footer>
    </div>
  </form>
</template>

<script>
  import { mapGetters } from "vuex";
  import { required, minLength, requiredUnless} from "vuelidate/lib/validators";

  export default {
    name: 'create_party',
    data() {
      return {
        formProps: {
          name: '',
          host_id: null,
          comments: '',
          guests: [],
          guest: {
            first_name: '',
            last_name: '',
            email: '',
          },
        },
        submitted: false,
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
        this.submitted = true;

        this.$v.$touch();
        if (this.$v.$invalid) {
          this.$v.$touch();
          return;
        }

        try {
          await this.$store.dispatch("party/createParty", this.formProps);
          this.$parent.close();
        } catch (error) {
          console.log(error);
        }
      },
      addGuest() {
        if (this.$v.formProps.guest.first_name.$invalid ||
          this.$v.formProps.guest.last_name.$invalid ||
          this.$v.formProps.guest.email.$invalid) {
          return;
        }
        this.formProps.guests.push({ ...this.formProps.guest });
        //this.formProps.guests.push(this.formProps.guest);
        this.formProps.guest.first_name = '';
        this.formProps.guest.last_name= '';
        this.formProps.guest.email = '';
      },
      deleteGuest(index) {
        this.formProps.guests.splice(this.formProps.guests.indexOf(index), 1);
      },
      setTBDGuest() {
        this.formProps.guest.first_name = 'TBD';
        this.formProps.guest.last_name= 'TBD';
        this.formProps.guest.email = 'not_known_yet@tbd.com';
      }
    },
    validations: {
      formProps: {
        name: { required },
        host_id: { required },
        guests: {
          required,
          minLength: minLength(1),
        },
        guest: {
          first_name: requiredUnless('hasGuests'),
          last_name: requiredUnless('hasGuests'),
          email: requiredUnless('hasGuests'),
        },
      }
    },
    computed: {
      ...mapGetters({
        users: "users/users",
      }),
      hasGuests() {
        if (this.formProps.guests.length > 0) {
          console.log("validation passed so chill?");
          return true;
        } else {
          return false;
        }
      }
    },
    created() {
      this.getHosts();
    },
  };
</script>

<style lang="scss">
.guests {
  background-color: #fffaff;
}

.guest {
  margin-top: 10px;
}

</style>
