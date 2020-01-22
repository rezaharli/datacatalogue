<style>
  @import '../../../assets/styles/dashboard.css';
</style>

<style>
.dd-filter .v-text-field--box .v-label, .v-text-field--box .v-label {
  font-size: 14px;
  color: black;
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

.xs1.koma.lima{
  flex-basis: 12.5%;
  max-width: 12.5%;
}

.v-alert{
  width: 100%;
}

.action-button-wrapper .btn{
  max-width: 35px;
  max-height: 35px;
  width: 35px;
  height: 35px;
  margin-left: 10px;
}
</style>

<style>
.fixed-header{ display: flex; flex-direction: column; height: 100%; }

.fixed-header table{ table-layout: fixed; width: unset; }

.fixed-header th{ position: sticky; top: 0; z-index: 5; }
.fixed-header th::after{ content: ''; position: absolute; left: 0; bottom: 0; width: 100%; }

.fixed-header tr.v-datatable__progress th{ height: 1px; }

.fixed-header .v-table__overflow{ flex-grow: 1; flex-shrink: 1; overflow-x: auto; overflow-y: auto; }

.fixed-header > .v-datatable.v-table{ flex-grow: 0; flex-shrink: 1; }
.fixed-header > .v-datatable.v-table .v-datatable__actions{ flex-wrap: nowrap }
.fixed-header > .v-datatable.v-table .v-datatable__actions .v-datatable__actions__pagination{ white-space: nowrap; }
</style>

<template>
  <v-content>
    <v-container fluid grid-list-md box>
      <v-layout row>
        <v-flex d-flex xs12>
          <v-layout row wrap>
            <PageHeader />

            <v-flex d-flex xs2 koma lima>
              <global-filter-dropdown label="Country" 
                  v-model="store.dd.ddVal.ddCountrySelected"
                  :items="ddCountryOptions"
                  :disabled="activeTabStore.left.isLoading"
                />
            </v-flex>

            <v-flex d-flex xs2 koma lima>
              <global-filter-dropdown label="Business Segment" 
                v-model="store.dd.ddVal.ddBusinessSegmentSelected"
                :items="ddBusinessSegmentOptions"
                :disabled="activeTabStore.left.isLoading"
                />
            </v-flex>

            <v-flex d-flex xs2 koma lima>
              <global-filter-dropdown label="Source System" 
                  v-model="store.dd.ddVal.ddSourceSystemSelected"
                  :items="ddSourceSystemOptions"
                  :disabled="activeTabStore.left.isLoading"
                />
            </v-flex>

            <v-flex d-flex xs2 koma lima>
              <global-filter-dropdown label="Tier" 
                  v-model="store.dd.ddVal.ddTierSelected"
                  :items="ddTierOptions"
                  :disabled="activeTabStore.left.isLoading"
                />
            </v-flex>

            <v-flex d-flex xs2 koma lima>
              <v-layout row wrap justify-end>
                <v-flex d-flex xs12>
                  &nbsp;
                </v-flex>
              </v-layout>
            </v-flex>

            <v-flex d-flex xs2 koma lima>
              <v-layout row wrap align-center justify-end class="action-button-wrapper">
                <!-- <v-flex d-flex xs12> -->
                  <b-button class="float-right red-neon icon-only shadow-sm" @click="resetFilter"><i class="fa fa-filter"></i></b-button>

                  <page-export class="float-right icon-only shadow-sm" :storeName="activeTabStoreName" :leftTableCols="activeTabStore.leftHeaders" :rightTableCols="[]" :rowSelectInvolved="true" />
                <!-- </v-flex> -->
              </v-layout>
            </v-flex>

            <v-flex d-flex xs12>
              <v-layout row wrap>
                <v-tabs id="page-tab" class="page-tab" v-model="activeTab">
                    <v-tab v-for="tab in tabs" class="px-2 mx-5" :id="'tab-' + tab.id" :key="tab.key" :to="addressPath + '/' + tab.key" :ref="tab.id">{{ tab.name }}</v-tab>
                </v-tabs>
              </v-layout>
            </v-flex>
              
            <v-flex d-flex xs12>
              <v-layout row wrap>
                <transition name="fade" mode="out-in">
                  <router-view />
                </transition>
              </v-layout>
            </v-flex>
              
            <v-flex d-flex xs12>
              <v-layout row wrap>
                <p v-if="!showNote">&nbsp;</p>
                <p v-if="showNote">Note: Tables with column count exceeded with more than 100 will be displayed on the next page.</p>
              </v-layout>
            </v-flex>
          </v-layout>
        </v-flex>
      </v-layout>
    </v-container>
  </v-content>
</template>

<script>
import PageHeader from '@/components/PageHeader';
import pageExport from "@/components/PageExport.vue";
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
        ddTierOptions: [],
        activeTab: '',
        tabs: [
            { id: 'technical', key: 'technical-metadata', name: 'Technical Metadata', route: this.addressPath + '/technical-metadata' },
            { id: 'business', key: 'business-metadata', name: 'Business Metadata', route: this.addressPath + '/business-metadata' },
            // { id: 'consumption', key: 'consumption-apps', name: 'Consumption Applications', route: this.addressPath + '/consumption-apps' },
        ],
        showNote: true,
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
        return this.$store.getters[this.storeName + "/isDdGlobalFilterEmpty"];
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
      $route() {},
      activeTab(val, oldVal){
        if (oldVal) {
          this.store.dd.isNewPage = false;
        } else {
          this.store.dd.isNewPage = true;
        }
      },
      'store.dd.ddVal.ddCountrySelected'(val) {
        this.activeTabStore.left.display = [];
        this.activeTabStore.left.source = [];
        this.activeTabStore.left.totalItems = 0;

        if(this.store.dd.firstload) return;
        this.activeTabStore.left.isLoading = true;

        this.setDdBusinessSegmentOptions();
        this.setDdSourceSystemOptions();
        this.setDdTierOptions();

        setTimeout(() => {
          this.refreshActiveTabTable("COUNTRY", val);
        }, 0);
      },
      'store.dd.ddVal.ddBusinessSegmentSelected'(val) {
        this.activeTabStore.left.display = [];
        this.activeTabStore.left.source = [];
        this.activeTabStore.left.totalItems = 0;

        if(this.store.dd.firstload) return;
        this.activeTabStore.left.isLoading = true;
        
        this.setDdCountryOptions();
        this.setDdSourceSystemOptions();
        this.setDdTierOptions();

        setTimeout(() => {
          this.refreshActiveTabTable("BUSINESS_SEGMENT", val);
        }, 0);
      },
      'store.dd.ddVal.ddSourceSystemSelected'(val) {
        this.activeTabStore.left.display = [];
        this.activeTabStore.left.source = [];
        this.activeTabStore.left.totalItems = 0;

        if(this.store.dd.firstload) return;
        this.activeTabStore.left.isLoading = true;

        this.setDdCountryOptions();
        this.setDdBusinessSegmentOptions();
        this.setDdTierOptions();

        setTimeout(() => {
          this.refreshActiveTabTable("EDM_SOURCE_SYSTEM_NAME", val);
        }, 0);
      },
      'store.dd.ddVal.ddTierSelected'(val) {
        this.activeTabStore.left.display = [];
        this.activeTabStore.left.source = [];
        this.activeTabStore.left.totalItems = 0;

        if(this.store.dd.firstload) return;
        this.activeTabStore.left.isLoading = true;

        this.setDdCountryOptions();
        this.setDdBusinessSegmentOptions();
        this.setDdSourceSystemOptions();

        setTimeout(() => {
          this.refreshActiveTabTable("TIER", val);
        }, 0);
      },
    },
    mounted() {
      this.store.tabName = this.storeName;
      this.store.system = this.$route.params.system;

      this.getDropdownOpts();

      setInterval(() => {
        this.showNote = !this.showNote;
      }, 500);
    },
    methods: {
      setNewDropdownOpts() {
        this.ddCountryOptions = this._.sortedUniq(this._.sortBy(this._.map(this.store.dd.DDSource, (v) => v.COUNTRY.toString()), [function(o) { return o; }]));
        this.ddBusinessSegmentOptions = this._.sortedUniq(this._.sortBy(this._.map(this.store.dd.DDSource, (v) => v.BUSINESS_SEGMENT.toString()), [function(o) { return o; }]));
        this.ddSourceSystemOptions = this._.sortedUniq(this._.sortBy(this._.map(this.store.dd.DDSource, (v) => v.EDM_SOURCE_SYSTEM_NAME.toString()), [function(o) { return o; }]));
        this.ddTierOptions = this._.sortedUniq(this._.sortBy(this._.map(this.store.dd.DDSource, (v) => v.TIER.toString()), [function(o) { return o; }]));
      },
      getDropdownOpts() {
        this.$store.dispatch(`${this.storeName}/getDdDropdownOpts`).then(() => {
          this.setNewDropdownOpts();
        });
      },
      refreshActiveTabTable(updatedAttr, val) {
        this.store.dd.firstload = true;
        if( ! this.store.dd.globalFilters.filterTypes) this.store.dd.globalFilters.filterTypes = {};

        this.store.dd.globalFilters[updatedAttr] = val;
        this.store.dd.globalFilters.filterTypes[updatedAttr] = "eq";

        this.activeTabStore.filters.left = {}

        if(this.isGlobalFilterEmpty) {
          this.activeTabStore.left.isLoading = false;
          this.store.dd.firstload = false;
          this.store.dd.displayTable = false;
        } else {
          // Remove table component from the DOM
          this.store.dd.displayTable = false;
          this.$nextTick().then(() => {
            // Add the component back in
            this.store.dd.displayTable = true;
          });
        }
      },
      setDdCountryOptions () {
        var self = this;
        var filtered = this._.filter(self.store.dd.DDSource, (v) => {
          if(this.isGlobalFilterEmpty) return true;

          return (self.store.dd.ddVal.ddCountrySelected.length > 0 ? (self.store.dd.ddVal.ddCountrySelected.includes(v.COUNTRY)) : false)
          || (
            (self.store.dd.ddVal.ddBusinessSegmentSelected.length > 0 ? (self.store.dd.ddVal.ddBusinessSegmentSelected.includes(v.BUSINESS_SEGMENT)) : true) 
            && (self.store.dd.ddVal.ddSourceSystemSelected.length > 0 ? (self.store.dd.ddVal.ddSourceSystemSelected.includes(v.EDM_SOURCE_SYSTEM_NAME)) : true)
            && (self.store.dd.ddVal.ddTierSelected.length > 0 ? (self.store.dd.ddVal.ddTierSelected.includes(v.TIER)) : true)
          );
        });
        
        this.ddCountryOptions = this._.sortedUniq(this._.sortBy(this._.map(filtered, (v) => v.COUNTRY.toString()), [function(o) { return o; }]));
      },
      setDdBusinessSegmentOptions () {
        var self = this;
        var filtered = this._.filter(self.store.dd.DDSource, (v) => {
          if(this.isGlobalFilterEmpty) return true;

          return (self.store.dd.ddVal.ddBusinessSegmentSelected.length > 0 ? (self.store.dd.ddVal.ddBusinessSegmentSelected.includes(v.BUSINESS_SEGMENT)) : false)
          || (
            (self.store.dd.ddVal.ddCountrySelected.length > 0 ? (self.store.dd.ddVal.ddCountrySelected.includes(v.COUNTRY)) : true) 
            && (self.store.dd.ddVal.ddSourceSystemSelected.length > 0 ? (self.store.dd.ddVal.ddSourceSystemSelected.includes(v.EDM_SOURCE_SYSTEM_NAME)) : true)
            && (self.store.dd.ddVal.ddTierSelected.length > 0 ? (self.store.dd.ddVal.ddTierSelected.includes(v.TIER)) : true)
          );
        });
        
        this.ddBusinessSegmentOptions = this._.sortedUniq(this._.sortBy(this._.map(filtered, (v) => v.BUSINESS_SEGMENT.toString()), [function(o) { return o; }]));
      },
      setDdSourceSystemOptions () {
        var self = this;
        var filtered = this._.filter(self.store.dd.DDSource, (v) => {
          if(this.isGlobalFilterEmpty) return true;
          
          return (self.store.dd.ddVal.ddSourceSystemSelected.length > 0 ? (self.store.dd.ddVal.ddSourceSystemSelected.includes(v.EDM_SOURCE_SYSTEM_NAME)) : false)
          || (
            (self.store.dd.ddVal.ddCountrySelected.length > 0 ? (self.store.dd.ddVal.ddCountrySelected.includes(v.COUNTRY)) : true) 
            && (self.store.dd.ddVal.ddBusinessSegmentSelected.length > 0 ? (self.store.dd.ddVal.ddBusinessSegmentSelected.includes(v.BUSINESS_SEGMENT)) : true)
            && (self.store.dd.ddVal.ddTierSelected.length > 0 ? (self.store.dd.ddVal.ddTierSelected.includes(v.TIER)) : true)
          );
        });
        
        this.ddSourceSystemOptions = this._.sortedUniq(this._.sortBy(this._.map(filtered, (v) => v.EDM_SOURCE_SYSTEM_NAME.toString()), [function(o) { return o; }]));
      },
      setDdTierOptions () {
        var self = this;
        var filtered = this._.filter(self.store.dd.DDSource, (v) => {
          if(this.isGlobalFilterEmpty) return true;
          
          return (self.store.dd.ddVal.ddTierSelected.length > 0 ? (self.store.dd.ddVal.ddTierSelected.includes(v.TIER)) : false)
            || (
              (self.store.dd.ddVal.ddCountrySelected.length > 0 ? (self.store.dd.ddVal.ddCountrySelected.includes(v.COUNTRY)) : true) 
              && (self.store.dd.ddVal.ddBusinessSegmentSelected.length > 0 ? (self.store.dd.ddVal.ddBusinessSegmentSelected.includes(v.BUSINESS_SEGMENT)) : true)
              && (self.store.dd.ddVal.ddSourceSystemSelected.length > 0 ? (self.store.dd.ddVal.ddSourceSystemSelected.includes(v.EDM_SOURCE_SYSTEM_NAME)) : true)
            );
        });
        
        this.ddTierOptions = this._.sortedUniq(this._.sortBy(this._.map(filtered, (v) => v.TIER.toString()), [function(o) { return o; }]));
      },
      updateRouter(val){
        this.$router.push(val);
      },
      resetFilter() {
        this.store.dd.globalFilters = {}

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
        this.store.dd.ddVal.ddTierSelected = [];
      },
    },
}
</script>