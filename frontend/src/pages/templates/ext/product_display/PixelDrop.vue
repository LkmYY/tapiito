<template>
  <div class="pdx-root min-h-screen bg-[#141026] text-[#f4f1ff]">
    <div class="mx-auto w-full max-w-md px-4 pb-14 pt-5">
      <header class="pdx-in flex items-center justify-between border-4 border-[#f4f1ff] bg-[#2a2150] px-3 py-2 font-mono text-[9px] uppercase tracking-[0.14em]">
        <span class="flex items-center gap-2"><Gamepad2 class="h-4 w-4 text-[#7cf29b]" /> LOOT_04</span>
        <span class="text-[#ffd166]">NFC:{{ nfcId || '????' }}</span>
      </header>

      <section class="pdx-in pdx-d1 pdx-stage mt-4">
        <div class="pdx-stars"></div>
        <div class="pdx-floor"></div>
        <span class="pdx-coin pdx-coin-a">+10</span>
        <span class="pdx-coin pdx-coin-b">+25</span>
        <div class="pdx-item">
          <div class="pdx-cube">
            <span class="pdx-cube-f pdx-cube-front"><i></i></span>
            <span class="pdx-cube-f pdx-cube-top"></span>
            <span class="pdx-cube-f pdx-cube-side"></span>
          </div>
          <span class="pdx-shadow"></span>
        </div>
        <span class="pdx-rarity">★ RARE DROP ★</span>
        <span class="pdx-level">LV.01</span>
      </section>

      <div class="pdx-in pdx-d2 mt-4 border-4 border-[#f4f1ff] bg-[#211a3f] p-4">
        <p class="font-mono text-[9px] uppercase tracking-[0.18em] text-[#7cf29b]">&gt; {{ p.subtitle || '—' }}</p>
        <h1 class="mt-2 text-[30px] font-black uppercase leading-[1.02] tracking-[-0.02em]">{{ p.title }}</h1>
        <div class="mt-3 flex items-end justify-between border-t-4 border-dashed border-[#f4f1ff]/25 pt-3">
          <div>
            <p class="font-mono text-[8px] uppercase tracking-[0.14em] text-[#f4f1ff]/45">BUY PRICE</p>
            <p class="pdx-price mt-1 text-[32px] font-black leading-none tabular-nums text-[#ffd166]">¥{{ p.price || '—' }}</p>
          </div>
          <div class="text-right font-mono text-[9px]">
            <p v-if="p.originalPrice" class="text-[#f4f1ff]/35 line-through">¥{{ p.originalPrice }}</p>
            <p v-if="discount" class="mt-1 inline-block bg-[#7cf29b] px-1.5 py-0.5 font-bold text-[#141026]">-{{ discount }}%</p>
            <p class="mt-1 text-[#f4f1ff]/35">NO REFUNDS</p>
          </div>
        </div>
      </div>

      <section class="pdx-in pdx-d3 mt-4 border-4 border-[#f4f1ff] bg-[#2a2150] p-3">
        <h2 class="flex items-center justify-between font-mono text-[9px] uppercase tracking-[0.18em]">
          <span class="text-[#7cf29b]">▸ ITEM PERKS</span><span class="text-[#f4f1ff]/35">03/04</span>
        </h2>
        <ul class="mt-2 grid gap-1.5">
          <li v-for="(f, i) in features" :key="i" class="flex items-center gap-2 border-2 border-[#f4f1ff]/20 bg-[#141026] px-2.5 py-2">
            <span class="text-[10px] text-[#ffd166]">◆</span>
            <span class="flex-1 text-[11px] leading-snug">{{ f }}</span>
            <span class="font-mono text-[8px] text-[#7cf29b]">+{{ (i + 1) * 5 }}</span>
          </li>
          <li v-if="!features.length" class="border-2 border-[#f4f1ff]/20 px-2.5 py-2 text-[11px] text-[#f4f1ff]/40"></li>
        </ul>
      </section>

      <section class="pdx-in pdx-d3 mt-4 border-4 border-[#f4f1ff] bg-[#211a3f] p-3">
        <h2 class="font-mono text-[9px] uppercase tracking-[0.18em] text-[#7cf29b]">▸ STAT SHEET</h2>
        <dl class="mt-2">
          <div v-for="[k, v] in specs" :key="k" class="flex items-center gap-2 border-b-2 border-dotted border-[#f4f1ff]/20 py-1.5 last:border-0">
            <dt class="w-20 flex-none font-mono text-[8.5px] uppercase tracking-[0.08em] text-[#f4f1ff]/45">{{ k }}</dt>
            <span class="flex-1 text-[10px] text-[#7cf29b]/50">··········</span>
            <dd class="font-mono text-[10px] font-bold">{{ v }}</dd>
          </div>
          <div v-if="!specs.length" class="py-1.5 text-[11px] text-[#f4f1ff]/40"></div>
        </dl>
      </section>

      <p v-if="p.description" class="pdx-in pdx-d3 mt-3 border-l-4 border-[#ffd166] pl-2.5 font-mono text-[10px] leading-[1.9] text-[#f4f1ff]/55">{{ p.description }}</p>

      <button class="pdx-in pdx-d4 pdx-buy mt-4 flex w-full items-center justify-between border-4 border-[#141026] px-4 py-3.5" @click="buy">
        <span class="text-[14px] font-black uppercase tracking-[0.04em]">{{ p.ctaText || '立即购买' }}</span>
        <span class="flex items-center gap-1 font-mono text-[10px] font-bold">[ ENTER ] <ArrowRight class="h-4 w-4" /></span>
      </button>

      <p class="mt-4 text-center font-mono text-[8px] uppercase tracking-[0.2em] text-[#f4f1ff]/30">INSERT COIN · PRESS START · 1 PLAYER</p>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { ArrowRight, Gamepad2 } from 'lucide-vue-next'

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
.pdx-root { font-family: ui-monospace, 'Courier New', monospace; }
.pdx-stage {
  position: relative;
  aspect-ratio: 1 / 0.92;
  overflow: hidden;
  border: 4px solid #f4f1ff;
  background: linear-gradient(180deg, #241b4a 0%, #171233 62%, #0e0a20 100%);
  image-rendering: pixelated;
}
.pdx-stars {
  position: absolute;
  inset: 0;
  background-image: radial-gradient(#f4f1ff 1px, transparent 1px), radial-gradient(#7cf29b 1px, transparent 1px);
  background-size: 37px 31px, 53px 47px;
  background-position: 0 0, 14px 12px;
  opacity: 0.55;
}
.pdx-floor {
  position: absolute;
  left: 0;
  right: 0;
  bottom: 0;
  height: 30%;
  background: repeating-linear-gradient(90deg, rgba(124, 242, 155, 0.18) 0 2px, transparent 2px 22px), linear-gradient(180deg, rgba(124, 242, 155, 0.08), transparent);
  transform: perspective(120px) rotateX(28deg);
  transform-origin: bottom;
}
.pdx-item { position: absolute; left: 50%; top: 47%; width: 40%; height: 46%; transform: translate(-50%, -50%); animation: pdxBob 3.2s steps(4, end) infinite; }
.pdx-cube { position: absolute; inset: 0; transform: rotate(-8deg); }
.pdx-cube-f { position: absolute; border: 3px solid #0e0a20; }
.pdx-cube-front { left: 0; bottom: 0; width: 74%; height: 64%; background: #7cf29b; box-shadow: inset -8px -8px 0 rgba(20, 16, 38, 0.22); }
.pdx-cube-front i { position: absolute; inset: 24%; border: 3px solid #141026; }
.pdx-cube-front i::after { content: ''; position: absolute; left: 50%; top: 50%; width: 26%; height: 26%; transform: translate(-50%, -50%); background: #141026; }
.pdx-cube-top { left: 18%; top: 0; width: 66%; height: 28%; background: #b5ffc5; transform: skewX(-44deg); transform-origin: bottom; }
.pdx-cube-side { right: 0; bottom: 0; width: 26%; height: 64%; background: #3fbf67; transform: skewY(-46deg); transform-origin: left; }
.pdx-shadow { position: absolute; left: 12%; right: 10%; bottom: -10%; height: 9%; background: rgba(0, 0, 0, 0.45); filter: blur(5px); }
.pdx-coin { position: absolute; font-family: ui-monospace, monospace; font-size: 11px; font-weight: 700; color: #ffd166; text-shadow: 2px 2px 0 #141026; }
.pdx-coin-a { left: 14%; top: 24%; animation: pdxCoin 3s steps(3, end) infinite; }
.pdx-coin-b { right: 12%; top: 40%; animation: pdxCoin 3.6s steps(3, end) infinite 0.6s; }
.pdx-rarity { position: absolute; left: 50%; bottom: 8%; transform: translateX(-50%); white-space: nowrap; background: #ffd166; color: #141026; padding: 2px 8px; font-size: 9px; font-weight: 700; letter-spacing: 0.1em; }
.pdx-level { position: absolute; right: 8px; top: 8px; border: 2px solid #f4f1ff; padding: 1px 5px; font-size: 9px; color: #f4f1ff; }
.pdx-price { text-shadow: 3px 3px 0 rgba(124, 242, 155, 0.4); }
.pdx-buy {
  background: #7cf29b;
  color: #141026;
  box-shadow: 6px 6px 0 #f4f1ff;
  transition: transform 0.12s steps(2, end), box-shadow 0.12s steps(2, end), background 0.12s steps(2, end);
}
.pdx-buy:hover { background: #a7ffbd; transform: translate(-2px, -2px); box-shadow: 8px 8px 0 #f4f1ff; }
.pdx-buy:active { transform: translate(4px, 4px); box-shadow: 2px 2px 0 #f4f1ff; }
.pdx-in { animation: pdxIn 0.5s steps(6, end) both; }
.pdx-d1 { animation-delay: 0.05s; }
.pdx-d2 { animation-delay: 0.1s; }
.pdx-d3 { animation-delay: 0.15s; }
.pdx-d4 { animation-delay: 0.2s; }
@keyframes pdxIn { from { opacity: 0; transform: translateY(10px); } to { opacity: 1; transform: none; } }
@keyframes pdxBob { 0%, 100% { transform: translate(-50%, -50%); } 50% { transform: translate(-50%, -56%); } }
@keyframes pdxCoin { 0%, 100% { opacity: 0; transform: translateY(0); } 40% { opacity: 1; } 70% { opacity: 0; transform: translateY(-18px); } }
@media (prefers-reduced-motion: reduce) {
  .pdx-in, .pdx-item, .pdx-coin { animation: none !important; }
  .pdx-buy { transition: none; }
}
</style>
