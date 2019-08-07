import Vue from 'vue';
import Vuex from 'vuex';

import { alert } from './alert.module';
import { account } from './account.module';
import { crypto } from './crypto.module';
import { config } from './config.module';
import { users } from './users.module';
import { usersusage } from './usersusage.module';
import { dsc } from './dsc.module';
import { dscmy } from './dscmy.module';
import { dscall } from './dscall.module';
import { dsccde } from './dsccde.module';
import { dsccdp } from './dsccdp.module';
import { dsccdpcde } from './dsccdpcde.module';
import { dsciarc } from './dsciarc.module';
import { dsciarcpersonal } from './dsciarcpersonal.module';
import { dscinterfaces } from './dscinterfaces.module';
import { dscinterfacescde } from './dscinterfacescde.module';
import { dscdd } from './dscdd.module';
import { dscddTechnical } from './dscdd-technical.module';
import { dscddBusiness } from './dscdd-business.module';
import { dscddPolicy } from './dscdd-policy.module';
import { dpo } from './dpo.module';
import { dpodataelements } from './dpodataelements.module';
import { dpodatalineage } from './dpodatalineage.module';
import { dpomy } from './dpomy.module';
import { dpoall } from './dpoall.module';
import { ddo } from './ddo.module';
import { ddomy } from './ddomy.module';
import { ddoall } from './ddoall.module';
import { ddobusinessterm } from './ddobusinessterm.module';
import { ddosystems } from './ddosystems.module';
import { ddosystemsbusinessterm } from './ddosystemsbusinessterm.module';
import { ddodownstream } from './ddodownstream.module';
import { ddodownstreambusinessterm } from './ddodownstreambusinessterm.module';
import { rfo } from './rfo.module';
import { rfomy } from './rfomy.module';
import { rfoall } from './rfoall.module';
import { rfosummary } from './rfosummary.module';
import { rfopriority } from './rfopriority.module';
import { edmp } from './edmp.module';
import { edmpddTechnical } from './edmpdd-technical.module';
import { edmpddBusiness } from './edmpdd-business.module';
import { edmpddOriginal } from './edmpdd-original.module';
import { edmpddPresence } from './edmpdd-presence.module';
import { edmpddProtection } from './edmpdd-protection.module';
import { edmpddAttribute } from './edmpdd-attribute.module';

Vue.use(Vuex);

export const store = new Vuex.Store({
    modules: {
        alert,
        account,
        crypto,
        config,
        users, usersusage,
        dsc, dscmy, dscall, dsccde, dsccdp, dsccdpcde, dscinterfaces, dscinterfacescde, dsciarc, dsciarcpersonal, dscdd, dscddTechnical, dscddBusiness, dscddPolicy,
        dpo, dpomy, dpoall, dpodataelements, dpodatalineage,
        ddo, ddomy, ddoall, ddobusinessterm, ddosystems, ddosystemsbusinessterm, ddodownstream, ddodownstreambusinessterm,
        rfo, rfomy, rfoall, rfosummary, rfopriority,
        edmp, edmpddTechnical, edmpddBusiness, edmpddOriginal, edmpddPresence, edmpddProtection, edmpddAttribute,
    }
});