<template>
  <v-flex d-flex xs12>
    <v-layout row wrap>
      <!-- Main content -->

      <v-flex d-flex xs12>
        <v-layout row wrap justify-center class="table-v2-title">Technical Metadata</v-layout>
      </v-flex>

      <v-flex d-flex xs12>
        <v-layout row wrap>
          <v-alert :value="isGlobalFilterEmpty" type="warning">
            Choose the global filters for the data.
          </v-alert>
        </v-layout>
      </v-flex>
      

      <v-flex d-flex xs12>
        <v-layout row wrap style="overflow: auto">
          <v-data-table
              v-model="store.selected"
              select-all
              v-if="( ! isGlobalFilterEmpty) && edmpStore.dd.displayTable"
              :headers="displayedHeaders"
              :items="store.left.display"
              :pagination.sync="store.left.pagination"
              :total-items="store.left.totalItems"
              :loading="store.left.isLoading"
              :expand="false"
              :must-sort="true"
              :rows-per-page-items="[100]"
              item-key="ID"
              class="elevation-1 table-v2 fixed-header v-table__overflow"
              style="max-height: calc(100vh - 160px); backface-visibility: hidden;"
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
              <v-alert :value="store.left.isLoading" type="info">Please wait while data is loading</v-alert>
              <v-alert :value="!store.left.isLoading" type="error">Sorry, nothing to display here</v-alert>
            </template>

            <template slot="items" slot-scope="props">
              <tr :active="props.selected">
                <td>
                  <v-checkbox :input-value="props.selected" primary hide-details @click="props.selected = !props.selected"></v-checkbox></td>

                <td v-bind:style="{ width: store.left.colWidth['Details'] + 'px' }">
                  <v-btn small exact target="blank" :href="linkDataXray(props.item.DATA_XRAY)" :disabled="dataXrayDisabled(props.item.DATA_XRAY)" class="green-tosca-gradient icon-only">
                    <i class="fa fa-fw fa-external-link-alt"></i></v-btn></td>

                <td v-bind:style="{ width: store.left.colWidth['ITAM'] + 'px' }">
                  <tablecell :fulltext="props.item.ITAM.toString().trim() ? props.item.ITAM : 'NA'" showOn="click"></tablecell></td>

                <td v-bind:style="{ width: store.left.colWidth['EDM_SOURCE_SYSTEM_NAME'] + 'px' }">
                  <tablecell :fulltext="props.item.EDM_SOURCE_SYSTEM_NAME" showOn="click"></tablecell></td>
                
                <td v-bind:style="{ width: store.left.colWidth['DATABASE_NAME'] + 'px' }">
                  <tablecell :fulltext="props.item.DATABASE_NAME" showOn="click"></tablecell></td>
                
                <td v-bind:style="{ width: store.left.colWidth['CERTIFIED'] + 'px' }">
                  <tablecell :fulltext="props.item.CERTIFIED.toString().trim() ? props.item.CERTIFIED : 'NA'" showOn="click"></tablecell></td>
                
                <td v-bind:style="{ width: store.left.colWidth['TABLE_NAME'] + 'px' }">
                  <span>{{ props.item.TABLE_NAME.toString().trim() ? props.item.TABLE_NAME : 'NA' }}</span>
                </td>

                <td v-bind:style="{ width: store.left.colWidth['COLUMN_NAME'] + 'px' }">
                  <span>{{ props.item.COLUMN_NAME.toString().trim() ? props.item.COLUMN_NAME : 'NA' }}</span>
                </td>

                <td v-bind:style="{ width: store.left.colWidth['DATA_TYPE'] + 'px' }">
                  <tablecell showOn="hover" :fulltext="props.item.DATA_TYPE.toString().trim() ? props.item.DATA_TYPE : 'NA'"></tablecell></td>

                <td v-bind:style="{ width: store.left.colWidth['COLUMN_LENGTH'] + 'px' }">
                  <tablecell showOn="hover" :fulltext="props.item.COLUMN_LENGTH.toString().trim() ? props.item.COLUMN_LENGTH : 'NA'"></tablecell></td>

                <td v-bind:style="{ width: store.left.colWidth['NULLABLE'] + 'px' }">
                  <tablecell showOn="hover" :fulltext="props.item.NULLABLE.toString().trim() ? props.item.NULLABLE : 'NA'"></tablecell>
                </td>

                <td v-bind:style="{ width: store.left.colWidth['PRIMARY_KEY'] + 'px' }">
                    <tablecell showOn="hover" :fulltext="props.item.PRIMARY_KEY.toString().trim() ? props.item.PRIMARY_KEY : 'NA'"></tablecell>
                </td>

                <td v-bind:style="{ width: store.left.colWidth['PII'] + 'px' }">
                    <tablecell showOn="hover" :fulltext="props.item.PII.toString().trim() ? props.item.PII : 'NA'"></tablecell>
                </td>

                <td v-bind:style="{ width: store.left.colWidth['DATA_LINEAGE'] + 'px' }">
                  <tablecell showOn="hover" :fulltext="props.item.DATA_LINEAGE.toString().trim() ? props.item.DATA_LINEAGE : 'NA'"></tablecell>
                </td>

                <td v-bind:style="{ width: store.left.colWidth['CDE'] + 'px' }">
                  <tablecell showOn="hover" :fulltext="props.item.CDE.toString().trim() ? props.item.CDE : 'NA'"></tablecell>
                </td>
              </tr>
              
              <template>
                <tr :key="row.TMTID" v-for="row in props.item.Tables">
                  <td>&nbsp;</td>
                  
                  <td v-bind:style="{ width: store.left.colWidth['Details'] + 'px' }"></td>

                  <td v-bind:style="{ width: store.left.colWidth['ITAM'] + 'px' }">
                    <tablecell :fulltext="row.ITAM.toString().trim() ? row.ITAM : 'NA'" showOn="click"></tablecell></td>

                  <td v-bind:style="{ width: store.left.colWidth['EDM_SOURCE_SYSTEM_NAME'] + 'px' }">
                    <tablecell :fulltext="row.EDM_SOURCE_SYSTEM_NAME" showOn="click"></tablecell></td>

                  <td v-bind:style="{ width: store.left.colWidth['DATABASE_NAME'] + 'px' }">
                    <tablecell :fulltext="row.DATABASE_NAME" showOn="click"></tablecell></td>

                  <td v-bind:style="{ width: store.left.colWidth['CERTIFIED'] + 'px' }">
                    <tablecell :fulltext="row.CERTIFIED" showOn="click"></tablecell></td>

                  <td v-bind:style="{ width: store.left.colWidth['TABLE_NAME'] + 'px' }">
                    <span>{{ row.TABLE_NAME.toString().trim() ? row.TABLE_NAME : 'NA' }}</span></td>

                  <td v-bind:style="{ width: store.left.colWidth['COLUMN_NAME'] + 'px' }">
                    <span>{{ row.COLUMN_NAME.toString().trim() ? row.COLUMN_NAME : 'NA' }}</span></td>

                  <td v-bind:style="{ width: store.left.colWidth['DATA_TYPE'] + 'px' }">
                      <tablecell showOn="hover" :fulltext="row.DATA_TYPE.toString().trim() ? row.DATA_TYPE : 'NA'"></tablecell></td>

                  <td v-bind:style="{ width: store.left.colWidth['COLUMN_LENGTH'] + 'px' }">
                      <tablecell showOn="hover" :fulltext="row.COLUMN_LENGTH.toString().trim() ? row.COLUMN_LENGTH : 'NA'"></tablecell></td>

                  <td v-bind:style="{ width: store.left.colWidth['NULLABLE'] + 'px' }">
                      <tablecell showOn="hover" :fulltext="row.NULLABLE.toString().trim() ? row.NULLABLE : 'NA'"></tablecell></td>

                  <td v-bind:style="{ width: store.left.colWidth['PRIMARY_KEY'] + 'px' }">
                      <tablecell showOn="hover" :fulltext="row.PRIMARY_KEY.toString().trim() ? row.PRIMARY_KEY : 'NA'"></tablecell></td>

                  <td v-bind:style="{ width: store.left.colWidth['PII'] + 'px' }">
                      <tablecell showOn="hover" :fulltext="row.PII.toString().trim() ? row.PII : 'NA'"></tablecell></td>

                  <td v-bind:style="{ width: store.left.colWidth['DATA_LINEAGE'] + 'px' }">
                    <tablecell showOn="hover" :fulltext="row.DATA_LINEAGE.toString().trim() ? row.DATA_LINEAGE : 'NA'"></tablecell></td>

                  <td v-bind:style="{ width: store.left.colWidth['CDE'] + 'px' }">
                    <tablecell showOn="hover" :fulltext="row.CDE.toString().trim() ? row.CDE : 'NA'"></tablecell></td>
                </tr>
              </template>
            </template>
          </v-data-table>
        </v-layout>
      </v-flex>
    </v-layout>
  </v-flex>
</template>

<script>
import Vue from "vue";

import JsonExcel from "vue-json-excel";
import tableheader from "./TableHeader.vue";
import tablecell from "@/components/Tablecell.vue";

Vue.component("downloadExcel", JsonExcel);

export default {
  name: "EdmpDdTechnical",
  components: {
    tableheader, tablecell
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
        return this.$store.getters[this.edmpStoreName + "/isDdGlobalFilterEmpty"];
    },
  },
  watch: {
    $route() {},
    "store.left.pagination": {
      handler(val, oldVal) {
        if (Object.keys(oldVal).length > 0 || ( ! this.edmpStore.dd.isNewPage)) {
          this.getLeftTable();
        }
      },
      deep: true
    },
  },
  mounted() {},
  updated() {},
  methods: {
    getLeftTable() {
      this.store.system = this.$route.params.system;
      this.edmpStore.dd.firstload = true;

      if( ! this.isGlobalFilterEmpty) {
        this.$store.dispatch(`${this.storeName}/getLeftTable`).then(() => {
          this.edmpStore.dd.firstload = false;
        }) 
      } else {
        this.store.dd.firstload = false;
        this.store.left.isLoading = false;
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
