<template>
  <section class="section">
    <div class="container">
      <h1 class="title">Settings</h1>

      <b-tabs type="is-toggle">
        <b-tab-item label="Event" icon="calendar-alt">
          <Event />
        </b-tab-item>
        <b-tab-item label="Invite Design" icon="envelope">
          <Invite />
        </b-tab-item>
        <b-tab-item label="App Settings" icon="cog">
          <App />
        </b-tab-item>
        <b-tab-item label="Manage Users" icon="users">
          <Users />
        </b-tab-item>
        <b-tab-item label="Show Env Settings" icon="tree">
          <Environment />
        </b-tab-item>
      </b-tabs>

    </div>
  </section>
</template>

<script>
import Users from "@/components/Settings/Users.vue";
import Invite from "@/components/Settings/Invite.vue";
import Event from "@/components/Settings/Event.vue";
import App from "@/components/Settings/App.vue";
import Environment from "@/components/Settings/Environment.vue";

export default {
  name: "Settings",
  components: { Users, Invite, Event, App, Environment },
}
</script>
