import Vue from 'vue';
import VueRouter from 'vue-router';

import Login from '../pages/Login';
import Crypto from '../pages/Crypto';

import Dsc from '../pages/dsc/Dsc';
import DscMenu from '../pages/dsc/Dsc-menu';
import DscCde from '../pages/dsc/Dsc-cde';
import DscCdp from '../pages/dsc/Dsc-cdp';
import DscCdpCde from '../pages/dsc/Dsc-cdp-cde';
import DscIarc from '../pages/dsc/Dsc-iarc';
import DscIarcInformation from '../pages/dsc/Dsc-iarc-information';
import DscIarcPersonal from '../pages/dsc/Dsc-iarc-personal';
import DscInterfaces from '../pages/dsc/Dsc-interfaces';
import DscInterfacesCde from '../pages/dsc/Dsc-interfaces-cde';
import DscDd from '../pages/dsc/Dsc-dd';
import DscDdTechnical from '../pages/dsc/Dsc-dd-technical';
import DscDdBusiness from '../pages/dsc/Dsc-dd-business';
import DscDdPolicy from '../pages/dsc/Dsc-dd-policy';
import DscDetails from '../pages/dsc/Dsc-details';

import Edmp from '../pages/dsc/edmp/Edmp';
import EdmpDd from '../pages/dsc/edmp/Edmp-dd';
import EdmpDdTechnical from '../pages/dsc/edmp/Edmp-dd-technical';
import EdmpDdBusiness from '../pages/dsc/edmp/Edmp-dd-business';
import EdmpIarc from '../pages/dsc/edmp/Edmp-iarc';
import EdmpIarcPersonal from '../pages/dsc/edmp/Edmp-iarc-personal';

import Access from '../pages/access/Access';
import AccessUsers from '../pages/access/Access-users';
import AccessRoles from '../pages/access/Access-roles';
import AccessUsage from '../pages/access/Access-usage';

import {store} from '../_store/index'

Vue.use(VueRouter);

const router = new VueRouter({
  mode: 'history',
  routes: [{ // home
    path: '/', name: 'landingpage', redirect: { name: 'dsc' }
  }, { // dsc
    path: '/dsc', name: 'dsc', component: Dsc, 
    meta: { 
      title: "DSC - Data Catalogue",
      permission: "DSC"
    },
  }, { // dsc.edmp
    path: '/dsc/ENTERPRISE DATA MGMT PLATFORM', name: 'dsc.edmp', component: Edmp, 
    meta: { 
      title: "EDMp - Data Catalogue",
      permission: "DSC"
    },
  }, { // dsc.edmp.dd
    path: '/dsc/ENTERPRISE DATA MGMT PLATFORM/dd', name: 'dsc.edmp.dd', component: EdmpDd, 
    meta: { 
      title: "EDMp - Data Catalogue",
      permission: "DSC"
    },
    children: [
      { 
        path: '', name: 'dsc.edmp.dd.default', redirect: { name: 'dsc.edmp.dd.technical' }
      }, { // dsc.edmp.dd.technical
        path: '/dsc/ENTERPRISE DATA MGMT PLATFORM/dd/technical-metadata', name: 'dsc.edmp.dd.technical', component: EdmpDdTechnical, 
        meta: { 
          title: "EDMp - Data Catalogue",
          permission: "DSC"
        },
      }, { // dsc.edmp.dd.business
        path: '/dsc/ENTERPRISE DATA MGMT PLATFORM/dd/business-metadata', name: 'dsc.edmp.dd.business', component: EdmpDdBusiness, 
        meta: { 
          title: "EDMp - Data Catalogue",
          permission: "DSC"
        },
      }, 
    ]
  }, { // dsc.edmp.iarc
    path: '/dsc/ENTERPRISE DATA MGMT PLATFORM/iarc', name: 'dsc.edmp.iarc', component: EdmpIarc, 
    meta: { 
      title: "EDMp - Data Catalogue",
      permission: "DSC"
    },
    children: [
      { 
        path: '', name: 'dsc.edmp.iarc', redirect: { name: 'dsc.edmp.iarc.personal' }
      }, { // dsc.iarc.personal
        path: '/dsc/ENTERPRISE DATA MGMT PLATFORM/iarc/personal', name: 'dsc.edmp.iarc.personal', component: EdmpIarcPersonal, 
        meta: { 
          title: "DSC - Data Catalogue",
          permission: "DSC"
        },
      }, 
    ]
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
  }, { // dsc.iarc
    path: '/dsc/iarc/:system', name: 'dsc.iarc', component: DscIarc, 
    meta: { 
      title: "DSC - Data Catalogue",
      permission: "DSC"
    },
    children: [
      { 
        path: '', name: 'dsc.iarc', redirect: { name: 'dsc.iarc.information' }
      }, { // dsc.iarc.information
        path: '/dsc/iarc/information/:system', name: 'dsc.iarc.information', component: DscIarcInformation, 
        meta: { 
          title: "DSC - Data Catalogue",
          permission: "DSC"
        },
      }, { // dsc.iarc.personal
        path: '/dsc/iarc/personal/:system', name: 'dsc.iarc.personal', component: DscIarcPersonal, 
        meta: { 
          title: "DSC - Data Catalogue",
          permission: "DSC"
        },
      },
    ]
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
    children: [
      { 
        path: '', name: 'dsc.dd', redirect: { name: 'dsc.dd.technical' }
      }, { // dsc.dd.technical
        path: '/dsc/dd/technical-metadata/:system', name: 'dsc.dd.technical', component: DscDdTechnical, 
        meta: { 
          title: "DSC - Data Catalogue",
          permission: "DSC"
        },
      }, { // dsc.dd.business
        path: '/dsc/dd/business-metadata/:system', name: 'dsc.dd.business', component: DscDdBusiness, 
        meta: { 
          title: "DSC - Data Catalogue",
          permission: "DSC"
        },
      }, { // dsc.dd.policy
        path: '/dsc/dd/policy-related/:system', name: 'dsc.dd.policy', component: DscDdPolicy, 
        meta: { 
          title: "DSC - Data Catalogue",
          permission: "DSC"
        },
      }, 
    ]
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

router.beforeEach((to, _from, next) => {
  document.title = (to.meta.title || '')
  next()
});

var checkSession = () => {
  return store.dispatch(`account/checkSession`);
}

var saveUsage = function(isAuth, to){
  var saveLog = (param) => { store.dispatch(`users/saveLog`, param); }
  
  var user = store.state.account.user;
  if(user)
    saveLog({
      Username: user.USERNAME,
      Fullname: user.NAME,
      Role: user.ROLE,
      Module: to.name,
      Description: isAuth ? "SUCCESS" : "Permission Denied",
      ResourceUrl: to.fullPath
    });
}

router.beforeEach((to, from, next) => {
  checkSession().then(() => {
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
            if(user.ROLE.toLowerCase().split(",").indexOf("Admin".toLowerCase()) == -1){
              saveUsage(false, to, from, next);
              return next('/');
            }
          } else {
            if(user.ROLE.toLowerCase().split(",").indexOf(to.name.split(".")[0].toLowerCase()) == -1){
              saveUsage(false, to, from, next);
              return next('/');
            }
          }
      }

      saveUsage(true, to, from, next)
      next();
    }
  })
})

export default router