<template>
  <div class="relative min-h-screen overflow-hidden bg-[#0a0b09] text-[#edf0e6]">
    <div class="ml-glow pointer-events-none absolute inset-x-0 top-0 h-72"></div>
    <div class="ml-diag pointer-events-none absolute inset-x-0 top-0 h-1.5"></div>

    <div class="relative z-10 mx-auto w-full max-w-md px-5 pb-14 pt-8">
      <div class="ml-rise flex items-center justify-between font-mono text-[8px] uppercase tracking-[0.24em] text-[#8b957f]">
        <span>Macro Kitchen · 02</span>
        <span v-if="nfcId">NFC {{ nfcId }}</span>
      </div>

      <header class="ml-rise ml-d1 mt-6">
        <div class="flex items-end justify-between gap-3">
          <div class="min-w-0 flex-1">
            <p class="ml-cond text-[10px] uppercase tracking-[0.28em] text-[#a3e635]">Fuel your training</p>
            <h1 class="ml-cond mt-2 text-[51px] font-bold uppercase leading-[0.9] tracking-[-0.02em] text-[#f4f7ea]">
              {{ payload?.shopName || '增肌便当' }}
            </h1>
          </div>
          <span class="ml-stamp">HIGH<br />PROTEIN</span>
        </div>
        <p v-if="payload?.tagline" class="mt-3 max-w-[19rem] text-[11px] leading-relaxed text-[#9ba690]">{{ payload.tagline }}</p>
      </header>

      <div class="ml-rise ml-d2 mt-6 grid grid-cols-3 gap-2">
        <div class="ml-stat">
          <span class="ml-stat-k">DISHES</span>
          <span class="ml-stat-v">{{ pad(total) }}</span>
          <span class="ml-stat-u">道在售</span>
        </div>
        <div class="ml-stat">
          <span class="ml-stat-k">PLANT</span>
          <span class="ml-stat-v">{{ pad(vegCount) }}</span>
          <span class="ml-stat-u">植物系</span>
        </div>
        <div class="ml-stat ml-stat-hl">
          <span class="ml-stat-k">SIGNATURE</span>
          <span class="ml-stat-v">{{ pad(sigCount) }}</span>
          <span class="ml-stat-u">招牌款</span>
        </div>
      </div>

      <div class="ml-rise ml-d2 mt-5 flex items-center gap-2 border-y border-[#2a3126] py-2">
        <span class="ml-cond text-[10px] uppercase tracking-[0.2em] text-[#a3e635]">Today's Box</span>
        <span class="h-px flex-1 bg-[#2a3126]"></span>
        <span class="font-mono text-[8px] uppercase tracking-[0.1em] text-[#7d8974]">配比 · 现做 · 少油</span>
      </div>

      <section v-if="groups.length" class="mt-1">
        <div v-for="(group, gi) in groups" :key="gi" class="ml-rise ml-group" :style="{ animationDelay: 0.14 + gi * 0.06 + 's' }">
          <div class="flex items-center gap-2">
            <span class="ml-idx">{{ pad(gi + 1) }}</span>
            <h2 class="ml-cond text-[16px] uppercase tracking-[0.1em] text-[#e6ecd9]">{{ group.name }}</h2>
            <span class="ml-bar flex-1"></span>
            <span class="font-mono text-[8px] text-[#7d8974]">{{ pad(group.items.length) }}</span>
          </div>
          <ul>
            <li v-for="(item, i) in group.items" :key="i" class="ml-row">
              <div class="min-w-0 flex-1">
                <div class="flex flex-wrap items-center gap-x-2 gap-y-0.5">
                  <span class="text-[13.5px] font-semibold text-[#edf0e6]">{{ item?.name }}</span>
                  <span v-if="signature(item)" class="ml-tag ml-tag-lime">SIGNATURE</span>
                  <span v-if="veg(item)" class="ml-tag ml-tag-green">PLANT</span>
                  <span v-if="spicy(item)" class="inline-flex text-[#e0a35a]"><Flame v-for="n in spicy(item)" :key="n" class="h-2.5 w-2.5" /></span>
                </div>
                <p v-if="item?.desc" class="mt-0.5 max-w-[16rem] text-[9.5px] leading-[1.5] text-[#8b957f]">{{ item.desc }}</p>
              </div>
              <span v-if="hasPrice(item)" class="ml-price">¥{{ item.price }}</span>
              <span v-else class="font-mono text-[9px] text-[#5f6a58]">—</span>
            </li>
          </ul>
        </div>
      </section>

      <div v-else class="ml-rise mt-5 border border-dashed border-[#2a3126] px-6 py-12 text-center">
        <Dumbbell class="mx-auto h-6 w-6 text-[#7d8974]" />
        <p class="ml-cond mt-3 text-[13px] uppercase tracking-[0.12em] text-[#8b957f]">Meal prep in progress</p>
      </div>

      <div class="ml-rise ml-d3 mt-5 grid grid-cols-[1fr_auto] items-center gap-3 border border-[#2a3126] bg-[#10130e] px-4 py-3">
        <div class="min-w-0">
          <p class="font-mono text-[7px] uppercase tracking-[0.2em] text-[#7d8974]">Kitchen note</p>
          <p class="mt-1 text-[10px] leading-relaxed text-[#a8b39c]">{{ payload?.notice || '每份现做，主食可换糙米或藜麦。' }}</p>
        </div>
        <span class="ml-chip">NO<br />SUGAR</span>
      </div>

      <a v-if="payload?.phone" :href="'tel:' + payload.phone" class="ml-rise ml-d3 ml-call mt-4">
        <Phone class="h-4 w-4" />
        <span class="ml-cond text-[14px] uppercase tracking-[0.08em]">Order your box</span>
        <span class="ml-auto font-mono text-[9px] tracking-[0.04em]">{{ payload.phone }}</span>
      </a>

      <div class="mt-4 flex items-center justify-between gap-3 font-mono text-[8px] text-[#7d8974]">
        <span v-if="payload?.hours" class="inline-flex items-center gap-1"><Clock class="h-3 w-3" />{{ payload.hours }}</span>
        <span v-if="payload?.address" class="inline-flex items-center gap-1 text-right"><MapPin class="h-3 w-3 flex-shrink-0" />{{ payload.address }}</span>
      </div>

      <p class="mt-7 text-center font-mono text-[6.5px] uppercase tracking-[0.34em] text-[#5f6a58]">Train hard · Eat clean · Repeat</p>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { Phone, Flame, Dumbbell, Clock, MapPin } from 'lucide-vue-next'

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
const all = computed(() => groups.value.flatMap((g) => g.items))
const total = computed(() => all.value.length)
const vegCount = computed(() => all.value.filter(veg).length)
const sigCount = computed(() => all.value.filter(signature).length)
</script>

<style scoped>
.ml-glow { background: radial-gradient(60% 90% at 80% 0%, rgba(163, 230, 53, 0.14), transparent 70%); }
.ml-diag { background: repeating-linear-gradient(115deg, #a3e635 0 12px, #0a0b09 12px 24px); opacity: 0.8; }
.ml-cond { font-family: 'Oswald', 'Arial Narrow', 'PingFang SC', sans-serif; }
.ml-stamp {
  flex-shrink: 0;
  border: 1.5px solid #a3e635;
  padding: 4px 7px;
  text-align: center;
  font-family: ui-monospace, monospace;
  font-size: 6.5px;
  font-weight: 700;
  line-height: 1.4;
  letter-spacing: 0.12em;
  color: #a3e635;
  transform: rotate(-3deg);
}
.ml-stat { border: 1px solid #2a3126; border-radius: 4px; background: #10130e; padding: 8px; }
.ml-stat-hl { border-color: rgba(163, 230, 53, 0.5); background: rgba(163, 230, 53, 0.07); }
.ml-stat-k { display: block; font-family: ui-monospace, monospace; font-size: 6px; letter-spacing: 0.16em; color: #7d8974; }
.ml-stat-v { display: block; margin-top: 2px; font-family: 'Oswald', 'Arial Narrow', sans-serif; font-size: 27px; font-weight: 700; line-height: 1; color: #e8f2d8; }
.ml-stat-hl .ml-stat-v { color: #a3e635; }
.ml-stat-u { display: block; margin-top: 2px; font-size: 8px; color: #7d8974; }
.ml-group { padding-top: 18px; }
.ml-idx { display: inline-flex; align-items: center; justify-content: center; height: 17px; width: 17px; border-radius: 2px; background: #a3e635; color: #0a0b09; font-family: ui-monospace, monospace; font-size: 8px; font-weight: 800; }
.ml-bar { height: 1px; background: #2a3126; }
.ml-row { display: flex; align-items: center; gap: 10px; border-bottom: 1px solid #222820; padding: 10px 0; }
.ml-row:last-child { border-bottom: none; }
.ml-tag { padding: 0 3px; font-family: ui-monospace, monospace; font-size: 5.5px; font-weight: 700; letter-spacing: 0.06em; }
.ml-tag-lime { background: #a3e635; color: #0a0b09; }
.ml-tag-green { border: 1px solid #5d7a46; color: #a6bf8a; }
.ml-price { flex-shrink: 0; border-radius: 2px; border: 1px solid #3a4632; background: #161b13; padding: 3px 7px; font-family: ui-monospace, monospace; font-size: 12px; font-weight: 700; color: #d8e6c4; }
.ml-chip { display: inline-flex; height: 34px; width: 34px; align-items: center; justify-content: center; border-radius: 9999px; border: 1px solid #a3e635; text-align: center; font-family: ui-monospace, monospace; font-size: 6px; line-height: 1.3; letter-spacing: 0.06em; color: #a3e635; transform: rotate(8deg); }
.ml-call {
  display: flex;
  align-items: center;
  gap: 9px;
  border-radius: 3px;
  background: #a3e635;
  padding: 12px 14px;
  color: #0a0b09;
  transition: filter 0.2s ease, transform 0.15s ease;
}
.ml-call:hover { filter: brightness(1.06); }
.ml-call:active { transform: scale(0.985); }
.ml-rise { animation: ml-in 0.62s cubic-bezier(0.2, 0.9, 0.25, 1) both; }
.ml-d1 { animation-delay: 0.07s; }
.ml-d2 { animation-delay: 0.13s; }
.ml-d3 { animation-delay: 0.3s; }
@keyframes ml-in {
  from { opacity: 0; transform: translateY(12px); }
  to { opacity: 1; transform: none; }
}
@media (prefers-reduced-motion: reduce) {
  .ml-rise, .ml-call { animation: none !important; transition: none !important; }
}
</style>
