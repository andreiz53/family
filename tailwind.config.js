const { default: daisyui } = require("daisyui");

/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./web/**/*.{templ,go}"],
  theme: {
    extend: {},
  },
  daisyui: {
    themes: ["night"],
  },
  plugins: [require("daisyui")],
};
