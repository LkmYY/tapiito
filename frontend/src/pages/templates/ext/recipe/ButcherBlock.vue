<template>
  <div class="relative min-h-screen overflow-hidden bg-[#1e130b] text-[#f4e6cf]">
    <div class="bb-wood pointer-events-none absolute inset-0"></div>
    <div class="bb-vignette pointer-events-none absolute inset-0"></div>

    <div class="relative z-10 mx-auto w-full max-w-md px-5 pb-14 pt-8">
      <header class="bb-rise text-center">
        <div class="bb-brand">HAND CUT · DAILY</div>
        <h1 class="bb-slab mt-3 text-[38px] font-black uppercase leading-[0.98] tracking-tight text-[#f7e8cd] drop-shadow-[0_2px_0_rgba(0,0,0,0.45)]">
          {{ payload?.shopName || '木砧板食堂' }}
        </h1>
        <p v-if="payload?.tagline" class="mx-auto mt-2 max-w-[19rem] text-[11.5px] leading-relaxed text-[#c9ab84]">{{ payload.tagline }}</p>
        <div class="mt-4 flex items-center justify-center gap-4 font-mono text-[8.5px] uppercase tracking-[0.16em] text-[#a8885f]">
          <span v-if="payload?.hours" class="inline-flex items-center gap-1"><Clock class="h-3 w-3" />{{ payload.hours }}</span>
          <span v-if="nfcId">No. {{ nfcId }}</span>
        </div>
      </header>

      <div class="bb-rise bb-d1 mt-5 flex items-center gap-2">
        <span class="h-px flex-1 bg-[#7a5734]/70"></span>
        <UtensilsCrossed class="h-3.5 w-3.5 text-[#c89556]" />
        <span class="h-px flex-1 bg-[#7a5734]/70"></span>
      </div>

      <section v-if="groups.length" class="mt-5 space-y-7">
        <div v-for="(group, gi) in groups" :key="gi" class="bb-rise" :style="{ animationDelay: 0.08 + gi * 0.07 + 's' }">
          <div class="flex items-baseline justify-between">
            <h2 class="bb-slab text-[19px] font-black uppercase tracking-[0.06em] text-[#e9c98f]">{{ group.name }}</h2>
            <span class="font-mono text-[8px] uppercase tracking-[0.16em] text-[#a8885f]">{{ pad(gi + 1) }} / {{ pad(group.items.length) }}</span>
          </div>

          <div class="mt-2 grid grid-cols-2 gap-2.5">
            <article v-for="(item, i) in group.items" :key="i" class="bb-tag">
              <span class="bb-string"></span>
              <span class="bb-hole"></span>
              <div class="flex items-start justify-between gap-1">
                <span class="font-mono text-[7px] uppercase tracking-[0.12em] text-[#9a7a54]">CUT {{ pad(i + 1) }}</span>
                <span v-if="signature(item)" class="bb-pick">招牌</span>
              </div>
              <h3 class="bb-slab mt-1 text-[14.5px] font-bold leading-tight text-[#362314]">{{ item?.name }}</h3>
              <p v-if="item?.desc" class="mt-1 line-clamp-2 min-h-[27px] text-[9.5px] leading-[1.45] text-[#7b5d3f]">{{ item.desc }}</p>
              <div class="mt-2 flex items-end justify-between gap-1 border-t border-dashed border-[#c7aa82] pt-1.5">
                <div class="flex items-center gap-1">
                  <Flame v-for="n in spicy(item)" :key="n" class="h-2.5 w-2.5 text-[#b4472b]" />
                  <Leaf v-if="veg(item)" class="h-3 w-3 text-[#5f7545]" />
                </div>
                <span v-if="hasPrice(item)" class="bb-price">¥{{ item.price }}</span>
                <span v-else class="font-mono text-[9px] text-[#a88c69]">—</span>
              </div>
            </article>
          </div>
        </div>
      </section>

      <div v-else class="bb-rise mt-6 border border-dashed border-[#8d6741] bg-[#2a1a0e]/70 px-6 py-12 text-center">
        <UtensilsCrossed class="mx-auto h-6 w-6 text-[#a8885f]" />
        <p class="mt-3 text-[12px] text-[#c9ab84]">砧板已擦净，菜单正在备料。</p>
      </div>

      <div v-if="payload?.notice" class="bb-rise mt-6 border-l-4 border-[#c89556] bg-[#2a1a0e]/60 px-3.5 py-3">
        <p class="text-[10.5px] leading-relaxed text-[#c9ab84]">{{ payload.notice }}</p>
      </div>

      <div class="bb-rise bb-d2 mt-6 flex items-center justify-between gap-3">
        <a v-if="payload?.phone" :href="'tel:' + payload.phone" class="bb-call">
          <Phone class="h-4 w-4" />
          <span class="bb-slab text-[14px] font-bold">留座电话</span>
          <span class="ml-auto font-mono text-[10px]">{{ payload.phone }}</span>
        </a>
      </div>

      <div v-if="payload?.address" class="mt-4 flex items-start justify-center gap-1.5 text-center text-[9.5px] leading-relaxed text-[#a8885f]">
        <MapPin class="mt-0.5 h-3 w-3 flex-shrink-0" />
        <span>{{ payload.address }}</span>
      </div>

      <p class="mt-7 text-center font-mono text-[7.5px] uppercase tracking-[0.3em] text-[#7f603f]">Slow Food · Sharp Knives</p>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { Clock, MapPin, Phone, Flame, Leaf, UtensilsCrossed } from 'lucide-vue-next'

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
.bb-wood {
  background-color: #2b1a0d;
  background-image:
    repeating-linear-gradient(93deg, rgba(255, 214, 160, 0.045) 0 2px, transparent 2px 9px),
    repeating-linear-gradient(88deg, rgba(0, 0, 0, 0.22) 0 1px, transparent 1px 26px),
    repeating-linear-gradient(90deg, #2b1a0d 0 72px, #352112 72px 150px, #2a190c 150px 240px);
}
.bb-vignette { box-shadow: inset 0 0 120px 30px rgba(0, 0, 0, 0.55); }
.bb-brand {
  display: inline-block;
  border: 1px solid #a8885f;
  padding: 3px 10px;
  font-family: ui-monospace, monospace;
  font-size: 8px;
  letter-spacing: 0.34em;
  color: #d8b785;
}
.bb-slab { font-family: 'Rockwell', 'Roboto Slab', 'Songti SC', Georgia, serif; }
.bb-tag {
  position: relative;
  border-radius: 2px;
  background: #f5e8d0;
  padding: 12px 10px 9px;
  box-shadow: 0 8px 14px -10px rgba(0, 0, 0, 0.8), inset 0 -2px 0 rgba(154, 122, 84, 0.18);
  transition: transform 0.18s ease, box-shadow 0.18s ease;
}
.bb-tag:hover { transform: translateY(-3px) rotate(-0.6deg); box-shadow: 0 14px 20px -12px rgba(0, 0, 0, 0.9); }
.bb-string {
  position: absolute;
  top: -6px;
  left: 50%;
  height: 7px;
  width: 1px;
  background: #c9ab84;
}
.bb-hole {
  position: absolute;
  top: 3px;
  left: 50%;
  height: 7px;
  width: 7px;
  transform: translateX(-50%);
  border-radius: 9999px;
  background: #2b1a0d;
  box-shadow: inset 0 1px 2px #000, 0 0 0 1.5px #d8bd94;
}
.bb-pick {
  border: 1px solid #a8452c;
  padding: 0 3px;
  font-size: 7px;
  font-weight: 700;
  color: #a8452c;
  transform: rotate(2deg);
}
.bb-price { font-family: ui-monospace, monospace; font-size: 14px; font-weight: 800; color: #8f3f24; }
.bb-call {
  display: flex;
  align-items: center;
  gap: 9px;
  width: 100%;
  border: 1px solid #c89556;
  border-radius: 2px;
  background: #3a2412;
  padding: 12px 14px;
  color: #f1ddbe;
  transition: background-color 0.2s ease, transform 0.15s ease;
}
.bb-call:hover { background: #4a2f18; }
.bb-call:active { transform: scale(0.985); }
.bb-rise { animation: bb-in 0.65s cubic-bezier(0.2, 0.9, 0.25, 1) both; }
.bb-d1 { animation-delay: 0.08s; }
.bb-d2 { animation-delay: 0.32s; }
@keyframes bb-in {
  from { opacity: 0; transform: translateY(12px); }
  to { opacity: 1; transform: none; }
}
@media (prefers-reduced-motion: reduce) {
  .bb-rise, .bb-tag, .bb-call { animation: none !important; transition: none !important; }
  .bb-tag:hover { transform: none; }
}
</style>
