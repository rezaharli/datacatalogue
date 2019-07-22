<style scoped>
ol {
    -webkit-column-count: 3;
       -moz-column-count: 3;
            column-count: 3;
    -webkit-column-gap: 20px;
       -moz-column-gap: 20px;
            column-gap: 20px;
}
</style>

<template>
  <v-content class="mx-4 my-5">
    <PageHeader title="Risk Framework Owner View" />
    <page-loader v-if="store.left.isLoading" />
    
    <b-container fluid class="row-kasijarak">

      <!-- <b-row>
          <b-col sm=12 md=3>
              <div class="card card-v2 transition">
                  <h6 class="title-1">Principal Risk Type</h6>
                  <h3 class="title-2 text-capitalize">{{ store.left.display[0] ? store.left.display[0].PRINCIPAL_RISK : "" }}</h3>
              </div>
          </b-col>
          
          <b-col sm=12 md=3>
              <div class="card card-v2 transition">
                  <h6 class="title-1">Risk Sub Type</h6>
                  <h3 class="title-2 text-capitalize">{{ $route.params.type }}</h3>
              </div>
          </b-col>
      </b-row> -->

      <b-row>
        <b-col cols="3">           
          <!-- <div class="card card-v2 transition mb-4">
              <h6 class="title-1">Principal Risk Type</h6>
              <h3 class="title-2 text-capitalize">{{ store.left.display[0] ? store.left.display[0].PRINCIPAL_RISK : "" }}</h3>
          </div>

          <div class="card card-v2 transition">
              <h6 class="title-1">Risk Sub Type</h6>
              <h3 class="title-2 text-capitalize">{{ $route.params.type }}</h3>
          </div> -->
          
          <b-card tag="article" class="mb-2">
            <b-media class="left-card-media mb-4" >
              <h6 class="left-card-title">Principal Risk Type</h6>
              <!-- <text-wrap-dialog :fulltext="store.left.display[0] ? store.left.display[0].PRINCIPAL_RISK: ''"></text-wrap-dialog> -->
              {{ store.left.display[0] ? store.left.display[0].PRINCIPAL_RISK : "" }}
            </b-media>
            
            <b-media class="left-card-media">
              <h6 class="left-card-title">Risk Sub Type</h6>
              <!-- <text-wrap-dialog :fulltext="store.left.display[0] ? store.left.display[0].RISK_SUB: ''"></text-wrap-dialog> -->
              {{ store.left.display[0] ? store.left.display[0].RISK_SUB : "" }}
            </b-media>
          </b-card>
        </b-col>

        <b-col cols="9"> 
          
          <b-row>
            <b-col>
              <b-card tag="article" class="mb-2">
                <h4 class="card-title border-0 mb-0 pb-0" v-b-toggle.collapse-1>
                  Reports
                  <i class="when-opened float-right fa fa-chevron-up"></i>
                  <i class="when-closed float-right fa fa-chevron-down"></i>
                </h4>
                <b-collapse id="collapse-1" class="mt-3 pt-4 border-top" visible>
                  <p class="card-text pl-2">
                    <ol>
                      <li v-bind:key="i" v-for="(pr,i) in rfoSummaries.PRIORITY_REPORTs">
                        {{ pr }}
                      </li>
                    </ol>
                  </p>
                </b-collapse>
              </b-card>
            </b-col>
          </b-row>

          <b-row>
            <b-col>
              <b-card tag="article" class="mb-2">
                <h4 class="card-title border-0 mb-0 pb-0" v-b-toggle.collapse-2>
                  Critical Risk Measures
                  <i class="when-opened float-right fa fa-chevron-up"></i>
                  <i class="when-closed float-right fa fa-chevron-down"></i>
                </h4>
                <b-collapse id="collapse-2" class="mt-3 pt-4 border-top">
                  <p class="card-text pl-2">
                    <ol>
                      <li v-bind:key="i" v-for="(crm,i) in rfoSummaries.CRM_NAMEs">
                        {{ crm }}
                      </li>
                    </ol>
                  </p>
                </b-collapse>
              </b-card>
            </b-col>
          </b-row>

          <b-row>
            <b-col>
              <b-card tag="article" class="mb-2">
                <h4 class="card-title border-0 mb-0 pb-0" v-b-toggle.collapse-3>
                  Critical Data Elements 
                  <i class="when-opened float-right fa fa-chevron-up"></i>
                  <i class="when-closed float-right fa fa-chevron-down"></i>
                </h4>
                <b-collapse id="collapse-3" class="mt-3 pt-4 border-top">
                  <p class="card-text pl-2">
                    <ol>
                      <li v-bind:key="i" v-for="(cde,i) in rfoSummaries.CDE_NAMEs">
                        {{ cde }}
                      </li>
                    </ol>
                  </p>
                </b-collapse>
              </b-card>
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
  name: "RfoSummary",
  components: {
    PageHeader, pageSearch, pageExport, tableheader, tablecell, pageLoader
  },
  data() {
    return {
      storeName: "rfosummary",
      rfoSummaries: {
        CDE_NAMEs : [],
        CRM_NAMEs : [],
        PRIORITY_REPORTs : [],
      },
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
    this.store.system = this.$route.params.type;
    this.getLeftTable();     
  },
  updated() {
  },
  methods: {
    getLeftTable() {
      var _self = this;
      var getLeftTableVal = this.$store.dispatch(`${this.storeName}/getLeftTable`);
      getLeftTableVal.then(res => {
        this.setData();
      });
    },
    setData() {
      var self = this;

      var CDE_NAMEs = [];
      var CRM_NAMEs = [];
      var PRIORITY_REPORTs = [];
      this.store.left.display.forEach(x => {  
        CDE_NAMEs.push(x.CDE_NAME);
        CRM_NAMEs.push(x.CRM_NAME);
        PRIORITY_REPORTs.push(x.PRIORITY_REPORT);
      });

      var uniqCDE_NAMEs = _.uniq(CDE_NAMEs);
      var uniqCRM_NAMEs = _.uniq(CRM_NAMEs);
      var uniqPRIORITY_REPORTs = _.uniq(PRIORITY_REPORTs);
      var rfoSummaries = {
        CDE_NAMEs : {},
        CRM_NAMEs : {},
        PRIORITY_REPORTs : {},
      };
      var i = 0;
      for (i = 0; i < uniqCDE_NAMEs.length; i++) {
        var item = uniqCDE_NAMEs[i];
        rfoSummaries.CDE_NAMEs[i] = item;
      }
      for (i = 0; i < uniqCRM_NAMEs.length; i++) {
        var item = uniqCRM_NAMEs[i];
        rfoSummaries.CRM_NAMEs[i] = item;
      }
      for (i = 0; i < uniqPRIORITY_REPORTs.length; i++) {
        var item = uniqPRIORITY_REPORTs[i];
        rfoSummaries.PRIORITY_REPORTs[i] = item;
      }
      self.rfoSummaries = rfoSummaries;
    },
  }
};
</script>
