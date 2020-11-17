<template>
    <div>
      <h1>About</h1>
      <v-btn @click="logout">Logout</v-btn>
    </div>
</template>

<script>
    export default {
        name: "about",
        methods: {
          logout() {
            this.$auth.logout();
          },
        }
    }
</script>

<style scoped>

</style>
