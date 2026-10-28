<template>
  <div class="relative min-h-screen overflow-hidden bg-[#fbf8ef] text-[#191817]">
    <div class="mp-paper pointer-events-none absolute inset-0"></div>

    <div class="relative z-10 mx-auto w-full max-w-md px-5 pb-12 pt-6">
      <div class="mp-rise flex items-center justify-between border-b border-[#191817] pb-1 font-mono text-[7px] uppercase tracking-[0.18em] text-[#6f6a60]">
        <span>2026 年 10 月 20 日 · 星期二</span>
        <span>零售价 ¥2.00</span>
      </div>

      <header class="mp-rise mp-d1 mt-3 text-center">
        <div class="flex items-center justify-center gap-3">
          <span class="h-px w-8 bg-[#b3261e]"></span>
          <span class="font-mono text-[8px] uppercase tracking-[0.3em] text-[#b3261e]">Morning Edition</span>
          <span class="h-px w-8 bg-[#b3261e]"></span>
        </div>
        <h1 class="mp-masthead mt-1 text-[54px] font-black leading-none tracking-[-0.06em] text-[#b3261e]">晨报食单</h1>
        <div class="mt-1 flex items-center justify-between border-y-2 border-[#191817] py-1 font-mono text-[7px] uppercase tracking-[0.14em] text-[#5f5a52]">
          <span>第 1286 期</span>
          <span v-if="payload?.hours">营业 {{ payload.hours }}</span>
          <span>今日八版</span>
        </div>
      </header>

      <section class="mp-rise mp-d2 mt-4">
        <div class="flex items-center gap-2">
          <span class="mp-flag">头版</span>
          <span class="h-px flex-1 bg-[#191817]"></span>
          <span v-if="nfcId" class="font-mono text-[7px] text-[#8a8478]">NFC {{ nfcId }}</span>
        </div>

        <div class="mt-2 grid grid-cols-[1.35fr_1fr] gap-3">
          <div class="min-w-0">
            <h2 class="mp-serif text-[27px] font-bold leading-[1.08] tracking-[-0.02em]">{{ payload?.shopName }}</h2>
            <p v-if="payload?.tagline" class="mp-serif mt-2 text-[11px] leading-[1.65] text-[#5f5a52]">{{ payload.tagline }}</p>
            <div v-if="payload?.address" class="mt-2 flex items-start gap-1 font-mono text-[7px] leading-relaxed text-[#7c766b]">
              <MapPin class="mt-0.5 h-2.5 w-2.5 flex-shrink-0" />{{ payload.address }}
            </div>
          </div>
          <div class="mp-photo">
            <svg viewBox="0 0 80 80" fill="none" aria-hidden="true">
              <path d="M19 45h42c-2 12-9 18-21 18s-19-6-21-18Z" fill="#f6ead2" stroke="#191817" stroke-width="1" />
              <path d="M13 45h54" stroke="#191817" stroke-width="1" />
              <path d="M31 35c0-4 4-4 4-8M40 35c0-4 4-4 4-8M49 35c0-4 4-4 4-8" stroke="#b3261e" stroke-width="1.2" stroke-linecap="round" />
              <path d="M30 50c3 3 17 3 20 0" stroke="#b3261e" stroke-width="1" opacity="0.5" />
            </svg>
            <span class="mp-caption">今晨 · 一碗热汤</span>
          </div>
        </div>
      </section>

      <div class="mp-rise mp-d2 mt-4 flex items-center gap-2">
        <span class="h-px flex-1 bg-[#191817]"></span>
        <span class="mp-serif text-[11px] font-bold uppercase tracking-[0.18em]">今日菜单 · THE MENU</span>
        <span class="h-px flex-1 bg-[#191817]"></span>
      </div>

      <section v-if="groups.length" class="mp-rise mp-d3 mt-3 columns-2 gap-4">
        <div v-for="(group, gi) in groups" :key="gi" class="mb-4 break-inside-avoid">
          <div class="flex items-center gap-1.5 border-b border-[#191817] pb-1">
            <span class="h-1.5 w-1.5 bg-[#b3261e]"></span>
            <h3 class="mp-serif text-[13px] font-bold">{{ group.name }}</h3>
          </div>
          <ul class="mt-1">
            <li v-for="(item, i) in group.items" :key="i" class="mp-item">
              <div class="flex items-baseline gap-1">
                <span class="mp-serif text-[12.5px] font-bold leading-tight">{{ item?.name }}</span>
                <span v-if="signature(item)" class="mp-pick">荐</span>
                <span v-if="veg(item)" class="mp-veg">素</span>
              </div>
              <p v-if="item?.desc" class="mt-0.5 text-[9px] leading-[1.45] text-[#6f6a60]">{{ item.desc }}</p>
              <div class="mt-0.5 flex items-center gap-1">
                <span class="mp-leader"></span>
                <span v-if="hasPrice(item)" class="mp-serif text-[10.5px] font-bold">¥{{ item.price }}</span>
                <span v-else class="font-mono text-[8px] text-[#9a9488]">—</span>
              </div>
            </li>
          </ul>
        </div>
      </section>

      <div v-else class="mp-rise mt-4 border border-[#191817] px-5 py-10 text-center">
        <Newspaper class="mx-auto h-5 w-5 text-[#b3261e]" />
        <p class="mp-serif mt-2 text-[12px]">菜单付印中，请稍候。</p>
      </div>

      <div v-if="payload?.notice" class="mp-rise mt-4 border-y border-[#191817] py-2">
        <p class="mp-serif text-[10px] leading-[1.7] text-[#5f5a52]">{{ payload.notice }}</p>
      </div>

      <div class="mp-rise mt-4 flex items-center justify-between gap-3">
        <div class="min-w-0">
          <p class="font-mono text-[6.5px] uppercase tracking-[0.2em] text-[#8a8478]">订餐热线</p>
          <a v-if="payload?.phone" :href="'tel:' + payload.phone" class="mp-serif text-[15px] font-bold text-[#b3261e] hover:underline">{{ payload.phone }}</a>
          <span v-else class="mp-serif text-[12px] text-[#9a9488]">—</span>
        </div>
        <span class="mp-seal">今日<br />好味</span>
      </div>

      <p class="mt-5 text-center font-mono text-[6.5px] uppercase tracking-[0.3em] text-[#9a9488]">Breakfast · Lunch · Dinner</p>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { MapPin, Newspaper } from 'lucide-vue-next'

const props = defineProps(['payload', 'nfcId'])

const hasPrice = (item) => item?.price !== undefined && item?.price !== null && item?.price !== ''
const signature = (item) => item?.signature === true || (Array.isArray(item?.tags) && item.tags.some((t) => /招牌|必点|signature/i.test(String(t))))
const veg = (item) => item?.vegetarian === true || (Array.isArray(item?.tags) && item.tags.some((t) => /素|vegan|vegetarian/i.test(String(t))))

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
.mp-paper {
  opacity: 0.45;
  background-image: radial-gradient(rgba(25, 24, 23, 0.1) 0.6px, transparent 0.6px);
  background-size: 5px 5px;
}
.mp-masthead { font-family: 'Songti SC', 'STSong', 'SimSun', Georgia, serif; }
.mp-serif { font-family: 'Songti SC', 'STSong', 'SimSun', Georgia, serif; }
.mp-flag { border: 1px solid #b3261e; background: #b3261e; padding: 1px 5px; font-size: 8px; font-weight: 700; letter-spacing: 0.1em; color: #fbf8ef; }
.mp-photo {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  border: 1px solid #191817;
  background: repeating-linear-gradient(45deg, #f0e7d4 0 4px, #ece1ca 4px 8px);
  padding: 6px;
}
.mp-photo svg { width: 100%; height: auto; }
.mp-caption { margin-top: 3px; border-top: 1px dotted #8a8478; padding-top: 2px; font-family: ui-monospace, monospace; font-size: 6px; letter-spacing: 0.08em; color: #6f6a60; }
.mp-item { padding: 5px 0 4px; border-bottom: 1px dotted #b9b2a5; }
.mp-item:last-child { border-bottom: none; }
.mp-leader { flex: 1; border-bottom: 1px dotted #8a8478; transform: translateY(-2px); }
.mp-pick { border: 1px solid #b3261e; padding: 0 2px; font-size: 6.5px; font-weight: 700; color: #b3261e; }
.mp-veg { border: 1px solid #607a4e; padding: 0 2px; font-size: 6.5px; color: #607a4e; }
.mp-seal {
  display: inline-flex;
  height: 42px;
  width: 42px;
  align-items: center;
  justify-content: center;
  border: 2px solid #b3261e;
  border-radius: 9999px;
  text-align: center;
  font-family: 'Songti SC', serif;
  font-size: 9px;
  font-weight: 700;
  line-height: 1.3;
  color: #b3261e;
  transform: rotate(-8deg);
}
.mp-rise { animation: mp-in 0.7s cubic-bezier(0.2, 0.9, 0.25, 1) both; }
.mp-d1 { animation-delay: 0.08s; }
.mp-d2 { animation-delay: 0.15s; }
.mp-d3 { animation-delay: 0.22s; }
@keyframes mp-in {
  from { opacity: 0; transform: translateY(12px); }
  to { opacity: 1; transform: none; }
}
@media (prefers-reduced-motion: reduce) {
  .mp-rise { animation: none !important; }
}
</style>
