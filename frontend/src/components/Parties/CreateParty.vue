<template>
  <form @submit.prevent="isFormValid" novalidate>
    <div class="modal-card">
      <header class="modal-card-head">
        <p class="modal-card-title">Create New Party</p>
      </header>
      <section class="modal-card-body">

        <div class="field is-grouped">
          <b-field label="Party Name">
            <b-input
              placeholder="Party Name"
              v-model="formProps.name"
              :use-html5-validation="false"
              required />
          </b-field>

          <b-field label="Host">
            <b-select placeholder="Select Host" v-model="selectedHost">
              <option
                v-for="user in users"
                :value="user.id"
                :key="user.id"
                :v-bind:value="user.id">
              {{ user.name }}
              </option>
            </b-select>
          </b-field>
        </div>

        <div class="columns">
          <div class="column guests is-two-fifths">
            <b-field label="Guests" />

            <b-field label="First Name">
              <b-input
                placeholder="First Name"
                v-model="formProps.guest.first_name"
                :use-html5-validation="false"
                required />
            </b-field>

            <b-field label="Last Name">
              <b-input
                placeholder="Last Name"
                v-model="formProps.guest.last_name"
                :use-html5-validation="false"
                required />
            </b-field>

            <b-field label="Email">
              <b-input
                type="email"
                placeholder="e.g. john@doe.com"
                v-model.trim="formProps.guest.email"
                :use-html5-validation="false"
                required />
            </b-field>

            <nav class="level">
              <div class="level-left">
                <div class="level-item">
                  <b-field label="Plus One">
                    <b-select v-model="formProps.guest.plus_one" placeholder="Plus One">
                      <option :value="false">No</option>
                      <option :value="true">Yes</option>
                    </b-select>
                  </b-field>
                </div>
              </div>

              <div class="level-right">
                <div class="level-item">
                  <b-field label="Invite Language">
                    <b-select v-model="formProps.guest.invitation_lang" placeholder="Invite Language">
                      <option value="es">Espanol</option>
                      <option value="en">English</option>
                    </b-select>
                  </b-field>
                </div>
              </div>
            </nav>

            <b-field>
              <p class="buttons">
              <button class="button is-success" :disabled="!isGuestFormValid" @click="addGuest()">Add Guest</button>
              </p>
            </b-field>
          </div>

          <div class="column">
            <div v-if="formProps.guests.length >0" class="guestlist">
              <h3>Party Guests:</h3>
              <ul>
                <li class="guest" v-for="(guest, index) in formProps.guests" :key="index">
                  <b-icon icon="user" />
                  {{ guest.first_name }} {{ guest.last_name }}

                  <span v-if="guest.plus_one" class="has-text-success">
                    (+1)
                  </span>

                  <b-button @click="deleteGuest(index)" style="margin-left: 8px" type="is-danger" icon-right="trash-alt" size="is-small" />
                </li>
              </ul>
            </div>
          </div>
        </div>

        <b-field label="Comments">
          <b-input maxlength="200" rows="2" v-model="formProps.comments" type="textarea" />
        </b-field>

        <footer class="modal-card-foot">
          <button class="button is-success" :disabled="!isFormValid()" @click="createParty()">Create</button>
          <button class="button" @click="closeModal()">Close</button>
        </footer>
      </section>
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
        guest: {
          first_name: '',
          last_name: '',
          email: '',
          plus_one: false,
          invitation_lang: 'es',
        },
      },
      selectedHost: null,
    }
  },
  computed: {
    ...mapGetters({
      users: "users/users",
      user: "user/user",
    }),
    isGuestFormValid () {
      if (
        this.formProps.guest.first_name !== '' &&
        this.formProps.guest.last_name !== '' &&
        this.isValidEmail(this.formProps.guest.email)
      ) {
        return true;
      }
      return false;
    },
  },
  watch: {
    selectedHost() {
      this.formProps.host_id = this.selectedHost;
    },
  },
  created() {
    this.getHosts();
    this.selectedHost = this.user.sub;
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
      if (!this.isFormValid) {
        return;
      }
      try {
        await this.$store.dispatch("party/createParty", this.formProps);
        // close modal dialog
        this.closeModal();
        const msg = `Successfully created party ${this.formProps.name}`
        this.$buefy.toast.open({
          message: msg,
          type: 'is-success',
          position: 'is-bottom',
          duration: 3000,
        })
      } catch (error) {
        console.log(error);
        const msg = `There was an error creating party ${this.formProps.name}`
        this.$buefy.toast.open({
          message: msg,
          type: 'is-danger',
          position: 'is-bottom',
          duration: 3000,
        })
      }
    },
    addGuest() {
      this.formProps.guests.push({ ...this.formProps.guest });
      this.formProps.guest.first_name = '';
      this.formProps.guest.last_name= '';
      this.formProps.guest.email = '';
      this.formProps.guest.plus_one = false;
      this.formProps.guest.invitation_lang = 'es';
    },
    deleteGuest(index) {
      this.formProps.guests.splice(this.formProps.guests.indexOf(index), 1);
    },
    isFormValid() {
      if (
        this.formProps.guests.length > 0 &&
        this.formProps.name !== '' &&
        this.formProps.host_id
      ) {
        return true;
      }
      return false;
    },
    closeModal() {
      this.$parent.close();
    },
    isValidEmail(email) {
      var re = /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
      return re.test(email);
    }
  },
};
</script>

<style lang="scss">

.guests {
  background-color: #fffaff;
}

.guestlist {
  background-color: #fffaaf;
}

.guest {
  margin-top: 10px;
}

</style>
