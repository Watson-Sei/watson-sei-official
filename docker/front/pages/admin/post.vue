<template>
  <div class="container">
    <v-md-editor
      v-model="text"
      height="600px"
      :disabled-menus="[]"
      @upload-image="handleUploadImage"
    />
  </div>
</template>

<script>

export default {
  name: "post",
  data() {
    return {
      text: ''
    }
  },
  methods: {
    handleUploadImage: function (event, insertImage, files) {
      // ここで画像をファイルサーバーに転送し、返り値としてURLを受け取れるようにする。
      const params = new FormData();
      params.append('file', files[0]);
      const { v4: uuid_v4 } = require('uuid');
      params.append('uuid', uuid_v4())
      this.$axios
        .$post(
          'https://localhost/api/v1/upload/image',
          params,
          {
            headers: {
              'content-type': 'multipart/form-data'
            }
          }
        )
        .then((response) => {
          console.log('response data', response.url)
          // これで埋め込む
          insertImage({
            // url: 'https://zukan.pokemon.co.jp/zukan-api/up/images/index/f8d806f32ee833db68f00e2c50b136be.png',
            url: 'https://localhost/api/' + response.url,
            desc: 'desc',
          });
        })
        .catch(error => {
          console.log('response error', error)
        })
    },
  },
};
</script>

<style scoped>
  .container {
    margin: 0 auto;
    min-height: 100vh;
    display: flex;
    justify-content: center;
    align-items: center;
  }
</style>
