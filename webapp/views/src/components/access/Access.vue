<style>
  @import '../../assets/styles/dashboard.css';
</style>

<template>
    <v-content>
        <PageHeader />

        <b-container fluid>
            <b-row>
                <b-col sm=12 md=12 lg=12>
                    <v-tabs left class="page-tab">
                        <v-tab to="/access/users">Users</v-tab>
                        <v-tab to="/access/roles">Roles</v-tab>
                        <!-- <v-tab to="/access/roles">Application Usage Detail</v-tab>
                        <v-tab to="/access/roles">User Audit Trail</v-tab> -->
                    </v-tabs>
                </b-col>
            </b-row>
            
            <transition name="fade" mode="out-in">
                <router-view></router-view>
            </transition>


            <!-- <transition name="fade" mode="out-in">
                <router-view></router-view>
            </transition> -->
        </b-container>
    </v-content>
</template>

<script>
import CToolbar from '../ComponentToolbar';

import PageHeader from '../PageHeader';

export default {
    components: { CToolbar, PageHeader },
    data () {
      return {}
    },
}
</script>