<template>
  <section class="section is-fullheight">
    <div class="container">

      <div v-if="Object.entries(this.invite).length <= 0">
        Invalid Invitation ID
      </div>

      <div class="columns is-multiline" v-if="Object.entries(this.invite).length > 0">
        <div class="column is-half">
          <figure class="image">
            <img :src="invite.invite_image" @click="isModalActive = true" class="is-shady" />
          </figure>

          <b-modal :active.sync="isModalActive">
            <p class="image is-4by5">
            <img :src="invite.invite_image" />
            </p>
          </b-modal>

        </div>

        <!-- English -->
        <div v-if="invite.invite_lang == 'en'" class="column is-half">
          <div class="box is-shady">
            <p>Hi {{ invite.guest.first_name }}, <br/><br/>

            The wedding will be on {{ invite.event_date }} at {{ invite.event_location }}.
            <br/><br/>

            For more information please visit our website at: <a :href="invite.wedding_website" alt="Wedding Website">{{ invite.wedding_website }}</a>

            <br/><br/>

            Love,<br/>
            {{ invite.bride_name }} &amp; {{ invite.groom_name }}<br/><br/>
            <img class="signature-img" :src="invite.signature_image" />
            </p>
          </div>

          <div class="box is-shady" v-if="!this.isModal && !this.confirmed">
            <h1 class="title">RSVP</h1>
            <span>Are you planning to attend our wedding?</span>
            <br/><br/>

            <div class="rsvp">
              <b-select icon="calendar-check" v-model="attending" rounded placeholder="Select an option">
                <option :value="true">Yes</option>
                <option :value="false">No</option>
              </b-select>
              <br/>

              <div v-if="this.invite.guest.plus_one">
                <span>Plus One?</span>
                <br/><br/>
                <b-select icon="plus" v-model="plus_one" rounded placeholder="Select an option">
                  <option :value="true">Yes</option>
                  <option :value="false">No</option>
                </b-select>
              </div>

              <br/>
              <b-button @click="saveChanges" rounded class="confirm-btn" type="is-success">
                Confirm
              </b-button>
            </div>
          </div>

          <div class="box is-shady" v-if="this.confirmed">
            Thanks for your confirmation. <br/>
            See you on {{ invite.event_date }} <br/>
          </div>

        </div>


        <!-- Spanish -->
        <div v-if="invite.invite_lang == 'es'" class="column is-half">
          <div class="box is-shady">
            <p>Hola {{ invite.guest.first_name }}, <br/><br/>
            El matrimonio será el día {{ invite.event_date }} en {{ invite.event_location }}.
            <br/><br/>

            Para más información revisa nuestra página web: <a :href="invite.wedding_website" alt="Wedding Website">{{ invite.wedding_website }}</a>

            <br/><br/>

            Love,<br/>
            {{ invite.bride_name }} &amp; {{ invite.groom_name }}<br/><br/>
            <img class="signature-img" :src="invite.signature_image" />
            </p>
          </div>

          <div class="box is-shady" v-if="!this.isModal && !this.confirmed">
            <h1 class="title">RSVP</h1>
            <span>Estas planeando venir a nuestro matrimonio?</span>
            <br/><br/>

            <b-select icon="calendar-check" v-model="attending" rounded placeholder="Select an option">
              <option :value="true">Si</option>
              <option :value="false">No</option>
            </b-select>
            <br/>

            <div v-if="this.invite.guest.plus_one">
              <span>Vendrás con alguien?</span>
              <br/><br/>

              <b-select icon="plus" v-model="plus_one" rounded placeholder="Select an option">
                <option :value="true">Si</option>
                <option :value="false">No</option>
              </b-select>
              <br/>
            </div>

            <br/>
            <div class="has-text-centered">
              <b-button @click="saveChanges" rounded class="confirm-btn" type="is-success">
                Confirmar
              </b-button>
            </div>
          </div>

          <div class="box is-shady" v-if="this.confirmed">
            Gracias por confirmar. <br/>
            Nos vemos el {{ invite.event_date }} <br/>
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
  props: ['id', 'isModal'],
  data() {
    return {
      attending: null,
      plus_one: null,
      isModalActive: false,
      confirmed: false,
    }
  },
  computed: {
    ...mapGetters({
      invite: "invite/invite",
      user: "user/user",
    }),
  },
  created() {
    this.getInvite();
  },
  methods: {
    async getInvite() {
      try {
        if (this.user && this.user.username && this.user.sub && this.id.toString().length < 10 && !isNaN(parseInt(this.id))) {
          await this.$store.dispatch("invite/getInviteProtected", this.id);
        } else if (this.id.toString().length > 20) {
          await this.$store.dispatch("invite/getInvite", this.id);
          // update invitation_opened if this is the first time it's being opened
          if (this.invite && this.id.length > 10 && this.invite.guest.invitation_opened == null) {
            this.setInviteOpened();
          }
          // sync local attending state with the invite
          this.attending = this.invite.guest.is_attending;
          this.plus_one = this.invite.guest.plus_one;
        }
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
    async saveChanges() {
      try {
        const data = {
          action: "attending",
          invitation_id: this.id,
          is_attending: this.attending,
          plus_one: this.plus_one
        };
        await this.$store.dispatch("invite/updateInvite", data);
        this.confirmed = true;

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

.signature-img {
  height: 50px;
}

</style>
