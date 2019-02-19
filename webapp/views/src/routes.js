import Vue from 'vue';
import VueRouter from 'vue-router';

import Home from './components/Home';

import Login from './components/Login';

import Dsc from './components/dsc/Dsc';
import DscDetails from './components/dsc/Dsc-details';
import DscMy from './components/dsc/Dsc-my';
import DscAll from './components/dsc/Dsc-all';
import DscInterfaces from './components/dsc/Dsc-interfaces';

import Dpo from './components/dpo/Dpo';
import DpoDetails from './components/dpo/Dpo-details';
import DpoMy from './components/dpo/Dpo-my';
import DpoAll from './components/dpo/Dpo-all';

import Ddo from './components/Ddo';
import Rfo from './components/Rfo';

import Access from './components/access/Access';
import AccessUsers from './components/access/Access-users';
import AccessRoles from './components/access/Access-roles';

Vue.use(VueRouter);

const router = new VueRouter({
  mode: 'history',
  routes: [{ // home
    path: '/', name: 'landingpage', component: Home, 
    meta: { 
      title: "Home - Data Catalogue" 
    },
    children: [{ // dsc
      path: '/dsc', component: Dsc, 
      meta: { 
        title: "DSC - Data Catalogue",
        permission: "DSC"
      }, 
      children: [{ 
        path: '', name: 'dsc', redirect: { name: 'dsc.my' }
      }, { //dsc.my
        path: 'my', 
        name: 'dsc.my', 
        component: DscMy, 
        meta: { 
          title: "DSC - Data Catalogue",
          showModal: false,
          permission: "DSC"
        } 
      }, { // dsc.my.system
        path: 'my/:system', name: 'dsc.my', component: DscMy, 
        meta: { 
          title: "DSC - Data Catalogue",
          showModal: false,
          permission: "DSC"
        }, 
        children: [{ // dsc.my.system.details
          path: ':details', name: 'dsc.my.details', component: DscDetails,
          meta: { 
            title: "DSC Details - Data Catalogue",
            showModal: true,
            permission: "DSC"
          } 
        }] 
      }, { // dsc.all
        path: 'all', name: 'dsc.all', component: DscAll, 
        meta: { 
          title: "DSC - Data Catalogue" ,
          showModal: false,
          permission: "DSC"
        } 
      }, { 
        path: 'all/:system', name: 'dsc.all', component: DscAll, 
        meta: { 
          title: "DSC - Data Catalogue" ,
          showModal: false,
          permission: "DSC"
        }, 
        children: [{ // dsc.all.system.details
          path: ':details', name: 'dsc.all.details', component: DscDetails,
          meta: { 
            title: "DSC Details - Data Catalogue",
            showModal: true,
            permission: "DSC"
          } 
        }]
      }, { // dsc.interfaces
        path: 'interfaces', name: 'dsc.interfaces', component: DscInterfaces, 
        meta: { 
          title: "DSC - Data Catalogue" ,
          showModal: false,
          permission: "DSC"
        } 
      }, { 
        path: 'interfaces/:system', name: 'dsc.interfaces', component: DscInterfaces, 
        meta: { 
          title: "DSC - Data Catalogue" ,
          showModal: false,
          permission: "DSC"
        }, 
        children: [{ // dsc.interfaces.system.details
          path: ':details', name: 'dsc.interfaces.details', component: DscDetails,
          meta: { 
            title: "DSC Details - Data Catalogue",
            showModal: true,
            permission: "DSC"
          } 
        }]
      }]
    }, { // dpo
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
    }, { 
      path: '/ddo', component: Ddo, 
      meta: { 
        title: "DDO - Data Catalogue" ,
        permission: "DDO"
      } 
    }, { 
      path: '/rfo', component: Rfo, 
      meta: { 
        title: "RFO - Data Catalogue" ,
        permission: "RFO"
      } 
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
      }]
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

router.beforeEach((to, from, next) => {
  // redirect to login page if not logged in and trying to access a restricted page
  const publicPages = ['/login'];
  const authRequired = !publicPages.includes(to.path);
  const loggedIn = localStorage.getItem('user');
  
  if (authRequired && !loggedIn) {
    return next('/login');
  }

  next();
})

export default router