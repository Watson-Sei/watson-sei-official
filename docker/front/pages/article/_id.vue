<template>
  <div>
    <h1>{{ articleTitle }}</h1>
    <v-md-preview :text="articleText"></v-md-preview>
  </div>
</template>

<script>
export default {
  name: "_id",
  auth: false,
  data() {
    return {
      articleTitle: '',
      articleText: '',
    }
  },
  mounted() {
    this.$axios.get('https://localhost/api/v1/user/detail/' + this.$route.params.id)
    .then(res => {
      this.articleTitle = res.data.title
      this.articleText = res.data.text
    })
    .catch(err => {
      console.log(err)
    })
  }
}
</script>

<style scoped>

</style>
