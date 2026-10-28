<template>
  <div class="pdg-root min-h-screen bg-[#f3f3f0] text-[#161615]">
    <div class="mx-auto w-full max-w-md px-6 pb-14 pt-7">
      <header class="pdg-in flex items-center justify-between border-b border-[#161615]/15 pb-3 font-mono text-[10px] uppercase tracking-[0.22em] text-[#161615]/50">
        <span>Object / 01</span>
        <span>NFC {{ nfcId || '—' }}</span>
      </header>

      <section class="pdg-in pdg-d1 pdg-stage mt-6">
        <div class="pdg-grid"></div>
        <span class="pdg-index">GS—27</span>
        <div class="pdg-halo"></div>
        <div class="pdg-band"></div>
        <div class="pdg-cup pdg-cup-l"><i></i></div>
        <div class="pdg-cup pdg-cup-r"><i></i></div>
        <span class="pdg-note">ACOUSTIC OBJECT</span>
      </section>

      <div class="pdg-in pdg-d2 mt-7 flex items-start justify-between gap-4">
        <div class="min-w-0">
          <h1 class="text-[34px] font-semibold leading-[1.02] tracking-[-0.03em]">{{ p.title }}</h1>
          <p class="mt-2 text-[12px] leading-relaxed text-[#161615]/55">{{ p.subtitle || '—' }}</p>
        </div>
        <span class="mt-1 flex h-9 w-9 flex-none items-center justify-center rounded-full border border-[#161615]/20">
          <ArrowUpRight class="h-4 w-4" />
        </span>
      </div>

      <div class="pdg-in pdg-d2 mt-7 flex items-end justify-between border-y border-[#161615] py-5">
        <div class="flex items-baseline gap-3">
          <span class="text-[11px] font-medium">CNY</span>
          <span class="text-[46px] font-semibold leading-none tracking-[-0.05em] tabular-nums">¥{{ p.price || '—' }}</span>
        </div>
        <div class="text-right">
          <p v-if="p.originalPrice" class="font-mono text-[11px] text-[#161615]/40 line-through">¥{{ p.originalPrice }}</p>
          <p v-if="discount" class="mt-1 font-mono text-[10px] uppercase tracking-[0.12em] text-[#c2410c]">Save {{ discount }}%</p>
        </div>
      </div>

      <section class="pdg-in pdg-d3 mt-7">
        <div class="flex items-baseline justify-between">
          <h2 class="font-mono text-[10px] uppercase tracking-[0.22em]">Features</h2>
          <span class="font-mono text-[10px] text-[#161615]/35">{{ String(features.length).padStart(2, '0') }}</span>
        </div>
        <ul class="mt-3">
          <li v-for="(f, i) in features" :key="i" class="flex gap-4 border-t border-[#161615]/12 py-3">
            <span class="font-mono text-[10px] text-[#161615]/35">{{ String(i + 1).padStart(2, '0') }}</span>
            <span class="text-[13px] leading-snug">{{ f }}</span>
          </li>
          <li v-if="!features.length" class="border-t border-[#161615]/12 py-3 text-[13px] text-[#161615]/40"></li>
        </ul>
      </section>

      <section class="pdg-in pdg-d4 mt-7">
        <h2 class="font-mono text-[10px] uppercase tracking-[0.22em]">Specification</h2>
        <dl class="mt-3 grid grid-cols-2 border-l border-t border-[#161615]/15">
          <div v-for="[k, v] in specs" :key="k" class="border-b border-r border-[#161615]/15 p-3.5">
            <dt class="font-mono text-[9px] uppercase tracking-[0.12em] text-[#161615]/40">{{ k }}</dt>
            <dd class="mt-1.5 text-[13px] font-medium">{{ v }}</dd>
          </div>
          <div v-if="!specs.length" class="col-span-2 border-b border-r border-[#161615]/15 p-3.5 text-[13px] text-[#161615]/40"></div>
        </dl>
      </section>

      <p v-if="p.description" class="pdg-in pdg-d4 mt-7 border-l-2 border-[#161615] pl-4 text-[12px] leading-[1.8] text-[#161615]/60">{{ p.description }}</p>

      <button class="pdg-in pdg-d5 pdg-buy mt-8 flex w-full items-center justify-between px-6 py-4 text-[13px] font-medium tracking-wide" @click="buy">
        <span>{{ p.ctaText || '立即购买' }}</span>
        <span class="pdg-buy-ico"><ArrowUpRight class="h-4 w-4" /></span>
      </button>

      <p class="mt-6 text-center font-mono text-[9px] uppercase tracking-[0.3em] text-[#161615]/30">Industrial design · 2026</p>
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
.pdg-stage {
  position: relative;
  aspect-ratio: 1 / 1.06;
  overflow: hidden;
  border: 1px solid rgba(22, 22, 21, 0.14);
  background: radial-gradient(circle at 50% 42%, #ffffff 0%, #ebeae5 55%, #dedcd4 100%);
}
.pdg-grid {
  position: absolute;
  inset: 0;
  background-image: linear-gradient(rgba(22, 22, 21, 0.055) 1px, transparent 1px),
    linear-gradient(90deg, rgba(22, 22, 21, 0.055) 1px, transparent 1px);
  background-size: 28px 28px;
  mask-image: radial-gradient(circle at 50% 50%, black, transparent 78%);
}
.pdg-index {
  position: absolute;
  left: 14px;
  top: 12px;
  font-family: ui-monospace, monospace;
  font-size: 9px;
  letter-spacing: 0.18em;
  color: rgba(22, 22, 21, 0.4);
}
.pdg-note {
  position: absolute;
  right: 14px;
  bottom: 12px;
  font-family: ui-monospace, monospace;
  font-size: 8px;
  letter-spacing: 0.24em;
  color: rgba(22, 22, 21, 0.32);
}
.pdg-halo {
  position: absolute;
  left: 50%;
  top: 50%;
  width: 66%;
  aspect-ratio: 1;
  transform: translate(-50%, -50%);
  border-radius: 999px;
  background: radial-gradient(circle, rgba(255, 255, 255, 0.9), rgba(255, 255, 255, 0) 70%);
  box-shadow: 0 0 0 1px rgba(22, 22, 21, 0.06), 0 0 0 44px rgba(22, 22, 21, 0.025);
}
.pdg-band {
  position: absolute;
  left: 50%;
  top: 24%;
  width: 44%;
  height: 42%;
  transform: translateX(-50%);
  border: 9px solid #232320;
  border-bottom: 0;
  border-radius: 999px 999px 0 0;
  box-shadow: inset 0 3px 0 rgba(255, 255, 255, 0.18);
}
.pdg-cup {
  position: absolute;
  top: 50%;
  width: 27%;
  height: 34%;
  transform: translateY(-32%);
  border-radius: 22% 22% 26% 26%;
  background: linear-gradient(145deg, #34342f 0%, #191917 46%, #0c0c0b 100%);
  box-shadow: inset 0 1px 0 rgba(255, 255, 255, 0.22), inset -6px 0 12px rgba(0, 0, 0, 0.5), 0 18px 26px -18px rgba(0, 0, 0, 0.6);
  animation: pdgFloat 6s ease-in-out infinite;
}
.pdg-cup-l { left: 16%; }
.pdg-cup-r { right: 16%; animation-delay: -3s; }
.pdg-cup i {
  position: absolute;
  left: 50%;
  top: 50%;
  width: 58%;
  height: 62%;
  transform: translate(-50%, -50%);
  border-radius: 18%;
  background: repeating-linear-gradient(0deg, rgba(255, 255, 255, 0.07) 0 1px, transparent 1px 4px), linear-gradient(160deg, #3d3d37, #131311);
  box-shadow: inset 0 0 0 1px rgba(255, 255, 255, 0.08);
}
.pdg-buy {
  background: #161615;
  color: #f5f4ef;
  transition: background 0.25s ease, transform 0.18s ease;
}
.pdg-buy:hover { background: #2c2c29; }
.pdg-buy:active { transform: scale(0.985); }
.pdg-buy-ico {
  display: grid;
  place-items: center;
  width: 34px;
  height: 34px;
  border-radius: 999px;
  background: #f5f4ef;
  color: #161615;
  transition: transform 0.25s ease;
}
.pdg-buy:hover .pdg-buy-ico { transform: rotate(45deg); }

.pdg-in { animation: pdgIn 0.7s cubic-bezier(0.22, 1, 0.36, 1) both; }
.pdg-d1 { animation-delay: 0.06s; }
.pdg-d2 { animation-delay: 0.13s; }
.pdg-d3 { animation-delay: 0.2s; }
.pdg-d4 { animation-delay: 0.27s; }
.pdg-d5 { animation-delay: 0.34s; }

@keyframes pdgIn {
  from { opacity: 0; transform: translateY(16px); }
  to { opacity: 1; transform: none; }
}
@keyframes pdgFloat {
  0%, 100% { transform: translateY(-32%); }
  50% { transform: translateY(-36%); }
}
@media (prefers-reduced-motion: reduce) {
  .pdg-in, .pdg-cup { animation: none !important; }
  .pdg-buy, .pdg-buy-ico { transition: none; }
}
</style>
