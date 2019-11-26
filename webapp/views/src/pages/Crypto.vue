<style>
  @import '../assets/styles/login.css';
</style>

<style scoped>
div#app div#login div#form {
    width: 600px;
}
</style>


<template>
  <v-content>
    <v-container fluid fill-height back>
        <div id="login">
            <div id="form">
                <b-form>
                    <b-form-group id="input-plaintext"
                            label=" "
                            label-for="plainText">
                        <div class="passw-container">
                            <!-- <i class="fa fas" v-bind:class="{'fa-eye' : hidePassword, 'fa-eye-slash' : !hidePassword }" @click="hidePassword = !hidePassword"></i> -->
                            <b-form-input id="plainText"
                                :type="passwordType"
                                v-model="store.plainText"
                                placeholder="Enter Password"></b-form-input>
                        </div>
                    </b-form-group>

                    <b-button style="float: right" size="xs" :disabled="store.isLoading" @click="encrypt">Encrypt</b-button>
                    <img v-show="store.isLoading" src="data:image/gif;base64,R0lGODlhEAAQAPIAAP///wAAAMLCwkJCQgAAAGJiYoKCgpKSkiH/C05FVFNDQVBFMi4wAwEAAAAh/hpDcmVhdGVkIHdpdGggYWpheGxvYWQuaW5mbwAh+QQJCgAAACwAAAAAEAAQAAADMwi63P4wyklrE2MIOggZnAdOmGYJRbExwroUmcG2LmDEwnHQLVsYOd2mBzkYDAdKa+dIAAAh+QQJCgAAACwAAAAAEAAQAAADNAi63P5OjCEgG4QMu7DmikRxQlFUYDEZIGBMRVsaqHwctXXf7WEYB4Ag1xjihkMZsiUkKhIAIfkECQoAAAAsAAAAABAAEAAAAzYIujIjK8pByJDMlFYvBoVjHA70GU7xSUJhmKtwHPAKzLO9HMaoKwJZ7Rf8AYPDDzKpZBqfvwQAIfkECQoAAAAsAAAAABAAEAAAAzMIumIlK8oyhpHsnFZfhYumCYUhDAQxRIdhHBGqRoKw0R8DYlJd8z0fMDgsGo/IpHI5TAAAIfkECQoAAAAsAAAAABAAEAAAAzIIunInK0rnZBTwGPNMgQwmdsNgXGJUlIWEuR5oWUIpz8pAEAMe6TwfwyYsGo/IpFKSAAAh+QQJCgAAACwAAAAAEAAQAAADMwi6IMKQORfjdOe82p4wGccc4CEuQradylesojEMBgsUc2G7sDX3lQGBMLAJibufbSlKAAAh+QQJCgAAACwAAAAAEAAQAAADMgi63P7wCRHZnFVdmgHu2nFwlWCI3WGc3TSWhUFGxTAUkGCbtgENBMJAEJsxgMLWzpEAACH5BAkKAAAALAAAAAAQABAAAAMyCLrc/jDKSatlQtScKdceCAjDII7HcQ4EMTCpyrCuUBjCYRgHVtqlAiB1YhiCnlsRkAAAOwAAAAAAAAAAAA==" />

                    <b-form-group id="input-encrypted"
                            label=" "
                            label-for="encrypted">
                        <b-form-input id="encrypted"
                            disabled
                            type="text"
                            v-model="store.encrypted"
                            placeholder="Enter Encrypted Text"></b-form-input>
                    </b-form-group>
                </b-form>
            </div>
        </div>
    </v-container>
  </v-content>
</template>

<script>
import jQuery from 'jquery'
import { mapState, mapActions } from 'vuex'

export default {
    data() {
        return {
            submitted: false,
            hidePassword: true,
        };
    },
    computed: {
        store () { return this.$store.state.crypto.all },
        passwordType() {
            return this.hidePassword ? 'password' : 'text'
        },
    },
    methods: {
        encrypt () {
            this.$store.dispatch(`crypto/encrypt`)
        },
        decrypt () {
            this.$store.dispatch(`crypto/decrypt`)
        },
    }
}
</script>