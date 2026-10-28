<template>
  <div class="pdc-root min-h-screen bg-[#f2e8da] text-[#4a2f24]">
    <div class="pdc-sand pointer-events-none"></div>
    <div class="mx-auto w-full max-w-md px-6 pb-14 pt-8">
      <header class="pdc-in flex items-center justify-between font-mono text-[9px] uppercase tracking-[0.2em] text-[#4a2f24]/50">
        <span>Clay & Co.</span>
        <span>{{ nfcId ? 'Piece ' + nfcId : 'Handmade' }}</span>
      </header>

      <section class="pdc-in pdc-d1 pdc-stage mt-6">
        <div class="pdc-arch"></div>
        <div class="pdc-vase">
          <span class="pdc-vase-rim"></span>
          <span class="pdc-vase-body"></span>
          <span class="pdc-vase-wrap"></span>
          <span class="pdc-vase-grain"></span>
        </div>
        <span class="pdc-badge">HAND<br />THROWN</span>
        <div class="pdc-floor"></div>
        <span class="pdc-anno">STONEWARE · 2026</span>
      </section>

      <div class="pdc-in pdc-d2 mt-7 text-center">
        <p class="font-mono text-[9px] uppercase tracking-[0.28em] text-[#a8582f]">{{ p.subtitle || '—' }}</p>
        <h1 class="mt-2 font-serif text-[36px] leading-[1.1]">{{ p.title }}</h1>
        <p class="mt-3 text-[12px] italic text-[#4a2f24]/55">slow made · one of a kind</p>
      </div>

      <div class="pdc-in pdc-d2 mt-6 flex items-center justify-center gap-5">
        <span class="text-[36px] font-semibold leading-none tracking-tight tabular-nums">¥{{ p.price || '—' }}</span>
        <span v-if="p.originalPrice" class="font-mono text-[12px] text-[#4a2f24]/35 line-through">¥{{ p.originalPrice }}</span>
        <span v-if="discount" class="rounded-full bg-[#a8582f] px-2.5 py-1 font-mono text-[8px] uppercase tracking-[0.12em] text-[#f9f3e9]">−{{ discount }}%</span>
      </div>

      <p v-if="p.description" class="pdc-in pdc-d3 mt-5 px-2 text-center text-[12.5px] leading-[2] text-[#4a2f24]/65">{{ p.description }}</p>

      <section class="pdc-in pdc-d3 mt-7">
        <div class="flex items-center gap-3">
          <span class="h-px flex-1 bg-[#4a2f24]/25"></span>
          <h2 class="font-mono text-[8.5px] uppercase tracking-[0.26em] text-[#a8582f]">Little details</h2>
          <span class="h-px flex-1 bg-[#4a2f24]/25"></span>
        </div>
        <ul class="mt-3 grid gap-2">
          <li v-for="(f, i) in features" :key="i" class="pdc-feature">
            <span class="pdc-feature-no">{{ String(i + 1).padStart(2, '0') }}</span>
            <span class="text-[12.5px] leading-relaxed">{{ f }}</span>
            <Leaf class="ml-auto h-3.5 w-3.5 flex-none text-[#a8582f]/50" />
          </li>
          <li v-if="!features.length" class="pdc-feature text-[12px] text-[#4a2f24]/40"></li>
        </ul>
      </section>

      <section class="pdc-in pdc-d3 mt-6">
        <h2 class="text-center font-mono text-[8.5px] uppercase tracking-[0.26em] text-[#a8582f]">Made to measure</h2>
        <dl class="mt-3 grid grid-cols-2 gap-2">
          <div v-for="[k, v] in specs" :key="k" class="pdc-spec">
            <dt class="font-mono text-[8px] uppercase tracking-[0.12em] text-[#4a2f24]/45">{{ k }}</dt>
            <dd class="mt-1 font-serif text-[14px]">{{ v }}</dd>
          </div>
          <div v-if="!specs.length" class="pdc-spec col-span-2 text-[12px] text-[#4a2f24]/40"></div>
        </dl>
      </section>

      <button class="pdc-in pdc-d4 pdc-buy mt-7 flex w-full items-center justify-center gap-3 py-4" @click="buy">
        <span class="text-[13px] font-semibold tracking-[0.04em]">{{ p.ctaText || '立即购买' }}</span>
        <ArrowUpRight class="h-4 w-4" />
      </button>

      <p class="mt-5 text-center font-mono text-[8px] uppercase tracking-[0.24em] text-[#4a2f24]/35">Fired in small batches · Jingdezhen clay</p>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { ArrowUpRight, Leaf } from 'lucide-vue-next'

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
.pdc-root { position: relative; overflow: hidden; }
.pdc-sand {
  position: absolute;
  inset: 0;
  background-image: radial-gradient(rgba(74, 47, 36, 0.07) 0.8px, transparent 0.9px);
  background-size: 9px 9px;
  opacity: 0.7;
  pointer-events: none;
}
.pdc-stage {
  position: relative;
  aspect-ratio: 1 / 1.12;
  overflow: hidden;
  border-radius: 48% 48% 18px 18px / 26% 26% 18px 18px;
  background: linear-gradient(180deg, #e8d5bf 0%, #e2cbb1 62%, #d4b696 100%);
  box-shadow: inset 0 0 0 1px rgba(74, 47, 36, 0.12), 0 26px 50px -32px rgba(74, 47, 36, 0.5);
}
.pdc-arch {
  position: absolute;
  left: 50%;
  top: 12%;
  width: 72%;
  aspect-ratio: 1 / 1.25;
  transform: translateX(-50%);
  border-radius: 999px 999px 0 0;
  background: rgba(249, 243, 233, 0.38);
  border: 1px solid rgba(74, 47, 36, 0.12);
}
.pdc-floor {
  position: absolute;
  left: 0;
  right: 0;
  bottom: 0;
  height: 22%;
  background: linear-gradient(180deg, rgba(74, 47, 36, 0.04), rgba(74, 47, 36, 0.14));
}
.pdc-vase { position: absolute; left: 50%; top: 52%; width: 34%; height: 56%; transform: translate(-50%, -50%); animation: pdcRise 0.9s cubic-bezier(0.22, 1, 0.36, 1) both; }
.pdc-vase-body {
  position: absolute;
  inset: 6% 0 0;
  border-radius: 46% 46% 34% 34% / 22% 22% 28% 28%;
  background: linear-gradient(100deg, #c97b4e 0%, #b9673c 40%, #8f4a28 100%);
  box-shadow: inset -14px -6px 28px rgba(64, 30, 14, 0.35), inset 12px 8px 20px rgba(255, 226, 199, 0.35), 0 22px 30px -20px rgba(74, 47, 36, 0.75);
}
.pdc-vase-rim {
  position: absolute;
  left: 50%;
  top: 1%;
  width: 46%;
  height: 9%;
  transform: translateX(-50%);
  border-radius: 50%;
  background: linear-gradient(180deg, #d9966c, #9c5028);
  border: 1px solid rgba(74, 47, 36, 0.3);
}
.pdc-vase-wrap {
  position: absolute;
  left: 18%;
  right: 18%;
  top: 54%;
  height: 12%;
  border-radius: 999px;
  background: rgba(249, 243, 233, 0.85);
  box-shadow: 0 3px 10px rgba(74, 47, 36, 0.18);
  transform: rotate(-3deg);
}
.pdc-vase-grain {
  position: absolute;
  inset: 6% 0 0;
  border-radius: 46% 46% 34% 34% / 22% 22% 28% 28%;
  opacity: 0.35;
  background-image: radial-gradient(rgba(255, 255, 255, 0.35) 0.7px, transparent 0.8px);
  background-size: 5px 5px;
}
.pdc-badge {
  position: absolute;
  right: 16px;
  top: 16px;
  display: grid;
  place-items: center;
  width: 60px;
  height: 60px;
  border-radius: 999px;
  border: 1px dashed rgba(74, 47, 36, 0.5);
  font-family: ui-monospace, monospace;
  font-size: 7.5px;
  line-height: 1.5;
  letter-spacing: 0.14em;
  text-align: center;
  color: #4a2f24;
  transform: rotate(8deg);
  background: rgba(249, 243, 233, 0.45);
}
.pdc-anno { position: absolute; left: 16px; bottom: 12px; font-family: ui-monospace, monospace; font-size: 7.5px; letter-spacing: 0.2em; color: rgba(74, 47, 36, 0.5); }
.pdc-feature {
  display: flex;
  align-items: center;
  gap: 10px;
  border-radius: 999px;
  border: 1px solid rgba(74, 47, 36, 0.14);
  background: rgba(249, 243, 233, 0.6);
  padding: 10px 14px;
  transition: background 0.2s ease, transform 0.2s ease;
}
.pdc-feature:hover { background: #f9f3e9; transform: translateX(2px); }
.pdc-feature-no { font-family: ui-monospace, monospace; font-size: 8px; color: #a8582f; }
.pdc-spec { border-radius: 14px; border: 1px solid rgba(74, 47, 36, 0.16); background: rgba(249, 243, 233, 0.5); padding: 11px 13px; text-align: center; }
.pdc-buy {
  border-radius: 999px;
  background: #4a2f24;
  color: #f9f3e9;
  box-shadow: 0 16px 26px -16px rgba(74, 47, 36, 0.8);
  transition: background 0.25s ease, transform 0.18s ease;
}
.pdc-buy:hover { background: #6b4534; }
.pdc-buy:active { transform: scale(0.985); }
.pdc-in { animation: pdcIn 0.75s cubic-bezier(0.22, 1, 0.36, 1) both; }
.pdc-d1 { animation-delay: 0.06s; }
.pdc-d2 { animation-delay: 0.13s; }
.pdc-d3 { animation-delay: 0.2s; }
.pdc-d4 { animation-delay: 0.27s; }
@keyframes pdcIn { from { opacity: 0; transform: translateY(16px); } to { opacity: 1; transform: none; } }
@keyframes pdcRise { from { opacity: 0; transform: translate(-50%, -44%); } to { opacity: 1; transform: translate(-50%, -50%); } }
@media (prefers-reduced-motion: reduce) {
  .pdc-in, .pdc-vase { animation: none !important; }
  .pdc-buy, .pdc-feature { transition: none; }
}
</style>
