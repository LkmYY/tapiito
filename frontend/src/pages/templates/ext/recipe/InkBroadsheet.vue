<template>
  <div class="relative min-h-screen overflow-hidden bg-[#f4f1e8] text-[#131211]">
    <div class="ib-paper pointer-events-none absolute inset-0"></div>

    <div class="relative z-10 mx-auto w-full max-w-md px-5 pb-12 pt-7">
      <header class="ib-rise border-y-4 border-[#131211] py-2">
        <div class="flex items-end justify-between">
          <h1 class="ib-masthead text-[40px] font-black leading-none tracking-[-0.04em]">每日食报</h1>
          <div class="text-right font-mono text-[7.5px] uppercase leading-tight tracking-[0.12em] text-[#5f5c55]">
            <p>The Daily Table</p>
            <p>No. 1286 · 晴</p>
            <p v-if="nfcId">NFC {{ nfcId }}</p>
          </div>
        </div>
      </header>

      <div class="ib-rise ib-d1 flex items-center justify-between border-b border-[#131211] py-1.5 font-mono text-[7.5px] uppercase tracking-[0.16em] text-[#5f5c55]">
        <span>晨间六时出刊</span>
        <span v-if="payload?.hours">{{ payload.hours }}</span>
        <span v-else>营业时间 —</span>
        <span>第 01 版</span>
      </div>

      <section class="ib-rise ib-d2 mt-5 text-center">
        <p class="ib-kicker">今日头版 · 一碗人间烟火</p>
        <h2 class="ib-serif mx-auto mt-2 max-w-[19rem] text-[30px] font-bold leading-[1.12]">
          {{ payload?.shopName }}
        </h2>
        <p v-if="payload?.tagline" class="ib-serif mx-auto mt-2 max-w-[19rem] text-[12px] leading-[1.7] text-[#55524c]">
          {{ payload.tagline }}
        </p>
        <div class="mt-3 flex items-center justify-center gap-2 font-mono text-[8px] uppercase tracking-[0.14em] text-[#77736a]">
          <span v-if="payload?.address" class="inline-flex items-center gap-1"><MapPin class="h-2.5 w-2.5" />{{ payload.address }}</span>
        </div>
      </section>

      <div class="ib-rise ib-d2 mt-5 flex items-center gap-2">
        <span class="h-px flex-1 bg-[#131211]"></span>
        <span class="font-mono text-[8px] font-bold uppercase tracking-[0.28em]">菜单要闻</span>
        <span class="h-px flex-1 bg-[#131211]"></span>
      </div>

      <section v-if="groups.length" class="ib-rise ib-d3 mt-4 columns-2 gap-4 border-b border-[#131211] pb-5">
        <div v-for="(group, gi) in groups" :key="gi" class="mb-5 break-inside-avoid">
          <h3 class="ib-cat">{{ group.name }}</h3>
          <ul>
            <li v-for="(item, i) in group.items" :key="i" class="ib-item">
              <div class="flex items-baseline gap-1">
                <span class="ib-serif text-[13.5px] font-bold leading-tight">{{ item?.name }}</span>
                <span v-if="signature(item)" class="ib-seal">招</span>
              </div>
              <p v-if="item?.desc" class="mt-0.5 text-[9.5px] leading-[1.5] text-[#625f58]">{{ item.desc }}</p>
              <div class="mt-0.5 flex items-center gap-1.5">
                <span class="ib-leader"></span>
                <span v-if="hasPrice(item)" class="ib-serif text-[11.5px] font-bold">¥{{ item.price }}</span>
                <span v-else class="font-mono text-[9px] text-[#8a867c]">—</span>
              </div>
              <div class="mt-0.5 flex items-center gap-2 font-mono text-[7px] uppercase tracking-[0.1em] text-[#7a766d]">
                <span v-if="veg(item)">Vegetarian</span>
                <span v-if="spicy(item)">Spicy {{ spicy(item) }}/3</span>
              </div>
            </li>
          </ul>
        </div>
      </section>

      <div v-else class="ib-rise mt-5 border border-[#131211] px-5 py-10 text-center">
        <Newspaper class="mx-auto h-6 w-6" />
        <p class="ib-serif mt-2 text-[13px]">本期菜单尚在采写，明日见报。</p>
      </div>

      <div v-if="payload?.notice" class="ib-rise mt-4 border-l-2 border-[#131211] pl-3">
        <p class="text-[10px] leading-relaxed text-[#5f5c55]">{{ payload.notice }}</p>
      </div>

      <div v-if="payload?.phone" class="ib-rise mt-5 flex items-center justify-between border border-[#131211] px-3 py-2.5">
        <div>
          <p class="font-mono text-[7px] uppercase tracking-[0.2em] text-[#77736a]">订餐热线 · Classified</p>
          <a :href="'tel:' + payload.phone" class="ib-serif text-[16px] font-bold hover:underline">{{ payload.phone }}</a>
        </div>
        <Phone class="h-5 w-5" />
      </div>

      <div class="mt-6 flex items-center justify-between border-t-4 border-[#131211] pt-2 font-mono text-[7px] uppercase tracking-[0.18em] text-[#7a766d]">
        <span>Printed in the kitchen</span>
        <span>¥ 2.00</span>
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { MapPin, Phone, Newspaper } from 'lucide-vue-next'

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
.ib-paper {
  opacity: 0.5;
  background-image: radial-gradient(rgba(19, 18, 17, 0.10) 0.7px, transparent 0.7px);
  background-size: 4px 4px;
}
.ib-masthead { font-family: 'Songti SC', 'STSong', 'SimSun', Georgia, serif; letter-spacing: -0.04em; }
.ib-serif { font-family: 'Songti SC', 'STSong', 'SimSun', Georgia, serif; }
.ib-kicker { font-family: ui-monospace, monospace; font-size: 8px; letter-spacing: 0.3em; color: #6d6961; }
.ib-cat {
  border-bottom: 3px double #131211;
  padding-bottom: 2px;
  font-family: 'Songti SC', 'STSong', serif;
  font-size: 14px;
  font-weight: 900;
  letter-spacing: 0.04em;
}
.ib-item { padding: 7px 0 6px; border-bottom: 1px dotted #b9b3a7; }
.ib-item:last-child { border-bottom: none; }
.ib-seal { border: 1px solid #131211; padding: 0 2px; font-size: 7px; font-weight: 700; line-height: 1.3; }
.ib-leader { flex: 1; border-bottom: 1px dotted #8d887e; transform: translateY(-2px); }
.ib-rise { animation: ib-in 0.65s ease-out both; }
.ib-d1 { animation-delay: 0.07s; }
.ib-d2 { animation-delay: 0.13s; }
.ib-d3 { animation-delay: 0.2s; }
@keyframes ib-in {
  from { opacity: 0; transform: translateY(12px); }
  to { opacity: 1; transform: none; }
}
@media (prefers-reduced-motion: reduce) {
  .ib-rise { animation: none !important; }
}
</style>
