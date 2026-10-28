<template>
  <div class="pdf-root min-h-screen bg-[#dfe8f4] text-[#10233c]">
    <div class="pdf-blob pdf-blob-a"></div>
    <div class="pdf-blob pdf-blob-b"></div>
    <div class="relative mx-auto w-full max-w-md px-5 pb-14 pt-6">
      <header class="pdf-in flex items-center justify-between rounded-full border border-white/70 bg-white/35 px-4 py-2 font-mono text-[9px] uppercase tracking-[0.2em] text-[#10233c]/60 backdrop-blur-md">
        <span>Aero Studio</span>
        <span>{{ nfcId ? 'NFC ' + nfcId : 'Clear series' }}</span>
      </header>

      <section class="pdf-in pdf-d1 pdf-stage mt-5">
        <div class="pdf-shine"></div>
        <div class="pdf-vessel">
          <div class="pdf-vessel-top"></div>
          <div class="pdf-vessel-body"></div>
          <div class="pdf-vessel-glow"></div>
          <span class="pdf-vessel-label">AERO<br /><b>01</b></span>
        </div>
        <div class="pdf-stage-meta">
          <span>TRANSPARENT FORM</span>
          <span>2026 / EDITION</span>
        </div>
      </section>

      <div class="pdf-in pdf-d2 pdf-glass mt-5 px-5 py-5">
        <div class="flex items-start justify-between gap-4">
          <div class="min-w-0">
            <p class="font-mono text-[9px] uppercase tracking-[0.22em] text-[#0f766e]">The clear object</p>
            <h1 class="mt-2 text-[30px] font-light leading-[1.06] tracking-[-0.03em]">{{ p.title }}</h1>
            <p class="mt-2 text-[12px] leading-relaxed text-[#10233c]/55">{{ p.subtitle || '—' }}</p>
          </div>
          <div class="pdf-price">
            <span class="font-mono text-[8px] uppercase tracking-[0.16em] text-[#10233c]/40">Price</span>
            <span class="mt-1 block text-[24px] font-medium leading-none tabular-nums">¥{{ p.price || '—' }}</span>
            <span v-if="p.originalPrice" class="mt-1 block font-mono text-[9px] text-[#10233c]/35 line-through">¥{{ p.originalPrice }}</span>
          </div>
        </div>
        <div v-if="discount" class="mt-4 inline-flex items-center gap-1.5 rounded-full bg-[#0f766e] px-3 py-1 font-mono text-[9px] uppercase tracking-[0.14em] text-white">
          <Sparkles class="h-3 w-3" /> Save {{ discount }}%
        </div>
      </div>

      <section class="pdf-in pdf-d3 pdf-glass mt-3 px-5 py-5">
        <div class="flex items-center justify-between">
          <h2 class="font-mono text-[9px] uppercase tracking-[0.24em] text-[#10233c]/50">Inside the object</h2>
          <span class="font-mono text-[9px] text-[#10233c]/30">{{ String(features.length).padStart(2, '0') }}</span>
        </div>
        <ul class="mt-3 flex flex-wrap gap-2">
          <li v-for="(f, i) in features" :key="i" class="pdf-pill">
            <span class="pdf-pill-dot"></span>{{ f }}
          </li>
          <li v-if="!features.length" class="text-[12px] text-[#10233c]/40"></li>
        </ul>
        <p v-if="p.description" class="mt-4 border-t border-white/70 pt-3 text-[12px] leading-[1.85] text-[#10233c]/60">{{ p.description }}</p>
      </section>

      <section class="pdf-in pdf-d3 pdf-glass mt-3 px-5 py-5">
        <h2 class="font-mono text-[9px] uppercase tracking-[0.24em] text-[#10233c]/50">Technical sheet</h2>
        <dl class="mt-3 grid grid-cols-2 gap-2">
          <div v-for="[k, v] in specs" :key="k" class="rounded-2xl border border-white/70 bg-white/35 px-3 py-2.5">
            <dt class="font-mono text-[8px] uppercase tracking-[0.12em] text-[#10233c]/40">{{ k }}</dt>
            <dd class="mt-1 text-[12px] font-medium">{{ v }}</dd>
          </div>
          <div v-if="!specs.length" class="col-span-2 rounded-2xl border border-white/70 bg-white/35 px-3 py-2.5 text-[12px] text-[#10233c]/40"></div>
        </dl>
      </section>

      <button class="pdf-in pdf-d4 pdf-buy mt-4 flex w-full items-center justify-center gap-2.5 py-4" @click="buy">
        <ShoppingBag class="h-4 w-4" />
        <span class="text-[13px] font-semibold tracking-wide">{{ p.ctaText || '立即购买' }}</span>
        <ArrowUpRight class="h-4 w-4" />
      </button>

      <p class="mt-5 text-center font-mono text-[8px] uppercase tracking-[0.28em] text-[#10233c]/30">Designed to disappear · Aero Studio</p>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { ArrowUpRight, ShoppingBag, Sparkles } from 'lucide-vue-next'

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
.pdf-root { position: relative; overflow: hidden; }
.pdf-blob { position: absolute; border-radius: 999px; filter: blur(40px); opacity: 0.55; pointer-events: none; }
.pdf-blob-a { width: 300px; height: 300px; top: -90px; left: -100px; background: #a5b4fc; animation: pdfDrift 16s ease-in-out infinite; }
.pdf-blob-b { width: 260px; height: 260px; top: 220px; right: -110px; background: #99f6e4; animation: pdfDrift 20s ease-in-out infinite reverse; }
.pdf-glass {
  border: 1px solid rgba(255, 255, 255, 0.8);
  border-radius: 26px;
  background: rgba(255, 255, 255, 0.42);
  box-shadow: 0 18px 40px -26px rgba(16, 35, 60, 0.45), inset 0 1px 0 rgba(255, 255, 255, 0.9);
  backdrop-filter: blur(16px);
}
.pdf-stage {
  position: relative;
  aspect-ratio: 1 / 1.02;
  overflow: hidden;
  border-radius: 30px;
  border: 1px solid rgba(255, 255, 255, 0.85);
  background:
    radial-gradient(60% 50% at 50% 44%, rgba(255, 255, 255, 0.9), transparent 72%),
    linear-gradient(150deg, rgba(255, 255, 255, 0.55), rgba(190, 214, 240, 0.4));
  box-shadow: inset 0 0 60px rgba(255, 255, 255, 0.5), 0 24px 50px -32px rgba(16, 35, 60, 0.5);
}
.pdf-shine {
  position: absolute;
  inset: 0;
  background: linear-gradient(115deg, transparent 24%, rgba(255, 255, 255, 0.55) 40%, transparent 58%);
  transform: translateX(-30%);
  animation: pdfShine 7s ease-in-out infinite;
}
.pdf-vessel { position: absolute; left: 50%; top: 50%; width: 42%; height: 60%; transform: translate(-50%, -50%); }
.pdf-vessel-top {
  position: absolute;
  left: 22%;
  top: -6%;
  width: 56%;
  height: 18%;
  border-radius: 10px 10px 4px 4px;
  background: linear-gradient(100deg, rgba(255, 255, 255, 0.9), rgba(173, 203, 232, 0.75));
  border: 1px solid rgba(255, 255, 255, 0.95);
  box-shadow: inset -6px 0 10px rgba(120, 160, 200, 0.25);
}
.pdf-vessel-body {
  position: absolute;
  inset: 10% 0 0;
  border-radius: 30% 30% 22% 22% / 14% 14% 18% 18%;
  background: linear-gradient(105deg, rgba(255, 255, 255, 0.92) 0%, rgba(207, 228, 247, 0.7) 42%, rgba(140, 182, 220, 0.55) 100%);
  border: 1px solid rgba(255, 255, 255, 0.9);
  box-shadow: inset -14px -8px 24px rgba(87, 132, 176, 0.28), inset 10px 8px 22px rgba(255, 255, 255, 0.8), 0 22px 34px -22px rgba(16, 35, 60, 0.55);
  backdrop-filter: blur(2px);
}
.pdf-vessel-glow {
  position: absolute;
  left: 50%;
  bottom: 2%;
  width: 90%;
  height: 14%;
  transform: translateX(-50%);
  border-radius: 50%;
  background: rgba(13, 148, 136, 0.28);
  filter: blur(14px);
}
.pdf-vessel-label {
  position: absolute;
  left: 50%;
  top: 52%;
  transform: translate(-50%, -50%);
  text-align: center;
  font-family: ui-monospace, monospace;
  font-size: 9px;
  line-height: 1.5;
  letter-spacing: 0.2em;
  color: rgba(16, 35, 60, 0.55);
}
.pdf-vessel-label b { font-size: 14px; letter-spacing: 0.1em; }
.pdf-stage-meta {
  position: absolute;
  inset-inline: 16px;
  bottom: 12px;
  display: flex;
  justify-content: space-between;
  font-family: ui-monospace, monospace;
  font-size: 7.5px;
  letter-spacing: 0.2em;
  color: rgba(16, 35, 60, 0.42);
}
.pdf-price { flex: none; text-align: right; }
.pdf-pill {
  display: inline-flex;
  align-items: center;
  gap: 7px;
  border-radius: 999px;
  border: 1px solid rgba(255, 255, 255, 0.9);
  background: rgba(255, 255, 255, 0.5);
  padding: 6px 11px;
  font-size: 11px;
  color: rgba(16, 35, 60, 0.72);
  transition: background 0.2s ease, transform 0.2s ease;
}
.pdf-pill:hover { background: rgba(255, 255, 255, 0.85); transform: translateY(-1px); }
.pdf-pill-dot { width: 4px; height: 4px; border-radius: 99px; background: #0d9488; flex: none; }
.pdf-buy {
  border-radius: 999px;
  background: linear-gradient(100deg, #0f766e, #0e7490);
  color: #f0fdfa;
  box-shadow: 0 16px 30px -16px rgba(15, 118, 110, 0.8), inset 0 1px 0 rgba(255, 255, 255, 0.35);
  transition: box-shadow 0.25s ease, transform 0.18s ease, filter 0.25s ease;
}
.pdf-buy:hover { filter: brightness(1.08); box-shadow: 0 20px 38px -16px rgba(15, 118, 110, 0.9); }
.pdf-buy:active { transform: scale(0.985); }
.pdf-in { animation: pdfIn 0.8s cubic-bezier(0.22, 1, 0.36, 1) both; }
.pdf-d1 { animation-delay: 0.06s; }
.pdf-d2 { animation-delay: 0.14s; }
.pdf-d3 { animation-delay: 0.22s; }
.pdf-d4 { animation-delay: 0.3s; }
@keyframes pdfIn { from { opacity: 0; transform: translateY(18px); } to { opacity: 1; transform: none; } }
@keyframes pdfShine { 0%, 100% { transform: translateX(-35%); } 50% { transform: translateX(35%); } }
@keyframes pdfDrift { 0%, 100% { transform: translate(0, 0); } 50% { transform: translate(18px, 24px); } }
@media (prefers-reduced-motion: reduce) {
  .pdf-in, .pdf-shine, .pdf-blob { animation: none !important; }
  .pdf-buy, .pdf-pill { transition: none; }
}
</style>
