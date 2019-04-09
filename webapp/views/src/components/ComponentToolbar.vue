<style>
  @import '../assets/styles/dashboard.css';
</style>

<template>
    <div>
        <v-navigation-drawer app
                :clipped="drawer.clipped"
                :fixed="drawer.fixed"
                :permanent="drawer.permanent"
                :mini-variant="drawer.mini"
                v-model="drawer.open">
            <v-list>
                <v-list-tile to="/">
                    <v-list-tile-action><v-icon>dashboard</v-icon></v-list-tile-action>
                    <v-list-tile-content><v-list-tile-title>Home</v-list-tile-title></v-list-tile-content>
                </v-list-tile>

                <v-list-tile to="/dsc" v-if="isDSC">
                    <v-list-tile-action><v-icon>list</v-icon></v-list-tile-action>
                    <v-list-tile-content><v-list-tile-title>Dataset Custodian</v-list-tile-title></v-list-tile-content>
                </v-list-tile>

                <v-list-tile to="/dpo" v-if="isDPO">
                    <v-list-tile-action><v-icon>list</v-icon></v-list-tile-action>
                    <v-list-tile-content><v-list-tile-title>Downstream Process Owner</v-list-tile-title></v-list-tile-content>
                </v-list-tile>

                <v-list-tile to="/ddo" v-if="isDDO">
                    <v-list-tile-action><v-icon>list</v-icon></v-list-tile-action>
                    <v-list-tile-content><v-list-tile-title>Data Domain Owner</v-list-tile-title></v-list-tile-content>
                </v-list-tile>

                <v-list-tile to="/rfo" v-if="isRFO">
                    <v-list-tile-action><v-icon>list</v-icon></v-list-tile-action>
                    <v-list-tile-content><v-list-tile-title>Risk Framework Owner</v-list-tile-title></v-list-tile-content>
                </v-list-tile>

                <v-list-tile to="/access" v-if="isAdmin">
                    <v-list-tile-action><v-icon>list</v-icon></v-list-tile-action>
                    <v-list-tile-content><v-list-tile-title>User Access</v-list-tile-title></v-list-tile-content>
                </v-list-tile>
            </v-list>
        </v-navigation-drawer>

        <v-toolbar app
                :fixed="toolbar.fixed"
                :clipped-left="toolbar.clippedLeft">
            <v-toolbar-side-icon @click.stop="toggleDrawer"></v-toolbar-side-icon>
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
