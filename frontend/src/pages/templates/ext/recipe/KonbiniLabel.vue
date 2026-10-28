<template>
  <div class="relative min-h-screen overflow-hidden bg-[#eef0ea] text-[#17201a]">
    <div class="kl-grid pointer-events-none absolute inset-0"></div>

    <div class="relative z-10 mx-auto w-full max-w-md px-5 pb-14 pt-7">
      <div class="kl-rise flex items-center justify-between font-mono text-[9px] uppercase tracking-[0.22em] text-[#63705f]">
        <span class="inline-flex items-center gap-2"><span class="kl-dot"></span>Fresh Counter</span>
        <span v-if="nfcId">NFC-{{ nfcId }}</span>
      </div>

      <header class="kl-rise kl-d1 mt-5 border border-[#1f2b22] bg-white">
        <div class="flex items-stretch">
          <div class="flex w-14 flex-col items-center justify-center gap-0.5 bg-[#1f8a4c] text-white">
            <span class="font-mono text-[18px] font-black leading-none">24</span>
            <span class="font-mono text-[7px] tracking-[0.2em]">HOUR</span>
          </div>
          <div class="min-w-0 flex-1 px-4 py-3.5">
            <div class="flex items-center justify-between gap-2">
              <span class="font-mono text-[8.5px] font-bold uppercase tracking-[0.24em] text-[#1f8a4c]">Store Kitchen</span>
              <span class="font-mono text-[8px] text-[#8a9487]">NO. 031</span>
            </div>
            <h1 class="mt-1 truncate text-[26px] font-black leading-tight tracking-tight">{{ payload?.shopName || '街角食堂' }}</h1>
            <p v-if="payload?.tagline" class="mt-1 line-clamp-2 text-[10.5px] leading-relaxed text-[#6a7567]">{{ payload.tagline }}</p>
          </div>
        </div>
        <div class="kl-barcode-row">
          <div class="kl-barcode"></div>
          <div class="text-right font-mono text-[8px] leading-tight text-[#687364]">
            <p v-if="payload?.hours" class="tabular-nums">{{ payload.hours }}</p>
            <p v-if="payload?.address" class="max-w-[9rem] truncate">{{ payload.address }}</p>
          </div>
        </div>
      </header>

      <div class="kl-rise kl-d2 mt-4 flex items-center justify-between border-b-2 border-[#17201a] pb-1.5">
        <span class="text-[11px] font-bold uppercase tracking-[0.18em]">今日上架 · Today's Menu</span>
        <span class="font-mono text-[9px] text-[#1f8a4c]">{{ total }} SKU</span>
      </div>

      <section v-if="groups.length" class="mt-3 space-y-4">
        <div v-for="(group, gi) in groups" :key="gi" class="kl-rise" :style="{ animationDelay: 0.1 + gi * 0.06 + 's' }">
          <div class="flex items-center gap-2">
            <span class="kl-cat">{{ group.name }}</span>
            <span class="h-px flex-1 bg-[#cbd2c8]"></span>
            <span class="font-mono text-[8px] text-[#8a9487]">AISLE {{ pad(gi + 1) }}</span>
          </div>

          <div class="mt-2 grid grid-cols-2 gap-2">
            <article v-for="(item, i) in group.items" :key="i" class="kl-label">
              <div class="flex items-start justify-between gap-1">
                <span class="font-mono text-[7px] uppercase tracking-[0.14em] text-[#8a9487]">ITEM {{ pad(i + 1) }}</span>
                <span v-if="signature(item)" class="kl-best">推荐</span>
              </div>
              <h3 class="mt-1 text-[14px] font-bold leading-snug">{{ item?.name }}</h3>
              <p v-if="item?.desc" class="mt-0.5 line-clamp-2 min-h-[26px] text-[9.5px] leading-[1.4] text-[#778274]">{{ item.desc }}</p>
              <div class="mt-2 flex items-end justify-between gap-1">
                <div class="flex flex-wrap items-center gap-1">
                  <span v-if="veg(item)" class="kl-badge kl-badge-green">VEG</span>
                  <span v-if="spicy(item)" class="kl-badge kl-badge-red">SPICY {{ spicy(item) }}</span>
                </div>
                <span v-if="hasPrice(item)" class="kl-price"><small>¥</small>{{ item.price }}</span>
                <span v-else class="font-mono text-[10px] text-[#a3aca0]">—</span>
              </div>
              <div class="kl-mini-barcode"></div>
            </article>
          </div>
        </div>
      </section>

      <div v-else class="kl-rise mt-4 border border-dashed border-[#b7c0b3] bg-white px-6 py-12 text-center">
        <PackageOpen class="mx-auto h-6 w-6 text-[#8a9487]" />
        <p class="mt-3 font-mono text-[10px] uppercase tracking-[0.16em] text-[#778274]">货架补货中 · Restocking</p>
      </div>

      <div v-if="payload?.notice" class="kl-rise mt-5 border-l-2 border-[#e3c229] bg-[#fbf6d9] px-3 py-2.5">
        <p class="text-[10px] leading-relaxed text-[#6f6a45]">{{ payload.notice }}</p>
      </div>

      <a v-if="payload?.phone" :href="'tel:' + payload.phone" class="kl-rise kl-d3 kl-cta mt-6">
        <Phone class="h-4 w-4" />
        <span class="font-bold tracking-wide">电话预留</span>
        <span class="ml-auto font-mono text-[10px] tracking-[0.08em]">{{ payload.phone }}</span>
      </a>

      <div class="mt-6 flex items-center justify-between font-mono text-[7.5px] uppercase tracking-[0.2em] text-[#98a294]">
        <span>Scan · Select · Heat</span>
        <span>Open 24/7</span>
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { Phone, PackageOpen } from 'lucide-vue-next'

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
.kl-grid {
  background-image: linear-gradient(rgba(31, 43, 34, 0.045) 1px, transparent 1px), linear-gradient(90deg, rgba(31, 43, 34, 0.045) 1px, transparent 1px);
  background-size: 22px 22px;
}
.kl-dot { height: 6px; width: 6px; border-radius: 9999px; background: #1f8a4c; box-shadow: 0 0 0 3px rgba(31, 138, 76, 0.15); }
.kl-barcode-row {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 12px;
  border-top: 1px dashed #c7cfc4;
  padding: 8px 14px;
  background: #fbfcf8;
}
.kl-barcode {
  height: 20px;
  width: 88px;
  background: repeating-linear-gradient(90deg, #17201a 0 2px, transparent 2px 4px, #17201a 4px 5px, transparent 5px 9px);
  opacity: 0.8;
}
.kl-cat {
  border-radius: 2px;
  background: #17201a;
  padding: 3px 7px;
  font-size: 9.5px;
  font-weight: 700;
  letter-spacing: 0.08em;
  color: #f7f8f2;
}
.kl-label {
  position: relative;
  overflow: hidden;
  border: 1px solid #c7cfc4;
  border-radius: 3px;
  background: #fff;
  padding: 9px 9px 7px;
  box-shadow: 0 1px 0 rgba(23, 32, 26, 0.06);
  transition: transform 0.18s ease, box-shadow 0.18s ease;
}
.kl-label:hover { transform: translateY(-2px); box-shadow: 0 8px 18px -12px rgba(23, 32, 26, 0.5); }
.kl-best { border: 1px solid #1f8a4c; padding: 0 3px; font-size: 7px; font-weight: 700; color: #1f8a4c; }
.kl-price {
  flex-shrink: 0;
  border-radius: 1px;
  background: #f4d91f;
  padding: 1px 6px;
  font-family: ui-monospace, monospace;
  font-size: 15px;
  font-weight: 800;
  line-height: 1.3;
  color: #17201a;
  box-shadow: 1px 1px 0 rgba(23, 32, 26, 0.18);
}
.kl-price small { margin-right: 1px; font-size: 9px; font-weight: 600; }
.kl-badge { padding: 0 3px; font-family: ui-monospace, monospace; font-size: 6.5px; font-weight: 700; letter-spacing: 0.04em; }
.kl-badge-green { background: #dff1e5; color: #1f7a48; }
.kl-badge-red { background: #fbe4dc; color: #b34025; }
.kl-mini-barcode {
  margin-top: 7px;
  height: 9px;
  background: repeating-linear-gradient(90deg, #2b352d 0 1px, transparent 1px 2px, #2b352d 2px 4px, transparent 4px 7px);
  opacity: 0.55;
}
.kl-cta {
  display: flex;
  align-items: center;
  gap: 9px;
  border-radius: 3px;
  background: #17201a;
  padding: 13px 15px;
  font-size: 13px;
  color: #f4f6ef;
  transition: background-color 0.2s ease, transform 0.15s ease;
}
.kl-cta:hover { background: #1f8a4c; }
.kl-cta:active { transform: scale(0.985); }
.kl-rise { animation: kl-in 0.6s cubic-bezier(0.2, 0.9, 0.25, 1) both; }
.kl-d1 { animation-delay: 0.07s; }
.kl-d2 { animation-delay: 0.13s; }
.kl-d3 { animation-delay: 0.34s; }
@keyframes kl-in {
  from { opacity: 0; transform: translateY(10px); }
  to { opacity: 1; transform: none; }
}
@media (prefers-reduced-motion: reduce) {
  .kl-rise, .kl-label, .kl-cta { animation: none !important; transition: none !important; }
  .kl-label:hover { transform: none; }
}
</style>
