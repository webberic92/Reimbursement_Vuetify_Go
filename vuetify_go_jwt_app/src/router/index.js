import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'
import Register from '../views/Register.vue'
import Login from '../views/Login.vue'
import Dashboard from '../views/Dashboard.vue'
import store from "..//store/index";


Vue.use(VueRouter)


function useMustBeLoggedIn(to, from, next)
{
  if(store.state.isLoggedIn){
    next(); // allow to enter route
   } 
   else
   {
    next('/'); // go to '/';
   }
  }

function useAlreadyLoggedIn(to, from, next)
{
  if(!store.state.isLoggedIn){

    next();
   }
  }

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home,
    beforeEnter : useAlreadyLoggedIn,

  },
  {
    path: '/login',
    name: 'Login',
    component: Login,
    beforeEnter : useAlreadyLoggedIn
  },
  {
    path: '/register',
    name: 'Register',
    component: Register,
    beforeEnter : useAlreadyLoggedIn,

  },
  {
    path: '/dashboard',
    name: 'Dashboard',
    component: Dashboard,
    beforeEnter : useMustBeLoggedIn,
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

router.beforeEach((to, from, next) => {
  if(to.matched.some(record => record.meta.requiresAuth)) {
    if (store.getters.isLoggedIn) {
      next()
      return
    }
    next('/')
  } else {
    next()
  }
})

export default router
