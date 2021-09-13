<template>
  <!-- BEGIN ADMIN STUFF -->
  <div text-align: center v-if="this.$store.state.isAdmin">
    ADMIN STUFF HERE

    <v-container class="grey lighten-5" fill-height fluid>
      <v-card class="pa-md-4 mx-lg-auto" color="white" width="auto">
        <p text-align: center>
          <strong>
            Welcome {{ this.firstName }}
            <br />
            <br />
            You are currently logged in as a
            {{ this.userType }} <br />
          </strong>
          <br />
        </p>
        <p v-if="ApproveOrDenyMessage" text-align: center style="color:green">
          <strong>{{ ApproveOrDenyMessage }}</strong>
        </p>
        <!-- Get Open Reimbursements -->
        <p v-if="loadOfOpenRequest" text-align: center style="color:green">



        
          <v-data-table :headers="adminCurrentHeaders" :items="adminCurrent">
            <template v-slot:item="row">
              <tr>
                <td>{{ row.item.RequestId }}</td>
                <td>{{ row.item.userID }}</td>
                <td>{{ row.item.title }}</td>
                <td>{{ row.item.description }}</td>
                <td>{{ row.item.amount }}</td>
                <td>
                  <v-btn
                    color="success"
                    @click="approveOrDeny('A', row.item.RequestId)"
                    >Approve</v-btn
                  >
                </td>
                <td>
                  <v-btn
                    color="error"
                    @click="approveOrDeny('D', row.item.RequestId)"
                    >Deny</v-btn
                  >
                </td>
              </tr>
            </template>
          </v-data-table>
        </p>
        <!-- Get Open Reimbursements -->

        <v-btn @click="getAllOpenRequests">
          Get All Open Requests.
        </v-btn>
        <br />
        <br />
        <v-btn @click="clearAdmin">
          Clear
        </v-btn>
      </v-card>
    </v-container>
  </div>
  <!-- End ADMIN STUFF -->

  <!-- Regular User Stuff -->
  <div text-align: center v-else-if="!this.$store.state.isAdmin">
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
          <form @submit.prevent="submit" id="myForm">
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
        <v-btn @click="getOpenRequests">
          Get Current Requests.
        </v-btn>
        <br />
        <br />
        <!-- Messages -->
        <p v-if="errorMessage" text-align: center style="color:red">
          <strong>
            {{ errorMessage }}
          </strong>
        </p>
        <!-- New Reimbursment created table -->

        <p v-if="createReimbursmentMessage" text-align: center style="color:green">
          <strong>{{ createReimbursmentMessage }}</strong>
          <v-data-table
            v-if="this.submitted[0].rId != null"
            :headers="submittedHeaders"
            :items="submitted"
            :items-per-page="1"
            class="elevation-1"
          ></v-data-table>
          <!-- New Reimbursment created table -->
        </p>
        <br />

        <!-- Get Open Reimbursements -->
        <p v-if="getReimbursementMessage" text-align: center style="color:green">
          <strong>{{ getReimbursementMessage }}</strong>
          <v-data-table
            :headers="currentHeaders"
            :items="current"
            :items-per-page="10"
            class="elevation-1"
          >
          </v-data-table>
          <!-- Get Open Reimbursements -->
        </p>
      </v-card>
    </v-container>
  </div>
</template>

<script>
import Axios from "axios";
import axios from "axios";

import store from "..//store/index";
import { required, digits,  regex } from "vee-validate/dist/rules";
import {
  extend,
  ValidationObserver,
  ValidationProvider,
  setInteractionMode,
} from "vee-validate";
import {  numeric } from "vee-validate/dist/rules";
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

extend("regex", {
  ...regex,
  message: "{_field_} {_value_} does not match {regex}",
});


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
          store.commit("setUserId", this.form.userId);
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
      errorMessage: "",
      createReimbursmentMessage: "",
      ApproveOrDenyMessage: "",
      loadOfOpenRequest: false,
      getReimbursementMessage: "",
      submittedHeaders: [
        {
          text: "Reimbursment ID",
          align: "start",
          sortable: false,
          value: "rId",
        },
        { text: "Requestor", value: "requestor" },
        { text: "Title", value: "title" },
        { text: "Description", value: "description" },
        { text: "Amount", value: "amount" },
      ],
      submitted: [
        {
          rId: null,
          requestor: null,
          title: "",
          description: "",
          amount: null,
        },
      ],
      currentHeaders: [
        {
          text: "Reimbursment ID",
          align: "start",
          sortable: false,
          value: "RequestId",
        },
        { text: "Requestor", value: "userID" },
        { text: "Title", value: "title" },
        { text: "Description", value: "description" },
        { text: "Amount", value: "amount" },
      ],
      current: [
        {
          RequestId: null,
          userID: null,
          title: "",
          description: "",
          amount: null,
        },
      ],
      adminCurrentHeaders: [
        {
          text: "Reimbursment ID",
          align: "start",
          sortable: false,
          value: "RequestId",
        },
        { text: "Requestor", value: "userID" },
        { text: "Title", value: "title" },
        { text: "Description", value: "description" },
        { text: "Amount", value: "amount" },
        { text: "Approve" },
        { text: "Deny" },
      ],
      adminCurrent: [
        {
          RequestId: null,
          userID: null,
          title: "",
          description: "",
          amount: null,
          approveOrDeny: "",
        },
      ],
    };
  },
  methods: {
    clear() {
      this.submitted[0].rId = null;
      this.submitted[0].requestor = null;
      this.submitted[0].title = "";
      this.submitted[0].description = "";
      this.submitted[0].amount = null;
      this.createReimbursmentMessage = "";
      this.errorMessage = "";
      document.getElementById("myForm").reset();
      this.getReimbursementMessage = "";
    },
    clearAdmin() {
      this.adminCurrent = [];
      this.ApproveOrDenyMessage = "";
      this.loadOfOpenRequest = false;
    },
    submit() {
      this.form.userId = this.form.userId.toString();
      axios({
        method: "post",
        url: "http://localhost:8000/api/createReimbursment",
        data: this.form,
        withCredentials: true,
      })
        .then(async (response) => {
          this.errorMessage = "";
          this.submitted[0].rId = response.data.RequestId;
          this.submitted[0].requestor = response.data.userID;
          this.submitted[0].title = response.data.title;
          this.submitted[0].description = response.data.description;
          this.submitted[0].amount = "$" + response.data.amount;
          this.createReimbursmentMessage =
            "You successfully created a new Reimbursment Request.";
          document.getElementById("myForm").reset();
        })

        .catch((error) => {
          console.log(error.response.data.message);
          this.errorMessage = error.response.data.message;
          this.createReimbursmentMessage = "";
        });
    },
    getOpenRequests() {
      this.current = [];
      axios({
        method: "post",
        url: "http://localhost:8000/api/getReimbursments",
        withCredentials: true,
      })
        .then(async (response) => {
          for (var x of response.data) {
            x.amount = "$" + x.amount;

            this.current.push(x);
          }
          this.createReimbursmentMessage = "";
          this.getReimbursementMessage =
            "You successfully loaded  all Reimbursment Requests.";
        })
        .catch((error) => {
          console.log(error.response);
          this.errorMessage = error.response;
          this.successMessage = "";
        });
    },
    getAllOpenRequests() {
      this.adminCurrent = [];
      axios({
        method: "post",
        url: "http://localhost:8000/api/getAllOpenReimbursments",
        withCredentials: true,
      })
        .then(async (response) => {
          for (var x of response.data) {
            x.amount = "$" + x.amount;

            this.adminCurrent.push(x);
          }
          this.loadOfOpenRequest = true
        })
        .catch((error) => {
          console.log(error.response);
          this.errorMessage = error.response;
        });
    },
    approveOrDeny(approveOrDeny, requestId) {
      axios({
        method: "post",
        url: "http://localhost:8000/api/approveOrDeny",
        withCredentials: true,
        data: {
          requestId: requestId.toString(),
          approveOrDeny: approveOrDeny,
        },
      })
        .then(async (response) => {
          if (response.data.approvalStatus == "A") {
            this.ApproveOrDenyMessage =
              "You successfully approved Reimbursment Requests #" + response.data.RequestId;
          } else {
            this.ApproveOrDenyMessage =
              "You successfully denied Reimbursment Requests #" + response.data.RequestId;
          }
         this.getAllOpenRequests()
        })
        .catch((error) => {
          console.log(error);
          this.ApproveOrDenyMessage = error;
        });
        
    },
  },
};
</script>
