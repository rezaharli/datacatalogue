<style>
#table-dsc-dd-personal table.v-table tr th:nth-of-type(1){width: calc(100%/20) !important; display: table-cell;}
#table-dsc-dd-personal table.v-table tr th:nth-of-type(2){width: calc(100%/20) !important; display: table-cell;}
#table-dsc-dd-personal table.v-table tr th:nth-of-type(3){width: calc(100%/20) !important; display: table-cell;}
#table-dsc-dd-personal table.v-table tr th:nth-of-type(4){width: calc(100%/20) !important; display: table-cell;}
#table-dsc-dd-personal table.v-table tr th:nth-of-type(5){width: calc(100%/20) !important; display: table-cell;}
#table-dsc-dd-personal table.v-table tr th:nth-of-type(6){width: calc(100%/20) !important; display: table-cell;}
#table-dsc-dd-personal table.v-table tr th:nth-of-type(7){width: calc(100%/20) !important; display: table-cell;}
#table-dsc-dd-personal table.v-table tr th:nth-of-type(8){width: calc(100%/20) !important; display: table-cell;}
#table-dsc-dd-personal table.v-table tr th:nth-of-type(9){width: calc(100%/20) !important; display: table-cell;}
#table-dsc-dd-personal table.v-table tr th:nth-of-type(10){width: calc(100%/20) !important; display: table-cell;}
#table-dsc-dd-personal table.v-table tr th:nth-of-type(11){width: calc(100%/20) !important; display: table-cell;}
#table-dsc-dd-personal table.v-table tr th:nth-of-type(12){width: calc(100%/20) !important; display: table-cell;}
#table-dsc-dd-personal table.v-table tr th:nth-of-type(13){width: calc(100%/20) !important; display: table-cell;}

.row-action-buttons{ top: -59px; position: absolute; right: 15px; }
</style>

<template>
  <b-container fluid>
    <page-loader v-if="store.left.isLoading" />

    <b-row>
      <b-col>
        <b-row class="my-4">
            <b-col>
                <b-button class="float-right red-neon icon-only ml-3" @click="resetFilter">
                    <i class="fa fa-fw fa-filter"></i>
                </b-button>

                <page-export class="float-right" :storeName="storeName" :leftTableCols="store.leftHeaders" :rightTableCols="[]"/>
            </b-col>
        </b-row>

        <b-row class="my-4">
          <b-col>
            <!-- Main content -->
            <!-- <div class="table-v2-title">Policy Related Information</div> -->
            <v-data-table
                :headers="store.leftHeaders.filter(v => v.display == true)"
                :items="store.left.display"
                :pagination.sync="store.left.pagination"
                :total-items="store.left.totalItems"
                :loading="store.left.isLoading"
                :expand="false"
                :must-sort="true"
                :rows-per-page-items="[25, 50, 75, 100]"
                item-key="ID"
                class="table-v1"
                id="table-dsc-dd-personal">
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
                  <td v-bind:style="{ width: store.left.colWidth['TABLE_NAME'] + 'px' }">
                    <tablecell :fulltext="props.item.TABLE_NAME" showOn="click"></tablecell></td>
                  
                  <td v-bind:style="{ width: store.left.colWidth['COLUMN_NAME'] + 'px' }">
                    <tablecell :fulltext="props.item.COLUMN_NAME" showOn="click"></tablecell></td>
                  
                  <td v-bind:style="{ width: store.left.colWidth['BUSINESS_ALIAS_NAME'] + 'px' }">
                    <tablecell :fulltext="props.item.BUSINESS_ALIAS_NAME" showOn="click"></tablecell></td>
                  
                  <td v-bind:style="{ width: store.left.colWidth['BUSINESS_ALIAS_DESCRIPTION'] + 'px' }">
                    <tablecell :fulltext="props.item.BUSINESS_ALIAS_DESCRIPTION" showOn="click"></tablecell></td>
                  
                  <td v-bind:style="{ width: store.left.colWidth['CDE_YES_NO'] + 'px' }">
                    <tablecell :fulltext="props.item.CDE_YES_NO" showOn="click"></tablecell></td>
                  
                  <td v-bind:style="{ width: store.left.colWidth['PII_FLAG'] + 'px' }">
                    <tablecell :fulltext="props.item.PII_FLAG" showOn="click"></tablecell></td>
                </tr>
              </template>
            </v-data-table>
                  
          </b-col>
        </b-row>
      </b-col>
    </b-row>
  </b-container>
</template>

<script>
import Vue from "vue";
import { mapState, mapActions } from "vuex";

import JsonExcel from "vue-json-excel";
import pageExport from "../PageExport.vue";
import tableheader from "../TableHeader.vue";
import tablecell from "../Tablecell.vue";
import pageLoader from "../PageLoader.vue";

Vue.component("downloadExcel", JsonExcel);

export default {
  name: "DscIarcPersonal",
  components: {
    pageExport, tableheader, tablecell, pageLoader
  },
  data() {
    return {
      storeName: "dsciarcpersonal",
      systemSource: [],
      tablenameSource: []
    };
  },
  computed: {
    store() {
      return this.$store.state[this.storeName].all;
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
  updated() {
    this.setTableColumnsWidth($('#table-dsc-dd-personal'));
  },
  methods: {
    getLeftTable() {
      this.store.system = this.$route.params.system;
      var getLeftTableVal = this.$store.dispatch(`${this.storeName}/getLeftTable`);
      getLeftTableVal.then(res => {
        this.setTableColumnsWidth($('#table-dsc-dd-personal'));
      });
    },
    resetFilter (e) {
        if(Object.keys(this.store.filters.left).length > 0){
            this.store.filters.left = {};
            this.getLeftTable();
        }
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
};
</script>
