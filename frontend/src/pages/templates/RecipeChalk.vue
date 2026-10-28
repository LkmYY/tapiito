<template>
  <div class="relative min-h-screen overflow-hidden bg-[#16201c] text-[#e8e4d8]">
    <!-- board texture -->
    <div class="chalk-dust pointer-events-none absolute inset-0"></div>
    <div class="pointer-events-none absolute inset-x-0 top-0 h-80 bg-[radial-gradient(70%_100%_at_50%_0%,rgba(232,228,216,0.07),transparent_72%)]"></div>
    <!-- wooden frame -->
    <div class="pointer-events-none absolute inset-3 rounded-sm border border-[#3a2b1e]/70"></div>

    <div class="relative z-10 mx-auto w-full max-w-md px-7 pb-12 pt-10">
      <div class="rise flex items-center justify-between font-mono text-[10px] uppercase tracking-[0.26em] text-[#9aa89c]">
        <span>Today's Board</span>
        <span v-if="nfcId">NFC · {{ nfcId }}</span>
      </div>

      <header class="rise rise-1 mt-6 text-center">
        <p class="chalk-line mx-auto w-10"></p>
        <h1 class="chalk-title mt-4 text-balance text-[32px] leading-[1.15] tracking-wide text-[#f2eee2]">
          {{ payload?.shopName || '今日菜单' }}
        </h1>
        <p v-if="payload?.tagline" class="chalk-sub mt-3 text-[12.5px] leading-relaxed text-[#b8c2b6]">
          {{ payload.tagline }}
        </p>
        <p class="chalk-line mx-auto mt-5 w-16"></p>
      </header>

      <!-- hours -->
      <div v-if="payload?.hours || payload?.address" class="rise rise-2 mt-7 space-y-1.5 text-center font-mono text-[11px] tracking-[0.12em] text-[#9aa89c]">
        <p v-if="payload?.hours" class="tabular-nums">{{ payload.hours }}</p>
        <p v-if="payload?.address" class="text-[10.5px] leading-relaxed">{{ payload.address }}</p>
      </div>

      <!-- menu -->
      <section v-if="groups.length" class="mt-9 space-y-8">
        <div v-for="(group, gi) in groups" :key="gi" class="rise" :class="riseDelay(gi)">
          <div class="flex items-center gap-3">
            <h2 class="chalk-head text-[14px] tracking-[0.14em] text-[#f0d9a8]">{{ group.name }}</h2>
            <span class="chalk-dash flex-1"></span>
          </div>

          <ul class="mt-2">
            <li v-for="(item, i) in group.items" :key="i" class="ch-item">
              <div class="min-w-0 flex-1">
                <div class="flex flex-wrap items-baseline gap-x-2 gap-y-1">
                  <span class="text-[15.5px] font-medium text-[#eee9db]">{{ item?.name || '菜品' }}</span>
                  <span v-if="isSignature(item)" class="ch-tag">招牌</span>
                  <span v-if="spicyLevel(item)" class="inline-flex items-center gap-[1px] text-[#e08a5a]">
                    <Flame v-for="n in spicyLevel(item)" :key="n" class="h-3 w-3" />
                  </span>
                  <span v-if="isVegetarian(item)" class="ch-tag ch-tag-green">素</span>
                </div>
                <p v-if="item?.desc" class="mt-0.5 text-[11.5px] leading-relaxed text-[#9aa89c]">{{ item.desc }}</p>
              </div>
              <span class="ch-dots"></span>
              <span v-if="hasPrice(item)" class="flex-shrink-0 text-[15.5px] font-semibold tabular-nums text-[#f0d9a8]">
                {{ item.price }}
              </span>
            </li>
          </ul>
        </div>
      </section>

      <div v-else class="rise mt-10 rounded-lg border border-dashed border-[#4a5a50] px-6 py-12 text-center">
        <UtensilsCrossed class="mx-auto h-6 w-6 text-[#6b7d70]" />
        <p class="mt-3 text-[13px] text-[#9aa89c]">菜单更新中，请稍后再来</p>
      </div>

      <div v-if="payload?.notice" class="rise mt-8 text-center">
        <p class="font-mono text-[10.5px] leading-relaxed tracking-[0.08em] text-[#8fa094]">{{ payload.notice }}</p>
      </div>

      <div v-if="payload?.phone" class="rise rise-3 mt-8">
        <a :href="'tel:' + payload.phone" class="ch-cta">
          <Phone class="h-4 w-4" />
          <span>电话订座</span>
          <span class="ml-auto font-mono text-[11px] tracking-[0.12em] tabular-nums text-[#16201c]/60">{{ payload.phone }}</span>
        </a>
      </div>

      <p class="mt-9 text-center font-mono text-[9px] uppercase tracking-[0.3em] text-[#5d6b62]">Chalk &amp; Charcoal</p>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { Phone, Flame, UtensilsCrossed } from 'lucide-vue-next'

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
.chalk-dust {
  background-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' width='160' height='160'%3E%3Cfilter id='n'%3E%3CfeTurbulence type='fractalNoise' baseFrequency='0.8' numOctaves='4'/%3E%3C/filter%3E%3Crect width='160' height='160' filter='url(%23n)' opacity='0.055'/%3E%3C/svg%3E");
}

.chalk-title,
.chalk-head,
.chalk-sub {
  font-family: 'Geist', 'PingFang SC', 'Microsoft YaHei', sans-serif;
  font-weight: 500;
}
.chalk-title {
  text-shadow: 0 0 1px rgba(242, 238, 226, 0.5), 0 1px 0 rgba(0, 0, 0, 0.25);
  letter-spacing: 0.06em;
}
.chalk-head {
  text-shadow: 0 0 1px rgba(240, 217, 168, 0.55);
}

.chalk-line {
  display: block;
  height: 1px;
  background: linear-gradient(90deg, transparent, rgba(232, 228, 216, 0.45), transparent);
}
.chalk-dash {
  height: 1px;
  background: repeating-linear-gradient(90deg, rgba(232, 228, 216, 0.28) 0 6px, transparent 6px 11px);
}

.ch-item {
  display: flex;
  align-items: baseline;
  gap: 8px;
  padding: 11px 0;
  border-bottom: 1px dashed rgba(154, 168, 156, 0.14);
}
.ch-item:last-child {
  border-bottom: none;
}

.ch-dots {
  flex: 1;
  min-width: 12px;
  border-bottom: 1px dotted rgba(232, 228, 216, 0.22);
  transform: translateY(-2px);
}

.ch-tag {
  border: 1px solid rgba(240, 217, 168, 0.45);
  border-radius: 3px;
  padding: 0 4px;
  font-size: 8.5px;
  font-weight: 600;
  letter-spacing: 0.1em;
  color: #f0d9a8;
  transform: translateY(-1px);
}
.ch-tag-green {
  border-color: rgba(150, 200, 160, 0.4);
  color: #a8d5b0;
}

.ch-cta {
  display: flex;
  align-items: center;
  gap: 9px;
  width: 100%;
  border-radius: 10px;
  background: #e8e4d8;
  padding: 13px 18px;
  font-size: 14px;
  font-weight: 600;
  letter-spacing: 0.04em;
  color: #16201c;
  box-shadow: 0 16px 34px -22px rgba(0, 0, 0, 0.8);
  transition: background-color 0.2s ease, transform 0.15s ease;
}
.ch-cta:hover {
  background: #f6f3ea;
}
.ch-cta:active {
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
