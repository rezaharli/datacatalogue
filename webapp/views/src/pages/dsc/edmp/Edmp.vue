<style>
  @import '../../../assets/styles/dashboard.css';
</style>

<template>
    <v-content class="bg-2 card-v3-wrapper">
        <b-container>
            <PageHeader />

            <div class="">
                
                <b-row>
                    <b-col sm=12 md=2 lg=2 class="my-3">
                    </b-col>

                    <b-col sm=12 md=4 lg=4 class="my-3">
                        <div class="card card-v3 transition link" @click="goToEdmpMenuContent('iarc')">
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

                    <b-col sm=12 md=4 lg=4 class="my-3">
                        <div class="card card-v3 transition link" @click="goToEdmpMenuContent('dd')">
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

                    <b-col sm=12 md=2 lg=2 class="my-3">
                    </b-col>

                    <div class="clearfix"></div>
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
                cde: require('@/assets/images/icon-cde.png'),
                cdsp: require('@/assets/images/icon-cdsp.png'),
                iarc: require('@/assets/images/icon-all-system.png'),
                dd: require('@/assets/images/icon-dd.png'),
                imi: require('@/assets/images/icon-imi.png'),
            }
        }
    },
    computed: {
        addressPathParent (){
            var tmp = this.$route.path.split("/")
            return tmp.slice(0, 3).join("/")
        },
        store () {
            return this.$store.state.edmp.all
        },
    },
    mounted () {
        this.getCounts("ENTERPRISE DATA MGMT PLATFORM");
    },
    methods: {
        getCounts (param) {
            this.$store.dispatch(`edmp/getCounts`, param)
        },
        goToDscMenuContent(param){

            var tmp = this.$route.path.split("/")
            tmp.splice(2, 0, encodeURIComponent(param))
            
            this.$router.push(tmp.slice(0, 4).join("/"));
        },
        goToEdmpMenuContent(param){
            this.$router.push(this.addressPathParent + '/' + encodeURIComponent(param));
        },
    }
}
</script>
