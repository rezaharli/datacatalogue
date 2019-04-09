<style>
  @import '../../assets/styles/dashboard.css';
</style>

<template>
    <b-row>
        <b-col>
            <!-- <b-row>
                <b-col>
                    <v-tabs left class="page-tab">
                        <v-tab to="/dsc/my">My Systems</v-tab>
                        <v-tab to="/dsc/all">All Systems</v-tab>
                        <v-tab to="/dsc/interfaces">Interfaces</v-tab>
                    </v-tabs>
                </b-col>
            </b-row> -->

            <b-row>
                <b-col>
                    
                </b-col>
            </b-row>

            <b-row>
                <b-col cols=3 />

                <b-col cols=3>
                    <dsc-my />
                </b-col>

                <b-col cols=3>
                    <dsc-all />
                </b-col>

                <b-col cols=3 />
            </b-row>

            <!-- <transition name="fade" mode="out-in">
                <router-view></router-view>
            </transition> -->
        </b-col>
    </b-row>
</template>

<script>
import dscMy from './Dsc-my.vue'
import dscAll from './Dsc-all.vue'

export default {
    components: {
        dscMy, dscAll
    },
    data () {
        return {}
    },
}
</script>
