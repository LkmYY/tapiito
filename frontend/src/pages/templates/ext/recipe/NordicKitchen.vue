<template>
  <div class="relative min-h-screen overflow-hidden bg-[#eef0e6] text-[#2a3329]">
    <div class="pointer-events-none absolute -right-20 top-28 h-64 w-64 rounded-full bg-[#c7d2b8]/50 blur-3xl"></div>
    <div class="pointer-events-none absolute -left-16 bottom-20 h-52 w-52 rounded-full bg-[#e8dcc6]/60 blur-3xl"></div>

    <div class="relative z-10 mx-auto w-full max-w-md px-6 pb-14 pt-9">
      <div class="nk-rise flex items-center justify-between font-mono text-[8px] uppercase tracking-[0.26em] text-[#7e8a78]">
        <span>Nordic Kitchen</span>
        <span v-if="nfcId">NFC · {{ nfcId }}</span>
      </div>

      <header class="nk-rise nk-d1 mt-7">
        <span class="nk-line"></span>
        <h1 class="nk-serif mt-4 text-[43px] font-normal leading-[1.02] tracking-[-0.03em] text-[#283226]">
          {{ payload?.shopName || '北岸厨房' }}
        </h1>
        <p v-if="payload?.tagline" class="mt-3 max-w-[19rem] text-[12px] leading-[1.8] text-[#65715f]">{{ payload.tagline }}</p>

        <div class="mt-6 flex flex-wrap items-center gap-x-4 gap-y-2">
          <span v-if="payload?.hours" class="nk-pill"><Clock class="h-3 w-3" />{{ payload.hours }}</span>
          <span v-if="payload?.address" class="nk-address"><MapPin class="h-3 w-3" />{{ payload.address }}</span>
        </div>
      </header>

      <div class="nk-rise nk-d2 mt-9 flex items-baseline justify-between border-b border-[#a9b5a0] pb-2">
        <h2 class="font-mono text-[9px] uppercase tracking-[0.3em] text-[#5f6d59]">Today's Table</h2>
        <span class="nk-serif text-[11px] italic text-[#7e8a78]">{{ total }} dishes</span>
      </div>

      <section v-if="groups.length" class="mt-2">
        <div v-for="(group, gi) in groups" :key="gi" class="nk-rise nk-group" :style="{ animationDelay: 0.14 + gi * 0.07 + 's' }">
          <div class="flex items-center gap-2.5">
            <h3 class="nk-serif text-[20px] text-[#37452f]">{{ group.name }}</h3>
            <span class="h-px flex-1 bg-[#d3d9c8]"></span>
            <span class="font-mono text-[8px] tracking-[0.1em] text-[#93a08a]">{{ pad(group.items.length) }}</span>
          </div>
          <ul>
            <li v-for="(item, i) in group.items" :key="i" class="nk-item">
              <div class="min-w-0 flex-1">
                <div class="flex flex-wrap items-center gap-x-2 gap-y-1">
                  <span class="text-[14.5px] font-medium tracking-[0.01em] text-[#2e3a2b]">{{ item?.name }}</span>
                  <span v-if="signature(item)" class="nk-star">★</span>
                  <span v-if="veg(item)" class="nk-leaf"><Leaf class="h-2.5 w-2.5" />植</span>
                  <span v-if="spicy(item)" class="inline-flex text-[#a15a3d]"><Flame v-for="n in spicy(item)" :key="n" class="h-2.5 w-2.5" /></span>
                </div>
                <p v-if="item?.desc" class="mt-1 max-w-[17rem] text-[10.5px] leading-[1.65] text-[#7e8a78]">{{ item.desc }}</p>
              </div>
              <span v-if="hasPrice(item)" class="nk-price">{{ item.price }}</span>
              <span v-else class="font-mono text-[9px] text-[#a9b5a0]">—</span>
            </li>
          </ul>
        </div>
      </section>

      <div v-else class="nk-rise mt-6 rounded-[24px] bg-[#f7f8f2] px-6 py-12 text-center shadow-[0_20px_50px_-40px_rgba(42,51,41,0.4)]">
        <Sprout class="mx-auto h-6 w-6 text-[#93a08a]" />
        <p class="mt-3 text-[12px] text-[#7e8a78]">菜单正在生长，稍后再来。</p>
      </div>

      <div v-if="payload?.notice" class="nk-rise mt-7 rounded-2xl bg-[#e3e7d9] px-4 py-3.5">
        <p class="text-[10.5px] leading-[1.7] text-[#65715f]">{{ payload.notice }}</p>
      </div>

      <div class="nk-rise nk-d3 mt-8">
        <a v-if="payload?.phone" :href="'tel:' + payload.phone" class="nk-call">
          <span class="nk-call-label">Call</span>
          <span class="nk-serif text-[16px] text-[#283226]">{{ payload.phone }}</span>
          <ArrowRight class="ml-auto h-4 w-4" />
        </a>
      </div>

      <div class="mt-8 flex items-center justify-between border-t border-[#d3d9c8] pt-3 font-mono text-[7px] uppercase tracking-[0.22em] text-[#9daa94]">
        <span>Seasonal · Local · Simple</span>
        <span>Malmö / 2026</span>
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { Clock, MapPin, Flame, Leaf, Sprout, ArrowRight } from 'lucide-vue-next'

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
.nk-serif { font-family: 'Iowan Old Style', 'Palatino Linotype', 'Songti SC', Georgia, serif; }
.nk-line { display: block; height: 2px; width: 34px; background: #6f825f; }
.nk-pill {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  border: 1px solid #c4cfb7;
  border-radius: 9999px;
  background: #f6f8f0;
  padding: 5px 11px;
  font-family: ui-monospace, monospace;
  font-size: 8.5px;
  letter-spacing: 0.06em;
  color: #61705a;
}
.nk-address { display: inline-flex; align-items: center; gap: 5px; max-width: 13rem; font-size: 9.5px; line-height: 1.4; color: #7e8a78; }
.nk-group { padding-top: 24px; }
.nk-item {
  display: flex;
  align-items: flex-start;
  gap: 12px;
  padding: 13px 0;
  border-bottom: 1px solid #dfe4d5;
}
.nk-item:last-child { border-bottom: none; }
.nk-star { font-size: 9px; color: #b98b3a; }
.nk-leaf { display: inline-flex; align-items: center; gap: 2px; border-radius: 9999px; background: #e2ead7; padding: 1px 5px; font-size: 7.5px; color: #5f7350; }
.nk-price { flex-shrink: 0; border-radius: 9999px; background: #f7f8f2; padding: 4px 10px; font-family: ui-monospace, monospace; font-size: 12px; color: #4c5c45; box-shadow: inset 0 0 0 1px #d3d9c8; }
.nk-price::before { content: '¥'; margin-right: 1px; font-size: 8px; opacity: 0.7; }
.nk-call {
  display: flex;
  align-items: center;
  gap: 12px;
  border-radius: 9999px;
  background: #f7f8f2;
  padding: 12px 18px;
  box-shadow: 0 14px 30px -24px rgba(42, 51, 41, 0.6);
  transition: background-color 0.2s ease, transform 0.15s ease;
}
.nk-call:hover { background: #fbfcf7; }
.nk-call:active { transform: scale(0.985); }
.nk-call-label { font-family: ui-monospace, monospace; font-size: 7.5px; text-transform: uppercase; letter-spacing: 0.2em; color: #93a08a; }
.nk-rise { animation: nk-in 0.75s cubic-bezier(0.2, 0.9, 0.25, 1) both; }
.nk-d1 { animation-delay: 0.1s; }
.nk-d2 { animation-delay: 0.18s; }
.nk-d3 { animation-delay: 0.36s; }
@keyframes nk-in {
  from { opacity: 0; transform: translateY(14px); }
  to { opacity: 1; transform: none; }
}
@media (prefers-reduced-motion: reduce) {
  .nk-rise, .nk-call { animation: none !important; transition: none !important; }
}
</style>
