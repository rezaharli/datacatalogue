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
          <b-row>
            <b-col>
              <div class="input-group mb-3">
                <input v-model="dscinterfaces.searchMain" type="text" class="form-control" placeholder="Search" aria-label="Recipient's username" aria-describedby="basic-addon2">

                <div class="input-group-append">
                  <b-dropdown right id="ddown1" text="" ref="ddownSearch">
                    <b-container>
                      <b-form-row class="main-table-search-dropdown-form">
                        <b-col>
                          <b-form @submit="onSubmit" @reset="onReset">
                            <b-form-group horizontal :label-cols="4" breakpoint="md" label="System Name" label-for="systemName">
                              <b-form-input id="systemName" type="text" v-model="dscinterfaces.searchDropdown.SystemName"></b-form-input>
                            </b-form-group>

                            <b-form-group horizontal :label-cols="4" breakpoint="md" label="ITAM ID" label-for="itamID">
                              <b-form-input id="itamID" type="text" v-model="dscinterfaces.searchDropdown.ItamID"></b-form-input>
                            </b-form-group>

                            <b-form-group horizontal :label-cols="4" breakpoint="md" label="Table Name" label-for="tableName">
                              <b-form-select id="tableName" :options="tablenameMaster" v-model="dscinterfaces.searchDropdown.TableName"></b-form-select>
                            </b-form-group>

                            <b-form-group horizontal :label-cols="4" breakpoint="md" label="Column Name" label-for="columnName">
                              <b-form-select id="columnName" :options="columnNameMaster" v-model="dscinterfaces.searchDropdown.ColumnName"></b-form-select>
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
                  :items="dscinterfaces.left.display"
                  :pagination.sync="dscinterfaces.left.pagination"
                  :total-items="dscinterfaces.left.totalItems"
                  :loading="dscinterfaces.left.loading"
                  :expand="false"
                  item-key="ID"
                  class="elevation-1">

                <template slot="headerCell" slot-scope="props">
                  {{ props.header.text }} ({{ distinctData(props.header.value, dscinterfaces.left.source).length }})

                  <b-dropdown no-caret variant="link" class="header-filter-icon">
                    <template slot="button-content">
                      <!-- <i class="fa fa-filter text-muted"></i> -->
                      <v-icon small>filter_list</v-icon>
                    </template>

                    <b-dropdown-header>
                      <b-form-input type="text" placeholder="Filter" v-model="search['systems'][props.header.value]" @change="filterKeyup('systems', props.header)"></b-form-input>
                    </b-dropdown-header>

                    <b-dropdown-item v-for="item in distinctData(props.header.value, dscinterfaces.left.source)" :key="item" @click="columnFilter('systems', props.header, item)">
                      {{ item }}
                    </b-dropdown-item>
                  </b-dropdown>
                </template>

                <v-progress-linear slot="progress" color="blue" indeterminate></v-progress-linear>

                <template slot="no-data">
                  <v-alert :value="dscinterfaces.left.loading" type="info">
                    Please wait while data is loading
                  </v-alert>

                  <v-alert :value="!dscinterfaces.left.loading" type="error">
                    Sorry, nothing to display here
                  </v-alert>
                </template>

                <template slot="items" slot-scope="props">
                    <td><b-link :to="{ path: addressPath + '/' + props.item.ID }"><tablecell :fulltext="props.item.SYSTEM_NAME" :isklik="false"></tablecell></b-link></td>
                    <td><b-link @click="props.expanded = !props.expanded"><tablecell :fulltext="(_.uniq(_.map(props.item.Custodians, 'ITAM_ID').filter(Boolean)).join(', '))" :isklik="true"></tablecell></b-link></td>
                    <td><tablecell :fulltext="(_.uniq(_.map(props.item.Custodians, 'DATASET_CUSTODIAN').filter(Boolean)).join(', '))" :isklik="true"></tablecell></td>
                    <td><tablecell :fulltext="(_.uniq(_.map(props.item.Custodians, 'BANK_ID').filter(Boolean)).join(', '))" :isklik="true"></tablecell></td>
                </template>

                <template slot="expand" slot-scope="props">
                  <v-data-table
                    :headers="firstTableHeaders"
                    :items="props.item.Custodians"
                    class="elevation-1"
                    hide-actions
                    hide-headers
                  >
                    <template slot="items" slot-scope="props">
                      <td style="width: 25%">&nbsp;</td>
                      <td style="width: 25%">&nbsp;</td>
                      <td style="width: 25%">{{ props.item.DATASET_CUSTODIAN }}</td>
                      <td style="width: 25%">{{ props.item.BANK_ID }}</td>
                    </template>
                  </v-data-table>
                </template>
              </v-data-table>
            </b-col>
            
            <b-col class="scrollableasdf">
              <v-data-table
                  :headers="secondTableHeaders"
                  :items="dscinterfaces.right.display"
                  :pagination.sync="dscinterfaces.right.pagination"
                  :total-items="dscinterfaces.right.totalItems"
                  :loading="dscinterfaces.right.loading"
                  v-if="secondtable"
                  item-key="R__"
                  class="elevation-1">
                <v-progress-linear slot="progress" color="blue" indeterminate></v-progress-linear>
                  <template slot="no-data">
                    <v-alert :value="dscinterfaces.right.loading" type="info">
                    Please wait while data is loading
                  </v-alert>

                  <v-alert :value="!dscinterfaces.right.loading" type="error">
                    Sorry, nothing to display here
                  </v-alert>
                  </template>

                  <template slot="headerCell" slot-scope="props">
                  {{ props.header.text }} ({{ distinctData(props.header.value, dscinterfaces.right.source).length }})

                  <b-dropdown no-caret variant="link" class="header-filter-icon">
                    <template slot="button-content">
                      <!-- <i class="fa fa-filter text-muted"></i> -->
                      <v-icon small>filter_list</v-icon>
                    </template>

                    <b-dropdown-header>
                      <b-form-input type="text" placeholder="Filter" v-model="search['tablename'][props.header.value]" @change="filterKeyup('tablename', props.header)"></b-form-input>
                    </b-dropdown-header>

                    <b-dropdown-item v-for="item in distinctData(props.header.value, dscinterfaces.right.source)" :key="item" @click="columnFilter('tablename', props.header, item)">
                      {{ item }}
                    </b-dropdown-item>
                  </b-dropdown>
                </template>

                <template slot="items" slot-scope="props">
                  <td><b-link :to="{ path:'/dsc/interfaces/' + props.item.TSID + '/' + props.item.ID + '/' + props.item.COLID }" href="#foo" v-b-modal.modallg>{{ props.item.LISTOF_CDE }}</b-link></td>
                  <!-- <td><b-link :to="{ path:'/dsc/interfaces/' + route.params.system + "/details" }" v-b-modal.modallg>{{ props.item.name }}</b-link></td> -->
                  <td><tablecell :fulltext="(_.uniq(_.map(props.item.Values, 'IMM_PREC_SYSTEM_NAME').filter(Boolean)).join(', '))" :isklik="true"></tablecell></td>
                  <td><tablecell :fulltext="(_.uniq(_.map(props.item.Values, 'IMM_PREC_SYSTEM_SLA').filter(Boolean)).join(', '))" :isklik="true"></tablecell></td>
                  <td><tablecell :fulltext="(_.uniq(_.map(props.item.Values, 'IMM_PREC_SYSTEM_OLA').filter(Boolean)).join(', '))" :isklik="true"></tablecell></td>
                  <td><tablecell :fulltext="(_.uniq(_.map(props.item.Values, 'IMM_SUCC_SYSTEM_NAME').filter(Boolean)).join(', '))" :isklik="true"></tablecell></td>
                  <td><tablecell :fulltext="(_.uniq(_.map(props.item.Values, 'IMM_SUCC_SYSTEM_SLA').filter(Boolean)).join(', '))" :isklik="true"></tablecell></td>
                  <td><tablecell :fulltext="(_.uniq(_.map(props.item.Values, 'IMM_SUCC_SYSTEM_OLA').filter(Boolean)).join(', '))" :isklik="true"></tablecell></td>
                  <td><tablecell :fulltext="(_.uniq(_.map(props.item.Values, 'LIST_DOWNSTREAM_PROCESS').filter(Boolean)).join(', '))" :isklik="true"></tablecell></td>
                  <td><tablecell :fulltext="(_.uniq(_.map(props.item.Values, 'DOWNSTREAM_OWNER').filter(Boolean)).join(', '))" :isklik="true"></tablecell></td>
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
import tablecell from '../Tablecell.vue'
 
Vue.component('downloadExcel', JsonExcel)

export default {
    components: {
      tablecell
    },
    data () {
      return {
        search: {
          systems: {},
          tablename: {}
        },
        secondtable: false,
        systemSource: [],
        tablenameSource: [],
        searchMain: '',
        firstTableHeaders: [
          { text: 'System Name', align: 'left', value: 'SYSTEM_NAME', sortable: false },
          { text: 'ITAM ID', align: 'left', value: 'Custodians.ITAM_ID', sortable: false },
          { text: 'Dataset Custodian', align: 'left', value: 'Custodians.DATASET_CUSTODIAN', sortable: false },
          { text: 'Bank ID', align: 'left', value: 'Custodians.BANK_ID', sortable: false }
        ],
        secondTableHeaders: [
          { text: 'List of CDEs', align: 'left', sortable: false, value: 'LISTOF_CDE', width: "25%" },
          { text: 'Immediate Preceding System', align: 'left', sortable: false, value: 'Values.IMM_PREC_SYSTEM_NAME', width: "25%" },
          { text: 'SLA(Yes/No)', align: 'left', sortable: false, value: 'Values.IMM_PREC_SYSTEM_SLA', width: "25%" },
          { text: 'OLA(Yes/No)', align: 'left', sortable: false, value: 'Values.IMM_PREC_SYSTEM_OLA', width: "25%" },
          { text: 'Immediate Succeeding System', align: 'left', sortable: false, value: 'Values.IMM_SUCC_SYSTEM_NAME', width: "25%" },
          { text: 'SLA (Yes/No)', align: 'left', sortable: false, value: 'Values.IMM_SUCC_SYSTEM_SLA', width: "25%" },
          { text: 'OLA (Yes/No)', align: 'left', sortable: false, value: 'Values.IMM_SUCC_SYSTEM_OLA', width: "25%" },
          { text: 'List of Downstream Process', align: 'left', sortable: false, value: 'Values.LIST_DOWNSTREAM_PROCESS', width: "25%" },
          { text: 'Downstream Process Owner', align: 'left', sortable: false, value: 'Values.DOWNSTREAM_OWNER', width: "25%" },
        ],
      }
    },
    computed: {
      ...mapState({
        dscinterfaces: state => state.dscinterfaces.all
      }),
      addressPath (){
        var tmp = this.$route.path.split("/")
        return tmp.slice(0, 3).join("/")
      },
      tablenameMaster (){
        return this._.map(this.dscinterfaces.right.source, 'TABLE_NAME')
      },
      columnNameMaster (){
        return this._.map(this._.flattenDeep(this._.map(this.dscinterfaces.right.source, 'Columns')), 'COLUMN_NAME')
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

        this._.each(this.dscinterfaces.left.display, (system, i) => {
          var temp = {
            SYSTEM_NAME: system.SYSTEM_NAME,
            ITAM_ID: system.ITAM_ID,
            DATASET_CUSTODIAN: system.DATASET_CUSTODIAN,
            BANK_ID: system.BANK_ID,
          }

          var tables = this._.filter(this.dscinterfaces.right.display, (v) => v.TSID == system.ID)
          if(tables.length > 0){
            this._.each(tables, (table, i) => {
              var tableLevel = _.cloneDeep(temp);
              tableLevel.LISTOF_CDE = table.LISTOF_CDE;
              tableLevel.IMM_PREC_SYSTEM_NAME = table.IMM_PREC_SYSTEM_NAME;
              tableLevel.IMM_PREC_SYSTEM_SLA = table.IMM_PREC_SYSTEM_SLA;
              tableLevel.IMM_PREC_SYSTEM_OLA = table.IMM_PREC_SYSTEM_OLA;
              tableLevel.IMM_SUCC_SYSTEM_NAME = table.IMM_SUCC_SYSTEM_NAME;
              tableLevel.IMM_SUCC_SYSTEM_SLA = table.IMM_SUCC_SYSTEM_SLA;
              tableLevel.IMM_SUCC_SYSTEM_OLA = table.IMM_SUCC_SYSTEM_OLA;
              tableLevel.LIST_DOWNSTREAM_PROCESS = table.LIST_DOWNSTREAM_PROCESS;
              tableLevel.DOWNSTREAM_OWNER = table.DOWNSTREAM_OWNER;

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
      "dscinterfaces.left.pagination": {
        handler () {
          this.getLeftTable();
        },
        deep: true
      },
      "dscinterfaces.right.pagination": {
        handler () {
          if(this.secondtable){
            this.getRightTable(this.$route.params.system);
          }
        },
        deep: true
      },
      "dscinterfaces.searchMain" (val, oldVal){
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
      ...mapActions('dscinterfaces', {
          getLeftTable: 'getLeftTable',
          getRightTable: 'getRightTable',
      }),
      columnFilter (type, keyModel, val) {
        if(val == ""){
          if(type == "systems"){
            this.dscinterfaces.left.display = this.dscinterfaces.left.source;
          } else {
            this.dscinterfaces.right.display = this.dscinterfaces.right.source;
          }
          return
        }

        if(type == "systems"){
          this.dscinterfaces.left.display = _.filter(this.dscinterfaces.left.source, [keyModel.value, val]);
        } else {
          this.dscinterfaces.right.display = _.filter(this.dscinterfaces.right.source, [keyModel.value, val]);
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
        if(evt) evt.preventDefault();

        this.getLeftTable();

        if(this.secondtable){
          this.getRightTable(this.$route.params.system);
        }

        this.$refs.ddownSearch.hide(true);
        // this.searchForm.show = false;
      },
      onReset (evt) {
        evt.preventDefault();
        /* Reset our form values */
        this.dscinterfaces.searchDropdown.SystemName = '';
        this.dscinterfaces.searchDropdown.ItamID = '';
        this.dscinterfaces.searchDropdown.TableName = '';
        this.dscinterfaces.searchDropdown.ColumnName = '';

        this.onSubmit();

        // /* Trick to reset/clear native browser form validation state */
        // this.searchForm.show = false;
        // this.$nextTick(() => { this.searchForm.show = true });
      }
    }
}
</script>
