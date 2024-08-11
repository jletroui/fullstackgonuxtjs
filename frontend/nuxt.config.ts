// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  ssr: false,
  compatibilityDate: '2024-04-03',
  devtools: { enabled: true },
  modules: ["@nuxt/eslint"],
  typescript: {
    typeCheck: true
  },
  $production: {
    runtimeConfig: {
      public: {
        apiBaseUrl: '/api'
      }
    }
  },
  $development: {
    runtimeConfig: {
      public: {
        apiBaseUrl: 'http://localhost:8080/api'
      }
    }
  }
})
