<style>
/* #table-edmp-dd-consumption table.v-table tbody tr {display: block;} */
/* #table-edmp-dd-consumption table.v-table.v-datatable thead{
    width: unset;
    display: table-header-group;
    padding-right: unset;
}
#table-edmp-dd-consumption table.v-table.v-datatable tbody{
    display:table-row-group;
    overflow:auto;
    max-height:unset;
    width:unset;
}
#table-edmp-dd-consumption table.v-table.v-datatable tbody tr {display: table-row;} */
#table-edmp-dd-consumption table.v-table tr th:nth-of-type(1){width: calc(100%/20) !important; display: table-cell;}
#table-edmp-dd-consumption table.v-table tr th:nth-of-type(2){width: calc(100%/20) !important; display: table-cell;}
#table-edmp-dd-consumption table.v-table tr th:nth-of-type(3){width: calc(100%/20) !important; display: table-cell;}
#table-edmp-dd-consumption table.v-table tr th:nth-of-type(4){width: calc(100%/20) !important; display: table-cell;}
#table-edmp-dd-consumption table.v-table tr th:nth-of-type(5){width: calc(100%/20) !important; display: table-cell;}
</style>

<template>
    <b-row style="margin-top: 10px;margin-bottom: 10px;">
      <b-col>
        <!-- Main content -->
        <div class="table-v2-title">Consumption Apps</div>
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
            id="table-edmp-dd-consumption">
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
              
              <td v-bind:style="{ width: store.left.colWidth['TABLE_NAME'] + 'px' }" class="text-capitalize text-title">
                <b-link @click="props.expanded = !props.expanded" v-if="props.item.Tables.length > 0">
                  <tablecell :fulltext="props.item.TABLE_NAME.toString().trim() ? props.item.TABLE_NAME : 'NA'" showOn="hover"></tablecell>
                </b-link>

                <tablecell :fulltext="props.item.TABLE_NAME.toString().trim() ? props.item.TABLE_NAME : 'NA'" showOn="hover" v-if="props.item.Tables.length < 1"></tablecell>
              </td>

              <td v-bind:style="{ width: store.left.colWidth['COLUMN_NAME'] + 'px' }" class="text-capitalize">
                <tablecell showOn="hover" v-if="isMainLevelCellShowing(props)" :fulltext="props.item.COLUMN_NAME.toString().trim() ? props.item.COLUMN_NAME : 'NA'"></tablecell>
              </td>

              <td v-bind:style="{ width: store.left.colWidth['CONSUMING_APPLICATION'] + 'px' }" class="text-capitalize">
                <tablecell showOn="hover" v-if="isMainLevelCellShowing(props)" :fulltext="props.item.CONSUMING_APPLICATION.toString().trim() ? props.item.CONSUMING_APPLICATION : 'NA'"></tablecell>
              </td>

              <td v-bind:style="{ width: store.left.colWidth['CONSUMING_APPLICATION_ITAM'] + 'px' }" class="text-capitalize">
                <tablecell showOn="hover" v-if="isMainLevelCellShowing(props)" :fulltext="props.item.CONSUMING_APPLICATION_ITAM.toString().trim() ? props.item.CONSUMING_APPLICATION_ITAM : 'NA'"></tablecell>
              </td>

              <td v-bind:style="{ width: store.left.colWidth['CONSUMING_APPLICATION_OWNER'] + 'px' }" class="text-capitalize">
                <tablecell showOn="hover" v-if="isMainLevelCellShowing(props)" :fulltext="props.item.CONSUMING_APPLICATION_OWNER.toString().trim() ? props.item.CONSUMING_APPLICATION_OWNER : 'NA'"></tablecell>
              </td>

              <td v-bind:style="{ width: store.left.colWidth['CONSUMER_DESCRIPTION'] + 'px' }" class="text-capitalize">
                <tablecell showOn="hover" v-if="isMainLevelCellShowing(props)" :fulltext="props.item.CONSUMER_DESCRIPTION.toString().trim() ? props.item.CONSUMER_DESCRIPTION : 'NA'"></tablecell>
              </td>

              <td v-bind:style="{ width: store.left.colWidth['TECH_CONTACT'] + 'px' }" class="text-capitalize">
                <tablecell showOn="hover" v-if="isMainLevelCellShowing(props)" :fulltext="props.item.TECH_CONTACT.toString().trim() ? props.item.TECH_CONTACT : 'NA'"></tablecell>
              </td>

              <td v-bind:style="{ width: store.left.colWidth['BUSINESS_OWNERSHIP'] + 'px' }" class="text-capitalize">
                <tablecell showOn="hover" v-if="isMainLevelCellShowing(props)" :fulltext="props.item.BUSINESS_OWNERSHIP.toString().trim() ? props.item.BUSINESS_OWNERSHIP : 'NA'"></tablecell>
              </td>

              <td v-bind:style="{ width: store.left.colWidth['ACCESS_ROLE'] + 'px' }" class="text-capitalize">
                <tablecell showOn="hover" v-if="isMainLevelCellShowing(props)" :fulltext="props.item.ACCESS_ROLE.toString().trim() ? props.item.ACCESS_ROLE : 'NA'"></tablecell>
              </td>

              <td v-bind:style="{ width: store.left.colWidth['ROLE_DESCRIPTION'] + 'px' }" class="text-capitalize">
                <tablecell showOn="hover" v-if="isMainLevelCellShowing(props)" :fulltext="props.item.ROLE_DESCRIPTION.toString().trim() ? props.item.ROLE_DESCRIPTION : 'NA'"></tablecell>
              </td>

              <td v-bind:style="{ width: store.left.colWidth['CONSUMING_TECH_METADATA'] + 'px' }" class="text-capitalize">
                <tablecell showOn="hover" v-if="isMainLevelCellShowing(props)" :fulltext="props.item.CONSUMING_TECH_METADATA.toString().trim() ? props.item.CONSUMING_TECH_METADATA : 'NA'"></tablecell>
              </td>
            </tr>
          </template>

          <template slot="expand" slot-scope="props">
            <v-data-table
              :headers="store.leftHeaders.filter(v => v.display == true)"
              :items="props.item.Tables"
              item-key="TMTID"
              class=""
              hide-actions
              hide-headers
              @update:pagination="setExpandedTableColumnsWidth"
            >
              <template slot="items" slot-scope="props">
                <td class="text-capitalize" v-bind:style="{ width: store.left.colWidth['Details'] + 'px' }">&nbsp;</td>
                <td class="text-capitalize" v-bind:style="{ width: store.left.colWidth['EDM_SOURCE_SYSTEM_NAME'] + 'px' }">&nbsp;</td>
                <td class="text-capitalize" v-bind:style="{ width: store.left.colWidth['DATABASE_NAME'] + 'px' }">&nbsp;</td>
                <td class="text-capitalize" v-bind:style="{ width: store.left.colWidth['TABLE_NAME'] + 'px' }">&nbsp;</td>

                <td class="text-capitalize" v-bind:style="{ width: store.left.colWidth['COLUMN_NAME'] + 'px' }">
                  <tablecell :fulltext="props.item.COLUMN_NAME.toString().trim() ? props.item.COLUMN_NAME : 'NA'" showOn="hover"></tablecell>
                </td>
                <td class="text-capitalize" v-bind:style="{ width: store.left.colWidth['CONSUMING_APPLICATION'] + 'px' }">
                  <tablecell :fulltext="props.item.CONSUMING_APPLICATION.toString().trim() ? props.item.CONSUMING_APPLICATION : 'NA'" showOn="hover"></tablecell>
                </td>
                <td class="text-capitalize" v-bind:style="{ width: store.left.colWidth['CONSUMING_APPLICATION_ITAM'] + 'px' }">
                  <tablecell :fulltext="props.item.CONSUMING_APPLICATION_ITAM.toString().trim() ? props.item.CONSUMING_APPLICATION_ITAM : 'NA'" showOn="hover"></tablecell>
                </td>
                <td class="text-capitalize" v-bind:style="{ width: store.left.colWidth['CONSUMING_APPLICATION_OWNER'] + 'px' }">
                  <tablecell :fulltext="props.item.CONSUMING_APPLICATION_OWNER.toString().trim() ? props.item.CONSUMING_APPLICATION_OWNER : 'NA'" showOn="hover"></tablecell>
                </td>
                <td class="text-capitalize" v-bind:style="{ width: store.left.colWidth['CONSUMER_DESCRIPTION'] + 'px' }">
                  <tablecell :fulltext="props.item.CONSUMER_DESCRIPTION.toString().trim() ? props.item.CONSUMER_DESCRIPTION : 'NA'" showOn="hover"></tablecell>
                </td>
                <td class="text-capitalize" v-bind:style="{ width: store.left.colWidth['TECH_CONTACT'] + 'px' }">
                  <tablecell :fulltext="props.item.TECH_CONTACT.toString().trim() ? props.item.TECH_CONTACT : 'NA'" showOn="hover"></tablecell>
                </td>
                <td class="text-capitalize" v-bind:style="{ width: store.left.colWidth['BUSINESS_OWNERSHIP'] + 'px' }">
                  <tablecell :fulltext="props.item.BUSINESS_OWNERSHIP.toString().trim() ? props.item.BUSINESS_OWNERSHIP : 'NA'" showOn="hover"></tablecell>
                </td>
                <td class="text-capitalize" v-bind:style="{ width: store.left.colWidth['ACCESS_ROLE'] + 'px' }">
                  <tablecell :fulltext="props.item.ACCESS_ROLE.toString().trim() ? props.item.ACCESS_ROLE : 'NA'" showOn="hover"></tablecell>
                </td>
                <td class="text-capitalize" v-bind:style="{ width: store.left.colWidth['ROLE_DESCRIPTION'] + 'px' }">
                  <tablecell :fulltext="props.item.ROLE_DESCRIPTION.toString().trim() ? props.item.ROLE_DESCRIPTION : 'NA'" showOn="hover"></tablecell>
                </td>
                <td class="text-capitalize" v-bind:style="{ width: store.left.colWidth['CONSUMING_TECH_METADATA'] + 'px' }">
                  <tablecell :fulltext="props.item.CONSUMING_TECH_METADATA.toString().trim() ? props.item.CONSUMING_TECH_METADATA : 'NA'" showOn="hover"></tablecell>
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
  name: "EdmpDdConsumption",
  components: {
    PageHeader, pageSearch, pageExport, tableheader, tablecell, pageLoader
  },
  data() {
    return {
      storeName: "edmpddConsumption",
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
      this.setTableColumnsWidth($('#table-edmp-dd-consumption'));
    }, 300);

    $("#page-tab #tab-consumption").on('click', function(){
      setTimeout(() => {
        self.setTableColumnsWidth($('#table-edmp-dd-consumption'));
      }, 1);
    });
  },
  updated() {
    this.setTableColumnsWidth($('#table-edmp-dd-consumption'));
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
    isMainLevelCellShowing (props){
      if( ! props.expanded) return true;
      else {
        if(props.item.Tables.length > 0) {
          return true;
        }
        
        return false;
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
