/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./views/**/*.templ"],
  theme: {
    extend: {
      colors: {
        orange: "#ff6600",
        white: "#e5e5e5",
        yellow: "#f2de8a",
        "light-blue": "#4fb2c0",
        "dark-blue": "#002642",
      },
    },
  },
  plugins: [],
};
