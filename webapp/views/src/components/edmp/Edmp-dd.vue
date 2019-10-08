<style>
  @import '../../assets/styles/dashboard.css';
</style>

<template>
    <v-content>
        <b-container fluid>
            <PageHeader />
            
            <b-row class="my-4">
                <b-col class="ml-5 col-md-7">
                  <b-row class="ml-3">
                    <b-col>
                      <v-select
                        v-model="store.ddVal.ddCountrySelected"
                        :items="ddCountryOptions"
                        label="Country"
                        box
                      ></v-select>
                    </b-col>

                    <b-col>
                      <v-select
                        v-model="store.ddVal.ddBusinessSegmentSelected"
                        :items="ddBusinessSegmentOptions"
                        label="Business Segment"
                        box
                      ></v-select>
                    </b-col>

                    <b-col>
                      <v-select
                        v-model="store.ddVal.ddSourceSystemSelected"
                        :items="ddSourceSystemOptions"
                        label="Source System"
                        box
                      ></v-select>
                    </b-col>

                    <b-col>
                      <v-select
                        v-model="store.ddVal.ddClusterSelected"
                        :items="ddClusterOptions"
                        label="Cluster"
                        box
                      ></v-select>
                    </b-col>

                    <b-col>
                      <v-select
                        v-model="store.ddVal.ddTierSelected"
                        :items="ddTierOptions"
                        label="Tier"
                        box
                      ></v-select>
                    </b-col>

                    <b-col>
                      <v-select
                        v-model="store.ddVal.ddItamSelected"
                        :items="ddItamOptions"
                        label="ITAM"
                        box
                      ></v-select>
                    </b-col>
                  </b-row>
                  
                </b-col>
                
                <b-col>
                  <b-button class="float-right red-neon icon-only ml-3 shadow-sm" @click="resetFilter">
                    <i class="fa fa-filter"></i>
                  </b-button>

                  <page-export class="float-right shadow-sm" :storeName="storeName" :leftTableCols="store.leftHeaders" :rightTableCols="[]"/>
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
            { id: 'consumption', key: 'consumption-apps', name: 'Consumption Apps', route: this.addressPath + '/consumption-apps' },
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
      addressPath() {
        var tmp = this.$route.path.split("/");
        return tmp.slice(0, 4).join("/");
      },
      urlParam1() {
        return this.$route.params.system;
      },
      ddCountryOptions () {
        return _.uniq(_.map(this.store.DDSource, (v) => v.COUNTRY.toString())).filter(Boolean);
      },
      ddBusinessSegmentOptions () {
        var self = this;
        var filtered = _.filter(self.store.DDSource, (v) => {
          return self.store.ddVal.ddCountrySelected ? (v.COUNTRY == self.store.ddVal.ddCountrySelected) : true;
        });

        return _.uniq(_.map(filtered, (v) => v.BUSINESS_SEGMENT.toString())).filter(Boolean);
      },
      ddSourceSystemOptions () {
        var self = this;
        var filtered = _.filter(self.store.DDSource, (v) => {
          return self.store.ddVal.ddCountrySelected ? (v.COUNTRY == self.store.ddVal.ddCountrySelected) : true 
            && self.store.ddVal.ddBusinessSegmentSelected ? (v.BUSINESS_SEGMENT == self.store.ddVal.ddBusinessSegmentSelected) : true;
        });
        
        return _.uniq(_.map(filtered, (v) => v.SOURCE_SYSTEM.toString())).filter(Boolean);
      },
      ddClusterOptions () {
        var self = this;
        var filtered = _.filter(self.store.DDSource, (v) => {
          return self.store.ddVal.ddCountrySelected ? (v.COUNTRY == self.store.ddVal.ddCountrySelected) : true 
            && self.store.ddVal.ddBusinessSegmentSelected ? (v.BUSINESS_SEGMENT == self.store.ddVal.ddBusinessSegmentSelected) : true 
            && self.store.ddVal.ddSourceSystemSelected ? (v.SOURCE_SYSTEM == self.store.ddVal.ddSourceSystemSelected) : true;
        });
        
        return _.uniq(_.map(filtered, (v) => v.CLUSTER_NAME.toString())).filter(Boolean);
      },
      ddTierOptions () {
        var self = this;
        var filtered = _.filter(self.store.DDSource, (v) => {
          return self.store.ddVal.ddCountrySelected ? (v.COUNTRY == self.store.ddVal.ddCountrySelected) : true 
            && self.store.ddVal.ddBusinessSegmentSelected ? (v.BUSINESS_SEGMENT == self.store.ddVal.ddBusinessSegmentSelected) : true 
            && self.store.ddVal.ddSourceSystemSelected ? (v.SOURCE_SYSTEM == self.store.ddVal.ddSourceSystemSelected) : true
            && self.store.ddVal.ddClusterSelected ? (v.CLUSTER_NAME == self.store.ddVal.ddClusterSelected) : true;
        });
        
        return _.uniq(_.map(filtered, (v) => v.TIER.toString())).filter(Boolean);
      },
      ddItamOptions () {
        var self = this;
        var filtered = _.filter(self.store.DDSource, (v) => {
          return self.store.ddVal.ddCountrySelected ? (v.COUNTRY == self.store.ddVal.ddCountrySelected) : true 
            && self.store.ddVal.ddBusinessSegmentSelected ? (v.BUSINESS_SEGMENT == self.store.ddVal.ddBusinessSegmentSelected) : true 
            && self.store.ddVal.ddSourceSystemSelected ? (v.SOURCE_SYSTEM == self.store.ddVal.ddSourceSystemSelected) : true
            && self.store.ddVal.ddClusterSelected ? (v.CLUSTER_NAME == self.store.ddVal.ddClusterSelected) : true
            && self.store.ddVal.ddTierSelected ? (v.TIER == self.store.ddVal.ddTierSelected) : true;
        });
        
        return _.uniq(_.map(filtered, (v) => v.ITAM.toString())).filter(Boolean);
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
      },
      'store.ddVal.ddCountrySelected' () {
        if(this.store.firstload) return;

        this.store.ddVal.ddBusinessSegmentSelected = "";

        this.store.firstload = true;

        this.refreshActiveTabTable();
      },
      'store.ddVal.ddBusinessSegmentSelected' () {
        if(this.store.firstload) return;

        this.store.ddVal.ddSourceSystemSelected = "";

        this.store.firstload = true;

        this.refreshActiveTabTable();
      },
      'store.ddVal.ddSourceSystemSelected' () {
        if(this.store.firstload) return;

        this.store.ddVal.ddClusterSelected = "";
        
        this.store.firstload = true;

        this.refreshActiveTabTable();
      },
      'store.ddVal.ddClusterSelected' () {
        if(this.store.firstload) return;

        this.store.ddVal.ddTierSelected = "";

        this.store.firstload = true;

        this.refreshActiveTabTable();
      },
      'store.ddVal.ddTierSelected' () {
        if(this.store.firstload) return;

        this.store.ddVal.ddItamSelected = "";

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
        if(Object.keys(this.store.filters.left).length > 0){
          this.store.filters.left = {};
          this.getLeftTable();
        }
      },
    },
}
</script>