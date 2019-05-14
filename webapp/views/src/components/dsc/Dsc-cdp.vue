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
    <v-content class="mx-4 my-5">
        <b-container fluid>
            <PageHeader />

            <page-loader v-if="store.left.isLoading" />
            
            <b-row>
                <b-col sm=12 md=3>
                    <div class="card card-v2 transition">
                        <h6 class="title-1">System Name</h6>
                        <h3 class="title-2 text-capitalize">{{$route.params.system}}</h3>
                    </div>
                </b-col>
                <b-col sm=12 md=3>
                    <div class="card card-v2 transition">
                        <h6 class="title-1">Downstream Processes</h6>
                        <h3 class="title-2 text-capitalize">{{ store.left.source[0] ? store.left.source[0]["COUNT_DSP_NAME"] : "0" }}</h3>
                    </div>
                </b-col>
            </b-row>

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
                <!-- Dsc details -->
                <router-view/>

                <!-- Main content -->
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
                    <tr :class="{even: props.index % 2, odd: !(props.index % 2)}">
                      <td style="width: calc(100% / 6)" class="text-uppercase">{{props.item.DSP_NAME}}</td>
                      <td style="width: calc(100% / 6)" class="text-capitalize"><tablecell :fulltext="props.item.PROCESS_OWNER" showOn="click"></tablecell></td>
                      <td style="width: calc(100% / 6)" class="text-capitalize"><b-link @click.stop="showCDEs(props.item)"><tablecell :fulltext="props.item.CDE_COUNT" showOn="hover"></tablecell></b-link></td>
                    </tr>
                  </template>

                  <template slot="footer" >
                    <td :colspan="2">
                      <b>Total</b>
                    </td>
                    <td><b>{{ store.left.display[0] ? store.left.display[0].TOTAL : 0 }}</b></td>
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
      storeName: "dsccdp",
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
