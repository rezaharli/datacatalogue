<style>
  @import '../../assets/styles/dashboard.css';
</style>

<style>
.dd-filter .v-text-field--box .v-label {
  font-size: 14px;
}
.v-select__selection--comma {
  text-overflow: ellipsis;
  white-space: nowrap;
  overflow: hidden;
  display: inline-block;
}

.v-select__selections{
  max-width: 70%;
}

.v-select.v-text-field--enclosed:not(.v-text-field--single-line) .v-select__selections{
  padding-top: 30px;
  min-height: unset !important;
}

.v-select.v-select--chips input{
  position: absolute;
}
</style>

<template>
    <v-content>
        <b-container fluid>
            <PageHeader />
            
            <b-row class="my-4">
                <b-col class="ml-5 col-md-10">
                  <b-row class="ml-3 dd-filter">
                    <b-col cols="2">
                      <v-tooltip top>
                        <template slot="activator" slot-scope="{ on }">
                          <div v-on="on">
                            <v-select
                              v-model="store.dd.ddVal.ddCountrySelected"
                              :items="ddCountryOptions"
                              label="Country"
                              single-line
                              multiple
                              box
                              clearable
                            ></v-select>
                          </div>
                        </template>

                        <span>Country</span>
                      </v-tooltip>
                    </b-col>

                    <b-col cols="2">
                      <v-tooltip top>
                        <template slot="activator" slot-scope="{ on }">
                          <div v-on="on">
                            <v-select
                              v-model="store.dd.ddVal.ddBusinessSegmentSelected"
                              :items="ddBusinessSegmentOptions"
                              label="Business Segment"
                              single-line
                              multiple
                              box
                              clearable
                            ></v-select>
                          </div>
                        </template>
                        
                        <span>Business Segment</span>
                      </v-tooltip>
                    </b-col>

                    <b-col cols="2">
                      <v-tooltip top>
                        <template slot="activator" slot-scope="{ on }">
                          <div v-on="on">
                            <v-select
                              v-model="store.dd.ddVal.ddSourceSystemSelected"
                              :items="ddSourceSystemOptions"
                              label="Source System"
                              single-line
                              multiple
                              box
                              clearable
                            ></v-select>
                          </div>
                        </template>
                        
                        <span>Source System</span>
                      </v-tooltip>
                    </b-col>

                    <b-col cols="2">
                      <v-tooltip top>
                        <template slot="activator" slot-scope="{ on }">
                          <div v-on="on">
                            <v-select
                              v-model="store.dd.ddVal.ddClusterSelected"
                              :items="ddClusterOptions"
                              label="Cluster"
                              single-line
                              multiple
                              box
                              clearable
                            ></v-select>
                          </div>
                        </template>
                        
                        <span>Cluster</span>
                      </v-tooltip>
                    </b-col>

                    <b-col cols="2">
                      <v-tooltip top>
                        <template slot="activator" slot-scope="{ on }">
                          <div v-on="on">
                            <v-select
                              v-model="store.dd.ddVal.ddTierSelected"
                              :items="ddTierOptions"
                              label="Tier"
                              single-line
                              multiple
                              box
                              clearable
                            ></v-select>
                          </div>
                        </template>
                        
                        <span>Tier</span>
                      </v-tooltip>
                    </b-col>

                    <b-col cols="2">
                      <v-tooltip top>
                        <template slot="activator" slot-scope="{ on }">
                          <div v-on="on">
                            <v-select
                              v-model="store.dd.ddVal.ddItamSelected"
                              :items="ddItamOptions"
                              label="ITAM"
                              single-line
                              multiple
                              box
                              clearable
                            ></v-select>
                          </div>
                        </template>
                        
                        <span>ITAM</span>
                      </v-tooltip>
                    </b-col>
                  </b-row>
                  
                </b-col>
                
                <b-col>
                  <b-button class="float-right red-neon icon-only ml-3 shadow-sm" @click="resetFilter">
                    <i class="fa fa-filter"></i>
                  </b-button>

                  <page-export class="float-right shadow-sm" :storeName="activeTabStoreName" :leftTableCols="activeTabStore.leftHeaders" :rightTableCols="[]" :rowSelectInvolved="true" />
                </b-col>
            </b-row>

            <b-row>
                <b-col>
                    <v-tabs id="page-tab" class="page-tab" v-model="activeTab">
                        <v-tab v-for="tab in tabs" class="px-2 mx-5" :id="'tab-' + tab.id" :key="tab.key" :to="addressPath + '/' + tab.key" :ref="tab.id">{{ tab.name }}</v-tab>
                    </v-tabs>
                    
                    <transition name="fade" mode="out-in">
                      <router-view />
                    </transition>
                </b-col>
            </b-row>
        </b-container>
    </v-content>
</template>

<script>
import PageHeader from '../PageHeader';
import pageExport from "../PageExport.vue";

export default {
    components: { PageHeader, pageExport },
    data() {
      return {
        storeName: "edmp",
        technicalStoreName: "edmpddTechnical",
        businessStoreName: "edmpddBusiness",
        consumptionStoreName: "edmpddConsumption",
        activeTab: '',
        tabs: [
            { id: 'technical', key: 'technical-metadata', name: 'Technical Metadata', route: this.addressPath + '/technical-metadata' },
            { id: 'business', key: 'business-metadata', name: 'Business Metadata', route: this.addressPath + '/business-metadata' },
            { id: 'consumption', key: 'consumption-apps', name: 'Consumption Applications', route: this.addressPath + '/consumption-apps' },
        ]
      };
    },
    computed: {
      store() {
        return this.$store.state[this.storeName].all;
      },
      technicalStore() {
        return this.$store.state[this.technicalStoreName].all;
      },
      businessStore() {
        return this.$store.state[this.businessStoreName].all;
      },
      consumptionStore() {
        return this.$store.state[this.consumptionStoreName].all;
      },
      activeTabStoreName() {
        if(this.activeTab.indexOf("technical-metadata") != -1){
          return this.technicalStoreName;
        } else if(this.activeTab.indexOf("business-metadata") != -1){
          return this.businessStoreName;
        } else if(this.activeTab.indexOf("consumption-apps") != -1){
          return this.consumptionStoreName;
        } else {
          return this.technicalStoreName;
        }
      },
      activeTabStore() {
        return this.$store.state[this.activeTabStoreName].all;
      },
      addressPath() {
        var tmp = this.$route.path.split("/");
        return tmp.slice(0, 4).join("/");
      },
      urlParam1() {
        return this.$route.params.system;
      },
      ddCountryOptions () {
        return _.sortedUniq(_.sortBy(_.map(this.store.dd.DDSource, (v) => v.COUNTRY.toString()), [function(o) { return o; }]));
      },
      ddBusinessSegmentOptions () {
        var self = this;
        var filtered = _.filter(self.store.dd.DDSource, (v) => {
          return self.store.dd.ddVal.ddCountrySelected.length > 0 ? (self.store.dd.ddVal.ddCountrySelected.includes(v.COUNTRY)) : true;
        });

        return _.sortedUniq(_.sortBy(_.map(filtered, (v) => v.BUSINESS_SEGMENT.toString()), [function(o) { return o; }]));
      },
      ddSourceSystemOptions () {
        var self = this;
        var filtered = _.filter(self.store.dd.DDSource, (v) => {
          return self.store.dd.ddVal.ddCountrySelected.length > 0 ? (self.store.dd.ddVal.ddCountrySelected.includes(v.COUNTRY)) : true 
            && self.store.dd.ddVal.ddBusinessSegmentSelected.length > 0 ? (self.store.dd.ddVal.ddBusinessSegmentSelected.includes(v.BUSINESS_SEGMENT)) : true;
        });
        
        return _.sortedUniq(_.sortBy(_.map(filtered, (v) => v.EDM_SOURCE_SYSTEM_NAME.toString()), [function(o) { return o; }]));
      },
      ddClusterOptions () {
        var self = this;
        var filtered = _.filter(self.store.dd.DDSource, (v) => {
          return self.store.dd.ddVal.ddCountrySelected.length > 0 ? (self.store.dd.ddVal.ddCountrySelected.includes(v.COUNTRY)) : true 
            && self.store.dd.ddVal.ddBusinessSegmentSelected.length > 0 ? (self.store.dd.ddVal.ddBusinessSegmentSelected.includes(v.BUSINESS_SEGMENT)) : true 
            && self.store.dd.ddVal.ddSourceSystemSelected.length > 0 ? (self.store.dd.ddVal.ddSourceSystemSelected.includes(v.EDM_SOURCE_SYSTEM_NAME)) : true;
        });
        
        return _.sortedUniq(_.sortBy(_.map(filtered, (v) => v.CLUSTER_NAME.toString()), [function(o) { return o; }]));
      },
      ddTierOptions () {
        var self = this;
        var filtered = _.filter(self.store.dd.DDSource, (v) => {
          return self.store.dd.ddVal.ddCountrySelected.length > 0 ? (self.store.dd.ddVal.ddCountrySelected.includes(v.COUNTRY)) : true 
            && self.store.dd.ddVal.ddBusinessSegmentSelected.length > 0 ? (self.store.dd.ddVal.ddBusinessSegmentSelected.includes(v.BUSINESS_SEGMENT)) : true 
            && self.store.dd.ddVal.ddSourceSystemSelected.length > 0 ? (self.store.dd.ddVal.ddSourceSystemSelected.includes(v.EDM_SOURCE_SYSTEM_NAME)) : true
            && self.store.dd.ddVal.ddClusterSelected.length > 0 ? (self.store.dd.ddVal.ddClusterSelected.includes(v.CLUSTER_NAME)) : true;
        });
        
        return _.sortedUniq(_.sortBy(_.map(filtered, (v) => v.TIER.toString()), [function(o) { return o; }]));
      },
      ddItamOptions () {
        var self = this;
        var filtered = _.filter(self.store.dd.DDSource, (v) => {
          return self.store.dd.ddVal.ddCountrySelected.length > 0 ? (self.store.dd.ddVal.ddCountrySelected.includes(v.COUNTRY)) : true 
            && self.store.dd.ddVal.ddBusinessSegmentSelected.length > 0 ? (self.store.dd.ddVal.ddBusinessSegmentSelected.includes(v.BUSINESS_SEGMENT)) : true 
            && self.store.dd.ddVal.ddSourceSystemSelected.length > 0 ? (self.store.dd.ddVal.ddSourceSystemSelected.includes(v.EDM_SOURCE_SYSTEM_NAME)) : true
            && self.store.dd.ddVal.ddClusterSelected.length > 0 ? (self.store.dd.ddVal.ddClusterSelected.includes(v.CLUSTER_NAME)) : true
            && self.store.dd.ddVal.ddTierSelected.length > 0 ? (self.store.dd.ddVal.ddTierSelected.includes(v.TIER)) : true;
        });
        
        return _.sortedUniq(_.sortBy(_.map(filtered, (v) => v.ITAM.toString()), [function(o) { return o; }]));
      },
    },
    watch: {
      $route(to) {},
      "store.searchMain"(val, oldVal) {
        if (val || oldVal) {
          this.getLeftTable();
        }
      },
      'store.dd.ddVal.ddCountrySelected' () {
        if(this.store.dd.firstload) return;

        this.store.dd.ddVal.ddBusinessSegmentSelected = [];
        this.store.dd.ddVal.ddSourceSystemSelected = [];
        this.store.dd.ddVal.ddClusterSelected = [];
        this.store.dd.ddVal.ddTierSelected = [];
        this.store.dd.ddVal.ddItamSelected = [];

        this.store.dd.firstload = true;

        this.refreshActiveTabTable();
      },
      'store.dd.ddVal.ddBusinessSegmentSelected' () {
        if(this.store.dd.firstload) return;

        this.store.dd.ddVal.ddSourceSystemSelected = [];
        this.store.dd.ddVal.ddClusterSelected = [];
        this.store.dd.ddVal.ddTierSelected = [];
        this.store.dd.ddVal.ddItamSelected = [];

        this.store.dd.firstload = true;

        this.refreshActiveTabTable();
      },
      'store.dd.ddVal.ddSourceSystemSelected' () {
        if(this.store.dd.firstload) return;

        this.store.dd.ddVal.ddClusterSelected = [];
        this.store.dd.ddVal.ddTierSelected = [];
        this.store.dd.ddVal.ddItamSelected = [];
        
        this.store.dd.firstload = true;

        this.refreshActiveTabTable();
      },
      'store.dd.ddVal.ddClusterSelected' () {
        if(this.store.dd.firstload) return;

        this.store.dd.ddVal.ddTierSelected = [];
        this.store.dd.ddVal.ddItamSelected = [];

        this.store.dd.firstload = true;

        this.refreshActiveTabTable();
      },
      'store.dd.ddVal.ddTierSelected' () {
        if(this.store.dd.firstload) return;

        this.store.dd.ddVal.ddItamSelected = [];

        this.store.dd.firstload = true;

        this.refreshActiveTabTable();
      },
      'store.dd.ddVal.ddItamSelected' () {
        if(this.store.dd.firstload) return;

        this.store.dd.firstload = true;

        this.refreshActiveTabTable();
      },
    },
    mounted() {
      this.store.tabName = this.storeName;
      this.store.system = this.$route.params.system;

      this.getDropdownOpts();
    },
    methods: {
      getDropdownOpts() {
        this.$store.dispatch(`${this.storeName}/getDdDropdownOpts`);
      },
      getLeftTable() {
        this.$store.dispatch(`${this.storeName}/getLeftTable`);
      },
      refreshActiveTabTable() {
        if(this.activeTab.indexOf("technical-metadata") != -1){
          this.refreshTechnicalTable();
        }
        if(this.activeTab.indexOf("business-metadata") != -1){
          this.refreshBusinessTable();
        }
        if(this.activeTab.indexOf("consumption-apps") != -1){
          this.refreshConsumptionTable();
        }
      },
      refreshTechnicalTable() {
        if( ! this.technicalStore.filters.left.filterTypes) this.technicalStore.filters.left.filterTypes = {};

        this.technicalStore.filters.left["COUNTRY"] = this.store.dd.ddVal.ddCountrySelected;
        this.technicalStore.filters.left.filterTypes["COUNTRY"] = "eq";

        this.technicalStore.filters.left["BUSINESS_SEGMENT"] = this.store.dd.ddVal.ddBusinessSegmentSelected;
        this.technicalStore.filters.left.filterTypes["BUSINESS_SEGMENT"] = "eq";

        this.technicalStore.filters.left["EDM_SOURCE_SYSTEM_NAME"] = this.store.dd.ddVal.ddSourceSystemSelected;
        this.technicalStore.filters.left.filterTypes["EDM_SOURCE_SYSTEM_NAME"] = "eq";

        this.technicalStore.filters.left["CLUSTER_NAME"] = this.store.dd.ddVal.ddClusterSelected;
        this.technicalStore.filters.left.filterTypes["CLUSTER_NAME"] = "eq";

        this.technicalStore.filters.left["TIER"] = this.store.dd.ddVal.ddTierSelected;
        this.technicalStore.filters.left.filterTypes["TIER"] = "eq";

        this.technicalStore.filters.left["ITAM"] = this.store.dd.ddVal.ddItamSelected;
        this.technicalStore.filters.left.filterTypes["ITAM"] = "eq";

        this.$store.dispatch(`${this.technicalStoreName}/getLeftTable`).then(res => {
          this.store.dd.firstload = false;
        });
      },
      refreshBusinessTable() {
        if( ! this.businessStore.filters.left.filterTypes) this.businessStore.filters.left.filterTypes = {};

        this.businessStore.filters.left["COUNTRY"] = this.store.dd.ddVal.ddCountrySelected;
        this.businessStore.filters.left.filterTypes["COUNTRY"] = "eq";

        this.businessStore.filters.left["BUSINESS_SEGMENT"] = this.store.dd.ddVal.ddBusinessSegmentSelected;
        this.businessStore.filters.left.filterTypes["BUSINESS_SEGMENT"] = "eq";

        this.businessStore.filters.left["EDM_SOURCE_SYSTEM_NAME"] = this.store.dd.ddVal.ddSourceSystemSelected;
        this.businessStore.filters.left.filterTypes["EDM_SOURCE_SYSTEM_NAME"] = "eq";

        this.businessStore.filters.left["CLUSTER_NAME"] = this.store.dd.ddVal.ddClusterSelected;
        this.businessStore.filters.left.filterTypes["CLUSTER_NAME"] = "eq";

        this.businessStore.filters.left["TIER"] = this.store.dd.ddVal.ddTierSelected;
        this.businessStore.filters.left.filterTypes["TIER"] = "eq";

        this.businessStore.filters.left["ITAM"] = this.store.dd.ddVal.ddItamSelected;
        this.businessStore.filters.left.filterTypes["ITAM"] = "eq";

        this.$store.dispatch(`${this.businessStoreName}/getLeftTable`).then(res => {
          this.store.dd.firstload = false;
        });
      },
      refreshConsumptionTable() {
        if( ! this.consumptionStore.filters.left.filterTypes) this.consumptionStore.filters.left.filterTypes = {};

        this.consumptionStore.filters.left["COUNTRY"] = this.store.dd.ddVal.ddCountrySelected;
        this.consumptionStore.filters.left.filterTypes["COUNTRY"] = "eq";

        this.consumptionStore.filters.left["BUSINESS_SEGMENT"] = this.store.dd.ddVal.ddBusinessSegmentSelected;
        this.consumptionStore.filters.left.filterTypes["BUSINESS_SEGMENT"] = "eq";

        this.consumptionStore.filters.left["EDM_SOURCE_SYSTEM_NAME"] = this.store.dd.ddVal.ddSourceSystemSelected;
        this.consumptionStore.filters.left.filterTypes["EDM_SOURCE_SYSTEM_NAME"] = "eq";

        this.consumptionStore.filters.left["CLUSTER_NAME"] = this.store.dd.ddVal.ddClusterSelected;
        this.consumptionStore.filters.left.filterTypes["CLUSTER_NAME"] = "eq";

        this.consumptionStore.filters.left["TIER"] = this.store.dd.ddVal.ddTierSelected;
        this.consumptionStore.filters.left.filterTypes["TIER"] = "eq";

        this.consumptionStore.filters.left["ITAM"] = this.store.dd.ddVal.ddItamSelected;
        this.consumptionStore.filters.left.filterTypes["ITAM"] = "eq";

        this.$store.dispatch(`${this.consumptionStoreName}/getLeftTable`).then(res => {
          this.store.dd.firstload = false;
        });
      },
      updateRouter(val){
        this.$router.push(val);
      },
      resetFilter (e) {
        if(this.activeTab.indexOf("technical-metadata") != -1){
          if(Object.keys(this.technicalStore.filters.left).length > 0){
            this.technicalStore.filters.left = {};
          }
        }
        if(this.activeTab.indexOf("business-metadata") != -1){
          if(Object.keys(this.businessStore.filters.left).length > 0){
            this.businessStore.filters.left = {};
          }
        }
        if(this.activeTab.indexOf("consumption-apps") != -1){
          if(Object.keys(this.consumptionStore.filters.left).length > 0){
            this.consumptionStore.filters.left = {};
          }
        }

        this.store.dd.ddVal.ddCountrySelected = "";
        this.store.dd.ddVal.ddBusinessSegmentSelected = "";
        this.store.dd.ddVal.ddSourceSystemSelected = "";
        this.store.dd.ddVal.ddClusterSelected = "";
        this.store.dd.ddVal.ddTierSelected = "";
        this.store.dd.ddVal.ddItamSelected = "";

        this.refreshActiveTabTable();
      },
    },
}
</script>