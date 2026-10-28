<template>
  <div class="relative min-h-screen overflow-hidden bg-[#f7f7f4] text-[#101010]">
    <div class="sm-grid pointer-events-none absolute inset-0"></div>
    <div class="relative z-10 mx-auto w-full max-w-md px-6 pb-12 pt-7">
      <div class="sm-rise grid grid-cols-2 border-t-2 border-[#101010] font-mono text-[8px] uppercase tracking-[0.18em]">
        <div class="border-r border-[#c9c9c2] py-2">Menu / Speisekarte</div>
        <div class="py-2 pl-3 text-right">Zürich · 2026</div>
      </div>

      <header class="sm-rise sm-d1 mt-10">
        <div class="flex items-start justify-between gap-4">
          <div class="min-w-0 flex-1">
            <p class="font-mono text-[9px] uppercase tracking-[0.3em] text-[#e5322d]">01 — Kitchen</p>
            <h1 class="sm-display mt-3 text-[46px] font-bold leading-[0.92] tracking-[-0.045em]">
              {{ payload?.shopName }}
            </h1>
          </div>
          <span class="sm-dot mt-1"></span>
        </div>
        <p v-if="payload?.tagline" class="mt-5 max-w-[20rem] text-[12px] leading-[1.7] text-[#5c5c58]">
          {{ payload.tagline }}
        </p>
      </header>

      <div class="sm-rise sm-d2 mt-8 grid grid-cols-2 gap-x-4 border-y border-[#101010] py-3">
        <div>
          <p class="font-mono text-[7.5px] uppercase tracking-[0.2em] text-[#8a8a84]">Hours</p>
          <p class="mt-1 font-mono text-[11px] tabular-nums">{{ payload?.hours || '—' }}</p>
        </div>
        <div>
          <p class="font-mono text-[7.5px] uppercase tracking-[0.2em] text-[#8a8a84]">Address</p>
          <p class="mt-1 text-[10.5px] leading-snug">{{ payload?.address || '—' }}</p>
        </div>
      </div>

      <section v-if="groups.length" class="mt-2">
        <div v-for="(group, gi) in groups" :key="gi" class="sm-rise sm-group" :style="{ animationDelay: 0.16 + gi * 0.06 + 's' }">
          <div class="grid grid-cols-[34px_1fr] gap-x-3">
            <span class="pt-0.5 font-mono text-[10px] text-[#e5322d]">{{ pad(gi + 1) }}</span>
            <div>
              <h2 class="border-b border-[#101010] pb-1.5 text-[14px] font-bold uppercase tracking-[0.06em]">{{ group.name }}</h2>
              <ul>
                <li v-for="(item, i) in group.items" :key="i" class="sm-row">
                  <div class="min-w-0 flex-1">
                    <div class="flex items-baseline gap-2">
                      <span class="text-[14.5px] font-medium tracking-[-0.01em]">{{ item?.name }}</span>
                      <span v-if="signature(item)" class="sm-tag">★</span>
                      <span v-if="veg(item)" class="sm-tag sm-tag-g">V</span>
                    </div>
                    <p v-if="item?.desc" class="mt-0.5 max-w-[16rem] text-[10px] leading-[1.5] text-[#7c7c76]">{{ item.desc }}</p>
                  </div>
                  <span v-if="hasPrice(item)" class="flex-shrink-0 font-mono text-[12.5px] tabular-nums">{{ item.price }}</span>
                  <span v-else class="flex-shrink-0 font-mono text-[11px] text-[#a3a39d]">—</span>
                </li>
              </ul>
            </div>
          </div>
        </div>
      </section>

      <div v-else class="sm-rise mt-8 border border-[#101010] p-8 text-center">
        <span class="mx-auto block h-2 w-2 rounded-full bg-[#e5322d]"></span>
        <p class="mt-4 text-[11px] text-[#6d6d68]">菜单尚未配置</p>
      </div>

      <div v-if="payload?.notice" class="sm-rise mt-6 grid grid-cols-[34px_1fr] gap-x-3">
        <span class="font-mono text-[9px] text-[#e5322d]">NB</span>
        <p class="border-l border-[#c9c9c2] pl-3 text-[10px] leading-[1.6] text-[#6d6d68]">{{ payload.notice }}</p>
      </div>

      <div class="sm-rise sm-d3 mt-8 border-t-2 border-[#101010] pt-4">
        <a v-if="payload?.phone" :href="'tel:' + payload.phone" class="sm-call">
          <span class="font-mono text-[8px] uppercase tracking-[0.2em] text-[#8a8a84]">Call</span>
          <span class="ml-auto font-mono text-[14px] font-semibold tracking-[0.02em]">{{ payload.phone }}</span>
          <ArrowUpRight class="h-4 w-4" />
        </a>
        <div class="mt-4 flex items-center justify-between font-mono text-[7px] uppercase tracking-[0.22em] text-[#a3a39d]">
          <span>Print · Scan · Eat</span>
          <span v-if="nfcId">ID {{ nfcId }}</span>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { ArrowUpRight } from 'lucide-vue-next'

const props = defineProps(['payload', 'nfcId'])

const pad = (n) => String(n).padStart(2, '0')
const hasPrice = (item) => item?.price !== undefined && item?.price !== null && item?.price !== ''
const signature = (item) => item?.signature === true || (Array.isArray(item?.tags) && item.tags.some((t) => /招牌|必点|signature/i.test(String(t))))
const veg = (item) => item?.vegetarian === true || (Array.isArray(item?.tags) && item.tags.some((t) => /素|vegan|vegetarian/i.test(String(t))))

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
.sm-grid {
  background-image: linear-gradient(rgba(16, 16, 16, 0.04) 1px, transparent 1px), linear-gradient(90deg, rgba(16, 16, 16, 0.04) 1px, transparent 1px);
  background-size: 33px 33px;
  mask-image: linear-gradient(180deg, #000 0%, transparent 70%);
}
.sm-display { font-family: 'Helvetica Neue', Helvetica, Arial, 'PingFang SC', sans-serif; }
.sm-dot { height: 26px; width: 26px; flex-shrink: 0; border-radius: 9999px; background: #e5322d; }
.sm-group { padding: 18px 0 4px; }
.sm-row {
  display: flex;
  align-items: baseline;
  gap: 10px;
  padding: 10px 0;
  border-bottom: 1px solid #e2e2dc;
}
.sm-row:last-child { border-bottom: none; }
.sm-tag { font-size: 8px; color: #e5322d; }
.sm-tag-g { color: #4d6b47; }
.sm-call { display: flex; align-items: center; gap: 10px; transition: color 0.2s ease; }
.sm-call:hover { color: #e5322d; }
.sm-rise { animation: sm-in 0.6s cubic-bezier(0.2, 0.9, 0.25, 1) both; }
.sm-d1 { animation-delay: 0.08s; }
.sm-d2 { animation-delay: 0.15s; }
.sm-d3 { animation-delay: 0.32s; }
@keyframes sm-in {
  from { opacity: 0; transform: translateY(10px); }
  to { opacity: 1; transform: none; }
}
@media (prefers-reduced-motion: reduce) {
  .sm-rise { animation: none !important; }
  .sm-call { transition: none !important; }
}
</style>
