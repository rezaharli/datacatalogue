<style>
  @import '../assets/styles/dashboard.css';
</style>

<template>
    <div>
        <v-toolbar app
                :fixed="toolbar.fixed"
                :clipped-left="toolbar.clippedLeft">

            <b-dropdown no-caret id="ddowntoolbar" v-if="!($route.name == 'dsc')">
                <template slot="button-content">
                    <v-btn icon>
                        <v-icon>home</v-icon>
                    </v-btn>
                </template>
            </b-dropdown>

            <v-toolbar-title to="/" class="app-title text-capitalize">
                <router-link to="/" class="toolbar-title">
                    {{ $route.name.indexOf('dsc.') != -1 ? ($route.name.indexOf('dsc.edmp') != -1 ? "ENTERPRISE DATA MGMT PLATFORM" : $route.params.system) : (title ? title : "Data Catalogue") }}
                </router-link>
            </v-toolbar-title>
            
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
    props: ["title"],
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
