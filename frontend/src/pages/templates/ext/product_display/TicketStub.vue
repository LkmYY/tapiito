<template>
  <div class="pdts-root min-h-screen bg-[#e7e2d6] text-[#18201c]">
    <div class="mx-auto w-full max-w-md px-5 pb-14 pt-7">
      <p class="pdts-in text-center font-mono text-[9px] uppercase tracking-[0.36em] text-[#18201c]/40">Admit one · good object</p>

      <section class="pdts-in pdts-d1 pdts-ticket mt-4">
        <div class="pdts-head">
          <div class="flex items-center justify-between">
            <span class="font-mono text-[9px] uppercase tracking-[0.2em]">No. {{ nfcId || '000128' }}</span>
            <span class="font-mono text-[9px] uppercase tracking-[0.2em]">2026 / 10 / 20</span>
          </div>
          <h1 class="pdts-title mt-5 text-center text-[38px] font-black uppercase leading-[0.95]">{{ p.title }}</h1>
          <p class="mt-2 text-center font-mono text-[9px] uppercase tracking-[0.18em] text-[#18201c]/50">{{ p.subtitle || '—' }}</p>
        </div>

        <div class="pdts-perf"></div>

        <div class="pdts-body">
          <div class="pdts-stub">
            <span class="pdts-stub-label">PRICE</span>
            <span class="pdts-stub-price">¥{{ p.price || '—' }}</span>
            <span v-if="p.originalPrice" class="pdts-stub-old">¥{{ p.originalPrice }}</span>
            <span v-if="discount" class="pdts-stub-off">−{{ discount }}%</span>
          </div>
          <div class="pdts-art">
            <div class="pdts-rings"><i></i><i></i><i></i></div>
            <div class="pdts-obj">
              <span class="pdts-obj-lid"></span>
              <span class="pdts-obj-body"></span>
              <span class="pdts-obj-tag">A1</span>
            </div>
            <span class="pdts-art-meta">SEAT 07 · ROW B</span>
          </div>
        </div>

        <div class="pdts-perf pdts-perf-dash"></div>

        <div class="pdts-info">
          <div class="flex items-center justify-between">
            <span class="font-mono text-[8px] uppercase tracking-[0.2em] text-[#18201c]/40">ITEM DETAILS</span>
            <span class="font-mono text-[8px] text-[#18201c]/40">×1</span>
          </div>
          <ul class="mt-2 grid grid-cols-2 gap-x-4">
            <li v-for="(f, i) in features" :key="i" class="flex gap-1.5 border-b border-[#18201c]/15 py-1.5 text-[11px] leading-snug">
              <span class="text-[#c2410c]">✓</span>{{ f }}
            </li>
            <li v-if="!features.length" class="col-span-2 border-b border-[#18201c]/15 py-1.5 text-[11px] text-[#18201c]/40"></li>
          </ul>
          <p v-if="p.description" class="mt-3 text-[11px] leading-[1.8] text-[#18201c]/55">{{ p.description }}</p>
        </div>

        <div class="pdts-perf"></div>

        <div class="pdts-specs">
          <div v-for="[k, v] in specs" :key="k" class="flex items-baseline justify-between border-b border-dashed border-[#18201c]/20 py-1.5">
            <span class="font-mono text-[9px] uppercase tracking-[0.12em] text-[#18201c]/45">{{ k }}</span>
            <span class="font-mono text-[10px]">{{ v }}</span>
          </div>
          <div v-if="!specs.length" class="border-b border-dashed border-[#18201c]/20 py-1.5 font-mono text-[10px] text-[#18201c]/40"></div>
        </div>

        <div class="pdts-barcode-row">
          <span class="pdts-barcode" aria-hidden="true"><i v-for="n in 32" :key="n" :style="{ width: (n % 4) + 1 + 'px' }"></i></span>
          <span class="font-mono text-[8px] tracking-[0.18em] text-[#18201c]/45">{{ nfcId || '20261020' }}</span>
        </div>
      </section>

      <button class="pdts-in pdts-d2 pdts-buy mt-4 flex w-full items-center justify-between px-5 py-4" @click="buy">
        <span class="font-mono text-[11px] font-bold uppercase tracking-[0.2em]">{{ p.ctaText || '立即购买' }}</span>
        <span class="flex items-center gap-2"><span class="font-mono text-[9px] uppercase tracking-[0.1em]">TEAR HERE</span><Scissors class="h-4 w-4" /></span>
      </button>

      <p class="mt-5 text-center font-mono text-[8px] uppercase tracking-[0.28em] text-[#18201c]/35">No refunds · keep this stub · object gallery 07</p>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { Scissors } from 'lucide-vue-next'

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
.pdts-ticket {
  position: relative;
  overflow: hidden;
  border: 1px solid rgba(24, 32, 28, 0.35);
  background: #f3efe4;
  box-shadow: 0 22px 40px -30px rgba(24, 32, 28, 0.65);
}
.pdts-ticket::before,
.pdts-ticket::after {
  content: '';
  position: absolute;
  top: 58.5%;
  width: 22px;
  height: 22px;
  border-radius: 999px;
  background: #e7e2d6;
  border: 1px solid rgba(24, 32, 28, 0.35);
  z-index: 2;
}
.pdts-ticket::before { left: -12px; }
.pdts-ticket::after { right: -12px; }
.pdts-head { padding: 16px 18px 18px; }
.pdts-title { font-family: 'Arial Black', 'Helvetica Neue', sans-serif; letter-spacing: -0.03em; }
.pdts-perf { height: 1px; margin-inline: 14px; background-image: linear-gradient(90deg, rgba(24, 32, 28, 0.45) 50%, transparent 0); background-size: 9px 1px; }
.pdts-perf-dash { opacity: 0.6; }
.pdts-body { display: flex; min-height: 190px; }
.pdts-stub {
  width: 34%;
  flex: none;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  gap: 2px;
  border-right: 1px dashed rgba(24, 32, 28, 0.35);
  background: rgba(194, 65, 12, 0.06);
}
.pdts-stub-label { font-family: ui-monospace, monospace; font-size: 8px; letter-spacing: 0.2em; color: rgba(24, 32, 28, 0.45); }
.pdts-stub-price { margin-top: 2px; font-size: 28px; font-weight: 900; letter-spacing: -0.05em; line-height: 1; }
.pdts-stub-old { font-family: ui-monospace, monospace; font-size: 9px; text-decoration: line-through; color: rgba(24, 32, 28, 0.35); }
.pdts-stub-off { margin-top: 4px; background: #c2410c; color: #f3efe4; padding: 1px 6px; font-family: ui-monospace, monospace; font-size: 8px; }
.pdts-art { position: relative; flex: 1; overflow: hidden; background: linear-gradient(150deg, #d7d2c3, #c3bdab); }
.pdts-rings { position: absolute; inset: 0; display: grid; place-items: center; }
.pdts-rings i { position: absolute; border: 1px solid rgba(24, 32, 28, 0.15); border-radius: 999px; }
.pdts-rings i:nth-child(1) { width: 58%; aspect-ratio: 1; }
.pdts-rings i:nth-child(2) { width: 86%; aspect-ratio: 1; }
.pdts-rings i:nth-child(3) { width: 116%; aspect-ratio: 1; }
.pdts-obj { position: absolute; left: 50%; top: 50%; width: 42%; height: 60%; transform: translate(-50%, -50%) rotate(3deg); animation: pdtsFloat 6s ease-in-out infinite; }
.pdts-obj-lid { position: absolute; left: 18%; top: 0; width: 64%; height: 12%; border-radius: 3px 3px 0 0; background: #18201c; }
.pdts-obj-body {
  position: absolute;
  inset: 11% 0 0;
  border-radius: 6px 6px 16px 16px;
  background: linear-gradient(100deg, #faf6e9, #ded7c4 55%, #b9b19c);
  box-shadow: inset -8px 0 14px rgba(24, 32, 28, 0.18), 8px 10px 0 rgba(24, 32, 28, 0.15);
}
.pdts-obj-tag { position: absolute; left: 50%; top: 56%; transform: translate(-50%, -50%); display: grid; place-items: center; width: 42%; height: 26%; border: 1px solid #18201c; font-family: ui-monospace, monospace; font-size: 13px; font-weight: 700; }
.pdts-art-meta { position: absolute; left: 10px; bottom: 7px; font-family: ui-monospace, monospace; font-size: 7px; letter-spacing: 0.16em; color: rgba(24, 32, 28, 0.55); }
.pdts-info { padding: 12px 18px 14px; }
.pdts-specs { padding: 10px 18px; }
.pdts-barcode-row { display: flex; align-items: center; justify-content: space-between; gap: 10px; padding: 10px 18px 16px; }
.pdts-barcode { display: flex; align-items: stretch; gap: 2px; height: 26px; flex: 1; }
.pdts-barcode i { background: #18201c; display: block; }
.pdts-buy { background: #18201c; color: #f3efe4; transition: background 0.25s ease, gap 0.2s ease, transform 0.16s ease; }
.pdts-buy:hover { background: #c2410c; }
.pdts-buy:active { transform: translateY(1px); }
.pdts-in { animation: pdtsIn 0.7s cubic-bezier(0.22, 1, 0.36, 1) both; }
.pdts-d1 { animation-delay: 0.08s; }
.pdts-d2 { animation-delay: 0.16s; }
@keyframes pdtsIn { from { opacity: 0; transform: translateY(16px); } to { opacity: 1; transform: none; } }
@keyframes pdtsFloat { 0%, 100% { transform: translate(-50%, -50%) rotate(3deg); } 50% { transform: translate(-50%, -53%) rotate(0deg); } }
@media (prefers-reduced-motion: reduce) {
  .pdts-in, .pdts-obj { animation: none !important; }
  .pdts-buy { transition: none; }
}
</style>
