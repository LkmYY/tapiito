<template>
  <div class="relative min-h-screen overflow-hidden bg-[#f3efe3] text-[#1c1a17]">
    <div class="iw-paper pointer-events-none absolute inset-0"></div>
    <svg class="pointer-events-none absolute -right-10 top-10 h-56 w-56 opacity-[0.09]" viewBox="0 0 200 200" fill="none" aria-hidden="true">
      <circle cx="100" cy="100" r="70" stroke="#1c1a17" stroke-width="16" opacity="0.5" />
      <circle cx="112" cy="92" r="68" stroke="#1c1a17" stroke-width="4" />
      <path d="M30 140c40-10 70-6 100-26s40-34 46-58" stroke="#1c1a17" stroke-width="8" stroke-linecap="round" opacity="0.5" />
    </svg>

    <div class="relative z-10 mx-auto w-full max-w-md px-7 pb-14 pt-9">
      <div class="iw-rise flex items-center justify-between font-mono text-[8px] uppercase tracking-[0.24em] text-[#8a8274]">
        <span>Ink &amp; Rice · 食单</span>
        <span v-if="nfcId">NFC {{ nfcId }}</span>
      </div>

      <header class="iw-rise iw-d1 mt-8 flex items-start justify-between gap-5">
        <div class="min-w-0 flex-1">
          <p class="font-mono text-[8px] uppercase tracking-[0.3em] text-[#a8452c]">今日 · 手书</p>
          <h1 class="iw-brush mt-3 text-[43px] font-bold leading-[1.04] tracking-[0.04em] text-[#1c1a17]">{{ payload?.shopName || '水墨食单' }}</h1>
          <p v-if="payload?.tagline" class="iw-brush mt-3 max-w-[16rem] text-[15px] leading-[1.7] text-[#5d574c]">{{ payload.tagline }}</p>
          <span class="iw-seal mt-4">食<br />单</span>
        </div>
        <p class="iw-vertical flex-shrink-0 text-[11px] tracking-[0.3em] text-[#8a8274]">一箸一世界</p>
      </header>

      <div class="iw-rise iw-d2 mt-7 flex items-center gap-3">
        <span class="iw-stroke"></span>
        <span class="iw-brush text-[12px] tracking-[0.2em] text-[#a8452c]">菜 单</span>
        <span class="iw-stroke flex-1"></span>
      </div>

      <section v-if="groups.length" class="mt-2">
        <div v-for="(group, gi) in groups" :key="gi" class="iw-rise iw-group" :style="{ animationDelay: 0.14 + gi * 0.07 + 's' }">
          <div class="flex items-end justify-between border-b border-[#1c1a17]/20 pb-1">
            <h2 class="iw-brush text-[20px] font-semibold tracking-[0.12em] text-[#1c1a17]">{{ group.name }}</h2>
            <span class="mb-0.5 font-mono text-[7.5px] tracking-[0.16em] text-[#a8452c]">{{ pad(gi + 1) }} / {{ pad(group.items.length) }}</span>
          </div>
          <ul class="mt-1">
            <li v-for="(item, i) in group.items" :key="i" class="iw-row">
              <div class="min-w-0 flex-1">
                <div class="flex flex-wrap items-baseline gap-x-2">
                  <span class="iw-brush text-[16px] text-[#26231e]">{{ item?.name }}</span>
                  <span v-if="signature(item)" class="iw-tag">招牌</span>
                  <span v-if="veg(item)" class="iw-tag iw-tag-g">素</span>
                  <span v-if="spicy(item)" class="inline-flex text-[#a8452c]"><Flame v-for="n in spicy(item)" :key="n" class="h-2.5 w-2.5" /></span>
                </div>
                <p v-if="item?.desc" class="mt-0.5 max-w-[16rem] text-[9.5px] leading-[1.65] text-[#7b7365]">{{ item.desc }}</p>
              </div>
              <span class="iw-dots"></span>
              <span v-if="hasPrice(item)" class="iw-price">{{ item.price }}</span>
              <span v-else class="font-mono text-[9px] text-[#a39a8a]">—</span>
            </li>
          </ul>
        </div>
      </section>

      <div v-else class="iw-rise mt-6 border border-[#1c1a17]/20 px-6 py-12 text-center">
        <Brush class="mx-auto h-5 w-5 text-[#8a8274]" />
        <p class="iw-brush mt-3 text-[13px] text-[#7b7365]">砚台已磨，菜单待书。</p>
      </div>

      <div v-if="payload?.notice" class="iw-rise mt-6 flex items-start gap-2.5">
        <span class="mt-1 h-1.5 w-1.5 flex-shrink-0 rounded-full bg-[#a8452c]"></span>
        <p class="iw-brush text-[13px] leading-[1.7] text-[#6b6457]">{{ payload.notice }}</p>
      </div>

      <div class="iw-rise iw-d2 mt-7 flex items-end justify-between gap-4 border-t border-[#1c1a17]/15 pt-4">
        <div>
          <p class="font-mono text-[7px] uppercase tracking-[0.24em] text-[#a39a8a]">订座 · Calligraphy</p>
          <a v-if="payload?.phone" :href="'tel:' + payload.phone" class="iw-brush mt-1 inline-flex items-center gap-2 text-[17px] text-[#1c1a17] hover:text-[#a8452c]">
            <Phone class="h-3.5 w-3.5" />{{ payload.phone }}
          </a>
        </div>
        <div class="text-right font-mono text-[7.5px] leading-[1.7] text-[#8a8274]">
          <p v-if="payload?.hours">{{ payload.hours }}</p>
          <p v-if="payload?.address" class="max-w-[9rem]">{{ payload.address }}</p>
        </div>
      </div>

      <p class="iw-brush mt-8 text-center text-[12px] tracking-[0.3em] text-[#a39a8a]">— 食无定味，适口者珍 —</p>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { Flame, Phone, Brush } from 'lucide-vue-next'

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
.iw-paper {
  opacity: 0.6;
  background-image: repeating-linear-gradient(5deg, rgba(100, 85, 60, 0.035) 0 1px, transparent 1px 5px),
    radial-gradient(rgba(100, 85, 60, 0.08) 0.6px, transparent 0.6px);
  background-size: auto, 8px 8px;
}
.iw-brush { font-family: 'STKaiti', 'Kaiti SC', 'KaiTi', 'Songti SC', serif; }
.iw-vertical { writing-mode: vertical-rl; text-orientation: upright; }
.iw-seal {
  display: inline-flex;
  height: 44px;
  width: 26px;
  align-items: center;
  justify-content: center;
  border: 1.5px solid #a8452c;
  border-radius: 2px;
  text-align: center;
  font-family: 'STKaiti', 'KaiTi', serif;
  font-size: 11px;
  line-height: 1.4;
  color: #a8452c;
  transform: rotate(-3deg);
}
.iw-stroke {
  display: block;
  height: 2px;
  width: 28px;
  background: linear-gradient(90deg, transparent, #1c1a17 20%, #1c1a17 80%, transparent);
  opacity: 0.5;
}
.iw-group { padding-top: 22px; }
.iw-row { display: flex; align-items: baseline; gap: 8px; padding: 9px 0; border-bottom: 1px dotted rgba(28, 26, 23, 0.22); }
.iw-row:last-child { border-bottom: none; }
.iw-dots { flex: 1; min-width: 10px; border-bottom: 1px dotted rgba(28, 26, 23, 0.28); transform: translateY(-2px); }
.iw-price { flex-shrink: 0; font-family: 'STKaiti', 'KaiTi', serif; font-size: 16px; color: #a8452c; }
.iw-price::before { content: '¥'; margin-right: 1px; font-size: 9px; opacity: 0.7; }
.iw-tag { border: 1px solid rgba(168, 69, 44, 0.5); padding: 0 3px; font-size: 7px; letter-spacing: 0.06em; color: #a8452c; }
.iw-tag-g { border-color: rgba(90, 110, 70, 0.5); color: #5a6e46; }
.iw-rise { animation: iw-in 0.75s cubic-bezier(0.2, 0.9, 0.25, 1) both; }
.iw-d1 { animation-delay: 0.1s; }
.iw-d2 { animation-delay: 0.18s; }
@keyframes iw-in {
  from { opacity: 0; transform: translateY(14px); }
  to { opacity: 1; transform: none; }
}
@media (prefers-reduced-motion: reduce) {
  .iw-rise { animation: none !important; }
  a { transition: none !important; }
}
</style>
