<template>
  <p align="center">
    <strong>
      Welcome {{ this.firstName }}
      <br />
      <br />
      You are currently logged in as a
      {{ this.userType }} user<br />
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
  </p>
</template>

<script>
import Axios from "axios";

export default {
       created() {  
      console.log("get user button hit");
      Axios.get("http://localhost:8000/api/user", {
        withCredentials: true,
        headers: {
          "Access-Control-Allow-Origin": "*",
          "Content-Type": "application/json",
        },
      })
        .then(async (response) => {
          if (!response.ok) {
            this.firstName = response.data.firstName;

            if (response.data.userType == "R") {
              this.userType = "Regular";
            }

            console.log(response.data);
          }
        })
        .catch((error) => {
          console.log(error.response);
           this.$router.push("/");
        });
  },
  data() {
    return {
      firstName: "",
      userType: "",
    };
  },
  methods: {
  },
};
</script>
