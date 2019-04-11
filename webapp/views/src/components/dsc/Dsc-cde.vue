<style>
table.v-table thead th > div.btn-group {
  	width: auto;
}

.header-filter-icon .dropdown-menu {
	overflow: scroll;
	height: 200px;
}
</style>

<template>
    <v-content>
        <b-container fluid>
            <PageHeader />

            <b-row>
                <b-col>
                    <page-loader v-if="store.left.isLoading" />

                    <b-row style="margin-top: 10px; margin-bottom: 20px;">
                        <b-col>
                            <v-btn class="float-right" color="red-neon" @click.native="resetFilter" dark>
                                <!-- <v-icon dark>filter_list</v-icon> -->
                                <i class="fa fa-filter"></i>
                            </v-btn>

                            <page-export class="float-right" :storeName="storeName" :leftTableCols="store.leftHeaders" :rightTableCols="[]"/>
                        </b-col>
                    </b-row>

                    <b-row style="margin-top: 10px;margin-bottom: 10px;">
                      <b-col>
                        <!-- Dsc details -->
                        <router-view/>

                        <!-- Main content -->
                        <b-row>
                          <b-col>
                            <page-loader v-if="store.left.isLoading"/>

                            <b-row>
                              <b-col cols="12">
                                <v-data-table
                                    :headers="store.leftHeaders.filter(v => v.display == true)"
                                    :items="store.left.display"
                                    :pagination.sync="store.left.pagination"
                                    :total-items="store.left.totalItems"
                                    :loading="store.left.isLoading"
                                    item-key="ID"
                                    class="card-content">
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
                                    <td>
                                      <b-link @click="showDetails(props.item)">
                                        <tablecell :fulltext="props.item.CDE" :isklik="false"></tablecell>
                                      </b-link>
                                    </td>
                                    <td><tablecell :fulltext="props.item.DESCRIPTION" :isklik="true"></tablecell></td>
                                    <td><tablecell :fulltext="props.item.TABLE_NAME" :isklik="true"></tablecell></td>
                                    <td><tablecell :fulltext="props.item.COLUMN_NAME" :isklik="true"></tablecell></td>
                                    <td><tablecell :fulltext="props.item.DSP_NAME" :isklik="true"></tablecell></td>
                                    <td><tablecell :fulltext="props.item.PROCESS_OWNER" :isklik="true"></tablecell></td>
                                  </template>
                                </v-data-table>
                              </b-col>
                            </b-row>
                          </b-col>
                        </b-row>
                      </b-col>
                    </b-row>
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

Vue.component("downloadExcel", JsonExcel);

export default {
  components: {
    PageHeader, pageSearch, pageExport, tableheader, tablecell, pageLoader
  },
  data() {
    return {
      storeName: "dsccde",
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
  },
  methods: {
    getLeftTable() {
      this.$store.dispatch(`${this.storeName}/getLeftTable`, this.$route.params.system);
    },
    systemRowClick(evt) {
      evt.preventDefault();
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
        this.addressPath + "/" + param.TSID + "/" + param.ID + "/" + param.COLID
      );
    }
  }
};
</script>
