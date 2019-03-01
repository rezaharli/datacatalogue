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
          <b-card tag="article" class="mb-2" v-if="selectedDetails">
            <b-media class="left-card-media" >
              <h6 class="left-card-title">System Name</h6>
              <p v-html="selectedDetails.SYSTEM_NAME"></p>
            </b-media>
            
            <b-media class="left-card-media">
              <h6 class="left-card-title">ITAM ID</h6>
              <p v-html="selectedDetails.ITAM_ID"></p>
            </b-media>
            
            <b-media class="left-card-media">
              <h6 class="left-card-title">Dataset Custodian</h6>
              <p v-html="selectedDetails.FIRST_NAME"></p>
            </b-media>
            
            <b-media class="left-card-media">
              <h6 class="left-card-title">Bank ID</h6>
              <p v-html="selectedDetails.BANK_ID"></p>
            </b-media>
            
            <b-media class="left-card-media">
              <h6 class="left-card-title">Business Alias Name</h6>
              <p v-html="selectedDetails.ALIAS_NAME"></p>
            </b-media>
            
            <b-media class="left-card-media">
              <h6 class="left-card-title">Table Name</h6>
              <p v-html="selectedDetails.TABLE_NAME"></p>
            </b-media>
            
            <b-media class="left-card-media">
              <h6 class="left-card-title">Column Name</h6>
              <p v-html="selectedDetails.COLUMN_NAME"></p>
            </b-media>
          </b-card>
        </b-col>

        <b-col cols="8"> 
          <b-row>
            <b-col>
              <b-card title="Technical Metadata From System" tag="article" class="mb-2" v-if="selectedDetails">
                <p class="card-text">
                  <b-form>
                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Table Name" label-for="tableName">
                      <b-form-select id="tableName" class="col-8" v-model="selectedDetails.TABLE_NAME" :options="[selectedDetails.TABLE_NAME]">
                      </b-form-select>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Column Name" label-for="columnName">
                      <b-form-select id="columnName" class="col-8" v-model="selectedDetails.COLUMN_NAME" :options="[selectedDetails.COLUMN_NAME]"></b-form-select>
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
                      <p v-html="selectedDetails.CDE"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Status*">
                      <p v-html="selectedDetails.STATUS"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Data Type">
                      <p v-html="selectedDetails.DATA_TYPE"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Data Format">
                      <p v-html="selectedDetails.DATA_FORMAT"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Data Length">
                      <p v-html="selectedDetails.DATA_LENGTH"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Example">
                      <p v-html="selectedDetails.EXAMPLE"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Derived (Yes/No)*">
                      <p v-html="selectedDetails.DERIVED"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Derivation logic*">
                      <p v-html="selectedDetails.DERIVATION_LOGIC"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Sourced from Upstream (Yes/No)*">
                      <p v-html="selectedDetails.SOURCED_FROM_UPSTREAM"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="System Checks*">
                      <p v-html="selectedDetails.SYSTEM_CHECKS"></p>
                    </b-form-group>
                  </b-form>
                </p>
              </b-card>
            </b-col>
          </b-row>

          <b-row>
            <b-col>
              <b-card title="Business Metadata From Domain" tag="article" class="mb-2" v-if="selectedDetails">
                <p class="card-text">
                  <b-form>
                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Domain">
                      <p v-html="selectedDetails.DOMAIN"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Sub Domain">
                      <p v-html="selectedDetails.SUBDOMAIN"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Domain Owner">
                      <p></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Business Term*">
                      <p v-html="selectedDetails.BT_NAME"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Business Term Description">
                      <p v-html="selectedDetails.BUSINESS_DESCRIPTION"></p>
                    </b-form-group>
                  </b-form>
                </p>
              </b-card>
            </b-col>
          </b-row>

          <b-row>
            <b-col>
              <b-card title="Policy Related Information" tag="article" class="mb-2" v-if="selectedDetails">
                <p class="card-text">
                  <b-form>
                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Information Asset Names">
                      <p v-html="selectedDetails.INFO_ASSET_NAME"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Information Asset Description">
                      <p v-html="selectedDetails.INFO_ASSET_DESCRIPTION"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="C - Confidentiality">
                      <p v-html="selectedDetails.CONFIDENTIALITY"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="I - Integrity">
                      <p v-html="selectedDetails.INTEGRITY"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="A - Availability">
                      <p v-html="selectedDetails.AVAILABILITY"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Overall CIA Rating">
                      <p v-html="selectedDetails.OVERALL_CIA_RATING"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Record Categories">
                      <p v-html="selectedDetails.RECORD_CATEGORY"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="PII Flag">
                      <p v-html="selectedDetails.PII_FLAG"></p>
                    </b-form-group>
                  </b-form>
                </p>
              </b-card>
            </b-col>
          </b-row>

          <b-row>
            <b-col>
              <b-card title="Interfaces" tag="article" class="mb-2" v-if="selectedDetails">
                <p class="card-text">
                  <b-form>
                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Immediate Preceding System*">
                      <p v-html="selectedDetails.IMM_PREC_SYSTEM_ID"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Immediate Succeeding System*">
                      <p v-html="selectedDetails.IMM_SUCC_SYSTEM_ID"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="DQ Standards | Threshold*">
                      <p v-html="selectedDetails.THRESHOLD"></p>
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
      selectedDetails: null,
      ddTable: {
        selected: null,
      },
      ddColumn: {
        selected: null,
      },
      excelFields: {
        'System Name': "selectedDetails.SYSTEM_NAME",
        'ITAM ID': "selectedDetails.ITAM_ID",
        'Dataset Custodian': 'selectedDetails.FIRST_NAME',
        'Bank ID' : 'selectedDetails.BANK_ID',
        'Business Alias Name': "selectedDetails.ALIAS_NAME",
        'Table Name': 'selectedDetails.TABLE_NAME',
        'Column Name': 'selectedDetails.COLUMN_NAME',
        'Screen Label Name*': 'asdf',
        'CDE (yes/no)': 'selectedDetails.CDE',
        'Status*': 'selectedDetails.STATUS',
        'Data Type': 'selectedDetails.DATA_TYPE',
        'Data Format': 'selectedDetails.DATA_FORMAT',
        "Data Length": "selectedDetails.DATA_LENGTH",
        "Example": "selectedDetails.EXAMPLE",
        "Derived (Yes/No)*": "selectedDetails.DERIVED",
        "Derivation logic*": "selectedDetails.DERIVATION_LOGIC",
        "Sourced from Upstream (Yes/No)*": "selectedDetails.SOURCED_FROM_UPSTREAM",
        "System Checks*": "selectedDetails.SYSTEM_CHECKS",
        "Domain": "selectedDetails.DOMAIN",
        "Sub Domain": "selectedDetails.SUBDOMAIN",
        "Domain Owner": "asdf",
        "Business Term*": "selectedDetails.BT_NAME",
        "Business Term Description": "selectedDetails.BUSINESS_DESCRIPTION",
        "Information Asset Names": "selectedDetails.INFO_ASSET_NAME",
        "Information Asset Description": "selectedDetails.INFO_ASSET_DESCRIPTION",
        "C - Confidentiality": "selectedDetails.CONFIDENTIALITY",
        "I - Integrity": "selectedDetails.INTEGRITY",
        "A - Availability": "selectedDetails.AVAILABILITY",
        "Overall CIA Rating": "selectedDetails.OVERALL_CIA_RATING",
        "Record Categories": "selectedDetails.RECORD_CATEGORY",
        "PII Flag": "selectedDetails.PII_FLAG",
        "Immediate Preceding System*": "selectedDetails.IMM_PREC_SYSTEM_ID",
        "Immediate Succeeding System*": "selectedDetails.IMM_SUCC_SYSTEM_ID",
        "DQ Standards | Threshold*": "selectedDetails.THRESHOLD"
      }
    }
  },
  computed: {
    ...mapState({
      dscmy: state => state.dscmy.all
    }),
    ddTableOptions () {
      this.ddTable.selected = this.selectedDetails ? this.selectedDetails.Name : null;
      return this.selectedDetails ? [{ value: this.selectedDetails.Name, text: this.selectedDetails.Name }] : [];
    },
    ddColumnOptions () {
      return this.selectedDetails ? _.map(this.selectedDetails.Columns, function(v) { return { value: v.ID, text: v.Name } }) : [];
    },
    exportDatas () {
      if(this.selectedDetails){
        return [{
          selectedDetails: this.selectedDetails,
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
    "dscmy.detailsSource"() {
      if (this.dscmy.detailsSource.length > 0)
        this.selectedDetails = this.dscmy.detailsSource[0];
      else 
        this.selectedDetails = null;
    },
  },
  mounted() {
    this.$refs.modalDetails.show();

    var param = {
      left: parseInt(this.$route.params.system),
      right: parseInt(this.$route.params.details)
    };

    this.getDetails(param);
    // this.selectedDetails = _.find(this.dscmy.systemsSource, ['ID', parseInt(this.$route.params.system)])
    // this.selectedDetails = _.find(this.dscmy.tableSource, ['ID', parseInt(this.$route.params.details)])
  },
  methods: {
    ...mapActions("dscmy", {
      getDetails: "getDetails"
    }),
    handleClose () {
      this.$router.go(-1)
    }
  },
}
</script>