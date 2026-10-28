<template>
  <nav class="sticky top-0 z-40 w-full border-b border-[#14181d]/[0.07] bg-[#fbfaf7]/85 backdrop-blur-md">
    <div class="mx-auto flex h-[60px] max-w-[1600px] items-center px-3 sm:px-6">
      <router-link to="/" class="group mr-3 flex items-center gap-2.5 sm:mr-7">
        <span class="flex h-8 w-8 items-center justify-center rounded-[10px]  text-[#ff6b35] transition-transform duration-300 group-hover:-rotate-6 group-hover:scale-105 bg-[#15151B]">
          <img src="/tapiito-icon-white.png" alt="Tapiito" class="h-5 w-5" draggable="false" />
        </span>
        <span class="text-[18px] font-semibold tracking-tight text-[#14181d]">Tapiito</span>
      </router-link>

      <div class="flex flex-1 items-center gap-1">
        <router-link
          to="/admin/nfc"
          class="nav-item"
          :class="isActive('/admin/nfc') ? 'nav-item-active' : ''"
          aria-label="标签管理"
        >
          <Tags class="h-3.5 w-3.5" />
          <span class="hidden lg:inline">标签管理</span>
        </router-link>
        <router-link
          to="/admin/orders"
          class="nav-item"
          :class="isActive('/admin/orders') ? 'nav-item-active' : ''"
          aria-label="订单管理"
        >
          <ClipboardList class="h-3.5 w-3.5" />
          <span class="hidden lg:inline">订单管理</span>
        </router-link>
        <router-link
          to="/admin/stats"
          class="nav-item"
          :class="isActive('/admin/stats') ? 'nav-item-active' : ''"
          aria-label="访问统计"
        >
          <BarChart2 class="h-3.5 w-3.5" />
          <span class="hidden lg:inline">访问统计</span>
        </router-link>
        <router-link
          to="/admin/feedback"
          class="nav-item"
          :class="isActive('/admin/feedback') ? 'nav-item-active' : ''"
          aria-label="用户意见"
        >
          <MessageSquare class="h-3.5 w-3.5" />
          <span class="hidden lg:inline">用户意见</span>
        </router-link>
      </div>

      <div class="flex items-center gap-2">
        <router-link to="/feedback" class="nav-about hidden lg:inline-flex" aria-label="我有意见">
          <MessageSquare class="h-3.5 w-3.5" />
          我有意见
        </router-link>
        <router-link to="/about" class="nav-about hidden lg:inline-flex" aria-label="关于我们">
          <Info class="h-3.5 w-3.5" />
          关于我们
        </router-link>
        <span class="hidden items-center gap-2 rounded-full border border-[#14181d]/[0.08] bg-white/60 px-3 py-1.5 sm:inline-flex">
          <span class="h-1.5 w-1.5 rounded-full " />
          <span class="text-[12px] font-medium text-[#14181d]/60">{{ authStore.username }}</span>
        </span>
        <button
          @click="authStore.logout()"
          aria-label="退出登录"
          class="inline-flex items-center gap-1.5 rounded-full px-3 py-1.5 text-[12.5px] font-medium text-[#14181d]/45 transition-colors duration-200 hover:bg-[#ff6b35]/[0.08] hover:text-[#d94f1a]"
        >
          <LogOut class="h-3.5 w-3.5" />
          <span class="hidden sm:inline">退出</span>
        </button>
      </div>
    </div>
  </nav>
</template>

<script setup>
import { useRoute } from 'vue-router'
import { Nfc, Tags, BarChart2, ClipboardList, LogOut, Info, MessageSquare } from 'lucide-vue-next'
import { useAuthStore } from '../stores/auth'

const route = useRoute()
const authStore = useAuthStore()

const isActive = (path) => route.path.startsWith(path)
</script>

<style scoped>
.nav-item {
  display: inline-flex;
  align-items: center;
  gap: 0.4rem;
  border-radius: 9999px;
  padding: 0.45rem 0.85rem;
  font-size: 13px;
  font-weight: 500;
  color: rgba(20, 24, 29, 0.55);
  transition: background-color 0.18s ease, color 0.18s ease;
}
.nav-item:hover {
  background: rgba(20, 24, 29, 0.045);
  color: #14181d;
}
.nav-item-active {
  background: #14181d;
  color: #fff;
}
.nav-item-active:hover {
  background: #14181d;
  color: #fff;
}
.nav-about {
  align-items: center;
  gap: 0.35rem;
  height: 32px;
  padding: 0 0.7rem;
  border-radius: 999px;
  font-size: 12.5px;
  font-weight: 500;
  color: rgba(20, 24, 29, 0.5);
  transition: background-color 0.18s ease, color 0.18s ease;
}
.nav-about:hover {
  background: rgba(20, 24, 29, 0.045);
  color: #14181d;
}

@media (max-width: 1023px) {
  .nav-item {
    padding: 0.45rem 0.6rem;
  }
}
</style>
