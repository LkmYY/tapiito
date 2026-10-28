<template>
  <header
    class="z-40 w-full transition-colors duration-300"
    :class="transparent
      ? 'absolute inset-x-0 top-0'
      : 'sticky top-0 border-b border-[#14181d]/[0.07] bg-[#fbfaf7]/85 backdrop-blur-md'"
  >
    <div class="mx-auto flex h-16 max-w-[1400px] items-center justify-between px-4 sm:px-6">
      <!-- 品牌 -->
      <router-link to="/" class="group flex shrink-0 items-center gap-2.5">
        <span
          class="flex h-8 w-8 items-center justify-center rounded-[10px] bg-[#15151B] transition-transform duration-300 group-hover:-rotate-6 group-hover:scale-105"
        >
          <img src="/tapiito-icon-white.png" alt="Tapiito" class="h-5 w-5" draggable="false" />
        </span>
        <span class="text-[18px] font-semibold tracking-[-0.02em] text-[#14181d]">Tapiito</span>
      </router-link>

      <!-- 桌面菜单 -->
      <nav class="hidden items-center gap-1 sm:flex">
        <router-link v-if="backTo" :to="backTo" class="nav-item">
          <ArrowLeft class="h-3.5 w-3.5" />
          {{ backLabel }}
        </router-link>
        <router-link to="/about" class="nav-item">关于我们</router-link>
        <router-link to="/feedback" class="nav-item">我有意见</router-link>
        <router-link :to="ctaTo" class="nav-cta">
          {{ ctaLabel }}
          <ArrowRight class="h-3.5 w-3.5" />
        </router-link>
      </nav>

      <!-- 移动端：文字 + 箭头，不用裸图标 -->
      <button
        class="mobile-trigger sm:hidden"
        :aria-expanded="open"
        aria-controls="site-menu"
        @click="open = !open"
      >
        <span>{{ open ? '收起' : '菜单' }}</span>
        <ChevronDown class="h-4 w-4 transition-transform duration-200" :class="open ? 'rotate-180' : ''" />
      </button>
    </div>

    <!-- 移动端下拉面板 -->
    <Transition name="menu-drop">
      <div v-if="open" id="site-menu" class="mobile-panel sm:hidden">
        <router-link v-if="backTo" :to="backTo" class="mobile-item" @click="close">
          <ArrowLeft class="h-4 w-4 opacity-60" />
          {{ backLabel }}
        </router-link>
        <router-link to="/about" class="mobile-item" @click="close">
          <Info class="h-4 w-4 opacity-60" />
          关于我们
        </router-link>
        <router-link to="/feedback" class="mobile-item" @click="close">
          <MessageSquare class="h-4 w-4 opacity-60" />
          我有意见
        </router-link>
        <router-link :to="ctaTo" class="mobile-item mobile-item-cta" @click="close">
          <ShoppingBag class="h-4 w-4" />
          {{ ctaLabel }}
        </router-link>
      </div>
    </Transition>
  </header>
</template>

<script setup>
import { ref, watch } from 'vue'
import { useRoute } from 'vue-router'
import { ArrowLeft, ArrowRight, ChevronDown, Info, MessageSquare, ShoppingBag } from 'lucide-vue-next'

const props = defineProps({
  /** 首页 hero 上用绝对定位透明头部，其余页面用 sticky 白底 */
  transparent: { type: Boolean, default: false },
  /** 可选返回入口，如 /showroom */
  backTo: { type: String, default: '' },
  backLabel: { type: String, default: '返回' },
  ctaLabel: { type: String, default: '预约咨询' },
  /** CTA 目标，展厅页用它指向下单页 */
  ctaTo: { type: String, default: '/consult' },
})

const open = ref(false)
const route = useRoute()
const close = () => { open.value = false }
// 路由变化时收起，避免切页后面板还挂着
watch(() => route.fullPath, close)
</script>

<style scoped>
.nav-item {
  display: inline-flex;
  align-items: center;
  gap: 0.375rem;
  padding: 0.4rem 0.75rem;
  border-radius: 9px;
  font-size: 13.5px;
  color: rgba(20, 24, 29, 0.62);
  transition: color 0.2s ease, background-color 0.2s ease;
}
.nav-item:hover {
  color: #14181d;
  background: rgba(20, 24, 29, 0.04);
}
.nav-cta {
  display: inline-flex;
  align-items: center;
  gap: 0.375rem;
  margin-left: 0.375rem;
  padding: 0.5rem 0.95rem;
  border-radius: 9px;
  background: #15151b;
  color: #fff;
  font-size: 13.5px;
  font-weight: 500;
  transition: opacity 0.2s ease, transform 0.15s ease;
}
.nav-cta:hover {
  opacity: 0.9;
}
.nav-cta:active {
  transform: scale(0.98);
}

/* 移动端触发按钮：带文字，避免用户看不懂纯图标 */
.mobile-trigger {
  display: inline-flex;
  align-items: center;
  gap: 0.3rem;
  padding: 0.4rem 0.7rem;
  border-radius: 9px;
  border: 1px solid rgba(20, 24, 29, 0.12);
  background: #fff;
  font-size: 13px;
  color: #14181d;
}

.mobile-panel {
  border-top: 1px solid rgba(20, 24, 29, 0.07);
  background: #fbfaf7;
  padding: 0.5rem 1rem 0.85rem;
}
.mobile-item {
  display: flex;
  align-items: center;
  gap: 0.6rem;
  padding: 0.7rem 0.5rem;
  border-radius: 9px;
  font-size: 14.5px;
  color: #14181d;
}
.mobile-item:active {
  background: rgba(20, 24, 29, 0.05);
}
.mobile-item-cta {
  margin-top: 0.35rem;
  background: #15151b;
  color: #fff;
  justify-content: center;
  font-weight: 500;
}

/* scoped 的 display 优先级高于 Tailwind 的 sm:hidden，桌面上显式隐藏 */
@media (min-width: 640px) {
  .mobile-trigger,
  .mobile-panel {
    display: none !important;
  }
}

.menu-drop-enter-active,
.menu-drop-leave-active {
  transition: opacity 0.18s ease, transform 0.18s ease;
}
.menu-drop-enter-from,
.menu-drop-leave-to {
  opacity: 0;
  transform: translateY(-6px);
}
</style>
