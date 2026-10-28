<template>
  <div class="relative min-h-screen overflow-hidden bg-[#0c1b14] text-[#f0e3c6]">
    <div class="ct-arc pointer-events-none absolute inset-x-0 top-0 h-72"></div>
    <div class="ct-grain pointer-events-none absolute inset-0"></div>

    <div class="relative z-10 mx-auto w-full max-w-md px-6 pb-14 pt-9">
      <div class="ct-rise flex items-center justify-between font-mono text-[8px] uppercase tracking-[0.28em] text-[#9e9371]">
        <span>The Amber List</span>
        <span v-if="nfcId">No. {{ nfcId }}</span>
      </div>

      <header class="ct-rise ct-d1 mt-7 text-center">
        <div class="ct-fan mx-auto">
          <span v-for="n in 7" :key="n" :style="{ transform: 'rotate(' + (n - 4) * 12 + 'deg)' }"></span>
        </div>
        <h1 class="ct-serif mt-4 text-[42px] font-medium leading-[1.05] tracking-[0.03em] text-[#f5e7c6]">
          {{ payload?.shopName || '琥珀酒单' }}
        </h1>
        <p v-if="payload?.tagline" class="mx-auto mt-2 max-w-[18rem] text-[10.5px] leading-[1.8] text-[#b0a582]">{{ payload.tagline }}</p>
        <div class="mt-4 flex items-center justify-center gap-3 font-mono text-[8px] tracking-[0.16em] text-[#9e9371]">
          <span v-if="payload?.hours">{{ payload.hours }}</span>
          <span v-if="payload?.hours" class="h-1 w-1 rounded-full bg-[#c9a84c]"></span>
          <span>SPIRITS · WINE · NIGHT</span>
        </div>
      </header>

      <div class="ct-rise ct-d2 mt-6 flex items-center gap-2">
        <span class="h-px flex-1 bg-[#c9a84c]/40"></span>
        <span class="ct-diamond"></span>
        <span class="h-px flex-1 bg-[#c9a84c]/40"></span>
      </div>

      <section v-if="groups.length" class="ct-ticket mt-5">
        <span class="ct-perf ct-perf-l"></span>
        <span class="ct-perf ct-perf-r"></span>
        <div class="ct-stub">
          <p class="font-mono text-[6.5px] uppercase tracking-[0.16em] text-[#9e9371]">Tab</p>
          <p class="ct-serif mt-1 text-[20px] leading-none text-[#e2cd9a]">{{ gi2 }}</p>
          <p class="mt-1 font-mono text-[6px] tracking-[0.1em] text-[#7d7459]">NIGHT<br />LIST</p>
        </div>
        <div class="ct-body">
          <div v-for="(group, gi) in groups" :key="gi" class="ct-rise ct-group" :style="{ animationDelay: 0.12 + gi * 0.06 + 's' }">
            <div class="flex items-baseline gap-2">
              <span class="font-mono text-[7px] text-[#c9a84c]">{{ pad(gi + 1) }}</span>
              <h2 class="ct-serif text-[14px] tracking-[0.14em] text-[#e2cd9a]">{{ group.name }}</h2>
              <span class="h-px flex-1 bg-[#c9a84c]/20"></span>
            </div>
            <ul class="mt-1">
              <li v-for="(item, i) in group.items" :key="i" class="ct-row">
                <div class="min-w-0 flex-1">
                  <div class="flex flex-wrap items-baseline gap-x-2">
                    <span class="ct-serif text-[13.5px] text-[#f0e3c6]">{{ item?.name }}</span>
                    <span v-if="signature(item)" class="ct-pick">SIGNATURE</span>
                    <span v-if="veg(item)" class="ct-veg">V</span>
                    <span v-if="spicy(item)" class="inline-flex text-[#cf8a54]"><Flame v-for="n in spicy(item)" :key="n" class="h-2.5 w-2.5" /></span>
                  </div>
                  <p v-if="item?.desc" class="mt-0.5 max-w-[15rem] text-[9px] leading-[1.55] text-[#9e9371]">{{ item.desc }}</p>
                </div>
                <span class="ct-dots"></span>
                <span v-if="hasPrice(item)" class="flex-shrink-0 font-mono text-[11.5px] tabular-nums text-[#d8bd80]">¥{{ item.price }}</span>
                <span v-else class="font-mono text-[9px] text-[#7d7459]">—</span>
              </li>
            </ul>
          </div>
        </div>
      </section>

      <div v-if="!groups.length" class="ct-rise mt-5 border border-[#c9a84c]/30 px-6 py-12 text-center">
        <Wine class="mx-auto h-6 w-6 text-[#9e9371]" />
        <p class="ct-serif mt-3 text-[12px] italic text-[#b0a582]">酒杯已擦亮，酒单正在写。</p>
      </div>

      <div v-if="payload?.notice" class="ct-rise mt-5 text-center">
        <p class="text-[9.5px] leading-[1.8] text-[#9e9371]">{{ payload.notice }}</p>
      </div>

      <div class="ct-rise ct-d3 mt-7">
        <a v-if="payload?.phone" :href="'tel:' + payload.phone" class="ct-call">
          <span class="ct-call-line"></span>
          <span class="font-mono text-[7px] uppercase tracking-[0.24em] text-[#9e9371]">Reserve a stool</span>
          <span class="ct-serif text-[16px] tracking-[0.06em] text-[#e8d09a]">{{ payload.phone }}</span>
          <span class="ct-call-line"></span>
        </a>
      </div>

      <div v-if="payload?.address" class="mt-4 flex items-start justify-center gap-1.5 text-center text-[9px] leading-relaxed text-[#8b8267]">
        <MapPin class="mt-0.5 h-3 w-3 flex-shrink-0" />
        <span>{{ payload.address }}</span>
      </div>

      <p class="mt-7 text-center font-mono text-[6.5px] uppercase tracking-[0.4em] text-[#6d654e]">Stirred · Shaken · Poured</p>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { MapPin, Flame, Wine } from 'lucide-vue-next'

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
const gi2 = computed(() => pad(groups.value.length || 0))
</script>

<style scoped>
.ct-arc { background: radial-gradient(60% 100% at 50% 0%, rgba(201, 168, 76, 0.14), transparent 68%); }
.ct-grain {
  opacity: 0.35;
  background-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' width='130' height='130'%3E%3Cfilter id='n'%3E%3CfeTurbulence type='fractalNoise' baseFrequency='0.9' numOctaves='3'/%3E%3C/filter%3E%3Crect width='130' height='130' filter='url(%23n)' opacity='0.07'/%3E%3C/svg%3E");
}
.ct-serif { font-family: 'Didot', 'Bodoni MT', 'Songti SC', Georgia, serif; }
.ct-fan { display: flex; align-items: flex-end; justify-content: center; height: 22px; gap: 1px; }
.ct-fan span { display: block; height: 18px; width: 1px; margin-left: -1px; transform-origin: bottom center; background: linear-gradient(180deg, #c9a84c, transparent); opacity: 0.65; }
.ct-diamond { height: 7px; width: 7px; transform: rotate(45deg); border: 1px solid #c9a84c; }
.ct-ticket {
  position: relative;
  display: flex;
  border: 1px solid rgba(201, 168, 76, 0.45);
  border-radius: 4px;
  background: linear-gradient(160deg, rgba(21, 43, 32, 0.92), rgba(10, 25, 18, 0.95));
  box-shadow: 0 24px 50px -34px rgba(0, 0, 0, 0.9);
}
.ct-stub {
  display: flex;
  width: 54px;
  flex-shrink: 0;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  border-right: 1px dashed rgba(201, 168, 76, 0.4);
  padding: 10px 4px;
  text-align: center;
}
.ct-body { min-width: 0; flex: 1; padding: 12px 14px; }
.ct-perf { position: absolute; top: 50%; height: 12px; width: 12px; transform: translateY(-50%); border-radius: 9999px; background: #0c1b14; }
.ct-perf-l { left: -7px; border: 1px solid rgba(201, 168, 76, 0.45); }
.ct-perf-r { right: -7px; border: 1px solid rgba(201, 168, 76, 0.45); }
.ct-group { padding-top: 10px; }
.ct-row { display: flex; align-items: baseline; gap: 7px; padding: 6px 0; border-bottom: 1px solid rgba(201, 168, 76, 0.1); }
.ct-row:last-child { border-bottom: none; }
.ct-dots { flex: 1; min-width: 8px; border-bottom: 1px dotted rgba(201, 168, 76, 0.28); transform: translateY(-2px); }
.ct-pick { border: 1px solid rgba(201, 168, 76, 0.5); padding: 0 3px; font-family: ui-monospace, monospace; font-size: 5.5px; letter-spacing: 0.1em; color: #d8bd80; }
.ct-veg { border: 1px solid rgba(150, 170, 120, 0.45); padding: 0 3px; font-size: 6px; color: #b3c49a; }
.ct-call { display: flex; flex-direction: column; align-items: center; gap: 3px; }
.ct-call-line { height: 1px; width: 44px; background: #c9a84c; opacity: 0.45; }
.ct-rise { animation: ct-in 0.75s cubic-bezier(0.2, 0.9, 0.25, 1) both; }
.ct-d1 { animation-delay: 0.08s; }
.ct-d2 { animation-delay: 0.16s; }
.ct-d3 { animation-delay: 0.34s; }
@keyframes ct-in {
  from { opacity: 0; transform: translateY(14px); }
  to { opacity: 1; transform: none; }
}
@media (prefers-reduced-motion: reduce) {
  .ct-rise { animation: none !important; }
}
</style>
