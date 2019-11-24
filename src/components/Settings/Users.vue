<template>
  <section>

    <b-modal
      :active.sync="editUserActive"
      has-modal-card
      trap-focus
      aria-role="dialog"
      aria-modal>
      <EditUser :user_id="selected.id" />
    </b-modal>

    <b-modal
      :active.sync="createUserActive"
      has-modal-card
      trap-focus
      aria-role="dialog"
      aria-modal>
      <CreateUser/>
    </b-modal>

    <b-modal
      :active.sync="deleteUserActive"
      has-modal-card
      trap-focus
      aria-role="dialog"
      aria-modal>
      <DeleteUser :user_id="selected.id" :user_name="selected.name" />
    </b-modal>

    <div class="box">
      <h1 class="title">Manage Users</h1>
      <div class="buttons">
        <b-button @click="createUserActive = true" type="is-success" style="margin-left: 10px;" icon-left="user-plus">Create User</b-button>
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
                  <b-button @click="deleteUser(user.id, user.name)" type="is-danger" icon-left="trash-alt" size="is-small"></b-button>
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
  import CreateUser from "@/components/Settings/CreateUser.vue";
  import EditUser from "@/components/Settings/EditUser.vue";
  import DeleteUser from "@/components/Settings/DeleteUser.vue";

  export default {
    name: "settings_user",
    components: { CreateUser, EditUser, DeleteUser, },
    computed: {
      ...mapGetters({
        users: "users/users",
      })
    },
    data() {
      return {
        createUserActive: false,
        editUserActive: false,
        deleteUserActive: false,
        selected: {
          id: null,
          name: '',
        },
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
        this.selected.id = id;
        this.editUserActive = true;
      },
      deleteUser(id, name) {
        this.selected.id = id;
        this.selected.name = name
        this.deleteUserActive = true;
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
