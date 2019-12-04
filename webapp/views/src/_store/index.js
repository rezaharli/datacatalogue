import Vue from 'vue';
import Vuex from 'vuex';

import { alert } from './alert.module';
import { account } from './account.module';
import { crypto } from './crypto.module';
import { config } from './config.module';
import { header } from './header.module';
import { exportData } from './exportdata.module';

import { users } from './users.module';
import { usersusage } from './usersusage.module';

import { edmp } from './edmp.module';
import { edmpddTechnical } from './edmpdd-technical.module';
import { edmpddBusiness } from './edmpdd-business.module';
import { edmpddConsumption } from './edmpdd-consumption.module';
import { edmpIarcPersonal } from './edmpiarc-personal.module';

Vue.use(Vuex);

export const store = new Vuex.Store({
    modules: {
        alert,
        account,
        crypto,
        config,
        header,
        exportData,
        users, usersusage,
        edmp, edmpIarcPersonal, edmpddTechnical, edmpddBusiness, edmpddConsumption,
    }
});