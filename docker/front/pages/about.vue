<template>
  <div>
    <v-container style="margin-top: 60px;">
      <v-row justify="center">
        <v-avatar size="200">
          <img
            src="~/assets/image/me.jpg"
            alt="Watson Sei"
          >
        </v-avatar>
      </v-row>
      <v-row justify="center">
        <h1>Watson-Sei</h1>
      </v-row>
      <v-row justify="center">
        <p style="padding: 12px; max-width: 500px;">Go/TypeScriptを使って個人開発をしている高校生です。</p>
      </v-row>
      <v-row justify="space-around" style="max-width: 500px; margin: auto">
        <v-btn icon href="https://github.com/Watson-Sei" target="_blank">
          <v-icon large>mdi-github</v-icon>
        </v-btn>
        <v-btn icon href="https://twitter.com/watson_sei" target="_blank">
          <v-icon large>mdi-twitter</v-icon>
        </v-btn>
        <v-btn icon href="mailto:info#64;seinabehack@gmail.com">
          <v-icon large>mdi-email</v-icon>
        </v-btn>
      </v-row>
    </v-container>
  </div>
</template>

<script>
export default {
  name: "about"
}
</script>

<style scoped>

</style>
