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

import IdleVue from 'idle-vue'
 
const eventsHub = new Vue()
 
Vue.use(IdleVue, {
  eventEmitter: eventsHub,
  idleTime: 15 * 60 * 1000 //15 minutes
})

new Vue({
  router,
  store,
  render: h => h(App),
  onIdle() {
    localStorage.removeItem('user');
    location.reload(true);
  },
  watch: {
    isAppIdle() {
    }
  }
}).$mount('#app')