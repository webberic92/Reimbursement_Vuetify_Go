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

      <!-- Not logged in menu -->
      <div v-if="!this.$store.state.isLoggedIn">
        <v-btn text to="register">Register</v-btn>
        <v-btn text to="login">Login</v-btn>
      </div>

      <!-- IS logged in menu -->
      <div v-if="this.$store.state.isLoggedIn">
        <v-btn text to="Dashboard">Dashboard</v-btn>

        <v-btn text to="History">History</v-btn>

        <v-btn text @click="logout">Logout</v-btn>
      </div>
    </v-app-bar>

    <v-main>
      <router-view />
    </v-main>
  </v-app>
</template>

<script>
import axios from "axios";
import store from "../src/store/index";

export default {
  name: "App",
  methods: {
    logout() {
      axios("http://localhost:8000/api/logout", {
        method: "get",
        withCredentials: true,
      })
        .then(async (response) => {
          if (!response.ok) {
            store.commit("logUserOut");
            this.$router.push("/");
          }
        })
        .catch((error) => {
          console.log(error.response);
        });
    },
  },
};
</script>
