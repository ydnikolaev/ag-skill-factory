<script setup lang="ts">
const steps = [
  {
    step: '01',
    title: 'Install CLI',
    code: 'go install github.com/ydnikolaev/antigravity-factory@latest',
    desc: 'One command. Zero dependencies. Works anywhere Go runs.',
  },
  {
    step: '02',
    title: 'Select Preset',
    code: 'factory install --preset=backend',
    desc: 'Choose from 6 curated bundles: all, core, backend, frontend, tma, cli.',
  },
  {
    step: '03',
    title: 'Build Anything',
    code: 'Your agents are ready.',
    desc: '20+ skills, TDD protocols, team structure, and architecture tests — deployed to .agent/',
  },
]
</script>

<template>
  <section class="py-24 bg-slate-50 dark:bg-slate-950 border-t border-slate-200 dark:border-slate-800">
    <div class="container-strict">
      <!-- Header -->
      <div class="text-center mb-16">
        <div class="inline-flex items-center gap-2 px-3 py-1 mb-4 rounded-full bg-sky-100 dark:bg-sky-900/30 text-sky-600 dark:text-sky-400 text-xs font-medium uppercase tracking-wider">
          <span class="w-1.5 h-1.5 rounded-full bg-sky-500"></span>
          Quick Start
        </div>
        <h2 class="text-3xl lg:text-4xl font-bold text-slate-900 dark:text-white mb-4">
          Three Steps to Production
        </h2>
        <p class="text-lg text-slate-600 dark:text-slate-400 max-w-2xl mx-auto">
          From zero to full agentic infrastructure in under 60 seconds.
        </p>
      </div>

      <!-- Steps -->
      <div class="grid md:grid-cols-3 gap-8">
        <div 
          v-for="(item, i) in steps" 
          :key="i"
          class="relative"
        >
          <!-- Connector Line (hidden on mobile, visible on desktop) -->
          <div 
            v-if="i < steps.length - 1" 
            class="hidden md:block absolute top-8 left-full w-full h-0.5 bg-gradient-to-r from-sky-500 to-sky-500/0"
          ></div>

          <!-- Step Card -->
          <div class="group relative p-6 rounded-2xl bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-800 hover:border-sky-400 dark:hover:border-sky-500 transition-all duration-300 hover:shadow-lg">
            <!-- Step Number -->
            <div class="inline-flex items-center justify-center w-10 h-10 rounded-xl bg-sky-500 text-white font-bold text-sm mb-4">
              {{ item.step }}
            </div>
            
            <!-- Title -->
            <h3 class="text-xl font-bold text-slate-900 dark:text-white mb-3">
              {{ item.title }}
            </h3>
            
            <!-- Code Block -->
            <div class="rounded-lg bg-slate-900 dark:bg-slate-950 px-4 py-3 mb-4 overflow-x-auto">
              <code class="text-sm text-sky-400 font-mono whitespace-nowrap">{{ item.code }}</code>
            </div>
            
            <!-- Description -->
            <p class="text-sm text-slate-600 dark:text-slate-400 leading-relaxed">
              {{ item.desc }}
            </p>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>
