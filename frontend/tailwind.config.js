/** @type {import('tailwindcss').Config} */
export default {
  content: ["./index.html", "./src/**/*.{vue,js,ts,jsx,tsx}"],
  theme: {
    extend: {
      colors: {
        "main-light-grey": "#EDF6FF",
        "main-gold": "#B9732F",
        "main-dark-blue": "#437FC7",
        "main-blue": "#6DAFFE",
        "main-white": "#F8F9FE",
        "main-grey": "#2E363F",
      },
    },
    fontFamily: {
      redhat: ["Red Hat Display", "sans-serif"],
    },
  },
  plugins: [],
};
