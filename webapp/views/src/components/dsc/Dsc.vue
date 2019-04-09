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
                            <v-btn class="float-right" color="red" dark>
                                <v-icon dark>filter_list</v-icon>
                            </v-btn>
                            
                            <v-btn class="float-right" color="success" dark>
                                <v-icon dark>file_copy</v-icon>
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
import PageHeader from '../PageHeader';

import dscMy from './Dsc-my.vue';
import dscAll from './Dsc-all.vue';

export default {
    components: { PageHeader, dscMy, dscAll },
    data () {
      return {}
    },
}
</script>