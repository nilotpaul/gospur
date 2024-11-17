/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ['web/**/*.html'],
  theme: {
    extend: {},
  },
  darkMode: 'class',
  plugins: [
    require('@tailwindcss/forms'),
    require('@tailwindcss/typography'),
    require('preline/plugin'),
  ],
}
