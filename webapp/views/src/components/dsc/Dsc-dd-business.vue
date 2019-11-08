<style>
/* #table-dsc-dd-business table.v-table tr {display: block;} */
/* #table-dsc-dd-business table.v-table.v-datatable thead{
    width: unset;
    display: table-header-group;
    padding-right: unset;
}
#table-dsc-dd-business table.v-table.v-datatable tbody{
    display:table-row-group;
    overflow:auto;
    max-height:unset;
    width:unset;
}
#table-dsc-dd-business table.v-table.v-datatable tbody tr {display: table-row;} */
#table-dsc-dd-business table.v-table tr th:nth-of-type(1){width: calc(100%/20) !important; display: table-cell;}
#table-dsc-dd-business table.v-table tr th:nth-of-type(2){width: calc(100%/20) !important; display: table-cell;}
#table-dsc-dd-business table.v-table tr th:nth-of-type(3){width: calc(100%/20) !important; display: table-cell;}
#table-dsc-dd-business table.v-table tr th:nth-of-type(4){width: calc(100%/20) !important; display: table-cell;}
#table-dsc-dd-business table.v-table tr th:nth-of-type(5){width: calc(100%/20) !important; display: table-cell;}
#table-dsc-dd-business table.v-table tr th:nth-of-type(6){width: calc(100%/20) !important; display: table-cell;}
#table-dsc-dd-business table.v-table tr th:nth-of-type(7){width: calc(100%/20) !important; display: table-cell;}
#table-dsc-dd-business table.v-table tr th:nth-of-type(8){width: calc(100%/20) !important; display: table-cell;}
#table-dsc-dd-business table.v-table tr th:nth-of-type(9){width: calc(100%/20) !important; display: table-cell;}
#table-dsc-dd-business table.v-table tr th:nth-of-type(10){width: calc(100%/20) !important; display: table-cell;}
</style>

<template>
    <b-row style="margin-top: 10px;margin-bottom: 10px;">
      <b-col>
        <!-- Main content -->
        <div class="table-v2-title">Business Metadata</div>
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
            class="table-v2"
            id="table-dsc-dd-business">
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
              <!-- <td v-bind:style="{ width: store.left.colWidth['SYSTEM_NAME'] + 'px' }">
                <tablecell :fulltext="props.item.SYSTEM_NAME" showOn="click"></tablecell></td>
              
              <td v-bind:style="{ width: store.left.colWidth['ITAM_ID'] + 'px' }">
                <tablecell :fulltext="props.item.ITAM_ID" showOn="click"></tablecell></td> -->

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
                
              <td v-bind:style="{ width: store.left.colWidth['DOMAIN'] + 'px' }">
                <tablecell :fulltext="props.item.DOMAIN" showOn="click"></tablecell></td>
              
              <td v-bind:style="{ width: store.left.colWidth['SUBDOMAIN'] + 'px' }">
                <tablecell :fulltext="props.item.SUBDOMAIN" showOn="click"></tablecell></td>
              
              <td v-bind:style="{ width: store.left.colWidth['DOMAIN_OWNER'] + 'px' }">
                <tablecell :fulltext="props.item.DOMAIN_OWNER" showOn="click"></tablecell></td>
              
              <td v-bind:style="{ width: store.left.colWidth['BUSINESS_TERM'] + 'px' }">
                <tablecell :fulltext="props.item.BUSINESS_TERM" showOn="click"></tablecell></td>
              
              <td v-bind:style="{ width: store.left.colWidth['BUSINESS_TERM_DESCRIPTION'] + 'px' }">
                <tablecell :fulltext="props.item.BUSINESS_TERM_DESCRIPTION" showOn="click"></tablecell></td>
            </tr>
          </template>
        </v-data-table>
              
      </b-col>
    </b-row>
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
  name: "DscDdBusiness",
  components: {
    PageHeader, pageSearch, pageExport, tableheader, tablecell, pageLoader
  },
  data() {
    return {
      storeName: "dscddBusiness",
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
    var self = this;

    $("#page-tab #tab-business").on('click', function(){
      setTimeout(() => {
        self.setTableColumnsWidth($('#table-dsc-dd-business'));
      }, 1);
    });
  },
  updated() {
    this.setTableColumnsWidth($('#table-dsc-dd-business'));
  },
  methods: {
    getLeftTable() {
      this.store.system = this.$route.params.system;
      var getLeftTableVal = this.$store.dispatch(`${this.storeName}/getLeftTable`);
      getLeftTableVal.then(res => {
        this.$store.dispatch(`header/getRowCount`, this.store.param).then(res => {
          this.store.left.totalItems = res.Data;
        });

        this.setTableColumnsWidth($('#table-dsc-dd-business'));
      });
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
