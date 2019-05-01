<style>
  @import '../../assets/styles/dashboard.css';
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
                <td><tablecell :fulltext="props.item.DATA_DOMAIN" showOn="hover"></tablecell></td>
                <td><b-link @click="showRightTable(props.item)"><tablecell :fulltext="props.item.SUB_DOMAINS" showOn="hover"></tablecell></b-link></td>
                <td><tablecell :fulltext="props.item.SUB_DOMAIN_OWNER" showOn="click"></tablecell></td>
                <td><tablecell :fulltext="props.item.BANK_ID" showOn="click"></tablecell></td>
            </template>
            </v-data-table>

            <!-- <div class="card-circle transition"></div> -->
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
        storeName: "ddomy",
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
        return tmp.slice(0, 2).join("/")
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
      },
      "store.left.pagination": {
        handler () {
          this.doGetLeftTable();
        },
        deep: true
      },
      "store.right.pagination": {
        handler () {},
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
        this.$router.push(this.addressPath + '/' + param.SUB_DOMAINS);
      },
    }
}
</script>
