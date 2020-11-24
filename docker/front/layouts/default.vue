<template>
  <v-app>
    <Navbar/>
    <Footer/>
  </v-app>
</template>

<script>
import Navbar from "@/components/Navbar";
import Footer from "@/components/Footer";
export default {
  components: {Navbar, Footer},
  middleware({ route, store }) {
    if (route.path.match(/admin/)) {
      if (route.path.match(/login/)) {
        console.log('ログイン画面です。')
        store.commit('check/adminFalse')
        console.log(store.getters['check/adminStatus'])
      } else {
        console.log('admin画面です。')
        store.commit('check/adminTrue')
        console.log(store.getters['check/adminStatus'])
      }
    } else {
      console.log('通常画面です。')
      store.commit('check/adminFalse')
      console.log(store.getters['check/adminStatus'])
    }
  },
  computed: {
    admin () {
      return this.$store.getters['check/adminStatus']
    }
  }
}
</script>

<style>
html {
  font-family:
    'Source Sans Pro',
    -apple-system,
    BlinkMacSystemFont,
    'Segoe UI',
    Roboto,
    'Helvetica Neue',
    Arial,
    sans-serif;
  font-size: 16px;
  word-spacing: 1px;
  -ms-text-size-adjust: 100%;
  -webkit-text-size-adjust: 100%;
  -moz-osx-font-smoothing: grayscale;
  -webkit-font-smoothing: antialiased;
  box-sizing: border-box;
}

*,
*::before,
*::after {
  box-sizing: border-box;
  margin: 0;
}

.button--green {
  display: inline-block;
  border-radius: 4px;
  border: 1px solid #3b8070;
  color: #3b8070;
  text-decoration: none;
  padding: 10px 30px;
}

.button--green:hover {
  color: #fff;
  background-color: #3b8070;
}

.button--grey {
  display: inline-block;
  border-radius: 4px;
  border: 1px solid #35495e;
  color: #35495e;
  text-decoration: none;
  padding: 10px 30px;
  margin-left: 15px;
}

.button--grey:hover {
  color: #fff;
  background-color: #35495e;
}
</style>
