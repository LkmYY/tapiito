<template>
  <div class="pdbh-root min-h-screen bg-[#f1ece1] text-[#141414]">
    <div class="mx-auto w-full max-w-md px-5 pb-14 pt-5">
      <header class="pdbh-in flex items-center justify-between border-b-2 border-[#141414] pb-2">
        <span class="text-[13px] font-black uppercase tracking-[0.16em]">Formwerk</span>
        <span class="font-mono text-[9px]">BAUHAUS / {{ nfcId || '09' }}</span>
      </header>

      <section class="pdbh-in pdbh-d1 pdbh-stage mt-4">
        <span class="pdbh-block pdbh-red"></span>
        <span class="pdbh-block pdbh-blue"></span>
        <span class="pdbh-block pdbh-yellow"></span>
        <div class="pdbh-circle"></div>
        <div class="pdbh-object">
          <span class="pdbh-obj-bar"></span>
          <span class="pdbh-obj-disc"></span>
          <span class="pdbh-obj-stem"></span>
        </div>
        <span class="pdbh-num">09</span>
        <span class="pdbh-caption">KONSTRUKTION · 1926 / 2026</span>
      </section>

      <div class="pdbh-in pdbh-d2 mt-5 grid grid-cols-12 gap-3">
        <h1 class="col-span-8 text-[38px] font-black uppercase leading-[0.9] tracking-[-0.03em]">{{ p.title }}</h1>
        <div class="col-span-4 flex items-end justify-end">
          <span class="font-mono text-[9px] leading-tight">DESIGNED<br />FOR USE</span>
        </div>
        <p class="col-span-12 text-[12px] leading-relaxed text-[#141414]/65">{{ p.subtitle || '—' }}</p>
      </div>

      <div class="pdbh-in pdbh-d2 mt-4 grid grid-cols-3 border-2 border-[#141414]">
        <div class="col-span-2 bg-[#141414] px-3 py-3 text-[#f1ece1]">
          <p class="font-mono text-[8px] uppercase tracking-[0.18em]">PREIS / PRICE</p>
          <p class="mt-1 text-[34px] font-black leading-none tabular-nums">¥{{ p.price || '—' }}</p>
        </div>
        <div class="flex flex-col justify-center bg-[#e63b2e] px-3 py-3 text-[#f1ece1]">
          <p class="font-mono text-[8px] uppercase tracking-[0.1em]">WAS</p>
          <p class="mt-0.5 text-[15px] font-bold tabular-nums line-through">¥{{ p.originalPrice || '—' }}</p>
          <p class="mt-1 font-mono text-[8px] leading-tight">{{ discount ? '−' + discount + '%' : 'FIXPREIS' }}</p>
        </div>
      </div>

      <section class="pdbh-in pdbh-d3 mt-5 grid grid-cols-12 gap-4">
        <div class="col-span-1">
          <span class="block h-8 w-8 rounded-full bg-[#f2c230]"></span>
        </div>
        <div class="col-span-11">
          <h2 class="font-mono text-[9px] uppercase tracking-[0.22em]">FUNKTION</h2>
          <ul class="mt-1">
            <li v-for="(f, i) in features" :key="i" class="flex items-start gap-2 border-b border-[#141414]/20 py-2">
              <span class="mt-1 h-2 w-2 flex-none rounded-full" :class="['bg-[#e63b2e]', 'bg-[#1d4ed8]', 'bg-[#f2c230]'][i % 3]"></span>
              <span class="text-[12px] leading-snug">{{ f }}</span>
            </li>
            <li v-if="!features.length" class="border-b border-[#141414]/20 py-2 text-[12px] text-[#141414]/40"></li>
          </ul>
        </div>
      </section>

      <section class="pdbh-in pdbh-d3 mt-5 grid grid-cols-12 gap-4">
        <div class="col-span-4">
          <span class="block h-16 w-full bg-[#1d4ed8]"></span>
          <p class="mt-1 font-mono text-[8px] uppercase tracking-[0.14em]">MASS / UNIT</p>
        </div>
        <div class="col-span-8">
          <dl class="grid grid-cols-2 gap-x-4">
            <div v-for="[k, v] in specs" :key="k" class="border-b border-[#141414]/20 py-2">
              <dt class="font-mono text-[8px] uppercase text-[#141414]/45">{{ k }}</dt>
              <dd class="text-[12px] font-bold">{{ v }}</dd>
            </div>
            <div v-if="!specs.length" class="col-span-2 border-b border-[#141414]/20 py-2 text-[12px] text-[#141414]/40"></div>
          </dl>
        </div>
      </section>

      <p v-if="p.description" class="pdbh-in pdbh-d3 mt-4 text-[11.5px] leading-[1.8] text-[#141414]/60">{{ p.description }}</p>

      <button class="pdbh-in pdbh-d4 pdbh-buy mt-5 flex w-full items-center justify-between border-2 border-[#141414] px-4 py-3.5" @click="buy">
        <span class="text-[13px] font-black uppercase tracking-[0.08em]">{{ p.ctaText || '立即购买' }}</span>
        <span class="grid h-7 w-7 place-items-center rounded-full bg-[#f2c230]"><ArrowRight class="h-4 w-4" /></span>
      </button>

      <p class="mt-4 text-center font-mono text-[8px] uppercase tracking-[0.24em] text-[#141414]/40">Form follows function · Formwerk</p>
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
.pdbh-stage { position: relative; aspect-ratio: 1 / 1.02; overflow: hidden; border: 2px solid #141414; background: #e8e1d3; }
.pdbh-block { position: absolute; display: block; }
.pdbh-red { left: 0; top: 0; width: 34%; height: 34%; background: #e63b2e; }
.pdbh-blue { right: 0; bottom: 0; width: 28%; height: 26%; background: #1d4ed8; }
.pdbh-yellow { left: 10%; bottom: 0; width: 34%; height: 9%; background: #f2c230; }
.pdbh-circle { position: absolute; left: 50%; top: 50%; width: 62%; aspect-ratio: 1; transform: translate(-50%, -50%); border-radius: 999px; background: #f2c230; }
.pdbh-object { position: absolute; left: 50%; top: 50%; width: 40%; height: 66%; transform: translate(-50%, -50%) rotate(-18deg); animation: pdbhTurn 9s ease-in-out infinite; }
.pdbh-obj-bar { position: absolute; left: 6%; top: 8%; width: 88%; height: 12%; background: #141414; border-radius: 999px; }
.pdbh-obj-disc { position: absolute; left: 50%; top: 16%; width: 78%; aspect-ratio: 1; transform: translateX(-50%); border-radius: 999px; background: #f7f2e7; border: 10px solid #141414; box-shadow: inset 0 0 0 6px #e63b2e; }
.pdbh-obj-disc::after { content: ''; position: absolute; left: 50%; top: 50%; width: 22%; aspect-ratio: 1; transform: translate(-50%, -50%); border-radius: 999px; background: #141414; }
.pdbh-obj-stem { position: absolute; left: 50%; bottom: 0; width: 14%; height: 44%; transform: translateX(-50%); background: #141414; }
.pdbh-num { position: absolute; right: 8px; bottom: 6px; font-size: 72px; font-weight: 900; line-height: 0.8; color: rgba(20, 20, 20, 0.16); letter-spacing: -0.06em; }
.pdbh-caption { position: absolute; left: 8px; bottom: 8px; font-family: ui-monospace, monospace; font-size: 7px; letter-spacing: 0.18em; color: #141414; }
.pdbh-buy { background: #141414; color: #f1ece1; transition: background 0.2s ease, letter-spacing 0.25s ease; }
.pdbh-buy:hover { background: #1d4ed8; letter-spacing: 0.14em; }
.pdbh-buy:active { transform: scale(0.99); }
.pdbh-in { animation: pdbhIn 0.7s cubic-bezier(0.22, 1, 0.36, 1) both; }
.pdbh-d1 { animation-delay: 0.06s; }
.pdbh-d2 { animation-delay: 0.12s; }
.pdbh-d3 { animation-delay: 0.18s; }
.pdbh-d4 { animation-delay: 0.24s; }
@keyframes pdbhIn { from { opacity: 0; transform: translateY(14px); } to { opacity: 1; transform: none; } }
@keyframes pdbhTurn { 0%, 100% { transform: translate(-50%, -50%) rotate(-18deg); } 50% { transform: translate(-50%, -50%) rotate(-10deg); } }
@media (prefers-reduced-motion: reduce) {
  .pdbh-in, .pdbh-object { animation: none !important; }
  .pdbh-buy { transition: none; }
}
</style>
