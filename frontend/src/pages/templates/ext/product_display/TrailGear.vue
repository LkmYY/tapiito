<template>
  <div class="pdt-root min-h-screen bg-[#101b15] text-[#e6efe3]">
    <div class="pdt-topo pointer-events-none"></div>
    <div class="mx-auto w-full max-w-md px-5 pb-14 pt-6">
      <header class="pdt-in flex items-center justify-between border-b border-[#b9d17a]/25 pb-3 font-mono text-[9px] uppercase tracking-[0.22em] text-[#b9d17a]/80">
        <span class="flex items-center gap-2"><Mountain class="h-3.5 w-3.5" /> Field Supply</span>
        <span class="text-[#e6efe3]/35">{{ nfcId ? 'TAG ' + nfcId : 'FW 26' }}</span>
      </header>

      <section class="pdt-in pdt-d1 pdt-stage mt-5">
        <div class="pdt-grid"></div>
        <span class="pdt-coord">N 31°14' · E 121°29'</span>
        <span class="pdt-coord pdt-coord-r">ELEV 1,864 M</span>
        <div class="pdt-pack">
          <span class="pdt-pack-strap"></span>
          <span class="pdt-pack-body">
            <i class="pdt-pack-zip"></i>
            <i class="pdt-pack-label">FS<br /><b>32L</b></i>
            <i class="pdt-pack-buckle"></i>
          </span>
        </div>
        <span class="pdt-scale">SCALE 1:6<br />REF. 004</span>
      </section>

      <div class="pdt-in pdt-d2 mt-6 flex items-start justify-between gap-4">
        <div>
          <p class="font-mono text-[9px] uppercase tracking-[0.24em] text-[#b9d17a]">{{ p.subtitle || '—' }}</p>
          <h1 class="mt-2 text-[36px] font-semibold uppercase leading-[0.98] tracking-[-0.02em]">{{ p.title }}</h1>
        </div>
        <span class="mt-1 flex h-10 w-10 flex-none items-center justify-center rounded-full border border-[#b9d17a]/40 text-[#b9d17a]">
          <Compass class="h-5 w-5" />
        </span>
      </div>

      <div class="pdt-in pdt-d2 mt-5 flex items-end justify-between border-y border-[#b9d17a]/25 py-4">
        <div class="flex items-baseline gap-2">
          <span class="font-mono text-[10px] text-[#b9d17a]">CNY</span>
          <span class="text-[42px] font-semibold leading-none tabular-nums">{{ p.price || '—' }}</span>
        </div>
        <div class="text-right font-mono text-[9px] leading-relaxed">
          <p v-if="p.originalPrice" class="text-[#e6efe3]/35 line-through">¥{{ p.originalPrice }}</p>
          <p v-if="discount" class="text-[#b9d17a]">FIELD PRICE −{{ discount }}%</p>
          <p class="text-[#e6efe3]/30">VAT INCLUDED</p>
        </div>
      </div>

      <section class="pdt-in pdt-d3 mt-5">
        <div class="flex items-center justify-between font-mono text-[9px] uppercase tracking-[0.2em] text-[#b9d17a]/70">
          <span>01 / Field notes</span>
          <span>{{ String(features.length).padStart(2, '0') }}</span>
        </div>
        <ul class="mt-2">
          <li v-for="(f, i) in features" :key="i" class="flex items-start gap-3 border-b border-[#e6efe3]/10 py-3">
            <span class="mt-0.5 grid h-4 w-4 flex-none place-items-center border border-[#b9d17a]/50 font-mono text-[7px] text-[#b9d17a]">✓</span>
            <span class="flex-1 text-[12.5px] leading-snug text-[#e6efe3]/80">{{ f }}</span>
            <span class="font-mono text-[8px] text-[#e6efe3]/20">0{{ i + 1 }}</span>
          </li>
          <li v-if="!features.length" class="border-b border-[#e6efe3]/10 py-3 text-[12px] text-[#e6efe3]/35"></li>
        </ul>
      </section>

      <section class="pdt-in pdt-d3 mt-5">
        <div class="flex items-center justify-between font-mono text-[9px] uppercase tracking-[0.2em] text-[#b9d17a]/70">
          <span>02 / Technical</span>
          <span>SI UNITS</span>
        </div>
        <dl class="mt-2 border border-[#b9d17a]/25">
          <div v-for="([k, v], i) in specs" :key="k" :class="['grid grid-cols-[1fr_1.2fr] text-[11.5px]', i ? 'border-t border-[#b9d17a]/15' : '']">
            <dt class="border-r border-[#b9d17a]/15 px-3 py-2.5 font-mono uppercase tracking-[0.1em] text-[#e6efe3]/40">{{ k }}</dt>
            <dd class="px-3 py-2.5 text-[#e6efe3]/85">{{ v }}</dd>
          </div>
          <div v-if="!specs.length" class="px-3 py-2.5 text-[12px] text-[#e6efe3]/35"></div>
        </dl>
      </section>

      <p v-if="p.description" class="pdt-in pdt-d3 mt-4 text-[11.5px] leading-[1.85] text-[#e6efe3]/50">{{ p.description }}</p>

      <button class="pdt-in pdt-d4 pdt-buy mt-5 flex w-full items-center justify-between px-4 py-4" @click="buy">
        <span class="font-mono text-[10px] uppercase tracking-[0.22em]">{{ p.ctaText || '立即购买' }}</span>
        <span class="flex items-center gap-2 font-mono text-[9px] uppercase tracking-[0.16em]"><Tent class="h-4 w-4" /> Add to kit <ArrowUpRight class="h-3.5 w-3.5" /></span>
      </button>

      <p class="mt-5 text-center font-mono text-[8px] uppercase tracking-[0.24em] text-[#e6efe3]/25">Tested above the treeline · Field Supply Co.</p>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { ArrowUpRight, Compass, Mountain, Tent } from 'lucide-vue-next'

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
.pdt-root { position: relative; overflow: hidden; }
.pdt-topo {
  position: absolute;
  inset: 0;
  opacity: 0.5;
  background-image: repeating-radial-gradient(circle at 88% 8%, transparent 0 26px, rgba(185, 209, 122, 0.1) 27px 28px),
    repeating-radial-gradient(circle at 8% 92%, transparent 0 34px, rgba(185, 209, 122, 0.08) 35px 36px);
  pointer-events: none;
}
.pdt-stage {
  position: relative;
  aspect-ratio: 4 / 4.4;
  overflow: hidden;
  border: 1px solid rgba(185, 209, 122, 0.32);
  background: linear-gradient(155deg, #17251d, #0a130e 75%);
}
.pdt-grid {
  position: absolute;
  inset: 0;
  background-image: linear-gradient(rgba(185, 209, 122, 0.1) 1px, transparent 1px), linear-gradient(90deg, rgba(185, 209, 122, 0.1) 1px, transparent 1px);
  background-size: 26px 26px;
  mask-image: radial-gradient(circle at 50% 52%, black 30%, transparent 80%);
}
.pdt-coord { position: absolute; left: 10px; top: 9px; font-family: ui-monospace, monospace; font-size: 7px; letter-spacing: 0.16em; color: rgba(185, 209, 122, 0.6); }
.pdt-coord-r { left: auto; right: 10px; }
.pdt-pack { position: absolute; left: 50%; top: 52%; width: 46%; height: 58%; transform: translate(-50%, -50%); animation: pdtFloat 6s ease-in-out infinite; }
.pdt-pack-strap {
  position: absolute;
  left: 24%;
  top: -4%;
  width: 52%;
  height: 24%;
  border: 3px solid #2c3b32;
  border-bottom: 0;
  border-radius: 40px 40px 0 0;
  box-shadow: inset 0 0 0 1px rgba(185, 209, 122, 0.25);
}
.pdt-pack-body {
  position: absolute;
  inset: 12% 0 0;
  border-radius: 26% 26% 20% 20% / 18% 18% 14% 14%;
  background: linear-gradient(105deg, #5b6b4a, #37452f 42%, #1f2a21 100%);
  box-shadow: inset 0 0 0 1px rgba(185, 209, 122, 0.3), inset -14px -8px 26px rgba(0, 0, 0, 0.5), 0 24px 32px -20px rgba(0, 0, 0, 0.8);
}
.pdt-pack-zip {
  position: absolute;
  left: 10%;
  right: 10%;
  top: 30%;
  height: 2px;
  background: rgba(185, 209, 122, 0.5);
  box-shadow: 0 -6px 0 rgba(0, 0, 0, 0.25);
}
.pdt-pack-label {
  position: absolute;
  left: 50%;
  top: 52%;
  transform: translate(-50%, -50%);
  display: grid;
  place-items: center;
  width: 58%;
  padding: 7px 0;
  border: 1px solid rgba(185, 209, 122, 0.5);
  background: rgba(10, 19, 14, 0.55);
  font-family: ui-monospace, monospace;
  font-size: 8px;
  line-height: 1.4;
  letter-spacing: 0.18em;
  text-align: center;
  color: #d6e4c2;
}
.pdt-pack-label b { font-size: 11px; color: #b9d17a; }
.pdt-pack-buckle {
  position: absolute;
  left: 50%;
  top: 72%;
  width: 26%;
  height: 6%;
  transform: translateX(-50%);
  border: 2px solid #b9d17a;
  border-radius: 3px;
}
.pdt-scale { position: absolute; right: 12px; bottom: 10px; font-family: ui-monospace, monospace; font-size: 7px; line-height: 1.7; letter-spacing: 0.14em; text-align: right; color: rgba(230, 239, 227, 0.35); }
.pdt-buy {
  background: #b9d17a;
  color: #101b15;
  clip-path: polygon(0 0, calc(100% - 14px) 0, 100% 14px, 100% 100%, 0 100%);
  transition: background 0.25s ease, transform 0.16s ease;
}
.pdt-buy:hover { background: #d0e69a; }
.pdt-buy:active { transform: translateY(1px); }
.pdt-in { animation: pdtIn 0.7s cubic-bezier(0.22, 1, 0.36, 1) both; }
.pdt-d1 { animation-delay: 0.06s; }
.pdt-d2 { animation-delay: 0.12s; }
.pdt-d3 { animation-delay: 0.18s; }
.pdt-d4 { animation-delay: 0.24s; }
@keyframes pdtIn { from { opacity: 0; transform: translateY(16px); } to { opacity: 1; transform: none; } }
@keyframes pdtFloat { 0%, 100% { transform: translate(-50%, -50%); } 50% { transform: translate(-50%, -52%); } }
@media (prefers-reduced-motion: reduce) {
  .pdt-in, .pdt-pack { animation: none !important; }
  .pdt-buy { transition: none; }
}
</style>
