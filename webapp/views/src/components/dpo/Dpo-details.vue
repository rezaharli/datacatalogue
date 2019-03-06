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
              <h6 class="left-card-title">Downstream Process Name</h6>
              <text-wrap-dialog :fulltext="selectedDetails.DOWNSTREAM_PROCESS"></text-wrap-dialog>
            </b-media>
            
            <b-media class="left-card-media">
              <h6 class="left-card-title">Process Owner</h6>
              <text-wrap-dialog :fulltext="selectedDetails.PROCESS_OWNER"></text-wrap-dialog>
            </b-media>
            
            <b-media class="left-card-media">
              <h6 class="left-card-title">Bank ID</h6>
              <text-wrap-dialog :fulltext="selectedDetails.BANK_ID"></text-wrap-dialog>
            </b-media>
            
            <b-media class="left-card-media">
              <h6 class="left-card-title">CDE Name</h6>
              <text-wrap-dialog :fulltext="selectedDetails.BUSINESS_TERM_ID"></text-wrap-dialog>
            </b-media>
            
            <b-media class="left-card-media">
              <h6 class="left-card-title">CDE Rationale</h6>
              <text-wrap-dialog :fulltext="selectedDetails.CDE_RATIONALE"></text-wrap-dialog>
            </b-media>
          </b-card>
        </b-col>

        <b-col cols="8"> 
          <b-row>
            <b-col>
              <b-card title="Immediate Preceding System Details" tag="article" class="mb-2" v-if="selectedDetails">
                <p class="card-text">
                  <b-form>
                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="System Name" label-for="systemName">
                      <b-form-select id="systemName" class="col-8" v-model="selectedDetails.SYSTEM_NAME" :options="[selectedDetails.SYSTEM_NAME]">
                      </b-form-select>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="ITAM ID" label-for="itamid">
                      <b-form-select id="itamid" class="col-8" v-model="selectedDetails.ITAM_ID" :options="[selectedDetails.ITAM_ID]"></b-form-select>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="TableName" label-for="tablename">
                      <b-form-select id="tablename" class="col-8" v-model="selectedDetails.TABLE_NAME" :options="[selectedDetails.TABLE_NAME]"></b-form-select>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Column Name" label-for="columnname">
                      <b-form-select id="columnname" class="col-8" v-model="selectedDetails.COLUMN_NAME" :options="[selectedDetails.COLUMN_NAME]"></b-form-select>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Screen Label Name" label-for="screenlabelname">
                      <b-form-select id="screenlabelname" class="col-8"></b-form-select>
                    </b-form-group>
                    
                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Business Description">
                      <text-wrap-dialog :fulltext="selectedDetails.BUSINESS_DESCRIPTION"></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Derived (Yes/No)">
                      <text-wrap-dialog :fulltext="selectedDetails.DERIVED"></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Derivation Logic">
                      <text-wrap-dialog :fulltext="selectedDetails.DERIVATION_LOGIC"></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Data Quality Requirements">
                      <text-wrap-dialog :fulltext="selectedDetails.DQ_REQUIREMENTS"></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Thresholds">
                      <text-wrap-dialog :fulltext="selectedDetails.THRESHOLD"></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Data SLA signed?">
                      <text-wrap-dialog :fulltext="selectedDetails.DATA_SLA_SIGNED"></text-wrap-dialog>
                    </b-form-group>
                  </b-form>
                </p>
              </b-card>
            </b-col>
          </b-row>

          <b-row>
            <b-col>
              <b-card title="Ultimate Source System Details" tag="article" class="mb-2" v-if="selectedDetails">
                <p class="card-text">
                  <b-form>
                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="System Name" label-for="systemName">
                      <b-form-select id="systemName" class="col-8" v-model="selectedDetails.SYSTEM_NAME" :options="[selectedDetails.SYSTEM_NAME]">
                      </b-form-select>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="ITAM ID" label-for="itamid">
                      <b-form-select id="itamid" class="col-8" v-model="selectedDetails.ITAM_ID" :options="[selectedDetails.ITAM_ID]"></b-form-select>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="TableName" label-for="tablename">
                      <b-form-select id="tablename" class="col-8" v-model="selectedDetails.TABLE_NAME" :options="[selectedDetails.TABLE_NAME]"></b-form-select>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Column Name" label-for="columnname">
                      <b-form-select id="columnname" class="col-8" v-model="selectedDetails.COLUMN_NAME" :options="[selectedDetails.COLUMN_NAME]"></b-form-select>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Screen Label Name" label-for="screenlabelname">
                      <b-form-select id="screenlabelname" class="col-8"></b-form-select>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Business Description">
                      <text-wrap-dialog :fulltext="selectedDetails.BUSINESS_DESCRIPTION"></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Derived?">
                      <text-wrap-dialog :fulltext="selectedDetails.DERIVED"></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Derivation Logic">
                      <text-wrap-dialog :fulltext="selectedDetails.DERIVATION_LOGIC"></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Data Quality requirements">
                      <text-wrap-dialog :fulltext="selectedDetails.DQ_STANDARDS"></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Thresholds">
                      <text-wrap-dialog :fulltext="selectedDetails.THRESHOLD"></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Golden Source (Yes/No)">
                      <text-wrap-dialog :fulltext="selectedDetails.GOLDEN_SOURCE"></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Golden Source System Name">
                      <text-wrap-dialog :fulltext="selectedDetails.GOLDEN_SOURCE_SYSTEM_ID"></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Golden Source ITAM ID">
                      <p></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Golden Source Table Name">
                      <p></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Golden Source Column Name">
                      <p></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Golden Source Screen Name">
                      <p></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Golden Source Business Description">
                      <p></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Golden Source Derived (Yes/No)">
                      <p ></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Golden Source Derivation Logic">
                      <p></p>
                    </b-form-group>
                  </b-form>
                </p>
              </b-card>
            </b-col>
          </b-row>

          <b-row>
            <b-col>
              <b-card title="Domain View" tag="article" class="mb-2" v-if="selectedDetails">
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

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Business Term">
                      <text-wrap-dialog :fulltext="selectedDetails.BUSINESS_TERM"></text-wrap-dialog>
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
              <b-card title="Data Standards" tag="article" class="mb-2" v-if="selectedDetails">
                <p class="card-text">
                  <b-form>
                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="DQ Standards set by DPO">
                      <text-wrap-dialog :fulltext="selectedDetails.DPO_DQ_STANDARDS"></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="DQ Standards set at Business Term Level">
                      <text-wrap-dialog :fulltext="selectedDetails.DQ_STANDARDS"></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Thresholds defined by DPO*">
                      <text-wrap-dialog :fulltext="selectedDetails.DPO_THRESHOLD"></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Thresholds set at Business term level">
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
      showModal: this.$route.meta.showModal,
      selectedDetails: null,
      ddSystem: {
        selected: null,
      },
      ddItam: {
        selected: null,
      },
      ddTableName: {
        selected: null,
      },
      ddColumnName: {
        selected: null,
      },
      excelFields: {
        'Downstream Process Name': "selectedDetails.DOWNSTREAM_PROCESS",
        'Process Owner': "selectedDetails.PROCESS_OWNER",
        'CDE Name': 'selectedDetails.BUSINESS_TERM_ID',
        'CDE Rationale': 'selectedDetails.CDE_RATIONALE',
        'System Name': 'selectedDetails.SYSTEM_NAME',
        'ITAM ID': 'selectedDetails.ITAM_ID',
        'TableName': 'selectedDetails.TABLE_NAME',
        "Column Name": "selectedDetails.COLUMN_NAME",
        "Screen Label Name": "asdf",
        "Business Description": "selectedDetails.BUSINESS_DESCRIPTION",
        "Derived (Yes/No)": "selectedDetails.DERIVED",
        "Derivation Logic": "selectedDetails.DERIVATION_LOGIC",
        "Data Quality Requirements": "selectedDetails.DQ_REQUIREMENTS",
        "Thresholds": "selectedDetails.THRESHOLD",
        "Data SLA signed?": "selectedDetails.DATA_SLA_SIGNED",
        'System Name': 'selectedDetails.SYSTEM_NAME',
        'ITAM ID': 'selectedDetails.ITAM_ID',
        'TableName': 'selectedDetails.TABLE_NAME',
        "Column Name": "selectedDetails.COLUMN_NAME",
        "Screen Label Name": "asdf",
        "Business Description": "selectedDetails.BUSINESS_DESCRIPTION",
        "Derived?": "selectedDetails.DERIVED",
        "Derivation Logic": "selectedDetails.DERIVATION_LOGIC",
        "Data Quality requirements": "selectedDetails.DQ_REQUIREMENTS",
        "Thresholds": "selectedDetails.THRESHOLD",
        "Golden Source (Yes/No)": "selectedDetails.GOLDEN_SOURCE",
        "Golden Source System Name": "selectedDetails.GOLDEN_SOURCE_SYSTEM_ID",
        "Golden Source ITAM ID": "ASDF",
        "Golden Source Table Name": "ASDF",
        "Golden Source Column Name": "ASDF",
        "Golden Source Screen Name": "asdf",
        "Golden Source Business Description": "asdf",
        "Golden Source Derived (Yes/No)": "asdf",
        "Golden Source Derivation Logic": "asdf",
        "Domain": "selectedDetails.DOMAIN",
        "Sub Domain": "selectedDetails.SUBDOMAIN",
        "Domain Owner": "selectedDetails.DOMAIN_OWNER",
        "Business Term": "selectedDetails.BUSINESS_TERM",
        "Business Term Description": "selectedDetails.BUSINESS_DESCRIPTION",
        "DQ Standards set by DPO": "selectedDetails.DPO_DQ_Standards",
        "DQ Standards set at Business Term Level": "selectedDetails.DQ_Standards",
        "Thresholds defined by DPO*": "selectedDetails.DPO_Threshold"
      }
    }
  },
  computed: {
    ...mapState({
      dpomy: state => state.dpomy.all
    }),
    ddSystemOptions () {
      this.ddSystem.selected = this.selectedDetails ? this.selectedDetails.Name : null;
      return this.selectedDetails ? [{ value: this.selectedDetails.Name, text: this.selectedDetails.Name }] : [];
    },
    ddItamOptions () {
      this.ddItam.selected = this.selectedDetails ? this.selectedDetails.ITAM_ID : null;
      return this.selectedDetails ? [{ value: this.selectedDetails.ITAM_ID, text: this.selectedDetails.ITAM_ID }] : [];
    },
    ddTableOptions () {
      this.ddTableName.selected = this.selectedDetails ? this.selectedDetails.Name : null;
      return this.selectedDetails ? [{ value: this.selectedDetails.Name, text: this.selectedDetails.Name }] : [];
    },
    ddColumnNameOptions () {
      this.ddColumnName.selected = this.selectedDetails ? this.selectedDetails.Name : null;
      return this.selectedDetails ? [{ value: this.selectedDetails.Name, text: this.selectedDetails.Name }] : [];
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
    selectedDetails (){
      this.ddItam.selected = this.selectedDetails ? this.selectedDetails.ITAM_ID : null;
    },
    "dpomy.detailsSource"() {
      if (this.dpomy.detailsSource.length > 0)
        this.selectedDetails = this.dpomy.detailsSource[0];
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
    // this.selectedDetails = _.find(this.dpomy.systemsSource, ['ID', parseInt(this.$route.params.system)])
    // this.selectedDetails = _.find(this.dpomy.tableSource, ['ID', parseInt(this.$route.params.details)])
  },
  methods: {
    ...mapActions("dpomy", {
      getDetails: "getDetails"
    }),
    handleClose () {
      this.$router.go(-1)
    }
  },
}
</script>