<style>
  @import '../assets/styles/dashboard.css';
</style>

<template>
    <v-app>
        <CToolbar></CToolbar>
        <v-content>
            <v-container fluid>
                <v-tabs fixed-tabs>
                    <v-tab
                    v-for="n in 3"
                    :key="n"
                    >
                    Item {{ n }}
                    </v-tab>
                </v-tabs>
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
