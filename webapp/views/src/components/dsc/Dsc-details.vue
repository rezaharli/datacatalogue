<style>
.left-card-media{
  margin-bottom: 10px;
}
.left-card-title{
  font-weight: bolder;
}
.card-title{
  font-weight: bolder;
  border-bottom: 1px solid #cecece;
  padding-bottom: 9px;
}
.form-group {
  margin-bottom: 10px !important;
}
.col-form-label {
    font-weight: bolder;
}
</style>


<template>
  <b-modal hide-footer header-class="modal-details-header setbackground" body-class="setbackground" v-if="showModal" id="modal-details" ref="modalDetails" size="lg" @hidden="handleClose">
    <b-container fluid class="row-kasijarak">
      <b-row>
        <b-col> 
          <b-col>
              <download-excel
                  :data   = "exportDatas"
                  :fields = "excelFields"
                  worksheet = "My Worksheet"
                  name    = "filename.xls">
              
                  <b-btn size="sm" class="float-right" variant="success">Export</b-btn>
              </download-excel>
            </b-col>
        </b-col>
      </b-row>
    <!-- </b-container>

    <b-container> -->
      <b-row>
        <b-col cols="4"> 
          <b-card tag="article" class="mb-2" v-if="selectedSystem && selectedTable.Columns.length > 0">
            <b-media class="left-card-media" >
              <h6 class="left-card-title">System Name</h6>
              <p v-html="selectedSystem.System_Name"></p>
            </b-media>
            
            <b-media class="left-card-media">
              <h6 class="left-card-title">ITAM ID</h6>
              <p v-html="selectedSystem.ITAM_ID"></p>
            </b-media>
            
            <b-media class="left-card-media">
              <h6 class="left-card-title">Dataset Custodian</h6>
              <p>&nbsp;</p>
            </b-media>
            
            <b-media class="left-card-media">
              <h6 class="left-card-title">Bank ID</h6>
              <p>&nbsp;</p>
            </b-media>
            
            <b-media class="left-card-media">
              <h6 class="left-card-title">Business Alias Name</h6>
              <p v-html="selectedTable.Columns[0].Alias_Name"></p>
            </b-media>
            
            <b-media class="left-card-media">
              <h6 class="left-card-title">Table Name</h6>
              <p v-html="selectedTable.Name"></p>
            </b-media>
            
            <b-media class="left-card-media">
              <h6 class="left-card-title">Column Name</h6>
              <p v-html="selectedTable.Columns[0].Name"></p>
            </b-media>
          </b-card>
        </b-col>

        <b-col cols="8"> 
          <b-row>
            <b-col>
              <b-card title="Technical Metadata From System" tag="article" class="mb-2" v-if="selectedColumn">
                <p class="card-text">
                  <b-form>
                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Table Name" label-for="tableName">
                      <b-form-select id="tableName" class="col-8" v-model="ddTable.selected" :options="ddTableOptions">
                      </b-form-select>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Column Name" label-for="columnName">
                      <b-form-select id="columnName" class="col-8" v-model="ddColumn.selected" :options="ddColumnOptions"></b-form-select>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Screen Label Name*" label-for="screenLabelName">
                      <b-form-select id="screenLabelName" class="col-8"></b-form-select>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Business Description*">
                      <v-dialog v-model="dialog" width="500">
                        <p slot="activator">This field describes the full name of the client and used for capturing lengthy names where the customer name field is unable to accommodat. It is also used for sub-funds.<b-link>[more]</b-link></p>

                        <v-card>
                          <v-card-text>
                            Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.
                          </v-card-text>
                        </v-card>
                      </v-dialog>
                    </b-form-group>
                    
                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="CDE (yes/no)">
                      <p v-html="selectedColumn.CDE"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Status*">
                      <p v-html="selectedColumn.Status"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Data Type">
                      <p v-html="selectedColumn.Data_Type"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Data Format">
                      <p v-html="selectedColumn.Data_Format"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Data Length">
                      <p v-html="selectedColumn.Data_Length"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Example">
                      <p v-html="selectedColumn.Example"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Derived (Yes/No)*">
                      <p v-html="selectedColumn.Derived"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Derivation logic*">
                      <p v-html="selectedColumn.Derivation_Logic"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Sourced from Upstream (Yes/No)*">
                      <p v-html="selectedColumn.Sourced_from_Upstream"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="System Checks*">
                      <p v-html="selectedColumn.System_Checks"></p>
                    </b-form-group>
                  </b-form>
                </p>
              </b-card>
            </b-col>
          </b-row>

          <b-row>
            <b-col>
              <b-card title="Business Metadata From Domain" tag="article" class="mb-2" v-if="selectedColumn">
                <p class="card-text">
                  <b-form>
                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Domain">
                      <p v-html="selectedColumn.BusinessTerms.SubCategory.Category.Name"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Sub Domain">
                      <p v-html="selectedColumn.BusinessTerms.SubCategory.Name"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Domain Owner">
                      <p></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Business Term*">
                      <p v-html="selectedColumn.BusinessTerms.BT_Name"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Business Term Description">
                      <p v-html="selectedColumn.BusinessTerms.Description"></p>
                    </b-form-group>
                  </b-form>
                </p>
              </b-card>
            </b-col>
          </b-row>

          <b-row>
            <b-col>
              <b-card title="Policy Related Information" tag="article" class="mb-2" v-if="selectedColumn">
                <p class="card-text">
                  <b-form>
                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Information Asset Names">
                      <p v-html="selectedColumn.BusinessTerms.Policy.Info_Asset_Name"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Information Asset Description">
                      <p v-html="selectedColumn.BusinessTerms.Policy.Description"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="C - Confidentiality">
                      <p v-html="selectedColumn.BusinessTerms.Policy.Confidentiality"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="I - Integrity">
                      <p v-html="selectedColumn.BusinessTerms.Policy.Integrity"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="A - Availability">
                      <p v-html="selectedColumn.BusinessTerms.Policy.Availability"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Overall CIA Rating">
                      <p v-html="selectedColumn.BusinessTerms.Policy.Overall_CIA_Rating"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Record Categories">
                      <p v-html="selectedColumn.BusinessTerms.Policy.Record_Category"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="PII Flag">
                      <p v-html="selectedColumn.BusinessTerms.Policy.PII_Flag"></p>
                    </b-form-group>
                  </b-form>
                </p>
              </b-card>
            </b-col>
          </b-row>

          <b-row>
            <b-col>
              <b-card title="Interfaces" tag="article" class="mb-2" v-if="selectedColumn">
                <p class="card-text">
                  <b-form>
                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Immediate Preceding System*">
                      <p v-html="selectedColumn.Imm_Prec_System_ID"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Immediate Succeeding System*">
                      <p v-html="selectedColumn.Imm_Succ_System_ID"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="DQ Standards | Threshold*">
                      <p v-html="selectedColumn.Threshold"></p>
                    </b-form-group>
                  </b-form>
                </p>
              </b-card>
            </b-col>
          </b-row>
        </b-col>
      </b-row>

      <b-row>
      </b-row>
    </b-container>
  </b-modal>
</template>

<script>
import { mapState, mapActions } from 'vuex'

export default {
  data () {
    return {
      dialog: false,
      showModal: this.$route.meta.showModal,
      selectedSystem: null,
      selectedTable: null,
      selectedColumn: null,
      ddTable: {
        selected: null,
      },
      ddColumn: {
        selected: null,
      },
      excelFields: {
        'System Name': "selectedSystem.System_Name",
        'ITAM ID': "selectedSystem.ITAM_ID",
        'Dataset Custodian': 'asdf',
        'Bank ID' : 'asdf',
        'Business Alias Name': "selectedColumn.Alias_Name",
        'Table Name': 'selectedTable.Name',
        'Column Name': 'selectedColumn.Name',
        'Screen Label Name*': 'asdf',
        'CDE (yes/no)': 'selectedColumn.CDE',
        'Status*': 'selectedColumn.Status',
        'Data Type': 'selectedColumn.Data_Type',
        'Data Format': 'selectedColumn.Data_Format',
        "Data Length": "selectedColumn.Data_Length",
        "Example": "selectedColumn.Example",
        "Derived (Yes/No)*": "selectedColumn.Derived",
        "Derivation logic*": "selectedColumn.Derivation_Logic",
        "Sourced from Upstream (Yes/No)*": "selectedColumn.Sourced_from_Upstream",
        "System Checks*": "selectedColumn.System_Checks",
        "Domain": "selectedColumn.BusinessTerms.SubCategory.Category.Name",
        "Sub Domain": "selectedColumn.BusinessTerms.SubCategory.Name",
        "Domain Owner": "asdf",
        "Business Term*": "selectedColumn.BusinessTerms.BT_Name",
        "Business Term Description": "selectedColumn.BusinessTerms.Description",
        "Information Asset Names": "selectedColumn.BusinessTerms.Policy.Info_Asset_Name",
        "Information Asset Description": "selectedColumn.BusinessTerms.Policy.Description",
        "C - Confidentiality": "selectedColumn.BusinessTerms.Policy.Confidentiality",
        "I - Integrity": "selectedColumn.BusinessTerms.Policy.Integrity",
        "A - Availability": "selectedColumn.BusinessTerms.Policy.Availability",
        "Overall CIA Rating": "selectedColumn.BusinessTerms.Policy.Overall_CIA_Rating",
        "Record Categories": "selectedColumn.BusinessTerms.Policy.Record_Category",
        "PII Flag": "selectedColumn.BusinessTerms.Policy.PII_Flag",
        "Immediate Preceding System*": "selectedColumn.Imm_Prec_System_ID",
        "Immediate Succeeding System*": "selectedColumn.Imm_Succ_System_ID",
        "DQ Standards | Threshold*": "selectedColumn.Threshold"
      }
    }
  },
  computed: {
    ...mapState({
      dscmy: state => state.dscmy.all
    }),
    count () {
      return this.dscmy.tableSource.length
    },
    ddTableOptions () {
      this.ddTable.selected = this.selectedTable ? this.selectedTable.Name : null;
      return this.selectedTable ? [{ value: this.selectedTable.Name, text: this.selectedTable.Name }] : [];
    },
    ddColumnOptions () {
      return this.selectedTable ? _.map(this.selectedTable.Columns, function(v) { return { value: v.ID, text: v.Name } }) : [];
    },
    exportDatas () {
      if(this.selectedSystem && this.selectedTable && this.selectedColumn){
        return [{
          selectedSystem: this.selectedSystem,
          selectedTable: this.selectedTable,
          selectedColumn: this.selectedColumn
        }]
      } else {
        return [];
      }
    }
  },
  watch: {
    '$route.meta' ({showModal}) {
      this.showModal = showModal;
    },
    count (newCount, oldCount) {
      this.selectedSystem = _.find(this.dscmy.systemsSource, ['ID', parseInt(this.$route.params.system)])
      this.selectedTable = _.find(this.dscmy.tableSource, ['ID', parseInt(this.$route.params.details)])
    },
    selectedTable (){
      this.ddColumn.selected = this.selectedTable ? (this.selectedTable.Columns[0] ? this.selectedTable.Columns[0].ID : null) : null;
      this.selectedColumn = this.selectedTable ? _.find(this.selectedTable.Columns, ['ID', parseInt(this.ddColumn.selected)]) : null;
    }
  },
  mounted() {
    this.$refs.modalDetails.show();
    
    this.selectedSystem = _.find(this.dscmy.systemsSource, ['ID', parseInt(this.$route.params.system)])
    this.selectedTable = _.find(this.dscmy.tableSource, ['ID', parseInt(this.$route.params.details)])
  },
  methods: {
    handleClose () {
      this.$router.go(-1)
    }
  },
}
</script>