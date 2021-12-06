<template>
  <!-- BEGIN ADMIN STUFF -->
  <div style="text-align: center" v-if="this.$store.state.isAdmin">
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

        <v-btn @click="adminGetHistory"> Get History </v-btn>
        <br /><br />

        <v-btn @click="clear"> Clear </v-btn>

        <!-- Get Open Reimbursements -->
        <p v-if="getReimbursementMessage" style="text-align: center">
          <br />
          <strong style="color: green">{{ getReimbursementMessage }}</strong>
          <v-data-table
            v-if="showTable"
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
  <div style="text-align: center" v-else-if="!this.$store.state.isAdmin">
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

        <v-btn @click="getHistory"> Get History </v-btn>
        <br /><br />

        <v-btn @click="clear"> Clear </v-btn>

        <!-- Get Open Reimbursements -->
        <p v-if="getReimbursementMessage" style="text-align: center">
          <strong style="color: green">{{ getReimbursementMessage }}</strong>
          <v-data-table
            v-if="showTable"
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
      showTable: false,
      firstName: "",
      userType: "",
      errorMessage: "",
      getReimbursementMessage: "",
      currentHeaders: [
        {
          text: "Reimbursment ID",
          align: "start",
          sortable: false,
          value: "request_id",
        },
        { text: "Requestor ID", value: "user_id" },
        { text: "Requestor Name", value: "submitter" },

        { text: "Title", value: "title" },
        { text: "Description", value: "description" },
        { text: "Amount", value: "amount" },
        {
          text: "Reviewer ID",
          value: "approved_by",
        },
        {
          text: "Reviewer Name",
          value: "approver",
        },
        { text: "Date Reviewed", value: "date_approved" },
        { text: "Status", value: "approved_status" },
      ],
      current: [
        {
          RequestId: null,
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
      this.getReimbursementMessage = "";
    },
    getHistory() {
      this.current = [];
      axios({
        method: "get",
        url: "http://localhost:8000/api/getHistory",
        withCredentials: true,
      })
        .then(async (response) => {
          if (response.data.length != 0) {
            for (var x of response.data) {
              x.amount = "$" + x.amount;

              this.current.push(x);
            }
            this.getReimbursementMessage =
              "You successfully loaded your completed requests.";
            this.showTable = true;
          } else {
            this.getReimbursementMessage = "No Current History for this User.";
            this.showTable = false;
          }
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
        method: "get",
        url: "http://localhost:8000/api/getAllHistory",
        withCredentials: true,
      })
        .then(async (response) => {
          if (response.data.length != 0) {
            for (var x of response.data) {
              x.amount = "$" + x.amount;
              this.current.push(x);
            }

            this.getReimbursementMessage =
              "You successfully loaded your completed requests.";
            this.showTable = true;
          } else {
            this.getReimbursementMessage = "No completed requests available.";
            this.showTable = false;
          }
        })
        .catch((error) => {
          this.errorMessage = error.response;
          this.successMessage = "";
        });
    },
  },
};
</script>
