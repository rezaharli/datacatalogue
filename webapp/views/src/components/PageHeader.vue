<style>
  @import '../assets/styles/dashboard.css';
</style>

<template>
    <div>
        <v-toolbar app
                :fixed="toolbar.fixed"
                :clipped-left="toolbar.clippedLeft">

            <b-dropdown no-caret id="ddowntoolbar">
                <template slot="button-content">
                    <v-btn icon>
                        <v-icon>home</v-icon>
                    </v-btn>
                </template>

                <b-dropdown-item>
                    <router-link to="/" class="standard-a">Home</router-link>
                </b-dropdown-item>

                <b-dropdown-item v-if="$route.name == 'dsc.menu' || $route.name.indexOf('dsc.') != -1">
                    <router-link to="/dsc" class="standard-a">DSC System</router-link>
                </b-dropdown-item>

                <b-dropdown-item v-if="
                    $route.name == 'dsc.cde' || 
                    $route.name == 'dsc.cdp' || 
                    $route.name == 'dsc.interfaces' || 
                    $route.name == 'dsc.cdp.cde' || 
                    $route.name == 'dsc.interfaces.cde'
                    ">
                    <router-link :to="goToDscMenu" class="standard-a">System Landing Page</router-link>
                </b-dropdown-item>

                <b-dropdown-item v-if="$route.name == 'dsc.cdp.cde'">
                    <router-link :to="goToLevel3" class="standard-a">Critical Downstream Process View</router-link>
                </b-dropdown-item>

                <b-dropdown-item v-if="$route.name == 'dsc.interfaces.cde'">
                    <router-link :to="goToLevel3" class="standard-a">Immediate Interface View</router-link>
                </b-dropdown-item>
            </b-dropdown>

            <v-toolbar-title to="/" class="app-title text-capitalize"><router-link to="/" class="toolbar-title">{{ $route.name.indexOf('dsc.') != -1 ? $route.params.system : "Data Catalogue" }}</router-link></v-toolbar-title>
            <v-spacer></v-spacer>

            <v-toolbar-title class="login-name">Hi, {{ account.user.Name }}</v-toolbar-title>

            <b-dropdown right no-caret id="ddowntoolbar">
                <template slot="button-content">
                    <v-btn icon>
                        <v-icon>more_vert</v-icon>
                    </v-btn>
                </template>

                <b-dropdown-item @click="doLogout">Logout</b-dropdown-item>
            </b-dropdown>
        </v-toolbar>
    </div>
</template>

<script>
import { mapState, mapActions } from 'vuex'

export default {
    data () {
        return {
            drawer: {
                clipped: true,
                fixed: false,
                mini: false,
                open: false,
                permanent: false,
            }, 
            toolbar: {
                fixed: true,
                clippedLeft: true
            },
        }
    },
    computed: {
        ...mapState({
            account: state => state.account,
        }),
        ...mapState('account', ['user']),
        isAdmin () {
            return this.user.Role.split(",").indexOf("Admin") != -1
        },
        isDSC () {
            return this.user.Role.split(",").indexOf("DSC") != -1
        },
        isDPO () {
            return this.user.Role.split(",").indexOf("DPO") != -1
        },
        isDDO () {
            return this.user.Role.split(",").indexOf("DDO") != -1
        },
        isRFO () {
            return this.user.Role.split(",").indexOf("RFO") != -1
        },
        goToDscMenu(){
            var tmp = this.$route.path.split("/");
            var tmp2 = tmp.slice(0, 2).join("/");
            var tmp3 = tmp2 + '/' + this.$route.params.system;
            return tmp3;
        },
        goToLevel3(){
            var tmp = this.$route.path.split("/");
            var tmp2 = tmp.slice(0, 3).join("/");
            var tmp3 = tmp2 + '/' + this.$route.params.system;
            return tmp3;
        },
    },
    methods: {
        ...mapActions('account', ['logout']),
        doLogout() {
            this.logout();
        },
        toggleDrawer () {
            if (this.drawer.permanent) {
                this.drawer.permanent = !this.drawer.permanent
                this.drawer.clipped = true
                this.toolbar.clippedLeft = true
            } else {
                this.drawer.open = !this.drawer.open
            }
        }
    }
}
</script>
