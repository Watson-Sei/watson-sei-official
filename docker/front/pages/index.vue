<template>
  <div>
    <h1 style="text-align: center; margin-top: 20px">新着記事！</h1>
    <!-- card sample -->
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
          :to="`/article/` + post.id"
          nuxt
        >
          記事を読む
        </v-btn>
      </v-card-actions>
    </v-card>
    <!-- end -->
  </div>
</template>

<script>
export default {
  auth: false,
  data() {
    return {
      posts: null,
      meta: {
        title: "",
        description: "",
        type: "",
        url: "",
      }
    }
  },
  head() {
    const path = this.$route.path;
    this.meta.title = '一覧ページ'
    this.meta.description = 'ブログサイトの記事を全て表示しています。'
    this.meta.url = "https://www.watson-sei.tokyo" + path;

    return {
      title: this.meta.title,
      meta: [
        { hid: "description", name: "description", content: this.meta.description },
        { hid: "keywords", name: "keywords", content: "watson-sei, blog, programming"},
        { hid: "og:title", property: "og:title", content: this.meta.title },
        { hid: "og:description", property: "og:description", content: this.meta.description },
        { hid: "og:type", property: "og:type", content: "article"},
        { hid: "og:url", property: "og:url", content: this.meta.url},
        { hid: "twitter:title", property: "twitter:title", content: this.meta.title },
        { hid: "twitter:description", property: "twitter:description", content: this.meta.description },
        { hid: "twitter:image", property: "twitter:image", content: "https://cdn.vuetifyjs.com/images/cards/docks.jpg"},
      ]
    }
  },
  mounted() {
    this.$axios.get(`/v1/article/list`)
    .then(res => {
      this.posts = res.data.reverse()
    })
    .catch(err => {
      console.log(err)
    })
  },
}
</script>

<style>

</style>
