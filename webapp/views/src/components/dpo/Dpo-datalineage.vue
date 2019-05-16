<style> 
#table-dpo-datalineage table.v-table tr th:nth-of-type(1){width: 10% !important;}
#table-dpo-datalineage table.v-table tr th:nth-of-type(2){width: 10% !important;}
</style>

<template>
    <v-content class="mx-4 my-5">
        <b-container fluid>
            <PageHeader title="Data Flow / Lineage View (Upstream)" />

            <page-loader v-if="store.left.isLoading" />
            
            <b-row>
                <b-col sm=12 md=3>
                    <div class="card card-v2 transition">
                        <h6 class="title-1">Downstream Process Name</h6>
                        <h3 class="title-2 text-capitalize">{{ store.left.source[0] ? store.left.source[0]["PROCESS_NAME"] : " " }}</h3>
                    </div>
                </b-col>

                <b-col sm=12 md=3>
                    <div class="card card-v2 transition">
                        <h6 class="title-1">Upstream Data Providers</h6>
                        <h3 class="title-2 text-capitalize">{{ store.left.source[0] ? store.left.source[0]["UPSTREAM_COUNT"] : "0" }}</h3>
                    </div>
                </b-col>
            </b-row>

            <b-row style="margin-top: 10px; margin-bottom: 20px;">
                <b-col>
                    <b-button class="float-right red-neon icon-only ml-3" @click="resetFilter">
                        <i class="fa fa-fw fa-filter"></i>
                    </b-button>

                    <page-export class="float-right" :storeName="storeName" :leftTableCols="store.leftHeaders" :rightTableCols="[]"/>

                    <b-button class="float-right primary icon-only mr-3" @click="linkDataLineage">
                        <i class="fa fa-fw fa-external-link-alt"></i> Go to data lineage
                    </b-button>
                </b-col>
            </b-row>

            <b-row style="margin-top: 10px;margin-bottom: 10px;">
              <b-col>
                <!-- Ddo details -->
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
                    class="table-v1"
                    id="table-dpo-datalineage">
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

                  <!-- <template slot="items" slot-scope="props">
                    <tr :class="{even: props.index % 2, odd: !(props.index % 2)}">
                      <td v-bind:style="{ width: store.left.colWidth['Details'] + 'px' }" class="text-capitalize text-title">
                        <b-button size="sm" class="green-tosca-gradient icon-only" @click="showDetails(props.item)">
                          <i class="fa fa-fw fa-external-link-alt"></i></b-button></td>

                      <td v-bind:style="{ width: store.left.colWidth['BT_NAME'] + 'px' }" class="text-uppercase">{{props.item.BT_NAME}}</td>
                      <td v-bind:style="{ width: store.left.colWidth['TABLE_NAME'] + 'px' }" class="text-capitalize"><tablecell :fulltext="props.item.TABLE_NAME" showOn="click"></tablecell></td>
                      <td v-bind:style="{ width: store.left.colWidth['COLUMN_NAME'] + 'px' }" class="text-capitalize"><tablecell :fulltext="props.item.COLUMN_NAME" showOn="click"></tablecell></td>
                    </tr>
                  </template> -->

                  <template slot="items" slot-scope="props">
                    <tr :class="{even: props.index % 2, odd: !(props.index % 2)}">
                      <td v-if="displayPRName" v-bind:style="{ width: store.left.colWidth['PR_NAME'] + 'px' }" class="text-uppercase">
                        <tablecell showOn="hover" :fulltext="props.item.PR_NAME"></tablecell></td>

                      <td v-bind:style="{ width: store.left.colWidth['SYSTEM_NAME'] + 'px' }" class="text-uppercase">
                        <tablecell showOn="hover" :fulltext="props.item.SYSTEM_NAME"></tablecell></td>

                      <td v-bind:style="{ width: store.left.colWidth['CDE_COUNT'] + 'px' }" class="text-uppercase">
                        <tablecell showOn="hover" :fulltext="props.item.CDE_COUNT"></tablecell></td>
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
      storeName: "dpodatalineage",
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
    displayPRName() {
      var prNameHeader = this.store.leftHeaders.find(v => v.value == "PR_NAME");

      if(parseInt(this.store.left.source[0] ? this.store.left.source[0]["PR_NAME_COUNT"] : "0") == 0){
        prNameHeader.display = false;
        prNameHeader.exportable = false;
        return false;
      } else {
        prNameHeader.display = true;
        prNameHeader.exportable = true;
        return true;
      }
    }
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
    this.store.dspname = this.$route.params.dspname;

    this.resetFilter();
  },
  methods: {
    getLeftTable() {
      this.$store.dispatch(`${this.storeName}/getLeftTable`);
    },
    isMainLevelCellShowing (props){
      if( ! props.expanded) return true;
      else {
        if(props.item.Tables.length > 0) {
          if(props.item.Tables[0].Columns.length == 0) return true;
        }
        
        return false;
      }
    },
    isTableLevelCellShowing (props){
      if( ! props.expanded) return true;
      else {
        if(props.item.Columns.length > 0) {
          return true;
        }
        
        return false;
      }
    },
    systemRowClick(evt) {
      evt.preventDefault();
    },
    linkDataLineage(){
      var processName = store.left.source[0] ? store.left.source[0]["PROCESS_NAME"] : "";

      switch (processName) {
        case "USFRR":
          window.open("http://hklpadmdm002.global.standardchartered.com:9999/data/view/id/535#!tab-data-relationships");
          break;
        case "GNS":
          window.open("http://hklpadmdm002.global.standardchartered.com:9999/data/view/id/532#!tab-data-relationships");
          break;
        default:
          break;
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
    getCDEConclusion(cdes) {
      return cdes
        .filter(Boolean)
        .join(", ")
        .indexOf("Yes") != -1
        ? "Yes"
        : "No";
    },
    showDetails(param) {
      this.$router.push(
        this.addressPath + "/" + this.$route.params.dspname + "/" + param.ALIAS_NAME
      );
    }
  }
};
</script>
