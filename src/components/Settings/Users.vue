<template>
  <section>

    <b-modal
      :active.sync="editUserActive"
      has-modal-card
      trap-focus
      aria-role="dialog"
      aria-modal>
      <EditUser :user_id="selected" />
    </b-modal>

    <div class="box">
      <h1 class="title">Manage Users</h1>
      <div class="buttons">
        <b-button type="is-success" style="margin-left: 10px;" icon-left="user-plus">Create User</b-button>
      </div>
      <p>

      <div v-for="user in users" :key="user.id" class="userlist">
        <ul>
          <li>

            <div class="level">

              <div class="level-left">
                <div class="level-item">
                  {{ user.name }}
                </div>
              </div>

              <div class="level-right">
                <div class="level-item">
                  <b-button @click="editUser(user.id)" type="is-info" icon-left="user-edit" size="is-small"></b-button>
                </div>
                <div class="level-item">
                  <b-button type="is-danger" icon-left="trash-alt" size="is-small"></b-button>
                </div>
              </div>

            </div>

          </li>
        </ul>
      </div>

    </div>
  </section>
</template>

<script>
  import { mapGetters } from "vuex";
  import EditUser from "@/components/Settings/EditUser.vue";

  export default {
    name: "settings_user",
    components: { EditUser },
    computed: {
      ...mapGetters({
        users: "users/users",
      })
    },
    data() {
      return {
        editUserActive: false,
        selected: null,
      }
    },
    methods: {
      async getUsers() {
        try {
          const response = await this.$store.dispatch("users/getUsers");
          return response.data;
        } catch (error) {
          console.log(error);
        }
      },
      editUser(id) {
        this.editUserActive = true;
        this.selected = id;
      }
    },
    created() {
      this.getUsers();
    }
  }
</script>

<style lang="scss">
@import "@/variables";

.userlist {
  width: 300px;
}

.userlist li {
  margin: 10px 0;
}

</style>
