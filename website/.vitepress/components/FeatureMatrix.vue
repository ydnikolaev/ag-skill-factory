<script setup lang="ts">
import { withBase } from 'vitepress'

const features = [
  {
    icon: `<svg class="w-6 h-6" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M2 20a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V8l-7-7H4a2 2 0 0 0-2 2v17Z"/><path d="M17 2v6h6"/><path d="m9 18 3-3-3-3"/><path d="m15 18-3-3 3-3"/></svg>`,
    title: 'Factory CLI',
    desc: 'Native Go CLI for bootstrapping and syncing agent workspaces. Interactive TUI with preset selection.',
    color: 'text-sky-500',
  },
  {
    icon: `<svg class="w-6 h-6" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="10"/><path d="M12 16v-4"/><path d="M12 8h.01"/><path d="M8 12h8"/><path d="M12 8v8"/></svg>`,
    title: '20+ Expert Skills',
    desc: 'Pre-built agents for Go, Nuxt, DevOps, TMA, CLI/TUI, and System Design. Battle-tested in production.',
    color: 'text-violet-500',
  },
  {
    icon: `<svg class="w-6 h-6" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M12 22s8-4 8-10V5l-8-3-8 3v7c0 6 8 10 8 10"/><path d="m9 12 2 2 4-4"/></svg>`,
    title: 'AST Enforcement',
    desc: 'Static analysis ensures Go Modern standards. No Any, No Log, No Globals — enforced at build time.',
    color: 'text-emerald-500',
  },
  {
    icon: `<svg class="w-6 h-6" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M21 12a9 9 0 1 1-9-9c2.52 0 4.93 1 6.74 2.74L21 8"/><path d="M21 3v5h-5"/></svg>`,
    title: 'Self-Evolution',
    desc: 'Meta-skills like @skill-updater allow the factory to upgrade its own blueprints autonomously.',
    color: 'text-amber-500',
  },
  {
    icon: `<svg class="w-6 h-6" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M13 2 3 14h9l-1 8 10-12h-9l1-8z"/></svg>`,
    title: 'Go 1.25 Native',
    desc: 'Zero-dependency architecture. Built on latest Go stdlib features including iter, unique, and optimized maps.',
    color: 'text-cyan-500',
  },
  {
    icon: `<svg class="w-6 h-6" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"/><polyline points="14 2 14 8 20 8"/><line x1="16" y1="13" x2="8" y2="13"/><line x1="16" y1="17" x2="8" y2="17"/></svg>`,
    title: 'Dual-Write Memory',
    desc: 'Two-phase commit protocol separating ephemeral thought (brain/) from permanent docs (project/).',
    color: 'text-rose-500',
  },
]
</script>

<template>
  <section class="py-24 bg-white dark:bg-slate-900 border-t border-slate-100 dark:border-slate-800/50">
    <div class="container-strict">
      <!-- Header -->
      <div class="text-center mb-16">
        <div class="inline-flex items-center gap-2 px-3 py-1 mb-4 rounded-full bg-slate-100 dark:bg-slate-800 text-slate-600 dark:text-slate-400 text-xs font-medium uppercase tracking-wider">
          <span class="w-1.5 h-1.5 rounded-full bg-emerald-500"></span>
          Core Capabilities
        </div>
        <h2 class="text-3xl lg:text-4xl font-bold text-slate-900 dark:text-white mb-4">
          Industrial-Grade Infrastructure
        </h2>
        <p class="text-lg text-slate-600 dark:text-slate-400 max-w-2xl mx-auto">
          We don't just provide prompts. We provide the entire operating system for building agentic software.
        </p>
      </div>

      <!-- Feature Grid -->
      <div class="grid md:grid-cols-2 lg:grid-cols-3 gap-6">
        <div 
          v-for="(feature, i) in features" 
          :key="i" 
          class="group relative p-6 rounded-2xl bg-slate-50 dark:bg-slate-800/40 border border-slate-200 dark:border-slate-700/50 hover:border-sky-400 dark:hover:border-sky-500 transition-all duration-300 hover:shadow-xl hover:shadow-sky-500/5 hover:-translate-y-1"
        >
          <!-- Icon -->
          <div 
            class="w-12 h-12 rounded-xl flex items-center justify-center mb-4 bg-white dark:bg-slate-800 border border-slate-200 dark:border-slate-700 group-hover:scale-110 transition-transform duration-300"
            :class="feature.color"
            v-html="feature.icon"
          ></div>
          
          <!-- Content -->
          <h3 class="text-lg font-bold text-slate-900 dark:text-white mb-2 group-hover:text-sky-600 dark:group-hover:text-sky-400 transition-colors">
            {{ feature.title }}
          </h3>
          <p class="text-slate-600 dark:text-slate-400 text-sm leading-relaxed">
            {{ feature.desc }}
          </p>
        </div>
      </div>
    </div>
  </section>
</template>
