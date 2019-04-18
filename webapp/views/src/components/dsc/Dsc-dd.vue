<style>
  @import '../../assets/styles/dashboard.css';
</style>

<template>
    <v-content>
        <b-container fluid>
            <PageHeader />
            
            <b-row>
                <b-col>
                    <v-tabs left class="page-tab">
                      <v-tab class="px-2 mx-5" key="technical-metadata">Technical Metadata</v-tab>
                      <v-tab class="px-2 mx-5" key="business-metadata">Business Metadata</v-tab>
                      <v-tab class="px-2 mx-5" key="policy-related">Policy Related Information</v-tab>
                      <v-tab class="px-2 mx-5" key="interfaces">Interfaces</v-tab>

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
        </b-container>
    </v-content>
</template>

<script>
import PageHeader from '../PageHeader';
import DscDdTechnical from './Dsc-dd-technical';
import DscDdBusiness from './Dsc-dd-business';
import DscDdPolicy from './Dsc-dd-policy';
import DscDdInterfaces from './Dsc-dd-interfaces';

export default {
    components: { PageHeader, DscDdTechnical, DscDdBusiness, DscDdPolicy, DscDdInterfaces },
    data () {
      return {}
    },
}
</script>