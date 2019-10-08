<style>
/* #table-edmp-dd-technical table.v-table tbody tr {display: block;} */
/* #table-edmp-dd-technical table.v-table.v-datatable thead{
    width: unset;
    display: table-header-group;
    padding-right: unset;
}
#table-edmp-dd-technical table.v-table.v-datatable tbody{
    display:table-row-group;
    overflow:auto;
    max-height:unset;
    width:unset;
}
#table-edmp-dd-technical table.v-table.v-datatable tbody tr {display: table-row;} */
#table-edmp-dd-technical table.v-table tr th:nth-of-type(1){width: calc(100%/20) !important; display: table-cell;}
#table-edmp-dd-technical table.v-table tr th:nth-of-type(2){width: calc(100%/20) !important; display: table-cell;}
#table-edmp-dd-technical table.v-table tr th:nth-of-type(3){width: calc(100%/20) !important; display: table-cell;}
#table-edmp-dd-technical table.v-table tr th:nth-of-type(4){width: calc(100%/20) !important; display: table-cell;}
#table-edmp-dd-technical table.v-table tr th:nth-of-type(5){width: calc(100%/20) !important; display: table-cell;}
#table-edmp-dd-technical table.v-table tr th:nth-of-type(6){width: calc(100%/20) !important; display: table-cell;}
#table-edmp-dd-technical table.v-table tr th:nth-of-type(7){width: calc(100%/20) !important; display: table-cell;}
#table-edmp-dd-technical table.v-table tr th:nth-of-type(8){width: calc(100%/20) !important; display: table-cell;}
#table-edmp-dd-technical table.v-table tr th:nth-of-type(9){width: calc(100%/20) !important; display: table-cell;}
#table-edmp-dd-technical table.v-table tr th:nth-of-type(10){width: calc(100%/20) !important; display: table-cell;}
#table-edmp-dd-technical table.v-table tr th:nth-of-type(11){width: calc(100%/20) !important; display: table-cell;}
#table-edmp-dd-technical table.v-table tr th:nth-of-type(12){width: calc(100%/20) !important; display: table-cell;}
#table-edmp-dd-technical table.v-table tr th:nth-of-type(13){width: calc(100%/20) !important; display: table-cell;}
</style>

<template>
    <b-row style="margin-top: 10px;margin-bottom: 10px;">
      <b-col>
        <!-- Main content -->
        <div class="table-v2-title">Technical Metadata</div>
        
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
            id="table-edmp-dd-technical">
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

              <td v-bind:style="{ width: store.left.colWidth['EDM_SOURCE_SYSTEM_NAME'] + 'px' }">
                <tablecell :fulltext="props.item.EDM_SOURCE_SYSTEM_NAME" showOn="click"></tablecell></td>
              
              <td v-bind:style="{ width: store.left.colWidth['DATABASE_NAME'] + 'px' }">
                <tablecell :fulltext="props.item.DATABASE_NAME" showOn="click"></tablecell></td>
              
              <td v-bind:style="{ width: store.left.colWidth['TABLE_NAME'] + 'px' }">
                <tablecell :fulltext="props.item.TABLE_NAME" showOn="click"></tablecell></td>
              
              <td v-bind:style="{ width: store.left.colWidth['COLUMN_NAME'] + 'px' }">
                <tablecell :fulltext="props.item.COLUMN_NAME" showOn="click"></tablecell></td>
              
              <td v-bind:style="{ width: store.left.colWidth['DATA_TYPE'] + 'px' }">
                <tablecell :fulltext="props.item.DATA_TYPE" showOn="click"></tablecell></td>
              
              <td v-bind:style="{ width: store.left.colWidth['COLUMN_LENGTH'] + 'px' }">
                <tablecell :fulltext="props.item.COLUMN_LENGTH" showOn="click"></tablecell></td>
              
              <td v-bind:style="{ width: store.left.colWidth['NULLABLE'] + 'px' }">
                <tablecell :fulltext="props.item.NULLABLE" showOn="click"></tablecell></td>
              
              <td v-bind:style="{ width: store.left.colWidth['PRIMARY_KEY'] + 'px' }">
                <tablecell :fulltext="props.item.PRIMARY_KEY" showOn="click"></tablecell></td>
              
              <td v-bind:style="{ width: store.left.colWidth['CERTIFIED'] + 'px' }">
                <tablecell :fulltext="props.item.CERTIFIED" showOn="click"></tablecell></td>
              
              <td v-bind:style="{ width: store.left.colWidth['PROFILE_RESULTS'] + 'px' }">
                <tablecell :fulltext="props.item.PROFILE_RESULTS" showOn="click"></tablecell></td>
              
              <td v-bind:style="{ width: store.left.colWidth['DATA_LINEAGE'] + 'px' }">
                <tablecell :fulltext="props.item.DATA_LINEAGE" showOn="click"></tablecell></td>
              
              <td v-bind:style="{ width: store.left.colWidth['BUSINESS_ALIAS_NAME'] + 'px' }">
                <tablecell :fulltext="props.item.BUSINESS_ALIAS_NAME" showOn="click"></tablecell></td>
              
              <td v-bind:style="{ width: store.left.colWidth['BUSINESS_ALIAS_DESCRIPTION'] + 'px' }">
                <tablecell :fulltext="props.item.BUSINESS_ALIAS_DESCRIPTION" showOn="click"></tablecell></td>
              
              <td v-bind:style="{ width: store.left.colWidth['CDE'] + 'px' }">
                <tablecell :fulltext="props.item.CDE" showOn="click"></tablecell></td>
              
              <td v-bind:style="{ width: store.left.colWidth['DATA_LENGTH'] + 'px' }">
                <tablecell :fulltext="props.item.DATA_LENGTH" showOn="click"></tablecell></td>
              
              <td v-bind:style="{ width: store.left.colWidth['EXAMPLE'] + 'px' }">
                <tablecell :fulltext="props.item.EXAMPLE" showOn="click"></tablecell></td>
              
              <td v-bind:style="{ width: store.left.colWidth['DERIVED'] + 'px' }">
                <tablecell :fulltext="props.item.DERIVED" showOn="click"></tablecell></td>
              
              <td v-bind:style="{ width: store.left.colWidth['DERIVATION_LOGIC'] + 'px' }">
                <tablecell :fulltext="props.item.DERIVATION_LOGIC" showOn="click"></tablecell></td>
              
              <td v-bind:style="{ width: store.left.colWidth['SOURCED_FROM_UPSTREAM'] + 'px' }">
                <tablecell :fulltext="props.item.SOURCED_FROM_UPSTREAM" showOn="click"></tablecell></td>
              
              <td v-bind:style="{ width: store.left.colWidth['SYSTEM_CHECKS'] + 'px' }">
                <tablecell :fulltext="props.item.SYSTEM_CHECKS" showOn="click"></tablecell></td>
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
  name: "EdmpDdTechnical",
  components: {
    PageHeader, pageSearch, pageExport, tableheader, tablecell, pageLoader
  },
  data() {
    return {
      storeName: "edmpddTechnical",
      edmpStoreName: "edmp"
    };
  },
  computed: {
    store() {
      return this.$store.state[this.storeName].all;
    },
    edmpStore() {
      return this.$store.state[this.edmpStoreName].all;
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

    setTimeout(() => {
      this.setTableColumnsWidth($('#table-edmp-dd-technical'));
    }, 300);

    $("#page-tab #tab-technical").on('click', function(){
      setTimeout(() => {
        self.setTableColumnsWidth($('#table-edmp-dd-technical'));
      }, 1);
    });
  },
  updated() {
    this.setTableColumnsWidth($('#table-edmp-dd-technical'));
  },
  methods: {
    getLeftTable() {
      this.store.system = this.$route.params.system;

      if( ! this.store.filters.left.filterTypes) this.store.filters.left.filterTypes = {};

      this.store.filters.left["COUNTRY"] = this.edmpStore.ddVal.ddCountrySelected;
      this.store.filters.left.filterTypes["COUNTRY"] = "eq";

      this.store.filters.left["BUSINESS_SEGMENT"] = this.edmpStore.ddVal.ddBusinessSegmentSelected;
      this.store.filters.left.filterTypes["BUSINESS_SEGMENT"] = "eq";

      this.store.filters.left["SOURCE_SYSTEM"] = this.edmpStore.ddVal.ddSourceSystemSelected;
      this.store.filters.left.filterTypes["SOURCE_SYSTEM"] = "eq";

      this.store.filters.left["CLUSTER_NAME"] = this.edmpStore.ddVal.ddClusterSelected;
      this.store.filters.left.filterTypes["CLUSTER_NAME"] = "eq";

      this.store.filters.left["TIER"] = this.edmpStore.ddVal.ddTierSelected;
      this.store.filters.left.filterTypes["TIER"] = "eq";

      this.store.filters.left["ITAM"] = this.edmpStore.ddVal.ddItamSelected;
      this.store.filters.left.filterTypes["ITAM"] = "eq";

      this.$store.dispatch(`${this.storeName}/getLeftTable`);
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
