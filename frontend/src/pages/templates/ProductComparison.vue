<template>
  <div class="relative min-h-screen bg-[#f4f6f8] text-slate-900">
    <div class="relative z-10 px-6 pb-16 pt-10">
      <!-- Meta -->
      <div class="rise flex items-center justify-between font-mono text-[10px] uppercase tracking-[0.24em] text-slate-400">
        <span>Spec Sheet</span>
        <span v-if="nfcId">NFC · {{ nfcId }}</span>
      </div>

      <!-- Header -->
      <div class="rise rise-1 mt-6 flex items-center gap-5">
        <div class="relative h-24 w-24 flex-shrink-0 overflow-hidden rounded-xl bg-slate-200 ring-1 ring-slate-900/5">
          <img
            v-if="hasImage"
            :src="payload.mainImage"
            alt="商品图片"
            class="h-full w-full object-cover"
            @error="imgError = true"
          />
          <div v-else class="flex h-full w-full items-center justify-center bg-[linear-gradient(140deg,#e2e8f0,#cbd5e1)]">
            <Package class="h-6 w-6 text-slate-400" />
          </div>
        </div>
        <div class="min-w-0 flex-1">
          <h1 class="text-balance text-[22px] font-semibold leading-snug tracking-tight">{{ payload.title || '未命名商品' }}</h1>
          <p v-if="payload.subtitle" class="mt-1.5 truncate text-[12px] text-slate-500">{{ payload.subtitle }}</p>
        </div>
      </div>

      <!-- Description -->
      <p v-if="payload.description" class="rise rise-2 mt-7 whitespace-pre-line text-[13px] leading-[1.9] text-slate-500">{{ payload.description }}</p>

      <!-- Price panel -->
      <div v-if="payload.price" class="rise rise-2 mt-7 rounded-2xl bg-white p-5 shadow-[0_20px_45px_-30px_rgba(15,23,42,0.45)] ring-1 ring-slate-900/5">
        <div class="flex items-end justify-between gap-4">
          <div>
            <p class="font-mono text-[9px] uppercase tracking-[0.22em] text-slate-400">Price</p>
            <div class="mt-2 flex items-baseline gap-3">
              <span class="text-[28px] font-semibold leading-none tracking-tight tabular-nums">¥{{ payload.price }}</span>
              <span v-if="payload.originalPrice" class="text-[13px] text-slate-400 line-through tabular-nums">¥{{ payload.originalPrice }}</span>
            </div>
          </div>
          <div v-if="discount" class="rounded-xl bg-blue-600/[0.07] px-3.5 py-2.5 text-right ring-1 ring-blue-600/15">
            <p class="font-mono text-[9px] uppercase tracking-[0.16em] text-blue-700">Save</p>
            <p class="mt-1 text-[13px] font-semibold leading-none tabular-nums text-blue-700">¥{{ discount.amount }}</p>
          </div>
        </div>
      </div>

      <!-- Included features -->
      <section v-if="features.length" class="rise rise-2 mt-8">
        <div class="flex items-baseline justify-between">
          <h2 class="font-mono text-[10px] uppercase tracking-[0.24em] text-slate-400">Included · 本款标配</h2>
          <span class="font-mono text-[10px] tabular-nums text-slate-400">{{ features.length }} 项</span>
        </div>
        <ul class="mt-3 overflow-hidden rounded-2xl bg-white ring-1 ring-slate-900/5">
          <li
            v-for="(feature, index) in features"
            :key="index"
            class="flex items-center gap-3.5 px-5 py-3.5"
            :class="index !== 0 ? 'border-t border-slate-100' : ''"
          >
            <span class="flex h-5 w-5 flex-shrink-0 items-center justify-center rounded-md bg-blue-600/10">
              <Check class="h-3 w-3 text-blue-700" />
            </span>
            <span class="flex-1 text-[13px] text-slate-700">{{ feature }}</span>
          </li>
        </ul>
      </section>

      <!-- Specification table -->
      <section v-if="specs.length" class="rise rise-3 mt-8">
        <h2 class="font-mono text-[10px] uppercase tracking-[0.24em] text-slate-400">Specification</h2>
        <div class="mt-3 overflow-hidden rounded-2xl bg-white ring-1 ring-slate-900/5">
          <div class="grid grid-cols-[1fr_1.25fr] items-center bg-slate-900 px-5 py-3">
            <span class="font-mono text-[9px] uppercase tracking-[0.2em] text-slate-400">项目</span>
            <span class="flex items-center gap-2 font-mono text-[9px] uppercase tracking-[0.2em] text-white">
              <span class="h-1.5 w-1.5 rounded-full bg-blue-400"></span>
              本款
            </span>
          </div>
          <div
            v-for="([key, value], index) in specs"
            :key="key"
            class="grid grid-cols-[1fr_1.25fr] items-center px-5 py-3.5"
            :class="index !== 0 ? 'border-t border-slate-100' : ''"
          >
            <span class="text-[12px] text-slate-500">{{ key }}</span>
            <span class="text-[13px] font-medium tabular-nums text-slate-900">{{ value }}</span>
          </div>
        </div>
      </section>

      <!-- CTA -->
      <button
        v-if="payload.ctaText"
        @click="handleCta"
        class="rise rise-3 mt-9 flex w-full items-center justify-center gap-2.5 rounded-2xl bg-blue-600 py-4 text-[14px] font-semibold text-white shadow-[0_18px_35px_-16px_rgba(37,99,235,0.55)] transition-all hover:bg-blue-700 active:scale-[0.98]"
      >
        <ShoppingBag class="h-4 w-4" />
        {{ payload.ctaText }}
      </button>
    </div>
  </div>
</template>

<script setup>
import { computed, ref } from 'vue'
import { Check, Package, ShoppingBag } from 'lucide-vue-next'

const props = defineProps(['payload', 'nfcId'])

const imgError = ref(false)
const hasImage = computed(() => Boolean(props.payload?.mainImage) && !imgError.value)
const features = computed(() => (Array.isArray(props.payload?.features) ? props.payload.features.filter(Boolean) : []))
const specs = computed(() => Object.entries(props.payload?.specifications || {}))

const discount = computed(() => {
  if (!props.payload?.price || !props.payload?.originalPrice) return null
  const price = Number(props.payload.price)
  const original = Number(props.payload.originalPrice)
  if (!Number.isFinite(price) || !Number.isFinite(original) || original <= price) return null
  return {
    amount: Math.round((original - price) * 100) / 100,
    percent: Math.round(((original - price) / original) * 100)
  }
})

/**
 * 安全处理 CTA 按钮点击，仅允许 http/https 协议的 URL 跳转
 */
function handleCta() {
  const action = props.payload?.ctaAction
  if (!action) return
  try {
    const url = new URL(action)
    if (['http:', 'https:'].includes(url.protocol)) {
      window.location.href = url.href
    }
  } catch {
    // 非合法 URL，忽略
  }
}
</script>

<style scoped>
@keyframes rise {
  from {
    opacity: 0;
    transform: translateY(14px);
  }
  to {
    opacity: 1;
    transform: none;
  }
}
.rise {
  animation: rise 0.65s cubic-bezier(0.22, 1, 0.36, 1) both;
}
.rise-1 {
  animation-delay: 0.08s;
}
.rise-2 {
  animation-delay: 0.16s;
}
.rise-3 {
  animation-delay: 0.24s;
}
</style>
