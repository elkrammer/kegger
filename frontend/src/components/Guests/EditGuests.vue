<template>
    <form v-on:submit.prevent>
        <div class="modal-card" style="width: 500px; height: 900px;">
            <header class="modal-card-head">
                <p class="modal-card-title">Edit Guests</p>
            </header>
            <section class="modal-card-body">

                <div style="margin-bottom: 5px" class="buttons">
                    <b-button @click="addGuest" type="is-success">Add</b-button>
                </div>

                <b-collapse
                    class="card"
                    v-for="(guest, index) in guests"
                    :key="index"
                    :open="isOpen == index"
                    @open="isOpen = index"
                    >
                    <div
                        slot="trigger"
                        slot-scope="props"
                        class="card-header"
                        role="button">
                        <p class="card-header-title">
                            {{ guest.first_name }} {{ guest.last_name }}
                        </p>
                    <a class="card-header-icon">
                        <b-icon
                            :icon="props.open ? 'caret-up' : 'caret-down'">
                        </b-icon>
                    </a>
                    </div>
                    <div class="card-content">
                        <div class="content">
                            <b-field label="First Name">
                                <b-input placeholder="First Name" v-model="guest.first_name"></b-input>
                            </b-field>

                            <b-field label="Last Name">
                                <b-input placeholder="Last Name" v-model="guest.last_name"></b-input>
                            </b-field>

                            <b-field label="Email">
                                <b-input placeholder="Email" v-model="guest.email"></b-input>
                            </b-field>

                            <b-switch v-model="guest.plus_one" type="is-active">Plus One?</b-switch>

                            <div class="is-pulled-right">
                                <b-button @click="deleteGuest(index)" type="is-danger" icon-right="trash-alt" size="is-small" />
                            </div>
                        </div>
                    </div>
                </b-collapse>

            </section>
            <footer class="modal-card-foot">
                <button class="button is-success" @click="editGuests">Save</button>
                <button class="button" @click="$parent.close()">Close</button>
            </footer>
        </div>
    </form>
</template>

<script>
import { mapGetters } from "vuex";
export default {
    name: 'edit_guest',
    props: ['party_id'],
    data() {
        return {
            guest: [],
            isOpen: null,
        }
    },
    methods: {
        async getGuests() {
            try {
                const response = await this.$store.dispatch("guest/getGuests", this.party_id);
                return response.data;
            } catch (error) {
                console.log(error);
            }
        },
        async editGuests() {
            try {
                await this.$store.dispatch("guest/editGuests", this.guests);
                await this.$store.dispatch("party/getParties", null, { root: true });
                const msg = `Changes saved!`
                this.$buefy.toast.open({
                    message: msg,
                    type: 'is-success',
                    position: 'is-bottom',
                    duration: 3000,
                })
                this.$parent.close();
            } catch (error) {
                const msg = `There was an error editing guests: ${error}`
                this.$buefy.toast.open({
                    message: msg,
                    type: 'is-danger',
                    position: 'is-bottom',
                    duration: 3000,
                })
                console.log(error);
            }
        },
        addGuest() {
            const guest = {
                first_name:  '',
                last_name:    '',
                email:  '',
                party_refer: this.party_id,
            };
            this.guests.unshift(guest);
            this.isOpen = 0;
        },
        deleteGuest(index) {
            this.$delete(this.guests, index);
        },
    },
    computed: {
        ...mapGetters({
            guests: "guest/guests",
        })
    },
    getters: {
        guests(state) {
            return state.guests;
        }
    },
    created() {
        this.getGuests();
    },
}
</script>
