<template>
  <div class="relative min-h-screen overflow-hidden bg-[#f3f1e3] text-[#2f3928]">
    <div class="hg-paper pointer-events-none absolute inset-0"></div>
    <div class="pointer-events-none absolute -left-12 top-40 h-48 w-48 rounded-full bg-[#dce4c8]/50 blur-2xl"></div>

    <div class="relative z-10 mx-auto w-full max-w-md px-6 pb-14 pt-9">
      <div class="hg-rise flex items-center justify-between font-mono text-[7.5px] uppercase tracking-[0.24em] text-[#8a9577]">
        <span>Herbarium · No. 25</span>
        <span v-if="nfcId">NFC {{ nfcId }}</span>
      </div>

      <header class="hg-rise hg-d1 mt-5 text-center">
        <svg class="mx-auto h-12 w-16 text-[#6f8557]" viewBox="0 0 64 48" fill="none" aria-hidden="true">
          <path d="M32 45V13" stroke="currentColor" stroke-width="1" />
          <path d="M32 30c-8-1-13-5-14-12 8 1 13 5 14 12ZM32 32c8-1 13-5 14-12-8 1-13 5-14 12ZM32 20c-6-1-10-4-11-9 6 1 10 4 11 9ZM32 22c6-1 10-4 11-9-6 1-10 4-11 9Z" fill="currentColor" opacity="0.35" />
          <circle cx="32" cy="8" r="3" stroke="currentColor" stroke-width="1" />
        </svg>
        <p class="hg-serif mt-1 text-[11px] italic tracking-[0.14em] text-[#8a9577]">A menu gathered from the garden</p>
        <h1 class="hg-serif mt-1 text-[41px] font-medium leading-[1.03] tracking-[-0.015em] text-[#2f3928]">{{ payload?.shopName || '香草庭园' }}</h1>
        <p v-if="payload?.tagline" class="hg-serif mx-auto mt-2 max-w-[18rem] text-[12px] italic leading-[1.7] text-[#67725a]">{{ payload.tagline }}</p>
        <div class="mt-4 flex items-center justify-center gap-2 font-mono text-[7.5px] uppercase tracking-[0.14em] text-[#8a9577]">
          <span v-if="payload?.hours" class="inline-flex items-center gap-1"><Clock class="h-2.5 w-2.5" />{{ payload.hours }}</span>
          <span v-if="payload?.hours" class="h-3 w-px bg-[#c1cab0]"></span>
          <span>Seasonal · 2026</span>
        </div>
      </header>

      <div class="hg-rise hg-d2 mt-6 flex items-center gap-3">
        <span class="hg-rule"></span>
        <span class="hg-serif text-[10px] italic tracking-[0.24em] text-[#6f8557]">the garden beds</span>
        <span class="hg-rule"></span>
      </div>

      <section v-if="groups.length" class="mt-3 space-y-4">
        <article v-for="(group, gi) in groups" :key="gi" class="hg-rise hg-card" :style="{ animationDelay: 0.12 + gi * 0.07 + 's' }">
          <div class="hg-illust">
            <svg viewBox="0 0 44 100" fill="none" aria-hidden="true" class="h-full w-full">
              <path d="M22 96V20" stroke="#6f8557" stroke-width="0.8" />
              <path d="M22 40c-9-2-15-7-16-16 9 2 15 7 16 16ZM22 44c9-2 15-7 16-16-9 2-15 7-16 16ZM22 62c-7-2-12-6-13-13 7 2 12 6 13 13ZM22 66c7-2 12-6 13-13-7 2-12 6-13 13ZM22 24c-5-2-8-5-9-10 5 2 8 5 9 10Z" fill="#8fa578" opacity="0.45" />
              <circle cx="22" cy="13" r="4" stroke="#6f8557" stroke-width="0.8" />
            </svg>
            <span class="hg-num">{{ pad(gi + 1) }}</span>
          </div>

          <div class="min-w-0 flex-1 py-3 pr-3.5">
            <div class="flex items-baseline justify-between gap-2 border-b border-[#c1cab0] pb-1">
              <h2 class="hg-serif text-[17px] font-semibold tracking-[0.04em] text-[#37432e]">{{ group.name }}</h2>
              <span class="font-mono text-[7px] uppercase tracking-[0.12em] text-[#8a9577]">{{ pad(group.items.length) }} sp.</span>
            </div>
            <ul class="mt-1">
              <li v-for="(item, i) in group.items" :key="i" class="hg-row">
                <span class="hg-leaf">✦</span>
                <div class="min-w-0 flex-1">
                  <div class="flex flex-wrap items-baseline gap-x-1.5">
                    <span class="hg-serif text-[13px] font-medium text-[#2f3928]">{{ item?.name }}</span>
                    <span v-if="signature(item)" class="hg-tag">园主推荐</span>
                    <span v-if="veg(item)" class="hg-tag hg-tag-g">蔬食</span>
                    <span v-if="spicy(item)" class="inline-flex text-[#a35a3a]"><Flame v-for="n in spicy(item)" :key="n" class="h-2.5 w-2.5" /></span>
                  </div>
                  <p v-if="item?.desc" class="mt-0.5 text-[9px] leading-[1.5] text-[#7d8770]">{{ item.desc }}</p>
                </div>
                <span v-if="hasPrice(item)" class="hg-price">{{ item.price }}</span>
                <span v-else class="font-mono text-[9px] text-[#a8b098]">—</span>
              </li>
            </ul>
          </div>
        </article>
      </section>

      <div v-if="!groups.length" class="hg-rise mt-5 border border-[#c1cab0] bg-[#faf9ef] px-6 py-12 text-center">
        <Sprout class="mx-auto h-6 w-6 text-[#8a9577]" />
        <p class="hg-serif mt-3 text-[12px] italic text-[#7d8770]">The beds are resting. Menu returns with the season.</p>
      </div>

      <div v-if="payload?.notice" class="hg-rise mt-5 border-l border-[#8fa578] pl-3">
        <p class="hg-serif text-[12px] italic leading-[1.7] text-[#67725a]">{{ payload.notice }}</p>
      </div>

      <div class="hg-rise hg-d3 mt-7 flex items-end justify-between gap-4 border-t border-[#c1cab0] pt-4">
        <div class="min-w-0">
          <p class="font-mono text-[6.5px] uppercase tracking-[0.22em] text-[#8a9577]">Reserve a table</p>
          <a v-if="payload?.phone" :href="'tel:' + payload.phone" class="hg-serif mt-1 inline-flex items-center gap-2 text-[16px] text-[#37432e] hover:text-[#6f8557]">
            <Phone class="h-3.5 w-3.5" />{{ payload.phone }}
          </a>
        </div>
        <div v-if="payload?.address" class="max-w-[10rem] text-right">
          <p class="font-mono text-[6px] uppercase tracking-[0.18em] text-[#8a9577]">Find the garden</p>
          <p class="mt-1 flex items-start justify-end gap-1 text-[8.5px] leading-[1.5] text-[#7d8770]"><MapPin class="mt-0.5 h-2.5 w-2.5 flex-shrink-0" />{{ payload.address }}</p>
        </div>
      </div>

      <p class="hg-serif mt-7 text-center text-[11px] italic tracking-[0.18em] text-[#a8b098]">— picked this morning —</p>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { Clock, MapPin, Phone, Flame, Sprout } from 'lucide-vue-next'

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
</script>

<style scoped>
.hg-paper {
  opacity: 0.55;
  background-image: repeating-linear-gradient(4deg, rgba(90, 100, 70, 0.04) 0 1px, transparent 1px 6px),
    radial-gradient(rgba(90, 100, 70, 0.09) 0.6px, transparent 0.6px);
  background-size: auto, 9px 9px;
}
.hg-serif { font-family: 'Cormorant Garamond', 'Iowan Old Style', 'Palatino Linotype', 'Songti SC', Georgia, serif; }
.hg-rule { flex: 1; height: 1px; background: linear-gradient(90deg, transparent, #b6c0a4, transparent); }
.hg-card {
  display: flex;
  overflow: hidden;
  border: 1px solid #c1cab0;
  border-radius: 3px;
  background: #faf9ee;
  box-shadow: 0 14px 30px -28px rgba(47, 57, 40, 0.5);
  transition: box-shadow 0.2s ease, transform 0.2s ease;
}
.hg-card:hover { transform: translateY(-2px); box-shadow: 0 18px 30px -24px rgba(47, 57, 40, 0.55); }
.hg-illust { position: relative; width: 48px; flex-shrink: 0; border-right: 1px dotted #c1cab0; background: #f0f2e2; padding: 8px 4px 4px; }
.hg-num { position: absolute; left: 50%; bottom: 3px; transform: translateX(-50%); font-family: ui-monospace, monospace; font-size: 6px; letter-spacing: 0.1em; color: #8a9577; }
.hg-row { display: flex; align-items: baseline; gap: 6px; padding: 7px 0; border-bottom: 1px dotted rgba(193, 202, 176, 0.8); }
.hg-row:last-child { border-bottom: none; }
.hg-leaf { font-size: 6px; color: #8fa578; }
.hg-tag { border: 1px solid rgba(111, 133, 87, 0.5); padding: 0 2px; font-size: 6px; letter-spacing: 0.05em; color: #6f8557; }
.hg-tag-g { border-color: rgba(140, 160, 110, 0.5); color: #7d925c; }
.hg-price { flex-shrink: 0; font-family: 'Cormorant Garamond', 'Songti SC', serif; font-size: 14px; color: #6f8557; }
.hg-price::before { content: '¥'; margin-right: 1px; font-size: 8px; opacity: 0.7; }
.hg-rise { animation: hg-in 0.75s cubic-bezier(0.2, 0.9, 0.25, 1) both; }
.hg-d1 { animation-delay: 0.08s; }
.hg-d2 { animation-delay: 0.16s; }
.hg-d3 { animation-delay: 0.34s; }
@keyframes hg-in {
  from { opacity: 0; transform: translateY(12px); }
  to { opacity: 1; transform: none; }
}
@media (prefers-reduced-motion: reduce) {
  .hg-rise, .hg-card { animation: none !important; transition: none !important; }
  .hg-card:hover { transform: none; }
}
</style>
