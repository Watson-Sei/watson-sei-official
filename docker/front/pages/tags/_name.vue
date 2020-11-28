<template>
  <div>
    <h1>タグ: {{ this.$route.params.name }}</h1>
    <ul>
      <li v-for="(post, index) in posts" :key="index">
        <a :href="`/article/` + post.id">{{ post.title }}</a>
      </li>
    </ul>
  </div>
</template>
<script>
export default {
  name: "_name",
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
