<template>
  <form v-on:submit.prevent>
    <div class="modal-card" style="width: auto">
      <header class="modal-card-head">
        <p class="modal-card-title">Delete Party</p>
      </header>
      <section class="modal-card-body">
        <p>Are you sure that you want to delete Party {{ party_name }}?</p>
      </section>
      <footer class="modal-card-foot">
        <button class="button is-danger" @click="deleteParty">Delete</button>
        <button class="button" @click="$parent.close()">Close</button>
      </footer>
    </div>
  </form>
</template>

<script>
  import { mapGetters } from "vuex";
  export default {
    name: 'delete_party',
    props: ['party_id', 'party_name'],
    methods: {
      async deleteParty() {
        try {
          await this.$store.dispatch("party/deleteParty", this.party_id);
          this.$parent.close();
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
  }
</script>
