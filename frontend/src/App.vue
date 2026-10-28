<template>
  <div :class="{ dark: isDark }">
    <AppNavbar v-if="showNavbar" />
    <div class="app-container">
      <router-view />
    </div>
    <button
      v-if="showThemeToggle"
      @click="toggleDark"
      class="fixed bottom-6 right-6 z-50 w-10 h-10 rounded-full shadow-lg flex items-center justify-center transition-colors bg-gray-200 dark:bg-gray-700 text-gray-800 dark:text-gray-200 hover:scale-110"
      aria-label="切换暗色模式"
    >
      <Sun v-if="isDark" class="w-5 h-5" />
      <Moon v-else class="w-5 h-5" />
    </button>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { Sun, Moon } from 'lucide-vue-next'
import { useAuthStore } from './stores/auth'
import AppNavbar from './components/AppNavbar.vue'

const route = useRoute()
const authStore = useAuthStore()

const isPublicPage = computed(() => {
  const p = route.path || ''
  return p === '/' || p === '/about' || p === '/privacy' || p === '/disclaimer' || p === '/feedback' || p === '/login' || p === '/preview' || p === '/showroom' || p === '/order' || p.startsWith('/p/')
})

const showNavbar = computed(() => authStore.isLoggedIn && !isPublicPage.value)

const showThemeToggle = computed(() => !isPublicPage.value)

const isDark = ref(false)

onMounted(() => {
  isDark.value = localStorage.getItem('theme') === 'dark'
  document.documentElement.classList.toggle('dark', isDark.value)
})

const toggleDark = () => {
  isDark.value = !isDark.value
  localStorage.setItem('theme', isDark.value ? 'dark' : 'light')
  document.documentElement.classList.toggle('dark', isDark.value)
}
</script>
