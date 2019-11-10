<template>
  <form action="">
    <div class="modal-card" style="width: auto">
      <header class="modal-card-head">
        <p class="modal-card-title">Create New Party</p>
      </header>
      <section class="modal-card-body">
        <b-field label="Party Name">
          <b-input
            placeholder="Party Name"
            required>
          </b-input>
        </b-field>

        <b-field label="Host">
          <b-select placeholder="Select Host">
            <option
              v-for="user in users"
              :value="user.name"
              :key="user.id">
            {{ user.name }}
            </option>
          </b-select>
        </b-field>

      </section>
      <footer class="modal-card-foot">
        <button class="button is-success">Create</button>
        <button class="button" type="button" @click="$parent.close()">Close</button>
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
          partyName: '',
        }
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
    },
    computed: {
      ...mapGetters({
        users: "users/users"
      })
    },
    created() {
      this.getHosts();
    },
  }
</script>
