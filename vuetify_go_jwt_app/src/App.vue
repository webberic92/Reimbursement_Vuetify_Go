<template>
  <v-app>
    <v-app-bar app color="primary" dark>
      <div class="d-flex align-center">
        <router-link to="/">
          <img
            to="/"
            class="mr-3"
            :src="require('@/assets/cpbLogoDark.png')"
            height="40"
          />
        </router-link>

        <v-toolbar-title
          style="cursor: pointer"
          @click="$router.push('/').catch((err) => {})"
          >CPB Reimbusment System</v-toolbar-title
        >

        <v-spacer></v-spacer>
      </div>

      <v-spacer></v-spacer>

      <v-btn text  to="register">Register</v-btn>
      <v-btn text  to="login">Login</v-btn>
      <v-btn text  to="home">LoggedInUser</v-btn>

      <v-btn text  @click="logout">Logout</v-btn>
    </v-app-bar>

    <v-main>
      <router-view />
    </v-main>
  </v-app>
</template>

<script>
import axios from "axios";

export default {
  name: "App",
   methods: {
    logout() {
      axios({
        method: "post",
        url: "http://localhost:8000/api/logout",
        headers: {'Content-Type': 'application/json'},
        credentials: 'include'
      })
        .then(async (response) => {
          if (!response.ok) {
            console.log(response.data.message);

          }
        })
        .catch((error) => {
          console.log(error.response.data.message);


        });
    },

  data: () => ({
    //
  }),
   }}
</script>
