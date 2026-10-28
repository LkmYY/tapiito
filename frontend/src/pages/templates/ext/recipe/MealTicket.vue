<template>
  <div class="relative min-h-screen overflow-hidden bg-[#d9c9a3] text-[#34291a]">
    <div class="mt-cloth pointer-events-none absolute inset-0"></div>

    <div class="relative z-10 mx-auto w-full max-w-md px-5 pb-14 pt-8">
      <div class="mt-rise flex items-center justify-between font-mono text-[8px] uppercase tracking-[0.2em] text-[#6e5d3f]">
        <span>Canteen · 内部饭票</span>
        <span v-if="nfcId">票号 {{ nfcId }}</span>
      </div>

      <header class="mt-rise mt-d1 mt-5 flex items-end justify-between gap-4">
        <div class="min-w-0 flex-1">
          <p class="font-mono text-[8px] uppercase tracking-[0.3em] text-[#8a3a22]">Good for one meal</p>
          <h1 class="mt-serif mt-2 text-[37px] font-bold leading-[1.02] text-[#34291a]">{{ payload?.shopName || '食堂饭票' }}</h1>
          <p v-if="payload?.tagline" class="mt-2 max-w-[17rem] text-[10.5px] leading-[1.7] text-[#6e5d3f]">{{ payload.tagline }}</p>
        </div>
        <div class="mt-seal">凭票<br />取餐</div>
      </header>

      <div class="mt-rise mt-d2 mt-5 flex items-center justify-between border-y border-dashed border-[#6e5d3f]/60 py-2 font-mono text-[7.5px] uppercase tracking-[0.12em] text-[#6e5d3f]">
        <span v-if="payload?.hours">供应 {{ payload.hours }}</span>
        <span v-else>供应时间 —</span>
        <span>盖章有效</span>
      </div>

      <section v-if="groups.length" class="mt-4 space-y-3">
        <article v-for="(group, gi) in groups" :key="gi" class="mt-rise mt-ticket" :style="{ animationDelay: 0.12 + gi * 0.07 + 's' }">
          <div class="mt-stub">
            <span class="mt-stub-k">MEAL</span>
            <span class="mt-stub-v">{{ pad(gi + 1) }}</span>
            <span class="mt-stub-k">{{ group.items.length }} ITEMS</span>
          </div>
          <div class="mt-perf"></div>
          <div class="mt-main">
            <div class="flex items-baseline justify-between gap-2">
              <h2 class="mt-serif text-[17px] font-bold">{{ group.name }}</h2>
              <span class="font-mono text-[7px] tracking-[0.12em] text-[#94805e]">NO. {{ pad(gi + 1) }}00</span>
            </div>
            <ul class="mt-1">
              <li v-for="(item, i) in group.items" :key="i" class="mt-item">
                <span class="font-mono text-[7px] text-[#94805e]">{{ pad(i + 1) }}</span>
                <div class="min-w-0 flex-1">
                  <div class="flex flex-wrap items-baseline gap-x-1.5">
                    <span class="text-[12.5px] font-medium text-[#34291a]">{{ item?.name }}</span>
                    <span v-if="signature(item)" class="mt-tag">加菜</span>
                    <span v-if="veg(item)" class="mt-tag mt-tag-g">素</span>
                    <span v-if="spicy(item)" class="inline-flex text-[#a8512f]"><Flame v-for="n in spicy(item)" :key="n" class="h-2.5 w-2.5" /></span>
                  </div>
                  <p v-if="item?.desc" class="mt-0.5 text-[9px] leading-[1.45] text-[#7c6a4b]">{{ item.desc }}</p>
                </div>
                <span v-if="hasPrice(item)" class="mt-price">{{ item.price }}</span>
                <span v-else class="font-mono text-[9px] text-[#94805e]">—</span>
              </li>
            </ul>
            <div class="mt-cutline"></div>
            <p class="mt-thanks">请将饭票交至窗口 · THANK YOU</p>
          </div>
        </article>
      </section>

      <div v-else class="mt-rise mt-4 border border-dashed border-[#6e5d3f] bg-[#efe3c6]/70 px-6 py-12 text-center">
        <Ticket class="mx-auto h-6 w-6 text-[#8a7652]" />
        <p class="mt-serif mt-3 text-[12px]">今日饭票尚未开印。</p>
      </div>

      <div v-if="payload?.notice" class="mt-rise mt-4 flex items-start gap-2 rounded-sm bg-[#efe3c6] px-3 py-2.5">
        <Info class="mt-0.5 h-3.5 w-3.5 flex-shrink-0 text-[#8a3a22]" />
        <p class="text-[9.5px] leading-[1.65] text-[#6e5d3f]">{{ payload.notice }}</p>
      </div>

      <div class="mt-rise mt-d3 mt-5 flex items-center justify-between gap-3">
        <a v-if="payload?.phone" :href="'tel:' + payload.phone" class="mt-call">
          <Phone class="h-3.5 w-3.5" />
          <span class="font-mono text-[9px] tracking-[0.08em]">电话订餐</span>
          <span class="ml-auto font-mono text-[9px]">{{ payload.phone }}</span>
        </a>
        <p v-if="payload?.address" class="max-w-[9rem] text-right text-[8px] leading-[1.5] text-[#7c6a4b]">{{ payload.address }}</p>
      </div>

      <p class="mt-7 text-center font-mono text-[6.5px] uppercase tracking-[0.3em] text-[#8a7652]">One ticket · One warm meal</p>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { Flame, Info, Phone, Ticket } from 'lucide-vue-next'

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
.mt-cloth {
  opacity: 0.7;
  background-image: repeating-linear-gradient(0deg, rgba(90, 70, 40, 0.07) 0 1px, transparent 1px 3px),
    repeating-linear-gradient(90deg, rgba(255, 250, 235, 0.16) 0 1px, transparent 1px 4px);
}
.mt-serif { font-family: 'Songti SC', 'STSong', 'SimSun', Georgia, serif; }
.mt-seal {
  display: flex;
  height: 50px;
  width: 50px;
  flex-shrink: 0;
  align-items: center;
  justify-content: center;
  border: 1.5px solid #8a3a22;
  border-radius: 3px;
  text-align: center;
  font-family: 'Songti SC', serif;
  font-size: 9px;
  font-weight: 700;
  line-height: 1.3;
  letter-spacing: 0.1em;
  color: #8a3a22;
  transform: rotate(-5deg);
  opacity: 0.8;
}
.mt-ticket {
  position: relative;
  display: flex;
  overflow: hidden;
  border-radius: 4px;
  background: #f3e7c9;
  box-shadow: 0 3px 0 rgba(110, 93, 63, 0.25), 0 16px 30px -26px rgba(52, 41, 26, 0.8);
  transition: transform 0.16s ease;
}
.mt-ticket:hover { transform: translateY(-2px) rotate(-0.25deg); }
.mt-stub {
  display: flex;
  width: 50px;
  flex-shrink: 0;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  gap: 3px;
  background: #8a3a22;
  color: #f3e7c9;
}
.mt-stub-k { font-family: ui-monospace, monospace; font-size: 5.5px; letter-spacing: 0.12em; opacity: 0.75; }
.mt-stub-v { font-family: 'Songti SC', serif; font-size: 22px; line-height: 1; }
.mt-perf {
  width: 9px;
  flex-shrink: 0;
  background: radial-gradient(circle at 0 5px, #d9c9a3 3.5px, transparent 4px) 0 0 / 9px 10px repeat-y;
}
.mt-main { min-width: 0; flex: 1; padding: 10px 12px 8px 6px; }
.mt-item { display: flex; align-items: baseline; gap: 6px; padding: 6px 0; border-bottom: 1px dotted rgba(110, 93, 63, 0.28); }
.mt-item:last-child { border-bottom: none; }
.mt-tag { border: 1px solid #8a3a22; padding: 0 2px; font-size: 6px; color: #8a3a22; }
.mt-tag-g { border-color: #5f7348; color: #5f7348; }
.mt-price { flex-shrink: 0; font-family: ui-monospace, monospace; font-size: 12px; font-weight: 700; color: #8a3a22; }
.mt-price::before { content: '¥'; margin-right: 1px; font-size: 7px; }
.mt-cutline { margin-top: 4px; border-top: 1px dashed rgba(110, 93, 63, 0.5); }
.mt-thanks { margin-top: 4px; font-family: ui-monospace, monospace; font-size: 5.5px; letter-spacing: 0.16em; color: #94805e; }
.mt-call {
  display: flex;
  flex: 1;
  align-items: center;
  gap: 7px;
  border: 1px solid #6e5d3f;
  border-radius: 3px;
  background: #efe3c6;
  padding: 9px 11px;
  color: #4a3b26;
  transition: background-color 0.2s ease, transform 0.15s ease;
}
.mt-call:hover { background: #f7edd5; }
.mt-call:active { transform: scale(0.985); }
.mt-rise { animation: mt-in 0.65s cubic-bezier(0.2, 0.9, 0.25, 1) both; }
.mt-d1 { animation-delay: 0.08s; }
.mt-d2 { animation-delay: 0.15s; }
.mt-d3 { animation-delay: 0.34s; }
@keyframes mt-in {
  from { opacity: 0; transform: translateY(12px); }
  to { opacity: 1; transform: none; }
}
@media (prefers-reduced-motion: reduce) {
  .mt-rise, .mt-ticket, .mt-call { animation: none !important; transition: none !important; }
  .mt-ticket:hover { transform: none; }
}
</style>
