<template>
  <div class="relative min-h-screen bg-[#fbfbfa] text-stone-900">
    <!-- Ambient light -->
    <div class="pointer-events-none absolute inset-x-0 top-0 h-72 bg-[radial-gradient(70%_100%_at_50%_0%,rgba(5,150,105,0.07),transparent_72%)]"></div>

    <div class="relative z-10 px-6 pt-10">
      <!-- Meta -->
      <div class="rise flex items-center justify-between font-mono text-[10px] uppercase tracking-[0.24em] text-stone-400">
        <span>Product</span>
        <span v-if="nfcId">NFC · {{ nfcId }}</span>
      </div>

      <!-- Main image -->
      <div class="rise rise-1 mt-6">
        <div class="relative aspect-[4/5] w-full overflow-hidden rounded-[22px] bg-stone-100 shadow-[0_28px_60px_-32px_rgba(28,25,23,0.4)] ring-1 ring-stone-900/5">
          <img
            v-if="hasImage"
            :src="payload.mainImage"
            alt="商品图片"
            class="h-full w-full object-cover"
            @error="imgError = true"
          />
          <div v-else class="placeholder flex h-full w-full flex-col items-center justify-center gap-3">
            <Package class="h-9 w-9 text-stone-300" />
            <p class="font-mono text-[9px] uppercase tracking-[0.3em] text-stone-400">No Image</p>
          </div>
        </div>
      </div>

      <!-- Title -->
      <div class="rise rise-1 mt-7">
        <h1 class="text-balance text-[26px] font-semibold leading-[1.15] tracking-tight">{{ payload.title || '未命名商品' }}</h1>
        <p v-if="payload.subtitle" class="mt-2 text-[13px] tracking-wide text-stone-500">{{ payload.subtitle }}</p>
      </div>

      <!-- Price -->
      <div v-if="payload.price" class="rise rise-2 mt-6 flex items-baseline gap-3 border-t border-stone-200 pt-5">
        <span class="text-[30px] font-semibold leading-none tracking-tight tabular-nums">¥{{ payload.price }}</span>
        <span v-if="payload.originalPrice" class="text-[14px] text-stone-400 line-through tabular-nums">¥{{ payload.originalPrice }}</span>
        <span v-if="discount" class="ml-auto font-mono text-[10px] uppercase tracking-[0.16em] text-emerald-700">-{{ discount.percent }}%</span>
      </div>

      <!-- Description -->
      <section v-if="payload.description" class="rise rise-2 mt-9">
        <p class="font-mono text-[10px] uppercase tracking-[0.24em] text-stone-400">01 / 描述</p>
        <p class="mt-3 whitespace-pre-line text-[14px] leading-[1.85] text-stone-600">{{ payload.description }}</p>
      </section>

      <!-- Features -->
      <section v-if="features.length" class="rise rise-2 mt-10">
        <p class="font-mono text-[10px] uppercase tracking-[0.24em] text-stone-400">02 / 特点</p>
        <ul class="mt-4 border-t border-stone-200">
          <li v-for="(feature, index) in features" :key="index" class="flex items-start gap-3.5 border-b border-stone-200/80 py-4">
            <span class="mt-0.5 flex h-5 w-5 flex-shrink-0 items-center justify-center rounded-full bg-emerald-50 ring-1 ring-emerald-600/15">
              <Check class="h-3 w-3 text-emerald-600" />
            </span>
            <span class="flex-1 text-[14px] leading-relaxed text-stone-700">{{ feature }}</span>
          </li>
        </ul>
      </section>

      <!-- Specifications -->
      <section v-if="specs.length" class="rise rise-3 mt-10">
        <p class="font-mono text-[10px] uppercase tracking-[0.24em] text-stone-400">03 / 规格</p>
        <dl class="mt-4 grid grid-cols-2 gap-x-6">
          <div v-for="[key, value] in specs" :key="key" class="border-t border-stone-200 py-3.5">
            <dt class="truncate font-mono text-[10px] uppercase tracking-[0.14em] text-stone-400">{{ key }}</dt>
            <dd class="mt-1 text-[13px] font-medium text-stone-800 tabular-nums">{{ value }}</dd>
          </div>
        </dl>
      </section>
    </div>

    <!-- Sticky purchase bar -->
    <div v-if="payload.ctaText" class="sticky bottom-0 z-20 mt-12 border-t border-stone-200 bg-[#fbfbfa]">
      <div class="flex items-center gap-4 px-6 py-4">
        <div v-if="payload.price" class="min-w-0">
          <p class="font-mono text-[9px] uppercase tracking-[0.2em] text-stone-400">到手价</p>
          <p class="mt-0.5 text-[17px] font-semibold leading-none tabular-nums">¥{{ payload.price }}</p>
        </div>
        <button
          @click="handleCtaClick"
          class="ml-auto flex items-center gap-2 rounded-full bg-stone-900 px-7 py-3.5 text-sm font-semibold text-white transition-all hover:bg-stone-800 active:scale-[0.98]"
        >
          {{ payload.ctaText }}
          <ArrowUpRight class="h-4 w-4" />
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed, ref } from 'vue'
import { ArrowUpRight, Check, Package } from 'lucide-vue-next'

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
  return { percent: Math.round(((original - price) / original) * 100) }
})

/**
 * 安全处理 CTA 按钮点击，仅允许 http/https 协议的 URL 跳转
 */
function handleCtaClick() {
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
.placeholder {
  background-image:
    radial-gradient(rgba(28, 25, 23, 0.05) 1px, transparent 1px),
    linear-gradient(160deg, #f5f5f4, #e7e5e4);
  background-size: 14px 14px, 100% 100%;
}

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
