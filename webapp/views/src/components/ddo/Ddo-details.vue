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
legend.col-form-label, label.col-form-labell {
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
              <p v-html="selectedDetails.DATA_DOMAIN"></p>
            </b-media>

            <b-media class="left-card-media">
              <h6 class="left-card-title">Sub Domain</h6>
              <p v-html="selectedDetails.SUB_DOMAIN"></p>
            </b-media>

            <b-media class="left-card-media">
              <h6 class="left-card-title">Sub Domain Owner</h6>
              <p v-html="selectedDetails.SUB_DOMAIN_OWNER"></p>
            </b-media>

            <b-media class="left-card-media">
              <h6 class="left-card-title">Bank ID</h6>
              <p v-html="selectedDetails.BANK_ID"></p>
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
                      <p v-html="selectedDetails.BUSINESS_TERM_DESCRIPTION"></p>
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
                      <p v-html="selectedDetails.CDE"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Data Quality standards | DQ Thresholds">
                      <p v-html="selectedDetails.DQ_STANDARDS"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Policy guidance (if any)">
                      <p v-html="selectedDetails.POLICY_GUIDANCE"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Mandatory (Yes/No)">
                      <p v-html="selectedDetails.MANDATORY"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Golden Source System">
                      <p v-html="selectedDetails.GOLDEN_SOURCE_SYSTEM_ID"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Golden Source ITAM ID">
                      <p v-html="selectedDetails.ITAM_ID"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Golden Source Table Name">
                      <p v-html="selectedDetails.GOLDEN_SOURCE_TABLE_NAME"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Golden Source Column Name">
                      <p v-html="selectedDetails.GOLDEN_SOURCE_COLUMN_NAME"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Target Golden Source">
                      <p v-html="selectedDetails.ITAM_ID"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Target golden source ITAM ID">
                      <p v-html="selectedDetails.ITAM_ID"></p>
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
                      <p v-html="selectedDetails.INFO_ASSET_NAME"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Information Asset Descriptions">
                      <p v-html="selectedDetails.INFO_ASSET_DESC"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="C - Confidentiality">
                      <p v-html="selectedDetails.CONFIDENTIALITY"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="I - Integrity">
                      <p v-html="selectedDetails.INTEGRITY"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="A - Availability">
                      <p v-html="selectedDetails.AVAILABILITY"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Overall CIA Rating">
                      <p v-html="selectedDetails.OVERALL_CIA_RATING"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Record Categories">
                      <p v-html="selectedDetails.RECORD_CATEGORY"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="PII Flag (Yes/No)">
                      <p v-html="selectedDetails.PII_FLAG"></p>
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
                      <p v-html="selectedDetails.DOWNSTREAM_PROCESS_OWNER"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Is data taken from Golden Source">
                      <p v-html="selectedDetails.ASDF"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="System Name">
                      <p v-html="selectedDetails.SYSTEM_NAME"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="ITAM ID">
                      <p v-html="selectedDetails.ITAM_ID"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Table Name">
                      <p v-html="selectedDetails.GOLDEN_SOURCE_TABLE_NAME"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Column Name">
                      <p v-html="selectedDetails.COLUMN_NAME"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="DQ Standards | Thresholds set by DDO">
                      <p v-html="selectedDetails.DDO_THRESHOLD"></p>
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

export default {
  data() {
    return {
      dialog: false,
      showModal: this.$route.meta.showModal,
      selectedDetails: null,
      selectedColumn: null,
      ddTable: {
        selected: null
      },
      ddColumn: {
        selected: null
      },
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
    exportDatas() {
      if (this.selectedDetails) {
        return [
          {
            selectedDetails: this.selectedDetails,
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
    "ddomy.detailsSource"() {
      if (this.ddomy.detailsSource.length > 0)
        this.selectedDetails = this.ddomy.detailsSource[0];
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
  },
  methods: {
    ...mapActions("ddomy", {
      getDetails: "getDetails"
    }),
    handleClose() {
      this.$router.go(-1);
    }
  }
};
</script>