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
                <input v-model="dpomy.searchMain" type="text" class="form-control" placeholder="Search" aria-label="Recipient's username" aria-describedby="basic-addon2">
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
                  :items="dpomy.left.display"
                  :pagination.sync="dpomy.left.pagination"
                  :total-items="dpomy.left.totalItems"
                  :loading="dpomy.left.loading"
                  class="elevation-1 fixed-header">

                <template slot="headerCell" slot-scope="props">
                  {{ props.header.text }} ({{ distinctData(props.header.value, dpomy.left.source).length }})

                  <b-dropdown no-caret variant="link" class="header-filter-icon">
                    <template slot="button-content">
                      <!-- <i class="fa fa-filter text-muted"></i> -->
                      <v-icon small>filter_list</v-icon>
                    </template>

                    <b-dropdown-header>
                      <b-form-input type="text" placeholder="Filter" v-model="search['systems'][props.header.value]" @change="filterKeyup('systems', props.header)"></b-form-input>
                    </b-dropdown-header>

                    <b-dropdown-item v-for="item in distinctData(props.header.value, dpomy.left.source)" :key="item" @click="columnFilter('systems', props.header, item)">
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
                    <td><b-link :to="{ path: addressPath + '/' + props.item.ID }">{{ props.item.DOWNSTREAM_PROCESS }}</b-link></td>
                    <td>{{ props.item.PROCESS_OWNER }}</td>
                    <td>{{ props.item.BANK_ID }}</td>
                </template>
              </v-data-table>
            </b-col>
            
            <b-col class="scrollableasdf">
              <v-data-table
                  :headers="secondTableHeaders"
                  :items="dpomy.right.display"
                  :pagination.sync="dpomy.right.pagination"
                  :total-items="dpomy.right.totalItems"
                  :loading="dpomy.right.loading"
                  v-if="secondtable"
                  item-key="ID"
                  class="elevation-1">
                <v-progress-linear slot="progress" color="blue" indeterminate></v-progress-linear>
                  <template slot="no-data">
                    <v-alert :value="true" color="error" icon="warning">
                      Sorry, nothing to display here :(
                    </v-alert>
                  </template>

                <template slot="headerCell" slot-scope="props">
                  {{ props.header.text }} ({{ distinctData(props.header.value, dpomy.right.source).length }})

                  <b-dropdown no-caret variant="link" class="header-filter-icon">
                    <template slot="button-content">
                      <!-- <i class="fa fa-filter text-muted"></i> -->
                      <v-icon small>filter_list</v-icon>
                    </template>

                    <b-dropdown-header>
                      <b-form-input type="text" placeholder="Filter" v-model="search['tablename'][props.header.value]" @change="filterKeyup('tablename', props.header)"></b-form-input>
                    </b-dropdown-header>

                    <b-dropdown-item v-for="item in distinctData(props.header.value, dpomy.right.source)" :key="item" @click="columnFilter('tablename', props.header, item)">
                      {{ item }}
                    </b-dropdown-item>
                  </b-dropdown>
                </template>

                <template slot="items" slot-scope="props">
                  <tr @click="props.expanded = !props.expanded">
                    <td><b-link @click="showDetails(props.item.ID)">{{ props.item.CDE_NAME }}</b-link></td>
                    <td>{{ props.item.SEGMENT }}</td>
                    <td>{{ props.item.IMM_PREC_SYSTEM_ID }}</td>
                    <td>{{ props.item.ULTIMATE_SOURCE_SYSTEM_ID }}</td>
                    <td>{{ props.item.BUSINESS_DESCRIPTION }}</td>
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
          processName: '',
          processOwner: '',
          country: '',
          countryMaster: [],
          cdeName: ''
        },
        firstTableHeaders: [
          { text: 'Downstream Processes', align: 'left', value: 'DOWNSTREAM_PROCESS', sortable: false },
          { text: 'Process Owner', align: 'left', value: 'PROCESS_OWNER', sortable: false },
          { text: 'Bank ID', align: 'left', value: 'BANK_ID', sortable: false }
        ],
        secondTableHeaders: [
          { text: 'CDE Name', align: 'left', sortable: false, value: 'CDE_NAME', width: "25%" },
          { text: 'Segment', align: 'left', sortable: false, value: 'SEGMENT', width: "25%" },
          { text: 'Immediate Preceding System', align: 'left', sortable: false, value: 'IMM_PREC_SYSTEM_ID', width: "25%" },
          { text: 'Ultimate Source System', align: 'left', sortable: false, value: 'ULTIMATE_SOURCE_SYSTEM_ID', width: "25%" },
          { text: 'Business Description', align: 'left', sortable: false, value: 'BUSINESS_DESCRIPTION', width: "25%" },
          { text: 'CDE Rationale', align: 'left', sortable: false, value: 'CDE_RATIONALE', width: "25%" },
        ],
        excelFields: {
          'Downstream Processes': 'DOWNSTREAM_PROCESS',
          'Process Owner': 'PROCESS_OWNER',
          'Bank ID': 'BANK_ID',
          'CDE Name': 'CDE_NAME',
          'Segment': 'SEGMENT',
          'Immediate Preceding System': 'IMM_PREC_SYSTEM_ID',
          'Ultimate Source System': 'ULTIMATE_SOURCE_SYSTEM_ID',
          'Business Description': 'BUSINESS_DESCRIPTION',
          'CDE Rationale': 'CDE_RATIONALE',
        }
      }
    },
    computed: {
      ...mapState({
        dpomy: state => state.dpomy.all
      }),
      addressPath (){
        var tmp = this.$route.path.split("/")
        return tmp.slice(0, 3).join("/")
      },
      excelData () {
        var res = [];

        this._.each(this.dpomy.left.display, (system, i) => {
          var temp = {
            DOWNSTREAM_PROCESS: system.DOWNSTREAM_PROCESS,
            PROCESS_OWNER: system.PROCESS_OWNER,
            BANK_ID: system.BANK_ID
          }

          var tables = this._.filter(this.dpomy.right.display, (v) => v.TDPID == system.ID)
          if(tables.length > 0){
            this._.each(tables, (table, i) => {
              var tableLevel = _.cloneDeep(temp);
              tableLevel.CDE_NAME = table.CDE_NAME;
              tableLevel.SEGMENT = table.SEGMENT;
              tableLevel.IMM_PREC_SYSTEM_ID = table.IMM_PREC_SYSTEM_ID;
              tableLevel.ULTIMATE_SOURCE_SYSTEM_ID = table.ULTIMATE_SOURCE_SYSTEM_ID;
              tableLevel.BUSINESS_DESCRIPTION = table.BUSINESS_DESCRIPTION;
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
      "dpomy.left.pagination": {
        handler () {
          this.getLeftTable();
        },
        deep: true
      },
      "dpomy.right.pagination": {
        handler () {
          if(this.secondtable){
            this.getRightTable(this.$route.params.system);
          }
        },
        deep: true
      },
      "dpomy.searchMain" (val, oldVal){
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
      ...mapActions('dpomy', {
          getLeftTable: 'getLeftTable',
          getRightTable: 'getRightTable',
      }),
      columnFilter (type, keyModel, val) {
        if(val == ""){
          if(type == "systems"){
            this.dpomy.left.display = this.dpomy.left.source;
          } else {
            this.dpomy.right.display = this.dpomy.right.source;
          }
          return
        }

        if(type == "systems"){
          this.dpomy.left.display = _.filter(this.dpomy.left.source, [keyModel.value, val]);
        } else {
          this.dpomy.right.display = _.filter(this.dpomy.right.source, [keyModel.value, val]);
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

        this.dpomy.left.display = this.dpomy.left.source;
        this.dpomy.right.display = this.dpomy.right.source;

        if(this.searchForm.processName)
          this.dpomy.left.display = this._.filter(this.dpomy.left.display, (val) => val.DOWNSTREAM_PROCESS.indexOf(this.searchForm.processName) != -1);
        if(this.searchForm.processOwner)
          this.dpomy.left.display = this._.filter(this.dpomy.left.display, (val) => val.PROCESS_OWNER.toString().indexOf(this.searchForm.processOwner) != -1);

        if(this.searchForm.cdeName)
          this.dpomy.right.display = this._.filter(this.dpomy.right.display, (val) => val.CDE_NAME.toString().indexOf(this.searchForm.cdeName) != -1);

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
        this.$router.push(this.addressPath + "/" + this.$route.params.system + '/' + id)
      }
    }
}
</script>
