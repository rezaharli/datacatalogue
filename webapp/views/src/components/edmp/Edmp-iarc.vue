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
                <b-col class="ml-5 col-md-10" v-if="activeTabStoreName == personalStoreName">
                  <b-row class="ml-3 dd-filter">
                    <b-col cols="2">
                      <v-tooltip top>
                        <template slot="activator" slot-scope="{ on }">
                          <div v-on="on">
                            <v-select
                              v-model="store.iarc.ddVal.ddCountrySelected"
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
                              v-model="store.iarc.ddVal.ddSourceSystemSelected"
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
                              v-model="store.iarc.ddVal.ddItamSelected"
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
        personalStoreName: "edmpIarcPersonal",
        informationStoreName: "dsciarc",
        activeTab: '',
        tabs: [
            { id: 'information', key: 'information', name: 'Information Asset Category', route: this.addressPath + '/information' },
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
      addressPath() {
        var tmp = this.$route.path.split("/");
        return tmp.slice(0, 4).join("/");
      },
      urlParam1() {
        return this.$route.params.system;
      },
      ddCountryOptions () {
        return _.sortedUniq(_.sortBy(_.map(this.store.iarc.DDSource, (v) => v.COUNTRY.toString()), [function(o) { return o; }]));
      },
      ddSourceSystemOptions () {
        var self = this;
        var filtered = _.filter(self.store.iarc.DDSource, (v) => {
          return self.store.iarc.ddVal.ddCountrySelected.length > 0 ? (self.store.iarc.ddVal.ddCountrySelected.includes(v.COUNTRY)) : true;
        });
        
        return _.sortedUniq(_.sortBy(_.map(filtered, (v) => v.EDM_SOURCE_SYSTEM_NAME.toString()), [function(o) { return o; }]));
      },
      ddItamOptions () {
        var self = this;
        var filtered = _.filter(self.store.iarc.DDSource, (v) => {
          return self.store.iarc.ddVal.ddCountrySelected.length > 0 ? (self.store.iarc.ddVal.ddCountrySelected.includes(v.COUNTRY)) : true  
            && self.store.iarc.ddVal.ddSourceSystemSelected.length > 0 ? (self.store.iarc.ddVal.ddSourceSystemSelected.includes(v.EDM_SOURCE_SYSTEM_NAME)) : true;
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
      'store.iarc.ddVal.ddCountrySelected' () {
        if(this.store.iarc.firstload) return;

        this.store.iarc.ddVal.ddSourceSystemSelected = [];
        this.store.iarc.ddVal.ddItamSelected = [];

        this.store.iarc.firstload = true;

        this.refreshActiveTabTable();
      },
      'store.iarc.ddVal.ddSourceSystemSelected' () {
        if(this.store.iarc.firstload) return;

        this.store.iarc.ddVal.ddItamSelected = [];
        
        this.store.iarc.firstload = true;

        this.refreshActiveTabTable();
      },
      'store.iarc.ddVal.ddItamSelected' () {
        if(this.store.iarc.firstload) return;

        this.store.iarc.firstload = true;

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
        this.$store.dispatch(`${this.storeName}/getIarcDropdownOpts`);
      },
      getLeftTable() {
        this.$store.dispatch(`${this.storeName}/getLeftTable`);
      },
      refreshActiveTabTable() {
        if(this.activeTab.indexOf("personal") != -1){
          this.refreshPersonalTable();
        }
        if(this.activeTab.indexOf("information") != -1){
          this.refreshInformationTable();
        }
      },
      refreshPersonalTable() {
        if( ! this.personalStore.filters.left.filterTypes) this.personalStore.filters.left.filterTypes = {};

        this.personalStore.filters.left["COUNTRY"] = this.store.iarc.ddVal.ddCountrySelected;
        this.personalStore.filters.left.filterTypes["COUNTRY"] = "eq";

        this.personalStore.filters.left["EDM_SOURCE_SYSTEM_NAME"] = this.store.iarc.ddVal.ddSourceSystemSelected;
        this.personalStore.filters.left.filterTypes["EDM_SOURCE_SYSTEM_NAME"] = "eq";

        this.personalStore.filters.left["ITAM"] = this.store.iarc.ddVal.ddItamSelected;
        this.personalStore.filters.left.filterTypes["ITAM"] = "eq";

        this.$store.dispatch(`${this.personalStoreName}/getLeftTable`).then(res => {
          this.store.iarc.firstload = false;
        });
      },
      refreshInformationTable() {
        this.$store.dispatch(`${this.informationStoreName}/getLeftTable`).then(res => {
          this.store.iarc.firstload = false;
        });
      },
      updateRouter(val){
        this.$router.push(val);
      },
      resetFilter (e) {
        if(this.activeTab.indexOf("personal-data") != -1){
          if(Object.keys(this.personalStore.filters.left).length > 0){
            this.personalStore.filters.left = {};
          }
        }

        this.store.iarc.ddVal.ddCountrySelected = "";
        this.store.iarc.ddVal.ddSourceSystemSelected = "";
        this.store.iarc.ddVal.ddItamSelected = "";

        this.refreshActiveTabTable();
      },
    },
}
</script>