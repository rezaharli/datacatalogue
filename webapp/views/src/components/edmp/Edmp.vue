<style>
  @import '../../assets/styles/dashboard.css';
</style>

<template>
    <v-content class="bg-2 card-v3-wrapper">
        <PageHeader />
        <b-container>

            <div class="text-center">
                
                <!-- <b-row>
                    <b-col sm=12 md=1 lg=1 />
                    <b-col sm=12 md=5 lg=5 class="my-3">
                        <div class="card card-v3 transition link" @click="goToMenuContent('dd')">
                            <b-row>
                                <b-col cols=4 class="transition">
                                    <v-img :src="images.dd" :contain="true" class="card-icon"></v-img>
                                </b-col>

                                <b-col cols=8 class="">
                                    <div class="card-content">
                                        <div class="text">
                                            <h6 class="title-1">EDMp Data Dictionary</h6>
                                        </div>
                                    </div>
                                </b-col>
                            </b-row>

                            <h4 class="title-2">{{ store.counts ? store.counts.CDE_COUNT : 0 }}</h4>
                            <div class="card-circle-left transition"></div>
                        </div>
                    </b-col>
                
                    <b-col sm=12 md=5 lg=5 class="my-3">
                        <div class="card card-v3 transition link" @click="goToMenuContent('cr')">
                            <b-row>
                                <b-col cols=4 class="transition">
                                    <v-img :src="images.cr" :contain="true" class="card-icon"></v-img>
                                </b-col>

                                <b-col cols=8 class="">
                                    <div class="card-content">
                                        <div class="text">
                                            <h6 class="title-1">EDMp Consumption Report</h6>
                                        </div>
                                    </div>
                                </b-col>
                            </b-row>

                            <h4 class="title-2">{{ store.counts ? store.counts.DSP_COUNT : 0 }}</h4>
                            <div class="card-circle-left transition"></div>
                        </div>
                    </b-col>
                    <b-col sm=12 md=1 lg=1 />
                    <div class="clearfix"></div>
                </b-row> -->

                <b-row>
                    <b-col sm=12 md=4 lg=4 ></b-col>
                    <b-col sm=12 md=4 lg=4 >
                        <v-img :src="images.searchBg" :contain="true" class="card-icon"></v-img>
                        <div class="home-title">
                            <!-- <h3 class="title-1 mb-0">Welcome to</h3> -->
                            <h4 class="title-2 mt-5 mb-4">Metadata Search</h4>
                        </div>
                        <v-text-field
                            v-model="message"
                            :append-icon="message ? 'fa-arrow-right' : 'fa-search'"
                            solo
                            label="Search"
                            type="text"
                            @click:append="sendMessage"
                            @click:clear="clearMessage"
                            class="edmp-search"
                        ></v-text-field>

                        <div class="mt-5">
                            <b-button size="sm" class="dark-blue icon-only" @click="goToMenuContent('dd')">EDMp Data Dictionary</b-button>
                        </div>
                    </b-col>
                    <b-col sm=12 md=4 lg=4 ></b-col>
                </b-row>
                
            </div>

        </b-container>
    </v-content>
</template>

<script>
import PageHeader from '../PageHeader';
import { mapState, mapActions } from 'vuex'

export default {
    components: { PageHeader },
    data () {
        return {
            images: {
                dd: require('../../assets/images/icon-edmp-dd.png'),
                cr: require('../../assets/images/icon-edmp-cr.png'),
                searchBg: require('../../assets/images/icon-edmp-search.png'),
            },
            message: '',
        }
    },
    computed: {
        addressPathParent (){
            var tmp = this.$route.path.split("/")
            return tmp.slice(0, 2).join("/")
        },
        store () {
            return this.$store.state.edmp.all
        },
        icon () {
            return this.icons[this.iconIndex]
        },
    },
    mounted () {
        this.getCounts(this.$route.params.system);
    },
    methods: {
        getCounts (param) {
            this.$store.dispatch(`edmp/getCounts`, param)
        },
        goToMenuContent(param){
            this.$router.push(this.addressPathParent + '/' + encodeURIComponent(param));
        },

        sendMessage () {
            this.resetIcon()
            this.clearMessage()
        },
        clearMessage () {
            this.message = ''
        },
        resetIcon () {
            this.iconIndex = 0
        },
    }
}
</script>
