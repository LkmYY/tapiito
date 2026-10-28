<template>
  <div class="relative min-h-screen overflow-hidden bg-[#fdfbf5] text-[#243027]">
    <div class="pointer-events-none absolute inset-x-0 top-0 h-72 bg-[radial-gradient(75%_100%_at_50%_0%,rgba(234,179,8,0.14),transparent_70%)]"></div>
    <div class="bf-dots pointer-events-none absolute inset-0"></div>

    <div class="relative z-10 mx-auto w-full max-w-md px-6 pb-12 pt-9">
      <div class="rise flex items-center justify-between font-mono text-[10px] uppercase tracking-[0.22em] text-[#a8a08c]">
        <span class="flex items-center gap-1.5">
          <Sun class="h-3 w-3 text-[#ca8a04]" />
          Morning Menu
        </span>
        <span v-if="nfcId">NFC · {{ nfcId }}</span>
      </div>

      <header class="rise rise-1 mt-7">
        <h1 class="text-balance text-[33px] font-extrabold leading-[1.1] tracking-tight text-[#243027]">
          {{ payload?.shopName || '早餐铺' }}
        </h1>
        <p v-if="payload?.tagline" class="mt-2.5 text-[13px] leading-relaxed text-[#7d8471]">{{ payload.tagline }}</p>
      </header>

      <!-- service note -->
      <div v-if="payload?.hours" class="rise rise-2 mt-5 inline-flex items-center gap-2 rounded-lg border border-[#e6e0cb] bg-white px-3 py-2">
        <Clock class="h-3.5 w-3.5 text-[#ca8a04]" />
        <span class="text-[12px] font-medium tabular-nums tracking-wide text-[#5d6553]">{{ payload.hours }}</span>
      </div>

      <section v-if="groups.length" class="mt-8 space-y-8">
        <div v-for="(group, gi) in groups" :key="gi" class="rise" :class="riseDelay(gi)">
          <div class="flex items-center gap-2">
            <h2 class="rounded-md bg-[#243027] px-2.5 py-1 text-[11px] font-bold tracking-[0.12em] text-[#fdfbf5]">{{ group.name }}</h2>
            <span class="h-px flex-1 bg-[#e6e0cb]"></span>
          </div>

          <ul class="mt-2">
            <li v-for="(item, i) in group.items" :key="i" class="bf-item">
              <div class="min-w-0 flex-1">
                <div class="flex flex-wrap items-center gap-x-2 gap-y-1">
                  <span class="text-[15px] font-semibold text-[#243027]">{{ item?.name || '菜品' }}</span>
                  <span v-if="isSignature(item)" class="bf-tag">招牌</span>
                  <span v-if="isVegetarian(item)" class="bf-tag bf-tag-green">素</span>
                  <span v-if="spicyLevel(item)" class="inline-flex items-center gap-[1px] text-[#dc2626]">
                    <Flame v-for="n in spicyLevel(item)" :key="n" class="h-3 w-3" />
                  </span>
                </div>
                <p v-if="item?.desc" class="mt-0.5 text-[11.5px] leading-relaxed text-[#8b9280]">{{ item.desc }}</p>
              </div>
              <span v-if="hasPrice(item)" class="flex-shrink-0 text-[15px] font-bold tabular-nums text-[#243027]">
                ¥{{ item.price }}
              </span>
            </li>
          </ul>
        </div>
      </section>

      <div v-else class="rise mt-9 rounded-xl border border-dashed border-[#ddd8c4] bg-white px-6 py-11 text-center">
        <Coffee class="mx-auto h-6 w-6 text-[#c4bfa9]" />
        <p class="mt-3 text-[13px] text-[#9a9c88]">早点正在准备，请稍候</p>
      </div>

      <div v-if="payload?.notice" class="rise mt-7 rounded-lg bg-[#f5f1e3] px-4 py-3">
        <p class="text-[11.5px] leading-relaxed text-[#7d8471]">{{ payload.notice }}</p>
      </div>

      <div v-if="payload?.address" class="rise mt-5 flex items-start gap-2.5">
        <MapPin class="mt-0.5 h-3.5 w-3.5 flex-shrink-0 text-[#a8a08c]" />
        <p class="text-[11.5px] leading-relaxed text-[#8b9280]">{{ payload.address }}</p>
      </div>

      <div v-if="payload?.phone" class="rise rise-3 mt-7">
        <a :href="'tel:' + payload.phone" class="bf-cta">
          <Phone class="h-4 w-4" />
          电话预订
          <span class="ml-auto font-mono text-[11px] font-normal tracking-[0.1em] tabular-nums text-white/70">{{ payload.phone }}</span>
        </a>
      </div>

      <p class="mt-8 text-center font-mono text-[9px] uppercase tracking-[0.28em] text-[#bfbba6]">Served Fresh Daily</p>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { Sun, Clock, MapPin, Phone, Flame, Coffee } from 'lucide-vue-next'

const props = defineProps(['payload', 'nfcId'])

const hasPrice = (item) => item?.price !== undefined && item?.price !== null && item?.price !== ''
const riseDelay = (i) => (i === 0 ? 'rise-1' : i === 1 ? 'rise-2' : 'rise-3')

const isSignature = (item) => {
  if (item?.signature === true) return true
  const tags = Array.isArray(item?.tags) ? item.tags : []
  return tags.some((t) => String(t).includes('招牌') || String(t).toLowerCase() === 'signature')
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
.bf-dots {
  background-image: radial-gradient(rgba(36, 48, 39, 0.055) 1px, transparent 1px);
  background-size: 18px 18px;
}

.bf-item {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 13px 2px;
  border-bottom: 1px solid #efeade;
}
.bf-item:last-child {
  border-bottom: none;
}

.bf-tag {
  border-radius: 3px;
  background: rgba(202, 138, 4, 0.14);
  padding: 1px 5px;
  font-size: 9px;
  font-weight: 700;
  letter-spacing: 0.06em;
  color: #a16207;
}
.bf-tag-green {
  background: rgba(22, 101, 52, 0.1);
  color: #15803d;
}

.bf-cta {
  display: flex;
  align-items: center;
  gap: 9px;
  width: 100%;
  border-radius: 11px;
  background: #243027;
  padding: 14px 18px;
  font-size: 14px;
  font-weight: 600;
  letter-spacing: 0.04em;
  color: #fdfbf5;
  box-shadow: 0 18px 36px -22px rgba(36, 48, 39, 0.7);
  transition: background-color 0.2s ease, transform 0.15s ease;
}
.bf-cta:hover {
  background: #36453a;
}
.bf-cta:active {
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
