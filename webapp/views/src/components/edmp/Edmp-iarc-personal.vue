<style>
/* #table-edmp-iarc-personal table.v-table tbody tr {display: block;} */
/* #table-edmp-iarc-personal table.v-table.v-datatable thead{
    width: unset;
    display: table-header-group;
    padding-right: unset;
}
#table-edmp-iarc-personal table.v-table.v-datatable tbody{
    display:table-row-group;
    overflow:auto;
    max-height:unset;
    width:unset;
}
#table-edmp-iarc-personal table.v-table.v-datatable tbody tr {display: table-row;} */
#table-edmp-iarc-personal table.v-table tr th:nth-of-type(1){width: calc(100%/20) !important; display: table-cell;}
#table-edmp-iarc-personal table.v-table tr th:nth-of-type(2){width: calc(100%/20) !important; display: table-cell;}
#table-edmp-iarc-personal table.v-table tr th:nth-of-type(3){width: calc(100%/20) !important; display: table-cell;}
#table-edmp-iarc-personal table.v-table tr th:nth-of-type(4){width: calc(100%/20) !important; display: table-cell;}
#table-edmp-iarc-personal table.v-table tr th:nth-of-type(5){width: calc(100%/20) !important; display: table-cell;}
#table-edmp-iarc-personal table.v-table tr th:nth-of-type(6){width: calc(100%/20) !important; display: table-cell;}
#table-edmp-iarc-personal table.v-table tr th:nth-of-type(7){width: calc(100%/20) !important; display: table-cell;}
#table-edmp-iarc-personal table.v-table tr th:nth-of-type(8){width: calc(100%/20) !important; display: table-cell;}
#table-edmp-iarc-personal table.v-table tr th:nth-of-type(9){width: calc(100%/20) !important; display: table-cell;}
#table-edmp-iarc-personal table.v-table tr th:nth-of-type(10){width: calc(100%/20) !important; display: table-cell;}
#table-edmp-iarc-personal table.v-table tr th:nth-of-type(11){width: calc(100%/20) !important; display: table-cell;}
#table-edmp-iarc-personal table.v-table tr th:nth-of-type(12){width: calc(100%/20) !important; display: table-cell;}
#table-edmp-iarc-personal table.v-table tr th:nth-of-type(13){width: calc(100%/20) !important; display: table-cell;}
</style>

<template>
    <b-row style="margin-top: 10px;margin-bottom: 10px;">
      <b-col>
        <!-- Main content -->
        <div class="table-v2-title">Personal Data</div>
        
        <v-data-table
            :headers="displayedHeaders"
            :items="store.left.display"
            :pagination.sync="store.left.pagination"
            :total-items="store.left.totalItems"
            :loading="store.left.isLoading"
            :expand="false"
            :must-sort="true"
            :rows-per-page-items="[25, 50, 75, 100]"
            item-key="ID"
            class="elevation-1 table-v2"
            id="table-edmp-iarc-personal"
            @update:pagination="setTableColumnsWidth">

          <template slot="headers" slot-scope="props">
            <tr>
              <template v-for="header in props.headers">
                <th
                  v-if="header.sortable == true"
                  :key="header.text"
                  :class="['column sortable text-xs-left', store.left.pagination.descending ? 'desc' : 'asc', header.value === store.left.pagination.sortBy ? 'active' : '']"
                  @click="changeSort(header.value)"
                >
                  <tableheader :storeName="storeName" :props="header" :which="'left'" :fromHeaderLoop="true" />
                  <v-icon small>arrow_upward</v-icon>
                </th>

                <th
                  v-if="header.sortable == false"
                  :key="header.text"
                  :class="['column sortable text-xs-left', store.left.pagination.descending ? 'desc' : 'asc', header.value === store.left.pagination.sortBy ? 'active' : '']"
                >
                  <tableheader :storeName="storeName" :props="header" :which="'left'" :fromHeaderLoop="true" />
                </th>
              </template>
            </tr>
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
                  <td v-bind:style="{ width: store.left.colWidth['ITAM'] + 'px' }">
                    <tablecell :fulltext="props.item.ITAM.toString().trim() ? props.item.ITAM : 'NA'" showOn="click"></tablecell></td>

                  <td v-bind:style="{ width: store.left.colWidth['EDM_SOURCE_SYSTEM_NAME'] + 'px' }">
                    <tablecell :fulltext="props.item.EDM_SOURCE_SYSTEM_NAME.toString().trim() ? props.item.EDM_SOURCE_SYSTEM_NAME : 'NA'" showOn="click"></tablecell></td>
                  
                  <td v-bind:style="{ width: store.left.colWidth['COUNTRY'] + 'px' }">
                    <tablecell :fulltext="props.item.COUNTRY.toString().trim() ? props.item.COUNTRY : 'NA'" showOn="click"></tablecell></td>
                  
                  <td v-bind:style="{ width: store.left.colWidth['DATABASE_NAME'] + 'px' }">
                    <tablecell :fulltext="props.item.DATABASE_NAME.toString().trim() ? props.item.DATABASE_NAME : 'NA'" showOn="click"></tablecell></td>
                  
                  <td v-bind:style="{ width: store.left.colWidth['TABLE_NAME'] + 'px' }" class="text-capitalize">
                    <div class="ini wrapper-showmore d-inline-block">
                      <span>{{ props.item.TABLE_NAME.toString().trim() ? props.item.TABLE_NAME : 'NA' }}</span>
                    </div>
                  </td>
                  
                  <td v-bind:style="{ width: store.left.colWidth['COLUMN_NAME'] + 'px' }" class="text-capitalize">
                    <div class="ini wrapper-showmore d-inline-block">
                        <span>{{ props.item.COLUMN_NAME.toString().trim() ? props.item.COLUMN_NAME : 'NA' }}</span>
                      </div>
                  </td>
                  
                  <td v-bind:style="{ width: store.left.colWidth['BUSINESS_ALIAS_NAME'] + 'px' }">
                    <tablecell :fulltext="props.item.BUSINESS_ALIAS_NAME.toString().trim() ? props.item.BUSINESS_ALIAS_NAME : 'NA'" showOn="click"></tablecell></td>
                  
                  <td v-bind:style="{ width: store.left.colWidth['BUSINESS_ALIAS_DESCRIPTION'] + 'px' }">
                    <tablecell :fulltext="props.item.BUSINESS_ALIAS_DESCRIPTION.toString().trim() ? props.item.BUSINESS_ALIAS_DESCRIPTION : 'NA'" showOn="click"></tablecell></td>
                  
                  <td v-bind:style="{ width: store.left.colWidth['CDE'] + 'px' }">
                    <tablecell :fulltext="props.item.CDE.toString().trim() ? props.item.CDE : 'NA'" showOn="click"></tablecell></td>
                    
                  <td v-bind:style="{ width: store.left.colWidth['PII'] + 'px' }">
                    <tablecell :fulltext="props.item.PII.toString().trim() ? props.item.PII : 'NA'" showOn="click"></tablecell></td>
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
  name: "EdmpIarcPersonal",
  components: {
    PageHeader, pageSearch, pageExport, tableheader, tablecell, pageLoader
  },
  data() {
    return {
      storeName: "edmpIarcPersonal",
      edmpStoreName: "edmp",
    };
  },
  computed: {
    store() {
      return this.$store.state[this.storeName].all;
    },
    edmpStore() {
      return this.$store.state[this.edmpStoreName].all;
    },
    displayedHeaders() {
      return this.store.leftHeaders.filter(v => v.display == true);
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
      this.setTableColumnsWidth();
    }, 10);

    $("#page-tab #tab-personal").on('click', function(){
      setTimeout(() => {
        self.setTableColumnsWidth();
      }, 1);
    });
  },
  updated() {
    this.setTableColumnsWidth();
  },
  methods: {
    getLeftTable() {
      this.store.system = this.$route.params.system;

      if( ! this.store.filters.left.filterTypes) this.store.filters.left.filterTypes = {};

      this.store.filters.left["COUNTRY"] = this.edmpStore.dd.ddVal.ddCountrySelected;
      this.store.filters.left.filterTypes["COUNTRY"] = "eq";

      this.store.filters.left["EDM_SOURCE_SYSTEM_NAME"] = this.edmpStore.dd.ddVal.ddSourceSystemSelected;
      this.store.filters.left.filterTypes["EDM_SOURCE_SYSTEM_NAME"] = "eq";

      this.store.filters.left["ITAM"] = this.edmpStore.dd.ddVal.ddItamSelected;
      this.store.filters.left.filterTypes["ITAM"] = "eq";

      this.$store.dispatch(`${this.storeName}/getLeftTable`).then(v => { 
        setTimeout(() => {
          this.setTableColumnsWidth() 
        }, 10);
      })
    },
    isMainLevelCellShowing (props){
      if( ! props.expanded) return true;
      else {
        if(props.item.Tables.length > 0) {
          return true;
        }
        
        return false;
      }
    },
    setTableColumnsWidth(){
      var elem = $('#table-edmp-iarc-personal');
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
    },
    toggleAll () {
      if (this.store.selected.length) this.store.selected = []
      else this.store.selected = this.store.left.display.slice()
    },
    changeSort (column) {
      if (this.store.left.pagination.sortBy === column) {
        this.store.left.pagination.descending = !this.store.left.pagination.descending
      } else {
        this.store.left.pagination.sortBy = column
        this.store.left.pagination.descending = false
      }
    },
  }
};
</script>
