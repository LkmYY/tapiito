<template>
  <div class="pdi-root min-h-screen bg-[#f3efe6] text-[#23211d]">
    <div class="pdi-paper pointer-events-none"></div>
    <div class="relative mx-auto w-full max-w-md px-7 pb-14 pt-9">
      <header class="pdi-in flex items-center justify-between font-mono text-[9px] uppercase tracking-[0.28em] text-[#23211d]/45">
        <span>墨物 · INKWARE</span>
        <span>{{ nfcId || '無題' }}</span>
      </header>

      <section class="pdi-in pdi-d1 pdi-stage mt-7">
        <div class="pdi-moon"></div>
        <div class="pdi-mountain pdi-mountain-a"></div>
        <div class="pdi-mountain pdi-mountain-b"></div>
        <div class="pdi-object">
          <span class="pdi-object-body"></span>
          <span class="pdi-object-lid"></span>
          <span class="pdi-object-line"></span>
        </div>
        <span class="pdi-seal">墨<br />物</span>
        <span class="pdi-vertical">一器一世界</span>
      </section>

      <div class="pdi-in pdi-d2 mt-8 flex items-start gap-5">
        <h1 class="pdi-title flex-none text-[40px] font-normal leading-[1.15] tracking-[0.06em]">
          <span v-for="(ch, i) in titleChars" :key="i" class="block">{{ ch }}</span>
        </h1>
        <div class="flex-1 pt-1">
          <p class="text-[12px] leading-[2] text-[#23211d]/60">{{ p.subtitle || '—' }}</p>
          <div class="mt-4 flex items-baseline gap-3 border-t border-[#23211d]/25 pt-3">
            <span class="text-[30px] font-light leading-none tabular-nums">¥{{ p.price || '—' }}</span>
            <span v-if="p.originalPrice" class="font-mono text-[10px] text-[#23211d]/35 line-through">¥{{ p.originalPrice }}</span>
          </div>
          <p v-if="discount" class="mt-1.5 font-mono text-[9px] uppercase tracking-[0.12em] text-[#8a3b2e]">省 {{ discount }}%</p>
          <p class="mt-4 text-[11px] leading-[1.9] text-[#23211d]/50">{{ p.description }}</p>
        </div>
      </div>

      <section class="pdi-in pdi-d3 mt-9">
        <div class="flex items-center gap-3">
          <span class="h-px w-8 bg-[#8a3b2e]"></span>
          <h2 class="font-mono text-[9px] uppercase tracking-[0.26em] text-[#8a3b2e]">特 · 性</h2>
        </div>
        <ul class="mt-3">
          <li v-for="(f, i) in features" :key="i" class="flex items-baseline gap-4 border-b border-[#23211d]/12 py-3.5">
            <span class="font-serif text-[11px] text-[#8a3b2e]">{{ ['一', '二', '三', '四', '五'][i] || i + 1 }}</span>
            <span class="flex-1 text-[12.5px] leading-relaxed">{{ f }}</span>
          </li>
          <li v-if="!features.length" class="border-b border-[#23211d]/12 py-3.5 text-[12px] text-[#23211d]/40"></li>
        </ul>
      </section>

      <section class="pdi-in pdi-d3 mt-7">
        <div class="flex items-center gap-3">
          <span class="h-px w-8 bg-[#8a3b2e]"></span>
          <h2 class="font-mono text-[9px] uppercase tracking-[0.26em] text-[#8a3b2e]">器 · 度</h2>
        </div>
        <dl class="mt-3 grid grid-cols-2 gap-x-8">
          <div v-for="[k, v] in specs" :key="k" class="border-b border-[#23211d]/15 py-3">
            <dt class="font-mono text-[8.5px] uppercase tracking-[0.16em] text-[#23211d]/40">{{ k }}</dt>
            <dd class="mt-1 text-[13px]">{{ v }}</dd>
          </div>
          <div v-if="!specs.length" class="col-span-2 border-b border-[#23211d]/15 py-3 text-[12px] text-[#23211d]/40"></div>
        </dl>
      </section>

      <button class="pdi-in pdi-d4 pdi-buy mt-8 flex w-full items-center justify-between border-y border-[#23211d] px-1 py-4" @click="buy">
        <span class="text-[13px] tracking-[0.14em]">{{ p.ctaText || '立即购买' }}</span>
        <span class="flex h-8 w-8 items-center justify-center rounded-full border border-[#8a3b2e] text-[#8a3b2e]"><ArrowUpRight class="h-4 w-4" /></span>
      </button>

      <p class="mt-6 text-center font-mono text-[8px] uppercase tracking-[0.3em] text-[#23211d]/30">Handcrafted in Jingdezhen · 2026</p>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { ArrowUpRight } from 'lucide-vue-next'

const props = defineProps(['payload', 'nfcId'])
const p = computed(() => props.payload || {})
const titleChars = computed(() => String(p.value.title).slice(0, 6).split(''))
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
.pdi-root { position: relative; overflow: hidden; }
.pdi-paper {
  position: absolute;
  inset: 0;
  background-image: radial-gradient(rgba(35, 33, 29, 0.13) 0.6px, transparent 0.7px);
  background-size: 6px 6px;
  opacity: 0.5;
  pointer-events: none;
}
.pdi-stage { position: relative; aspect-ratio: 4 / 4.6; overflow: hidden; border: 1px solid rgba(35, 33, 29, 0.2); background: #eee9dd; }
.pdi-moon { position: absolute; right: 14%; top: 12%; width: 62px; height: 62px; border-radius: 999px; background: #f7f3e9; box-shadow: 0 0 0 1px rgba(35, 33, 29, 0.08); }
.pdi-mountain { position: absolute; bottom: 0; background: linear-gradient(180deg, rgba(35, 33, 29, 0.16), rgba(35, 33, 29, 0.05)); }
.pdi-mountain-a { left: -10%; width: 80%; height: 38%; clip-path: polygon(0 100%, 34% 18%, 56% 62%, 72% 36%, 100% 100%); }
.pdi-mountain-b { right: -12%; width: 86%; height: 30%; clip-path: polygon(0 100%, 28% 44%, 48% 70%, 70% 30%, 100% 100%); opacity: 0.7; }
.pdi-object { position: absolute; left: 50%; top: 52%; width: 39%; height: 45%; transform: translate(-50%, -50%); animation: pdiInk 8s ease-in-out infinite; }
.pdi-object-body {
  position: absolute;
  inset: 14% 0 0;
  border-radius: 44% 44% 38% 38% / 18% 18% 28% 28%;
  background: radial-gradient(80% 60% at 36% 24%, #4a4740, #211f1b 62%, #0f0e0c);
  box-shadow: 18px 24px 26px -18px rgba(35, 33, 29, 0.55), inset -14px -10px 26px rgba(0, 0, 0, 0.55);
}
.pdi-object-lid {
  position: absolute;
  left: 50%;
  top: 8%;
  width: 38%;
  height: 9%;
  transform: translateX(-50%);
  border-radius: 4px;
  background: #2b2924;
}
.pdi-object-line {
  position: absolute;
  left: 12%;
  right: 12%;
  top: 42%;
  height: 1px;
  background: rgba(243, 239, 230, 0.22);
}
.pdi-seal {
  position: absolute;
  right: 14px;
  bottom: 16px;
  display: grid;
  place-items: center;
  width: 38px;
  height: 38px;
  border: 1.5px solid #8a3b2e;
  color: #8a3b2e;
  font-family: serif;
  font-size: 11px;
  line-height: 1.15;
  text-align: center;
}
.pdi-vertical {
  position: absolute;
  left: 14px;
  top: 16px;
  writing-mode: vertical-rl;
  font-family: serif;
  font-size: 11px;
  letter-spacing: 0.28em;
  color: rgba(35, 33, 29, 0.5);
}
.pdi-title { font-family: 'Songti SC', 'Noto Serif SC', Georgia, serif; }
.pdi-buy { background: transparent; transition: background 0.3s ease, color 0.3s ease, padding 0.3s ease; }
.pdi-buy:hover { background: #23211d; color: #f3efe6; padding-inline: 14px; }
.pdi-buy:hover .pdi-buy span:last-child { border-color: #f3efe6; color: #f3efe6; }
.pdi-buy:active { transform: scale(0.99); }
.pdi-in { animation: pdiRise 0.9s cubic-bezier(0.22, 1, 0.36, 1) both; }
.pdi-d1 { animation-delay: 0.08s; }
.pdi-d2 { animation-delay: 0.16s; }
.pdi-d3 { animation-delay: 0.24s; }
.pdi-d4 { animation-delay: 0.32s; }
@keyframes pdiRise { from { opacity: 0; transform: translateY(18px); } to { opacity: 1; transform: none; } }
@keyframes pdiInk { 0%, 100% { transform: translate(-50%, -50%) scale(1); filter: blur(0); } 50% { transform: translate(-50%, -51%) scale(1.012); filter: blur(0.3px); } }
@media (prefers-reduced-motion: reduce) {
  .pdi-in, .pdi-object { animation: none !important; }
  .pdi-buy { transition: none; }
}
</style>
