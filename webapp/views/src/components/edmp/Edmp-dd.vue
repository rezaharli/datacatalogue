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
                    <v-tabs left id="page-tab" class="page-tab">
                      <v-tab class="px-2 mx-5" key="technical-metadata" id="tab-technical" ref="technical">TECHNICAL METADATA</v-tab>
                      <v-tab class="px-2 mx-5" key="business-mapping" id="tab-business" ref="business">BUSINESS MAPPING</v-tab>
                      <v-tab class="px-2 mx-5" key="original-to" id="tab-original" ref="policy">ORIGINAL TO EDM SOURCE SYSTEM MAPPING</v-tab>
                      <v-tab class="px-2 mx-5" key="data-presence" id="tab-presence" ref="policy">DATA PRESENCE</v-tab>
                      <v-tab class="px-2 mx-5" key="data-protection" id="tab-protection" ref="policy">DATA PROTECTION</v-tab>
                      <v-tab class="px-2 mx-5" key="attribute-business" id="tab-attribute" ref="policy">ATTRIBUTE BUSINESS PROPERTY</v-tab>
                      <!-- <v-tab class="px-2 mx-5" key="interfaces">Interfaces</v-tab> -->

                      <v-tab-item key="technical-metadata">
                        <edmp-dd-technical />
                      </v-tab-item>

                      <v-tab-item key="business-mapping">
                        <edmp-dd-business />
                      </v-tab-item>

                      <v-tab-item key="original-to">
                        <edmp-dd-original />
                      </v-tab-item>

                      <v-tab-item key="data-presence">
                        <edmp-dd-presence />
                      </v-tab-item>

                      <v-tab-item key="data-protection">
                        <edmp-dd-protection />
                      </v-tab-item>

                      <v-tab-item key="attribute-business">
                        <edmp-dd-attribute />
                      </v-tab-item>
                    </v-tabs>
                </b-col>
            </b-row>
        </b-container>
    </v-content>
</template>

<script>
import PageHeader from '../PageHeader';
import pageExport from "../PageExport.vue";
import EdmpDdTechnical from './Edmp-dd-technical';
import EdmpDdBusiness from './Edmp-dd-business';
import EdmpDdOriginal from './Edmp-dd-original';
import EdmpDdPresence from './Edmp-dd-presence';
import EdmpDdProtection from './Edmp-dd-protection';
import EdmpDdAttribute from './Edmp-dd-attribute';

export default {
    components: { PageHeader, pageExport, EdmpDdTechnical, EdmpDdBusiness, EdmpDdOriginal, EdmpDdPresence, EdmpDdProtection, EdmpDdAttribute, 
    // DscDdBusiness, DscDdPolicy, DscDdInterfaces 
    },
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