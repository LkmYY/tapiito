<template>
  <div class="pdbo-root min-h-screen bg-[#f0f2e6] text-[#243527]">
    <div class="pdbo-leaf-deco pointer-events-none">
      <i class="pdbo-leaf pdbo-leaf-a"></i>
      <i class="pdbo-leaf pdbo-leaf-b"></i>
      <i class="pdbo-leaf pdbo-leaf-c"></i>
    </div>
    <div class="relative mx-auto w-full max-w-md px-6 pb-14 pt-8">
      <header class="pdbo-in flex items-center justify-between font-mono text-[9px] uppercase tracking-[0.24em] text-[#243527]/50">
        <span class="flex items-center gap-2 font-serif text-[13px] normal-case italic tracking-normal text-[#243527]"><Leaf class="h-4 w-4 text-[#4d7c45]" /> Herbier</span>
        <span>{{ nfcId || 'BOTANICAL 07' }}</span>
      </header>

      <section class="pdbo-in pdbo-d1 pdbo-stage mt-6">
        <div class="pdbo-green"></div>
        <span class="pdbo-ring"></span>
        <div class="pdbo-vessel">
          <span class="pdbo-dropper"></span>
          <span class="pdbo-bottle">
            <i class="pdbo-liquid"></i>
            <i class="pdbo-label">H<br /><small>07</small></i>
          </span>
        </div>
        <span class="pdbo-sprig pdbo-sprig-a"></span>
        <span class="pdbo-sprig pdbo-sprig-b"></span>
        <span class="pdbo-latin">OLEUM HERBARUM</span>
      </section>

      <div class="pdbo-in pdbo-d2 mt-6 flex items-start justify-between gap-4">
        <div class="min-w-0">
          <p class="font-mono text-[8.5px] uppercase tracking-[0.22em] text-[#4d7c45]">{{ p.subtitle || '—' }}</p>
          <h1 class="mt-2 font-serif text-[34px] leading-[1.05]">{{ p.title }}</h1>
        </div>
        <div class="flex-none rounded-full border border-[#243527]/20 px-3 py-2 text-center">
          <p class="font-mono text-[7px] uppercase tracking-[0.1em] text-[#243527]/40">From</p>
          <p class="mt-0.5 text-[18px] font-semibold leading-none tabular-nums">¥{{ p.price || '—' }}</p>
        </div>
      </div>

      <div class="pdbo-in pdbo-d2 mt-4 flex items-center gap-3 border-y border-[#243527]/20 py-2.5 font-mono text-[9px] uppercase tracking-[0.14em]">
        <span class="flex items-center gap-1.5"><Sprout class="h-3.5 w-3.5 text-[#4d7c45]" /> Small batch</span>
        <span class="h-3 w-px bg-[#243527]/20"></span>
        <span>Cold pressed</span>
        <span v-if="p.originalPrice" class="ml-auto text-[#243527]/40 line-through">¥{{ p.originalPrice }}</span>
        <span v-if="discount" class="rounded-full bg-[#4d7c45] px-2 py-0.5 text-[8px] text-[#f0f2e6]">−{{ discount }}%</span>
      </div>

      <p v-if="p.description" class="pdbo-in pdbo-d3 mt-5 font-serif text-[13px] leading-[2] text-[#243527]/65">{{ p.description }}</p>

      <section class="pdbo-in pdbo-d3 mt-6">
        <div class="flex items-center gap-3">
          <h2 class="font-mono text-[8.5px] uppercase tracking-[0.24em] text-[#4d7c45]">In the bottle</h2>
          <span class="h-px flex-1 bg-[#243527]/20"></span>
        </div>
        <ul class="mt-2 grid gap-1.5">
          <li v-for="(f, i) in features" :key="i" class="flex items-center gap-3 rounded-xl border border-[#4d7c45]/20 bg-[#e7ecd9]/60 px-3 py-2.5">
            <span class="flex h-5 w-5 flex-none items-center justify-center rounded-full bg-[#4d7c45] font-mono text-[7px] text-[#f0f2e6]">{{ String(i + 1).padStart(2, '0') }}</span>
            <span class="flex-1 text-[11.5px] leading-snug">{{ f }}</span>
          </li>
          <li v-if="!features.length" class="rounded-xl border border-[#4d7c45]/20 px-3 py-2.5 text-[11.5px] text-[#243527]/40"></li>
        </ul>
      </section>

      <section class="pdbo-in pdbo-d3 mt-5 grid grid-cols-2 gap-x-6 gap-y-0">
        <div v-for="[k, v] in specs" :key="k" class="border-b border-[#243527]/15 py-2.5">
          <dt class="font-mono text-[7.5px] uppercase tracking-[0.12em] text-[#243527]/45">{{ k }}</dt>
          <dd class="mt-0.5 font-serif text-[13px]">{{ v }}</dd>
        </div>
        <div v-if="!specs.length" class="col-span-2 border-b border-[#243527]/15 py-2.5 text-[11.5px] text-[#243527]/40"></div>
      </section>

      <button class="pdbo-in pdbo-d4 pdbo-buy mt-6 flex w-full items-center justify-between px-5 py-4" @click="buy">
        <span class="flex items-center gap-2 text-[12px] font-medium tracking-[0.06em]"><Leaf class="h-4 w-4" /> {{ p.ctaText || '立即购买' }}</span>
        <ArrowUpRight class="h-4 w-4" />
      </button>

      <p class="mt-5 text-center font-mono text-[8px] uppercase tracking-[0.24em] text-[#243527]/30">Harvested by hand · distilled slowly · Herbier</p>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { ArrowUpRight, Leaf, Sprout } from 'lucide-vue-next'

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
.pdbo-root { position: relative; overflow: hidden; }
.pdbo-leaf-deco { position: absolute; inset: 0; }
.pdbo-leaf { position: absolute; display: block; width: 90px; height: 46px; border-radius: 100% 0 100% 0; background: rgba(77, 124, 69, 0.1); }
.pdbo-leaf-a { right: -28px; top: 44px; transform: rotate(-24deg); }
.pdbo-leaf-b { left: -40px; bottom: 140px; transform: rotate(38deg) scale(1.3); background: rgba(77, 124, 69, 0.07); }
.pdbo-leaf-c { right: -30px; bottom: 40px; transform: rotate(12deg) scale(0.8); }
.pdbo-stage { position: relative; aspect-ratio: 1 / 1.08; overflow: hidden; border-radius: 999px 999px 26px 26px; background: linear-gradient(160deg, #dfe7cd, #cbd9b7 65%, #b9caa4); box-shadow: inset 0 0 0 1px rgba(36, 53, 39, 0.12), 0 26px 46px -32px rgba(36, 53, 39, 0.55); }
.pdbo-green { position: absolute; left: 50%; top: 50%; width: 84%; aspect-ratio: 1; transform: translate(-50%, -50%); border-radius: 999px; background: radial-gradient(circle, #8fae76 0%, #6c9159 58%, transparent 72%); opacity: 0.85; }
.pdbo-ring { position: absolute; left: 50%; top: 50%; width: 66%; aspect-ratio: 1; transform: translate(-50%, -50%); border: 1px solid rgba(240, 242, 230, 0.6); border-radius: 999px; }
.pdbo-vessel { position: absolute; left: 50%; top: 50%; width: 34%; height: 58%; transform: translate(-50%, -50%); animation: pdboFloat 7s ease-in-out infinite; }
.pdbo-dropper { position: absolute; left: 32%; top: 0; width: 36%; height: 22%; border-radius: 5px 5px 0 0; background: #243527; }
.pdbo-dropper::after { content: ''; position: absolute; left: 50%; bottom: -14%; width: 2px; height: 24%; transform: translateX(-50%); background: #243527; }
.pdbo-bottle {
  position: absolute;
  inset: 22% 0 0;
  border-radius: 12px 12px 20px 20px;
  background: linear-gradient(100deg, rgba(251, 252, 244, 0.95), rgba(219, 227, 201, 0.75) 55%, rgba(160, 181, 138, 0.65));
  border: 1px solid rgba(36, 53, 39, 0.2);
  box-shadow: inset -12px 0 20px rgba(36, 53, 39, 0.18), inset 8px 0 16px rgba(255, 255, 255, 0.7), 12px 18px 24px -16px rgba(36, 53, 39, 0.6);
}
.pdbo-liquid { position: absolute; left: 10%; right: 10%; bottom: 6%; height: 42%; border-radius: 6px 6px 15px 15px; background: linear-gradient(180deg, rgba(149, 176, 108, 0.75), rgba(94, 122, 62, 0.85)); }
.pdbo-label { position: absolute; left: 50%; top: 34%; transform: translate(-50%, -50%); display: grid; place-items: center; width: 56%; height: 30%; background: rgba(240, 242, 230, 0.9); font-family: Georgia, serif; font-size: 16px; line-height: 1; color: #243527; }
.pdbo-label small { font-size: 8px; letter-spacing: 0.1em; }
.pdbo-sprig { position: absolute; width: 74px; height: 150px; border-left: 1px solid rgba(36, 53, 39, 0.4); }
.pdbo-sprig::before, .pdbo-sprig::after { content: ''; position: absolute; width: 30px; height: 16px; border-radius: 100% 0 100% 0; background: rgba(36, 53, 39, 0.45); }
.pdbo-sprig::before { left: 0; top: 30px; transform: rotate(-24deg); }
.pdbo-sprig::after { left: 0; top: 80px; transform: rotate(24deg) scaleX(-1); }
.pdbo-sprig-a { left: 12%; bottom: -22px; transform: rotate(-8deg); }
.pdbo-sprig-b { right: 12%; bottom: -30px; transform: rotate(172deg); opacity: 0.7; }
.pdbo-latin { position: absolute; left: 50%; bottom: 12px; transform: translateX(-50%); font-family: ui-monospace, monospace; font-size: 7px; letter-spacing: 0.26em; color: rgba(36, 53, 39, 0.5); white-space: nowrap; }
.pdbo-buy { border-radius: 999px; background: #243527; color: #f0f2e6; transition: background 0.25s ease, transform 0.16s ease; }
.pdbo-buy:hover { background: #4d7c45; }
.pdbo-buy:active { transform: scale(0.985); }
.pdbo-in { animation: pdboIn 0.8s cubic-bezier(0.22, 1, 0.36, 1) both; }
.pdbo-d1 { animation-delay: 0.07s; }
.pdbo-d2 { animation-delay: 0.14s; }
.pdbo-d3 { animation-delay: 0.21s; }
.pdbo-d4 { animation-delay: 0.28s; }
@keyframes pdboIn { from { opacity: 0; transform: translateY(16px); } to { opacity: 1; transform: none; } }
@keyframes pdboFloat { 0%, 100% { transform: translate(-50%, -50%); } 50% { transform: translate(-50%, -53%); } }
@media (prefers-reduced-motion: reduce) {
  .pdbo-in, .pdbo-vessel { animation: none !important; }
  .pdbo-buy { transition: none; }
}
</style>
