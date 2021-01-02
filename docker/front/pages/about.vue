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
        <p style="padding: 12px; max-width: 500px;">Go/NuxtでWeb開発しています。</p>
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
  name: "about",
  auth: false,
  data() {
    return {
      meta: {
        title: "",
        description: "",
        url: "",
      }
    }
  },
  head() {
    this.meta.title = "About"
    this.meta.description = "自己紹介ページです。";
    this.meta.url = "https://www.watson-sei.tokyo" + this.$route.path;
    return {
      title: this.meta.title,
      meta: [
        { hid: 'description', name: 'description', content: this.meta.description },
        { hid: 'keywords', name: 'keywords', content: 'about, 自己紹介'},
        { hid: 'og:title', property: 'og:title', content: this.meta.title },
        { hid: 'og:description', property: 'og:description', content: this.meta.description },
        { hid: 'og:type', property: 'og:type', content: 'website'},
        { hid: 'og.url', property: 'og.url', content: this.meta.url },
        { hid: "twitter:title", property: "twitter:title", content: this.meta.title },
        { hid: "twitter:description", property: "twitter:description", content: this.meta.description },
        { hid: "twitter:image", property: "twitter:image", content: "https://cdn.vuetifyjs.com/images/cards/docks.jpg"},
      ]
    }
  }
}
</script>

<style scoped>

</style>
