<style>
#table-dsc-cde table.v-table .table-header-wrapper {font-size: 0.97em; }
#table-dsc-cde table.v-table .table-header-wrapper .dropdown-button-wrapper .btn {padding: 0;}
#table-dsc-cde table.v-table thead tr th:first-of-type,
#table-dsc-cde table.v-table tbody tr td:first-of-type{padding-left: 20px;}
#table-dsc-cde table.v-table thead tr th:last-of-type,
#table-dsc-cde table.v-table tbody tr td:last-of-type{padding-right: 20px;}
#table-dsc-cde table.v-table thead tr th,
#table-dsc-cde table.v-table tbody tr td {padding-left: 10px; padding-right: 10px; }
#table-dsc-cde table.v-table tr th:nth-of-type(1){width: 8% !important;}
#table-dsc-cde table.v-table tr th:nth-of-type(2){width: 25% !important;}
#table-dsc-cde table.v-table tr th:nth-of-type(3){width: 19% !important;}
#table-dsc-cde table.v-table tr th:nth-of-type(4){width: 13% !important;}
#table-dsc-cde table.v-table tr th:nth-of-type(5){width: 13% !important;}
#table-dsc-cde table.v-table tr th:nth-of-type(6){width: 10% !important;}
#table-dsc-cde table.v-table tr th:nth-of-type(7){width: 12% !important;}
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
                        <h6 class="title-1">Critical Data Elements</h6>
                        <h3 class="title-2 text-capitalize">{{ store.left.source[0] ? store.left.source[0]["COUNT_CDE"] : "0" }}</h3>
                    </div>
                </b-col>
                <b-col sm=12 md=3>
                    <div class="card card-v2 transition">
                        <h6 class="title-1">Downstream Process</h6>
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
                    class="table-v1"
                    id="table-dsc-cde">
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
                      <td v-bind:style="{ width: store.left.colWidth['Details'] + 'px' }" class="text-capitalize text-title">
                        <b-button size="sm" class="green-tosca-gradient icon-only" @click="showDetails(props.item)">
                          <i class="fa fa-fw fa-external-link-alt"></i></b-button></td>

                      <td v-bind:style="{ width: store.left.colWidth['CDE'] + 'px' }" class="text-capitalize text-title">
                        <b-link @click="props.expanded = !props.expanded" v-if="props.item.Tables.length > 0">
                          {{props.item.CDE}}
                        </b-link>

                        <span v-if="props.item.Tables.length < 1">{{props.item.CDE}}</span>
                      </td>

                      <td v-bind:style="{ width: store.left.colWidth['DESCRIPTION'] + 'px' }" class="text-description">
                        <tablecell :fulltext="props.item.DESCRIPTION" showOn="hover"></tablecell></td>

                      <td v-bind:style="{ width: store.left.colWidth['TABLE_NAME'] + 'px' }" class="text-uppercase">
                        <tablecell showOn="hover" v-if="isMainLevelCellShowing(props)" :fulltext="props.item.TABLE_NAME"></tablecell>
                      </td>

                      <td v-bind:style="{ width: store.left.colWidth['COLUMN_NAME'] + 'px' }" class="text-uppercase">
                        <tablecell showOn="hover" v-if="isMainLevelCellShowing(props)" :fulltext="props.item.COLUMN_NAME"></tablecell>
                      </td>

                      <td v-bind:style="{ width: store.left.colWidth['DSP_NAME'] + 'px' }" class="text-uppercase">
                        <tablecell showOn="hover" v-if="isMainLevelCellShowing(props)" :fulltext="props.item.DSP_NAME"></tablecell></td>
                        
                      <td v-bind:style="{ width: store.left.colWidth['PROCESS_OWNER'] + 'px' }">
                        <tablecell showOn="hover" v-if="isMainLevelCellShowing(props)" :fulltext="props.item.PROCESS_OWNER"></tablecell></td>
                        <!-- <tablecell v-if="!props.expanded" 
                          :fulltext="props.item.Tables[0] ? (props.item.Tables[0].Columns[0] ? (props.item.Tables[0].Columns[0].Dsps[0] ? _.uniq(_.map(props.item.Tables[0].Columns[0].Dsps[0].DspsVal, 'PROCESS_OWNER')).filter(Boolean).join(', ') : props.item.PROCESS_OWNER) : props.item.PROCESS_OWNER) : props.item.PROCESS_OWNER" showOn="hover"></tablecell></td> -->
                    </tr>
                  </template>

                  <template slot="expand" slot-scope="props">
                    <!-- <table-rows-sub :storeName="storeName" :props="props" /> -->
                    <v-data-table
                      :headers="store.leftHeaders.filter(v => v.display == true)"
                      :items="props.item.Tables"
                      :expand="false"
                      class=""
                      item-key="TMTID"
                      hide-actions
                      hide-headers
                    >
                      <template slot="items" slot-scope="props">
                        <td v-bind:style="{ width: store.left.colWidth['Details'] + 'px' }">&nbsp;</td>
                        <td v-bind:style="{ width: store.left.colWidth['CDE'] + 'px' }">&nbsp;</td>
                        <td v-bind:style="{ width: store.left.colWidth['DESCRIPTION'] + 'px' }">&nbsp;</td>

                        <td v-bind:style="{ width: store.left.colWidth['TABLE_NAME'] + 'px' }">
                          <b-link @click="props.expanded = !props.expanded" v-if="props.item.Columns.length >= 1">
                            <tablecell :fulltext="props.item.TABLE_NAME" showOn="hover"></tablecell>
                          </b-link>

                          <tablecell :fulltext="props.item.TABLE_NAME" showOn="hover" v-if="props.item.Columns.length < 1"></tablecell>
                        </td>

                        <td class="text-uppercase" v-bind:style="{ width: store.left.colWidth['COLUMN_NAME'] + 'px' }">
                          <tablecell showOn="hover" v-if="isTableLevelCellShowing(props)" :fulltext="props.item.COLUMN_NAME"></tablecell>
                        </td>

                        <td class="text-uppercase" v-bind:style="{ width: store.left.colWidth['DSP_NAME'] + 'px' }">
                          <tablecell showOn="hover" v-if="isTableLevelCellShowing(props)" :fulltext="props.item.DSP_NAME"></tablecell>  
                        </td>

                        <td v-bind:style="{ width: store.left.colWidth['PROCESS_OWNER'] + 'px' }">
                          <tablecell showOn="hover" v-if="isTableLevelCellShowing(props)" :fulltext="props.item.PROCESS_OWNER"></tablecell></td>
                          <!-- <tablecell v-if="!props.expanded" showOn="hover" :fulltext="props.item.Columns[0] ? (props.item.Columns[0].Dsps[0] ? _.uniq(_.map(props.item.Columns[0].Dsps[0].DspsVal, 'PROCESS_OWNER')).filter(Boolean).join(', ') : props.item.PROCESS_OWNER) : props.item.PROCESS_OWNER"></tablecell></td> -->
                      </template>

                      <template slot="expand" slot-scope="props">
                        <v-data-table
                          :headers="store.leftHeaders.filter(v => v.display == true)"
                          :items="props.item.Columns"
                          item-key="COLID"
                          class=""
                          hide-actions
                          hide-headers
                        >
                          <template slot="items" slot-scope="props">
                            <td v-bind:style="{ width: store.left.colWidth['Details'] + 'px' }">&nbsp;</td>
                            <td v-bind:style="{ width: store.left.colWidth['CDE'] + 'px' }">&nbsp;</td>
                            <td v-bind:style="{ width: store.left.colWidth['DESCRIPTION'] + 'px' }">&nbsp;</td>
                            <td v-bind:style="{ width: store.left.colWidth['TABLE_NAME'] + 'px' }">&nbsp;</td>
                            <td v-bind:style="{ width: store.left.colWidth['COLUMN_NAME'] + 'px' }">
                              <b-link @click="props.expanded = !props.expanded" v-if="props.item.Dsps.length >= 1">
                                <tablecell :fulltext="props.item.COLUMN_NAME" showOn="hover"></tablecell>
                              </b-link>

                              <tablecell :fulltext="props.item.COLUMN_NAME" showOn="hover" v-if="props.item.Dsps.length < 1"></tablecell>
                            </td>

                            <td class="text-uppercase" v-bind:style="{ width: store.left.colWidth['DSP_NAME'] + 'px' }">
                              <tablecell :fulltext="props.item.DSP_NAME" showOn="hover"></tablecell></td>

                            <td v-bind:style="{ width: store.left.colWidth['PROCESS_OWNER'] + 'px' }">
                              <tablecell showOn="hover" :fulltext="props.item.PROCESS_OWNER"></tablecell></td>
                              <!-- <tablecell showOn="hover" :fulltext="props.item.Dsps[0] ? _.uniq(_.map(props.item.Dsps[0].DspsVal, 'PROCESS_OWNER')).filter(Boolean).join(', ') : props.item.PROCESS_OWNER"></tablecell></td> -->
                          </template>

                          <template slot="expand" slot-scope="props">
                            <v-data-table
                              :headers="store.leftHeaders.filter(v => v.display == true)"
                              :items="props.item.Dsps"
                              class=""
                              hide-actions
                              hide-headers
                            >
                              <template slot="items" slot-scope="props">
                                <td v-bind:style="{ width: store.left.colWidth['Details'] + 'px' }">&nbsp;</td>
                                <td v-bind:style="{ width: store.left.colWidth['CDE'] + 'px' }">&nbsp;</td>
                                <td v-bind:style="{ width: store.left.colWidth['DESCRIPTION'] + 'px' }">&nbsp;</td>
                                <td v-bind:style="{ width: store.left.colWidth['TABLE_NAME'] + 'px' }">&nbsp;</td>
                                <td v-bind:style="{ width: store.left.colWidth['COLUMN_NAME'] + 'px' }">&nbsp;</td>

                                <td v-bind:style="{ width: store.left.colWidth['DSP_NAME'] + 'px' }">
                                  <tablecell :fulltext="props.item.DSP_NAME" showOn="hover"></tablecell></td>
                                
                                <td v-bind:style="{ width: store.left.colWidth['PROCESS_OWNER'] + 'px' }">
                                  <tablecell :fulltext="(_.uniq(_.map(props.item.DspsVal, 'PROCESS_OWNER')).filter(Boolean).join(', '))" showOn="hover"></tablecell></td>
                              </template>
                            </v-data-table>
                          </template>
                        </v-data-table>
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
      storeName: "dsccde",
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
          if(props.item.Columns[0].Dsps.length == 0) return true;
        }
        
        return false;
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
    showDetails(param) {
      this.$router.push(
        this.addressPath + "/" + param.SYSTEM_NAME + "/" + param.TSID + "/" + param.TMTID + "/" + param.COLID
      );
    }
  }
};
</script>
