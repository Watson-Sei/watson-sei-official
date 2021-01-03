const enviroment = process.env.NODE_ENV
require('dotenv').config({path: `config/.env.${enviroment}`})
const { API_URL, API_URL_BROWSER,BASE_URL } = process.env
console.log('Nuxtjs ENV:',process.env.NODE_ENV)
export default {
  /*
  ** Nuxt rendering mode
  ** See https://nuxtjs.org/api/configuration-mode
  */
  mode: 'universal',
  /*
  ** Nuxt target
  ** See https://nuxtjs.org/api/configuration-target
  */
  target: 'server',
  /*
  ** Headers of the page
  ** See https://nuxtjs.org/api/configuration-head
  */
  head: {
    // title: process.env.npm_package_name || '',
    htmlAttrs: {
      prefix: 'og: http://ogp.me/ns#'
    },
    title: "watson-sei-blog",
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      { hid: 'og:image', property: 'og:image', content: 'https://cdn.vuetifyjs.com/images/cards/docks.jpg'},
      { hid: 'twitter:card', name: 'twitter:card', content: 'summary'},
      { hid: 'twitter:site', name: 'twitter:site', content: '@watson_sei'}
    ],
    link: [
      { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' }
    ]
  },
  /*
  ** Global CSS
  */
  css: [
  ],
  vuetify: {
    customVariables: ['~/assets/common/common.scss'],
    treeShake: true,
  },
  /*
  ** Plugins to load before mounting the App
  ** https://nuxtjs.org/guide/plugins
  */
  plugins: [
    { src: '@/plugins/editor.js', mode: 'client' },
  ],
  /*
  ** Auto import components
  ** See https://nuxtjs.org/api/configuration-components
  */
  components: true,
  /*
  ** Nuxt.js dev-modules
  */
  buildModules: [
    '@nuxtjs/vuetify',
  ],
  /*
  ** Nuxt.js modules
  */
  modules: [
    '@nuxtjs/axios',
    '@nuxtjs/auth',
    '@nuxtjs/auth-next',
  ],
  /*
  ** Build configuration
  ** See https://nuxtjs.org/api/configuration-build/
  */
  build: {
  },
  axios: {
    baseURL: process.env.API_URL,
    browserBaseURL: process.env.API_URL_BROWSER
  },
  auth: {
    redirect: {
      login: '/admin/login',
      logout: '/admin/login',
      callback: false,
      home: '/admin'
    },
    strategies: {
      local: {
        scheme: 'refresh',
        token: {
          property: 'access_token',
        },
        refreshToken: {
          property: 'refresh_token',
        },
        endpoints: {
          login: {url: '/admin/login', method: 'post'},
          logout: {url: '/admin/logout', method: 'get'},
          refresh: {url: '/admin/refresh', method: 'get'},
          user: false
        },
      }
    }
  },
  router: {
    middleware: ['auth']
  },
  env: {
    API_URL,
    API_URL_BROWSER,
    BASE_URL
  }
}
