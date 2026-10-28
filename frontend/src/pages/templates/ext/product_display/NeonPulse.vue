<template>
  <div class="pdn-root min-h-screen bg-[#06050d] text-[#e8f7ff]">
    <div class="pdn-scan pointer-events-none"></div>
    <div class="mx-auto w-full max-w-md px-5 pb-16 pt-6">
      <header class="pdn-in flex items-center justify-between font-mono text-[10px] uppercase tracking-[0.2em] text-[#5eead4]/70">
        <span class="flex items-center gap-2"><i class="pdn-dot"></i>Pulse Audio</span>
        <span class="text-[#e8f7ff]/35">ID {{ nfcId || '0000' }}</span>
      </header>

      <section class="pdn-in pdn-d1 pdn-stage mt-6">
        <div class="pdn-lines"></div>
        <span class="pdn-hud pdn-hud-tl">MODE / ANC</span>
        <span class="pdn-hud pdn-hud-tr">45dB</span>
        <span class="pdn-hud pdn-hud-bl">BAT 40H</span>
        <span class="pdn-hud pdn-hud-br">LDAC</span>
        <div class="pdn-orb">
          <div class="pdn-ring pdn-ring-a"></div>
          <div class="pdn-ring pdn-ring-b"></div>
          <div class="pdn-core"></div>
          <div class="pdn-cup pdn-cup-l"></div>
          <div class="pdn-cup pdn-cup-r"></div>
        </div>
        <p class="pdn-stage-tag">NEXT-GEN AUDIO OBJECT</p>
      </section>

      <div class="pdn-in pdn-d2 mt-7">
        <p class="font-mono text-[10px] uppercase tracking-[0.28em] text-[#f472b6]">// {{ p.subtitle || '—' }}</p>
        <h1 class="pdn-glitch mt-3 text-[38px] font-black uppercase leading-[0.94] tracking-[-0.04em]">{{ p.title }}</h1>
      </div>

      <div class="pdn-in pdn-d2 mt-6 grid grid-cols-[1fr_auto] items-end gap-4 border-y border-[#5eead4]/25 py-5">
        <div>
          <p class="font-mono text-[9px] uppercase tracking-[0.24em] text-[#e8f7ff]/40">Unit price</p>
          <p class="mt-1 text-[44px] font-bold leading-none tabular-nums text-[#5eead4]"><span class="text-[20px] align-top">¥</span>{{ p.price || '—' }}</p>
        </div>
        <div class="text-right font-mono text-[10px] leading-relaxed text-[#e8f7ff]/45">
          <p v-if="p.originalPrice" class="line-through">¥{{ p.originalPrice }}</p>
          <p v-if="discount" class="mt-1 text-[#f472b6]">-{{ discount }}% OFF</p>
          <p class="mt-1 text-[#e8f7ff]/25">FREE SHIP</p>
        </div>
      </div>

      <section class="pdn-in pdn-d3 mt-6">
        <p class="font-mono text-[9px] uppercase tracking-[0.28em] text-[#5eead4]/60">01 · CORE FEATURES</p>
        <ul class="mt-3 grid grid-cols-2 gap-2">
          <li v-for="(f, i) in features" :key="i" class="pdn-chip">
            <span class="font-mono text-[8px] text-[#f472b6]">0{{ i + 1 }}</span>
            <span class="text-[11.5px] leading-snug">{{ f }}</span>
          </li>
          <li v-if="!features.length" class="pdn-chip col-span-2 text-[12px] text-[#e8f7ff]/35"></li>
        </ul>
      </section>

      <section class="pdn-in pdn-d4 mt-6">
        <p class="font-mono text-[9px] uppercase tracking-[0.28em] text-[#5eead4]/60">02 · SPEC MATRIX</p>
        <dl class="mt-3 border border-[#5eead4]/20 font-mono text-[11px]">
          <div v-for="([k, v], i) in specs" :key="k" :class="['flex items-center justify-between border-[#5eead4]/12 px-3 py-2.5', i ? 'border-t' : '']">
            <dt class="uppercase tracking-[0.14em] text-[#e8f7ff]/40">{{ k }}</dt>
            <dd class="text-[#e8f7ff]/90">{{ v }}</dd>
          </div>
          <div v-if="!specs.length" class="px-3 py-2.5 text-[#e8f7ff]/35"></div>
        </dl>
      </section>

      <p v-if="p.description" class="pdn-in pdn-d4 mt-6 text-[12px] leading-[1.9] text-[#e8f7ff]/55">{{ p.description }}</p>

      <button class="pdn-in pdn-d5 pdn-cta mt-7 flex w-full items-center justify-between px-5 py-4" @click="buy">
        <span class="font-mono text-[12px] font-bold uppercase tracking-[0.24em]">{{ p.ctaText || '立即购买' }}</span>
        <span class="pdn-cta-arrow"><ArrowUpRight class="h-4 w-4" /></span>
      </button>

      <p class="mt-5 text-center font-mono text-[8px] uppercase tracking-[0.32em] text-[#e8f7ff]/20">Encrypted checkout · NFC verified</p>
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
.pdn-root { position: relative; overflow: hidden; }
.pdn-scan {
  position: fixed;
  inset: 0;
  background: repeating-linear-gradient(0deg, rgba(255, 255, 255, 0.025) 0 1px, transparent 1px 3px);
  mix-blend-mode: screen;
}
.pdn-dot { width: 5px; height: 5px; border-radius: 99px; background: #5eead4; box-shadow: 0 0 10px #5eead4; animation: pdnBlink 2s steps(2, end) infinite; }
.pdn-stage {
  position: relative;
  aspect-ratio: 1 / 1;
  overflow: hidden;
  border: 1px solid rgba(94, 234, 212, 0.3);
  background:
    radial-gradient(55% 55% at 50% 45%, rgba(94, 234, 212, 0.16), transparent 70%),
    radial-gradient(40% 40% at 72% 22%, rgba(244, 114, 182, 0.22), transparent 70%),
    linear-gradient(150deg, #0b0a18, #06050d 70%);
  clip-path: polygon(0 0, 100% 0, 100% calc(100% - 22px), calc(100% - 22px) 100%, 0 100%);
}
.pdn-lines {
  position: absolute;
  inset: 0;
  background-image: linear-gradient(rgba(94, 234, 212, 0.12) 1px, transparent 1px), linear-gradient(90deg, rgba(94, 234, 212, 0.12) 1px, transparent 1px);
  background-size: 34px 34px;
  mask-image: radial-gradient(circle at 50% 50%, black, transparent 82%);
}
.pdn-hud { position: absolute; font-family: ui-monospace, monospace; font-size: 8px; letter-spacing: 0.18em; color: rgba(94, 234, 212, 0.8); }
.pdn-hud-tl { left: 12px; top: 10px; }
.pdn-hud-tr { right: 12px; top: 10px; color: #f472b6; }
.pdn-hud-bl { left: 12px; bottom: 26px; }
.pdn-hud-br { right: 12px; bottom: 26px; }
.pdn-orb { position: absolute; left: 50%; top: 50%; width: 76%; aspect-ratio: 1; transform: translate(-50%, -50%); display: grid; place-items: center; }
.pdn-ring { position: absolute; border-radius: 999px; border: 1px solid rgba(94, 234, 212, 0.3); }
.pdn-ring-a { inset: 4%; animation: pdnSpin 18s linear infinite; border-style: dashed; }
.pdn-ring-b { inset: 16%; border-color: rgba(244, 114, 182, 0.35); animation: pdnSpin 12s linear infinite reverse; }
.pdn-core {
  position: relative;
  width: 46%;
  aspect-ratio: 1;
  border-radius: 999px;
  background: radial-gradient(circle at 36% 30%, #d8fff8, #5eead4 36%, #0f766e 70%, #04201d);
  box-shadow: 0 0 40px rgba(94, 234, 212, 0.45), inset -8px -10px 20px rgba(0, 0, 0, 0.45);
  animation: pdnPulse 4s ease-in-out infinite;
}
.pdn-cup { position: absolute; width: 22%; height: 30%; border-radius: 40% 40% 46% 46%; background: linear-gradient(150deg, #1f2937, #06050d); border: 1px solid rgba(94, 234, 212, 0.45); box-shadow: 0 0 22px rgba(94, 234, 212, 0.25); }
.pdn-cup-l { left: 9%; }
.pdn-cup-r { right: 9%; }
.pdn-stage-tag { position: absolute; left: 50%; bottom: 14px; transform: translateX(-50%); white-space: nowrap; font-family: ui-monospace, monospace; font-size: 7.5px; letter-spacing: 0.3em; color: rgba(232, 247, 255, 0.3); }
.pdn-glitch { position: relative; text-shadow: 2px 0 rgba(244, 114, 182, 0.55), -2px 0 rgba(94, 234, 212, 0.55); }
.pdn-chip {
  display: flex;
  gap: 8px;
  align-items: flex-start;
  padding: 10px;
  border: 1px solid rgba(94, 234, 212, 0.18);
  background: rgba(94, 234, 212, 0.035);
  transition: border-color 0.2s ease, background 0.2s ease;
}
.pdn-chip:hover { border-color: rgba(244, 114, 182, 0.6); background: rgba(244, 114, 182, 0.06); }
.pdn-cta {
  color: #06050d;
  background: linear-gradient(95deg, #5eead4, #7dd3fc 55%, #f472b6);
  clip-path: polygon(0 0, 100% 0, 100% calc(100% - 12px), calc(100% - 12px) 100%, 0 100%);
  transition: filter 0.25s ease, transform 0.15s ease;
}
.pdn-cta:hover { filter: brightness(1.15) saturate(1.15); }
.pdn-cta:active { transform: translateY(1px) scale(0.99); }
.pdn-cta-arrow { display: grid; place-items: center; width: 30px; height: 30px; border: 1px solid rgba(6, 5, 13, 0.5); }
.pdn-in { animation: pdnIn 0.7s cubic-bezier(0.16, 1, 0.3, 1) both; }
.pdn-d1 { animation-delay: 0.06s; }
.pdn-d2 { animation-delay: 0.12s; }
.pdn-d3 { animation-delay: 0.18s; }
.pdn-d4 { animation-delay: 0.24s; }
.pdn-d5 { animation-delay: 0.3s; }
@keyframes pdnIn { from { opacity: 0; transform: translateY(14px); } to { opacity: 1; transform: none; } }
@keyframes pdnBlink { 50% { opacity: 0.25; } }
@keyframes pdnSpin { to { transform: rotate(360deg); } }
@keyframes pdnPulse { 0%, 100% { transform: scale(1); } 50% { transform: scale(1.04); } }
@media (prefers-reduced-motion: reduce) {
  .pdn-in, .pdn-dot, .pdn-ring, .pdn-core { animation: none !important; }
  .pdn-cta, .pdn-chip { transition: none; }
}
</style>
