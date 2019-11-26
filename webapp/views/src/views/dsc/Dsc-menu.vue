<style>
  @import '../../assets/styles/dashboard.css';
</style>

<template>
    <v-content class="bg-2 card-v3-wrapper">
        <b-container>
            <PageHeader />

            <!-- <b-row>
                <b-col sm=12 md=6>
                    <div class="card card-v2 transition">
                        <h6 class="title-1">System Name</h6>
                        <h3 class="title-2 text-capitalize">{{$route.params.system}}</h3>
                    </div>
                </b-col>
            </b-row> -->
            
            <!-- <b-row>
                <b-col md=12 class="mb-3 mt-5">
                    <h5><b>Select <span class="text-capitalize">{{$route.params.system}}</span> Type</b></h5>
                </b-col>
            </b-row> -->
            <div class="">
                
                <b-row>
                    <b-col sm=12 md=4 lg=4 class="my-3">
                        <div class="card card-v3 transition link" @click="goToDscMenuContent('cde')">
                            <b-row>
                                <b-col cols=5 class="transition">
                                    <v-img :src="images.cde" :contain="true" class="card-icon"></v-img>
                                </b-col>

                                <b-col cols=7 class="">
                                    <div class="card-content">
                                        <div class="text">
                                            <h6 class="title-1">Critical Data Elements</h6>
                                        </div>
                                    </div>
                                </b-col>
                            </b-row>

                            <h4 class="title-2">{{ store.counts ? store.counts.CDE_COUNT : 0 }}</h4>
                            <div class="card-circle-left transition"></div>
                        </div>
                    </b-col>
                
                    <b-col sm=12 md=4 lg=4 class="my-3">
                        <div class="card card-v3 transition link" @click="goToDscMenuContent('cdp')">
                            <b-row>
                                <b-col cols=5 class="transition">
                                    <v-img :src="images.cdsp" :contain="true" class="card-icon"></v-img>
                                </b-col>

                                <b-col cols=7 class="">
                                    <div class="card-content">
                                        <div class="text">
                                            <h6 class="title-1">Critical Downstream Processes</h6>
                                        </div>
                                    </div>
                                </b-col>
                            </b-row>

                            <h4 class="title-2">{{ store.counts ? store.counts.DSP_COUNT : 0 }}</h4>
                            <div class="card-circle-left transition"></div>
                        </div>
                    </b-col>

                    <b-col sm=12 md=4 lg=4 class="my-3">
                        <div class="card card-v3 transition link" @click="goToDscMenuContent('iarc')">
                            <b-row>
                                <b-col cols=5 class="transition">
                                    <v-img :src="images.iarc" :contain="true" class="card-icon"></v-img>
                                </b-col>

                                <b-col cols=7 class="">
                                    <div class="card-content">
                                        <div class="text">
                                            <h6 class="title-1">Data Classification (Information Asset, Personal Data &amp; Record Categories)</h6>
                                        </div>
                                    </div>
                                </b-col>
                            </b-row>

                            <h4 class="title-2">{{ store.counts ? store.counts.POLICY_COUNT : 0 }}</h4>
                            <div class="card-circle-left transition"></div>
                        </div>
                    </b-col>

                    <div class="clearfix"></div>
                </b-row><b-row>

                    <b-col sm=12 md=4 lg=4 class="my-3">
                        <div class="card card-v3 transition link" @click="goToDscMenuContent('interfaces')">
                            <b-row>
                                <b-col cols=5 class="transition">
                                    <v-img :src="images.imi" :contain="true" class="card-icon"></v-img>
                                </b-col>

                                <b-col cols=7 class="">
                                    <div class="card-content">
                                        <div class="text">
                                            <h6 class="title-1">Immediate Interfaces</h6>
                                        </div>
                                    </div>
                                </b-col>
                            </b-row>

                            <h4 class="title-2">{{ store.counts ? store.counts.INTERFACE_COUNT : 0 }}</h4>
                            <div class="card-circle-left transition"></div>
                        </div>
                    </b-col>

                    <b-col sm=12 md=4 lg=4 class="my-3">
                        <div class="card card-v3 transition link" @click="goToDscMenuContent('dd')">
                            <b-row>
                                <b-col cols=5 class="transition">
                                    <v-img :src="images.dd" :contain="true" class="card-icon"></v-img>
                                </b-col>

                                <b-col cols=7 class="">
                                    <div class="card-content">
                                        <div class="text">
                                            <h6 class="title-1">Data Dictionary</h6>
                                        </div>
                                    </div>
                                </b-col>
                            </b-row>

                            <div class="card-circle-left transition"></div>
                        </div>
                    </b-col>
                    <b-col sm=12 md=4 lg=4 />
                </b-row>
            </div>

        </b-container>
    </v-content>
</template>

<script>
import PageHeader from '@/components/PageHeader';
import { mapState, mapActions } from 'vuex'

export default {
    components: { PageHeader},
    data () {
        return {
            images: {
                cde: require('../../assets/images/icon-cde.png'),
                cdsp: require('../../assets/images/icon-cdsp.png'),
                iarc: require('../../assets/images/icon-all-system.png'),
                dd: require('../../assets/images/icon-dd.png'),
                imi: require('../../assets/images/icon-imi.png'),
            }
        }
    },
    computed: {
        addressPathParent (){
            var tmp = this.$route.path.split("/")
            return tmp.slice(0, 2).join("/")
        },
        store () {
            return this.$store.state.dsc.all
        },
    },
    mounted () {
        this.getCounts(this.$route.params.system);
    },
    methods: {
        getCounts (param) {
            this.$store.dispatch(`dsc/getCounts`, param)
        },
        goToDscMenuContent(param){
            this.$router.push(this.addressPathParent + '/' + encodeURIComponent(param) + '/' + encodeURIComponent(this.$route.params.system));
        },
    }
}
</script>
