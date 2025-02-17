/** @type {import('tailwindcss').Config} */
module.exports = {
    content: ["./src/**/*.{html,js,jsx}"],
    theme: {
        extend: {
            colors: {
                primary: "#010409",
                editor: "#1e1e1e",
            },
            fontFamily: {
                sans: ['"Fira Code"'],
            },
        },
    },
    plugins: [],
};
