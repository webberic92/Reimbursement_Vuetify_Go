<template>




<!-- ADMIN STUFF -->
  <div v-if="this.$store.state.isAdmin">ADMIN STUFF HERE

 <p align="center">
    <strong>
      Welcome {{ this.firstName }}
      <br />
      <br />
      You are currently logged in as a
      {{ this.userType }} <br />
    </strong>
    <br />
  </p>


  </div>

<!-- Regular User Stuff -->
  <div align="center" v-else-if="!this.$store.state.isAdmin">REGULAR USER STUFF HERE





   <v-container class="grey lighten-5" fill-height fluid>





    <v-card class="pa-md-4 mx-lg-auto" color="white" width="auto">
     <p>
    <strong>
      Welcome {{ this.firstName }}
      <br />
      <br /> 
      You are currently logged in as a
      {{ this.userType }} User <br />
    </strong>
    <br />
  </p>

      <validation-observer ref="observer" v-slot="{ invalid }">
        <form @submit.prevent="submit">
          <!-- Title   -->
          <validation-provider
            v-slot="{ errors }"
            name="Title"
            rules="required|max:20"
          >
            <v-text-field
              v-model="form.title"
              :counter="20"
              :error-messages="errors"
              label="Title"
              required
            ></v-text-field>
          </validation-provider>

          <!-- Description   -->
          <validation-provider
            v-slot="{ errors }"
            name="Description"
            rules="required|max:20"
          >
            <v-text-field
              v-model="form.description"
              :counter="20"
              :error-messages="errors"
              label="Description"
              required
            ></v-text-field>
          </validation-provider>

                    <!-- Amount   -->
          <validation-provider
            v-slot="{ errors }"
            name="Amount"
            rules="required|max:20"
          >
            <v-text-field
              v-model="form.amount"
              :counter="20"
              :error-messages="errors"
              label="Amount"
              required
            ></v-text-field>
          </validation-provider>

          <v-btn class="mr-4" type="submit" :disabled="invalid">
            submit
          </v-btn>
          <v-btn @click="clear">
            clear
          </v-btn>
        </form>
      </validation-observer>
    </v-card>
  </v-container>


  </div>

</template>

<script>
import Axios from "axios";
import store from "..//store/index";
import { required, digits, email, max, regex } from "vee-validate/dist/rules";
import {
  extend,
  ValidationObserver,
  ValidationProvider,
  setInteractionMode,
} from "vee-validate";

setInteractionMode("eager");

extend("digits", {
  ...digits,
  message: "{_field_} needs to be {length} digits. ({_value_})",
});

extend("required", {
  ...required,
  message: "{_field_} can not be empty",
});

extend("max", {
  ...max,
  message: "{_field_} may not be greater than {length} characters",
});

extend("regex", {
  ...regex,
  message: "{_field_} {_value_} does not match {regex}",
});

extend("email", {
  ...email,
  message: "Email must be valid",
});



export default {
  components: {
      ValidationProvider,
    ValidationObserver,
  },
  created() {
    Axios.get("http://localhost:8000/api/user", {
      withCredentials: true,
      headers: {
        "Access-Control-Allow-Origin": "*",
        "Content-Type": "application/json",
      },
    })
      .then(async (response) => {
        if (!response.ok) {
          this.id = response.data.id
          this.firstName = response.data.firstName;
          if (response.data.userType == "R") {
            this.userType = "Regular";
          } else if (response.data.userType == "A") {
            this.userType = "Admin";
            store.commit("makeUserAdmin");
          }
        }
      })
      .catch((error) => {
        console.log(error.response);
        this.$router.push("/");
      });
  },
  data() {
    return {
      form:{
            title: "",
            description: "",
            amount: null
      },
      id: null,
      firstName: "",
      userType:""
    };
  },
  methods: {
 clear() {
      this.firstName = "";
    },


  },
};
</script>
