<template>
  <v-container>
    <v-row>
      <v-text-field  v-model="title" label="タイトル" solo></v-text-field>

      <v-flex xs12>
        <v-combobox multiple
                    v-model="select"
                    label="Tags"
                    chips
                    class="tag-input"
                   >
        </v-combobox>
      </v-flex>

      <p>概要</p>
      <v-textarea v-model="overview"></v-textarea>

      <v-md-editor
        v-model="text"
        height="600px"
        :disabled-menus="[]"
        @upload-image="handleUploadImage"
      />
      <div class="d-flex justify-end">
        <v-btn @click="articleUpdate">記事を更新</v-btn>
      </div>
    </v-row>
  </v-container>
</template>

<script>
export default {
  name: "_id",
  data() {
    return {
      checked: false,
      title: '',
      overview: '',
      text: '',
      select: [],
    }
  },
  mounted() {
    this.$axios.$get(`/v1/article/detail/` + this.$route.params.id)
    .then(res => {
      this.title = res.title
      this.overview = res.overview
      this.text = res.text
      this.selectAdd(res.tags)
    })
    .catch(err => {
      console.log(err)
      this.$router.push('/admin')
    })
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
          `/v1/upload/image`,
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
            url: `${process.env.BASE_URL}/api/` + response.url,
            desc: 'desc',
          });
        })
        .catch(error => {
          console.log('response error', error)
        })
    },
    articleUpdate: async function () {
      this.$axios.$put(`/v1/article/update/` + this.$route.params.id, {
        title: this.title,
        overview: this.overview,
        text: this.text,
      })
        .then((response) => {
          console.log("success")
          this.$router.push('/admin')
        })
        .catch((err) => {
          console.log(err)
        })
    },
    selectAdd: function (tags) {
      for (let i = 0; i < tags.length; i++) {
        this.select.push(tags[i].Name)
      }
    },
  }
}
</script>

<style scoped>

</style>
