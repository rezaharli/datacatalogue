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
                <input v-model="dscmy.searchMain" type="text" class="form-control" placeholder="Search" aria-label="Recipient's username" aria-describedby="basic-addon2">
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
                  :items="dscmy.left.display"
                  :pagination.sync="dscmy.left.pagination"
                  :total-items="dscmy.left.totalItems"
                  :loading="dscmy.left.loading"
                  class="elevation-1 fixed-header">

                <template slot="headerCell" slot-scope="props">
                  {{ props.header.text }} ({{ distinctData(props.header.value, dscmy.left.source).length }})

                  <b-dropdown no-caret variant="link" class="header-filter-icon">
                    <template slot="button-content">
                      <!-- <i class="fa fa-filter text-muted"></i> -->
                      <v-icon small>filter_list</v-icon>
                    </template>

                    <b-dropdown-header>
                      <b-form-input type="text" placeholder="Filter" v-model="search['systems'][props.header.value]" @change="filterKeyup('systems', props.header)"></b-form-input>
                    </b-dropdown-header>

                    <b-dropdown-item v-for="item in distinctData(props.header.value, dscmy.left.source)" :key="item" @click="columnFilter('systems', props.header, item)">
                      {{ item }}
                    </b-dropdown-item>
                  </b-dropdown>
                </template>

                <v-progress-linear slot="progress" color="blue" indeterminate></v-progress-linear>

                <template slot="no-data">
                  <v-alert :value="dscmy.left.loading" type="info">
                    Please wait while data is loading
                  </v-alert>

                  <v-alert :value="!dscmy.left.loading" type="error">
                    Sorry, nothing to display here
                  </v-alert>
                </template>

                <template slot="items" slot-scope="props">
                    <td><b-link :to="{ path: addressPath + '/' + props.item.ID }"><tablecell :fulltext="props.item.SYSTEM_NAME" :isklik="false"></tablecell></b-link></td>
                    <td>{{ props.item.ITAM_ID }}</td>
                    <td><tablecell :fulltext="props.item.FIRST_NAME" :isklik="true"></tablecell></td>
                    <td>{{ props.item.BANK_ID }}</td>
                </template>
              </v-data-table>
            </b-col>
            
            <b-col>
              <v-data-table
                  :headers="secondTableHeaders"
                  :items="dscmy.right.display"
                  :pagination.sync="dscmy.right.pagination"
                  :total-items="dscmy.right.totalItems"
                  :loading="dscmy.right.loading"
                  :expand="false"
                  v-if="secondtable"
                  item-key="ID"
                  class="elevation-1">
                <v-progress-linear slot="progress" color="blue" indeterminate></v-progress-linear>

                <template slot="no-data">
                  <v-alert :value="dscmy.right.loading" type="info">
                    Please wait while data is loading
                  </v-alert>

                  <v-alert :value="!dscmy.right.loading" type="error">
                    Sorry, nothing to display here
                  </v-alert>
                </template>

                <template slot="headerCell" slot-scope="props">
                  {{ props.header.text }} ({{ distinctData(props.header.value, dscmy.right.source).length }})

                  <b-dropdown no-caret variant="link" class="header-filter-icon">
                    <template slot="button-content">
                      <!-- <i class="fa fa-filter text-muted"></i> -->
                      <v-icon small>filter_list</v-icon>
                    </template>

                    <b-dropdown-header>
                      <b-form-input type="text" placeholder="Filter" v-model="search['tablename'][props.header.value]" @change="filterKeyup('tablename', props.header)"></b-form-input>
                    </b-dropdown-header>

                    <b-dropdown-item v-for="item in distinctData(props.header.value, dscmy.right.source)" :key="item" @click="columnFilter('tablename', props.header, item)">
                      {{ item }}
                    </b-dropdown-item>
                  </b-dropdown>
                </template>

                <template slot="items" slot-scope="props">
                  <tr>
                    <td><b-link @click="props.expanded = !props.expanded"><tablecell :fulltext="props.item.TABLE_NAME" :isklik="false"></tablecell></b-link></td>
                    <td><tablecell :fulltext="(_.map(props.item.Columns, 'COLUMN_NAME').filter(Boolean).join(', '))" :isklik="true"></tablecell></td>
                    <td><tablecell :fulltext="(_.map(props.item.Columns, 'ALIAS_NAME').filter(Boolean).join(', '))" :isklik="true"></tablecell></td>
                    <td><tablecell :fulltext="(_.map(props.item.Columns, 'CDE').filter(Boolean).join(', '))" :isklik="true"></tablecell></td>
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
                      <td style="width: 25%"><b-link @click="showDetails(props.item.TSID, props.item.ID, props.item.COLID)">{{ props.item.COLUMN_NAME }}</b-link></td>
                      <td style="width: 25%">{{ props.item.ALIAS_NAME }}</td>
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
        searchForm: {
          systemName: '',
          itamID: '',
          country: '',
          countryMaster: ['a', 'c', 'd'],
          tableName: '',
          columnName: '',
        },
        firstTableHeaders: [
          { text: 'System Name', align: 'left', value: 'SYSTEM_NAME', sortable: false },
          { text: 'ITAM ID', align: 'left', value: 'ITAM_ID', sortable: false },
          { text: 'Dataset Custodian', align: 'left', value: 'FIRST_NAME', sortable: false },
          { text: 'Bank ID', align: 'left', value: 'BANK_ID', sortable: false }
        ],
        secondTableHeaders: [
          { text: 'Table Name', align: 'left', sortable: false, value: 'TABLE_NAME', width: "25%" },
          { text: 'Column Name', align: 'left', sortable: false, value: 'Columns.COLUMN_NAME', width: "25%" },
          { text: 'Business Alias Name', align: 'left', sortable: false, value: 'Columns.ALIAS_NAME', width: "25%" },
          { text: 'CDE (Yes/No)', align: 'left', sortable: false, value: 'Columns.CDE', width: "25%" }
        ],
      }
    },
    computed: {
      ...mapState({
        dscmy: state => state.dscmy.all
      }),
      addressPath (){
        var tmp = this.$route.path.split("/")
        return tmp.slice(0, 3).join("/")
      },
      tablenameMaster (){
        return this._.map(this.dscmy.right.source, 'TABLE_NAME')
      },
      columnNameMaster (){
        return this._.map(this._.flattenDeep(this._.map(this.dscmy.right.source, 'Columns')), 'COLUMN_NAME')
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

        this._.each(this.dscmy.left.display, (system, i) => {
          var temp = {
            SYSTEM_NAME: system.SYSTEM_NAME,
            ITAM_ID: system.ITAM_ID,
            FIRST_NAME: system.FIRST_NAME,
            BANK_ID: system.BANK_ID,
          }

          var tables = this._.filter(this.dscmy.right.display, (v) => v.TSID == system.ID)
          if(tables.length > 0){
            this._.each(tables, (table, i) => {
              var tableLevel = _.cloneDeep(temp);
              tableLevel.TABLE_NAME = table.TABLE_NAME;

              if(table.Columns.length > 0){
                this._.each(table.Columns, (column, j) => {
                  var colLevel = _.cloneDeep(tableLevel);
                  colLevel.COLUMN_NAME = column.COLUMN_NAME;
                  colLevel.ALIAS_NAME = column.ALIAS_NAME;
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
          this.getRightTable(this.$route.params.system);
        }
      },
      "dscmy.left.pagination": {
        handler () {
          this.getLeftTable();
        },
        deep: true
      },
      "dscmy.right.pagination": {
        handler () {
          if(this.secondtable){
            this.getRightTable(this.$route.params.system);
          }
        },
        deep: true
      },
      "dscmy.searchMain" (val, oldVal){
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
      ...mapActions('dscmy', {
          getLeftTable: 'getLeftTable',
          getRightTable: 'getRightTable',
      }),
      columnFilter (type, keyModel, val) {
        if(val == ""){
          if(type == "systems"){
            this.dscmy.left.display = this.dscmy.left.source;
          } else {
            this.dscmy.right.display = this.dscmy.right.source;
          }
          return
        }

        if(type == "systems"){
          this.dscmy.left.display = _.filter(this.dscmy.left.source, [keyModel.value, val]);
        } else {
          this.dscmy.right.display = _.filter(this.dscmy.right.source, [keyModel.value, val]);
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

        this.dscmy.left.display = this.dscmy.left.source;
        this.dscmy.right.display = this.dscmy.right.source;
        if(this.searchForm.systemName)
          this.dscmy.left.display = this._.filter(this.dscmy.left.display, (val) => val.SYSTEM_NAME.toString().toUpperCase().indexOf(this.searchForm.systemName.toString().toUpperCase()) != -1);
        if(this.searchForm.itamID)
          this.dscmy.left.display = this._.filter(this.dscmy.left.display, (val) => val.ITAM_ID.toString().indexOf(this.searchForm.itamID.toString().toUpperCase()) != -1);
        if(this.searchForm.tableName)
          this.dscmy.right.display = this._.filter(this.dscmy.right.display, (val) => val.TABLE_NAME.toString().indexOf(this.searchForm.tableName.toString().toUpperCase()) != -1);
        if(this.searchForm.columnName) {
          this._.each(this.dscmy.right.display, (v, i) => {
            this.dscmy.right.display[i].Columns = this._.filter(this.dscmy.right.display[i].Columns, (w) => w.COLUMN_NAME.toString().toUpperCase().indexOf(this.searchForm.columnName.toString().toUpperCase()) != -1);
            this.dscmy.right.display = this._.filter(this.dscmy.right.display, (w) => w.Columns.length > 0)
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
        this.searchForm.columnName = '';

        this.onSubmit();

        // /* Trick to reset/clear native browser form validation state */
        // this.searchForm.show = false;
        // this.$nextTick(() => { this.searchForm.show = true });
      },
      showDetails (tsid, id, colid) {
        this.$router.push(this.addressPath + "/" + tsid + '/' + id + '/' + colid)
      }
    }
}
</script>
