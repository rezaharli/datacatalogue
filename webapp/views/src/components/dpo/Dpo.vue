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
                                <v-tab to="/dpo/my">My Processes</v-tab>
                                <v-tab to="/dpo/all">All Processes</v-tab>
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