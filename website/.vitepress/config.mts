import { defineConfig } from 'vitepress'

export default defineConfig({
  title: 'Antigravity Factory',
  description: 'AI Agent Skills Catalog',
  base: '/antigravity-factory/',
  ignoreDeadLinks: true,
  
  themeConfig: {
    nav: [
      { text: 'Home', link: '/' },
      { text: 'Skills', link: '/skills/' },
      { text: 'GitHub', link: 'https://github.com/ydnikolaev/antigravity-factory' }
    ],

    sidebar: {
      '/skills/': [
        {
          text: 'Core Pipeline',
          items: [
            { text: 'idea-interview', link: '/skills/idea-interview' },
            { text: 'product-analyst', link: '/skills/product-analyst' },
            { text: 'bmad-architect', link: '/skills/bmad-architect' },
            { text: 'tech-spec-writer', link: '/skills/tech-spec-writer' },
            { text: 'qa-lead', link: '/skills/qa-lead' },
          ]
        },
        {
          text: 'Backend',
          items: [
            { text: 'backend-go-expert', link: '/skills/backend-go-expert' },
            { text: 'mcp-expert', link: '/skills/mcp-expert' },
            { text: 'debugger', link: '/skills/debugger' },
          ]
        },
        {
          text: 'Frontend',
          items: [
            { text: 'frontend-nuxt', link: '/skills/frontend-nuxt' },
            { text: 'ux-designer', link: '/skills/ux-designer' },
            { text: 'ui-implementor', link: '/skills/ui-implementor' },
          ]
        },
        {
          text: 'CLI/TUI',
          items: [
            { text: 'cli-architect', link: '/skills/cli-architect' },
            { text: 'tui-charm-expert', link: '/skills/tui-charm-expert' },
          ]
        },
        {
          text: 'Telegram',
          items: [
            { text: 'tma-expert', link: '/skills/tma-expert' },
            { text: 'telegram-mechanic', link: '/skills/telegram-mechanic' },
          ]
        },
        {
          text: 'DevOps & Utility',
          items: [
            { text: 'devops-sre', link: '/skills/devops-sre' },
            { text: 'project-bro', link: '/skills/project-bro' },
            { text: 'refactor-architect', link: '/skills/refactor-architect' },
            { text: 'doc-janitor', link: '/skills/doc-janitor' },
            { text: 'feature-fit', link: '/skills/feature-fit' },
          ]
        }
      ]
    },

    socialLinks: [
      { icon: 'github', link: 'https://github.com/ydnikolaev/antigravity-factory' }
    ],
    
    footer: {
      message: 'AI Agent Skills for Modern Development',
      copyright: '© 2026 Antigravity Factory'
    }
  }
})
