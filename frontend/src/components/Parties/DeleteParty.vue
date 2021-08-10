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
    computed: {
      ...mapGetters({
        parties: "party/parties"
      })
    },
    methods: {
      async deleteParty() {
        try {
            await this.$store.dispatch("party/deleteParty", this.party_id);
            const msg = `Successfully deleted party ${this.party_name}`
            this.$buefy.toast.open({
                message: msg,
                type: 'is-success',
                position: 'is-bottom',
                duration: 3000,
            })
            this.$parent.close();
        } catch (error) {
            const msg = `There was an error deleting party ${this.party_name}`
            this.$buefy.toast.open({
                message: msg,
                type: 'is-danger',
                position: 'is-bottom',
                duration: 3000,
            })
            console.log(error);
        }
      },
    },
  }
</script>
