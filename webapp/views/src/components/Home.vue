<style>
  @import '../assets/styles/home.css';
</style>

<template>

  <v-content>
    <b-container fluid v-if="!isHome">
      <CToolbar></CToolbar>
      <router-view></router-view>
    </b-container>

    <v-container fluid fill-height back v-if="isHome">
      <v-layout align-center justify-center>
        <v-flex xs12 sm12 md12 card-container>
          <v-container grid-list-md text-xs-center>
            <v-layout row wrap align-center>
              <v-flex xs12>
                <v-card class="force-elevation-0" color="transparent">
                  <v-card-title primary-title class="justify-center white--text">
                    <div>
                      <h1 class="mb-0">Welcome to</h1>
                      <h1 class="mb-0">Data Catalogue</h1>
                    </div>
                  </v-card-title>
                </v-card>
              </v-flex>

              <v-flex xs3 sm3 md3 card-container>
                <v-card flat class="rounded-card" 
                  :class="{ [`elevation-${elevation.dsc}`]: true }" 
                  @click="dscClick()" 
                  @mouseover="cardMouseover('dsc')"
                  @mouseleave="cardMouseleave('dsc')">
                  <v-img :src="images.dsc" class="menu-image"></v-img>
                  
                  <v-card-title primary-title class="justify-center">
                    <div>
                      <h3 class="mb-0">Dataset</h3>
                      <h3 class="mb-0">Custodian</h3>
                    </div>
                  </v-card-title>
                </v-card>
              </v-flex>

              <v-flex xs3 sm3 md3 card-container>
                <v-card flat class="rounded-card" 
                  :class="{ [`elevation-${elevation.dpo}`]: true }" 
                  @click="dpoClick()" 
                  @mouseover="cardMouseover('dpo')"
                  @mouseleave="cardMouseleave('dpo')">
                  <v-img :src="images.dpo" class="menu-image"></v-img>

                  <v-card-title primary-title class="justify-center">
                    <div>
                      <h3 class="mb-0">Downstream</h3>
                      <h3 class="mb-0">Process Owner</h3>
                    </div>
                  </v-card-title>
                </v-card>
              </v-flex>

              <v-flex xs3 sm3 md3 card-container>
                <v-card flat class="rounded-card" 
                  :class="{ [`elevation-${elevation.ddo}`]: true }" 
                  @click="ddoClick()" 
                  @mouseover="cardMouseover('ddo')"
                  @mouseleave="cardMouseleave('ddo')">
                  <v-img :src="images.ddo" class="menu-image"></v-img>

                  <v-card-title primary-title class="justify-center">
                    <div>
                      <h3 class="mb-0">Data Domain</h3>
                      <h3 class="mb-0">Owner</h3>
                    </div>
                  </v-card-title>
                </v-card>
              </v-flex>

              <v-flex xs3 sm3 md3 card-container>
                <v-card flat class="rounded-card" 
                  :class="{ [`elevation-${elevation.rfo}`]: true }" 
                  @click="rfoClick()" 
                  @mouseover="cardMouseover('rfo')"
                  @mouseleave="cardMouseleave('rfo')">
                  <v-img :src="images.rfo" class="menu-image"></v-img>

                  <v-card-title primary-title class="justify-center">
                    <div>
                      <h3 class="mb-0">Risk Framework</h3>
                      <h3 class="mb-0">Owner</h3>
                    </div>
                  </v-card-title>
                </v-card>
              </v-flex>
            </v-layout>
          </v-container>
        </v-flex>
      </v-layout>
    </v-container>
  </v-content>
</template>

<script>
import CToolbar from './ComponentToolbar';

export default {
  components: { CToolbar },
  data() {
    return {
      isHome: this.decideIsHome(),
      elevation: {
        dsc: 12,
        dpo: 12,
        ddo: 12,
        rfo: 12,
      },
      images: {
        dsc: require('../assets/images/u23.png'),
        dpo: require('../assets/images/u27.png'),
        ddo: require('../assets/images/u29.png'),
        rfo: require('../assets/images/u31.png'),
      }
    };
  },
  created () {
    this.decideIsHome();
  },
  watch: {
      $route (to){
        this.decideIsHome();
      },
    },
  methods: {
    decideIsHome: function(){
      this.isHome = this.$route.name == "landingpage" ? true : false;
    },
    dscClick: function () {
      this.$router.push('/dsc');
    },
    dpoClick: function () {
      this.$router.push('/dpo');
    },
    ddoClick: function () {
      this.$router.push('/ddo');
    },
    rfoClick: function () {
      this.$router.push('/rfo');
    },
    cardMouseover: function(type){
      this.elevation[type] = 24;
    },
    cardMouseleave: function(type){
      this.elevation[type] = 12;
    },
  }
}
</script>