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
      <!-- Ddo details -->
      <router-view/>

      <!-- Main content -->
      <b-row>
        <b-col>
          <b-row>
            <b-col>
              <div class="input-group mb-3">
                <input v-model="ddoall.searchMain" type="text" class="form-control" placeholder="Search" aria-label="Recipient's username" aria-describedby="basic-addon2">

                <div class="input-group-append">
                  <b-dropdown right id="ddown1" text="" ref="ddownSearch">
                    <b-container>
                      <b-form-row class="main-table-search-dropdown-form">
                        <b-col>
                          <b-form @submit="onSubmit" @reset="onReset">
                            <b-form-group horizontal :label-cols="4" breakpoint="md" label="Data Domain" label-for="DataDomain">
                              <b-form-input id="DataDomain" type="text" v-model="ddoall.searchDropdown.DataDomain"></b-form-input>
                            </b-form-group>

                            <b-form-group horizontal :label-cols="4" breakpoint="md" label="Sub Data Domain" label-for="SubDataDomain">
                              <b-form-input id="SubDataDomain" type="text" v-model="ddoall.searchDropdown.SubDataDomain"></b-form-input>
                            </b-form-group>

                            <b-form-group horizontal :label-cols="4" breakpoint="md" label="Sub Data Domain Owner" label-for="SubDataDomainOwner">
                              <b-form-input id="SubDataDomainOwner" type="text" v-model="ddoall.searchDropdown.SubDataDomainOwner"></b-form-input>
                            </b-form-group>

                            <b-form-group horizontal :label-cols="4" breakpoint="md" label="Business Term" label-for="BusinessTerm">
                              <b-form-select id="BusinessTerm" :options="businessTermMaster" v-model="ddoall.searchDropdown.BusinessTerm"></b-form-select>
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
                  :items="ddoall.left.display"
                  :pagination.sync="ddoall.left.pagination"
                  :total-items="ddoall.left.totalItems"
                  :loading="ddoall.left.loading"
                  :expand="false"
                  item-key="ID"
                  class="elevation-1 fixed-header">

                <template slot="headerCell" slot-scope="props">
                  {{ props.header.text }} ({{ ddoall.left.source[0] ? ddoall.left.source[0]["COUNT_" + props.header.value.split(".").reverse()[0]] : 0 }})

                  <b-dropdown no-caret variant="link" class="header-filter-icon">
                    <template slot="button-content">
                      <!-- <i class="fa fa-filter text-muted"></i> -->
                      <v-icon small>filter_list</v-icon>
                    </template>

                    <b-dropdown-header>
                      <b-form-input type="text" placeholder="Filter" v-model="search['systems'][props.header.value]" @change="filterKeyup('systems', props.header)"></b-form-input>
                    </b-dropdown-header>

                    <b-dropdown-item v-for="item in distinctData(props.header.value, ddoall.left.source)" :key="item" @click="columnFilter('systems', props.header, item)">
                      {{ item }}
                    </b-dropdown-item>
                  </b-dropdown>
                </template>

                <v-progress-linear slot="progress" color="blue" indeterminate></v-progress-linear>

                <template slot="no-data">
                  <v-alert :value="ddoall.left.loading" type="info">
                    Please wait while data is loading
                  </v-alert>

                  <v-alert :value="!ddoall.left.loading" type="error">
                    Sorry, nothing to display here
                  </v-alert>
                </template>

                <template slot="items" slot-scope="props">
                  <td><b-link :to="{ path: addressPath + '/' + props.item.ID }"><tablecell :fulltext="props.item.SUB_DOMAINS" :isklik="false"></tablecell></b-link></td>
                  <td><b-link @click="props.expanded = !props.expanded"><tablecell :fulltext="(_.uniq(_.map(props.item.Values, 'DATA_DOMAIN').filter(Boolean)).join(', '))" :isklik="false"></tablecell></b-link></td>
                  <td><tablecell :fulltext="(_.uniq(_.map(props.item.Values, 'SUB_DOMAIN_OWNER').filter(Boolean)).join('; '))" :isklik="true"></tablecell></td>
                  <td><tablecell :fulltext="(_.uniq(_.map(props.item.Values, 'BANK_ID').filter(Boolean)).join('; '))" :isklik="true"></tablecell></td>
                </template>

                <template slot="expand" slot-scope="props">
                  <v-data-table
                    :headers="firstTableHeaders"
                    :items="props.item.Values"
                    class="elevation-1"
                    hide-actions
                    hide-headers
                  >
                    <template slot="items" slot-scope="props">
                      <td style="width: 25%">&nbsp;</td>
                      <td style="width: 25%">&nbsp;</td>
                      <td style="width: 25%">{{ props.item.SUB_DOMAIN_OWNER }}</td>
                      <td style="width: 25%">{{ props.item.BANK_ID }}</td>
                    </template>
                  </v-data-table>
                </template>
              </v-data-table>
            </b-col>
            
            <b-col class="scrollableasdf">
              <v-data-table
                  :headers="secondTableHeaders"
                  :items="ddoall.right.display"
                  :pagination.sync="ddoall.right.pagination"
                  :total-items="ddoall.right.totalItems"
                  :loading="ddoall.right.loading"
                  v-if="secondtable"
                  item-key="ID"
                  class="elevation-1">
                <v-progress-linear slot="progress" color="blue" indeterminate></v-progress-linear>

                <template slot="no-data">
                  <v-alert :value="ddoall.right.loading" type="info">
                    Please wait while data is loading
                  </v-alert>

                  <v-alert :value="!ddoall.right.loading" type="error">
                    Sorry, nothing to display here
                  </v-alert>
                </template>

                <template slot="headerCell" slot-scope="props">
                  {{ props.header.text }} {{ props.header.displayCount ? "(" + (ddoall.right.source[0] ? ddoall.right.source[0]["COUNT_" + props.header.value.split(".").reverse()[0]] : 0) + ")" : "" }}

                  <b-dropdown no-caret variant="link" class="header-filter-icon">
                    <template slot="button-content">
                      <!-- <i class="fa fa-filter text-muted"></i> -->
                      <v-icon small>filter_list</v-icon>
                    </template>

                    <b-dropdown-header>
                      <b-form-input type="text" placeholder="Filter" v-model="search['tablename'][props.header.value]" @change="filterKeyup('tablename', props.header)"></b-form-input>
                    </b-dropdown-header>

                    <b-dropdown-item v-for="item in distinctData(props.header.value, ddoall.right.source)" :key="item" @click="columnFilter('tablename', props.header, item)">
                      {{ item }}
                    </b-dropdown-item>
                  </b-dropdown>
                </template>

                <template slot="items" slot-scope="props">
                  <tr>
                    <td><b-link @click="showDetails(props.item)"><tablecell :fulltext="props.item.BUSINESS_TERM" :isklik="false"></tablecell></b-link></td>
                    <td><tablecell :fulltext="props.item.BT_DESCRIPTION" :isklik="true"></tablecell></td>
                    <td><tablecell :fulltext="props.item.CDE_YES_NO" :isklik="true"></tablecell></td>
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
          { text: 'Sub Domains', align: 'left', value: 'SUB_DOMAINS', sortable: false },
          { text: 'Data Domain', align: 'left', value: 'DATA_DOMAIN', sortable: false },
          { text: 'Sub Domain Owner', align: 'left', value: 'SUB_DOMAIN_OWNER', sortable: false },
          { text: 'Bank ID', align: 'left', value: 'BANK_ID', sortable: false },
        ],
        secondTableHeaders: [
          { text: 'Business Term', align: 'left', sortable: false, value: 'BUSINESS_TERM', displayCount: true, width: "25%" },
          { text: 'Business Term Description', align: 'left', sortable: false, value: 'BT_DESCRIPTION', displayCount: true, width: "25%" },
          { text: 'CDE (Yes/No)', align: 'left', sortable: false, value: 'CDE_YES_NO', displayCount: true, width: "25%" },
        ],
      }
    },
    computed: {
      ...mapState({
        ddoall: state => state.ddoall.all
      }),
      addressPath (){
        var tmp = this.$route.path.split("/")
        return tmp.slice(0, 3).join("/")
      },
      businessTermMaster (){
        return this._.map(this.ddoall.right.source, 'BUSINESS_TERM')
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

        this._.each(this.ddoall.left.display, (system, i) => {
          var temp = {
            SUB_DOMAINS: system.SUB_DOMAINS,
            DATA_DOMAIN: system.DATA_DOMAIN,
            SUB_DOMAIN_OWNER: system.SUB_DOMAIN_OWNER,
            BANK_ID: system.BANK_ID,
          }
          
          var tables = this._.filter(this.ddoall.right.display, (v) => v.TSCID == system.ID)
          if(this.secondtable && tables.length > 0){
            this._.each(tables, (table, i) => {
              var tableLevel = _.cloneDeep(temp);
              tableLevel.BUSINESS_TERM = table.BUSINESS_TERM;
              tableLevel.BT_DESCRIPTION = table.BT_DESCRIPTION;
              tableLevel.CDE_YES_NO = table.CDE_YES_NO;

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
      "ddoall.left.pagination": {
        handler () {
          this.getLeftTable();
        },
        deep: true
      },
      "ddoall.right.pagination": {
        handler () {
          if(this.secondtable){
            this.getRightTable(this.$route.params.system);
          }
        },
        deep: true
      },
      "ddoall.searchMain" (val, oldVal){
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
      ...mapActions('ddoall', {
          getLeftTable: 'getLeftTable',
          getRightTable: 'getRightTable',
      }),
      columnFilter (type, keyModel, val) {
        if(val == ""){
          if(type == "systems"){
            this.ddoall.left.display = this.ddoall.left.source;
          } else {
            this.ddoall.right.display = this.ddoall.right.source;
          }
          
          return
        }

        if(type == "systems"){
          if(keyModel.value.split(".")[1]){
            this.ddoall.left.display = _.cloneDeep(this.ddoall.left.source);
            this.ddoall.left.display = this.ddoall.left.display.filter(
              v => {
                var key = keyModel.value.split(".")[0];
                
                v[key] = v[key].filter(w => w[keyModel.value.split(".")[1]].toString().toUpperCase().includes(val.toString().toUpperCase()))

                return v[key].length > 0;
              }
            );
          } else {
            this.ddoall.left.display = _.cloneDeep(this.ddoall.left.source);
            this.ddoall.left.display = _.filter(this.ddoall.left.display, (v) => {
              return v[keyModel.value].toString().toUpperCase().includes(val.toString().toUpperCase());
            });
          }
        } else {
          if(keyModel.value.split(".")[1]){
            this.ddoall.right.display = _.cloneDeep(this.ddoall.right.source);
            this.ddoall.right.display = this.ddoall.right.display.filter(
              v => {
                var key = keyModel.value.split(".")[0];
                
                v[key] = v[key].filter(w => w[keyModel.value.split(".")[1]].toString().toUpperCase().includes(val.toString().toUpperCase()))

                return v[key].length > 0;
              }
            );
          } else {
            this.ddoall.right.display = _.cloneDeep(this.ddoall.right.source);
            this.ddoall.right.display = _.filter(this.ddoall.right.display, (v) => {
              return v[keyModel.value].toString().toUpperCase().includes(val.toString().toUpperCase());
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
      onSubmit (evt) {
        if(evt) evt.preventDefault();

        this.getLeftTable();

        if(this.secondtable){
          this.getRightTable(this.$route.params.system);
        }

        this.$refs.ddownSearch.hide(true);
        // this.ddoall.searchDropdown.show = false;
      },
      onReset (evt) {
        evt.preventDefault();
        /* Reset our form values */
        this.ddoall.searchDropdown.SubDataDomain = '';
        this.ddoall.searchDropdown.DataDomain = '';
        this.ddoall.searchDropdown.SubDataDomainOwner = '';
        this.ddoall.searchDropdown.BusinessTerm = '';

        this.onSubmit();

        // /* Trick to reset/clear native browser form validation state */
        // this.searchForm.show = false;
        // this.$nextTick(() => { this.searchForm.show = true });
      },
      showDetails (param) {
        this.$router.push(this.addressPath + "/" + param.TSCID + '/' + param.ID)
      }
    }
}
</script>
