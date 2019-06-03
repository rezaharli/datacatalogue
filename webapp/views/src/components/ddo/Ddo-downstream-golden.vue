<style scoped>
/* .golden-source-level-1 table.v-table tbody tr td, .table-v1 table.v-table tfoot tr td{padding: 0;} */
.golden-source-level-1 table.v-table tbody tr td:first-of-type, 
.golden-source-level-1 table.v-table tfoot tr td:first-of-type {
  padding-left: 30px;
  /* border-left: none; */
}
.golden-source-level-1 table.v-table tbody tr td:last-of-type, 
.golden-source-level-1 table.v-table tfoot tr td:last-of-type {
  padding-right: 30px;
  /* border-right: none; */
}
</style>

<template>
  <v-data-table
      :headers="store.leftHeaders.filter(v => v.display == true && v.forInnerTable == true)"
      :items="props.item.AliasName"
      :expand="false"
      item-key="ID"
      class="golden-source-level-1"
      hide-headers
      hide-actions>

    <template slot="items" slot-scope="props">
      <tr :class="{even: props.index % 2, odd: !(props.index % 2)}">
        <td v-bind:style="{ width: store.left.colWidth['ALIAS_NAME'] + 'px' }" class="text-capitalize text-title">
          <b-link @click="props.expanded = !props.expanded" v-if="props.item.GoldenSource.length > 0">
            <tablecell :fulltext="props.item.ALIAS_NAME" showOn="hover"></tablecell>
          </b-link>

          <tablecell :fulltext="props.item.ALIAS_NAME" showOn="hover" v-if="props.item.GoldenSource.length < 1"></tablecell>
        </td>

        <td v-bind:style="{ width: store.left.colWidth['GOLDEN_SOURCE'] + 'px' }" class="text-uppercase">
          <tablecell showOn="hover" v-if="isGSMainLevelCellShowing(props)" :fulltext="props.item.GOLDEN_SOURCE"></tablecell>
        </td>

        <td v-bind:style="{ width: store.left.colWidth['GS_SYSTEM_NAME'] + 'px' }" class="text-uppercase">
          <tablecell showOn="hover" v-if="isGSMainLevelCellShowing(props) && props.item.GOLDEN_SOURCE.toLowerCase() == 'no'" :fulltext="props.item.GS_SYSTEM_NAME"></tablecell>
        </td>

        <td v-bind:style="{ width: store.left.colWidth['GS_TABLE_NAME'] + 'px' }" class="text-uppercase">
          <tablecell showOn="hover" v-if="isGSMainLevelCellShowing(props) && props.item.GOLDEN_SOURCE.toLowerCase() == 'no'" :fulltext="props.item.GS_TABLE_NAME"></tablecell>
        </td>

        <td v-bind:style="{ width: store.left.colWidth['GS_COLUMN_NAME'] + 'px' }" class="text-uppercase">
          <tablecell showOn="hover" v-if="isGSMainLevelCellShowing(props) && props.item.GOLDEN_SOURCE.toLowerCase() == 'no'" :fulltext="props.item.GS_COLUMN_NAME"></tablecell>
        </td>
      </tr>
    </template>

    <template slot="expand" slot-scope="props">
      <v-data-table
        :headers="store.leftHeaders.filter(v => v.display == true && v.forInnerTable == true)"
        :items="props.item.GoldenSource"
        :expand="false"
        class="golden-source-level-2"
        item-key="GSID"
        hide-actions
        hide-headers
        @update:pagination="setExpandedTableColumnsWidthGS"
      >
        <template slot="items" slot-scope="props">
          <td v-bind:style="{ width: store.left.colWidth['ALIAS_NAME'] + 'px' }">&nbsp;</td>

          <td v-bind:style="{ width: store.left.colWidth['GOLDEN_SOURCE'] + 'px' }">
            <template v-if="props.item.GOLDEN_SOURCE.toLowerCase() == 'no'">
              <b-link @click="props.expanded = !props.expanded" v-if="props.item.GSSystemName.length >= 1">
                <tablecell :fulltext="props.item.GOLDEN_SOURCE" showOn="hover"></tablecell>
              </b-link>

              <tablecell :fulltext="props.item.GOLDEN_SOURCE" showOn="hover" v-if="props.item.GSSystemName.length < 1"></tablecell>
            </template>

            <template v-if="props.item.GOLDEN_SOURCE.toLowerCase() == 'yes'">
              <tablecell :fulltext="props.item.GOLDEN_SOURCE" showOn="hover"></tablecell>
            </template>
          </td>

          <td class="text-uppercase" v-bind:style="{ width: store.left.colWidth['GS_SYSTEM_NAME'] + 'px' }">
            <tablecell showOn="hover" v-if="isGSSecondLevelCellShowing(props) && props.item.GOLDEN_SOURCE.toLowerCase() == 'no'" :fulltext="props.item.GS_SYSTEM_NAME"></tablecell>
          </td>

          <td class="text-uppercase" v-bind:style="{ width: store.left.colWidth['GS_TABLE_NAME'] + 'px' }">
            <tablecell showOn="hover" v-if="isGSSecondLevelCellShowing(props) && props.item.GOLDEN_SOURCE.toLowerCase() == 'no'" :fulltext="props.item.GS_TABLE_NAME"></tablecell>
          </td>

          <td class="text-uppercase" v-bind:style="{ width: store.left.colWidth['GS_COLUMN_NAME'] + 'px' }">
            <tablecell showOn="hover" v-if="isGSSecondLevelCellShowing(props) && props.item.GOLDEN_SOURCE.toLowerCase() == 'no'" :fulltext="props.item.GS_COLUMN_NAME"></tablecell>
          </td>
        </template>

        <template slot="expand" slot-scope="props">
          <v-data-table
            :headers="store.leftHeaders.filter(v => v.display == true && v.forInnerTable == true)"
            :items="props.item.GSSystemName"
            :expand="false"
            class="golden-source-level-3"
            item-key="SYSID"
            hide-actions
            hide-headers
            @update:pagination="setExpandedTableColumnsWidthGS"
          >
            <template slot="items" slot-scope="props">
              <td v-bind:style="{ width: store.left.colWidth['ALIAS_NAME'] + 'px' }">&nbsp;</td>
              <td v-bind:style="{ width: store.left.colWidth['GOLDEN_SOURCE'] + 'px' }">&nbsp;</td>

              <td v-bind:style="{ width: store.left.colWidth['GS_SYSTEM_NAME'] + 'px' }">
                <b-link @click="props.expanded = !props.expanded" v-if="props.item.GSTableName.length >= 1">
                  <tablecell :fulltext="props.item.GS_SYSTEM_NAME" showOn="hover"></tablecell>
                </b-link>

                <tablecell :fulltext="props.item.GS_SYSTEM_NAME" showOn="hover" v-if="props.item.GSTableName.length < 1"></tablecell>
              </td>

              <td class="text-uppercase" v-bind:style="{ width: store.left.colWidth['GS_TABLE_NAME'] + 'px' }">
                <tablecell showOn="hover" v-if="isGSThirdLevelCellShowing(props)" :fulltext="props.item.GS_TABLE_NAME"></tablecell>
              </td>

              <td class="text-uppercase" v-bind:style="{ width: store.left.colWidth['GS_COLUMN_NAME'] + 'px' }">
                <tablecell showOn="hover" v-if="isGSThirdLevelCellShowing(props)" :fulltext="props.item.GS_COLUMN_NAME"></tablecell>
              </td>
            </template>

            <template slot="expand" slot-scope="props">
              <v-data-table
                :headers="store.leftHeaders.filter(v => v.display == true && v.forInnerTable == true)"
                :items="props.item.GSTableName"
                :expand="false"
                class="golden-source-level-4"
                item-key="TBTID"
                hide-actions
                hide-headers
                @update:pagination="setExpandedTableColumnsWidthGS"
              >
                <template slot="items" slot-scope="props">
                  <td v-bind:style="{ width: store.left.colWidth['ALIAS_NAME'] + 'px' }">&nbsp;</td>
                  <td v-bind:style="{ width: store.left.colWidth['GOLDEN_SOURCE'] + 'px' }">&nbsp;</td>
                  <td v-bind:style="{ width: store.left.colWidth['GS_SYSTEM_NAME'] + 'px' }">&nbsp;</td>

                  <td v-bind:style="{ width: store.left.colWidth['GS_TABLE_NAME'] + 'px' }">
                    <b-link @click="props.expanded = !props.expanded" v-if="props.item.GSColumnName.length >= 1">
                      <tablecell :fulltext="props.item.GS_TABLE_NAME" showOn="hover"></tablecell>
                    </b-link>

                    <tablecell :fulltext="props.item.GS_TABLE_NAME" showOn="hover" v-if="props.item.GSColumnName.length < 1"></tablecell>
                  </td>

                  <td class="text-uppercase" v-bind:style="{ width: store.left.colWidth['GS_COLUMN_NAME'] + 'px' }">
                    <tablecell showOn="hover" v-if="isGSTableLevelCellShowing(props)" :fulltext="props.item.GS_COLUMN_NAME"></tablecell>
                  </td>
                </template>

                <template slot="expand" slot-scope="props">
                  <v-data-table
                    :headers="store.leftHeaders.filter(v => v.display == true && v.forInnerTable == true)"
                    :items="props.item.GSColumnName"
                    item-key="COLID"
                    class="golden-source-level-5"
                    hide-actions
                    hide-headers
                    @update:pagination="setExpandedTableColumnsWidthGS"
                  >
                    <template slot="items" slot-scope="props">
                      <td v-bind:style="{ width: store.left.colWidth['ALIAS_NAME'] + 'px' }">&nbsp;</td>
                      <td v-bind:style="{ width: store.left.colWidth['GOLDEN_SOURCE'] + 'px' }">&nbsp;</td>
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
      this.setExpandedTableColumnsWidthGS();
    }, 300);
  },
  updated() {
    this.setExpandedTableColumnsWidthGS();
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
        if(props.item.GoldenSource.length > 0) {
          if(props.item.GoldenSource[0].GSSystemName.length == 0) return true;
        }
        
        return false;
      }
    },
    isGSSecondLevelCellShowing (props){
      if( ! props.expanded) return true;
      else {
        if(props.item.GSSystemName.length > 0) {
          if(props.item.GSSystemName[0].GSTableName.length == 0) return true;
        }
        
        return false;
      }
    },
    isGSThirdLevelCellShowing (props){
      if( ! props.expanded) return true;
      else {
        if(props.item.GSTableName.length > 0) {
          if(props.item.GSTableName[0].GSColumnName.length == 0) return true;
        }
        
        return false;
      }
    },
    isGSTableLevelCellShowing (props){
      if( ! props.expanded) return true;
      else {
        if(props.item.GSColumnName.length > 0) {
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
    setExpandedTableColumnsWidthGS(){
      setTimeout(() => {
        var elem = $('.v-datatable__expand-row');
        var tableElem = elem.find('.v-datatable__expand-content:first table.v-table:first');
        var theadTR = elem.closest('.table-v1').find('table.v-table:first thead tr:first');
        var THs = theadTR.find('th');
        var tbodyTR = tableElem.children('tbody').children('tr');;
        var thWidths = [];
        THs.each(function (thIndex) {
          thWidths[thIndex] = $(this).outerWidth();
        });
        
        var tableLv1 = $('.golden-source-level-1');
        var tableLv1TRs = tableLv1.find('table.v-table > tbody > tr');
        tableLv1TRs.each(function () {
          $(this).children('td:not([colspan])').each(function (tdIndex2) {
            var colWidth = thWidths[parseInt(tdIndex2)+5];
            colWidth = colWidth - 60; // untuk mengurangi additional width yang datang tiba2 seperti syaiton, xixixi
            $(this).width(colWidth);
            $(this).addClass('tdindex-egs-1-'+tdIndex2+'-'+colWidth);
          });
        });

        var tableLv2 = $('.golden-source-level-2');
        var tableLv2TRs = tableLv2.find('table.v-table > tbody > tr');
        tableLv2TRs.each(function () {
          $(this).children('td:not([colspan])').each(function (tdIndex2) {
            var colWidth = thWidths[parseInt(tdIndex2)+5];
            colWidth = colWidth - 60; // untuk mengurangi additional width yang datang tiba2 seperti syaiton, xixixi
            $(this).width(colWidth);
            $(this).addClass('tdindex-egs-2-'+tdIndex2+'-'+colWidth);
          });
        });

        var tableLv3 = $('.golden-source-level-3');
        var tableLv3TRs = tableLv3.find('table.v-table > tbody > tr');
        tableLv3TRs.each(function () {
          $(this).children('td:not([colspan])').each(function (tdIndex2) {
            var colWidth = thWidths[parseInt(tdIndex2)+5];
            colWidth = colWidth - 60; // untuk mengurangi additional width yang datang tiba2 seperti syaiton, xixixi
            $(this).width(colWidth);
            $(this).addClass('tdindex-egs-3-'+tdIndex2+'-'+colWidth);
          });
        });

        var tableLv4 = $('.golden-source-level-4');
        var tableLv4TRs = tableLv4.find('table.v-table > tbody > tr');
        tableLv4TRs.each(function () {
          $(this).children('td:not([colspan])').each(function (tdIndex2) {
            var colWidth = thWidths[parseInt(tdIndex2)+5];
            colWidth = colWidth - 60; // untuk mengurangi additional width yang datang tiba2 seperti syaiton, xixixi
            $(this).width(colWidth);
            $(this).addClass('tdindex-egs-4-'+tdIndex2+'-'+colWidth);
          });
        });

        var tableLv5 = $('.golden-source-level-5');
        var tableLv5TRs = tableLv5.find('table.v-table > tbody > tr');
        tableLv5TRs.each(function () {
          $(this).children('td:not([colspan])').each(function (tdIndex2) {
            var colWidth = thWidths[parseInt(tdIndex2)+5];
            colWidth = colWidth - 60; // untuk mengurangi additional width yang datang tiba2 seperti syaiton, xixixi
            $(this).width(colWidth);
            $(this).addClass('tdindex-egs-5-'+tdIndex2+'-'+colWidth);
          });
        });
      }, 10);
    }
  }
};
</script>
