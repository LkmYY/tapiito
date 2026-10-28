<template>
  <div class="pdsc-root min-h-screen bg-[#e9e4d8] text-[#2b2a26]">
    <div class="mx-auto w-full max-w-md px-5 pb-14 pt-7">
      <header class="pdsc-in flex items-center justify-between px-1 font-mono text-[9px] uppercase tracking-[0.22em] text-[#2b2a26]/45">
        <span>Stack / Studio</span>
        <span>{{ nfcId ? 'CARD ' + nfcId : 'SET 05' }}</span>
      </header>

      <div class="pdsc-in pdsc-d1 pdsc-stack mt-5">
        <div class="pdsc-card pdsc-card-back-2"></div>
        <div class="pdsc-card pdsc-card-back-1"></div>

        <article class="pdsc-card pdsc-main">
          <div class="flex items-center justify-between">
            <span class="rounded-full border border-[#2b2a26]/20 px-2.5 py-1 font-mono text-[8px] uppercase tracking-[0.14em]">Object card</span>
            <span class="font-mono text-[9px] text-[#2b2a26]/35">01 / 06</span>
          </div>
          <div class="pdsc-visual mt-3">
            <span class="pdsc-arch"></span>
            <div class="pdsc-object">
              <span class="pdsc-obj-lid"></span>
              <span class="pdsc-obj-body"></span>
              <span class="pdsc-obj-dot"></span>
            </div>
            <span class="pdsc-visual-cap">MATTE STONE · 320 ML</span>
          </div>
          <p class="mt-3 font-mono text-[8px] uppercase tracking-[0.18em] text-[#a16207]">{{ p.subtitle || '—' }}</p>
          <h1 class="mt-1 text-[26px] font-semibold leading-[1.05] tracking-[-0.02em]">{{ p.title }}</h1>
          <div class="mt-3 flex items-end justify-between border-t border-[#2b2a26]/15 pt-3">
            <span class="text-[30px] font-semibold leading-none tabular-nums">¥{{ p.price || '—' }}</span>
            <div class="text-right">
              <p v-if="p.originalPrice" class="font-mono text-[9px] text-[#2b2a26]/35 line-through">¥{{ p.originalPrice }}</p>
              <p v-if="discount" class="font-mono text-[8px] uppercase tracking-[0.1em] text-[#a16207]">−{{ discount }}% today</p>
            </div>
          </div>
        </article>

        <article class="pdsc-card pdsc-sub">
          <div class="flex items-center justify-between">
            <h2 class="font-mono text-[9px] uppercase tracking-[0.2em]">Why keep it</h2>
            <span class="font-mono text-[9px] text-[#2b2a26]/35">02</span>
          </div>
          <ul class="mt-2 grid grid-cols-2 gap-x-4">
            <li v-for="(f, i) in features" :key="i" class="flex gap-1.5 border-b border-[#2b2a26]/12 py-2 text-[11px] leading-snug">
              <span class="text-[#a16207]">·</span>{{ f }}
            </li>
            <li v-if="!features.length" class="col-span-2 border-b border-[#2b2a26]/12 py-2 text-[11px] text-[#2b2a26]/40"></li>
          </ul>
          <p v-if="p.description" class="mt-2.5 text-[11px] leading-[1.8] text-[#2b2a26]/55">{{ p.description }}</p>
        </article>

        <article class="pdsc-card pdsc-sub">
          <div class="flex items-center justify-between">
            <h2 class="font-mono text-[9px] uppercase tracking-[0.2em]">Spec card</h2>
            <span class="font-mono text-[9px] text-[#2b2a26]/35">03</span>
          </div>
          <dl class="mt-2 grid grid-cols-2 gap-x-4">
            <div v-for="[k, v] in specs" :key="k" class="flex items-baseline justify-between border-b border-[#2b2a26]/12 py-1.5">
              <dt class="font-mono text-[8px] uppercase tracking-[0.1em] text-[#2b2a26]/45">{{ k }}</dt>
              <dd class="text-[10.5px] font-medium">{{ v }}</dd>
            </div>
            <div v-if="!specs.length" class="col-span-2 border-b border-[#2b2a26]/12 py-1.5 text-[11px] text-[#2b2a26]/40"></div>
          </dl>
        </article>
      </div>

      <button class="pdsc-in pdsc-d2 pdsc-buy mt-4 flex w-full items-center justify-between px-5 py-4" @click="buy">
        <span class="text-[13px] font-semibold">{{ p.ctaText || '立即购买' }}</span>
        <span class="flex items-center gap-2 font-mono text-[9px] uppercase tracking-[0.16em]">ADD TO STACK <ArrowUpRight class="h-4 w-4" /></span>
      </button>

      <p class="mt-5 text-center font-mono text-[8px] uppercase tracking-[0.26em] text-[#2b2a26]/30">Six cards · one object · stack studio</p>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { ArrowUpRight } from 'lucide-vue-next'

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
.pdsc-stack { position: relative; padding-bottom: 10px; }
.pdsc-card { border-radius: 20px; border: 1px solid rgba(43, 42, 38, 0.16); background: #f7f3e9; box-shadow: 0 18px 30px -26px rgba(43, 42, 38, 0.6); }
.pdsc-card-back-2 { position: absolute; left: 20px; right: 20px; top: 18px; height: 60px; transform: rotate(-3deg); background: #d6cfbf; }
.pdsc-card-back-1 { position: absolute; left: 10px; right: 10px; top: 9px; height: 60px; transform: rotate(2deg); background: #e4ddce; }
.pdsc-main { position: relative; z-index: 2; padding: 14px; }
.pdsc-sub { position: relative; z-index: 2; margin-top: 10px; padding: 13px 14px; }
.pdsc-visual {
  position: relative;
  aspect-ratio: 16 / 10;
  overflow: hidden;
  border-radius: 14px;
  background: linear-gradient(160deg, #dfd7c6, #c8bea9 70%);
}
.pdsc-arch { position: absolute; left: 50%; top: 10%; width: 54%; height: 102%; transform: translateX(-50%); border-radius: 999px 999px 0 0; background: rgba(247, 243, 233, 0.55); border: 1px solid rgba(43, 42, 38, 0.1); }
.pdsc-object { position: absolute; left: 50%; top: 52%; width: 24%; height: 62%; transform: translate(-50%, -50%); animation: pdscLift 6s ease-in-out infinite; }
.pdsc-obj-lid { position: absolute; left: 30%; top: 0; width: 40%; height: 9%; border-radius: 3px; background: #3d3b35; }
.pdsc-obj-body {
  position: absolute;
  inset: 8% 0 0;
  border-radius: 12px 12px 22px 22px;
  background: linear-gradient(100deg, #f9f5eb 0%, #ddd3c0 55%, #b6a990 100%);
  box-shadow: inset -8px 0 14px rgba(43, 42, 38, 0.18), inset 6px 0 10px rgba(255, 255, 255, 0.6), 8px 12px 18px -12px rgba(43, 42, 38, 0.5);
}
.pdsc-obj-dot { position: absolute; left: 50%; top: 48%; width: 20%; aspect-ratio: 1; transform: translate(-50%, -50%); border-radius: 999px; border: 1px solid rgba(43, 42, 38, 0.35); }
.pdsc-visual-cap { position: absolute; left: 8px; bottom: 6px; font-family: ui-monospace, monospace; font-size: 7px; letter-spacing: 0.14em; color: rgba(43, 42, 38, 0.55); }
.pdsc-buy { border-radius: 14px; background: #2b2a26; color: #f7f3e9; transition: background 0.25s ease, transform 0.16s ease; }
.pdsc-buy:hover { background: #a16207; }
.pdsc-buy:active { transform: translateY(1px); }
.pdsc-in { animation: pdscIn 0.75s cubic-bezier(0.22, 1, 0.36, 1) both; }
.pdsc-d1 { animation-delay: 0.08s; }
.pdsc-d2 { animation-delay: 0.16s; }
@keyframes pdscIn { from { opacity: 0; transform: translateY(16px); } to { opacity: 1; transform: none; } }
@keyframes pdscLift { 0%, 100% { transform: translate(-50%, -50%); } 50% { transform: translate(-50%, -53%); } }
@media (prefers-reduced-motion: reduce) {
  .pdsc-in, .pdsc-object { animation: none !important; }
  .pdsc-buy { transition: none; }
}
</style>
