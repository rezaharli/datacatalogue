import Vue from 'vue';
import VueRouter from 'vue-router';

import Home from './components/Home';

import Login from './components/Login';

import Dsc from './components/dsc/Dsc';
import DscMenu from './components/dsc/Dsc-menu';
import DscCde from './components/dsc/Dsc-cde';
import DscCdp from './components/dsc/Dsc-cdp';
import DscCdpCde from './components/dsc/Dsc-cdp-cde';
import DscDetails from './components/dsc/Dsc-details';
import DscMy from './components/dsc/Dsc-my';
import DscAll from './components/dsc/Dsc-all';
import DscInterfaces from './components/dsc/Dsc-interfaces';

import Dpo from './components/dpo/Dpo';
import DpoDetails from './components/dpo/Dpo-details';
import DpoMy from './components/dpo/Dpo-my';
import DpoAll from './components/dpo/Dpo-all';

import Ddo from './components/ddo/Ddo';
import DdoDetails from './components/ddo/Ddo-details';
import DdoMy from './components/ddo/Ddo-my';
import DdoAll from './components/ddo/Ddo-all';

import Rfo from './components/rfo/Rfo';
import RfoMy from './components/rfo/Rfo-my';
import RfoAll from './components/rfo/Rfo-all';

import Access from './components/access/Access';
import AccessUsers from './components/access/Access-users';
import AccessRoles from './components/access/Access-roles';
import AccessUsage from './components/access/Access-usage';

Vue.use(VueRouter);

var getDetailsRoute = (page, tab) => {
  return {
    path: ':details', name: page + '.' + tab + '.details', component: global.DscDetails,
    meta: { 
      title: _.upperCase(page) + " Details - Data Catalogue",
      showModal: true,
      permission: _.upperCase(page)
    } 
  }
}

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
      path: ':details/:column', name: 'dsc.details', component: DscDetails,
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
    },
    children: [{ // dsc.details
      path: ':details/:column', name: 'dsc.details', component: DscDetails,
      meta: { 
        title: "DSC Details - Data Catalogue",
        showModal: true,
        permission: "DSC"
      }
    }] 
  }, { // dsc.cdp.cde
    path: '/dsc/cdp/:system/:dspName', name: 'dsc.cdp.cde', component: DscCdpCde, 
    meta: { 
      title: "DSC - Data Catalogue",
      permission: "DSC"
    },
    children: [{ // dsc.details
      path: ':details/:column', name: 'dsc.details', component: DscDetails,
      meta: { 
        title: "DSC Details - Data Catalogue",
        showModal: true,
        permission: "DSC"
      }
    }] 
  },
      // children: [{ 
      //   path: '', name: 'dsc', redirect: { name: 'dsc.my' }
      // }, { //dsc.my
      //   path: 'my', 
      //   name: 'dsc.my', 
      //   component: DscMy, 
      //   meta: { 
      //     title: "DSC - Data Catalogue",
      //     showModal: false,
      //     permission: "DSC"
      //   } 
      // }, { // dsc.my.system
      //   path: 'my/:system', name: 'dsc.my', component: DscMy, 
      //   meta: { 
      //     title: "DSC - Data Catalogue",
      //     showModal: false,
      //     permission: "DSC"
      //   }, 
      //   children: [{ // dsc.my.system.details
      //     path: ':details', name: 'dsc.my.details', component: DscDetails,
      //     meta: { 
      //       title: "DSC Details - Data Catalogue",
      //       showModal: true,
      //       permission: "DSC"
      //     }, 
      //     children: [{ // dsc.my.system.details.col
      //       path: ':column', name: 'dsc.my.details', component: DscDetails,
      //       meta: { 
      //         title: "DSC Details - Data Catalogue",
      //         showModal: true,
      //         permission: "DSC"
      //       }
      //     }] 
      //   }] 
      // }, { // dsc.all
      //   path: 'all', name: 'dsc.all', component: DscAll, 
      //   meta: { 
      //     title: "DSC - Data Catalogue" ,
      //     showModal: false,
      //     permission: "DSC"
      //   } 
      // }, { // dsc.all.system
      //   path: 'all/:system', name: 'dsc.all', component: DscAll, 
      //   meta: { 
      //     title: "DSC - Data Catalogue" ,
      //     showModal: false,
      //     permission: "DSC"
      //   }, 
      //   children: [{ // dsc.all.system.details
      //     path: ':details', name: 'dsc.all.details', component: DscDetails,
      //     meta: { 
      //       title: "DSC Details - Data Catalogue",
      //       showModal: true,
      //       permission: "DSC"
      //     }, 
      //     children: [{ // dsc.all.system.details.col
      //       path: ':column', name: 'dsc.all.details', component: DscDetails,
      //       meta: { 
      //         title: "DSC Details - Data Catalogue",
      //         showModal: true,
      //         permission: "DSC"
      //       }
      //     }]
      //   }]
      // }, { // dsc.interfaces
      //   path: 'interfaces', name: 'dsc.interfaces', component: DscInterfaces, 
      //   meta: { 
      //     title: "DSC - Data Catalogue" ,
      //     showModal: false,
      //     permission: "DSC"
      //   } 
      // }, { // dsc.interfaces.system
      //   path: 'interfaces/:system', name: 'dsc.interfaces', component: DscInterfaces, 
      //   meta: { 
      //     title: "DSC - Data Catalogue" ,
      //     showModal: false,
      //     permission: "DSC"
      //   }, 
      //   children: [{ // dsc.interfaces.system.details
      //     path: ':details', name: 'dsc.interfaces.details', component: DscDetails,
      //     meta: { 
      //       title: "DSC Details - Data Catalogue",
      //       showModal: true,
      //       permission: "DSC"
      //     }, 
      //     children: [{ // dsc.interfaces.system.details.col
      //       path: ':column', name: 'dsc.interfaces.details', component: DscDetails,
      //       meta: { 
      //         title: "DSC Details - Data Catalogue",
      //         showModal: true,
      //         permission: "DSC"
      //       }
      //     }] 
      //   }]
      // }]
    // }, 
  { // dpo
    path: '/dpo', component: Dpo, 
    meta: { 
      title: "DPO - Data Catalogue",
      permission: "DPO"
    }, 
    children: [{ 
      path: '', name: 'dpo', redirect: { name: 'dpo.my' }
    }, { //dpo.my
      path: 'my', 
      name: 'dpo.my', 
      component: DpoMy, 
      meta: { 
        title: "DPO - Data Catalogue",
        showModal: false,
        permission: "DPO"
      } 
    }, { // dpo.my.system
      path: 'my/:system', name: 'dpo.my', component: DpoMy, 
      meta: { 
        title: "DPO - Data Catalogue",
        showModal: false,
        permission: "DPO"
      }, 
      children: [{ // dpo.my.system.details
        path: ':details', name: 'dpo.my.details', component: DpoDetails,
        meta: { 
          title: "DPO Details - Data Catalogue",
          showModal: true,
          permission: "DPO"
        } 
      }] 
    }, { // dpo.all
      path: 'all', name: 'dpo.all', component: DpoAll, 
      meta: { 
        title: "DPO - Data Catalogue" ,
        showModal: false,
        permission: "DPO"
      } 
    }, { 
      path: 'all/:system', name: 'dpo.all', component: DpoAll, 
      meta: { 
        title: "DPO - Data Catalogue" ,
        showModal: false,
        permission: "DPO"
      }, 
      children: [{ // dpo.all.system.details
        path: ':details', name: 'dpo.all.details', component: DpoDetails,
        meta: { 
          title: "DPO Details - Data Catalogue",
          showModal: true,
          permission: "DPO"
        } 
      }]
    }]
  }, { // ddo
    path: '/ddo', component: Ddo, 
    meta: { 
      title: "DDO - Data Catalogue",
      permission: "DDO"
    }, 
    children: [{ 
      path: '', name: 'ddo', redirect: { name: 'ddo.my' }
    }, { //ddo.my
      path: 'my', 
      name: 'ddo.my', 
      component: DdoMy, 
      meta: { 
        title: "DDO - Data Catalogue",
        showModal: false,
        permission: "DDO"
      } 
    }, { // ddo.my.system
      path: 'my/:system', name: 'ddo.my', component: DdoMy, 
      meta: { 
        title: "DDO - Data Catalogue",
        showModal: false,
        permission: "DDO"
      }, 
      children: [{ // ddo.my.system.details
        path: ':details', name: 'ddo.my.details', component: DdoDetails,
        meta: { 
          title: "DDO Details - Data Catalogue",
          showModal: true,
          permission: "DDO"
        } 
      }] 
    }, { // ddo.all
      path: 'all', name: 'ddo.all', component: DdoAll, 
      meta: { 
        title: "DDO - Data Catalogue" ,
        showModal: false,
        permission: "DDO"
      } 
    }, { 
      path: 'all/:system', name: 'ddo.all', component: DdoAll, 
      meta: { 
        title: "DDO - Data Catalogue" ,
        showModal: false,
        permission: "DDO"
      }, 
      children: [{ // ddo.all.system.details
        path: ':details', name: 'ddo.all.details', component: DdoDetails,
        meta: { 
          title: "DDO Details - Data Catalogue",
          showModal: true,
          permission: "DDO"
        } 
      }]
    }]
  }, { // rfo
    path: '/rfo', component: Rfo, 
    meta: { 
      title: "RFO - Data Catalogue",
      permission: "RFO"
    }, 
    children: [{ 
      path: '', name: 'rfo', redirect: { name: 'rfo.my' }
    }, { //rfo.my
      path: 'my', 
      name: 'rfo.my', 
      component: RfoMy, 
      meta: { 
        title: "RFO - Data Catalogue",
        showModal: false,
        permission: "RFO"
      } 
    }, { // rfo.my.system
      path: 'my/:system', name: 'rfo.my', component: RfoMy, 
      meta: { 
        title: "RFO - Data Catalogue",
        showModal: false,
        permission: "RFO"
      },
    }, { // rfo.all
      path: 'all', name: 'rfo.all', component: RfoAll, 
      meta: { 
        title: "RFO - Data Catalogue" ,
        showModal: false,
        permission: "RFO"
      } 
    }, { 
      path: 'all/:system', name: 'rfo.all', component: RfoAll, 
      meta: { 
        title: "RFO - Data Catalogue" ,
        showModal: false,
        permission: "RFO"
      },
    }]
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
    }, { //access.users
      path: 'roles', name: 'access.roles', component: AccessRoles, 
      meta: { 
        title: "Roles - Data Catalogue",
        permission: "Admin"
      } 
    }, { //access.users
      path: 'usage', name: 'access.usage', component: AccessUsage, 
      meta: { 
        title: "Usage Detail - Data Catalogue",
        permission: "Admin"
      } 
    }]
  }, {
    path: '/login', name: 'login', component: Login, 
  },
  { path: '*', redirect: '/' }]
});

router.beforeEach((to, from, next) => {
  document.title = (to.meta.title || '')
  next()
});

var saveUsage = function(isAuth, to, from, next){
  // console.log(isAuth, to, from, next);
  
}

router.beforeEach((to, from, next) => {
  // redirect to login page if not logged in and trying to access a restricted page
  const publicPages = ['/login'];
  const authRequired = !publicPages.includes(to.path);
  const loggedIn = localStorage.getItem('user');
  
  if (authRequired && !loggedIn) {
    saveUsage(false, to, from, next);
    return next('/login');
  }

  saveUsage(true, to, from, next)
  next();
})

export default router