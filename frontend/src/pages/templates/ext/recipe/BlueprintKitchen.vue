<template>
  <div class="relative min-h-screen overflow-hidden bg-[#0a2a45] text-[#e7f4ff]">
    <div class="bp-grid pointer-events-none absolute inset-0"></div>
    <div class="bp-vignette pointer-events-none absolute inset-0"></div>

    <div class="relative z-10 mx-auto w-full max-w-md px-5 pb-14 pt-8">
      <div class="bp-rise flex items-start justify-between font-mono text-[8px] uppercase tracking-[0.2em] text-[#8ec5e8]">
        <div>
          <p>Kitchen Drawing</p>
          <p class="mt-0.5 text-[#5f96bb]">DWG NO. RCP-025</p>
        </div>
        <div class="text-right">
          <p>SCALE 1:1</p>
          <p class="mt-0.5 text-[#5f96bb]">REV. 03</p>
        </div>
      </div>

      <div class="bp-rise bp-d1 mt-4 border border-[#e7f4ff]/60 p-4">
        <div class="flex items-start justify-between gap-4">
          <div class="min-w-0 flex-1">
            <p class="font-mono text-[7.5px] uppercase tracking-[0.3em] text-[#7fb9de]">Plan View · Menu Board</p>
            <h1 class="mt-2 text-[31px] font-semibold leading-[1.05] tracking-[-0.01em] text-[#f2faff]">{{ payload?.shopName }}</h1>
            <p v-if="payload?.tagline" class="mt-2 max-w-[16rem] text-[10px] leading-[1.7] text-[#a9d2eb]">{{ payload.tagline }}</p>
          </div>
          <div class="bp-plate flex-shrink-0">
            <svg viewBox="0 0 72 72" fill="none" aria-hidden="true">
              <circle cx="36" cy="36" r="26" stroke="currentColor" stroke-width="0.8" stroke-dasharray="2 3" />
              <circle cx="36" cy="36" r="19" stroke="currentColor" stroke-width="0.8" />
              <circle cx="36" cy="36" r="9" stroke="currentColor" stroke-width="0.8" stroke-dasharray="1 2" />
              <path d="M36 4v10M36 58v10M4 36h10M58 36h10" stroke="currentColor" stroke-width="0.8" />
              <path d="M21 22c5 4 10 6 15 6s10-2 15-6" stroke="currentColor" stroke-width="0.6" opacity="0.7" />
            </svg>
          </div>
        </div>

        <div class="bp-dim mt-4 flex items-center gap-1 font-mono text-[6.5px] tracking-[0.12em] text-[#7fb9de]">
          <span class="h-2 w-px bg-current"></span>
          <span class="h-px flex-1 bg-current opacity-60"></span>
          <span>240 mm · COUNTER</span>
          <span class="h-px flex-1 bg-current opacity-60"></span>
          <span class="h-2 w-px bg-current"></span>
        </div>
      </div>

      <div class="bp-rise bp-d2 mt-3 grid grid-cols-3 gap-2 font-mono text-[7px] uppercase tracking-[0.1em]">
        <div class="bp-meta"><span class="text-[#6ba6cb]">Hours</span><span class="mt-1 block text-[9px] text-[#d8ecf8]">{{ payload?.hours || '—' }}</span></div>
        <div class="bp-meta col-span-2"><span class="text-[#6ba6cb]">Address</span><span class="mt-1 block truncate text-[9px] text-[#d8ecf8]">{{ payload?.address || '—' }}</span></div>
      </div>

      <div class="bp-rise bp-d2 mt-5 flex items-center gap-2 border-b border-[#e7f4ff]/50 pb-1.5">
        <span class="font-mono text-[8px] uppercase tracking-[0.26em] text-[#8ec5e8]">Schedule of Dishes</span>
        <span class="h-px flex-1 bg-[#e7f4ff]/20"></span>
        <span class="font-mono text-[8px] text-[#7fb9de]">{{ pad(total) }} POS.</span>
      </div>

      <section v-if="groups.length" class="mt-2">
        <div v-for="(group, gi) in groups" :key="gi" class="bp-rise bp-group" :style="{ animationDelay: 0.16 + gi * 0.06 + 's' }">
          <div class="flex items-baseline gap-2">
            <span class="bp-node"></span>
            <h2 class="font-mono text-[10px] font-semibold uppercase tracking-[0.16em] text-[#9fd3f0]">{{ group.name }}</h2>
            <span class="h-px flex-1 bg-[#e7f4ff]/15"></span>
            <span class="font-mono text-[7px] text-[#6ba6cb]">{{ pad(gi + 1) }}-{{ pad(group.items.length) }}</span>
          </div>
          <ul class="mt-1 border-l border-dashed border-[#7fb9de]/35 pl-3">
            <li v-for="(item, i) in group.items" :key="i" class="bp-row">
              <span class="font-mono text-[7px] text-[#5f96bb]">{{ pad(i + 1) }}</span>
              <div class="min-w-0 flex-1">
                <div class="flex flex-wrap items-baseline gap-x-2">
                  <span class="text-[12.5px] font-medium text-[#eaf6ff]">{{ item?.name }}</span>
                  <span v-if="signature(item)" class="bp-tag">KEY</span>
                  <span v-if="veg(item)" class="bp-tag bp-tag-green">VEG</span>
                  <span v-if="spicy(item)" class="font-mono text-[7px] text-[#e6b27b]">S{{ spicy(item) }}</span>
                </div>
                <p v-if="item?.desc" class="mt-0.5 max-w-[16rem] text-[9px] leading-[1.55] text-[#9fc6dd]">{{ item.desc }}</p>
              </div>
              <span class="bp-leader"></span>
              <span v-if="hasPrice(item)" class="flex-shrink-0 border border-[#7fb9de]/40 px-1.5 py-0.5 font-mono text-[9.5px] tabular-nums text-[#d8ecf8]">¥{{ item.price }}</span>
              <span v-else class="font-mono text-[9px] text-[#5f96bb]">—</span>
            </li>
          </ul>
        </div>
      </section>

      <div v-else class="bp-rise mt-5 border border-dashed border-[#7fb9de]/50 px-6 py-12 text-center">
        <Ruler class="mx-auto h-6 w-6 text-[#7fb9de]" />
        <p class="mt-3 font-mono text-[9px] uppercase tracking-[0.16em] text-[#9fc6dd]">Drawing pending · 菜单待绘制</p>
      </div>

      <div v-if="payload?.notice" class="bp-rise mt-5 border-l-2 border-[#e6b27b] bg-[#0d3554]/80 px-3 py-2.5">
        <p class="font-mono text-[8px] leading-[1.7] text-[#a9d2eb]">{{ payload.notice }}</p>
      </div>

      <div class="bp-rise bp-d3 mt-5 flex items-center justify-between gap-4 border-t border-[#e7f4ff]/40 pt-3">
        <a v-if="payload?.phone" :href="'tel:' + payload.phone" class="bp-call">
          <Phone class="h-3.5 w-3.5" />
          <span class="font-mono text-[8px] uppercase tracking-[0.16em]">Telephone</span>
          <span class="font-mono text-[10px] text-[#d8ecf8]">{{ payload.phone }}</span>
        </a>
        <span class="font-mono text-[6.5px] uppercase tracking-[0.14em] text-[#5f96bb]">Checked by<br />Chef · {{ nfcId || '—' }}</span>
      </div>

      <p class="mt-7 text-center font-mono text-[6.5px] uppercase tracking-[0.3em] text-[#4c7d9e]">Do not scale this drawing</p>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { Phone, Ruler } from 'lucide-vue-next'

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
.bp-grid {
  background-image: linear-gradient(rgba(231, 244, 255, 0.075) 1px, transparent 1px), linear-gradient(90deg, rgba(231, 244, 255, 0.075) 1px, transparent 1px),
    linear-gradient(rgba(231, 244, 255, 0.04) 1px, transparent 1px), linear-gradient(90deg, rgba(231, 244, 255, 0.04) 1px, transparent 1px);
  background-size: 44px 44px, 44px 44px, 11px 11px, 11px 11px;
}
.bp-vignette { box-shadow: inset 0 0 120px 20px rgba(2, 16, 28, 0.5); }
.bp-plate { height: 66px; width: 66px; color: #9fd3f0; opacity: 0.75; }
.bp-meta { border: 1px solid rgba(127, 185, 222, 0.3); background: rgba(8, 34, 54, 0.6); padding: 6px 7px; }
.bp-node { height: 7px; width: 7px; flex-shrink: 0; border: 1px solid #9fd3f0; background: #0a2a45; transform: rotate(45deg); }
.bp-group { padding-top: 16px; }
.bp-row { display: flex; align-items: baseline; gap: 7px; padding: 8px 0; border-bottom: 1px solid rgba(231, 244, 255, 0.1); }
.bp-row:last-child { border-bottom: none; }
.bp-leader { flex: 1; min-width: 8px; border-bottom: 1px dotted rgba(127, 185, 222, 0.35); transform: translateY(-2px); }
.bp-tag { border: 1px solid rgba(159, 211, 240, 0.5); padding: 0 3px; font-family: ui-monospace, monospace; font-size: 5.5px; letter-spacing: 0.08em; color: #9fd3f0; }
.bp-tag-green { border-color: rgba(150, 190, 140, 0.5); color: #b4d3a4; }
.bp-call { display: flex; align-items: center; gap: 7px; transition: color 0.2s ease; }
.bp-call:hover { color: #e6b27b; }
.bp-rise { animation: bp-in 0.7s cubic-bezier(0.2, 0.9, 0.25, 1) both; }
.bp-d1 { animation-delay: 0.08s; }
.bp-d2 { animation-delay: 0.15s; }
.bp-d3 { animation-delay: 0.32s; }
@keyframes bp-in {
  from { opacity: 0; transform: translateY(12px); }
  to { opacity: 1; transform: none; }
}
@media (prefers-reduced-motion: reduce) {
  .bp-rise { animation: none !important; }
  .bp-call { transition: none !important; }
}
</style>
