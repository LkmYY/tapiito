<template>
  <div class="relative min-h-screen overflow-hidden bg-[#0d0d0d] text-[#e6e0d2]">
    <div class="fc-grain pointer-events-none absolute inset-0"></div>
    <div class="pointer-events-none absolute inset-y-0 left-0 w-4 fc-sprockets"></div>
    <div class="pointer-events-none absolute inset-y-0 right-0 w-4 fc-sprockets"></div>

    <div class="relative z-10 mx-auto w-full max-w-md px-7 pb-12 pt-7">
      <div class="fc-rise flex items-center justify-between font-mono text-[7.5px] uppercase tracking-[0.24em] text-[#8f8a7c]">
        <span>Film 400 · 35mm</span>
        <span v-if="nfcId">Frame {{ nfcId }}</span>
      </div>

      <header class="fc-rise fc-d1 mt-4 border-y border-[#3a3832] py-3">
        <div class="flex items-end justify-between">
          <div class="min-w-0 flex-1">
            <p class="font-mono text-[7px] uppercase tracking-[0.28em] text-[#c7a86a]">Contact Sheet · Kitchen Roll</p>
            <h1 class="mt-1 truncate text-[27px] font-semibold uppercase tracking-[-0.01em] text-[#f1ebdd]">{{ payload?.shopName }}</h1>
          </div>
          <div class="text-right font-mono text-[6.5px] uppercase leading-tight tracking-[0.1em] text-[#8f8a7c]">
            <p>EXP. 2026</p>
            <p>{{ pad(total) }} EXP</p>
            <p>ISO 400</p>
          </div>
        </div>
        <p v-if="payload?.tagline" class="mt-1.5 max-w-[20rem] text-[9.5px] leading-[1.6] text-[#a39d8e]">{{ payload.tagline }}</p>
        <div class="mt-2 flex items-center gap-3 font-mono text-[7px] text-[#8f8a7c]">
          <span v-if="payload?.hours">{{ payload.hours }}</span>
          <span v-if="payload?.address" class="truncate">{{ payload.address }}</span>
        </div>
      </header>

      <section v-if="groups.length" class="mt-4 space-y-4">
        <div v-for="(group, gi) in groups" :key="gi" class="fc-rise" :style="{ animationDelay: 0.12 + gi * 0.07 + 's' }">
          <div class="flex items-center gap-2">
            <span class="fc-roll">ROLL {{ pad(gi + 1) }}</span>
            <span class="h-px flex-1 bg-[#3a3832]"></span>
            <h2 class="font-mono text-[8px] uppercase tracking-[0.18em] text-[#c7a86a]">{{ group.name }}</h2>
          </div>

          <div class="mt-1.5 grid grid-cols-2 gap-x-1.5 gap-y-2">
            <article v-for="(item, i) in group.items" :key="i" class="fc-frame">
              <div class="fc-photo" :class="'fc-tone-' + ((i + gi) % 3)">
                <svg viewBox="0 0 100 64" fill="none" aria-hidden="true">
                  <ellipse cx="50" cy="41" rx="28" ry="8" fill="rgba(255,255,255,0.15)" stroke="rgba(255,255,255,0.5)" stroke-width="0.7" />
                  <path d="M34 37c0-7 7-12 16-12s16 5 16 12" fill="rgba(255,255,255,0.1)" stroke="rgba(255,255,255,0.5)" stroke-width="0.7" />
                  <path d="M42 27c0-4 3-4 3-7M51 26c0-4 3-4 3-7" stroke="rgba(255,255,255,0.45)" stroke-width="0.7" stroke-linecap="round" />
                  <rect x="4" y="4" width="92" height="56" stroke="rgba(255,255,255,0.16)" stroke-width="0.5" />
                </svg>
                <span class="fc-no">{{ pad(i + 1) }}A</span>
                <span v-if="signature(item)" class="fc-star">★</span>
              </div>
              <div class="fc-cap">
                <div class="flex items-baseline justify-between gap-1">
                  <span class="truncate text-[10.5px] font-medium text-[#f1ebdd]">{{ item?.name }}</span>
                  <span v-if="hasPrice(item)" class="flex-shrink-0 font-mono text-[9px] text-[#c7a86a]">¥{{ item.price }}</span>
                </div>
                <p v-if="item?.desc" class="mt-0.5 line-clamp-2 min-h-[22px] text-[8px] leading-[1.35] text-[#a39d8e]">{{ item.desc }}</p>
                <div class="mt-1 flex items-center gap-1 font-mono text-[6px] uppercase tracking-[0.1em] text-[#7d786c]">
                  <span v-if="veg(item)">VEG</span>
                  <span v-if="spicy(item)">SPICY-{{ spicy(item) }}</span>
                  <span v-if="!veg(item) && !spicy(item) && !signature(item)">—</span>
                </div>
              </div>
            </article>
          </div>
        </div>
      </section>

      <div v-if="!groups.length" class="fc-rise mt-4 border border-dashed border-[#3a3832] px-6 py-12 text-center">
        <Aperture class="mx-auto h-6 w-6 text-[#8f8a7c]" />
        <p class="mt-3 font-mono text-[8px] uppercase tracking-[0.2em] text-[#a39d8e]">Roll not developed</p>
      </div>

      <div v-if="payload?.notice" class="fc-rise mt-4 border-l border-[#c7a86a]/50 pl-3">
        <p class="font-mono text-[8px] leading-[1.7] text-[#a39d8e]">{{ payload.notice }}</p>
      </div>

      <div class="fc-rise fc-d3 mt-5 flex items-center justify-between gap-3 border-t border-[#3a3832] pt-3">
        <a v-if="payload?.phone" :href="'tel:' + payload.phone" class="fc-call">
          <Phone class="h-3.5 w-3.5" />
          <span class="font-mono text-[8px] uppercase tracking-[0.16em]">Develop &amp; order</span>
          <span class="font-mono text-[9px] text-[#c7a86a]">{{ payload.phone }}</span>
        </a>
        <span v-else class="font-mono text-[8px] text-[#7d786c]">—</span>
        <span class="text-right font-mono text-[6px] uppercase leading-tight tracking-[0.12em] text-[#7d786c]">Hand printed<br />in the kitchen</span>
      </div>

      <p class="mt-6 text-center font-mono text-[6px] uppercase tracking-[0.36em] text-[#5f5b52]">Do not expose to direct sunlight</p>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { Phone, Aperture } from 'lucide-vue-next'

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
.fc-grain {
  opacity: 0.5;
  background-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' width='150' height='150'%3E%3Cfilter id='n'%3E%3CfeTurbulence type='fractalNoise' baseFrequency='0.9' numOctaves='4'/%3E%3C/filter%3E%3Crect width='150' height='150' filter='url(%23n)' opacity='0.09'/%3E%3C/svg%3E");
}
.fc-sprockets {
  background: repeating-linear-gradient(180deg, transparent 0 9px, #d8d2c4 9px 21px, transparent 21px 30px);
  opacity: 0.35;
}
.fc-roll { border: 1px solid #c7a86a; padding: 1px 4px; font-family: ui-monospace, monospace; font-size: 6px; letter-spacing: 0.12em; color: #c7a86a; }
.fc-frame { border: 1px solid #2e2c27; background: #151515; padding: 3px; }
.fc-photo { position: relative; height: 78px; overflow: hidden; border: 1px solid #2e2c27; }
.fc-photo svg { position: absolute; inset: 0; height: 100%; width: 100%; }
.fc-tone-0 { background: linear-gradient(140deg, #3a4034, #171b16); }
.fc-tone-1 { background: linear-gradient(140deg, #4a3a2c, #1b1510); }
.fc-tone-2 { background: linear-gradient(140deg, #343846, #14161b); }
.fc-no { position: absolute; left: 3px; bottom: 2px; font-family: ui-monospace, monospace; font-size: 5.5px; letter-spacing: 0.08em; color: rgba(255, 255, 255, 0.6); }
.fc-star { position: absolute; right: 4px; top: 3px; font-size: 9px; color: #e6c67c; }
.fc-cap { padding: 5px 2px 2px; }
.fc-call { display: flex; align-items: center; gap: 7px; transition: color 0.2s ease; }
.fc-call:hover { color: #c7a86a; }
.fc-rise { animation: fc-in 0.65s cubic-bezier(0.2, 0.9, 0.25, 1) both; }
.fc-d1 { animation-delay: 0.08s; }
.fc-d3 { animation-delay: 0.3s; }
@keyframes fc-in {
  from { opacity: 0; transform: translateY(10px); }
  to { opacity: 1; transform: none; }
}
@media (prefers-reduced-motion: reduce) {
  .fc-rise { animation: none !important; }
  .fc-call { transition: none !important; }
}
</style>
