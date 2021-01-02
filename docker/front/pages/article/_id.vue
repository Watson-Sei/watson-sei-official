<template>
  <v-container>
    <h1 style="text-align: center">{{ articleTitle }}</h1>
    <v-md-preview :text="articleText"></v-md-preview>
  </v-container>
</template>

<script>
export default {
  auth: false,
  data() {
    return {
      articleTitle: '',
      articleText: '',
      meta: {
        title: '',
        overview: '',
        url: '',
        tag: ''
      }
    }
  },
  head() {
    this.meta.url = "https://www.watson-sei.tokyo" + this.$route.path;
    return {
      title: this.meta.title,
      meta: [
        { hid: 'description', name: 'description', content: this.meta.overview },
        { hid: 'keywords', name: 'keywords', content: this.meta.tag },
        { hid: 'og:title', property: 'og:title', content: this.meta.title },
        { hid: 'og:type', property: 'og:type', content: 'article'},
        { hid: 'og:url', property: 'og:url', content: this.meta.url },
        { hid: 'og:description', property: 'og:description', content: this.meta.overview },
        { hid: "twitter:title", property: "twitter:title", content: this.meta.title },
        { hid: "twitter:description", property: "twitter:description", content: this.meta.description },
        { hid: "twitter:image", property: "twitter:image", content: "https://cdn.vuetifyjs.com/images/cards/docks.jpg"},
      ]
    }
  },
  mounted() {
    this.$axios.get(`${process.env.API}/v1/article/detail/` + this.$route.params.id)
    .then(res => {
      this.articleTitle = res.data.title
      this.articleText = res.data.text
      this.meta.title = res.data.title
      this.meta.overview = res.data.overview
      for(let i = 0; i < res.data.tags.length; i++) {
        this.meta.tag += res.data.tags[i].Name + ","
      }
    })
    .catch(err => {
      console.log(err)
    })
  }
}
</script>

<style scoped>

</style>
