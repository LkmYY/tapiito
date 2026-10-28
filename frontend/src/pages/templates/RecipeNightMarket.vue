<template>
  <div class="relative min-h-screen overflow-hidden bg-[#0b0908] text-[#f5efe6]">
    <div class="pointer-events-none absolute inset-x-0 top-0 h-96 bg-[radial-gradient(80%_100%_at_50%_0%,rgba(249,115,22,0.20),transparent_70%)]"></div>
    <div class="pointer-events-none absolute inset-x-0 bottom-0 h-80 bg-[radial-gradient(70%_100%_at_50%_100%,rgba(220,38,38,0.16),transparent_72%)]"></div>
    <div class="nb-grain pointer-events-none absolute inset-0"></div>

    <div class="relative z-10 mx-auto w-full max-w-md px-6 pb-12 pt-9">
      <div class="rise flex items-center justify-between font-mono text-[10px] uppercase tracking-[0.24em] text-[#8a7a68]">
        <span class="flex items-center gap-1.5">
          <span class="nb-live"></span>
          Late Night
        </span>
        <span v-if="nfcId">NFC · {{ nfcId }}</span>
      </div>

      <header class="rise rise-1 mt-7">
        <h1 class="text-balance text-[38px] font-black leading-[1.02] tracking-tight text-[#fff7ea]">
          {{ payload?.shopName || '深夜食堂' }}
        </h1>
        <p v-if="payload?.tagline" class="mt-3 text-[13px] leading-relaxed text-[#b9a892]">{{ payload.tagline }}</p>
        <div v-if="payload?.hours" class="mt-4 inline-flex items-center gap-2 rounded-full border border-[#f97316]/30 bg-[#f97316]/10 px-3 py-1.5">
          <Clock class="h-3 w-3 text-[#fb923c]" />
          <span class="font-mono text-[11px] tracking-[0.1em] tabular-nums text-[#fdba74]">{{ payload.hours }}</span>
        </div>
      </header>

      <section v-if="groups.length" class="mt-9 space-y-8">
        <div v-for="(group, gi) in groups" :key="gi" class="rise" :class="riseDelay(gi)">
          <div class="flex items-baseline gap-3">
            <h2 class="text-[13px] font-bold uppercase tracking-[0.2em] text-[#fb923c]">{{ group.name }}</h2>
            <span class="h-px flex-1 bg-gradient-to-r from-[#f97316]/35 to-transparent"></span>
            <span class="font-mono text-[10px] tabular-nums text-[#7a6a58]">{{ pad(group.items.length) }}</span>
          </div>

          <ul class="mt-1.5">
            <li v-for="(item, i) in group.items" :key="i" class="nb-item">
              <div class="min-w-0 flex-1">
                <div class="flex flex-wrap items-center gap-x-2 gap-y-1">
                  <span class="text-[16px] font-bold text-[#fff7ea]">{{ item?.name || '菜品' }}</span>
                  <span v-if="isSignature(item)" class="nb-tag">必点</span>
                  <span v-if="spicyLevel(item)" class="inline-flex items-center gap-[1px] text-[#ef4444]">
                    <Flame v-for="n in spicyLevel(item)" :key="n" class="h-3.5 w-3.5" />
                  </span>
                  <span v-if="isVegetarian(item)" class="nb-tag nb-tag-green">素</span>
                </div>
                <p v-if="item?.desc" class="mt-1 text-[11.5px] leading-relaxed text-[#9a8a78]">{{ item.desc }}</p>
              </div>
              <span v-if="hasPrice(item)" class="nb-price">
                <span class="text-[10px] font-normal opacity-70">¥</span>{{ item.price }}
              </span>
            </li>
          </ul>
        </div>
      </section>

      <div v-else class="rise mt-10 rounded-2xl border border-dashed border-[#3a2f26] bg-[#12100e] px-6 py-12 text-center">
        <Flame class="mx-auto h-6 w-6 text-[#5c4a38]" />
        <p class="mt-3 text-[13px] text-[#8a7a68]">正在备料，菜单马上就来</p>
      </div>

      <div v-if="payload?.address" class="rise mt-8 flex items-start gap-2.5">
        <MapPin class="mt-0.5 h-3.5 w-3.5 flex-shrink-0 text-[#7a6a58]" />
        <p class="text-[11.5px] leading-relaxed text-[#9a8a78]">{{ payload.address }}</p>
      </div>

      <div v-if="payload?.notice" class="rise mt-4 rounded-xl border border-[#f97316]/20 bg-[#f97316]/[0.07] px-4 py-3">
        <p class="text-[11.5px] leading-relaxed text-[#d9b28a]">{{ payload.notice }}</p>
      </div>

      <div v-if="payload?.phone" class="rise rise-3 mt-8">
        <a :href="'tel:' + payload.phone" class="nb-cta">
          <Phone class="h-4 w-4" />
          马上订位
          <span class="ml-auto font-mono text-[11px] font-normal tracking-[0.1em] tabular-nums opacity-75">{{ payload.phone }}</span>
        </a>
      </div>

      <p class="mt-9 text-center font-mono text-[9px] uppercase tracking-[0.3em] text-[#5c4a38]">Open Till Late</p>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { Clock, MapPin, Phone, Flame } from 'lucide-vue-next'

const props = defineProps(['payload', 'nfcId'])

const pad = (n) => String(n).padStart(2, '0')
const hasPrice = (item) => item?.price !== undefined && item?.price !== null && item?.price !== ''
const riseDelay = (i) => (i === 0 ? 'rise-1' : i === 1 ? 'rise-2' : 'rise-3')

const isSignature = (item) => {
  if (item?.signature === true) return true
  const tags = Array.isArray(item?.tags) ? item.tags : []
  return tags.some((t) => String(t).includes('招牌') || String(t).includes('必点') || String(t).toLowerCase() === 'signature')
}
const isVegetarian = (item) => {
  if (item?.vegetarian === true) return true
  const tags = Array.isArray(item?.tags) ? item.tags : []
  return tags.some((t) => String(t).includes('素'))
}
const spicyLevel = (item) => {
  const n = Number(item?.spicy)
  return Number.isFinite(n) ? Math.max(0, Math.min(3, Math.round(n))) : 0
}

const groups = computed(() => {
  const p = props.payload || {}
  const clean = (list) => (Array.isArray(list) ? list.filter((i) => i && (i.name || hasPrice(i))) : [])

  if (Array.isArray(p.categories) && p.categories.length) {
    return p.categories
      .filter((c) => c && (c.name || clean(c.items).length))
      .map((c) => ({ name: c.name || '其他', items: clean(c.items) }))
      .filter((c) => c.items.length)
  }

  const flat = clean(p.menu)
  if (!flat.length) return []
  const map = new Map()
  for (const item of flat) {
    const key = item.category || '菜单'
    if (!map.has(key)) map.set(key, [])
    map.get(key).push(item)
  }
  return [...map.entries()].map(([name, items]) => ({ name, items }))
})
</script>

<style scoped>
.nb-grain {
  background-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' width='140' height='140'%3E%3Cfilter id='n'%3E%3CfeTurbulence type='fractalNoise' baseFrequency='0.85' numOctaves='3'/%3E%3C/filter%3E%3Crect width='140' height='140' filter='url(%23n)' opacity='0.05'/%3E%3C/svg%3E");
}

.nb-live {
  display: inline-block;
  height: 5px;
  width: 5px;
  border-radius: 9999px;
  background: #fb923c;
  box-shadow: 0 0 0 3px rgba(251, 146, 60, 0.18);
  animation: live 2.4s ease-in-out infinite;
}
@keyframes live {
  0%,
  100% {
    opacity: 1;
    box-shadow: 0 0 0 3px rgba(251, 146, 60, 0.18);
  }
  50% {
    opacity: 0.55;
    box-shadow: 0 0 0 6px rgba(251, 146, 60, 0.06);
  }
}

.nb-item {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 14px 2px;
  border-bottom: 1px solid rgba(122, 106, 88, 0.18);
}
.nb-item:last-child {
  border-bottom: none;
}

.nb-tag {
  border-radius: 3px;
  background: #f97316;
  padding: 1px 5px;
  font-size: 9px;
  font-weight: 700;
  letter-spacing: 0.08em;
  color: #0b0908;
}
.nb-tag-green {
  background: #4ade80;
}

.nb-price {
  flex-shrink: 0;
  border-radius: 8px;
  border: 1px solid rgba(249, 115, 22, 0.32);
  background: rgba(249, 115, 22, 0.09);
  padding: 5px 10px;
  font-size: 17px;
  font-weight: 800;
  font-variant-numeric: tabular-nums;
  color: #fdba74;
}

.nb-cta {
  display: flex;
  align-items: center;
  gap: 9px;
  width: 100%;
  border-radius: 12px;
  background: linear-gradient(100deg, #f97316, #dc2626);
  padding: 14px 18px;
  font-size: 14px;
  font-weight: 700;
  letter-spacing: 0.04em;
  color: #fff7ea;
  box-shadow: 0 20px 44px -22px rgba(249, 115, 22, 0.85);
  transition: filter 0.2s ease, transform 0.15s ease;
}
.nb-cta:hover {
  filter: brightness(1.08);
}
.nb-cta:active {
  transform: scale(0.985);
}

@keyframes rise {
  from {
    opacity: 0;
    transform: translateY(14px);
  }
  to {
    opacity: 1;
    transform: none;
  }
}
.rise {
  animation: rise 0.65s cubic-bezier(0.22, 1, 0.36, 1) both;
}
.rise-1 {
  animation-delay: 0.08s;
}
.rise-2 {
  animation-delay: 0.16s;
}
.rise-3 {
  animation-delay: 0.24s;
}
</style>
