<style>
  @import '../../assets/styles/dashboard.css';
</style>

<template>
    <v-content>
        <b-container fluid>
            <PageHeader />

            <b-row>
                <b-col>
                    <b-row style="margin-top: 10px; margin-bottom: 20px;">
                        <b-col>
                            <v-btn class="float-right" color="red-neon" @click.native="resetFilter" dark>
                                <!-- <v-icon dark>filter_list</v-icon> -->
                                <i class="fa fa-filter"></i>
                            </v-btn>
                            
                            <v-btn class="float-right" color="green-tosca" dark>
                                <!-- <v-icon dark>file_copy</v-icon> -->
                                <i class="fa fa-file-excel-o"></i>
                            </v-btn>
                        </b-col>
                    </b-row>

                    <b-row>
                        <b-col cols=3 />

                        <b-col cols=3>
                            <dsc-my />
                        </b-col>

                        <b-col cols=3>
                            <dsc-all />
                        </b-col>

                        <b-col cols=3 />
                    </b-row>

                    <!-- <transition name="fade" mode="out-in">
                        <router-view></router-view>
                    </transition> -->
                </b-col>
            </b-row>
        </b-container>
    </v-content>
</template>

<script>
import { mapState, mapActions } from 'vuex'

import PageHeader from '../PageHeader';

import dscMy from './Dsc-my.vue';
import dscAll from './Dsc-all.vue';

export default {
    components: { PageHeader, dscMy, dscAll },
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
            this.myStore.filters.left = {};
            this.getMyLeftTable();

            // this.myStore.filters.right = {}
            // this.getMyRightTable(this.$route.params.system);

            this.allStore.filters.left = {}
            this.getAllLeftTable();

            // this.allStore.filters.right = {}
            // this.getAllRightTable(this.$route.params.system);
        }
    },
}
</script>