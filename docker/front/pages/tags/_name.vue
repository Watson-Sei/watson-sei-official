<template>
  <div>
    <h1 style="text-align: center; margin-top: 20px;">タグ: {{ this.$route.params.name }}</h1>
    <v-card
      v-for="(post, index) in posts"
      :key="index"
      class="mx-auto"
      style="margin-top: 15px; margin-bottom: 15px;"
      max-width="400"
    >
      <v-img
        class="white--text align-end"
        height="200px"
        src="https://cdn.vuetifyjs.com/images/cards/docks.jpg"
      >
        <v-card-title>
          <span class="d-inline-block text-truncate" style="max-width: 400px;">{{ post.title }}</span>
        </v-card-title>
      </v-img>
      <v-card-subtitle class="pb-0">
        <v-chip
          v-for="(tag, index) in post.tags"
          :key="index"
          class="ma-2"
          color="primary"
          @click="$router.push('/tags/' + tag.Name)"
        >
          {{ tag.Name }}
        </v-chip>
      </v-card-subtitle>
      <v-card-text class="text--primary">
        <div v-if="post.overview == ''">
          概要なし
        </div>
        <div v-else>
          {{ post.overview }}
        </div>
      </v-card-text>
      <v-card-actions>
        <v-btn
          color="green lighten-1"
          :href="`/article/` + post.id"
        >
          記事を読む
        </v-btn>
      </v-card-actions>
    </v-card>
  </div>
</template>
<script>
export default {
  name: "_name",
  auth: false,
  data() {
    return {
      posts: null,
      meta: {
        title: '',
        description: '',
        url: '',
        keywords: ''
      }
    }
  },
  head() {
    this.meta.title = 'Tag - ' + this.$route.params.name;
    this.meta.description = 'タグの詳細ページです。';
    this.meta.url = 'https://www.watson-sei.tokyo' + this.$route.path;
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
    this.$axios.get(`/v1/article/tags/` + this.$route.params.name)
    .then(res => {
      this.posts = res.data
      this.meta.keywords = this.$route.params.name
    })
    .catch(err => {
      console.log(err)
    })
  }
}
</script>

<style scoped>

</style>
