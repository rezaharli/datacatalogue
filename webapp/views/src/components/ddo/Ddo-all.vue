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
                  <b-dropdown right id="ddown1" text="">
                    <b-container>
                      <b-form-row class="main-table-search-dropdown-form">
                        <b-col>
                          <b-form @submit="onSubmit" @reset="onReset">
                            <b-form-group horizontal :label-cols="4" breakpoint="md" label="Data Domain" label-for="dataDomain">
                              <b-form-input id="dataDomain" type="text" v-model="searchForm.dataDomain"></b-form-input>
                            </b-form-group>

                            <b-form-group horizontal :label-cols="4" breakpoint="md" label="Sub Data Domain" label-for="subDataDomain">
                              <b-form-input id="subDataDomain" type="text" v-model="searchForm.subDataDomain"></b-form-input>
                            </b-form-group>

                            <b-form-group horizontal :label-cols="4" breakpoint="md" label="Sub Data Domain Owner" label-for="subDataDomainOwner">
                              <b-form-input id="subDataDomainOwner" type="text" v-model="searchForm.subDataDomainOwner"></b-form-input>
                            </b-form-group>

                            <b-form-group horizontal :label-cols="4" breakpoint="md" label="Business Term" label-for="businessTerm">
                              <b-form-select id="businessTerm" :options="tablenameMaster" v-model="searchForm.businessTerm"></b-form-select>
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
                  class="elevation-1 fixed-header">

                <template slot="headerCell" slot-scope="props">
                  {{ props.header.text }} ({{ distinctData(props.header.value, ddomy.left.source).length }})

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
                  <v-alert :value="true" color="error" icon="warning">
                    Sorry, nothing to display here :(
                  </v-alert>
                </template>

                <template slot="items" slot-scope="props">
                    <td><b-link :to="{ path: addressPath + '/' + props.item.ID }">{{ props.item.DOMAIN }}</b-link></td>
                    <td>{{ props.item.SUB_DOMAIN }}</td>
                </template>
              </v-data-table>
            </b-col>
            
            <b-col>
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
                  <v-alert :value="true" color="error" icon="warning">
                    Sorry, nothing to display here :(
                  </v-alert>
                </template>

                <template slot="headerCell" slot-scope="props">
                  {{ props.header.text }} ({{ distinctData(props.header.value, ddomy.right.source).length }})

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
                    <td><b-link @click="showDetails(props.item.ID)">{{ props.item.BT_NAME }}</b-link></td>
                    <td>{{ props.item.DESCRIPTION }}</td>
                    <td>{{ props.item.CDE }}</td>
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
          dataDomain: '',
          subDataDomain: '',
          subDataDomainOwner: '',
          businessTerm: ''
        },
        firstTableHeaders: [
          { text: 'Data Domain', align: 'left', value: 'DOMAIN', sortable: false },
          { text: 'Sub Domains', align: 'left', value: 'SUB_DOMAIN', sortable: false },
        ],
        secondTableHeaders: [
          { text: 'Business Term', align: 'left', sortable: false, value: 'BT_NAME', width: "25%" },
          { text: 'Business Term Description', align: 'left', sortable: false, value: 'DESCRIPTION', width: "25%" },
          { text: 'CDE (Yes/No)', align: 'left', sortable: false, value: 'CDE', width: "25%" },
        ],
        excelFields: {
          'Data Domain': 'DOMAIN',
          'Sub Domains': 'SUB_DOMAIN',
          'Business Term': 'BT_NAME',
          'Business Term Description': 'DESCRIPTION',
          'CDE (Yes/No)': 'CDE'
        }
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
      tablenameMaster (){
        return this._.map(this.ddomy.right.source, 'BT_NAME')
      },
      columnNameMaster (){
        return this._.map(this._.flattenDeep(this._.map(this.ddomy.right.source, 'Columns')), 'Name')
      },
      excelData () {
        var res = [];

        this._.each(this.ddomy.left.display, (system, i) => {
          var temp = {
            DOMAIN: system.DOMAIN,
            SUB_DOMAIN: system.SUB_DOMAIN,
          }
          
          var tables = this._.filter(this.ddomy.right.display, (v) => v.TSCID == system.ID)
          if(tables.length > 0){
            this._.each(tables, (table, i) => {
              var tableLevel = _.cloneDeep(temp);
              tableLevel.BT_NAME = table.BT_NAME;
              tableLevel.DESCRIPTION = table.DESCRIPTION;
              tableLevel.CDE = table.CDE;

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
          this.ddomy.left.display = _.filter(this.ddomy.left.source, [keyModel.value, val]);
        } else {
          this.ddomy.right.display = _.filter(this.ddomy.right.source, [keyModel.value, val]);
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

        this.ddomy.left.display = this.ddomy.left.source;
        this.ddomy.right.display = this.ddomy.right.source;
        if(this.searchForm.dataDomain)
          this.ddomy.left.display = this._.filter(this.ddomy.left.display, (val) => val.DOMAIN.toString().indexOf(this.searchForm.dataDomain) != -1);
        if(this.searchForm.subDataDomain)
          this.ddomy.left.display = this._.filter(this.ddomy.left.display, (val) => val.SUB_DOMAIN.toString().indexOf(this.searchForm.subDataDomain) != -1);

        if(this.searchForm.businessTerm)
          this.ddomy.right.display = this._.filter(this.ddomy.right.display, (val) => val.BT_NAME.indexOf(this.searchForm.businessTerm) != -1);
        // if(this.searchForm.businessTerm) {
        //   this._.each(this.ddomy.right.display, (v, i) => {
        //     this.ddomy.right.display[i].Columns = this._.filter(this.ddomy.right.display[i].Columns, (w) => w.Name.indexOf(this.searchForm.businessTerm) != -1);
        //     this.ddomy.right.display = this._.filter(this.ddomy.right.display, (w) => w.Columns.length > 0)
        //   });
        // }

        this.searchForm.show = false;
      },
      onReset (evt) {
        evt.preventDefault();
        /* Reset our form values */
        this.searchForm.dataDomain = '';
        this.searchForm.subDataDomain = '';
        this.searchForm.subDataDomainOwner = '';
        this.searchForm.businessTerm = '';

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
