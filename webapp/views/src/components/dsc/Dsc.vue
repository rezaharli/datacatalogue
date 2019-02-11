<style>
  @import '../../assets/styles/dashboard.css';
</style>

<template>
    <div>
        <b-row>
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
    </div>
</template>

<script>
export default {
    data () {
      return {}
    },
}
</script>
