/** @type {import('tailwindcss').Config} */
export default {
	content: [
		"./index.html",
		"./src/**/*.{vue,js,ts,jsx,tsx}"
	],
	theme: {
		extend: {
			colors: {
				brown: {
					DEFAULT: '#e6b999',
				},
			}
		}
	},
	plugins: [],
}
