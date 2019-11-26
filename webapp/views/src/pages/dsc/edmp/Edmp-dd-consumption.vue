<template>
    <b-row style="margin-top: 10px;margin-bottom: 10px;">
      <b-col>
        <!-- Main content -->
        <div class="table-v2-title">Consumption Applications</div>

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
            id="table-edmp-dd-consumption">

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
                  
              <td v-bind:style="{ width: store.left.colWidth['EDM_SOURCE_SYSTEM_NAME'] + 'px' }">
                <tablecell :fulltext="props.item.EDM_SOURCE_SYSTEM_NAME" showOn="click"></tablecell></td>
                  
              <td v-bind:style="{ width: store.left.colWidth['DATABASE_NAME'] + 'px' }">
                <tablecell :fulltext="props.item.DATABASE_NAME" showOn="click"></tablecell></td>
              
              <td v-bind:style="{ width: store.left.colWidth['TABLE_NAME'] + 'px' }" class="text-capitalize">
                <b-link @click="props.expanded = !props.expanded" v-if="props.item.Tables.length > 0">
                  <span class="ini">{{ props.item.TABLE_NAME.toString().trim() ? props.item.TABLE_NAME : 'NA' }}</span>
                </b-link>

                <span class="ini" v-if="props.item.Tables.length < 1">{{ props.item.TABLE_NAME.toString().trim() ? props.item.TABLE_NAME : 'NA' }}</span>
              </td>

              <td v-bind:style="{ width: store.left.colWidth['COLUMN_NAME'] + 'px' }" class="text-capitalize">
                <span class="ini" v-if="isMainLevelCellShowing(props)">{{ props.item.COLUMN_NAME.toString().trim() ? props.item.COLUMN_NAME : 'NA' }}</span>
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
            
            <template v-if="props.item.Tables.length > 0 && props.expanded">
              <tr :key="row.TMTID" v-for="row in props.item.Tables">
                <td class="text-capitalize">&nbsp;</td>
                <td class="text-capitalize" v-bind:style="{ width: store.left.colWidth['EDM_SOURCE_SYSTEM_NAME'] + 'px' }">&nbsp;</td>
                <td class="text-capitalize" v-bind:style="{ width: store.left.colWidth['DATABASE_NAME'] + 'px' }">&nbsp;</td>
                <td class="text-capitalize" v-bind:style="{ width: store.left.colWidth['TABLE_NAME'] + 'px' }">&nbsp;</td>

                <td class="text-capitalize" v-bind:style="{ width: store.left.colWidth['COLUMN_NAME'] + 'px' }">
                  <span class="ini" v-if="isMainLevelCellShowing(props)">{{ row.COLUMN_NAME.toString().trim() ? row.COLUMN_NAME : 'NA' }}</span>
                </td>
                <td class="text-capitalize" v-bind:style="{ width: store.left.colWidth['CONSUMING_APPLICATION'] + 'px' }">
                  <tablecell :fulltext="row.CONSUMING_APPLICATION.toString().trim() ? row.CONSUMING_APPLICATION : 'NA'" showOn="hover"></tablecell>
                </td>
                <td class="text-capitalize" v-bind:style="{ width: store.left.colWidth['CONSUMING_APPLICATION_ITAM'] + 'px' }">
                  <tablecell :fulltext="row.CONSUMING_APPLICATION_ITAM.toString().trim() ? row.CONSUMING_APPLICATION_ITAM : 'NA'" showOn="hover"></tablecell>
                </td>
                <td class="text-capitalize" v-bind:style="{ width: store.left.colWidth['CONSUMING_APPLICATION_OWNER'] + 'px' }">
                  <tablecell :fulltext="row.CONSUMING_APPLICATION_OWNER.toString().trim() ? row.CONSUMING_APPLICATION_OWNER : 'NA'" showOn="hover"></tablecell>
                </td>
                <td class="text-capitalize" v-bind:style="{ width: store.left.colWidth['CONSUMER_DESCRIPTION'] + 'px' }">
                  <tablecell :fulltext="row.CONSUMER_DESCRIPTION.toString().trim() ? row.CONSUMER_DESCRIPTION : 'NA'" showOn="hover"></tablecell>
                </td>
                <td class="text-capitalize" v-bind:style="{ width: store.left.colWidth['TECH_CONTACT'] + 'px' }">
                  <tablecell :fulltext="row.TECH_CONTACT.toString().trim() ? row.TECH_CONTACT : 'NA'" showOn="hover"></tablecell>
                </td>
                <td class="text-capitalize" v-bind:style="{ width: store.left.colWidth['BUSINESS_OWNERSHIP'] + 'px' }">
                  <tablecell :fulltext="row.BUSINESS_OWNERSHIP.toString().trim() ? row.BUSINESS_OWNERSHIP : 'NA'" showOn="hover"></tablecell>
                </td>
                <td class="text-capitalize" v-bind:style="{ width: store.left.colWidth['ACCESS_ROLE'] + 'px' }">
                  <tablecell :fulltext="row.ACCESS_ROLE.toString().trim() ? row.ACCESS_ROLE : 'NA'" showOn="hover"></tablecell>
                </td>
                <td class="text-capitalize" v-bind:style="{ width: store.left.colWidth['ROLE_DESCRIPTION'] + 'px' }">
                  <tablecell :fulltext="row.ROLE_DESCRIPTION.toString().trim() ? row.ROLE_DESCRIPTION : 'NA'" showOn="hover"></tablecell>
                </td>
                <td class="text-capitalize" v-bind:style="{ width: store.left.colWidth['CONSUMING_TECH_METADATA'] + 'px' }">
                  <tablecell :fulltext="row.CONSUMING_TECH_METADATA.toString().trim() ? row.CONSUMING_TECH_METADATA : 'NA'" showOn="hover"></tablecell>
                </td>
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

import PageHeader from '@/components/PageHeader';

import JsonExcel from "vue-json-excel";
import pageSearch from "@/components/PageSearch.vue";
import pageExport from "@/components/PageExport.vue";
import tableheader from "./TableHeader.vue";
import tablecell from "@/components/Tablecell.vue";
import pageLoader from "@/components/PageLoader.vue";

Vue.component("downloadExcel", JsonExcel);

export default {
  name: "EdmpDdConsumption",
  components: {
    PageHeader, pageSearch, pageExport, tableheader, tablecell, pageLoader
  },
  data() {
    return {
      storeName: "edmpddConsumption",
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
  mounted() {},
  updated() {},
  methods: {
    getLeftTable() {
      this.store.system = this.$route.params.system;

      if( ! this.isGlobalFilterEmpty) {
        this.$store.dispatch(`${this.storeName}/getLeftTable`).then(v => { 
          this.edmpStore.dd.firstload = false;
        })
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
  }
};
</script>
