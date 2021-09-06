import Vue from 'vue';
import Vuex from 'vuex';
Vue.use(Vuex);


const store = new Vuex.Store({
  state: {
    isLoggedIn : false,
    isAdmin : false,
    userId : null
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
      state.isAdmin = null

    },
    setUserId (state, loggedInUserId){
      state.userId = loggedInUserId
    }

  },
  getters: {
    loggedIn : state =>  state.isLoggedIn,
    admin : state => state.isAdmin,
    userId : state => state.userId
  }
});

export default store;