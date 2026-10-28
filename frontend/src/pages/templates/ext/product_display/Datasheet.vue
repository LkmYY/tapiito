<template>
  <div class="pdd-root min-h-screen bg-[#e9edf2] text-[#0f1b2a]">
    <div class="mx-auto w-full max-w-md pb-14">
      <header class="pdd-in flex items-center justify-between bg-[#0f1b2a] px-5 py-3 text-[#dbe7f3]">
        <span class="flex items-center gap-2 font-mono text-[10px] uppercase tracking-[0.18em]"><Cpu class="h-4 w-4 text-[#67e8f9]" /> SYSTEM / 04</span>
        <span class="font-mono text-[9px] text-[#67e8f9]">{{ nfcId ? 'UNIT ' + nfcId : 'REV. B' }}</span>
      </header>

      <div class="px-5">
        <section class="pdd-in pdd-d1 pdd-stage mt-4">
          <div class="pdd-grid"></div>
          <span class="pdd-axis pdd-axis-x"></span>
          <span class="pdd-axis pdd-axis-y"></span>
          <div class="pdd-device">
            <span class="pdd-screen"><i class="pdd-wave"></i></span>
            <span class="pdd-btn"></span>
            <span class="pdd-vent"></span>
          </div>
          <span class="pdd-callout pdd-callout-a">A. 12MM DRIVER</span>
          <span class="pdd-callout pdd-callout-b">B. USB-C</span>
          <span class="pdd-dim">48 × 62 × 24 mm</span>
        </section>

        <div class="pdd-in pdd-d2 mt-5 flex items-start justify-between gap-4">
          <div>
            <p class="font-mono text-[9px] uppercase tracking-[0.2em] text-[#0e7490]">PRODUCT DATASHEET</p>
            <h1 class="mt-1.5 text-[30px] font-semibold leading-[1.05] tracking-[-0.02em]">{{ p.title }}</h1>
            <p class="mt-1 text-[11.5px] text-[#0f1b2a]/55">{{ p.subtitle || '—' }}</p>
          </div>
          <div class="pdd-barcode">
            <i v-for="n in 12" :key="n" :style="{ width: (n % 3) + 1 + 'px' }"></i>
          </div>
        </div>

        <div class="pdd-in pdd-d2 mt-4 grid grid-cols-12 items-end border-y border-[#0f1b2a]/20 py-3">
          <div class="col-span-7">
            <p class="font-mono text-[8px] uppercase tracking-[0.18em] text-[#0f1b2a]/45">MSRP / CNY</p>
            <p class="mt-1 text-[38px] font-semibold leading-none tabular-nums">¥{{ p.price || '—' }}</p>
          </div>
          <div class="col-span-5 text-right font-mono text-[9px] leading-relaxed">
            <p v-if="p.originalPrice" class="text-[#0f1b2a]/40 line-through">¥{{ p.originalPrice }}</p>
            <p v-if="discount" class="mt-1 inline-block bg-[#0e7490] px-2 py-0.5 text-white">Δ −{{ discount }}%</p>
            <p class="mt-1 text-[#0f1b2a]/35">INCL. VAT</p>
          </div>
        </div>

        <section class="pdd-in pdd-d3 mt-4">
          <h2 class="flex items-center justify-between font-mono text-[9px] uppercase tracking-[0.2em]">
            <span>01 — FEATURE MATRIX</span><span class="text-[#0f1b2a]/35">N={{ features.length }}</span>
          </h2>
          <ul class="mt-2 border border-[#0f1b2a]/20 bg-white/60">
            <li v-for="(f, i) in features" :key="i" class="flex items-center gap-3 border-b border-[#0f1b2a]/12 px-3 py-2.5 last:border-0">
              <span class="font-mono text-[8.5px] text-[#0e7490]">F{{ String(i + 1).padStart(2, '0') }}</span>
              <span class="flex-1 text-[11.5px]">{{ f }}</span>
              <span class="font-mono text-[8px] text-[#0f1b2a]/30">OK</span>
            </li>
            <li v-if="!features.length" class="px-3 py-2.5 text-[11.5px] text-[#0f1b2a]/40"></li>
          </ul>
        </section>

        <section class="pdd-in pdd-d3 mt-4">
          <h2 class="font-mono text-[9px] uppercase tracking-[0.2em]">02 — SPECIFICATION TABLE</h2>
          <table class="mt-2 w-full border-collapse text-left">
            <tbody>
              <tr v-for="([k, v], i) in specs" :key="k" :class="i % 2 ? 'bg-[#0f1b2a]/5' : 'bg-white/55'">
                <th class="w-1/2 border border-[#0f1b2a]/15 px-3 py-2 font-mono text-[8.5px] font-normal uppercase tracking-[0.12em] text-[#0f1b2a]/45">{{ k }}</th>
                <td class="border border-[#0f1b2a]/15 px-3 py-2 font-mono text-[10.5px]">{{ v }}</td>
              </tr>
              <tr v-if="!specs.length">
                <td class="border border-[#0f1b2a]/15 px-3 py-2 text-[11px] text-[#0f1b2a]/40" colspan="2"></td>
              </tr>
            </tbody>
          </table>
        </section>

        <section v-if="p.description" class="pdd-in pdd-d3 mt-4 border border-[#0f1b2a]/20 bg-white/55 p-3">
          <h2 class="font-mono text-[8.5px] uppercase tracking-[0.18em] text-[#0f1b2a]/45">03 — NOTES</h2>
          <p class="mt-1.5 font-mono text-[10.5px] leading-[1.85] text-[#0f1b2a]/70">{{ p.description }}</p>
        </section>

        <button class="pdd-in pdd-d4 pdd-buy mt-5 flex w-full items-center justify-between px-4 py-3.5" @click="buy">
          <span class="font-mono text-[11px] font-semibold uppercase tracking-[0.16em]">{{ p.ctaText || '立即购买' }}</span>
          <span class="flex items-center gap-2 font-mono text-[9px] uppercase">ORDER <ArrowUpRight class="h-4 w-4" /></span>
        </button>

        <p class="mt-4 text-center font-mono text-[8px] uppercase tracking-[0.2em] text-[#0f1b2a]/35">Document generated · rev 2.4 · NFC verified</p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { ArrowUpRight, Cpu } from 'lucide-vue-next'

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
.pdd-stage { position: relative; aspect-ratio: 4 / 3.4; overflow: hidden; border: 1px solid rgba(15, 27, 42, 0.2); background: #f8fafc; }
.pdd-grid {
  position: absolute;
  inset: 0;
  background-image: linear-gradient(rgba(15, 27, 42, 0.08) 1px, transparent 1px), linear-gradient(90deg, rgba(15, 27, 42, 0.08) 1px, transparent 1px);
  background-size: 22px 22px;
}
.pdd-axis { position: absolute; background: rgba(14, 116, 144, 0.5); }
.pdd-axis-x { left: 0; right: 0; top: 50%; height: 1px; }
.pdd-axis-y { top: 0; bottom: 0; left: 50%; width: 1px; }
.pdd-device { position: absolute; left: 50%; top: 50%; width: 38%; height: 58%; transform: translate(-50%, -50%) rotate(-6deg); animation: pddFloat 6s ease-in-out infinite; }
.pdd-screen {
  position: absolute;
  inset: 0 0 14%;
  border-radius: 22% 22% 30% 30% / 16% 16% 24% 24%;
  background: linear-gradient(145deg, #1e293b, #0f172a 60%, #020617);
  box-shadow: inset 0 1px 0 rgba(255, 255, 255, 0.2), 0 20px 30px -20px rgba(2, 6, 23, 0.7);
  display: grid;
  place-items: center;
  overflow: hidden;
}
.pdd-wave { width: 52%; height: 14%; background: repeating-linear-gradient(90deg, #67e8f9 0 2px, transparent 2px 5px); clip-path: polygon(0 45%, 16% 45%, 22% 0, 30% 100%, 38% 30%, 46% 70%, 54% 15%, 62% 85%, 70% 40%, 78% 60%, 84% 45%, 100% 45%, 100% 55%, 40% 55%, 0 55%); }
.pdd-btn { position: absolute; right: 14%; bottom: 2%; width: 12%; height: 4%; border-radius: 3px; background: #67e8f9; }
.pdd-vent { position: absolute; left: 50%; bottom: 19%; width: 34%; height: 2px; transform: translateX(-50%); background: rgba(148, 163, 184, 0.6); box-shadow: 0 -4px 0 rgba(148, 163, 184, 0.3); }
.pdd-callout { position: absolute; font-family: ui-monospace, monospace; font-size: 7px; letter-spacing: 0.12em; color: #0e7490; border-bottom: 1px solid rgba(14, 116, 144, 0.5); padding-bottom: 2px; }
.pdd-callout-a { left: 8%; top: 18%; }
.pdd-callout-b { right: 8%; bottom: 18%; }
.pdd-callout-a::after { content: ''; position: absolute; right: -26px; top: 100%; width: 26px; height: 1px; background: rgba(14, 116, 144, 0.5); transform: rotate(28deg); transform-origin: left; }
.pdd-dim { position: absolute; left: 50%; bottom: 8px; transform: translateX(-50%); font-family: ui-monospace, monospace; font-size: 7.5px; letter-spacing: 0.14em; color: rgba(15, 27, 42, 0.45); }
.pdd-barcode { display: flex; align-items: stretch; gap: 2px; height: 30px; padding: 3px; border: 1px solid rgba(15, 27, 42, 0.3); background: #fff; }
.pdd-barcode i { display: block; background: #0f1b2a; }
.pdd-buy { background: #0e7490; color: #ecfeff; transition: background 0.25s ease, transform 0.16s ease; }
.pdd-buy:hover { background: #0891b2; }
.pdd-buy:active { transform: translateY(1px); }
.pdd-in { animation: pddIn 0.7s cubic-bezier(0.22, 1, 0.36, 1) both; }
.pdd-d1 { animation-delay: 0.05s; }
.pdd-d2 { animation-delay: 0.1s; }
.pdd-d3 { animation-delay: 0.16s; }
.pdd-d4 { animation-delay: 0.22s; }
@keyframes pddIn { from { opacity: 0; transform: translateY(14px); } to { opacity: 1; transform: none; } }
@keyframes pddFloat { 0%, 100% { transform: translate(-50%, -50%) rotate(-6deg); } 50% { transform: translate(-50%, -52%) rotate(-3deg); } }
@media (prefers-reduced-motion: reduce) {
  .pdd-in, .pdd-device { animation: none !important; }
  .pdd-buy { transition: none; }
}
</style>
