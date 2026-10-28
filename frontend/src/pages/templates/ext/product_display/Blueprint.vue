<template>
  <div class="pdbp-root min-h-screen bg-[#0e3a5c] text-[#d7efff]">
    <div class="pdbp-grid pointer-events-none"></div>
    <div class="relative mx-auto w-full max-w-md px-5 pb-12 pt-6">
      <header class="pdbp-in flex items-center justify-between border border-[#7dd3fc]/40 px-3 py-2 font-mono text-[9px] uppercase tracking-[0.2em] text-[#7dd3fc]">
        <span>DWG NO. PD-25 / REV C</span>
        <span>{{ nfcId ? 'NFC:' + nfcId : 'SHEET 01' }}</span>
      </header>

      <section class="pdbp-in pdbp-d1 pdbp-sheet mt-4">
        <div class="pdbp-sheet-head">
          <span>ORTHOGRAPHIC VIEW — FRONT</span>
          <span>SCALE 1:2</span>
        </div>
        <div class="pdbp-draw">
          <div class="pdbp-crosshair"></div>
          <div class="pdbp-product">
            <span class="pdbp-part pdbp-part-a"></span>
            <span class="pdbp-part pdbp-part-b"></span>
            <span class="pdbp-part pdbp-part-c"></span>
            <span class="pdbp-part pdbp-part-d"></span>
            <span class="pdbp-center"></span>
          </div>
          <span class="pdbp-dim pdbp-dim-h"><i></i><b>128.0</b></span>
          <span class="pdbp-dim pdbp-dim-v"><i></i><b>96.4</b></span>
          <span class="pdbp-note pdbp-note-a">A</span>
          <span class="pdbp-note pdbp-note-b">B</span>
          <span class="pdbp-note pdbp-note-c">C</span>
          <span class="pdbp-leader pdbp-leader-a"></span>
          <span class="pdbp-leader pdbp-leader-b"></span>
          <span class="pdbp-draw-label">SECTION A–A</span>
          <span class="pdbp-tol">TOL ±0.2</span>
        </div>
        <div class="pdbp-sheet-foot">
          <span>PROJECTION: FIRST ANGLE</span>
          <span>UNITS: MM</span>
          <span>DRAWN: STUDIO 25</span>
        </div>
      </section>

      <div class="pdbp-in pdbp-d2 mt-5 flex items-start justify-between gap-3">
        <div>
          <p class="font-mono text-[8.5px] uppercase tracking-[0.22em] text-[#7dd3fc]">{{ p.subtitle || '—' }}</p>
          <h1 class="mt-1.5 text-[29px] font-semibold leading-[1.05] tracking-[-0.01em]">{{ p.title }}</h1>
        </div>
        <span class="border border-[#7dd3fc]/40 px-2 py-1 font-mono text-[8px] uppercase leading-tight text-[#7dd3fc]">APPROVED<br />FOR SALE</span>
      </div>

      <div class="pdbp-in pdbp-d2 mt-4 grid grid-cols-[1.3fr_1fr] border border-[#7dd3fc]/40">
        <div class="pdbp-price-cell">
          <p class="font-mono text-[8px] uppercase tracking-[0.18em] text-[#7dd3fc]/70">UNIT COST / CNY</p>
          <p class="mt-1 font-mono text-[38px] font-semibold leading-none tabular-nums">¥{{ p.price || '—' }}</p>
        </div>
        <div class="border-l border-[#7dd3fc]/40 px-3 py-2.5 font-mono text-[9px] leading-relaxed">
          <p class="text-[#7dd3fc]/70">REF. {{ p.originalPrice || '—' }}</p>
          <p v-if="discount" class="mt-1 text-[#fde68a]">ADJ −{{ discount }}%</p>
          <p class="mt-1 text-[#d7efff]/40">FOB SHENZHEN</p>
        </div>
      </div>

      <section class="pdbp-in pdbp-d3 mt-4">
        <div class="flex items-center justify-between border-b border-[#7dd3fc]/40 pb-1 font-mono text-[8.5px] uppercase tracking-[0.18em] text-[#7dd3fc]">
          <span>PART LIST — FEATURES</span><span>{{ String(features.length).padStart(2, '0') }} ITEMS</span>
        </div>
        <ul class="mt-1.5">
          <li v-for="(f, i) in features" :key="i" class="flex items-start gap-3 border-b border-[#d7efff]/12 py-2">
            <span class="font-mono text-[8px] text-[#7dd3fc]/70">{{ String.fromCharCode(65 + i) }}{{ String(i + 1).padStart(2, '0') }}</span>
            <span class="flex-1 text-[11.5px] leading-snug text-[#d7efff]/85">{{ f }}</span>
            <span class="mt-1 h-1.5 w-1.5 rounded-full border border-[#7dd3fc]/60"></span>
          </li>
          <li v-if="!features.length" class="border-b border-[#d7efff]/12 py-2 text-[11.5px] text-[#d7efff]/40"></li>
        </ul>
      </section>

      <section class="pdbp-in pdbp-d3 mt-4">
        <div class="flex items-center justify-between border-b border-[#7dd3fc]/40 pb-1 font-mono text-[8.5px] uppercase tracking-[0.18em] text-[#7dd3fc]">
          <span>DIMENSIONS &amp; MATERIALS</span><span>TABLE 02</span>
        </div>
        <dl class="mt-1.5 grid grid-cols-2 gap-px bg-[#7dd3fc]/25 border border-[#7dd3fc]/25">
          <div v-for="[k, v] in specs" :key="k" class="bg-[#0e3a5c] px-2.5 py-2">
            <dt class="font-mono text-[7px] uppercase tracking-[0.12em] text-[#7dd3fc]/65">{{ k }}</dt>
            <dd class="mt-0.5 font-mono text-[10px] text-[#d7efff]/90">{{ v }}</dd>
          </div>
          <div v-if="!specs.length" class="col-span-2 bg-[#0e3a5c] px-2.5 py-2 font-mono text-[10px] text-[#d7efff]/40"></div>
        </dl>
      </section>

      <p v-if="p.description" class="pdbp-in pdbp-d3 mt-3 border-l border-[#7dd3fc]/50 pl-3 font-mono text-[9.5px] leading-[1.9] text-[#d7efff]/55">{{ p.description }}</p>

      <button class="pdbp-in pdbp-d4 pdbp-buy mt-5 flex w-full items-center justify-between border border-[#7dd3fc] px-4 py-3.5" @click="buy">
        <span class="font-mono text-[10px] font-semibold uppercase tracking-[0.2em]">{{ p.ctaText || '立即购买' }}</span>
        <span class="flex items-center gap-2 font-mono text-[9px] uppercase tracking-[0.14em]">ISSUE ORDER <ArrowUpRight class="h-4 w-4" /></span>
      </button>

      <div class="mt-5 flex items-center justify-between border border-[#7dd3fc]/30 px-2.5 py-1.5 font-mono text-[7.5px] uppercase tracking-[0.14em] text-[#7dd3fc]/60">
        <span>CHECKED: QC-04</span><span>SHEET 1 OF 1</span><span>SCALE 1:2</span>
      </div>
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
.pdbp-root { position: relative; overflow: hidden; }
.pdbp-grid {
  position: fixed;
  inset: 0;
  background-image: linear-gradient(rgba(125, 211, 252, 0.09) 1px, transparent 1px), linear-gradient(90deg, rgba(125, 211, 252, 0.09) 1px, transparent 1px);
  background-size: 20px 20px;
  pointer-events: none;
}
.pdbp-sheet { border: 1px solid rgba(125, 211, 252, 0.45); }
.pdbp-sheet-head, .pdbp-sheet-foot { display: flex; justify-content: space-between; padding: 5px 8px; font-family: ui-monospace, monospace; font-size: 7px; letter-spacing: 0.16em; color: rgba(125, 211, 252, 0.75); }
.pdbp-sheet-head { border-bottom: 1px solid rgba(125, 211, 252, 0.3); }
.pdbp-sheet-foot { border-top: 1px solid rgba(125, 211, 252, 0.3); color: rgba(215, 239, 255, 0.4); }
.pdbp-draw { position: relative; aspect-ratio: 4 / 3.05; overflow: hidden; background: rgba(6, 35, 58, 0.35); }
.pdbp-draw::before {
  content: '';
  position: absolute;
  inset: 0;
  background-image: linear-gradient(rgba(125, 211, 252, 0.07) 1px, transparent 1px), linear-gradient(90deg, rgba(125, 211, 252, 0.07) 1px, transparent 1px);
  background-size: 12px 12px;
}
.pdbp-crosshair { position: absolute; left: 50%; top: 50%; width: 84%; height: 84%; transform: translate(-50%, -50%); border: 1px dashed rgba(125, 211, 252, 0.12); }
.pdbp-product { position: absolute; left: 50%; top: 50%; width: 40%; height: 56%; transform: translate(-50%, -50%); animation: pdbpIn 1s cubic-bezier(0.22, 1, 0.36, 1) both; }
.pdbp-part { position: absolute; border: 1.5px solid #7dd3fc; background: rgba(125, 211, 252, 0.04); }
.pdbp-part-a { left: 6%; top: 0; width: 88%; height: 12%; border-radius: 4px 4px 0 0; }
.pdbp-part-b { left: 0; top: 12%; width: 100%; height: 58%; border-radius: 6px; }
.pdbp-part-c { left: 18%; bottom: 0; width: 64%; height: 30%; border-radius: 0 0 10px 10px; }
.pdbp-part-d { left: 30%; top: 30%; width: 40%; height: 22%; border-style: dashed; border-color: rgba(253, 230, 138, 0.75); border-radius: 999px; }
.pdbp-center { position: absolute; left: 50%; top: 41%; width: 6px; height: 6px; transform: translate(-50%, -50%); border: 1px solid #7dd3fc; border-radius: 999px; }
.pdbp-center::before, .pdbp-center::after { content: ''; position: absolute; background: rgba(125, 211, 252, 0.7); }
.pdbp-center::before { left: 50%; top: -10px; width: 1px; height: 24px; }
.pdbp-center::after { top: 50%; left: -10px; height: 1px; width: 24px; }
.pdbp-dim { position: absolute; font-family: ui-monospace, monospace; font-size: 7px; color: rgba(125, 211, 252, 0.8); }
.pdbp-dim i { position: absolute; background: rgba(125, 211, 252, 0.5); }
.pdbp-dim b { position: absolute; font-weight: 400; }
.pdbp-dim-h { left: 30%; right: 30%; bottom: 7%; height: 1px; }
.pdbp-dim-h i { left: 0; right: 0; top: 0; height: 1px; }
.pdbp-dim-h i::before, .pdbp-dim-h i::after { content: ''; position: absolute; top: -3px; width: 1px; height: 7px; background: rgba(125, 211, 252, 0.6); }
.pdbp-dim-h i::before { left: 0; }
.pdbp-dim-h i::after { right: 0; }
.pdbp-dim-h b { left: 50%; top: -12px; transform: translateX(-50%); }
.pdbp-dim-v { top: 30%; bottom: 24%; right: 5%; width: 1px; }
.pdbp-dim-v i { top: 0; bottom: 0; left: 0; width: 1px; }
.pdbp-dim-v i::before, .pdbp-dim-v i::after { content: ''; position: absolute; left: -3px; width: 7px; height: 1px; background: rgba(125, 211, 252, 0.6); }
.pdbp-dim-v i::before { top: 0; }
.pdbp-dim-v i::after { bottom: 0; }
.pdbp-dim-v b { left: 6px; top: 50%; writing-mode: vertical-rl; transform: translateY(-50%); }
.pdbp-note { position: absolute; display: grid; place-items: center; width: 14px; height: 14px; border: 1px solid rgba(125, 211, 252, 0.7); border-radius: 999px; font-family: ui-monospace, monospace; font-size: 7px; color: #7dd3fc; }
.pdbp-note-a { left: 18%; top: 18%; }
.pdbp-note-b { right: 18%; top: 34%; }
.pdbp-note-c { left: 26%; bottom: 24%; }
.pdbp-leader { position: absolute; width: 46px; height: 1px; background: rgba(125, 211, 252, 0.45); }
.pdbp-leader-a { left: 21%; top: 24%; transform: rotate(18deg); transform-origin: left; }
.pdbp-leader-b { right: 21%; top: 39%; transform: rotate(162deg); transform-origin: left; }
.pdbp-draw-label { position: absolute; left: 8px; top: 7px; font-family: ui-monospace, monospace; font-size: 7px; letter-spacing: 0.14em; color: rgba(125, 211, 252, 0.7); }
.pdbp-tol { position: absolute; right: 8px; bottom: 7px; font-family: ui-monospace, monospace; font-size: 7px; color: rgba(215, 239, 255, 0.4); }
.pdbp-price-cell { padding: 9px 10px; }
.pdbp-buy { background: rgba(125, 211, 252, 0.1); color: #d7efff; transition: background 0.25s ease, color 0.25s ease, transform 0.16s ease; }
.pdbp-buy:hover { background: #7dd3fc; color: #0e3a5c; }
.pdbp-buy:active { transform: translateY(1px); }
.pdbp-in { animation: pdbpRise 0.75s cubic-bezier(0.22, 1, 0.36, 1) both; }
.pdbp-d1 { animation-delay: 0.06s; }
.pdbp-d2 { animation-delay: 0.13s; }
.pdbp-d3 { animation-delay: 0.2s; }
.pdbp-d4 { animation-delay: 0.27s; }
@keyframes pdbpRise { from { opacity: 0; transform: translateY(16px); } to { opacity: 1; transform: none; } }
@keyframes pdbpIn { from { opacity: 0; transform: translate(-50%, -46%); } to { opacity: 1; transform: translate(-50%, -50%); } }
@media (prefers-reduced-motion: reduce) {
  .pdbp-in, .pdbp-product { animation: none !important; }
  .pdbp-buy { transition: none; }
}
</style>
