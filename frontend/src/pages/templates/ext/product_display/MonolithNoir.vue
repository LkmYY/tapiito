<template>
  <div class="pdm-root min-h-screen bg-[#0a0a0a] text-[#ededed]">
    <div class="mx-auto w-full max-w-md px-6 pb-16 pt-7">
      <header class="pdm-in flex items-center justify-between font-mono text-[9px] uppercase tracking-[0.26em] text-[#ededed]/40">
        <span>Monolith</span>
        <span>{{ nfcId || 'M-01' }}</span>
      </header>

      <section class="pdm-in pdm-d1 pdm-stage mt-8">
        <div class="pdm-slab">
          <span class="pdm-groove pdm-groove-a"></span>
          <span class="pdm-groove pdm-groove-b"></span>
          <span class="pdm-slab-mark">01</span>
          <div class="pdm-object">
            <span class="pdm-obj-top"></span>
            <span class="pdm-obj-body"></span>
            <span class="pdm-obj-line"></span>
          </div>
        </div>
        <span class="pdm-shadow"></span>
        <span class="pdm-side">FORM / FUNCTION / SILENCE</span>
      </section>

      <div class="pdm-in pdm-d2 mt-9">
        <p class="font-mono text-[9px] uppercase tracking-[0.3em] text-[#ededed]/35">{{ p.subtitle || '—' }}</p>
        <h1 class="mt-3 text-[42px] font-light leading-[0.98] tracking-[-0.04em]">{{ p.title }}</h1>
      </div>

      <div class="pdm-in pdm-d2 mt-8">
        <div class="flex items-start gap-2">
          <span class="mt-2 text-[15px] font-light text-[#ededed]/40">¥</span>
          <span class="pdm-price text-[80px] font-extralight leading-[0.78] tracking-[-0.07em] tabular-nums">{{ p.price || '—' }}</span>
        </div>
        <div class="mt-4 flex items-center justify-between border-t border-[#ededed]/15 pt-3 font-mono text-[9px] uppercase tracking-[0.18em]">
          <span class="text-[#ededed]/35">CNY · tax included</span>
          <span v-if="p.originalPrice" class="text-[#ededed]/30">was <b class="font-normal line-through">¥{{ p.originalPrice }}</b></span>
          <span v-if="discount" class="text-[#c7c7a0]">−{{ discount }}%</span>
        </div>
      </div>

      <p v-if="p.description" class="pdm-in pdm-d3 mt-8 max-w-[300px] text-[12px] font-light leading-[2.1] text-[#ededed]/45">{{ p.description }}</p>

      <section class="pdm-in pdm-d3 mt-10">
        <div class="flex items-baseline justify-between border-b border-[#ededed]/15 pb-2">
          <h2 class="font-mono text-[9px] uppercase tracking-[0.24em] text-[#ededed]/40">Qualities</h2>
          <span class="font-mono text-[9px] text-[#ededed]/25">Ⅰ</span>
        </div>
        <ul class="mt-1">
          <li v-for="(f, i) in features" :key="i" class="flex items-baseline gap-5 border-b border-[#ededed]/10 py-4">
            <span class="font-mono text-[9px] text-[#ededed]/25">{{ String(i + 1).padStart(2, '0') }}</span>
            <span class="flex-1 text-[13px] font-light leading-snug text-[#ededed]/80">{{ f }}</span>
          </li>
          <li v-if="!features.length" class="border-b border-[#ededed]/10 py-4 text-[12px] font-light text-[#ededed]/30"></li>
        </ul>
      </section>

      <section class="pdm-in pdm-d3 mt-8">
        <div class="flex items-baseline justify-between border-b border-[#ededed]/15 pb-2">
          <h2 class="font-mono text-[9px] uppercase tracking-[0.24em] text-[#ededed]/40">Measure</h2>
          <span class="font-mono text-[9px] text-[#ededed]/25">Ⅱ</span>
        </div>
        <dl class="mt-1">
          <div v-for="[k, v] in specs" :key="k" class="flex items-baseline justify-between gap-4 border-b border-[#ededed]/10 py-3">
            <dt class="font-mono text-[9px] uppercase tracking-[0.14em] text-[#ededed]/30">{{ k }}</dt>
            <dd class="text-[12px] font-light text-[#ededed]/75">{{ v }}</dd>
          </div>
          <div v-if="!specs.length" class="border-b border-[#ededed]/10 py-3 text-[12px] font-light text-[#ededed]/30"></div>
        </dl>
      </section>

      <button class="pdm-in pdm-d4 pdm-buy mt-9 flex w-full items-center justify-between border border-[#ededed]/30 px-5 py-4" @click="buy">
        <span class="font-mono text-[9px] uppercase tracking-[0.26em]">{{ p.ctaText || '立即购买' }}</span>
        <ArrowRight class="h-4 w-4" />
      </button>

      <p class="mt-7 text-center font-mono text-[8px] uppercase tracking-[0.3em] text-[#ededed]/20">Nothing added · nothing missing</p>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { ArrowRight } from 'lucide-vue-next'

const props = defineProps(['payload', 'nfcId'])
const p = computed(() => props.payload || {})
const features = computed(() => (Array.isArray(p.value.features) ? p.value.features.filter(Boolean) : []))
const specs = computed(() => Object.entries(p.value.specifications || {}))

const discount = computed(() => {
  const price = Number(p.value.price)
  const original = Number(p.value.originalPrice)
  if (!Number.isFinite(price) || !Number.isFinite(original) || original <= price) return null
  return Math.round(((original - price) / original) * 100)
})

function buy() {
  const action = p.value.ctaAction
  if (!action) return
  try {
    const url = new URL(action)
    if (['http:', 'https:'].includes(url.protocol)) window.location.href = url.href
  } catch {
    // 链接不合法时忽略
  }
}
</script>

<style scoped>
.pdm-stage { position: relative; aspect-ratio: 1 / 1.18; }
.pdm-slab {
  position: relative;
  width: 100%;
  height: 100%;
  overflow: hidden;
  background: linear-gradient(160deg, #1c1c1c 0%, #111 55%, #050505 100%);
  box-shadow: inset 0 0 0 1px rgba(237, 237, 237, 0.1), 0 30px 60px -40px #000;
}
.pdm-groove { position: absolute; left: 8%; right: 8%; height: 1px; background: rgba(237, 237, 237, 0.1); }
.pdm-groove-a { top: 18%; }
.pdm-groove-b { bottom: 18%; }
.pdm-slab-mark { position: absolute; right: 6%; top: 4%; font-size: 72px; font-weight: 200; line-height: 0.8; letter-spacing: -0.06em; color: rgba(237, 237, 237, 0.07); }
.pdm-object { position: absolute; left: 50%; top: 50%; width: 36%; height: 56%; transform: translate(-50%, -50%); animation: pdmLevitate 8s ease-in-out infinite; }
.pdm-obj-top {
  position: absolute;
  left: 34%;
  top: 0;
  width: 32%;
  height: 6%;
  border-radius: 2px 2px 0 0;
  background: #2d2d2d;
  box-shadow: inset 0 1px 0 rgba(255, 255, 255, 0.16);
}
.pdm-obj-body {
  position: absolute;
  inset: 5% 0 0;
  border-radius: 20% 20% 14% 14% / 8% 8% 6% 6%;
  background: linear-gradient(100deg, #3a3a3a 0%, #1d1d1d 38%, #080808 100%);
  box-shadow: inset -14px 0 26px rgba(0, 0, 0, 0.8), inset 10px 0 18px rgba(255, 255, 255, 0.09), 0 30px 44px -26px #000;
}
.pdm-obj-line { position: absolute; left: 14%; right: 14%; top: 38%; height: 1px; background: rgba(237, 237, 237, 0.16); }
.pdm-shadow { position: absolute; left: 20%; right: 20%; bottom: 8%; height: 3%; border-radius: 50%; background: rgba(0, 0, 0, 0.85); filter: blur(10px); }
.pdm-side { position: absolute; right: -84px; top: 50%; transform: translateY(-50%); font-family: ui-monospace, monospace; font-size: 7px; letter-spacing: 0.3em; color: rgba(237, 237, 237, 0.28); white-space: nowrap; }
.pdm-price { font-feature-settings: 'tnum'; }
.pdm-buy { background: transparent; color: #ededed; transition: background 0.35s ease, color 0.35s ease, padding 0.3s ease; }
.pdm-buy:hover { background: #ededed; color: #0a0a0a; padding-inline: 26px; }
.pdm-buy:active { transform: scale(0.99); }
.pdm-in { animation: pdmIn 0.9s cubic-bezier(0.22, 1, 0.36, 1) both; }
.pdm-d1 { animation-delay: 0.08s; }
.pdm-d2 { animation-delay: 0.16s; }
.pdm-d3 { animation-delay: 0.24s; }
.pdm-d4 { animation-delay: 0.32s; }
@keyframes pdmIn { from { opacity: 0; transform: translateY(20px); } to { opacity: 1; transform: none; } }
@keyframes pdmLevitate { 0%, 100% { transform: translate(-50%, -50%); } 50% { transform: translate(-50%, -53%); } }
@media (prefers-reduced-motion: reduce) {
  .pdm-in, .pdm-object { animation: none !important; }
  .pdm-buy { transition: none; }
}
</style>
