<style>
/* table.v-table thead th > div.btn-group {
  	width: auto;
}

.header-filter-icon .dropdown-menu {
	overflow: scroll;
	height: 200px;
} */
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
                                    {{ props.item.USERNAME }}</td>

                                <td v-bind:style="{ width: store.left.colWidth['FULLNAME'] + 'px' }" class="text-capitalize text-title">
                                    {{ props.item.FULLNAME }}</td>

                                <td v-bind:style="{ width: store.left.colWidth['ROLE'] + 'px' }" class="text-capitalize text-title">
                                    {{ props.item.ROLE }}</td>

                                <td v-bind:style="{ width: store.left.colWidth['MODULE'] + 'px' }" class="text-capitalize text-title">
                                    {{ props.item.MODULE }}</td>

                                <td v-bind:style="{ width: store.left.colWidth['DESCRIPTION'] + 'px' }" class="text-capitalize text-title">
                                    {{ props.item.DESCRIPTION }}</td>

                                <td v-bind:style="{ width: store.left.colWidth['TIME'] + 'px' }" class="text-capitalize text-title">
                                    {{ props.item.TIME }}</td>

                                <td v-bind:style="{ width: store.left.colWidth['RESOURCEURL'] + 'px' }" class="text-capitalize text-title">
                                    {{ props.item.RESOURCEURL }}</td>
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

Vue.component("downloadExcel", JsonExcel);

export default {
  components: {
    PageHeader, pageSearch, pageExport, tableheader, tablecell, pageLoader
  },
  data() {
    return {
      storeName: "usersusage",
      systemSource: [],
      tablenameSource: []
    };
  },
  computed: {
    store() {
      return this.$store.state[this.storeName].all;
    },
    addressPath() {
      var tmp = this.$route.path.split("/");
      return tmp.slice(0, 3).join("/");
    },
  },
  watch: {
    $route(to) {},
    "store.left.pagination": {
      handler() {
        this.getLeftTable();
      },
      deep: true
    },
    "store.searchMain"(val, oldVal) {
      if (val || oldVal) {
        this.getLeftTable();
      }
    }
  },
  mounted() {
    this.store.tabName = this.storeName;
    this.store.system = this.$route.params.system;
    this.resetFilter();
  },
  methods: {
    getLeftTable() {
      this.$store.dispatch(`${this.storeName}/getLeftTable`);
    },
    systemRowClick(evt) {
      evt.preventDefault();
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
    getCDEConclusion(cdes) {
      return cdes
        .filter(Boolean)
        .join(", ")
        .indexOf("Yes") != -1
        ? "Yes"
        : "No";
    },
    showCDEs(param) {
      this.$router.push(
        this.addressPath + "/" + param.SYSTEM_NAME + "/" + param.DSP_NAME
      );
    },
    showDetails(param) {
      this.$router.push(
        this.addressPath + "/" + param.TSID + "/" + param.ID + "/" + param.COLID
      );
    }
  }
};
</script>
