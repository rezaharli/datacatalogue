<style>
  @import '../../assets/styles/dashboard.css';
</style>

<style>
table.v-table thead th > div.btn-group {
    width: auto;
}
</style>

<template>
    <b-row>
        <b-col>
        <!-- Users details -->
            <router-view/>

            <!-- Main content -->
            <b-row>
                <b-col>
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
                                                <v-alert :value="users.error" color="error" icon="warning">
                                                    {{ users.error }}
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

                    <v-data-table
                            :headers="headers"
                            :items="users.items"
                            :loading="users.isLoading"
                            class="elevation-1">
                        <v-progress-linear slot="progress" color="blue" indeterminate></v-progress-linear>

                        <template slot="no-data">
                            <v-alert :value="!users.isLoading" color="error" icon="warning">
                                Sorry, nothing to display here :(
                            </v-alert>
                        </template>

                        <template slot="headers" slot-scope="props">
                            <tr>
                                <th v-for="header in props.headers" :key="header.text">
                                    {{ header.text }} ({{ distinctData(header.value, users.items).length }})

                                    <b-dropdown v-if="header.filter"  no-caret variant="link" class="header-filter-icon">
                                        <template slot="button-content">
                                            <!-- <i class="fa fa-filter text-muted"></i> -->
                                            <v-icon small>filter_list</v-icon>
                                        </template>

                                        <b-dropdown-header><b-form-input type="text" placeholder="Filter"></b-form-input></b-dropdown-header>

                                        <b-dropdown-item v-for="item in distinctData(header.value, users.items)" v-bind:key="item">{{ item }}</b-dropdown-item>
                                    </b-dropdown>
                                </th>
                            </tr>
                        </template>

                        <template slot="items" slot-scope="props">
                            <tr :active="props.selected" @click="props.selected = !props.selected">
                                <td>{{ props.item.Username }}</td>
                                <td>{{ props.item.Email }}</td>
                                <td>{{ props.item.Name }}</td>
                                <td>{{ props.item.Role }}</td>
                                <td>{{ props.item.Status }}</td>
                                <td>{{ props.item.CreatedAt }}</td>
                                <td>{{ props.item.UpdatedAt }}</td>
                                <td class="justify-center layout px-0">
                                    <!-- <v-icon
                                        small
                                        class="mr-2"
                                        @click="editItem(props.item)"
                                    >edit</v-icon>

                                    <v-icon
                                        small
                                        @click="deleteItem(props.item)"
                                    >delete</v-icon> -->
                                </td>
                            </tr>
                        </template>
                    </v-data-table>
                </b-col>
            </b-row>
        </b-col>
    </b-row>
</template>

<script>
import { mapState, mapActions } from 'vuex'

export default {
    data () {
        return {
            valid: true,
            rules: {
                required: value => !!value || 'Required.',
            },
            dialog: false,
            rolesMaster: ["Admin", "DSC", "DPO", "DDO", "RFO"],
            headers: [
                { text: 'Username', align: 'left', value: 'Username', sortable: false, filter: true },
                { text: 'Email', align: 'left', value: 'Email', sortable: false, filter: true },
                { text: 'Name', align: 'left', value: 'Name', sortable: false, filter: true },
                { text: 'Role', align: 'left', value: 'Role', sortable: false, filter: true },
                { text: 'Status', align: 'left', value: 'Status', sortable: false, filter: true },
                { text: 'Created At', align: 'left', value: 'CreatedAt', sortable: false, filter: true },
                { text: 'Updated At', align: 'left', value: 'UpdatedAt', sortable: false, filter: true },
                { text: 'Actions', value: 'name', sortable: false, filter: false }
            ],
            editedIndex: -1,
            editedItem: {},
            defaultItem: {
                Username: '',
                Password: '',
                Email: '',
                Name: '',
                Role: [],
                Status: true
            }
        }
    },
    computed: {
        ...mapState({
            users: state => state.users.all
        }),
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
    created() {
        this.getAllUsers();
    },
    methods: {
        ...mapActions('users', {
            getAllUsers: 'getAll',
            deleteUser: 'delete',
            updateUser: 'update',
            registerUser: 'register'
        }),
        distinctData (col, datax) {
            return this._.uniq(
                    this._.map(
                    this._.sortBy(datax, col), 
                    col
                    )
                );
        },
        editItem (item) {
            this.editedIndex = this.users.items.indexOf(item)

            this.editedItem = Object.assign({}, item)
            this.editedItem.Role = this.editedItem.Role.split(",");
            this.editedItem.Status = this.editedItem.Status == 1 ? true : false;
            this.editedItem.Password = ''

            this.users.error = null
            this.dialog = true
        },
        deleteItem (item) {
            confirm('Are you sure you want to delete this user?') && 
                this.deleteUser(item.Username).then(
                    res => this.getAllUsers(), 
                    err => this.getAllUsers()
                )
        },
        close () {
            this.dialog = false
            setTimeout(() => {
                this.editedItem = Object.assign({}, this.defaultItem)
                this.editedIndex = -1
                this.reset()
            }, 300)
        },
        reset () {
            this.$refs.form.reset()
        },
        save () {
            this.users.error = null;
            if (this.editedIndex > -1) { //edit
                this.updateUser(this.editedItem).then(res => {
                    if(!this.users.error) {
                        this.close()
                    }
                    this.getAllUsers()
                }, err => this.getAllUsers());
            } else { //add
                this.registerUser(this.editedItem).then(res => {
                    if(!this.users.error) {
                        this.close()
                    }
                    this.getAllUsers()
                }, err => {});
            }
        }
    }
}
</script>
