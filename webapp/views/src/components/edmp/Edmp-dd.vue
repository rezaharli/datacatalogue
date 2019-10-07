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
                        :items="selectCountry"
                        label="Country"
                        box
                      ></v-select>
                    </b-col>

                    <b-col>
                      <v-select
                        :items="selectBusinessSegment"
                        label="System Name"
                        box
                      ></v-select>
                    </b-col>

                    <b-col>
                      <v-select
                        :items="selectSourceSystem"
                        label="Tier"
                        box
                      ></v-select>
                    </b-col>

                    <b-col>
                      <v-select
                        :items="selectCluster"
                        label="Country"
                        box
                      ></v-select>
                    </b-col>

                    <b-col>
                      <v-select
                        :items="selectTier"
                        label="Tier"
                        box
                      ></v-select>
                    </b-col>

                    <b-col>
                      <v-select
                        :items="selectITAM"
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
                    <v-tabs left id="page-tab" class="page-tab">
                      <v-tab class="px-2 mx-5" key="technical-metadata" id="tab-technical" ref="technical">TECHNICAL METADATA</v-tab>
                      <v-tab class="px-2 mx-5" key="business-metadata" id="tab-business" ref="business">BUSINESS METADATA</v-tab>
                      <v-tab class="px-2 mx-5" key="consumption-apps" id="tab-consumption" ref="policy">CONSUMPTION APPS</v-tab>

                      <v-tab-item key="technical-metadata">
                        <edmp-dd-technical />
                      </v-tab-item>

                      <v-tab-item key="business-metadata">
                        <edmp-dd-business />
                      </v-tab-item>

                      <v-tab-item key="consumption-apps">
                        <edmp-dd-consumption />
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
import EdmpDdConsumption from './Edmp-dd-consumption';

export default {
    components: { PageHeader, pageExport, EdmpDdTechnical, EdmpDdBusiness, EdmpDdConsumption, 
    },
    data() {
      return {
        storeName: "dscdd",
        selectCountry: ["3128","1531","2211","1456","1279","1290","1489"],
        selectBusinessSegment: ["a","adiuw","jikut","lorem","iosum","dolor","sitamet"],
        selectSourceSystem: ["Tier 1","Tier 2","Tier 3"],
        selectCluster: ["Singapore","Malaysia","Indonesia","India","USA"],
        selectTier: ["Product 1","Product 2","Product 3","Product 4","Product 5","Product 6"],
        selectITAM: ["Product 1","Product 2","Product 3","Product 4","Product 5","Product 6"]
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