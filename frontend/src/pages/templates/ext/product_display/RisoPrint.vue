<template>
  <div class="pdr-root min-h-screen bg-[#f4ead7] text-[#1c2a4a]">
    <div class="pdr-noise pointer-events-none"></div>
    <div class="mx-auto w-full max-w-md px-5 pb-14 pt-6">
      <header class="pdr-in flex items-end justify-between border-b-2 border-[#1c2a4a] pb-2">
        <div>
          <p class="font-mono text-[8px] uppercase tracking-[0.3em] text-[#e04a2f]">Riso edition / No. 12</p>
          <h2 class="mt-1 text-[15px] font-bold uppercase tracking-[0.14em]">Paper Goods Co.</h2>
        </div>
        <span class="font-mono text-[9px] text-[#1c2a4a]/50">{{ nfcId || '—' }}</span>
      </header>

      <section class="pdr-in pdr-d1 pdr-stage mt-5">
        <span class="pdr-reg pdr-reg-a">PRINT<br />OFFSET</span>
        <span class="pdr-reg pdr-reg-b">2 COLORS</span>
        <div class="pdr-halftone"></div>
        <div class="pdr-vessel">
          <span class="pdr-vessel-neck"></span>
          <span class="pdr-vessel-body"></span>
          <span class="pdr-vessel-print">RISO<br />INK</span>
        </div>
        <span class="pdr-overlay">{{ discount ? '−' + discount + '%' : 'EDITION 12' }}</span>
      </section>

      <div class="pdr-in pdr-d2 mt-6">
        <p class="font-mono text-[9px] uppercase tracking-[0.24em] text-[#e04a2f]">{{ p.subtitle || '—' }}</p>
        <h1 class="pdr-title mt-2 text-[46px] font-black uppercase leading-[0.88] tracking-[-0.045em]">{{ p.title }}</h1>
      </div>

      <div class="pdr-in pdr-d2 mt-5 flex items-end justify-between border-y-2 border-[#1c2a4a] py-4">
        <div>
          <p class="font-mono text-[8px] uppercase tracking-[0.2em] text-[#1c2a4a]/45">Direct price</p>
          <p class="mt-1 text-[42px] font-black leading-none tabular-nums text-[#e04a2f]">¥{{ p.price || '—' }}</p>
        </div>
        <div class="text-right">
          <p v-if="p.originalPrice" class="font-mono text-[11px] text-[#1c2a4a]/45 line-through">¥{{ p.originalPrice }}</p>
          <p class="mt-1 font-mono text-[8px] uppercase leading-relaxed tracking-[0.1em] text-[#1c2a4a]/55">Ink on paper<br />Ships flat</p>
        </div>
      </div>

      <section class="pdr-in pdr-d3 mt-5 grid grid-cols-12 gap-4">
        <div class="col-span-7">
          <h2 class="inline-block bg-[#1c2a4a] px-2 py-0.5 font-mono text-[8.5px] uppercase tracking-[0.2em] text-[#f4ead7]">What's inside</h2>
          <ul class="mt-2">
            <li v-for="(f, i) in features" :key="i" class="flex gap-2 border-b border-[#1c2a4a]/20 py-2.5 text-[12px] leading-snug">
              <span class="font-mono text-[9px] font-bold text-[#e04a2f]">{{ String(i + 1).padStart(2, '0') }}</span>
              <span>{{ f }}</span>
            </li>
            <li v-if="!features.length" class="border-b border-[#1c2a4a]/20 py-2.5 text-[12px] text-[#1c2a4a]/40"></li>
          </ul>
        </div>
        <div class="col-span-5">
          <h2 class="inline-block bg-[#e04a2f] px-2 py-0.5 font-mono text-[8.5px] uppercase tracking-[0.2em] text-[#f4ead7]">Specs</h2>
          <dl class="mt-2">
            <div v-for="[k, v] in specs" :key="k" class="border-b border-[#1c2a4a]/20 py-2.5">
              <dt class="font-mono text-[8px] uppercase tracking-[0.14em] text-[#1c2a4a]/45">{{ k }}</dt>
              <dd class="mt-0.5 text-[12px] font-bold">{{ v }}</dd>
            </div>
            <div v-if="!specs.length" class="border-b border-[#1c2a4a]/20 py-2.5 text-[11px] text-[#1c2a4a]/40"></div>
          </dl>
        </div>
      </section>

      <p v-if="p.description" class="pdr-in pdr-d3 mt-5 border border-[#1c2a4a] p-3 text-[11.5px] leading-[1.85] text-[#1c2a4a]/75">{{ p.description }}</p>

      <button class="pdr-in pdr-d4 pdr-buy mt-5 flex w-full items-center justify-between border-2 border-[#1c2a4a] px-4 py-3.5" @click="buy">
        <span class="font-mono text-[11px] font-bold uppercase tracking-[0.2em]">{{ p.ctaText || '立即购买' }}</span>
        <span class="flex items-center gap-2 font-mono text-[9px] uppercase tracking-[0.12em]">Add to bag <ArrowUpRight class="h-4 w-4" /></span>
      </button>

      <p class="mt-5 text-center font-mono text-[8px] uppercase tracking-[0.22em] text-[#1c2a4a]/40">Risograph printed · spot colors · edition of 200</p>
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
.pdr-root { position: relative; overflow: hidden; }
.pdr-noise {
  position: absolute;
  inset: 0;
  opacity: 0.5;
  background-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' width='120' height='120'%3E%3Cfilter id='n'%3E%3CfeTurbulence type='fractalNoise' baseFrequency='0.9' numOctaves='2'/%3E%3C/filter%3E%3Crect width='100%25' height='100%25' filter='url(%23n)' opacity='0.18'/%3E%3C/svg%3E");
  pointer-events: none;
}
.pdr-stage {
  position: relative;
  aspect-ratio: 1 / 1.05;
  overflow: hidden;
  border: 2px solid #1c2a4a;
  background: #efe2c9;
}
.pdr-halftone {
  position: absolute;
  inset: 0;
  background-image: radial-gradient(#e04a2f 1.1px, transparent 1.2px);
  background-size: 7px 7px;
  opacity: 0.5;
  mask-image: radial-gradient(circle at 50% 45%, black 20%, transparent 72%);
}
.pdr-vessel {
  position: absolute;
  left: 50%;
  top: 52%;
  width: 40%;
  height: 62%;
  transform: translate(-50%, -50%);
  animation: pdrLift 7s ease-in-out infinite;
}
.pdr-vessel-body {
  position: absolute;
  inset: 14% 0 0;
  border: 2px solid #1c2a4a;
  border-radius: 6px 6px 18px 18px;
  background: #b9d9d2;
  box-shadow: 8px 8px 0 rgba(224, 74, 47, 0.9);
}
.pdr-vessel-neck {
  position: absolute;
  left: 32%;
  top: 2%;
  width: 36%;
  height: 15%;
  border: 2px solid #1c2a4a;
  border-bottom: 0;
  background: #f4ead7;
}
.pdr-vessel-print {
  position: absolute;
  left: 50%;
  top: 56%;
  transform: translate(-50%, -50%) rotate(-4deg);
  font-family: ui-monospace, monospace;
  font-size: 14px;
  font-weight: 700;
  line-height: 1.15;
  letter-spacing: 0.08em;
  text-align: center;
  color: #1c2a4a;
}
.pdr-reg { position: absolute; font-family: ui-monospace, monospace; font-size: 8px; line-height: 1.5; letter-spacing: 0.16em; color: #1c2a4a; z-index: 2; }
.pdr-reg-a { left: 10px; top: 10px; }
.pdr-reg-b { right: 10px; top: 10px; color: #e04a2f; text-align: right; }
.pdr-overlay {
  position: absolute;
  left: 50%;
  bottom: 18px;
  transform: translateX(-50%) rotate(-2deg);
  padding: 4px 12px;
  background: #e04a2f;
  color: #f4ead7;
  font-family: ui-monospace, monospace;
  font-size: 10px;
  letter-spacing: 0.2em;
  mix-blend-mode: multiply;
}
.pdr-title { text-shadow: 3px 3px 0 rgba(224, 74, 47, 0.55); }
.pdr-buy { background: #1c2a4a; color: #f4ead7; transition: background 0.25s ease, transform 0.16s ease; }
.pdr-buy:hover { background: #2c3f6b; }
.pdr-buy:active { transform: translate(1px, 1px); }
.pdr-in { animation: pdrIn 0.7s cubic-bezier(0.22, 1, 0.36, 1) both; }
.pdr-d1 { animation-delay: 0.06s; }
.pdr-d2 { animation-delay: 0.12s; }
.pdr-d3 { animation-delay: 0.18s; }
.pdr-d4 { animation-delay: 0.24s; }
@keyframes pdrIn { from { opacity: 0; transform: translateY(16px); } to { opacity: 1; transform: none; } }
@keyframes pdrLift { 0%, 100% { transform: translate(-50%, -50%); } 50% { transform: translate(-50%, -53%); } }
@media (prefers-reduced-motion: reduce) {
  .pdr-in, .pdr-vessel { animation: none !important; }
  .pdr-buy { transition: none; }
}
</style>
