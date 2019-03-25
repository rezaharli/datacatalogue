<style>
.left-card-media {
  margin-bottom: 10px;
}
.left-card-title {
  font-weight: bolder;
}
.card-title {
  font-weight: bolder;
  border-bottom: 1px solid #cecece;
  padding-bottom: 9px;
}
.form-group {
  margin-bottom: 10px !important;
}
legend.col-form-label,
label.col-form-label {
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
    <page-loader v-if="ddomy.detailsLoading" />
    
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
              <v-btn small color="success" class="float-right">Export</v-btn>
            </download-excel>
          </b-col>
        </b-col>
      </b-row>
      <!-- </b-container>

      <b-container>-->
      <b-row>
        <b-col cols="4">
          <b-card tag="article" class="mb-2">
            <b-media class="left-card-media">
              <h6 class="left-card-title">Data Domain</h6>
              <text-wrap-dialog
                :fulltext="selectedDetailsBusinessMetadata ? selectedDetailsBusinessMetadata.DATA_DOMAIN : ''"
              ></text-wrap-dialog>
            </b-media>

            <b-media class="left-card-media">
              <h6 class="left-card-title">Sub Domain</h6>
              <text-wrap-dialog
                :fulltext="selectedDetailsBusinessMetadata ? selectedDetailsBusinessMetadata.SUB_DOMAIN : ''"
              ></text-wrap-dialog>
            </b-media>

            <b-media class="left-card-media">
              <h6 class="left-card-title">Sub Domain Owner</h6>
              <text-wrap-dialog
                :fulltext="selectedDetailsBusinessMetadata ? selectedDetailsBusinessMetadata.SUBDOMAIN_OWNER : ''"
              ></text-wrap-dialog>
            </b-media>

            <b-media class="left-card-media">
              <h6 class="left-card-title">Bank ID</h6>
              <text-wrap-dialog
                :fulltext="selectedDetailsBusinessMetadata ? selectedDetailsBusinessMetadata.BANK_ID : ''"
              ></text-wrap-dialog>
            </b-media>
          </b-card>
        </b-col>

        <b-col cols="8">
          <b-row>
            <b-col>
              <b-card title="Technical Metadata From System" tag="article" class="mb-2">
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
                        v-model="ddBusinessTermSelected"
                        :options="ddBusinessTermOptions"
                      ></b-form-select>
                    </b-form-group>

                    <b-form-group
                      horizontal
                      :label-cols="4"
                      breakpoint="md"
                      label="Business Term Description"
                    >
                      <text-wrap-dialog
                        :fulltext="selectedDetailsBusinessMetadata ? selectedDetailsBusinessMetadata.BT_DESCRIPTION : ''"
                      ></text-wrap-dialog>
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
                        v-model="ddBusinessAliasSelected"
                        :options="ddBusinessAliasOptions"
                      ></b-form-select>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="CDE (Yes/No)">
                      <text-wrap-dialog
                        :fulltext="selectedDetailsBusinessMetadata ? selectedDetailsBusinessMetadata.CDE_YES_NO : ''"
                      ></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group
                      horizontal
                      :label-cols="4"
                      breakpoint="md"
                      label="Data Quality standards | DQ Thresholds"
                    >
                      <text-wrap-dialog
                        :fulltext="selectedDetailsBusinessMetadata ? selectedDetailsBusinessMetadata.DQ_STANDARDS : ''"
                      ></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group
                      horizontal
                      :label-cols="4"
                      breakpoint="md"
                      label="Policy guidance (if any)"
                    >
                      <text-wrap-dialog
                        :fulltext="selectedDetailsBusinessMetadata ? selectedDetailsBusinessMetadata.POLICY_GUIDANCE : ''"
                      ></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group
                      horizontal
                      :label-cols="4"
                      breakpoint="md"
                      label="Mandatory (Yes/No)"
                    >
                      <text-wrap-dialog
                        :fulltext="selectedDetailsBusinessMetadata ? selectedDetailsBusinessMetadata.MANDATORY : ''"
                      ></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group
                      horizontal
                      :label-cols="4"
                      breakpoint="md"
                      label="Golden Source System"
                    >
                      <text-wrap-dialog
                        :fulltext="selectedDetailsBusinessMetadata ? selectedDetailsBusinessMetadata.GS_SYSTEM : ''"
                      ></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group
                      horizontal
                      :label-cols="4"
                      breakpoint="md"
                      label="Golden Source ITAM ID"
                    >
                      <text-wrap-dialog
                        :fulltext="selectedDetailsBusinessMetadata ? selectedDetailsBusinessMetadata.GS_ITAM_ID : ''"
                      ></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group
                      horizontal
                      :label-cols="4"
                      breakpoint="md"
                      label="Golden Source Table Name"
                    >
                      <text-wrap-dialog
                        :fulltext="selectedDetailsBusinessMetadata ? selectedDetailsBusinessMetadata.GS_TABLE_NAME : ''"
                      ></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group
                      horizontal
                      :label-cols="4"
                      breakpoint="md"
                      label="Golden Source Column Name"
                    >
                      <text-wrap-dialog
                        :fulltext="selectedDetailsBusinessMetadata ? selectedDetailsBusinessMetadata.GS_COLUMN_NAME : ''"
                      ></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group
                      horizontal
                      :label-cols="4"
                      breakpoint="md"
                      label="Target Golden Source"
                    >
                      <text-wrap-dialog
                        :fulltext="selectedDetailsBusinessMetadata ? selectedDetailsBusinessMetadata.TGS_SYSTEM : ''"
                      ></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group
                      horizontal
                      :label-cols="4"
                      breakpoint="md"
                      label="Target golden source ITAM ID"
                    >
                      <text-wrap-dialog
                        :fulltext="selectedDetailsBusinessMetadata ? selectedDetailsBusinessMetadata.TGS_ITAM_ID : ''"
                      ></text-wrap-dialog>
                    </b-form-group>
                  </b-form>
                </p>
              </b-card>
            </b-col>
          </b-row>

          <b-row>
            <b-col>
              <b-card title="Policy Related Information" tag="article" class="mb-2">
                <p class="card-text">
                  <b-form>
                    <b-form-group
                      horizontal
                      :label-cols="4"
                      breakpoint="md"
                      label="Information Asset Names"
                    >
                      <text-wrap-dialog
                        :fulltext="selectedDetailsBusinessMetadata ? selectedDetailsBusinessMetadata.INFO_ASSET_NAMES : ''"
                      ></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group
                      horizontal
                      :label-cols="4"
                      breakpoint="md"
                      label="Information Asset Descriptions"
                    >
                      <text-wrap-dialog
                        :fulltext="selectedDetailsBusinessMetadata ? selectedDetailsBusinessMetadata.INFO_ASSET_DESC : ''"
                      ></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group
                      horizontal
                      :label-cols="4"
                      breakpoint="md"
                      label="C - Confidentiality"
                    >
                      <text-wrap-dialog
                        :fulltext="selectedDetailsBusinessMetadata ? selectedDetailsBusinessMetadata.CONFIDENTIALITY : ''"
                      ></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="I - Integrity">
                      <text-wrap-dialog
                        :fulltext="selectedDetailsBusinessMetadata ? selectedDetailsBusinessMetadata.INTEGRITY : ''"
                      ></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group
                      horizontal
                      :label-cols="4"
                      breakpoint="md"
                      label="A - Availability"
                    >
                      <text-wrap-dialog
                        :fulltext="selectedDetailsBusinessMetadata ? selectedDetailsBusinessMetadata.AVAILABILITY : ''"
                      ></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group
                      horizontal
                      :label-cols="4"
                      breakpoint="md"
                      label="Overall CIA Rating"
                    >
                      <text-wrap-dialog
                        :fulltext="selectedDetailsBusinessMetadata ? selectedDetailsBusinessMetadata.OVERALL_CIA_RATING : ''"
                      ></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group
                      horizontal
                      :label-cols="4"
                      breakpoint="md"
                      label="Record Categories"
                    >
                      <text-wrap-dialog
                        :fulltext="selectedDetailsBusinessMetadata ? selectedDetailsBusinessMetadata.RECORD_CATEGORIES : ''"
                      ></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group
                      horizontal
                      :label-cols="4"
                      breakpoint="md"
                      label="PII Flag (Yes/No)"
                    >
                      <text-wrap-dialog
                        :fulltext="selectedDetailsBusinessMetadata ? selectedDetailsBusinessMetadata.PII_FLAG : ''"
                      ></text-wrap-dialog>
                    </b-form-group>
                  </b-form>
                </p>
              </b-card>
            </b-col>
          </b-row>

          <b-row>
            <b-col>
              <b-card title="Downstream Usage of Business Term" tag="article" class="mb-2">
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
                        v-model="ddDownstreamProcessNameSelected"
                        :options="ddDownstreamProcessNameOptions"
                      ></b-form-select>
                    </b-form-group>

                    <b-form-group
                      horizontal
                      :label-cols="4"
                      breakpoint="md"
                      label="Downstream Process Owner"
                    >
                      <text-wrap-dialog
                        :fulltext="selectedDetailsDownstreamUsage ? selectedDetailsDownstreamUsage.DOWNSTREAM_PROCESS_OWNER : ''"
                      ></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group
                      horizontal
                      :label-cols="4"
                      breakpoint="md"
                      label="Is data taken from Golden Source"
                    >
                      <text-wrap-dialog
                        :fulltext="selectedDetailsDownstreamUsage ? selectedDetailsDownstreamUsage.DATA_FROM_GS : ''"
                      ></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="System Name">
                      <text-wrap-dialog
                        :fulltext="selectedDetailsDownstreamUsage ? selectedDetailsDownstreamUsage.SYSTEM_NAME : ''"
                      ></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="ITAM ID">
                      <text-wrap-dialog
                        :fulltext="selectedDetailsDownstreamUsage ? selectedDetailsDownstreamUsage.ITAM_ID : ''"
                      ></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Table Name">
                      <text-wrap-dialog
                        :fulltext="selectedDetailsDownstreamUsage ? selectedDetailsDownstreamUsage.TABLE_NAME : ''"
                      ></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Column Name">
                      <text-wrap-dialog
                        :fulltext="selectedDetailsDownstreamUsage ? selectedDetailsDownstreamUsage.COLUMN_NAME : ''"
                      ></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group
                      horizontal
                      :label-cols="4"
                      breakpoint="md"
                      label="DQ Standards | Thresholds set by DDO"
                    >
                      <text-wrap-dialog
                        :fulltext="selectedDetailsDownstreamUsage ? selectedDetailsDownstreamUsage.DQ_STANDARDS : ''"
                      ></text-wrap-dialog>
                    </b-form-group>
                  </b-form>
                </p>
              </b-card>
            </b-col>
          </b-row>

          <b-row>
            <b-col>
              <b-card title="Business Term Residing In Other Systems" tag="article" class="mb-2">
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
                        v-model="ddSystemNameSelected"
                        :options="ddSystemNameOptions"
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
                        v-model="ddItamIdSelected"
                        :options="ddItamIdOptions"
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
                        v-model="ddTableNameSelected"
                        :options="ddTableNameOptions"
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
                        v-model="ddColumnNameSelected"
                        :options="ddColumnNameOptions"
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
import pageLoader from '../PageLoader.vue'
import textWrapDialog from "../TextWrapDialog.vue";

export default {
  components: {
    pageLoader, textWrapDialog
  },
  data() {
    return {
      firstload: true,
      showModal: this.$route.meta.showModal,
      selectedDetailsBusinessMetadata: null,
      selectedDetailsDownstreamUsage: null,
      selectedDetailsBTResiding: null,
      ddBusinessTermSelected: null,
      ddBusinessAliasSelected: null,
      ddDownstreamProcessNameSelected: null,
      ddSystemNameSelected: null,
      ddItamIdSelected: null,
      ddTableNameSelected: null,
      ddColumnNameSelected: null,
      excelFields: {
        "Data Domain": "selectedDetails.DATA_DOMAIN",
        "Sub Domain": "selectedDetails.SUB_DOMAIN",
        "Sub Domain Owner": "selectedDetails.SUB_DOMAIN_OWNER",
        "Bank ID": "selectedDetails.BANK_ID",
        "Business Term": "selectedDetails.BUSINESS_TERM",
        "Business Term Description": "selectedDetails.BT_DESCRIPTION",
        "Business Alias": "selectedDetails.BUSINESS_ALIAS",
        "CDE (Yes/No)": "selectedDetails.CDE_YES_NO",
        "Data Quality standards | DQ Thresholds": "selectedDetails.DQ_STANDARDS",
        "Policy guidance (if any)": "selectedDetails.POLICY_GUIDANCE",
        "Mandatory (Yes/No)": "selectedDetails.MANDATORY",
        "Golden Source System": "selectedDetails.GOLDEN_SYSTEM",
        "Golden Source ITAM ID": "selectedDetails.GOLDEN_ITAM_ID",
        "Golden Source Table Name": "selectedDetails.GOLDEN_TABLE_NAME",
        "Golden Source Column Name": "selectedDetails.GOLDEN_COLUMN_NAME",
        "Target Golden Source": "selectedDetails.ITAM_ID",
        "Target golden source ITAM ID": "selectedDetails.ITAM_ID",
        "Information Asset Names": "selectedDetails.INFO_ASSET_NAMES",
        "Information Asset Descriptions": "selectedDetails.INFO_ASSET_DESC",
        "C - Confidentiality": "selectedDetails.CONFIDENTIALITY",
        "I - Integrity": "selectedDetails.INTEGRITY",
        "A - Availability": "selectedDetails.AVAILABILITY",
        "Overall CIA Rating": "selectedDetails.OVERALL_CIA_RATING",
        "Record Categories": "selectedDetails.RECORD_CATEGORIES",
        "PII Flag (Yes/No)": "selectedDetails.PII_FLAG",
        "Downstream Process Name": "selectedDetails.DOWNSTREAM_PROCESS_NAME",
        "Downstream Process Owner": "selectedDetails.DOWNSTREAM_PROCESS_OWNER",
        "Is data taken from Golden Source": "selectedDetails.ASDF",
        "System Name": "selectedDetails.SYSTEM_NAME",
        "ITAM ID": "selectedDetails.ITAM_ID",
        "Table Name": "selectedDetails.TABLE_NAME",
        "Column Name": "selectedDetails.COLUMN_NAME",
        "DQ Standards | Thresholds set by DDO": "selectedDetails.DQ_STANDARDS_BY_DDO",
        "System Name": "selectedDetails.SYSTEM_NAME_DD",
        "ITAM ID": "selectedDetails.ITAM_ID_DD",
        "Table Name": "selectedDetails.TABLE_NAME_DD",
        "Column Name": "selectedDetails.COLUMN_NAME_DD"
      }
    };
  },
  computed: {
    ...mapState({
      ddomy: state => state.ddomy.all
    }),
    ddBusinessTermOptions() {
      return _.uniq(
        _.map(this.ddomy.DDSourceBusinessMetadata, "BUSINESS_TERM")
      );
    },
    ddBusinessAliasOptions() {
      var self = this;
      var filtered = _.filter(self.ddomy.DDSourceBusinessMetadata, v => {
        return v.BUSINESS_TERM == self.ddBusinessTermSelected;
      });

      return _.uniq(_.map(filtered, "BUSINESS_ALIAS"));
    },
    ddDownstreamProcessNameOptions() {
      return _.uniq(
        _.map(this.ddomy.DetailsDownstreamUsage, "DOWNSTREAM_PROCESS_NAME")
      );
      // var self = this;
      // var filtered = _.filter(self.ddomy.DetailsDownstreamUsage, (v) => {
      //   return v.BUSINESS_TERM == self.ddBusinessTermSelected
      //     && v.BUSINESS_ALIAS == self.ddBusinessAliasSelected;
      // });

      // return _.uniq(_.map(filtered, "DOWNSTREAM_PROCESS_NAME"));
    },
    ddSystemNameOptions() {
      return _.uniq(_.map(this.ddomy.DetailsBTResiding, "SYSTEM_NAME"));
      // var self = this;
      // var filtered = _.filter(self.ddomy.DDSource, (v) => {
      //   return v.BUSINESS_TERM == self.ddBusinessTermSelected
      //     && v.BUSINESS_ALIAS == self.ddBusinessAliasSelected
      //     && v.DOWNSTREAM_PROCESS_NAME == self.ddDownstreamProcessNameSelected;
      // });

      // return _.uniq(_.map(filtered, "SYSTEM_NAME_DD"));
    },
    ddItamIdOptions() {
      var self = this;
      var filtered = _.filter(self.ddomy.DetailsBTResiding, v => {
        return v.SYSTEM_NAME == self.ddSystemNameSelected;
      });

      return _.uniq(_.map(filtered, "ITAM_ID"));
    },
    ddTableNameOptions() {
      var self = this;
      var filtered = _.filter(self.ddomy.DetailsBTResiding, v => {
        return (
          v.SYSTEM_NAME == self.ddSystemNameSelected &&
          v.ITAM_ID == self.ddItamIdSelected
        );
      });

      return _.uniq(_.map(filtered, "TABLE_NAME"));
    },
    ddColumnNameOptions() {
      var self = this;
      var filtered = _.filter(self.ddomy.DetailsBTResiding, v => {
        return (
          v.SYSTEM_NAME == self.ddSystemNameSelected &&
          v.ITAM_ID == self.ddItamIdSelected &&
          v.TABLE_NAME == self.ddTableNameSelected
        );
      });

      return _.uniq(_.map(filtered, "COLUMN_NAME"));
    },
    exportDatas() {
      if (
        this.selectedDetailsBusinessMetadata &&
        this.selectedDetailsDownstreamUsage &&
        this.selectedDetailsBTResiding
      ) {
        return [
          {
            selectedDetailsBusinessMetadata: this.selectedDetailsBusinessMetadata,
            selectedDetailsDownstreamUsage: this.selectedDetailsDownstreamUsage,
            selectedDetailsBTResiding: this.selectedDetailsBTResiding
          }
        ];
      } else {
        return [];
      }
    }
  },
  watch: {
    "$route.meta"({ showModal }) {
      this.showModal = showModal;
    },
    ddBusinessTermSelected() {
      if (this.firstload) return;

      if (this.ddBusinessAliasOptions[0])
        this.ddBusinessAliasSelected = this.ddBusinessAliasOptions[0];
      else this.ddBusinessAliasSelected = "";
    },
    ddBusinessAliasSelected() {
      if (this.firstload) return;

      if (!this.firstload) {
        var param = {
          BusinessTerm: this.ddBusinessTermSelected.toString()
        };

        if (this.ddBusinessAliasSelected && this.ddBusinessAliasSelected != "NA") {
          param.BusinessAlias = this.ddBusinessAliasSelected.toString();
        }

        this.runGetDetails(param);
      }
    },
    ddDownstreamProcessNameSelected() {
      if (this.firstload) return;

      if (!this.firstload) {
        var param = {
          BusinessTerm: this.ddBusinessTermSelected.toString()
        };

        if (this.ddBusinessAliasSelected && this.ddBusinessAliasSelected != "NA") {
          param.BusinessAlias = this.ddBusinessAliasSelected.toString();
        }

        if (this.ddDownstreamProcessNameSelected && this.ddDownstreamProcessNameSelected != "NA") {
          param.DownstreamProcessName = this.ddDownstreamProcessNameSelected.toString();
        }

        this.runGetDetails(param);
      }
    },
    ddSystemNameSelected() {
      if (this.firstload) return;

      if (!this.firstload) {
        var param = {
          BusinessTerm: this.ddBusinessTermSelected.toString()
        };

        if (this.ddBusinessAliasSelected && this.ddBusinessAliasSelected != "NA") {
          param.BusinessAlias = this.ddBusinessAliasSelected.toString();
        }

        if (this.ddDownstreamProcessNameSelected && this.ddDownstreamProcessNameSelected != "NA") {
          param.DownstreamProcessName = this.ddDownstreamProcessNameSelected.toString();
        }

        if (this.ddSystemNameSelected && this.ddSystemNameSelected != "NA") {
          param.SystemName = this.ddSystemNameSelected.toString();
        }

        this.runGetDetails(param);
      }
    },
    ddItamIdSelected() {
      if (this.firstload) return;

      if (!this.firstload) {
        var param = {
          BusinessTerm: this.ddBusinessTermSelected.toString()
        };

        if (this.ddBusinessAliasSelected && this.ddBusinessAliasSelected != "NA") {
          param.BusinessAlias = this.ddBusinessAliasSelected.toString();
        }

        if (this.ddDownstreamProcessNameSelected && this.ddDownstreamProcessNameSelected != "NA") {
          param.DownstreamProcessName = this.ddDownstreamProcessNameSelected.toString();
        }

        if (this.ddSystemNameSelected && this.ddSystemNameSelected != "NA") {
          param.SystemName = this.ddSystemNameSelected.toString();
        }

        if (this.ddItamIdSelected && this.ddItamIdSelected != "NA") {
          param.ItamId = this.ddItamIdSelected.toString();
        }

        this.runGetDetails(param);
      }
    },
    ddTableNameSelected() {
      if (this.firstload) return;

      if (!this.firstload) {
        var param = {
          BusinessTerm: this.ddBusinessTermSelected.toString()
        };

        if (this.ddBusinessAliasSelected && this.ddBusinessAliasSelected != "NA") {
          param.BusinessAlias = this.ddBusinessAliasSelected.toString();
        }

        if (this.ddDownstreamProcessNameSelected && this.ddDownstreamProcessNameSelected != "NA") {
          param.DownstreamProcessName = this.ddDownstreamProcessNameSelected.toString();
        }

        if (this.ddSystemNameSelected && this.ddSystemNameSelected != "NA") {
          param.SystemName = this.ddSystemNameSelected.toString();
        }

        if (this.ddItamIdSelected && this.ddItamIdSelected != "NA") {
          param.ItamId = this.ddItamIdSelected.toString();
        }

        if (this.ddTableNameSelected && this.ddTableNameSelected != "NA") {
          param.TableName = this.ddTableNameSelected.toString();
        }

        this.runGetDetails(param);
      }
    },
    ddColumnNameSelected() {
      if (this.firstload) return;

      if (!this.firstload) {
        var param = {
          BusinessTerm: this.ddBusinessTermSelected.toString()
        };

        if (this.ddBusinessAliasSelected && this.ddBusinessAliasSelected != "NA") {
          param.BusinessAlias = this.ddBusinessAliasSelected.toString();
        }

        if (this.ddDownstreamProcessNameSelected && this.ddDownstreamProcessNameSelected != "NA") {
          param.DownstreamProcessName = this.ddDownstreamProcessNameSelected.toString();
        }

        if (this.ddSystemNameSelected && this.ddSystemNameSelected != "NA") {
          param.SystemName = this.ddSystemNameSelected.toString();
        }

        if (this.ddItamIdSelected && this.ddItamIdSelected != "NA") {
          param.ItamId = this.ddItamIdSelected.toString();
        }

        if (this.ddTableNameSelected && this.ddTableNameSelected != "NA") {
          param.TableName = this.ddTableNameSelected.toString();
        }
        
        if (this.ddColumnNameSelected && this.ddColumnNameSelected != "NA") {
          param.ColumnName = this.ddColumnNameSelected.toString();
        }

        this.runGetDetails(param);
      }
    }
  },
  mounted() {
    this.$refs.modalDetails.show();

    var param = {};
    this.runGetDetails(param);
  },
  methods: {
    ...mapActions("ddomy", {
      getDetails: "getDetails"
    }),
    handleClose() {
      this.$router.go(-1);
    },
    runGetDetails(param) {
      var self = this;

      param.Which = self.$route.name;
      param.Left = parseInt(self.$route.params.system).toString();
      param.Right = parseInt(self.$route.params.details).toString();

      return self.getDetails(param).then(
        res => {
          this.firstload = true;

          if (self.ddomy.DetailsBusinessMetadata.length > 0) {
            var tmp = self.ddomy.DetailsBusinessMetadata[0].Values[0];

            self.selectedDetailsBusinessMetadata = {};
            _.each(Object.keys(tmp), function(v, i) {
              self.selectedDetailsBusinessMetadata[v] = _.uniq(
                _.map(self.ddomy.DetailsBusinessMetadata[0].Values, val => val[v] ? val[v].toString().trim() : "").filter(Boolean)
              );
              
              if(v == "SUBDOMAIN_OWNER" || v == "BANK_ID"){
                self.selectedDetailsBusinessMetadata[v] = self.selectedDetailsBusinessMetadata[v].join('; ');
              } else {
                self.selectedDetailsBusinessMetadata[v] = self.selectedDetailsBusinessMetadata[v].join(', ');
              }
            });

            // make falsy NA
            Object.keys(self.selectedDetailsBusinessMetadata).forEach(val => {
              self.selectedDetailsBusinessMetadata[val] = !!self.selectedDetailsBusinessMetadata[val].trim()
                ? self.selectedDetailsBusinessMetadata[val]
                : "NA";
            });

            setTimeout(() => {
              self.ddBusinessTermSelected = self.selectedDetailsBusinessMetadata.BUSINESS_TERM;
              self.ddBusinessAliasSelected = self.selectedDetailsBusinessMetadata.BUSINESS_ALIAS;

              setTimeout(() => {
                this.firstload = false;
              }, 100);
            }, 100);
          } else {
            this.selectedDetailsBusinessMetadata = null;
          }

          if (self.ddomy.DetailsDownstreamUsage.length > 0) {
            var tmp = self.ddomy.DetailsDownstreamUsage[0].Values[0];

            self.selectedDetailsDownstreamUsage = {};
            _.each(Object.keys(tmp), function(v, i) {
              self.selectedDetailsDownstreamUsage[v] = _.uniq(
                _.map(self.ddomy.DetailsDownstreamUsage[0].Values, val => val[v] ? val[v].toString().trim() : "").filter(Boolean)
              ).join(", ");
            });

            // make falsy NA
            Object.keys(self.selectedDetailsDownstreamUsage).forEach(val => {
              self.selectedDetailsDownstreamUsage[val] = !!self.selectedDetailsDownstreamUsage[val].trim()
                ? self.selectedDetailsDownstreamUsage[val]
                : "NA";
            });

            setTimeout(() => {
              self.ddDownstreamProcessNameSelected = self.selectedDetailsDownstreamUsage.DOWNSTREAM_PROCESS_NAME;

              setTimeout(() => {
                this.firstload = false;
              }, 100);
            }, 100);
          } else {
            this.selectedDetailsDownstreamUsage = null;
          }

          if (self.ddomy.DetailsBTResiding.length > 0) {
            var tmp = self.ddomy.DetailsBTResiding[0].Values[0];

            self.selectedDetailsBTResiding = {};
            _.each(Object.keys(tmp), function(v, i) {
              self.selectedDetailsBTResiding[v] = _.uniq(
                _.map(self.ddomy.DetailsBTResiding[0].Values, val => val[v] ? val[v].toString().trim() : "").filter(Boolean)
              ).join(", ");
            });

            // make falsy NA
            Object.keys(self.selectedDetailsBTResiding).forEach(val => {
              self.selectedDetailsBTResiding[val] = !!self.selectedDetailsBTResiding[val].trim()
                ? self.selectedDetailsBTResiding[val]
                : "NA";
            });

            setTimeout(() => {
              self.ddSystemNameSelected = self.selectedDetailsBTResiding.SYSTEM_NAME;
              self.ddItamIdSelected = self.selectedDetailsBTResiding.ITAM_ID;
              self.ddTableNameSelected = self.selectedDetailsBTResiding.TABLE_NAME;
              self.ddColumnNameSelected = self.selectedDetailsBTResiding.COLUMN_NAME;

              setTimeout(() => {
                this.firstload = false;
              }, 100);
            }, 100);
          } else {
            this.selectedDetailsBTResiding = null;
          }
        },
        err => err
      );
    }
  }
};
</script>