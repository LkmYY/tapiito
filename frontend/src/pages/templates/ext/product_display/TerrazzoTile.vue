<template>
  <div class="pdz-root min-h-screen bg-[#f7f1e7] text-[#3c3340]">
    <div class="mx-auto w-full max-w-md px-4 pb-14 pt-5">
      <header class="pdz-in flex items-center justify-between rounded-2xl border-2 border-[#3c3340] bg-white px-3 py-2">
        <span class="text-[13px] font-black uppercase tracking-[0.1em]">Tessera</span>
        <span class="font-mono text-[8px] uppercase tracking-[0.14em] text-[#3c3340]/50">STONE OBJECTS / {{ nfcId || '№14' }}</span>
      </header>

      <section class="pdz-in pdz-d1 pdz-grid mt-4">
        <div class="pdz-tile pdz-tile-a">
          <span class="pdz-chip pdz-chip-1"></span>
          <span class="pdz-chip pdz-chip-2"></span>
          <span class="pdz-chip pdz-chip-3"></span>
          <span class="pdz-chip pdz-chip-4"></span>
          <span class="pdz-grid-label">MATERIAL<br />STUDY</span>
        </div>
        <div class="pdz-tile pdz-tile-b">
          <div class="pdz-obj">
            <span class="pdz-obj-top"></span>
            <span class="pdz-obj-body"></span>
            <span class="pdz-obj-dot"></span>
          </div>
        </div>
        <div class="pdz-tile pdz-tile-c">
          <span class="pdz-bigprice">¥{{ p.price || '—' }}</span>
          <span class="pdz-bigprice-label">TAX IN</span>
        </div>
        <div class="pdz-tile pdz-tile-d">
          <span class="flex items-center gap-1.5 font-mono text-[8px] uppercase tracking-[0.12em]"><Sparkles class="h-3 w-3" /> NEW STONE</span>
          <span class="mt-auto text-[22px] font-black leading-none">14</span>
        </div>
      </section>

      <div class="pdz-in pdz-d2 mt-5 flex items-start justify-between gap-3">
        <div>
          <p class="font-mono text-[8.5px] uppercase tracking-[0.2em] text-[#9d5c63]">{{ p.subtitle || '—' }}</p>
          <h1 class="mt-1 text-[33px] font-black leading-[0.98] tracking-[-0.03em]">{{ p.title }}</h1>
        </div>
        <div class="pdz-price-card">
          <span class="font-mono text-[7.5px] uppercase tracking-[0.1em]">PRICE</span>
          <span class="text-[17px] font-black leading-none">¥{{ p.price || '—' }}</span>
          <span v-if="p.originalPrice" class="font-mono text-[8px] text-[#3c3340]/40 line-through">¥{{ p.originalPrice }}</span>
          <span v-if="discount" class="mt-0.5 rounded-full bg-[#9d5c63] px-1.5 font-mono text-[7px] text-white">−{{ discount }}%</span>
        </div>
      </div>

      <section class="pdz-in pdz-d2 mt-4 grid grid-cols-2 gap-2">
        <div v-for="(f, i) in features" :key="i" class="pdz-feature" :class="'pdz-f' + (i % 3)">
          <span class="font-mono text-[8px] opacity-50">0{{ i + 1 }}</span>
          <span class="text-[11px] font-semibold leading-snug">{{ f }}</span>
        </div>
        <div v-if="!features.length" class="pdz-feature col-span-2 text-[11px]"></div>
      </section>

      <section class="pdz-in pdz-d3 mt-3 rounded-2xl border-2 border-[#3c3340] bg-white p-3">
        <h2 class="flex items-center justify-between font-mono text-[8.5px] uppercase tracking-[0.18em]">
          <span>SPEC TILES</span><span class="text-[#3c3340]/35">04</span>
        </h2>
        <dl class="mt-2 grid grid-cols-2 gap-1.5">
          <div v-for="[k, v] in specs" :key="k" class="rounded-xl bg-[#f3ead9] px-2.5 py-2">
            <dt class="font-mono text-[7px] uppercase tracking-[0.1em] text-[#3c3340]/45">{{ k }}</dt>
            <dd class="mt-0.5 text-[10.5px] font-bold">{{ v }}</dd>
          </div>
          <div v-if="!specs.length" class="col-span-2 rounded-xl bg-[#f3ead9] px-2.5 py-2 text-[11px] text-[#3c3340]/40"></div>
        </dl>
      </section>

      <p v-if="p.description" class="pdz-in pdz-d3 mt-3 rounded-2xl bg-[#e9dfcd] px-3.5 py-3 text-[11px] leading-[1.8] text-[#3c3340]/65">{{ p.description }}</p>

      <button class="pdz-in pdz-d4 pdz-buy mt-4 flex w-full items-center justify-between rounded-2xl border-2 border-[#3c3340] px-4 py-3.5" @click="buy">
        <span class="text-[12.5px] font-black uppercase tracking-[0.05em]">{{ p.ctaText || '立即购买' }}</span>
        <span class="flex items-center gap-2 font-mono text-[8px] uppercase tracking-[0.12em]">TAKE A PIECE <ArrowUpRight class="h-4 w-4" /></span>
      </button>

      <p class="mt-4 text-center font-mono text-[7.5px] uppercase tracking-[0.22em] text-[#3c3340]/35">Tessera · stone & glass · Mediterranean</p>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { ArrowUpRight, Sparkles } from 'lucide-vue-next'

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
.pdz-grid { display: grid; grid-template-columns: 1.2fr 0.8fr; grid-template-rows: 150px 92px; gap: 8px; }
.pdz-tile { position: relative; overflow: hidden; border: 2px solid #3c3340; border-radius: 16px; padding: 10px; }
.pdz-tile-a { background: #efe3cf; }
.pdz-tile-b { background: #e7d9df; grid-row: span 2; }
.pdz-tile-c { background: #cfdcd2; display: flex; flex-direction: column; justify-content: center; }
.pdz-tile-d { background: #f2c9a8; display: flex; flex-direction: column; }
.pdz-chip { position: absolute; border-radius: 999px; opacity: 0.9; }
.pdz-chip-1 { left: 12%; top: 20%; width: 34px; height: 24px; background: #9d5c63; transform: rotate(24deg); }
.pdz-chip-2 { right: 14%; top: 30%; width: 22px; height: 22px; background: #4f6d5a; }
.pdz-chip-3 { left: 30%; bottom: 16%; width: 18px; height: 26px; background: #d9a441; transform: rotate(-30deg); }
.pdz-chip-4 { right: 26%; bottom: 24%; width: 12px; height: 12px; background: #3c3340; }
.pdz-grid-label { position: absolute; left: 10px; bottom: 8px; font-family: ui-monospace, monospace; font-size: 7px; line-height: 1.5; letter-spacing: 0.16em; color: rgba(60, 51, 64, 0.55); }
.pdz-obj { position: absolute; left: 50%; top: 50%; width: 58%; height: 74%; transform: translate(-50%, -50%); animation: pdzFloat 6s ease-in-out infinite; }
.pdz-obj-top { position: absolute; left: 26%; top: 0; width: 48%; height: 10%; border-radius: 999px; background: #9d5c63; }
.pdz-obj-body { position: absolute; inset: 8% 0 0; border-radius: 36% 36% 24% 24% / 16% 16% 14% 14%; background: linear-gradient(100deg, #fdf8ef, #d9c9b5 55%, #b5a18a); box-shadow: inset -10px 0 18px rgba(60, 51, 64, 0.2), 8px 12px 18px -12px rgba(60, 51, 64, 0.5); }
.pdz-obj-dot { position: absolute; left: 50%; top: 46%; width: 34%; aspect-ratio: 1; transform: translate(-50%, -50%); border-radius: 999px; background: #9d5c63; }
.pdz-bigprice { font-size: 30px; font-weight: 900; letter-spacing: -0.05em; line-height: 1; }
.pdz-bigprice-label { margin-top: 4px; font-family: ui-monospace, monospace; font-size: 7px; letter-spacing: 0.18em; color: rgba(60, 51, 64, 0.45); }
.pdz-price-card { display: flex; flex-direction: column; align-items: flex-end; gap: 1px; flex: none; border: 2px solid #3c3340; border-radius: 12px; background: #f2c9a8; padding: 7px 10px; }
.pdz-feature { display: flex; flex-direction: column; gap: 2px; min-height: 56px; border: 2px solid #3c3340; border-radius: 12px; padding: 8px 10px; }
.pdz-f0 { background: #f2c9a8; }
.pdz-f1 { background: #cfdcd2; }
.pdz-f2 { background: #e7d9df; }
.pdz-buy { background: #3c3340; color: #f7f1e7; transition: background 0.25s ease, transform 0.16s ease; }
.pdz-buy:hover { background: #9d5c63; }
.pdz-buy:active { transform: translateY(1px); }
.pdz-in { animation: pdzIn 0.7s cubic-bezier(0.22, 1, 0.36, 1) both; }
.pdz-d1 { animation-delay: 0.06s; }
.pdz-d2 { animation-delay: 0.13s; }
.pdz-d3 { animation-delay: 0.2s; }
.pdz-d4 { animation-delay: 0.27s; }
@keyframes pdzIn { from { opacity: 0; transform: translateY(14px); } to { opacity: 1; transform: none; } }
@keyframes pdzFloat { 0%, 100% { transform: translate(-50%, -50%); } 50% { transform: translate(-50%, -53%); } }
@media (prefers-reduced-motion: reduce) {
  .pdz-in, .pdz-obj { animation: none !important; }
  .pdz-buy { transition: none; }
}
</style>
