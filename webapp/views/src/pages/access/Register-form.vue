<template>
    <v-dialog v-model="store.dialog" max-width="500px">
        <!-- <v-btn slot="activator" color="primary" dark class="mb-2">New Item</v-btn> -->
        <button slot="activator" type="button" class="btn green-tosca-gradient icon-only btn-secondary btn-md mb-4"><i class="fa fa-fw fa-plus-circle"></i> &nbsp; New Item</button>
        <v-card>
            <v-form
                ref="form"
                v-model="valid"
                >
                <v-card-title>
                    <span class="headline">{{ formTitle }}</span>
                </v-card-title>

                <v-card-text>
                    <v-container grid-list-md>
                        <v-layout wrap>
                            <v-flex xs12 sm12 md12>
                                <v-alert :value="store.error" color="error" icon="warning">
                                    {{ store.error }}
                                </v-alert>
                            </v-flex>

                            <v-flex xs12 sm6 md4>
                                <v-text-field type="number" :readonly="store.editedIndex > -1" :rules="[rules.required]" v-model="store.editedItem.USERNAME" label="Username" ></v-text-field>
                            </v-flex>

                            <v-flex xs12 sm6 md4>
                                <v-text-field type="email" :rules="[rules.required]" v-model="store.editedItem.EMAIL" label="Email"></v-text-field>
                            </v-flex>

                            <v-flex xs12 sm6 md4>
                                <v-text-field type="password" v-model="store.editedItem.PASSWORD" label="Password"></v-text-field>
                            </v-flex>

                            <v-flex xs12 sm6 md4>
                                <v-text-field :rules="[rules.required]" v-model="store.editedItem.NAME" label="Name"></v-text-field>
                            </v-flex>

                            <v-flex xs12 sm6 md4>
                                <v-select
                                    v-model="store.editedItem.ROLE"
                                    :items="rolesMaster"
                                    :rules="[rules.required]" 
                                    label="Select"
                                    multiple
                                    chips
                                    hint="What are the target regions"
                                    persistent-hint
                                ></v-select>
                            </v-flex>

                            <v-flex xs12 sm6 md4>
                                <v-switch :label="`Status`" v-model="store.editedItem.STATUS"></v-switch>
                                <!-- <v-text-field v-model="store.editedItem.Status" label="Status"></v-text-field> -->
                            </v-flex>
                        </v-layout>
                    </v-container>
                </v-card-text>

                <v-card-actions>
                    <v-spacer></v-spacer>
                    <v-btn color="blue darken-1" flat @click="close">Cancel</v-btn>
                    <v-btn color="blue darken-1" flat @click="save" :disabled="!formIsValid">Save</v-btn>
                </v-card-actions>
            </v-form>
        </v-card>
    </v-dialog>
</template>

<script>
export default {
    name: "registerForm",
    props: [],
    data() {
        return {
            valid: true,
            rules: {
                required: value => !!value || 'Required.',
            },
            rolesMaster: ["Admin", "DSC", "DPO", "DDO", "RFO"],
        };
    },
    computed: {
        store () { return this.$store.state.users.all },
        formTitle () {
            return this.store.editedIndex === -1 ? 'New Item' : 'Edit Item'
        },
        formIsValid () {
            return (
                this.store.editedItem.USERNAME &&
                this.store.editedItem.EMAIL &&
                this.store.editedItem.NAME &&
                this.store.editedItem.ROLE
            )
        }
    },
    methods: {
        getAllUsers() {
            return this.$store.dispatch(`users/getLeftTable`);
        },
        registerUser(param) {
            return this.$store.dispatch(`users/register`, param);
        },
        updateUser(param) {
            return this.$store.dispatch(`users/update`, param);
        },
        close () {
            this.store.dialog = false
            setTimeout(() => {
                this.store.editedItem = Object.assign({}, this.defaultItem)
                this.store.editedIndex = -1
                this.reset()
            }, 300)
        },
        save () {
            this.store.error = null;
            if (this.store.editedIndex > -1) { //edit
                this.updateUser(this.store.editedItem).then(res => {
                    if(!this.store.error) {
                        this.close()
                    }
                    this.getAllUsers()
                }, err => this.getAllUsers());
            } else { //add
                this.registerUser(this.store.editedItem).then(res => {
                    if(!this.store.error) {
                        this.close()
                    }
                    this.getAllUsers()
                }, err => {});
            }
        },
        reset () {
            this.$refs.form.reset()
        },
    }
};
</script>