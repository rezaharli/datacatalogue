<template>
    <b-row style="margin-top: 10px;margin-bottom: 10px;">
      <b-col>
        <!-- Main content -->
        <div class="table-v2-title">Technical Metadata</div>

        <v-alert
          :value="isGlobalFilterEmpty"
          type="warning"
        >
          Choose the global filters for the data.
        </v-alert>
        
        <v-data-table
            v-model="store.selected"
            select-all
            v-if="!isGlobalFilterEmpty"
            :headers="displayedHeaders"
            :items="store.left.display"
            :pagination.sync="store.left.pagination"
            :total-items="store.left.totalItems"
            :loading="store.left.isLoading"
            :expand="false"
            :must-sort="true"
            :rows-per-page-items="[100]"
            item-key="ID"
            class="elevation-1 table-v2"
            id="table-edmp-dd-technical">

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
                  <div class="th-wrapper" v-bind:style="{ width: header.width ? header.width : 'unset' }">
                    <tableheader :storeName="storeName" :props="header" :which="'left'" :fromHeaderLoop="true" />
                    <v-icon small>arrow_upward</v-icon>
                  </div>
                </th>

                <th
                  v-if="header.sortable == false"
                  :key="header.text"
                  :class="['column sortable text-xs-left', store.left.pagination.descending ? 'desc' : 'asc', header.value === store.left.pagination.sortBy ? 'active' : '']"
                >
                  <div class="th-wrapper" v-bind:style="{ width: header.width ? header.width : 'unset' }">
                    <tableheader :storeName="storeName" :props="header" :which="'left'" :fromHeaderLoop="true" />
                  </div>
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
                  <span>{{ props.item.TABLE_NAME.toString().trim() ? props.item.TABLE_NAME : 'NA' }}</span>
                </b-link>

                <span v-if="props.item.Tables.length < 1">{{ props.item.TABLE_NAME.toString().trim() ? props.item.TABLE_NAME : 'NA' }}</span>
              </td>

              <td v-bind:style="{ width: store.left.colWidth['COLUMN_NAME'] + 'px' }" class="text-capitalize">
                <span v-if="isMainLevelCellShowing(props)">{{ props.item.COLUMN_NAME.toString().trim() ? props.item.COLUMN_NAME : 'NA' }}</span>
              </td>

              <td v-bind:style="{ width: store.left.colWidth['DATA_TYPE'] + 'px' }" class="text-capitalize">
                <tablecell showOn="hover" v-if="isMainLevelCellShowing(props)" :fulltext="props.item.DATA_TYPE.toString().trim() ? props.item.DATA_TYPE : 'NA'"></tablecell></td>

              <td v-bind:style="{ width: store.left.colWidth['COLUMN_LENGTH'] + 'px' }" class="text-capitalize">
                <tablecell showOn="hover" v-if="isMainLevelCellShowing(props)" :fulltext="props.item.COLUMN_LENGTH.toString().trim() ? props.item.COLUMN_LENGTH : 'NA'"></tablecell></td>

              <td v-bind:style="{ width: store.left.colWidth['NULLABLE'] + 'px' }" class="text-capitalize">
                <tablecell showOn="hover" v-if="isMainLevelCellShowing(props)" :fulltext="props.item.NULLABLE.toString().trim() ? props.item.NULLABLE : 'NA'"></tablecell>
              </td>

              <td v-bind:style="{ width: store.left.colWidth['PRIMARY_KEY'] + 'px' }" class="text-capitalize">
                  <tablecell showOn="hover" v-if="isMainLevelCellShowing(props)" :fulltext="props.item.PRIMARY_KEY.toString().trim() ? props.item.PRIMARY_KEY : 'NA'"></tablecell>
              </td>

              <td v-bind:style="{ width: store.left.colWidth['PII'] + 'px' }" class="text-capitalize">
                  <tablecell showOn="hover" v-if="isMainLevelCellShowing(props)" :fulltext="props.item.PII.toString().trim() ? props.item.PII : 'NA'"></tablecell>
              </td>

              <td v-bind:style="{ width: store.left.colWidth['DATA_LINEAGE'] + 'px' }" class="text-capitalize">
                <tablecell showOn="hover" v-if="isMainLevelCellShowing(props)" :fulltext="props.item.DATA_LINEAGE.toString().trim() ? props.item.DATA_LINEAGE : 'NA'"></tablecell>
              </td>

              <td v-bind:style="{ width: store.left.colWidth['CDE'] + 'px' }" class="text-capitalize">
                  <tablecell showOn="hover" v-if="isMainLevelCellShowing(props)" :fulltext="props.item.CDE.toString().trim() ? props.item.CDE : 'NA'"></tablecell>
              </td>
            </tr>
            
            <template v-if="props.item.Tables.length > 0 && props.expanded">
              <tr :key="row.TMTID" v-for="row in props.item.Tables">
                <td class="text-capitalize">&nbsp;</td>
                
                <td class="text-capitalize" v-bind:style="{ width: store.left.colWidth['Details'] + 'px' }"></td>

                <td class="text-capitalize" v-bind:style="{ width: store.left.colWidth['ITAM'] + 'px' }">&nbsp;</td>
                <td class="text-capitalize" v-bind:style="{ width: store.left.colWidth['EDM_SOURCE_SYSTEM_NAME'] + 'px' }">&nbsp;</td>
                <td class="text-capitalize" v-bind:style="{ width: store.left.colWidth['DATABASE_NAME'] + 'px' }">&nbsp;</td>
                <td class="text-capitalize" v-bind:style="{ width: store.left.colWidth['CERTIFIED'] + 'px' }">&nbsp;</td>
                <td class="text-capitalize" v-bind:style="{ width: store.left.colWidth['TABLE_NAME'] + 'px' }">&nbsp;</td>

                <td class="text-capitalize" v-bind:style="{ width: store.left.colWidth['COLUMN_NAME'] + 'px' }">
                  <span>{{ row.COLUMN_NAME.toString().trim() ? row.COLUMN_NAME : 'NA' }}</span></td>

                <td class="text-capitalize" v-bind:style="{ width: store.left.colWidth['DATA_TYPE'] + 'px' }">
                    <tablecell :fulltext="row.DATA_TYPE.toString().trim() ? row.DATA_TYPE : 'NA'" showOn="hover"></tablecell></td>

                <td class="text-capitalize" v-bind:style="{ width: store.left.colWidth['COLUMN_LENGTH'] + 'px' }">
                    <tablecell :fulltext="row.COLUMN_LENGTH.toString().trim() ? row.COLUMN_LENGTH : 'NA'" showOn="hover"></tablecell></td>

                <td class="text-capitalize" v-bind:style="{ width: store.left.colWidth['NULLABLE'] + 'px' }">
                    <tablecell :fulltext="row.NULLABLE.toString().trim() ? row.NULLABLE : 'NA'" showOn="hover"></tablecell></td>

                <td class="text-capitalize" v-bind:style="{ width: store.left.colWidth['PRIMARY_KEY'] + 'px' }">
                    <tablecell :fulltext="row.PRIMARY_KEY.toString().trim() ? row.PRIMARY_KEY : 'NA'" showOn="hover"></tablecell></td>

                <td class="text-capitalize" v-bind:style="{ width: store.left.colWidth['PII'] + 'px' }">
                    <tablecell :fulltext="row.PII.toString().trim() ? row.PII : 'NA'" showOn="hover"></tablecell></td>

                <td class="text-capitalize" v-bind:style="{ width: store.left.colWidth['DATA_LINEAGE'] + 'px' }">
                  <tablecell :fulltext="row.DATA_LINEAGE.toString().trim() ? row.DATA_LINEAGE : 'NA'" showOn="hover"></tablecell></td>

                <td class="text-capitalize" v-bind:style="{ width: store.left.colWidth['CDE'] + 'px' }">
                    <tablecell :fulltext="row.CDE.toString().trim() ? row.CDE : 'NA'" showOn="hover"></tablecell></td>
              </tr>
            </template>
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
import tableheader from "./TableHeader.vue";
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
    isGlobalFilterEmpty() {
      return this.edmpStore.dd.ddVal.ddCountrySelected.length == 0
        && this.edmpStore.dd.ddVal.ddBusinessSegmentSelected.length == 0
        && this.edmpStore.dd.ddVal.ddSourceSystemSelected.length == 0
        && this.edmpStore.dd.ddVal.ddClusterSelected.length == 0
        && this.edmpStore.dd.ddVal.ddTierSelected.length == 0
        && this.edmpStore.dd.ddVal.ddItamSelected.length == 0;
    },
  },
  watch: {
    $route(to) {},
    "store.left.pagination": {
      handler() {
        if( ! this.edmpStore.dd.firstload) {
          this.getLeftTable();
        }
      },
      deep: true
    },
    "store.searchMain"(val, oldVal) {
      if (val || oldVal) {
        this.getLeftTable();
      }
    }
  },
  mounted() {},
  updated() {},
  methods: {
    getLeftTable() {
      this.store.system = this.$route.params.system;

      if( ! this.isGlobalFilterEmpty) {
        this.$store.dispatch(`${this.storeName}/getLeftTable`).then(v => {
          this.edmpStore.dd.firstload = false;
        }) 
      } else {
        this.store.left.isLoading = false;
      }
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
