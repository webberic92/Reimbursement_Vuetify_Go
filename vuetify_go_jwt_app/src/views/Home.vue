<template>
  <hello-world />
</template>

<script>
  import HelloWorld from '../components/HelloWorld'
// import store from "..//store/index"

  export default {
     created() {
            console.log(this.$store.state.isLoggedIn)
             console.log(this.$store.getters.loggedIn)
     },
    name: 'Home',

    components: {
      HelloWorld,
    },
  }
</script>
