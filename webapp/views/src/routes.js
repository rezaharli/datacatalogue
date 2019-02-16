import Vue from 'vue';
import VueRouter from 'vue-router';

import Home from './components/Home';

import Login from './components/Login';

import Dsc from './components/dsc/Dsc';
import DscDetails from './components/dsc/Dsc-details';
import DscMy from './components/dsc/Dsc-my';
import DscAll from './components/dsc/Dsc-all';
import DscInterfaces from './components/dsc/Dsc-interfaces';

import Dpo from './components/Dpo';
import Ddo from './components/Ddo';
import Rfo from './components/Rfo';

import Access from './components/access/Access';
import AccessUsers from './components/access/Access-users';

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
        title: "DSC - Data Catalogue" 
      }, 
      children: [{ 
        path: '', name: 'dsc', redirect: { name: 'dsc.my' }
      }, { //dsc.my
        path: 'my', 
        name: 'dsc.my', 
        component: DscMy, 
        meta: { 
          title: "DSC - Data Catalogue",
          showModal: false
        } 
      }, { // dsc.my.system
        path: 'my/:system', name: 'dsc.my', component: DscMy, 
        meta: { 
          title: "DSC - Data Catalogue",
          showModal: false
        }, 
        children: [{ // dsc.my.system.details
          path: 'details', name: 'dsc.my.details', component: DscDetails, 
          meta: { 
            title: "DSC Details - Data Catalogue",
            showModal: true
          } 
        }] 
      }, { // dsc.all
        path: 'all', name: 'dsc.all', component: DscAll, 
        meta: { 
          title: "DSC - Data Catalogue" 
        } 
      }, { 
        path: 'all/:system', name: 'dsc.all', component: DscAll, 
        meta: { 
          title: "DSC - Data Catalogue" 
        }, 
        children: [{ 
          path: 'details', name: 'dsc.all.details', component: DscDetails, 
          meta: { 
            title: "DSC Details - Data Catalogue" 
          }
        }] 
      }, {
        path: 'interfaces', name: 'dsc.interfaces', component: DscInterfaces, 
        meta: { 
          title: "DSC - Data Catalogue" 
        } 
      }, { 
        path: 'interfaces/:system', name: 'dsc.interfaces', component: DscInterfaces, 
        meta: { 
          title: "DSC - Data Catalogue" 
        }, 
        children: [{ 
          path: 'details', name: 'dsc.interfaces.details', component: DscDetails, 
          meta: { 
            title: "DSC Details - Data Catalogue" 
          } 
        }] 
      }]
    }, { 
      path: '/dpo', component: Dpo, 
      meta: { 
        title: "DPO - Data Catalogue" 
      } 
    }, { 
      path: '/ddo', component: Ddo, 
      meta: { 
        title: "DDO - Data Catalogue" 
      } 
    }, { 
      path: '/rfo', component: Rfo, 
      meta: { 
        title: "RFO - Data Catalogue" 
      } 
    }, { // access
      path: '/access', component: Access, 
      meta: { 
        title: "Access - Data Catalogue" 
      }, 
      children: [{ 
        path: '', name: 'access', redirect: { name: 'access.users' }
      }, { //access.users
        path: 'users', name: 'access.users', component: AccessUsers, 
        meta: { 
          title: "Users - Data Catalogue",
          showModal: false
        } 
      }, { // access.users
        path: 'users/:id', name: 'access.users', component: AccessUsers, 
        meta: { 
          title: "DSC - Data Catalogue",
          showModal: false
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