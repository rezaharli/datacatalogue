<style>
  @import '../assets/styles/dashboard.css';
</style>

<template>
    <v-app>
        <CToolbar></CToolbar>
        <v-content>
            <v-container fluid>
            <router-view></router-view>
            </v-container>
        </v-content>
    </v-app>
</template>

<script>
import CToolbar from './ComponentToolbar';

export default {
    components: {
        CToolbar,
    }
}
</script>
