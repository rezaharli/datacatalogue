<style>
  @import '../../../assets/styles/dashboard.css';
</style>

<style>
.dd-filter .v-text-field--box .v-label {
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
                        v-model="store.iarc.ddVal.ddCountrySelected"
                        :items="ddCountryOptions"
                        :disabled="activeTabStore.left.isLoading"
                      />
                    </b-col>

                    <b-col cols="2">
                      <global-filter-dropdown label="Source System" 
                        v-model="store.iarc.ddVal.ddSourceSystemSelected"
                        :items="ddSourceSystemOptions"
                        :disabled="activeTabStore.left.isLoading"
                      />
                    </b-col>

                    <b-col cols="2">
                      <global-filter-dropdown label="ITAM" 
                        v-model="store.iarc.ddVal.ddItamSelected"
                        :items="ddItamOptions"
                        :disabled="activeTabStore.left.isLoading"
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
import PageHeader from '@/components/PageHeader';
import pageExport from "@/components/PageExport.vue";
import globalFilterDropdown from "./GlobalFilterDropdown.vue";

export default {
    components: { PageHeader, pageExport, globalFilterDropdown },
    data() {
      return {
        storeName: "edmp",
        personalStoreName: "edmpIarcPersonal",
        informationStoreName: "dsciarc",
        ddCountryOptions: [],
        ddSourceSystemOptions: [],
        ddItamOptions: [],
        activeTab: '',
        tabs: [
            { id: 'personal', key: 'personal', name: 'Personal Data', route: this.addressPath + '/personal' },
        ]
      };
    },
    computed: {
      store() {
        return this.$store.state[this.storeName].all;
      },
      personalStore() {
        return this.$store.state[this.personalStoreName].all;
      },
      informationStore() {
        return this.$store.state[this.informationStoreName].all;
      },
      activeTabStoreName() {
        if(this.activeTab.indexOf("personal") != -1){
          return this.personalStoreName;
        } else {
          return this.informationStoreName;
        }
      },
      activeTabStore() {
        return this.$store.state[this.activeTabStoreName].all;
      },
      isGlobalFilterEmpty() {
        return this.$store.getters[this.storeName + "/isIarcGlobalFilterEmpty"];
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
      activeTab(val, oldVal){
        if (oldVal) {
          this.store.iarc.isNewPage = false;
        } else {
          this.store.iarc.isNewPage = true;
        }
      },
      'store.iarc.ddVal.ddCountrySelected'(val) {
        this.activeTabStore.left.display = [];
        this.activeTabStore.left.source = [];
        this.activeTabStore.left.totalItems = 0;

        if(this.store.iarc.firstload) return;
        this.activeTabStore.left.isLoading = true;

        this.setDdSourceSystemOptions();
        this.setDdItamOptions();

        setTimeout(() => {
          this.refreshActiveTabTable("COUNTRY", val);
        }, 0);
      },
      'store.iarc.ddVal.ddSourceSystemSelected'(val) {
        this.activeTabStore.left.display = [];
        this.activeTabStore.left.source = [];
        this.activeTabStore.left.totalItems = 0;

        if(this.store.iarc.firstload) return;
        this.activeTabStore.left.isLoading = true;

        this.setDdCountryOptions();
        this.setDdItamOptions();

        setTimeout(() => {
          this.refreshActiveTabTable("EDM_SOURCE_SYSTEM_NAME", val);
        }, 0);
      },
      'store.iarc.ddVal.ddItamSelected'(val) {
        this.activeTabStore.left.display = [];
        this.activeTabStore.left.source = [];
        this.activeTabStore.left.totalItems = 0;

        if(this.store.iarc.firstload) return;
        this.activeTabStore.left.isLoading = true;

        this.setDdCountryOptions();
        this.setDdSourceSystemOptions();

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
        this.ddCountryOptions = _.sortedUniq(_.sortBy(_.map(this.store.iarc.DDSource, (v) => v.COUNTRY.toString()), [function(o) { return o; }]));
        this.ddSourceSystemOptions = _.sortedUniq(_.sortBy(_.map(this.store.iarc.DDSource, (v) => v.EDM_SOURCE_SYSTEM_NAME.toString()), [function(o) { return o; }]));
        this.ddItamOptions = _.sortedUniq(_.sortBy(_.map(this.store.iarc.DDSource, (v) => v.ITAM.toString()), [function(o) { return o; }]));
      },
      getDropdownOpts() {
        this.$store.dispatch(`${this.storeName}/getIarcDropdownOpts`).then(() => {
          this.setNewDropdownOpts();
        });
      },
      refreshActiveTabTable(updatedAttr, val) {
        this.store.iarc.firstload = true;
        if( ! this.store.iarc.globalFilters.filterTypes) this.store.iarc.globalFilters.filterTypes = {};
        
        this.store.iarc.globalFilters[updatedAttr] = val;
        this.store.iarc.globalFilters.filterTypes[updatedAttr] = "eq";

        this.activeTabStore.filters.left = {}

        if(this.isGlobalFilterEmpty) {
          this.activeTabStore.left.isLoading = false;
        } else {
          // Remove table component from the DOM
          this.store.iarc.displayTable = false;
          this.$nextTick().then(() => {
            // Add the component back in
            this.store.iarc.displayTable = true;
          });
        }
      },
      setDdCountryOptions () {
        var self = this;
        var filtered = _.filter(self.store.iarc.DDSource, (v) => {
          if(this.isGlobalFilterEmpty) return true;

          return (self.store.iarc.ddVal.ddCountrySelected.length > 0 ? (self.store.iarc.ddVal.ddCountrySelected.includes(v.COUNTRY)) : false) 
            || (
              (self.store.iarc.ddVal.ddSourceSystemSelected.length > 0 ? (self.store.iarc.ddVal.ddSourceSystemSelected.includes(v.EDM_SOURCE_SYSTEM_NAME)) : true)
              && (self.store.iarc.ddVal.ddItamSelected.length > 0 ? (self.store.iarc.ddVal.ddItamSelected.includes(v.ITAM)) : true)
            );
        });
        
        this.ddCountryOptions = _.sortedUniq(_.sortBy(_.map(filtered, (v) => v.COUNTRY.toString()), [function(o) { return o; }]));
      },
      setDdSourceSystemOptions () {
        var self = this;
        var filtered = _.filter(self.store.iarc.DDSource, (v) => {
          if(this.isGlobalFilterEmpty) return true;
          
          return (self.store.iarc.ddVal.ddSourceSystemSelected.length > 0 ? (self.store.iarc.ddVal.ddSourceSystemSelected.includes(v.EDM_SOURCE_SYSTEM_NAME)) : false)
            || (
              (self.store.iarc.ddVal.ddCountrySelected.length > 0 ? (self.store.iarc.ddVal.ddCountrySelected.includes(v.COUNTRY)) : true) 
              && (self.store.iarc.ddVal.ddItamSelected.length > 0 ? (self.store.iarc.ddVal.ddItamSelected.includes(v.ITAM)) : true)
            );
        });
        
        this.ddSourceSystemOptions = _.sortedUniq(_.sortBy(_.map(filtered, (v) => v.EDM_SOURCE_SYSTEM_NAME.toString()), [function(o) { return o; }]));
      },
      setDdItamOptions () {
        var self = this;
        var filtered = _.filter(self.store.iarc.DDSource, (v) => {
          if(this.isGlobalFilterEmpty) return true;
          
          return (self.store.iarc.ddVal.ddItamSelected.length > 0 ? (self.store.iarc.ddVal.ddItamSelected.includes(v.ITAM)) : false)
            || (
              (self.store.iarc.ddVal.ddCountrySelected.length > 0 ? (self.store.iarc.ddVal.ddCountrySelected.includes(v.COUNTRY)) : true) 
              && (self.store.iarc.ddVal.ddSourceSystemSelected.length > 0 ? (self.store.iarc.ddVal.ddSourceSystemSelected.includes(v.EDM_SOURCE_SYSTEM_NAME)) : true)
            );
        });
        
        this.ddItamOptions = _.sortedUniq(_.sortBy(_.map(filtered, (v) => v.ITAM.toString()), [function(o) { return o; }]));
      },
      updateRouter(val){
        this.$router.push(val);
      },
      resetFilter (e) {
        this.store.iarc.globalFilters = {}

        if(this.activeTab.indexOf("personal") != -1){
          if(Object.keys(this.personalStore.filters.left).length > 0){
            this.personalStore.filters.left = {};
          }
        }

        this.store.iarc.ddVal.ddCountrySelected = [];
        this.store.iarc.ddVal.ddSourceSystemSelected = [];
        this.store.iarc.ddVal.ddItamSelected = [];
      },
    },
}
</script>