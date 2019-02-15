<style>
  @import './assets/styles/style.css';
  @import './assets/vendors/font-awesome-4.7.0/css/font-awesome.min.css';
</style>

<template>
  <v-app>
    <transition name="fade" mode="out-in">
      <router-view />
    </transition>

    <Footer />
  </v-app>
</template>

<script>
import Footer from '@/components/Footer.vue'
import { mapState, mapActions } from 'vuex'


export default {
  name: 'App',
  components: { Footer },
  data() {
    return {}
  },
  methods: {
      ...mapActions({
          clearAlert: 'alert/clear' 
      })
  },
  watch: {
      $route (to, from){
          // clear alert on location change
          this.clearAlert();
      }
  } 
}
</script>