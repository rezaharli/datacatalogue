<style>
#table-dsc-interfaces table.v-table tr th:nth-of-type(1){width: 40% !important;}
#table-dsc-interfaces table.v-table tr th:nth-of-type(2){width: 32% !important;}
#table-dsc-interfaces table.v-table tr th:nth-of-type(3){width: 38% !important;}
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
                        <h6 class="title-1">Immediate Interfaces</h6>
                        <h3 class="title-2 text-capitalize">{{ store.left.source[0] ? store.left.source[0]["COUNT_IMM_INTERFACE"] : "0" }}</h3>
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
                    class="table-v1"
                    id="table-dsc-interfaces">
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
                      <td style="width: calc(100% / 6)" class="text-capitalize">{{props.item.IMM_INTERFACE}}</td>
                      <td style="width: calc(100% / 6)" class="text-capitalize"><b-link @click.stop="showCDEs(props.item)"><tablecell :fulltext="props.item.CDE_COUNT" showOn="hover"></tablecell></b-link></td>
                      <td style="width: calc(100% / 6)" class="text-capitalize"><tablecell :fulltext="props.item.PROCESS_OWNER" showOn="hover"></tablecell></td>
                    </tr>
                  </template> -->

                  <template slot="items" slot-scope="props">
                    <tr :class="{even: props.index % 2, odd: !(props.index % 2)}">
                      
                      <td v-bind:style="{ width: store.left.colWidth['IMM_INTERFACE'] + 'px' }" class="text-capitalize text-title">
                        <b-link @click="props.expanded = !props.expanded" v-if="props.item.Owners.length > 0">
                          <tablecell :fulltext="props.item.IMM_INTERFACE" showOn="hover"></tablecell>
                        </b-link>

                        <tablecell :fulltext="props.item.IMM_INTERFACE" showOn="hover" v-if="props.item.Owners.length < 1"></tablecell>
                      </td>

                      <td v-bind:style="{ width: store.left.colWidth['CDE_COUNT'] + 'px' }" class="text-uppercase">
                        <b-link @click.stop="showCDEs(props.item)">
                          <tablecell showOn="hover" v-if="isMainLevelCellShowing(props)" :fulltext="props.item.CDE_COUNT"></tablecell>
                        </b-link>
                      </td>

                      <td v-bind:style="{ width: store.left.colWidth['PROCESS_OWNER'] + 'px' }" class="text-uppercase">
                        <tablecell showOn="hover" v-if="isMainLevelCellShowing(props)" :fulltext="props.item.PROCESS_OWNER"></tablecell>
                      </td>
                    </tr>
                  </template>

                  <template slot="expand" slot-scope="props">
                    <v-data-table
                      :headers="store.leftHeaders.filter(v => v.display == true)"
                      :items="props.item.Owners"
                      item-key="COLID"
                      class=""
                      hide-actions
                      hide-headers
                    >
                      <template slot="items" slot-scope="props">
                        <td v-bind:style="{ width: store.left.colWidth['IMM_INTERFACE'] + 'px' }">&nbsp;</td>
                        <td v-bind:style="{ width: store.left.colWidth['CDE_COUNT'] + 'px' }">&nbsp;</td>
                        <td v-bind:style="{ width: store.left.colWidth['PROCESS_OWNER'] + 'px' }">
                          <tablecell :fulltext="props.item.PROCESS_OWNER" showOn="hover"></tablecell>
                        </td>
                      </template>
                    </v-data-table>
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
      storeName: "dscinterfaces",
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
    isMainLevelCellShowing (props){
      if( ! props.expanded) return true;
      else {
        if(props.item.Owners.length > 0) {
          return true;
        }
        
        return false;
      }
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
        this.addressPath + "/" + encodeURIComponent(param.SYSTEM_NAME) + "/" + encodeURIComponent(param.IMM_INTERFACE)
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
