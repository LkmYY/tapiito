<template>
  <div class="relative min-h-screen overflow-hidden bg-[#f2eee5] text-[#191714]">
    <div class="fe-grain pointer-events-none absolute inset-0"></div>
    <div class="pointer-events-none absolute -right-16 top-24 select-none font-serif text-[280px] leading-none text-[#191714]/[0.035]">食</div>

    <div class="relative z-10 mx-auto w-full max-w-md px-6 pb-14 pt-8">
      <div class="fe-rise flex items-center justify-between border-b border-[#191714] pb-2 font-mono text-[9px] uppercase tracking-[0.26em]">
        <span>Gourmet Journal</span>
        <span>No. 07 / 2026</span>
      </div>

      <header class="fe-rise fe-d1 mt-8">
        <div class="flex items-start justify-between gap-4">
          <div class="min-w-0">
            <p class="font-mono text-[10px] uppercase tracking-[0.32em] text-[#a33d28]">Taste · Memory · Table</p>
            <h1 class="fe-display mt-2 text-[52px] font-black leading-[0.92] tracking-[-0.03em]">
              {{ payload?.shopName || '一张餐桌' }}
            </h1>
          </div>
          <div class="mt-1 flex flex-col items-center border border-[#191714] px-2 py-1.5 text-center">
            <span class="font-mono text-[8px] tracking-[0.16em]">VOL</span>
            <span class="font-serif text-[22px] leading-none">07</span>
          </div>
        </div>

        <p v-if="payload?.tagline" class="fe-serif mt-4 max-w-[20rem] border-l-2 border-[#a33d28] pl-3 text-[15px] italic leading-relaxed text-[#5e574c]">
          {{ payload.tagline }}
        </p>

        <div class="mt-5 flex flex-wrap items-center gap-x-4 gap-y-1.5 font-mono text-[9.5px] uppercase tracking-[0.13em] text-[#6b6458]">
          <span v-if="payload?.hours" class="inline-flex items-center gap-1.5"><Clock class="h-3 w-3 text-[#a33d28]" />{{ payload.hours }}</span>
          <span v-if="payload?.address" class="inline-flex items-center gap-1.5"><MapPin class="h-3 w-3 text-[#a33d28]" />{{ payload.address }}</span>
          <span v-if="nfcId">NFC · {{ nfcId }}</span>
        </div>
      </header>

      <div class="fe-rise fe-d2 mt-8 flex items-center gap-3">
        <span class="h-px flex-1 bg-[#191714]"></span>
        <span class="fe-serif text-[11px] italic tracking-[0.18em] text-[#a33d28]">The Menu</span>
        <span class="h-px flex-1 bg-[#191714]"></span>
      </div>

      <section v-if="groups.length" class="mt-5 columns-2 gap-6">
        <div v-for="(group, gi) in groups" :key="gi" class="fe-rise mb-7 break-inside-avoid" :style="{ animationDelay: 0.16 + gi * 0.06 + 's' }">
          <div class="flex items-center gap-2 border-b-2 border-[#191714] pb-1.5">
            <span class="fe-serif text-[13px] font-bold">{{ group.name }}</span>
            <span class="ml-auto font-mono text-[9px] text-[#a33d28]">{{ pad(group.items.length) }}</span>
          </div>
          <ul class="mt-2 space-y-3.5">
            <li v-for="(item, i) in group.items" :key="i" class="fe-item">
              <div class="flex items-baseline gap-1.5">
                <h3 class="fe-serif text-[14.5px] font-semibold leading-snug">{{ item?.name }}</h3>
                <span v-if="signature(item)" class="fe-star">★</span>
              </div>
              <p v-if="item?.desc" class="mt-0.5 text-[10.5px] leading-[1.55] text-[#7b7367]">{{ item.desc }}</p>
              <div class="mt-1 flex items-center gap-2">
                <span v-if="hasPrice(item)" class="fe-serif text-[12px] text-[#a33d28]">¥ {{ item.price }}</span>
                <span v-if="veg(item)" class="fe-tag">VEG</span>
                <span v-if="spicy(item)" class="inline-flex text-[#a33d28]"><Flame v-for="n in spicy(item)" :key="n" class="h-2.5 w-2.5" /></span>
                <span v-if="!hasPrice(item)" class="font-mono text-[8px] text-[#b2a99b]">—</span>
              </div>
            </li>
          </ul>
        </div>
      </section>

      <div v-else class="fe-rise mt-8 border border-[#191714] px-6 py-12 text-center">
        <UtensilsCrossed class="mx-auto h-5 w-5 text-[#a33d28]" />
        <p class="fe-serif mt-3 text-sm italic text-[#7b7367]">本期菜单正在排版，敬请期待。</p>
      </div>

      <div v-if="payload?.notice" class="fe-rise mt-4 border-t border-[#c9c1b2] pt-4">
        <p class="fe-serif text-[11.5px] italic leading-relaxed text-[#7b7367]">{{ payload.notice }}</p>
      </div>

      <div v-if="payload?.phone" class="fe-rise fe-d3 mt-7 flex items-center justify-between border-y-2 border-[#191714] py-3">
        <span class="font-mono text-[9px] uppercase tracking-[0.2em] text-[#7b7367]">Reservations</span>
        <a :href="'tel:' + payload.phone" class="fe-serif text-[16px] font-semibold tracking-wide hover:text-[#a33d28]">
          {{ payload.phone }}
        </a>
      </div>

      <p class="mt-8 text-center font-mono text-[8px] uppercase tracking-[0.4em] text-[#a8a093]">Eat Slowly · Live Fully</p>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { Clock, MapPin, Flame, UtensilsCrossed } from 'lucide-vue-next'

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
.fe-grain {
  opacity: 0.35;
  background-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' width='140' height='140'%3E%3Cfilter id='n'%3E%3CfeTurbulence type='fractalNoise' baseFrequency='0.95' numOctaves='3'/%3E%3C/filter%3E%3Crect width='140' height='140' filter='url(%23n)' opacity='0.08'/%3E%3C/svg%3E");
}
.fe-display { font-family: 'Playfair Display', 'Songti SC', 'STSong', Georgia, serif; }
.fe-serif { font-family: 'Songti SC', 'STSong', Georgia, 'Times New Roman', serif; }
.fe-star { font-size: 9px; color: #a33d28; transform: translateY(-4px); }
.fe-tag {
  border: 1px solid #9aa88b;
  padding: 0 3px;
  font-family: ui-monospace, monospace;
  font-size: 7.5px;
  letter-spacing: 0.08em;
  color: #66785a;
}
.fe-item { position: relative; }
.fe-item::after {
  content: '';
  display: block;
  margin-top: 10px;
  height: 1px;
  background: repeating-linear-gradient(90deg, #c9c1b2 0 2px, transparent 2px 5px);
  opacity: 0.6;
}
.fe-item:last-child::after { display: none; }
.fe-rise { animation: fe-in 0.75s cubic-bezier(0.2, 0.9, 0.25, 1) both; }
.fe-d1 { animation-delay: 0.08s; }
.fe-d2 { animation-delay: 0.14s; }
.fe-d3 { animation-delay: 0.4s; }
@keyframes fe-in {
  from { opacity: 0; transform: translateY(16px); }
  to { opacity: 1; transform: none; }
}
@media (prefers-reduced-motion: reduce) {
  .fe-rise { animation: none !important; }
  .fe-serif a, a { transition: none !important; }
}
</style>
