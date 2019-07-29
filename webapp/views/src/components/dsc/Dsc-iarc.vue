<template>
  <v-content class="mx-4 my-5">
    <b-container fluid>
      <PageHeader />
      
      <b-row>
          <b-col sm=12 md=3>
              <div class="card card-v2 transition">
                  <h6 class="title-1">System Name</h6>
                  <h3 class="title-2 text-capitalize">{{$route.params.system}}</h3>
              </div>
          </b-col>
      </b-row>

      <b-row>
        <b-col>
          <b-card tag="article" class="mb-2">
            <h5 class="card-title border-0 mb-0 pb-0" v-b-toggle.collapse-1>
              Information Asset Category
              <i class="when-opened float-right fa fa-chevron-up"></i>
              <i class="when-closed float-right fa fa-chevron-down"></i>
            </h5>

            <b-collapse id="collapse-1" class="mt-3 pt-4 border-top" visible>
              <p class="card-text">
                <dsc-iarc-information />
              </p>
            </b-collapse>
          </b-card>
        </b-col>
      </b-row>

      <b-row>
        <b-col>
          <b-card tag="article" class="mb-2">
            <h5 class="card-title border-0 mb-0 pb-0" v-b-toggle.collapse-2>
              Personal Data
              <i class="when-opened float-right fa fa-chevron-up"></i>
              <i class="when-closed float-right fa fa-chevron-down"></i>
            </h5>

            <b-collapse id="collapse-2" class="mt-3 pt-4 border-top" visible>
              <p class="card-text">
                <dsc-iarc-personal />
              </p>
            </b-collapse>
          </b-card>
        </b-col>
      </b-row>
    </b-container>
  </v-content>
</template>

<script>
import Vue from "vue";
import { mapState, mapActions } from "vuex";

import PageHeader from '../PageHeader';

import DscIarcInformation from './Dsc-iarc-information.vue';
import DscIarcPersonal from './Dsc-iarc-personal.vue';

export default {
  name: "DscIarc",
  components: {
    PageHeader, DscIarcInformation, DscIarcPersonal
  },
  data() {
    return {};
  },
};
</script>
