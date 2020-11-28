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
      posts: null
    }
  },
  mounted() {
    this.$axios.get('https://localhost/api/v1/article/tags/' + this.$route.params.name)
    .then(res => {
      this.posts = res.data
    })
    .catch(err => {
      console.log(err)
    })
  }
}
</script>

<style scoped>

</style>
