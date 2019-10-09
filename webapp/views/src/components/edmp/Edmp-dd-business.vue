<style>
/* #table-edmp-dd-business table.v-table tbody tr {display: block;} */
/* #table-edmp-dd-business table.v-table.v-datatable thead{
    width: unset;
    display: table-header-group;
    padding-right: unset;
}
#table-edmp-dd-business table.v-table.v-datatable tbody{
    display:table-row-group;
    overflow:auto;
    max-height:unset;
    width:unset;
}
#table-edmp-dd-business table.v-table.v-datatable tbody tr {display: table-row;} */
#table-edmp-dd-business table.v-table tr th:nth-of-type(1){width: calc(100%/20) !important; display: table-cell;}
#table-edmp-dd-business table.v-table tr th:nth-of-type(2){width: calc(100%/20) !important; display: table-cell;}
#table-edmp-dd-business table.v-table tr th:nth-of-type(3){width: calc(100%/20) !important; display: table-cell;}
#table-edmp-dd-business table.v-table tr th:nth-of-type(4){width: calc(100%/20) !important; display: table-cell;}
</style>

<template>
    <b-row style="margin-top: 10px;margin-bottom: 10px;">
      <b-col>
        <!-- Main content -->
        <div class="table-v2-title">Business Metadata</div>
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
            id="table-edmp-dd-business">

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

              <td v-bind:style="{ width: store.left.colWidth['TABLE_DESCRIPTION'] + 'px' }" class="text-capitalize">
                <tablecell showOn="hover" v-if="isMainLevelCellShowing(props)" :fulltext="props.item.TABLE_DESCRIPTION.toString().trim() ? props.item.TABLE_DESCRIPTION : 'NA'"></tablecell>
              </td>

              <td v-bind:style="{ width: store.left.colWidth['COLUMN_DESCRIPTION'] + 'px' }" class="text-capitalize">
                <tablecell showOn="hover" v-if="isMainLevelCellShowing(props)" :fulltext="props.item.COLUMN_DESCRIPTION.toString().trim() ? props.item.COLUMN_DESCRIPTION : 'NA'"></tablecell>
              </td>

              <td v-bind:style="{ width: store.left.colWidth['BUSINESS_TERM'] + 'px' }" class="text-capitalize">
                <tablecell showOn="hover" v-if="isMainLevelCellShowing(props)" :fulltext="props.item.BUSINESS_TERM.toString().trim() ? props.item.BUSINESS_TERM : 'NA'"></tablecell>
              </td>

              <td v-bind:style="{ width: store.left.colWidth['BUSINESS_DESCRIPTION'] + 'px' }" class="text-capitalize">
                <tablecell showOn="hover" v-if="isMainLevelCellShowing(props)" :fulltext="props.item.BUSINESS_DESCRIPTION.toString().trim() ? props.item.BUSINESS_DESCRIPTION : 'NA'"></tablecell>
              </td>

              <td v-bind:style="{ width: store.left.colWidth['DETERMINES_CLIENT_LOCATION'] + 'px' }" class="text-capitalize">
                <tablecell showOn="hover" v-if="isMainLevelCellShowing(props)" :fulltext="props.item.DETERMINES_CLIENT_LOCATION.toString().trim() ? props.item.DETERMINES_CLIENT_LOCATION : 'NA'"></tablecell>
              </td>

              <td v-bind:style="{ width: store.left.colWidth['DETERMINES_ACCOUNT'] + 'px' }" class="text-capitalize">
                <tablecell showOn="hover" v-if="isMainLevelCellShowing(props)" :fulltext="props.item.DETERMINES_ACCOUNT.toString().trim() ? props.item.DETERMINES_ACCOUNT : 'NA'"></tablecell>
              </td>

              <td v-bind:style="{ width: store.left.colWidth['BUSINESS_SEGMENT'] + 'px' }" class="text-capitalize">
                <tablecell showOn="hover" v-if="isMainLevelCellShowing(props)" :fulltext="props.item.BUSINESS_SEGMENT.toString().trim() ? props.item.BUSINESS_SEGMENT : 'NA'"></tablecell>
              </td>

              <td v-bind:style="{ width: store.left.colWidth['PRODUCT_CATEGORY'] + 'px' }" class="text-capitalize">
                <tablecell showOn="hover" v-if="isMainLevelCellShowing(props)" :fulltext="props.item.PRODUCT_CATEGORY.toString().trim() ? props.item.PRODUCT_CATEGORY : 'NA'"></tablecell>
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

              <td v-bind:style="{ width: store.left.colWidth['DOMAIN'] + 'px' }" class="text-capitalize">
                <tablecell showOn="hover" v-if="isMainLevelCellShowing(props)" :fulltext="props.item.DOMAIN.toString().trim() ? props.item.DOMAIN : 'NA'"></tablecell>
              </td>

              <td v-bind:style="{ width: store.left.colWidth['SUBDOMAIN'] + 'px' }" class="text-capitalize">
                <tablecell showOn="hover" v-if="isMainLevelCellShowing(props)" :fulltext="props.item.SUBDOMAIN.toString().trim() ? props.item.SUBDOMAIN : 'NA'"></tablecell>
              </td>

              <td v-bind:style="{ width: store.left.colWidth['DOMAIN_OWNER'] + 'px' }" class="text-capitalize">
                <tablecell showOn="hover" v-if="isMainLevelCellShowing(props)" :fulltext="props.item.DOMAIN_OWNER.toString().trim() ? props.item.DOMAIN_OWNER : 'NA'"></tablecell>
              </td>

              <td v-bind:style="{ width: store.left.colWidth['BUSINESS_TERM_DESCRIPTION'] + 'px' }" class="text-capitalize">
                <tablecell showOn="hover" v-if="isMainLevelCellShowing(props)" :fulltext="props.item.BUSINESS_TERM_DESCRIPTION.toString().trim() ? props.item.BUSINESS_TERM_DESCRIPTION : 'NA'"></tablecell>
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
                <td class="text-capitalize" v-bind:style="{ width: store.left.colWidth['Details'] + 'px' }">&nbsp;</td>
                <td class="text-capitalize" v-bind:style="{ width: store.left.colWidth['EDM_SOURCE_SYSTEM_NAME'] + 'px' }">&nbsp;</td>
                <td class="text-capitalize" v-bind:style="{ width: store.left.colWidth['DATABASE_NAME'] + 'px' }">&nbsp;</td>
                <td class="text-capitalize" v-bind:style="{ width: store.left.colWidth['TABLE_NAME'] + 'px' }">&nbsp;</td>

                <td class="text-capitalize" v-bind:style="{ width: store.left.colWidth['COLUMN_NAME'] + 'px' }">
                  <tablecell :fulltext="props.item.COLUMN_NAME.toString().trim() ? props.item.COLUMN_NAME : 'NA'" showOn="hover"></tablecell>
                </td>
                <td class="text-capitalize" v-bind:style="{ width: store.left.colWidth['TABLE_DESCRIPTION'] + 'px' }">
                  <tablecell :fulltext="props.item.TABLE_DESCRIPTION.toString().trim() ? props.item.TABLE_DESCRIPTION : 'NA'" showOn="hover"></tablecell>
                </td>
                <td class="text-capitalize" v-bind:style="{ width: store.left.colWidth['COLUMN_DESCRIPTION'] + 'px' }">
                  <tablecell :fulltext="props.item.COLUMN_DESCRIPTION.toString().trim() ? props.item.COLUMN_DESCRIPTION : 'NA'" showOn="hover"></tablecell>
                </td>
                <td class="text-capitalize" v-bind:style="{ width: store.left.colWidth['BUSINESS_TERM'] + 'px' }">
                  <tablecell :fulltext="props.item.BUSINESS_TERM.toString().trim() ? props.item.BUSINESS_TERM : 'NA'" showOn="hover"></tablecell>
                </td>
                <td class="text-capitalize" v-bind:style="{ width: store.left.colWidth['BUSINESS_DESCRIPTION'] + 'px' }">
                  <tablecell :fulltext="props.item.BUSINESS_DESCRIPTION.toString().trim() ? props.item.BUSINESS_DESCRIPTION : 'NA'" showOn="hover"></tablecell>
                </td>
                <td class="text-capitalize" v-bind:style="{ width: store.left.colWidth['DETERMINES_CLIENT_LOCATION'] + 'px' }">
                  <tablecell :fulltext="props.item.DETERMINES_CLIENT_LOCATION.toString().trim() ? props.item.DETERMINES_CLIENT_LOCATION : 'NA'" showOn="hover"></tablecell>
                </td>
                <td class="text-capitalize" v-bind:style="{ width: store.left.colWidth['DETERMINES_ACCOUNT'] + 'px' }">
                  <tablecell :fulltext="props.item.DETERMINES_ACCOUNT.toString().trim() ? props.item.DETERMINES_ACCOUNT : 'NA'" showOn="hover"></tablecell>
                </td>
                <td class="text-capitalize" v-bind:style="{ width: store.left.colWidth['BUSINESS_SEGMENT'] + 'px' }">
                  <tablecell :fulltext="props.item.BUSINESS_SEGMENT.toString().trim() ? props.item.BUSINESS_SEGMENT : 'NA'" showOn="hover"></tablecell>
                </td>
                <td class="text-capitalize" v-bind:style="{ width: store.left.colWidth['PRODUCT_CATEGORY'] + 'px' }">
                  <tablecell :fulltext="props.item.PRODUCT_CATEGORY.toString().trim() ? props.item.PRODUCT_CATEGORY : 'NA'" showOn="hover"></tablecell>
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
                <td class="text-capitalize" v-bind:style="{ width: store.left.colWidth['DOMAIN'] + 'px' }">
                  <tablecell :fulltext="props.item.DOMAIN.toString().trim() ? props.item.DOMAIN : 'NA'" showOn="hover"></tablecell>
                </td>
                <td class="text-capitalize" v-bind:style="{ width: store.left.colWidth['SUBDOMAIN'] + 'px' }">
                  <tablecell :fulltext="props.item.SUBDOMAIN.toString().trim() ? props.item.SUBDOMAIN : 'NA'" showOn="hover"></tablecell>
                </td>
                <td class="text-capitalize" v-bind:style="{ width: store.left.colWidth['DOMAIN_OWNER'] + 'px' }">
                  <tablecell :fulltext="props.item.DOMAIN_OWNER.toString().trim() ? props.item.DOMAIN_OWNER : 'NA'" showOn="hover"></tablecell>
                </td>
                <td class="text-capitalize" v-bind:style="{ width: store.left.colWidth['BUSINESS_TERM_DESCRIPTION'] + 'px' }">
                  <tablecell :fulltext="props.item.BUSINESS_TERM_DESCRIPTION.toString().trim() ? props.item.BUSINESS_TERM_DESCRIPTION : 'NA'" showOn="hover"></tablecell>
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
  name: "EdmpDdBusiness",
  components: {
    PageHeader, pageSearch, pageExport, tableheader, tablecell, pageLoader
  },
  data() {
    return {
      storeName: "edmpddBusiness",
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
      this.setTableColumnsWidth($('#table-edmp-dd-business'));
    }, 300);

    $("#page-tab #tab-business").on('click', function(){
      setTimeout(() => {
        self.setTableColumnsWidth($('#table-edmp-dd-business'));
      }, 1);
    });
  },
  updated() {
    this.setTableColumnsWidth($('#table-edmp-dd-business'));
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

  }
};
</script>
