<style>
.button-footer-logout{
    position: absolute;
    right: 10px;

    color: #666666 !important;
    background-color: #ffffff !important;
    border-color: #ffffff !important;
    font-weight: bold !important;
}
</style>


<template>
  <v-footer class="pa-3 white--text" color="teal">
    <div style="font-weight: bold">Eaciit &copy; {{ new Date().getFullYear() }}</div>
   
    <!-- <b-link v-if="showLogout" @click="doLogout">Logout</b-link> -->

    <!-- <v-btn small right v-if="showLogout" @click="doLogout">Logout</v-btn> -->

    <b-button size="sm" variant="secondary" class="button-footer-logout" v-if="showLogout" @click="doLogout">Logout</b-button>
 </v-footer>
</template>

<script>
import { mapState, mapActions } from 'vuex'

export default {
    data() {
      return {}
    },
    computed: {
      ...mapState('account', ['status']),
      showLogout () {
        var user = localStorage.getItem('user');
        var show = this.$route.name == "landingpage";
        
        return user && show ? true : false;
      }
    },
    methods: {
      ...mapActions('account', ['logout']),
      doLogout() {
        this.logout();
      }
    }
}
</script>