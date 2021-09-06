<template>
  <!-- ADMIN STUFF -->
  <div v-if="this.$store.state.isAdmin">
    ADMIN STUFF HERE

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
  <div align="center" v-else-if="!this.$store.state.isAdmin">
    REGULAR USER STUFF HERE

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
              rules="required|max:200"
            >
              <v-text-field
                v-model="form.description"
                :counter="200"
                :error-messages="errors"
                label="Description"
                required
              ></v-text-field>
            </validation-provider>

            <!-- Amount   -->

            <validation-provider
              v-slot="{ errors }"
              name="Reimbursment Amount"
              rules="required|digits_between:1,8"
            >
              <v-text-field
                v-model="form.amount"
                :error-messages="errors"
                label="Reimbursment Amount"
                required
              ></v-text-field>
            </validation-provider>

            <br />
            <br />

            <v-btn class="mr-4" type="submit" :disabled="invalid">
              submit
            </v-btn>
            <v-btn @click="clear">
              clear
            </v-btn>
          </form>
        </validation-observer>
        <br />
        <v-btn>
          Get Current Requests.
        </v-btn>
      </v-card>
    </v-container>
  </div>
</template>

<script>
import Axios from "axios";
import axios from "axios";

import store from "..//store/index";
import { required, digits,  max, regex } from "vee-validate/dist/rules";
import {
  extend,
  ValidationObserver,
  ValidationProvider,
  setInteractionMode,
} from "vee-validate";
import { min, numeric } from "vee-validate/dist/rules";
import { validate } from "vee-validate";

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

extend("min", min);
extend("max", max);
extend("numeric", numeric);

extend("digits_between", {
  async validate(value, { min, max }) {
    const res = await validate(value, `numeric|min:${min}|max:${max}`);
    return res.valid;
  },
  params: ["min", "max"],
  message: "The {_field_} must be between {min} and {max} digits",
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
          this.form.userId = response.data.id;
          store.commit("setUserId", this.form.userId)
          console.log(" this is the user id : ", store.getters.userId)
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
      form: {
        userId: "",
        title: "",
        description: "",
        amount: "",
      },
      firstName: "",
      userType: "",
    };
  },
  methods: {
    clear() {
      console.log(JSON.stringify(this.form));
    },
     submit() { 
       this.form.userId = this.form.userId.toString()
      console.log(JSON.stringify(this.form));
      axios({
        method: "post",
        url: "http://localhost:8000/api/createReimbursment",
        data: this.form,
        withCredentials: true})
        .then(async (response) => {
          if (!response.ok) {
            console.log(response);
            this.errorMessage = "";
            this.successMessage = response.data.message;
          }
        })
        .catch((error) => {
          console.log(error.response.data.message);
          this.errorMessage = error.response.data.message;
          this.successMessage = "";
        });
    }
  },
};
</script>
