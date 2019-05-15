<style>
  @import '../../assets/styles/dashboard.css';
</style>

<template>
    <v-content>
        <PageHeader />

        <b-container fluid>
            <b-row>
                <b-col>
                    <v-tabs left class="page-tab">
                        <v-tab class="px-2 mx-3" to="/access/users">Users</v-tab>
                        <v-tab class="px-2 mx-3" to="/access/roles">Roles</v-tab>
                        <!-- <v-tab class="px-2 mx-3" to="/access/roles">Application Usage Detail</v-tab> -->
                        <v-tab class="px-2 mx-3" to="/access/usage">User Audit Trail</v-tab>
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