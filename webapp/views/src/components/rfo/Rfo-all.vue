<style>
  @import '../../assets/styles/dashboard.css';
</style>

<template>
  <b-row style="margin-top: 10px;margin-bottom: 10px;">
    <b-col>
      <!-- Dsc details -->
      <router-view/>

      <!-- Main content -->
      <b-row>
        <b-col>
          <page-loader v-if="store.left.isLoading || (store.isRightTable && store.right.isLoading)" />

              <div class="card card-v1 transition">
                <div class="title-wrapper transition">
                  <v-img :src="images.all" :contain="true" class="card-icon transition"></v-img>
                  <h2 class="transition title-1">All Risk Types</h2>
                  <b-link v-on:click="toggleFieldDisplay()" class="toggle-field-display">
                    <span v-if="hiddenFields">show all columns &raquo;</span>
                    <span v-else>&laquo; show less columns</span>
                  </b-link>
                </div>
              
                <v-data-table
                    :headers="store.leftHeaders.filter(v => v.display == true)"
                    :items="store.left.display"
                    :pagination.sync="store.left.pagination"
                    :total-items="store.left.totalItems"
                    :loading="store.left.isLoading"
                    :expand="false"
                    item-key="ID"
                    class="card-content">

                  <template slot="headerCell" slot-scope="props">
                    <tableheader :storeName="storeName" :props="props" :which="'left'" />
                  </template>

                  <v-progress-linear slot="progress" color="blue" indeterminate></v-progress-linear>

                  <template slot="no-data">
                    <v-alert :value="store.left.isLoading" type="info">
                      Please wait while data is loading
                    </v-alert>

                    <v-alert :value="!store.left.isLoading" type="error">
                      Sorry, nothing to display here
                    </v-alert>
                  </template>

                  <template slot="items" slot-scope="props">
                    <tr>
                      <td v-if="isDisplayed('PRINCIPAL_RISK_TYPES')" :rowspan="props.item.RISK_SUB_TYPEs.length + 1">
                        <tablecell :fulltext="props.item.PRINCIPAL_RISK_TYPES" showOn="hover"></tablecell></td>
                      
                      <td v-if="isDisplayed('RISK_SUB_TYPE')">
                        <b-link @click.stop="showRightTable(props.item)">
                          <tablecell :fulltext="props.item.RISK_SUB_TYPE" showOn="hover"></tablecell></b-link></td>
                      
                      <td v-if="isDisplayed('RISK_FRAMEWORK_OWNER')">
                        <tablecell :fulltext="props.item.RISK_FRAMEWORK_OWNER" showOn="click"></tablecell></td>
                      
                      <td v-if="isDisplayed('RISK_REPORTING_LEAD')">
                        <tablecell :fulltext="props.item.RISK_REPORTING_LEAD" showOn="click"></tablecell></td>
                      
                      <td v-if="isDisplayed('PR_COUNT')">
                        <b-link @click.stop="showRightTable(props.item)">
                          <tablecell :fulltext="props.item.PR_COUNT" showOn="click"></tablecell></b-link></td>
                      
                      <td v-if="isDisplayed('CRM_COUNT')">
                        <tablecell :fulltext="props.item.CRM_COUNT" showOn="click"></tablecell></td>
                      
                      <td v-if="isDisplayed('CDE_COUNT')">
                        <tablecell :fulltext="props.item.CDE_COUNT" showOn="click"></tablecell></td>
                    </tr>

                    <tr :key="props.item.ID + '' + i" v-for="(item, i) in props.item.RISK_SUB_TYPEs">
                      <td v-if="isDisplayed('RISK_SUB_TYPE')">
                        <b-link @click.stop="showRightTable(item)">
                          <tablecell :fulltext="item.RISK_SUB_TYPE" showOn="hover"></tablecell></b-link></td>
                      
                      <td v-if="isDisplayed('RISK_FRAMEWORK_OWNER')">
                        <tablecell :fulltext="item.RISK_FRAMEWORK_OWNER" showOn="click"></tablecell></td>
                      
                      <td v-if="isDisplayed('RISK_REPORTING_LEAD')">
                        <tablecell :fulltext="item.RISK_REPORTING_LEAD" showOn="click"></tablecell></td>
                      
                      <td v-if="isDisplayed('PR_COUNT')">
                        <tablecell :fulltext="item.PR_COUNT" showOn="click"></tablecell></td>
                      
                      <td v-if="isDisplayed('CRM_COUNT')">
                        <tablecell :fulltext="item.CRM_COUNT" showOn="click"></tablecell></td>
                      
                      <td v-if="isDisplayed('CDE_COUNT')">
                        <tablecell :fulltext="item.CDE_COUNT" showOn="click"></tablecell></td>
                    </tr>
                  </template>
                </v-data-table>
              </div>
        </b-col>
      </b-row>
    </b-col>
  </b-row>
</template>

<script>
import Vue from 'vue'
import { mapState, mapActions } from 'vuex'
import JsonExcel from 'vue-json-excel'
import pageSearch from '../PageSearch.vue'
import pageExport from '../PageExport.vue'
import tableheader from '../TableHeader.vue'
import tablecell from '../Tablecell.vue'
import pageLoader from '../PageLoader.vue'
 
Vue.component('downloadExcel', JsonExcel)

export default {
    components: {
      pageSearch, pageExport, tableheader, tablecell, pageLoader
    },
    data () {
      return {
        storeName: "rfoall",
        systemSource: [],
        tablenameSource: [],
        images: {
          all: require('../../assets/images/icon-all-system.png'),
        },
        hiddenFields: true
      }
    },
    computed: {
      store () { return this.$store.state[this.storeName].all },
      dscStore () { return this.$store.state.dsc.all },
      addressPath (){
        var tmp = this.$route.path.split("/")
        return tmp.slice(0, 2).join("/")
      },
      searchDropdownInputs () {
        return [
          { type: "text", label: "System Name", source: "SystemName" },
          { type: "text", label: "ITAM ID", source: "ItamID" },
          { type: "dropdown", label: "Table Name", source: "TableName", options: this._.map(this.store.right.source, 'TABLE_NAME') },
          { type: "dropdown", label: "Column Name", source: "ColumnName", options: this._.map(this._.flattenDeep(this._.map(this.store.right.source, 'Columns')), 'COLUMN_NAME') },
        ]
      },
    },
    watch: {
      $route (to){
        this.store.isRightTable = false;

        if (to.params != undefined) {
          this.store.isRightTable = to.params.system; 
        }
      },
      "store.left.pagination": {
        handler () {
          this.doGetLeftTable();
        },
        deep: true
      },
      "store.right.pagination": {
        handler () {
        },
        deep: true
      },
      "store.searchMain" (val, oldVal){
        if(val || oldVal) {
          this.doGetLeftTable();
        }
      }
    },
    mounted() {
      this.store.tabName = this.storeName;
    },
    methods: {
      getLeftTable () {
        this.$store.dispatch(`${this.storeName}/getLeftTable`)
      },
      doGetLeftTable () {
        this.getLeftTable();
      },
      showRightTable(param){
        this.$router.push(this.addressPath + '/' + param.RISK_SUB_TYPE);
      },
      toggleFieldDisplay(){
        var toggleFieldName = (name) => {
          var opts = this.store.leftHeaders.find(v => v.value == name);
          opts.display = !opts.display;
          this.hiddenFields = !this.hiddenFields; 
        }

        toggleFieldName('RISK_REPORTING_LEAD');
        toggleFieldName('PR_COUNT');
        toggleFieldName('CRM_COUNT');
        toggleFieldName('CDE_COUNT');
      },
      isDisplayed(name){
        var opts = this.store.leftHeaders.find(v => v.value == name);
        return opts ? opts.display : false;
      }
    }
}
</script>
