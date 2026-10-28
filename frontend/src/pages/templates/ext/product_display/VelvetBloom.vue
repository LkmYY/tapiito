<template>
  <div class="pdv-root min-h-screen bg-[#fbeef1] text-[#4a2233]">
    <div class="pdv-halo pointer-events-none"></div>
    <div class="mx-auto w-full max-w-md px-6 pb-14 pt-8">
      <header class="pdv-in flex items-center justify-between text-[9px] uppercase tracking-[0.26em] text-[#a3566f]">
        <span class="font-serif italic normal-case tracking-[0.08em] text-[13px]">Maison Velvet</span>
        <span class="font-mono">{{ nfcId ? 'N° ' + nfcId : 'Édition' }}</span>
      </header>

      <section class="pdv-in pdv-d1 pdv-stage mt-6">
        <div class="pdv-arch"></div>
        <div class="pdv-pedestal"></div>
        <div class="pdv-bottle">
          <span class="pdv-cap"></span>
          <span class="pdv-neck"></span>
          <span class="pdv-body">
            <i class="pdv-face">VE<br /><small>L</small></i>
          </span>
          <span class="pdv-ribbon"></span>
        </div>
        <span class="pdv-petal pdv-petal-a"></span>
        <span class="pdv-petal pdv-petal-b"></span>
        <span class="pdv-petal pdv-petal-c"></span>
        <p class="pdv-engrave">EAU DE PARFUM · 50 ML</p>
      </section>

      <div class="pdv-in pdv-d2 mt-7 text-center">
        <p class="font-mono text-[8.5px] uppercase tracking-[0.3em] text-[#c08497]">{{ p.subtitle || '—' }}</p>
        <h1 class="mt-2 font-serif text-[37px] leading-[1.08]">{{ p.title }}</h1>
        <p class="mt-2 font-serif text-[13px] italic text-[#7a4354]/60">softness, bottled</p>
      </div>

      <div class="pdv-in pdv-d2 mt-5 flex flex-col items-center">
        <div class="flex items-baseline gap-3">
          <span class="font-serif text-[40px] leading-none tabular-nums">¥{{ p.price || '—' }}</span>
          <span v-if="p.originalPrice" class="font-mono text-[11px] text-[#7a4354]/35 line-through">¥{{ p.originalPrice }}</span>
        </div>
        <span v-if="discount" class="mt-2 rounded-full border border-[#c08497]/40 bg-white/50 px-3 py-0.5 font-mono text-[8px] uppercase tracking-[0.16em] text-[#a3566f]">Rituel −{{ discount }}%</span>
      </div>

      <p v-if="p.description" class="pdv-in pdv-d3 mt-5 px-3 text-center text-[12px] leading-[2] text-[#4a2233]/60">{{ p.description }}</p>

      <section class="pdv-in pdv-d3 mt-7 rounded-[22px] border border-[#d8a3b4]/50 bg-white/45 p-5">
        <h2 class="text-center font-serif text-[15px] italic text-[#a3566f]">The ritual</h2>
        <ul class="mt-3">
          <li v-for="(f, i) in features" :key="i" class="flex items-center gap-3 border-b border-[#d8a3b4]/35 py-2.5 last:border-0">
            <span class="flex h-5 w-5 flex-none items-center justify-center rounded-full bg-[#f3d3dc] font-mono text-[7.5px] text-[#a3566f]">{{ String(i + 1).padStart(2, '0') }}</span>
            <span class="flex-1 text-[12px] leading-relaxed">{{ f }}</span>
            <Sparkle class="h-3 w-3 flex-none text-[#c08497]" />
          </li>
          <li v-if="!features.length" class="py-2 text-center text-[12px] text-[#7a4354]/40"></li>
        </ul>
      </section>

      <section class="pdv-in pdv-d3 mt-4">
        <div class="flex items-center gap-3">
          <span class="h-px flex-1 bg-[#c08497]/30"></span>
          <h2 class="font-mono text-[8px] uppercase tracking-[0.26em] text-[#a3566f]">Composition</h2>
          <span class="h-px flex-1 bg-[#c08497]/30"></span>
        </div>
        <dl class="mt-3 grid grid-cols-2 gap-x-5">
          <div v-for="[k, v] in specs" :key="k" class="flex items-baseline justify-between border-b border-[#c08497]/30 py-2.5">
            <dt class="font-mono text-[8px] uppercase tracking-[0.1em] text-[#7a4354]/50">{{ k }}</dt>
            <dd class="font-serif text-[13px]">{{ v }}</dd>
          </div>
          <div v-if="!specs.length" class="col-span-2 border-b border-[#c08497]/30 py-2.5 text-center text-[11px] text-[#7a4354]/40"></div>
        </dl>
      </section>

      <button class="pdv-in pdv-d4 pdv-buy mt-6 flex w-full items-center justify-center gap-2.5 py-4" @click="buy">
        <span class="text-[12px] font-medium tracking-[0.1em]">{{ p.ctaText || '立即购买' }}</span>
        <ArrowUpRight class="h-3.5 w-3.5" />
      </button>

      <p class="mt-5 text-center font-mono text-[8px] uppercase tracking-[0.24em] text-[#7a4354]/30">Paris · Grasse · Cruelty free</p>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { ArrowUpRight, Sparkle } from 'lucide-vue-next'

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
.pdv-root { position: relative; overflow: hidden; }
.pdv-halo {
  position: absolute;
  top: -90px;
  left: 50%;
  width: 420px;
  height: 420px;
  transform: translateX(-50%);
  border-radius: 999px;
  background: radial-gradient(circle, rgba(232, 168, 188, 0.55), transparent 68%);
  filter: blur(18px);
}
.pdv-stage {
  position: relative;
  aspect-ratio: 1 / 1.1;
  overflow: hidden;
  border-radius: 130px 130px 22px 22px;
  background: linear-gradient(180deg, #f7dbe3 0%, #f2c9d5 55%, #e7b2c2 100%);
  box-shadow: inset 0 0 0 1px rgba(163, 86, 111, 0.12), 0 28px 50px -32px rgba(122, 67, 84, 0.5);
}
.pdv-arch {
  position: absolute;
  left: 50%;
  top: 8%;
  width: 74%;
  height: 84%;
  transform: translateX(-50%);
  border-radius: 999px 999px 0 0;
  background: linear-gradient(180deg, rgba(255, 255, 255, 0.55), rgba(255, 255, 255, 0.05));
  border: 1px solid rgba(255, 255, 255, 0.7);
}
.pdv-pedestal {
  position: absolute;
  left: 50%;
  bottom: 13%;
  width: 52%;
  height: 13%;
  transform: translateX(-50%);
  border-radius: 50%;
  background: linear-gradient(180deg, #f9e2e8, #e4aebb);
  box-shadow: 0 15px 22px -12px rgba(122, 67, 84, 0.45);
}
.pdv-bottle { position: absolute; left: 50%; bottom: 20%; width: 34%; height: 52%; transform: translateX(-50%); animation: pdvFloat 6s ease-in-out infinite; }
.pdv-cap {
  position: absolute;
  left: 28%;
  top: 0;
  width: 44%;
  height: 16%;
  border-radius: 5px 5px 2px 2px;
  background: linear-gradient(100deg, #e9c7a0, #c69d72 45%, #a87e55);
  box-shadow: inset 0 1px 0 rgba(255, 255, 255, 0.6);
}
.pdv-neck {
  position: absolute;
  left: 37%;
  top: 15%;
  width: 26%;
  height: 9%;
  background: rgba(255, 255, 255, 0.6);
  border-inline: 1px solid rgba(163, 86, 111, 0.15);
}
.pdv-body {
  position: absolute;
  inset: 22% 0 0;
  border-radius: 16% 16% 12% 12% / 8% 8% 10% 10%;
  background: linear-gradient(100deg, rgba(255, 255, 255, 0.95) 0%, rgba(252, 228, 236, 0.8) 40%, rgba(216, 163, 180, 0.65) 100%);
  border: 1px solid rgba(255, 255, 255, 0.9);
  box-shadow: inset -12px -6px 22px rgba(163, 86, 111, 0.25), inset 10px 8px 20px rgba(255, 255, 255, 0.9), 0 18px 26px -16px rgba(122, 67, 84, 0.55);
}
.pdv-face {
  position: absolute;
  left: 50%;
  top: 42%;
  transform: translate(-50%, -50%);
  display: flex;
  align-items: baseline;
  justify-content: center;
  width: 52%;
  height: 34%;
  border: 1px solid rgba(163, 86, 111, 0.35);
  background: rgba(255, 246, 249, 0.65);
  font-family: Georgia, serif;
  font-size: 16px;
  letter-spacing: 0.06em;
  color: #8b4a60;
  text-align: center;
}
.pdv-face small { font-size: 9px; }
.pdv-ribbon {
  position: absolute;
  left: 50%;
  bottom: 20%;
  width: 78%;
  height: 7%;
  transform: translateX(-50%) rotate(-2deg);
  background: rgba(163, 86, 111, 0.55);
  border-radius: 999px;
}
.pdv-petal { position: absolute; width: 46px; height: 26px; border-radius: 100% 0 100% 0; background: rgba(214, 134, 158, 0.45); }
.pdv-petal-a { left: 12%; top: 22%; transform: rotate(-24deg); animation: pdvDrift 9s ease-in-out infinite; }
.pdv-petal-b { right: 10%; top: 38%; transform: rotate(160deg) scale(0.8); animation: pdvDrift 11s ease-in-out infinite reverse; }
.pdv-petal-c { left: 20%; bottom: 22%; transform: rotate(40deg) scale(0.65); animation: pdvDrift 13s ease-in-out infinite; }
.pdv-engrave {
  position: absolute;
  left: 50%;
  bottom: 8%;
  transform: translateX(-50%);
  white-space: nowrap;
  font-family: ui-monospace, monospace;
  font-size: 7.5px;
  letter-spacing: 0.26em;
  color: rgba(122, 67, 84, 0.5);
}
.pdv-buy {
  border-radius: 999px;
  background: #4a2233;
  color: #fbeef1;
  box-shadow: 0 16px 28px -16px rgba(74, 34, 51, 0.75);
  transition: background 0.25s ease, transform 0.18s ease;
}
.pdv-buy:hover { background: #6d334a; }
.pdv-buy:active { transform: scale(0.985); }
.pdv-in { animation: pdvIn 0.8s cubic-bezier(0.22, 1, 0.36, 1) both; }
.pdv-d1 { animation-delay: 0.06s; }
.pdv-d2 { animation-delay: 0.14s; }
.pdv-d3 { animation-delay: 0.22s; }
.pdv-d4 { animation-delay: 0.3s; }
@keyframes pdvIn { from { opacity: 0; transform: translateY(18px); } to { opacity: 1; transform: none; } }
@keyframes pdvFloat { 0%, 100% { transform: translateX(-50%) translateY(0); } 50% { transform: translateX(-50%) translateY(-7px); } }
@keyframes pdvDrift { 0%, 100% { translate: 0 0; } 50% { translate: 5px 9px; } }
@media (prefers-reduced-motion: reduce) {
  .pdv-in, .pdv-bottle, .pdv-petal { animation: none !important; }
  .pdv-buy { transition: none; }
}
</style>
