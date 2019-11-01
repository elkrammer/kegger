<template>
  <section class="section">
    <div class="container">
      <h1 class="title">Parties</h1>
      <div v-for="party in parties" :key="party.id">
        <ul style="list-style-type:disc;">
          <li> {{ party.name }} </li>
        </ul>
      </div>
    </div>
  </section>
</template>

<script>
  import { mapGetters } from "vuex";

  export default {
    name: 'listparties',
    data() {
      return {

      };
    },
    methods: {
      async loadParties() {
        try {
          const response = await this.$store.dispatch("party/getParties");
          console.log(response);
          return response.data;
        } catch (error) {
          console.log(error);
        }
      },
    },
    computed: {
      ...mapGetters({
        parties: "party/parties"
      })
    },
    created() {
      this.loadParties();
    }
  };
</script>

<style lang="scss">
</style>
