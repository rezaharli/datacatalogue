<style>
  @import '../../assets/styles/dashboard.css';
</style>

<template>
    <v-content>
        <CToolbar></CToolbar>

        <b-container fluid>
            <b-row v-if="showTabs">
                <b-col>
                    <v-tabs left class="page-tab">
                        <v-tab to="/dsc/my">My Systems</v-tab>
                        <v-tab to="/dsc/all">All System</v-tab>
                        <v-tab to="/dsc/interfaces">Interfaces</v-tab>
                    </v-tabs>
                </b-col>
            </b-row>

            <b-row>
                <b-col>
                    <b-container fluid class="container-fluid-needpadding">
                        <transition name="fade" mode="out-in">
                            <router-view></router-view>
                        </transition>
                    </b-container>
                </b-col>
            </b-row>
        </b-container>
    </v-content>
</template>

<script>
import CToolbar from '../ComponentToolbar';

export default {
    components: { CToolbar },
    data () {
      return {
          showTabs: this.isShowTabs(),
      }
    },
    watch: {
        $route (){
            this.showTabs = this.isShowTabs()
        }
    },
    methods: {
        isShowTabs (){
            return (this.$route.name == "dsc.my.details") ? false : true
        }
    }
}
</script>
