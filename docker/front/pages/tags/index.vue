<template>
  <div>
    <h1 style="text-align: center; margin-top: 20px;">タグリスト</h1>
    <v-container style="display: flex; flex-wrap: wrap; justify-content: flex-start; max-width: 500px;">
      <div v-for="(tag,index) in tags" :key="index">
        <nuxt-link :to="`/tags/` + tag.Name">{{ tag.Name }}</nuxt-link>
      </div>
    </v-container>
  </div>
</template>

<script>
export default {
  name: "index",
  auth: false,
  data() {
    return {
      tags: null,
      meta: {
        title: '',
        description: '',
        url: '',
        keywords: ''
      }
    }
  },
  head() {
    this.meta.title = "Tag一覧";
    this.meta.description = "タグ一覧ページです。";
    this.meta.url = "https://www.watson-sei.tokyo" + this.$route.path;
    return {
      title: this.meta.title,
      meta: [
        { hid: 'description', name: 'description', content: this.meta.description },
        { hid: 'keywords', name: 'keywords', content: this.meta.keywords },
        { hid: 'og:title', property: 'og:title', content: this.meta.title },
        { hid: 'og:type', property: 'og:type', content: 'tag' },
        { hid: 'og:url', property: 'og:url', content: this.meta.url },
        { hid: 'og:description', property: 'og:description', content: this.meta.description },
        { hid: "twitter:title", property: "twitter:title", content: this.meta.title },
        { hid: "twitter:description", property: "twitter:description", content: this.meta.description },
        { hid: "twitter:image", property: "twitter:image", content: "https://cdn.vuetifyjs.com/images/cards/docks.jpg"},
      ]
    }
  },
  mounted() {
    this.$axios.get(`/v1/article/tags`)
    .then(res => {
      this.tags = res.data
      for(let i = 0; i < this.tags.length; i++) {
        this.meta.keywords += this.tags[i].Name + ","
      }
    })
    .catch(err => {
      console.log(err)
    })
  }
}
</script>

<style scoped>
a {
  display: inline-block;
  margin: 0 .1em .6em 0;
  padding: .6em;
  line-height: 1;
  text-decoration: none;
  color: #0000ee;
  background-color: #fff;
  border: 1px solid #0000ee;
  border-radius: 2em;
}
</style>
