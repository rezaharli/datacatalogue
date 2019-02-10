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
          <b-card tag="article" class="mb-2">
            <b-media>
              <h6>System Name</h6>
              <p>SCI</p>
            </b-media>
            
            <b-media>
              <h6>ITAM ID</h6>
              <p>14768</p>
            </b-media>
            
            <b-media>
              <h6>Dataset Custodian</h6>
              <p>Samuel Rajakumar</p>
            </b-media>
            
            <b-media>
              <h6>Bank ID</h6>
              <p>1159683</p>
            </b-media>
            
            <b-media>
              <h6>Business Alias Name</h6>
              <p>Actual Legal Name</p>
            </b-media>
            
            <b-media>
              <h6>Table Name</h6>
              <p>P44_LE_MAIN_PROFILE</p>
            </b-media>
            
            <b-media>
              <h6>Column Name</h6>
              <p>LMP_ACTUAL_LEGAL_NAME</p>
            </b-media>
          </b-card>
        </b-col>

        <b-col cols="8"> 
          <b-row>
            <b-col>
              <b-card title="Technical Metadata From System" tag="article" class="mb-2">
                <p class="card-text">
                  <b-form>
                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="System Name" label-for="systemName">
                      <b-form-select id="systemName" class="col-8"></b-form-select>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Column Name" label-for="columnName">
                      <b-form-select id="columnName" class="col-8"></b-form-select>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Screen Label Name*" label-for="screenLabelName">
                      <b-form-select id="screenLabelName" class="col-8"></b-form-select>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Screen Label Name*">
                      <p>This field describes the full name of the client and used for capturing lengthy names where the customer name field is unable to accommodat. It is also used for sub-funds. [more]</p>
                    </b-form-group>
                    
                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="CDE (yes/no)">
                      <p>Yes</p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Status*">
                      <p>Active</p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Data Type">
                      <p>VARCHAR2</p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Data Format">
                      <p>Active</p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Data Length">
                      <p>Active</p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Example">
                      <p>Active</p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Derived (Yes/No)*">
                      <p>Active</p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Derivation logic*">
                      <p>Active</p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Sourced from Upstream (Yes/No)*">
                      <p>Active</p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="System Checks*">
                      <p>Active</p>
                    </b-form-group>
                  </b-form>
                </p>
              </b-card>
            </b-col>
          </b-row>

          <b-row>
            <b-col>
              <b-card title="Business Metadata From Domain" tag="article" class="mb-2">
                <p class="card-text">
                  <b-form>
                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Domain">
                      <p>Active</p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Sub Domain">
                      <p>Active</p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Domain Owner">
                      <p>Active</p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Business Term*">
                      <p>Active</p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Business Term Description">
                      <p>Active</p>
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
                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Information Asset Names">
                      <p>Active</p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Information Asset Description">
                      <p>Active</p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="C - Confidentiality">
                      <p>Active</p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="I - Integrity">
                      <p>Active</p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="A - Availability">
                      <p>Active</p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Overall CIA Rating">
                      <p>Active</p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Record Categories">
                      <p>Active</p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="PII Flag">
                      <p>Active</p>
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
                      <p>Active</p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="Immediate Succeeding System*">
                      <p>Active</p>
                    </b-form-group>

                    <b-form-group horizontal :label-cols="4" breakpoint="md" label="DQ Standards | Threshold*">
                      <p>Active</p>
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
export default {
  data () {
    return {
      showModal: this.$route.meta.showModal,
    }
  },
  watch: {
    '$route.meta' ({showModal}) {
      this.showModal = showModal;
    }
  },
  mounted() {
    this.$refs.modalDetails.show();
  },
  methods: {
    handleClose () {
      this.$router.go(-1)
    }
  },
}
</script>