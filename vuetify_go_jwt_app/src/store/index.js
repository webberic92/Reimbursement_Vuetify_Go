import Vue from 'vue';
import Vuex from 'vuex';
Vue.use(Vuex);


const store = new Vuex.Store({
  state: {
    isLoggedIn : false,
    isAdmin : false
  },
  mutations: {
    logUserIn(state){
      state.isLoggedIn = true
    },
    makeUserAdmin(state){
      state.isAdmin = true
    },
    logUserOut(state){
      state.isLoggedIn = false
      state.isAdmin = false
    }

  },
  getters: {
    loggedIn : state =>  state.isLoggedIn,
    admin : state => state.isAdmin
  }
});

export default store;