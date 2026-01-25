<script setup lang="ts">
const cases = [
  {
    icon: `<svg class="w-8 h-8" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"><path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"/><circle cx="12" cy="7" r="4"/></svg>`,
    title: 'Solo Developers',
    desc: 'Full agentic stack for personal projects. No more reinventing the wheel — get TDD, Git protocols, and 20+ expert skills out of the box.',
    accent: 'from-violet-500 to-purple-600',
  },
  {
    icon: `<svg class="w-8 h-8" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"><path d="M13 2L3 14h9l-1 8 10-12h-9l1-8z"/></svg>`,
    title: 'Startups',
    desc: 'Consistent team practices from day one. Ship faster with pre-configured pipelines, code review standards, and architecture enforcement.',
    accent: 'from-sky-500 to-cyan-600',
  },
  {
    icon: `<svg class="w-8 h-8" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"><rect x="2" y="7" width="20" height="14" rx="2" ry="2"/><path d="M16 21V5a2 2 0 0 0-2-2h-4a2 2 0 0 0-2 2v16"/></svg>`,
    title: 'Enterprises',
    desc: 'Audit-ready skills with full traceability. ARTIFACT_REGISTRY, tech debt protocols, and dual-write patterns for compliance.',
    accent: 'from-emerald-500 to-teal-600',
  },
]
</script>

<template>
  <section class="py-24 bg-gradient-to-b from-white to-slate-50 dark:from-slate-900 dark:to-slate-950 border-t border-slate-100 dark:border-slate-800/50">
    <div class="container-strict">
      <!-- Header -->
      <div class="text-center mb-16">
        <div class="inline-flex items-center gap-2 px-3 py-1 mb-4 rounded-full bg-violet-100 dark:bg-violet-900/30 text-violet-600 dark:text-violet-400 text-xs font-medium uppercase tracking-wider">
          <span class="w-1.5 h-1.5 rounded-full bg-violet-500"></span>
          Who Is It For
        </div>
        <h2 class="text-3xl lg:text-4xl font-bold text-slate-900 dark:text-white mb-4">
          Built for Every Scale
        </h2>
        <p class="text-lg text-slate-600 dark:text-slate-400 max-w-2xl mx-auto">
          From side projects to enterprise deployments, Antigravity Factory adapts to your needs.
        </p>
      </div>

      <!-- Use Case Cards -->
      <div class="grid md:grid-cols-3 gap-8">
        <div 
          v-for="(item, i) in cases" 
          :key="i"
          class="group relative overflow-hidden rounded-2xl bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-800 hover:border-transparent transition-all duration-300 hover:shadow-2xl"
        >
          <!-- Gradient Top Bar -->
          <div 
            class="h-1 w-full bg-gradient-to-r"
            :class="item.accent"
          ></div>
          
          <!-- Content -->
          <div class="p-8">
            <!-- Icon -->
            <div 
              class="w-14 h-14 rounded-2xl flex items-center justify-center mb-6 bg-gradient-to-br text-white"
              :class="item.accent"
              v-html="item.icon"
            ></div>
            
            <!-- Title -->
            <h3 class="text-xl font-bold text-slate-900 dark:text-white mb-3">
              {{ item.title }}
            </h3>
            
            <!-- Description -->
            <p class="text-slate-600 dark:text-slate-400 leading-relaxed">
              {{ item.desc }}
            </p>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>
