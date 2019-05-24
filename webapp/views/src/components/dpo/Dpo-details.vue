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
              <text-wrap-dialog :fulltext="selectedDetailsImmediatePrecedingSystem ? selectedDetailsImmediatePrecedingSystem.PROCESS_NAME: ''"></text-wrap-dialog>
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
              <text-wrap-dialog :fulltext="selectedDetailsImmediatePrecedingSystem ? selectedDetailsImmediatePrecedingSystem.DATA_ELEMENT: ''"></text-wrap-dialog>
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
                        <b-form-select id="systemName" class="col-8" v-model="ddImmSystemNameSelected" :options="ddImmSystemNameOptions"></b-form-select>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="ITAM ID" label-for="itamid">
                        <b-form-select id="itamid" class="col-8" v-model="ddImmItamIDSelected" :options="ddImmItamIDOptions"></b-form-select>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="TableName" label-for="tablename">
                        <b-form-select id="tablename" class="col-8" v-model="ddImmTableNameSelected" :options="ddImmTableNameOptions"></b-form-select>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Column Name" label-for="columnname">
                        <b-form-select id="columnname" class="col-8" v-model="ddImmColumnNameSelected" :options="ddImmColumnNameOptions"></b-form-select>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Business Alias Name" label-for="screenlabelname">
                        <b-form-select id="screenlabelname" class="col-8" v-model="ddImmScreenLabelSelected" :options="ddImmScreenLabelOptions"></b-form-select>
                      </b-form-group>
                      
                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Business Alias Description">
                        <text-wrap-dialog :fulltext="selectedDetailsImmediatePrecedingSystem ? selectedDetailsImmediatePrecedingSystem.DESCRIPTION : ''"></text-wrap-dialog>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Derived (Yes/No)">
                        <text-wrap-dialog :fulltext="selectedDetailsImmediatePrecedingSystem ? selectedDetailsImmediatePrecedingSystem.DERIVED : ''"></text-wrap-dialog>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Derivation Logic">
                        <text-wrap-dialog :fulltext="selectedDetailsImmediatePrecedingSystem ? selectedDetailsImmediatePrecedingSystem.DERIVATION_LOGIC : ''"></text-wrap-dialog>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Data Quality Requirements">
                        <text-wrap-dialog :fulltext="selectedDetailsImmediatePrecedingSystem ? selectedDetailsImmediatePrecedingSystem.DQ_STANDARDS : ''"></text-wrap-dialog>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Thresholds">
                        <text-wrap-dialog :fulltext="selectedDetailsImmediatePrecedingSystem ? selectedDetailsImmediatePrecedingSystem.THRESHOLD : ''"></text-wrap-dialog>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Data SLA signed?">
                        <text-wrap-dialog :fulltext="selectedDetailsImmediatePrecedingSystem ? selectedDetailsImmediatePrecedingSystem.DATA_SLA : ''"></text-wrap-dialog>
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
                        <b-form-select id="systemName" class="col-8" v-model="ddUltSystemNameSelected" :options="ddUltSystemNameOptions">
                        </b-form-select>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="ITAM ID" label-for="itamid">
                        <b-form-select id="itamid" class="col-8" v-model="ddUltItamIDSelected" :options="ddUltItamIDOptions"></b-form-select>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="TableName" label-for="tablename">
                        <b-form-select id="tablename" class="col-8" v-model="ddUltTableNameSelected" :options="ddUltTableNameOptions"></b-form-select>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Column Name" label-for="columnname">
                        <b-form-select id="columnname" class="col-8" v-model="ddUltColumnNameSelected" :options="ddUltColumnNameOptions"></b-form-select>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Business Alias Name" label-for="screenlabelname">
                        <b-form-select id="screenlabelname" class="col-8" v-model="ddUltScreenLabelSelected" :options="ddUltScreenLabelOptions"></b-form-select>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Business Alias Description">
                        <text-wrap-dialog :fulltext="selectedDetailsUltimateSourceSystem ? selectedDetailsUltimateSourceSystem.DESCRIPTION : ''"></text-wrap-dialog>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Derived?">
                        <text-wrap-dialog :fulltext="selectedDetailsUltimateSourceSystem ? selectedDetailsUltimateSourceSystem.DERIVED : ''"></text-wrap-dialog>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Derivation Logic">
                        <text-wrap-dialog :fulltext="selectedDetailsUltimateSourceSystem ? selectedDetailsUltimateSourceSystem.DERIVATION_LOGIC : ''"></text-wrap-dialog>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Data Quality requirements">
                        <text-wrap-dialog :fulltext="selectedDetailsUltimateSourceSystem ? selectedDetailsUltimateSourceSystem.DQ_STANDARDS : ''"></text-wrap-dialog>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Thresholds">
                        <text-wrap-dialog :fulltext="selectedDetailsUltimateSourceSystem ? selectedDetailsUltimateSourceSystem.THRESHOLD : ''"></text-wrap-dialog>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Golden Source (Yes/No)">
                        <text-wrap-dialog :fulltext="selectedDetailsUltimateSourceSystem ? selectedDetailsUltimateSourceSystem.GOLDEN_SOURCE : ''"></text-wrap-dialog>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Golden Source System Name">
                        <text-wrap-dialog :fulltext="selectedDetailsUltimateSourceSystem ? selectedDetailsUltimateSourceSystem.GS_SYSTEM_NAME : ''"></text-wrap-dialog>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Golden Source ITAM ID">
                        <text-wrap-dialog :fulltext="selectedDetailsUltimateSourceSystem ? selectedDetailsUltimateSourceSystem.GF_ITAM_ID : ''"></text-wrap-dialog>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Golden Source Table Name">
                        <text-wrap-dialog :fulltext="selectedDetailsUltimateSourceSystem ? selectedDetailsUltimateSourceSystem.GS_TABLE_NAME : ''"></text-wrap-dialog>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Golden Source Column Name">
                        <text-wrap-dialog :fulltext="selectedDetailsUltimateSourceSystem ? selectedDetailsUltimateSourceSystem.GS_COLUMN_NAME : ''"></text-wrap-dialog>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Golden Source Screen Name">
                        <text-wrap-dialog :fulltext="selectedDetailsUltimateSourceSystem ? selectedDetailsUltimateSourceSystem.GS_DATA_ELEMENT : ''"></text-wrap-dialog>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Golden Source Business Description">
                        <text-wrap-dialog :fulltext="selectedDetailsUltimateSourceSystem ? selectedDetailsUltimateSourceSystem.GS_DESCRIPTION : ''"></text-wrap-dialog>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Golden Source Derived (Yes/No)">
                        <text-wrap-dialog :fulltext="selectedDetailsUltimateSourceSystem ? selectedDetailsUltimateSourceSystem.GS_DERIVED : ''"></text-wrap-dialog>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Golden Source Derivation Logic">
                        <text-wrap-dialog :fulltext="selectedDetailsUltimateSourceSystem ? selectedDetailsUltimateSourceSystem.GS_DERIVATION_LOGIC : ''"></text-wrap-dialog>
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
                        <text-wrap-dialog :fulltext="selectedDetailsDomainView ? selectedDetailsDomainView.DOMAIN_NAME : ''"></text-wrap-dialog>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Sub Domain">
                        <text-wrap-dialog :fulltext="selectedDetailsDomainView ? selectedDetailsDomainView.SUBDOMAIN_NAME : ''"></text-wrap-dialog>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Domain Owner">
                        <text-wrap-dialog :fulltext="selectedDetailsDomainView ? selectedDetailsDomainView.SUBDOMAIN_OWNER : ''"></text-wrap-dialog>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Business Term">
                        <text-wrap-dialog :fulltext="selectedDetailsDomainView ? selectedDetailsDomainView.BUSINESS_TERM : ''"></text-wrap-dialog>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Business Term Description">
                        <text-wrap-dialog :fulltext="selectedDetailsDomainView ? selectedDetailsDomainView.BUSINESS_TERM_DESCRIPTION : ''"></text-wrap-dialog>
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
                        <text-wrap-dialog :fulltext="selectedDetailsDataStandards ? selectedDetailsDataStandards.DPO_DQ_STANDARDS : ''"></text-wrap-dialog>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="DQ Standards set at Business Term Level">
                        <text-wrap-dialog :fulltext="selectedDetailsDataStandards ? selectedDetailsDataStandards.DQ_STANDARDS_BT_LEVEL : ''"></text-wrap-dialog>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Thresholds defined by DPO">
                        <text-wrap-dialog :fulltext="selectedDetailsDataStandards ? selectedDetailsDataStandards.DPO_THRESHOLD : ''"></text-wrap-dialog>
                      </b-form-group>

                      <b-form-group horizontal :label-cols="4" breakpoint="md" label="Thresholds set at Business term level">
                        <text-wrap-dialog :fulltext="selectedDetailsDataStandards ? selectedDetailsDataStandards.THRESHOLD_BT_LEVEL : ''"></text-wrap-dialog>
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
      selectedDetailsImmediatePrecedingSystem: null,
      selectedDetailsUltimateSourceSystem: null,
      selectedDetailsDomainView: null,
      selectedDetailsDataStandards: null,
      ddImmSystemNameSelected: null,
      ddImmItamIDSelected: null,
      ddImmTableNameSelected: null,
      ddImmColumnNameSelected: null,
      ddImmScreenLabelSelected: null,
      ddUltSystemNameSelected: null,
      ddUltItamIDSelected: null,
      ddUltTableNameSelected: null,
      ddUltColumnNameSelected: null,
      ddUltScreenLabelSelected: null,
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
      return _.uniq(_.map(this.store.DDSourceImmediatePrecedingSystem, (v) => v.SYSTEM_NAME.toString().trim())).filter(Boolean);
    },
    ddImmItamIDOptions () {
      var self = this;
      var filtered = _.filter(self.store.DDSourceImmediatePrecedingSystem, (v) => {
        return v.SYSTEM_NAME == self.ddImmSystemNameSelected;
      });

      return _.uniq(_.map(filtered, (v) => v.ITAM_ID.toString().trim())).filter(Boolean);
    },
    ddImmTableNameOptions () {
      var self = this;
      var filtered = _.filter(self.store.DDSourceImmediatePrecedingSystem, (v) => {
        return v.SYSTEM_NAME == self.ddImmSystemNameSelected 
          && v.ITAM_ID == self.ddImmItamIDSelected
      });
      
      return _.uniq(_.map(filtered, (v) => v.TABLE_NAME.toString().trim())).filter(Boolean);
    },
    ddImmColumnNameOptions () {
      var self = this;
      var filtered = _.filter(self.store.DDSourceImmediatePrecedingSystem, (v) => {
        return v.SYSTEM_NAME == self.ddImmSystemNameSelected 
          && v.ITAM_ID == self.ddImmItamIDSelected
          && v.TABLE_NAME == self.ddImmTableNameSelected
      });
      
      return _.uniq(_.map(filtered, (v) => v.COLUMN_NAME.toString().trim())).filter(Boolean);
    },
    ddImmScreenLabelOptions () {
      var self = this;
      var filtered = _.filter(self.store.DDSourceImmediatePrecedingSystem, (v) => {
        return v.SYSTEM_NAME == self.ddImmSystemNameSelected 
          && v.ITAM_ID == self.ddImmItamIDSelected
          && v.TABLE_NAME == self.ddImmTableNameSelected
          && v.COLUMN_NAME == self.ddImmColumnNameSelected
      });
      
      return _.uniq(_.map(filtered, (v) => v.DATA_ELEMENT.toString().trim())).filter(Boolean);
    },
    ddUltSystemNameOptions () {
      return _.uniq(_.map(this.store.DDSourceUltimateSourceSystem, (v) => v.SYSTEM_NAME.toString().trim())).filter(Boolean);
    },
    ddUltItamIDOptions () {
      var self = this;
      var filtered = _.filter(self.store.DDSourceUltimateSourceSystem, (v) => {
        return v.SYSTEM_NAME == self.ddUltSystemNameSelected;
      });

      return _.uniq(_.map(filtered, (v) => v.ITAM_ID.toString().trim())).filter(Boolean);
    },
    ddUltTableNameOptions () {
      var self = this;
      var filtered = _.filter(self.store.DDSourceUltimateSourceSystem, (v) => {
        return v.SYSTEM_NAME == self.ddUltSystemNameSelected 
          && v.ITAM_ID == self.ddUltItamIDSelected
      });
      
      return _.uniq(_.map(filtered, (v) => v.TABLE_NAME.toString().trim())).filter(Boolean);
    },
    ddUltColumnNameOptions () {
      var self = this;
      var filtered = _.filter(self.store.DDSourceUltimateSourceSystem, (v) => {
        return v.SYSTEM_NAME == self.ddUltSystemNameSelected 
          && v.ITAM_ID == self.ddUltItamIDSelected
          && v.TABLE_NAME == self.ddUltTableNameSelected
      });
      
      return _.uniq(_.map(filtered, (v) => v.COLUMN_NAME.toString().trim())).filter(Boolean);
    },
    ddUltScreenLabelOptions () {
      var self = this;
      var filtered = _.filter(self.store.DDSourceUltimateSourceSystem, (v) => {
        return v.SYSTEM_NAME == self.ddUltSystemNameSelected 
          && v.ITAM_ID == self.ddUltItamIDSelected
          && v.TABLE_NAME == self.ddUltTableNameSelected
          && v.COLUMN_NAME == self.ddUltColumnNameSelected
      });
      
      return _.uniq(_.map(filtered, (v) => v.DATA_ELEMENT.toString().trim())).filter(Boolean);
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
    ddImmSystemNameSelected () {
      if(this.firstload) return;

      if (!this.firstload) {
        var param = {};

        if (this.ddImmSystemNameSelected && this.ddImmSystemNameSelected != "NA") {
          param.ImmSystemName = this.ddImmSystemNameSelected.toString();
        }

        this.runGetDetails(param);
      }
    },
    ddImmItamIDSelected () {
      if(this.firstload) return;
      
      if (!this.firstload) {
        var param = {};

        if (this.ddImmSystemNameSelected && this.ddImmSystemNameSelected != "NA") {
          param.ImmSystemName = this.ddImmSystemNameSelected.toString();
        }
        if (this.ddImmItamIDSelected && this.ddImmItamIDSelected != "NA") {
          param.ImmItamID = this.ddImmItamIDSelected.toString();
        }

        this.runGetDetails(param);
      }
    },
    ddImmTableNameSelected () {
      if(this.firstload) return;
      
      if (!this.firstload) {
        var param = {};

        if (this.ddImmSystemNameSelected && this.ddImmSystemNameSelected != "NA") {
          param.ImmSystemName = this.ddImmSystemNameSelected.toString();
        }
        if (this.ddImmItamIDSelected && this.ddImmItamIDSelected != "NA") {
          param.ImmItamID = this.ddImmItamIDSelected.toString();
        }
        if (this.ddImmTableNameSelected && this.ddImmTableNameSelected != "NA") {
          param.ImmTableName = this.ddImmTableNameSelected.toString();
        }

        this.runGetDetails(param);
      }
    },
    ddImmColumnNameSelected () {
      if(this.firstload) return;
      
      if (!this.firstload) {
        var param = {};

        if (this.ddImmSystemNameSelected && this.ddImmSystemNameSelected != "NA") {
          param.ImmSystemName = this.ddImmSystemNameSelected.toString();
        }
        if (this.ddImmItamIDSelected && this.ddImmItamIDSelected != "NA") {
          param.ImmItamID = this.ddImmItamIDSelected.toString();
        }
        if (this.ddImmTableNameSelected && this.ddImmTableNameSelected != "NA") {
          param.ImmTableName = this.ddImmTableNameSelected.toString();
        }
        if (this.ddImmColumnNameSelected && this.ddImmColumnNameSelected != "NA") {
          param.ImmColumnName = this.ddImmColumnNameSelected.toString();
        }

        this.runGetDetails(param);
      }
    },
    ddImmScreenLabelSelected () {
      if(this.firstload) return;
      
      if (!this.firstload) {
        var param = {};

        if (this.ddImmSystemNameSelected && this.ddImmSystemNameSelected != "NA") {
          param.ImmSystemName = this.ddImmSystemNameSelected.toString();
        }
        if (this.ddImmItamIDSelected && this.ddImmItamIDSelected != "NA") {
          param.ImmItamID = this.ddImmItamIDSelected.toString();
        }
        if (this.ddImmTableNameSelected && this.ddImmTableNameSelected != "NA") {
          param.ImmTableName = this.ddImmTableNameSelected.toString();
        }
        if (this.ddImmColumnNameSelected && this.ddImmColumnNameSelected != "NA") {
          param.ImmColumnName = this.ddImmColumnNameSelected.toString();
        }
        if (this.ddImmScreenLabelSelected && this.ddImmScreenLabelSelected != "NA") {
          param.ImmScreenLabel = this.ddImmScreenLabelSelected.toString();
        }

        this.runGetDetails(param);
      }
    },
    ddUltSystemNameSelected () {
      if(this.firstload) return;
      
      if (!this.firstload) {
        var param = {};

        if (this.ddImmSystemNameSelected && this.ddImmSystemNameSelected != "NA") {
          param.ImmSystemName = this.ddImmSystemNameSelected.toString();
        }
        if (this.ddImmItamIDSelected && this.ddImmItamIDSelected != "NA") {
          param.ImmItamID = this.ddImmItamIDSelected.toString();
        }
        if (this.ddImmTableNameSelected && this.ddImmTableNameSelected != "NA") {
          param.ImmTableName = this.ddImmTableNameSelected.toString();
        }
        if (this.ddImmColumnNameSelected && this.ddImmColumnNameSelected != "NA") {
          param.ImmColumnName = this.ddImmColumnNameSelected.toString();
        }
        if (this.ddImmScreenLabelSelected && this.ddImmScreenLabelSelected != "NA") {
          param.ImmScreenLabel = this.ddImmScreenLabelSelected.toString();
        }
        if (this.ddUltSystemNameSelected && this.ddUltSystemNameSelected != "NA") {
          param.UltSystemName = this.ddUltSystemNameSelected.toString();
        }

        this.runGetDetails(param);
      }
    },
    ddUltItamIDSelected () {
      if(this.firstload) return;
      
      if (!this.firstload) {
        var param = {};

        if (this.ddImmSystemNameSelected && this.ddImmSystemNameSelected != "NA") {
          param.ImmSystemName = this.ddImmSystemNameSelected.toString();
        }
        if (this.ddImmItamIDSelected && this.ddImmItamIDSelected != "NA") {
          param.ImmItamID = this.ddImmItamIDSelected.toString();
        }
        if (this.ddImmTableNameSelected && this.ddImmTableNameSelected != "NA") {
          param.ImmTableName = this.ddImmTableNameSelected.toString();
        }
        if (this.ddImmColumnNameSelected && this.ddImmColumnNameSelected != "NA") {
          param.ImmColumnName = this.ddImmColumnNameSelected.toString();
        }
        if (this.ddImmScreenLabelSelected && this.ddImmScreenLabelSelected != "NA") {
          param.ImmScreenLabel = this.ddImmScreenLabelSelected.toString();
        }
        if (this.ddUltSystemNameSelected && this.ddUltSystemNameSelected != "NA") {
          param.UltSystemName = this.ddUltSystemNameSelected.toString();
        }
        if (this.ddUltItamIDSelected && this.ddUltItamIDSelected != "NA") {
          param.UltItamID = this.ddUltItamIDSelected.toString();
        }

        this.runGetDetails(param);
      }
    },
    ddUltTableNameSelected () {
      if(this.firstload) return;
      
      if (!this.firstload) {
        var param = {};

        if (this.ddImmSystemNameSelected && this.ddImmSystemNameSelected != "NA") {
          param.ImmSystemName = this.ddImmSystemNameSelected.toString();
        }
        if (this.ddImmItamIDSelected && this.ddImmItamIDSelected != "NA") {
          param.ImmItamID = this.ddImmItamIDSelected.toString();
        }
        if (this.ddImmTableNameSelected && this.ddImmTableNameSelected != "NA") {
          param.ImmTableName = this.ddImmTableNameSelected.toString();
        }
        if (this.ddImmColumnNameSelected && this.ddImmColumnNameSelected != "NA") {
          param.ImmColumnName = this.ddImmColumnNameSelected.toString();
        }
        if (this.ddImmScreenLabelSelected && this.ddImmScreenLabelSelected != "NA") {
          param.ImmScreenLabel = this.ddImmScreenLabelSelected.toString();
        }
        if (this.ddUltSystemNameSelected && this.ddUltSystemNameSelected != "NA") {
          param.UltSystemName = this.ddUltSystemNameSelected.toString();
        }
        if (this.ddUltItamIDSelected && this.ddUltItamIDSelected != "NA") {
          param.UltItamID = this.ddUltItamIDSelected.toString();
        }
        if (this.ddUltTableNameSelected && this.ddUltTableNameSelected != "NA") {
          param.UltTableName = this.ddUltTableNameSelected.toString();
        }

        this.runGetDetails(param);
      }
    },
    ddUltColumnNameSelected () {
      if(this.firstload) return;
      
      if (!this.firstload) {
        var param = {};

        if (this.ddImmSystemNameSelected && this.ddImmSystemNameSelected != "NA") {
          param.ImmSystemName = this.ddImmSystemNameSelected.toString();
        }
        if (this.ddImmItamIDSelected && this.ddImmItamIDSelected != "NA") {
          param.ImmItamID = this.ddImmItamIDSelected.toString();
        }
        if (this.ddImmTableNameSelected && this.ddImmTableNameSelected != "NA") {
          param.ImmTableName = this.ddImmTableNameSelected.toString();
        }
        if (this.ddImmColumnNameSelected && this.ddImmColumnNameSelected != "NA") {
          param.ImmColumnName = this.ddImmColumnNameSelected.toString();
        }
        if (this.ddImmScreenLabelSelected && this.ddImmScreenLabelSelected != "NA") {
          param.ImmScreenLabel = this.ddImmScreenLabelSelected.toString();
        }
        if (this.ddUltSystemNameSelected && this.ddUltSystemNameSelected != "NA") {
          param.UltSystemName = this.ddUltSystemNameSelected.toString();
        }
        if (this.ddUltItamIDSelected && this.ddUltItamIDSelected != "NA") {
          param.UltItamID = this.ddUltItamIDSelected.toString();
        }
        if (this.ddUltTableNameSelected && this.ddUltTableNameSelected != "NA") {
          param.UltTableName = this.ddUltTableNameSelected.toString();
        }
        if (this.ddUltColumnNameSelected && this.ddUltColumnNameSelected != "NA") {
          param.UltColumnName = this.ddUltColumnNameSelected.toString();
        }

        this.runGetDetails(param);
      }
    },
    ddUltScreenLabelSelected (){
      if(this.firstload) return;

      if (!this.firstload) {
        var param = {};

        if (this.ddImmSystemNameSelected && this.ddImmSystemNameSelected != "NA") {
          param.ImmSystemName = this.ddImmSystemNameSelected.toString();
        }
        if (this.ddImmItamIDSelected && this.ddImmItamIDSelected != "NA") {
          param.ImmItamID = this.ddImmItamIDSelected.toString();
        }
        if (this.ddImmTableNameSelected && this.ddImmTableNameSelected != "NA") {
          param.ImmTableName = this.ddImmTableNameSelected.toString();
        }
        if (this.ddImmColumnNameSelected && this.ddImmColumnNameSelected != "NA") {
          param.ImmColumnName = this.ddImmColumnNameSelected.toString();
        }
        if (this.ddImmScreenLabelSelected && this.ddImmScreenLabelSelected != "NA") {
          param.ImmScreenLabel = this.ddImmScreenLabelSelected.toString();
        }
        if (this.ddUltSystemNameSelected && this.ddUltSystemNameSelected != "NA") {
          param.UltSystemName = this.ddUltSystemNameSelected.toString();
        }
        if (this.ddUltItamIDSelected && this.ddUltItamIDSelected != "NA") {
          param.UltItamID = this.ddUltItamIDSelected.toString();
        }
        if (this.ddUltTableNameSelected && this.ddUltTableNameSelected != "NA") {
          param.UltTableName = this.ddUltTableNameSelected.toString();
        }
        if (this.ddUltColumnNameSelected && this.ddUltColumnNameSelected != "NA") {
          param.UltColumnName = this.ddUltColumnNameSelected.toString();
        }
        if (this.ddUltScreenLabelSelected && this.ddUltScreenLabelSelected != "NA") {
          param.UltScreenLabel = this.ddUltScreenLabelSelected.toString();
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

      return self.getDetails(param).then(
        res => {
          this.firstload = true;

          // if (self.store.detailsSource.length > 0){
          //   var tmp = self.store.detailsSource[0].Values[0];

          //   self.selectedDetails = {}
          //   _.each(Object.keys(tmp), function(v, i){
          //       self.selectedDetails[v] = _.uniq(
          //         _.map(self.store.detailsSource[0].Values, (val) => val[v] ? val[v].toString().trim() : "").filter(Boolean)
          //       );

          //       self.selectedDetails[v] = self.selectedDetails[v].join(", ");
          //   });

          //   // make falsy NA
          //   Object.keys(self.selectedDetails).forEach((val) => {
          //     self.selectedDetails[val] = !!self.selectedDetails[val].trim() ? self.selectedDetails[val] : "NA";
          //   });

          //   self.store.DDSource.map(function(v){
          //     Object.keys(v).forEach(function(key){
          //       v[key] = v[key] ? v[key] : "NA"
          //     })

          //     return v
          //   })
            
          //   setTimeout(() => {
          //     self.ddImmSystemNameSelected = self.selectedDetails.SYSTEM_NAME;
          //     self.ddImmItamIDSelected = self.selectedDetails.ITAM_ID;
          //     self.ddImmTableNameSelected = self.selectedDetails.TABLE_NAME;
          //     self.ddImmColumnNameSelected = self.selectedDetails.COLUMN_NAME;
          //     self.ddImmScreenLabelSelected = self.selectedDetails.DATA_ELEMENT;
          //     self.ddUltSystemNameSelected = self.selectedDetails.ULT_SYSTEM;
          //     self.ddUltItamIDSelected = self.selectedDetails.ULT_ITAM_ID;
          //     self.ddUltTableNameSelected = self.selectedDetails.ULT_TABLE_NAME;
          //     self.ddUltColumnNameSelected = self.selectedDetails.ULT_COLUMN_NAME;
          //     self.ddUltScreenLabelSelected = self.selectedDetails.ULT_SCREEN_LABEL_NAME;

          //     setTimeout(() => {
          //       this.firstload = false;
          //     }, 100);
          //   }, 100);
          // } else {
          //   this.selectedDetails = null;
          // }

          if (self.store.DetailsImmediatePrecedingSystem.length > 0){
            var tmp = self.store.DetailsImmediatePrecedingSystem[0].Values[0];

            self.selectedDetailsImmediatePrecedingSystem = {}
            _.each(Object.keys(tmp), function(v, i){
                self.selectedDetailsImmediatePrecedingSystem[v] = _.uniq(
                  _.map(self.store.DetailsImmediatePrecedingSystem[0].Values, (val) => val[v] ? val[v].toString().trim() : "").filter(Boolean)
                );

                self.selectedDetailsImmediatePrecedingSystem[v] = self.selectedDetailsImmediatePrecedingSystem[v].join(", ");
            });

            // make falsy NA
            Object.keys(self.selectedDetailsImmediatePrecedingSystem).forEach((val) => {
              self.selectedDetailsImmediatePrecedingSystem[val] = !!self.selectedDetailsImmediatePrecedingSystem[val].trim() ? self.selectedDetailsImmediatePrecedingSystem[val] : "NA";
            });

            self.store.DDSourceImmediatePrecedingSystem.map(function(v){
              Object.keys(v).forEach(function(key){
                v[key] = v[key] ? v[key] : "NA"
              })

              return v
            })
            
            setTimeout(() => {
              self.ddImmSystemNameSelected = self.selectedDetailsImmediatePrecedingSystem.SYSTEM_NAME;
              self.ddImmItamIDSelected = self.selectedDetailsImmediatePrecedingSystem.ITAM_ID;
              self.ddImmTableNameSelected = self.selectedDetailsImmediatePrecedingSystem.TABLE_NAME;
              self.ddImmColumnNameSelected = self.selectedDetailsImmediatePrecedingSystem.COLUMN_NAME;
              self.ddImmScreenLabelSelected = self.selectedDetailsImmediatePrecedingSystem.DATA_ELEMENT;

              setTimeout(() => {
                this.firstload = false;
              }, 100);
            }, 100);
          } else {
            this.selectedDetailsImmediatePrecedingSystem = null;
          }

          if (self.store.DetailsUltimateSourceSystem.length > 0){
            var tmp = self.store.DetailsUltimateSourceSystem[0].Values[0];

            self.selectedDetailsUltimateSourceSystem = {}
            _.each(Object.keys(tmp), function(v, i){
                self.selectedDetailsUltimateSourceSystem[v] = _.uniq(
                  _.map(self.store.DetailsUltimateSourceSystem[0].Values, (val) => val[v] ? val[v].toString().trim() : "").filter(Boolean)
                );

                self.selectedDetailsUltimateSourceSystem[v] = self.selectedDetailsUltimateSourceSystem[v].join(", ");
            });

            // make falsy NA
            Object.keys(self.selectedDetailsUltimateSourceSystem).forEach((val) => {
              self.selectedDetailsUltimateSourceSystem[val] = !!self.selectedDetailsUltimateSourceSystem[val].trim() ? self.selectedDetailsUltimateSourceSystem[val] : "NA";
            });

            self.store.DDSourceImmediatePrecedingSystem.map(function(v){
              Object.keys(v).forEach(function(key){
                v[key] = v[key] ? v[key] : "NA"
              })

              return v
            })
            
            setTimeout(() => {
              self.ddUltSystemNameSelected = self.selectedDetailsUltimateSourceSystem.SYSTEM_NAME;
              self.ddUltItamIDSelected = self.selectedDetailsUltimateSourceSystem.ITAM_ID;
              self.ddUltTableNameSelected = self.selectedDetailsUltimateSourceSystem.TABLE_NAME;
              self.ddUltColumnNameSelected = self.selectedDetailsUltimateSourceSystem.COLUMN_NAME;
              self.ddUltScreenLabelSelected = self.selectedDetailsUltimateSourceSystem.DATA_ELEMENT;

              setTimeout(() => {
                this.firstload = false;
              }, 100);
            }, 100);
          } else {
            this.selectedDetailsUltimateSourceSystem = null;
          }

          if (self.store.DetailsDomainView.length > 0){
            var tmp = self.store.DetailsDomainView[0].Values[0];

            self.selectedDetailsDomainView = {}
            _.each(Object.keys(tmp), function(v, i){
                self.selectedDetailsDomainView[v] = _.uniq(
                  _.map(self.store.DetailsDomainView[0].Values, (val) => val[v] ? val[v].toString().trim() : "").filter(Boolean)
                );
                
                self.selectedDetailsDomainView[v] = self.selectedDetailsDomainView[v].join(", ");
            });

            // make falsy NA
            Object.keys(self.selectedDetailsDomainView).forEach((val) => {
              self.selectedDetailsDomainView[val] = !!self.selectedDetailsDomainView[val].trim() ? self.selectedDetailsDomainView[val] : "NA";
            });

            self.store.DDSourceDomainView.map(function(v){
              Object.keys(v).forEach(function(key){
                v[key] = v[key] ? v[key] : "NA"
              })

              return v
            })
          } else {
            this.selectedDetailsDomainView = null;
          }

          if (self.store.DetailsDataStandards.length > 0){
            var tmp = self.store.DetailsDataStandards[0].Values[0];

            self.selectedDetailsDataStandards = {}
            _.each(Object.keys(tmp), function(v, i){
                self.selectedDetailsDataStandards[v] = _.uniq(
                  _.map(self.store.DetailsDataStandards[0].Values, (val) => val[v] ? val[v].toString().trim() : "").filter(Boolean)
                );
                
                self.selectedDetailsDataStandards[v] = self.selectedDetailsDataStandards[v].join(", ");
            });

            // make falsy NA
            Object.keys(self.selectedDetailsDataStandards).forEach((val) => {
              self.selectedDetailsDataStandards[val] = !!self.selectedDetailsDataStandards[val].trim() ? self.selectedDetailsDataStandards[val] : "NA";
            });

            self.store.DDSourceDataStandards.map(function(v){
              Object.keys(v).forEach(function(key){
                v[key] = v[key] ? v[key] : "NA"
              })

              return v
            })
          } else {
            this.selectedDetailsDataStandards = null;
          }
        },
        err => err
      );
    }
  },
}
</script>