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

                <b-dropdown-item v-if="$route.params.system">
                    <router-link to="/dsc" class="standard-a">System</router-link>
                </b-dropdown-item>
            </b-dropdown>

            <v-toolbar-title to="/" class="app-title"><router-link to="/" class="toolbar-title">Data Catalogue</router-link></v-toolbar-title>
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
        }
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
