<style>
#table-ddo-downstream table.v-table tr th:nth-of-type(1){width: 10% !important;}
#table-ddo-downstream table.v-table tr th:nth-of-type(2){width: 10% !important;}
#table-ddo-downstream table.v-table tr th:nth-of-type(3){width: 10% !important;}
</style>

<template>
    <v-content class="mx-4 my-5">
        <b-container fluid>
            <PageHeader />

            <page-loader v-if="store.left.isLoading" />
            
            <b-row>
                <b-col sm=12 md=3>
                    <div class="card card-v2 transition">
                        <h6 class="title-1">Data Domain</h6>
                        <h3 class="title-2 text-capitalize">{{ store.left.source[0] ? store.left.source[0]["DATA_DOMAIN"] : " " }}</h3>
                    </div>
                </b-col>

                <b-col sm=12 md=3>
                    <div class="card card-v2 transition">
                        <h6 class="title-1">Sub-domain</h6>
                        <h3 class="title-2 text-capitalize">{{ store.left.source[0] ? store.left.source[0]["SUB_DOMAINS"] : " " }}</h3>
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
                    id="table-ddo-downstream">
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
                      <td v-bind:style="{ width: store.left.colWidth['DP_NAME'] + 'px' }" class="text-uppercase">{{ props.item.DP_NAME.toString().trim() ? props.item.DP_NAME : 'NA' }}</td>
                      <td v-bind:style="{ width: store.left.colWidth['PROCESS_OWNER'] + 'px' }" class="text-uppercase">{{ props.item.PROCESS_OWNER.toString().trim() ? props.item.PROCESS_OWNER : 'NA' }}</td>
                      <td v-bind:style="{ width: store.left.colWidth['BT_COUNT'] + 'px' }" class="text-capitalize">
                        <b-link @click.stop="showBusinessterms(props.item)">
                          <tablecell :fulltext="props.item.BT_COUNT" showOn="click"></tablecell></b-link></td>
                    </tr>
                  </template>

                  <template slot="footer" >
                    <td v-bind:style="{ width: (store.left.colWidth['DP_NAME']+store.left.colWidth['PROCESS_OWNER']) + 'px' }" colspan="2">
                      <b>Total</b>
                    </td>
                    <td v-bind:style="{ width: store.left.colWidth['BT_COUNT'] + 'px' }"><b>{{ store.left.display[0] ? store.left.display[0].TOTAL : 0 }}</b></td>
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
      storeName: "ddodownstream",
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
    this.store.system = this.$route.params.subdomain;
    this.resetFilter();
  },
  updated() {
    this.setTableColumnsWidth($('#table-ddo-downstream'));
    this.setTableFooterColumnsWidth($('#table-ddo-downstream'));
  },
  methods: {
    getLeftTable() {
      var getLeftTableVal = this.$store.dispatch(`${this.storeName}/getLeftTable`);
      getLeftTableVal.then(res => {
        this.setTableColumnsWidth($('#table-ddo-downstream'));
        this.setTableFooterColumnsWidth($('#table-ddo-downstream'));
      });
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
    showBusinessterms(param) {
      this.$router.push(
        this.addressPath + "/" + encodeURIComponent(this.$route.params.subdomain) + "/" + encodeURIComponent(param.DP_NAME)
      );
    },
    showDetails(param) {
      this.$router.push(
        this.addressPath + "/" + encodeURIComponent(param.TSID) + "/" + encodeURIComponent(param.ID) + "/" + encodeURIComponent(param.COLID)
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
    setTableFooterColumnsWidth(elem){
      var tableElem = elem.find('.v-table__overflow > table.v-table');
      var theadTR = tableElem.find('thead tr:first');
      var THs = theadTR.find('th');
      var tfootTR = tableElem.find('tfoot tr');
      var thWidths = [];
      THs.each(function (thIndex) {
        thWidths[thIndex] = $(this).outerWidth();
      });
      
      var TDs = tfootTR.find('td');
      var colspan = 1;
      var usedIndex = 0;
      TDs.each(function (tdIndex) {
        
        if ($(this)[0].hasAttribute('colspan')){
          colspan = $(this).attr('colspan');
          var colWidth = thWidths[tdIndex] + thWidths[(tdIndex+colspan)-1];
        }else{
          if(colspan>1){
            usedIndex = parseInt(tdIndex) + parseInt(colspan);
          }else{
            usedIndex = parseInt(tdIndex);
          }
          var colWidth = thWidths[usedIndex];
        }
        colWidth = colWidth - 75; // untuk mengurangi additional width yang datang tiba2 seperti syaiton, xixixi
        $(this).width(colWidth);
      });
    }
  }
};
</script>
