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
    <page-loader v-if="store.detailsLoading" />
    
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
              <text-wrap-dialog :fulltext="store.selectedDetails ? store.selectedDetails.SYSTEM_NAME: ''"></text-wrap-dialog>
            </b-media>
            
            <b-media class="left-card-media">
              <h6 class="left-card-title">ITAM ID</h6>
              <text-wrap-dialog :fulltext="store.selectedDetails ? store.selectedDetails.ITAM_ID: ''"></text-wrap-dialog>
            </b-media>
            
            <b-media class="left-card-media">
              <h6 class="left-card-title">Dataset Custodian</h6>
              <text-wrap-dialog :fulltext="store.selectedDetails ? store.selectedDetails.DATASET_CUSTODIAN: ''"></text-wrap-dialog>
            </b-media>
            
            <b-media class="left-card-media">
              <h6 class="left-card-title">Bank ID</h6>
              <text-wrap-dialog :fulltext="store.selectedDetails ? store.selectedDetails.BANK_ID: ''" ></text-wrap-dialog>
            </b-media>
            
            <b-media class="left-card-media">
              <h6 class="left-card-title">Business Alias Name</h6>
              <text-wrap-dialog :fulltext="store.selectedDetails ? store.selectedDetails.BUSINESS_ALIAS_NAME: ''"></text-wrap-dialog>
            </b-media>
            
            <b-media class="left-card-media">
              <h6 class="left-card-title">Table Name</h6>
              <text-wrap-dialog :fulltext="store.selectedDetails ? store.selectedDetails.TABLE_NAME: ''"></text-wrap-dialog>
            </b-media>
            
            <b-media class="left-card-media">
              <h6 class="left-card-title">Column Name</h6>
              <text-wrap-dialog :fulltext="store.selectedDetails ? store.selectedDetails.COLUMN_NAME: ''"></text-wrap-dialog>
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
                        <b-form-select id="tableName" class="col-8" v-model="store.ddVal.ddTableSelected" :options="ddTableOptions"></b-form-select>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Column Name" label-for="columnName">
                        <b-form-select id="columnName" class="col-8" v-model="store.ddVal.ddColumnSelected" :options="ddColumnOptions"></b-form-select>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Business Alias Name*" label-for="screenLabelName">
                        <b-form-select id="screenLabelName" class="col-8" v-model="store.ddVal.ddScreenLabelSelected" :options="ddScreenLabelOptions"></b-form-select>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Business Alias Description">
                        <text-wrap-dialog :fulltext="store.selectedDetails ? store.selectedDetails.BUSINESS_ALIAS_DESCRIPTION : ''"></text-wrap-dialog>
                      </b-form-group>
                      
                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="CDE (yes/no)">
                        <text-wrap-dialog :fulltext="store.selectedDetails ? store.selectedDetails.CDE_YES_NO : ''"></text-wrap-dialog>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Status*">
                        <text-wrap-dialog :fulltext="store.selectedDetails ? store.selectedDetails.STATUS : ''"></text-wrap-dialog>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Data Type">
                        <text-wrap-dialog :fulltext="store.selectedDetails ? store.selectedDetails.DATA_TYPE : ''"></text-wrap-dialog>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Data Format">
                        <text-wrap-dialog :fulltext="store.selectedDetails ? store.selectedDetails.DATA_FORMAT : ''"></text-wrap-dialog>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Data Length">
                        <text-wrap-dialog :fulltext="store.selectedDetails ? store.selectedDetails.DATA_LENGTH : ''"></text-wrap-dialog>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Example">
                        <text-wrap-dialog :fulltext="store.selectedDetails ? store.selectedDetails.EXAMPLE : ''"></text-wrap-dialog>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Derived (Yes/No)*">
                        <text-wrap-dialog :fulltext="store.selectedDetails ? store.selectedDetails.DERIVED_YES_NO : ''"></text-wrap-dialog>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Derivation logic*">
                        <text-wrap-dialog :fulltext="store.selectedDetails ? store.selectedDetails.DERIVATION_LOGIC : ''"></text-wrap-dialog>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Sourced from Upstream (Yes/No)*">
                        <text-wrap-dialog :fulltext="store.selectedDetails ? store.selectedDetails.SOURCED_FROM_UPSTREAM_YES_NO : ''"></text-wrap-dialog>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="System Checks*">
                        <text-wrap-dialog :fulltext="store.selectedDetails ? store.selectedDetails.SYSTEM_CHECKS : ''"></text-wrap-dialog>
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
                        <text-wrap-dialog :fulltext="store.selectedDetails ? store.selectedDetails.DOMAIN : ''"></text-wrap-dialog>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Sub Domain">
                        <text-wrap-dialog :fulltext="store.selectedDetails ? store.selectedDetails.SUBDOMAIN : ''"></text-wrap-dialog>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Domain Owner">
                        <text-wrap-dialog :fulltext="store.selectedDetails ? store.selectedDetails.DOMAIN_OWNER : ''"></text-wrap-dialog>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Business Term*">
                        <!-- <text-wrap-dialog :fulltext="store.selectedDetails ? store.selectedDetails.BUSINESS_TERM : ''"></text-wrap-dialog> -->
                        <b-form-select id="businessTerm" class="col-8" v-model="store.ddVal.ddBusinessTermSelected" :options="ddBusinessTermOptions"></b-form-select>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Business Term Description">
                        <text-wrap-dialog :fulltext="store.selectedDetails ? store.selectedDetails.BUSINESS_TERM_DESCRIPTION : ''"></text-wrap-dialog>
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
                        <b-form-select id="prec" class="col-8" v-model="store.ddVal.ddPrecSelected" :options="ddPrecOptions"></b-form-select>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Incoming Data Element" label-for="precIncoming">
                        <b-form-select id="precIncoming" class="col-8" v-model="store.ddVal.ddPrecIncomingSelected" :options="ddPrecIncomingOptions"></b-form-select>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Incoming Derived (Yes/No)*">
                        <text-wrap-dialog :fulltext="store.selectedDetails ? store.selectedDetails.IMM_PREC_DERIVED : ''"></text-wrap-dialog>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Incoming Derivation Logic*">
                        <text-wrap-dialog :fulltext="store.selectedDetails ? store.selectedDetails.IMM_PREC_DERIVATION_LOGIC : ''"></text-wrap-dialog>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Immediate Succeeding System*">
                        <b-form-select id="succ" class="col-8" v-model="store.ddVal.ddSuccSelected" :options="ddSuccOptions"></b-form-select>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Outgoing Data Element" label-for="succIncoming">
                        <b-form-select id="succIncoming" class="col-8" v-model="store.ddVal.ddSuccIncomingSelected" :options="ddSuccIncomingOptions"></b-form-select>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Outgoing Derived (Yes/No)*">
                        <text-wrap-dialog :fulltext="store.selectedDetails ? store.selectedDetails.IMM_SUCC_DERIVED : ''"></text-wrap-dialog>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Outgoing Derivation Logic*">
                        <text-wrap-dialog :fulltext="store.selectedDetails ? store.selectedDetails.IMM_SUCC_DERIVATION_LOGIC : ''"></text-wrap-dialog>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="DQ Standards | Threshold*">
                        <text-wrap-dialog :fulltext="store.selectedDetails ? store.selectedDetails.THRESHOLD : ''"></text-wrap-dialog>
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
                        <text-wrap-dialog :fulltext="store.selectedDetails ? store.selectedDetails.INFORMATION_ASSET_NAMES : ''"></text-wrap-dialog>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Information Asset Description">
                        <text-wrap-dialog :fulltext="store.selectedDetails ? store.selectedDetails.INFORMATION_ASSET_DESCRIPTION : ''"></text-wrap-dialog>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="C - Confidentiality">
                        <text-wrap-dialog :fulltext="store.selectedDetails ? store.selectedDetails.CONFIDENTIALITY : ''"></text-wrap-dialog>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="I - Integrity">
                        <text-wrap-dialog :fulltext="store.selectedDetails ? store.selectedDetails.INTEGRITY : ''"></text-wrap-dialog>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="A - Availability">
                        <text-wrap-dialog :fulltext="store.selectedDetails ? store.selectedDetails.AVAILABILITY : ''"></text-wrap-dialog>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Overall CIA Rating">
                        <text-wrap-dialog :fulltext="store.selectedDetails ? store.selectedDetails.OVERALL_CIA_RATING : ''"></text-wrap-dialog>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Record Categories">
                        <text-wrap-dialog :fulltext="store.selectedDetails ? store.selectedDetails.RECORD_CATEGORIES : ''"></text-wrap-dialog>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="PII Flag">
                        <text-wrap-dialog :fulltext="store.selectedDetails ? store.selectedDetails.PII_FLAG : ''"></text-wrap-dialog>
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
      showModal: this.$route.meta.showModal,
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
        "Immediate Preceeding System*": "selectedDetails.IMM_PRECEEDING_SYSTEM",
        "Incoming Data elements": "selectedDetails.IMM_PREC_INCOMING",
        "Incoming Derived (Yes/No)": "selectedDetails.IMM_PREC_DERIVED",
        "Incoming Derivation Logic": "selectedDetails.IMM_PREC_DERIVATION_LOGIC",
        "Immediate Succeeding System*": "selectedDetails.IMM_SUCCEEDING_SYSTEM",
        "Outgoing Data Element": "selectedDetails.IMM_SUCC_INCOMING",
        "Outgoing Derived(Yes/No)": "selectedDetails.IMM_SUCC_DERIVED",
        "Outgoing Derivation Logic": "selectedDetails.IMM_SUCC_DERIVATION_LOGIC",
        "DQ Standards | Threshold*": "selectedDetails.THRESHOLD",
        "Information Asset Names": "selectedDetails.INFORMATION_ASSET_NAMES",
        "Information Asset Description": "selectedDetails.INFORMATION_ASSET_DESCRIPTION",
        "C - Confidentiality": "selectedDetails.CONFIDENTIALITY",
        "I - Integrity": "selectedDetails.INTEGRITY",
        "A - Availability": "selectedDetails.AVAILABILITY",
        "Overall CIA Rating": "selectedDetails.OVERALL_CIA_RATING",
        "Record Categories": "selectedDetails.RECORD_CATEGORIES",
        "PII Flag": "selectedDetails.PII_FLAG",
      }
    }
  },
  computed: {
    ...mapState({
      store: state => state.dscmy.all
    }),
    ddTableOptions () {
      return _.uniq(_.map(this.store.DDSource, (v) => v.TABLE_NAME.toString().trim())).filter(Boolean);
    },
    ddColumnOptions () {
      var self = this;
      var filtered = _.filter(self.store.DDSource, (v) => {
        return v.TABLE_NAME == self.store.ddVal.ddTableSelected;
      });

      return _.uniq(_.map(filtered, (v) => v.COLUMN_NAME.toString().trim())).filter(Boolean);
    },
    ddScreenLabelOptions () {
      var self = this;
      var filtered = _.filter(self.store.DDSource, (v) => {
        return v.TABLE_NAME == self.store.ddVal.ddTableSelected 
          && v.COLUMN_NAME == self.store.ddVal.ddColumnSelected;
      });
      
      return _.uniq(_.map(filtered, (v) => v.BUSINESS_ALIAS_NAME.toString().trim())).filter(Boolean);
    },
    ddBusinessTermOptions () {
      var self = this;
      var filtered = _.filter(self.store.DDSource, (v) => {
        return v.TABLE_NAME == self.store.ddVal.ddTableSelected 
          && v.COLUMN_NAME == self.store.ddVal.ddColumnSelected 
          && v.BUSINESS_ALIAS_NAME == self.store.ddVal.ddScreenLabelSelected;
      });
      
      return _.uniq(_.map(filtered, (v) => v.BUSINESS_TERM.toString().trim())).filter(Boolean);
    },
    ddPrecOptions () {
      var self = this;
      var filtered = _.filter(self.store.DDSource, (v) => {
        return v.TABLE_NAME == self.store.ddVal.ddTableSelected 
          && v.COLUMN_NAME == self.store.ddVal.ddColumnSelected 
          && v.BUSINESS_ALIAS_NAME == self.store.ddVal.ddScreenLabelSelected
          && v.BUSINESS_TERM == self.store.ddVal.ddBusinessTermSelected;
      });
      
      return _.uniq(_.map(filtered, (v) => v.IMM_PRECEEDING_SYSTEM.toString().trim())).filter(Boolean);
    },
    ddPrecIncomingOptions () {
      var self = this;
      var filtered = _.filter(self.store.DDSource, (v) => {
        return v.TABLE_NAME == self.store.ddVal.ddTableSelected 
          && v.COLUMN_NAME == self.store.ddVal.ddColumnSelected 
          && v.BUSINESS_ALIAS_NAME == self.store.ddVal.ddScreenLabelSelected
          && v.BUSINESS_TERM == self.store.ddVal.ddBusinessTermSelected
          && v.IMM_PRECEEDING_SYSTEM == self.store.ddVal.ddPrecSelected;
      });
      
      return _.uniq(_.map(filtered, (v) => v.IMM_PREC_INCOMING.toString().trim())).filter(Boolean);
    },
    ddSuccOptions () {
      var self = this;
      var filtered = _.filter(self.store.DDSource, (v) => {
        return v.TABLE_NAME == self.store.ddVal.ddTableSelected 
          && v.COLUMN_NAME == self.store.ddVal.ddColumnSelected 
          && v.BUSINESS_ALIAS_NAME == self.store.ddVal.ddScreenLabelSelected
          && v.BUSINESS_TERM == self.store.ddVal.ddBusinessTermSelected
          && v.IMM_PRECEEDING_SYSTEM == self.store.ddVal.ddPrecSelected
          && v.IMM_PREC_INCOMING == self.store.ddVal.ddPrecIncomingSelected;
      });
      
      return _.uniq(_.map(filtered, (v) => v.IMM_SUCCEEDING_SYSTEM.toString().trim())).filter(Boolean);
    },
    ddSuccIncomingOptions () {
      var self = this;
      var filtered = _.filter(self.store.DDSource, (v) => {
        return v.TABLE_NAME == self.store.ddVal.ddTableSelected 
          && v.COLUMN_NAME == self.store.ddVal.ddColumnSelected 
          && v.BUSINESS_ALIAS_NAME == self.store.ddVal.ddScreenLabelSelected
          && v.BUSINESS_TERM == self.store.ddVal.ddBusinessTermSelected
          && v.IMM_PRECEEDING_SYSTEM == self.store.ddVal.ddPrecSelected
          && v.IMM_PREC_INCOMING == self.store.ddVal.ddPrecIncomingSelected
          && v.IMM_SUCCEEDING_SYSTEM == self.store.ddVal.ddSuccSelected;
      });
      
      return _.uniq(_.map(filtered, (v) => v.IMM_SUCC_INCOMING.toString().trim())).filter(Boolean);
    },
    exportDatas () {
      if(this.store.selectedDetails){
        return [{
          selectedDetails: this.store.selectedDetails,
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
    'store.ddVal.ddTableSelected' () {
      if(this.store.firstload) return;

      if(this.ddColumnOptions[0]) this.store.ddVal.ddColumnSelected = this.ddColumnOptions[0];
      else this.store.ddVal.ddColumnSelected = "";
    },
    'store.ddVal.ddColumnSelected' () {
      if(this.store.firstload) return;
      if( ! this.store.firstload){
        var param = {
          ColumnName: this.store.ddVal.ddColumnSelected.toString(),
          TableName: this.store.ddVal.ddTableSelected.toString(),
        };
        
        this.runGetDetails(param);
      }
    },
    'store.ddVal.ddScreenLabelSelected' (){
      if(this.store.firstload) return;
      if( ! this.store.firstload){
        var param = {
          ColumnName: this.store.ddVal.ddColumnSelected.toString(),
          TableName: this.store.ddVal.ddTableSelected.toString(),
        };

        if(this.store.ddVal.ddScreenLabelSelected && this.store.ddVal.ddScreenLabelSelected != "NA"){
          param.ScreenLabel = this.store.ddVal.ddScreenLabelSelected.toString();
        } 
        
        this.runGetDetails(param);
      }
    },
    'store.ddVal.ddBusinessTermSelected' (){
      if(this.store.firstload) return;
      if( ! this.store.firstload){
        var param = {
          ColumnName: this.store.ddVal.ddColumnSelected.toString(),
          TableName: this.store.ddVal.ddTableSelected.toString(),
        };

        if(this.store.ddVal.ddScreenLabelSelected && this.store.ddVal.ddScreenLabelSelected != "NA"){
          param.ScreenLabel = this.store.ddVal.ddScreenLabelSelected.toString();
        } 

        if(this.store.ddVal.ddBusinessTermSelected && this.store.ddVal.ddBusinessTermSelected != "NA"){
          param.BusinessTerm = this.store.ddVal.ddBusinessTermSelected.toString();
        } 
        
        this.runGetDetails(param);
      }
    },
    'store.ddVal.ddPrecSelected' (){
      if(this.store.firstload) return;
      if( ! this.store.firstload){
        var param = {
          ColumnName: this.store.ddVal.ddColumnSelected.toString(),
          TableName: this.store.ddVal.ddTableSelected.toString(),
        };

        if(this.store.ddVal.ddScreenLabelSelected && this.store.ddVal.ddScreenLabelSelected != "NA"){
          param.ScreenLabel = this.store.ddVal.ddScreenLabelSelected.toString();
        } 

        if(this.store.ddVal.ddBusinessTermSelected && this.store.ddVal.ddBusinessTermSelected != "NA"){
          param.BusinessTerm = this.store.ddVal.ddBusinessTermSelected.toString();
        } 

        if(this.store.ddVal.ddPrecSelected && this.store.ddVal.ddPrecSelected != "NA"){
          param.Prec = this.store.ddVal.ddPrecSelected.toString();
        } 
        
        this.runGetDetails(param);
      }
    },
    'store.ddVal.ddPrecIncomingSelected' (){
      if(this.store.firstload) return;
      if( ! this.store.firstload){
        var param = {
          ColumnName: this.store.ddVal.ddColumnSelected.toString(),
          TableName: this.store.ddVal.ddTableSelected.toString(),
        };

        if(this.store.ddVal.ddScreenLabelSelected && this.store.ddVal.ddScreenLabelSelected != "NA"){
          param.ScreenLabel = this.store.ddVal.ddScreenLabelSelected.toString();
        } 

        if(this.store.ddVal.ddBusinessTermSelected && this.store.ddVal.ddBusinessTermSelected != "NA"){
          param.BusinessTerm = this.store.ddVal.ddBusinessTermSelected.toString();
        } 

        if(this.store.ddVal.ddPrecSelected && this.store.ddVal.ddPrecSelected != "NA"){
          param.Prec = this.store.ddVal.ddPrecSelected.toString();
        } 

        if(this.store.ddVal.ddPrecIncomingSelected && this.store.ddVal.ddPrecIncomingSelected != "NA"){
          param.PrecIncoming = this.store.ddVal.ddPrecIncomingSelected.toString();
        } 
        
        this.runGetDetails(param);
      }
    },
    'store.ddVal.ddSuccSelected' (){
      if(this.store.firstload) return;
      if( ! this.store.firstload){
        var param = {
          ColumnName: this.store.ddVal.ddColumnSelected.toString(),
          TableName: this.store.ddVal.ddTableSelected.toString(),
        };

        if(this.store.ddVal.ddScreenLabelSelected && this.store.ddVal.ddScreenLabelSelected != "NA"){
          param.ScreenLabel = this.store.ddVal.ddScreenLabelSelected.toString();
        } 

        if(this.store.ddVal.ddBusinessTermSelected && this.store.ddVal.ddBusinessTermSelected != "NA"){
          param.BusinessTerm = this.store.ddVal.ddBusinessTermSelected.toString();
        } 

        if(this.store.ddVal.ddPrecSelected && this.store.ddVal.ddPrecSelected != "NA"){
          param.Prec = this.store.ddVal.ddPrecSelected.toString();
        } 

        if(this.store.ddVal.ddPrecIncomingSelected && this.store.ddVal.ddPrecIncomingSelected != "NA"){
          param.PrecIncoming = this.store.ddVal.ddPrecIncomingSelected.toString();
        } 

        if(this.store.ddVal.ddSuccSelected && this.store.ddVal.ddSuccSelected != "NA"){
          param.Succ = this.store.ddVal.ddSuccSelected.toString();
        } 
        
        this.runGetDetails(param);
      }
    },
    'store.ddVal.ddSuccIncomingSelected' (){
      if(this.store.firstload) return;
      if( ! this.store.firstload){
        var param = {
          ColumnName: this.store.ddVal.ddColumnSelected.toString(),
          TableName: this.store.ddVal.ddTableSelected.toString(),
        };

        if(this.store.ddVal.ddScreenLabelSelected && this.store.ddVal.ddScreenLabelSelected != "NA"){
          param.ScreenLabel = this.store.ddVal.ddScreenLabelSelected.toString();
        } 

        if(this.store.ddVal.ddBusinessTermSelected && this.store.ddVal.ddBusinessTermSelected != "NA"){
          param.BusinessTerm = this.store.ddVal.ddBusinessTermSelected.toString();
        } 

        if(this.store.ddVal.ddPrecSelected && this.store.ddVal.ddPrecSelected != "NA"){
          param.Prec = this.store.ddVal.ddPrecSelected.toString();
        } 

        if(this.store.ddVal.ddPrecIncomingSelected && this.store.ddVal.ddPrecIncomingSelected != "NA"){
          param.PrecIncoming = this.store.ddVal.ddPrecIncomingSelected.toString();
        } 

        if(this.store.ddVal.ddSuccSelected && this.store.ddVal.ddSuccSelected != "NA"){
          param.Succ = this.store.ddVal.ddSuccSelected.toString();
        } 

        if(this.store.ddVal.ddSuccIncomingSelected && this.store.ddVal.ddSuccIncomingSelected != "NA"){
          param.PrecIncoming = this.store.ddVal.ddSuccIncomingSelected.toString();
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
      param.Left = parseInt(self.$route.params.systemID).toString();
      param.Right = parseInt(self.$route.params.details).toString();
      param.Column = parseInt(self.$route.params.column).toString();

      return self.getDetails(param);
    }
  },
}
</script>