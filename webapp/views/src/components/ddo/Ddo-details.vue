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
  <b-modal
    hide-footer
    header-class="modal-details-header setbackground"
    body-class="setbackground"
    v-if="showModal"
    id="modal-details"
    ref="modalDetails"
    size="lg"
    @hidden="handleClose"
  >
    <b-container fluid class="row-kasijarak">
      <b-row>
        <b-col>
          <b-col>
            <download-excel
              :data="exportDatas"
              :fields="excelFields"
              worksheet="My Worksheet"
              name="filename.xls"
            >
              <b-btn size="sm" class="float-right" variant="success">Export</b-btn>
            </download-excel>
          </b-col>
        </b-col>
      </b-row>
      <!-- </b-container>

      <b-container>-->
      <b-row>
        <b-col cols="4">
          <b-card tag="article" class="mb-2" v-if="selectedDetails">
            <b-media class="left-card-media">
              <h6 class="left-card-title">Data Domain</h6>
              <text-wrap-dialog :fulltext="selectedDetails.DATA_DOMAIN"></text-wrap-dialog>
            </b-media>

            <b-media class="left-card-media">
              <h6 class="left-card-title">Sub Domain</h6>
              <text-wrap-dialog :fulltext="selectedDetails.SUB_DOMAIN"></text-wrap-dialog>
            </b-media>

            <b-media class="left-card-media">
              <h6 class="left-card-title">Sub Domain Owner</h6>
              <text-wrap-dialog :fulltext="selectedDetails.SUB_DOMAIN_OWNER"></text-wrap-dialog>
            </b-media>

            <b-media class="left-card-media">
              <h6 class="left-card-title">Bank ID</h6>
              <text-wrap-dialog :fulltext="selectedDetails.BANK_ID"></text-wrap-dialog>
            </b-media>
          </b-card>
        </b-col>

        <b-col cols="8">
          <b-row>
            <b-col>
              <b-card
                title="Technical Metadata From System"
                tag="article"
                class="mb-2"
                v-if="selectedDetails"
              >
                <p class="card-text">
                  <b-form>
                    <b-form-group
                      horizontal
                      :label-cols="4"
                      breakpoint="md"
                      label="Business Term"
                      label-for="businessterm"
                    >
                      <b-form-select
                        id="businessterm"
                        class="col-8"
                        v-model="selectedDetails.BT_NAME"
                        :options="[selectedDetails.BT_NAME]"
                      ></b-form-select>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Business Term Description">
                      <text-wrap-dialog :fulltext="selectedDetails.BUSINESS_TERM_DESCRIPTION"></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group
                      horizontal
                      :label-cols="4"
                      breakpoint="md"
                      label="Business Alias"
                      label-for="businessalias"
                    >
                      <b-form-select
                        id="businessalias"
                        class="col-8"
                        v-model="selectedDetails.ALIAS_NAME"
                        :options="[selectedDetails.ALIAS_NAME]"
                      ></b-form-select>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="CDE (Yes/No)">
                      <text-wrap-dialog :fulltext="selectedDetails.CDE"></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Data Quality standards | DQ Thresholds">
                      <text-wrap-dialog :fulltext="selectedDetails.DQ_STANDARDS"></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Policy guidance (if any)">
                      <text-wrap-dialog :fulltext="selectedDetails.POLICY_GUIDANCE"></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Mandatory (Yes/No)">
                      <text-wrap-dialog :fulltext="selectedDetails.MANDATORY"></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Golden Source System">
                      <text-wrap-dialog :fulltext="selectedDetails.GOLDEN_SOURCE_SYSTEM_ID"></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Golden Source ITAM ID">
                      <text-wrap-dialog :fulltext="selectedDetails.ITAM_ID"></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Golden Source Table Name">
                      <text-wrap-dialog :fulltext="selectedDetails.GOLDEN_SOURCE_TABLE_NAME"></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Golden Source Column Name">
                      <text-wrap-dialog :fulltext="selectedDetails.GOLDEN_SOURCE_COLUMN_NAME"></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Target Golden Source">
                      <text-wrap-dialog :fulltext="selectedDetails.ITAM_ID"></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Target golden source ITAM ID">
                      <text-wrap-dialog :fulltext="selectedDetails.ITAM_ID"></text-wrap-dialog>
                    </b-form-group>
                  </b-form>
                </p>
              </b-card>
            </b-col>
          </b-row>

          <b-row>
            <b-col>
              <b-card
                title="Policy Related Information"
                tag="article"
                class="mb-2"
                v-if="selectedDetails"
              >
                <p class="card-text">
                  <b-form>
                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Information Asset Names">
                      <text-wrap-dialog :fulltext="selectedDetails.INFO_ASSET_NAME"></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Information Asset Descriptions">
                      <text-wrap-dialog :fulltext="selectedDetails.INFO_ASSET_DESC"></text-wrap-dialog>
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

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="PII Flag (Yes/No)">
                      <text-wrap-dialog :fulltext="selectedDetails.PII_FLAG"></text-wrap-dialog>
                    </b-form-group>
                  </b-form>
                </p>
              </b-card>
            </b-col>
          </b-row>

          <b-row>
            <b-col>
              <b-card title="Downstream Usage of Business Term" tag="article" class="mb-2" v-if="selectedDetails">
                <p class="card-text">
                  <b-form>
                    <b-form-group
                      horizontal
                      :label-cols="4"
                      breakpoint="md"
                      label="Downstream Process Name"
                      label-for="downstreamprocessname"
                    >
                      <b-form-select
                        id="downstreamprocessname"
                        class="col-8"
                        v-model="selectedDetails.DOWNSTREAM_PROCESS_NAME"
                        :options="[selectedDetails.DOWNSTREAM_PROCESS_NAME]"
                      ></b-form-select>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Downstream Process Owner">
                      <text-wrap-dialog :fulltext="selectedDetails.DOWNSTREAM_PROCESS_OWNER"></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Is data taken from Golden Source">
                      <text-wrap-dialog :fulltext="selectedDetails.ASDF"></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="System Name">
                      <text-wrap-dialog :fulltext="selectedDetails.SYSTEM_NAME"></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="ITAM ID">
                      <text-wrap-dialog :fulltext="selectedDetails.ITAM_ID"></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Table Name">
                      <text-wrap-dialog :fulltext="selectedDetails.GOLDEN_SOURCE_TABLE_NAME"></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Column Name">
                      <text-wrap-dialog :fulltext="selectedDetails.COLUMN_NAME"></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="DQ Standards | Thresholds set by DDO">
                      <text-wrap-dialog :fulltext="selectedDetails.DDO_THRESHOLD"></text-wrap-dialog>
                    </b-form-group>
                  </b-form>
                </p>
              </b-card>
            </b-col>
          </b-row>

          <b-row>
            <b-col>
              <b-card title="Business Term Residing In Other Systems" tag="article" class="mb-2" v-if="selectedDetails">
                <p class="card-text">
                  <b-form>
                    <b-form-group
                      horizontal
                      :label-cols="4"
                      breakpoint="md"
                      label="System Name"
                      label-for="systemname"
                    >
                      <b-form-select
                        id="systemname"
                        class="col-8"
                        v-model="selectedDetails.SYSTEM_NAME"
                        :options="[selectedDetails.SYSTEM_NAME]"
                      ></b-form-select>
                    </b-form-group>

                    <b-form-group
                      horizontal
                      :label-cols="4"
                      breakpoint="md"
                      label="ITAM ID"
                      label-for="ITAMID"
                    >
                      <b-form-select
                        id="ITAMID"
                        class="col-8"
                        v-model="selectedDetails.ITAM_ID"
                        :options="[selectedDetails.ITAM_ID]"
                      ></b-form-select>
                    </b-form-group>

                    <b-form-group
                      horizontal
                      :label-cols="4"
                      breakpoint="md"
                      label="Table Name"
                      label-for="tablename"
                    >
                      <b-form-select
                        id="tablename"
                        class="col-8"
                        v-model="selectedDetails.GOLDEN_SOURCE_TABLE_NAME"
                        :options="[selectedDetails.GOLDEN_SOURCE_TABLE_NAME]"
                      ></b-form-select>
                    </b-form-group>

                    <b-form-group
                      horizontal
                      :label-cols="4"
                      breakpoint="md"
                      label="Column Name"
                      label-for="colname"
                    >
                      <b-form-select
                        id="colname"
                        class="col-8"
                        v-model="selectedDetails.COLUMN_NAME"
                        :options="[selectedDetails.COLUMN_NAME]"
                      ></b-form-select>
                    </b-form-group>
                  </b-form>
                </p>
              </b-card>
            </b-col>
          </b-row>
        </b-col>
      </b-row>

      <b-row></b-row>
    </b-container>
  </b-modal>
</template>

<script>
import { mapState, mapActions } from "vuex";
import textWrapDialog from '../TextWrapDialog.vue'

export default {
  components: {
    textWrapDialog
  },
  data() {
    return {
      firstload: true,
      showModal: this.$route.meta.showModal,
      selectedDetails: null,
      ddTableSelected: null,
      ddColumnSelected: null,
      excelFields: {
        "Data Domain": "selectedDetails.DATA_DOMAIN",
        "Sub Domain": "selectedDetails.SUB_DOMAIN",
        "Sub Domain Owner": "selectedDetails.SUB_DOMAIN_OWNER",
        "Bank ID": "selectedDetails.BANK_ID",
        "Business Term": "selectedDetails.BT_NAME",
        "Business Alias": "selectedDetails.ALIAS_NAME",
        "CDE (Yes/No)": "selectedDetails.CDE",
        "Data Quality standards | DQ Thresholds": "selectedDetails.DQ_STANDARDS",
        "Policy guidance (if any)": "selectedDetails.POLICY_GUIDANCE",
        "Mandatory (Yes/No)": "selectedDetails.MANDATORY",
        "Golden Source System": "selectedDetails.GOLDEN_SOURCE_SYSTEM_ID",
        "Golden Source ITAM ID": "selectedDetails.ITAM_ID",
        "Golden Source Table Name": "selectedDetails.GOLDEN_SOURCE_TABLE_NAME",
        "Golden Source Column Name": "selectedDetails.GOLDEN_SOURCE_COLUMN_NAME",
        "Target Golden Source": "selectedDetails.ITAM_ID",
        "Target golden source ITAM ID": "selectedDetails.ITAM_ID",
        "Information Asset Names": "selectedDetails.INFO_ASSET_NAME",
        "Information Asset Descriptions": "selectedDetails.INFO_ASSET_DESC",
        "C - Confidentiality": "selectedDetails.CONFIDENTIALITY",
        "I - Integrity": "selectedDetails.INTEGRITY",
        "A - Availability": "selectedDetails.AVAILABILITY",
        "Overall CIA Rating": "selectedDetails.OVERALL_CIA_RATING",
        "Record Categories": "selectedDetails.RECORD_CATEGORY",
        "PII Flag (Yes/No)": "selectedDetails.PII_FLAG",
        "Downstream Process Name": "selectedDetails.DOWNSTREAM_PROCESS_NAME",
        "Downstream Process Owner": "selectedDetails.DOWNSTREAM_PROCESS_OWNER",
        "Is data taken from Golden Source": "selectedDetails.ASDF",
        "System Name": "selectedDetails.SYSTEM_NAME",
        "ITAM ID": "selectedDetails.ITAM_ID",
        "Table Name": "selectedDetails.GOLDEN_SOURCE_TABLE_NAME",
        "Column Name": "selectedDetails.COLUMN_NAME",
        "DQ Standards | Thresholds set by DDO": "selectedDetails.DDO_THRESHOLD",
        "System Name": "selectedDetails.SYSTEM_NAME",
        "ITAM ID": "selectedDetails.ITAM_ID",
        "Table Name": "selectedDetails.GOLDEN_SOURCE_TABLE_NAME",
        "Column Name": "selectedDetails.COLUMN_NAME"
      }
    };
  },
  computed: {
    ...mapState({
      ddomy: state => state.ddomy.all
    }),
    ddTableOptions () {
      return _.uniq(_.map(this.ddomy.DDSource, "TABLE_NAME"))
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
    "$route.meta"({ showModal }) {
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
    ddScreenLabelSelected (){
      if( ! this.firstload){
        var param = {
          ScreenLabel: this.ddScreenLabelSelected.toString(),
          ColumnName: this.ddColumnSelected.toString(),
          TableName: this.ddTableSelected.toString(),
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
    ...mapActions("ddomy", {
      getDetails: "getDetails"
    }),
    handleClose() {
      this.$router.go(-1);
    },
    runGetDetails(param){
      var self = this;

      param.Which = self.$route.name;
      param.Left = parseInt(self.$route.params.system).toString();
      param.Right = parseInt(self.$route.params.details).toString();

      self.getDetails(param).then(
        res => {
          if (self.ddomy.detailsSource.length > 0){
            // self.selectedDetails = self.ddomy.detailsSource[0];
            var tmp = self.ddomy.detailsSource[0].Values[0];

            self.selectedDetails = {}
            _.each(Object.keys(tmp), function(v, i){
                self.selectedDetails[v] = _.uniq(
                  _.map(self.ddomy.detailsSource[0].Values, (val) => val[v].toString().trim()).filter(Boolean)).join(', ');
            });

            // make falsy NA
            Object.keys(self.selectedDetails).forEach((val) => {
              self.selectedDetails[val] = !!self.selectedDetails[val].trim() ? self.selectedDetails[val] : "NA";
            });
            
            setTimeout(() => {
              self.ddTableSelected = self.selectedDetails.TABLE_NAME;
              self.ddColumnSelected = self.selectedDetails.COLUMN_NAME;
              self.ddScreenLabelSelected = self.selectedDetails.BUSINESS_ALIAS_NAME;
            }, 100);
          } else {
            this.selectedDetails = null;
          }
        },
        err => err
      );
    }
  }
};
</script>