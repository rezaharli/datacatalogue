<style>
  @import '../assets/styles/login.css';
</style>

<template>
  <v-content>
    <v-container fluid fill-height back>
        <div id="login">
            <div id="description">
                <h1>Data Catalogue</h1>
                <p>Sign in to your account.</p>
            </div>

            <div id="form">
                <b-form @submit.prevent="doLogin">
                    <b-form-group id="exampleInputGroup1"
                            label="ID"
                            label-for="userid">
                        <b-form-input id="userid"
                            type="number"
                            v-model="userid"
                            required
                            placeholder="Enter ID"></b-form-input>
                    </b-form-group>

                    <b-form-group id="exampleInputGroup2"
                            label="Password"
                            label-for="password">
                        <div class="passw-container">
                            <v-icon color="#95a5a6" class="fas" v-text="passwordIcon" @click="hidePassword = !hidePassword">visibility</v-icon>
                            <b-form-input id="password"
                                :type="passwordType"
                                v-model="password"
                                required
                                placeholder="Enter password"></b-form-input>
                        </div>
                    </b-form-group>

                    <b-button type="submit" size="xs">Login</b-button>
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
            userid: '',
            password: '',
            hidePassword: true,
            submitted: false
        };
    },
    computed: {
        ...mapState('account', ['status']),
        passwordType() {
            return this.hidePassword ? 'password' : 'text'
        },
        passwordIcon() {
            return this.hidePassword ? 'visibility' : 'visibility_off'
        },
    },
    created () {
        jQuery('form').on('focus', 'input[type=number]', function () {
            jQuery(this).on('mousewheel.disableScroll', function (e) {
                e.preventDefault()
            })
        })

        jQuery('form').on('blur', 'input[type=number]', function () {
            jQuery(this).off('mousewheel.disableScroll')
        })

        var user = localStorage.getItem('user');
        if(user){
            this.$router.push("/")
        }
    },
    methods: {
        ...mapActions('account', ['login']),
        doLogin() {
            this.submitted = true;
            const { userid, password } = this;

            if (userid && password) {
                this.login({ userid, password })
            }
        }
    }
}
</script>