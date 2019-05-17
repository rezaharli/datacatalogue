<template>
    <v-content class="mx-4 my-5">
        <b-container fluid>
            <PageHeader title="Downstream Process Owner View" />

            <page-loader v-if="store.left.isLoading" />
            
            <b-row>
                <b-col sm=12 md=3>
                    <div class="card card-v2 transition">
                        <h6 class="title-1">Downstream Process Name</h6>
                        <h3 class="title-2 text-capitalize">{{ store.left.source[0] ? store.left.source[0]["DSP_NAME"] : " " }}</h3>
                    </div>
                </b-col>

                <b-col sm=12 md=3>
                    <div class="card card-v2 transition">
                        <h6 class="title-1">Data Elements</h6>
                        <h3 class="title-2 text-capitalize">{{ store.left.source[0] ? store.left.source[0]["COUNT_DATA_ELEMENTS"] : "0" }}</h3>
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
                    id="table-dpo-dataelements">
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
                      <td v-bind:style="{ width: store.left.colWidth['Details'] + 'px' }" class="text-capitalize text-title">
                        <b-button size="sm" class="green-tosca-gradient icon-only" @click="showDetails(props.item)">
                          <i class="fa fa-fw fa-external-link-alt"></i></b-button></td>

                      <td v-bind:style="{ width: store.left.colWidth['SYSTEM_NAME'] + 'px' }" class="text-uppercase">
                        <tablecell showOn="hover" :fulltext="props.item.SYSTEM_NAME"></tablecell></td>

                      <td v-bind:style="{ width: store.left.colWidth['ITAM_ID'] + 'px' }" class="text-uppercase">
                        <tablecell showOn="hover" :fulltext="props.item.ITAM_ID"></tablecell></td>

                      <td v-bind:style="{ width: store.left.colWidth['ALIAS_NAME'] + 'px' }" class="text-capitalize text-title">
                        <b-link @click="props.expanded = !props.expanded" v-if="props.item.Tables.length > 0">
                          <tablecell :fulltext="props.item.ALIAS_NAME" showOn="hover"></tablecell>
                        </b-link>

                        <tablecell :fulltext="props.item.BT_NAME" showOn="hover" v-if="props.item.Tables.length < 1"></tablecell>
                      </td>

                      <td v-bind:style="{ width: store.left.colWidth['CDE'] + 'px' }" class="text-uppercase">
                        {{ props.item.CDE }}</td>

                      <td v-bind:style="{ width: store.left.colWidth['TABLE_NAME'] + 'px' }" class="text-uppercase">
                        <tablecell showOn="hover" v-if="isMainLevelCellShowing(props)" :fulltext="props.item.TABLE_NAME"></tablecell>
                      </td>

                      <td v-bind:style="{ width: store.left.colWidth['COLUMN_NAME'] + 'px' }" class="text-uppercase">
                        <tablecell showOn="hover" v-if="isMainLevelCellShowing(props)" :fulltext="props.item.COLUMN_NAME"></tablecell>
                      </td>

                      <td v-bind:style="{ width: store.left.colWidth['ULT_SYSTEM_NAME'] + 'px' }" class="text-uppercase">
                        <tablecell showOn="hover" v-if="isMainLevelCellShowing(props)" :fulltext="props.item.ULT_SYSTEM_NAME"></tablecell>
                      </td>

                      <td v-bind:style="{ width: store.left.colWidth['DATA_SLA'] + 'px' }" class="text-uppercase">
                        <tablecell showOn="hover" v-if="isMainLevelCellShowing(props)" :fulltext="props.item.DATA_SLA"></tablecell>
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
                        <td v-bind:style="{ width: store.left.colWidth['SYSTEM_NAME'] + 'px' }">&nbsp;</td>
                        <td v-bind:style="{ width: store.left.colWidth['ITAM_ID'] + 'px' }">&nbsp;</td>
                        <td v-bind:style="{ width: store.left.colWidth['ALIAS_NAME'] + 'px' }">&nbsp;</td>
                        <td v-bind:style="{ width: store.left.colWidth['CDE'] + 'px' }">&nbsp;</td>

                        <td v-bind:style="{ width: store.left.colWidth['TABLE_NAME'] + 'px' }">
                          <b-link @click="props.expanded = !props.expanded" v-if="props.item.Columns.length >= 1">
                            <tablecell :fulltext="props.item.TABLE_NAME" showOn="hover"></tablecell>
                          </b-link>

                          <tablecell :fulltext="props.item.TABLE_NAME" showOn="hover" v-if="props.item.Columns.length < 1"></tablecell>
                        </td>

                        <td class="text-uppercase" v-bind:style="{ width: store.left.colWidth['COLUMN_NAME'] + 'px' }">
                          <tablecell showOn="hover" v-if="isTableLevelCellShowing(props)" :fulltext="props.item.COLUMN_NAME"></tablecell>
                        </td>

                        <td class="text-uppercase" v-bind:style="{ width: store.left.colWidth['ULT_SYSTEM_NAME'] + 'px' }">
                          <tablecell showOn="hover" v-if="isTableLevelCellShowing(props)" :fulltext="props.item.ULT_SYSTEM_NAME"></tablecell>
                        </td>

                        <td class="text-uppercase" v-bind:style="{ width: store.left.colWidth['DATA_SLA'] + 'px' }">
                          <tablecell showOn="hover" v-if="isTableLevelCellShowing(props)" :fulltext="props.item.DATA_SLA"></tablecell>
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
                            <td v-bind:style="{ width: store.left.colWidth['SYSTEM_NAME'] + 'px' }">&nbsp;</td>
                            <td v-bind:style="{ width: store.left.colWidth['ITAM_ID'] + 'px' }">&nbsp;</td>
                            <td v-bind:style="{ width: store.left.colWidth['ALIAS_NAME'] + 'px' }">&nbsp;</td>
                            <td v-bind:style="{ width: store.left.colWidth['CDE'] + 'px' }">&nbsp;</td>
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
      storeName: "dpodataelements",
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
    this.store.dspname = this.$route.params.dspname;
    this.resetFilter();
    setTimeout(() => {
      this.setTableColumnsWidth($('#table-dpo-dataelements'));
    }, 300);
  },
  updated() {
    this.setTableColumnsWidth($('#table-dpo-dataelements'));
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
    setExpandedTableColumnsWidth(){
      setTimeout(() => {
        var elem = $('.v-datatable__expand-row');
        var elemExpandedTable = elem.find('.v-datatable__expand-content table.v-table');
        var THs = elem.closest('table.v-table').find('thead tr:first th');
        var tbodyTR = elemExpandedTable.find('tbody tr');
        THs.each(function (thIndex) {
          $(this).css({'color': 'red'});
          var thWidth = $(this).width();
          tbodyTR.each(function (tdIndex) {
            var TDs = $(this).find('td:not([colspan])');
            TDs.eq(thIndex).width(thWidth);
          });
        });
      }, 10);
    }
  }
};
</script>
