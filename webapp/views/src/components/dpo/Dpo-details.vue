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
          <b-card tag="article" class="mb-2" v-if="selectedProcess">
            <b-media class="left-card-media" >
              <h6 class="left-card-title">Downstream Process Name</h6>
              <p v-html="selectedProcess.Name"></p>
            </b-media>
            
            <b-media class="left-card-media">
              <h6 class="left-card-title">Process Owner</h6>
              <p v-html="selectedProcess.Owner_ID"></p>
            </b-media>
            
            <b-media class="left-card-media">
              <h6 class="left-card-title">Bank ID</h6>
              <p></p>
            </b-media>
            
            <b-media class="left-card-media">
              <h6 class="left-card-title">CDE Name</h6>
              <p v-html="selectedProcessDet.Business_Term_ID"></p>
            </b-media>
            
            <b-media class="left-card-media">
              <h6 class="left-card-title">CDE Rationale</h6>
              <p v-html="selectedProcessDet.CDE_Rationale"></p>
            </b-media>
          </b-card>
        </b-col>

        <b-col cols="8"> 
          <b-row>
            <b-col>
              <b-card title="Immediate Preceding System Details" tag="article" class="mb-2" v-if="selectedProcessDet">
                <p class="card-text">
                  <b-form>
                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="System Name" label-for="systemName">
                      <b-form-select id="systemName" class="col-8" v-model="ddSystem.selected" :options="ddSystemOptions">
                      </b-form-select>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="ITAM ID" label-for="itamid">
                      <b-form-select id="itamid" class="col-8" v-model="ddItam.selected" :options="ddItamOptions"></b-form-select>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="TableName" label-for="tablename">
                      <b-form-select id="tablename" class="col-8" v-model="ddTableName.selected" :options="ddTableOptions"></b-form-select>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Column Name" label-for="columnname">
                      <b-form-select id="columnname" class="col-8" v-model="ddColumnName.selected" :options="ddColumnNameOptions"></b-form-select>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Screen Label Name" label-for="screenlabelname">
                      <b-form-select id="screenlabelname" class="col-8"></b-form-select>
                    </b-form-group>
                    
                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Business Description">
                      <p v-html="selectedProcessDet.BusinessTerm.Description"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Derived (Yes/No)">
                      <p v-html="selectedProcessDet.System.Columns.Derived"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Derivation Logic">
                      <p v-html="selectedProcessDet.System.Columns.Derivation_Logic"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Data Quality Requirements">
                      <p v-html="selectedProcessDet.System.Columns.DQ_Standards"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Thresholds">
                      <p v-html="selectedProcessDet.System.Columns.Threshold"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Data SLA signed?">
                      <p v-html="selectedProcessDet.System.Columns.Data_SLA_Signed"></p>
                    </b-form-group>
                  </b-form>
                </p>
              </b-card>
            </b-col>
          </b-row>

          <b-row>
            <b-col>
              <b-card title="Ultimate Source System Details" tag="article" class="mb-2" v-if="selectedProcessDet">
                <p class="card-text">
                  <b-form>
                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="System Name" label-for="systemName">
                      <b-form-select id="systemName" class="col-8" v-model="ddSystem.selected" :options="ddSystemOptions">
                      </b-form-select>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="ITAM ID" label-for="itamid">
                      <b-form-select id="itamid" class="col-8" v-model="ddItam.selected" :options="ddItamOptions"></b-form-select>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="TableName" label-for="tablename">
                      <b-form-select id="tablename" class="col-8" v-model="ddTableName.selected" :options="ddTableOptions"></b-form-select>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Column Name" label-for="columnname">
                      <b-form-select id="columnname" class="col-8" v-model="ddColumnName.selected" :options="ddColumnNameOptions"></b-form-select>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Screen Label Name" label-for="screenlabelname">
                      <b-form-select id="screenlabelname" class="col-8"></b-form-select>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Business Description">
                      <p v-html="selectedProcessDet.BusinessTerm.Description"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Derived?">
                      <p v-html="selectedProcessDet.System.Columns.Derived"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Derivation Logic">
                      <p v-html="selectedProcessDet.System.Columns.Derivation_Logic"></p> 
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Data Quality requirements">
                      <p v-html="selectedProcessDet.System.Columns.DQ_Standards"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Thresholds">
                      <p v-html="selectedProcessDet.System.Columns.Threshold"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Golden Source (Yes/No)">
                      <p v-html="selectedProcessDet.System.Columns.Golden_Source"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Golden Source System Name">
                      <p v-html="selectedProcessDet.BusinessTerm.Golden_Source_ID"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Golden Source ITAM ID">
                      <p v-html="selectedProcessDet.System.ITAM_ID"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Golden Source Table Name">
                      <p v-html="selectedProcessDet.System.Columns.MDTable.Business_Term_ID"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Golden Source Column Name">
                      <p v-html="selectedProcessDet.System.Columns.Name"></p>
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
              <b-card title="Domain View" tag="article" class="mb-2" v-if="selectedProcessDet">
                <p class="card-text">
                  <b-form>
                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Domain">
                      <p v-html="selectedProcessDet.BusinessTerm.SubCategory.Category.Name"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Sub Domain">
                      <p v-html="selectedProcessDet.BusinessTerm.SubCategory.Name"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Domain Owner">
                      <p></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Business Term">
                      <p v-html="selectedProcessDet.BusinessTerm.BT_Name"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Business Term Description">
                      <p v-html="selectedProcessDet.BusinessTerm.Description"></p>
                    </b-form-group>
                  </b-form>
                </p>
              </b-card>
            </b-col>
          </b-row>

          <b-row>
            <b-col>
              <b-card title="Data Standards" tag="article" class="mb-2" v-if="selectedProcessDet">
                <p class="card-text">
                  <b-form>
                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="DQ Standards set by DPO">
                      <p v-html="selectedProcessDet.System.Columns.DPO_DQ_Standards"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="DQ Standards set at Business Term Level">
                      <p v-html="selectedProcessDet.BusinessTerm.DQ_Standards"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Thresholds defined by DPO*">
                      <p v-html="selectedProcessDet.System.Columns.DPO_Threshold"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Thresholds set at Business term level">
                      <p v-html="selectedProcessDet.System.Columns.Threshold"></p>
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
      showModal: this.$route.meta.showModal,
      selectedProcess: null,
      selectedProcessDet: null,
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
        'Downstream Process Name': "selectedProcess.Name",
        'Process Owner': "selectedProcess.Owner_ID",
        'CDE Name': 'selectedProcessDet.Business_Term_ID',
        'CDE Rationale': 'selectedProcessDet.CDE_Rationale',
        'System Name': 'ddSystem.selected',
        'ITAM ID': 'ddItam.selected',
        'TableName': 'ddTableName.selected',
        "Column Name": "ddColumnName.selected",
        "Screen Label Name": "asdf",
        "Business Description": "selectedProcessDet.BusinessTerm.Description",
        "Derived (Yes/No)": "selectedProcessDet.System.Columns.Derived",
        "Derivation Logic": "selectedProcessDet.System.Columns.Derivation_Logic",
        "Data Quality Requirements": "selectedProcessDet.System.Columns.DQ_Standards",
        "Thresholds": "selectedProcessDet.System.Columns.Threshold",
        "Data SLA signed?": "selectedProcessDet.System.Columns.Data_SLA_Signed",
        'System Name': 'ddSystem.selected',
        'ITAM ID': 'ddItam.selected',
        'TableName': 'ddTableName.selected',
        "Column Name": "ddColumnName.selected",
        "Screen Label Name": "asdf",
        "Business Description": "selectedProcessDet.BusinessTerm.Description",
        "Derived?": "selectedProcessDet.System.Columns.Derived",
        "Derivation Logic": "selectedProcessDet.System.Columns.Derivation_Logic",
        "Data Quality requirements": "selectedProcessDet.System.Columns.DQ_Standards",
        "Thresholds": "selectedProcessDet.System.Columns.Threshold",
        "Golden Source (Yes/No)": "selectedProcessDet.System.Columns.Golden_Source",
        "Golden Source System Name": "selectedProcessDet.BusinessTerm.Golden_Source_ID",
        "Golden Source ITAM ID": "selectedProcessDet.System.ITAM_ID",
        "Golden Source Table Name": "selectedProcessDet.System.Columns.MDTable.Business_Term_ID",
        "Golden Source Column Name": "selectedProcessDet.System.Columns.Name",
        "Golden Source Screen Name": "asdf",
        "Golden Source Business Description": "asdf",
        "Golden Source Derived (Yes/No)": "asdf",
        "Golden Source Derivation Logic": "asdf",
        "Domain": "selectedProcessDet.BusinessTerm.SubCategory.Category.Name",
        "Sub Domain": "selectedProcessDet.BusinessTerm.SubCategory.Name",
        "Domain Owner": "asdf",
        "Business Term": "selectedProcessDet.BusinessTerm.BT_Name",
        "Business Term Description": "selectedProcessDet.BusinessTerm.Description",
        "DQ Standards set by DPO": "selectedProcessDet.System.Columns.DPO_DQ_Standards",
        "DQ Standards set at Business Term Level": "selectedProcessDet.BusinessTerm.DQ_Standards",
        "Thresholds defined by DPO*": "selectedProcessDet.System.Columns.DPO_Threshold"
      }
    }
  },
  computed: {
    ...mapState({
      dpomy: state => state.dpomy.all
    }),
    count () {
      return this.dpomy.tableSource.length
    },
    ddSystemOptions () {
      this.ddSystem.selected = this.selectedProcess ? this.selectedProcess.Name : null;
      return this.selectedProcess ? [{ value: this.selectedProcess.Name, text: this.selectedProcess.Name }] : [];
    },
    ddItamOptions () {
      this.ddItam.selected = this.selectedProcessDet ? this.selectedProcessDet.System.ITAM_ID : null;
      return this.selectedProcessDet.System ? [{ value: this.selectedProcessDet.System.ITAM_ID, text: this.selectedProcessDet.System.ITAM_ID }] : [];
    },
    ddTableOptions () {
      this.ddTableName.selected = this.selectedProcessDet ? this.selectedProcessDet.System.Columns.MDTable.Name : null;
      return this.selectedProcessDet.System ? [{ value: this.selectedProcessDet.System.Columns.MDTable.Name, text: this.selectedProcessDet.System.Columns.MDTable.Name }] : [];
    },
    ddColumnNameOptions () {
      this.ddColumnName.selected = this.selectedProcessDet ? this.selectedProcessDet.System.Columns.Name : null;
      return this.selectedProcessDet.System ? [{ value: this.selectedProcessDet.System.Columns.Name, text: this.selectedProcessDet.System.Columns.Name }] : [];
    },
    exportDatas () {
      if(this.selectedProcess && this.selectedProcessDet && this.selectedProcessDet){
        return [{
          selectedProcess: this.selectedProcess,
          selectedProcessDet: this.selectedProcessDet,
          ddSystem: this.ddSystem,
          ddItam: this.ddItam,
          ddTableName: this.ddTableName,
          ddColumnName: this.ddColumnName
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
      this.selectedProcess = _.find(this.dpomy.systemsSource, ['ID', parseInt(this.$route.params.system)])
      this.selectedProcessDet = _.find(this.dpomy.tableSource, ['ID', parseInt(this.$route.params.details)])
    },
    selectedProcessDet (){
      this.ddItam.selected = this.selectedProcessDet ? this.selectedProcessDet.System.ITAM_ID : null;
    }
  },
  mounted() {
    this.$refs.modalDetails.show();
    
    this.selectedProcess = _.find(this.dpomy.systemsSource, ['ID', parseInt(this.$route.params.system)])
    this.selectedProcessDet = _.find(this.dpomy.tableSource, ['ID', parseInt(this.$route.params.details)])
  },
  methods: {
    handleClose () {
      this.$router.go(-1)
    }
  },
}
</script>