<style>
/* #table-ddo-downstream-businessterm table.v-table.v-datatable tbody tr {display: table-row;} */

.table-v1 .v-table__overflow{
  overflow-y: hidden !important;
}
</style>


<template>
    <v-content class="mx-4 my-5">
        <b-container fluid>
            <PageHeader />

            <page-loader v-if="store.left.isLoading" />
            
            <b-row>
                <b-col sm=12 md=3>
                    <div class="card card-v2 transition">
                        <h6 class="title-1">Data Domain</h6>
                        <h3 class="title-2 text-capitalize">{{ store.left.source[0] ? store.left.source[0]["DATA_DOMAIN"] : " " }}</h3>
                    </div>
                </b-col>

                <b-col sm=12 md=3>
                    <div class="card card-v2 transition">
                        <h6 class="title-1">Sub-domain</h6>
                        <h3 class="title-2 text-capitalize">{{ store.left.source[0] ? store.left.source[0]["SUB_DOMAINS"] : " " }}</h3>
                    </div>
                </b-col>
            </b-row>

            <b-row class="my-4">
                <b-col>
                    <b-button class="float-right red-neon icon-only ml-3" @click="resetFilter">
                        <i class="fa fa-fw fa-filter"></i>
                    </b-button>

                    <page-export class="float-right" :storeName="storeName" :leftTableCols="store.leftHeaders" :rightTableCols="[]"/>
                </b-col>
            </b-row>

            <b-row style="margin-top: 10px;margin-bottom: 10px;">
              <b-col>
                <!-- Ddo details -->
                <router-view/>

                <!-- Main content -->
                <v-data-table
                    :headers="store.leftHeaders.filter(v => v.display == true)"
                    :items="store.left.display"
                    :pagination.sync="store.left.pagination"
                    :total-items="store.left.totalItems"
                    :loading="store.left.isLoading"
                    :expand="false"
                    :must-sort="true"
                    item-key="ID"
                    class="table-v1"
                    id="table-ddo-downstream-businessterm">
                  <template slot="headerCell" slot-scope="props">
                    <tableheader :storeName="storeName" :props="props" :which="'left'"/>
                  </template>

                  <v-progress-linear slot="progress" color="blue" indeterminate></v-progress-linear>

                  <template slot="no-data">
                    <v-alert
                        :value="store.left.isLoading"
                        type="info"
                      >Please wait while data is loading</v-alert>

                    <v-alert
                        :value="!store.left.isLoading"
                        type="error"
                      >Sorry, nothing to display here</v-alert>
                  </template>

                  <template slot="items" slot-scope="props">
                    <tr :class="{even: props.index % 2, odd: !(props.index % 2)}">
                      <td v-bind:style="{ width: store.left.colWidth['Details'] + 'px' }" class="text-capitalize text-title">
                        <b-button size="sm" class="green-tosca-gradient icon-only" @click="showDetails(props.item)">
                          <i class="fa fa-fw fa-external-link-alt"></i></b-button></td>

                      <td v-bind:style="{ width: store.left.colWidth['BT_NAME'] + 'px' }" class="text-capitalize text-title">
                        <b-link @click="props.expanded = !props.expanded" v-if="props.item.ALIAS_NAMEs.length > 0">
                          {{ props.item.BT_NAME }}
                        </b-link>

                        <span v-if="props.item.ALIAS_NAMEs.length < 1">{{ props.item.BT_NAME }}</span>
                      </td>

                      <td v-bind:style="{ width: store.left.colWidth['ALIAS_NAME'] + 'px' }" class="text-uppercase">
                        <tablecell showOn="hover" v-if="isMainLevelCellShowing(props)" :fulltext="props.item.ALIAS_NAME"></tablecell>
                      </td>

                      <td v-bind:style="{ width: store.left.colWidth['SYSTEM_CONSUMED'] + 'px' }" class="text-uppercase">
                        <tablecell showOn="hover" v-if="isMainLevelCellShowing(props)" :fulltext="props.item.SYSTEM_CONSUMED"></tablecell>
                      </td>

                      <td v-bind:style="{ width: store.left.colWidth['TABLE_NAME'] + 'px' }" class="text-uppercase">
                        <tablecell showOn="hover" v-if="isMainLevelCellShowing(props)" :fulltext="props.item.TABLE_NAME"></tablecell>
                      </td>

                      <td v-bind:style="{ width: store.left.colWidth['COLUMN_NAME'] + 'px' }" class="text-uppercase">
                        <tablecell showOn="hover" v-if="isMainLevelCellShowing(props)" :fulltext="props.item.COLUMN_NAME"></tablecell>
                      </td>

                      <td v-bind:style="{ width: store.left.colWidth['GOLDEN_SOURCE'] + 'px' }" class="text-uppercase">
                        <tablecell showOn="hover" v-if="isMainLevelCellShowing(props)" :fulltext="props.item.GOLDEN_SOURCE"></tablecell>
                      </td>

                      <td v-bind:style="{ width: store.left.colWidth['GS_SYSTEM_NAME'] + 'px' }" class="text-uppercase">
                        <tablecell showOn="hover" v-if="isMainLevelCellShowing(props)" :fulltext="props.item.GS_SYSTEM_NAME"></tablecell>
                      </td>

                      <td v-bind:style="{ width: store.left.colWidth['GS_TABLE_NAME'] + 'px' }" class="text-uppercase">
                        <tablecell showOn="hover" v-if="isMainLevelCellShowing(props)" :fulltext="props.item.GS_TABLE_NAME"></tablecell>
                      </td>

                      <td v-bind:style="{ width: store.left.colWidth['GS_COLUMN_NAME'] + 'px' }" class="text-uppercase">
                        <tablecell showOn="hover" v-if="isMainLevelCellShowing(props)" :fulltext="props.item.GS_COLUMN_NAME"></tablecell>
                      </td>

                      <!-- <td v-bind:style="{ width: (store.left.colWidth['ALIAS_NAME'] + store.left.colWidth['GOLDEN_SOURCE'] + store.left.colWidth['GS_SYSTEM_NAME'] + store.left.colWidth['GS_TABLE_NAME'] + store.left.colWidth['GS_COLUMN_NAME']) + 'px' }">
                        <downstream-g-s :props="props"></downstream-g-s>
                      </td> -->
                    </tr>
                  </template>

                  <template slot="expand" slot-scope="props">
                    <!-- <table-rows-sub :storeName="storeName" :props="props" /> -->
                    <v-data-table
                      :headers="store.leftHeaders.filter(v => v.display == true)"
                      :items="props.item.ALIAS_NAMEs"
                      :expand="false"
                      class="expanded-table-level-1"
                      item-key="ALIAS_NAMEID"
                      hide-actions
                      hide-headers
                      @update:pagination="setExpandedTableColumnsWidthDBT"
                    >
                      <template slot="items" slot-scope="props">
                        <td v-bind:style="{ width: store.left.colWidth['Details'] + 'px' }">&nbsp;</td>
                        <td v-bind:style="{ width: store.left.colWidth['BT_NAME'] + 'px' }">&nbsp;</td>

                        <td v-bind:style="{ width: store.left.colWidth['ALIAS_NAME'] + 'px' }">
                          <b-link @click="props.expanded = !props.expanded" v-if="props.item.SYSTEM_CONSUMEDs.length >= 1">
                            <tablecell :fulltext="props.item.ALIAS_NAME" showOn="hover"></tablecell>
                          </b-link>

                          <tablecell :fulltext="props.item.ALIAS_NAME" showOn="hover" v-if="props.item.SYSTEM_CONSUMEDs.length < 1"></tablecell>
                        </td>

                        <td class="text-uppercase" v-bind:style="{ width: store.left.colWidth['SYSTEM_CONSUMED'] + 'px' }">
                          <tablecell showOn="hover" v-if="isSecondLevelCellShowing(props)" :fulltext="props.item.SYSTEM_CONSUMED"></tablecell>
                        </td>

                        <td class="text-uppercase" v-bind:style="{ width: store.left.colWidth['TABLE_NAME'] + 'px' }">
                          <tablecell showOn="hover" v-if="isSecondLevelCellShowing(props)" :fulltext="props.item.TABLE_NAME"></tablecell>
                        </td>

                        <td class="text-uppercase" v-bind:style="{ width: store.left.colWidth['COLUMN_NAME'] + 'px' }">
                          <tablecell showOn="hover" v-if="isSecondLevelCellShowing(props)" :fulltext="props.item.COLUMN_NAME"></tablecell>
                        </td>

                        <td class="text-uppercase" v-bind:style="{ width: store.left.colWidth['GOLDEN_SOURCE'] + 'px' }">
                          <tablecell showOn="hover" v-if="isSecondLevelCellShowing(props)" :fulltext="props.item.GOLDEN_SOURCE"></tablecell>
                        </td>

                        <td class="text-uppercase" v-bind:style="{ width: store.left.colWidth['GS_SYSTEM_NAME'] + 'px' }">
                          <tablecell showOn="hover" v-if="isSecondLevelCellShowing(props)" :fulltext="props.item.GS_SYSTEM_NAME"></tablecell>
                        </td>

                        <td class="text-uppercase" v-bind:style="{ width: store.left.colWidth['GS_TABLE_NAME'] + 'px' }">
                          <tablecell showOn="hover" v-if="isSecondLevelCellShowing(props)" :fulltext="props.item.GS_TABLE_NAME"></tablecell>
                        </td>

                        <td class="text-uppercase" v-bind:style="{ width: store.left.colWidth['GS_COLUMN_NAME'] + 'px' }">
                          <tablecell showOn="hover" v-if="isSecondLevelCellShowing(props)" :fulltext="props.item.GS_COLUMN_NAME"></tablecell>
                        </td>
                      </template>

                      <template slot="expand" slot-scope="props">
                        <v-data-table
                          :headers="store.leftHeaders.filter(v => v.display == true)"
                          :items="props.item.SYSTEM_CONSUMEDs"
                          item-key="SYSTEM_CONSUMEDID"
                          class="expanded-table-level-2"
                          hide-actions
                          hide-headers
                          @update:pagination="setExpandedTableColumnsWidthDBT"
                        >
                          <template slot="items" slot-scope="props">
                            <td v-bind:style="{ width: store.left.colWidth['Details'] + 'px' }">&nbsp;</td>
                            <td v-bind:style="{ width: store.left.colWidth['BT_NAME'] + 'px' }">&nbsp;</td>
                            <td v-bind:style="{ width: store.left.colWidth['ALIAS_NAME'] + 'px' }">&nbsp;</td>

                            <td v-bind:style="{ width: store.left.colWidth['SYSTEM_CONSUMED'] + 'px' }">
                              <b-link @click="props.expanded = !props.expanded" v-if="props.item.TABLE_NAMEs.length >= 1">
                                <tablecell :fulltext="props.item.SYSTEM_CONSUMED" showOn="hover"></tablecell>
                              </b-link>

                              <tablecell :fulltext="props.item.SYSTEM_CONSUMED" showOn="hover" v-if="props.item.TABLE_NAMEs.length < 1"></tablecell>
                            </td>

                            <td class="text-uppercase" v-bind:style="{ width: store.left.colWidth['TABLE_NAME'] + 'px' }">
                              <tablecell showOn="hover" v-if="isThirdLevelCellShowing(props)" :fulltext="props.item.TABLE_NAME"></tablecell>
                            </td>

                            <td class="text-uppercase" v-bind:style="{ width: store.left.colWidth['COLUMN_NAME'] + 'px' }">
                              <tablecell showOn="hover" v-if="isThirdLevelCellShowing(props)" :fulltext="props.item.COLUMN_NAME"></tablecell>
                            </td>

                            <td class="text-uppercase" v-bind:style="{ width: store.left.colWidth['GOLDEN_SOURCE'] + 'px' }">
                              <tablecell showOn="hover" v-if="isThirdLevelCellShowing(props)" :fulltext="props.item.GOLDEN_SOURCE"></tablecell>
                            </td>

                            <td class="text-uppercase" v-bind:style="{ width: store.left.colWidth['GS_SYSTEM_NAME'] + 'px' }">
                              <tablecell showOn="hover" v-if="isThirdLevelCellShowing(props)" :fulltext="props.item.GS_SYSTEM_NAME"></tablecell>
                            </td>

                            <td class="text-uppercase" v-bind:style="{ width: store.left.colWidth['GS_TABLE_NAME'] + 'px' }">
                              <tablecell showOn="hover" v-if="isThirdLevelCellShowing(props)" :fulltext="props.item.GS_TABLE_NAME"></tablecell>
                            </td>

                            <td class="text-uppercase" v-bind:style="{ width: store.left.colWidth['GS_COLUMN_NAME'] + 'px' }">
                              <tablecell showOn="hover" v-if="isThirdLevelCellShowing(props)" :fulltext="props.item.GS_COLUMN_NAME"></tablecell>
                            </td>
                          </template>

                          <template slot="expand" slot-scope="props">
                            <v-data-table
                              :headers="store.leftHeaders.filter(v => v.display == true)"
                              :items="props.item.TABLE_NAMEs"
                              item-key="TABLE_NAMEID"
                              class="expanded-table-level-2"
                              hide-actions
                              hide-headers
                              @update:pagination="setExpandedTableColumnsWidthDBT"
                            >
                              <template slot="items" slot-scope="props">
                                <td v-bind:style="{ width: store.left.colWidth['Details'] + 'px' }">&nbsp;</td>
                                <td v-bind:style="{ width: store.left.colWidth['BT_NAME'] + 'px' }">&nbsp;</td>
                                <td v-bind:style="{ width: store.left.colWidth['ALIAS_NAME'] + 'px' }">&nbsp;</td>
                                <td v-bind:style="{ width: store.left.colWidth['SYSTEM_CONSUMED'] + 'px' }">&nbsp;</td>

                                <td v-bind:style="{ width: store.left.colWidth['TABLE_NAME'] + 'px' }">
                                  <b-link @click="props.expanded = !props.expanded" v-if="props.item.COLUMN_NAMEs.length >= 1">
                                    <tablecell :fulltext="props.item.TABLE_NAME" showOn="hover"></tablecell>
                                  </b-link>

                                  <tablecell :fulltext="props.item.TABLE_NAME" showOn="hover" v-if="props.item.COLUMN_NAMEs.length < 1"></tablecell>
                                </td>

                                <td class="text-uppercase" v-bind:style="{ width: store.left.colWidth['COLUMN_NAME'] + 'px' }">
                                  <tablecell showOn="hover" v-if="isFourthLevelCellShowing(props)" :fulltext="props.item.COLUMN_NAME"></tablecell>
                                </td>

                                <td class="text-uppercase" v-bind:style="{ width: store.left.colWidth['GOLDEN_SOURCE'] + 'px' }">
                                  <tablecell showOn="hover" v-if="isFourthLevelCellShowing(props)" :fulltext="props.item.GOLDEN_SOURCE"></tablecell>
                                </td>

                                <td class="text-uppercase" v-bind:style="{ width: store.left.colWidth['GS_SYSTEM_NAME'] + 'px' }">
                                  <tablecell showOn="hover" v-if="isFourthLevelCellShowing(props)" :fulltext="props.item.GS_SYSTEM_NAME"></tablecell>
                                </td>

                                <td class="text-uppercase" v-bind:style="{ width: store.left.colWidth['GS_TABLE_NAME'] + 'px' }">
                                  <tablecell showOn="hover" v-if="isFourthLevelCellShowing(props)" :fulltext="props.item.GS_TABLE_NAME"></tablecell>
                                </td>

                                <td class="text-uppercase" v-bind:style="{ width: store.left.colWidth['GS_COLUMN_NAME'] + 'px' }">
                                  <tablecell showOn="hover" v-if="isFourthLevelCellShowing(props)" :fulltext="props.item.GS_COLUMN_NAME"></tablecell>
                                </td>
                              </template>

                              <template slot="expand" slot-scope="props">
                                <v-data-table
                                  :headers="store.leftHeaders.filter(v => v.display == true)"
                                  :items="props.item.COLUMN_NAMEs"
                                  item-key="COLUMN_NAMEID"
                                  class="expanded-table-level-2"
                                  hide-actions
                                  hide-headers
                                  @update:pagination="setExpandedTableColumnsWidthDBT"
                                >
                                  <template slot="items" slot-scope="props">
                                    <td v-bind:style="{ width: store.left.colWidth['Details'] + 'px' }">&nbsp;</td>
                                    <td v-bind:style="{ width: store.left.colWidth['BT_NAME'] + 'px' }">&nbsp;</td>
                                    <td v-bind:style="{ width: store.left.colWidth['ALIAS_NAME'] + 'px' }">&nbsp;</td>
                                    <td v-bind:style="{ width: store.left.colWidth['SYSTEM_CONSUMED'] + 'px' }">&nbsp;</td>
                                    <td v-bind:style="{ width: store.left.colWidth['TABLE_NAME'] + 'px' }">&nbsp;</td>

                                    <td v-bind:style="{ width: store.left.colWidth['COLUMN_NAME'] + 'px' }">
                                      <b-link @click="props.expanded = !props.expanded" v-if="props.item.GOLDEN_SOURCEs.length >= 1">
                                        <tablecell :fulltext="props.item.COLUMN_NAME" showOn="hover"></tablecell>
                                      </b-link>

                                      <tablecell :fulltext="props.item.COLUMN_NAME" showOn="hover" v-if="props.item.GOLDEN_SOURCEs.length < 1"></tablecell>
                                    </td>

                                    <td class="text-uppercase" v-bind:style="{ width: store.left.colWidth['GOLDEN_SOURCE'] + 'px' }">
                                      <tablecell showOn="hover" v-if="isFifthLevelCellShowing(props)" :fulltext="props.item.GOLDEN_SOURCE"></tablecell>
                                    </td>

                                    <td class="text-uppercase" v-bind:style="{ width: store.left.colWidth['GS_SYSTEM_NAME'] + 'px' }">
                                      <tablecell showOn="hover" v-if="isFifthLevelCellShowing(props)" :fulltext="props.item.GS_SYSTEM_NAME"></tablecell>
                                    </td>

                                    <td class="text-uppercase" v-bind:style="{ width: store.left.colWidth['GS_TABLE_NAME'] + 'px' }">
                                      <tablecell showOn="hover" v-if="isFifthLevelCellShowing(props)" :fulltext="props.item.GS_TABLE_NAME"></tablecell>
                                    </td>

                                    <td class="text-uppercase" v-bind:style="{ width: store.left.colWidth['GS_COLUMN_NAME'] + 'px' }">
                                      <tablecell showOn="hover" v-if="isFifthLevelCellShowing(props)" :fulltext="props.item.GS_COLUMN_NAME"></tablecell>
                                    </td>
                                  </template>

                                  <template slot="expand" slot-scope="props">
                                    <v-data-table
                                      :headers="store.leftHeaders.filter(v => v.display == true)"
                                      :items="props.item.GOLDEN_SOURCEs"
                                      item-key="GOLDEN_SOURCEID"
                                      class="expanded-table-level-2"
                                      hide-actions
                                      hide-headers
                                      @update:pagination="setExpandedTableColumnsWidthDBT"
                                    >
                                      <template slot="items" slot-scope="props">
                                        <td v-bind:style="{ width: store.left.colWidth['Details'] + 'px' }">&nbsp;</td>
                                        <td v-bind:style="{ width: store.left.colWidth['BT_NAME'] + 'px' }">&nbsp;</td>
                                        <td v-bind:style="{ width: store.left.colWidth['ALIAS_NAME'] + 'px' }">&nbsp;</td>
                                        <td v-bind:style="{ width: store.left.colWidth['SYSTEM_CONSUMED'] + 'px' }">&nbsp;</td>
                                        <td v-bind:style="{ width: store.left.colWidth['TABLE_NAME'] + 'px' }">&nbsp;</td>
                                        <td v-bind:style="{ width: store.left.colWidth['COLUMN_NAME'] + 'px' }">&nbsp;</td>

                                        <td v-bind:style="{ width: store.left.colWidth['GOLDEN_SOURCE'] + 'px' }">
                                          <b-link @click="props.expanded = !props.expanded" v-if="props.item.GS_SYSTEM_NAMEs.length >= 1">
                                            <tablecell :fulltext="props.item.GOLDEN_SOURCE" showOn="hover"></tablecell>
                                          </b-link>

                                          <tablecell :fulltext="props.item.GOLDEN_SOURCE" showOn="hover" v-if="props.item.GS_SYSTEM_NAMEs.length < 1"></tablecell>
                                        </td>

                                        <td class="text-uppercase" v-bind:style="{ width: store.left.colWidth['GS_SYSTEM_NAME'] + 'px' }">
                                          <tablecell showOn="hover" v-if="isSixthLevelCellShowing(props)" :fulltext="props.item.GS_SYSTEM_NAME"></tablecell>
                                        </td>

                                        <td class="text-uppercase" v-bind:style="{ width: store.left.colWidth['GS_TABLE_NAME'] + 'px' }">
                                          <tablecell showOn="hover" v-if="isSixthLevelCellShowing(props)" :fulltext="props.item.GS_TABLE_NAME"></tablecell>
                                        </td>

                                        <td class="text-uppercase" v-bind:style="{ width: store.left.colWidth['GS_COLUMN_NAME'] + 'px' }">
                                          <tablecell showOn="hover" v-if="isSixthLevelCellShowing(props)" :fulltext="props.item.GS_COLUMN_NAME"></tablecell>
                                        </td>
                                      </template>

                                      <template slot="expand" slot-scope="props">
                                        <v-data-table
                                          :headers="store.leftHeaders.filter(v => v.display == true)"
                                          :items="props.item.GS_SYSTEM_NAMEs"
                                          item-key="GS_SYSTEM_NAMEID"
                                          class="expanded-table-level-2"
                                          hide-actions
                                          hide-headers
                                          @update:pagination="setExpandedTableColumnsWidthDBT"
                                        >
                                          <template slot="items" slot-scope="props">
                                            <td v-bind:style="{ width: store.left.colWidth['Details'] + 'px' }">&nbsp;</td>
                                            <td v-bind:style="{ width: store.left.colWidth['BT_NAME'] + 'px' }">&nbsp;</td>
                                            <td v-bind:style="{ width: store.left.colWidth['ALIAS_NAME'] + 'px' }">&nbsp;</td>
                                            <td v-bind:style="{ width: store.left.colWidth['SYSTEM_CONSUMED'] + 'px' }">&nbsp;</td>
                                            <td v-bind:style="{ width: store.left.colWidth['TABLE_NAME'] + 'px' }">&nbsp;</td>
                                            <td v-bind:style="{ width: store.left.colWidth['COLUMN_NAME'] + 'px' }">&nbsp;</td>
                                            <td v-bind:style="{ width: store.left.colWidth['GOLDEN_SOURCE'] + 'px' }">&nbsp;</td>

                                            <td v-bind:style="{ width: store.left.colWidth['GS_SYSTEM_NAME'] + 'px' }">
                                              <b-link @click="props.expanded = !props.expanded" v-if="props.item.GS_TABLE_NAMEs.length >= 1">
                                                <tablecell :fulltext="props.item.GS_SYSTEM_NAME" showOn="hover"></tablecell>
                                              </b-link>

                                              <tablecell :fulltext="props.item.GS_SYSTEM_NAME" showOn="hover" v-if="props.item.GS_TABLE_NAMEs.length < 1"></tablecell>
                                            </td>

                                            <td class="text-uppercase" v-bind:style="{ width: store.left.colWidth['GS_TABLE_NAME'] + 'px' }">
                                              <tablecell showOn="hover" v-if="isSeventhLevelCellShowing(props)" :fulltext="props.item.GS_TABLE_NAME"></tablecell>
                                            </td>

                                            <td class="text-uppercase" v-bind:style="{ width: store.left.colWidth['GS_COLUMN_NAME'] + 'px' }">
                                              <tablecell showOn="hover" v-if="isSeventhLevelCellShowing(props)" :fulltext="props.item.GS_COLUMN_NAME"></tablecell>
                                            </td>
                                          </template>

                                          <template slot="expand" slot-scope="props">
                                            <v-data-table
                                              :headers="store.leftHeaders.filter(v => v.display == true)"
                                              :items="props.item.GS_TABLE_NAMEs"
                                              item-key="GS_TABLE_NAMEID"
                                              class="expanded-table-level-2"
                                              hide-actions
                                              hide-headers
                                              @update:pagination="setExpandedTableColumnsWidthDBT"
                                            >
                                              <template slot="items" slot-scope="props">
                                                <td v-bind:style="{ width: store.left.colWidth['Details'] + 'px' }">&nbsp;</td>
                                                <td v-bind:style="{ width: store.left.colWidth['BT_NAME'] + 'px' }">&nbsp;</td>
                                                <td v-bind:style="{ width: store.left.colWidth['ALIAS_NAME'] + 'px' }">&nbsp;</td>
                                                <td v-bind:style="{ width: store.left.colWidth['SYSTEM_CONSUMED'] + 'px' }">&nbsp;</td>
                                                <td v-bind:style="{ width: store.left.colWidth['TABLE_NAME'] + 'px' }">&nbsp;</td>
                                                <td v-bind:style="{ width: store.left.colWidth['COLUMN_NAME'] + 'px' }">&nbsp;</td>
                                                <td v-bind:style="{ width: store.left.colWidth['GOLDEN_SOURCE'] + 'px' }">&nbsp;</td>
                                                <td v-bind:style="{ width: store.left.colWidth['GS_SYSTEM_NAME'] + 'px' }">&nbsp;</td>

                                                <td v-bind:style="{ width: store.left.colWidth['GS_TABLE_NAME'] + 'px' }">
                                                  <b-link @click="props.expanded = !props.expanded" v-if="props.item.GS_COLUMN_NAMEs.length >= 1">
                                                    <tablecell :fulltext="props.item.GS_TABLE_NAME" showOn="hover"></tablecell>
                                                  </b-link>

                                                  <tablecell :fulltext="props.item.GS_TABLE_NAME" showOn="hover" v-if="props.item.GS_COLUMN_NAMEs.length < 1"></tablecell>
                                                </td>

                                                <td class="text-uppercase" v-bind:style="{ width: store.left.colWidth['GS_COLUMN_NAME'] + 'px' }">
                                                  <tablecell showOn="hover" v-if="isEighthLevelCellShowing(props)" :fulltext="props.item.GS_COLUMN_NAME"></tablecell>
                                                </td>
                                              </template>

                                              <template slot="expand" slot-scope="props">
                                                <v-data-table
                                                  :headers="store.leftHeaders.filter(v => v.display == true)"
                                                  :items="props.item.GS_COLUMN_NAMEs"
                                                  item-key="GS_COLUMN_NAMEID"
                                                  class=""
                                                  hide-actions
                                                  hide-headers
                                                  @update:pagination="setExpandedTableColumnsWidthDBT"
                                                >
                                                  <template slot="items" slot-scope="props">
                                                    <td v-bind:style="{ width: store.left.colWidth['Details'] + 'px' }">&nbsp;</td>
                                                    <td v-bind:style="{ width: store.left.colWidth['BT_NAME'] + 'px' }">&nbsp;</td>
                                                    <td v-bind:style="{ width: store.left.colWidth['ALIAS_NAME'] + 'px' }">&nbsp;</td>
                                                    <td v-bind:style="{ width: store.left.colWidth['SYSTEM_CONSUMED'] + 'px' }">&nbsp;</td>
                                                    <td v-bind:style="{ width: store.left.colWidth['TABLE_NAME'] + 'px' }">&nbsp;</td>
                                                    <td v-bind:style="{ width: store.left.colWidth['COLUMN_NAME'] + 'px' }">&nbsp;</td>
                                                    <td v-bind:style="{ width: store.left.colWidth['GOLDEN_SOURCE'] + 'px' }">&nbsp;</td>
                                                    <td v-bind:style="{ width: store.left.colWidth['GS_SYSTEM_NAME'] + 'px' }">&nbsp;</td>
                                                    <td v-bind:style="{ width: store.left.colWidth['GS_TABLE_NAME'] + 'px' }">&nbsp;</td>

                                                    <td class="text-uppercase" v-bind:style="{ width: store.left.colWidth['GS_COLUMN_NAME'] + 'px' }">
                                                      <tablecell showOn="hover" :fulltext="props.item.GS_COLUMN_NAME"></tablecell></td>
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
                            </v-data-table>
                          </template>
                        </v-data-table>
                      </template>
                    </v-data-table>
                  </template>
                </v-data-table>
                      
              </b-col>
            </b-row>
              
        </b-container>
    </v-content>
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
import downstreamGS from "./Ddo-downstream-golden.vue";

Vue.component("downloadExcel", JsonExcel);

export default {
  components: {
    PageHeader, pageSearch, pageExport, tableheader, tablecell, pageLoader, downstreamGS
  },
  data() {
    return {
      storeName: "ddodownstreambusinessterm",
      systemSource: [],
      tablenameSource: []
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
      this.setTableColumnsWidthDBT($('#table-ddo-downstream-businessterm'));
    }, 300);
  },
  updated() {
    this.setTableColumnsWidthDBT($('#table-ddo-downstream-businessterm'));
  },
  methods: {
    getLeftTable() {
      this.$store.dispatch(`${this.storeName}/getLeftTable`);
    },
    isMainLevelCellShowing (props){
      if( ! props.expanded) return true;
      else {
        if(props.item.ALIAS_NAMEs.length > 0) {
          if(props.item.ALIAS_NAMEs[0].SYSTEM_CONSUMEDs.length == 0) return true;
        }
        
        return false;
      }
    },
    isSecondLevelCellShowing (props){
      if( ! props.expanded) return true;
      else {
        if(props.item.SYSTEM_CONSUMEDs.length > 0) {
          if(props.item.SYSTEM_CONSUMEDs[0].TABLE_NAMEs.length == 0) return true;
        }
        
        return false;
      }
    },
    isThirdLevelCellShowing (props){
      if( ! props.expanded) return true;
      else {
        if(props.item.TABLE_NAMEs.length > 0) {
          if(props.item.TABLE_NAMEs[0].COLUMN_NAMEs.length == 0) return true;
        }
        
        return false;
      }
    },
    isFourthLevelCellShowing (props){
      if( ! props.expanded) return true;
      else {
        if(props.item.COLUMN_NAMEs.length > 0) {
          if(props.item.COLUMN_NAMEs[0].GOLDEN_SOURCEs.length == 0) return true;
        }
        
        return false;
      }
    },
    isFifthLevelCellShowing (props){
      if( ! props.expanded) return true;
      else {
        if(props.item.GOLDEN_SOURCEs.length > 0) {
          if(props.item.GOLDEN_SOURCEs[0].GS_SYSTEM_NAMEs.length == 0) return true;
        }
        
        return false;
      }
    },
    isSixthLevelCellShowing (props){
      if( ! props.expanded) return true;
      else {
        if(props.item.GS_SYSTEM_NAMEs.length > 0) {
          if(props.item.GS_SYSTEM_NAMEs[0].GS_TABLE_NAMEs.length == 0) return true;
        }
        
        return false;
      }
    },
    isSeventhLevelCellShowing (props){
      if( ! props.expanded) return true;
      else {
        if(props.item.GS_TABLE_NAMEs.length > 0) {
          if(props.item.GS_TABLE_NAMEs[0].GS_COLUMN_NAMEs.length == 0) return true;
        }
        
        return false;
      }
    },
    isEighthLevelCellShowing (props){
      if( ! props.expanded) return true;
      else {
        if(props.item.GS_COLUMN_NAMEs.length > 0) {
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
    setTableColumnsWidthDBT(elem){
      var tableElem = elem.find('.v-table__overflow:first > table.v-table:first');

      var theadTR = elem.closest('.table-v1').find('table.v-table:first thead tr:first');
      var THs = theadTR.find('th');
      var tbodyTR = tableElem.children('tbody').children('tr');
      var thWidths = [];
      THs.each(function (thIndex) {
        thWidths[thIndex] = $(this).outerWidth();
      });

      tbodyTR.each(function () {
        $(this).children('td:not([colspan])').each(function (tdIndex2) {
          if(tdIndex2==5){
            var colWidth = thWidths[parseInt(tdIndex2)] + thWidths[parseInt(tdIndex2)+1] + thWidths[parseInt(tdIndex2)+2] + thWidths[parseInt(tdIndex2)+3] + thWidths[parseInt(tdIndex2)+4];
            $(this).css({'padding': '0'});
          }else{
            var colWidth = thWidths[parseInt(tdIndex2)];
          }
          if(tdIndex2==0){
            colWidth = colWidth - 70; // untuk mengurangi additional width yang datang tiba2 seperti syaiton, xixixi
          }else if(tdIndex2==5){
            colWidth = colWidth - 0; // untuk mengurangi additional width yang datang tiba2 seperti syaiton, xixixi
          }else{
            colWidth = colWidth - 60; // untuk mengurangi additional width yang datang tiba2 seperti syaiton, xixixi
          }
          $(this).width(colWidth);
          $(this).addClass('tdindex-dbt-'+tdIndex2+'-'+colWidth);
        });
      });

    },
    setExpandedTableColumnsWidthDBT(){
      setTimeout(() => {
        var elem = $('.v-datatable__expand-row');
        var tableElem = elem.find('.v-datatable__expand-content:first table.v-table:first');
        var theadTR = elem.closest('.table-v1').find('table.v-table:first thead tr:first');
        var THs = theadTR.find('th');
        var tbodyTR = tableElem.children('tbody').children('tr');
        var thWidths = [];
        THs.each(function (thIndex) {
          thWidths[thIndex] = $(this).outerWidth();
        });

        var tableLv1 = $('.expanded-table-level-1');
        var tableLv1TRs = tableLv1.find('table.v-table > tbody > tr');
        tableLv1TRs.each(function () {
          $(this).children('td:not([colspan])').each(function (tdIndex2) {
            if(tdIndex2==5){
              var colWidth = thWidths[parseInt(tdIndex2)] + thWidths[parseInt(tdIndex2)+1] + thWidths[parseInt(tdIndex2)+2] + thWidths[parseInt(tdIndex2)+3] + thWidths[parseInt(tdIndex2)+4];
            }else{
              var colWidth = thWidths[parseInt(tdIndex2)];
            }
            if(tdIndex2==0){
              colWidth = colWidth - 75; // untuk mengurangi additional width yang datang tiba2 seperti syaiton, xixixi
            }else{
              colWidth = colWidth - 60; // untuk mengurangi additional width yang datang tiba2 seperti syaiton, xixixi
            }
            $(this).width(colWidth);
            $(this).addClass('tdindex-dbt-lv1-'+tdIndex2+'-'+colWidth);
          });
        });

        var tableLv2 = $('.expanded-table-level-2');
        var tableLv2TRs = tableLv2.find('table.v-table > tbody > tr');
        tableLv2TRs.each(function () {
          $(this).children('td:not([colspan])').each(function (tdIndex2) {
            if(tdIndex2==5){
              var colWidth = thWidths[parseInt(tdIndex2)] + thWidths[parseInt(tdIndex2)+1] + thWidths[parseInt(tdIndex2)+2] + thWidths[parseInt(tdIndex2)+3] + thWidths[parseInt(tdIndex2)+4];
            }else{
              var colWidth = thWidths[parseInt(tdIndex2)];
            }
            if(tdIndex2==0){
              colWidth = colWidth - 75; // untuk mengurangi additional width yang datang tiba2 seperti syaiton, xixixi
            }else{
              colWidth = colWidth - 60; // untuk mengurangi additional width yang datang tiba2 seperti syaiton, xixixi
            }
            $(this).width(colWidth);
            $(this).addClass('tdindex-dbt-lv2-'+tdIndex2+'-'+colWidth);
          });
        });
      }, 10);
    }
  }
};
</script>
