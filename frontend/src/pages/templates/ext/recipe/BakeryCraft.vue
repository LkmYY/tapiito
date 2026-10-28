<template>
  <div class="relative min-h-screen overflow-hidden bg-[#cfa877] text-[#392614]">
    <div class="bk-kraft pointer-events-none absolute inset-0"></div>
    <div class="pointer-events-none absolute inset-x-0 top-0 h-40 bg-[radial-gradient(60%_100%_at_50%_0%,rgba(255,244,214,0.35),transparent_70%)]"></div>

    <div class="relative z-10 mx-auto w-full max-w-md px-5 pb-14 pt-8">
      <div class="bk-rise flex items-center justify-between">
        <span class="bk-stamp">BAKERY · DAILY</span>
        <span v-if="nfcId" class="font-mono text-[8px] uppercase tracking-[0.2em] text-[#6e4d2b]">NFC {{ nfcId }}</span>
      </div>

      <header class="bk-rise bk-d1 relative mt-5 overflow-hidden rounded-t-[26px] bg-[#fdf3d4] px-6 pb-6 pt-7 shadow-[0_20px_40px_-32px_rgba(57,38,20,0.8)]">
        <span class="bk-scallop"></span>
        <p class="font-mono text-[8px] uppercase tracking-[0.34em] text-[#a2763b]">Freshly baked · 每日现烤</p>
        <h1 class="bk-serif mt-2 text-[39px] font-bold leading-[1.03] tracking-[-0.02em] text-[#4a2f16]">
          {{ payload?.shopName || '面包房手记' }}
        </h1>
        <p v-if="payload?.tagline" class="mt-2.5 max-w-[18rem] text-[11.5px] leading-relaxed text-[#7d5a34]">{{ payload.tagline }}</p>
        <div class="mt-4 flex flex-wrap items-center gap-x-3 gap-y-1.5 border-t border-dashed border-[#cfa877] pt-3 font-mono text-[9px] tracking-[0.04em] text-[#8a6236]">
          <span v-if="payload?.hours" class="inline-flex items-center gap-1"><Clock class="h-3 w-3" />{{ payload.hours }}</span>
          <span v-if="payload?.address" class="inline-flex items-center gap-1"><MapPin class="h-3 w-3" />{{ payload.address }}</span>
        </div>
      </header>

      <div class="bk-rise bk-d1 flex items-center justify-center gap-2 bg-[#b5844f] py-1.5 text-[#fdf3d4]">
        <Wheat class="h-3 w-3" />
        <span class="font-mono text-[7.5px] uppercase tracking-[0.3em]">Baked with butter &amp; time</span>
        <Wheat class="h-3 w-3" />
      </div>

      <section v-if="groups.length" class="mt-5 space-y-5">
        <div v-for="(group, gi) in groups" :key="gi" class="bk-rise" :style="{ animationDelay: 0.12 + gi * 0.07 + 's' }">
          <div class="flex items-center justify-between">
            <span class="bk-label">{{ group.name }}</span>
            <span class="font-mono text-[8px] uppercase tracking-[0.14em] text-[#6e4d2b]">{{ pad(group.items.length) }} items</span>
          </div>

          <div class="mt-2 space-y-2">
            <article v-for="(item, i) in group.items" :key="i" class="bk-card">
              <span class="bk-notch bk-notch-l"></span>
              <span class="bk-notch bk-notch-r"></span>
              <div class="flex items-start gap-3">
                <div class="min-w-0 flex-1">
                  <div class="flex flex-wrap items-center gap-x-2 gap-y-1">
                    <h3 class="bk-serif text-[16.5px] font-bold leading-tight text-[#402915]">{{ item?.name }}</h3>
                    <span v-if="signature(item)" class="bk-best">本店招牌</span>
                    <span v-if="veg(item)" class="bk-veg">素</span>
                    <span v-if="spicy(item)" class="inline-flex text-[#a9602d]"><Flame v-for="n in spicy(item)" :key="n" class="h-2.5 w-2.5" /></span>
                  </div>
                  <p v-if="item?.desc" class="mt-1 text-[10px] leading-[1.55] text-[#85653f]">{{ item.desc }}</p>
                </div>
                <div class="flex flex-col items-end gap-1">
                  <span v-if="hasPrice(item)" class="bk-price">{{ item.price }}</span>
                  <span v-else class="font-mono text-[10px] text-[#a98b66]">—</span>
                  <span class="font-mono text-[6.5px] uppercase tracking-[0.12em] text-[#b09470]">No. {{ pad(i + 1) }}</span>
                </div>
              </div>
            </article>
          </div>
        </div>
      </section>

      <div v-else class="bk-rise mt-5 rounded-lg bg-[#fdf3d4]/80 px-6 py-12 text-center">
        <Croissant class="mx-auto h-6 w-6 text-[#a2763b]" />
        <p class="mt-3 text-[12px] text-[#8a6236]">烤箱预热中，菜单很快出炉。</p>
      </div>

      <div v-if="payload?.notice" class="bk-rise mt-5 rounded-lg border border-dashed border-[#6e4d2b]/40 bg-[#f7e8c2]/70 px-4 py-3">
        <p class="text-[10px] leading-relaxed text-[#755029]">{{ payload.notice }}</p>
      </div>

      <a v-if="payload?.phone" :href="'tel:' + payload.phone" class="bk-rise bk-d2 bk-call mt-6">
        <span class="bk-call-ico"><Phone class="h-4 w-4" /></span>
        <div>
          <p class="font-mono text-[7px] uppercase tracking-[0.22em] text-[#e8cfa4]">Reserve a loaf</p>
          <p class="font-mono text-[12px] tracking-[0.04em] text-[#fdf3d4]">{{ payload.phone }}</p>
        </div>
        <ArrowUpRight class="ml-auto h-4 w-4 text-[#e8cfa4]" />
      </a>

      <div class="mt-7 flex items-center justify-between font-mono text-[7px] uppercase tracking-[0.24em] text-[#795733]">
        <span>Small batch · Hand shaped</span>
        <span>Est. 2018</span>
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { Clock, MapPin, Phone, Flame, Wheat, Croissant, ArrowUpRight } from 'lucide-vue-next'

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
.bk-kraft {
  opacity: 0.8;
  background-image: repeating-linear-gradient(94deg, rgba(255, 240, 210, 0.08) 0 1px, transparent 1px 5px),
    repeating-linear-gradient(2deg, rgba(80, 50, 20, 0.06) 0 1px, transparent 1px 7px);
}
.bk-stamp {
  border: 1.5px solid #392614;
  padding: 3px 7px;
  font-family: ui-monospace, monospace;
  font-size: 7.5px;
  font-weight: 700;
  letter-spacing: 0.16em;
  color: #392614;
  transform: rotate(-2deg);
}
.bk-serif { font-family: 'Iowan Old Style', 'Palatino Linotype', 'Songti SC', Georgia, serif; }
.bk-scallop {
  position: absolute;
  inset: 0 0 auto;
  height: 10px;
  background: radial-gradient(circle at 5px 0, transparent 4.5px, #b5844f 5px) 0 0 / 10px 10px repeat-x;
  opacity: 0.5;
}
.bk-label {
  display: inline-block;
  border-radius: 2px;
  background: #4a2f16;
  padding: 3px 9px;
  font-size: 10px;
  font-weight: 700;
  letter-spacing: 0.12em;
  color: #fdf3d4;
  box-shadow: 1px 1px 0 rgba(57, 38, 20, 0.3);
}
.bk-card {
  position: relative;
  overflow: hidden;
  border-radius: 10px;
  background: #fbecc9;
  padding: 11px 14px;
  box-shadow: 0 2px 0 rgba(120, 82, 40, 0.22);
  transition: transform 0.16s ease;
}
.bk-card:hover { transform: translateX(2px); }
.bk-notch {
  position: absolute;
  top: 50%;
  height: 10px;
  width: 10px;
  transform: translateY(-50%);
  border-radius: 9999px;
  background: #cfa877;
  box-shadow: inset 0 1px 2px rgba(80, 50, 20, 0.35);
}
.bk-notch-l { left: -5px; }
.bk-notch-r { right: -5px; }
.bk-best { border-radius: 2px; background: #c47a33; padding: 0 4px; font-size: 7px; font-weight: 700; letter-spacing: 0.05em; color: #fff8e8; }
.bk-veg { border: 1px solid #78874f; padding: 0 3px; font-size: 7px; color: #64743f; }
.bk-price {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  min-width: 42px;
  border-radius: 9999px;
  background: #f4d47a;
  padding: 4px 8px;
  font-family: ui-monospace, monospace;
  font-size: 14px;
  font-weight: 800;
  color: #4a2f16;
  box-shadow: 0 2px 0 #d2ab55;
}
.bk-price::before { content: '¥'; margin-right: 1px; font-size: 8px; opacity: 0.65; }
.bk-call {
  display: flex;
  align-items: center;
  gap: 11px;
  border-radius: 12px;
  background: #402915;
  padding: 11px 14px;
  box-shadow: 0 18px 30px -24px rgba(57, 38, 20, 0.9);
  transition: background-color 0.2s ease, transform 0.15s ease;
}
.bk-call:hover { background-color: #55371c; }
.bk-call:active { transform: scale(0.985); }
.bk-call-ico { display: inline-flex; align-items: center; justify-content: center; height: 32px; width: 32px; border-radius: 9999px; background: #f4d47a; color: #402915; }
.bk-rise { animation: bk-in 0.7s cubic-bezier(0.2, 0.9, 0.25, 1) both; }
.bk-d1 { animation-delay: 0.08s; }
.bk-d2 { animation-delay: 0.32s; }
@keyframes bk-in {
  from { opacity: 0; transform: translateY(12px); }
  to { opacity: 1; transform: none; }
}
@media (prefers-reduced-motion: reduce) {
  .bk-rise, .bk-card, .bk-call { animation: none !important; transition: none !important; }
  .bk-card:hover { transform: none; }
}
</style>
