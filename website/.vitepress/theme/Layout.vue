<script setup lang="ts">
import DefaultTheme from 'vitepress/theme'
import { useData } from 'vitepress'
import HomeHero from '../components/HomeHero.vue'

const { Layout } = DefaultTheme
const { frontmatter } = useData()
</script>

<template>
  <Layout>
    <template #home-hero-before>
      <div v-if="frontmatter.customHero" class="relative z-10 w-full overflow-hidden bg-dark-bg">
        <HomeHero />
      </div>
    </template>
    
    <!-- Hide Default Hero if Custom is Active -->
    <template #home-hero-after v-if="frontmatter.customHero" />
  </Layout>
</template>
