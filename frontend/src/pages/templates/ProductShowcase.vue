<template>
  <div class="relative min-h-screen bg-[#0e0e0f] text-white">
    <!-- Stage -->
    <div class="rise relative h-[72vh] min-h-[420px] w-full overflow-hidden">
      <img
        v-if="hasImage"
        :src="payload.mainImage"
        alt="商品图片"
        class="h-full w-full object-cover"
        @error="imgError = true"
      />
      <div v-else class="stage-ph flex h-full w-full flex-col items-center justify-center gap-5">
        <Package class="h-12 w-12 text-white/20" />
        <p class="font-mono text-[9px] uppercase tracking-[0.34em] text-white/30">No Image</p>
      </div>
      <!-- Vignette -->
      <div class="pointer-events-none absolute inset-0 bg-[linear-gradient(180deg,rgba(14,14,15,0.4)_0%,transparent_28%,transparent_52%,#0e0e0f_100%)]"></div>

      <!-- Stage meta -->
      <div class="absolute inset-x-0 top-0 flex items-center justify-between px-7 pt-10 font-mono text-[10px] uppercase tracking-[0.28em] text-white/50">
        <span>Showcase</span>
        <span v-if="nfcId">NFC · {{ nfcId }}</span>
      </div>

      <!-- Overlay title -->
      <div class="absolute inset-x-0 bottom-24 px-7">
        <p v-if="payload.subtitle" class="font-mono text-[10px] uppercase tracking-[0.3em] text-white/60">{{ payload.subtitle }}</p>
        <h1 class="mt-3 text-balance text-[36px] font-semibold leading-[1.05] tracking-tight">{{ payload.title || '未命名商品' }}</h1>
      </div>
    </div>

    <!-- Bottom sheet -->
    <div class="rise rise-1 relative z-10 -mt-12 rounded-t-[28px] bg-[#151517] px-7 pb-12 pt-4 ring-1 ring-white/10">
      <div class="mx-auto h-1 w-10 rounded-full bg-white/15"></div>

      <!-- Price -->
      <div class="mt-7 flex items-end justify-between gap-5">
        <div v-if="payload.price">
          <p class="font-mono text-[9px] uppercase tracking-[0.24em] text-white/35">Price</p>
          <div class="mt-2.5 flex items-baseline gap-3">
            <span class="text-[30px] font-semibold leading-none tracking-tight tabular-nums">¥{{ payload.price }}</span>
            <span v-if="payload.originalPrice" class="text-[13px] text-white/30 line-through tabular-nums">¥{{ payload.originalPrice }}</span>
          </div>
        </div>
        <span v-if="discount" class="mb-1 rounded-full px-3 py-1.5 font-mono text-[10px] uppercase tracking-[0.18em] text-white/70 ring-1 ring-white/20">
          -{{ discount.percent }}%
        </span>
      </div>

      <!-- Description -->
      <p v-if="payload.description" class="mt-7 whitespace-pre-line text-[13px] leading-[1.95] text-white/55">{{ payload.description }}</p>

      <!-- Features -->
      <div v-if="features.length" class="mt-9 border-t border-white/10">
        <div v-for="(feature, index) in features" :key="index" class="flex items-start gap-4 border-b border-white/10 py-4">
          <span class="pt-0.5 font-mono text-[10px] tabular-nums text-white/30">{{ String(index + 1).padStart(2, '0') }}</span>
          <span class="flex-1 text-[13px] leading-relaxed text-white/75">{{ feature }}</span>
        </div>
      </div>

      <!-- Specifications -->
      <div v-if="specs.length" class="mt-9">
        <p class="font-mono text-[9px] uppercase tracking-[0.26em] text-white/35">Spec</p>
        <dl class="mt-3 grid grid-cols-2 gap-x-6">
          <div v-for="[key, value] in specs" :key="key" class="border-t border-white/10 py-3">
            <dt class="truncate font-mono text-[9px] uppercase tracking-[0.16em] text-white/35">{{ key }}</dt>
            <dd class="mt-1 text-[12px] tabular-nums text-white/80">{{ value }}</dd>
          </div>
        </dl>
      </div>

      <!-- CTA -->
      <button
        v-if="payload.ctaText"
        @click="handleCta"
        class="mt-10 flex w-full items-center justify-center gap-2.5 rounded-full bg-white py-4 text-[14px] font-semibold text-[#0e0e0f] transition-all hover:bg-white/90 active:scale-[0.98]"
      >
        <ShoppingBag class="h-4 w-4" />
        {{ payload.ctaText }}
      </button>
      <p class="mt-6 text-center font-mono text-[9px] uppercase tracking-[0.28em] text-white/25">Scan · View · Own</p>
    </div>
  </div>
</template>

<script setup>
import { computed, ref } from 'vue'
import { Package, ShoppingBag } from 'lucide-vue-next'

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
.stage-ph {
  background-image:
    linear-gradient(rgba(255, 255, 255, 0.035) 1px, transparent 1px),
    linear-gradient(90deg, rgba(255, 255, 255, 0.035) 1px, transparent 1px),
    linear-gradient(160deg, #17171a, #101012);
  background-size: 44px 44px, 44px 44px, 100% 100%;
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
