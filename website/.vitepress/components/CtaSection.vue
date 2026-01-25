<script setup lang="ts">
import { withBase } from 'vitepress'
</script>

<template>
  <section class="relative py-24 overflow-hidden">
    <!-- Dark Gradient Background -->
    <div class="absolute inset-0 bg-gradient-to-br from-slate-900 via-slate-900 to-sky-950"></div>
    
    <!-- Subtle Grid Pattern -->
    <div class="absolute inset-0 opacity-[0.03]" style="background-image: url('data:image/svg+xml,<svg xmlns=&quot;http://www.w3.org/2000/svg&quot; viewBox=&quot;0 0 32 32&quot; width=&quot;32&quot; height=&quot;32&quot; fill=&quot;none&quot; stroke=&quot;white&quot;><path d=&quot;M0 .5H31.5V32&quot;/></svg>')"></div>
    
    <!-- Gradient Orbs -->
    <div class="absolute top-0 left-1/4 w-96 h-96 bg-sky-500/20 rounded-full blur-3xl"></div>
    <div class="absolute bottom-0 right-1/4 w-96 h-96 bg-violet-500/20 rounded-full blur-3xl"></div>
    
    <!-- Content -->
    <div class="relative container-strict text-center">
      <h2 class="text-4xl lg:text-5xl font-bold text-white mb-6 leading-tight">
        Ready to <span class="text-transparent bg-clip-text bg-gradient-to-r from-sky-400 to-cyan-300">Standardize</span>?
      </h2>
      <p class="text-lg text-slate-400 max-w-2xl mx-auto mb-10">
        Join the teams building the next generation of AI-powered software with production-grade agent infrastructure.
      </p>
      
      <!-- CTA Buttons -->
      <div class="flex flex-wrap justify-center gap-4 mb-12">
        <a 
          :href="withBase('/skills/')" 
          class="px-8 py-4 rounded-xl bg-sky-500 hover:bg-sky-400 text-white font-semibold shadow-lg shadow-sky-500/25 transition-all duration-300 hover:shadow-xl hover:shadow-sky-500/30 hover:-translate-y-0.5"
        >
          Explore Catalog
        </a>
        <a 
          href="https://github.com/ydnikolaev/antigravity-factory"
          class="px-8 py-4 rounded-xl bg-white/10 hover:bg-white/20 text-white font-semibold border border-white/20 transition-all duration-300 hover:-translate-y-0.5 backdrop-blur-sm"
        >
          View on GitHub
        </a>
      </div>
      
      <!-- Trust Badges -->
      <div class="flex flex-wrap justify-center gap-8 text-sm text-slate-500">
        <div class="flex items-center gap-2">
          <svg class="w-5 h-5 text-sky-400" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M13 2L3 14h9l-1 8 10-12h-9l1-8z"/></svg>
          <span>Go 1.25 Native</span>
        </div>
        <div class="flex items-center gap-2">
          <svg class="w-5 h-5 text-emerald-400" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M12 22s8-4 8-10V5l-8-3-8 3v7c0 6 8 10 8 10"/><path d="m9 12 2 2 4-4"/></svg>
          <span>100% Open Source</span>
        </div>
        <div class="flex items-center gap-2">
          <svg class="w-5 h-5 text-violet-400" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="10"/><path d="m9 12 2 2 4-4"/></svg>
          <span>MIT License</span>
        </div>
      </div>
    </div>
  </section>
</template>
