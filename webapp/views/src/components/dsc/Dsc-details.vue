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
legend.col-form-label, label.col-form-label {
    font-weight: bolder;
}
</style>


<template>
  <b-modal hide-footer header-class="modal-details-header setbackground" body-class="setbackground" v-if="showModal" id="modal-details" ref="modalDetails" size="lg" @hidden="handleClose">
    <b-container fluid class="row-kasijarak">
      <b-row v-if="selectedDetails">
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
        <b-col cols="12" v-if="dscmy.detailsLoading">
          <v-progress-linear :indeterminate="true"></v-progress-linear>
          <v-alert :value="true" type="info">
            Please wait while data is loading
          </v-alert>
        </b-col>

        <b-col cols="4" v-if="!dscmy.detailsLoading"> 
          <b-card tag="article" class="mb-2">
            <b-media class="left-card-media" >
              <h6 class="left-card-title">System Name</h6>
              <text-wrap-dialog :fulltext="selectedDetails.SYSTEM_NAME" v-if="selectedDetails"></text-wrap-dialog>
            </b-media>
            
            <b-media class="left-card-media">
              <h6 class="left-card-title">ITAM ID</h6>
              <text-wrap-dialog :fulltext="selectedDetails.ITAM_ID" v-if="selectedDetails"></text-wrap-dialog>
            </b-media>
            
            <b-media class="left-card-media">
              <h6 class="left-card-title">Dataset Custodian</h6>
              <text-wrap-dialog :fulltext="selectedDetails.FIRST_NAME" v-if="selectedDetails"></text-wrap-dialog>
            </b-media>
            
            <b-media class="left-card-media">
              <h6 class="left-card-title">Bank ID</h6>
              <text-wrap-dialog :fulltext="selectedDetails.BANK_ID" v-if="selectedDetails"></text-wrap-dialog>
            </b-media>
            
            <b-media class="left-card-media">
              <h6 class="left-card-title">Business Alias Name</h6>
              <text-wrap-dialog :fulltext="selectedDetails.ALIAS_NAME" v-if="selectedDetails"></text-wrap-dialog>
            </b-media>
            
            <b-media class="left-card-media">
              <h6 class="left-card-title">Table Name</h6>
              <text-wrap-dialog :fulltext="selectedDetails.TABLE_NAME" v-if="selectedDetails"></text-wrap-dialog>
            </b-media>
            
            <b-media class="left-card-media">
              <h6 class="left-card-title">Column Name</h6>
              <text-wrap-dialog :fulltext="selectedDetails.COLUMN_NAME" v-if="selectedDetails"></text-wrap-dialog>
            </b-media>
          </b-card>
        </b-col>

        <b-col cols="8" v-if="!dscmy.detailsLoading"> 
          <b-row>
            <b-col>
              <b-card title="Technical Metadata From System" tag="article" class="mb-2">
                <p class="card-text">
                  <b-form>
                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Table Name" label-for="tableName">
                      <b-form-select id="tableName" class="col-8" v-model="ddTableSelected" :options="ddTableOptions" @change="ddTableChanged"></b-form-select>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Column Name" label-for="columnName">
                      <b-form-select id="columnName" class="col-8" v-model="ddColumnSelected" :options="ddColumnOptions" @change="ddColumnChanged"></b-form-select>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Business Alias Name*" label-for="screenLabelName">
                      <b-form-select id="columnName" class="col-8" v-model="ddScreenLabelSelected" :options="ddScreenLabelOptions" @change="ddCScreenLabelChanged"></b-form-select>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Business Description*">
                      <text-wrap-dialog :fulltext="selectedDetails.BUSINESS_DESCRIPTION" v-if="selectedDetails"></text-wrap-dialog>
                      <!-- <v-dialog v-model="dialog" width="500" v-if="selectedDetails">
                        <p slot="activator" class="col-form-label" v-html="selectedDetails.BUSINESS_DESCRIPTION"><b-link>[more]</b-link></p>

                        <v-card>
                          <v-card-text v-html="selectedDetails.BUSINESS_DESCRIPTION">
                            
                          </v-card-text>
                        </v-card>
                      </v-dialog> -->
                    </b-form-group>
                    
                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="CDE (yes/no)">
                      <text-wrap-dialog :fulltext="selectedDetails.CDE" v-if="selectedDetails"></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Status*">
                      <text-wrap-dialog :fulltext="selectedDetails.STATUS" v-if="selectedDetails"></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Data Type">
                      <text-wrap-dialog :fulltext="selectedDetails.DATA_TYPE" v-if="selectedDetails"></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Data Format">
                      <text-wrap-dialog :fulltext="selectedDetails.DATA_FORMAT" v-if="selectedDetails"></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Data Length">
                      <text-wrap-dialog :fulltext="selectedDetails.DATA_LENGTH" v-if="selectedDetails"></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Example">
                      <text-wrap-dialog :fulltext="selectedDetails.EXAMPLE" v-if="selectedDetails"></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Derived (Yes/No)*">
                      <text-wrap-dialog :fulltext="selectedDetails.DERIVED" v-if="selectedDetails"></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Derivation logic*">
                      <text-wrap-dialog :fulltext="selectedDetails.DERIVATION_LOGIC" v-if="selectedDetails"></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Sourced from Upstream (Yes/No)*">
                      <text-wrap-dialog :fulltext="selectedDetails.SOURCED_FROM_UPSTREAM" v-if="selectedDetails"></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="System Checks*">
                      <text-wrap-dialog :fulltext="selectedDetails.SYSTEM_CHECKS" v-if="selectedDetails"></text-wrap-dialog>
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
                      <text-wrap-dialog :fulltext="selectedDetails.DOMAIN"></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Sub Domain">
                      <text-wrap-dialog :fulltext="selectedDetails.SUBDOMAIN"></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Domain Owner">
                      <text-wrap-dialog :fulltext="selectedDetails.DOMAIN_OWNER"></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Business Term*">
                      <text-wrap-dialog :fulltext="selectedDetails.BT_NAME"></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Business Term Description">
                      <text-wrap-dialog :fulltext="selectedDetails.BUSINESS_DESCRIPTION"></text-wrap-dialog>
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
                      <text-wrap-dialog :fulltext="selectedDetails.INFO_ASSET_NAME"></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Information Asset Description">
                      <text-wrap-dialog :fulltext="selectedDetails.INFO_ASSET_DESCRIPTION"></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="C - Confidentiality">
                      <text-wrap-dialog :fulltext="selectedDetails.CONFIDENTIALITY"></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="I - Integrity">
                      <text-wrap-dialog :fulltext="selectedDetails.INTEGRITY"></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="A - Availability">
                      <text-wrap-dialog :fulltext="selectedDetails.AVAILABILITY"></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Overall CIA Rating">
                      <text-wrap-dialog :fulltext="selectedDetails.OVERALL_CIA_RATING"></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Record Categories">
                      <text-wrap-dialog :fulltext="selectedDetails.RECORD_CATEGORY"></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="PII Flag">
                      <text-wrap-dialog :fulltext="selectedDetails.PII_FLAG"></text-wrap-dialog>
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
                      <text-wrap-dialog :fulltext="selectedDetails.IMM_PREC_SYSTEM_NAME"></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Immediate Succeeding System*">
                      <text-wrap-dialog :fulltext="selectedDetails.IMM_SUCC_SYSTEM_NAME"></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="DQ Standards | Threshold*">
                      <text-wrap-dialog :fulltext="selectedDetails.THRESHOLD"></text-wrap-dialog>
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
import textWrapDialog from '../TextWrapDialog.vue'

export default {
  components: {
    textWrapDialog
  },
  data () {
    return {
      firstload: true,
      showModal: this.$route.meta.showModal,
      selectedDetails: null,
      ddTableSelected: null,
      ddColumnSelected: null,
      ddScreenLabelSelected: null,
      excelFields: {
        'System Name': "selectedDetails.SYSTEM_NAME",
        'ITAM ID': "selectedDetails.ITAM_ID",
        'Dataset Custodian': 'selectedDetails.FIRST_NAME',
        'Bank ID' : 'selectedDetails.BANK_ID',
        'Business Alias Name': "selectedDetails.ALIAS_NAME",
        'Table Name': 'selectedDetails.TABLE_NAME',
        'Column Name': 'selectedDetails.COLUMN_NAME',
        'Business Alias Name*': 'selectedDetails.ALIAS_NAME',
        'Business Description*': 'selectedDetails.BUSINESS_DESCRIPTION',
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
        "Domain Owner": "selectedDetails.DOMAIN_OWNER",
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
        "Immediate Preceding System*": "selectedDetails.IMM_PREC_SYSTEM_NAME",
        "Immediate Succeeding System*": "selectedDetails.IMM_SUCC_SYSTEM_NAME",
        "DQ Standards | Threshold*": "selectedDetails.THRESHOLD"
      }
    }
  },
  computed: {
    ...mapState({
      dscmy: state => state.dscmy.all
    }),
    ddTableOptions () {
      return _.uniq(_.map(this.dscmy.DDSource, "TABLE_NAME"))
    },
    ddColumnOptions () {
      var self = this;
      var filtered = _.filter(self.dscmy.DDSource, (v) => {
        return v.TABLE_NAME == self.ddTableSelected;
      });

      return _.uniq(_.map(filtered, "COLUMN_NAME"));
    },
    ddScreenLabelOptions () {
      var self = this;
      var filtered = _.filter(self.dscmy.DDSource, (v) => {
        return v.TABLE_NAME == self.ddTableSelected && v.COLUMN_NAME == self.ddColumnSelected;
      });

      return _.uniq(_.map(filtered, "ALIAS_NAME"));
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
    ddTableSelected () {
      if( ! this.ddColumnSelected)
        this.ddColumnSelected = this.ddColumnOptions[0];
      else {
        if(this.ddColumnOptions.indexOf(this.ddColumnSelected) == -1)
          this.ddColumnSelected = this.ddColumnOptions[0];
      }
    },
    ddColumnSelected () {
      if( ! this.ddScreenLabelSelected)
        this.ddScreenLabelSelected = this.ddScreenLabelOptions[0];
      else {
        if(this.ddScreenLabelOptions.indexOf(this.ddScreenLabelSelected) == -1)
          this.ddScreenLabelSelected = this.ddScreenLabelOptions[0];
      }
    },
    'ddScreenLabelSelected' (){
      if( ! this.firstload){
        var param = {
          ScreenLabel: this.ddScreenLabelSelected,
          ColumnName: this.ddColumnSelected,
          TableName: this.ddTableSelected,
        };

        this.runGetDetails(param);
      }

      this.firstload = false;
    }
  },
  mounted() {
    this.$refs.modalDetails.show();

    var param = {};
    this.runGetDetails(param)
  },
  methods: {
    ...mapActions("dscmy", {
      getDetails: "getDetails"
    }),
    handleClose () {
      this.$router.go(-1)
    },
    ddTableChanged () {
      
    },
    ddColumnChanged () {
      console.log(this.ddColumnSelected);
      
    },
    ddCScreenLabelChanged () {

    },
    runGetDetails (param){
      var self = this;

      param.Which = self.$route.name;
      param.Left = parseInt(self.$route.params.system);
      param.Right = parseInt(self.$route.params.details);
      param.Column = parseInt(self.$route.params.column);

      self.getDetails(param).then(
        res => {
          if (self.dscmy.detailsSource.length > 0){
            self.selectedDetails = self.dscmy.detailsSource[0];
            self.selectedDetails.CDE = self.selectedDetails.CDE != 0 ? "Yes" : "No";
            self.selectedDetails.STATUS = self.selectedDetails.STATUS != 0 ? "ACTIVE" : "INACTIVE";
            self.selectedDetails.DERIVED = self.selectedDetails.DERIVED != 0 ? "Yes" : "No";
            self.selectedDetails.SOURCED_FROM_UPSTREAM = self.selectedDetails.SOURCED_FROM_UPSTREAM != 0 ? "Yes" : "No";

            Object.keys(self.selectedDetails).forEach((val) => {
              self.selectedDetails[val] = !!self.selectedDetails[val] ? self.selectedDetails[val] : "NA";
            });
            
            setTimeout(() => {
              console.log(self.selectedDetails.TABLE_NAME, self.selectedDetails.COLUMN_NAME, self.selectedDetails.ALIAS_NAME);

              self.ddTableSelected = self.selectedDetails.TABLE_NAME;
              self.ddColumnSelected = self.selectedDetails.COLUMN_NAME;
              self.ddScreenLabelSelected = self.selectedDetails.ALIAS_NAME;

              console.log(self.ddTableSelected, self.ddColumnSelected, self.ddScreenLabelSelected);
            }, 100);
          } else {
            this.selectedDetails = null;
          }
        },
        err => err
      );
    }
  },
}
</script>