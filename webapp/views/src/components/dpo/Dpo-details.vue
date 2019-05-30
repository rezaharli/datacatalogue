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
              <h6 class="left-card-title">Downstream Process Name</h6>
              <text-wrap-dialog :fulltext="store.selectedDetailsImmediatePrecedingSystem ? store.selectedDetailsImmediatePrecedingSystem.PROCESS_NAME: ''"></text-wrap-dialog>
            </b-media>
            
            <!-- <b-media class="left-card-media">
              <h6 class="left-card-title">Process Owner</h6>
              <text-wrap-dialog :fulltext="selectedDetails ? selectedDetails.PROCESS_OWNER: ''"></text-wrap-dialog>
            </b-media> -->
            
            <!-- <b-media class="left-card-media">
              <h6 class="left-card-title">Bank ID</h6>
              <text-wrap-dialog :fulltext="selectedDetails ? selectedDetails.BANK_ID: ''"></text-wrap-dialog>
            </b-media> -->
            
            <b-media class="left-card-media">
              <h6 class="left-card-title">Data Element</h6>
              <text-wrap-dialog :fulltext="store.selectedDetailsImmediatePrecedingSystem ? store.selectedDetailsImmediatePrecedingSystem.DATA_ELEMENT: ''"></text-wrap-dialog>
            </b-media>
            
            <!-- <b-media class="left-card-media">
              <h6 class="left-card-title">CDE Rationale</h6>
              <text-wrap-dialog :fulltext="selectedDetails ? selectedDetails.CDE_RATIONALE: ''"></text-wrap-dialog>
            </b-media> -->

            <!-- <b-media class="left-card-media">
              <h6 class="left-card-title" style="text-decoration: underline;"><b-link>CDE Lineage</b-link></h6>
            </b-media> -->
          </b-card>
        </b-col>

        <b-col cols="8"> 
          <b-row>
            <b-col>
              <b-card tag="article" class="mb-2">
                <h4 class="card-title border-0 mb-0 pb-0" v-b-toggle.collapse-1>
                  Immediate Preceding System Details
                  <i class="when-opened float-right fa fa-chevron-up"></i>
                  <i class="when-closed float-right fa fa-chevron-down"></i>
                </h4>
                <b-collapse id="collapse-1" class="mt-3 pt-4 border-top" visible>
                  <p class="card-text">
                    <b-form>
                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="System Name" label-for="systemName">
                        <b-form-select id="systemName" class="col-8" v-model="store.ddValImmediatePrecedingSystem.ddImmSystemNameSelected" :options="ddImmSystemNameOptions"></b-form-select>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="ITAM ID" label-for="itamid">
                        <b-form-select id="itamid" class="col-8" v-model="store.ddValImmediatePrecedingSystem.ddImmItamIDSelected" :options="ddImmItamIDOptions"></b-form-select>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="TableName" label-for="tablename">
                        <b-form-select id="tablename" class="col-8" v-model="store.ddValImmediatePrecedingSystem.ddImmTableNameSelected" :options="ddImmTableNameOptions"></b-form-select>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Column Name" label-for="columnname">
                        <b-form-select id="columnname" class="col-8" v-model="store.ddValImmediatePrecedingSystem.ddImmColumnNameSelected" :options="ddImmColumnNameOptions"></b-form-select>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Business Alias Name" label-for="screenlabelname">
                        <b-form-select id="screenlabelname" class="col-8" v-model="store.ddValImmediatePrecedingSystem.ddImmScreenLabelSelected" :options="ddImmScreenLabelOptions"></b-form-select>
                      </b-form-group>
                      
                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Business Alias Description">
                        <text-wrap-dialog :fulltext="store.selectedDetailsImmediatePrecedingSystem ? store.selectedDetailsImmediatePrecedingSystem.DESCRIPTION : ''"></text-wrap-dialog>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Derived (Yes/No)">
                        <text-wrap-dialog :fulltext="store.selectedDetailsImmediatePrecedingSystem ? store.selectedDetailsImmediatePrecedingSystem.DERIVED : ''"></text-wrap-dialog>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Derivation Logic">
                        <text-wrap-dialog :fulltext="store.selectedDetailsImmediatePrecedingSystem ? store.selectedDetailsImmediatePrecedingSystem.DERIVATION_LOGIC : ''"></text-wrap-dialog>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Data Quality Requirements">
                        <text-wrap-dialog :fulltext="store.selectedDetailsImmediatePrecedingSystem ? store.selectedDetailsImmediatePrecedingSystem.DQ_STANDARDS : ''"></text-wrap-dialog>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Thresholds">
                        <text-wrap-dialog :fulltext="store.selectedDetailsImmediatePrecedingSystem ? store.selectedDetailsImmediatePrecedingSystem.THRESHOLD : ''"></text-wrap-dialog>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Data SLA signed?">
                        <text-wrap-dialog :fulltext="store.selectedDetailsImmediatePrecedingSystem ? store.selectedDetailsImmediatePrecedingSystem.DATA_SLA : ''"></text-wrap-dialog>
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
                  Ultimate Source System Details
                  <i class="when-opened float-right fa fa-chevron-up"></i>
                  <i class="when-closed float-right fa fa-chevron-down"></i>
                </h4>
                <b-collapse id="collapse-2" class="mt-3 pt-4 border-top">
                  <p class="card-text">
                    <b-form>
                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="System Name" label-for="systemName">
                        <b-form-select id="systemName" class="col-8" v-model="store.ddValUltimateSourceSystem.ddUltSystemNameSelected" :options="ddUltSystemNameOptions">
                        </b-form-select>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="ITAM ID" label-for="itamid">
                        <b-form-select id="itamid" class="col-8" v-model="store.ddValUltimateSourceSystem.ddUltItamIDSelected" :options="ddUltItamIDOptions"></b-form-select>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="TableName" label-for="tablename">
                        <b-form-select id="tablename" class="col-8" v-model="store.ddValUltimateSourceSystem.ddUltTableNameSelected" :options="ddUltTableNameOptions"></b-form-select>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Column Name" label-for="columnname">
                        <b-form-select id="columnname" class="col-8" v-model="store.ddValUltimateSourceSystem.ddUltColumnNameSelected" :options="ddUltColumnNameOptions"></b-form-select>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Business Alias Name" label-for="screenlabelname">
                        <b-form-select id="screenlabelname" class="col-8" v-model="store.ddValUltimateSourceSystem.ddUltScreenLabelSelected" :options="ddUltScreenLabelOptions"></b-form-select>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Business Alias Description">
                        <text-wrap-dialog :fulltext="store.selectedDetailsUltimateSourceSystem ? store.selectedDetailsUltimateSourceSystem.DESCRIPTION : 'NA'"></text-wrap-dialog>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Derived?">
                        <text-wrap-dialog :fulltext="store.selectedDetailsUltimateSourceSystem ? store.selectedDetailsUltimateSourceSystem.DERIVED : 'NA'"></text-wrap-dialog>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Derivation Logic">
                        <text-wrap-dialog :fulltext="store.selectedDetailsUltimateSourceSystem ? store.selectedDetailsUltimateSourceSystem.DERIVATION_LOGIC : 'NA'"></text-wrap-dialog>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Data Quality requirements">
                        <text-wrap-dialog :fulltext="store.selectedDetailsUltimateSourceSystem ? store.selectedDetailsUltimateSourceSystem.DQ_STANDARDS : 'NA'"></text-wrap-dialog>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Thresholds">
                        <text-wrap-dialog :fulltext="store.selectedDetailsUltimateSourceSystem ? store.selectedDetailsUltimateSourceSystem.THRESHOLD : 'NA'"></text-wrap-dialog>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Golden Source (Yes/No)">
                        <text-wrap-dialog :fulltext="store.selectedDetailsUltimateSourceSystem ? store.selectedDetailsUltimateSourceSystem.GOLDEN_SOURCE : 'NA'"></text-wrap-dialog>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Golden Source System Name">
                        <text-wrap-dialog :fulltext="store.selectedDetailsUltimateSourceSystem ? store.selectedDetailsUltimateSourceSystem.GS_SYSTEM_NAME : 'NA'"></text-wrap-dialog>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Golden Source ITAM ID">
                        <text-wrap-dialog :fulltext="store.selectedDetailsUltimateSourceSystem ? store.selectedDetailsUltimateSourceSystem.GF_ITAM_ID : 'NA'"></text-wrap-dialog>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Golden Source Table Name">
                        <text-wrap-dialog :fulltext="store.selectedDetailsUltimateSourceSystem ? store.selectedDetailsUltimateSourceSystem.GS_TABLE_NAME : 'NA'"></text-wrap-dialog>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Golden Source Column Name">
                        <text-wrap-dialog :fulltext="store.selectedDetailsUltimateSourceSystem ? store.selectedDetailsUltimateSourceSystem.GS_COLUMN_NAME : 'NA'"></text-wrap-dialog>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Golden Source Screen Name">
                        <text-wrap-dialog :fulltext="store.selectedDetailsUltimateSourceSystem ? store.selectedDetailsUltimateSourceSystem.GS_DATA_ELEMENT : 'NA'"></text-wrap-dialog>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Golden Source Business Description">
                        <text-wrap-dialog :fulltext="store.selectedDetailsUltimateSourceSystem ? store.selectedDetailsUltimateSourceSystem.GS_DESCRIPTION : 'NA'"></text-wrap-dialog>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Golden Source Derived (Yes/No)">
                        <text-wrap-dialog :fulltext="store.selectedDetailsUltimateSourceSystem ? store.selectedDetailsUltimateSourceSystem.GS_DERIVED : 'NA'"></text-wrap-dialog>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Golden Source Derivation Logic">
                        <text-wrap-dialog :fulltext="store.selectedDetailsUltimateSourceSystem ? store.selectedDetailsUltimateSourceSystem.GS_DERIVATION_LOGIC : 'NA'"></text-wrap-dialog>
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
                  Domain View
                  <i class="when-opened float-right fa fa-chevron-up"></i>
                  <i class="when-closed float-right fa fa-chevron-down"></i>
                </h4>
                <b-collapse id="collapse-3" class="mt-3 pt-4 border-top">
                  <p class="card-text">
                    <b-form>
                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Domain">
                        <text-wrap-dialog :fulltext="store.selectedDetailsDomainView ? store.selectedDetailsDomainView.DOMAIN_NAME : 'NA'"></text-wrap-dialog>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Sub Domain">
                        <text-wrap-dialog :fulltext="store.selectedDetailsDomainView ? store.selectedDetailsDomainView.SUBDOMAIN_NAME : 'NA'"></text-wrap-dialog>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Domain Owner">
                        <text-wrap-dialog :fulltext="store.selectedDetailsDomainView ? store.selectedDetailsDomainView.SUBDOMAIN_OWNER : 'NA'"></text-wrap-dialog>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Business Term">
                        <text-wrap-dialog :fulltext="store.selectedDetailsDomainView ? store.selectedDetailsDomainView.BUSINESS_TERM : 'NA'"></text-wrap-dialog>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Business Term Description">
                        <text-wrap-dialog :fulltext="store.selectedDetailsDomainView ? store.selectedDetailsDomainView.BUSINESS_TERM_DESCRIPTION : 'NA'"></text-wrap-dialog>
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
                  Data Standards
                  <i class="when-opened float-right fa fa-chevron-up"></i>
                  <i class="when-closed float-right fa fa-chevron-down"></i>
                </h4>
                <b-collapse id="collapse-4" class="mt-3 pt-4 border-top">
                  <p class="card-text">
                    <b-form>
                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="DQ Standards set by DPO">
                        <text-wrap-dialog :fulltext="store.selectedDetailsDataStandards ? store.selectedDetailsDataStandards.DPO_DQ_STANDARDS : ''"></text-wrap-dialog>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="DQ Standards set at Business Term Level">
                        <text-wrap-dialog :fulltext="store.selectedDetailsDataStandards ? store.selectedDetailsDataStandards.DQ_STANDARDS_BT_LEVEL : ''"></text-wrap-dialog>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Thresholds defined by DPO">
                        <text-wrap-dialog :fulltext="store.selectedDetailsDataStandards ? store.selectedDetailsDataStandards.DPO_THRESHOLD : ''"></text-wrap-dialog>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Thresholds set at Business term level">
                        <text-wrap-dialog :fulltext="store.selectedDetailsDataStandards ? store.selectedDetailsDataStandards.THRESHOLD_BT_LEVEL : ''"></text-wrap-dialog>
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
        'Downstream Process Name': "selectedDetails.DOWNSTREAM_PROCESS",
        'Process Owner': "selectedDetails.PROCESS_OWNER",
        'CDE Name': 'selectedDetails.CDE_NAME',
        'CDE Rationale': 'selectedDetails.CDE_RATIONALE',
        'System Name': 'selectedDetails.SYSTEM_NAME',
        'ITAM ID': 'selectedDetails.ITAM_ID',
        'TableName': 'selectedDetails.TABLE_NAME',
        "Column Name": "selectedDetails.COLUMN_NAME",
        "Business Alias Name": "selectedDetails.DATA_ELEMENT",
        "Business Alias Description": "selectedDetails.IMM_BUSINESS_DESCRIPTION",
        "Derived (Yes/No)": "selectedDetails.IMM_DERIVED_YES_NO",
        "Derivation Logic": "selectedDetails.IMM_DERIVATION_LOGIC",
        "Data Quality Requirements": "selectedDetails.IMM_DQ_REQUIREMENTS",
        "Thresholds": "selectedDetails.IMM_THRESHOLD",
        "Data SLA signed?": "selectedDetails.IMM_DATA_SLA_SIGNED",
        'System Name ': 'selectedDetails.ULT_SYSTEM',
        'ITAM ID ': 'selectedDetails.ULT_ITAM_ID',
        'TableName ': 'selectedDetails.ULT_TABLE_NAME',
        "Column Name ": "selectedDetails.ULT_COLUMN_NAME",
        "Business Alias Name ": "selectedDetails.ULT_SCREEN_LABEL_NAME",
        "Business Alias Description ": "selectedDetails.ULT_BUSINESS_DESCRIPTION",
        "Derived?": "selectedDetails.ULT_DERIVED_YES_NO",
        "Derivation Logic ": "selectedDetails.ULT_DERIVATION_LOGIC",
        "Data Quality requirements ": "selectedDetails.ULT_DQ_REQUIREMENTS",
        "Thresholds ": "selectedDetails.ULT_THRESHOLD",
        "Golden Source (Yes/No)": "selectedDetails.GOLDEN_SOURCE",
        "Golden Source System Name": "selectedDetails.GS_SYSTEM_NAME",
        "Golden Source ITAM ID": "selectedDetails.GS_ITAM_ID",
        "Golden Source Table Name": "selectedDetails.GS_TABLE_NAME",
        "Golden Source Column Name": "selectedDetails.GS_COLUMN_NAME",
        "Golden Source Screen Name": "selectedDetails.GS_SCREEN_NAME",
        "Golden Source Business Description": "selectedDetails.GS_BUSINESS_DESC",
        "Golden Source Derived (Yes/No)": "selectedDetails.GS_DERIVED_YES_NO",
        "Golden Source Derivation Logic": "selectedDetails.GS_DERIVATION_LOGIC",
        "Domain": "selectedDetails.DOMAIN",
        "Sub Domain": "selectedDetails.SUBDOMAIN",
        "Domain Owner": "selectedDetails.DOMAIN_OWNER",
        "Business Term": "selectedDetails.BUSINESS_TERM",
        "Business Term Description": "selectedDetails.BUSINESS_TERM_DESCRIPTION",
        "DQ Standards set by DPO": "selectedDetails.DPO_DQ_STANDARDS",
        "DQ Standards set at Business Term Level": "selectedDetails.DQ_STANDARDS_BT_LEVEL",
        "Thresholds defined by DPO": "selectedDetails.DPO_THRESHOLD"
      }
    }
  },
  computed: {
    ...mapState({
      store: state => state.dpo.all
    }),
    ddImmSystemNameOptions () {
      return _.uniq(_.map(this.store.DDSourceImmediatePrecedingSystem, (v) => v.SYSTEM_NAME.toString())).filter(Boolean);
    },
    ddImmItamIDOptions () {
      var self = this;
      var filtered = _.filter(self.store.DDSourceImmediatePrecedingSystem, (v) => {
        return v.SYSTEM_NAME == self.store.ddValImmediatePrecedingSystem.ddImmSystemNameSelected;
      });

      return _.uniq(_.map(filtered, (v) => v.ITAM_ID.toString())).filter(Boolean);
    },
    ddImmTableNameOptions () {
      var self = this;
      var filtered = _.filter(self.store.DDSourceImmediatePrecedingSystem, (v) => {
        return v.SYSTEM_NAME == self.store.ddValImmediatePrecedingSystem.ddImmSystemNameSelected 
          && v.ITAM_ID == self.store.ddValImmediatePrecedingSystem.ddImmItamIDSelected
      });
      
      return _.uniq(_.map(filtered, (v) => v.TABLE_NAME.toString())).filter(Boolean);
    },
    ddImmColumnNameOptions () {
      var self = this;
      var filtered = _.filter(self.store.DDSourceImmediatePrecedingSystem, (v) => {
        return v.SYSTEM_NAME == self.store.ddValImmediatePrecedingSystem.ddImmSystemNameSelected 
          && v.ITAM_ID == self.store.ddValImmediatePrecedingSystem.ddImmItamIDSelected
          && v.TABLE_NAME == self.store.ddValImmediatePrecedingSystem.ddImmTableNameSelected
      });
      
      return _.uniq(_.map(filtered, (v) => v.COLUMN_NAME.toString())).filter(Boolean);
    },
    ddImmScreenLabelOptions () {
      var self = this;
      var filtered = _.filter(self.store.DDSourceImmediatePrecedingSystem, (v) => {
        return v.SYSTEM_NAME == self.store.ddValImmediatePrecedingSystem.ddImmSystemNameSelected 
          && v.ITAM_ID == self.store.ddValImmediatePrecedingSystem.ddImmItamIDSelected
          && v.TABLE_NAME == self.store.ddValImmediatePrecedingSystem.ddImmTableNameSelected
          && v.COLUMN_NAME == self.store.ddValImmediatePrecedingSystem.ddImmColumnNameSelected
      });
      
      return _.uniq(_.map(filtered, (v) => v.DATA_ELEMENT.toString())).filter(Boolean);
    },
    ddUltSystemNameOptions () {
      var ret = _.uniq(_.map(this.store.DDSourceUltimateSourceSystem, (v) => v.SYSTEM_NAME.toString())).filter(Boolean);
      return ret.length > 0 ? ret : ["NA"];
    },
    ddUltItamIDOptions () {
      var self = this;
      var filtered = _.filter(self.store.DDSourceUltimateSourceSystem, (v) => {
        return v.SYSTEM_NAME == self.store.ddValUltimateSourceSystem.ddUltSystemNameSelected;
      });

      var ret = _.uniq(_.map(filtered, (v) => v.ITAM_ID.toString())).filter(Boolean);
      return ret.length > 0 ? ret : ["NA"];
    },
    ddUltTableNameOptions () {
      var self = this;
      var filtered = _.filter(self.store.DDSourceUltimateSourceSystem, (v) => {
        return v.SYSTEM_NAME == self.store.ddValUltimateSourceSystem.ddUltSystemNameSelected 
          && v.ITAM_ID == self.store.ddValUltimateSourceSystem.ddUltItamIDSelected
      });
      
      var ret = _.uniq(_.map(filtered, (v) => v.TABLE_NAME.toString())).filter(Boolean);
      return ret.length > 0 ? ret : ["NA"];
    },
    ddUltColumnNameOptions () {
      var self = this;
      var filtered = _.filter(self.store.DDSourceUltimateSourceSystem, (v) => {
        return v.SYSTEM_NAME == self.store.ddValUltimateSourceSystem.ddUltSystemNameSelected 
          && v.ITAM_ID == self.store.ddValUltimateSourceSystem.ddUltItamIDSelected
          && v.TABLE_NAME == self.store.ddValUltimateSourceSystem.ddUltTableNameSelected
      });
      
      var ret = _.uniq(_.map(filtered, (v) => v.COLUMN_NAME.toString())).filter(Boolean);
      return ret.length > 0 ? ret : ["NA"];
    },
    ddUltScreenLabelOptions () {
      var self = this;
      var filtered = _.filter(self.store.DDSourceUltimateSourceSystem, (v) => {
        return v.SYSTEM_NAME == self.store.ddValUltimateSourceSystem.ddUltSystemNameSelected 
          && v.ITAM_ID == self.store.ddValUltimateSourceSystem.ddUltItamIDSelected
          && v.TABLE_NAME == self.store.ddValUltimateSourceSystem.ddUltTableNameSelected
          && v.COLUMN_NAME == self.store.ddValUltimateSourceSystem.ddUltColumnNameSelected
      });
      
      var ret = _.uniq(_.map(filtered, (v) => v.DATA_ELEMENT.toString())).filter(Boolean);
      return ret.length > 0 ? ret : ["NA"];
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
    'store.ddValImmediatePrecedingSystem.ddImmSystemNameSelected' () {
      if(this.firstload) return;

      if (!this.firstload) {
        var param = {};

        if (this.store.ddValImmediatePrecedingSystem.ddImmSystemNameSelected && this.store.ddValImmediatePrecedingSystem.ddImmSystemNameSelected != "NA") {
          param.ImmSystemName = this.store.ddValImmediatePrecedingSystem.ddImmSystemNameSelected.toString();
        }

        this.runGetDetails(param);
      }
    },
    'store.ddValImmediatePrecedingSystem.ddImmItamIDSelected' () {
      if(this.firstload) return;
      
      if (!this.firstload) {
        var param = {};

        if (this.store.ddValImmediatePrecedingSystem.ddImmSystemNameSelected && this.store.ddValImmediatePrecedingSystem.ddImmSystemNameSelected != "NA") {
          param.ImmSystemName = this.store.ddValImmediatePrecedingSystem.ddImmSystemNameSelected.toString();
        }
        if (this.store.ddValImmediatePrecedingSystem.ddImmItamIDSelected && this.store.ddValImmediatePrecedingSystem.ddImmItamIDSelected != "NA") {
          param.ImmItamID = this.store.ddValImmediatePrecedingSystem.ddImmItamIDSelected.toString();
        }

        this.runGetDetails(param);
      }
    },
    'store.ddValImmediatePrecedingSystem.ddImmTableNameSelected' () {
      if(this.firstload) return;
      
      if (!this.firstload) {
        var param = {};

        if (this.store.ddValImmediatePrecedingSystem.ddImmSystemNameSelected && this.store.ddValImmediatePrecedingSystem.ddImmSystemNameSelected != "NA") {
          param.ImmSystemName = this.store.ddValImmediatePrecedingSystem.ddImmSystemNameSelected.toString();
        }
        if (this.store.ddValImmediatePrecedingSystem.ddImmItamIDSelected && this.store.ddValImmediatePrecedingSystem.ddImmItamIDSelected != "NA") {
          param.ImmItamID = this.store.ddValImmediatePrecedingSystem.ddImmItamIDSelected.toString();
        }
        if (this.store.ddValImmediatePrecedingSystem.ddImmTableNameSelected && this.store.ddValImmediatePrecedingSystem.ddImmTableNameSelected != "NA") {
          param.ImmTableName = this.store.ddValImmediatePrecedingSystem.ddImmTableNameSelected.toString();
        }

        this.runGetDetails(param);
      }
    },
    'store.ddValImmediatePrecedingSystem.ddImmColumnNameSelected' () {
      if(this.firstload) return;
      
      if (!this.firstload) {
        var param = {};

        if (this.store.ddValImmediatePrecedingSystem.ddImmSystemNameSelected && this.store.ddValImmediatePrecedingSystem.ddImmSystemNameSelected != "NA") {
          param.ImmSystemName = this.store.ddValImmediatePrecedingSystem.ddImmSystemNameSelected.toString();
        }
        if (this.store.ddValImmediatePrecedingSystem.ddImmItamIDSelected && this.store.ddValImmediatePrecedingSystem.ddImmItamIDSelected != "NA") {
          param.ImmItamID = this.store.ddValImmediatePrecedingSystem.ddImmItamIDSelected.toString();
        }
        if (this.store.ddValImmediatePrecedingSystem.ddImmTableNameSelected && this.store.ddValImmediatePrecedingSystem.ddImmTableNameSelected != "NA") {
          param.ImmTableName = this.store.ddValImmediatePrecedingSystem.ddImmTableNameSelected.toString();
        }
        if (this.store.ddValImmediatePrecedingSystem.ddImmColumnNameSelected && this.store.ddValImmediatePrecedingSystem.ddImmColumnNameSelected != "NA") {
          param.ImmColumnName = this.store.ddValImmediatePrecedingSystem.ddImmColumnNameSelected.toString();
        }

        this.runGetDetails(param);
      }
    },
    'store.ddValImmediatePrecedingSystem.ddImmScreenLabelSelected' () {
      if(this.firstload) return;
      
      if (!this.firstload) {
        var param = {};

        if (this.store.ddValImmediatePrecedingSystem.ddImmSystemNameSelected && this.store.ddValImmediatePrecedingSystem.ddImmSystemNameSelected != "NA") {
          param.ImmSystemName = this.store.ddValImmediatePrecedingSystem.ddImmSystemNameSelected.toString();
        }
        if (this.store.ddValImmediatePrecedingSystem.ddImmItamIDSelected && this.store.ddValImmediatePrecedingSystem.ddImmItamIDSelected != "NA") {
          param.ImmItamID = this.store.ddValImmediatePrecedingSystem.ddImmItamIDSelected.toString();
        }
        if (this.store.ddValImmediatePrecedingSystem.ddImmTableNameSelected && this.store.ddValImmediatePrecedingSystem.ddImmTableNameSelected != "NA") {
          param.ImmTableName = this.store.ddValImmediatePrecedingSystem.ddImmTableNameSelected.toString();
        }
        if (this.store.ddValImmediatePrecedingSystem.ddImmColumnNameSelected && this.store.ddValImmediatePrecedingSystem.ddImmColumnNameSelected != "NA") {
          param.ImmColumnName = this.store.ddValImmediatePrecedingSystem.ddImmColumnNameSelected.toString();
        }
        if (this.store.ddValImmediatePrecedingSystem.ddImmScreenLabelSelected && this.store.ddValImmediatePrecedingSystem.ddImmScreenLabelSelected != "NA") {
          param.ImmScreenLabel = this.store.ddValImmediatePrecedingSystem.ddImmScreenLabelSelected.toString();
        }

        this.runGetDetails(param);
      }
    },
    'store.ddValUltimateSourceSystem.ddUltSystemNameSelected' () {
      if(this.firstload) return;
      
      if (!this.firstload) {
        var param = {};

        if (this.store.ddValImmediatePrecedingSystem.ddImmSystemNameSelected && this.store.ddValImmediatePrecedingSystem.ddImmSystemNameSelected != "NA") {
          param.ImmSystemName = this.store.ddValImmediatePrecedingSystem.ddImmSystemNameSelected.toString();
        }
        if (this.store.ddValImmediatePrecedingSystem.ddImmItamIDSelected && this.store.ddValImmediatePrecedingSystem.ddImmItamIDSelected != "NA") {
          param.ImmItamID = this.store.ddValImmediatePrecedingSystem.ddImmItamIDSelected.toString();
        }
        if (this.store.ddValImmediatePrecedingSystem.ddImmTableNameSelected && this.store.ddValImmediatePrecedingSystem.ddImmTableNameSelected != "NA") {
          param.ImmTableName = this.store.ddValImmediatePrecedingSystem.ddImmTableNameSelected.toString();
        }
        if (this.store.ddValImmediatePrecedingSystem.ddImmColumnNameSelected && this.store.ddValImmediatePrecedingSystem.ddImmColumnNameSelected != "NA") {
          param.ImmColumnName = this.store.ddValImmediatePrecedingSystem.ddImmColumnNameSelected.toString();
        }
        if (this.store.ddValImmediatePrecedingSystem.ddImmScreenLabelSelected && this.store.ddValImmediatePrecedingSystem.ddImmScreenLabelSelected != "NA") {
          param.ImmScreenLabel = this.store.ddValImmediatePrecedingSystem.ddImmScreenLabelSelected.toString();
        }
        if (this.store.ddValUltimateSourceSystem.ddUltSystemNameSelected && this.store.ddValUltimateSourceSystem.ddUltSystemNameSelected != "NA") {
          param.UltSystemName = this.store.ddValUltimateSourceSystem.ddUltSystemNameSelected.toString();
        }

        this.runGetDetails(param);
      }
    },
    'store.ddValUltimateSourceSystem.ddUltItamIDSelected' () {
      if(this.firstload) return;
      
      if (!this.firstload) {
        var param = {};

        if (this.store.ddValImmediatePrecedingSystem.ddImmSystemNameSelected && this.store.ddValImmediatePrecedingSystem.ddImmSystemNameSelected != "NA") {
          param.ImmSystemName = this.store.ddValImmediatePrecedingSystem.ddImmSystemNameSelected.toString();
        }
        if (this.store.ddValImmediatePrecedingSystem.ddImmItamIDSelected && this.store.ddValImmediatePrecedingSystem.ddImmItamIDSelected != "NA") {
          param.ImmItamID = this.store.ddValImmediatePrecedingSystem.ddImmItamIDSelected.toString();
        }
        if (this.store.ddValImmediatePrecedingSystem.ddImmTableNameSelected && this.store.ddValImmediatePrecedingSystem.ddImmTableNameSelected != "NA") {
          param.ImmTableName = this.store.ddValImmediatePrecedingSystem.ddImmTableNameSelected.toString();
        }
        if (this.store.ddValImmediatePrecedingSystem.ddImmColumnNameSelected && this.store.ddValImmediatePrecedingSystem.ddImmColumnNameSelected != "NA") {
          param.ImmColumnName = this.store.ddValImmediatePrecedingSystem.ddImmColumnNameSelected.toString();
        }
        if (this.store.ddValImmediatePrecedingSystem.ddImmScreenLabelSelected && this.store.ddValImmediatePrecedingSystem.ddImmScreenLabelSelected != "NA") {
          param.ImmScreenLabel = this.store.ddValImmediatePrecedingSystem.ddImmScreenLabelSelected.toString();
        }
        if (this.store.ddValUltimateSourceSystem.ddUltSystemNameSelected && this.store.ddValUltimateSourceSystem.ddUltSystemNameSelected != "NA") {
          param.UltSystemName = this.store.ddValUltimateSourceSystem.ddUltSystemNameSelected.toString();
        }
        if (this.store.ddValUltimateSourceSystem.ddUltItamIDSelected && this.store.ddValUltimateSourceSystem.ddUltItamIDSelected != "NA") {
          param.UltItamID = this.store.ddValUltimateSourceSystem.ddUltItamIDSelected.toString();
        }

        this.runGetDetails(param);
      }
    },
    'store.ddValUltimateSourceSystem.ddUltTableNameSelected' () {
      if(this.firstload) return;
      
      if (!this.firstload) {
        var param = {};

        if (this.store.ddValImmediatePrecedingSystem.ddImmSystemNameSelected && this.store.ddValImmediatePrecedingSystem.ddImmSystemNameSelected != "NA") {
          param.ImmSystemName = this.store.ddValImmediatePrecedingSystem.ddImmSystemNameSelected.toString();
        }
        if (this.store.ddValImmediatePrecedingSystem.ddImmItamIDSelected && this.store.ddValImmediatePrecedingSystem.ddImmItamIDSelected != "NA") {
          param.ImmItamID = this.store.ddValImmediatePrecedingSystem.ddImmItamIDSelected.toString();
        }
        if (this.store.ddValImmediatePrecedingSystem.ddImmTableNameSelected && this.store.ddValImmediatePrecedingSystem.ddImmTableNameSelected != "NA") {
          param.ImmTableName = this.store.ddValImmediatePrecedingSystem.ddImmTableNameSelected.toString();
        }
        if (this.store.ddValImmediatePrecedingSystem.ddImmColumnNameSelected && this.store.ddValImmediatePrecedingSystem.ddImmColumnNameSelected != "NA") {
          param.ImmColumnName = this.store.ddValImmediatePrecedingSystem.ddImmColumnNameSelected.toString();
        }
        if (this.store.ddValImmediatePrecedingSystem.ddImmScreenLabelSelected && this.store.ddValImmediatePrecedingSystem.ddImmScreenLabelSelected != "NA") {
          param.ImmScreenLabel = this.store.ddValImmediatePrecedingSystem.ddImmScreenLabelSelected.toString();
        }
        if (this.store.ddValUltimateSourceSystem.ddUltSystemNameSelected && this.store.ddValUltimateSourceSystem.ddUltSystemNameSelected != "NA") {
          param.UltSystemName = this.store.ddValUltimateSourceSystem.ddUltSystemNameSelected.toString();
        }
        if (this.store.ddValUltimateSourceSystem.ddUltItamIDSelected && this.store.ddValUltimateSourceSystem.ddUltItamIDSelected != "NA") {
          param.UltItamID = this.store.ddValUltimateSourceSystem.ddUltItamIDSelected.toString();
        }
        if (this.store.ddValUltimateSourceSystem.ddUltTableNameSelected && this.store.ddValUltimateSourceSystem.ddUltTableNameSelected != "NA") {
          param.UltTableName = this.store.ddValUltimateSourceSystem.ddUltTableNameSelected.toString();
        }

        this.runGetDetails(param);
      }
    },
    'store.ddValUltimateSourceSystem.ddUltColumnNameSelected' () {
      if(this.firstload) return;
      
      if (!this.firstload) {
        var param = {};

        if (this.store.ddValImmediatePrecedingSystem.ddImmSystemNameSelected && this.store.ddValImmediatePrecedingSystem.ddImmSystemNameSelected != "NA") {
          param.ImmSystemName = this.store.ddValImmediatePrecedingSystem.ddImmSystemNameSelected.toString();
        }
        if (this.store.ddValImmediatePrecedingSystem.ddImmItamIDSelected && this.store.ddValImmediatePrecedingSystem.ddImmItamIDSelected != "NA") {
          param.ImmItamID = this.store.ddValImmediatePrecedingSystem.ddImmItamIDSelected.toString();
        }
        if (this.store.ddValImmediatePrecedingSystem.ddImmTableNameSelected && this.store.ddValImmediatePrecedingSystem.ddImmTableNameSelected != "NA") {
          param.ImmTableName = this.store.ddValImmediatePrecedingSystem.ddImmTableNameSelected.toString();
        }
        if (this.store.ddValImmediatePrecedingSystem.ddImmColumnNameSelected && this.store.ddValImmediatePrecedingSystem.ddImmColumnNameSelected != "NA") {
          param.ImmColumnName = this.store.ddValImmediatePrecedingSystem.ddImmColumnNameSelected.toString();
        }
        if (this.store.ddValImmediatePrecedingSystem.ddImmScreenLabelSelected && this.store.ddValImmediatePrecedingSystem.ddImmScreenLabelSelected != "NA") {
          param.ImmScreenLabel = this.store.ddValImmediatePrecedingSystem.ddImmScreenLabelSelected.toString();
        }
        if (this.store.ddValUltimateSourceSystem.ddUltSystemNameSelected && this.store.ddValUltimateSourceSystem.ddUltSystemNameSelected != "NA") {
          param.UltSystemName = this.store.ddValUltimateSourceSystem.ddUltSystemNameSelected.toString();
        }
        if (this.store.ddValUltimateSourceSystem.ddUltItamIDSelected && this.store.ddValUltimateSourceSystem.ddUltItamIDSelected != "NA") {
          param.UltItamID = this.store.ddValUltimateSourceSystem.ddUltItamIDSelected.toString();
        }
        if (this.store.ddValUltimateSourceSystem.ddUltTableNameSelected && this.store.ddValUltimateSourceSystem.ddUltTableNameSelected != "NA") {
          param.UltTableName = this.store.ddValUltimateSourceSystem.ddUltTableNameSelected.toString();
        }
        if (this.store.ddValUltimateSourceSystem.ddUltColumnNameSelected && this.store.ddValUltimateSourceSystem.ddUltColumnNameSelected != "NA") {
          param.UltColumnName = this.store.ddValUltimateSourceSystem.ddUltColumnNameSelected.toString();
        }

        this.runGetDetails(param);
      }
    },
    'store.ddValUltimateSourceSystem.ddUltScreenLabelSelected' (){
      if(this.firstload) return;

      if (!this.firstload) {
        var param = {};

        if (this.store.ddValImmediatePrecedingSystem.ddImmSystemNameSelected && this.store.ddValImmediatePrecedingSystem.ddImmSystemNameSelected != "NA") {
          param.ImmSystemName = this.store.ddValImmediatePrecedingSystem.ddImmSystemNameSelected.toString();
        }
        if (this.store.ddValImmediatePrecedingSystem.ddImmItamIDSelected && this.store.ddValImmediatePrecedingSystem.ddImmItamIDSelected != "NA") {
          param.ImmItamID = this.store.ddValImmediatePrecedingSystem.ddImmItamIDSelected.toString();
        }
        if (this.store.ddValImmediatePrecedingSystem.ddImmTableNameSelected && this.store.ddValImmediatePrecedingSystem.ddImmTableNameSelected != "NA") {
          param.ImmTableName = this.store.ddValImmediatePrecedingSystem.ddImmTableNameSelected.toString();
        }
        if (this.store.ddValImmediatePrecedingSystem.ddImmColumnNameSelected && this.store.ddValImmediatePrecedingSystem.ddImmColumnNameSelected != "NA") {
          param.ImmColumnName = this.store.ddValImmediatePrecedingSystem.ddImmColumnNameSelected.toString();
        }
        if (this.store.ddValImmediatePrecedingSystem.ddImmScreenLabelSelected && this.store.ddValImmediatePrecedingSystem.ddImmScreenLabelSelected != "NA") {
          param.ImmScreenLabel = this.store.ddValImmediatePrecedingSystem.ddImmScreenLabelSelected.toString();
        }
        if (this.store.ddValUltimateSourceSystem.ddUltSystemNameSelected && this.store.ddValUltimateSourceSystem.ddUltSystemNameSelected != "NA") {
          param.UltSystemName = this.store.ddValUltimateSourceSystem.ddUltSystemNameSelected.toString();
        }
        if (this.store.ddValUltimateSourceSystem.ddUltItamIDSelected && this.store.ddValUltimateSourceSystem.ddUltItamIDSelected != "NA") {
          param.UltItamID = this.store.ddValUltimateSourceSystem.ddUltItamIDSelected.toString();
        }
        if (this.store.ddValUltimateSourceSystem.ddUltTableNameSelected && this.store.ddValUltimateSourceSystem.ddUltTableNameSelected != "NA") {
          param.UltTableName = this.store.ddValUltimateSourceSystem.ddUltTableNameSelected.toString();
        }
        if (this.store.ddValUltimateSourceSystem.ddUltColumnNameSelected && this.store.ddValUltimateSourceSystem.ddUltColumnNameSelected != "NA") {
          param.UltColumnName = this.store.ddValUltimateSourceSystem.ddUltColumnNameSelected.toString();
        }
        if (this.store.ddValUltimateSourceSystem.ddUltScreenLabelSelected && this.store.ddValUltimateSourceSystem.ddUltScreenLabelSelected != "NA") {
          param.UltScreenLabel = this.store.ddValUltimateSourceSystem.ddUltScreenLabelSelected.toString();
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
    ...mapActions("dpo", {
      getDetails: "getDetails"
    }),
    handleClose () {
      this.$router.go(-1)
    },
    runGetDetails (param){
      var self = this;

      param.Which = self.$route.name;
      param.Left = self.$route.params.dspname;
      param.Right = self.$route.params.cdename;

      return self.getDetails(param);
    }
  },
}
</script>