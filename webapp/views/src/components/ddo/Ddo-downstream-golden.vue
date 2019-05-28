<template>
  <v-data-table
      :headers="store.leftHeaders.filter(v => v.display == true && v.forInnerTable == true)"
      :items="props.item.GSSystems"
      :expand="false"
      item-key="ID"
      class=""
      hide-actions>

    <template slot="items" slot-scope="props">
      <tr :class="{even: props.index % 2, odd: !(props.index % 2)}">
        <td v-bind:style="{ width: store.left.colWidth['GS_SYSTEM_NAME'] + 'px' }" class="text-capitalize text-title">
          <b-link @click="props.expanded = !props.expanded" v-if="props.item.GSTables.length > 0">
            <tablecell :fulltext="props.item.GS_SYSTEM_NAME" showOn="hover"></tablecell>
          </b-link>

          <tablecell :fulltext="props.item.GS_SYSTEM_NAME" showOn="hover" v-if="props.item.GSTables.length < 1"></tablecell>
        </td>

        <td v-bind:style="{ width: store.left.colWidth['GS_TABLE_NAME'] + 'px' }" class="text-uppercase">
          <tablecell showOn="hover" v-if="isGSMainLevelCellShowing(props)" :fulltext="props.item.GS_TABLE_NAME"></tablecell>
        </td>

        <td v-bind:style="{ width: store.left.colWidth['GS_COLUMN_NAME'] + 'px' }" class="text-uppercase">
          <tablecell showOn="hover" v-if="isGSMainLevelCellShowing(props)" :fulltext="props.item.GS_COLUMN_NAME"></tablecell>
        </td>
      </tr>
    </template>

    <template slot="expand" slot-scope="props">
      <v-data-table
        :headers="store.leftHeaders.filter(v => v.display == true && v.forInnerTable == true)"
        :items="props.item.GSTables"
        :expand="false"
        class=""
        item-key="TMTID"
        hide-actions
        hide-headers
        @update:pagination="setExpandedTableColumnsWidth"
      >
        <template slot="items" slot-scope="props">
          <td v-bind:style="{ width: store.left.colWidth['GS_SYSTEM_NAME'] + 'px' }">&nbsp;</td>

          <td v-bind:style="{ width: store.left.colWidth['GS_TABLE_NAME'] + 'px' }">
            <b-link @click="props.expanded = !props.expanded" v-if="props.item.GSColumns.length >= 1">
              <tablecell :fulltext="props.item.GS_TABLE_NAME" showOn="hover"></tablecell>
            </b-link>

            <tablecell :fulltext="props.item.GS_TABLE_NAME" showOn="hover" v-if="props.item.GSColumns.length < 1"></tablecell>
          </td>

          <td class="text-uppercase" v-bind:style="{ width: store.left.colWidth['GS_COLUMN_NAME'] + 'px' }">
            <tablecell showOn="hover" v-if="isGSTableLevelCellShowing(props)" :fulltext="props.item.GS_COLUMN_NAME"></tablecell>
          </td>
        </template>

        <template slot="expand" slot-scope="props">
          <v-data-table
            :headers="store.leftHeaders.filter(v => v.display == true && v.forInnerTable == true)"
            :items="props.item.GSColumns"
            item-key="COLID"
            class=""
            hide-actions
            hide-headers
            @update:pagination="setExpandedTableColumnsWidth"
          >
            <template slot="items" slot-scope="props">
              <td v-bind:style="{ width: store.left.colWidth['GS_SYSTEM_NAME'] + 'px' }">&nbsp;</td>
              <td v-bind:style="{ width: store.left.colWidth['GS_TABLE_NAME'] + 'px' }">&nbsp;</td>

              <td v-bind:style="{ width: store.left.colWidth['GS_COLUMN_NAME'] + 'px' }">
                <tablecell :fulltext="props.item.GS_COLUMN_NAME" showOn="hover"></tablecell>
              </td>
            </template>
          </v-data-table>
        </template>
      </v-data-table>
    </template>
  </v-data-table>
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
  name: "downstreamGS",
  props: ["props"],
  components: {
    PageHeader, pageSearch, pageExport, tableheader, tablecell, pageLoader
  },
  data() {
    return {
      storeName: "ddodownstreambusinessterm",
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
    this.store.subdomain = this.$route.params.subdomain;
    this.store.system = this.$route.params.system;
    this.resetFilter();
    setTimeout(() => {
      this.setTableColumnsWidth($('#table-ddo-downstream-businessterm'));
    }, 300);
  },
  updated() {
    this.setTableColumnsWidth($('#table-ddo-downstream-businessterm'));
  },
  methods: {
    getLeftTable() {
      this.$store.dispatch(`${this.storeName}/getLeftTable`);
    },
    isMainLevelCellShowing (props){
      if( ! props.expanded) return true;
      else {
        if(props.item.Systems.length > 0) {
          if(props.item.Systems[0].Tables.length == 0) return true;
        }
        
        return false;
      }
    },
    isSystemLevelCellShowing (props){
      if( ! props.expanded) return true;
      else {
        if(props.item.Tables.length > 0) {
          if(props.item.Tables[0].Columns.length == 0) return true;
        }
        
        return false;
      }
    },
    isTableLevelCellShowing (props){
      if( ! props.expanded) return true;
      else {
        if(props.item.Columns.length > 0) {
          return true;
        }
        
        return false;
      }
    },
    isGSMainLevelCellShowing (props){
      if( ! props.expanded) return true;
      else {
        if(props.item.GSTables.length > 0) {
          if(props.item.GSTables[0].GSColumns.length == 0) return true;
        }
        
        return false;
      }
    },
    isGSTableLevelCellShowing (props){
      if( ! props.expanded) return true;
      else {
        if(props.item.GSColumns.length > 0) {
          return true;
        }
        
        return false;
      }
    },
    systemRowClick(evt) {
      evt.preventDefault();
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
    getCDEConclusion(cdes) {
      return cdes
        .filter(Boolean)
        .join(", ")
        .indexOf("Yes") != -1
        ? "Yes"
        : "No";
    },
    showDetails(param) {
      this.$router.push(
        this.addressPath + "/" + encodeURIComponent(this.$route.params.subdomain) + "/" + encodeURIComponent(this.$route.params.system) + "/" + encodeURIComponent(param.BT_NAME)
      );
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
    setExpandedTableColumnsWidth(){
      setTimeout(() => {
        var elem = $('.v-datatable__expand-row');
        var elemExpandedTable = elem.find('.v-datatable__expand-content table.v-table');
        var THs = elem.closest('table.v-table').find('thead tr:first th');
        var tbodyTR = elemExpandedTable.find('tbody tr');
        THs.each(function (thIndex) {
          $(this).css({'color': 'red'});
          var thWidth = $(this).width();
          tbodyTR.each(function (tdIndex) {
            var TDs = $(this).find('td:not([colspan])');
            TDs.eq(thIndex).width(thWidth);
          });
        });
      }, 10);
    }
  }
};
</script>
