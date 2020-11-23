<template>
  <div>
    <h1>Admin Page</h1>
    <v-app>
      <v-data-table
        :headers="headers"
        :items="posts"
        sort-by="calories"
        class="elevation-1"
      >
        <template v-slot:top>
          <v-toolbar
            flat
          >
            <v-toolbar-title>記事管理</v-toolbar-title>
            <v-divider
              class="mx-4"
              inset
              vertical
            ></v-divider>
            <v-spacer></v-spacer>
            <v-btn
              color="primary"
              dark
              class="mb-2"
              @click="$router.push('/admin/post')"
            >
              New Item
            </v-btn>
            <v-dialog v-model="dialogDelete" max-width="500px">
              <v-card>
                <v-card-title class="headline">Are you sure you want to delete this item?</v-card-title>
                <v-card-actions>
                  <v-spacer></v-spacer>
                  <v-btn color="blue darken-1" text @click="closeDelete">Cancel</v-btn>
                  <v-btn color="blue darken-1" text @click="deleteItemConfirm">OK</v-btn>
                  <v-spacer></v-spacer>
                </v-card-actions>
              </v-card>
            </v-dialog>
          </v-toolbar>
        </template>
        <template v-slot:item.actions="{ item }">
          <v-icon
            small
            class="mr-2"
            @click="editItem(item)"
          >
            mdi-pencil
          </v-icon>
          <v-icon
            small
            @click="deleteItem(item)"
          >
            mdi-delete
          </v-icon>
        </template>
        <template v-slot:no-data>
          <v-btn
            color="primary"
            @click="initialize"
          >
            Reset
          </v-btn>
        </template>
      </v-data-table>
    </v-app>
  </div>
</template>

<script>
export default {
  name: "index",
  data() {
    return {
      dialog: false,
      dialogDelete: false,
      editedIndex: -1,
      itemId: 0,
      headers: [
        {
          text: 'ID',
          align: 'start',
          sortable: false,
          value: 'id',
        },
        { text: 'Title', value: 'title' },
        { text: 'Actions', value: 'actions', sortable: false },
      ],
      posts: [],
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
  },
  watch: {
    dialog (val) {
      val || this.close()
    },
    dialogDelete (val) {
      val || this.closeDelete()
    },
  },
  created() {
    this.initialize()
  },
  methods: {
    initialize () {
      this.posts = []
    },
    deleteItem (item) {
      this.editedIndex = this.posts.indexOf(item)
      this.itemId = item.id
      this.dialogDelete = true
    },
    deleteItemConfirm () {
      this.$axios.$delete(
        'https://localhost/api/v1/user/delete/' + this.itemId
      )
      .then(res => {
        console.log(res)
        this.posts.splice(this.editedIndex, 1)
      })
      .catch(err => {
        console.log('削除に失敗しました。')
      })
      this.closeDelete()
    },
    closeDelete () {
      this.dialogDelete = false
    },
    editItem (item) {
      this.$router.push('/admin/post/' + item.id)
    }
  },
}
</script>

<style scoped>

</style>
