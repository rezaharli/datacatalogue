import Vue from 'vue'
import './plugins/vuetify'
import App from './App.vue'

import router from './routes'
import { store } from './_store'

import VueLodash from 'vue-lodash'

import BootstrapVue from 'bootstrap-vue'
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'

require("expose-loader?$!jquery");
require("expose-loader?math!mathjs");

Vue.config.productionTip = process.env.NODE_ENV == 'production';
Vue.use(BootstrapVue);
Vue.use(VueLodash);

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')