<template>
  <b-modal hide-footer header-class="modal-details-header setbackground" body-class="setbackground" v-if="showModal" id="modal-details" ref="modalDetails" size="lg" @hidden="handleClose">
    <b-container fluid class="row-kasijarak">
      <b-row>
        <b-col> 
          <b-btn size="sm" class="float-right" variant="success">Export</b-btn>
        </b-col>
      </b-row>
    <!-- </b-container>

    <b-container> -->
      <b-row>
        <b-col cols="4"> 
          <b-card tag="article" class="mb-2" v-if="selectedSystem && selectedTable.Columns.length > 0">
            <b-media>
              <h6>System Name</h6>
              <p v-html="selectedSystem.System_Name"></p>
            </b-media>
            
            <b-media>
              <h6>ITAM ID</h6>
              <p v-html="selectedSystem.ITAM_ID"></p>
            </b-media>
            
            <b-media>
              <h6>Dataset Custodian</h6>
              <p></p>
            </b-media>
            
            <b-media>
              <h6>Bank ID</h6>
              <p></p>
            </b-media>
            
            <b-media>
              <h6>Business Alias Name</h6>
              <p v-html="selectedTable.Columns[0].Alias_Name"></p>
            </b-media>
            
            <b-media>
              <h6>Table Name</h6>
              <p v-html="selectedTable.Name"></p>
            </b-media>
            
            <b-media>
              <h6>Column Name</h6>
              <p v-html="selectedTable.Columns[0].Name"></p>
            </b-media>
          </b-card>
        </b-col>

        <b-col cols="8"> 
          <b-row>
            <b-col>
              <b-card title="Technical Metadata From System" tag="article" class="mb-2" v-if="selectedColumn">
                <p class="card-text">
                  <b-form>
                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="System Name" label-for="systemName">
                      <b-form-select id="systemName" class="col-8" v-model="ddSystem.selected" :options="ddSystemOptions">
                      </b-form-select>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Column Name" label-for="columnName">
                      <b-form-select id="columnName" class="col-8" v-model="ddColumn.selected" :options="ddColumnOptions"></b-form-select>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Screen Label Name*" label-for="screenLabelName">
                      <b-form-select id="screenLabelName" class="col-8"></b-form-select>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Screen Label Name*">
                      <p>This field describes the full name of the client and used for capturing lengthy names where the customer name field is unable to accommodat. It is also used for sub-funds. [more]</p>
                    </b-form-group>
                    
                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="CDE (yes/no)">
                      <p v-html="selectedColumn.CDE"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Status*">
                      <p v-html="selectedColumn.Status"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Data Type">
                      <p v-html="selectedColumn.Data_Type"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Data Format">
                      <p v-html="selectedColumn.Data_Format"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Data Length">
                      <p v-html="selectedColumn.Data_Length"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Example">
                      <p v-html="selectedColumn.Example"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Derived (Yes/No)*">
                      <p v-html="selectedColumn.Derived"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Derivation logic*">
                      <p v-html="selectedColumn.Derivation_Logic"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Sourced from Upstream (Yes/No)*">
                      <p v-html="selectedColumn.Sourced_from_Upstream"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="System Checks*">
                      <p v-html="selectedColumn.System_Checks"></p>
                    </b-form-group>
                  </b-form>
                </p>
              </b-card>
            </b-col>
          </b-row>

          <b-row>
            <b-col>
              <b-card title="Business Metadata From Domain" tag="article" class="mb-2" v-if="selectedColumn">
                <p class="card-text">
                  <b-form>
                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Domain">
                      <p v-html="selectedColumn.BusinessTerms.SubCategory.Category.Name"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Sub Domain">
                      <p v-html="selectedColumn.BusinessTerms.SubCategory.Name"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Domain Owner">
                      <p></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Business Term*">
                      <p v-html="selectedColumn.BusinessTerms.BT_Name"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Business Term Description">
                      <p v-html="selectedColumn.BusinessTerms.Description"></p>
                    </b-form-group>
                  </b-form>
                </p>
              </b-card>
            </b-col>
          </b-row>

          <b-row>
            <b-col>
              <b-card title="Policy Related Information" tag="article" class="mb-2" v-if="selectedColumn">
                <p class="card-text">
                  <b-form>
                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Information Asset Names">
                      <p v-html="selectedColumn.BusinessTerms.Policy.Info_Asset_Name"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Information Asset Description">
                      <p v-html="selectedColumn.BusinessTerms.Policy.Description"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="C - Confidentiality">
                      <p v-html="selectedColumn.BusinessTerms.Policy.Confidentiality"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="I - Integrity">
                      <p v-html="selectedColumn.BusinessTerms.Policy.Integrity"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="A - Availability">
                      <p v-html="selectedColumn.BusinessTerms.Policy.Availability"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Overall CIA Rating">
                      <p v-html="selectedColumn.BusinessTerms.Policy.Overall_CIA_Rating"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Record Categories">
                      <p v-html="selectedColumn.BusinessTerms.Policy.Record_Category"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="PII Flag">
                      <p v-html="selectedColumn.BusinessTerms.Policy.PII_Flag"></p>
                    </b-form-group>
                  </b-form>
                </p>
              </b-card>
            </b-col>
          </b-row>

          <b-row>
            <b-col>
              <b-card title="Interfaces" tag="article" class="mb-2">
                <p class="card-text">
                  <b-form>
                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Immediate Preceding System*">
                      <p v-html="selectedColumn.Imm_Prec_System_ID"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Immediate Succeeding System*">
                      <p v-html="selectedColumn.Imm_Succ_System_ID"></p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="DQ Standards | Threshold*">
                      <p v-html="selectedColumn.Threshold"></p>
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
      selectedSystem: null,
      selectedTable: null,
      selectedColumn: null,
      ddSystem: {
        selected: null,
      },
      ddColumn: {
        selected: null,
      }
    }
  },
  computed: {
    ...mapState({
      dscmy: state => state.dscmy.all
    }),
    count () {
      return this.dscmy.tableSource.length
    },
    ddSystemOptions () {
      this.ddSystem.selected = this.selectedSystem ? this.selectedSystem.System_Name : null;
      return this.selectedSystem ? [{ value: this.selectedSystem.System_Name, text: this.selectedSystem.System_Name }] : [];
    },
    ddColumnOptions () {
      return this.selectedTable ? _.map(this.selectedTable.Columns, function(v) { return { value: v.ID, text: v.Name } }) : [];
    },
  },
  watch: {
    '$route.meta' ({showModal}) {
      this.showModal = showModal;
    },
    count (newCount, oldCount) {
      this.selectedSystem = _.find(this.dscmy.systemsSource, ['ID', parseInt(this.$route.params.system)])
      this.selectedTable = _.find(this.dscmy.tableSource, ['ID', parseInt(this.$route.params.details)])
    },
    selectedTable (){
      this.ddColumn.selected = this.selectedTable ? (this.selectedTable.Columns[0] ? this.selectedTable.Columns[0].ID : null) : null;
      this.selectedColumn = this.selectedTable ? _.find(this.selectedTable.Columns, ['ID', parseInt(this.ddColumn.selected)]) : null;
    }
  },
  mounted() {
    this.$refs.modalDetails.show();
    
    this.selectedSystem = _.find(this.dscmy.systemsSource, ['ID', parseInt(this.$route.params.system)])
    this.selectedTable = _.find(this.dscmy.tableSource, ['ID', parseInt(this.$route.params.details)])
  },
  methods: {
    handleClose () {
      this.$router.go(-1)
    }
  },
}
</script>