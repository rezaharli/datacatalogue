<style>
  @import '../../assets/styles/dashboard.css';
</style>

<template>
    <b-row>
        <b-col>
            <b-row>
                <b-col>
                    <v-tabs left class="page-tab">
                        <v-tab to="/ddo/my">My Domains</v-tab>
                        <v-tab to="/ddo/all">All Domains</v-tab>
                    </v-tabs>
                </b-col>
            </b-row>

            <transition name="fade" mode="out-in">
                <router-view></router-view>
            </transition>
            <!-- <b-row>
                <v-flex xs12>
                    <transition name="fade" mode="out-in">
                        <router-view></router-view>
                    </transition>
                </v-flex> -->

                <!-- <b-row>
                    <b-col>
                        <v-container fluid fill-height class="container-fluid-needpadding">
                            <transition name="fade" mode="out-in">
                                <router-view></router-view>
                            </transition>
                        </v-container>
                    </b-col>
                </b-row> -->
            <!-- </b-row> -->
        </b-col>
    </b-row>
</template>

<script>
export default {
    data () {
      return {}
    },
}
</script>
