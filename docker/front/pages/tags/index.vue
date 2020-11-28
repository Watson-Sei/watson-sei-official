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
    }
  },
  mounted() {
    this.$axios.get('https://localhost/api/v1/article/tags')
    .then(res => {
      this.tags = res.data
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
