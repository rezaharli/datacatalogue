<style>
@import '../../assets/styles/dashboard.css';
#table-dpo-my table.v-table tr th:nth-of-type(1){width: 55% !important;}
#table-dpo-my table.v-table tr th:nth-of-type(2){width: 45% !important;}
</style>

<template>
  <b-row style="margin-top: 10px;margin-bottom: 10px;">
    <b-col>
      <!-- Dpo details -->
      <router-view/>

      <!-- Main content -->
      <b-row>
        <b-col>
          <page-loader v-if="store.left.isLoading || (store.isRightTable && store.right.isLoading)" />

          <div class="card card-v1 transition">
            <div class="title-wrapper transition">
              <v-img :src="images.all" :contain="true" class="card-icon transition"></v-img>
              <h2 class="transition title-1">My Processes</h2>
            </div>
          
            <v-data-table
                :headers="store.leftHeaders.filter(v => v.display == true)"
                :items="store.left.display"
                :pagination.sync="store.left.pagination"
                :total-items="store.left.totalItems"
                :loading="store.left.isLoading"
                :expand="false"
                :must-sort="true"
                item-key="ID"
                class="card-content"
                id="table-dpo-my">

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
                  <td v-bind:style="{ width: store.left.colWidth['DSP_NAME'] + 'px' }"><b-link @click="showRightTable(props.item)"><tablecell :fulltext="props.item.DSP_NAME" showOn="hover"></tablecell></b-link></td>
                  <td v-bind:style="{ width: store.left.colWidth['DSP_OWNER'] + 'px' }">{{ props.item.DSP_OWNER }}</td>
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
        storeName: "dpomy",
        systemSource: [],
        tablenameSource: [],
        images: {
          all: require('../../assets/images/icon-my-system.png'),
        },
        hiddenFields: true
      }
    },
    computed: {
      store () { return this.$store.state[this.storeName].all },
      addressPath (){
        var tmp = this.$route.path.split("/")
        return tmp.slice(0, 2).join("/")
      },
      searchDropdownInputs () {
        return [
          { type: "text", label: "Process Name", source: "ProcessName" },
          { type: "text", label: "Process Owner", source: "ProcessOwner" },
          { type: "text", label: "CDE Name", source: "CDEName" },
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
    updated() {
      this.setTableColumnsWidth($('#table-dpo-my'));
    },
    methods: {
      getLeftTable () {
        var getLeftTableVal = this.$store.dispatch(`${this.storeName}/getLeftTable`);
        getLeftTableVal.then(res => {
          this.removeHypenOnEmptyTables($("#table-dpo-my"));
          this.setTableColumnsWidth($('#table-dpo-my'));
        });
      },
      doGetLeftTable () {
        this.getLeftTable();
      },
      showRightTable(param){
        this.$router.push(this.addressPath + '/' + encodeURIComponent(param.DSP_NAME));
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
      },
      removeHypenOnEmptyTables(elem){
        var paginationElem = elem.find('.v-datatable__actions .v-datatable__actions__range-controls .v-datatable__actions__pagination');
        paginationElem.each(function () {
          var paginationText = $(this).text();
          if(paginationText=="–"){
            $(this).hide();
          }else{
            $(this).show();
          }
        });
      }
    }
}
</script>
