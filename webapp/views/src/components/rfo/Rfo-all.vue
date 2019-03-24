<style>
  @import '../../assets/styles/dashboard.css';
</style>

<style>
table.v-table thead th > div.btn-group {
  width: auto;
}

.header-filter-icon{

}
.header-filter-icon .dropdown-menu{
  overflow: scroll;
  height: 200px;
}
</style>


<template>
  <b-row style="margin-top: 10px;margin-bottom: 10px;">
    <b-col>
      <!-- Dsc details -->
      <router-view/>

      <!-- Main content -->
      <b-row>
        <b-col>
          <page-loader v-if="store.left.isLoading || (store.isRightTable && store.right.isLoading)" />
          
          <b-row>
            <b-col>
              <page-search :storeName="storeName" :searchDDInputs="searchDropdownInputs"/>
            </b-col>

            <b-col></b-col>
            <b-col></b-col>
            <b-col>
              <page-export :storeName="storeName" :leftTableCols="firstTableHeaders" :rightTableCols="secondTableHeaders"/>
            </b-col>
          </b-row>

          <b-row>
            <b-col cols="6">
              <v-data-table
                  :headers="firstTableHeaders"
                  :items="store.left.display"
                  :pagination.sync="store.left.pagination"
                  :total-items="store.left.totalItems"
                  :loading="store.left.isLoading"
                  item-key="ID"
                  class="elevation-1 fixed-header">

                <template slot="headerCell" slot-scope="props">
                  {{ props.header.text }} ({{ store.left.source[0] ? store.left.source[0]["COUNT_" + props.header.value.split(".").reverse()[0]] : 0 }})

                  <b-dropdown no-caret variant="link" class="header-filter-icon">
                    <template slot="button-content">
                      <!-- <i class="fa fa-filter text-muted"></i> -->
                      <v-icon small>filter_list</v-icon>
                    </template>

                    <b-dropdown-header>
                      <b-form-input type="text" placeholder="Filter" v-model="store.filters['left'][props.header.value.split('.').reverse()[0]]" @change="filterKeyup('left', props.header)"></b-form-input>
                    </b-dropdown-header>

                    <b-dropdown-item v-for="item in distinctData(props.header.value, store.left.source)" :key="item" @click="filterClick('left', props.header, item)">
                      {{ item }}
                    </b-dropdown-item>
                  </b-dropdown>
                </template>

                <v-progress-linear slot="progress" color="blue" indeterminate></v-progress-linear>

                <template slot="no-data">
                  <v-alert :value="store.left.isLoading" type="info">
                    Please wait while data is loading
                  </v-alert>

                  <v-alert :value="!store.left.isLoading" type="error">
                    Sorry, nothing to display here
                  </v-alert>
                </template>

                <template slot="items" slot-scope="props">
                    <td><b-link :to="{ path: addressPath + '/' + props.item.ID }"><tablecell :fulltext="props.item.PRIORITY_REPORT" :isklik="false"></tablecell></b-link></td>
                    <td><tablecell :fulltext="props.item.RR_LEAD" :isklik="true"></tablecell></td>
                    <td><tablecell :fulltext="props.item.BANK_ID" :isklik="true"></tablecell></td>
                </template>
              </v-data-table>
            </b-col>
            
            <b-col cols="6">
              <v-data-table
                  :headers="secondTableHeaders"
                  :items="store.right.display"
                  :pagination.sync="store.right.pagination"
                  :total-items="store.right.totalItems"
                  :loading="store.right.isLoading"
                  :expand="false"
                  v-if="store.isRightTable"
                  item-key="ID"
                  class="elevation-1">
                <v-progress-linear slot="progress" color="blue" indeterminate></v-progress-linear>

                <template slot="no-data">
                  <v-alert :value="store.right.isLoading" type="info">
                    Please wait while data is loading
                  </v-alert>

                  <v-alert :value="!store.right.isLoading" type="error">
                    Sorry, nothing to display here
                  </v-alert>
                </template>

                <template slot="headerCell" slot-scope="props">
                  {{ props.header.text }} {{ props.header.displayCount ? "(" + (store.right.source[0] ? store.right.source[0]["COUNT_" + props.header.value.split(".").reverse()[0]] : 0) + ")" : "" }}

                  <b-dropdown no-caret variant="link" class="header-filter-icon">
                    <template slot="button-content">
                      <!-- <i class="fa fa-filter text-muted"></i> -->
                      <v-icon small>filter_list</v-icon>
                    </template>

                    <b-dropdown-header>
                      <b-form-input type="text" placeholder="Filter" v-model="store.filters['right'][props.header.value.split('.').reverse()[0]]" @change="filterKeyup('right', props.header)"></b-form-input>
                    </b-dropdown-header>

                    <b-dropdown-item v-for="item in distinctData(props.header.value, store.right.source)" :key="item" @click="filterClick('right', props.header, item)">
                      {{ item }}
                    </b-dropdown-item>
                  </b-dropdown>
                </template>

                <template slot="items" slot-scope="props">
                  <tr>
                    <td style="width: calc(100% / 7)"><b-link @click="props.expanded = !props.expanded"><tablecell :fulltext="props.item.PRINCIPAL_RISK" :isklik="false"></tablecell></b-link></td>
                    <td style="width: calc(100% / 7)"><tablecell :fulltext="(_.uniq(_.map(props.item.RiskSubTypesVal, 'RISK_SUB')).filter(Boolean).join(', '))" :isklik="true"></tablecell></td>
                    <td style="width: calc(100% / 7)"><tablecell :fulltext="(_.uniq(_.map(props.item.RiskSubTypesVal, 'PR_RATIONALE')).filter(Boolean).join(', '))" :isklik="true"></tablecell></td>
                    <td style="width: calc(100% / 7)"><tablecell :fulltext="(_.uniq(_.map(props.item.RiskSubTypesVal, 'CRM_NAME')).filter(Boolean).join(', '))" :isklik="true"></tablecell></td>
                    <td style="width: calc(100% / 7)"><tablecell :fulltext="(_.uniq(_.map(props.item.RiskSubTypesVal, 'CRM_RATIONALE')).filter(Boolean).join(', '))" :isklik="true"></tablecell></td>
                    <td style="width: calc(100% / 7)"><tablecell :fulltext="(_.uniq(_.map(props.item.RiskSubTypesVal, 'ASSOC_CDES')).filter(Boolean).join(', '))" :isklik="true"></tablecell></td>
                    <td style="width: calc(100% / 7)"><tablecell :fulltext="(_.uniq(_.map(props.item.RiskSubTypesVal, 'CDE_RATIONALE')).filter(Boolean).join(', '))" :isklik="true"></tablecell></td>
                  </tr>
                </template>
                
                <template slot="expand" slot-scope="props">
                  <v-data-table
                    :headers="secondTableHeaders"
                    :items="props.item.RiskSubTypes"
                    class="elevation-1"
                    item-key="ID"
                    hide-actions
                    hide-headers
                  >
                    <template slot="items" slot-scope="props">
                      <td style="width: calc(100% / 7)">&nbsp;</td>
                      <td style="width: calc(100% / 7)"><b-link @click="props.expanded = !props.expanded"><tablecell :fulltext="props.item.RISK_SUB" :isklik="false"></tablecell></b-link></td>
                      <td style="width: calc(100% / 7)"><tablecell :fulltext="(_.uniq(_.map(props.item.PriorityReportsVal, 'PR_RATIONALE')).filter(Boolean).join(', '))" :isklik="true"></tablecell></td>
                      <td style="width: calc(100% / 7)"><tablecell :fulltext="(_.uniq(_.map(props.item.PriorityReportsVal, 'CRM_NAME')).filter(Boolean).join(', '))" :isklik="true"></tablecell></td>
                      <td style="width: calc(100% / 7)"><tablecell :fulltext="(_.uniq(_.map(props.item.PriorityReportsVal, 'CRM_RATIONALE')).filter(Boolean).join(', '))" :isklik="true"></tablecell></td>
                      <td style="width: calc(100% / 7)"><tablecell :fulltext="(_.uniq(_.map(props.item.PriorityReportsVal, 'ASSOC_CDES')).filter(Boolean).join(', '))" :isklik="true"></tablecell></td>
                      <td style="width: calc(100% / 7)"><tablecell :fulltext="(_.uniq(_.map(props.item.PriorityReportsVal, 'CDE_RATIONALE')).filter(Boolean).join(', '))" :isklik="true"></tablecell></td>
                    </template>

                    <template slot="expand" slot-scope="props">
                      <v-data-table
                        :headers="secondTableHeaders"
                        :items="props.item.PriorityReports"
                        class="elevation-1"
                        item-key="ID"
                        hide-actions
                        hide-headers
                      >
                        <template slot="items" slot-scope="props">
                          <td style="width: calc(100% / 7)">&nbsp;</td>
                          <td style="width: calc(100% / 7)">&nbsp;</td>
                          <td style="width: calc(100% / 7)"><b-link @click="props.expanded = !props.expanded"><tablecell :fulltext="props.item.PR_RATIONALE" :isklik="false"></tablecell></b-link></td>
                          <td style="width: calc(100% / 7)"><tablecell :fulltext="(_.uniq(_.map(props.item.CRMNamesVal, 'CRM_NAME')).filter(Boolean).join(', '))" :isklik="true"></tablecell></td>
                          <td style="width: calc(100% / 7)"><tablecell :fulltext="(_.uniq(_.map(props.item.CRMNamesVal, 'CRM_RATIONALE')).filter(Boolean).join(', '))" :isklik="true"></tablecell></td>
                          <td style="width: calc(100% / 7)"><tablecell :fulltext="(_.uniq(_.map(props.item.CRMNamesVal, 'ASSOC_CDES')).filter(Boolean).join(', '))" :isklik="true"></tablecell></td>
                          <td style="width: calc(100% / 7)"><tablecell :fulltext="(_.uniq(_.map(props.item.CRMNamesVal, 'CDE_RATIONALE')).filter(Boolean).join(', '))" :isklik="true"></tablecell></td>
                        </template>

                        <template slot="expand" slot-scope="props">
                          <v-data-table
                            :headers="secondTableHeaders"
                            :items="props.item.CRMNames"
                            class="elevation-1"
                            item-key="ID"
                            hide-actions
                            hide-headers
                          >
                            <template slot="items" slot-scope="props">
                              <td style="width: calc(100% / 7)">&nbsp;</td>
                              <td style="width: calc(100% / 7)">&nbsp;</td>
                              <td style="width: calc(100% / 7)">&nbsp;</td>
                              <td style="width: calc(100% / 7)"><b-link @click="props.expanded = !props.expanded"><tablecell :fulltext="props.item.CRM_NAME" :isklik="false"></tablecell></b-link></td>
                              <td style="width: calc(100% / 7)"><tablecell :fulltext="(_.uniq(_.map(props.item.CDEsVal, 'CRM_RATIONALE')).filter(Boolean).join(', '))" :isklik="true"></tablecell></td>
                              <td style="width: calc(100% / 7)"><tablecell :fulltext="(_.uniq(_.map(props.item.CDEsVal, 'ASSOC_CDES')).join(', '))" :isklik="true"></tablecell></td>
                              <td style="width: calc(100% / 7)"><tablecell :fulltext="(_.uniq(_.map(props.item.CDEsVal, 'CDE_RATIONALE')).filter(Boolean).join(', '))" :isklik="true"></tablecell></td>
                            </template>

                            <template slot="expand" slot-scope="props">
                              <v-data-table
                                :headers="secondTableHeaders"
                                :items="props.item.CDEs"
                                class="elevation-1"
                                item-key="ID"
                                hide-actions
                                hide-headers
                              >
                                <template slot="items" slot-scope="props">
                                  <td style="width: calc(100% / 7)">&nbsp;</td>
                                  <td style="width: calc(100% / 7)">&nbsp;</td>
                                  <td style="width: calc(100% / 7)">&nbsp;</td>
                                  <td style="width: calc(100% / 7)">&nbsp;</td>
                                  <td style="width: calc(100% / 7)"><tablecell :fulltext="(_.uniq(_.map(props.item.Rationales, 'CRM_RATIONALE')).filter(Boolean).join(', '))" :isklik="true"></tablecell></td>
                                  <td style="width: calc(100% / 7)"><b-link @click="props.expanded = !props.expanded"><tablecell :fulltext="props.item.ASSOC_CDES" :isklik="false"></tablecell></b-link></td>
                                  <td style="width: calc(100% / 7)"><tablecell :fulltext="(_.uniq(_.map(props.item.Rationales, 'CDE_RATIONALE')).filter(Boolean).join(', '))" :isklik="true"></tablecell></td>
                                </template>

                                <template slot="expand" slot-scope="props">
                                  <v-data-table
                                    :headers="secondTableHeaders"
                                    :items="props.item.Rationales"
                                    class="elevation-1"
                                    item-key="ID"
                                    hide-actions
                                    hide-headers
                                  >
                                    <template slot="items" slot-scope="props">
                                      <td style="width: calc(100% / 7)">&nbsp;</td>
                                      <td style="width: calc(100% / 7)">&nbsp;</td>
                                      <td style="width: calc(100% / 7)">&nbsp;</td>
                                      <td style="width: calc(100% / 7)">&nbsp;</td>
                                      <td style="width: calc(100% / 7)"><tablecell :fulltext="props.item.CRM_RATIONALE" :isklik="true"></tablecell></td>
                                      <td style="width: calc(100% / 7)">&nbsp;</td>
                                      <td style="width: calc(100% / 7)"><tablecell :fulltext="props.item.CDE_RATIONALE" :isklik="true"></tablecell></td>
                                    </template>
                                  </v-data-table>
                                </template>
                              </v-data-table>
                            </template>
                          </v-data-table>
                        </template>
                      </v-data-table>
                    </template>
                  </v-data-table>
                </template>
              </v-data-table>
            </b-col>
          </b-row>
        </b-col>
      </b-row>
    </b-col>
  </b-row>
</template>

<script>
import Vue from 'vue'
import { mapState, mapActions } from 'vuex'
import JsonExcel from 'vue-json-excel'
import pageSearch from '../PageSearch.vue'
import pageExport from '../PageExport.vue'
import tablecell from '../Tablecell.vue'
import pageLoader from '../PageLoader.vue'
 
Vue.component('downloadExcel', JsonExcel)

export default {
    components: {
      pageSearch, pageExport, tablecell, pageLoader
    },
    data () {
      return {
        storeName: "rfoall",
        systemSource: [],
        tablenameSource: [],
        firstTableHeaders: [
          { text: 'Priority Report Names', align: 'left', value: 'PRIORITY_REPORT', sortable: false },
          { text: 'Risk Reporting Load', align: 'left', value: 'RR_LEAD', sortable: false },
          { text: 'Bank ID', align: 'left', value: 'BANK_ID', sortable: false },
        ],
        secondTableHeaders: [
          { text: 'Principal Risk Type', align: 'left', sortable: false, value: 'PRINCIPAL_RISK', width: "25%" },
          { text: 'Risk Sub Type', align: 'left', sortable: false, value: 'RiskSubTypesVal.RISK_SUB', width: "25%" },
          { text: 'Priority Report Rationale', align: 'left', sortable: false, value: 'RiskSubTypesVal.PR_RATIONALE', width: "25%" },
          { text: 'CRM Name', align: 'left', sortable: false, value: 'RiskSubTypesVal.CRM_NAME', width: "25%" },
          { text: 'CRM Rationale', align: 'left', sortable: false, value: 'RiskSubTypesVal.CRM_RATIONALE', width: "25%" },
          { text: 'Associated CDEs', align: 'left', sortable: false, value: 'RiskSubTypesVal.ASSOC_CDES', width: "25%" },
          { text: 'CDE Rationale', align: 'left', sortable: false, value: 'RiskSubTypesVal.CDE_RATIONALE', width: "25%" },
        ],
      }
    },
    computed: {
      store () {
        return this.$store.state[this.storeName].all
      },
      addressPath (){
        var tmp = this.$route.path.split("/")
        return tmp.slice(0, 3).join("/")
      },
      principalRiskMaster () {
        return this._.uniq(this._.map(this.rfomy.right.source, 'PRINCIPAL_RISK'))
      },
      subRiskMaster () {
        return this._.uniq(this._.map(this.rfomy.right.source, 'RISK_SUB'))
      },
      searchDropdownInputs () {
        return [
          { type: "text", label: "Priority Report Name", source: "PriorityReportName" },
          { type: "text", label: "Risk Reporting Lead", source: "RiskReportingLead" },
          { type: "dropdown", label: "Principal Risk Type", source: "PrincipalRiskType", options: this._.map(this.store.right.source, 'PRINCIPAL_RISK') },
          { type: "dropdown", label: "Sub Risk Type", source: "SubRiskType", options: this._.map(this.store.right.source, 'RISK_SUB') },
        ]
      },
    },
    watch: {
      $route (to){
        this.store.isRightTable = false;

        if (to.params != undefined) {
          this.store.isRightTable = to.params.system; 
        }

        if(this.store.isRightTable){
          this.doGetRightTable(this.$route.params.system);
        }
      },
      "store.left.pagination": {
        handler () {
          this.doGetLeftTable();
        },
        deep: true
      },
      "store.right.pagination": {
        handler () {
          if(this.store.isRightTable){
            this.doGetRightTable(this.$route.params.system);
          }
        },
        deep: true
      },
      "store.searchMain" (val, oldVal){
        if(val || oldVal) {
          this.doGetLeftTable();

          if(this.store.isRightTable){
            this.doGetRightTable(this.$route.params.system);
          }
        }
      }
    },
    mounted() {
      this.store.tabName = this.storeName;
      this.store.isRightTable = this.$route.params.system;
    },
    methods: {
      getLeftTable () {
        this.$store.dispatch(`${this.storeName}/getLeftTable`)
      },
      getRightTable (id) {
        this.$store.dispatch(`${this.storeName}/getRightTable`, id)
      },
      doGetLeftTable () {
        this.getLeftTable();
      },
      doGetRightTable (id) {
        this.getRightTable(id);
      },
      filterKeyup (type, keyModel) {
        // this.columnFilter(type, keyModel);
        if(type == "left") this.doGetLeftTable()
        else this.doGetRightTable(this.$route.params.system)
      },
      filterClick (type, keyModel, val) {
        this.store.filters[type][keyModel.value.split('.').reverse()[0]] = val;

        // this.columnFilter(type, keyModel);
        if(type == "left") this.doGetLeftTable()
        else this.doGetRightTable(this.$route.params.system)
      },
      distinctData (col, datax) {
        var cols = col.split(".")
        if(cols.length > 1){
          var a = datax;

          cols.forEach((c, i) => {
            a = this._.flattenDeep(this._.map(this._.sortBy(a, c), c));
          });

          return this._.uniq(a).filter(Boolean);
        }
        
        return this._.uniq(
            this._.map(this._.sortBy(datax, col), col)
          ).filter(Boolean);
      },
      systemRowClick (evt) {
        evt.preventDefault();
        this.store.isRightTable = true;
      },
      getCDEConclusion (cdes) {
        return cdes.filter(Boolean).join(', ').indexOf("Yes") != -1 ? "Yes" : "No";
      },
    }
}
</script>
