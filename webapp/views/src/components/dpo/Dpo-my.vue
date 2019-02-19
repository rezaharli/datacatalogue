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
      <!-- Dpo details -->
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
                            <b-form-group horizontal :label-cols="4" breakpoint="md" label="Process Name" label-for="processName">
                              <b-form-input id="processName" type="text" v-model="searchForm.processName"></b-form-input>
                            </b-form-group>

                            <b-form-group horizontal :label-cols="4" breakpoint="md" label="Process Owner" label-for="processOwner">
                              <b-form-input id="processOwner" type="text" v-model="searchForm.processOwner"></b-form-input>
                            </b-form-group>

                            <b-form-group horizontal :label-cols="4" breakpoint="md" label="Country" label-for="country">
                              <b-form-select id="country" :options="searchForm.countryMaster" v-model="searchForm.country"></b-form-select>
                            </b-form-group>

                            <b-form-group horizontal :label-cols="4" breakpoint="md" label="CDE Name" label-for="cdeName">
                              <b-form-input id="cdeName" type="text" v-model="searchForm.cdeName"></b-form-input>
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
                  :items="dpomy.systemsDisplay"
                  :loading="dpomy.systemsLoading"
                  :search="searchMain"
                  class="elevation-1 fixed-header">

                <template slot="headerCell" slot-scope="props">
                  {{ props.header.text }} ({{ distinctData(props.header.value, dpomy.systemsSource).length }})

                  <b-dropdown no-caret variant="link" class="header-filter-icon">
                    <template slot="button-content">
                      <i class="fa fa-filter text-muted"></i>
                    </template>

                    <b-dropdown-header>
                      <b-form-input type="text" placeholder="Filter" v-model="search['systems'][props.header.value]" @change="filterKeyup('systems', props.header)"></b-form-input>
                    </b-dropdown-header>

                    <b-dropdown-item v-for="item in distinctData(props.header.value, dpomy.systemsSource)" :key="item" @click="columnFilter('systems', props.header, item)">
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
                    <td><b-link :to="{ path:'/dpo/my/' + props.item.ID }">{{ props.item.Name }}</b-link></td>
                    <td>{{ props.item.Owner_ID }}</td>
                    <td>{{ props.item.fat }}</td>
                </template>
              </v-data-table>
            </b-col>
            
            <b-col class="scrollableasdf">
              <v-data-table
                :headers="secondTableHeaders"
                :items="dpomy.tableDisplay"
                :loading="dpomy.tableLoading"
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
                  {{ props.header.text }} ({{ distinctData(props.header.value, dpomy.tableSource).length }})

                  <b-dropdown no-caret variant="link" class="header-filter-icon">
                    <template slot="button-content">
                      <i class="fa fa-filter text-muted"></i>
                    </template>

                    <b-dropdown-header>
                      <b-form-input type="text" placeholder="Filter" v-model="search['tablename'][props.header.value]" @change="filterKeyup('tablename', props.header)"></b-form-input>
                    </b-dropdown-header>

                    <b-dropdown-item v-for="item in distinctData(props.header.value, dpomy.tableSource)" :key="item" @click="columnFilter('tablename', props.header, item)">
                      {{ item }}
                    </b-dropdown-item>
                  </b-dropdown>
                </template>

                <template slot="items" slot-scope="props">
                  <tr @click="props.expanded = !props.expanded">
                    <td><b-link @click="showDetails(props.item.ID)">{{ props.item.Business_Term_ID }}</b-link></td>
                    <td>{{ props.item.Segment_ID }}</td>
                    <td>{{ props.item.Imm_Prec_System_ID }}</td>
                    <td>{{ props.item.Ultimate_Source_System_ID }}</td>
                    <td>{{ props.item.BusinessTerm.Description }}</td>
                    <td>{{ props.item.CDE_Rationale }}</td>
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
        searchMain: '',
        searchForm: {
          processName: '',
          processOwner: '',
          country: '',
          countryMaster: [],
          cdeName: ''
        },
        firstTableHeaders: [
          { text: 'Downstream Processes', align: 'left', value: 'Name', sortable: false },
          { text: 'Process Owner', align: 'left', value: 'Owner_ID', sortable: false },
          { text: 'Bank ID', align: 'left', value: 'carbs', sortable: false }
        ],
        secondTableHeaders: [
          { text: 'CDE Name', align: 'left', sortable: false, value: 'Business_Term_ID', width: "25%" },
          { text: 'Segment', align: 'left', sortable: false, value: 'Segment_ID', width: "25%" },
          { text: 'Immediate Preceding System', align: 'left', sortable: false, value: 'Imm_Prec_System_ID', width: "25%" },
          { text: 'Ultimate Source System', align: 'left', sortable: false, value: 'Ultimate_Source_System_ID', width: "25%" },
          { text: 'Business Description', align: 'left', sortable: false, value: 'BusinessTerm.Description', width: "25%" },
          { text: 'CDE Rationale', align: 'left', sortable: false, value: 'CDE_Rationale', width: "25%" },
        ],
        excelFields: {
          'Downstream Processes': 'Name',
          'Process Owner': 'Owner_ID',
          'Bank ID': 'asdf',
          'CDE Name': 'Business_Term_ID',
          'Segment': 'Segment_ID',
          'Immediate Preceding System': 'Imm_Prec_System_ID',
          'Ultimate Source System': 'Ultimate_Source_System_ID',
          'Business Description': 'BusinessTermDescription',
          'CDE Rationale': 'CDE_Rationale',
        }
      }
    },
    computed: {
      ...mapState({
        dpomy: state => state.dpomy.all
      }),
      excelData () {
        var res = [];

        this._.each(this.dpomy.systemsDisplay, (system, i) => {
          var temp = {
            Name: system.Name,
            Owner_ID: system.Owner_ID,
          }

          var tables = this._.filter(this.dpomy.tableDisplay, (v) => v.Process_ID == system.ID)
          if(tables.length > 0){
            this._.each(tables, (table, i) => {
              var tableLevel = _.cloneDeep(temp);
              tableLevel.Business_Term_ID = table.Business_Term_ID;
              tableLevel.Segment_ID = table.Segment_ID;
              tableLevel.Imm_Prec_System_ID = table.Imm_Prec_System_ID;
              tableLevel.Ultimate_Source_System_ID = table.Ultimate_Source_System_ID;
              tableLevel.BusinessTermDescription = table.BusinessTerm.Description;
              tableLevel.CDE_Rationale = table.CDE_Rationale;

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
      ...mapActions('dpomy', {
          getAllSystem: 'getAllSystem',
          getTableName: 'getTableName',
      }),
      columnFilter (type, keyModel, val) {
        if(val == ""){
          if(type == "systems"){
            this.dpomy.systemsDisplay = this.dpomy.systemsSource;
          } else {
            this.dpomy.tableDisplay = this.dpomy.tableSource;
          }
          return
        }

        if(type == "systems"){
          this.dpomy.systemsDisplay = _.filter(this.dpomy.systemsSource, [keyModel.value, val]);
        } else {
          this.dpomy.tableDisplay = _.filter(this.dpomy.tableSource, [keyModel.value, val]);
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

        this.dpomy.systemsDisplay = this.dpomy.systemsSource;
        this.dpomy.tableDisplay = this.dpomy.tableSource;

        if(this.searchForm.processName)
          this.dpomy.systemsDisplay = this._.filter(this.dpomy.systemsDisplay, (val) => val.Name.indexOf(this.searchForm.processName) != -1);
        if(this.searchForm.processOwner)
          this.dpomy.systemsDisplay = this._.filter(this.dpomy.systemsDisplay, (val) => val.Owner_ID.toString().indexOf(this.searchForm.processOwner) != -1);
        if(this.searchForm.cdeName)
          this.dpomy.tableDisplay = this._.filter(this.dpomy.tableDisplay, (val) => val.Business_Term_ID.toString().indexOf(this.searchForm.cdeName) != -1);

        this.searchForm.show = false;
      },
      onReset (evt) {
        evt.preventDefault();
        /* Reset our form values */
        this.searchForm.processName = '';
        this.searchForm.processOwner = '';
        this.searchForm.country = '';
        this.searchForm.cdeName = '';

        // /* Trick to reset/clear native browser form validation state */
        // this.searchForm.show = false;
        // this.$nextTick(() => { this.searchForm.show = true });
      },
      showDetails (id) {
        this.$router.push('/dpo/my/' + this.$route.params.system + '/' + id)
      }
    }
}
</script>
