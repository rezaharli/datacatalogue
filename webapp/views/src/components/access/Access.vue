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