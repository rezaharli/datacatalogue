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
                        <b-form-select id="screenlabelname" class="col-8" v-model="store.ddValUltimateSourceSystem.ddUltGoldenSourceSelected" :options="ddUltGoldenSourceOptions"></b-form-select>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Golden Source System Name">
                        <b-form-select id="screenlabelname" class="col-8" v-model="store.ddValUltimateSourceSystem.ddUltGsSystemNameSelected" :options="ddUltGsSystemNameOptions"></b-form-select>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Golden Source ITAM ID">
                        <b-form-select id="screenlabelname" class="col-8" v-model="store.ddValUltimateSourceSystem.ddUltGsItamIdSelected" :options="ddUltGsItamIdOptions"></b-form-select>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Golden Source Table Name">
                        <b-form-select id="screenlabelname" class="col-8" v-model="store.ddValUltimateSourceSystem.ddUltGsTableNameSelected" :options="ddUltGsTableNameOptions"></b-form-select>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Golden Source Column Name">
                        <b-form-select id="screenlabelname" class="col-8" v-model="store.ddValUltimateSourceSystem.ddUltGsColumnNameSelected" :options="ddUltGsColumnNameOptions"></b-form-select>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Golden Source Screen Name">
                        <b-form-select id="screenlabelname" class="col-8" v-model="store.ddValUltimateSourceSystem.ddUltGsDataElementSelected" :options="ddUltGsDataElementOptions"></b-form-select>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Golden Source Business Description">
                        <b-form-select id="screenlabelname" class="col-8" v-model="store.ddValUltimateSourceSystem.ddUltGsDescriptionSelected" :options="ddUltGsDescriptionOptions"></b-form-select>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Golden Source Derived (Yes/No)">
                        <b-form-select id="screenlabelname" class="col-8" v-model="store.ddValUltimateSourceSystem.ddUltGsDerivedSelected" :options="ddUltGsDerivedOptions"></b-form-select>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Golden Source Derivation Logic">
                        <b-form-select id="screenlabelname" class="col-8" v-model="store.ddValUltimateSourceSystem.ddUltGsDerivationLogicSelected" :options="ddUltGsDerivationLogicOptions"></b-form-select>
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
    ddUltGoldenSourceOptions () {
      var self = this;
      var filtered = _.filter(self.store.DDSourceUltimateSourceSystem, (v) => {
        return v.SYSTEM_NAME == self.store.ddValUltimateSourceSystem.ddUltSystemNameSelected 
          && v.ITAM_ID == self.store.ddValUltimateSourceSystem.ddUltItamIDSelected
          && v.TABLE_NAME == self.store.ddValUltimateSourceSystem.ddUltTableNameSelected
          && v.COLUMN_NAME == self.store.ddValUltimateSourceSystem.ddUltColumnNameSelected
          && v.DATA_ELEMENT == self.store.ddValUltimateSourceSystem.ddUltScreenLabelSelected
      });
      
      var ret = _.uniq(_.map(filtered, (v) => v.GOLDEN_SOURCE.toString())).filter(Boolean);
      return ret.length > 0 ? ret : ["NA"];
    },
    ddUltGsSystemNameOptions () {
      var self = this;
      var filtered = _.filter(self.store.DDSourceUltimateSourceSystem, (v) => {
        return v.SYSTEM_NAME == self.store.ddValUltimateSourceSystem.ddUltSystemNameSelected 
          && v.ITAM_ID == self.store.ddValUltimateSourceSystem.ddUltItamIDSelected
          && v.TABLE_NAME == self.store.ddValUltimateSourceSystem.ddUltTableNameSelected
          && v.COLUMN_NAME == self.store.ddValUltimateSourceSystem.ddUltColumnNameSelected
          && v.DATA_ELEMENT == self.store.ddValUltimateSourceSystem.ddUltScreenLabelSelected
          && v.GOLDEN_SOURCE == self.store.ddValUltimateSourceSystem.ddUltGoldenSourceSelected
      });
      
      var ret = _.uniq(_.map(filtered, (v) => v.GS_SYSTEM_NAME.toString())).filter(Boolean);
      return ret.length > 0 ? ret : ["NA"];
    },
    ddUltGsItamIdOptions () {
      var self = this;
      var filtered = _.filter(self.store.DDSourceUltimateSourceSystem, (v) => {
        return v.SYSTEM_NAME == self.store.ddValUltimateSourceSystem.ddUltSystemNameSelected 
          && v.ITAM_ID == self.store.ddValUltimateSourceSystem.ddUltItamIDSelected
          && v.TABLE_NAME == self.store.ddValUltimateSourceSystem.ddUltTableNameSelected
          && v.COLUMN_NAME == self.store.ddValUltimateSourceSystem.ddUltColumnNameSelected
          && v.DATA_ELEMENT == self.store.ddValUltimateSourceSystem.ddUltScreenLabelSelected
          && v.GOLDEN_SOURCE == self.store.ddValUltimateSourceSystem.ddUltGoldenSourceSelected
          && v.GS_SYSTEM_NAME == self.store.ddValUltimateSourceSystem.ddUltGsSystemNameSelected
      });
      
      var ret = _.uniq(_.map(filtered, (v) => v.GF_ITAM_ID.toString())).filter(Boolean);
      return ret.length > 0 ? ret : ["NA"];
    },
    ddUltGsTableNameOptions () {
      var self = this;
      var filtered = _.filter(self.store.DDSourceUltimateSourceSystem, (v) => {
        return v.SYSTEM_NAME == self.store.ddValUltimateSourceSystem.ddUltSystemNameSelected 
          && v.ITAM_ID == self.store.ddValUltimateSourceSystem.ddUltItamIDSelected
          && v.TABLE_NAME == self.store.ddValUltimateSourceSystem.ddUltTableNameSelected
          && v.COLUMN_NAME == self.store.ddValUltimateSourceSystem.ddUltColumnNameSelected
          && v.DATA_ELEMENT == self.store.ddValUltimateSourceSystem.ddUltScreenLabelSelected
          && v.GOLDEN_SOURCE == self.store.ddValUltimateSourceSystem.ddUltGoldenSourceSelected
          && v.GS_SYSTEM_NAME == self.store.ddValUltimateSourceSystem.ddUltGsSystemNameSelected
          && v.GF_ITAM_ID == self.store.ddValUltimateSourceSystem.ddUltGsItamIdSelected
      });
      
      var ret = _.uniq(_.map(filtered, (v) => v.GS_TABLE_NAME.toString())).filter(Boolean);
      return ret.length > 0 ? ret : ["NA"];
    },
    ddUltGsColumnNameOptions () {
      var self = this;
      var filtered = _.filter(self.store.DDSourceUltimateSourceSystem, (v) => {
        return v.SYSTEM_NAME == self.store.ddValUltimateSourceSystem.ddUltSystemNameSelected 
          && v.ITAM_ID == self.store.ddValUltimateSourceSystem.ddUltItamIDSelected
          && v.TABLE_NAME == self.store.ddValUltimateSourceSystem.ddUltTableNameSelected
          && v.COLUMN_NAME == self.store.ddValUltimateSourceSystem.ddUltColumnNameSelected
          && v.DATA_ELEMENT == self.store.ddValUltimateSourceSystem.ddUltScreenLabelSelected
          && v.GOLDEN_SOURCE == self.store.ddValUltimateSourceSystem.ddUltGoldenSourceSelected
          && v.GS_SYSTEM_NAME == self.store.ddValUltimateSourceSystem.ddUltGsSystemNameSelected
          && v.GF_ITAM_ID == self.store.ddValUltimateSourceSystem.ddUltGsItamIdSelected
          && v.GS_TABLE_NAME == self.store.ddValUltimateSourceSystem.ddUltGsTableNameSelected
      });
      
      var ret = _.uniq(_.map(filtered, (v) => v.GS_COLUMN_NAME.toString())).filter(Boolean);
      return ret.length > 0 ? ret : ["NA"];
    },
    ddUltGsDataElementOptions () {
      var self = this;
      var filtered = _.filter(self.store.DDSourceUltimateSourceSystem, (v) => {
        return v.SYSTEM_NAME == self.store.ddValUltimateSourceSystem.ddUltSystemNameSelected 
          && v.ITAM_ID == self.store.ddValUltimateSourceSystem.ddUltItamIDSelected
          && v.TABLE_NAME == self.store.ddValUltimateSourceSystem.ddUltTableNameSelected
          && v.COLUMN_NAME == self.store.ddValUltimateSourceSystem.ddUltColumnNameSelected
          && v.DATA_ELEMENT == self.store.ddValUltimateSourceSystem.ddUltScreenLabelSelected
          && v.GOLDEN_SOURCE == self.store.ddValUltimateSourceSystem.ddUltGoldenSourceSelected
          && v.GS_SYSTEM_NAME == self.store.ddValUltimateSourceSystem.ddUltGsSystemNameSelected
          && v.GF_ITAM_ID == self.store.ddValUltimateSourceSystem.ddUltGsItamIdSelected
          && v.GS_TABLE_NAME == self.store.ddValUltimateSourceSystem.ddUltGsTableNameSelected
          && v.GS_COLUMN_NAME == self.store.ddValUltimateSourceSystem.ddUltGsColumnNameSelected
      });
      
      var ret = _.uniq(_.map(filtered, (v) => v.GS_DATA_ELEMENT.toString())).filter(Boolean);
      return ret.length > 0 ? ret : ["NA"];
    },
    ddUltGsDescriptionOptions () {
      var self = this;
      var filtered = _.filter(self.store.DDSourceUltimateSourceSystem, (v) => {
        return v.SYSTEM_NAME == self.store.ddValUltimateSourceSystem.ddUltSystemNameSelected 
          && v.ITAM_ID == self.store.ddValUltimateSourceSystem.ddUltItamIDSelected
          && v.TABLE_NAME == self.store.ddValUltimateSourceSystem.ddUltTableNameSelected
          && v.COLUMN_NAME == self.store.ddValUltimateSourceSystem.ddUltColumnNameSelected
          && v.DATA_ELEMENT == self.store.ddValUltimateSourceSystem.ddUltScreenLabelSelected
          && v.GOLDEN_SOURCE == self.store.ddValUltimateSourceSystem.ddUltGoldenSourceSelected
          && v.GS_SYSTEM_NAME == self.store.ddValUltimateSourceSystem.ddUltGsSystemNameSelected
          && v.GF_ITAM_ID == self.store.ddValUltimateSourceSystem.ddUltGsItamIdSelected
          && v.GS_TABLE_NAME == self.store.ddValUltimateSourceSystem.ddUltGsTableNameSelected
          && v.GS_COLUMN_NAME == self.store.ddValUltimateSourceSystem.ddUltGsColumnNameSelected
          && v.GS_DATA_ELEMENT == self.store.ddValUltimateSourceSystem.ddUltGsDataElementSelected
      });
      
      var ret = _.uniq(_.map(filtered, (v) => v.GS_DESCRIPTION.toString())).filter(Boolean);
      return ret.length > 0 ? ret : ["NA"];
    },
    ddUltGsDerivedOptions () {
      var self = this;
      var filtered = _.filter(self.store.DDSourceUltimateSourceSystem, (v) => {
        return v.SYSTEM_NAME == self.store.ddValUltimateSourceSystem.ddUltSystemNameSelected 
          && v.ITAM_ID == self.store.ddValUltimateSourceSystem.ddUltItamIDSelected
          && v.TABLE_NAME == self.store.ddValUltimateSourceSystem.ddUltTableNameSelected
          && v.COLUMN_NAME == self.store.ddValUltimateSourceSystem.ddUltColumnNameSelected
          && v.DATA_ELEMENT == self.store.ddValUltimateSourceSystem.ddUltScreenLabelSelected
          && v.GOLDEN_SOURCE == self.store.ddValUltimateSourceSystem.ddUltGoldenSourceSelected
          && v.GS_SYSTEM_NAME == self.store.ddValUltimateSourceSystem.ddUltGsSystemNameSelected
          && v.GF_ITAM_ID == self.store.ddValUltimateSourceSystem.ddUltGsItamIdSelected
          && v.GS_TABLE_NAME == self.store.ddValUltimateSourceSystem.ddUltGsTableNameSelected
          && v.GS_COLUMN_NAME == self.store.ddValUltimateSourceSystem.ddUltGsColumnNameSelected
          && v.GS_DATA_ELEMENT == self.store.ddValUltimateSourceSystem.ddUltGsDataElementSelected
          && v.GS_DESCRIPTION == self.store.ddValUltimateSourceSystem.ddUltGsDescriptionSelected
      });
      
      var ret = _.uniq(_.map(filtered, (v) => v.GS_DERIVED.toString())).filter(Boolean);
      return ret.length > 0 ? ret : ["NA"];
    },
    ddUltGsDerivationLogicOptions () {
      var self = this;
      var filtered = _.filter(self.store.DDSourceUltimateSourceSystem, (v) => {
        return v.SYSTEM_NAME == self.store.ddValUltimateSourceSystem.ddUltSystemNameSelected 
          && v.ITAM_ID == self.store.ddValUltimateSourceSystem.ddUltItamIDSelected
          && v.TABLE_NAME == self.store.ddValUltimateSourceSystem.ddUltTableNameSelected
          && v.COLUMN_NAME == self.store.ddValUltimateSourceSystem.ddUltColumnNameSelected
          && v.DATA_ELEMENT == self.store.ddValUltimateSourceSystem.ddUltScreenLabelSelected
          && v.GOLDEN_SOURCE == self.store.ddValUltimateSourceSystem.ddUltGoldenSourceSelected
          && v.GS_SYSTEM_NAME == self.store.ddValUltimateSourceSystem.ddUltGsSystemNameSelected
          && v.GF_ITAM_ID == self.store.ddValUltimateSourceSystem.ddUltGsItamIdSelected
          && v.GS_TABLE_NAME == self.store.ddValUltimateSourceSystem.ddUltGsTableNameSelected
          && v.GS_COLUMN_NAME == self.store.ddValUltimateSourceSystem.ddUltGsColumnNameSelected
          && v.GS_DATA_ELEMENT == self.store.ddValUltimateSourceSystem.ddUltGsDataElementSelected
          && v.GS_DESCRIPTION == self.store.ddValUltimateSourceSystem.ddUltGsDescriptionSelected
          && v.GS_DERIVED == self.store.ddValUltimateSourceSystem.ddUltGsDerivedSelected
      });
      
      var ret = _.uniq(_.map(filtered, (v) => v.GS_DERIVATION_LOGIC.toString())).filter(Boolean);
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
      if(this.store.firstload) return;

      if (!this.store.firstload) {
        var param = {};

        var setParam = (ddValGroup, names) => {
          names.forEach(name => {
            if (this.store[ddValGroup]["dd" + name + "Selected"] && this.store[ddValGroup]["dd" + name + "Selected"] != "NA") {
              param[name] = this.store[ddValGroup]["dd" + name + "Selected"].toString();
            }
          });
        }
        
        setParam("ddValImmediatePrecedingSystem", ["ImmSystemName"]);

        this.runGetDetails(param);
      }
    },
    'store.ddValImmediatePrecedingSystem.ddImmItamIDSelected' () {
      if(this.store.firstload) return;
      
      if (!this.store.firstload) {
        var param = {};

        var setParam = (ddValGroup, names) => {
          names.forEach(name => {
            if (this.store[ddValGroup]["dd" + name + "Selected"] && this.store[ddValGroup]["dd" + name + "Selected"] != "NA") {
              param[name] = this.store[ddValGroup]["dd" + name + "Selected"].toString();
            }
          });
        }
        
        setParam("ddValImmediatePrecedingSystem", ["ImmSystemName", "ImmItamID"]);

        this.runGetDetails(param);
      }
    },
    'store.ddValImmediatePrecedingSystem.ddImmTableNameSelected' () {
      if(this.store.firstload) return;
      
      if (!this.store.firstload) {
        var param = {};

        var setParam = (ddValGroup, names) => {
          names.forEach(name => {
            if (this.store[ddValGroup]["dd" + name + "Selected"] && this.store[ddValGroup]["dd" + name + "Selected"] != "NA") {
              param[name] = this.store[ddValGroup]["dd" + name + "Selected"].toString();
            }
          });
        }
        
        setParam("ddValImmediatePrecedingSystem", ["ImmSystemName", "ImmItamID", "ImmTableName"]);

        this.runGetDetails(param);
      }
    },
    'store.ddValImmediatePrecedingSystem.ddImmColumnNameSelected' () {
      if(this.store.firstload) return;
      
      if (!this.store.firstload) {
        var param = {};

        var setParam = (ddValGroup, names) => {
          names.forEach(name => {
            if (this.store[ddValGroup]["dd" + name + "Selected"] && this.store[ddValGroup]["dd" + name + "Selected"] != "NA") {
              param[name] = this.store[ddValGroup]["dd" + name + "Selected"].toString();
            }
          });
        }
        
        setParam("ddValImmediatePrecedingSystem", ["ImmSystemName", "ImmItamID", "ImmTableName", "ImmColumnName"]);

        this.runGetDetails(param);
      }
    },
    'store.ddValImmediatePrecedingSystem.ddImmScreenLabelSelected' () {
      if(this.store.firstload) return;
      
      if (!this.store.firstload) {
        var param = {};

        var setParam = (ddValGroup, names) => {
          names.forEach(name => {
            if (this.store[ddValGroup]["dd" + name + "Selected"] && this.store[ddValGroup]["dd" + name + "Selected"] != "NA") {
              param[name] = this.store[ddValGroup]["dd" + name + "Selected"].toString();
            }
          });
        }
        
        setParam("ddValImmediatePrecedingSystem", ["ImmSystemName", "ImmItamID", "ImmTableName", "ImmColumnName", "ImmScreenLabel"]);

        this.runGetDetails(param);
      }
    },
    'store.ddValUltimateSourceSystem.ddUltSystemNameSelected' () {
      if(this.store.firstload) return;
      
      if (!this.store.firstload) {
        var param = {};

        var setParam = (ddValGroup, names) => {
          names.forEach(name => {
            if (this.store[ddValGroup]["dd" + name + "Selected"] && this.store[ddValGroup]["dd" + name + "Selected"] != "NA") {
              param[name] = this.store[ddValGroup]["dd" + name + "Selected"].toString();
            }
          });
        }
        
        setParam("ddValImmediatePrecedingSystem", ["ImmSystemName", "ImmItamID", "ImmTableName", "ImmColumnName", "ImmScreenLabel"]);
        setParam("ddValUltimateSourceSystem", ["UltSystemName"]);

        this.runGetDetails(param);
      }
    },
    'store.ddValUltimateSourceSystem.ddUltItamIDSelected' () {
      if(this.store.firstload) return;
      
      if (!this.store.firstload) {
        var param = {};

        var setParam = (ddValGroup, names) => {
          names.forEach(name => {
            if (this.store[ddValGroup]["dd" + name + "Selected"] && this.store[ddValGroup]["dd" + name + "Selected"] != "NA") {
              param[name] = this.store[ddValGroup]["dd" + name + "Selected"].toString();
            }
          });
        }
        
        setParam("ddValImmediatePrecedingSystem", ["ImmSystemName", "ImmItamID", "ImmTableName", "ImmColumnName", "ImmScreenLabel"]);
        setParam("ddValUltimateSourceSystem", ["UltSystemName", "UltItamID"]);

        this.runGetDetails(param);
      }
    },
    'store.ddValUltimateSourceSystem.ddUltTableNameSelected' () {
      if(this.store.firstload) return;
      
      if (!this.store.firstload) {
        var param = {};

        var setParam = (ddValGroup, names) => {
          names.forEach(name => {
            if (this.store[ddValGroup]["dd" + name + "Selected"] && this.store[ddValGroup]["dd" + name + "Selected"] != "NA") {
              param[name] = this.store[ddValGroup]["dd" + name + "Selected"].toString();
            }
          });
        }
        
        setParam("ddValImmediatePrecedingSystem", ["ImmSystemName", "ImmItamID", "ImmTableName", "ImmColumnName", "ImmScreenLabel"]);
        setParam("ddValUltimateSourceSystem", ["UltSystemName", "UltItamID", "UltTableName"]);

        this.runGetDetails(param);
      }
    },
    'store.ddValUltimateSourceSystem.ddUltColumnNameSelected' () {
      if(this.store.firstload) return;
      
      if (!this.store.firstload) {
        var param = {};

        var setParam = (ddValGroup, names) => {
          names.forEach(name => {
            if (this.store[ddValGroup]["dd" + name + "Selected"] && this.store[ddValGroup]["dd" + name + "Selected"] != "NA") {
              param[name] = this.store[ddValGroup]["dd" + name + "Selected"].toString();
            }
          });
        }
        
        setParam("ddValImmediatePrecedingSystem", ["ImmSystemName", "ImmItamID", "ImmTableName", "ImmColumnName", "ImmScreenLabel"]);
        setParam("ddValUltimateSourceSystem", ["UltSystemName", "UltItamID", "UltTableName", "UltColumnName"]);

        this.runGetDetails(param);
      }
    },
    'store.ddValUltimateSourceSystem.ddUltScreenLabelSelected' (){
      if(this.store.firstload) return;

      if (!this.store.firstload) {
        var param = {};

        var setParam = (ddValGroup, names) => {
          names.forEach(name => {
            if (this.store[ddValGroup]["dd" + name + "Selected"] && this.store[ddValGroup]["dd" + name + "Selected"] != "NA") {
              param[name] = this.store[ddValGroup]["dd" + name + "Selected"].toString();
            }
          });
        }
        
        setParam("ddValImmediatePrecedingSystem", ["ImmSystemName", "ImmItamID", "ImmTableName", "ImmColumnName", "ImmScreenLabel"]);
        setParam("ddValUltimateSourceSystem", ["UltSystemName", "UltItamID", "UltTableName", "UltColumnName", "UltScreenLabel"]);

        this.runGetDetails(param);
      }
    },
    'store.ddValUltimateSourceSystem.ddUltGoldenSourceSelected' (){
      if(this.store.firstload) return;

      if (!this.store.firstload) {
        var param = {};

        var setParam = (ddValGroup, names) => {
          names.forEach(name => {
            if (this.store[ddValGroup]["dd" + name + "Selected"] && this.store[ddValGroup]["dd" + name + "Selected"] != "NA") {
              param[name] = this.store[ddValGroup]["dd" + name + "Selected"].toString();
            }
          });
        }
        
        setParam("ddValImmediatePrecedingSystem", ["ImmSystemName", "ImmItamID", "ImmTableName", "ImmColumnName", "ImmScreenLabel"]);
        setParam("ddValUltimateSourceSystem", ["UltSystemName", "UltItamID", "UltTableName", "UltColumnName", "UltScreenLabel", "UltGoldenSource"]);

        this.runGetDetails(param);
      }
    },
    'store.ddValUltimateSourceSystem.ddUltGsSystemNameSelected' (){
      if(this.store.firstload) return;

      if (!this.store.firstload) {
        var param = {};

        var setParam = (ddValGroup, names) => {
          names.forEach(name => {
            if (this.store[ddValGroup]["dd" + name + "Selected"] && this.store[ddValGroup]["dd" + name + "Selected"] != "NA") {
              param[name] = this.store[ddValGroup]["dd" + name + "Selected"].toString();
            }
          });
        }
        
        setParam("ddValImmediatePrecedingSystem", ["ImmSystemName", "ImmItamID", "ImmTableName", "ImmColumnName", "ImmScreenLabel"])
        setParam("ddValUltimateSourceSystem", ["UltSystemName", "UltItamID", "UltTableName", "UltColumnName", "UltScreenLabel", "UltGoldenSource", "UltGsSystemName"])

        this.runGetDetails(param);
      }
    },
    'store.ddValUltimateSourceSystem.ddUltGsItamIdSelected' (){
      if(this.store.firstload) return;

      if (!this.store.firstload) {
        var param = {};

        var setParam = (ddValGroup, names) => {
          names.forEach(name => {
            if (this.store[ddValGroup]["dd" + name + "Selected"] && this.store[ddValGroup]["dd" + name + "Selected"] != "NA") {
              param[name] = this.store[ddValGroup]["dd" + name + "Selected"].toString();
            }
          });
        }
        
        setParam("ddValImmediatePrecedingSystem", ["ImmSystemName", "ImmItamID", "ImmTableName", "ImmColumnName", "ImmScreenLabel"])
        setParam("ddValUltimateSourceSystem", ["UltSystemName", "UltItamID", "UltTableName", "UltColumnName", "UltScreenLabel", "UltGoldenSource", "UltGsSystemName", "UltGsItamId"])

        this.runGetDetails(param);
      }
    },
    'store.ddValUltimateSourceSystem.ddUltGsTableNameSelected' (){
      if(this.store.firstload) return;

      if (!this.store.firstload) {
        var param = {};

        var setParam = (ddValGroup, names) => {
          names.forEach(name => {
            if (this.store[ddValGroup]["dd" + name + "Selected"] && this.store[ddValGroup]["dd" + name + "Selected"] != "NA") {
              param[name] = this.store[ddValGroup]["dd" + name + "Selected"].toString();
            }
          });
        }
        
        setParam("ddValImmediatePrecedingSystem", ["ImmSystemName", "ImmItamID", "ImmTableName", "ImmColumnName", "ImmScreenLabel"])
        setParam("ddValUltimateSourceSystem", ["UltSystemName", "UltItamID", "UltTableName", "UltColumnName", "UltScreenLabel", "UltGoldenSource", "UltGsSystemName", "UltGsItamId", "UltGsTableName"])

        this.runGetDetails(param);
      }
    },
    'store.ddValUltimateSourceSystem.ddUltGsColumnNameSelected' (){
      if(this.store.firstload) return;

      if (!this.store.firstload) {
        var param = {};

        var setParam = (ddValGroup, names) => {
          names.forEach(name => {
            if (this.store[ddValGroup]["dd" + name + "Selected"] && this.store[ddValGroup]["dd" + name + "Selected"] != "NA") {
              param[name] = this.store[ddValGroup]["dd" + name + "Selected"].toString();
            }
          });
        }
        
        setParam("ddValImmediatePrecedingSystem", ["ImmSystemName", "ImmItamID", "ImmTableName", "ImmColumnName", "ImmScreenLabel"])
        setParam("ddValUltimateSourceSystem", ["UltSystemName", "UltItamID", "UltTableName", "UltColumnName", "UltScreenLabel", "UltGoldenSource", "UltGsSystemName", "UltGsItamId", "UltGsTableName", "UltGsColumnName"])

        this.runGetDetails(param);
      }
    },
    'store.ddValUltimateSourceSystem.ddUltGsDataElementSelected' (){
      if(this.store.firstload) return;

      if (!this.store.firstload) {
        var param = {};

        var setParam = (ddValGroup, names) => {
          names.forEach(name => {
            if (this.store[ddValGroup]["dd" + name + "Selected"] && this.store[ddValGroup]["dd" + name + "Selected"] != "NA") {
              param[name] = this.store[ddValGroup]["dd" + name + "Selected"].toString();
            }
          });
        }
        
        setParam("ddValImmediatePrecedingSystem", ["ImmSystemName", "ImmItamID", "ImmTableName", "ImmColumnName", "ImmScreenLabel"])
        setParam("ddValUltimateSourceSystem", ["UltSystemName", "UltItamID", "UltTableName", "UltColumnName", "UltScreenLabel", "UltGoldenSource", "UltGsSystemName", "UltGsItamId", "UltGsTableName", "UltGsColumnName", "UltGsDataElement"])

        this.runGetDetails(param);
      }
    },
    'store.ddValUltimateSourceSystem.ddUltGsDescriptionSelected' (){
      if(this.store.firstload) return;

      if (!this.store.firstload) {
        var param = {};

        var setParam = (ddValGroup, names) => {
          names.forEach(name => {
            if (this.store[ddValGroup]["dd" + name + "Selected"] && this.store[ddValGroup]["dd" + name + "Selected"] != "NA") {
              param[name] = this.store[ddValGroup]["dd" + name + "Selected"].toString();
            }
          });
        }
        
        setParam("ddValImmediatePrecedingSystem", ["ImmSystemName", "ImmItamID", "ImmTableName", "ImmColumnName", "ImmScreenLabel"])
        setParam("ddValUltimateSourceSystem", ["UltSystemName", "UltItamID", "UltTableName", "UltColumnName", "UltScreenLabel", "UltGoldenSource", "UltGsSystemName", "UltGsItamId", "UltGsTableName", "UltGsColumnName", "UltGsDataElement", "UltGsDescription"])

        this.runGetDetails(param);
      }
    },
    'store.ddValUltimateSourceSystem.ddUltGsDerivedSelected' (){
      if(this.store.firstload) return;

      if (!this.store.firstload) {
        var param = {};

        var setParam = (ddValGroup, names) => {
          names.forEach(name => {
            if (this.store[ddValGroup]["dd" + name + "Selected"] && this.store[ddValGroup]["dd" + name + "Selected"] != "NA") {
              param[name] = this.store[ddValGroup]["dd" + name + "Selected"].toString();
            }
          });
        }
        
        setParam("ddValImmediatePrecedingSystem", ["ImmSystemName", "ImmItamID", "ImmTableName", "ImmColumnName", "ImmScreenLabel"])
        setParam("ddValUltimateSourceSystem", ["UltSystemName", "UltItamID", "UltTableName", "UltColumnName", "UltScreenLabel", "UltGoldenSource", "UltGsSystemName", "UltGsItamId", "UltGsTableName", "UltGsColumnName", "UltGsDataElement", "UltGsDescription", "UltGsDerived"])

        this.runGetDetails(param);
      }
    },
    'store.ddValUltimateSourceSystem.ddUltGsDerivationLogicSelected' (){
      if(this.store.firstload) return;

      if (!this.store.firstload) {
        var param = {};

        var setParam = (ddValGroup, names) => {
          names.forEach(name => {
            if (this.store[ddValGroup]["dd" + name + "Selected"] && this.store[ddValGroup]["dd" + name + "Selected"] != "NA") {
              param[name] = this.store[ddValGroup]["dd" + name + "Selected"].toString();
            }
          });
        }
        
        setParam("ddValImmediatePrecedingSystem", ["ImmSystemName", "ImmItamID", "ImmTableName", "ImmColumnName", "ImmScreenLabel"])
        setParam("ddValUltimateSourceSystem", ["UltSystemName", "UltItamID", "UltTableName", "UltColumnName", "UltScreenLabel", "UltGoldenSource", "UltGsSystemName", "UltGsItamId", "UltGsTableName", "UltGsColumnName", "UltGsDataElement", "UltGsDescription", "UltGsDerived", "UltGsDerivationLogic"])

        this.runGetDetails(param);
      }
    },
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