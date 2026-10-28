<template>
  <div class="relative min-h-screen bg-[#0b0a08] text-[#f4efe6]">
    <!-- Ambient gold light -->
    <div class="pointer-events-none absolute inset-x-0 top-0 h-[420px] bg-[radial-gradient(60%_70%_at_50%_0%,rgba(200,164,95,0.13),transparent_72%)]"></div>
    <!-- Print frame -->
    <div class="pointer-events-none absolute inset-3 border border-[#c8a45f]/15"></div>
    <!-- Film grain -->
    <div class="grain pointer-events-none absolute inset-0 opacity-[0.28]"></div>

    <div class="relative z-10 px-9 pb-16 pt-14">
      <!-- Meta -->
      <div class="rise flex items-center justify-between font-mono text-[9px] uppercase tracking-[0.34em] text-[#c8a45f]/70">
        <span>Atelier</span>
        <span v-if="nfcId">No. {{ nfcId }}</span>
      </div>

      <!-- Framed image -->
      <div class="rise rise-1 mt-12">
        <div class="relative aspect-[4/5] w-full overflow-hidden bg-[#141210] ring-1 ring-[#c8a45f]/25">
          <img
            v-if="hasImage"
            :src="payload.mainImage"
            alt="商品图片"
            class="h-full w-full object-cover"
            @error="imgError = true"
          />
          <div v-else class="relative flex h-full w-full items-center justify-center bg-[linear-gradient(160deg,#16130f,#0d0b09)]">
            <div class="absolute inset-3 border border-[#c8a45f]/15"></div>
            <Gem class="h-10 w-10 text-[#c8a45f]/35" />
          </div>
        </div>
        <p v-if="payload.subtitle" class="mt-6 text-center font-mono text-[10px] uppercase tracking-[0.34em] text-[#c8a45f]/85">{{ payload.subtitle }}</p>
      </div>

      <!-- Title -->
      <h1 class="rise rise-2 mt-8 text-balance text-center font-serif text-[34px] leading-[1.18] tracking-tight">{{ payload.title || '未命名臻品' }}</h1>
      <div class="rise rise-2 mx-auto mt-8 h-px w-16 bg-[#c8a45f]/50"></div>

      <!-- Description -->
      <p v-if="payload.description" class="rise rise-2 mt-8 whitespace-pre-line text-center text-[13px] leading-[2.05] text-[#f4efe6]/60">{{ payload.description }}</p>

      <!-- Price -->
      <div v-if="payload.price" class="rise rise-3 mt-12 text-center">
        <p class="font-mono text-[9px] uppercase tracking-[0.34em] text-[#f4efe6]/35">Price</p>
        <div class="mt-5 flex items-baseline justify-center gap-4">
          <span class="font-serif text-[38px] leading-none tracking-tight tabular-nums"><span class="mr-1 text-[20px] text-[#c8a45f]">¥</span>{{ payload.price }}</span>
          <span v-if="payload.originalPrice" class="text-[15px] text-[#f4efe6]/35 line-through tabular-nums">¥{{ payload.originalPrice }}</span>
        </div>
      </div>

      <!-- Highlights -->
      <section v-if="features.length" class="rise rise-3 mt-14">
        <p class="text-center font-mono text-[9px] uppercase tracking-[0.34em] text-[#c8a45f]/70">Highlights</p>
        <ul class="mt-6">
          <li v-for="(feature, index) in features" :key="index" class="flex items-baseline gap-5 border-t border-white/10 py-5 last:border-b">
            <span class="font-mono text-[10px] tabular-nums tracking-[0.2em] text-[#c8a45f]/80">{{ String(index + 1).padStart(2, '0') }}</span>
            <span class="flex-1 text-[13px] leading-relaxed text-[#f4efe6]/75">{{ feature }}</span>
          </li>
        </ul>
      </section>

      <!-- Specification -->
      <section v-if="specs.length" class="rise rise-3 mt-14">
        <p class="text-center font-mono text-[9px] uppercase tracking-[0.34em] text-[#c8a45f]/70">Specification</p>
        <dl class="mt-6">
          <div v-for="[key, value] in specs" :key="key" class="flex items-baseline justify-between gap-6 border-b border-white/10 py-4 first:border-t">
            <dt class="font-mono text-[10px] uppercase tracking-[0.22em] text-[#f4efe6]/40">{{ key }}</dt>
            <dd class="text-right text-[13px] tabular-nums text-[#f4efe6]/85">{{ value }}</dd>
          </div>
        </dl>
      </section>

      <!-- CTA -->
      <button
        v-if="payload.ctaText"
        @click="handleCta"
        class="rise rise-3 mt-12 flex w-full items-center justify-center gap-3 bg-[#c8a45f] py-5 font-mono text-[11px] uppercase tracking-[0.3em] text-[#0b0a08] transition-all hover:bg-[#d8b56f] active:scale-[0.99]"
      >
        {{ payload.ctaText }}
        <ArrowUpRight class="h-4 w-4" />
      </button>

      <!-- Footer -->
      <p class="mt-7 text-center font-mono text-[9px] uppercase tracking-[0.3em] text-[#f4efe6]/25">
        Certified<span v-if="nfcId"> · NFC {{ nfcId }}</span>
      </p>
    </div>
  </div>
</template>

<script setup>
import { computed, ref } from 'vue'
import { ArrowUpRight, Gem } from 'lucide-vue-next'

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
