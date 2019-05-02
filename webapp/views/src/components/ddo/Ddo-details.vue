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
              <text-wrap-dialog :fulltext="selectedDetails ? selectedDetails.DATA_DOMAIN : ''"></text-wrap-dialog>
            </b-media>

            <b-media class="left-card-media">
              <h6 class="left-card-title">Sub Domain</h6>
              <text-wrap-dialog :fulltext="selectedDetails ? selectedDetails.SUB_DOMAINS : ''"></text-wrap-dialog>
            </b-media>

            <b-media class="left-card-media">
              <h6 class="left-card-title">Sub Domain Owner</h6>
              <text-wrap-dialog :fulltext="selectedDetails ? selectedDetails.SUB_DOMAIN_OWNER : ''"></text-wrap-dialog>
            </b-media>

            <b-media class="left-card-media">
              <h6 class="left-card-title">Bank ID</h6>
              <text-wrap-dialog :fulltext="selectedDetails ? selectedDetails.BANK_ID : ''"></text-wrap-dialog>
            </b-media>
          </b-card>
        </b-col>

        <b-col cols="8">
          <b-row>
            <b-col>
              <b-card title="" tag="article" class="mb-2">
                <p class="card-text">
                  <b-form>
                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Business Term">
                      <text-wrap-dialog :fulltext="selectedDetails ? selectedDetails.BT_NAME : ''"></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Description">
                      <text-wrap-dialog :fulltext="selectedDetails ? selectedDetails.BT_DESCRIPTION : ''"></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Policy Guidance (if any)">
                      <text-wrap-dialog :fulltext="selectedDetails ? selectedDetails.POLICY_GUIDANCE : ''"></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Business Rules (Documents, Domain Values, Dérivation Logic, etc.)">
                      <text-wrap-dialog :fulltext="selectedDetails ? selectedDetails.BUSINESS_RULES : ''"></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Mandatory (Y/N)">
                      <text-wrap-dialog :fulltext="selectedDetails ? selectedDetails.MANDATORY : ''"></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Golden Source(s) ITAM ID">
                      <text-wrap-dialog :fulltext="selectedDetails ? selectedDetails.GS_ITAM_ID : ''"></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Golden Source(s)">
                      <text-wrap-dialog :fulltext="selectedDetails ? selectedDetails.GS_SYSTEM_NAME : ''"></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Target Golden Source ITAM Id">
                      <text-wrap-dialog :fulltext="selectedDetails ? selectedDetails.TGS_ITAM_ID : ''"></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Target Golden Source">
                      <text-wrap-dialog :fulltext="selectedDetails ? selectedDetails.TGS_SYSTEM_NAME : ''"></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Remarks">
                      <text-wrap-dialog :fulltext="selectedDetails ? selectedDetails.REMARKS : ''"></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="System ITAM ID">
                      <text-wrap-dialog :fulltext="selectedDetails ? selectedDetails.ITAM_ID : ''"></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="System Name" label-for="systemname">
                      <b-form-select id="systemname" class="col-8" v-model="ddSystemNameSelected" :options="ddSystemNameOptions"></b-form-select>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Table Name" label-for="tablename">
                      <b-form-select id="tablename" class="col-8" v-model="ddTableNameSelected" :options="ddTableNameOptions"></b-form-select>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Column Name">
                      <text-wrap-dialog :fulltext="selectedDetails ? selectedDetails.COLUMN_NAME : ''"></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Business Alias Name" label-for="businessaliasname">
                      <b-form-select id="businessaliasname" class="col-8" v-model="ddBusinessAliasNameSelected" :options="ddBusinessAliasNameOptions"></b-form-select>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="CDE (Yes/No)">
                      <text-wrap-dialog :fulltext="selectedDetails ? selectedDetails.CDE : ''"></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Golden source (Yes / No)">
                      <text-wrap-dialog :fulltext="selectedDetails ? selectedDetails.GOLDEN_SOURCE : ''"></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="PII">
                      <text-wrap-dialog :fulltext="selectedDetails ? selectedDetails.PII_FLAG : ''"></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="CIA ratings (overall)">
                      <text-wrap-dialog :fulltext="selectedDetails ? selectedDetails.CIA_RATING : ''"></text-wrap-dialog>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Record Category">
                      <text-wrap-dialog :fulltext="selectedDetails ? selectedDetails.RECORD_CATEGORY : ''"></text-wrap-dialog>
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
      ddSystemNameSelected: null,
      ddTableNameSelected: null,
      ddBusinessAliasNameSelected: null,
      excelFields: {
        'Data Domain': "selectedDetails.DATA_DOMAIN",
        'Sub Domain': "selectedDetails.SUB_DOMAINS",
        'Sub Domain Owner': "selectedDetails.SUB_DOMAIN_OWNER",
        'Bank ID': "selectedDetails.BANK_ID",
        'Business Term': "selectedDetails.BT_NAME",
        'Description': "selectedDetails.BT_DESCRIPTION",
        'Policy Guidance (if any)': "selectedDetails.POLICY_GUIDANCE",
        'Business Rules (Documents, Domain Values, Dérivation Logic, etc.)': "selectedDetails.BUSINESS_RULES",
        'Mandatory (Y/N)': "selectedDetails.MANDATORY",
        'Golden Source(s) ITAM ID': "selectedDetails.GS_ITAM_ID",
        'Golden Source(s)': "selectedDetails.GS_SYSTEM_NAME",
        'Target Golden Source ITAM Id': "selectedDetails.TGS_ITAM_ID",
        'Target Golden Source': "selectedDetails.TGS_SYSTEM_NAME",
        'Remarks': "selectedDetails.REMARKS",
        'System ITAM ID': "selectedDetails.ITAM_ID",
        'System Name': "selectedDetails.SYSTEM_NAME",
        'Table Name': "selectedDetails.TABLE_NAME",
        'Column Name': "selectedDetails.COLUMN_NAME",
        'Business Alias Name': "selectedDetails.ALIAS_NAME",
        'CDE (Yes/No)': "selectedDetails.CDE",
        'Golden source (Yes / No)': "selectedDetails.GOLDEN_SOURCE",
        'PII': "selectedDetails.PII_FLAG",
        'CIA ratings (overall)': "selectedDetails.CIA_RATING",
        'Record Category': "selectedDetails.RECORD_CATEGORY",
      }
    }
  },
  computed: {
    ...mapState({
      ddomy: state => state.ddomy.all
    }),
    ddSystemNameOptions () {
      return _.uniq(_.map(this.ddomy.DDSource, (v) => v.SYSTEM_NAME.toString().trim())).filter(Boolean);
    },
    ddTableNameOptions () {
      var self = this;
      var filtered = _.filter(self.ddomy.DDSource, (v) => {
        return v.SYSTEM_NAME == self.ddSystemNameSelected;
      });

      return _.uniq(_.map(filtered, (v) => v.TABLE_NAME.toString().trim())).filter(Boolean);
    },
    ddBusinessAliasNameOptions () {
      var self = this;
      var filtered = _.filter(self.ddomy.DDSource, (v) => {
        return v.SYSTEM_NAME == self.ddSystemNameSelected 
          && v.TABLE_NAME == self.ddTableNameSelected;
      });
      
      return _.uniq(_.map(filtered, (v) => v.ALIAS_NAME.toString().trim())).filter(Boolean);
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
    ddSystemNameSelected () {
      if(this.firstload) return;

      if (!this.firstload) {
        var param = {};

        if (this.ddSystemNameSelected && this.ddSystemNameSelected != "NA") {
          param.SystemName = this.ddSystemNameSelected.toString();
        }

        this.runGetDetails(param);
      }
    },
    ddTableNameSelected () {
      if(this.firstload) return;
      
      if (!this.firstload) {
        var param = {};

        if (this.ddSystemNameSelected && this.ddSystemNameSelected != "NA") {
          param.SystemName = this.ddSystemNameSelected.toString();
        }

        if (this.ddTableNameSelected && this.ddTableNameSelected != "NA") {
          param.ImmItamID = this.ddTableNameSelected.toString();
        }

        this.runGetDetails(param);
      }
    },
    ddBusinessAliasNameSelected () {
      if(this.firstload) return;
      
      if (!this.firstload) {
        var param = {};

        if (this.ddSystemNameSelected && this.ddSystemNameSelected != "NA") {
          param.SystemName = this.ddSystemNameSelected.toString();
        }

        if (this.ddTableNameSelected && this.ddTableNameSelected != "NA") {
          param.TableName = this.ddTableNameSelected.toString();
        }

        if (this.ddBusinessAliasNameSelected && this.ddBusinessAliasNameSelected != "NA") {
          param.BusinessAliasName = this.ddBusinessAliasNameSelected.toString();
        }

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
    ...mapActions("ddomy", {
      getDetails: "getDetails"
    }),
    handleClose () {
      this.$router.go(-1)
    },
    runGetDetails (param){
      var self = this;

      param.Which = self.$route.name;
      param.Subdomain = self.$route.params.subdomain.toString();
      param.BTname = self.$route.params.btname.toString();

      return self.getDetails(param).then(
        res => {
          this.firstload = true;

          if (self.ddomy.detailsSource.length > 0){
            var tmp = self.ddomy.detailsSource[0].Values[0];

            self.selectedDetails = {}
            _.each(Object.keys(tmp), function(v, i){
                self.selectedDetails[v] = _.uniq(
                  _.map(self.ddomy.detailsSource[0].Values, (val) => val[v] ? val[v].toString().trim() : "").filter(Boolean)
                );
                if(v == "DATASET_CUSTODIAN" || v == "BANK_ID"){
                  self.selectedDetails[v] = self.selectedDetails[v].join('; ');
                } else {
                  self.selectedDetails[v] = self.selectedDetails[v].join(', ');
                }
            });
            
            // interrupt
            var doInterrupt = (v, cond, expectedres) => v == cond[1] ? expectedres[1] : (v == cond[0] ? expectedres[0] : v);

            self.selectedDetails.MANDATORY = doInterrupt(self.selectedDetails.MANDATORY, [0, 1], ["No", "Yes"]);
            self.selectedDetails.CDE = doInterrupt(self.selectedDetails.CDE, [0, 1], ["No", "Yes"]);

            // make falsy NA
            Object.keys(self.selectedDetails).forEach((val) => {
              self.selectedDetails[val] = !!self.selectedDetails[val].trim() ? self.selectedDetails[val] : "NA";
            });
            
            self.ddomy.DDSource.map(function(v){
              Object.keys(v).forEach(function(key){
                v[key] = v[key] ? v[key] : "NA"
              })

              return v
            })
            
            setTimeout(() => {
              self.ddSystemNameSelected = self.selectedDetails.SYSTEM_NAME.split(", ")[0];
              self.ddTableNameSelected = self.selectedDetails.TABLE_NAME.split(", ")[0];
              self.ddBusinessAliasNameSelected = self.selectedDetails.ALIAS_NAME.split(", ")[0];

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