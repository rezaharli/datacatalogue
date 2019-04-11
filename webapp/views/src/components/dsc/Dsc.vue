<style>
  @import '../../assets/styles/dashboard.css';
</style>

<template>
    <v-content>
        <PageHeader />

        <b-container fluid>
            <b-row style="margin-top: 10px; margin-bottom: 20px;">
                <b-col>
                    <v-btn class="float-right" color="red-neon" @click.native="resetFilter" dark>
                        <!-- <v-icon dark>filter_list</v-icon> -->
                        <i class="fa fa-filter"></i>
                    </v-btn>

                    <page-export class="float-right" storeName="dscmy" :leftTableCols="myStore.leftHeaders" :rightTableCols="myStore.rightHeaders"/>
                </b-col>
            </b-row>
        </b-container>

        <b-container>
            <b-row>
                <b-col sm=12 md=1 lg=1 />

                <b-col sm=12 md=5 lg=5>
                    <dsc-my />
                </b-col>

                <b-col sm=12 md=5 lg=5>
                    <dsc-all />
                </b-col>

                <b-col sm=12 md=1 lg=1 />
            </b-row>

            <!-- <transition name="fade" mode="out-in">
                <router-view></router-view>
            </transition> -->
                
        </b-container>
    </v-content>
</template>

<script>
import { mapState, mapActions } from 'vuex'

import PageHeader from '../PageHeader';
import pageExport from '../PageExport.vue'

import dscMy from './Dsc-my.vue';
import dscAll from './Dsc-all.vue';

export default {
    components: { PageHeader, pageExport, dscMy, dscAll },
    data () {
        return {}
    },
    computed: {
        myStore () { return this.$store.state.dscmy.all },
        allStore () { return this.$store.state.dscall.all }
    },
    methods: {
        getMyLeftTable () { return this.$store.dispatch(`dscmy/getLeftTable`) },
        getMyRightTable () { return this.$store.dispatch(`dscmy/getRightTable`) },
        getAllLeftTable () { return this.$store.dispatch(`dscall/getLeftTable`) },
        getAllRightTable () { return this.$store.dispatch(`dscall/getRightTable`) },
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