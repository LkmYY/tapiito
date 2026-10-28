<template>
  <div class="relative min-h-screen overflow-hidden bg-[#fbf7ee] text-[#1f1a14]">
    <div class="rp-grain pointer-events-none absolute inset-0"></div>
    <div class="pointer-events-none absolute inset-x-0 top-0 h-72 bg-[radial-gradient(70%_100%_at_50%_0%,rgba(185,28,28,0.10),transparent_72%)]"></div>

    <div class="relative z-10 mx-auto w-full max-w-md px-6 pb-12 pt-9">
      <!-- Meta -->
      <div class="rise flex items-center justify-between font-mono text-[10px] uppercase tracking-[0.24em] text-[#a89880]">
        <span>Menu · 今日供应</span>
        <span v-if="nfcId">NFC · {{ nfcId }}</span>
      </div>

      <!-- Identity -->
      <header class="rise rise-1 mt-7">
        <h1 class="text-balance text-[34px] font-black leading-[1.05] tracking-tight text-[#1f1a14]">
          {{ payload?.shopName || '小馆菜单' }}
        </h1>
        <p v-if="payload?.tagline" class="mt-2.5 text-[13px] leading-relaxed text-[#7a6a55]">
          {{ payload.tagline }}
        </p>
      </header>

      <!-- Info strip -->
      <div v-if="payload?.hours || payload?.address" class="rp-strip rise rise-2 mt-6">
        <div v-if="payload?.hours" class="flex items-center gap-3">
          <Clock class="h-4 w-4 flex-shrink-0 text-[#b91c1c]" />
          <div class="min-w-0 flex-1">
            <p class="text-[10px] uppercase tracking-[0.2em] text-[#a89880]">营业时间</p>
            <p class="mt-0.5 text-[13px] font-semibold tabular-nums tracking-wide">{{ payload.hours }}</p>
          </div>
        </div>
        <div v-if="payload?.address" class="mt-3 flex items-start gap-3 border-t border-[#efe4d2] pt-3">
          <MapPin class="mt-0.5 h-4 w-4 flex-shrink-0 text-[#b91c1c]" />
          <div class="min-w-0 flex-1">
            <p class="text-[10px] uppercase tracking-[0.2em] text-[#a89880]">门店地址</p>
            <p class="mt-0.5 text-[12.5px] leading-relaxed text-[#5f5240]">{{ payload.address }}</p>
          </div>
        </div>
      </div>

      <!-- Categories -->
      <section v-if="groups.length" class="mt-9 space-y-9">
        <div v-for="(group, gi) in groups" :key="gi" class="rise" :class="riseDelay(gi)">
          <div class="flex items-center gap-2.5">
            <span class="h-4 w-1 rounded-full bg-[#b91c1c]"></span>
            <h2 class="text-[15px] font-bold tracking-tight text-[#1f1a14]">{{ group.name }}</h2>
            <span class="h-px flex-1 bg-[#e7dcc8]"></span>
            <span class="font-mono text-[10px] tracking-[0.18em] text-[#b8a88f]">{{ pad(group.items.length) }}</span>
          </div>

          <ul class="mt-1">
            <li v-for="(item, i) in group.items" :key="i" class="rp-item">
              <div class="min-w-0 flex-1">
                <div class="flex flex-wrap items-center gap-x-1.5 gap-y-1">
                  <span class="text-[15px] font-semibold text-[#1f1a14]">{{ item?.name || '菜品' }}</span>
                  <span v-if="isSignature(item)" class="rp-tag">招牌</span>
                  <span v-if="spicyLevel(item)" class="inline-flex items-center text-[#c2410c]" :title="'辣度 ' + spicyLevel(item)">
                    <Flame v-for="n in spicyLevel(item)" :key="n" class="h-3 w-3" />
                  </span>
                  <span v-if="isVegetarian(item)" class="rp-tag rp-tag-green">素</span>
                </div>
                <p v-if="item?.desc" class="mt-1 text-xs leading-relaxed text-[#8a7860]">{{ item.desc }}</p>
              </div>
              <span class="rp-leader"></span>
              <span v-if="hasPrice(item)" class="flex-shrink-0 text-[16px] font-bold tabular-nums text-[#b91c1c]">
                ¥{{ item.price }}
              </span>
            </li>
          </ul>
        </div>
      </section>

      <!-- Empty -->
      <div v-else class="rise mt-10 rounded-2xl border border-dashed border-[#ddcfb6] bg-[#fffdf7] px-6 py-12 text-center">
        <UtensilsCrossed class="mx-auto h-6 w-6 text-[#c9b99f]" />
        <p class="mt-3 text-sm text-[#a89275]">菜单整理中，欢迎电话咨询</p>
      </div>

      <!-- Notice -->
      <div v-if="payload?.notice" class="rise mt-9 flex items-start gap-2.5 rounded-xl bg-[#f4ecdc] px-4 py-3">
        <Info class="mt-0.5 h-3.5 w-3.5 flex-shrink-0 text-[#a89275]" />
        <p class="text-[11.5px] leading-relaxed text-[#8a7860]">{{ payload.notice }}</p>
      </div>

      <!-- CTA -->
      <div v-if="payload?.phone" class="rise rise-3 mt-8">
        <a :href="'tel:' + payload.phone" class="rp-cta">
          <Phone class="h-4 w-4" />
          打电话订餐
          <span class="ml-auto font-mono text-[11px] font-normal tracking-[0.12em] tabular-nums text-white/70">{{ payload.phone }}</span>
        </a>
      </div>

      <p class="mt-8 text-center font-mono text-[9px] uppercase tracking-[0.28em] text-[#c2b49c]">
        Scan · Order · Enjoy
      </p>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { Clock, MapPin, Phone, Flame, Info, UtensilsCrossed } from 'lucide-vue-next'

const props = defineProps(['payload', 'nfcId'])

const pad = (n) => String(n).padStart(2, '0')
const hasPrice = (item) => item?.price !== undefined && item?.price !== null && item?.price !== ''

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

const riseDelay = (i) => (i === 0 ? 'rise-1' : i === 1 ? 'rise-2' : 'rise-3')

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
.rp-grain {
  opacity: 0.5;
  background-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' width='120' height='120'%3E%3Cfilter id='n'%3E%3CfeTurbulence type='fractalNoise' baseFrequency='0.9' numOctaves='3'/%3E%3C/filter%3E%3Crect width='120' height='120' filter='url(%23n)' opacity='0.06'/%3E%3C/svg%3E");
}

.rp-strip {
  border: 1px solid #e7dcc8;
  border-radius: 14px;
  background: #fffdf7;
  padding: 14px 16px;
  box-shadow: 0 18px 40px -34px rgba(95, 82, 64, 0.55);
}

.rp-item {
  display: flex;
  align-items: flex-end;
  gap: 10px;
  padding: 14px 2px;
  border-bottom: 1px solid #f0e7d6;
}
.rp-item:last-child {
  border-bottom: none;
}

.rp-leader {
  flex: 1;
  min-width: 14px;
  margin-bottom: 5px;
  border-bottom: 2px dotted rgba(185, 28, 28, 0.22);
}

.rp-tag {
  display: inline-flex;
  align-items: center;
  border: 1px solid rgba(185, 28, 28, 0.35);
  border-radius: 4px;
  padding: 1px 5px;
  font-size: 9px;
  font-weight: 600;
  letter-spacing: 0.08em;
  color: #b91c1c;
}
.rp-tag-green {
  border-color: rgba(21, 128, 61, 0.32);
  color: #15803d;
}

.rp-cta {
  display: flex;
  align-items: center;
  gap: 9px;
  width: 100%;
  border-radius: 12px;
  background: #b91c1c;
  padding: 14px 18px;
  font-size: 14px;
  font-weight: 600;
  letter-spacing: 0.04em;
  color: #fff8f0;
  box-shadow: 0 18px 36px -20px rgba(185, 28, 28, 0.65);
  transition: background-color 0.2s ease, transform 0.15s ease;
}
.rp-cta:hover {
  background: #991b1b;
}
.rp-cta:active {
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
