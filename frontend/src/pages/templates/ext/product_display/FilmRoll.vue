<template>
  <div class="pdfr-root min-h-screen bg-[#17130f] text-[#efe6d6]">
    <div class="pdfr-grain pointer-events-none"></div>
    <div class="mx-auto w-full max-w-md px-5 pb-14 pt-6">
      <header class="pdfr-in flex items-center justify-between font-mono text-[9px] uppercase tracking-[0.24em] text-[#d9b27c]">
        <span class="flex items-center gap-2"><Film class="h-4 w-4" /> KODA / 400</span>
        <span class="text-[#efe6d6]/35">ROLL {{ nfcId || '04A' }}</span>
      </header>

      <section class="pdfr-in pdfr-d1 pdfr-frame mt-4">
        <div class="pdfr-holes pdfr-holes-t"><i v-for="n in 8" :key="n"></i></div>
        <div class="pdfr-holes pdfr-holes-b"><i v-for="n in 8" :key="n"></i></div>
        <div class="pdfr-shot">
          <div class="pdfr-sun"></div>
          <div class="pdfr-horizon"></div>
          <div class="pdfr-product">
            <span class="pdfr-prod-body"></span>
            <span class="pdfr-prod-band"></span>
            <span class="pdfr-prod-cap"></span>
          </div>
          <span class="pdfr-shot-no">12A</span>
          <span class="pdfr-shot-meta">1/250 · f/2.8 · ISO 400</span>
          <span class="pdfr-date">'26 07 18</span>
        </div>
        <div class="pdfr-contact">
          <span>35mm</span><span>COLOR NEGATIVE</span><span>12 EXP</span>
        </div>
      </section>

      <div class="pdfr-in pdfr-d2 mt-5 flex items-start justify-between gap-4">
        <div>
          <p class="font-mono text-[9px] uppercase tracking-[0.22em] text-[#d9b27c]">{{ p.subtitle || '—' }}</p>
          <h1 class="mt-1.5 font-serif text-[32px] leading-[1.08]">{{ p.title }}</h1>
        </div>
        <span class="pdfr-iso">400</span>
      </div>

      <div class="pdfr-in pdfr-d2 mt-4 flex items-end justify-between border-y border-[#d9b27c]/30 py-3.5">
        <div class="flex items-baseline gap-2">
          <span class="font-mono text-[9px] text-[#d9b27c]/80">JPY</span>
          <span class="font-serif text-[36px] leading-none tabular-nums">¥{{ p.price || '—' }}</span>
        </div>
        <div class="text-right font-mono text-[9px] leading-relaxed">
          <p v-if="p.originalPrice" class="text-[#efe6d6]/30 line-through">¥{{ p.originalPrice }}</p>
          <p v-if="discount" class="mt-1 text-[#d9b27c]">SAVE {{ discount }}%</p>
          <p class="text-[#efe6d6]/30">SHIPPING INCL.</p>
        </div>
      </div>

      <section class="pdfr-in pdfr-d3 mt-5">
        <div class="flex items-center gap-3">
          <h2 class="font-mono text-[9px] uppercase tracking-[0.22em] text-[#d9b27c]">FRAME NOTES</h2>
          <span class="h-px flex-1 bg-[#d9b27c]/25"></span>
          <span class="font-mono text-[8px] text-[#efe6d6]/25">04</span>
        </div>
        <ul class="mt-2">
          <li v-for="(f, i) in features" :key="i" class="flex items-baseline gap-3 border-b border-[#efe6d6]/8 py-2.5">
            <span class="font-mono text-[8px] text-[#d9b27c]/70">{{ String(i + 1).padStart(2, '0') }}A</span>
            <span class="flex-1 text-[12px] leading-relaxed text-[#efe6d6]/75">{{ f }}</span>
          </li>
          <li v-if="!features.length" class="border-b border-[#efe6d6]/8 py-2.5 text-[12px] text-[#efe6d6]/35"></li>
        </ul>
      </section>

      <section class="pdfr-in pdfr-d3 mt-4">
        <h2 class="font-mono text-[9px] uppercase tracking-[0.22em] text-[#d9b27c]">EXPOSURE DATA</h2>
        <dl class="mt-2 grid grid-cols-4 border border-[#d9b27c]/25">
          <div v-for="[k, v] in specs" :key="k" class="border-r border-[#d9b27c]/15 px-2 py-3 last:border-r-0">
            <dt class="font-mono text-[7px] uppercase tracking-[0.1em] text-[#efe6d6]/35">{{ k }}</dt>
            <dd class="mt-1 text-[11px] text-[#efe6d6]/85">{{ v }}</dd>
          </div>
          <div v-if="!specs.length" class="col-span-4 px-2 py-3 text-[11px] text-[#efe6d6]/35"></div>
        </dl>
      </section>

      <p v-if="p.description" class="pdfr-in pdfr-d3 mt-4 font-serif text-[12px] italic leading-[1.9] text-[#efe6d6]/50">{{ p.description }}</p>

      <button class="pdfr-in pdfr-d4 pdfr-buy mt-5 flex w-full items-center justify-between border border-[#d9b27c] px-4 py-3.5" @click="buy">
        <span class="font-mono text-[10px] uppercase tracking-[0.2em] text-[#d9b27c]">{{ p.ctaText || '立即购买' }}</span>
        <span class="flex items-center gap-2 font-mono text-[9px] uppercase tracking-[0.14em] text-[#efe6d6]/70">DEVELOP ORDER <ArrowUpRight class="h-3.5 w-3.5" /></span>
      </button>

      <p class="mt-5 text-center font-mono text-[8px] uppercase tracking-[0.24em] text-[#efe6d6]/25">Shot on 35mm · hand developed · Koda Studio</p>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { ArrowUpRight, Film } from 'lucide-vue-next'

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
.pdfr-root { position: relative; overflow: hidden; }
.pdfr-grain {
  position: absolute;
  inset: 0;
  opacity: 0.4;
  background-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' width='140' height='140'%3E%3Cfilter id='n'%3E%3CfeTurbulence type='fractalNoise' baseFrequency='0.8' numOctaves='2'/%3E%3C/filter%3E%3Crect width='100%25' height='100%25' filter='url(%23n)' opacity='0.3'/%3E%3C/svg%3E");
  pointer-events: none;
}
.pdfr-frame { position: relative; padding: 22px 16px; background: #0f0c09; border: 1px solid rgba(217, 178, 124, 0.25); }
.pdfr-holes { position: absolute; left: 6px; right: 6px; display: flex; justify-content: space-between; }
.pdfr-holes-t { top: 5px; }
.pdfr-holes-b { bottom: 5px; }
.pdfr-holes i { width: 11px; height: 7px; border-radius: 2px; background: rgba(239, 230, 214, 0.22); }
.pdfr-shot { position: relative; aspect-ratio: 4 / 4.4; overflow: hidden; background: linear-gradient(180deg, #9a6a42 0%, #c99a63 42%, #6b4a30 100%); }
.pdfr-sun { position: absolute; left: 62%; top: 16%; width: 26%; aspect-ratio: 1; border-radius: 999px; background: #f4d9a8; filter: blur(1px); opacity: 0.9; }
.pdfr-horizon { position: absolute; left: 0; right: 0; bottom: 0; height: 42%; background: linear-gradient(180deg, rgba(60, 38, 24, 0.15), rgba(36, 22, 14, 0.75)); }
.pdfr-product { position: absolute; left: 50%; top: 54%; width: 34%; height: 50%; transform: translate(-50%, -50%); animation: pdfrSettle 1s cubic-bezier(0.22, 1, 0.36, 1) both; }
.pdfr-prod-body {
  position: absolute;
  inset: 16% 0 0;
  border-radius: 14% 14% 20% 20% / 8% 8% 12% 12%;
  background: linear-gradient(100deg, #3a2a1d 0%, #1f150e 45%, #0e0906 100%);
  box-shadow: inset -10px 0 18px rgba(0, 0, 0, 0.6), inset 8px 0 14px rgba(244, 217, 168, 0.16), 16px 20px 26px -16px rgba(0, 0, 0, 0.6);
}
.pdfr-prod-band { position: absolute; left: 34%; top: 0; width: 32%; height: 24%; border-radius: 4px 4px 0 0; background: #c9a15f; }
.pdfr-prod-cap { position: absolute; left: 28%; top: -5%; width: 44%; height: 7%; border-radius: 999px; background: #e6cfa2; }
.pdfr-shot-no { position: absolute; left: 10px; top: 8px; font-family: ui-monospace, monospace; font-size: 9px; color: rgba(244, 217, 168, 0.8); }
.pdfr-shot-meta { position: absolute; right: 10px; top: 8px; font-family: ui-monospace, monospace; font-size: 7px; letter-spacing: 0.12em; color: rgba(244, 217, 168, 0.7); }
.pdfr-date { position: absolute; right: 10px; bottom: 8px; font-family: ui-monospace, monospace; font-size: 8px; color: rgba(244, 217, 168, 0.8); }
.pdfr-contact { display: flex; justify-content: space-between; margin-top: 8px; font-family: ui-monospace, monospace; font-size: 7px; letter-spacing: 0.16em; color: rgba(217, 178, 124, 0.6); }
.pdfr-iso { display: grid; place-items: center; width: 46px; height: 46px; border: 1px solid rgba(217, 178, 124, 0.4); font-family: ui-monospace, monospace; font-size: 11px; color: #d9b27c; }
.pdfr-buy { background: transparent; transition: background 0.3s ease, color 0.3s ease; }
.pdfr-buy:hover { background: #d9b27c; color: #17130f; }
.pdfr-buy:hover span { color: #17130f; }
.pdfr-buy:active { transform: scale(0.99); }
.pdfr-in { animation: pdfrIn 0.8s cubic-bezier(0.22, 1, 0.36, 1) both; }
.pdfr-d1 { animation-delay: 0.07s; }
.pdfr-d2 { animation-delay: 0.14s; }
.pdfr-d3 { animation-delay: 0.21s; }
.pdfr-d4 { animation-delay: 0.28s; }
@keyframes pdfrIn { from { opacity: 0; transform: translateY(16px); } to { opacity: 1; transform: none; } }
@keyframes pdfrSettle { from { opacity: 0; transform: translate(-50%, -46%) scale(0.97); } to { opacity: 1; transform: translate(-50%, -50%) scale(1); } }
@media (prefers-reduced-motion: reduce) {
  .pdfr-in, .pdfr-product { animation: none !important; }
  .pdfr-buy { transition: none; }
}
</style>
