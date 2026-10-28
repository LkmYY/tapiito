<template>
  <div class="pdpo-root min-h-screen bg-[#f4e7d3] text-[#1b1a17]">
    <div class="mx-auto w-full max-w-md pb-10">
      <header class="pdpo-in flex items-center justify-between px-5 pt-5 font-mono text-[9px] uppercase tracking-[0.2em]">
        <span>POSTER / STORE</span>
        <span>{{ nfcId || 'NO. 07' }}</span>
      </header>

      <section class="pdpo-in pdpo-d1 pdpo-poster mt-4">
        <div class="pdpo-orange">
          <span class="pdpo-outline">BUY</span>
          <span class="pdpo-outline pdpo-outline-2">THIS</span>
          <span class="pdpo-year">2026</span>
          <div class="pdpo-flower"><i v-for="n in 8" :key="n" :style="{ transform: 'rotate(' + n * 45 + 'deg)' }"></i><b></b></div>
          <span class="pdpo-side">LIMITED OBJECT · ONE COLOR · NO LOGO</span>
        </div>
        <div class="pdpo-cream">
          <h1 class="text-[34px] font-black uppercase leading-[0.86] tracking-[-0.045em]">{{ p.title }}</h1>
          <p class="mt-2 max-w-[85%] font-mono text-[8.5px] uppercase leading-relaxed tracking-[0.12em]">{{ p.subtitle || '—' }}</p>
          <div class="pdpo-thumb">
            <span class="pdpo-thumb-obj"></span>
            <span class="pdpo-thumb-shadow"></span>
          </div>
          <p class="mt-auto font-mono text-[8px] uppercase tracking-[0.18em] text-[#1b1a17]/50">TAP TO OWN →</p>
        </div>
      </section>

      <div class="pdpo-in pdpo-d2 flex items-end justify-between gap-4 px-5 pt-5">
        <div>
          <p class="font-mono text-[8px] uppercase tracking-[0.2em] text-[#1b1a17]/45">PRICE / 含税</p>
          <p class="pdpo-price mt-1 text-[52px] font-black leading-[0.85] tracking-[-0.06em] tabular-nums">¥{{ p.price || '—' }}</p>
          <p v-if="p.originalPrice" class="mt-2 font-mono text-[11px] text-[#1b1a17]/40">原价 <span class="line-through">¥{{ p.originalPrice }}</span></p>
        </div>
        <div v-if="discount" class="pdpo-badge">
          <span class="text-[20px] font-black leading-none">−{{ discount }}%</span>
          <span class="mt-1 font-mono text-[7px] uppercase tracking-[0.14em]">DIRECT</span>
        </div>
      </div>

      <section class="pdpo-in pdpo-d3 mt-5 px-5">
        <div class="border-t-2 border-[#1b1a17] pt-2">
          <h2 class="flex items-center justify-between font-mono text-[9px] uppercase tracking-[0.2em]">
            <span>WHY / 为什么</span><span class="font-sans text-[12px] font-black">01</span>
          </h2>
        </div>
        <ul class="mt-1 grid grid-cols-2 gap-x-5">
          <li v-for="(f, i) in features" :key="i" class="border-b border-[#1b1a17]/20 py-2.5">
            <span class="font-mono text-[8px] text-[#d84f1e]">{{ String(i + 1).padStart(2, '0') }}</span>
            <p class="mt-0.5 text-[11.5px] font-semibold leading-snug">{{ f }}</p>
          </li>
          <li v-if="!features.length" class="col-span-2 border-b border-[#1b1a17]/20 py-2.5 text-[11.5px] text-[#1b1a17]/40"></li>
        </ul>
      </section>

      <section class="pdpo-in pdpo-d3 mt-4 px-5">
        <div class="border-t-2 border-[#1b1a17] pt-2">
          <h2 class="flex items-center justify-between font-mono text-[9px] uppercase tracking-[0.2em]">
            <span>SPEC / 尺寸</span><span class="font-sans text-[12px] font-black">02</span>
          </h2>
        </div>
        <dl class="mt-1">
          <div v-for="[k, v] in specs" :key="k" class="flex items-baseline justify-between border-b border-[#1b1a17]/20 py-2">
            <dt class="font-mono text-[9px] uppercase tracking-[0.1em] text-[#1b1a17]/50">{{ k }}</dt>
            <dd class="text-[12px] font-bold">{{ v }}</dd>
          </div>
          <div v-if="!specs.length" class="border-b border-[#1b1a17]/20 py-2 text-[11.5px] text-[#1b1a17]/40"></div>
        </dl>
      </section>

      <p v-if="p.description" class="pdpo-in pdpo-d3 mx-5 mt-4 border-l-4 border-[#d84f1e] pl-3 text-[11.5px] leading-[1.85] text-[#1b1a17]/65">{{ p.description }}</p>

      <div class="pdpo-in pdpo-d4 px-5 pt-5">
        <button class="pdpo-buy flex w-full items-center justify-between bg-[#1b1a17] px-5 py-4 text-[#f4e7d3]" @click="buy">
          <span class="text-[15px] font-black uppercase tracking-[0.02em]">{{ p.ctaText || '立即购买' }}</span>
          <ArrowUpRight class="h-5 w-5" />
        </button>
        <p class="mt-4 text-center font-mono text-[8px] uppercase tracking-[0.26em] text-[#1b1a17]/40">Poster Store · editions of 100 · made loudly</p>
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
.pdpo-poster { display: grid; grid-template-columns: 1.05fr 0.95fr; margin-inline: 20px; border: 2px solid #1b1a17; background: #f4e7d3; box-shadow: 8px 8px 0 #1b1a17; overflow: hidden; }
.pdpo-orange { position: relative; min-height: 330px; background: #d84f1e; color: #f4e7d3; overflow: hidden; }
.pdpo-outline { position: absolute; left: -4px; top: 12px; font-size: 66px; font-weight: 900; line-height: 0.82; letter-spacing: -0.06em; color: transparent; -webkit-text-stroke: 1.5px #1b1a17; }
.pdpo-outline-2 { top: 66px; left: 12px; -webkit-text-stroke-color: #f4e7d3; }
.pdpo-year { position: absolute; left: 10px; bottom: 46px; font-family: ui-monospace, monospace; font-size: 10px; letter-spacing: 0.2em; }
.pdpo-flower { position: absolute; left: 52%; top: 50%; width: 74px; height: 74px; transform: translate(-50%, -50%); }
.pdpo-flower i { position: absolute; left: 50%; top: 50%; width: 12px; height: 34px; margin-left: -6px; border-radius: 999px; background: #1b1a17; transform-origin: 50% 100%; translate: 0 -100%; }
.pdpo-flower b { position: absolute; left: 50%; top: 50%; width: 23px; height: 23px; transform: translate(-50%, -50%); border-radius: 999px; background: #f4e7d3; }
.pdpo-side { position: absolute; left: 8px; bottom: 8px; font-family: ui-monospace, monospace; font-size: 7px; line-height: 1.5; letter-spacing: 0.14em; color: rgba(27, 26, 23, 0.7); }
.pdpo-cream { display: flex; flex-direction: column; padding: 14px 12px 12px; }
.pdpo-thumb { position: relative; margin-top: 12px; flex: 1; min-height: 120px; display: grid; place-items: center; border: 1px dashed rgba(27, 26, 23, 0.35); background: rgba(27, 26, 23, 0.04); }
.pdpo-thumb-obj { width: 46%; height: 56%; border-radius: 45% 45% 38% 38% / 24% 24% 20% 20%; background: linear-gradient(120deg, #3a3833, #1b1a17 65%); box-shadow: inset -8px -4px 12px rgba(0, 0, 0, 0.4), inset 6px 4px 8px rgba(255, 255, 255, 0.14); animation: pdpoSway 6s ease-in-out infinite; }
.pdpo-thumb-shadow { position: absolute; bottom: 12%; width: 40%; height: 5%; border-radius: 50%; background: rgba(27, 26, 23, 0.22); filter: blur(3px); }
.pdpo-price { font-family: 'Arial Black', 'Helvetica Neue', sans-serif; }
.pdpo-badge { display: flex; flex-direction: column; align-items: center; justify-content: center; width: 76px; height: 76px; border-radius: 999px; background: #d84f1e; color: #f4e7d3; transform: rotate(7deg); flex: none; }
.pdpo-buy { transition: background 0.25s ease, transform 0.16s ease; }
.pdpo-buy:hover { background: #d84f1e; }
.pdpo-buy:active { transform: translateY(2px); }
.pdpo-in { animation: pdpoIn 0.8s cubic-bezier(0.22, 1, 0.36, 1) both; }
.pdpo-d1 { animation-delay: 0.07s; }
.pdpo-d2 { animation-delay: 0.14s; }
.pdpo-d3 { animation-delay: 0.2s; }
.pdpo-d4 { animation-delay: 0.26s; }
@keyframes pdpoIn { from { opacity: 0; transform: translateY(18px); } to { opacity: 1; transform: none; } }
@keyframes pdpoSway { 0%, 100% { transform: rotate(-3deg); } 50% { transform: rotate(2deg); } }
@media (prefers-reduced-motion: reduce) {
  .pdpo-in, .pdpo-thumb-obj { animation: none !important; }
  .pdpo-buy { transition: none; }
}
</style>
