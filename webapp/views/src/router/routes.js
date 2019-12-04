import Vue from 'vue';
import VueRouter from 'vue-router';

import Login from '../pages/Login';
import Crypto from '../pages/Crypto';

import EdmpDd from '../pages/dsc/edmp/Edmp-dd';
import EdmpDdTechnical from '../pages/dsc/edmp/Edmp-dd-technical';
import EdmpDdBusiness from '../pages/dsc/edmp/Edmp-dd-business';

import Access from '../pages/access/Access';
import AccessUsers from '../pages/access/Access-users';
import AccessRoles from '../pages/access/Access-roles';
import AccessUsage from '../pages/access/Access-usage';

import {store} from '../_store/index'

Vue.use(VueRouter);

const router = new VueRouter({
  mode: 'history',
  routes: [{ // home
    path: '/', name: 'landingpage', redirect: { name: 'dsc.edmp.dd.technical' }
  }, { // dsc.edmp.dd
    path: '/dsc/edmp/dd', name: 'dsc.edmp.dd', component: EdmpDd, 
    meta: { 
      title: "EDMp - Data Catalogue",
      permission: "DSC"
    },
    children: [
      { 
        path: '', name: 'dsc.edmp.dd.default', redirect: { name: 'dsc.edmp.dd.technical' }
      }, { // dsc.edmp.dd.technical
        path: '/dsc/edmp/dd/technical-metadata', name: 'dsc.edmp.dd.technical', component: EdmpDdTechnical, 
        meta: { 
          title: "EDMp - Data Catalogue",
          permission: "DSC"
        },
      }, { // dsc.edmp.dd.business
        path: '/dsc/edmp/dd/business-metadata', name: 'dsc.edmp.dd.business', component: EdmpDdBusiness, 
        meta: { 
          title: "EDMp - Data Catalogue",
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