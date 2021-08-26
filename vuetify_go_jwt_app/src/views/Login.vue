<template>
  <v-container class="grey lighten-5" fill-height fluid>
    <v-card class="pa-md-4 mx-lg-auto" color="white" width="600px">
     
      <!-- Messages -->
      <p v-if="errorMessage" align="center" style="color:red">
        <strong>
          {{ errorMessage }}
        </strong>
      </p>
      <p v-if="successMessage" align="center" style="color:green">
        <strong>{{ successMessage }}</strong>
      </p>

      <p v-if="routingMessage" align="center">
        <strong> {{ routingMessage }} </strong>
      </p>
      <v-row>
        <v-col cols="12" sm="4"> </v-col>
      </v-row>

      <validation-observer ref="observer" v-slot="{ invalid }">
        <form @submit.prevent="submit">


          <!-- Email -->

          <validation-provider
            v-slot="{ errors }"
            name="email"
            rules="required|email"
          >
            <v-text-field
              v-model="form.email"
              :error-messages="errors"
              label="E-mail"
              required
            ></v-text-field>
          </validation-provider>

          <!-- Password -->

          <validation-provider
            v-slot="{ errors }"
            name="password"
            rules="required"
          >
            <v-text-field
              type="password"
              v-model="form.password"
              :error-messages="errors"
              label="Password"
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
</template>

<script>
import { required, email } from "vee-validate/dist/rules";
import {
  extend,
  ValidationObserver,
  ValidationProvider,
  setInteractionMode,
} from "vee-validate";
import axios from "axios";

setInteractionMode("eager");

extend("required", {
  ...required,
  message: "{_field_} can not be empty",
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
  data() {
    return {
      errorMessage: "",
      successMessage: "",
      routingMessage: "",
      form: {
        email: "",
        password: "",
      },

      typeOfUser: ["Regular User", "Administrator User"],
    };
  },
  methods: {
    submit() {
      this.$refs.observer.validate();
      console.log("submit button hit");
      console.log(JSON.stringify(this.form));
axios.post('http://localhost:8000/api/login', this.form, {
withCredentials: true
})
        .then(async (response) => {
          if (!response.ok) {
            console.log(response.data.message);
            this.errorMessage = "";
            this.successMessage = response.data.message;
            this.routingMessage = "Now routing you to login page";
            setTimeout(() => {
              this.$router.push("/home");
            }, 3000);
          }
        })
        .catch((error) => {
          console.log(error.response);
          this.errorMessage = error.response;
          this.successMessage - "";

        });
    },
    clear() {
      this.email = "";
      this.password = "";
      this.$refs.observer.reset();
      this.errorMessage = "";
      this.successMessage = "";
      this.routingMessage = "";
    },
  },
};
</script>
