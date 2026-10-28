<template>
  <div class="relative min-h-screen overflow-hidden bg-[#f2ecdf] text-[#2b2620]">
    <div class="wk-fiber pointer-events-none absolute inset-0"></div>
    <div class="pointer-events-none absolute inset-4 border border-[#b9ab94]/50"></div>

    <div class="relative z-10 mx-auto w-full max-w-md px-7 pb-14 pt-9">
      <div class="wk-rise flex items-start justify-between gap-4">
        <div class="min-w-0 flex-1">
          <p class="font-mono text-[8px] uppercase tracking-[0.3em] text-[#9a8a72]">Seasonal Menu · 献立</p>
          <p v-if="payload?.hours" class="mt-6 inline-flex items-center gap-1.5 text-[10px] tracking-[0.08em] text-[#7c6f5c]">
            <Clock class="h-3 w-3 text-[#9e3b32]" />{{ payload.hours }}
          </p>
          <p v-if="payload?.address" class="mt-1.5 max-w-[12rem] text-[9.5px] leading-relaxed text-[#8a7c68]">{{ payload.address }}</p>
          <span class="wk-hanko mt-5">旬</span>
        </div>
        <div class="flex flex-col items-end">
          <h1 class="wk-vertical text-[30px] font-semibold leading-none tracking-[0.12em] text-[#2b2620]">
            {{ payload?.shopName || '和纸献立' }}
          </h1>
          <span class="mt-3 h-16 w-px bg-[#9e3b32]/50"></span>
        </div>
      </div>

      <p v-if="payload?.tagline" class="wk-rise wk-d1 mt-6 max-w-[17rem] text-[11.5px] leading-[1.9] tracking-[0.04em] text-[#6f6352]">
        {{ payload.tagline }}
      </p>

      <div class="wk-rise wk-d1 mt-7 flex items-center gap-3">
        <span class="h-px flex-1 bg-[#b9ab94]"></span>
        <span class="text-[10px] tracking-[0.4em] text-[#9e3b32]">お品書き</span>
        <span class="h-px flex-1 bg-[#b9ab94]"></span>
      </div>

      <section v-if="groups.length" class="mt-3">
        <div v-for="(group, gi) in groups" :key="gi" class="wk-rise wk-group" :style="{ animationDelay: 0.12 + gi * 0.07 + 's' }">
          <div class="flex items-baseline justify-between border-b border-[#2b2620]/25 pb-1">
            <h2 class="text-[15px] font-semibold tracking-[0.18em]">{{ group.name }}</h2>
            <span class="font-mono text-[8px] tracking-[0.1em] text-[#9a8a72]">其の{{ ['一', '二', '三', '四', '五', '六', '七', '八', '九'][gi] || gi + 1 }}</span>
          </div>
          <ul class="mt-1">
            <li v-for="(item, i) in group.items" :key="i" class="wk-row">
              <div class="min-w-0 flex-1">
                <div class="flex flex-wrap items-baseline gap-x-2">
                  <span class="text-[13.5px] tracking-[0.04em] text-[#2b2620]">{{ item?.name }}</span>
                  <span v-if="signature(item)" class="wk-tag">おすすめ</span>
                  <span v-if="veg(item)" class="wk-tag wk-tag-g">精進</span>
                  <span v-if="spicy(item)" class="inline-flex text-[#9e3b32]"><Flame v-for="n in spicy(item)" :key="n" class="h-2.5 w-2.5" /></span>
                </div>
                <p v-if="item?.desc" class="mt-0.5 max-w-[15rem] text-[9.5px] leading-[1.6] text-[#8a7c68]">{{ item.desc }}</p>
              </div>
              <span class="wk-leader"></span>
              <span v-if="hasPrice(item)" class="flex-shrink-0 text-[12.5px] tabular-nums tracking-wide text-[#9e3b32]">¥{{ item.price }}</span>
              <span v-else class="font-mono text-[9px] text-[#b3a58f]">—</span>
            </li>
          </ul>
        </div>
      </section>

      <div v-else class="wk-rise mt-6 border border-[#b9ab94] px-6 py-12 text-center">
        <UtensilsCrossed class="mx-auto h-5 w-5 text-[#9a8a72]" />
        <p class="mt-3 text-[11px] tracking-[0.1em] text-[#8a7c68]">本日の献立は準備中です</p>
      </div>

      <div v-if="payload?.notice" class="wk-rise mt-6 border-l-2 border-[#9e3b32]/60 pl-3">
        <p class="text-[10px] leading-[1.8] tracking-[0.03em] text-[#7c6f5c]">{{ payload.notice }}</p>
      </div>

      <div class="wk-rise wk-d2 mt-7 flex items-end justify-between gap-4 border-t border-[#2b2620]/25 pt-4">
        <div>
          <p class="font-mono text-[7px] uppercase tracking-[0.26em] text-[#9a8a72]">Reservation</p>
          <a v-if="payload?.phone" :href="'tel:' + payload.phone" class="mt-1 inline-flex items-center gap-2 text-[15px] tracking-[0.06em] hover:text-[#9e3b32]">
            <Phone class="h-3.5 w-3.5" />{{ payload.phone }}
          </a>
        </div>
        <p class="text-right font-mono text-[7px] leading-relaxed tracking-[0.16em] text-[#a89980]">一期一会<br />NFC {{ nfcId || '—' }}</p>
      </div>

      <p class="mt-7 text-center text-[9px] tracking-[0.5em] text-[#b3a58f]">いただきます</p>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { Clock, Phone, Flame, UtensilsCrossed } from 'lucide-vue-next'

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
.wk-fiber {
  opacity: 0.55;
  background-image: repeating-linear-gradient(8deg, rgba(120, 100, 70, 0.05) 0 1px, transparent 1px 4px),
    radial-gradient(rgba(120, 100, 70, 0.08) 0.6px, transparent 0.6px);
  background-size: auto, 7px 7px;
}
.wk-vertical {
  writing-mode: vertical-rl;
  text-orientation: upright;
  font-family: 'Songti SC', 'STSong', 'Yu Mincho', 'Hiragino Mincho ProN', serif;
}
.wk-hanko {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  height: 27px;
  width: 27px;
  border: 1.5px solid #9e3b32;
  border-radius: 2px;
  color: #9e3b32;
  font-family: 'Songti SC', serif;
  font-size: 14px;
  transform: rotate(-4deg);
}
.wk-group { padding-top: 18px; }
.wk-row {
  display: flex;
  align-items: baseline;
  gap: 8px;
  padding: 8px 0;
  border-bottom: 1px dotted rgba(120, 100, 70, 0.28);
}
.wk-row:last-child { border-bottom: none; }
.wk-leader { flex: 1; min-width: 10px; border-bottom: 1px dotted rgba(120, 100, 70, 0.2); transform: translateY(-2px); }
.wk-tag { border: 1px solid rgba(158, 59, 50, 0.5); padding: 0 3px; font-size: 7px; letter-spacing: 0.06em; color: #9e3b32; }
.wk-tag-g { border-color: rgba(95, 115, 80, 0.5); color: #5f7350; }
.wk-rise { animation: wk-in 0.7s cubic-bezier(0.2, 0.9, 0.25, 1) both; }
.wk-d1 { animation-delay: 0.1s; }
.wk-d2 { animation-delay: 0.32s; }
@keyframes wk-in {
  from { opacity: 0; transform: translateY(12px); }
  to { opacity: 1; transform: none; }
}
@media (prefers-reduced-motion: reduce) {
  .wk-rise { animation: none !important; }
}
</style>
