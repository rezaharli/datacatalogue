<template>
    <v-content class="mx-4 my-5">
        <b-container fluid>
            <PageHeader title="Risk Framework Owner View" />

            <page-loader v-if="store.left.isLoading" />
            
            <b-row>
                <b-col sm=12 md=3>
                    <div class="card card-v2 transition">
                        <h6 class="title-1">Principal Risk Type</h6>
                        <h3 class="title-2 text-capitalize">{{ store.left.display[0] ? store.left.display[0].PRINCIPAL_RISK : "" }}</h3>
                    </div>
                </b-col>
                
                <b-col sm=12 md=3>
                    <div class="card card-v2 transition">
                        <h6 class="title-1">Risk Sub Type</h6>
                        <h3 class="title-2 text-capitalize">{{ $route.params.type }}</h3>
                    </div>
                </b-col>

                <b-col sm=12 md=3>
                    <div class="card card-v2 transition">
                        <h6 class="title-1">Priority Reports</h6>
                        <h3 class="title-2 text-capitalize">{{ store.left.display[0] ? store.left.display[0].PR_COUNT : "" }}</h3>
                    </div>
                </b-col>
            </b-row>

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
                <!-- Dsc details -->
                <router-view/>

                <!-- Main content -->
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
                    id="table-rfo-priority">
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
                    <tr>
                      <td v-bind:style="{ width: store.left.colWidth['PRIORITY_REPORT'] + 'px' }" 
                          :rowspan="props.item.expanded ? props.item.rowspanAcuan : 1" class="text-capitalize text-title">
                        <b-link @click="props.item.expanded = !props.item.expanded">
                          <tablecell showOn="hover" :fulltext="props.item.PRIORITY_REPORT"></tablecell></b-link></td>

                      <td 
                          v-bind:style="{ width: store.left.colWidth['PRIORITY_REPORT_RATIONALE'] + 'px' }" 
                          :rowspan="props.item.expanded ? (props.item.rowspanAcuan / props.item.PRIORITY_REPORT_RATIONALEs.length) : 1"
                          class="text-description">
                        <tablecell showOn="hover" :fulltext="props.item.PRIORITY_REPORT_RATIONALE"></tablecell></td>

                      <td 
                          v-bind:style="{ width: store.left.colWidth['CRM_NAME'] + 'px' }"
                          :rowspan="props.item.expanded ? (props.item.rowspanAcuan / props.item.CRM_NAMEs.length) : 1"
                          class="text-uppercase">
                        <b-link @click="toggleHighlightCDEName(props.item, props.item)">
                          <tablecell showOn="hover" :fulltext="props.item.CRM_NAME"></tablecell></b-link></td>

                      <td 
                          v-bind:style="{ width: store.left.colWidth['CRM_RATIONALE'] + 'px' }"
                          :rowspan="props.item.expanded ? (props.item.rowspanAcuan / props.item.CRM_RATIONALEs.length) : 1"
                          class="text-uppercase">
                        <tablecell showOn="hover" :fulltext="props.item.CRM_RATIONALE"></tablecell></td>

                      <td 
                          v-bind:style="{ width: store.left.colWidth['CDE_NAME'] + 'px' }" 
                          v-bind:class="{ 'is-highlighted': highlightedCDENames.indexOf(props.item.CDE_NAME) != -1 }"
                          :rowspan="props.item.expanded ? (props.item.rowspanAcuan / props.item.CDE_NAMEs.length) : 1"
                          class="text-uppercase">
                        <tablecell showOn="hover" :fulltext="props.item.CDE_NAME"></tablecell></td>
                        
                      <td 
                          v-bind:style="{ width: store.left.colWidth['CDE_RATIONALE'] + 'px' }"
                          :rowspan="props.item.expanded ? (props.item.rowspanAcuan / props.item.CDE_NAMEs.length) : 1">
                        <tablecell showOn="hover" :fulltext="props.item.CDE_RATIONALE"></tablecell></td>
                    </tr>

                    <template v-if="props.item.expanded">
                      <tr :key="props.item.ID + '' + i" v-for="(item, i) in props.item.childrenRow">
                        <td
                            v-if="item.PRIORITY_REPORT_RATIONALE" 
                            :rowspan="props.item.rowspanAcuan / props.item.PRIORITY_REPORT_RATIONALEs.length"
                            v-bind:style="{ width: store.left.colWidth['PRIORITY_REPORT_RATIONALE'] + 'px' }" 
                            class="text-description">
                          <tablecell showOn="hover" :fulltext="item.PRIORITY_REPORT_RATIONALE"></tablecell></td>

                        <td 
                            v-if="item.CRM_NAME" 
                            :rowspan="props.item.rowspanAcuan / props.item.CRM_NAMEs.length"
                            v-bind:style="{ width: store.left.colWidth['CRM_NAME'] + 'px' }" 
                            class="text-uppercase">
                          <b-link @click="toggleHighlightCDEName(props.item, item)">
                            <tablecell showOn="hover" :fulltext="item.CRM_NAME"></tablecell></b-link></td>

                        <td 
                            v-if="item.CRM_RATIONALE" 
                            :rowspan="props.item.rowspanAcuan / props.item.CRM_RATIONALEs.length"
                            v-bind:style="{ width: store.left.colWidth['CRM_RATIONALE'] + 'px' }" 
                            class="text-uppercase">
                          <tablecell showOn="hover" :fulltext="item.CRM_RATIONALE"></tablecell></td>

                        <td 
                            v-if="item.CDE_NAME" 
                            :rowspan="props.item.rowspanAcuan / props.item.CDE_NAMEs.length"
                            v-bind:class="{ 'is-highlighted': highlightedCDENames.indexOf(item.CDE_NAME) != -1 }"
                            v-bind:style="{ width: store.left.colWidth['CDE_NAME'] + 'px' }" class="text-uppercase">
                          <tablecell showOn="hover" :fulltext="item.CDE_NAME"></tablecell></td>
                          
                        <td 
                            v-if="item.CDE_RATIONALE" 
                            :rowspan="props.item.rowspanAcuan / props.item.CDE_NAMEs.length"
                            v-bind:style="{ width: store.left.colWidth['CDE_RATIONALE'] + 'px' }">
                          <tablecell showOn="hover" :fulltext="item.CDE_RATIONALE"></tablecell></td>
                      </tr>
                    </template>
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
  name: "RfoPriority",
  components: {
    PageHeader, pageSearch, pageExport, tableheader, tablecell, pageLoader
  },
  data() {
    return {
      storeName: "rfopriority",
      highlightedCDENames: []
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
    this.store.system = this.$route.params.type;
    this.resetFilter();
  },
  methods: {
    getLeftTable() {
      this.$store.dispatch(`${this.storeName}/getLeftTable`);
    },
    toggleHighlightCDEName (propsItem, item) {
      var cdes = propsItem.CDE_NAMEs.filter(v => {
        return v.CRM_NAME == item.CRM_NAME;
      }).map(v => v.CDE_NAME);

      this.highlightedCDENames = cdes;
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
  }
};
</script>
