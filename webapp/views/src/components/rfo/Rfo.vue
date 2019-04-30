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

        <b-container fluid>
            <b-row>
                <!-- <b-col sm=12 md=1 lg=1 /> -->

                <b-col sm=12 md=6 lg=6>
                    <rfo-my />
                </b-col>

                <b-col sm=12 md=6 lg=6>
                    <rfo-all />
                </b-col>

                <!-- <b-col sm=12 md=1 lg=1 /> -->
            </b-row>
        </b-container>
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