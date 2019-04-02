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
      <!-- Ddo details -->
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
              <page-export :storeName="storeName" :leftTableCols="firstTableHeaders" :rightTableCols="secondTableHeaders"/>
            </b-col>
          </b-row>

          <b-row>
            <b-col cols="6">
              <v-data-table
                  :headers="firstTableHeaders"
                  :items="store.left.display"
                  :pagination.sync="store.left.pagination"
                  :total-items="store.left.totalItems"
                  :loading="store.left.isLoading"
                  item-key="ID"
                  class="elevation-1 ">

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
                  <td><tablecell :fulltext="props.item.DATA_DOMAIN" :isklik="false"></tablecell></td>
                  <td><b-link @click="showRightTable(props.item)"><tablecell :fulltext="props.item.SUB_DOMAINS" :isklik="false"></tablecell></b-link></td>
                  <td><tablecell :fulltext="props.item.SUB_DOMAIN_OWNER" :isklik="true"></tablecell></td>
                  <td><tablecell :fulltext="props.item.BANK_ID" :isklik="true"></tablecell></td>
                </template>
              </v-data-table>
            </b-col>
            
            <b-col cols="6">
              <v-data-table
                  :headers="secondTableHeaders"
                  :items="store.right.display"
                  :pagination.sync="store.right.pagination"
                  :total-items="store.right.totalItems"
                  :loading="store.right.isLoading"
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
                    <td><b-link @click="showDetails(props.item)"><tablecell :fulltext="props.item.BUSINESS_TERM" :isklik="false"></tablecell></b-link></td>
                    <td><tablecell :fulltext="props.item.BT_DESCRIPTION" :isklik="true"></tablecell></td>
                    <td><tablecell :fulltext="props.item.CDE_YES_NO" :isklik="true"></tablecell></td>
                  </tr>
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
        storeName: "ddomy",
        systemSource: [],
        tablenameSource: [],
        firstTableHeaders: [
          { text: 'Data Domain', align: 'left', value: 'DATA_DOMAIN', displayCount: true, sortable: false },
          { text: 'Sub Domains', align: 'left', value: 'SUB_DOMAINS', displayCount: true, sortable: false },
          { text: 'Sub Domain Owner', align: 'left', value: 'SUB_DOMAIN_OWNER', displayCount: true, sortable: false },
          { text: 'Bank ID', align: 'left', value: 'BANK_ID', displayCount: true, sortable: false },
        ],
        secondTableHeaders: [
          { text: 'Business Term', align: 'left', sortable: false, value: 'BUSINESS_TERM', displayCount: false, width: "25%" },
          { text: 'Business Term Description', align: 'left', sortable: false, value: 'BT_DESCRIPTION', displayCount: true, width: "25%" },
          { text: 'CDE (Yes/No)', align: 'left', sortable: false, value: 'CDE_YES_NO', displayCount: true, width: "25%" },
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
          { type: "text", label: "Data Domain", source: "DataDomain" },
          { type: "text", label: "Sub Data Domain", source: "SubDataDomain" },
          { type: "text", label: "Sub Data Domain Owner", source: "SubDataDomainOwner" },
          { type: "dropdown", label: "Business Term", source: "BusinessTerm", options: this._.map(this.store.right.source, 'BUSINESS_TERM') },
        ]
      }
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
      showRightTable(param){
        //reset right table filter
        this.store.filters.right = {};

        this.$router.push(this.addressPath + '/' + param.ID);
      },
      showDetails (param) {
        this.$router.push(this.addressPath + "/" + param.TSCID + '/' + param.ID)
      }
    }
}
</script>
