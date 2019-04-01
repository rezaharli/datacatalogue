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
    <page-loader v-if="dpomy.detailsLoading" />
    
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
              <text-wrap-dialog :fulltext="selectedDetails ? selectedDetails.DOWNSTREAM_PROCESS: ''"></text-wrap-dialog>
            </b-media>
            
            <b-media class="left-card-media">
              <h6 class="left-card-title">Process Owner</h6>
              <text-wrap-dialog :fulltext="selectedDetails ? selectedDetails.PROCESS_OWNER: ''"></text-wrap-dialog>
            </b-media>
            
            <b-media class="left-card-media">
              <h6 class="left-card-title">Bank ID</h6>
              <text-wrap-dialog :fulltext="selectedDetails ? selectedDetails.BANK_ID: ''"></text-wrap-dialog>
            </b-media>
            
            <b-media class="left-card-media">
              <h6 class="left-card-title">CDE Name</h6>
              <text-wrap-dialog :fulltext="selectedDetails ? selectedDetails.CDE_NAME: ''"></text-wrap-dialog>
            </b-media>
            
            <b-media class="left-card-media">
              <h6 class="left-card-title">CDE Rationale</h6>
              <text-wrap-dialog :fulltext="selectedDetails ? selectedDetails.CDE_RATIONALE: ''"></text-wrap-dialog>
            </b-media>
          </b-card>
        </b-col>

        <b-col cols="8"> 
          <b-row>
            <b-col>
              <b-card title="Immediate Preceding System Details" tag="article" class="mb-2">
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

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Screen Label Name" label-for="screenlabelname">
                      <b-form-select id="screenlabelname" class="col-8" v-model="ddImmScreenLabelSelected" :options="ddImmScreenLabelOptions"></b-form-select>
                    </b-form-group>
                    
                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Business Description">
                      <text-wrap-dialog :fulltext="selectedDetails ? selectedDetails.IMM_BUSINESS_DESCRIPTION : ''"></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Derived (Yes/No)">
                      <text-wrap-dialog :fulltext="selectedDetails ? selectedDetails.IMM_DERIVED_YES_NO : ''"></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Derivation Logic">
                      <text-wrap-dialog :fulltext="selectedDetails ? selectedDetails.IMM_DERIVATION_LOGIC : ''"></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Data Quality Requirements">
                      <text-wrap-dialog :fulltext="selectedDetails ? selectedDetails.IMM_DQ_REQUIREMENTS : ''"></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Thresholds">
                      <text-wrap-dialog :fulltext="selectedDetails ? selectedDetails.IMM_THRESHOLD : ''"></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Data SLA signed?">
                      <text-wrap-dialog :fulltext="selectedDetails ? selectedDetails.IMM_DATA_SLA_SIGNED : ''"></text-wrap-dialog>
                    </b-form-group>
                  </b-form>
                </p>
              </b-card>
            </b-col>
          </b-row>

          <b-row>
            <b-col>
              <b-card title="Ultimate Source System Details" tag="article" class="mb-2">
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

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Screen Label Name" label-for="screenlabelname">
                      <b-form-select id="screenlabelname" class="col-8" v-model="ddUltScreenLabelSelected" :options="ddUltScreenLabelOptions"></b-form-select>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Business Description">
                      <text-wrap-dialog :fulltext="selectedDetails ? selectedDetails.ULT_BUSINESS_DESCRIPTION : ''"></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Derived?">
                      <text-wrap-dialog :fulltext="selectedDetails ? selectedDetails.ULT_DERIVED_YES_NO : ''"></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Derivation Logic">
                      <text-wrap-dialog :fulltext="selectedDetails ? selectedDetails.ULT_DERIVATION_LOGIC : ''"></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Data Quality requirements">
                      <text-wrap-dialog :fulltext="selectedDetails ? selectedDetails.ULT_DQ_REQUIREMENTS : ''"></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Thresholds">
                      <text-wrap-dialog :fulltext="selectedDetails ? selectedDetails.ULT_THRESHOLD : ''"></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Golden Source (Yes/No)">
                      <text-wrap-dialog :fulltext="selectedDetails ? selectedDetails.GOLDEN_SOURCE : ''"></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Golden Source System Name">
                      <text-wrap-dialog :fulltext="selectedDetails ? selectedDetails.GS_SYSTEM_NAME : ''"></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Golden Source ITAM ID">
                      <text-wrap-dialog :fulltext="selectedDetails ? selectedDetails.GS_ITAM_ID : ''"></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Golden Source Table Name">
                      <text-wrap-dialog :fulltext="selectedDetails ? selectedDetails.GS_TABLE_NAME : ''"></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Golden Source Column Name">
                      <text-wrap-dialog :fulltext="selectedDetails ? selectedDetails.GS_COLUMN_NAME : ''"></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Golden Source Screen Name">
                      <text-wrap-dialog :fulltext="selectedDetails ? selectedDetails.GS_SCREEN_NAME : ''"></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Golden Source Business Description">
                      <text-wrap-dialog :fulltext="selectedDetails ? selectedDetails.GS_BUSINESS_DESC : ''"></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Golden Source Derived (Yes/No)">
                      <text-wrap-dialog :fulltext="selectedDetails ? selectedDetails.GS_DERIVED_YES_NO : ''"></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Golden Source Derivation Logic">
                      <text-wrap-dialog :fulltext="selectedDetails ? selectedDetails.GS_DERIVATION_LOGIC : ''"></text-wrap-dialog>
                    </b-form-group>
                  </b-form>
                </p>
              </b-card>
            </b-col>
          </b-row>

          <b-row>
            <b-col>
              <b-card title="Domain View" tag="article" class="mb-2">
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

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Business Term">
                      <text-wrap-dialog :fulltext="selectedDetails ? selectedDetails.BUSINESS_TERM : ''"></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Business Term Description">
                      <text-wrap-dialog :fulltext="selectedDetails ? selectedDetails.BUSINESS_TERM_DESCRIPTION : ''"></text-wrap-dialog>
                    </b-form-group>
                  </b-form>
                </p>
              </b-card>
            </b-col>
          </b-row>

          <b-row>
            <b-col>
              <b-card title="Data Standards" tag="article" class="mb-2">
                <p class="card-text">
                  <b-form>
                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="DQ Standards set by DPO">
                      <text-wrap-dialog :fulltext="selectedDetails ? selectedDetails.DPO_DQ_STANDARDS : ''"></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="DQ Standards set at Business Term Level">
                      <text-wrap-dialog :fulltext="selectedDetails ? selectedDetails.DQ_STANDARDS_BT_LEVEL : ''"></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Thresholds defined by DPO*">
                      <text-wrap-dialog :fulltext="selectedDetails ? selectedDetails.DPO_THRESHOLD : ''"></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Thresholds set at Business term level">
                      <text-wrap-dialog :fulltext="selectedDetails ? selectedDetails.THRESHOLD_BT_LEVEL : ''"></text-wrap-dialog>
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
        'System Name': 'selectedDetails.IMM_SYSTEM',
        'ITAM ID': 'selectedDetails.IMM_ITAM_ID',
        'TableName': 'selectedDetails.IMM_TABLE_NAME',
        "Column Name": "selectedDetails.IMM_COLUMN_NAME",
        "Screen Label Name": "selectedDetails.IMM_SCREEN_LABEL_NAME",
        "Business Description": "selectedDetails.IMM_BUSINESS_DESCRIPTION",
        "Derived (Yes/No)": "selectedDetails.IMM_DERIVED_YES_NO",
        "Derivation Logic": "selectedDetails.IMM_DERIVATION_LOGIC",
        "Data Quality Requirements": "selectedDetails.IMM_DQ_REQUIREMENTS",
        "Thresholds": "selectedDetails.IMM_THRESHOLD",
        "Data SLA signed?": "selectedDetails.IMM_DATA_SLA_SIGNED",
        'System Name ': 'selectedDetails.ULT_SYSTEM',
        'ITAM ID ': 'selectedDetails.ULT_ITAM_ID',
        'TableName ': 'selectedDetails.ULT_TABLE_NAME',
        "Column Name ": "selectedDetails.ULT_COLUMN_NAME",
        "Screen Label Name ": "selectedDetails.ULT_SCREEN_LABEL_NAME",
        "Business Description ": "selectedDetails.ULT_BUSINESS_DESCRIPTION",
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
        "Thresholds defined by DPO*": "selectedDetails.DPO_THRESHOLD"
      }
    }
  },
  computed: {
    ...mapState({
      dpomy: state => state.dpomy.all
    }),
    ddImmSystemNameOptions () {
      return _.uniq(_.map(this.dpomy.DDSource, (v) => v.IMM_SYSTEM.toString().trim())).filter(Boolean);
    },
    ddImmItamIDOptions () {
      var self = this;
      var filtered = _.filter(self.dpomy.DDSource, (v) => {
        return v.IMM_SYSTEM == self.ddImmSystemNameSelected;
      });

      return _.uniq(_.map(filtered, (v) => v.IMM_ITAM_ID.toString().trim())).filter(Boolean);
    },
    ddImmTableNameOptions () {
      var self = this;
      var filtered = _.filter(self.dpomy.DDSource, (v) => {
        return v.IMM_SYSTEM == self.ddImmSystemNameSelected 
          && v.IMM_ITAM_ID == self.ddImmItamIDSelected
      });
      
      return _.uniq(_.map(filtered, (v) => v.IMM_TABLE_NAME.toString().trim())).filter(Boolean);
    },
    ddImmColumnNameOptions () {
      var self = this;
      var filtered = _.filter(self.dpomy.DDSource, (v) => {
        return v.IMM_SYSTEM == self.ddImmSystemNameSelected 
          && v.IMM_ITAM_ID == self.ddImmItamIDSelected
          && v.IMM_TABLE_NAME == self.ddImmTableNameSelected
      });
      
      return _.uniq(_.map(filtered, (v) => v.IMM_COLUMN_NAME.toString().trim())).filter(Boolean);
    },
    ddImmScreenLabelOptions () {
      var self = this;
      var filtered = _.filter(self.dpomy.DDSource, (v) => {
        return v.IMM_SYSTEM == self.ddImmSystemNameSelected 
          && v.IMM_ITAM_ID == self.ddImmItamIDSelected
          && v.IMM_TABLE_NAME == self.ddImmTableNameSelected
          && v.IMM_COLUMN_NAME == self.ddImmColumnNameSelected
      });
      
      return _.uniq(_.map(filtered, (v) => v.IMM_SCREEN_LABEL_NAME.toString().trim())).filter(Boolean);
    },
    ddUltSystemNameOptions () {
      var self = this;
      var filtered = _.filter(self.dpomy.DDSource, (v) => {
        return v.IMM_SYSTEM == self.ddImmSystemNameSelected 
          && v.IMM_ITAM_ID == self.ddImmItamIDSelected
          && v.IMM_TABLE_NAME == self.ddImmTableNameSelected
          && v.IMM_COLUMN_NAME == self.ddImmColumnNameSelected
          && v.IMM_SCREEN_LABEL_NAME == self.ddImmScreenLabelSelected
      });
      
      return _.uniq(_.map(filtered, (v) => v.ULT_SYSTEM.toString().trim())).filter(Boolean);
    },
    ddUltItamIDOptions () {
      var self = this;
      var filtered = _.filter(self.dpomy.DDSource, (v) => {
        return v.IMM_SYSTEM == self.ddImmSystemNameSelected 
          && v.IMM_ITAM_ID == self.ddImmItamIDSelected
          && v.IMM_TABLE_NAME == self.ddImmTableNameSelected
          && v.IMM_COLUMN_NAME == self.ddImmColumnNameSelected
          && v.IMM_SCREEN_LABEL_NAME == self.ddImmScreenLabelSelected
          && v.ULT_SYSTEM == self.ddUltSystemNameSelected
      });
      
      return _.uniq(_.map(filtered, (v) => v.ULT_ITAM_ID.toString().trim())).filter(Boolean);
    },
    ddUltTableNameOptions () {
      var self = this;
      var filtered = _.filter(self.dpomy.DDSource, (v) => {
        return v.IMM_SYSTEM == self.ddImmSystemNameSelected 
          && v.IMM_ITAM_ID == self.ddImmItamIDSelected
          && v.IMM_TABLE_NAME == self.ddImmTableNameSelected
          && v.IMM_COLUMN_NAME == self.ddImmColumnNameSelected
          && v.IMM_SCREEN_LABEL_NAME == self.ddImmScreenLabelSelected
          && v.ULT_SYSTEM == self.ddUltSystemNameSelected
          && v.ULT_ITAM_ID == self.ddUltItamIDSelected
      });
      
      return _.uniq(_.map(filtered, (v) => v.ULT_TABLE_NAME.toString().trim())).filter(Boolean);
    },
    ddUltColumnNameOptions () {
      var self = this;
      var filtered = _.filter(self.dpomy.DDSource, (v) => {
        return v.IMM_SYSTEM == self.ddImmSystemNameSelected 
          && v.IMM_ITAM_ID == self.ddImmItamIDSelected
          && v.IMM_TABLE_NAME == self.ddImmTableNameSelected
          && v.IMM_COLUMN_NAME == self.ddImmColumnNameSelected
          && v.IMM_SCREEN_LABEL_NAME == self.ddImmScreenLabelSelected
          && v.ULT_SYSTEM == self.ddUltSystemNameSelected
          && v.ULT_ITAM_ID == self.ddUltItamIDSelected
          && v.ULT_TABLE_NAME == self.ddUltTableNameSelected
      });
      
      return _.uniq(_.map(filtered, (v) => v.ULT_COLUMN_NAME.toString().trim())).filter(Boolean);
    },
    ddUltScreenLabelOptions () {
      var self = this;
      var filtered = _.filter(self.dpomy.DDSource, (v) => {
        return v.IMM_SYSTEM == self.ddImmSystemNameSelected 
          && v.IMM_ITAM_ID == self.ddImmItamIDSelected
          && v.IMM_TABLE_NAME == self.ddImmTableNameSelected
          && v.IMM_COLUMN_NAME == self.ddImmColumnNameSelected
          && v.IMM_SCREEN_LABEL_NAME == self.ddImmScreenLabelSelected
          && v.ULT_SYSTEM == self.ddUltSystemNameSelected
          && v.ULT_ITAM_ID == self.ddUltItamIDSelected
          && v.ULT_TABLE_NAME == self.ddUltTableNameSelected
          && v.ULT_COLUMN_NAME == self.ddUltColumnNameSelected
      });
      
      return _.uniq(_.map(filtered, (v) => v.ULT_SCREEN_LABEL_NAME.toString().trim())).filter(Boolean);
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
    ...mapActions("dpomy", {
      getDetails: "getDetails"
    }),
    handleClose () {
      this.$router.go(-1)
    },
    runGetDetails (param){
      var self = this;

      param.Which = self.$route.name;
      param.Left = parseInt(self.$route.params.system).toString();
      param.Right = self.$route.params.details.toString();

      return self.getDetails(param).then(
        res => {
          this.firstload = true;

          if (self.dpomy.detailsSource.length > 0){
            var tmp = self.dpomy.detailsSource[0].Values[0];

            self.selectedDetails = {}
            _.each(Object.keys(tmp), function(v, i){
                self.selectedDetails[v] = _.uniq(
                  _.map(self.dpomy.detailsSource[0].Values, (val) => val[v] ? val[v].toString().trim() : "").filter(Boolean)
                );

                self.selectedDetails[v] = self.selectedDetails[v].join(", ");
            });

            // make falsy NA
            Object.keys(self.selectedDetails).forEach((val) => {
              self.selectedDetails[val] = !!self.selectedDetails[val].trim() ? self.selectedDetails[val] : "NA";
            });
            
            setTimeout(() => {
              self.ddImmSystemNameSelected = self.selectedDetails.IMM_SYSTEM;
              self.ddImmItamIDSelected = self.selectedDetails.IMM_ITAM_ID;
              self.ddImmTableNameSelected = self.selectedDetails.IMM_TABLE_NAME;
              self.ddImmColumnNameSelected = self.selectedDetails.IMM_COLUMN_NAME;
              self.ddImmScreenLabelSelected = self.selectedDetails.IMM_SCREEN_LABEL_NAME;
              self.ddUltSystemNameSelected = self.selectedDetails.ULT_SYSTEM;
              self.ddUltItamIDSelected = self.selectedDetails.ULT_ITAM_ID;
              self.ddUltTableNameSelected = self.selectedDetails.ULT_TABLE_NAME;
              self.ddUltColumnNameSelected = self.selectedDetails.ULT_COLUMN_NAME;
              self.ddUltScreenLabelSelected = self.selectedDetails.ULT_SCREEN_LABEL_NAME;

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