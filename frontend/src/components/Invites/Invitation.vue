<template>
  <section class="section is-fullheight">
    <div class="container">

      <div v-if="Object.entries(this.invite).length <= 0">
        Invalid Invitation ID
      </div>

      <div class="columns is-multiline" v-if="Object.entries(this.invite).length > 0">
        <div class="column is-half">
          <figure class="image">
            <img :src="invite.invite_image" @click="isModalActive = true" class="is-shady">
          </figure>

          <b-modal :active.sync="isModalActive">
            <p class="image is-4by5">
            <img :src="invite.invite_image">
            </p>
          </b-modal>

        </div>

        <!-- English -->
        <div v-if="invite.invite_lang == 'en'" class="column is-half">
          <div class="box is-shady">
            <p>Hi {{ invite.guest.first_name }}, <br><br>

            Please join us to celebrate the wedding of {{ invite.bride_name }} &amp; {{ invite.groom_name }}.
            <br><br>

            The wedding will be on {{ invite.event_date | moment('timezone', invite.time_zone, 'MMM do YYYY H:mm') }} at {{ invite.event_location }}.
            <br>
            The dress code will be {{ invite.dress_code }}.

            <br><br>

            <span v-if="invite.guest.plus_one">
              You are allowed to bring a guest to the wedding.
            </span>
            <br>
            <br>

            Love,<br/>
            {{ invite.bride_name }} &amp; {{ invite.groom_name }}<br><br>
            <img :src="invite.signature_image">
            </p>
          </div>

          <div class="box is-shady">
            <h1 class="title">RSVP</h1>
            <span>Are you planning to attend our wedding?</span>
            <br><br>

            <div class="rsvp">
              <b-select icon="calendar-check" v-model="attending" rounded placeholder="Select an option">
                <option :value="true">Yes</option>
                <option :value="false">No</option>
              </b-select>
              <br>

              <b-button @click="setAttending" rounded class="confirm-btn" type="is-success">
                Confirm
              </b-button>
            </div>

          </div>
        </div>


        <!-- Spanish -->
        <div v-if="invite.invite_lang == 'es'" class="column is-half">
          <div class="box is-shady">
            <p>Hola {{ invite.guest.first_name }}, <br><br>
            El matrimonio será el día {{ invite.event_date | moment('timezone', invite.time_zone, 'MMM do YYYY H:mm') }} en {{ invite.event_location }}.
            <br><br>

            La vestimenta para el evento será {{ invite.dress_code }}.

            <br><br>

            <span v-if="invite.guest.plus_one">
              Puedes traer a un invitado al matrimonio.
            </span>
            <br>
            <br>

            Love,<br><br>
            {{ invite.bride_name }} &amp; {{ invite.groom_name }}<br><br>
            <img :src="invite.signature_image">
            </p>
          </div>

          <div class="box is-shady">
            <h1 class="title">RSVP</h1>
            <span>Estas planeando venir a nuestro matrimonio?</span>
            <br><br>

            <div class="rsvp">
              <b-select icon="calendar-check" v-model="attending" rounded placeholder="Select an option">
                <option :value="true">Si</option>
                <option :value="false">No</option>
              </b-select>
              <br>

              <b-button @click="setAttending" rounded class="confirm-btn" type="is-success">
                Confirmar
              </b-button>
            </div>

          </div>

        </div>
      </div>
    </div>

  </section>
</template>


<script>
import { mapGetters } from "vuex";

export default {
  name: "invitation",
  props: ['id'],
  data() {
    return {
      attending: null,
      isModalActive: false,
    }
  },
  computed: {
    ...mapGetters({
      invite: "invite/invite",
    }),
  },
  methods: {
    async getInvite() {
      try {
        await this.$store.dispatch("invite/getInvite", this.id);
        // update invitation_opened if this is the first time it's being opened
        if (this.invite && this.invite.guest.invitation_opened == null) {
          this.setInviteOpened();
        }
        // sync local attending state with the invite
        this.attending = this.invite.guest.is_attending;
      } catch (error) {
        console.log(error);
      }
    },
    async setInviteOpened() {
      try {
        const data = {
          action: "opened",
          invitation_id: this.id
        };
        await this.$store.dispatch("invite/updateInvite", data);
      } catch (error) {
        console.log(error);
      }
    },
    async setAttending() {
      try {
        const data = {
          action: "attending",
          invitation_id: this.id,
          is_attending: this.attending
        };
        await this.$store.dispatch("invite/updateInvite", data);

        const msg = `Successfully sent your RSVP`
        this.$buefy.toast.open({
          message: msg,
          type: 'is-success',
          position: 'is-bottom',
          duration: 3000,
        })
      } catch (error) {
        const msg = `There was an error updating the invite`
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
  created() {
    this.getInvite();
  },
}
</script>

<style lang="scss" scoped>

html, body {
  height: 100%;
}

.section {
  height: 100vh;
  background: #A1FFCE;  /* fallback for old browsers */
  background: -webkit-linear-gradient(to bottom, #A1FFCE, #FAFFD1);  /* Chrome 10-25, Safari 5.1-6 */
  background: linear-gradient(to bottom, #A1FFCE, #FAFFD1); /* W3C, IE 10+/ Edge, Firefox 16+, Chrome 26+, Opera 12+, Safari 7+ */
}

.invite {
  min-height: 101vh !important;
}

.rsvp {
  display: table;
  margin: 0 auto;
}

</style>
