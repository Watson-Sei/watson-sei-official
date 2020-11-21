<template>
  <div>
    <h1>Main Page</h1>
    <ul>
      <li v-for="(post, index) in posts" :key="index">
        <a :href="`/article/` + post.id" target="_blank" rel="noopener noreferrer">{{ post.title }}</a>
      </li>
    </ul>
  </div>
</template>

<script>
export default {
  auth: false,
  data() {
    return {
      posts: null,
    }
  },
  mounted() {
    this.$axios.get('https://localhost/api/v1/user/list')
    .then(res => {
      this.posts = res.data
    })
    .catch(err => {
      console.log(err)
    })
  }
}
</script>

<style>

</style>
