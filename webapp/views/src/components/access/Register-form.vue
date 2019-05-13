<template>
    <v-dialog v-model="dialog" max-width="500px">
        <v-btn slot="activator" color="primary" dark class="mb-2">New Item</v-btn>
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
                                <v-text-field type="number" :readonly="editedIndex > -1" :rules="[rules.required]" v-model="editedItem.Username" label="Username" ></v-text-field>
                            </v-flex>

                            <v-flex xs12 sm6 md4>
                                <v-text-field type="email" :rules="[rules.required]" v-model="editedItem.Email" label="Email"></v-text-field>
                            </v-flex>

                            <v-flex xs12 sm6 md4>
                                <v-text-field type="password" v-model="editedItem.Password" label="Password"></v-text-field>
                            </v-flex>

                            <v-flex xs12 sm6 md4>
                                <v-text-field :rules="[rules.required]" v-model="editedItem.Name" label="Name"></v-text-field>
                            </v-flex>

                            <v-flex xs12 sm6 md4>
                                <v-select
                                    v-model="editedItem.Role"
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
                                <v-switch :label="`Status`" v-model="editedItem.Status"></v-switch>
                                <!-- <v-text-field v-model="editedItem.Status" label="Status"></v-text-field> -->
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
            dialog: false,
            valid: true,
            editedIndex: -1,
            rules: {
                required: value => !!value || 'Required.',
            },
            editedItem: {},
            rolesMaster: ["Admin", "DSC", "DPO", "DDO", "RFO"],
        };
    },
    computed: {
        store () { return this.$store.state.users.all },
        formTitle () {
            return this.editedIndex === -1 ? 'New Item' : 'Edit Item'
        },
        formIsValid () {
            return (
                this.editedItem.Username &&
                this.editedItem.Email &&
                this.editedItem.Name &&
                this.editedItem.Role
            )
        }
    },
    methods: {
        close () {
            this.dialog = false
            setTimeout(() => {
                this.editedItem = Object.assign({}, this.defaultItem)
                this.editedIndex = -1
                this.reset()
            }, 300)
        },
        save () {
            this.store.error = null;
            if (this.editedIndex > -1) { //edit
                this.updateUser(this.editedItem).then(res => {
                    if(!this.users.error) {
                        this.close()
                    }
                    this.getAllUsers()
                }, err => this.getAllUsers());
            } else { //add
                this.registerUser(this.editedItem).then(res => {
                    if(!this.store.error) {
                        this.close()
                    }
                    this.getAllUsers()
                }, err => {});
            }
        },
    }
};
</script>