<style>
  @import '../../assets/styles/dashboard.css';
</style>

<template>
    <div>
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
