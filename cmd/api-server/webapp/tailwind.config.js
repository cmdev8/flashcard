/** @type {import('tailwindcss').Config} */
export default {
  content: ["./src/**/*.{html,js,ts,jsx,tsx}"],
  theme: {
    extend: {},
  },
  plugins: [
    // require('@tailwindcss/forms'),
    require('@tailwindcss/typography'),
  ],
}

