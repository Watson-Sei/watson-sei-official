<template>
  <div>
    <h1>タグリスト</h1>
    <ul>
      <li v-for="(tag, index) in tags" :key="index">
        <a :href="`/tags/` + tag.Name">{{ tag.Name }}</a>
      </li>
    </ul>
  </div>
</template>

<script>
export default {
  name: "index",
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

</style>
