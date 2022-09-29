// https://v3.nuxtjs.org/api/configuration/nuxt.config
export default defineNuxtConfig({
    // ssr: false,
    css: [
        //'bootstrap/dist/css/bootstrap.min.css',
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
    build: {
        transpile: [
            "@heroicons/vue",
        ]
    }
})