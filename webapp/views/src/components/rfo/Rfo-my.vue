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
      <!-- Rfo details -->
      <router-view/>

      <!-- Main content -->
      <b-row>
        <b-col>
          <b-row>
            <b-col>
              <div class="input-group mb-3">
                <input v-model="rfomy.searchMain" type="text" class="form-control" placeholder="Search" aria-label="Recipient's username" aria-describedby="basic-addon2">
                <div class="input-group-append">
                  <b-dropdown right id="ddown1" text="">
                    <b-container>
                      <b-form-row class="main-table-search-dropdown-form">
                        <b-col>
                          <b-form @submit="onSubmit" @reset="onReset">
                            <b-form-group horizontal :label-cols="4" breakpoint="md" label="Priority Report Name" label-for="priorityReport">
                              <b-form-input id="priorityReport" type="text" v-model="searchForm.priorityReport"></b-form-input>
                            </b-form-group>

                            <b-form-group horizontal :label-cols="4" breakpoint="md" label="Risk Reporting Load" label-for="riskReporting">
                              <b-form-input id="riskReporting" type="text" v-model="searchForm.riskReporting"></b-form-input>
                            </b-form-group>

                            <b-form-group horizontal :label-cols="4" breakpoint="md" label="Country" label-for="country">
                              <b-form-select id="country" :options="countryMaster" v-model="searchForm.country"></b-form-select>
                            </b-form-group>

                            <b-form-group horizontal :label-cols="4" breakpoint="md" label="Principal Risk Type" label-for="principalRisk">
                              <b-form-select id="principalRisk" :options="principalRiskMaster" v-model="searchForm.principalRisk"></b-form-select>
                            </b-form-group>

                            <b-form-group horizontal :label-cols="4" breakpoint="md" label="Sub Risk Type" label-for="subRisk">
                              <b-form-select id="subRisk" :options="subRiskMaster" v-model="searchForm.subRisk"></b-form-select>
                            </b-form-group>

                            <b-button-group class="mx-1 float-right">
                              <b-button type="reset" variant="danger">Reset</b-button>
                              <b-button type="submit" variant="primary">Submit</b-button>
                            </b-button-group>
                          </b-form>
                        </b-col>
                      </b-form-row>
                    </b-container>
                  </b-dropdown>
                </div>
              </div>
            </b-col>

            <b-col></b-col>
            <b-col></b-col>
            <b-col>
              <download-excel
                  :data   = "excelData"
                  :fields = "excelFields"
                  worksheet = "My Worksheet"
                  name    = "filename.xls">
              
                  <b-btn size="sm" class="float-right" variant="success">Export</b-btn>
              </download-excel>
            </b-col>
          </b-row>

          <b-row>
            <b-col>
              <v-data-table
                  :headers="firstTableHeaders"
                  :items="rfomy.left.display"
                  :pagination.sync="rfomy.left.pagination"
                  :total-items="rfomy.left.totalItems"
                  :loading="rfomy.left.loading"
                  class="elevation-1 fixed-header">

                <template slot="headerCell" slot-scope="props">
                  {{ props.header.text }} ({{ distinctData(props.header.value, rfomy.left.source).length }})

                  <b-dropdown no-caret variant="link" class="header-filter-icon">
                    <template slot="button-content">
                      <!-- <i class="fa fa-filter text-muted"></i> -->
                      <v-icon small>filter_list</v-icon>
                    </template>

                    <b-dropdown-header>
                      <b-form-input type="text" placeholder="Filter" v-model="search['systems'][props.header.value]" @change="filterKeyup('systems', props.header)"></b-form-input>
                    </b-dropdown-header>

                    <b-dropdown-item v-for="item in distinctData(props.header.value, rfomy.left.source)" :key="item" @click="columnFilter('systems', props.header, item)">
                      {{ item }}
                    </b-dropdown-item>
                  </b-dropdown>
                </template>

                <v-progress-linear slot="progress" color="blue" indeterminate></v-progress-linear>

                <template slot="no-data">
                  <v-alert :value="rfomy.left.loading" type="info">
                    Please wait while data is loading
                  </v-alert>

                  <v-alert :value="!rfomy.left.loading" type="error">
                    Sorry, nothing to display here
                  </v-alert>
                </template>

                <template slot="items" slot-scope="props">
                    <td><b-link :to="{ path: addressPath + '/' + props.item.ID }">{{ props.item.NAME }}</b-link></td>
                    <td>{{ props.item.OWNER_ID }}</td>
                </template>
              </v-data-table>
            </b-col>
            
            <b-col class="scrollableasdf">
              <v-data-table
                  :headers="secondTableHeaders"
                  :items="rfomy.right.display"
                  :pagination.sync="rfomy.right.pagination"
                  :total-items="rfomy.right.totalItems"
                  :loading="rfomy.right.loading"
                  v-if="secondtable"
                  item-key="ID"
                  class="elevation-1">
                <v-progress-linear slot="progress" color="blue" indeterminate></v-progress-linear>

                <template slot="no-data">
                  <v-alert :value="rfomy.right.loading" type="info">
                    Please wait while data is loading
                  </v-alert>

                  <v-alert :value="!rfomy.right.loading" type="error">
                    Sorry, nothing to display here
                  </v-alert>
                </template>

                <template slot="headerCell" slot-scope="props">
                  {{ props.header.text }} ({{ distinctData(props.header.value, rfomy.right.source).length }})

                  <b-dropdown no-caret variant="link" class="header-filter-icon">
                    <template slot="button-content">
                      <!-- <i class="fa fa-filter text-muted"></i> -->
                      <v-icon small>filter_list</v-icon>
                    </template>

                    <b-dropdown-header>
                      <b-form-input type="text" placeholder="Filter" v-model="search['tablename'][props.header.value]" @change="filterKeyup('tablename', props.header)"></b-form-input>
                    </b-dropdown-header>

                    <b-dropdown-item v-for="item in distinctData(props.header.value, rfomy.right.source)" :key="item" @click="columnFilter('tablename', props.header, item)">
                      {{ item }}
                    </b-dropdown-item>
                  </b-dropdown>
                </template>

                <template slot="items" slot-scope="props">
                  <tr>
                    <td>{{ props.item.PRINCIPAL_RISK_TYPE }}</td>
                    <td>{{ props.item.RISK_SUB_TYPE }}</td>
                    <td>{{ props.item.PRIORITY_REPORT_RATIONALE }}</td>
                    <td>{{ props.item.CRM_NAME }}</td>
                    <td>{{ props.item.CRM_RATIONALE }}</td>
                    <td>{{ props.item.ASSOCIATED_CDES }}</td>
                    <td>{{ props.item.CDE_RATIONALE }}</td>
                  </tr>
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
 
Vue.component('downloadExcel', JsonExcel)

export default {
    data () {
      return {
        search: {
          systems: {},
          tablename: {}
        },
        secondtable: false,
        systemSource: [],
        tablenameSource: [],
        searchForm: {
          priorityReport: '',
          riskReporting: '',
          country: '',
          principalRisk: '',
          subRisk: ''
        },
        firstTableHeaders: [
          { text: 'Priority Report Names', align: 'left', value: 'NAME', sortable: false },
          { text: 'Risk Reporting Load', align: 'left', value: 'OWNER_ID', sortable: false },
        ],
        secondTableHeaders: [
          { text: 'Principal Risk Type', align: 'left', sortable: false, value: 'PRINCIPAL_RISK_TYPE', width: "25%" },
          { text: 'Risk Sub Type', align: 'left', sortable: false, value: 'RISK_SUB_TYPE', width: "25%" },
          { text: 'Priority Report Rationale', align: 'left', sortable: false, value: 'PRIORITY_REPORT_RATIONALE', width: "25%" },
          { text: 'CRM Name', align: 'left', sortable: false, value: 'CRM_NAME', width: "25%" },
          { text: 'CRM Rationale', align: 'left', sortable: false, value: 'CRM_RATIONALE', width: "25%" },
          { text: 'Associated CDEs', align: 'left', sortable: false, value: 'ASSOCIATED_CDES', width: "25%" },
          { text: 'CDE Rationale', align: 'left', sortable: false, value: 'CDE_RATIONALE', width: "25%" },
        ],
      }
    },
    computed: {
      ...mapState({
        rfomy: state => state.rfomy.all
      }),
      addressPath (){
        var tmp = this.$route.path.split("/")
        return tmp.slice(0, 3).join("/")
      },
      countryMaster (){
        return []
      },
      principalRiskMaster () {
        return this._.uniq(this._.map(this.rfomy.right.source, 'PRINCIPAL_RISK_TYPE'))
      },
      subRiskMaster () {
        return this._.uniq(this._.map(this.rfomy.right.source, 'RISK_SUB_TYPE'))
      },
      excelFields (){
        var ret = {}

        _.each(this.firstTableHeaders, function(v){
          ret[v.text] = v.value.split(".").reverse()[0];
        })

        if(this.secondtable){
          _.each(this.secondTableHeaders, function(v){
            ret[v.text] = v.value.split(".").reverse()[0];
          })
        }

        return ret
      },
      excelData () {
        var res = [];

        this._.each(this.rfomy.left.display, (system, i) => {
          var temp = {
            NAME: system.NAME,
            OWNER_ID: system.OWNER_ID,
          }
          
          var tables = this._.filter(this.rfomy.right.display, (v) => v.ID == system.ID)
          if(tables.length > 0){
            this._.each(tables, (table, i) => {
              var tableLevel = _.cloneDeep(temp);
              tableLevel.PRINCIPAL_RISK_TYPE = table.PRINCIPAL_RISK_TYPE;
              tableLevel.RISK_SUB_TYPE = table.RISK_SUB_TYPE;
              tableLevel.PRIORITY_REPORT_RATIONALE = table.PRIORITY_REPORT_RATIONALE;
              tableLevel.CRM_NAME = table.CRM_NAME;
              tableLevel.CRM_RATIONALE = table.CRM_RATIONALE;
              tableLevel.ASSOCIATED_CDES = table.ASSOCIATED_CDES;
              tableLevel.CDE_RATIONALE = table.CDE_RATIONALE;

              res.push(_.cloneDeep(tableLevel));
            })
          } else {
            res.push(_.cloneDeep(temp));
          }
        });

        return res
      }
    },
    watch: {
      $route (to){
        this.secondtable = false;

        if (to.params != undefined) {
          this.secondtable = to.params.system; 
        }

        if(this.secondtable){
          this.getRightTable(this.$route.params.system);
        }
      },
      "rfomy.left.pagination": {
        handler () {
          this.getLeftTable();
        },
        deep: true
      },
      "rfomy.right.pagination": {
        handler () {
          if(this.secondtable){
            this.getRightTable(this.$route.params.system);
          }
        },
        deep: true
      },
      "rfomy.searchMain" (val, oldVal){
        if(val || oldVal) {
          this.getLeftTable();

          if(this.secondtable){
            this.getRightTable(this.$route.params.system);
          }
        }
      }
    },
    mounted() {
      this.secondtable = this.$route.params.system;
    },
    methods: {
      ...mapActions('rfomy', {
          getLeftTable: 'getLeftTable',
          getRightTable: 'getRightTable',
      }),
      columnFilter (type, keyModel, val) {
        if(val == ""){
          if(type == "systems"){
            this.rfomy.left.display = this.rfomy.left.source;
          } else {
            this.rfomy.right.display = this.rfomy.right.source;
          }
          return
        }

        if(type == "systems"){
          this.rfomy.left.display = _.filter(this.rfomy.left.source, [keyModel.value, val]);
        } else {
          this.rfomy.right.display = _.filter(this.rfomy.right.source, [keyModel.value, val]);
        }
      },
      filterKeyup (type, keyModel) {
        this.columnFilter(type, keyModel, this.search[type][keyModel.value]);
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
        this.secondtable = true;
      },
      onSubmit (evt) {
        evt.preventDefault();

        this.rfomy.left.display = this.rfomy.left.source;
        this.rfomy.right.display = this.rfomy.right.source;
        if(this.searchForm.priorityReport)
          this.rfomy.left.display = this._.filter(this.rfomy.left.display, (val) => val.NAME.toString().toUpperCase().indexOf(this.searchForm.priorityReport.toString().toUpperCase()) != -1);
        if(this.searchForm.riskReporting)
          this.rfomy.left.display = this._.filter(this.rfomy.left.display, (val) => val.OWNER_ID.toString().toUpperCase().indexOf(this.searchForm.riskReporting.toString().toUpperCase()) != -1);

        if(this.searchForm.principalRisk)
          this.rfomy.right.display = this._.filter(this.rfomy.right.display, (val) => val.PRINCIPAL_RISK_TYPE.toString().toUpperCase().indexOf(this.searchForm.principalRisk.toString().toUpperCase()) != -1);
        if(this.searchForm.subRisk)
          this.rfomy.right.display = this._.filter(this.rfomy.right.display, (val) => val.RISK_SUB_TYPE.toString().toUpperCase().indexOf(this.searchForm.subRisk.toString().toUpperCase()) != -1);

        this.searchForm.show = false;
      },
      onReset (evt) {
        evt.preventDefault();
        /* Reset our form values */
        this.searchForm.priorityReport = '';
        this.searchForm.riskReporting = '';
        this.searchForm.country = '';
        this.searchForm.principalRisk = '';
        this.searchForm.subRisk = '';

        // /* Trick to reset/clear native browser form validation state */
        // this.searchForm.show = false;
        // this.$nextTick(() => { this.searchForm.show = true });
      },
      showDetails (id) {
        this.$router.push(this.addressPath + "/" + this.$route.params.system + '/' + id)
      }
    }
}
</script>
