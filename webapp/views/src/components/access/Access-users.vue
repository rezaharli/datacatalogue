<template>
    <v-content class="mx-4 my-5">
        <b-container fluid>
            <PageHeader title="Risk Framework Owner View" />

            <page-loader v-if="store.left.isLoading" />

            <b-row class="my-4">
                <b-col>
                    <b-button class="float-right red-neon icon-only ml-3" @click="resetFilter">
                        <i class="fa fa-fw fa-filter"></i>
                    </b-button>

                    <page-export class="float-right" :storeName="storeName" :leftTableCols="store.leftHeaders" :rightTableCols="[]"/>
                </b-col>
            </b-row>

            <b-row style="margin-top: 10px;margin-bottom: 10px;">
                <b-col>
                    <register-form />

                    <v-data-table
                            :headers="store.leftHeaders.filter(v => v.display == true)"
                            :items="store.left.display"
                            :pagination.sync="store.left.pagination"
                            :total-items="store.left.totalItems"
                            :loading="store.left.isLoading"
                            :expand="false"
                            :must-sort="true"
                            item-key="ID"
                            class="table-v1">
                        <template slot="headerCell" slot-scope="props">
                            <tableheader :storeName="storeName" :props="props" :which="'left'"/>
                        </template>

                        <v-progress-linear slot="progress" color="blue" indeterminate></v-progress-linear>

                        <template slot="no-data">
                            <v-alert
                                :value="store.left.isLoading"
                                type="info"
                            >Please wait while data is loading</v-alert>

                            <v-alert
                                :value="!store.left.isLoading"
                                type="error"
                            >Sorry, nothing to display here</v-alert>
                        </template>

                        <template slot="items" slot-scope="props">
                            <tr>
                                <td v-bind:style="{ width: store.left.colWidth['USERNAME'] + 'px' }" class="text-capitalize text-title">
                                    <tablecell showOn="hover" :fulltext="props.item.USERNAME"></tablecell></td>

                                <td v-bind:style="{ width: store.left.colWidth['EMAIL'] + 'px' }" class="text-capitalize text-title">
                                    <tablecell showOn="hover" :fulltext="props.item.EMAIL"></tablecell></td>

                                <td v-bind:style="{ width: store.left.colWidth['NAME'] + 'px' }" class="text-capitalize text-title">
                                    <tablecell showOn="hover" :fulltext="props.item.NAME"></tablecell></td>

                                <td v-bind:style="{ width: store.left.colWidth['ROLE'] + 'px' }" class="text-capitalize text-title">
                                    <tablecell showOn="hover" :fulltext="props.item.ROLE"></tablecell></td>

                                <td v-bind:style="{ width: store.left.colWidth['STATUS'] + 'px' }" class="text-capitalize text-title">
                                    <tablecell showOn="hover" :fulltext="props.item.STATUS"></tablecell></td>

                                <td v-bind:style="{ width: store.left.colWidth['CREATEDAT'] + 'px' }" class="text-capitalize text-title">
                                    <tablecell showOn="hover" :fulltext="props.item.CREATEDAT"></tablecell></td>

                                <td v-bind:style="{ width: store.left.colWidth['UPDATEAT'] + 'px' }" class="text-capitalize text-title">
                                    <tablecell showOn="hover" :fulltext="props.item.UPDATEAT"></tablecell></td>

                                <td v-bind:style="{ width: store.left.colWidth['ACTIONS'] + 'px' }" class="text-capitalize text-title">
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
        </b-container>
    </v-content>
</template>

<script>
import Vue from "vue";
import { mapState, mapActions } from "vuex";

import PageHeader from '../PageHeader';

import JsonExcel from "vue-json-excel";
import pageSearch from "../PageSearch.vue";
import pageExport from "../PageExport.vue";
import tableheader from "../TableHeader.vue";
import tablecell from "../Tablecell.vue";
import pageLoader from "../PageLoader.vue";

import registerForm from "./Register-form.vue";

Vue.component("downloadExcel", JsonExcel);

export default {
    components: {
        PageHeader, pageSearch, pageExport, tableheader, tablecell, pageLoader, registerForm
    },
    data () {
        return {
            storeName: "users",
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
            store: state => state.users.all
        }),
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
            this.editedIndex = this.store.items.indexOf(item)

            this.editedItem = Object.assign({}, item)
            this.editedItem.Role = this.editedItem.Role.split(",");
            this.editedItem.Status = this.editedItem.Status == 1 ? true : false;
            this.editedItem.Password = ''

            this.store.error = null
            this.dialog = true
        },
        deleteItem (item) {
            confirm('Are you sure you want to delete this user?') && 
                this.deleteUser(item.Username).then(
                    res => this.getAllUsers(), 
                    err => this.getAllUsers()
                )
        },
        reset () {
            this.$refs.form.reset()
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
        resetFilter (e) {
            if(Object.keys(this.store.filters.left).length > 0){
                this.store.filters.left = {};
                this.getLeftTable();
            }

            // if(Object.keys(this.store.filters.right).length > 0){
            //     this.store.filters.right = {}
            //     this.getMyRightTable(this.$route.params.system);
            // }
        },
    }
}
</script>
