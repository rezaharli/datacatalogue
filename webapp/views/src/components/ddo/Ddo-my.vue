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
                <input v-model="ddomy.searchMain" type="text" class="form-control" placeholder="Search" aria-label="Recipient's username" aria-describedby="basic-addon2">

                <div class="input-group-append">
                  <b-dropdown right id="ddown1" text="" ref="ddownSearch">
                    <b-container>
                      <b-form-row class="main-table-search-dropdown-form">
                        <b-col>
                          <b-form @submit="onSubmit" @reset="onReset">
                            <b-form-group horizontal :label-cols="4" breakpoint="md" label="Data Domain" label-for="DataDomain">
                              <b-form-input id="DataDomain" type="text" v-model="ddomy.searchDropdown.DataDomain"></b-form-input>
                            </b-form-group>

                            <b-form-group horizontal :label-cols="4" breakpoint="md" label="Sub Data Domain" label-for="SubDataDomain">
                              <b-form-input id="SubDataDomain" type="text" v-model="ddomy.searchDropdown.SubDataDomain"></b-form-input>
                            </b-form-group>

                            <b-form-group horizontal :label-cols="4" breakpoint="md" label="Sub Data Domain Owner" label-for="SubDataDomainOwner">
                              <b-form-input id="SubDataDomainOwner" type="text" v-model="ddomy.searchDropdown.SubDataDomainOwner"></b-form-input>
                            </b-form-group>

                            <b-form-group horizontal :label-cols="4" breakpoint="md" label="Business Term" label-for="BusinessTerm">
                              <b-form-select id="BusinessTerm" :options="businessTermMaster" v-model="ddomy.searchDropdown.BusinessTerm"></b-form-select>
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
                  :items="ddomy.left.display"
                  :pagination.sync="ddomy.left.pagination"
                  :total-items="ddomy.left.totalItems"
                  :loading="ddomy.left.loading"
                  item-key="ID"
                  class="elevation-1 fixed-header">

                <template slot="headerCell" slot-scope="props">
                  {{ props.header.text }} ({{ ddomy.left.source[0] ? ddomy.left.source[0]["COUNT_" + props.header.value.split(".").reverse()[0]] : 0 }})

                  <b-dropdown no-caret variant="link" class="header-filter-icon">
                    <template slot="button-content">
                      <!-- <i class="fa fa-filter text-muted"></i> -->
                      <v-icon small>filter_list</v-icon>
                    </template>

                    <b-dropdown-header>
                      <b-form-input type="text" placeholder="Filter" v-model="search['systems'][props.header.value]" @change="filterKeyup('systems', props.header)"></b-form-input>
                    </b-dropdown-header>

                    <b-dropdown-item v-for="item in distinctData(props.header.value, ddomy.left.source)" :key="item" @click="columnFilter('systems', props.header, item)">
                      {{ item }}
                    </b-dropdown-item>
                  </b-dropdown>
                </template>

                <v-progress-linear slot="progress" color="blue" indeterminate></v-progress-linear>

                <template slot="no-data">
                  <v-alert :value="ddomy.left.loading" type="info">
                    Please wait while data is loading
                  </v-alert>

                  <v-alert :value="!ddomy.left.loading" type="error">
                    Sorry, nothing to display here
                  </v-alert>
                </template>

                <template slot="items" slot-scope="props">
                  <td><b-link :to="{ path: addressPath + '/' + props.item.ID }"><tablecell :fulltext="props.item.DATA_DOMAIN" :isklik="false"></tablecell></b-link></td>
                  <td><tablecell :fulltext="props.item.SUB_DOMAINS" :isklik="true"></tablecell></td>
                </template>
              </v-data-table>
            </b-col>
            
            <b-col class="scrollableasdf">
              <v-data-table
                  :headers="secondTableHeaders"
                  :items="ddomy.right.display"
                  :pagination.sync="ddomy.right.pagination"
                  :total-items="ddomy.right.totalItems"
                  :loading="ddomy.right.loading"
                  v-if="secondtable"
                  item-key="ID"
                  class="elevation-1">
                <v-progress-linear slot="progress" color="blue" indeterminate></v-progress-linear>

                <template slot="no-data">
                  <v-alert :value="ddomy.right.loading" type="info">
                    Please wait while data is loading
                  </v-alert>

                  <v-alert :value="!ddomy.right.loading" type="error">
                    Sorry, nothing to display here
                  </v-alert>
                </template>

                <template slot="headerCell" slot-scope="props">
                  {{ props.header.text }} {{ props.header.displayCount ? "(" + (ddomy.right.source[0] ? ddomy.right.source[0]["COUNT_" + props.header.value.split(".").reverse()[0]] : 0) + ")" : "" }}

                  <b-dropdown no-caret variant="link" class="header-filter-icon">
                    <template slot="button-content">
                      <!-- <i class="fa fa-filter text-muted"></i> -->
                      <v-icon small>filter_list</v-icon>
                    </template>

                    <b-dropdown-header>
                      <b-form-input type="text" placeholder="Filter" v-model="search['tablename'][props.header.value]" @change="filterKeyup('tablename', props.header)"></b-form-input>
                    </b-dropdown-header>

                    <b-dropdown-item v-for="item in distinctData(props.header.value, ddomy.right.source)" :key="item" @click="columnFilter('tablename', props.header, item)">
                      {{ item }}
                    </b-dropdown-item>
                  </b-dropdown>
                </template>

                <template slot="items" slot-scope="props">
                  <tr>
                    <td><b-link @click="showDetails(props.item)"><tablecell :fulltext="props.item.BUSINESS_TERM" :isklik="false"></tablecell></b-link></td>
                    <td>{{ props.item.BUSINESS_TERM_DESCRIPTION }}</td>
                    <td>{{ props.item.CDE_YES_NO }}</td>
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
          { text: 'Data Domain', align: 'left', value: 'DATA_DOMAIN', sortable: false },
          { text: 'Sub Domains', align: 'left', value: 'SUB_DOMAINS', sortable: false },
        ],
        secondTableHeaders: [
          { text: 'Business Term', align: 'left', sortable: false, value: 'BUSINESS_TERM', width: "25%" },
          { text: 'Business Term Description', align: 'left', sortable: false, value: 'BUSINESS_TERM_DESCRIPTION', width: "25%" },
          { text: 'CDE (Yes/No)', align: 'left', sortable: false, value: 'CDE_YES_NO', width: "25%" },
        ],
      }
    },
    computed: {
      ...mapState({
        ddomy: state => state.ddomy.all
      }),
      addressPath (){
        var tmp = this.$route.path.split("/")
        return tmp.slice(0, 3).join("/")
      },
      businessTermMaster (){
        return this._.map(this._.flattenDeep(this._.map(this.ddomy.right.source, 'Columns')), 'Name')
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

        this._.each(this.ddomy.left.display, (system, i) => {
          var temp = {
            DATA_DOMAIN: system.DATA_DOMAIN,
            SUB_DOMAINS: system.SUB_DOMAINS,
          }
          
          var tables = this._.filter(this.ddomy.right.display, (v) => v.TSCID == system.ID)
          if(tables.length > 0){
            this._.each(tables, (table, i) => {
              var tableLevel = _.cloneDeep(temp);
              tableLevel.BUSINESS_TERM = table.BUSINESS_TERM;
              tableLevel.BUSINESS_TERM_DESCRIPTION = table.BUSINESS_TERM_DESCRIPTION;
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
      "ddomy.left.pagination": {
        handler () {
          this.getLeftTable();
        },
        deep: true
      },
      "ddomy.right.pagination": {
        handler () {
          if(this.secondtable){
            this.getRightTable(this.$route.params.system);
          }
        },
        deep: true
      },
      "ddomy.searchMain" (val, oldVal){
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
      ...mapActions('ddomy', {
          getLeftTable: 'getLeftTable',
          getRightTable: 'getRightTable',
      }),
      columnFilter (type, keyModel, val) {
        if(val == ""){
          if(type == "systems"){
            this.ddomy.left.display = this.ddomy.left.source;
          } else {
            this.ddomy.right.display = this.ddomy.right.source;
          }
          return
        }

        if(type == "systems"){
          if(keyModel.value.split(".")[1]){
            this.ddomy.left.display = _.filter(this.ddomy.left.source, (v) => {
              return v[keyModel.value.split(".")[0]].find((w) => {
                return w[keyModel.value.split(".")[1]] == val
              })
            });
          } else {
            this.ddomy.left.display = _.filter(this.ddomy.left.source, [keyModel.value, val]);
          }
        } else {
          if(keyModel.value.split(".")[1]){
            this.ddomy.right.display = _.filter(this.ddomy.right.source, (v) => {
              return v[keyModel.value.split(".")[0]].find((w) => {
                return w[keyModel.value.split(".")[1]] == val
              })
            });
          } else {
            this.ddomy.right.display = _.filter(this.ddomy.right.source, [keyModel.value, val]);
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
        // this.ddomy.searchDropdown.show = false;
      },
      onReset (evt) {
        evt.preventDefault();
        /* Reset our form values */
        this.searchForm.dataDomain = '';
        this.searchForm.subDataDomain = '';
        this.searchForm.subDataDomainOwner = '';
        this.searchForm.businessTerm = '';

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
