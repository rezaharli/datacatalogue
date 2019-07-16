<style>
  @import '../../assets/styles/dashboard.css';
</style>

<template>
    <v-content class="bg-2 card-v3-wrapper px-1 py-5">
        <PageHeader />
        <!-- <b-container fluid class="mt-5">
            <b-row>
                <b-col sm=12 md=3>
                    <div class="card card-v2 transition">
                        <h6 class="title-1">Principle Risk Type</h6>
                        <h3 class="title-2 text-capitalize">{{$route.params.system}}</h3>
                    </div>
                </b-col>
                <b-col sm=12 md=3>
                    <div class="card card-v2 transition">
                        <h6 class="title-1">Sub-Risk Type</h6>
                        <h3 class="title-2 text-capitalize">{{$route.params.system}}</h3>
                    </div>
                </b-col>
            </b-row>
        </b-container> -->
        <b-container>
            <div class="">
                <b-row class="mb-3">
                    <b-col sm=12 md=1 lg=1 />
                    <b-col sm=12 md=5 lg=5>
                        <div class="card card-v2 transition">
                            <h6 class="title-1">Principle Risk Type</h6>
                            <h3 class="title-2 text-capitalize">{{ store.counts ? store.counts.PRINCIPAL_RISK : "" }}</h3>
                        </div>
                    </b-col>
                    <b-col sm=12 md=5 lg=5>
                        <div class="card card-v2 transition">
                            <h6 class="title-1">Sub-Risk Type</h6>
                            <h3 class="title-2 text-capitalize">{{ store.counts ? store.counts.RISK_SUB : "" }}</h3>
                        </div>
                    </b-col>
                    <b-col sm=12 md=1 lg=1 />
                </b-row>
                <b-row>
                    <b-col sm=12 md=1 lg=1 />
                    <b-col sm=12 md=5 lg=5 class="my-3">
                        <div class="card card-v3 transition link" @click="goToRfoMenuContent('summary')">
                            <b-row>
                                <b-col cols=5 class="transition">
                                    <v-img :src="images.summary" :contain="true" class="card-icon"></v-img>
                                </b-col>

                                <b-col cols=7 class="">
                                    <div class="card-content">
                                        <div class="text">
                                            <h6 class="title-1">Summary View</h6>
                                        </div>
                                    </div>
                                </b-col>
                            </b-row>

                            <!-- <h4 class="title-2">{{ store.counts ? store.counts.CDE_COUNT : 0 }}</h4> -->
                            <div class="card-circle-left transition"></div>
                        </div>
                    </b-col>
                
                    <b-col sm=12 md=5 lg=5 class="my-3">
                        <div class="card card-v3 transition link" @click="goToRfoMenuContent('hierarchy')">
                            <b-row>
                                <b-col cols=5 class="transition">
                                    <v-img :src="images.hierarchy" :contain="true" class="card-icon"></v-img>
                                </b-col>

                                <b-col cols=7 class="">
                                    <div class="card-content">
                                        <div class="text">
                                            <h6 class="title-1">Hierarchy View</h6>
                                        </div>
                                    </div>
                                </b-col>
                            </b-row>

                            <!-- <h4 class="title-2">{{ store.counts ? store.counts.DSP_COUNT : 0 }}</h4> -->
                            <div class="card-circle-left transition"></div>
                        </div>
                    </b-col>

                    <b-col sm=12 md=1 lg=1 />

                    <div class="clearfix"></div>
                </b-row>
            </div>

        </b-container>
    </v-content>
</template>

<script>
import PageHeader from '../PageHeader';
import { mapState, mapActions } from 'vuex'

export default {
    name: "RfoMenu",
    components: { PageHeader},
    data () {
        return {
            images: {
                summary: require('../../assets/images/icon-all-system.png'),
                hierarchy: require('../../assets/images/icon-all-system.png'),
            }
        }
    },
    computed: {
        addressPathParent (){
            var tmp = this.$route.path.split("/")
            return tmp.slice(0, 2).join("/")
        },
        store () {
            return this.$store.state.rfo.all
            // return this.$store.state[this.storeName].all;
        },
    },
    mounted () {
        this.getCounts(this.$route.params.type);
    },
    methods: {
        getCounts (param) {
            this.$store.dispatch(`rfo/getCounts`, param)
        },
        goToRfoMenuContent(param){
            this.$router.push(this.addressPathParent + '/' + encodeURIComponent(param) + '/' + encodeURIComponent(this.$route.params.type));
        },
    }
}
</script>
