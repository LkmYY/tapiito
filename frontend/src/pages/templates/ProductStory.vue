<template>
  <div class="relative min-h-screen bg-[#f6f2ea] text-[#221d17]">
    <!-- Paper grain -->
    <div class="grain pointer-events-none absolute inset-0 opacity-[0.35]"></div>

    <div class="relative z-10 px-8 pb-20 pt-12">
      <!-- Meta -->
      <div class="rise flex items-center justify-between font-mono text-[10px] uppercase tracking-[0.26em] text-[#221d17]/40">
        <span>Story</span>
        <span v-if="nfcId">NFC · {{ nfcId }}</span>
      </div>

      <!-- Title -->
      <header class="rise rise-1 mt-14">
        <p v-if="payload.subtitle" class="font-mono text-[10px] uppercase tracking-[0.3em] text-[#b4552d]">{{ payload.subtitle }}</p>
        <h1 class="mt-5 text-balance font-serif text-[42px] leading-[1.08] tracking-tight">{{ payload.title || '未命名商品' }}</h1>
        <div class="mt-8 h-px w-full bg-[#221d17]/15"></div>
        <p class="mt-3 font-mono text-[9px] uppercase tracking-[0.26em] text-[#221d17]/35">A closer look</p>
      </header>

      <!-- Plate -->
      <figure class="rise rise-2 mt-10">
        <div class="relative aspect-[3/4] w-full overflow-hidden bg-[#e9e2d4]">
          <img
            v-if="hasImage"
            :src="payload.mainImage"
            alt="商品图片"
            class="h-full w-full object-cover"
            @error="imgError = true"
          />
          <div v-else class="flex h-full w-full flex-col items-center justify-center gap-4 bg-[linear-gradient(160deg,#efe9dc,#e0d7c4)]">
            <Package class="h-9 w-9 text-[#221d17]/20" />
            <p class="font-mono text-[9px] uppercase tracking-[0.3em] text-[#221d17]/30">Plate 01</p>
          </div>
        </div>
        <figcaption class="mt-3 flex items-center justify-between gap-4 font-mono text-[9px] uppercase tracking-[0.22em] text-[#221d17]/35">
          <span>Plate 01</span>
          <span class="truncate text-right">{{ payload.subtitle || 'Editorial' }}</span>
        </figcaption>
      </figure>

      <!-- Chapter 01 · Description -->
      <section v-if="payload.description" class="rise rise-2 mt-14">
        <div class="flex items-baseline gap-4">
          <span class="font-mono text-[10px] tabular-nums tracking-[0.2em] text-[#b4552d]">01</span>
          <h2 class="font-serif text-[22px] tracking-tight">关于这件作品</h2>
        </div>
        <p class="mt-5 whitespace-pre-line text-[14px] leading-[2.05] text-[#221d17]/70">{{ payload.description }}</p>
      </section>

      <!-- Chapter 02 · Features -->
      <section v-if="features.length" class="rise rise-3 mt-14">
        <div class="flex items-baseline gap-4">
          <span class="font-mono text-[10px] tabular-nums tracking-[0.2em] text-[#b4552d]">02</span>
          <h2 class="font-serif text-[22px] tracking-tight">细节之处</h2>
        </div>
        <div class="mt-6 border-t border-[#221d17]/15">
          <div v-for="(feature, index) in features" :key="index" class="flex gap-5 border-b border-[#221d17]/15 py-5">
            <span class="font-serif text-[20px] italic leading-none tabular-nums text-[#b4552d]/80">{{ index + 1 }}</span>
            <p class="flex-1 text-[14px] leading-[1.9] text-[#221d17]/75">{{ feature }}</p>
          </div>
        </div>
      </section>

      <!-- Chapter 03 · Specifications -->
      <section v-if="specs.length" class="rise rise-3 mt-14">
        <div class="flex items-baseline gap-4">
          <span class="font-mono text-[10px] tabular-nums tracking-[0.2em] text-[#b4552d]">03</span>
          <h2 class="font-serif text-[22px] tracking-tight">规格一览</h2>
        </div>
        <dl class="mt-6 border-t border-[#221d17]/15">
          <div v-for="[key, value] in specs" :key="key" class="flex items-baseline justify-between gap-6 border-b border-[#221d17]/15 py-4">
            <dt class="font-mono text-[10px] uppercase tracking-[0.2em] text-[#221d17]/40">{{ key }}</dt>
            <dd class="text-right text-[13px] font-medium tabular-nums text-[#221d17]/85">{{ value }}</dd>
          </div>
        </dl>
      </section>

      <!-- Epilogue · Price + CTA -->
      <section class="rise rise-3 mt-16">
        <div class="border border-[#221d17]/15 bg-[#fbf8f2] p-8 text-center">
          <p class="font-mono text-[9px] uppercase tracking-[0.3em] text-[#221d17]/35">The Edition</p>
          <div v-if="payload.price" class="mt-6 flex items-baseline justify-center gap-4">
            <span class="font-serif text-[36px] leading-none tracking-tight tabular-nums">¥{{ payload.price }}</span>
            <span v-if="payload.originalPrice" class="text-[14px] text-[#221d17]/35 line-through tabular-nums">¥{{ payload.originalPrice }}</span>
          </div>
          <button
            v-if="payload.ctaText"
            @click="handleCta"
            class="mt-7 flex w-full items-center justify-center gap-2.5 bg-[#221d17] py-4 text-[13px] font-medium tracking-[0.14em] text-[#f6f2ea] transition-all hover:bg-[#b4552d] active:scale-[0.99]"
          >
            {{ payload.ctaText }}
            <ArrowUpRight class="h-4 w-4" />
          </button>
        </div>
        <p class="mt-7 text-center font-mono text-[9px] uppercase tracking-[0.3em] text-[#221d17]/30">
          Fin<span v-if="nfcId"> · NFC {{ nfcId }}</span>
        </p>
      </section>
    </div>
  </div>
</template>

<script setup>
import { computed, ref } from 'vue'
import { ArrowUpRight, Package } from 'lucide-vue-next'

const props = defineProps(['payload', 'nfcId'])

const imgError = ref(false)
const hasImage = computed(() => Boolean(props.payload?.mainImage) && !imgError.value)
const features = computed(() => (Array.isArray(props.payload?.features) ? props.payload.features.filter(Boolean) : []))
const specs = computed(() => Object.entries(props.payload?.specifications || {}))

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
.grain {
  background-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' width='140' height='140'%3E%3Cfilter id='n'%3E%3CfeTurbulence type='fractalNoise' baseFrequency='0.85' numOctaves='2' stitchTiles='stitch'/%3E%3C/filter%3E%3Crect width='100%25' height='100%25' filter='url(%23n)' opacity='0.45'/%3E%3C/svg%3E");
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
