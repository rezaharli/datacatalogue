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
          <v-tabs id="page-tab" class="page-tab" v-model="activeTab">
              <v-tab v-for="tab in tabs" class="px-2 mx-5" :id="'tab-' + tab.id" :key="tab.key" :to="addressPath + '/' + tab.key + '/' + urlParam1" :ref="tab.id">{{ tab.name }}</v-tab>
          </v-tabs>

          <transition name="fade" mode="out-in">
            <router-view />
          </transition>
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
    return {
      activeTab: '',
      tabs: [
          { id: 'information', key: 'information', name: 'Information Asset Category' },
          { id: 'personal', key: 'personal', name: 'Personal Data' },
      ]
    };
  },
  computed: {
    addressPath() {
      var tmp = this.$route.path.split("/");
      return tmp.slice(0, 3).join("/");
    },
    urlParam1() {
      return this.$route.params.system;
    }
  },
};
</script>
