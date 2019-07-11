<style>
  @import '../../assets/styles/dashboard.css';
  #table-rfo-all table.v-table.v-datatable tbody tr {display: table-row;}
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
                    :must-sort="true"
                    item-key="ID"
                    class="card-content"
                    id="table-rfo-all">

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
                      <td v-bind:style="{ width: store.left.colWidth['PRINCIPAL_RISK_TYPES'] + 'px' }" v-if="isDisplayed('PRINCIPAL_RISK_TYPES')" :rowspan="props.item.RISK_SUB_TYPEs.length + 1">
                        <tablecell :fulltext="props.item.PRINCIPAL_RISK_TYPES" showOn="hover"></tablecell></td>
                      
                      <td v-bind:style="{ width: store.left.colWidth['RISK_SUB_TYPE'] + 'px' }" v-if="isDisplayed('RISK_SUB_TYPE')">
                        <b-link @click.stop="showRightTable(props.item)">
                          <tablecell :fulltext="props.item.RISK_SUB_TYPE" showOn="hover"></tablecell></b-link></td>
                      
                      <td v-bind:style="{ width: store.left.colWidth['RISK_FRAMEWORK_OWNER'] + 'px' }" v-if="isDisplayed('RISK_FRAMEWORK_OWNER')">
                        <tablecell :fulltext="props.item.RISK_FRAMEWORK_OWNER" showOn="click"></tablecell></td>
                      
                      <td v-bind:style="{ width: store.left.colWidth['RISK_REPORTING_LEAD'] + 'px' }" v-if="isDisplayed('RISK_REPORTING_LEAD')">
                        <tablecell :fulltext="props.item.RISK_SUB_TYPEs[0] ? (props.item.RISK_SUB_TYPEs[0].RISK_SUB_TYPEsVal.map(v => v.RISK_REPORTING_LEAD).join(' | ')) : ''" showOn="click"></tablecell></td>
                      
                      <td v-bind:style="{ width: store.left.colWidth['PR_COUNT'] + 'px' }" v-if="isDisplayed('PR_COUNT')">
                        <b-link @click.stop="showRightTable(props.item)">
                          <tablecell :fulltext="props.item.RISK_SUB_TYPEs[0] ? (props.item.RISK_SUB_TYPEs[0].RISK_SUB_TYPEsVal[0].PR_COUNT) : ''" showOn="click"></tablecell></b-link></td>
                      
                      <td v-bind:style="{ width: store.left.colWidth['CRM_COUNT'] + 'px' }" v-if="isDisplayed('CRM_COUNT')">
                        <tablecell :fulltext="props.item.RISK_SUB_TYPEs[0] ? (props.item.RISK_SUB_TYPEs[0].RISK_SUB_TYPEsVal[0].CRM_COUNT) : ''" showOn="click"></tablecell></td>
                      
                      <td v-bind:style="{ width: store.left.colWidth['CDE_COUNT'] + 'px' }" v-if="isDisplayed('CDE_COUNT')">
                        <tablecell :fulltext="props.item.RISK_SUB_TYPEs[0] ? (props.item.RISK_SUB_TYPEs[0].RISK_SUB_TYPEsVal[0].CDE_COUNT) : ''" showOn="click"></tablecell></td>
                    </tr>

                    <tr :key="props.item.ID + '' + i" v-for="(item, i) in props.item.RISK_SUB_TYPEs">
                      <template v-if="i != 0">
                        <td v-bind:style="{ width: store.left.colWidth['RISK_SUB_TYPE'] + 'px' }" v-if="isDisplayed('RISK_SUB_TYPE')">
                          <b-link @click.stop="showRightTable(item)">
                            <tablecell :fulltext="item.RISK_SUB_TYPE" showOn="hover"></tablecell></b-link></td>
                        
                        <td v-bind:style="{ width: store.left.colWidth['RISK_FRAMEWORK_OWNER'] + 'px' }" v-if="isDisplayed('RISK_FRAMEWORK_OWNER')">
                          <tablecell :fulltext="item.RISK_FRAMEWORK_OWNER" showOn="click"></tablecell></td>
                        
                        <td v-bind:style="{ width: store.left.colWidth['RISK_REPORTING_LEAD'] + 'px' }" v-if="isDisplayed('RISK_REPORTING_LEAD')">
                          <tablecell :fulltext="item.RISK_SUB_TYPEsVal.map(v => v.RISK_REPORTING_LEAD).join(' | ')" showOn="click"></tablecell></td>
                        
                        <td v-bind:style="{ width: store.left.colWidth['PR_COUNT'] + 'px' }" v-if="isDisplayed('PR_COUNT')">
                          <b-link @click.stop="showRightTable(item)">
                            <tablecell :fulltext="props.item.RISK_SUB_TYPEs[0] ? (props.item.RISK_SUB_TYPEs[i].RISK_SUB_TYPEsVal[0].PR_COUNT) : ''" showOn="click"></tablecell></b-link></td>
                        
                        <td v-bind:style="{ width: store.left.colWidth['CRM_COUNT'] + 'px' }" v-if="isDisplayed('CRM_COUNT')">
                          <tablecell :fulltext="props.item.RISK_SUB_TYPEs[0] ? (props.item.RISK_SUB_TYPEs[i].RISK_SUB_TYPEsVal[0].CRM_COUNT) : ''" showOn="click"></tablecell></td>
                        
                        <td v-bind:style="{ width: store.left.colWidth['CDE_COUNT'] + 'px' }" v-if="isDisplayed('CDE_COUNT')">
                          <tablecell :fulltext="props.item.RISK_SUB_TYPEs[0] ? (props.item.RISK_SUB_TYPEs[i].RISK_SUB_TYPEsVal[0].CDE_COUNT) : ''" showOn="click"></tablecell></td>
                      </template>
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
      setTimeout(() => {
        this.makeTableAlertFull();
      }, 400);
      this.toggleFieldDisplay(false);
    },
    updated() {
      this.setTableColumnsWidthRowspan($('#table-rfo-all'));
    },
    methods: {
      getLeftTable () {
        var getLeftTableVal = this.$store.dispatch(`${this.storeName}/getLeftTable`);
        getLeftTableVal.then(res => {
          this.removeHypenOnEmptyTables($("#table-rfo-all"));
          this.setTableColumnsWidthRowspan($('#table-rfo-all'));
        });
      },
      doGetLeftTable () {
        this.getLeftTable();
      },
      showRightTable(param){
        this.$router.push(this.addressPath + '/' + encodeURIComponent(param.RISK_SUB_TYPE));
      },
      toggleFieldDisplay(expanded = true){
        var toggleFieldName = (name, expanded) => {
          var opts = this.store.leftHeaders.find(v => v.value == name);
          opts.display = !opts.display;
          if(expanded==false){
            opts.display = false;
          }
        }

        toggleFieldName('RISK_REPORTING_LEAD', expanded);
        toggleFieldName('PR_COUNT', expanded);
        toggleFieldName('CRM_COUNT', expanded);
        toggleFieldName('CDE_COUNT', expanded);
        
        this.hiddenFields = !this.hiddenFields; 
        if(expanded==false){
          this.hiddenFields = true;
        }
      },
      isDisplayed(name){
        var opts = this.store.leftHeaders.find(v => v.value == name);
        return opts ? opts.display : false;
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
      setTableColumnsWidthRowspan(elem){
        var tableElem = elem.find('.v-table__overflow > table.v-table');
        var theadTR = tableElem.find('thead tr:first');
        var THs = theadTR.find('th');
        var tbodyTR = tableElem.find('tbody tr');
        var thWidths = [];
        THs.each(function (thIndex) {
          thWidths[thIndex] = $(this).outerWidth();
        });
        tbodyTR.each(function (tdIndex) {
          var TDs = $(this).find('td');
          TDs.each(function (tdIndex) {
            if ($(this)[0].hasAttribute('rowspan')){
              $(this).closest('tr').addClass('has-rowspan');
            }
            var colWidth = thWidths[tdIndex];
            colWidth = colWidth - 47; // untuk mengurangi additional width yang datang tiba2 seperti syaiton, xixixi
            $(this).width(colWidth);
          });
        });

        //untuk menyesuaikan ulang ukuran kolom karena crash dengan rowspan
        setTimeout(() => {
          tbodyTR.not(".has-rowspan").each(function (tdIndex) {
            var TDs = $(this).find('td');
            TDs.each(function (tdIndex) {
              if ($(this)[0].hasAttribute('rowspan')){
                $(this).closest('tr').addClass('has-rowspan');
              }
              var colWidth = thWidths[tdIndex+1];
              colWidth = colWidth - 47; // untuk mengurangi additional width yang datang tiba2 seperti syaiton, xixixi
              $(this).width(colWidth);
            });
          });
        }, 1);
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
      },
      makeTableAlertFull(){
        var elemAlert = $('table.v-table > tbody > tr > td >  .v-alert');
        var theadWidth = elemAlert.closest('table.v-table').find('thead > tr:first').width() - 50;
        elemAlert.parent('td').width(theadWidth);
      },
    }
}
</script>
