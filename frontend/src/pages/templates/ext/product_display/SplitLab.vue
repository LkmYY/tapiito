<template>
  <div class="pds-root min-h-screen bg-[#d7f34a] text-[#101208]">
    <div class="mx-auto w-full max-w-md">
      <header class="pds-in flex items-center justify-between border-b border-[#101208]/30 px-5 py-3 font-mono text-[9px] uppercase tracking-[0.2em]">
        <span class="flex items-center gap-2"><FlaskConical class="h-3.5 w-3.5" /> SPLIT LAB</span>
        <span>{{ nfcId || 'SAMPLE 12' }}</span>
      </header>

      <section class="pds-in pds-d1 pds-hero">
        <div class="pds-left">
          <p class="font-mono text-[9px] uppercase tracking-[0.22em] text-[#d7f34a]/70">FORMULA 12</p>
          <h1 class="pds-outline mt-3 text-[42px] font-black uppercase leading-[0.88] tracking-[-0.04em]">{{ p.title }}</h1>
          <p class="mt-4 max-w-[180px] text-[11px] leading-relaxed text-[#d7f34a]/60">{{ p.subtitle || '—' }}</p>
          <div class="pds-art">
            <div class="pds-tube">
              <span class="pds-tube-cap"></span>
              <span class="pds-tube-glass"><i></i></span>
              <span class="pds-tube-label">L12</span>
            </div>
          </div>
        </div>
        <div class="pds-right">
          <span class="pds-vertical">TESTED / 2026</span>
          <div class="pds-price">
            <span class="font-mono text-[8px] uppercase tracking-[0.16em] text-[#101208]/45">PRICE</span>
            <span class="mt-1 block text-[40px] font-black leading-none tracking-[-0.05em] tabular-nums">¥{{ p.price || '—' }}</span>
            <span v-if="p.originalPrice" class="mt-1 block font-mono text-[10px] text-[#101208]/40 line-through">¥{{ p.originalPrice }}</span>
            <span v-if="discount" class="mt-2 inline-block bg-[#101208] px-2 py-0.5 font-mono text-[8px] uppercase tracking-[0.1em] text-[#d7f34a]">−{{ discount }}%</span>
          </div>
          <div class="pds-spec-mini">
            <div v-for="[k, v] in specs.slice(0, 3)" :key="k" class="border-t border-[#101208]/25 py-1.5">
              <p class="font-mono text-[7px] uppercase text-[#101208]/45">{{ k }}</p>
              <p class="text-[10px] font-semibold">{{ v }}</p>
            </div>
            <div v-if="!specs.length" class="border-t border-[#101208]/25 py-1.5 text-[10px] text-[#101208]/40"></div>
          </div>
        </div>
      </section>

      <div class="pds-in pds-d2 border-t border-[#101208]/30 px-5 py-4">
        <div class="flex items-center justify-between">
          <h2 class="font-mono text-[9px] uppercase tracking-[0.24em]">01 / ACTIVE INGREDIENTS</h2>
          <span class="font-mono text-[9px] text-[#101208]/40">{{ String(features.length).padStart(2, '0') }}</span>
        </div>
        <ul class="mt-2">
          <li v-for="(f, i) in features" :key="i" class="flex items-center gap-3 border-b border-[#101208]/15 py-2.5">
            <span class="grid h-5 w-5 flex-none place-items-center rounded-full border border-[#101208] font-mono text-[8px]">{{ i + 1 }}</span>
            <span class="flex-1 text-[12px] leading-snug">{{ f }}</span>
            <span class="font-mono text-[8px] text-[#101208]/35">A{{ String(i + 1).padStart(2, '0') }}</span>
          </li>
          <li v-if="!features.length" class="border-b border-[#101208]/15 py-2.5 text-[12px] text-[#101208]/40"></li>
        </ul>
      </div>

      <div class="pds-in pds-d2 flex border-t border-[#101208]/30">
        <div class="flex-1 px-5 py-4">
          <h2 class="font-mono text-[9px] uppercase tracking-[0.2em]">02 / METHOD</h2>
          <p class="mt-2 text-[11px] leading-[1.8] text-[#101208]/65">{{ p.description }}</p>
        </div>
        <div class="w-[38%] flex-none border-l border-[#101208]/30 bg-[#101208] px-4 py-4 text-[#d7f34a]">
          <h2 class="font-mono text-[9px] uppercase tracking-[0.16em]">FULL SPEC</h2>
          <dl class="mt-2">
            <div v-for="[k, v] in specs" :key="k" class="flex items-baseline justify-between gap-2 border-b border-[#d7f34a]/20 py-1.5">
              <dt class="font-mono text-[7.5px] uppercase text-[#d7f34a]/50">{{ k }}</dt>
              <dd class="text-[9.5px]">{{ v }}</dd>
            </div>
            <div v-if="!specs.length" class="py-1.5 text-[10px] text-[#d7f34a]/50"></div>
          </dl>
        </div>
      </div>

      <div class="pds-in pds-d3 px-5 pb-10 pt-4">
        <button class="pds-buy flex w-full items-center justify-between border border-[#101208] px-4 py-3.5" @click="buy">
          <span class="text-[13px] font-black uppercase tracking-[0.06em]">{{ p.ctaText || '立即购买' }}</span>
          <span class="flex items-center gap-2 font-mono text-[9px] uppercase tracking-[0.16em]">RUN ORDER <ArrowUpRight class="h-4 w-4" /></span>
        </button>
        <p class="mt-4 text-center font-mono text-[8px] uppercase tracking-[0.24em] text-[#101208]/40">Split Lab · practical objects division</p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { ArrowUpRight, FlaskConical } from 'lucide-vue-next'

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
.pds-hero { display: flex; min-height: 430px; }
.pds-left { flex: 1; position: relative; overflow: hidden; padding: 18px 0 16px 20px; background: #101208; color: #d7f34a; }
.pds-outline { color: transparent; -webkit-text-stroke: 1.5px #d7f34a; }
.pds-art { position: absolute; left: 50%; bottom: -18px; width: 90%; height: 200px; transform: translateX(-50%); }
.pds-tube { position: absolute; left: 50%; bottom: 0; width: 34%; height: 82%; transform: translateX(-50%) rotate(-5deg); animation: pdsFloat 7s ease-in-out infinite; }
.pds-tube-cap { position: absolute; left: 24%; top: 0; width: 52%; height: 12%; border-radius: 4px 4px 0 0; background: #d7f34a; }
.pds-tube-glass {
  position: absolute;
  inset: 10% 0 0;
  border-radius: 10px 10px 22px 22px;
  background: linear-gradient(100deg, rgba(215, 243, 74, 0.22), rgba(215, 243, 74, 0.05) 45%, rgba(215, 243, 74, 0.3));
  border: 1px solid rgba(215, 243, 74, 0.65);
  box-shadow: inset -8px 0 12px rgba(215, 243, 74, 0.16);
}
.pds-tube-glass i { position: absolute; left: 12%; right: 12%; bottom: 8%; height: 54%; border-radius: 4px 4px 16px 16px; background: rgba(215, 243, 74, 0.55); }
.pds-tube-label { position: absolute; left: 50%; top: 46%; transform: translate(-50%, -50%) rotate(90deg); font-family: ui-monospace, monospace; font-size: 9px; letter-spacing: 0.2em; color: #101208; }
.pds-right { width: 38%; flex: none; display: flex; flex-direction: column; justify-content: space-between; padding: 18px 16px 16px; background: #d7f34a; position: relative; }
.pds-vertical { position: absolute; right: 6px; top: 50%; transform: translateY(-50%) rotate(90deg); transform-origin: center; font-family: ui-monospace, monospace; font-size: 7px; letter-spacing: 0.24em; color: rgba(16, 18, 8, 0.4); white-space: nowrap; }
.pds-price { padding-top: 56px; }
.pds-spec-mini { margin-top: 20px; }
.pds-buy { background: #101208; color: #d7f34a; transition: background 0.25s ease, padding 0.2s ease; }
.pds-buy:hover { background: #2c301a; padding-inline: 20px; }
.pds-buy:active { transform: scale(0.99); }
.pds-in { animation: pdsIn 0.7s cubic-bezier(0.22, 1, 0.36, 1) both; }
.pds-d1 { animation-delay: 0.06s; }
.pds-d2 { animation-delay: 0.14s; }
.pds-d3 { animation-delay: 0.22s; }
@keyframes pdsIn { from { opacity: 0; transform: translateY(14px); } to { opacity: 1; transform: none; } }
@keyframes pdsFloat { 0%, 100% { transform: translateX(-50%) rotate(-5deg); } 50% { transform: translateX(-50%) translateY(-8px) rotate(-2deg); } }
@media (prefers-reduced-motion: reduce) {
  .pds-in, .pds-tube { animation: none !important; }
  .pds-buy { transition: none; }
}
</style>
