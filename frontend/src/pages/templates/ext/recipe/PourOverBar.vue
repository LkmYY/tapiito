<template>
  <div class="relative min-h-screen overflow-hidden bg-[#1b120d] text-[#efe0cd]">
    <div class="pb-glow pointer-events-none absolute inset-x-0 top-0 h-80"></div>
    <div class="pb-grain pointer-events-none absolute inset-0"></div>

    <div class="relative z-10 mx-auto w-full max-w-md px-6 pb-14 pt-8">
      <div class="pb-rise flex items-center justify-between font-mono text-[8.5px] uppercase tracking-[0.26em] text-[#a08670]">
        <span>Brew Bar · 手冲吧台</span>
        <span v-if="nfcId">NFC {{ nfcId }}</span>
      </div>

      <header class="pb-rise pb-d1 mt-7 flex items-start justify-between gap-4">
        <div class="min-w-0 flex-1">
          <div class="flex items-center gap-2">
            <span class="pb-marker"></span>
            <span class="font-mono text-[8px] uppercase tracking-[0.22em] text-[#c08457]">Single Origin</span>
          </div>
          <h1 class="pb-serif mt-3 text-[42px] font-medium leading-[1.02] tracking-[-0.02em] text-[#f8ecdc]">
            {{ payload?.shopName || '手冲吧台' }}
          </h1>
          <p v-if="payload?.tagline" class="mt-2 max-w-[17rem] text-[11px] leading-relaxed text-[#b39b83]">{{ payload.tagline }}</p>
        </div>
        <div class="pb-cup">
          <svg viewBox="0 0 48 48" fill="none" aria-hidden="true">
            <path d="M10 18h24v10a12 12 0 0 1-24 0V18Z" stroke="currentColor" stroke-width="1.4" />
            <path d="M34 21h4a5 5 0 0 1 0 10h-4" stroke="currentColor" stroke-width="1.4" />
            <path d="M16 12c0-2 2-2 2-4M22 12c0-2 2-2 2-4M28 12c0-2 2-2 2-4" stroke="currentColor" stroke-width="1.2" stroke-linecap="round" />
          </svg>
        </div>
      </header>

      <div class="pb-rise pb-d2 mt-6 grid grid-cols-3 gap-1.5">
        <div class="pb-stat"><span class="pb-stat-k">DOSE</span><span class="pb-stat-v">15<small>g</small></span></div>
        <div class="pb-stat"><span class="pb-stat-k">WATER</span><span class="pb-stat-v">225<small>g</small></span></div>
        <div class="pb-stat"><span class="pb-stat-k">TIME</span><span class="pb-stat-v">2:30</span></div>
      </div>
      <p class="pb-rise pb-d2 mt-2 text-center font-mono text-[8px] uppercase tracking-[0.2em] text-[#7d6a58]">92°C · 中细研磨 · 三段注水</p>

      <div class="pb-rise pb-d3 mt-7 flex items-center gap-3 border-b border-[#4a3527] pb-2">
        <h2 class="font-mono text-[9px] uppercase tracking-[0.28em] text-[#c08457]">Today's Bar</h2>
        <span class="h-px flex-1 bg-[#4a3527]"></span>
        <span class="font-mono text-[8px] text-[#8b7764]">{{ total }} 项</span>
      </div>

      <section v-if="groups.length" class="mt-1">
        <div v-for="(group, gi) in groups" :key="gi" class="pb-rise pb-group" :style="{ animationDelay: 0.16 + gi * 0.06 + 's' }">
          <div class="flex items-baseline justify-between">
            <h3 class="pb-serif text-[17px] text-[#f0dfc9]">{{ group.name }}</h3>
            <span class="font-mono text-[8px] tracking-[0.1em] text-[#8b7764]">{{ pad(gi + 1) }}</span>
          </div>
          <ul class="mt-1">
            <li v-for="(item, i) in group.items" :key="i" class="pb-row">
              <div class="min-w-0 flex-1">
                <div class="flex flex-wrap items-center gap-x-2 gap-y-0.5">
                  <span class="text-[14px] text-[#efe0cd]">{{ item?.name }}</span>
                  <span v-if="signature(item)" class="pb-tag">BAR PICK</span>
                  <span v-if="veg(item)" class="pb-tag pb-tag-green">PLANT</span>
                  <span v-if="spicy(item)" class="inline-flex text-[#c08457]"><Flame v-for="n in spicy(item)" :key="n" class="h-2.5 w-2.5" /></span>
                </div>
                <p v-if="item?.desc" class="mt-0.5 max-w-[16rem] text-[10px] leading-[1.55] text-[#9d8872]">{{ item.desc }}</p>
              </div>
              <div class="pb-leader"></div>
              <span v-if="hasPrice(item)" class="flex-shrink-0 font-mono text-[12px] tabular-nums text-[#e0b27f]">¥{{ item.price }}</span>
              <span v-else class="font-mono text-[9px] text-[#6f5d4c]">—</span>
            </li>
          </ul>
        </div>
      </section>

      <div v-else class="pb-rise mt-5 border border-dashed border-[#4a3527] px-6 py-12 text-center">
        <Coffee class="mx-auto h-6 w-6 text-[#8b7764]" />
        <p class="mt-3 text-[11px] text-[#a08670]">豆子还在养，菜单稍后更新。</p>
      </div>

      <div v-if="payload?.notice" class="pb-rise mt-5 border-l border-[#c08457]/60 pl-3">
        <p class="text-[10px] leading-[1.7] text-[#9d8872]">{{ payload.notice }}</p>
      </div>

      <div class="pb-rise pb-d3 mt-6 rounded-lg border border-[#4a3527] bg-[#241811] p-4">
        <div class="grid grid-cols-2 gap-x-4 gap-y-3 font-mono text-[8.5px] text-[#a08670]">
          <div>
            <p class="uppercase tracking-[0.18em] text-[#6f5d4c]">Hours</p>
            <p class="mt-1 text-[#d8c3ad]">{{ payload?.hours || '—' }}</p>
          </div>
          <div>
            <p class="uppercase tracking-[0.18em] text-[#6f5d4c]">Address</p>
            <p class="mt-1 leading-relaxed text-[#d8c3ad]">{{ payload?.address || '—' }}</p>
          </div>
        </div>
        <a v-if="payload?.phone" :href="'tel:' + payload.phone" class="pb-call mt-3">
          <Phone class="h-3.5 w-3.5" />
          <span class="font-mono text-[9px] tracking-[0.12em]">CALL THE BAR</span>
          <span class="ml-auto font-mono text-[10px] text-[#e0b27f]">{{ payload.phone }}</span>
        </a>
      </div>

      <p class="mt-7 text-center font-mono text-[7px] uppercase tracking-[0.32em] text-[#6f5d4c]">Grind · Bloom · Pour · Enjoy</p>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { Phone, Flame, Coffee } from 'lucide-vue-next'

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
.pb-glow { background: radial-gradient(60% 90% at 70% 0%, rgba(192, 132, 87, 0.22), transparent 70%); }
.pb-grain {
  opacity: 0.4;
  background-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' width='130' height='130'%3E%3Cfilter id='n'%3E%3CfeTurbulence type='fractalNoise' baseFrequency='0.9' numOctaves='3'/%3E%3C/filter%3E%3Crect width='130' height='130' filter='url(%23n)' opacity='0.07'/%3E%3C/svg%3E");
}
.pb-serif { font-family: 'Iowan Old Style', 'Palatino Linotype', 'Songti SC', Georgia, serif; }
.pb-marker { height: 7px; width: 7px; border-radius: 9999px; border: 1.5px solid #c08457; }
.pb-cup { height: 46px; width: 46px; flex-shrink: 0; color: #c08457; opacity: 0.9; }
.pb-stat {
  border: 1px solid #4a3527;
  border-radius: 6px;
  background: #241811;
  padding: 7px 8px;
}
.pb-stat-k { display: block; font-family: ui-monospace, monospace; font-size: 6.5px; letter-spacing: 0.2em; color: #7d6a58; }
.pb-stat-v { display: block; margin-top: 2px; font-family: ui-monospace, monospace; font-size: 17px; color: #e0b27f; }
.pb-stat-v small { margin-left: 1px; font-size: 8px; color: #8b7764; }
.pb-group { padding-top: 18px; }
.pb-row {
  display: flex;
  align-items: baseline;
  gap: 8px;
  padding: 9px 0;
  border-bottom: 1px solid #34251b;
}
.pb-row:last-child { border-bottom: none; }
.pb-leader { flex: 1; min-width: 10px; border-bottom: 1px dotted #4a3527; transform: translateY(-2px); }
.pb-tag { border: 1px solid rgba(192, 132, 87, 0.45); padding: 0 3px; font-family: ui-monospace, monospace; font-size: 6px; letter-spacing: 0.08em; color: #d8a878; }
.pb-tag-green { border-color: rgba(150, 170, 120, 0.4); color: #b3c49a; }
.pb-call {
  display: flex;
  align-items: center;
  gap: 8px;
  border-top: 1px solid #4a3527;
  padding-top: 10px;
  color: #d8c3ad;
  transition: color 0.2s ease;
}
.pb-call:hover { color: #e0b27f; }
.pb-rise { animation: pb-in 0.7s cubic-bezier(0.2, 0.9, 0.25, 1) both; }
.pb-d1 { animation-delay: 0.08s; }
.pb-d2 { animation-delay: 0.15s; }
.pb-d3 { animation-delay: 0.3s; }
@keyframes pb-in {
  from { opacity: 0; transform: translateY(12px); }
  to { opacity: 1; transform: none; }
}
@media (prefers-reduced-motion: reduce) {
  .pb-rise { animation: none !important; }
  .pb-call { transition: none !important; }
}
</style>
