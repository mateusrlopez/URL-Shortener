/** @type {import('tailwindcss').Config} */
module.exports = {
    content: ['./index.html', './src/**/*.{vue,ts}'],
    theme: {
        extend: {
            fontFamily: {
                lato: ['Lato', 'sans-serif'],
            },
        },
    },
};
