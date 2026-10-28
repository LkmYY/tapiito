<template>
  <div class="relative min-h-screen overflow-hidden bg-[#f6f4ef] text-[#22201c]">
    <div class="pointer-events-none absolute inset-x-0 top-0 h-96 bg-[radial-gradient(70%_100%_at_50%_0%,rgba(120,113,98,0.14),transparent_72%)]"></div>

    <div class="relative z-10 mx-auto w-full max-w-md px-7 pb-14 pt-11">
      <div class="rise flex items-center justify-between font-mono text-[9.5px] uppercase tracking-[0.3em] text-[#a09a8c]">
        <span>Carte</span>
        <span v-if="nfcId">NFC · {{ nfcId }}</span>
      </div>

      <header class="rise rise-1 mt-9 text-center">
        <p class="mx-auto h-px w-8 bg-[#c4b9a4]"></p>
        <h1 class="mt-5 text-balance font-serif text-[36px] leading-[1.15] tracking-tight text-[#22201c]">
          {{ payload?.shopName || '餐厅' }}
        </h1>
        <p v-if="payload?.tagline" class="mx-auto mt-3.5 max-w-[19rem] text-[12.5px] leading-[1.8] text-[#8a8377]">
          {{ payload.tagline }}
        </p>
        <p class="mx-auto mt-6 h-px w-16 bg-[#dcd5c6]"></p>
      </header>

      <section v-if="groups.length" class="mt-11 space-y-11">
        <div v-for="(group, gi) in groups" :key="gi" class="rise" :class="riseDelay(gi)">
          <div class="text-center">
            <h2 class="font-serif text-[15px] italic tracking-[0.14em] text-[#6b6255]">{{ group.name }}</h2>
            <p class="mx-auto mt-2 h-px w-6 bg-[#d8d1c2]"></p>
          </div>

          <ul class="mt-4">
            <li v-for="(item, i) in group.items" :key="i" class="fc-item">
              <div class="min-w-0 flex-1">
                <div class="flex flex-wrap items-baseline gap-x-2 gap-y-1">
                  <span class="text-[15px] font-medium tracking-wide text-[#22201c]">{{ item?.name || '菜品' }}</span>
                  <span v-if="isSignature(item)" class="fc-tag">推荐</span>
                  <span v-if="isVegetarian(item)" class="fc-tag fc-tag-green">素</span>
                  <span v-if="spicyLevel(item)" class="inline-flex items-center gap-[1px] text-[#b45309]">
                    <Flame v-for="n in spicyLevel(item)" :key="n" class="h-2.5 w-2.5" />
                  </span>
                </div>
                <p v-if="item?.desc" class="mt-1 max-w-[19rem] text-[11.5px] leading-[1.7] text-[#9a9384]">{{ item.desc }}</p>
              </div>
              <span v-if="hasPrice(item)" class="fc-price">¥{{ item.price }}</span>
            </li>
          </ul>
        </div>
      </section>

      <div v-else class="rise mt-11 border-y border-[#e2dbcb] py-14 text-center">
        <ChefHat class="mx-auto h-6 w-6 text-[#c4b9a4]" />
        <p class="mt-3 font-serif text-[13px] italic text-[#a09a8c]">菜单正在准备中</p>
      </div>

      <div v-if="payload?.hours || payload?.address || payload?.phone" class="rise mt-12 border-t border-[#e2dbcb] pt-7">
        <dl class="space-y-3.5">
          <div v-if="payload?.hours" class="flex items-baseline justify-between gap-4">
            <dt class="font-mono text-[9.5px] uppercase tracking-[0.22em] text-[#a09a8c]">Hours</dt>
            <dd class="text-[12.5px] tabular-nums tracking-wide text-[#4a4438]">{{ payload.hours }}</dd>
          </div>
          <div v-if="payload?.address" class="flex items-baseline justify-between gap-4">
            <dt class="flex-shrink-0 font-mono text-[9.5px] uppercase tracking-[0.22em] text-[#a09a8c]">Address</dt>
            <dd class="text-right text-[12px] leading-relaxed text-[#4a4438]">{{ payload.address }}</dd>
          </div>
          <div v-if="payload?.phone" class="flex items-baseline justify-between gap-4">
            <dt class="font-mono text-[9.5px] uppercase tracking-[0.22em] text-[#a09a8c]">Reserve</dt>
            <dd class="text-[12.5px] tabular-nums tracking-wide text-[#4a4438]">{{ payload.phone }}</dd>
          </div>
        </dl>
      </div>

      <div v-if="payload?.notice" class="rise mt-7">
        <p class="text-center font-serif text-[11.5px] italic leading-relaxed text-[#a09a8c]">{{ payload.notice }}</p>
      </div>

      <div v-if="payload?.phone" class="rise rise-3 mt-9">
        <a :href="'tel:' + payload.phone" class="fc-cta">
          <Phone class="h-3.5 w-3.5" />
          致电预订
        </a>
      </div>

      <p class="mt-10 text-center font-mono text-[9px] uppercase tracking-[0.32em] text-[#c4b9a4]">Bon Appétit</p>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { Phone, Flame, ChefHat } from 'lucide-vue-next'

const props = defineProps(['payload', 'nfcId'])

const hasPrice = (item) => item?.price !== undefined && item?.price !== null && item?.price !== ''
const riseDelay = (i) => (i === 0 ? 'rise-1' : i === 1 ? 'rise-2' : 'rise-3')

const isSignature = (item) => {
  if (item?.signature === true) return true
  const tags = Array.isArray(item?.tags) ? item.tags : []
  return tags.some((t) => String(t).includes('招牌') || String(t).includes('推荐') || String(t).toLowerCase() === 'signature')
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
.fc-item {
  display: flex;
  align-items: baseline;
  gap: 16px;
  padding: 14px 0;
  border-bottom: 1px solid #ece6d8;
}
.fc-item:last-child {
  border-bottom: none;
}

.fc-tag {
  font-family: 'Geist', 'PingFang SC', 'Microsoft YaHei', sans-serif;
  border: 1px solid rgba(180, 83, 9, 0.3);
  border-radius: 2px;
  padding: 0 4px;
  font-size: 8.5px;
  font-weight: 500;
  letter-spacing: 0.12em;
  color: #b45309;
}
.fc-tag-green {
  border-color: rgba(21, 128, 61, 0.28);
  color: #15803d;
}

.fc-price {
  flex-shrink: 0;
  font-family: 'Geist', 'PingFang SC', sans-serif;
  font-size: 13.5px;
  font-weight: 500;
  font-variant-numeric: tabular-nums;
  letter-spacing: 0.02em;
  color: #4a4438;
}

.fc-cta {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
  width: 100%;
  border: 1px solid #22201c;
  border-radius: 2px;
  background: transparent;
  padding: 13px 0;
  font-size: 11.5px;
  font-weight: 500;
  letter-spacing: 0.2em;
  color: #22201c;
  transition: background-color 0.25s ease, color 0.25s ease;
}
.fc-cta:hover {
  background: #22201c;
  color: #f6f4ef;
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
