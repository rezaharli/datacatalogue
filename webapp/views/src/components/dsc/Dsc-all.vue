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
                <input v-model="dscall.searchMain" type="text" class="form-control" placeholder="Search" aria-label="Recipient's username" aria-describedby="basic-addon2">
                <div class="input-group-append">
                  <b-dropdown right id="ddown1" text="" ref="ddownSearch">
                    <b-container>
                      <b-form-row class="main-table-search-dropdown-form">
                        <b-col>
                          <b-form @submit="onSubmit" @reset="onReset">
                            <b-form-group horizontal :label-cols="4" breakpoint="md" label="System Name" label-for="SystemName">
                              <b-form-input id="SystemName" type="text" v-model="dscall.searchDropdown.SystemName"></b-form-input>
                            </b-form-group>

                            <b-form-group horizontal :label-cols="4" breakpoint="md" label="ITAM ID" label-for="ItamID">
                              <b-form-input id="ItamID" type="text" v-model="dscall.searchDropdown.ItamID"></b-form-input>
                            </b-form-group>

                            <b-form-group horizontal :label-cols="4" breakpoint="md" label="Table Name" label-for="TableName">
                              <b-form-select id="TableName" :options="tablenameMaster" v-model="dscall.searchDropdown.TableName"></b-form-select>
                            </b-form-group>

                            <b-form-group horizontal :label-cols="4" breakpoint="md" label="Column Name" label-for="ColumnName">
                              <b-form-select id="ColumnName" :options="columnNameMaster" v-model="dscall.searchDropdown.ColumnName"></b-form-select>
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
            <b-col cols=6>
              <v-data-table
                  :headers="firstTableHeaders"
                  :items="dscall.left.display"
                  :pagination.sync="dscall.left.pagination"
                  :total-items="dscall.left.totalItems"
                  :loading="dscall.left.loading"
                  :expand="false"
                  item-key="ID"
                  class="elevation-1">

                <template slot="headerCell" slot-scope="props">
                  {{ props.header.text }} ({{ dscall.left.source[0] ? dscall.left.source[0]["COUNT_" + props.header.value.split(".").reverse()[0]] : 0 }})

                  <b-dropdown no-caret variant="link" class="header-filter-icon">
                    <template slot="button-content">
                      <!-- <i class="fa fa-filter text-muted"></i> -->
                      <v-icon small>filter_list</v-icon>
                    </template>

                    <b-dropdown-header>
                      <b-form-input type="text" placeholder="Filter" v-model="search['systems'][props.header.value]" @change="filterKeyup('systems', props.header)"></b-form-input>
                    </b-dropdown-header>

                    <b-dropdown-item v-for="item in distinctData(props.header.value, dscall.left.source)" :key="item" @click="columnFilter('systems', props.header, item)">
                      {{ item }}
                    </b-dropdown-item>
                  </b-dropdown>
                </template>

                <v-progress-linear slot="progress" color="blue" indeterminate></v-progress-linear>

                <template slot="no-data">
                  <v-alert :value="dscall.left.loading" type="info">
                    Please wait while data is loading
                  </v-alert>

                  <v-alert :value="!dscall.left.loading" type="error">
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
            
            <b-col cols=6>
              <v-data-table
                  :headers="secondTableHeaders"
                  :items="dscall.right.display"
                  :pagination.sync="dscall.right.pagination"
                  :total-items="dscall.right.totalItems"
                  :loading="dscall.right.loading"
                  :expand="false"
                  v-if="secondtable"
                  item-key="ID"
                  class="elevation-1">
                <v-progress-linear slot="progress" color="blue" indeterminate></v-progress-linear>

                <template slot="no-data">
                  <v-alert :value="dscall.right.loading" type="info">
                    Please wait while data is loading
                  </v-alert>

                  <v-alert :value="!dscall.right.loading" type="error">
                    Sorry, nothing to display here
                  </v-alert>
                </template>

                <template slot="headerCell" slot-scope="props">
                  {{ props.header.text }} {{ props.header.displayCount ? "(" + (dscall.right.source[0] ? dscall.right.source[0]["COUNT_" + props.header.value.split(".").reverse()[0]] : 0) + ")" : "" }}

                  <b-dropdown no-caret variant="link" class="header-filter-icon">
                    <template slot="button-content">
                      <!-- <i class="fa fa-filter text-muted"></i> -->
                      <v-icon small>filter_list</v-icon>
                    </template>

                    <b-dropdown-header>
                      <b-form-input type="text" placeholder="Filter" v-model="search['tablename'][props.header.value]" @change="filterKeyup('tablename', props.header)"></b-form-input>
                    </b-dropdown-header>

                    <b-dropdown-item v-for="item in distinctData(props.header.value, dscall.right.source)" :key="item" @click="columnFilter('tablename', props.header, item)">
                      {{ item }}
                    </b-dropdown-item>
                  </b-dropdown>
                </template>

                <template slot="items" slot-scope="props">
                  <tr>
                    <td><b-link @click="props.expanded = !props.expanded"><tablecell :fulltext="props.item.TABLE_NAME" :isklik="false"></tablecell></b-link></td>
                    <td><tablecell :fulltext="(_.map(props.item.Columns, 'COLUMN_NAME').filter(Boolean).join(', '))" :isklik="true"></tablecell></td>
                    <td><tablecell :fulltext="(_.map(props.item.Columns, 'BUSINESS_ALIAS_NAME').filter(Boolean).join(', '))" :isklik="true"></tablecell></td>
                    <td><tablecell :fulltext="getCDEConclusion(_.map(props.item.Columns, 'CDE_YES_NO'))" :isklik="true"></tablecell></td>
                  </tr>
                </template>
                
                <template slot="expand" slot-scope="props">
                  <v-data-table
                    :headers="secondTableHeaders"
                    :items="props.item.Columns"
                    class="elevation-1"
                    hide-actions
                    hide-headers
                  >
                    <template slot="items" slot-scope="props">
                      <td style="width: 25%">&nbsp;</td>
                      <td style="width: 25%"><b-link @click="showDetails(props.item)">{{ props.item.COLUMN_NAME }}</b-link></td>
                      <td style="width: 25%">{{ props.item.BUSINESS_ALIAS_NAME }}</td>
                      <td style="width: 25%">{{ props.item.CDE_YES_NO }}</td>
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
        firstTableHeaders: [
          { text: 'System Name', align: 'left', value: 'SYSTEM_NAME', sortable: false },
          { text: 'ITAM ID', align: 'left', value: 'Custodians.ITAM_ID', sortable: false },
          { text: 'Dataset Custodian', align: 'left', value: 'Custodians.DATASET_CUSTODIAN', sortable: false },
          { text: 'Bank ID', align: 'left', value: 'Custodians.BANK_ID', sortable: false }
        ],
        secondTableHeaders: [
          { text: 'Table Name', align: 'left', sortable: false, value: 'TABLE_NAME', displayCount: true, width: "25%" },
          { text: 'Column Name', align: 'left', sortable: false, value: 'Columns.COLUMN_NAME', displayCount: true, width: "25%" },
          { text: 'Business Alias Name', align: 'left', sortable: false, value: 'Columns.BUSINESS_ALIAS_NAME', displayCount: true, width: "25%" },
          { text: 'CDE (Yes/No)', align: 'left', sortable: false, value: 'Columns.CDE_YES_NO', displayCount: true, width: "25%" }
        ],
      }
    },
    computed: {
      ...mapState({
        dscall: state => state.dscall.all
      }),
      addressPath (){
        var tmp = this.$route.path.split("/")
        return tmp.slice(0, 3).join("/")
      },
      tablenameMaster (){
        return this._.map(this.dscall.right.source, 'TABLE_NAME')
      },
      columnNameMaster (){
        return this._.map(this._.flattenDeep(this._.map(this.dscall.right.source, 'Columns')), 'COLUMN_NAME')
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

        this._.each(this.dscall.left.display, (system, i) => {
          var temp = {
            SYSTEM_NAME: system.SYSTEM_NAME,
            ITAM_ID: _.uniq(_.map(system.Custodians, "ITAM_ID").filter(Boolean)).join(', '),
            DATASET_CUSTODIAN: _.uniq(_.map(system.Custodians, "DATASET_CUSTODIAN").filter(Boolean)).join(', '),
            BANK_ID: _.uniq(_.map(system.Custodians, "BANK_ID").filter(Boolean)).join(', '),
          }

          var tables = this._.filter(this.dscall.right.display, (v) => v.TSID == system.ID)
          if(tables.length > 0){
            this._.each(tables, (table, i) => {
              var tableLevel = _.cloneDeep(temp);
              tableLevel.TABLE_NAME = table.TABLE_NAME;

              if(table.Columns.length > 0){
                this._.each(table.Columns, (column, j) => {
                  var colLevel = _.cloneDeep(tableLevel);
                  colLevel.COLUMN_NAME = column.COLUMN_NAME;
                  colLevel.BUSINESS_ALIAS_NAME = column.BUSINESS_ALIAS_NAME;
                  colLevel.CDE_YES_NO = column.CDE_YES_NO;
                  
                  res.push(_.cloneDeep(colLevel));
                })
              } else {
                res.push(_.cloneDeep(tableLevel));
              }
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
      "dscall.left.pagination": {
        handler () {
          this.getLeftTable();
        },
        deep: true
      },
      "dscall.right.pagination": {
        handler () {
          if(this.secondtable){
            this.getRightTable(this.$route.params.system);
          }
        },
        deep: true
      },
      "dscall.searchMain" (val, oldVal){
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
      ...mapActions('dscall', {
          getLeftTable: 'getLeftTable',
          getRightTable: 'getRightTable',
      }),
      columnFilter (type, keyModel, val) {
        if(val == ""){
          if(type == "systems"){
            this.dscall.left.display = this.dscall.left.source;
          } else {
            this.dscall.right.display = this.dscall.right.source;
          }
          
          return
        }

        if(type == "systems"){
          if(keyModel.value.split(".")[1]){
            this.dscall.left.display = _.cloneDeep(this.dscall.left.source);
            this.dscall.left.display = this.dscall.left.display.filter(
              v => {
                var key = keyModel.value.split(".")[0];
                
                v[key] = v[key].filter(
                  w => w[keyModel.value.split(".")[1]].toString().toUpperCase() == val.toString().toUpperCase()
                )

                return v[key].length > 0;
              }
            );
          } else {
            this.dscall.left.display = _.filter(this.dscall.left.source, (v) => {
              return v[keymodel.value].toString().toUpperCase() == val.toString().toUpperCase();
            });
          }
        } else {
          if(keyModel.value.split(".")[1]){
            this.dscall.right.display = _.cloneDeep(this.dscall.right.source);
            this.dscall.right.display = this.dscall.right.display.filter(
              v => {
                var key = keyModel.value.split(".")[0];
                
                v[key] = v[key].filter(
                  w => w[keyModel.value.split(".")[1]].toString().toUpperCase() == val.toString().toUpperCase()
                )

                return v[key].length > 0;
              }
            );
          } else {
            this.dscall.right.display = _.filter(this.dscall.right.source, (v) => {
              return v[keymodel.value].toString().toUpperCase() == val.toString().toUpperCase();
            });
          }
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
      getCDEConclusion (cdes) {
        return cdes.filter(Boolean).join(', ').indexOf("Yes") != -1 ? "Yes" : "No";
      },
      onSubmit (evt) {
        if(evt) evt.preventDefault();

        this.getLeftTable();

        if(this.secondtable){
          this.getRightTable(this.$route.params.system);
        }

        this.$refs.ddownSearch.hide(true);
        // this.dscall.searchDropdown.show = false;
      },
      onReset (evt) {
        evt.preventDefault();
        /* Reset our form values */
        this.dscall.searchDropdown.SystemName = '';
        this.dscall.searchDropdown.ItamID = '';
        this.dscall.searchDropdown.TableName = '';
        this.dscall.searchDropdown.ColumnName = '';

        this.onSubmit();

        // /* Trick to reset/clear native browser form validation state */
        // this.dscall.searchDropdown.show = false;
        // this.$nextTick(() => { this.dscall.searchDropdown.show = true });
      },
      showDetails (param) {
        this.$router.push(this.addressPath + "/" + param.TSID + '/' + param.ID + '/' + param.COLID)
      }
    }
}
</script>
