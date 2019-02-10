<style>
  @import '../../assets/styles/dashboard.css';
</style>

<template>
  <b-row>
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
            <b-col>
              <v-data-table
                :headers="firstTableHeaders"
                :items="desserts"
                :loading="true"
                class="elevation-1">
                <v-progress-linear slot="progress" color="blue" indeterminate></v-progress-linear>
                  <template slot="items" slot-scope="props">
                      <td><b-link to="/dsc/my/asdf" href="#foo">{{ props.item.name }}</b-link></td>
                      <td>{{ props.item.calories }}</td>
                      <td>{{ props.item.fat }}</td>
                      <td>{{ props.item.carbs }}</td>
                  </template>
              </v-data-table>
            </b-col>
            
            <b-col>
              <v-data-table
                :headers="secondTableHeaders"
                :items="desserts"
                :loading="true"
                v-if="secondtable"
                class="elevation-1">
                <v-progress-linear slot="progress" color="blue" indeterminate></v-progress-linear>
                  <template slot="items" slot-scope="props">
                      <!-- <td><b-link v-b-modal.modallg>{{ props.item.name }}</b-link></td> -->
                      <td><b-link to="/dsc/my/asdf/details" v-b-modal.modallg>{{ props.item.name }}</b-link></td>
                      <td>{{ props.item.calories }}</td>
                      <td>{{ props.item.fat }}</td>
                      <td>{{ props.item.carbs }}</td>
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
          { text: 'System Name (4)', align: 'left', sortable: false, value: 'name' },
          { text: 'ITAM ID (4)', align: 'left', sortable: false, value: 'calories' },
          { text: 'Dataset Custodian', align: 'left', sortable: false, value: 'fat' },
          { text: 'Bank ID', align: 'left', sortable: false, value: 'carbs' }
        ],
        secondTableHeaders: [
          { text: 'Table Name (2890)', align: 'left', sortable: false, value: 'name' },
          { text: 'Column Name (2890)', align: 'left', sortable: false, value: 'calories' },
          { text: 'Business Alias Name (1250)', align: 'left', sortable: false, value: 'fat' },
          { text: 'CDE (Yes/No)', align: 'left', sortable: false, value: 'carbs' }
        ],
        desserts: [
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
        ]
      }
    },
    watch: {
      $route (to){
        this.secondtable = false;

        if (to.params != undefined) {
          this.secondtable = to.params.system; 
        }
      },
    },
    created() {
      this.secondtable = this.$route.params.system; 
    },
    methods: {
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
