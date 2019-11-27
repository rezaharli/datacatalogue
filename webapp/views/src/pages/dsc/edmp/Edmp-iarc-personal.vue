<style scoped>
.table-v2 table.v-table tbody tr td:first-of-type{
  padding-left: 24px !important;
}
</style>

<template>
    <b-row style="margin-top: 10px;margin-bottom: 10px;">
      <b-col>
        <!-- Main content -->
        <div class="table-v2-title">Personal Data</div>

        <v-alert
          :value="isGlobalFilterEmpty"
          type="warning"
        >
          Choose the global filters for the data.
        </v-alert>
        
        <v-data-table
            v-model="store.selected"
            v-if="!isGlobalFilterEmpty"
            :headers="displayedHeaders"
            :items="store.left.display"
            :pagination.sync="store.left.pagination"
            :total-items="store.left.totalItems"
            :loading="store.left.isLoading"
            :must-sort="true"
            :rows-per-page-items="[100]"
            item-key="ID"
            class="elevation-1 table-v2"
            id="table-edmp-iarc-personal">

          <template slot="headers" slot-scope="props">
            <tr>
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
            <v-alert :value="store.left.isLoading" type="info">Please wait while data is loading</v-alert>
            <v-alert :value="!store.left.isLoading" type="error">Sorry, nothing to display here</v-alert>
          </template>

          <template slot="items" slot-scope="props">
            <tr :class="{even: props.index % 2, odd: !(props.index % 2)}">
              <td><tablecell :fulltext="props.item.ITAM.toString().trim() ? props.item.ITAM : 'NA'" showOn="click"></tablecell></td>

              <td><tablecell :fulltext="props.item.EDM_SOURCE_SYSTEM_NAME.toString().trim() ? props.item.EDM_SOURCE_SYSTEM_NAME : 'NA'" showOn="click"></tablecell></td>
              
              <td><tablecell :fulltext="props.item.COUNTRY.toString().trim() ? props.item.COUNTRY : 'NA'" showOn="click"></tablecell></td>
              
              <td><tablecell :fulltext="props.item.DATABASE_NAME.toString().trim() ? props.item.DATABASE_NAME : 'NA'" showOn="click"></tablecell></td>
              
              <td>
                <div class="ini wrapper-showmore d-inline-block">
                  <span>{{ props.item.TABLE_NAME.toString().trim() ? props.item.TABLE_NAME : 'NA' }}</span>
                </div>
              </td>
              
              <td>
                <div class="ini wrapper-showmore d-inline-block">
                  <span>{{ props.item.COLUMN_NAME.toString().trim() ? props.item.COLUMN_NAME : 'NA' }}</span>
                </div>
              </td>
              
              <td><tablecell :fulltext="props.item.CDE.toString().trim() ? props.item.CDE : 'NA'" showOn="click"></tablecell></td>
                
              <td><tablecell :fulltext="props.item.PII.toString().trim() ? props.item.PII : 'NA'" showOn="click"></tablecell></td>
            </tr>
          </template>
        </v-data-table>
              
      </b-col>
    </b-row>
</template>

<script>
import Vue from "vue";
import { mapState, mapActions } from "vuex";

import PageHeader from '@/components/PageHeader';

import JsonExcel from "vue-json-excel";
import pageSearch from "@/components/PageSearch.vue";
import pageExport from "@/components/PageExport.vue";
import tableheader from "./TableHeader.vue";
import tablecell from "@/components/Tablecell.vue";
import pageLoader from "@/components/PageLoader.vue";

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
    isGlobalFilterEmpty() {
      return this.$store.getters[this.edmpStoreName + "/isIarcGlobalFilterEmpty"];
    },
  },
  watch: {
    $route(to) {},
    "store.left.pagination": {
      handler() {
        if( ! this.edmpStore.iarc.firstload) {
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
      this.edmpStore.dd.firstload = true;

      if( ! this.isGlobalFilterEmpty) {
        this.$store.dispatch(`${this.storeName}/getLeftTable`).then(v => {
          this.edmpStore.iarc.firstload = false;
        })
      } else {
        this.store.dd.firstload = false;
        this.store.left.isLoading = false;
      }
    },
    isMainLevelCellShowing (props){
      if( ! props.expanded) {
        return true;
      } else {
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
  }
};
</script>
