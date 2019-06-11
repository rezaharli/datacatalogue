<style>
#table-access-usage table.v-table thead tr th:first-of-type,
#table-access-usage table.v-table tbody tr td:first-of-type{padding-left: 20px !important;}
#table-access-usage table.v-table thead tr th:last-of-type,
#table-access-usage table.v-table tbody tr td:last-of-type{padding-right: 20px;}
#table-access-usage table.v-table thead tr th,
#table-access-usage table.v-table tbody tr td {padding-left: 10px; padding-right: 10px; }
#table-access-usage table.v-table tr th:nth-of-type(1){width: 10% !important;}
#table-access-usage table.v-table tr th:nth-of-type(2){width: 12% !important;}
#table-access-usage table.v-table tr th:nth-of-type(3){width: 12% !important;}
#table-access-usage table.v-table tr th:nth-of-type(4){width: 12% !important;}
#table-access-usage table.v-table tr th:nth-of-type(5){width: 10% !important;}
#table-access-usage table.v-table tr th:nth-of-type(6){width: 10% !important;}
#table-access-usage table.v-table tr th:nth-of-type(7){width: 10% !important;}
/* #table-access-usage table.v-table.v-datatable thead{
    width: unset;
    display: table-header-group;
    padding-right: unset;
}
#table-access-usage table.v-table.v-datatable tbody{
    display:table-row-group;
    overflow:auto;
    max-height:unset;
    width:unset;
}
#table-access-usage table.v-table.v-datatable tbody tr {display: table-row;} */
</style>

<template>
    <v-content class="mx-4 my-0 py-0">
        <b-container fluid>
            <page-loader v-if="store.left.isLoading" />

            <!-- <b-row class="my-4">
                <b-col>
                    <b-button class="float-right red-neon icon-only ml-3" @click="resetFilter">
                        <i class="fa fa-fw fa-filter"></i>
                    </b-button>

                    <page-export class="float-right" :storeName="storeName" :leftTableCols="store.leftHeaders" :rightTableCols="[]"/>
                </b-col>
            </b-row> -->

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
                            id="table-access-usage">
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
                                <td v-bind:style="{ width: store.left.colWidth['USERNAME'] + 'px' }" class="text-capitalize text-title">
                                  <tablecell :fulltext="props.item.USERNAME" showOn="hover"></tablecell></td>

                                <td v-bind:style="{ width: store.left.colWidth['FULLNAME'] + 'px' }" class="text-capitalize text-title">
                                  <tablecell :fulltext="props.item.FULLNAME" showOn="hover"></tablecell></td>

                                <td v-bind:style="{ width: store.left.colWidth['ROLE'] + 'px' }" class="text-capitalize text-title">
                                  <tablecell :fulltext="props.item.ROLE" showOn="hover"></tablecell></td>

                                <td v-bind:style="{ width: store.left.colWidth['MODULE'] + 'px' }" class="text-capitalize text-title">
                                  <tablecell :fulltext="props.item.MODULE" showOn="hover"></tablecell></td>

                                <td v-bind:style="{ width: store.left.colWidth['DESCRIPTION'] + 'px' }" class="text-capitalize text-title">
                                  <tablecell :fulltext="props.item.DESCRIPTION" showOn="hover"></tablecell></td>

                                <td v-bind:style="{ width: store.left.colWidth['THEDATE'] + 'px' }" class="text-capitalize text-title">
                                  <tablecell :fulltext="props.item.THEDATE" showOn="hover"></tablecell></td>

                                <td v-bind:style="{ width: store.left.colWidth['TIME'] + 'px' }" class="text-capitalize text-title">
                                  <tablecell :fulltext="props.item.TIME" showOn="hover"></tablecell></td>

                                <td v-bind:style="{ width: store.left.colWidth['RESOURCEURL'] + 'px' }" class="text-capitalize text-title">
                                  <tablecell :fulltext="props.item.RESOURCEURL" showOn="hover"></tablecell></td>
                            </tr>
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
      storeName: "usersusage",
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
    this.resetFilter();
    setTimeout(() => {
      this.setTableColumnsWidth($('#table-access-usage'));
      }, 300);
  },
  updated() {
      this.setTableColumnsWidth($('#table-access-usage'));
  },
  methods: {
    getLeftTable() {
      this.$store.dispatch(`${this.storeName}/getLeftTable`);
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
    showCDEs(param) {
      this.$router.push(
        this.addressPath + "/" + encodeURIComponent(param.SYSTEM_NAME) + "/" + encodeURIComponent(param.DSP_NAME)
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
  }
};
</script>
