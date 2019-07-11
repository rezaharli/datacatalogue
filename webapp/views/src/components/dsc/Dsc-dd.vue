<style>
  @import '../../assets/styles/dashboard.css';
</style>

<template>
    <v-content>
        <b-container fluid>
            <PageHeader />
            
            <b-row class="my-4">
                <b-col>
                    <b-button class="float-right red-neon icon-only ml-3" @click="resetFilter">
                        <i class="fa fa-filter"></i>
                    </b-button>

                    <page-export class="float-right" :storeName="storeName" :leftTableCols="store.leftHeaders" :rightTableCols="[]"/>
                </b-col>
            </b-row>

            <b-row>
                <b-col>
                    <v-tabs id="page-tab" class="page-tab" v-model="activeTab">
                        <v-tab v-for="tab in tabs" class="px-2 mx-5" :id="'tab-' + tab.id" :key="tab.key" :to="addressPath + '/' + tab.key + '/' + urlParam1" :ref="tab.id">{{ tab.name }}</v-tab>
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
import DscDdTechnical from './Dsc-dd-technical';
import DscDdBusiness from './Dsc-dd-business';
import DscDdPolicy from './Dsc-dd-policy';
import DscDdInterfaces from './Dsc-dd-interfaces';

export default {
    components: { PageHeader, pageExport, DscDdTechnical, DscDdBusiness, DscDdPolicy, DscDdInterfaces },
    data() {
      return {
        storeName: "dscdd",
        activeTab: '',
        tabs: [
            { id: 'technical', key: 'technical-metadata', name: 'Technical Metadata', route: this.addressPath + '/technical-metadata/' + this.urlParam1 },
            { id: 'business', key: 'business-metadata', name: 'Business Metadata', route: this.addressPath + '/business-metadata/' + this.urlParam1 },
            { id: 'policy', key: 'policy-related', name: 'Policy Related Information', route: this.addressPath + '/policy-related/' + this.urlParam1 },
        ]
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
      urlParam1() {
        return this.$route.params.system;
      }
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
      this.store.system = this.$route.params.system;
    },
    methods: {
      getLeftTable() {
        this.$store.dispatch(`${this.storeName}/getLeftTable`);
      },
      updateRouter(val){
        this.$router.push(val);
      },
      resetFilter (e) {
        if(Object.keys(this.store.filters.left).length > 0){
          this.store.filters.left = {};
          this.getLeftTable();
        }

        // if(Object.keys(this.store.filters.right).length > 0){
        //     this.store.filters.right = {}
        //     this.getMyRightTable(this.$route.params.system);
        // }
      },
    },
}
</script>