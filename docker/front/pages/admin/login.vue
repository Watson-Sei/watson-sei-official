<template>
  <v-card width="600" style="margin: 20px auto;">
    <v-card-title>
      <h1 class="display-1">ログイン</h1>
    </v-card-title>
    <v-card-text>
      <v-form>
        <v-text-field v-model="username" prepend-icon="mdi-account-circle" label="ユーザー名" />
        <v-text-field v-model="password" prepend-icon="mdi-lock"  append-icon="mdi-eye-off" type="password" label="パスワード" />
        <v-card-actions>
          <v-btn @click="loginWithAuthModule">ログイン</v-btn>
        </v-card-actions>
      </v-form>
    </v-card-text>
  </v-card>
</template>

<script>
export default {
  name: "login",
  data() {
    return {
      username: '',
      password: ''
    }
  },
  methods: {
    async loginWithAuthModule() {
      await this.$auth.loginWith('local',{
        data: {
          username: this.username,
          password: this.password
        }
      })
      .then((response) => {
        return response
      })
      .catch(error => {
        console.log("response error", error)
      })
    }
  }
}
</script>

<style scoped>

</style>
