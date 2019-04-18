<style>
  @import '../../assets/styles/dashboard.css';
</style>

<template>
    <v-content>
        <b-container fluid>
            <CToolbar></CToolbar>
            
            <b-row>
                <b-col>
                    <b-row>
                        <b-col>
                            <v-tabs left class="page-tab">
                              <v-tab key="technical-metadata">Technical Metadata From System</v-tab>
                              <v-tab key="business-metadata">Business Metadata From Domain</v-tab>
                              <v-tab key="policy-related">Policy Related Information</v-tab>
                              <v-tab key="interfaces">Interfaces</v-tab>

                              <v-tab-item key="technical-metadata">
                                <dsc-dd-technical />
                              </v-tab-item>

                              <v-tab-item key="business-metadata">
                                <dsc-dd-business />
                              </v-tab-item>

                              <v-tab-item key="policy-related">
                                <dsc-dd-policy />
                              </v-tab-item>

                              <v-tab-item key="interfaces">
                                <dsc-dd-interfaces />
                              </v-tab-item>
                            </v-tabs>
                        </b-col>
                    </b-row>
                </b-col>
            </b-row>
        </b-container>
    </v-content>
</template>

<script>
import CToolbar from '../ComponentToolbar';
import DscDdTechnical from './Dsc-dd-technical';
import DscDdBusiness from './Dsc-dd-business';
import DscDdPolicy from './Dsc-dd-policy';
import DscDdInterfaces from './Dsc-dd-interfaces';

export default {
    components: { CToolbar, DscDdTechnical, DscDdBusiness, DscDdPolicy, DscDdInterfaces },
    data () {
      return {}
    },
}
</script>