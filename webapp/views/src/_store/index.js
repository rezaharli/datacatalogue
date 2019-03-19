import Vue from 'vue';
import Vuex from 'vuex';

import { alert } from './alert.module';
import { account } from './account.module';
import { users } from './users.module';
import { dscmy } from './dscmy.module';
import { dscall } from './dscall.module';
import { dscinterfaces } from './dscinterfaces.module';
import { dpomy } from './dpomy.module';
import { ddomy } from './ddomy.module';
import { ddoall } from './ddoall.module';
import { rfomy } from './rfomy.module';

Vue.use(Vuex);

export const store = new Vuex.Store({
    modules: {
        alert,
        account,
        users,
        dscmy, dscall, dscinterfaces,
        dpomy,
        ddomy,
        ddoall,
        rfomy
    }
});