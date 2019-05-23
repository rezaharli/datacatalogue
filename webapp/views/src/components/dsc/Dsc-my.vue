<style>
@import '../../assets/styles/dashboard.css';
#table-dsc-my table.v-table tr th:nth-of-type(1){width: 60% !important;}
#table-dsc-my table.v-table tr th:nth-of-type(2){width: 40% !important;}
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
          
          <!-- <b-row>
            <b-col>
              <page-search :storeName="storeName" :searchDDInputs="searchDropdownInputs"/>
            </b-col>

            <b-col></b-col>
            <b-col></b-col>
            <b-col>
              <page-export :storeName="storeName" :leftTableCols="store.leftHeaders" :rightTableCols="store.rightHeaders"/>
            </b-col>
          </b-row> -->

          
              <div class="card card-v1 transition">
                <div class="title-wrapper transition">
                  <v-img :src="images.my" :contain="true" class="card-icon transition"></v-img>
                  <h2 class="transition title-1">My System</h2>
                </div>

                <v-data-table
                    :headers="store.leftHeaders.filter(v => v.display == true)"
                    :items="store.left.display"
                    :pagination.sync="store.left.pagination"
                    :total-items="store.left.totalItems"
                    :loading="store.left.isLoading"
                    :must-sort="true"
                    item-key="ID"
                    class="card-content"
                    id="table-dsc-my">

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
                      <td v-bind:style="{ width: store.left.colWidth['SYSTEM_NAME'] + 'px' }"><b-link @click="showRightTable(props.item)"><tablecell :fulltext="props.item.SYSTEM_NAME" showOn="hover"></tablecell></b-link></td>
                      <td v-bind:style="{ width: store.left.colWidth['ITAM_ID'] + 'px' }">
                        <v-tooltip bottom slot="activator">
                          <tablecell slot="activator" :fulltext="(_.uniq(_.map(props.item.Custodians, 'ITAM_ID').filter(Boolean)).join(', '))" showOn="click"></tablecell>
                          <span>{{ (_.uniq(_.map(props.item.Custodians, 'DATASET_CUSTODIAN').filter(Boolean)).join('; ')) }} ,  {{ (_.uniq(_.map(props.item.Custodians, 'BANK_ID').filter(Boolean)).join('; ')) }}</span>
                        </v-tooltip>
                      </td>
                      <!-- <td><tablecell :fulltext="(_.uniq(_.map(props.item.Custodians, 'DATASET_CUSTODIAN').filter(Boolean)).join('; '))" showOn="click"></tablecell></td>
                      <td><tablecell :fulltext="(_.uniq(_.map(props.item.Custodians, 'BANK_ID').filter(Boolean)).join('; '))" showOn="click"></tablecell></td> -->
                  </template>
                </v-data-table>

                <!-- <div class="card-circle transition"></div> -->
              </div>
            
            <!-- <b-col cols="6">
              <v-data-table
                  :headers="store.rightHeaders"
                  :items="store.right.display"
                  :pagination.sync="store.right.pagination"
                  :total-items="store.right.totalItems"
                  :loading="store.right.isLoading"
                  :expand="false"
                  v-if="store.isRightTable"
                  item-key="ID"
                  class="elevation-1">
                <v-progress-linear slot="progress" color="blue" indeterminate></v-progress-linear>

                <template slot="no-data">
                  <v-alert :value="store.right.isLoading" type="info">
                    Please wait while data is loading
                  </v-alert>

                  <v-alert :value="!store.right.isLoading" type="error">
                    Sorry, nothing to display here
                  </v-alert>
                </template>

                <template slot="headerCell" slot-scope="props">
                  <tableheader :storeName="storeName" :props="props" :which="'right'" />
                </template>

                <template slot="items" slot-scope="props">
                  <tr>
                    <td><b-link @click="props.expanded = !props.expanded"><tablecell :fulltext="props.item.TABLE_NAME" showOn="hover"></tablecell></b-link></td>
                    <td><tablecell :fulltext="(_.map(props.item.Columns, 'COLUMN_NAME').filter(Boolean).join(', '))" showOn="click"></tablecell></td>
                    <td><tablecell :fulltext="(_.map(props.item.Columns, 'BUSINESS_ALIAS_NAME').filter(Boolean).join(', '))" showOn="click"></tablecell></td>
                    <td><tablecell :fulltext="getCDEConclusion(_.map(props.item.Columns, 'CDE_YES_NO'))" showOn="click"></tablecell></td>
                  </tr>
                </template>
                
                <template slot="expand" slot-scope="props">
                  <v-data-table
                    :headers="store.rightHeaders"
                    :items="props.item.Columns"
                    class="elevation-1"
                    hide-actions
                    hide-headers
                  >
                    <template slot="items" slot-scope="props">
                      <td style="width: 25%">&nbsp;</td>
                      <td style="width: 25%"><b-link @click="showDetails(props.item)">{{ props.item.COLUMN_NAME }}</b-link></td>
                      <td style="width: 25%">{{ props.item.BUSINESS_ALIAS_NAME }}</td>
                      <td style="width: 25%">{{ props.item.CDE_YES_NO }}</td>
                    </template>
                  </v-data-table>
                </template>
              </v-data-table>
            </b-col> -->
          
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
        storeName: "dscmy",
        systemSource: [],
        tablenameSource: [],
        images: {
          my: require('../../assets/images/icon-my-system.png'),
        }
      }
    },
    computed: {
      store () {
        return this.$store.state[this.storeName].all
      },
      addressPath (){
        var tmp = this.$route.path.split("/")
        return tmp.slice(0, 3).join("/")
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

        if(this.store.isRightTable){
          this.doGetRightTable(this.$route.params.system);
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
          if(this.store.isRightTable){
            this.doGetRightTable(this.$route.params.system);
          }
        },
        deep: true
      },
      "store.searchMain" (val, oldVal){
        if(val || oldVal) {
          this.doGetLeftTable();

          if(this.store.isRightTable){
            this.doGetRightTable(this.$route.params.system);
          }
        }
      }
    },
    mounted() {
      this.store.tabName = this.storeName;
      this.store.isRightTable = this.$route.params.system;
      setTimeout(() => {
        this.setTableColumnsWidth($("#table-dsc-my"));
      }, 300);
    },
    updated() {
      this.setTableColumnsWidth($("#table-dsc-my"));
    },
    methods: {
      getLeftTable () {
        this.$store.dispatch(`${this.storeName}/getLeftTable`)
      },
      getRightTable (id) {
        this.$store.dispatch(`${this.storeName}/getRightTable`, id)
      },
      doGetLeftTable () {
        this.getLeftTable();
      },
      doGetRightTable (id) {
        this.getRightTable(id);
      },
      systemRowClick (evt) {
        evt.preventDefault();
        this.store.isRightTable = true;
      },
      getCDEConclusion (cdes) {
        return cdes.filter(Boolean).join(', ').indexOf("Yes") != -1 ? "Yes" : "No";
      },
      showRightTable(param){
        //reset right table filter
        this.store.filters.right = {};

        this.$router.push(this.addressPath + '/' + encodeURIComponent(param.SYSTEM_NAME));
      },
      showDetails (param) {
        this.$router.push(this.addressPath + "/" + encodeURIComponent(param.TSID) + '/' + encodeURIComponent(param.ID) + '/' + encodeURIComponent(param.COLID))
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
      }
    }
}
</script>
