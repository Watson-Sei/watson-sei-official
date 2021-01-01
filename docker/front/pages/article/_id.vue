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
        url: ''
      }
    }
  },
  head() {
    this.meta.url = "https://www.watson-sei.tokyo" + this.$route.path;
    return {
      title: this.meta.title,
      meta: [
        { hid: 'description', name: 'description', content: this.meta.overview },
        { hid: 'og:title', name: 'og:title', content: this.meta.title },
        { hid: 'og:type', name: 'og:type', content: 'article'},
        { hid: 'og:url', name: 'og:url', content: this.meta.url },
        { hid: 'og:description', name: 'og:description', content: this.meta.overview }
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
    })
    .catch(err => {
      console.log(err)
    })
  }
}
</script>

<style scoped>

</style>
