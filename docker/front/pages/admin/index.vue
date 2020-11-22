<template>
  <v-container>
    <h1>Admin Page</h1>
    <v-row>
      <v-col cols="10">
        <v-data-table
          :headers="headers"
          :items="posts"
        >
          <template v-slot:item.delete="{ item }">
            <v-btn
              small
              color="error"
              @click="deleteItem(item)"
            >
              delete
            </v-btn>
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
        </v-data-table>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
export default {
  name: "index",
  data() {
    return {
      posts: [],
      headers: [
        {
          text: 'ID',
          value: 'id',
          align: 'center',
        },
        {
          text: 'Title',
          value: 'title',
          align: 'center',
        },
        {
          text: 'Actions',
          value: 'actions',
          align: 'center',
          sortable: false
        },
      ],
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
  methods: {
    deleteItem (item) {
      console.log(item.id)
    }
  }
}
</script>

<style scoped>

</style>
