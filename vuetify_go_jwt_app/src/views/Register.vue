<template>

  <v-container class="grey lighten-5" fill-height fluid
    >
    <v-card class="pa-md-4 mx-lg-auto" color="white" width="auto">
      <v-row>
        <v-col cols="12" sm="4"> </v-col>
      </v-row>
<p>TODO: ADD PASSWORD LOGIC
</p>
      <validation-observer ref="observer" v-slot="{ invalid }">
        <form @submit.prevent="submit">
          <!-- First Name   -->
          <validation-provider
            v-slot="{ errors }"
            name="Firt Name"
            rules="required|max:20"
          >
            <v-text-field
              v-model="form.firstName"
              :counter="20"
              :error-messages="errors"
              label="First Name"
              required
            ></v-text-field>
          </validation-provider>

          <!-- Last Name -->
          <validation-provider
            v-slot="{ errors }"
            name="Last Name"
            rules="required|max:20"
          >
            <v-text-field
              v-model="form.lastName"
              :counter="20"
              :error-messages="errors"
              label="Last Name"
              required
            ></v-text-field>
          </validation-provider>

          <!-- Phone Number -->

          <validation-provider
            v-slot="{ errors }"
            name="phoneNumber"
            :rules="{
              required: true,
              digits: 10,
              regex: '',
            }"
          >
            <v-text-field
              v-model="form.phoneNumber"
              :counter="10"
              :error-messages="errors"
              label="Phone Number"
              required
            ></v-text-field>
          </validation-provider>

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

          <!-- Type of -->

          <validation-provider
            v-slot="{ errors }"
            name="select"
            rules="required"
          >
            <v-select
              v-model="form.userType"
              :items="typeOfUser"
              :error-messages="errors"
              label="Select Type of User You Want to Register As."
              data-vv-name="select"
              required
            ></v-select>
          </validation-provider>

          <!-- SOU -->
          <validation-provider
            v-slot="{ errors }"
            rules="required"
            name="Statement of Understanding"
          >Statement of Understanding
            <v-checkbox
              v-model="form.sou"
              :error-messages="errors"
              value="true"
              label="By Checking this box you are agreeing to not use this app nefariously or out of context."
              type="sou"
              required
            ></v-checkbox>
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
import { required, digits, email, max, regex } from "vee-validate/dist/rules";
import {
  extend,
  ValidationObserver,
  ValidationProvider,
  setInteractionMode,
} from "vee-validate";
import axios from "axios";

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
  data() {
    return {
      form: {
        firstName: "",
        lastName: "",
        phoneNumber: "",
        email: "",
        userType: null,
        sou: null,
      },
      typeOfUser: ["Regular User", "Administrator User"],
    };
  },
  methods: {
    submit() {
      this.$refs.observer.validate();
      if (this.form.userType == "Regular User"){
        this.form.userType = "R"
      } 
      if (this.form.userType == "Administrator User"){
        this.form.userType = "A"
      } 
      console.log("submit button hit");
      console.log(JSON.stringify(this.form));
      axios({
        method: "post",
        url: "http://localhost:8000/api/register",
        data: this.form,
      }).then(function(response) {
        console.log(response);
      });
    },
     clear() {
      this.firstName = "";
      this.lastName = "";
      this.phoneNumber = "";
      this.email = "";
      this.userType = null;
      this.sou = null;
      this.$refs.observer.reset();
    },
  },
};
</script>
