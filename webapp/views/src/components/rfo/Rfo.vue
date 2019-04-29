<style>
  @import '../../assets/styles/dashboard.css';
</style>

<template>
    <v-content>
        <PageHeader />

        <b-container fluid>
            <b-row style="margin-top: 10px; margin-bottom: 20px;">
                <b-col>
                    <b-button class="float-right red-neon icon-only ml-3" @click="resetFilter">
                        <i class="fa fa-filter"></i>
                    </b-button>

                    <page-export class="float-right" storeName="rfoall" :leftTableCols="myStore.leftHeaders" :rightTableCols="myStore.rightHeaders"/>
                </b-col>
            </b-row>
        </b-container>

        <b-container>
            <b-row>
                <b-col sm=12 md=1 lg=1 />

                <b-col sm=12 md=5 lg=5>
                    <rfo-my />
                </b-col>

                <b-col sm=12 md=5 lg=5>
                    <rfo-all />
                </b-col>

                <b-col sm=12 md=1 lg=1 />
            </b-row>

            <!-- <transition name="fade" mode="out-in">
                <router-view></router-view>
            </transition> -->
        </b-container>

        <v-navigation-drawer v-model="store.drawer" right absolute temporary width="350">

            <b-row align-h="end" class="pr-4 pt-3 bg-light-grey">
                <b-button variant="light" class="float-right" @click.stop="store.drawer = !store.drawer">
                    <i class="fa fa-fw fa-arrow-right"></i>
                </b-button>
            </b-row>

            <v-list class="p-0 bg-light-grey">
                <!-- <v-list-tile avatar>
                    <v-list-tile-content>
                        <v-list-tile-title>{{ store.drawerContent.systemName }}</v-list-tile-title>
                    </v-list-tile-content>
                    <v-list-tile-content>
                        <v-list-tile-title>{{ store.drawerContent.itamID }}</v-list-tile-title>
                    </v-list-tile-content>
                    <v-list-tile-content>
                        <v-list-tile-title>{{ store.drawerContent.owners.length }}</v-list-tile-title>
                    </v-list-tile-content>
                </v-list-tile> -->
                <h1 class="px-4 py-2">{{ store.drawerContent.systemName }}</h1>
                <b-row class="px-4 pt-2 pb-4">
                    <b-col cols=12 sm=auto class="border-right border-dark">ITAM ID: {{ store.drawerContent.itamID }}</b-col>
                    <b-col cols=12 sm=auto>{{ store.drawerContent.owners.length }} Owners</b-col>
                </b-row>

            </v-list>
            

            <v-list class="px-0 pt-4">
                <!-- <v-divider></v-divider> -->

                <b-row align-v="center" class="px-4 py-2"
                    v-for="(owner, i) in store.drawerContent.owners"
                    :key="i">
                    <b-col cols="3" class="text-right">{{ owner.BANK_ID }}</b-col>

                    <b-col cols="9" class="text-capitalize border-left border-dark">{{ owner.DATASET_CUSTODIAN }}</b-col>
                </b-row>
            </v-list>
        </v-navigation-drawer>
    </v-content>
</template>

<script>
import { mapState, mapActions } from 'vuex'

import PageHeader from '../PageHeader';
import pageExport from '../PageExport.vue'

import rfoMy from './Rfo-my.vue';
import rfoAll from './Rfo-all.vue';

export default {
    components: { PageHeader, pageExport, rfoMy, rfoAll },
    data () {
        return {}
    },
    computed: {
        store () { return this.$store.state.rfo.all },
        myStore () { return this.$store.state.rfomy.all },
        allStore () { return this.$store.state.rfoall.all }
    },
    mounted() {
        this.resetFilter();
    },
    methods: {
        getMyLeftTable () { return this.$store.dispatch(`rfomy/getLeftTable`) },
        getMyRightTable () { return this.$store.dispatch(`rfomy/getRightTable`) },
        getAllLeftTable () { return this.$store.dispatch(`rfoall/getLeftTable`) },
        getAllRightTable () { return this.$store.dispatch(`rfoall/getRightTable`) },
        resetFilter (e) {
            if(Object.keys(this.myStore.filters.left).length > 0){
                this.myStore.filters.left = {};
                this.getMyLeftTable();
            }

            // if(Object.keys(this.myStore.filters.right).length > 0){
            //     this.myStore.filters.right = {}
            //     this.getMyRightTable(this.$route.params.system);
            // }

            if(Object.keys(this.allStore.filters.left).length > 0){
                this.allStore.filters.left = {}
                this.getAllLeftTable();
            }

            // if(Object.keys(this.allStore.filters.right).length > 0){
            //     this.allStore.filters.right = {}
            //     this.getAllRightTable(this.$route.params.system);
            // }
        }
    },
}
</script>