// https://v3.nuxtjs.org/api/configuration/nuxt.config
export default defineNuxtConfig({

    css: [
        // 'bootstrap/dist/css/bootstrap.min.css',
        '@/assets/css/main.css',
        'tailwindcss/dist/tailwind.min.css'
    ],
    script: [
        {
            src: 'bootstrap/dist/js/bootstrap.bundle.min.js'
        }
    ],
    vite: {
        define: {
            'process.env.DEBUG': false,
        },
    },
})