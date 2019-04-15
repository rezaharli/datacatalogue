import Vue from 'vue';
import Vuex from 'vuex';

import { alert } from './alert.module';
import { account } from './account.module';
import { users } from './users.module';
import { dsc } from './dsc.module';
import { dscmy } from './dscmy.module';
import { dscall } from './dscall.module';
import { dsccde } from './dsccde.module';
import { dsccdp } from './dsccdp.module';
import { dsccdpcde } from './dsccdpcde.module';
import { dscinterfaces } from './dscinterfaces.module';
import { dpomy } from './dpomy.module';
import { dpoall } from './dpoall.module';
import { ddomy } from './ddomy.module';
import { ddoall } from './ddoall.module';
import { rfomy } from './rfomy.module';
import { rfoall } from './rfoall.module';

Vue.use(Vuex);

export const store = new Vuex.Store({
    modules: {
        alert,
        account,
        users,
        dsc, dscmy, dscall, dsccde, dsccdp, dsccdpcde, dscinterfaces,
        dpomy, dpoall,
        ddomy, ddoall,
        rfomy, rfoall
    }
});