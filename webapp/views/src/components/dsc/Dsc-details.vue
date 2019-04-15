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
.collapsed > .when-opened,
:not(.collapsed) > .when-closed {
  display: none;
}
</style>

<template>
  <b-modal hide-footer header-class="modal-details-header setbackground" body-class="setbackground" v-if="showModal" id="modal-details" ref="modalDetails" size="lg" @hidden="handleClose">
    <page-loader v-if="dscmy.detailsLoading" />
    
    <b-container fluid class="row-kasijarak">
      <b-row>
          <b-col cols="4">
          </b-col>
          <b-col cols="8">
            <span class="float-left" style="margin-top: 5px;">*Must be available for CDEs</span>

            <download-excel
              :data   = "exportDatas"
              :fields = "excelFields"
              worksheet = "My Worksheet"
              name    = "filename.xls">
          
              <v-btn small color="success" class="float-right">Export</v-btn>
            </download-excel>
          </b-col>
      </b-row>
    <!-- </b-container>

    <b-container> -->
      <b-row>
        <b-col cols="4"> 
          <b-card tag="article" class="mb-2">
            <b-media class="left-card-media" >
              <h6 class="left-card-title">System Name</h6>
              <text-wrap-dialog :fulltext="selectedDetails ? selectedDetails.SYSTEM_NAME: ''"></text-wrap-dialog>
            </b-media>
            
            <b-media class="left-card-media">
              <h6 class="left-card-title">ITAM ID</h6>
              <text-wrap-dialog :fulltext="selectedDetails ? selectedDetails.ITAM_ID: ''"></text-wrap-dialog>
            </b-media>
            
            <b-media class="left-card-media">
              <h6 class="left-card-title">Dataset Custodian</h6>
              <text-wrap-dialog :fulltext="selectedDetails ? selectedDetails.DATASET_CUSTODIAN: ''"></text-wrap-dialog>
            </b-media>
            
            <b-media class="left-card-media">
              <h6 class="left-card-title">Bank ID</h6>
              <text-wrap-dialog :fulltext="selectedDetails ? selectedDetails.BANK_ID: ''" ></text-wrap-dialog>
            </b-media>
            
            <b-media class="left-card-media">
              <h6 class="left-card-title">Business Alias Name</h6>
              <text-wrap-dialog :fulltext="selectedDetails ? selectedDetails.BUSINESS_ALIAS_NAME: ''"></text-wrap-dialog>
            </b-media>
            
            <b-media class="left-card-media">
              <h6 class="left-card-title">Table Name</h6>
              <text-wrap-dialog :fulltext="selectedDetails ? selectedDetails.TABLE_NAME: ''"></text-wrap-dialog>
            </b-media>
            
            <b-media class="left-card-media">
              <h6 class="left-card-title">Column Name</h6>
              <text-wrap-dialog :fulltext="selectedDetails ? selectedDetails.COLUMN_NAME: ''"></text-wrap-dialog>
            </b-media>
          </b-card>
        </b-col>

        <b-col cols="8"> 
          <b-row>
            <b-col>
              <b-card tag="article" class="mb-2">
                <h4 class="card-title border-0 mb-0 pb-0" v-b-toggle.collapse-1>
                  Technical Metadata From System
                  <i class="when-opened float-right fa fa-chevron-up"></i>
                  <i class="when-closed float-right fa fa-chevron-down"></i>
                </h4>
                <b-collapse id="collapse-1" class="mt-3 pt-4 border-top" visible>
                  <p class="card-text">
                    <b-form>
                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Table Name" label-for="tableName">
                        <b-form-select id="tableName" class="col-8" v-model="ddTableSelected" :options="ddTableOptions"></b-form-select>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Column Name" label-for="columnName">
                        <b-form-select id="columnName" class="col-8" v-model="ddColumnSelected" :options="ddColumnOptions"></b-form-select>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Business Alias Name*" label-for="screenLabelName">
                        <b-form-select id="columnName" class="col-8" v-model="ddScreenLabelSelected" :options="ddScreenLabelOptions"></b-form-select>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Business Alias Description">
                        <text-wrap-dialog :fulltext="selectedDetails ? selectedDetails.BUSINESS_ALIAS_DESCRIPTION : ''"></text-wrap-dialog>
                      </b-form-group>
                      
                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="CDE (yes/no)">
                        <text-wrap-dialog :fulltext="selectedDetails ? selectedDetails.CDE_YES_NO : ''"></text-wrap-dialog>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Status*">
                        <text-wrap-dialog :fulltext="selectedDetails ? selectedDetails.STATUS : ''"></text-wrap-dialog>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Data Type">
                        <text-wrap-dialog :fulltext="selectedDetails ? selectedDetails.DATA_TYPE : ''"></text-wrap-dialog>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Data Format">
                        <text-wrap-dialog :fulltext="selectedDetails ? selectedDetails.DATA_FORMAT : ''"></text-wrap-dialog>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Data Length">
                        <text-wrap-dialog :fulltext="selectedDetails ? selectedDetails.DATA_LENGTH : ''"></text-wrap-dialog>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Example">
                        <text-wrap-dialog :fulltext="selectedDetails ? selectedDetails.EXAMPLE : ''"></text-wrap-dialog>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Derived (Yes/No)*">
                        <text-wrap-dialog :fulltext="selectedDetails ? selectedDetails.DERIVED_YES_NO : ''"></text-wrap-dialog>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Derivation logic*">
                        <text-wrap-dialog :fulltext="selectedDetails ? selectedDetails.DERIVATION_LOGIC : ''"></text-wrap-dialog>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Sourced from Upstream (Yes/No)*">
                        <text-wrap-dialog :fulltext="selectedDetails ? selectedDetails.SOURCED_FROM_UPSTREAM_YES_NO : ''"></text-wrap-dialog>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="System Checks*">
                        <text-wrap-dialog :fulltext="selectedDetails ? selectedDetails.SYSTEM_CHECKS : ''"></text-wrap-dialog>
                      </b-form-group>
                    </b-form>
                  </p>
                </b-collapse>
              </b-card>
            </b-col>
          </b-row>

          <b-row>
            <b-col>
              <b-card tag="article" class="mb-2">
                <h4 class="card-title border-0 mb-0 pb-0" v-b-toggle.collapse-2>
                  Business Metadata From Domain
                  <i class="when-opened float-right fa fa-chevron-up"></i>
                  <i class="when-closed float-right fa fa-chevron-down"></i>
                </h4>
                <b-collapse id="collapse-2" class="mt-3 pt-4 border-top">
                  <p class="card-text">
                    <b-form>
                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Domain">
                        <text-wrap-dialog :fulltext="selectedDetails ? selectedDetails.DOMAIN : ''"></text-wrap-dialog>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Sub Domain">
                        <text-wrap-dialog :fulltext="selectedDetails ? selectedDetails.SUBDOMAIN : ''"></text-wrap-dialog>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Domain Owner">
                        <text-wrap-dialog :fulltext="selectedDetails ? selectedDetails.DOMAIN_OWNER : ''"></text-wrap-dialog>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Business Term*">
                        <text-wrap-dialog :fulltext="selectedDetails ? selectedDetails.BUSINESS_TERM : ''"></text-wrap-dialog>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Business Term Description">
                        <text-wrap-dialog :fulltext="selectedDetails ? selectedDetails.BUSINESS_TERM_DESCRIPTION : ''"></text-wrap-dialog>
                      </b-form-group>
                    </b-form>
                  </p>
                </b-collapse>
              </b-card>
            </b-col>
          </b-row>

          <b-row>
            <b-col>
              <b-card tag="article" class="mb-2">
                <h4 class="card-title border-0 mb-0 pb-0" v-b-toggle.collapse-4>
                  Interfaces
                  <i class="when-opened float-right fa fa-chevron-up"></i>
                  <i class="when-closed float-right fa fa-chevron-down"></i>
                </h4>
                <b-collapse id="collapse-4" class="mt-3 pt-4 border-top">
                  <p class="card-text">
                    <b-form>
                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Immediate Preceding System*">
                        <text-wrap-dialog :fulltext="selectedDetails ? selectedDetails.IMM_PRECEEDING_SYSTEM : ''"></text-wrap-dialog>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Immediate Succeeding System*">
                        <text-wrap-dialog :fulltext="selectedDetails ? selectedDetails.IMM_SUCCEEDING_SYSTEM : ''"></text-wrap-dialog>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="DQ Standards | Threshold*">
                        <text-wrap-dialog :fulltext="selectedDetails ? selectedDetails.THRESHOLD : ''"></text-wrap-dialog>
                      </b-form-group>
                    </b-form>
                  </p>
                </b-collapse>
              </b-card>
            </b-col>
          </b-row>

          <b-row>
            <b-col>
              <b-card tag="article" class="mb-2">
                <h4 class="card-title border-0 mb-0 pb-0" v-b-toggle.collapse-3>
                  Policy Related Information
                  <i class="when-opened float-right fa fa-chevron-up"></i>
                  <i class="when-closed float-right fa fa-chevron-down"></i>
                </h4>
                <b-collapse id="collapse-3" class="mt-3 pt-4 border-top">
                  <p class="card-text">
                    <b-form>
                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Information Asset Names">
                        <text-wrap-dialog :fulltext="selectedDetails ? selectedDetails.INFORMATION_ASSET_NAMES : ''"></text-wrap-dialog>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Information Asset Description">
                        <text-wrap-dialog :fulltext="selectedDetails ? selectedDetails.INFORMATION_ASSET_DESCRIPTION : ''"></text-wrap-dialog>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="C - Confidentiality">
                        <text-wrap-dialog :fulltext="selectedDetails ? selectedDetails.CONFIDENTIALITY : ''"></text-wrap-dialog>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="I - Integrity">
                        <text-wrap-dialog :fulltext="selectedDetails ? selectedDetails.INTEGRITY : ''"></text-wrap-dialog>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="A - Availability">
                        <text-wrap-dialog :fulltext="selectedDetails ? selectedDetails.AVAILABILITY : ''"></text-wrap-dialog>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Overall CIA Rating">
                        <text-wrap-dialog :fulltext="selectedDetails ? selectedDetails.OVERALL_CIA_RATING : ''"></text-wrap-dialog>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Record Categories">
                        <text-wrap-dialog :fulltext="selectedDetails ? selectedDetails.RECORD_CATEGORIES : ''"></text-wrap-dialog>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="PII Flag">
                        <text-wrap-dialog :fulltext="selectedDetails ? selectedDetails.PII_FLAG : ''"></text-wrap-dialog>
                      </b-form-group>
                    </b-form>
                  </p>
                </b-collapse>
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
import pageLoader from '../PageLoader.vue'
import textWrapDialog from '../TextWrapDialog.vue'

export default {
  components: {
    pageLoader, textWrapDialog
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
        'Dataset Custodian': 'selectedDetails.DATASET_CUSTODIAN',
        'Bank ID' : 'selectedDetails.BANK_ID',
        'Business Alias Name': "selectedDetails.BUSINESS_ALIAS_NAME",
        'Table Name': 'selectedDetails.TABLE_NAME',
        'Column Name': 'selectedDetails.COLUMN_NAME',
        'Business Alias Name*': 'selectedDetails.BUSINESS_ALIAS_NAME',
        'Business Alias Description': 'selectedDetails.BUSINESS_ALIAS_DESCRIPTION',
        'CDE (yes/no)': 'selectedDetails.CDE_YES_NO',
        'Status*': 'selectedDetails.STATUS',
        'Data Type': 'selectedDetails.DATA_TYPE',
        'Data Format': 'selectedDetails.DATA_FORMAT',
        "Data Length": "selectedDetails.DATA_LENGTH",
        "Example": "selectedDetails.EXAMPLE",
        "Derived (Yes/No)*": "selectedDetails.DERIVED_YES_NO",
        "Derivation logic*": "selectedDetails.DERIVATION_LOGIC",
        "Sourced from Upstream (Yes/No)*": "selectedDetails.SOURCED_FROM_UPSTREAM_YES_NO",
        "System Checks*": "selectedDetails.SYSTEM_CHECKS",
        "Domain": "selectedDetails.DOMAIN",
        "Sub Domain": "selectedDetails.SUBDOMAIN",
        "Domain Owner": "selectedDetails.DOMAIN_OWNER",
        "Business Term*": "selectedDetails.BUSINESS_TERM",
        "Business Term Description": "selectedDetails.BUSINESS_TERM_DESCRIPTION",
        "Information Asset Names": "selectedDetails.INFORMATION_ASSET_NAMES",
        "Information Asset Description": "selectedDetails.INFORMATION_ASSET_DESCRIPTION",
        "C - Confidentiality": "selectedDetails.CONFIDENTIALITY",
        "I - Integrity": "selectedDetails.INTEGRITY",
        "A - Availability": "selectedDetails.AVAILABILITY",
        "Overall CIA Rating": "selectedDetails.OVERALL_CIA_RATING",
        "Record Categories": "selectedDetails.RECORD_CATEGORIES",
        "PII Flag": "selectedDetails.PII_FLAG",
        "Immediate Preceeding System*": "selectedDetails.IMM_PRECEEDING_SYSTEM",
        "Immediate Succeeding System*": "selectedDetails.IMM_SUCCEEDING_SYSTEM",
        "DQ Standards | Threshold*": "selectedDetails.THRESHOLD"
      }
    }
  },
  computed: {
    ...mapState({
      dscmy: state => state.dscmy.all
    }),
    ddTableOptions () {
      return _.uniq(_.map(this.dscmy.DDSource, (v) => v.TABLE_NAME.toString().trim())).filter(Boolean);
    },
    ddColumnOptions () {
      var self = this;
      var filtered = _.filter(self.dscmy.DDSource, (v) => {
        return v.TABLE_NAME == self.ddTableSelected;
      });

      return _.uniq(_.map(filtered, (v) => v.COLUMN_NAME.toString().trim())).filter(Boolean);
    },
    ddScreenLabelOptions () {
      var self = this;
      var filtered = _.filter(self.dscmy.DDSource, (v) => {
        return v.TABLE_NAME == self.ddTableSelected && v.COLUMN_NAME == self.ddColumnSelected;
      });
      
      return _.uniq(_.map(filtered, (v) => v.BUSINESS_ALIAS_NAME.toString().trim())).filter(Boolean);
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
      if(this.firstload) return;

      if(this.ddColumnOptions[0]) this.ddColumnSelected = this.ddColumnOptions[0];
      else this.ddColumnSelected = "";
    },
    ddColumnSelected () {
      if(this.firstload) return;
      
      if(this.ddScreenLabelOptions[0]) this.ddScreenLabelSelected = this.ddScreenLabelOptions[0];
      else this.ddScreenLabelSelected = "";
    },
    ddScreenLabelSelected (){
      if(this.firstload) return;
      if( ! this.firstload){
        var param = {
          ColumnName: this.ddColumnSelected.toString(),
          TableName: this.ddTableSelected.toString(),
        };

        if(this.ddScreenLabelSelected && this.ddScreenLabelSelected != "NA"){
          param.ScreenLabel = this.ddScreenLabelSelected.toString();
        } 
        
        this.runGetDetails(param);
      }
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
    runGetDetails (param){
      var self = this;

      param.Which = self.$route.name;
      param.Left = parseInt(self.$route.params.system).toString();
      param.Right = parseInt(self.$route.params.details).toString();
      param.Column = parseInt(self.$route.params.column).toString();

      return self.getDetails(param).then(
        res => {
          this.firstload = true;

          if (self.dscmy.detailsSource.length > 0){
            var tmp = self.dscmy.detailsSource[0].Values[0];

            self.selectedDetails = {}
            _.each(Object.keys(tmp), function(v, i){
                self.selectedDetails[v] = _.uniq(
                  _.map(self.dscmy.detailsSource[0].Values, (val) => val[v] ? val[v].toString().trim() : "").filter(Boolean)
                );
                if(v == "DATASET_CUSTODIAN" || v == "BANK_ID"){
                  self.selectedDetails[v] = self.selectedDetails[v].join('; ');
                } else {
                  self.selectedDetails[v] = self.selectedDetails[v].join(', ');
                }
            });
            
            // interrupt
            self.selectedDetails.CDE_YES_NO = self.selectedDetails.CDE_YES_NO != 0 ? "Yes" : "No";
            self.selectedDetails.STATUS = self.selectedDetails.STATUS != 0 ? "ACTIVE" : "INACTIVE";
            self.selectedDetails.DERIVED_YES_NO = self.selectedDetails.DERIVED_YES_NO != 0 ? "Yes" : "No";
            self.selectedDetails.SOURCED_FROM_UPSTREAM_YES_NO = self.selectedDetails.SOURCED_FROM_UPSTREAM_YES_NO != 0 ? "Yes" : "No";

            // make falsy NA
            Object.keys(self.selectedDetails).forEach((val) => {
              self.selectedDetails[val] = !!self.selectedDetails[val].trim() ? self.selectedDetails[val] : "NA";
            });

            self.dscmy.DDSource.map(function(v){
              Object.keys(v).forEach(function(key){
                v[key] = v[key] ? v[key] : "NA"
              })

              return v
            })
            
            setTimeout(() => {
              self.ddTableSelected = self.selectedDetails.TABLE_NAME;
              self.ddColumnSelected = self.selectedDetails.COLUMN_NAME;
              self.ddScreenLabelSelected = self.selectedDetails.BUSINESS_ALIAS_NAME;

              setTimeout(() => {
                this.firstload = false;
              }, 100);
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