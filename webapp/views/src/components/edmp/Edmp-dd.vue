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
                              v-model="store.ddVal.ddCountrySelected"
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
                              v-model="store.ddVal.ddBusinessSegmentSelected"
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
                              v-model="store.ddVal.ddSourceSystemSelected"
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
                              v-model="store.ddVal.ddClusterSelected"
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
                              v-model="store.ddVal.ddTierSelected"
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
                              v-model="store.ddVal.ddItamSelected"
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
        return _.sortedUniq(_.sortBy(_.map(this.store.DDSource, (v) => v.COUNTRY.toString()), [function(o) { return o; }]));
      },
      ddBusinessSegmentOptions () {
        var self = this;
        var filtered = _.filter(self.store.DDSource, (v) => {
          return self.store.ddVal.ddCountrySelected.length > 0 ? (self.store.ddVal.ddCountrySelected.includes(v.COUNTRY)) : true;
        });

        return _.sortedUniq(_.sortBy(_.map(filtered, (v) => v.BUSINESS_SEGMENT.toString()), [function(o) { return o; }]));
      },
      ddSourceSystemOptions () {
        var self = this;
        var filtered = _.filter(self.store.DDSource, (v) => {
          return self.store.ddVal.ddCountrySelected.length > 0 ? (self.store.ddVal.ddCountrySelected.includes(v.COUNTRY)) : true 
            && self.store.ddVal.ddBusinessSegmentSelected.length > 0 ? (self.store.ddVal.ddBusinessSegmentSelected.includes(v.BUSINESS_SEGMENT)) : true;
        });
        
        return _.sortedUniq(_.sortBy(_.map(filtered, (v) => v.SOURCE_SYSTEM.toString()), [function(o) { return o; }]));
      },
      ddClusterOptions () {
        var self = this;
        var filtered = _.filter(self.store.DDSource, (v) => {
          return self.store.ddVal.ddCountrySelected.length > 0 ? (self.store.ddVal.ddCountrySelected.includes(v.COUNTRY)) : true 
            && self.store.ddVal.ddBusinessSegmentSelected.length > 0 ? (self.store.ddVal.ddBusinessSegmentSelected.includes(v.BUSINESS_SEGMENT)) : true 
            && self.store.ddVal.ddSourceSystemSelected.length > 0 ? (self.store.ddVal.ddSourceSystemSelected.includes(v.SOURCE_SYSTEM)) : true;
        });
        
        return _.sortedUniq(_.sortBy(_.map(filtered, (v) => v.CLUSTER_NAME.toString()), [function(o) { return o; }]));
      },
      ddTierOptions () {
        var self = this;
        var filtered = _.filter(self.store.DDSource, (v) => {
          return self.store.ddVal.ddCountrySelected.length > 0 ? (self.store.ddVal.ddCountrySelected.includes(v.COUNTRY)) : true 
            && self.store.ddVal.ddBusinessSegmentSelected.length > 0 ? (self.store.ddVal.ddBusinessSegmentSelected.includes(v.BUSINESS_SEGMENT)) : true 
            && self.store.ddVal.ddSourceSystemSelected.length > 0 ? (self.store.ddVal.ddSourceSystemSelected.includes(v.SOURCE_SYSTEM)) : true
            && self.store.ddVal.ddClusterSelected.length > 0 ? (self.store.ddVal.ddClusterSelected.includes(v.CLUSTER_NAME)) : true;
        });
        
        return _.sortedUniq(_.sortBy(_.map(filtered, (v) => v.TIER.toString()), [function(o) { return o; }]));
      },
      ddItamOptions () {
        var self = this;
        var filtered = _.filter(self.store.DDSource, (v) => {
          return self.store.ddVal.ddCountrySelected.length > 0 ? (self.store.ddVal.ddCountrySelected.includes(v.COUNTRY)) : true 
            && self.store.ddVal.ddBusinessSegmentSelected.length > 0 ? (self.store.ddVal.ddBusinessSegmentSelected.includes(v.BUSINESS_SEGMENT)) : true 
            && self.store.ddVal.ddSourceSystemSelected.length > 0 ? (self.store.ddVal.ddSourceSystemSelected.includes(v.SOURCE_SYSTEM)) : true
            && self.store.ddVal.ddClusterSelected.length > 0 ? (self.store.ddVal.ddClusterSelected.includes(v.CLUSTER_NAME)) : true
            && self.store.ddVal.ddTierSelected.length > 0 ? (self.store.ddVal.ddTierSelected.includes(v.TIER)) : true;
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
      'store.ddVal.ddCountrySelected' () {
        if(this.store.firstload) return;

        this.store.ddVal.ddBusinessSegmentSelected = [];
        this.store.ddVal.ddSourceSystemSelected = [];
        this.store.ddVal.ddClusterSelected = [];
        this.store.ddVal.ddTierSelected = [];
        this.store.ddVal.ddItamSelected = [];

        this.store.firstload = true;

        this.refreshActiveTabTable();
      },
      'store.ddVal.ddBusinessSegmentSelected' () {
        if(this.store.firstload) return;

        this.store.ddVal.ddSourceSystemSelected = [];
        this.store.ddVal.ddClusterSelected = [];
        this.store.ddVal.ddTierSelected = [];
        this.store.ddVal.ddItamSelected = [];

        this.store.firstload = true;

        this.refreshActiveTabTable();
      },
      'store.ddVal.ddSourceSystemSelected' () {
        if(this.store.firstload) return;

        this.store.ddVal.ddClusterSelected = [];
        this.store.ddVal.ddTierSelected = [];
        this.store.ddVal.ddItamSelected = [];
        
        this.store.firstload = true;

        this.refreshActiveTabTable();
      },
      'store.ddVal.ddClusterSelected' () {
        if(this.store.firstload) return;

        this.store.ddVal.ddTierSelected = [];
        this.store.ddVal.ddItamSelected = [];

        this.store.firstload = true;

        this.refreshActiveTabTable();
      },
      'store.ddVal.ddTierSelected' () {
        if(this.store.firstload) return;

        this.store.ddVal.ddItamSelected = [];

        this.store.firstload = true;

        this.refreshActiveTabTable();
      },
      'store.ddVal.ddItamSelected' () {
        if(this.store.firstload) return;

        this.store.firstload = true;

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
        this.$store.dispatch(`${this.storeName}/getDropdownOpts`);
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

        this.technicalStore.filters.left["COUNTRY"] = this.store.ddVal.ddCountrySelected;
        this.technicalStore.filters.left.filterTypes["COUNTRY"] = "eq";

        this.technicalStore.filters.left["BUSINESS_SEGMENT"] = this.store.ddVal.ddBusinessSegmentSelected;
        this.technicalStore.filters.left.filterTypes["BUSINESS_SEGMENT"] = "eq";

        this.technicalStore.filters.left["SOURCE_SYSTEM"] = this.store.ddVal.ddSourceSystemSelected;
        this.technicalStore.filters.left.filterTypes["SOURCE_SYSTEM"] = "eq";

        this.technicalStore.filters.left["CLUSTER_NAME"] = this.store.ddVal.ddClusterSelected;
        this.technicalStore.filters.left.filterTypes["CLUSTER_NAME"] = "eq";

        this.technicalStore.filters.left["TIER"] = this.store.ddVal.ddTierSelected;
        this.technicalStore.filters.left.filterTypes["TIER"] = "eq";

        this.technicalStore.filters.left["ITAM"] = this.store.ddVal.ddItamSelected;
        this.technicalStore.filters.left.filterTypes["ITAM"] = "eq";

        this.$store.dispatch(`${this.technicalStoreName}/getLeftTable`).then(res => {
          this.store.firstload = false;
        });
      },
      refreshBusinessTable() {
        if( ! this.businessStore.filters.left.filterTypes) this.businessStore.filters.left.filterTypes = {};

        this.businessStore.filters.left["COUNTRY"] = this.store.ddVal.ddCountrySelected;
        this.businessStore.filters.left.filterTypes["COUNTRY"] = "eq";

        this.businessStore.filters.left["BUSINESS_SEGMENT"] = this.store.ddVal.ddBusinessSegmentSelected;
        this.businessStore.filters.left.filterTypes["BUSINESS_SEGMENT"] = "eq";

        this.businessStore.filters.left["SOURCE_SYSTEM"] = this.store.ddVal.ddSourceSystemSelected;
        this.businessStore.filters.left.filterTypes["SOURCE_SYSTEM"] = "eq";

        this.businessStore.filters.left["CLUSTER_NAME"] = this.store.ddVal.ddClusterSelected;
        this.businessStore.filters.left.filterTypes["CLUSTER_NAME"] = "eq";

        this.businessStore.filters.left["TIER"] = this.store.ddVal.ddTierSelected;
        this.businessStore.filters.left.filterTypes["TIER"] = "eq";

        this.businessStore.filters.left["ITAM"] = this.store.ddVal.ddItamSelected;
        this.businessStore.filters.left.filterTypes["ITAM"] = "eq";

        this.$store.dispatch(`${this.businessStoreName}/getLeftTable`).then(res => {
          this.store.firstload = false;
        });
      },
      refreshConsumptionTable() {
        if( ! this.consumptionStore.filters.left.filterTypes) this.consumptionStore.filters.left.filterTypes = {};

        this.consumptionStore.filters.left["COUNTRY"] = this.store.ddVal.ddCountrySelected;
        this.consumptionStore.filters.left.filterTypes["COUNTRY"] = "eq";

        this.consumptionStore.filters.left["BUSINESS_SEGMENT"] = this.store.ddVal.ddBusinessSegmentSelected;
        this.consumptionStore.filters.left.filterTypes["BUSINESS_SEGMENT"] = "eq";

        this.consumptionStore.filters.left["SOURCE_SYSTEM"] = this.store.ddVal.ddSourceSystemSelected;
        this.consumptionStore.filters.left.filterTypes["SOURCE_SYSTEM"] = "eq";

        this.consumptionStore.filters.left["CLUSTER_NAME"] = this.store.ddVal.ddClusterSelected;
        this.consumptionStore.filters.left.filterTypes["CLUSTER_NAME"] = "eq";

        this.consumptionStore.filters.left["TIER"] = this.store.ddVal.ddTierSelected;
        this.consumptionStore.filters.left.filterTypes["TIER"] = "eq";

        this.consumptionStore.filters.left["ITAM"] = this.store.ddVal.ddItamSelected;
        this.consumptionStore.filters.left.filterTypes["ITAM"] = "eq";

        this.$store.dispatch(`${this.consumptionStoreName}/getLeftTable`).then(res => {
          this.store.firstload = false;
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

        this.store.ddVal.ddCountrySelected = "";
        this.store.ddVal.ddBusinessSegmentSelected = "";
        this.store.ddVal.ddSourceSystemSelected = "";
        this.store.ddVal.ddClusterSelected = "";
        this.store.ddVal.ddTierSelected = "";
        this.store.ddVal.ddItamSelected = "";

        this.refreshActiveTabTable();
      },
    },
}
</script>