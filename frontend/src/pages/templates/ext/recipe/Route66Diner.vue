<template>
  <div class="relative min-h-screen overflow-hidden bg-[#f5e7cb] text-[#241b14]">
    <div class="rd-speckle pointer-events-none absolute inset-0"></div>
    <div class="rd-checker pointer-events-none absolute inset-x-0 top-0 h-3"></div>

    <div class="relative z-10 mx-auto w-full max-w-md px-5 pb-14 pt-8">
      <header class="rd-rise relative overflow-hidden rounded-t-[110px] border-[3px] border-[#b23a2a] bg-[#fbf3e1] px-6 pb-5 pt-9 text-center">
        <span class="rd-spark rd-spark-l"></span>
        <span class="rd-spark rd-spark-r"></span>
        <p class="font-mono text-[8px] uppercase tracking-[0.4em] text-[#1c6b68]">Est. 1966 · Route 66</p>
        <h1 class="rd-display mt-2 text-[36px] font-black uppercase leading-[0.95] tracking-[-0.01em] text-[#a52f22]">
          {{ payload?.shopName || '公路餐车' }}
        </h1>
        <p v-if="payload?.tagline" class="mx-auto mt-2 max-w-[18rem] text-[11px] leading-relaxed text-[#6c5843]">{{ payload.tagline }}</p>
        <div class="mt-3 flex items-center justify-center gap-3 font-mono text-[8px] uppercase tracking-[0.12em] text-[#8c7358]">
          <span v-if="payload?.hours" class="inline-flex items-center gap-1"><Clock class="h-3 w-3" />{{ payload.hours }}</span>
          <span class="h-1 w-1 rounded-full bg-[#b23a2a]"></span>
          <span>Dine-in · Take-out</span>
        </div>
        <span class="rd-badge">SINCE<br /><b>66</b></span>
      </header>

      <div class="rd-stripe"></div>

      <section v-if="groups.length" class="rd-rise rd-d1 mt-5">
        <div v-for="(group, gi) in groups" :key="gi" class="mb-5">
          <div class="flex items-center gap-2">
            <span class="rd-star">★</span>
            <h2 class="rd-display text-[17px] font-black uppercase tracking-[0.1em] text-[#a52f22]">{{ group.name }}</h2>
            <span class="rd-rule"></span>
            <span class="font-mono text-[8px] text-[#8c7358]">{{ pad(group.items.length) }}</span>
          </div>
          <ul class="mt-1.5 grid grid-cols-1 gap-x-4">
            <li v-for="(item, i) in group.items" :key="i" class="rd-item">
              <div class="min-w-0 flex-1">
                <div class="flex flex-wrap items-baseline gap-x-2">
                  <span class="rd-dish text-[15px] font-bold">{{ item?.name }}</span>
                  <span v-if="signature(item)" class="rd-pick">BEST</span>
                  <span v-if="veg(item)" class="rd-veg">V</span>
                  <span v-if="spicy(item)" class="inline-flex text-[#b23a2a]"><Flame v-for="n in spicy(item)" :key="n" class="h-3 w-3" /></span>
                </div>
                <p v-if="item?.desc" class="mt-0.5 text-[10px] leading-[1.45] text-[#7b6550]">{{ item.desc }}</p>
              </div>
              <span class="rd-dots"></span>
              <span v-if="hasPrice(item)" class="rd-price">{{ item.price }}</span>
              <span v-else class="font-mono text-[10px] text-[#b4a088]">—</span>
            </li>
          </ul>
        </div>
      </section>

      <div v-else class="rd-rise mt-4 rounded-lg border-2 border-dashed border-[#c9a883] bg-[#fbf3e1] px-6 py-12 text-center">
        <Coffee class="mx-auto h-6 w-6 text-[#b23a2a]" />
        <p class="rd-display mt-3 text-[13px] font-bold uppercase tracking-wider text-[#8c7358]">The grill is warming up</p>
      </div>

      <div v-if="payload?.notice" class="rd-rise mt-5 flex items-start gap-2 rounded border-2 border-[#1c6b68] bg-[#e3efec] px-3 py-2.5">
        <Info class="mt-0.5 h-3.5 w-3.5 flex-shrink-0 text-[#1c6b68]" />
        <p class="text-[10px] leading-relaxed text-[#2b5653]">{{ payload.notice }}</p>
      </div>

      <div class="rd-rise rd-d2 mt-6">
        <a v-if="payload?.phone" :href="'tel:' + payload.phone" class="rd-call">
          <span class="rd-call-ico"><Phone class="h-4 w-4" /></span>
          <span class="rd-display text-[15px] font-black uppercase tracking-[0.06em]">Call for a booth</span>
          <span class="ml-auto font-mono text-[10px] tracking-[0.04em]">{{ payload.phone }}</span>
        </a>
      </div>

      <div v-if="payload?.address" class="mt-4 flex items-start justify-center gap-1.5 text-center text-[9.5px] leading-relaxed text-[#8c7358]">
        <MapPin class="mt-0.5 h-3 w-3 flex-shrink-0" />
        <span>{{ payload.address }}</span>
      </div>

      <p class="mt-7 text-center font-mono text-[7.5px] uppercase tracking-[0.32em] text-[#a58a6d]">Good food · Cold drinks · Open late</p>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { Clock, MapPin, Phone, Flame, Coffee, Info } from 'lucide-vue-next'

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
.rd-speckle {
  opacity: 0.5;
  background-image: radial-gradient(rgba(90, 62, 34, 0.14) 0.6px, transparent 0.6px);
  background-size: 5px 5px;
}
.rd-checker {
  background-image: linear-gradient(45deg, #241b14 25%, transparent 25%, transparent 75%, #241b14 75%), linear-gradient(45deg, #241b14 25%, #f5e7cb 25%, #f5e7cb 75%, #241b14 75%);
  background-size: 12px 12px;
  background-position: 0 0, 6px 6px;
}
.rd-display { font-family: 'Archivo Black', 'Arial Black', 'PingFang SC', sans-serif; }
.rd-dish { font-family: 'Arial Black', 'PingFang SC', sans-serif; letter-spacing: -0.01em; }
.rd-spark { position: absolute; top: 14px; height: 22px; width: 22px; background: #f2c14e; clip-path: polygon(50% 0, 61% 35%, 98% 35%, 68% 57%, 79% 91%, 50% 70%, 21% 91%, 32% 57%, 2% 35%, 39% 35%); }
.rd-spark-l { left: 22px; transform: rotate(-12deg); }
.rd-spark-r { right: 22px; transform: rotate(14deg); }
.rd-badge {
  position: absolute;
  right: 14px;
  bottom: 12px;
  display: flex;
  height: 44px;
  width: 44px;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  border-radius: 9999px;
  border: 2px solid #1c6b68;
  background: #fbf3e1;
  font-family: ui-monospace, monospace;
  font-size: 6px;
  line-height: 1;
  letter-spacing: 0.1em;
  color: #1c6b68;
  transform: rotate(-8deg);
}
.rd-badge b { font-size: 16px; }
.rd-stripe {
  height: 9px;
  background: repeating-linear-gradient(90deg, #a52f22 0 18px, #f5e7cb 18px 30px, #1c6b68 30px 48px, #f5e7cb 48px 60px);
}
.rd-star { color: #f2c14e; font-size: 12px; text-shadow: 0 1px 0 #b98a1f; }
.rd-rule { flex: 1; height: 2px; background: repeating-linear-gradient(90deg, #b23a2a 0 10px, transparent 10px 15px); opacity: 0.35; }
.rd-item {
  display: flex;
  align-items: baseline;
  gap: 7px;
  border-bottom: 1px dashed rgba(140, 115, 88, 0.28);
  padding: 8px 0;
}
.rd-item:last-child { border-bottom: none; }
.rd-pick { border-radius: 2px; background: #a52f22; padding: 0 4px; font-family: ui-monospace, monospace; font-size: 6.5px; font-weight: 700; letter-spacing: 0.08em; color: #fbf3e1; }
.rd-veg { border: 1px solid #1c6b68; padding: 0 3px; font-size: 7px; font-weight: 700; color: #1c6b68; }
.rd-dots { flex: 1; min-width: 10px; border-bottom: 2px dotted rgba(165, 47, 34, 0.3); transform: translateY(-3px); }
.rd-price { flex-shrink: 0; font-family: ui-monospace, monospace; font-size: 14px; font-weight: 800; color: #a52f22; }
.rd-price::before { content: '$'; font-size: 9px; margin-right: 1px; opacity: 0.7; }
.rd-call {
  display: flex;
  align-items: center;
  gap: 10px;
  border: 2px solid #241b14;
  border-radius: 9999px;
  background: #1c6b68;
  padding: 8px 14px 8px 8px;
  color: #fbf3e1;
  box-shadow: 3px 3px 0 #241b14;
  transition: transform 0.15s ease, box-shadow 0.15s ease;
}
.rd-call:hover { transform: translate(-1px, -1px); box-shadow: 5px 5px 0 #241b14; }
.rd-call:active { transform: translate(2px, 2px); box-shadow: 1px 1px 0 #241b14; }
.rd-call-ico { display: inline-flex; align-items: center; justify-content: center; height: 30px; width: 30px; border-radius: 9999px; background: #f2c14e; color: #241b14; }
.rd-rise { animation: rd-in 0.65s cubic-bezier(0.2, 0.9, 0.25, 1) both; }
.rd-d1 { animation-delay: 0.1s; }
.rd-d2 { animation-delay: 0.3s; }
@keyframes rd-in {
  from { opacity: 0; transform: translateY(14px) scale(0.99); }
  to { opacity: 1; transform: none; }
}
@media (prefers-reduced-motion: reduce) {
  .rd-rise, .rd-call { animation: none !important; transition: none !important; }
}
</style>
