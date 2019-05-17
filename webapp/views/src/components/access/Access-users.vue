<style>
#table-access-users table.v-table tr th:nth-of-type(1){width: 10% !important;}
#table-access-users table.v-table tr th:nth-of-type(2){width: 12% !important;}
#table-access-users table.v-table tr th:nth-of-type(3){width: 12% !important;}
#table-access-users table.v-table tr th:nth-of-type(4){width: 12% !important;}
#table-access-users table.v-table tr th:nth-of-type(5){width: 10% !important;}
#table-access-users table.v-table tr th:nth-of-type(6){width: 10% !important;}
#table-access-users table.v-table tr th:nth-of-type(7){width: 10% !important;}
#table-access-users table.v-table tr th:nth-of-type(8){width: 5% !important;}
/* #table-access-users table.v-table.v-datatable thead{
    width: unset;
    display: table-header-group;
    padding-right: unset;
}
#table-access-users table.v-table.v-datatable tbody{
    display:table-row-group;
    overflow:auto;
    max-height:unset;
    width:unset;
} */
/* #table-access-users table.v-table.v-datatable tbody tr {display: table-row;} */
</style>

<template>
    <v-content class="mx-4 my-0 py-0">
        <b-container fluid>
            <page-loader v-if="store.left.isLoading" />

            <!-- <b-row class="my-4">
                <b-col>
                    <b-button class="float-right red-neon icon-only ml-3" @click="resetFilter">
                        <i class="fa fa-fw fa-filter"></i>
                    </b-button>

                    <page-export class="float-right" :storeName="storeName" :leftTableCols="store.leftHeaders" :rightTableCols="[]"/>
                </b-col>
            </b-row> -->

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
                            class="table-v1"
                            id="table-access-users">
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

                                <td v-bind:style="{ width: store.left.colWidth['UPDATEDAT'] + 'px' }" class="text-capitalize text-title">
                                    <tablecell showOn="hover" :fulltext="props.item.UPDATEDAT"></tablecell></td>

                                <td v-bind:style="{ width: store.left.colWidth['ACTIONS'] + 'px' }" class="text-capitalize text-title">
                                    <button @click="editItem(props.item)" v-if="isMyself(props.item) || amIAdmin()" type="button" class="btn btn-primary icon-only btn-md"><i class="fa fa-fw fa-pencil-alt"></i></button>

                                    <button @click="deleteItem(props.item)" v-if="( ! (isMyself(props.item))) && amIAdmin()" type="button" class="btn btn-danger icon-only btn-md"><i class="fa fa-fw fa-trash"></i></button>
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
    },
    created() {
        this.getAllUsers();
    },
    mounted() {
        setTimeout(() => {
        this.setTableColumnsWidth($('#table-access-users'));
        }, 300);
    },
    updated() {
        this.setTableColumnsWidth($('#table-access-users'));
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
        deleteUser(param) {
            return this.$store.dispatch(`users/delete`, param);
        },
        isMyself(item) {
            return JSON.parse(localStorage.getItem('user')).ID == item.ID;
        },
        amIAdmin() {
            return JSON.parse(localStorage.getItem('user')).Role.indexOf("Admin") != -1;
        },
        distinctData (col, datax) {
            var sorted = this._.sortBy(datax, col);

            return this._.uniq(
                this._.map(sorted, col));
        },
        editItem (item) {
            this.store.editedIndex = this.store.left.display.indexOf(item)

            this.store.editedItem = Object.assign({}, item)
            this.store.editedItem.ROLE = this.store.editedItem.ROLE.split(",");
            this.store.editedItem.STATUS = this.store.editedItem.STATUS == 1 ? true : false;
            this.store.editedItem.PASSWORD = ''

            this.store.error = null
            this.store.dialog = true
        },
        deleteItem (item) {
            confirm('Are you sure you want to delete this user?') && 
                this.deleteUser(item.USERNAME).then(
                    res => this.getAllUsers(), 
                    err => this.getAllUsers()
                )
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
        setTableColumnsWidth(elem){
            var tableElem = elem.find('.v-table__overflow > table.v-table');
            var THs = tableElem.find('thead tr th');
            var tbodyTR = tableElem.find('tbody tr');
            THs.each(function (thIndex) {
                var thWidth = $(this).width();
                tbodyTR.each(function (tdIndex) {
                var TDs = $(this).find('td:not([colspan])');
                TDs.eq(thIndex).width(thWidth);
                });
            });
        },
    }
}
</script>
