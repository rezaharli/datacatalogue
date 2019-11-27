<style>
@import "../../assets/styles/dashboard.css";
#table-dsc-all table.v-table tr th:nth-of-type(1) {
  width: 60% !important;
}
#table-dsc-all table.v-table tr th:nth-of-type(2) {
  width: 40% !important;
}
</style>

<template>
  <b-row style="margin-top: 10px;margin-bottom: 10px;">
    <b-col>
      <!-- Dsc details -->
      <router-view />

      <!-- Main content -->
      <b-row>
        <b-col>
          <page-loader
            v-if="store.left.isLoading || (store.isRightTable && store.right.isLoading)"
          />

          <div class="card card-v1 transition">
            <div class="title-wrapper transition">
              <v-img :src="images.all" :contain="true" class="card-icon transition"></v-img>
              <h2 class="transition title-1">All Systems</h2>
            </div>

            <table-component :items="store.left.display" :opts="tableOpts" />

            <!-- <v-data-table
                :headers="store.leftHeaders.filter(v => v.display == true)"
                :items="store.left.display"
                :pagination.sync="store.left.pagination"
                :total-items="store.left.totalItems"
                :loading="store.left.isLoading"
                :expand="false"
                :must-sort="true"
                item-key="ID"
                class="card-content"
                id="table-dsc-all"
              >
              <template slot="headerCell" slot-scope="props">
                <tableheader :storeName="storeName" :props="props" :which="'left'" />
              </template>

              <v-progress-linear slot="progress" color="blue" indeterminate></v-progress-linear>

              <template slot="no-data">
                <v-alert :value="store.left.isLoading" type="info">Please wait while data is loading</v-alert>

                <v-alert :value="!store.left.isLoading" type="error">Sorry, nothing to display here</v-alert>
              </template>

              <template slot="items" slot-scope="props">
                <td v-bind:style="{ width: store.left.colWidth['SYSTEM_NAME'] + 'px' }">
                  <b-link @click="showRightTable(props.item)">
                    <tablecell :fulltext="props.item.SYSTEM_NAME" showOn="hover"></tablecell>
                  </b-link>
                </td>

                <td v-bind:style="{ width: store.left.colWidth['ITAM_ID'] + 'px' }">
                  <b-link @click.stop="toggleDscDrawer(props.item)">
                    <tablecell
                      :fulltext="(_.uniq(_.map(props.item.Custodians, 'ITAM_ID').filter(Boolean)).join(', '))"
                      showOn="hover"
                    ></tablecell>
                  </b-link>
                </td>
              </template>
            </v-data-table> -->
          </div>
        </b-col>
      </b-row>
    </b-col>
  </b-row>
</template>

<script>
import Vue from "vue";
import { mapState, mapActions } from "vuex";
import JsonExcel from "vue-json-excel";

import pageLoader from "@/components/PageLoader.vue";
import pageSearch from "@/components/PageSearch.vue";
import pageExport from "@/components/PageExport.vue";

import tableComponent from "@/components/TableComponent.vue"
import tableheader from "@/components/TableHeader.vue";
import tablecell from "@/components/Tablecell.vue";

Vue.component("downloadExcel", JsonExcel);

export default {
  components: {
    pageSearch,
    pageExport,
    tableComponent,
    tableheader,
    tablecell,
    pageLoader
  },
  data() {
    return {
      storeName: "dscall",
      systemSource: [],
      tablenameSource: [],
      images: {
        all: require("../../assets/images/icon-all-system.png")
      }
    };
  },
  computed: {
    store() {
      return this.$store.state[this.storeName].all;
    },
    dscStore() {
      return this.$store.state.dsc.all;
    },
    tableOpts() {
      return {
        headers: this.store.leftHeaders,
        pagination: this.store.left.pagination,
        totalItems: this.store.left.totalItems,
        isLoading: this.store.left.isLoading
      }
    },
    addressPath() {
      var tmp = this.$route.path.split("/");
      return tmp.slice(0, 3).join("/");
    }
  },
  watch: {
    $route(to) {
      this.store.isRightTable = false;

      if (to.params != undefined) {
        this.store.isRightTable = to.params.system;
      }

      if (this.store.isRightTable) {
        this.doGetRightTable(this.$route.params.system);
      }
    },
    "store.left.pagination": {
      handler() {
        this.doGetLeftTable();
      },
      deep: true
    },
    "store.right.pagination": {
      handler() {
        if (this.store.isRightTable) {
          this.doGetRightTable(this.$route.params.system);
        }
      },
      deep: true
    },
    "store.searchMain"(val, oldVal) {
      if (val || oldVal) {
        this.doGetLeftTable();

        if (this.store.isRightTable) {
          this.doGetRightTable(this.$route.params.system);
        }
      }
    }
  },
  mounted() {
    this.store.tabName = this.storeName;
    this.store.isRightTable = this.$route.params.system;
  },
  updated() {},
  methods: {
    getLeftTable() {
      var getLeftTableVal = this.$store.dispatch(
        `${this.storeName}/getLeftTable`
      );
      getLeftTableVal.then(res => {
        this.removeHypenOnEmptyTables($("#table-dsc-all"));
      });
    },
    getRightTable(id) {
      this.$store.dispatch(`${this.storeName}/getRightTable`, id);
    },
    doGetLeftTable() {
      this.getLeftTable();
    },
    doGetRightTable(id) {
      this.getRightTable(id);
    },
    showRightTable(param) {
      //reset right table filter
      this.store.filters.right = {};

      if (
        param.SYSTEM_NAME.toUpperCase() ==
        "ENTERPRISE DATA MGMT PLATFORM".toUpperCase()
      ) {
        this.$router.push(this.addressPath + "/" + param.SYSTEM_NAME);
        return;
      }

      this.$router.push(
        this.addressPath + "/" + encodeURIComponent(param.SYSTEM_NAME)
      );
    },
    toggleDscDrawer(param) {
      this.dscStore.drawer = !this.dscStore.drawer;

      if (this.dscStore.drawer) {
        this.dscStore.drawerContent.systemName = param.SYSTEM_NAME;
        this.dscStore.drawerContent.itamID = param.ITAM_ID;
        this.dscStore.drawerContent.owners = param.Custodians.map(v => {
          return {
            BANK_ID: v.BANK_ID,
            DATASET_CUSTODIAN: v.DATASET_CUSTODIAN
          };
        });
      } else {
        this.dscStore.drawerContent.systemName = "";
        this.dscStore.drawerContent.itamID = "";
        this.dscStore.drawerContent.owners = [];
      }
    },
    removeHypenOnEmptyTables(elem) {
      var paginationElem = elem.find(
        ".v-datatable__actions .v-datatable__actions__range-controls .v-datatable__actions__pagination"
      );
      paginationElem.each(function() {
        var paginationText = $(this).text();
        if (paginationText == "–") {
          $(this).hide();
        } else {
          $(this).show();
        }
      });
    }
  }
};
</script>
