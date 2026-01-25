<script setup lang="ts">
import TerminalDemo from './TerminalDemo.vue'
import { withBase } from 'vitepress'
</script>

<template>
  <div class="relative w-full overflow-hidden bg-slate-50 dark:bg-slate-900 border-b border-slate-200 dark:border-slate-800 transition-colors duration-300">
    
    <!-- Subtle Grid Background -->
    <div class="absolute inset-0 bg-[url('/grid.svg')] opacity-[0.03] dark:opacity-[0.05]"></div>

    <div class="relative container-strict py-24 lg:py-32">
      <div class="grid lg:grid-cols-2 gap-16 items-center">
        
        <!-- Left: Text Content -->
        <div class="flex flex-col items-start text-left space-y-8">
          <!-- Tagline Badge -->
          <div class="inline-flex items-center gap-2 px-3 py-1 rounded-md bg-sky-100 dark:bg-sky-900/30 border border-sky-200 dark:border-sky-800 text-sky-700 dark:text-sky-300 text-xs font-mono font-medium tracking-wide uppercase">
            <span class="w-2 h-2 rounded-full bg-sky-500 animate-pulse"></span>
            Factory v3.0 • Production Ready
          </div>
          
          <h1 class="text-4xl lg:text-6xl font-bold tracking-tight text-slate-900 dark:text-white leading-[1.1]">
            The Unified Delivery Engine for 
            <span class="text-sky-600 dark:text-sky-400">Agent Infrastructure</span>
          </h1>
          
          <p class="text-lg text-slate-600 dark:text-slate-400 max-w-xl leading-relaxed">
            Deploy a complete autonomous ecosystem in seconds. 
            Includes <strong class="text-slate-900 dark:text-white">20+ Expert Skills</strong>, 
            <strong class="text-slate-900 dark:text-white">Factory CLI</strong>, and 
            <strong class="text-slate-900 dark:text-white">Architecture Enforcement</strong>.
            Powered by Go 1.25.
          </p>
          
          <div class="flex flex-wrap items-center gap-4 pt-2">
            <a :href="withBase('/skills/')" class="px-6 py-3 rounded-lg bg-sky-600 hover:bg-sky-700 text-white font-medium shadow-sm transition-all text-sm">
              Explore Catalog
            </a>
            <a href="https://github.com/ydnikolaev/antigravity-factory" class="px-6 py-3 rounded-lg bg-white dark:bg-slate-800 hover:bg-slate-50 dark:hover:bg-slate-700 text-slate-900 dark:text-white border border-slate-200 dark:border-slate-700 font-medium transition-all shadow-sm text-sm">
              GitHub Repository
            </a>
          </div>

          <!-- Quick Stats -->
          <div class="grid grid-cols-3 gap-8 pt-8 border-t border-slate-200 dark:border-slate-800 w-full">
            <div>
              <div class="text-2xl font-bold text-slate-900 dark:text-white font-mono">1.25</div>
              <div class="text-xs text-slate-500 dark:text-gray-500 uppercase tracking-wider mt-1">Go Native</div>
            </div>
            <div>
              <div class="text-2xl font-bold text-slate-900 dark:text-white font-mono">20+</div>
              <div class="text-xs text-slate-500 dark:text-gray-500 uppercase tracking-wider mt-1">Skills</div>
            </div>
             <div>
              <div class="text-2xl font-bold text-slate-900 dark:text-white font-mono">AST</div>
              <div class="text-xs text-slate-500 dark:text-gray-500 uppercase tracking-wider mt-1">Enforcement</div>
            </div>
          </div>
        </div>

        <!-- Right: Terminal Demo -->
        <div class="relative w-full max-w-2xl mx-auto lg:mx-0">
           <!-- Decorative backdrop -->
           <div class="absolute -inset-4 bg-gradient-to-tr from-sky-500/10 to-indigo-500/10 rounded-xl blur-xl"></div>
           <TerminalDemo />
        </div>

      </div>
    </div>
  </div>
</template>
