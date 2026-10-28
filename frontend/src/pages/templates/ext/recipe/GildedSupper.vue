<template>
  <div class="relative min-h-screen overflow-hidden bg-[#08080a] text-[#e8d6a8]">
    <div class="gs-glow pointer-events-none absolute inset-x-0 top-0 h-96"></div>
    <div class="gs-frame pointer-events-none absolute inset-3 border border-[#b08d45]/35"></div>
    <div class="gs-frame2 pointer-events-none absolute inset-5 border border-[#b08d45]/20"></div>

    <div class="relative z-10 mx-auto w-full max-w-md px-7 pb-14 pt-9">
      <div class="gs-rise flex items-center justify-between font-mono text-[7px] uppercase tracking-[0.24em] text-[#8d7b56]">
        <span>Maison · 1987</span>
        <span v-if="nfcId">NFC {{ nfcId }}</span>
      </div>

      <header class="gs-rise gs-d1 mt-4 text-center">
        <svg class="mx-auto h-8 w-24 text-[#c9a85f]" viewBox="0 0 96 32" fill="none" aria-hidden="true">
          <path d="M48 27 44 17h8l-4 10Z" fill="currentColor" opacity="0.85" />
          <path d="M48 24 38 8M48 24 58 8M48 24 30 13M48 24 66 13M48 24 24 20M48 24 72 20" stroke="currentColor" stroke-width="0.9" />
          <path d="M14 27h68" stroke="currentColor" stroke-width="0.8" />
        </svg>
        <p class="mt-2 font-mono text-[7.5px] uppercase tracking-[0.5em] text-[#a58a4f]">Fine Dining</p>
        <h1 class="gs-serif mt-2 text-[40px] font-medium leading-[1.06] tracking-[0.02em] text-[#f2e3bd]">{{ payload?.shopName || '金箔晚宴' }}</h1>
        <div class="mt-2 flex items-center justify-center gap-3">
          <span class="h-px w-10 bg-[#b08d45]"></span>
          <span class="h-1 w-1 rotate-45 bg-[#c9a85f]"></span>
          <span class="h-px w-10 bg-[#b08d45]"></span>
        </div>
        <p v-if="payload?.tagline" class="mx-auto mt-3 max-w-[17rem] text-[10px] leading-[1.8] text-[#b5a077]">{{ payload.tagline }}</p>
      </header>

      <section v-if="groups.length" class="mt-7">
        <div v-for="(group, gi) in groups" :key="gi" class="gs-rise gs-group" :style="{ animationDelay: 0.14 + gi * 0.08 + 's' }">
          <div class="flex items-center justify-center gap-3">
            <span class="h-px w-6 bg-[#b08d45]/60"></span>
            <h2 class="gs-serif text-[13px] uppercase tracking-[0.3em] text-[#d8bd80]">{{ group.name }}</h2>
            <span class="h-px w-6 bg-[#b08d45]/60"></span>
          </div>
          <ul class="mt-2">
            <li v-for="(item, i) in group.items" :key="i" class="gs-item">
              <div class="flex items-baseline justify-center gap-2">
                <span class="gs-serif text-[15px] tracking-[0.03em] text-[#f0e2c0]">{{ item?.name }}</span>
                <span v-if="signature(item)" class="gs-star">✦</span>
                <span v-if="veg(item)" class="gs-v">V</span>
              </div>
              <p v-if="item?.desc" class="mx-auto mt-0.5 max-w-[17rem] text-center text-[9px] leading-[1.6] text-[#9f8d6a]">{{ item.desc }}</p>
              <div class="mt-0.5 flex items-center justify-center gap-1.5">
                <span v-if="hasPrice(item)" class="font-mono text-[10px] tabular-nums tracking-[0.08em] text-[#c9a85f]">¥ {{ item.price }}</span>
                <span v-else class="font-mono text-[9px] text-[#6f644e]">—</span>
                <Flame v-for="n in spicy(item)" :key="n" class="h-2.5 w-2.5 text-[#b5713c]" />
              </div>
            </li>
          </ul>
        </div>
      </section>

      <div v-else class="gs-rise mt-7 border border-[#b08d45]/40 px-6 py-12 text-center">
        <UtensilsCrossed class="mx-auto h-5 w-5 text-[#8d7b56]" />
        <p class="gs-serif mt-3 text-[12px] italic text-[#b5a077]">The menu is being set.</p>
      </div>

      <div v-if="payload?.notice" class="gs-rise mt-6 text-center">
        <p class="gs-serif text-[9.5px] italic leading-[1.7] text-[#9f8d6a]">{{ payload.notice }}</p>
      </div>

      <div class="gs-rise gs-d3 mt-7 border-t border-[#b08d45]/35 pt-4 text-center">
        <p class="font-mono text-[6.5px] uppercase tracking-[0.32em] text-[#8d7b56]">Reservations</p>
        <a v-if="payload?.phone" :href="'tel:' + payload.phone" class="gs-serif mt-1 inline-block text-[17px] tracking-[0.08em] text-[#e8d6a8] hover:text-[#c9a85f]">{{ payload.phone }}</a>
        <div class="mt-3 flex flex-wrap items-center justify-center gap-x-3 gap-y-1 font-mono text-[7px] tracking-[0.1em] text-[#8d7b56]">
          <span v-if="payload?.hours">{{ payload.hours }}</span>
          <span v-if="payload?.hours && payload?.address" class="h-3 w-px bg-[#b08d45]/40"></span>
          <span v-if="payload?.address">{{ payload.address }}</span>
        </div>
      </div>

      <p class="mt-6 text-center font-mono text-[6px] uppercase tracking-[0.5em] text-[#6f644e]">Gold · Fire · Time</p>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { Flame, UtensilsCrossed } from 'lucide-vue-next'

const props = defineProps(['payload', 'nfcId'])

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
.gs-glow { background: radial-gradient(55% 70% at 50% 0%, rgba(201, 168, 95, 0.18), transparent 70%); }
.gs-frame { border-radius: 2px; pointer-events: none; }
.gs-frame2 { pointer-events: none; }
.gs-serif { font-family: 'Didot', 'Bodoni MT', 'Songti SC', 'STSong', Georgia, serif; }
.gs-group { padding-top: 24px; }
.gs-item {
  padding: 9px 0;
  border-bottom: 1px solid rgba(176, 141, 69, 0.14);
}
.gs-item:last-child { border-bottom: none; }
.gs-star { font-size: 9px; color: #c9a85f; }
.gs-v { border: 1px solid rgba(150, 170, 120, 0.45); padding: 0 3px; font-family: ui-monospace, monospace; font-size: 6px; color: #b3c49a; }
.gs-rise { animation: gs-in 0.9s cubic-bezier(0.2, 0.9, 0.25, 1) both; }
.gs-d1 { animation-delay: 0.1s; }
.gs-d3 { animation-delay: 0.34s; }
@keyframes gs-in {
  from { opacity: 0; transform: translateY(14px); }
  to { opacity: 1; transform: none; }
}
@media (prefers-reduced-motion: reduce) {
  .gs-rise { animation: none !important; }
  a { transition: none !important; }
}
</style>
