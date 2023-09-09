// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
	devtools: { enabled: true },
	modules: ['nuxt-quasar-ui', '@nuxtjs/tailwindcss', '@pinia/nuxt'],
	pinia: {
		autoImports: [
			// automatically imports `defineStore`
			'defineStore', // import { defineStore } from 'pinia'
			'acceptHMRUpdate', // import { acceptHMRUpdate } from 'pinia'
		],
	},
	imports: {
		dirs: ['stores'],
	},
});
