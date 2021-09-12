<template>
  <!-- BEGIN ADMIN STUFF -->
  <div align="center" v-if="this.$store.state.isAdmin">
    Admin USER STUFF HERE

    <v-container class="grey lighten-5" fill-height fluid>
      <v-card class="pa-md-4 mx-lg-auto" color="white" width="auto">
        <p>
          <strong>
            Welcome {{ this.firstName }}
            <br />
            <br />
            You are currently logged in as a
            {{ this.userType }} User <br />
            <br />
            To Get a History of All previous requests please click below.
          </strong>
          <br />
        </p>

        <v-btn @click="adminGetHistory">
          Get History
        </v-btn>
        <br /><br />

        <v-btn @click="clear">
          Clear
        </v-btn>

        <!-- Get Open Reimbursements -->
        <p v-if="getReimbursementMessage" align="center" style="color:green">
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
            <br />
            To Get a History of your previous requests please click below.
          </strong>
          <br />
        </p>

        <v-btn @click="getHistory">
          Get History
        </v-btn>
        <br /><br />

        <v-btn @click="clear">
          Clear
        </v-btn>

        <!-- Get Open Reimbursements -->
        <p v-if="getReimbursementMessage" align="center" style="color:green">
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

export default {
  components: {},
  created() {
    Axios.get("http://localhost:8000/api/user", {
      withCredentials: true,
      headers: {
        "Access-Control-Allow-Origin": "*",
        "Content-Type": "application/json",
      },
    })
      .then(async (response) => {
        this.userId = response.data.id;
        store.commit("setUserId", this.userId);
        this.firstName = response.data.firstName;
        if (response.data.userType == "R") {
          this.userType = "Regular";
        } else if (response.data.userType == "A") {
          this.userType = "Admin";
          store.commit("makeUserAdmin");
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
      errorMessage: "",
      getReimbursementMessage: "",
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
        { text: "Reviewed By", value: "approvedBy" },
        { text: "Date Reviewed", value: "dateApproved" },
        { text: "Status", value: "approvalStatus" },
      ],
      current: [
        {
          RequestId: null,
          userID: null,
          title: "",
          description: "",
          amount: null,
          approvedBy: null,
          dateApproved: "",
        },
      ],
    };
  },
  methods: {
    clear() {
      console.log("clear clicked");
      this.getReimbursementMessage = "";
    },
    getHistory() {
      this.current = [];
      axios({
        method: "post",
        url: "http://localhost:8000/api/getHistory",
        withCredentials: true,
      })
        .then(async (response) => {
          for (var x of response.data) {
            x.amount = "$" + x.amount;

            this.current.push(x);
          }
          console.log(response);

          this.getReimbursementMessage =
            "You successfully loaded your completed requests.";
        })
        .catch((error) => {
          console.log(error.response);
          this.errorMessage = error.response;
          this.successMessage = "";
        });
    },
        adminGetHistory() {
      this.current = [];
      axios({
        method: "post",
        url: "http://localhost:8000/api/getAllHistory",
        withCredentials: true,
      })
        .then(async (response) => {
          for (var x of response.data) {
            x.amount = "$" + x.amount;

            this.current.push(x);
          }
          console.log(response);

          this.getReimbursementMessage =
            "You successfully loaded your completed requests.";
        })
        .catch((error) => {
          console.log(error.response);
          this.errorMessage = error.response;
          this.successMessage = "";
        });
    },
  },
};
</script>
