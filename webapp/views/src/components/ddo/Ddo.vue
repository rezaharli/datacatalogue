<style>
  @import '../../assets/styles/dashboard.css';
</style>

<template>
    <v-content>
        <b-container fluid>
            <CToolbar></CToolbar>
            
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
      return {}
    },
}
</script>