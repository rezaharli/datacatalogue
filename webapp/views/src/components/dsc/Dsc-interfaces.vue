<style>
  @import '../../assets/styles/dashboard.css';
</style>

<style>
table.v-table thead th > div.btn-group {
  width: auto;
}

.header-filter-icon{

}
.header-filter-icon .dropdown-menu{
  overflow: scroll;
  height: 200px;
}
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
          
          <b-row>
            <b-col>
              <page-search :storeName="storeName" :searchDDInputs="searchDropdownInputs"/>
            </b-col>

            <b-col></b-col>
            <b-col></b-col>
            <b-col>
              <page-export :storeName="storeName" :leftTableCols="firstTableHeaders" :rightTableCols="secondTableHeaders" :forceRightAtFirstLevel="true" />
            </b-col>
          </b-row>

          <b-row>
            <b-col cols=6>
              <v-data-table
                  :headers="firstTableHeaders"
                  :items="store.left.display"
                  :pagination.sync="store.left.pagination"
                  :total-items="store.left.totalItems"
                  :loading="store.left.isLoading"
                  :expand="false"
                  item-key="ID"
                  class="elevation-1">

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
                    <td><b-link @click="showRightTable(props.item)"><tablecell :fulltext="props.item.SYSTEM_NAME" :isklik="false"></tablecell></b-link></td>
                    <td><b-link @click="props.expanded = !props.expanded"><tablecell :fulltext="(_.uniq(_.map(props.item.Custodians, 'ITAM_ID').filter(Boolean)).join(', '))" :isklik="true"></tablecell></b-link></td>
                    <td><tablecell :fulltext="(_.uniq(_.map(props.item.Custodians, 'DATASET_CUSTODIAN').filter(Boolean)).join('; '))" :isklik="true"></tablecell></td>
                    <td><tablecell :fulltext="(_.uniq(_.map(props.item.Custodians, 'BANK_ID').filter(Boolean)).join('; '))" :isklik="true"></tablecell></td>
                </template>

                <template slot="expand" slot-scope="props">
                  <v-data-table
                    :headers="firstTableHeaders"
                    :items="props.item.Custodians"
                    class="elevation-1"
                    hide-actions
                    hide-headers
                  >
                    <template slot="items" slot-scope="props">
                      <td style="width: 25%">&nbsp;</td>
                      <td style="width: 25%">&nbsp;</td>
                      <td style="width: 25%">{{ props.item.DATASET_CUSTODIAN }}</td>
                      <td style="width: 25%">{{ props.item.BANK_ID }}</td>
                    </template>
                  </v-data-table>
                </template>
              </v-data-table>
            </b-col>
            
            <b-col cols=6 class="scrollableasdf">
              <v-data-table
                  :headers="secondTableHeaders"
                  :items="store.right.display"
                  :pagination.sync="store.right.pagination"
                  :total-items="store.right.totalItems"
                  :loading="store.right.isLoading"
                  v-if="store.isRightTable"
                  item-key="R__"
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
                  <td><b-link :to="{ path:'/dsc/interfaces/' + props.item.TSID + '/' + props.item.ID + '/' + props.item.COLID }" href="#foo" v-b-modal.modallg>{{ props.item.LIST_OF_CDE }}</b-link></td>
                  <!-- <td><b-link :to="{ path:'/dsc/interfaces/' + route.params.system + "/details" }" v-b-modal.modallg>{{ props.item.name }}</b-link></td> -->
                  <td><tablecell :fulltext="(_.uniq(_.map(props.item.Values, 'IMM_PREC_SYSTEM_NAME').filter(Boolean)).join(', '))" :isklik="true"></tablecell></td>
                  <td><tablecell :fulltext="(_.uniq(_.map(props.item.Values, 'IMM_PREC_SYSTEM_SLA').filter(Boolean)).join(', '))" :isklik="true"></tablecell></td>
                  <td><tablecell :fulltext="(_.uniq(_.map(props.item.Values, 'IMM_PREC_SYSTEM_OLA').filter(Boolean)).join(', '))" :isklik="true"></tablecell></td>
                  <td><tablecell :fulltext="(_.uniq(_.map(props.item.Values, 'IMM_SUCC_SYSTEM_NAME').filter(Boolean)).join(', '))" :isklik="true"></tablecell></td>
                  <td><tablecell :fulltext="(_.uniq(_.map(props.item.Values, 'IMM_SUCC_SYSTEM_SLA').filter(Boolean)).join(', '))" :isklik="true"></tablecell></td>
                  <td><tablecell :fulltext="(_.uniq(_.map(props.item.Values, 'IMM_SUCC_SYSTEM_OLA').filter(Boolean)).join(', '))" :isklik="true"></tablecell></td>
                  <td><tablecell :fulltext="(_.uniq(_.map(props.item.Values, 'LIST_DOWNSTREAM_PROCESS').filter(Boolean)).join(', '))" :isklik="true"></tablecell></td>
                  <td><tablecell :fulltext="(_.uniq(_.map(props.item.Values, 'DOWNSTREAM_PROCESS_OWNER').filter(Boolean)).join(', '))" :isklik="true"></tablecell></td>
                </template>
              </v-data-table>
            </b-col>
          </b-row>
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
        storeName: "dscinterfaces",
        systemSource: [],
        tablenameSource: [],
        searchMain: '',
        firstTableHeaders: [
          { text: 'System Name', align: 'left', value: 'SYSTEM_NAME', sortable: false },
          { text: 'ITAM ID', align: 'left', value: 'Custodians.ITAM_ID', sortable: false },
          { text: 'Dataset Custodian', align: 'left', value: 'Custodians.DATASET_CUSTODIAN', sortable: false },
          { text: 'Bank ID', align: 'left', value: 'Custodians.BANK_ID', sortable: false }
        ],
        secondTableHeaders: [
          { text: 'List of CDEs', align: 'left', sortable: false, value: 'LIST_OF_CDE', displayCount: true, width: "25%" },
          { text: 'Immediate Preceding System', align: 'left', sortable: false, value: 'Values.IMM_PREC_SYSTEM_NAME', displayCount: true, width: "25%" },
          { text: 'SLA(Yes/No)', align: 'left', sortable: false, value: 'Values.IMM_PREC_SYSTEM_SLA', displayCount: false, width: "25%" },
          { text: 'OLA(Yes/No)', align: 'left', sortable: false, value: 'Values.IMM_PREC_SYSTEM_OLA', displayCount: false, width: "25%" },
          { text: 'Immediate Succeeding System', align: 'left', sortable: false, value: 'Values.IMM_SUCC_SYSTEM_NAME', displayCount: true, width: "25%" },
          { text: 'SLA (Yes/No)', align: 'left', sortable: false, value: 'Values.IMM_SUCC_SYSTEM_SLA', displayCount: false, width: "25%" },
          { text: 'OLA (Yes/No)', align: 'left', sortable: false, value: 'Values.IMM_SUCC_SYSTEM_OLA', displayCount: false, width: "25%" },
          { text: 'List of Downstream Process', align: 'left', sortable: false, value: 'Values.LIST_DOWNSTREAM_PROCESS', displayCount: true, width: "25%" },
          { text: 'Downstream Process Owner', align: 'left', sortable: false, value: 'Values.DOWNSTREAM_PROCESS_OWNER', displayCount: true, width: "25%" },
        ],
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
      showRightTable(param){
        //reset right table filter
        this.store.filters.right = {};

        this.$router.push(this.addressPath + '/' + param.ID);
      },
      systemRowClick (evt) {
        evt.preventDefault();
        this.store.isRightTable = true;
      },
    }
}
</script>
