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
                    <v-tabs left class="page-tab">
                      <v-tab class="px-2 mx-5" key="technical-metadata">Technical Metadata</v-tab>
                      <v-tab class="px-2 mx-5" key="business-metadata">Business Metadata</v-tab>
                      <v-tab class="px-2 mx-5" key="policy-related">Policy Related Information</v-tab>
                      <!-- <v-tab class="px-2 mx-5" key="interfaces">Interfaces</v-tab> -->

                      <v-tab-item key="technical-metadata">
                        <dsc-dd-technical />
                      </v-tab-item>

                      <v-tab-item key="business-metadata">
                        <dsc-dd-business />
                      </v-tab-item>

                      <v-tab-item key="policy-related">
                        <dsc-dd-policy />
                      </v-tab-item>

                      <!-- <v-tab-item key="interfaces">
                        <dsc-dd-interfaces />
                      </v-tab-item> -->
                    </v-tabs>
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
      };
    },
    computed: {
      store() {
        return this.$store.state[this.storeName].all;
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
      this.store.system = this.$route.params.system;
    },
    methods: {
      getLeftTable() {
        this.$store.dispatch(`${this.storeName}/getLeftTable`);
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