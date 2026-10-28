<template>
  <div class="relative min-h-screen overflow-hidden bg-[#12142a] text-[#e8e6d4]">
    <div class="px-stars pointer-events-none absolute inset-0"></div>
    <div class="pointer-events-none absolute inset-x-0 top-0 h-72 bg-[radial-gradient(60%_100%_at_50%_0%,rgba(250,204,21,0.12),transparent_70%)]"></div>

    <div class="relative z-10 mx-auto w-full max-w-md px-5 pb-14 pt-8">
      <div class="px-rise flex items-center justify-between font-mono text-[8px] uppercase tracking-[0.18em] text-[#7b7fa8]">
        <span>Player 1 · Ready</span>
        <span v-if="nfcId">ID {{ nfcId }}</span>
      </div>

      <header class="px-rise px-d1 mt-5">
        <div class="flex items-center justify-between gap-3">
          <div class="min-w-0">
            <p class="font-mono text-[8px] uppercase tracking-[0.3em] text-[#facc15]">Stage 01 · Kitchen</p>
            <h1 class="px-pixel mt-2 text-[28px] font-bold uppercase leading-[1.15] tracking-[0.02em] text-[#fff8db]">{{ payload?.shopName || 'PIXEL CANTEEN' }}</h1>
          </div>
          <div class="px-coin">
            <span class="px-coin-n">×{{ pad(total) }}</span>
            <span class="px-coin-l">DISH</span>
          </div>
        </div>
        <p v-if="payload?.tagline" class="mt-2 max-w-[19rem] text-[10px] leading-[1.7] text-[#a7a9c8]">{{ payload.tagline }}</p>
      </header>

      <div class="px-rise px-d2 mt-4 px-panel px-panel-y">
        <div class="flex items-center justify-between font-mono text-[7px] uppercase tracking-[0.14em] text-[#9a9dc0]">
          <span>Hunger</span>
          <span>100 / 100</span>
        </div>
        <div class="px-bar mt-1"><span class="px-bar-fill px-bar-food"></span></div>
        <div class="mt-2 flex items-center justify-between font-mono text-[7px] uppercase tracking-[0.14em] text-[#9a9dc0]">
          <span>Spice</span>
          <span>auto</span>
        </div>
        <div class="px-bar mt-1"><span class="px-bar-fill px-bar-spice"></span></div>
      </div>

      <div class="px-rise px-d2 mt-5 flex items-center gap-2 font-mono text-[8px] uppercase tracking-[0.2em] text-[#facc15]">
        <span class="px-cursor"></span>
        <span>Select your menu</span>
        <span class="px-dash flex-1"></span>
        <span class="text-[#7b7fa8]">{{ pad(groups.length) }} AREA</span>
      </div>

      <section v-if="groups.length" class="mt-2 space-y-3">
        <div v-for="(group, gi) in groups" :key="gi" class="px-rise" :style="{ animationDelay: 0.12 + gi * 0.07 + 's' }">
          <div class="px-panel">
            <div class="px-panel-head">
              <span class="px-tri">▶</span>
              <h2 class="font-mono text-[10px] font-bold uppercase tracking-[0.14em] text-[#fff8db]">{{ group.name }}</h2>
              <span class="ml-auto font-mono text-[7px] text-[#facc15]">LV.{{ pad(gi + 1) }}</span>
            </div>
            <ul class="px-panel-body">
              <li v-for="(item, i) in group.items" :key="i" class="px-row">
                <span class="px-bullet">·</span>
                <div class="min-w-0 flex-1">
                  <div class="flex flex-wrap items-center gap-x-1.5">
                    <span class="font-mono text-[11px] font-bold text-[#e8e6d4]">{{ item?.name }}</span>
                    <span v-if="signature(item)" class="px-tag px-tag-star">★</span>
                    <span v-if="veg(item)" class="px-tag px-tag-veg">V</span>
                    <span v-if="spicy(item)" class="px-tag px-tag-hot">H{{ spicy(item) }}</span>
                  </div>
                  <p v-if="item?.desc" class="mt-0.5 font-mono text-[8px] leading-[1.5] text-[#8f92b4]">{{ item.desc }}</p>
                </div>
                <span v-if="hasPrice(item)" class="px-price">{{ item.price }}<small>G</small></span>
                <span v-else class="font-mono text-[9px] text-[#5f6390]">--</span>
              </li>
            </ul>
          </div>
        </div>
      </section>

      <div v-if="!groups.length" class="px-rise mt-4 px-panel px-panel-body text-center">
        <Ghost class="mx-auto h-6 w-6 text-[#7b7fa8]" />
        <p class="mt-2 font-mono text-[9px] uppercase tracking-[0.14em] text-[#a7a9c8]">No menu data. Insert coin.</p>
      </div>

      <div class="px-rise px-d3 mt-4 px-panel px-panel-body">
        <p class="font-mono text-[7px] uppercase tracking-[0.16em] text-[#facc15]">&gt; SYSTEM MSG</p>
        <p class="mt-1 font-mono text-[9px] leading-[1.7] text-[#a7a9c8]">{{ payload?.notice || '本店菜单会随季节刷新，请以现场为准。' }}</p>
        <p v-if="payload?.hours || payload?.address" class="mt-2 border-t border-[#343a68] pt-2 font-mono text-[7.5px] leading-relaxed text-[#7b7fa8]">
          <span v-if="payload?.hours">{{ payload.hours }}</span>
          <span v-if="payload?.hours && payload?.address"> · </span>
          <span v-if="payload?.address">{{ payload.address }}</span>
        </p>
      </div>

      <a v-if="payload?.phone" :href="'tel:' + payload.phone" class="px-rise px-d3 px-call mt-4">
        <Phone class="h-3.5 w-3.5" />
        <span class="px-pixel text-[10px] uppercase tracking-[0.06em]">Call &amp; order</span>
        <span class="ml-auto font-mono text-[9px]">{{ payload.phone }}</span>
      </a>

      <p class="px-blink mt-6 text-center font-mono text-[8px] uppercase tracking-[0.2em] text-[#facc15]">Press start · 按此点单</p>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { Phone, Ghost } from 'lucide-vue-next'

const props = defineProps(['payload', 'nfcId'])

const pad = (n) => String(n).padStart(2, '0')
const hasPrice = (item) => item?.price !== undefined && item?.price !== null && item?.price !== ''
const signature = (item) => item?.signature === true || (Array.isArray(item?.tags) && item.tags.some((t) => /招牌|必点|signature/i.test(String(t))))
const veg = (item) => item?.vegetarian === true || (Array.isArray(item?.tags) && item.tags.some((t) => /素|vegan|vegetarian/i.test(String(t))))
const spicy = (item) => { const n = Number(item?.spicy); return Number.isFinite(n) ? Math.max(0, Math.min(3, Math.round(n))) : 0 }

const groups = computed(() => {
  const p = props.payload || {}
  const list = Array.isArray(p.categories) ? p.categories : []
  return list.map((c) => ({
    name: (c && c.name) || '未命名分类',
    items: (Array.isArray(c?.items) ? c.items : []).filter((i) => i && (i.name || i.price)),
  })).filter((g) => g.items.length)
})
const total = computed(() => groups.value.reduce((n, g) => n + g.items.length, 0))
</script>

<style scoped>
.px-stars {
  opacity: 0.7;
  background-image: radial-gradient(rgba(255, 248, 219, 0.5) 1px, transparent 1px), radial-gradient(rgba(250, 204, 21, 0.3) 1px, transparent 1px);
  background-size: 37px 37px, 61px 61px;
  background-position: 0 0, 12px 19px;
  mask-image: linear-gradient(180deg, #000, transparent 55%);
}
.px-pixel { font-family: ui-monospace, 'Courier New', monospace; text-shadow: 2px 2px 0 #343a68; }
.px-coin { flex-shrink: 0; border: 2px solid #facc15; background: #1b1e3c; padding: 4px 7px; text-align: center; box-shadow: 3px 3px 0 #0a0b18; }
.px-coin-n { display: block; font-family: ui-monospace, monospace; font-size: 15px; font-weight: 700; line-height: 1; color: #facc15; }
.px-coin-l { display: block; margin-top: 2px; font-family: ui-monospace, monospace; font-size: 5.5px; letter-spacing: 0.16em; color: #b7a44a; }
.px-panel { border: 2px solid #8b8fbd; background: #1b1e3c; box-shadow: 4px 4px 0 #0a0b18; }
.px-panel-y { border-color: #facc15; }
.px-panel-head { display: flex; align-items: center; gap: 6px; border-bottom: 2px solid #6e72a3; background: #2a2f5c; padding: 6px 8px; }
.px-panel-body { padding: 7px 9px 8px; }
.px-tri { font-size: 7px; color: #facc15; }
.px-bar { height: 8px; border: 1px solid #6e72a3; background: #0f1124; padding: 1px; }
.px-bar-fill { display: block; height: 100%; }
.px-bar-food { width: 100%; background: repeating-linear-gradient(90deg, #8be26a 0 7px, #6bc94c 7px 9px); }
.px-bar-spice { width: 62%; background: repeating-linear-gradient(90deg, #f8714a 0 7px, #d9502e 7px 9px); }
.px-cursor { height: 7px; width: 7px; background: #facc15; animation: px-blink 1.1s steps(2, end) infinite; }
.px-dash { height: 2px; background: repeating-linear-gradient(90deg, #6e72a3 0 4px, transparent 4px 8px); }
.px-row { display: flex; align-items: baseline; gap: 6px; border-bottom: 1px dashed #343a68; padding: 7px 0; }
.px-row:last-child { border-bottom: none; }
.px-bullet { color: #facc15; font-size: 10px; }
.px-tag { border: 1px solid #6e72a3; padding: 0 2px; font-family: ui-monospace, monospace; font-size: 5.5px; line-height: 1.4; }
.px-tag-star { border-color: #facc15; color: #facc15; }
.px-tag-veg { border-color: #8be26a; color: #8be26a; }
.px-tag-hot { border-color: #f8714a; color: #f8714a; }
.px-price { flex-shrink: 0; border: 1px solid #facc15; background: #24284a; padding: 2px 4px; font-family: ui-monospace, monospace; font-size: 10px; font-weight: 700; color: #facc15; }
.px-price small { margin-left: 1px; font-size: 5.5px; opacity: 0.7; }
.px-call {
  display: flex;
  align-items: center;
  gap: 8px;
  border: 2px solid #8be26a;
  background: #1b1e3c;
  padding: 9px 11px;
  color: #c7f5b5;
  box-shadow: 4px 4px 0 #0a0b18;
  transition: background-color 0.15s ease, transform 0.15s ease;
}
.px-call:hover { background: #26304e; }
.px-call:active { transform: translate(2px, 2px); box-shadow: 1px 1px 0 #0a0b18; }
.px-blink { animation: px-blink 1.4s steps(2, end) infinite; }
@keyframes px-blink { 0%, 45% { opacity: 1; } 50%, 100% { opacity: 0.25; } }
.px-rise { animation: px-in 0.5s steps(5, end) both; }
.px-d1 { animation-delay: 0.06s; }
.px-d2 { animation-delay: 0.12s; }
.px-d3 { animation-delay: 0.26s; }
@keyframes px-in {
  from { opacity: 0; transform: translateY(8px); }
  to { opacity: 1; transform: none; }
}
@media (prefers-reduced-motion: reduce) {
  .px-rise, .px-cursor, .px-blink, .px-call { animation: none !important; transition: none !important; }
}
</style>
