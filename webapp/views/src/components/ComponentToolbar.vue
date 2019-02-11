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

                <v-list-tile to="/dsc">
                    <v-list-tile-action><v-icon>list</v-icon></v-list-tile-action>
                    <v-list-tile-content><v-list-tile-title>Dataset Custodian</v-list-tile-title></v-list-tile-content>
                </v-list-tile>

                <v-list-tile to="/dpo">
                    <v-list-tile-action><v-icon>list</v-icon></v-list-tile-action>
                    <v-list-tile-content><v-list-tile-title>Downstream Process Owner</v-list-tile-title></v-list-tile-content>
                </v-list-tile>

                <v-list-tile to="/ddo">
                    <v-list-tile-action><v-icon>list</v-icon></v-list-tile-action>
                    <v-list-tile-content><v-list-tile-title>Data Domain Owner</v-list-tile-title></v-list-tile-content>
                </v-list-tile>

                <v-list-tile to="/rfo">
                    <v-list-tile-action><v-icon>list</v-icon></v-list-tile-action>
                    <v-list-tile-content><v-list-tile-title>Risk Framework Owner</v-list-tile-title></v-list-tile-content>
                </v-list-tile>
            </v-list>
        </v-navigation-drawer>

        <v-toolbar app
                :fixed="toolbar.fixed"
                :clipped-left="toolbar.clippedLeft">
            <v-toolbar-side-icon @click.stop="toggleDrawer"></v-toolbar-side-icon>
            <v-toolbar-title to="/" class="app-title"><router-link to="/" class="toolbar-title">Data Catalogue</router-link></v-toolbar-title>
            <v-spacer></v-spacer>

            <v-toolbar-title class="login-name">Hi, Samuel Rajakumar</v-toolbar-title>
            <v-btn icon>
                <v-icon>more_vert</v-icon>
            </v-btn>
        </v-toolbar>
    </div>
</template>

<script>
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
    methods: {
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
