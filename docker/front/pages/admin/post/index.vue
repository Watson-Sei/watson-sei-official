<template>
  <v-container>
    <v-row>
      <v-text-field  v-model="title" label="タイトル" solo></v-text-field>
      <v-md-editor
        v-model="text"
        height="600px"
        :disabled-menus="[]"
        @upload-image="handleUploadImage"
      />
      <div class="d-flex justify-end">
        <v-checkbox v-model="checked" :label="`そのまま公開`"></v-checkbox>
        <v-btn @click="articleSave">記事を投稿</v-btn>
      </div>
    </v-row>
  </v-container>
</template>

<script>

export default {
  data() {
    return {
      checked: false,
      title: '',
      text: '',
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
    articleSave: async function () {
      this.$axios.$post('https://localhost/api/v1/user/create', {
        title: this.title,
        text: this.text,
      })
      .then((response) => {
        console.log("success")
        this.$router.push('/admin')
      })
      .catch((err) => {
        console.log(err)
      })
    }
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
