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
 
 <v-toolbar-title style="cursor: pointer" @click="$router.push('/')" >CPB Reimbusment System</v-toolbar-title>
        
  <v-spacer></v-spacer>

      </div>

      <v-spacer></v-spacer>

    <v-btn text to="register">Register</v-btn>
    <v-btn text to="login">Login</v-btn>
    </v-app-bar>

    <v-main>
      <router-view/>
    </v-main>
  </v-app>
</template>

<script>

export default {
  name: 'App',

  data: () => ({
    //
  }),
};
</script>
