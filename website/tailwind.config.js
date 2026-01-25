/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./.vitepress/**/*.{js,ts,vue}",
    "./**/*.md",
    "!./node_modules/**",
  ],
  darkMode: 'class',
  theme: {
    extend: {
      colors: {
        brand: {
          light: '#38bdf8', // Sky 400
          DEFAULT: '#0ea5e9', // Sky 500
          dark: '#0284c7', // Sky 600
        },
        dark: {
          bg: '#0f172a', // Slate 900
          surface: '#1e293b', // Slate 800
        }
      },
      fontFamily: {
        sans: ['Inter', 'ui-sans-serif', 'system-ui'],
        mono: ['JetBrains Mono', 'ui-monospace', 'SFMono-Regular'],
      },
      animation: {
        'typewriter': 'typewriter 2s steps(40) 1s 1 normal both',
        'blink': 'blink 1s step-end infinite',
      },
      keyframes: {
        typewriter: {
          '0%': { width: '0%' },
          '100%': { width: '100%' },
        },
        blink: {
          '0%, 100%': { opacity: '1' },
          '50%': { opacity: '0' },
        }
      }
    },
  },
  plugins: [],
}
