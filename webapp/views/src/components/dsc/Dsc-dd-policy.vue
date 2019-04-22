<template>
    <b-row style="margin-top: 10px;margin-bottom: 10px;">
      <b-col>
        <!-- Main content -->
        <div class="table-v2-title">Policy Related Information</div>
        <v-data-table
            :headers="store.leftHeaders.policy.filter(v => v.display == true)"
            :items="store.left.display"
            :pagination.sync="store.left.pagination"
            :total-items="store.left.totalItems"
            :loading="store.left.isLoading"
            :expand="false"
            item-key="ID"
            class="table-v2">
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
              <td v-bind:style="{ width: store.left.colWidth['INFORMATION_ASSET_NAMES'] + 'px' }">
                <tablecell :fulltext="props.item.INFORMATION_ASSET_NAMES" showOn="click"></tablecell></td>
              
              <td v-bind:style="{ width: store.left.colWidth['INFORMATION_ASSET_DESCRIPTION'] + 'px' }">
                <tablecell :fulltext="props.item.INFORMATION_ASSET_DESCRIPTION" showOn="click"></tablecell></td>
              
              <td v-bind:style="{ width: store.left.colWidth['CONFIDENTIALITY'] + 'px' }">
                <tablecell :fulltext="props.item.CONFIDENTIALITY" showOn="click"></tablecell></td>
              
              <td v-bind:style="{ width: store.left.colWidth['INTEGRITY'] + 'px' }">
                <tablecell :fulltext="props.item.INTEGRITY" showOn="click"></tablecell></td>
              
              <td v-bind:style="{ width: store.left.colWidth['AVAILABILITY'] + 'px' }">
                <tablecell :fulltext="props.item.AVAILABILITY" showOn="click"></tablecell></td>
              
              <td v-bind:style="{ width: store.left.colWidth['OVERALL_CIA_RATING'] + 'px' }">
                <tablecell :fulltext="props.item.OVERALL_CIA_RATING" showOn="click"></tablecell></td>
              
              <td v-bind:style="{ width: store.left.colWidth['RECORD_CATEGORIES'] + 'px' }">
                <tablecell :fulltext="props.item.RECORD_CATEGORIES" showOn="click"></tablecell></td>
              
              <td v-bind:style="{ width: store.left.colWidth['PII_FLAG'] + 'px' }">
                <tablecell :fulltext="props.item.PII_FLAG" showOn="click"></tablecell></td>
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
  name: "DscDdPolicy",
  components: {
    PageHeader, pageSearch, pageExport, tableheader, tablecell, pageLoader
  },
  data() {
    return {
      storeName: "dscdd",
    };
  },
  computed: {
    store() {
      return this.$store.state[this.storeName].all;
    },
  },
};
</script>
