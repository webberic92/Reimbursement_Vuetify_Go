<template>
  <v-app>
    <v-app-bar
      app
      color="primary"
      dark
    >

    
      <div class="d-flex align-center">
        <router-link to="/">
 <img to="/" class="mr-3" :src="require('@/assets/cpbLogoDark.png')" height="40"/>
 </router-link>
 
 <v-toolbar-title style="cursor: pointer" @click="$router.push('/').catch(err => {})" >CPB Reimbusment System</v-toolbar-title>
        
  <v-spacer></v-spacer>

      </div>

      <v-spacer></v-spacer>

    <v-btn text  v-if="!isLoggedIn" to="register">Register</v-btn>
    <v-btn text  v-if="!isLoggedIn" to="login">Login</v-btn>

    <v-btn text v-if="isLoggedIn" @click='logout'>Logout</v-btn>
    </v-app-bar>

    <v-main>
      <router-view/>
    </v-main>
  </v-app>
</template>

<script>
import axios from "axios";

export default {

  name: 'App',
  computed : {
      isLoggedIn : function(){ return this.$store.getters.isLoggedIn}
    },
    methods : {

      logout({commit}) {
      console.log("Logout button hit");
      axios({
        method: "post",
        url: "http://localhost:8000/api/logout",
      })
        .then(async (response) => {
          if (!response.ok) {
            console.log(response.data.message);
            localStorage.removeItem('token')
            commit('logout')
            setTimeout(() => {
              this.$router.push("/");
            }, 3000);
          }
        })
        .catch((error) => {
          console.log(error.response.data.message);
          // this.errorMessage = error.response.data.message;
          // this.successMessage - "";

        });
    },
  data: () => ({
    //
  }),
    }}
</script>
