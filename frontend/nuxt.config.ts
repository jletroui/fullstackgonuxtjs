// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  ssr: false,
  compatibilityDate: '2024-04-03',
  devtools: { enabled: true },
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
    modules: ['@nuxt/eslint', '@nuxt/test-utils/module'],
    runtimeConfig: {
      public: {
        apiBaseUrl: 'http://localhost:8080/api'
      }
    }
  },
  $test: {
    modules: ['@nuxt/eslint', '@nuxt/test-utils/module'],
    runtimeConfig: {
      public: {
        apiBaseUrl: '' // Allow access to mocked endpoint by Nitro server: https://nuxt.com/docs/getting-started/testing#registerendpoint
      }
    }
  }
})
