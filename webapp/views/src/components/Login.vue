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
                            label="Email"
                            label-for="email">
                        <b-form-input id="email"
                            type="text"
                            v-model="email"
                            required
                            placeholder="Enter email"></b-form-input>
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
export default {
    data() {
        return {
            email: '',
            password: '',
            hidePassword: true
        };
    },
    computed: {
        passwordType() {
            return this.hidePassword ? 'password' : 'text'
        },
        passwordIcon() {
            return this.hidePassword ? 'visibility' : 'visibility_off'
        }
    },
    methods: {
        doLogin() {
            this.$router.push("/")
        }
    }
}
</script>