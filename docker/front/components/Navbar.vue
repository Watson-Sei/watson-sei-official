<template>
  <v-app>
    <v-app-bar
      app
      dark
    >
      <v-app-bar-nav-icon @click="drawer = true"></v-app-bar-nav-icon>
      <v-toolbar-title>Watson Sei Blog</v-toolbar-title>
      <!-- navbar tab -->
      <v-tabs>
        <v-tab
          v-for="(menuItem, index) in menuItems"
          :key="index"
          @click="push(menuItem.url)"
        >
          {{ menuItem.name }}
        </v-tab>
      </v-tabs>
    </v-app-bar>
    <v-navigation-drawer
      v-model="drawer"
      fixed
      temporary
    >
      <v-list
        nav
        dense
      >
        <v-list-item-group>
          <v-list-item
            v-for="(menuItem, index) in menuItems"
            :key="index"
            :href="menuItem.url"
          >
            <v-list-item-title>{{ menuItem.name }}</v-list-item-title>
          </v-list-item>
        </v-list-item-group>
      </v-list>
    </v-navigation-drawer>

    <v-main>
      <Nuxt />
    </v-main>
  </v-app>
</template>

<script>
export default {
  name: "Navbar",
  data() {
    return {
      drawer: false,
      menuItems: [
        {
          name: 'BLOG',
          url: '/',
        },
        {
          name: 'Tag',
          url: '/tags',
        },
        {
          name: 'ABOUT',
          url: '/about',
        },
      ]
    }
  },
  methods: {
    push: function (url) {
      this.$router.push(url)
    }
  }
}
</script>

<style lang="scss" scoped>
  .v-toolbar__title {
    overflow: visible !important;
    margin-right: 50px !important;
  }
  .v-app-bar__nav-icon {
    @include display_pc {
      display: none !important;
    }
  }
  .v-tabs {
    display: none;

    @include display_pc {
      display: block !important;
    }
  }
</style>
