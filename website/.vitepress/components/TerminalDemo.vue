<script setup lang="ts">
import { ref, onMounted } from 'vue'

const lines = [
  { text: '→ ~ > factory install', class: 'font-bold text-slate-100' },
  { text: '', class: 'h-1' },
  { text: '  Select preset to install:', class: 'text-sky-400 font-bold' },
  { text: '> all        - Full blueprint (21 skills)', class: 'text-sky-300' },
  { text: '  core       - Pipeline essentials', class: 'text-slate-500' },
  { text: '  backend    - Go backend development', class: 'text-slate-500' },
  { text: '  frontend   - Nuxt/Vue frontend', class: 'text-slate-500' },
  { text: '  tma        - Telegram Mini Apps', class: 'text-slate-500' },
  { text: '  cli        - CLI/TUI applications', class: 'text-slate-500' },
  { text: '', class: 'h-1' },
  { text: '• 21 skills selected', class: 'text-emerald-400 text-xs' },
  { text: '• 5 workflows included', class: 'text-emerald-400 text-xs' },
]

const visibleLines = ref<typeof lines>([])
const showPrompt = ref(false)
const isTyping = ref(false)

onMounted(() => {
  let i = 0
  isTyping.value = true
  
  const interval = setInterval(() => {
    if (i < lines.length) {
      visibleLines.value.push(lines[i])
      i++
    } else {
      clearInterval(interval)
      isTyping.value = false
      showPrompt.value = true
    }
  }, 120)
})
</script>

<template>
  <div class="relative w-full rounded-xl overflow-hidden shadow-2xl bg-[#0d1117] font-mono text-sm leading-relaxed border border-slate-700/50 select-none">
    <!-- Mac Window Header -->
    <div class="flex items-center justify-between px-4 py-3 bg-[#161b22] border-b border-slate-800/80">
      <div class="flex items-center gap-2">
        <div class="w-3 h-3 rounded-full bg-[#ff5f56] shadow-inner"></div>
        <div class="w-3 h-3 rounded-full bg-[#ffbd2e] shadow-inner"></div>
        <div class="w-3 h-3 rounded-full bg-[#27c93f] shadow-inner"></div>
      </div>
      <div class="text-xs text-slate-500 font-medium tracking-wide">factory — zsh</div>
      <div class="w-12"></div>
    </div>
    
    <!-- Terminal Content -->
    <div class="p-6 min-h-[340px] text-left bg-gradient-to-b from-[#0d1117] to-[#161b22]">
      <TransitionGroup name="line" tag="div">
        <div 
          v-for="(line, i) in visibleLines" 
          :key="i" 
          class="mb-1 whitespace-pre-wrap transition-all duration-200" 
          :class="line.class"
        >
          {{ line.text }}
        </div>
      </TransitionGroup>
      
      <!-- Blinking Prompt -->
      <Transition name="fade">
        <div v-if="showPrompt" class="mt-2 flex items-center">
          <span class="text-sky-500 mr-2">?</span>
          <span class="text-slate-300">Confirm installation? [Y/n]</span>
          <span class="inline-block w-2.5 h-5 ml-1 bg-sky-400 animate-pulse rounded-sm"></span>
        </div>
      </Transition>
      
      <!-- Typing indicator -->
      <div v-if="isTyping && visibleLines.length > 0" class="mt-2 flex items-center gap-1">
        <span class="w-1.5 h-1.5 bg-sky-400 rounded-full animate-bounce" style="animation-delay: 0ms"></span>
        <span class="w-1.5 h-1.5 bg-sky-400 rounded-full animate-bounce" style="animation-delay: 150ms"></span>
        <span class="w-1.5 h-1.5 bg-sky-400 rounded-full animate-bounce" style="animation-delay: 300ms"></span>
      </div>
    </div>
  </div>
</template>

<style scoped>
.line-enter-active {
  transition: all 0.2s ease-out;
}
.line-enter-from {
  opacity: 0;
  transform: translateX(-10px);
}

.fade-enter-active {
  transition: all 0.3s ease-out;
}
.fade-enter-from {
  opacity: 0;
}
</style>
