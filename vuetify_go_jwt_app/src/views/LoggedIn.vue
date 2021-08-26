<template>
  <p align="center">
    <strong>
      Logged in VUE
    </strong>
    <br />
    <strong>
      TODO: Introduce VUEX guarding
    </strong>
    <br />
    <strong>
      Differentiate between Admin users or Regular users.
    </strong>
    <br />
    <v-btn text  @click="getUser">Get Current User</v-btn>
  </p>
</template>

<script>

import Axios from "axios";



export default {

  methods: {
    getUser() {
      console.log("get user button hit");
      const token = JSON.parse(sessionStorage.getItem('jwt'));
      console.log(token)
Axios.get('http://localhost:8000/api/user',
 {
        withCredentials: true,
        headers: {'Access-Control-Allow-Origin': '*', 'Content-Type': 'application/json'
    }}).then(async (response) => {



          if (!response.ok) {
             

            console.log(response.data.message);
          }
        })
        .catch((error) => {
          console.log(error.response);

        });
    },
  },
};
</script>
