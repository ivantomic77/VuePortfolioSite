/** @type {import('tailwindcss').Config} */
module.exports = {
    content: ["./index.html", "./src/**/*.{vue,js,ts,jsx,tsx}"],
    theme: {
        extend: {},
        colors: {
            background: "#ACACAC",
            black: "#090909",
            white: "#F0F0F0",
            red: "#FF605C",
            orange: "#FFBD44",
            green: "#00CA4E",
        },
        fontFamily: {
            body: '"Open Sans"',
        },
    },
    plugins: [],
};
