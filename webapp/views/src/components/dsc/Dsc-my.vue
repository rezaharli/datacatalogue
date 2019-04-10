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
          
          <!-- <b-row>
            <b-col>
              <page-search :storeName="storeName" :searchDDInputs="searchDropdownInputs"/>
            </b-col>

            <b-col></b-col>
            <b-col></b-col>
            <b-col>
              <page-export :storeName="storeName" :leftTableCols="firstTableHeaders" :rightTableCols="secondTableHeaders"/>
            </b-col>
          </b-row> -->

          <b-row>
            <b-col cols="12">
              <div class="card transition">
                <h2 class="transition">My System</h2>

                <v-data-table
                    :headers="firstTableHeaders"
                    :items="store.left.display"
                    :pagination.sync="store.left.pagination"
                    :total-items="store.left.totalItems"
                    :loading="store.left.isLoading"
                    item-key="ID"
                    class="card-content ">

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
                      <td>
                        <v-tooltip bottom slot="activator">
                          <tablecell slot="activator" :fulltext="(_.uniq(_.map(props.item.Custodians, 'ITAM_ID').filter(Boolean)).join(', '))" :isklik="true"></tablecell>
                          <span>{{ (_.uniq(_.map(props.item.Custodians, 'DATASET_CUSTODIAN').filter(Boolean)).join('; ')) }} ,  {{ (_.uniq(_.map(props.item.Custodians, 'BANK_ID').filter(Boolean)).join('; ')) }}</span>
                        </v-tooltip>
                      </td>
                      <!-- <td><tablecell :fulltext="(_.uniq(_.map(props.item.Custodians, 'DATASET_CUSTODIAN').filter(Boolean)).join('; '))" :isklik="true"></tablecell></td>
                      <td><tablecell :fulltext="(_.uniq(_.map(props.item.Custodians, 'BANK_ID').filter(Boolean)).join('; '))" :isklik="true"></tablecell></td> -->
                  </template>
                </v-data-table>

                <div class="card_circle transition"></div>
              </div>
            </b-col>
            
            <!-- <b-col cols="6">
              <v-data-table
                  :headers="secondTableHeaders"
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
                    <td><b-link @click="props.expanded = !props.expanded"><tablecell :fulltext="props.item.TABLE_NAME" :isklik="false"></tablecell></b-link></td>
                    <td><tablecell :fulltext="(_.map(props.item.Columns, 'COLUMN_NAME').filter(Boolean).join(', '))" :isklik="true"></tablecell></td>
                    <td><tablecell :fulltext="(_.map(props.item.Columns, 'BUSINESS_ALIAS_NAME').filter(Boolean).join(', '))" :isklik="true"></tablecell></td>
                    <td><tablecell :fulltext="getCDEConclusion(_.map(props.item.Columns, 'CDE_YES_NO'))" :isklik="true"></tablecell></td>
                  </tr>
                </template>
                
                <template slot="expand" slot-scope="props">
                  <v-data-table
                    :headers="secondTableHeaders"
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
        storeName: "dscmy",
        systemSource: [],
        tablenameSource: [],
        firstTableHeaders: [
          { text: 'System Name', align: 'left', value: 'SYSTEM_NAME', displayCount: false, sortable: false },
          { text: 'ITAM ID', align: 'left', value: 'Custodians.ITAM_ID', displayCount: false, sortable: false },
          // { text: 'Dataset Custodian', align: 'left', value: 'Custodians.DATASET_CUSTODIAN', displayCount: true, sortable: false },
          // { text: 'Bank ID', align: 'left', value: 'Custodians.BANK_ID', displayCount: true, sortable: false }
        ],
        // secondTableHeaders: [
        //   { text: 'Table Name', align: 'left', sortable: false, value: 'TABLE_NAME', displayCount: true, width: "25%" },
        //   { text: 'Column Name', align: 'left', sortable: false, value: 'Columns.COLUMN_NAME', displayCount: true, width: "25%" },
        //   { text: 'Business Alias Name', align: 'left', sortable: false, value: 'Columns.BUSINESS_ALIAS_NAME', displayCount: false, width: "25%" },
        //   { text: 'CDE (Yes/No)', align: 'left', sortable: false, value: 'Columns.CDE_YES_NO', displayCount: true, width: "25%" }
        // ],
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

        this.$router.push(this.addressPath + '/' + param.ID);
      },
      showDetails (param) {
        this.$router.push(this.addressPath + "/" + param.TSID + '/' + param.ID + '/' + param.COLID)
      }
    }
}
</script>
