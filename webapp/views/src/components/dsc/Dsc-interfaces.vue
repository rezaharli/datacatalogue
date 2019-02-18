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
                <input type="text" class="form-control" placeholder="Search" aria-label="Recipient's username" aria-describedby="basic-addon2">
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
                              <b-form-select id="tableName" :options="searchForm.countryMaster" v-model="searchForm.tableName"></b-form-select>
                            </b-form-group>

                            <b-form-group horizontal :label-cols="4" breakpoint="md" label="Column Name" label-for="columnName">
                              <b-form-select id="columnName" :options="searchForm.countryMaster" v-model="searchForm.columnName"></b-form-select>
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
            <b-col></b-col>
          </b-row>

          <b-row>
            <b-col class="scrollableasdf">
              <v-data-table
                :headers="firstTableHeaders"
                :items="dscmy.systems"
                :loading="dscmy.loading"
                class="elevation-1">
                <v-progress-linear slot="progress" color="blue" indeterminate></v-progress-linear>
                  <template slot="no-data">
                    <v-alert :value="true" color="error" icon="warning">
                      Sorry, nothing to display here :(
                    </v-alert>
                  </template>

                  <template slot="headerCell" slot-scope="props">
                    {{ props.header.text }} ({{ distinctData(props.header.value, dscmy.systems).length }})

                    <b-dropdown no-caret variant="link" class="header-filter-icon">
                      <template slot="button-content">
                        <i class="fa fa-filter text-muted"></i>
                      </template>

                      <b-dropdown-header>
                        <b-form-input v-model="search" type="text" placeholder="Filter"></b-form-input>
                      </b-dropdown-header>

                      <b-dropdown-item v-for="item in distinctData(props.header.value, dscmy.systems)" v-bind:key="item">
                        {{ item }}
                      </b-dropdown-item>
                    </b-dropdown>
                  </template>

                  <template slot="items" slot-scope="props">
                      <td><b-link :to="{ path:'/dsc/interfaces/' + props.item.ID }" href="#foo">{{ props.item.System_Name }}</b-link></td>
                      <td>{{ props.item.ITAM_ID }}</td>
                      <td>{{ props.item.fat }}</td>
                      <td>{{ props.item.carbs }}</td>
                  </template>
              </v-data-table>
            </b-col>
            
            <b-col class="scrollableasdf">
              <v-data-table
                :headers="secondTableHeaders"
                :items="dscmy.table"
                :loading="dscmy.loading"
                v-if="secondtable"
                class="elevation-1">
                <v-progress-linear slot="progress" color="blue" indeterminate></v-progress-linear>
                <v-progress-linear slot="progress" color="blue" indeterminate></v-progress-linear>
                  <template slot="no-data">
                    <v-alert :value="true" color="error" icon="warning">
                      Sorry, nothing to display here :(
                    </v-alert>
                  </template>

                  <template slot="headerCell" slot-scope="props">
                    {{ props.header.text }} ({{ distinctData(props.header.value, dscmy.table).length }})

                    <b-dropdown no-caret variant="link" class="header-filter-icon">
                      <template slot="button-content">
                        <i class="fa fa-filter text-muted"></i>
                      </template>

                      <b-dropdown-header>
                        <b-form-input v-model="search" type="text" placeholder="Filter"></b-form-input>
                      </b-dropdown-header>

                      <b-dropdown-item v-for="item in distinctData(props.header.value, dscmy.table)" v-bind:key="item">
                        {{ item }}
                      </b-dropdown-item>
                    </b-dropdown>
                  </template>

                  <template slot="items" slot-scope="props">
                      <td><b-link :to="{ path:'/dsc/interfaces/' + $route.params.system + '/details' }" href="#foo" v-b-modal.modallg>{{ props.item.Name }}</b-link></td>
                      <!-- <td><b-link :to="{ path:'/dsc/interfaces/' + route.params.system + "/details" }" v-b-modal.modallg>{{ props.item.name }}</b-link></td> -->
                      <td>{{ props.item.calories }}</td>
                      <td>{{ props.item.fat }}</td>
                      <td>{{ props.item.carbs }}</td>
                      <td>{{ props.item.ITAM_ID }}</td>
                      <td>{{ props.item.fat }}</td>
                      <td>{{ props.item.carbs }}</td>
                      <td>{{ props.item.ITAM_ID }}</td>
                      <td>{{ props.item.fat }}</td>
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
import { mapState, mapActions } from 'vuex'

var dummy = [
  {
    name: 'Frozen Yogurt',
    calories: 159,
    fat: 6.0,
    carbs: 24,
  },
  {
    name: 'Ice cream sandwich',
    calories: 237,
    fat: 9.0,
    carbs: 37,
  },
  {
    name: 'Eclair',
    calories: 262,
    fat: 16.0,
    carbs: 23,
  },
  {
    name: 'Cupcake',
    calories: 305,
    fat: 3.7,
    carbs: 67,
  },
  {
    name: 'Gingerbread',
    calories: 356,
    fat: 16.0,
    carbs: 49,
  },
  {
    name: 'Jelly bean',
    calories: 375,
    fat: 0.0,
    carbs: 94,
  },
  {
    name: 'Lollipop',
    calories: 392,
    fat: 0.2,
    carbs: 98,
  },
  {
    name: 'Honeycomb',
    calories: 408,
    fat: 3.2,
    carbs: 87,
  },
  {
    name: 'Donut',
    calories: 452,
    fat: 25.0,
    carbs: 51,
  },
  {
    name: 'KitKat',
    calories: 518,
    fat: 26.0,
    carbs: 65,
  }
];

export default {
    data () {
      return {
        search: '',
        secondtable: false,
        systemId: "asdf",
        searchForm: {
          systemName: '',
          itamID: '',
          country: '',
          countryMaster: ['a', 'c', 'd'],
          tableName: '',
          colName: '',
        },
        firstTableHeaders: [
          { text: 'System Name', align: 'left', value: 'System_Name', sortable: false },
          { text: 'ITAM ID', align: 'left', value: 'ITAM_ID', sortable: false },
          { text: 'Dataset Custodian', align: 'left', value: 'fat', sortable: false },
          { text: 'Bank ID', align: 'left', value: 'carbs', sortable: false }
        ],
        secondTableHeaders: [
          { text: 'List of CDEs', align: 'left', sortable: false, value: 'Name' },
          { text: 'Immediate Preceding System', align: 'left', sortable: false, value: 'calories' },
          { text: 'SLA(Yes/No)', align: 'left', sortable: false, value: 'fat' },
          { text: 'OLA(Yes/No)', align: 'left', sortable: false, value: 'carbs' },
          { text: 'Immediate SucceedingSystem', align: 'left', sortable: false, value: 'carbs' },
          { text: 'SLA(Yes/No)', align: 'left', sortable: false, value: 'carbs' },
          { text: 'OLA(Yes/No)', align: 'left', sortable: false, value: 'carbs' },
          { text: 'List of Downstream Process', align: 'left', sortable: false, value: 'carbs' },
          { text: 'Downstream Process Owner', align: 'left', sortable: false, value: 'carbs' },
        ],
        datax: dummy
      }
    },
    computed: {
      ...mapState({
        dscmy: state => state.dscmy.all
      }),
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
      distinctData (col, datax) {
        return this._.uniq(
                this._.map(
                  this._.sortBy(datax, col), 
                  col
                )
              );
      },
      systemRowClick (evt) {
        evt.preventDefault();
        this.$router.push({ path: `my/${this.systemId}` });
        this.secondtable = true;
      },
      onSubmit (evt) {
        evt.preventDefault();
        alert(JSON.stringify(this.searchForm));
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
      }
    }
}
</script>
