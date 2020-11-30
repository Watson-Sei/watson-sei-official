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
    }
  },
  mounted() {
    this.$axios.get(`${process.env.API}/v1/article/detail/` + this.$route.params.id)
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
