<template>
  <section class="section">
    <div class="container">
      <h1 class="title">Settings</h1>

        <b-tabs type="is-toggle">
            <b-tab-item label="Manage Users" icon="users">
              <Users />
            </b-tab-item>
            <b-tab-item label="Foo" icon="dice">
              foo bar
            </b-tab-item>
        </b-tabs>

    </div>
  </section>
</template>

<script>
  import Users from "@/components/Settings/Users.vue";

  export default {
    name: "login",
    components: { Users },
  }
</script>
