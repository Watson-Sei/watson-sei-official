<template>
  <v-container>
    <v-row>
      <v-text-field  v-model="title" label="タイトル" solo></v-text-field>
      <v-flex xs12>
        <v-combobox multiple
                    v-model="select"
                    label="Tags"
                    chips
                    deletable-chips
                    class="tag-input"
                    :search-input.sync="search"
                    @keyup.tab="updateTags"
                    @paste="updateTags">
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
      <div>
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
      overview: '',
      text: '',
      select: [],
      search: "",
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
    articleSave: async function () {
      this.$axios.$post(`/v1/article/create`, {
        title: this.title,
        overview: this.overview,
        text: this.text,
        tags: this.selectArray(this.select),
      })
      .then((response) => {
        console.log("success")
        this.$router.push('/admin')
      })
      .catch((err) => {
        console.log(err)
      })
    },
    updateTags() {
      this.$nextTick(() => {
        this.select.push(...this.search.split(","));
        this.$nextTick(() => {
          this.search = "";
        });
      });
    },
    selectArray: function (tags) {
      let array = []
      for (let i = 0; i < tags.length; i++) {
        array.push({Name: tags[i]})
      }
      return array
    },
  },
};
</script>

<style scoped>
</style>
