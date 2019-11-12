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
                      <global-filter-dropdown label="Country" 
                        v-model="store.dd.ddVal.ddCountrySelected"
                        :items="ddCountryOptions"
                      />
                    </b-col>

                    <b-col cols="2">
                      <global-filter-dropdown label="Business Segment" 
                        v-model="store.dd.ddVal.ddBusinessSegmentSelected"
                        :items="ddBusinessSegmentOptions"
                      />
                    </b-col>

                    <b-col cols="2">
                      <global-filter-dropdown label="Source System" 
                        v-model="store.dd.ddVal.ddSourceSystemSelected"
                        :items="ddSourceSystemOptions"
                      />
                    </b-col>

                    <b-col cols="2">
                      <global-filter-dropdown label="Cluster" 
                        v-model="store.dd.ddVal.ddClusterSelected"
                        :items="ddClusterOptions"
                      />
                    </b-col>

                    <b-col cols="2">
                      <global-filter-dropdown label="Tier" 
                        v-model="store.dd.ddVal.ddTierSelected"
                        :items="ddTierOptions"
                      />
                    </b-col>

                    <b-col cols="2">
                      <global-filter-dropdown label="ITAM" 
                        v-model="store.dd.ddVal.ddItamSelected"
                        :items="ddItamOptions"
                      />
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
import globalFilterDropdown from "./GlobalFilterDropdown.vue";

export default {
    components: { PageHeader, pageExport, globalFilterDropdown },
    data() {
      return {
        storeName: "edmp",
        technicalStoreName: "edmpddTechnical",
        businessStoreName: "edmpddBusiness",
        consumptionStoreName: "edmpddConsumption",
        ddCountryOptions: [],
        ddBusinessSegmentOptions: [],
        ddSourceSystemOptions: [],
        ddClusterOptions: [],
        ddTierOptions: [],
        ddItamOptions: [],
        activeTab: '',
        tabs: [
            { id: 'technical', key: 'technical-metadata', name: 'Technical Metadata', route: this.addressPath + '/technical-metadata' },
            { id: 'business', key: 'business-metadata', name: 'Business Metadata', route: this.addressPath + '/business-metadata' },
            // { id: 'consumption', key: 'consumption-apps', name: 'Consumption Applications', route: this.addressPath + '/consumption-apps' },
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
      isGlobalFilterEmpty() {
        return this.store.dd.ddVal.ddCountrySelected.length == 0
          && this.store.dd.ddVal.ddBusinessSegmentSelected.length == 0
          && this.store.dd.ddVal.ddSourceSystemSelected.length == 0
          && this.store.dd.ddVal.ddClusterSelected.length == 0
          && this.store.dd.ddVal.ddTierSelected.length == 0
          && this.store.dd.ddVal.ddItamSelected.length == 0;
      },
      addressPath() {
        var tmp = this.$route.path.split("/");
        return tmp.slice(0, 4).join("/");
      },
      urlParam1() {
        return this.$route.params.system;
      },
    },
    watch: {
      $route(to) {},
      "store.searchMain"(val, oldVal) {
        if (val || oldVal) {
          this.getLeftTable();
        }
      },
      'store.dd.ddVal.ddCountrySelected'(val) {
        if(this.store.dd.firstload) return;

        this.setDdBusinessSegmentOptions();
        this.setDdSourceSystemOptions();
        this.setDdClusterOptions();
        this.setDdTierOptions();
        this.setDdItamOptions();

        this.store.dd.firstload = true;

        setTimeout(() => {
          this.refreshActiveTabTable("COUNTRY", val);
        }, 0);
      },
      'store.dd.ddVal.ddBusinessSegmentSelected'(val) {
        if(this.store.dd.firstload) return;
        
        this.setDdCountryOptions();
        this.setDdSourceSystemOptions();
        this.setDdClusterOptions();
        this.setDdTierOptions();
        this.setDdItamOptions();

        this.store.dd.firstload = true;

        setTimeout(() => {
          this.refreshActiveTabTable("BUSINESS_SEGMENT", val);
        }, 0);
      },
      'store.dd.ddVal.ddSourceSystemSelected'(val) {
        if(this.store.dd.firstload) return;

        this.setDdCountryOptions();
        this.setDdBusinessSegmentOptions();
        this.setDdClusterOptions();
        this.setDdTierOptions();
        this.setDdItamOptions();
        
        this.store.dd.firstload = true;

        setTimeout(() => {
          this.refreshActiveTabTable("EDM_SOURCE_SYSTEM_NAME", val);
        }, 0);
      },
      'store.dd.ddVal.ddClusterSelected'(val) {
        if(this.store.dd.firstload) return;

        this.setDdCountryOptions();
        this.setDdBusinessSegmentOptions();
        this.setDdSourceSystemOptions();
        this.setDdTierOptions();
        this.setDdItamOptions();

        this.store.dd.firstload = true;

        setTimeout(() => {
          this.refreshActiveTabTable("CLUSTER_NAME", val);
        }, 0);
      },
      'store.dd.ddVal.ddTierSelected'(val) {
        if(this.store.dd.firstload) return;

        this.setDdCountryOptions();
        this.setDdBusinessSegmentOptions();
        this.setDdSourceSystemOptions();
        this.setDdClusterOptions();
        this.setDdItamOptions();

        this.store.dd.firstload = true;

        setTimeout(() => {
          this.refreshActiveTabTable("TIER", val);
        }, 0);
      },
      'store.dd.ddVal.ddItamSelected'(val) {
        if(this.store.dd.firstload) return;

        this.setDdCountryOptions();
        this.setDdBusinessSegmentOptions();
        this.setDdSourceSystemOptions();
        this.setDdClusterOptions();
        this.setDdTierOptions();

        this.store.dd.firstload = true;

        setTimeout(() => {
          this.refreshActiveTabTable("ITAM", val);
        }, 0);
      },
    },
    mounted() {
      this.store.tabName = this.storeName;
      this.store.system = this.$route.params.system;

      this.getDropdownOpts();
    },
    methods: {
      setNewDropdownOpts() {
        this.ddCountryOptions = _.sortedUniq(_.sortBy(_.map(this.store.dd.DDSource, (v) => v.COUNTRY.toString()), [function(o) { return o; }]));
        this.ddBusinessSegmentOptions = _.sortedUniq(_.sortBy(_.map(this.store.dd.DDSource, (v) => v.BUSINESS_SEGMENT.toString()), [function(o) { return o; }]));
        this.ddSourceSystemOptions = _.sortedUniq(_.sortBy(_.map(this.store.dd.DDSource, (v) => v.EDM_SOURCE_SYSTEM_NAME.toString()), [function(o) { return o; }]));
        this.ddClusterOptions = _.sortedUniq(_.sortBy(_.map(this.store.dd.DDSource, (v) => v.CLUSTER_NAME.toString()), [function(o) { return o; }]));
        this.ddTierOptions = _.sortedUniq(_.sortBy(_.map(this.store.dd.DDSource, (v) => v.TIER.toString()), [function(o) { return o; }]));
        this.ddItamOptions = _.sortedUniq(_.sortBy(_.map(this.store.dd.DDSource, (v) => v.ITAM.toString()), [function(o) { return o; }]));
        
      },
      getDropdownOpts() {
        this.$store.dispatch(`${this.storeName}/getDdDropdownOpts`).then(() => {
          this.setNewDropdownOpts();
        });
      },
      getLeftTable() {
        this.$store.dispatch(`${this.storeName}/getLeftTable`);
      },
      refreshActiveTabTable(updatedAttr, val) {
        if( ! this.technicalStore.filters.left.filterTypes) this.technicalStore.filters.left.filterTypes = {};
        if( ! this.businessStore.filters.left.filterTypes) this.businessStore.filters.left.filterTypes = {};
        if( ! this.consumptionStore.filters.left.filterTypes) this.consumptionStore.filters.left.filterTypes = {};

        this.technicalStore.filters.left[updatedAttr] = val;
        this.technicalStore.filters.left.filterTypes[updatedAttr] = "eq";

        this.businessStore.filters.left[updatedAttr] = val;
        this.businessStore.filters.left.filterTypes[updatedAttr] = "eq";

        this.consumptionStore.filters.left[updatedAttr] = val;
        this.consumptionStore.filters.left.filterTypes[updatedAttr] = "eq";

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
        this.$store.dispatch(`${this.technicalStoreName}/getLeftTable`).then(res => {
          this.store.dd.firstload = false;
        });
      },
      refreshBusinessTable() {
        this.$store.dispatch(`${this.businessStoreName}/getLeftTable`).then(res => {
          this.store.dd.firstload = false;
        });
      },
      refreshConsumptionTable() {
        this.$store.dispatch(`${this.consumptionStoreName}/getLeftTable`).then(res => {
          this.store.dd.firstload = false;
        });
      },
      setDdCountryOptions () {
        var self = this;
        var filtered = _.filter(self.store.dd.DDSource, (v) => {
          if(this.isGlobalFilterEmpty) return true;

          return (self.store.dd.ddVal.ddCountrySelected.length > 0 ? (self.store.dd.ddVal.ddCountrySelected.includes(v.COUNTRY)) : false) 
            || (self.store.dd.ddVal.ddBusinessSegmentSelected.length > 0 ? (self.store.dd.ddVal.ddBusinessSegmentSelected.includes(v.BUSINESS_SEGMENT)) : false) 
            || (self.store.dd.ddVal.ddSourceSystemSelected.length > 0 ? (self.store.dd.ddVal.ddSourceSystemSelected.includes(v.EDM_SOURCE_SYSTEM_NAME)) : false)
            || (self.store.dd.ddVal.ddClusterSelected.length > 0 ? (self.store.dd.ddVal.ddClusterSelected.includes(v.CLUSTER_NAME)) : false)
            || (self.store.dd.ddVal.ddTierSelected.length > 0 ? (self.store.dd.ddVal.ddTierSelected.includes(v.TIER)) : false)
            || (self.store.dd.ddVal.ddItamSelected.length > 0 ? (self.store.dd.ddVal.ddItamSelected.includes(v.ITAM)) : false);
        });
        
        this.ddCountryOptions = _.sortedUniq(_.sortBy(_.map(filtered, (v) => v.COUNTRY.toString()), [function(o) { return o; }]));
      },
      setDdBusinessSegmentOptions () {
        var self = this;
        var filtered = _.filter(self.store.dd.DDSource, (v) => {
          if(this.isGlobalFilterEmpty) return true;

          return (self.store.dd.ddVal.ddCountrySelected.length > 0 ? (self.store.dd.ddVal.ddCountrySelected.includes(v.COUNTRY)) : false) 
            || (self.store.dd.ddVal.ddBusinessSegmentSelected.length > 0 ? (self.store.dd.ddVal.ddBusinessSegmentSelected.includes(v.BUSINESS_SEGMENT)) : false) 
            || (self.store.dd.ddVal.ddSourceSystemSelected.length > 0 ? (self.store.dd.ddVal.ddSourceSystemSelected.includes(v.EDM_SOURCE_SYSTEM_NAME)) : false)
            || (self.store.dd.ddVal.ddClusterSelected.length > 0 ? (self.store.dd.ddVal.ddClusterSelected.includes(v.CLUSTER_NAME)) : false)
            || (self.store.dd.ddVal.ddTierSelected.length > 0 ? (self.store.dd.ddVal.ddTierSelected.includes(v.TIER)) : false)
            || (self.store.dd.ddVal.ddItamSelected.length > 0 ? (self.store.dd.ddVal.ddItamSelected.includes(v.ITAM)) : false);
        });
        
        this.ddBusinessSegmentOptions = _.sortedUniq(_.sortBy(_.map(filtered, (v) => v.BUSINESS_SEGMENT.toString()), [function(o) { return o; }]));
      },
      setDdSourceSystemOptions () {
        var self = this;
        var filtered = _.filter(self.store.dd.DDSource, (v) => {
          if(this.isGlobalFilterEmpty) return true;
          
          return (self.store.dd.ddVal.ddCountrySelected.length > 0 ? (self.store.dd.ddVal.ddCountrySelected.includes(v.COUNTRY)) : false) 
            || (self.store.dd.ddVal.ddBusinessSegmentSelected.length > 0 ? (self.store.dd.ddVal.ddBusinessSegmentSelected.includes(v.BUSINESS_SEGMENT)) : false)
            || (self.store.dd.ddVal.ddSourceSystemSelected.length > 0 ? (self.store.dd.ddVal.ddSourceSystemSelected.includes(v.EDM_SOURCE_SYSTEM_NAME)) : false)
            || (self.store.dd.ddVal.ddClusterSelected.length > 0 ? (self.store.dd.ddVal.ddClusterSelected.includes(v.CLUSTER_NAME)) : false)
            || (self.store.dd.ddVal.ddTierSelected.length > 0 ? (self.store.dd.ddVal.ddTierSelected.includes(v.TIER)) : false)
            || (self.store.dd.ddVal.ddItamSelected.length > 0 ? (self.store.dd.ddVal.ddItamSelected.includes(v.ITAM)) : false);
        });
        
        this.ddSourceSystemOptions = _.sortedUniq(_.sortBy(_.map(filtered, (v) => v.EDM_SOURCE_SYSTEM_NAME.toString()), [function(o) { return o; }]));
      },
      setDdClusterOptions () {
        var self = this;
        var filtered = _.filter(self.store.dd.DDSource, (v) => {
          if(this.isGlobalFilterEmpty) return true;
          
          return (self.store.dd.ddVal.ddCountrySelected.length > 0 ? (self.store.dd.ddVal.ddCountrySelected.includes(v.COUNTRY)) : false) 
            || (self.store.dd.ddVal.ddBusinessSegmentSelected.length > 0 ? (self.store.dd.ddVal.ddBusinessSegmentSelected.includes(v.BUSINESS_SEGMENT)) : false)
            || (self.store.dd.ddVal.ddSourceSystemSelected.length > 0 ? (self.store.dd.ddVal.ddSourceSystemSelected.includes(v.EDM_SOURCE_SYSTEM_NAME)) : false)
            || (self.store.dd.ddVal.ddClusterSelected.length > 0 ? (self.store.dd.ddVal.ddClusterSelected.includes(v.CLUSTER_NAME)) : false)
            || (self.store.dd.ddVal.ddTierSelected.length > 0 ? (self.store.dd.ddVal.ddTierSelected.includes(v.TIER)) : false)
            || (self.store.dd.ddVal.ddItamSelected.length > 0 ? (self.store.dd.ddVal.ddItamSelected.includes(v.ITAM)) : false);
        });
        
        this.ddClusterOptions = _.sortedUniq(_.sortBy(_.map(filtered, (v) => v.CLUSTER_NAME.toString()), [function(o) { return o; }]));
      },
      setDdTierOptions () {
        var self = this;
        var filtered = _.filter(self.store.dd.DDSource, (v) => {
          if(this.isGlobalFilterEmpty) return true;
          
          return (self.store.dd.ddVal.ddCountrySelected.length > 0 ? (self.store.dd.ddVal.ddCountrySelected.includes(v.COUNTRY)) : false) 
            || (self.store.dd.ddVal.ddBusinessSegmentSelected.length > 0 ? (self.store.dd.ddVal.ddBusinessSegmentSelected.includes(v.BUSINESS_SEGMENT)) : false)
            || (self.store.dd.ddVal.ddSourceSystemSelected.length > 0 ? (self.store.dd.ddVal.ddSourceSystemSelected.includes(v.EDM_SOURCE_SYSTEM_NAME)) : false)
            || (self.store.dd.ddVal.ddClusterSelected.length > 0 ? (self.store.dd.ddVal.ddClusterSelected.includes(v.CLUSTER_NAME)) : false)
            || (self.store.dd.ddVal.ddTierSelected.length > 0 ? (self.store.dd.ddVal.ddTierSelected.includes(v.TIER)) : false)
            || (self.store.dd.ddVal.ddItamSelected.length > 0 ? (self.store.dd.ddVal.ddItamSelected.includes(v.ITAM)) : false);
        });
        
        this.ddTierOptions = _.sortedUniq(_.sortBy(_.map(filtered, (v) => v.TIER.toString()), [function(o) { return o; }]));
      },
      setDdItamOptions () {
        var self = this;
        var filtered = _.filter(self.store.dd.DDSource, (v) => {
          if(this.isGlobalFilterEmpty) return true;
          
          return (self.store.dd.ddVal.ddCountrySelected.length > 0 ? (self.store.dd.ddVal.ddCountrySelected.includes(v.COUNTRY)) : false) 
            || (self.store.dd.ddVal.ddBusinessSegmentSelected.length > 0 ? (self.store.dd.ddVal.ddBusinessSegmentSelected.includes(v.BUSINESS_SEGMENT)) : false) 
            || (self.store.dd.ddVal.ddSourceSystemSelected.length > 0 ? (self.store.dd.ddVal.ddSourceSystemSelected.includes(v.EDM_SOURCE_SYSTEM_NAME)) : false)
            || (self.store.dd.ddVal.ddClusterSelected.length > 0 ? (self.store.dd.ddVal.ddClusterSelected.includes(v.CLUSTER_NAME)) : false)
            || (self.store.dd.ddVal.ddTierSelected.length > 0 ? (self.store.dd.ddVal.ddTierSelected.includes(v.TIER)) : false)
            || (self.store.dd.ddVal.ddItamSelected.length > 0 ? (self.store.dd.ddVal.ddItamSelected.includes(v.ITAM)) : false);
        });
        
        this.ddItamOptions = _.sortedUniq(_.sortBy(_.map(filtered, (v) => v.ITAM.toString()), [function(o) { return o; }]));
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
        this.store.dd.ddVal.ddCountrySelected = [];
        this.store.dd.ddVal.ddBusinessSegmentSelected = [];
        this.store.dd.ddVal.ddSourceSystemSelected = [];
        this.store.dd.ddVal.ddClusterSelected = [];
        this.store.dd.ddVal.ddTierSelected = [];
        this.store.dd.ddVal.ddItamSelected = [];
      },
    },
}
</script>