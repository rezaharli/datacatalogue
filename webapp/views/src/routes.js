import Vue from 'vue';
import VueRouter from 'vue-router';

import Home from './components/Home';

import Login from './components/Login';
import Crypto from './components/Crypto';

import Dsc from './components/dsc/Dsc';
import DscMenu from './components/dsc/Dsc-menu';
import DscCde from './components/dsc/Dsc-cde';
import DscCdp from './components/dsc/Dsc-cdp';
import DscCdpCde from './components/dsc/Dsc-cdp-cde';
import DscInterfaces from './components/dsc/Dsc-interfaces';
import DscInterfacesCde from './components/dsc/Dsc-interfaces-cde';
import DscDd from './components/dsc/Dsc-dd';
import DscDetails from './components/dsc/Dsc-details';

import Dpo from './components/dpo/Dpo';
import DpoMenu from './components/dpo/Dpo-menu';
import DpoDataelements from './components/dpo/Dpo-dataelements';
import DpoDatalineage from './components/dpo/Dpo-datalineage';
import DpoDetails from './components/dpo/Dpo-details';

import Ddo from './components/ddo/Ddo';
import DdoMenu from './components/ddo/Ddo-menu';
import DdoBusinessterm from './components/ddo/Ddo-businessterm';
import DdoSystems from './components/ddo/Ddo-systems';
import DdoSystemsBusinessterm from './components/ddo/Ddo-systems-businessterm';
import DdoDownstream from './components/ddo/Ddo-downstream';
import DdoDownstreamBusinessterm from './components/ddo/Ddo-downstream-businessterm';
import DdoDetails from './components/ddo/Ddo-details';

import Rfo from './components/rfo/Rfo';
import RfoPriority from './components/rfo/Rfo-priority';

import Access from './components/access/Access';
import AccessUsers from './components/access/Access-users';
import AccessRoles from './components/access/Access-roles';
import AccessUsage from './components/access/Access-usage';

import {store} from './_store/index'

Vue.use(VueRouter);

const router = new VueRouter({
  mode: 'history',
  routes: [{ // home
    path: '/', name: 'landingpage', component: Home, 
    meta: { 
      title: "Home - Data Catalogue" 
    },
  }, { // dsc
    path: '/dsc', name: 'dsc', component: Dsc, 
    meta: { 
      title: "DSC - Data Catalogue",
      permission: "DSC"
    },
  }, { // dsc.menu
    path: '/dsc/:system', name: 'dsc.menu', component: DscMenu, 
    meta: { 
      title: "DSC - Data Catalogue",
      permission: "DSC"
    },
  }, { // dsc.cde
    path: '/dsc/cde/:system', name: 'dsc.cde', component: DscCde, 
    meta: { 
      title: "DSC - Data Catalogue",
      permission: "DSC"
    },
    children: [{ // dsc.details
      path: ':systemID/:details/:column', name: 'dsc.details', component: DscDetails,
      meta: { 
        title: "DSC Details - Data Catalogue",
        showModal: true,
        permission: "DSC"
      }
    }]
  }, { // dsc.cdp
    path: '/dsc/cdp/:system', name: 'dsc.cdp', component: DscCdp, 
    meta: { 
      title: "DSC - Data Catalogue",
      permission: "DSC"
    }
  }, { // dsc.cdp.cde
    path: '/dsc/cdp/:system/:dspName', name: 'dsc.cdp.cde', component: DscCdpCde, 
    meta: { 
      title: "DSC - Data Catalogue",
      permission: "DSC"
    },
    children: [{ // dsc.details
      path: ':systemID/:details/:column', name: 'dsc.details', component: DscDetails,
      meta: { 
        title: "DSC Details - Data Catalogue",
        showModal: true,
        permission: "DSC"
      }
    }] 
  }, { // dsc.interfaces
    path: '/dsc/interfaces/:system', name: 'dsc.interfaces', component: DscInterfaces, 
    meta: { 
      title: "DSC - Data Catalogue",
      permission: "DSC"
    }
  },{ // dsc.interfaces.cde
    path: '/dsc/interfaces/:system/:dspName', name: 'dsc.interfaces.cde', component: DscInterfacesCde, 
    meta: { 
      title: "DSC - Data Catalogue",
      permission: "DSC"
    },
    children: [{ // dsc.details
      path: ':systemID/:details/:column', name: 'dsc.details', component: DscDetails,
      meta: { 
        title: "DSC Details - Data Catalogue",
        showModal: true,
        permission: "DSC"
      }
    }] 
  }, { // dsc.dd
    path: '/dsc/dd/:system', name: 'dsc.dd', component: DscDd, 
    meta: { 
      title: "DSC - Data Catalogue",
      permission: "DSC"
    },
  },  
  { // dpo
    path: '/dpo', name: 'dpo', component: Dpo, 
    meta: { 
      title: "DPO - Data Catalogue",
      permission: "DPO"
    },
  }, { // dpo.menu
    path: '/dpo/:system', name: 'dpo.menu', component: DpoMenu, 
    meta: { 
      title: "DPO - Data Catalogue",
      permission: "DPO"
    },
  }, { // dpo.dataelements
    path: '/dpo/dataelements/:dspname', name: 'dpo.dataelements', component: DpoDataelements, 
    meta: { 
      title: "DPO - Data Catalogue",
      permission: "DPO"
    },
    children: [{ // dpo.details
      path: ':cdename', name: 'dpo.details', component: DpoDetails,
      meta: { 
        title: "DPO Details - Data Catalogue",
        showModal: true,
        permission: "DPO"
      } 
    }] 
  }, { // dpo.datalineage
    path: '/dpo/datalineage/:dspname', name: 'dpo.datalineage', component: DpoDatalineage, 
    meta: { 
      title: "DPO - Data Catalogue",
      permission: "DPO"
    },
  },
  { // ddo
    path: '/ddo', name: 'ddo', component: Ddo, 
    meta: { 
      title: "DDO - Data Catalogue",
      permission: "DSC"
    },
  }, { // ddo.menu
    path: '/ddo/:system', name: 'ddo.menu', component: DdoMenu, 
    meta: { 
      title: "DDO - Data Catalogue",
      permission: "DDO"
    },
  }, { // ddo.businessterm
    path: '/ddo/businessterm/:subdomain', name: 'ddo.businessterm', component: DdoBusinessterm, 
    meta: { 
      title: "DDO - Data Catalogue",
      permission: "DDO"
    },
    children: [{ // ddo.details
      path: ':btname', name: 'ddo.details', component: DdoDetails,
      meta: { 
        title: "DDO Details - Data Catalogue",
        showModal: true,
        permission: "DDO"
      } 
    }] 
  }, { // ddo.systems
    path: '/ddo/systems/:subdomain', name: 'ddo.systems', component: DdoSystems, 
    meta: { 
      title: "DDO - Data Catalogue",
      permission: "DDO"
    }
  }, { // ddo.systems.businessterm
    path: '/ddo/systems/:subdomain/:system', name: 'ddo.systems.businessterm', component: DdoSystemsBusinessterm, 
    meta: { 
      title: "DDO - Data Catalogue",
      permission: "DDO"
    },
    children: [{ // ddo.details
      path: ':btname', name: 'ddo.details', component: DdoDetails,
      meta: { 
        title: "DDO Details - Data Catalogue",
        showModal: true,
        permission: "DDO"
      } 
    }] 
  }, { // ddo.downstream
    path: '/ddo/downstream/:subdomain', name: 'ddo.downstream', component: DdoDownstream, 
    meta: { 
      title: "DDO - Data Catalogue",
      permission: "DDO"
    }
  }, { // dsc.downstream.businessterm
    path: '/ddo/downstream/:subdomain/:system', name: 'ddo.downstream.businessterm', component: DdoDownstreamBusinessterm, 
    meta: { 
      title: "DDO - Data Catalogue",
      permission: "DDO"
    },
    children: [{ // ddo.details
      path: ':btname', name: 'ddo.details', component: DdoDetails,
      meta: { 
        title: "DDO Details - Data Catalogue",
        showModal: true,
        permission: "DDO"
      } 
    }] 
  }, { // rfo
    path: '/rfo', name: 'rfo', component: Rfo, 
    meta: { 
      title: "RFO - Data Catalogue",
      permission: "RFO"
    },
  }, { // rfo.priority
    path: '/rfo/:type', name: 'rfo.priority', component: RfoPriority, 
    meta: { 
      title: "RFO - Data Catalogue",
      permission: "RFO"
    },
  }, { // access
    path: '/access', component: Access, 
    meta: { 
      title: "Access - Data Catalogue" ,
      permission: "Admin"
    }, 
    children: [{ 
      path: '', name: 'access', redirect: { name: 'access.users' }
    }, { //access.users
      path: 'users', name: 'access.users', component: AccessUsers, 
      meta: { 
        title: "Users - Data Catalogue",
        permission: "Admin"
      } 
    }, { //access.roles
      path: 'roles', name: 'access.roles', component: AccessRoles, 
      meta: { 
        title: "Roles - Data Catalogue",
        permission: "Admin"
      } 
    }, { //access.usage
      path: 'usage', name: 'access.usage', component: AccessUsage, 
      meta: { 
        title: "Usage Detail - Data Catalogue",
        permission: "Admin"
      } 
    }]
  }, {
    path: '/login', name: 'login', component: Login, 
  }, {
    path: '/crypto', name: 'crypto', component: Crypto, 
  },
  { path: '*', redirect: '/' }]
});

router.beforeEach((to, from, next) => {
  document.title = (to.meta.title || '')
  next()
});

var checkSession = () => {
  return store.dispatch(`account/checkSession`);
}

var saveUsage = function(isAuth, to, from, next){
  var saveLog = (param) => { store.dispatch(`users/saveLog`, param); }
  
  var user = store.state.account.user;
  if(user)
    saveLog({
      Username: user.Username,
      Fullname: user.Name,
      Role: user.Role,
      Module: to.name,
      Description: isAuth ? "SUCCESS" : "Permission Denied",
      ResourceUrl: to.fullPath
    });
}

router.beforeEach((to, from, next) => {
  checkSession().then(
    res => {
      var isSession = store.state.account.isSession;
      
      if( ! isSession){
        localStorage.removeItem('user');
      }

      // redirect to login page if not logged in and trying to access a restricted page
      const publicPages = ['/login', '/crypto'];
      const authRequired = !publicPages.find(v => to.path.indexOf(v) != -1);
      const loggedIn = localStorage.getItem('user');
      
      if (authRequired && !loggedIn) {
        saveUsage(false, to, from, next);
        return next('/login');
      } else {
        if(to.name != "landingpage" && authRequired){
          var user = store.state.account.user;
          if(user)
            if(to.name.split(".")[0].toLowerCase() == "access"){
              if(user.Role.toLowerCase().split(",").indexOf("Admin".toLowerCase()) == -1){
                saveUsage(false, to, from, next);
                return next('/');
              }
            } else {
              if(user.Role.toLowerCase().split(",").indexOf(to.name.split(".")[0].toLowerCase()) == -1){
                saveUsage(false, to, from, next);
                return next('/');
              }
            }
        }

        saveUsage(true, to, from, next)
        next();
      }
    }
  )
})

export default router