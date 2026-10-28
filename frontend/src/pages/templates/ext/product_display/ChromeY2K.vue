<template>
  <div class="pdcy-root min-h-screen bg-[#dfe3ef] text-[#26263a]">
    <div class="pdcy-glow pointer-events-none"></div>
    <div class="relative mx-auto w-full max-w-md px-5 pb-14 pt-6">
      <header class="pdcy-in flex items-center justify-between rounded-full border border-white/70 bg-white/40 px-4 py-2 font-mono text-[9px] uppercase tracking-[0.2em] backdrop-blur-md">
        <span class="flex items-center gap-1.5"><Sparkles class="h-3.5 w-3.5 text-[#6d5bd0]" /> CHROME 2000</span>
        <span class="text-[#26263a]/50">{{ nfcId || 'UNIT 01' }}</span>
      </header>

      <section class="pdcy-in pdcy-d1 pdcy-stage mt-4">
        <span class="pdcy-star pdcy-star-a">✦</span>
        <span class="pdcy-star pdcy-star-b">✦</span>
        <span class="pdcy-star pdcy-star-c">✧</span>
        <div class="pdcy-orb">
          <div class="pdcy-orb-core"></div>
          <div class="pdcy-ring"></div>
          <div class="pdcy-phone">
            <span class="pdcy-phone-screen"></span>
            <span class="pdcy-phone-cam"></span>
          </div>
        </div>
        <span class="pdcy-caption">DIGITAL / PHYSICAL / ETERNAL</span>
        <span class="pdcy-badge">Y2K<br /><b>26</b></span>
      </section>

      <div class="pdcy-in pdcy-d2 mt-5 text-center">
        <p class="font-mono text-[8.5px] uppercase tracking-[0.3em] text-[#6d5bd0]">{{ p.subtitle || '—' }}</p>
        <h1 class="pdcy-title mt-2 text-[36px] font-extrabold leading-[1.02] tracking-[-0.03em]">{{ p.title }}</h1>
        <div class="mt-4 inline-flex items-baseline gap-3 rounded-full border border-white/80 bg-white/45 px-5 py-2.5 shadow-[0_10px_24px_-18px_rgba(38,38,58,0.7)] backdrop-blur-sm">
          <span class="text-[30px] font-bold leading-none tabular-nums">¥{{ p.price || '—' }}</span>
          <span v-if="p.originalPrice" class="font-mono text-[10px] text-[#26263a]/40 line-through">¥{{ p.originalPrice }}</span>
          <span v-if="discount" class="rounded-full bg-[#6d5bd0] px-2 py-0.5 font-mono text-[8px] uppercase text-white">−{{ discount }}%</span>
        </div>
      </div>

      <p v-if="p.description" class="pdcy-in pdcy-d2 mx-auto mt-4 max-w-[320px] text-center text-[11.5px] leading-[1.85] text-[#26263a]/55">{{ p.description }}</p>

      <section class="pdcy-in pdcy-d3 mt-5 rounded-[24px] border border-white/80 bg-white/45 p-4 shadow-[0_16px_34px_-28px_rgba(38,38,58,0.8)] backdrop-blur-md">
        <div class="flex items-center justify-between">
          <h2 class="font-mono text-[9px] uppercase tracking-[0.22em] text-[#26263a]/55">✦ FEATURES</h2>
          <span class="font-mono text-[9px] text-[#26263a]/35">0{{ features.length }}</span>
        </div>
        <ul class="mt-2 grid gap-1.5">
          <li v-for="(f, i) in features" :key="i" class="flex items-center gap-2.5 rounded-full border border-white/80 bg-white/55 px-3 py-2">
            <span class="grid h-4 w-4 flex-none place-items-center rounded-full bg-[#c7bdf5] font-mono text-[7px] text-[#3c2f7a]">{{ i + 1 }}</span>
            <span class="flex-1 text-[11px] font-medium">{{ f }}</span>
          </li>
          <li v-if="!features.length" class="rounded-full border border-white/80 bg-white/55 px-3 py-2 text-[11px] text-[#26263a]/40"></li>
        </ul>
      </section>

      <section class="pdcy-in pdcy-d3 mt-3 rounded-[24px] border border-white/80 bg-white/40 p-4 backdrop-blur-md">
        <h2 class="font-mono text-[9px] uppercase tracking-[0.22em] text-[#26263a]/55">✦ SPECS</h2>
        <dl class="mt-2 grid grid-cols-2 gap-2">
          <div v-for="[k, v] in specs" :key="k" class="rounded-2xl border border-white/80 bg-white/45 px-3 py-2.5">
            <dt class="font-mono text-[7.5px] uppercase tracking-[0.1em] text-[#26263a]/40">{{ k }}</dt>
            <dd class="mt-1 text-[11px] font-semibold">{{ v }}</dd>
          </div>
          <div v-if="!specs.length" class="col-span-2 rounded-2xl border border-white/80 bg-white/45 px-3 py-2.5 text-[11px] text-[#26263a]/40"></div>
        </dl>
      </section>

      <button class="pdcy-in pdcy-d4 pdcy-buy mt-4 flex w-full items-center justify-between px-5 py-4" @click="buy">
        <span class="text-[12px] font-bold tracking-[0.06em]">{{ p.ctaText || '立即购买' }}</span>
        <span class="flex h-7 w-7 items-center justify-center rounded-full bg-[#26263a] text-white"><ArrowUpRight class="h-4 w-4" /></span>
      </button>

      <p class="mt-5 text-center font-mono text-[8px] uppercase tracking-[0.28em] text-[#26263a]/35">Future nostalgia · chrome edition</p>
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
.pdcy-root { position: relative; overflow: hidden; }
.pdcy-glow {
  position: absolute;
  top: -110px;
  left: 50%;
  width: 460px;
  height: 460px;
  transform: translateX(-50%);
  border-radius: 999px;
  background: conic-gradient(from 210deg, #b9c7ff, #f0c4f4, #c9f3ef, #d8c7ff, #b9c7ff);
  filter: blur(46px);
  opacity: 0.7;
}
.pdcy-stage { position: relative; aspect-ratio: 1 / 1.04; overflow: hidden; border-radius: 32px; border: 1px solid rgba(255, 255, 255, 0.85); background: linear-gradient(150deg, #eef1fb, #cfd6ea 55%, #b9c3dc); box-shadow: inset 0 0 50px rgba(255, 255, 255, 0.75), 0 24px 44px -30px rgba(38, 38, 58, 0.55); }
.pdcy-stage::before { content: ''; position: absolute; inset: 0; background: linear-gradient(120deg, transparent 30%, rgba(255, 255, 255, 0.55) 44%, transparent 58%); animation: pdcySweep 7s ease-in-out infinite; }
.pdcy-orb { position: absolute; left: 50%; top: 50%; width: 74%; aspect-ratio: 1; transform: translate(-50%, -50%); display: grid; place-items: center; }
.pdcy-orb-core { position: absolute; inset: 8%; border-radius: 999px; background: conic-gradient(from 40deg, #e7e9f3, #aeb8d6, #f2d9f7, #9fb0cf, #e7e9f3); box-shadow: inset 0 0 40px rgba(255, 255, 255, 0.9), 0 20px 40px -24px rgba(38, 38, 58, 0.65); }
.pdcy-ring { position: absolute; inset: 1%; border-radius: 999px; border: 1px solid rgba(255, 255, 255, 0.9); transform: rotateX(66deg); }
.pdcy-phone { position: absolute; width: 30%; height: 52%; transform: rotate(-12deg); border-radius: 10px; background: linear-gradient(150deg, #f9fafc, #9aa6c4 48%, #454d66 100%); box-shadow: inset 0 0 0 1px rgba(255, 255, 255, 0.55), 8px 14px 20px -12px rgba(38, 38, 58, 0.6); animation: pdcyTilt 7s ease-in-out infinite; }
.pdcy-phone-screen { position: absolute; inset: 8% 9% 14%; border-radius: 5px; background: linear-gradient(180deg, #d9d2ff, #9fe6ea 60%, #f8d6ef); box-shadow: inset 0 0 12px rgba(255, 255, 255, 0.9); }
.pdcy-phone-cam { position: absolute; left: 50%; top: 4%; width: 24%; height: 4%; transform: translateX(-50%); border-radius: 999px; background: #474e66; }
.pdcy-star { position: absolute; color: rgba(109, 91, 208, 0.6); animation: pdcyTwinkle 4s ease-in-out infinite; }
.pdcy-star-a { left: 12%; top: 16%; font-size: 18px; }
.pdcy-star-b { right: 14%; top: 26%; font-size: 12px; animation-delay: 1.2s; }
.pdcy-star-c { right: 22%; bottom: 16%; font-size: 22px; animation-delay: 2s; }
.pdcy-caption { position: absolute; left: 16px; bottom: 12px; font-family: ui-monospace, monospace; font-size: 7px; letter-spacing: 0.22em; color: rgba(38, 38, 58, 0.45); }
.pdcy-badge { position: absolute; right: 14px; top: 14px; display: grid; place-items: center; width: 48px; height: 48px; border-radius: 999px; border: 1px solid rgba(255, 255, 255, 0.9); background: rgba(255, 255, 255, 0.4); font-family: ui-monospace, monospace; font-size: 7px; line-height: 1.2; text-align: center; color: #6d5bd0; backdrop-filter: blur(4px); }
.pdcy-badge b { font-size: 11px; }
.pdcy-title { background: linear-gradient(100deg, #26263a 10%, #6d5bd0 50%, #26263a 90%); -webkit-background-clip: text; background-clip: text; color: transparent; }
.pdcy-buy {
  border-radius: 999px;
  background: linear-gradient(100deg, #c9c2f7, #e7c8f2 40%, #b9e7ea 70%, #c9c2f7);
  color: #26263a;
  box-shadow: 0 14px 28px -18px rgba(38, 38, 58, 0.65), inset 0 1px 0 rgba(255, 255, 255, 0.9);
  transition: filter 0.25s ease, transform 0.18s ease;
}
.pdcy-buy:hover { filter: brightness(1.06) saturate(1.15); transform: translateY(-1px); }
.pdcy-buy:active { transform: scale(0.985); }
.pdcy-in { animation: pdcyIn 0.8s cubic-bezier(0.22, 1, 0.36, 1) both; }
.pdcy-d1 { animation-delay: 0.06s; }
.pdcy-d2 { animation-delay: 0.13s; }
.pdcy-d3 { animation-delay: 0.2s; }
.pdcy-d4 { animation-delay: 0.27s; }
@keyframes pdcyIn { from { opacity: 0; transform: translateY(18px); } to { opacity: 1; transform: none; } }
@keyframes pdcySweep { 0%, 100% { transform: translateX(-30%); } 50% { transform: translateX(30%); } }
@keyframes pdcyTilt { 0%, 100% { transform: rotate(-12deg) translateY(0); } 50% { transform: rotate(-6deg) translateY(-5px); } }
@keyframes pdcyTwinkle { 0%, 100% { opacity: 0.3; transform: scale(0.85); } 50% { opacity: 0.9; transform: scale(1.15); } }
@media (prefers-reduced-motion: reduce) {
  .pdcy-in, .pdcy-stage::before, .pdcy-phone, .pdcy-star { animation: none !important; }
  .pdcy-buy { transition: none; }
}
</style>
