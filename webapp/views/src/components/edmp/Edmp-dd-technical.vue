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

.ini{
  max-width: 90%;
  word-break: break-word;
}

.transparent-tnya3{
  background-color: rgba(0, 0, 0, 0) !important;
  color: rgba(0, 0, 0, 0) !important;
}
</style>

<template>
    <b-row style="margin-top: 10px;margin-bottom: 10px;">
      <b-col>
        <!-- Main content -->
        <div class="table-v2-title">Technical Metadata</div>
        
        <v-data-table
            v-model="store.selected"
            select-all
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
            id="table-edmp-dd-technical"
            @update:pagination="setTableColumnsWidth">

          <template slot="headers" slot-scope="props">
            <tr>
              <th>
                <v-checkbox :input-value="props.all" :indeterminate="props.indeterminate" primary hide-details @click.stop="toggleAll"></v-checkbox>
              </th>

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
            <tr :class="{even: props.index % 2, odd: !(props.index % 2)}" :active="props.selected">
              <td>
                <v-checkbox :input-value="props.selected" primary hide-details @click="props.selected = !props.selected"></v-checkbox></td>

              <td v-bind:style="{ width: store.left.colWidth['Details'] + 'px' }" class="text-capitalize text-title">
                <v-btn small exact target="blank" :href="linkDataXray(props.item.DATA_XRAY)" :disabled="dataXrayDisabled(props.item.DATA_XRAY)" class="green-tosca-gradient icon-only">
                  <i class="fa fa-fw fa-external-link-alt"></i></v-btn></td>

              <td v-bind:style="{ width: store.left.colWidth['ITAM'] + 'px' }">
                <tablecell :fulltext="props.item.ITAM" showOn="click"></tablecell></td>

              <td v-bind:style="{ width: store.left.colWidth['EDM_SOURCE_SYSTEM_NAME'] + 'px' }">
                <tablecell :fulltext="props.item.EDM_SOURCE_SYSTEM_NAME" showOn="click"></tablecell></td>
              
              <td v-bind:style="{ width: store.left.colWidth['DATABASE_NAME'] + 'px' }">
                <tablecell :fulltext="props.item.DATABASE_NAME" showOn="click"></tablecell></td>
              
              <td v-bind:style="{ width: store.left.colWidth['CERTIFIED'] + 'px' }">
                <tablecell :fulltext="props.item.CERTIFIED" showOn="click"></tablecell></td>
              
              <td v-bind:style="{ width: store.left.colWidth['TABLE_NAME'] + 'px' }" class="text-capitalize">
                <b-link @click="props.expanded = !props.expanded" v-if="props.item.Tables.length > 0">
                  <span class="ini">{{ props.item.TABLE_NAME.toString().trim() ? props.item.TABLE_NAME : 'NA' }}</span>
                </b-link>

                <span class="ini" v-if="props.item.Tables.length < 1">{{ props.item.TABLE_NAME.toString().trim() ? props.item.TABLE_NAME : 'NA' }}</span>
              </td>

              <td v-bind:style="{ width: store.left.colWidth['COLUMN_NAME'] + 'px' }" class="text-capitalize">
                <span class="ini" v-if="isMainLevelCellShowing(props)">{{ props.item.COLUMN_NAME.toString().trim() ? props.item.COLUMN_NAME : 'NA' }}</span>
              </td>

              <td v-bind:style="{ width: store.left.colWidth['DATA_TYPE'] + 'px' }" class="text-capitalize">
                <tablecell showOn="hover" v-if="isMainLevelCellShowing(props)" :fulltext="props.item.DATA_TYPE.toString().trim() ? props.item.DATA_TYPE : 'NA'"></tablecell>
              </td>

              <td v-bind:style="{ width: store.left.colWidth['COLUMN_LENGTH'] + 'px' }" class="text-capitalize">
                <tablecell showOn="hover" v-if="isMainLevelCellShowing(props)" :fulltext="props.item.COLUMN_LENGTH.toString().trim() ? props.item.COLUMN_LENGTH : 'NA'"></tablecell>
              </td>

              <td v-bind:style="{ width: store.left.colWidth['NULLABLE'] + 'px' }" class="text-capitalize">
                <tablecell showOn="hover" v-if="isMainLevelCellShowing(props)" :fulltext="props.item.NULLABLE.toString().trim() ? props.item.NULLABLE : 'NA'"></tablecell>
              </td>

              <td v-bind:style="{ width: store.left.colWidth['PRIMARY_KEY'] + 'px' }" class="text-capitalize">
                <tablecell showOn="hover" v-if="isMainLevelCellShowing(props)" :fulltext="props.item.PRIMARY_KEY.toString().trim() ? props.item.PRIMARY_KEY : 'NA'"></tablecell>
              </td>

              <td v-bind:style="{ width: store.left.colWidth['DATA_LINEAGE'] + 'px' }" class="text-capitalize">
                <tablecell showOn="hover" v-if="isMainLevelCellShowing(props)" :fulltext="props.item.DATA_LINEAGE.toString().trim() ? props.item.DATA_LINEAGE : 'NA'"></tablecell>
              </td>

              <td v-bind:style="{ width: store.left.colWidth['BUSINESS_ALIAS_NAME'] + 'px' }" class="text-capitalize">
                <tablecell showOn="hover" v-if="isMainLevelCellShowing(props)" :fulltext="props.item.BUSINESS_ALIAS_NAME.toString().trim() ? props.item.BUSINESS_ALIAS_NAME : 'NA'"></tablecell>
              </td>

              <td v-bind:style="{ width: store.left.colWidth['BUSINESS_ALIAS_DESCRIPTION'] + 'px' }" class="text-capitalize">
                <tablecell showOn="hover" v-if="isMainLevelCellShowing(props)" :fulltext="props.item.BUSINESS_ALIAS_DESCRIPTION.toString().trim() ? props.item.BUSINESS_ALIAS_DESCRIPTION : 'NA'"></tablecell>
              </td>

              <td v-bind:style="{ width: store.left.colWidth['CDE'] + 'px' }" class="text-capitalize">
                <tablecell showOn="hover" v-if="isMainLevelCellShowing(props)" :fulltext="props.item.CDE.toString().trim() ? props.item.CDE : 'NA'"></tablecell>
              </td>

              <td v-bind:style="{ width: store.left.colWidth['EXAMPLE'] + 'px' }" class="text-capitalize">
                <tablecell showOn="hover" v-if="isMainLevelCellShowing(props)" :fulltext="props.item.EXAMPLE.toString().trim() ? props.item.EXAMPLE : 'NA'"></tablecell>
              </td>

              <td v-bind:style="{ width: store.left.colWidth['DERIVED'] + 'px' }" class="text-capitalize">
                <tablecell showOn="hover" v-if="isMainLevelCellShowing(props)" :fulltext="props.item.DERIVED.toString().trim() ? props.item.DERIVED : 'NA'"></tablecell>
              </td>

              <td v-bind:style="{ width: store.left.colWidth['DERIVATION_LOGIC'] + 'px' }" class="text-capitalize">
                <tablecell showOn="hover" v-if="isMainLevelCellShowing(props)" :fulltext="props.item.DERIVATION_LOGIC.toString().trim() ? props.item.DERIVATION_LOGIC : 'NA'"></tablecell>
              </td>

              <td v-bind:style="{ width: store.left.colWidth['SOURCED_FROM_UPSTREAM'] + 'px' }" class="text-capitalize">
                <tablecell showOn="hover" v-if="isMainLevelCellShowing(props)" :fulltext="props.item.SOURCED_FROM_UPSTREAM.toString().trim() ? props.item.SOURCED_FROM_UPSTREAM : 'NA'"></tablecell>
              </td>

              <td v-bind:style="{ width: store.left.colWidth['SYSTEM_CHECKS'] + 'px' }" class="text-capitalize">
                <tablecell showOn="hover" v-if="isMainLevelCellShowing(props)" :fulltext="props.item.SYSTEM_CHECKS.toString().trim() ? props.item.SYSTEM_CHECKS : 'NA'"></tablecell>
              </td>
            </tr>
          </template>

          <template slot="expand" slot-scope="props">
            <v-data-table
              :headers="displayedHeaders"
              :items="props.item.Tables"
              item-key="TMTID"
              class=""
              hide-actions
              hide-headers
              @update:pagination="setExpandedTableColumnsWidth"
            >
              <template slot="items" slot-scope="props">
                <td class="text-capitalize">&nbsp;</td>
                
                <td class="text-capitalize" v-bind:style="{ width: store.left.colWidth['Details'] + 'px' }">
                  <v-btn small exact disabled target="blank" :href="linkDataXray(props.item.DATA_XRAY)" class="icon-only" style="background-color: rgba(0, 0, 0, 0) !important; color: rgba(0, 0, 0, 0) !important;">
                    <i class="fa fa-fw fa-external-link-alt"></i></v-btn>
                    </td>
                <td class="text-capitalize" v-bind:style="{ width: store.left.colWidth['ITAM'] + 'px' }">&nbsp;</td>
                <td class="text-capitalize" v-bind:style="{ width: store.left.colWidth['EDM_SOURCE_SYSTEM_NAME'] + 'px' }">&nbsp;</td>
                <td class="text-capitalize" v-bind:style="{ width: store.left.colWidth['DATABASE_NAME'] + 'px' }">&nbsp;</td>
                <td class="text-capitalize" v-bind:style="{ width: store.left.colWidth['CERTIFIED'] + 'px' }">&nbsp;</td>
                <td class="text-capitalize" v-bind:style="{ width: store.left.colWidth['TABLE_NAME'] + 'px' }">&nbsp;</td>

                <td class="text-capitalize" v-bind:style="{ width: store.left.colWidth['COLUMN_NAME'] + 'px' }">
                  <span class="ini" v-if="isMainLevelCellShowing(props)">{{ props.item.COLUMN_NAME.toString().trim() ? props.item.COLUMN_NAME : 'NA' }}</span>
                </td>
                <td class="text-capitalize" v-bind:style="{ width: store.left.colWidth['DATA_TYPE'] + 'px' }">
                  <tablecell :fulltext="props.item.DATA_TYPE.toString().trim() ? props.item.DATA_TYPE : 'NA'" showOn="hover"></tablecell>
                </td>
                <td class="text-capitalize" v-bind:style="{ width: store.left.colWidth['COLUMN_LENGTH'] + 'px' }">
                  <tablecell :fulltext="props.item.COLUMN_LENGTH.toString().trim() ? props.item.COLUMN_LENGTH : 'NA'" showOn="hover"></tablecell>
                </td>
                <td class="text-capitalize" v-bind:style="{ width: store.left.colWidth['NULLABLE'] + 'px' }">
                  <tablecell :fulltext="props.item.NULLABLE.toString().trim() ? props.item.NULLABLE : 'NA'" showOn="hover"></tablecell>
                </td>
                <td class="text-capitalize" v-bind:style="{ width: store.left.colWidth['PRIMARY_KEY'] + 'px' }">
                  <tablecell :fulltext="props.item.PRIMARY_KEY.toString().trim() ? props.item.PRIMARY_KEY : 'NA'" showOn="hover"></tablecell>
                </td>
                <td class="text-capitalize" v-bind:style="{ width: store.left.colWidth['DATA_LINEAGE'] + 'px' }">
                  <tablecell :fulltext="props.item.DATA_LINEAGE.toString().trim() ? props.item.DATA_LINEAGE : 'NA'" showOn="hover"></tablecell>
                </td>
                <td class="text-capitalize" v-bind:style="{ width: store.left.colWidth['BUSINESS_ALIAS_NAME'] + 'px' }">
                  <tablecell :fulltext="props.item.BUSINESS_ALIAS_NAME.toString().trim() ? props.item.BUSINESS_ALIAS_NAME : 'NA'" showOn="hover"></tablecell>
                </td>
                <td class="text-capitalize" v-bind:style="{ width: store.left.colWidth['BUSINESS_ALIAS_DESCRIPTION'] + 'px' }">
                  <tablecell :fulltext="props.item.BUSINESS_ALIAS_DESCRIPTION.toString().trim() ? props.item.BUSINESS_ALIAS_DESCRIPTION : 'NA'" showOn="hover"></tablecell>
                </td>
                <td class="text-capitalize" v-bind:style="{ width: store.left.colWidth['CDE'] + 'px' }">
                  <tablecell :fulltext="props.item.CDE.toString().trim() ? props.item.CDE : 'NA'" showOn="hover"></tablecell>
                </td>
                <td class="text-capitalize" v-bind:style="{ width: store.left.colWidth['EXAMPLE'] + 'px' }">
                  <tablecell :fulltext="props.item.EXAMPLE.toString().trim() ? props.item.EXAMPLE : 'NA'" showOn="hover"></tablecell>
                </td>
                <td class="text-capitalize" v-bind:style="{ width: store.left.colWidth['DERIVED'] + 'px' }">
                  <tablecell :fulltext="props.item.DERIVED.toString().trim() ? props.item.DERIVED : 'NA'" showOn="hover"></tablecell>
                </td>
                <td class="text-capitalize" v-bind:style="{ width: store.left.colWidth['DERIVATION_LOGIC'] + 'px' }">
                  <tablecell :fulltext="props.item.DERIVATION_LOGIC.toString().trim() ? props.item.DERIVATION_LOGIC : 'NA'" showOn="hover"></tablecell>
                </td>
                <td class="text-capitalize" v-bind:style="{ width: store.left.colWidth['SOURCED_FROM_UPSTREAM'] + 'px' }">
                  <tablecell :fulltext="props.item.SOURCED_FROM_UPSTREAM.toString().trim() ? props.item.SOURCED_FROM_UPSTREAM : 'NA'" showOn="hover"></tablecell>
                </td>
                <td class="text-capitalize" v-bind:style="{ width: store.left.colWidth['SYSTEM_CHECKS'] + 'px' }">
                  <tablecell :fulltext="props.item.SYSTEM_CHECKS.toString().trim() ? props.item.SYSTEM_CHECKS : 'NA'" showOn="hover"></tablecell>
                </td>
              </template>
            </v-data-table>
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

    $("#page-tab #tab-technical").on('click', function(){
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

      this.store.filters.left["BUSINESS_SEGMENT"] = this.edmpStore.dd.ddVal.ddBusinessSegmentSelected;
      this.store.filters.left.filterTypes["BUSINESS_SEGMENT"] = "eq";

      this.store.filters.left["EDM_SOURCE_SYSTEM_NAME"] = this.edmpStore.dd.ddVal.ddSourceSystemSelected;
      this.store.filters.left.filterTypes["EDM_SOURCE_SYSTEM_NAME"] = "eq";

      this.store.filters.left["CLUSTER_NAME"] = this.edmpStore.dd.ddVal.ddClusterSelected;
      this.store.filters.left.filterTypes["CLUSTER_NAME"] = "eq";

      this.store.filters.left["TIER"] = this.edmpStore.dd.ddVal.ddTierSelected;
      this.store.filters.left.filterTypes["TIER"] = "eq";

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
      var elem = $('#table-edmp-dd-technical');
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
    linkDataXray(param){
      if(param.includes("https://")){
        return param;
      } else if (param.includes("http://")) {
        return param;
      } else {
        return "http://" + param;
      }
    },
    dataXrayDisabled(param){
      return param.trim() == "" ? true : false;
    },
  }
};
</script>
