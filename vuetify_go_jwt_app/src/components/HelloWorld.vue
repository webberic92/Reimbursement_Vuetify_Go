<template>
  <v-container>
    <v-row class="text-center">
      <v-col cols="12">
        <v-img
          :src="require('../assets/cpbLogo.png')"
          class="my-3"
          contain
          height="200"
        />
      </v-col>

      <v-col class="mb-4">
        <h1 class="display-2 font-weight-bold mb-3">
          Welcome to the CPB Reimbursment Application.
        </h1>

        <p>
          <br />

          This app will have two types of users and administrators.
          <br />
          Users will be able to Submit Reimbursment requests.
          <br />
          Admins will be able to apporve Reimbursment requests.
        </p>

        <br>
        <br>
        <p>Please
        <router-link to="register">Register</router-link>
        and
        <router-link to="login">Login</router-link>
        to view functionality.</p>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
export default {
  name: "HelloWorld",

  data: () => ({}),
};
</script>
