<style>
  @import '@/assets/styles/login.css';
</style>

<template>
  <v-content>
    <v-container fluid fill-height back>
        <div id="login">
            <div id="description">
                <h1>EDMP Metadata Catalogue</h1>
                <p>Sign in to your account.</p>
            </div>

            <div id="form">
        <!-- <v-alert :value="alert.message? true : false" :type="alert.type">{{alert.message}}</v-alert> -->
                <b-form @submit.prevent="doLogin">
                    <b-form-group id="exampleInputGroup1"
                            label="ID"
                            label-for="username">
                        <b-form-input id="username"
                            type="number"
                            v-model="username"
                            required
                            placeholder="Enter ID"></b-form-input>
                    </b-form-group>

                    <b-form-group id="exampleInputGroup2"
                            label="Password"
                            label-for="password">
                        <div class="passw-container">
                            <!-- <v-icon color="#95a5a6" class="fas" v-text="passwordIcon" @click="hidePassword = !hidePassword">visibility</v-icon> -->
                            <i class="fa fas" v-bind:class="{'fa-eye' : hidePassword, 'fa-eye-slash' : !hidePassword }" @click="hidePassword = !hidePassword"></i>
                            <b-form-input id="password"
                                :type="passwordType"
                                :disabled="isPasswordDisabled"
                                v-model="password"
                                placeholder="Enter password"></b-form-input>
                        </div>
                    </b-form-group>

                    <b-alert :show="alert.message? true : false" :variant="alert.type">{{alert.message}}</b-alert>
                    <b-button type="submit" size="xs" :disabled="status.loggingIn">Login</b-button>
                    <img v-show="status.loggingIn" src="data:image/gif;base64,R0lGODlhEAAQAPIAAP///wAAAMLCwkJCQgAAAGJiYoKCgpKSkiH/C05FVFNDQVBFMi4wAwEAAAAh/hpDcmVhdGVkIHdpdGggYWpheGxvYWQuaW5mbwAh+QQJCgAAACwAAAAAEAAQAAADMwi63P4wyklrE2MIOggZnAdOmGYJRbExwroUmcG2LmDEwnHQLVsYOd2mBzkYDAdKa+dIAAAh+QQJCgAAACwAAAAAEAAQAAADNAi63P5OjCEgG4QMu7DmikRxQlFUYDEZIGBMRVsaqHwctXXf7WEYB4Ag1xjihkMZsiUkKhIAIfkECQoAAAAsAAAAABAAEAAAAzYIujIjK8pByJDMlFYvBoVjHA70GU7xSUJhmKtwHPAKzLO9HMaoKwJZ7Rf8AYPDDzKpZBqfvwQAIfkECQoAAAAsAAAAABAAEAAAAzMIumIlK8oyhpHsnFZfhYumCYUhDAQxRIdhHBGqRoKw0R8DYlJd8z0fMDgsGo/IpHI5TAAAIfkECQoAAAAsAAAAABAAEAAAAzIIunInK0rnZBTwGPNMgQwmdsNgXGJUlIWEuR5oWUIpz8pAEAMe6TwfwyYsGo/IpFKSAAAh+QQJCgAAACwAAAAAEAAQAAADMwi6IMKQORfjdOe82p4wGccc4CEuQradylesojEMBgsUc2G7sDX3lQGBMLAJibufbSlKAAAh+QQJCgAAACwAAAAAEAAQAAADMgi63P7wCRHZnFVdmgHu2nFwlWCI3WGc3TSWhUFGxTAUkGCbtgENBMJAEJsxgMLWzpEAACH5BAkKAAAALAAAAAAQABAAAAMyCLrc/jDKSatlQtScKdceCAjDII7HcQ4EMTCpyrCuUBjCYRgHVtqlAiB1YhiCnlsRkAAAOwAAAAAAAAAAAA==" />
                </b-form>
            </div>
        </div>
    </v-container>
  </v-content>
</template>

<script>
import jQuery from 'jquery'

export default {
    data() {
        return {
            username: '',
            password: '',
            hidePassword: true,
            submitted: false
        };
    },
    computed: {
        status () { return this.$store.state.account.status },
        alert () { return this.$store.state.alert },
        configStore () { return this.$store.state.config.all },
        isPasswordDisabled () {
            var isLDAPEnabled = this.configStore.config.IsEnabled
            return isLDAPEnabled == "true" || isLDAPEnabled == "false" ? false : true;
        },
        passwordType() {
            return this.hidePassword ? 'password' : 'text'
        },
        passwordIcon() {
            return this.hidePassword ? 'visibility' : 'visibility_off'
        },
    },
    created () {
        this.getConfig("LDAP");

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
        login (param) { this.$store.dispatch(`account/login`, param) },
        getConfig (param) { this.$store.dispatch(`config/getConfig`, param) },
        doLogin() {
            this.submitted = true;
            const { username, password } = this;

            if (username) {
                this.login({ username, password })
            }
        }
    }
}
</script>