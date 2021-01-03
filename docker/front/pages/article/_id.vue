<template>
  <v-container>
    <h1 style="text-align: center">{{ post.title }}</h1>
    <v-md-preview :text="post.text"></v-md-preview>
  </v-container>
</template>

<script>
export default {
  auth: false,
  data() {
    return {
      meta: {
        title: '',
        overview: '',
        tag: '',
        url: ''
      }
    }
  },
  head() {
    this.meta.title = this.post.title;
    this.meta.overview = this.post.overview;
    this.meta.url = "https://www.watson-sei.tokyo" + this.$route.path;
    for(let i = 0; i < this.post.tags.length; i++) {
      this.meta.tag += this.post.tags[i].Name + ","
    };
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
  async asyncData(context) {
    const { data } = await context.$axios.get(`/v1/article/detail/${context.params.id}`);
    return {post: data}
  },
}
</script>

<style scoped>

</style>
