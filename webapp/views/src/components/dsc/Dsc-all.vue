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
                <input v-model="searchMain" type="text" class="form-control" placeholder="Search" aria-label="Recipient's username" aria-describedby="basic-addon2">
                <div class="input-group-append">
                  <b-dropdown right id="ddown1" text="">
                    <b-container>
                      <b-form-row class="main-table-search-dropdown-form">
                        <b-col>
                          <b-form @submit="onSubmit" @reset="onReset">
                            <b-form-group horizontal :label-cols="4" breakpoint="md" label="System Name" label-for="systemName">
                              <b-form-input id="systemName" type="text" v-model="searchForm.systemName"></b-form-input>
                            </b-form-group>

                            <b-form-group horizontal :label-cols="4" breakpoint="md" label="ITAM ID" label-for="itamID">
                              <b-form-input id="itamID" type="text" v-model="searchForm.itamID"></b-form-input>
                            </b-form-group>

                            <b-form-group horizontal :label-cols="4" breakpoint="md" label="Country" label-for="country">
                              <b-form-select id="country" :options="searchForm.countryMaster" v-model="searchForm.country"></b-form-select>
                            </b-form-group>

                            <b-form-group horizontal :label-cols="4" breakpoint="md" label="Table Name" label-for="tableName">
                              <b-form-select id="tableName" :options="tablenameMaster" v-model="searchForm.tableName"></b-form-select>
                            </b-form-group>

                            <b-form-group horizontal :label-cols="4" breakpoint="md" label="Column Name" label-for="columnName">
                              <b-form-select id="columnName" :options="columnNameMaster" v-model="searchForm.columnName"></b-form-select>
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
                  :items="dscmy.systemsDisplay"
                  :loading="dscmy.systemsLoading"
                  :search="searchMain"
                  class="elevation-1 fixed-header">

                <template slot="headerCell" slot-scope="props">
                  {{ props.header.text }} ({{ distinctData(props.header.value, dscmy.systemsSource).length }})

                  <b-dropdown no-caret variant="link" class="header-filter-icon">
                    <template slot="button-content">
                      <i class="fa fa-filter text-muted"></i>
                    </template>

                    <b-dropdown-header>
                      <b-form-input type="text" placeholder="Filter" v-model="search['systems'][props.header.value]" @change="filterKeyup('systems', props.header)"></b-form-input>
                    </b-dropdown-header>

                    <b-dropdown-item v-for="item in distinctData(props.header.value, dscmy.systemsSource)" :key="item" @click="columnFilter('systems', props.header, item)">
                      {{ item }}
                    </b-dropdown-item>
                  </b-dropdown>
                </template>

                <v-progress-linear slot="progress" color="blue" indeterminate></v-progress-linear>

                <template slot="no-data">
                  <v-alert :value="true" color="error" icon="warning">
                    Sorry, nothing to display here :(
                  </v-alert>
                </template>

                <template slot="items" slot-scope="props">
                    <td><b-link :to="{ path:'/dsc/all/' + props.item.ID }">{{ props.item.System_Name }}</b-link></td>
                    <td>{{ props.item.ITAM_ID }}</td>
                    <td>{{ props.item.fat }}</td>
                    <td>{{ props.item.carbs }}</td>
                </template>
              </v-data-table>
            </b-col>
            
            <b-col>
              <v-data-table
                :headers="secondTableHeaders"
                :items="dscmy.tableDisplay"
                :loading="dscmy.tableLoading"
                :expand="false"
                v-if="secondtable"
                item-key="ID"
                class="elevation-1">
                <v-progress-linear slot="progress" color="blue" indeterminate></v-progress-linear>
                <v-progress-linear slot="progress" color="blue" indeterminate></v-progress-linear>
                  <template slot="no-data">
                    <v-alert :value="true" color="error" icon="warning">
                      Sorry, nothing to display here :(
                    </v-alert>
                  </template>

                  <template slot="headerCell" slot-scope="props">
                  {{ props.header.text }} ({{ distinctData(props.header.value, dscmy.tableSource).length }})

                  <b-dropdown no-caret variant="link" class="header-filter-icon">
                    <template slot="button-content">
                      <i class="fa fa-filter text-muted"></i>
                    </template>

                    <b-dropdown-header>
                      <b-form-input type="text" placeholder="Filter" v-model="search['tablename'][props.header.value]" @change="filterKeyup('tablename', props.header)"></b-form-input>
                    </b-dropdown-header>

                    <b-dropdown-item v-for="item in distinctData(props.header.value, dscmy.tableSource)" :key="item" @click="columnFilter('tablename', props.header, item)">
                      {{ item }}
                    </b-dropdown-item>
                  </b-dropdown>
                </template>

                <template slot="items" slot-scope="props">
                  <tr>
                    <td><b-link @click="props.expanded = !props.expanded">{{ props.item.Name }}</b-link></td>
                    <td><b-link @click="showDetails(props.item.ID)">{{ _.map(props.item.Columns, "Name").join(", ") }}</b-link></td>
                    <td>{{ _.map(props.item.Columns, "Alias_Name").join(", ") }}</td>
                    <td>{{ _.map(props.item.Columns, "CDE").join(", ") }}</td>
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
                      <td style="width: 25%"><b-link @click="showDetails(props.item.Table_ID)">{{ props.item.Name }}</b-link></td>
                      <td style="width: 25%">{{ props.item.Alias_Name }}</td>
                      <td style="width: 25%">{{ props.item.CDE }}</td>
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
        searchMain: '',
        searchForm: {
          systemName: '',
          itamID: '',
          country: '',
          countryMaster: ['a', 'c', 'd'],
          tableName: '',
          columnName: '',
        },
        firstTableHeaders: [
          { text: 'System Name', align: 'left', value: 'System_Name', sortable: false },
          { text: 'ITAM ID', align: 'left', value: 'ITAM_ID', sortable: false },
          { text: 'Dataset Custodian', align: 'left', value: 'fat', sortable: false },
          { text: 'Bank ID', align: 'left', value: 'carbs', sortable: false }
        ],
        secondTableHeaders: [
          { text: 'Table Name', align: 'left', sortable: false, value: 'Name', width: "25%" },
          { text: 'Column Name', align: 'left', sortable: false, value: 'Columns.Name', width: "25%" },
          { text: 'Business Alias Name', align: 'left', sortable: false, value: 'Columns.Alias_Name', width: "25%" },
          { text: 'CDE (Yes/No)', align: 'left', sortable: false, value: 'Columns.CDE', width: "25%" }
        ],
        excelFields: {
          'System Name': 'System_Name',
          'ITAM ID': 'ITAM_ID',
          'Dataset Custodian': 'asdf',
          'Bank ID': 'carbs',
          'Table Name': 'Table_Name',
          'Column Name': 'Column_Name',
          'Business Alias Name': 'Alias_Name',
          'CDE (Yes/No)': 'CDE'
        }
      }
    },
    computed: {
      ...mapState({
        dscmy: state => state.dscmy.all
      }),
      tablenameMaster (){
        return this._.map(this.dscmy.tableSource, 'Name')
      },
      columnNameMaster (){
        return this._.map(this._.flattenDeep(this._.map(this.dscmy.tableSource, 'Columns')), 'Name')
      },
      excelData () {
        var res = [];

        this._.each(this.dscmy.systemsDisplay, (system, i) => {
          var temp = {
            System_Name: system.System_Name,
            ITAM_ID: system.ITAM_ID,
          }

          var tables = this._.filter(this.dscmy.tableDisplay, (v) => v.Imm_Prec_System_ID == system.ID)
          if(tables.length > 0){
            this._.each(tables, (table, i) => {
              var tableLevel = _.cloneDeep(temp);
              tableLevel.Table_Name = table.Name;

              if(table.Columns.length > 0){
                this._.each(table.Columns, (column, j) => {
                  var colLevel = _.cloneDeep(tableLevel);
                  colLevel.Column_Name = column.Name;
                  colLevel.Alias_Name = column.Alias_Name;
                  colLevel.CDE = column.CDE;
                  
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
            this.getTableName(this.$route.params.system);
          }
      },
    },
    created() {
      this.secondtable = this.$route.params.system;
      
      this.getAllSystem();

      if(this.secondtable){
        this.getTableName(this.$route.params.system);
      }
    },
    methods: {
      ...mapActions('dscmy', {
          getAllSystem: 'getAllSystem',
          getTableName: 'getTableName',
      }),
      columnFilter (type, keyModel, val) {
        if(val == ""){
          if(type == "systems"){
            this.dscmy.systemsDisplay = this.dscmy.systemsSource;
          } else {
            this.dscmy.tableDisplay = this.dscmy.tableSource;
          }
          return
        }

        if(type == "systems"){
          this.dscmy.systemsDisplay = _.filter(this.dscmy.systemsSource, [keyModel.value, val]);
        } else {
          this.dscmy.tableDisplay = _.filter(this.dscmy.tableSource, [keyModel.value, val]);
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

        this.dscmy.systemsDisplay = this.dscmy.systemsSource;
        this.dscmy.tableDisplay = this.dscmy.tableSource;
        if(this.searchForm.systemName)
          this.dscmy.systemsDisplay = this._.filter(this.dscmy.systemsDisplay, (val) => val.System_Name.indexOf(this.searchForm.systemName) != -1);
        if(this.searchForm.itamID)
          this.dscmy.systemsDisplay = this._.filter(this.dscmy.systemsDisplay, (val) => val.ITAM_ID.toString().indexOf(this.searchForm.itamID) != -1);
        if(this.searchForm.tableName)
          this.dscmy.tableDisplay = this._.filter(this.dscmy.tableDisplay, (val) => val.Name.indexOf(this.searchForm.tableName) != -1);
        if(this.searchForm.columnName) {
          this._.each(this.dscmy.tableDisplay, (v, i) => {
            this.dscmy.tableDisplay[i].Columns = this._.filter(this.dscmy.tableDisplay[i].Columns, (w) => w.Name.indexOf(this.searchForm.columnName) != -1);
            this.dscmy.tableDisplay = this._.filter(this.dscmy.tableDisplay, (w) => w.Columns.length > 0)
          });
        }

        this.searchForm.show = false;
      },
      onReset (evt) {
        evt.preventDefault();
        /* Reset our form values */
        this.searchForm.systemName = '';
        this.searchForm.itamID = '';
        this.searchForm.country = '';
        this.searchForm.tableName = '';
        this.searchForm.colName = '';

        // /* Trick to reset/clear native browser form validation state */
        // this.searchForm.show = false;
        // this.$nextTick(() => { this.searchForm.show = true });
      },
      showDetails (id) {
        this.$router.push('/dsc/all/' + this.$route.params.system + '/' + id)
      }
    }
}
</script>
