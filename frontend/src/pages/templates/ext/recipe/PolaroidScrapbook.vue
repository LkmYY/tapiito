<template>
  <div class="relative min-h-screen overflow-hidden bg-[#efe6d3] text-[#33291d]">
    <div class="ps-dots pointer-events-none absolute inset-0"></div>
    <div class="pointer-events-none absolute inset-x-0 top-0 h-56 bg-[radial-gradient(60%_100%_at_50%_0%,rgba(255,255,255,0.55),transparent_70%)]"></div>

    <div class="relative z-10 mx-auto w-full max-w-md px-5 pb-14 pt-8">
      <div class="ps-rise flex items-center justify-between font-mono text-[8px] uppercase tracking-[0.2em] text-[#8d7b60]">
        <span>My Kitchen Album</span>
        <span v-if="nfcId">No. {{ nfcId }}</span>
      </div>

      <header class="ps-rise ps-d1 relative mt-6 text-center">
        <span class="ps-tape mx-auto"></span>
        <h1 class="ps-hand mt-3 text-[39px] font-bold leading-[1.08] text-[#33291d]">{{ payload?.shopName || '拍立得食记' }}</h1>
        <p v-if="payload?.tagline" class="ps-hand mx-auto mt-1 max-w-[19rem] text-[14px] leading-[1.6] text-[#6f5d45]">{{ payload.tagline }}</p>
        <div class="mt-3 flex flex-wrap items-center justify-center gap-x-3 gap-y-1 font-mono text-[8px] uppercase tracking-[0.1em] text-[#8d7b60]">
          <span v-if="payload?.hours">营业 {{ payload.hours }}</span>
          <span v-if="payload?.hours" class="h-1 w-1 rounded-full bg-[#b98a4e]"></span>
          <span>胶卷 · 2026</span>
        </div>
      </header>

      <section v-if="groups.length" class="mt-5 space-y-6">
        <div v-for="(group, gi) in groups" :key="gi" class="ps-rise" :style="{ animationDelay: 0.12 + gi * 0.07 + 's' }">
          <div class="flex items-center gap-2">
            <span class="ps-label">{{ group.name }}</span>
            <span class="ps-line"></span>
            <span class="font-mono text-[8px] text-[#a29074]">{{ pad(group.items.length) }} SHOTS</span>
          </div>

          <div class="mt-2 grid grid-cols-2 gap-x-2 gap-y-3">
            <article v-for="(item, i) in group.items" :key="i" class="ps-polaroid" :class="'ps-tilt-' + ((i + gi) % 4)">
              <div class="ps-photo" :class="'ps-film-' + ((i + gi) % 4)">
                <svg viewBox="0 0 90 60" fill="none" aria-hidden="true">
                  <ellipse cx="45" cy="38" rx="26" ry="8" fill="rgba(255,255,255,0.28)" stroke="#fff" stroke-width="0.8" />
                  <path d="M30 34c0-7 6-12 15-12s15 5 15 12" fill="rgba(255,255,255,0.2)" stroke="#fff" stroke-width="0.8" />
                  <path d="M38 26c0-4 3-4 3-7M48 25c0-4 3-4 3-7M56 27c0-3 2-3 2-5" stroke="#fff" stroke-width="0.8" stroke-linecap="round" opacity="0.8" />
                  <circle cx="66" cy="16" r="5" fill="rgba(255,255,255,0.3)" />
                </svg>
                <span class="ps-frame-no">{{ pad(gi + 1) }}-{{ pad(i + 1) }}</span>
                <span v-if="signature(item)" class="ps-fav">♥</span>
              </div>
              <h3 class="ps-hand mt-1.5 truncate text-[16px] font-bold leading-tight text-[#33291d]">{{ item?.name }}</h3>
              <p v-if="item?.desc" class="mt-0.5 line-clamp-2 min-h-[25px] text-[9px] leading-[1.4] text-[#7d6a50]">{{ item.desc }}</p>
              <div class="mt-1 flex items-center justify-between border-t border-dashed border-[#d8c7a9] pt-1">
                <div class="flex items-center gap-1">
                  <span v-if="veg(item)" class="text-[8px] text-[#5f7a4a]">素</span>
                  <span v-if="spicy(item)" class="inline-flex text-[#b65b36]"><Flame v-for="n in spicy(item)" :key="n" class="h-2.5 w-2.5" /></span>
                </div>
                <span v-if="hasPrice(item)" class="ps-price">¥{{ item.price }}</span>
                <span v-else class="font-mono text-[8px] text-[#a29074]">—</span>
              </div>
            </article>
          </div>
        </div>
      </section>

      <div v-else class="ps-rise mt-5 rotate-[-1deg] border border-[#d8c7a9] bg-[#fbf6e9] px-6 py-10 text-center shadow-[2px_3px_0_rgba(141,123,96,0.2)]">
        <Camera class="mx-auto h-6 w-6 text-[#a29074]" />
        <p class="ps-hand mt-3 text-[14px] text-[#7d6a50]">相册还空着，等一张好菜的照片。</p>
      </div>

      <div v-if="payload?.notice" class="ps-rise mt-5 border-l-2 border-[#b98a4e]/60 pl-3">
        <p class="ps-hand text-[13px] leading-[1.6] text-[#6f5d45]">{{ payload.notice }}</p>
      </div>

      <div class="ps-rise ps-d2 mt-6 flex items-end justify-between gap-3">
        <a v-if="payload?.phone" :href="'tel:' + payload.phone" class="ps-call">
          <Phone class="h-3.5 w-3.5" />
          <span class="text-[11px] font-semibold">打电话订一份</span>
          <span class="ml-auto font-mono text-[9px] text-[#8d7b60]">{{ payload.phone }}</span>
        </a>
      </div>

      <div v-if="payload?.address" class="mt-3 flex items-start gap-1.5 text-[9px] leading-relaxed text-[#8d7b60]">
        <MapPin class="mt-0.5 h-3 w-3 flex-shrink-0" />
        <span>{{ payload.address }}</span>
      </div>

      <p class="ps-hand mt-7 text-center text-[12px] text-[#a29074]">— 好好吃饭，慢慢记录 —</p>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { Phone, MapPin, Flame, Camera } from 'lucide-vue-next'

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
.ps-dots {
  background-image: radial-gradient(rgba(120, 95, 55, 0.16) 1px, transparent 1px);
  background-size: 17px 17px;
}
.ps-hand { font-family: 'Kaiti SC', 'STKaiti', 'KaiTi', 'Songti SC', serif; }
.ps-tape {
  display: block;
  height: 22px;
  width: 92px;
  background: rgba(201, 172, 124, 0.5);
  transform: rotate(-2deg);
  box-shadow: inset 0 0 0 1px rgba(255, 255, 255, 0.2);
}
.ps-label {
  display: inline-block;
  background: #b98a4e;
  padding: 2px 8px;
  font-family: 'Kaiti SC', 'STKaiti', serif;
  font-size: 14px;
  font-weight: 700;
  color: #fbf6e9;
  transform: rotate(-1deg);
}
.ps-line { flex: 1; height: 1px; background: repeating-linear-gradient(90deg, #b8a583 0 5px, transparent 5px 9px); }
.ps-polaroid {
  border-radius: 2px;
  background: #fffdf6;
  padding: 5px 5px 6px;
  box-shadow: 1px 3px 8px -4px rgba(80, 60, 30, 0.45);
  transition: transform 0.18s ease, box-shadow 0.18s ease;
}
.ps-polaroid:hover { transform: translateY(-3px) rotate(0deg) !important; box-shadow: 1px 8px 14px -6px rgba(80, 60, 30, 0.5); }
.ps-tilt-0 { transform: rotate(-1.6deg); }
.ps-tilt-1 { transform: rotate(1.2deg); }
.ps-tilt-2 { transform: rotate(-0.7deg); }
.ps-tilt-3 { transform: rotate(1.8deg); }
.ps-photo { position: relative; overflow: hidden; border-radius: 1px; height: 78px; }
.ps-photo svg { position: absolute; inset: 0; height: 100%; width: 100%; }
.ps-film-0 { background: linear-gradient(140deg, #b98a4e, #6f4a28); }
.ps-film-1 { background: linear-gradient(140deg, #7c8d68, #3f4d36); }
.ps-film-2 { background: linear-gradient(140deg, #c27a5a, #7d3f2b); }
.ps-film-3 { background: linear-gradient(140deg, #8a8fa0, #4a4d5c); }
.ps-frame-no { position: absolute; right: 3px; bottom: 2px; font-family: ui-monospace, monospace; font-size: 5.5px; letter-spacing: 0.08em; color: rgba(255, 255, 255, 0.8); }
.ps-fav { position: absolute; left: 4px; top: 2px; font-size: 13px; color: #ffe8b8; text-shadow: 0 1px 1px rgba(0, 0, 0, 0.3); }
.ps-price { font-family: ui-monospace, monospace; font-size: 11px; font-weight: 700; color: #a8512f; }
.ps-call {
  display: flex;
  flex: 1;
  align-items: center;
  gap: 8px;
  border-radius: 9999px;
  background: #fbf6e9;
  padding: 9px 13px;
  color: #4a3b28;
  box-shadow: 0 10px 20px -18px rgba(80, 60, 30, 0.8);
  transition: background-color 0.2s ease, transform 0.15s ease;
}
.ps-call:hover { background: #fffdf6; }
.ps-call:active { transform: scale(0.985); }
.ps-rise { animation: ps-in 0.7s cubic-bezier(0.2, 0.9, 0.25, 1) both; }
.ps-d1 { animation-delay: 0.08s; }
.ps-d2 { animation-delay: 0.32s; }
@keyframes ps-in {
  from { opacity: 0; transform: translateY(14px) rotate(0.4deg); }
  to { opacity: 1; transform: none; }
}
@media (prefers-reduced-motion: reduce) {
  .ps-rise, .ps-polaroid, .ps-call { animation: none !important; transition: none !important; }
  .ps-polaroid:hover { transform: none !important; }
}
</style>
