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
                        <h6 class="title-1">Immediate Interface</h6>
                        <h3 class="title-2 text-capitalize">{{ $route.params.dspName }}</h3>
                    </div>
                </b-col>
                <b-col sm=12 md=3>
                    <div class="card card-v2 transition">
                        <h6 class="title-1">Critical Data Elements</h6>
                        <h3 class="title-2 text-capitalize">{{ store.left.source[0] ? store.left.source[0]["COUNT_CDE"] : "0" }}</h3>
                    </div>
                </b-col>
            </b-row>

            <b-row style="margin-top: 10px; margin-bottom: 20px;">
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
                    item-key="ID"
                    class="table-v1 table-w-button-left">
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
                          <tablecell :fulltext="props.item.CDE" showOn="hover"></tablecell>
                        </b-link>

                        <tablecell :fulltext="props.item.CDE" showOn="hover" v-if="props.item.Tables.length < 1"></tablecell>
                      </td>

                      <td v-bind:style="{ width: store.left.colWidth['DESCRIPTION'] + 'px' }" class="text-description">
                        <tablecell :fulltext="props.item.DESCRIPTION" showOn="hover"></tablecell></td>

                      <td v-bind:style="{ width: store.left.colWidth['TABLE_NAME'] + 'px' }" class="text-uppercase">
                        <tablecell showOn="hover" v-if="isMainLevelCellShowing(props)" :fulltext="props.item.TABLE_NAME"></tablecell>
                      </td>

                      <td v-bind:style="{ width: store.left.colWidth['COLUMN_NAME'] + 'px' }" class="text-uppercase">
                        <tablecell showOn="hover" v-if="isMainLevelCellShowing(props)" :fulltext="props.item.COLUMN_NAME"></tablecell>
                      </td>
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
                              <tablecell :fulltext="props.item.COLUMN_NAME" showOn="hover"></tablecell>
                            </td>
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
      storeName: "dscinterfacescde",
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
    this.store.dspName = this.$route.params.dspName;
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
        this.addressPath + "/" + this.store.system + "/" + this.store.dspName + "/" + param.TSID + "/" + param.TMTID + "/" + param.COLID
      );
    }
  }
};
</script>
