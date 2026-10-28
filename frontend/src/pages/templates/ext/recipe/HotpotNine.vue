<template>
  <div class="relative min-h-screen overflow-hidden bg-[#160b08] text-[#f6e7d4]">
    <div class="hp-heat pointer-events-none absolute inset-x-0 top-0 h-[380px]"></div>
    <div class="hp-grain pointer-events-none absolute inset-0"></div>

    <div class="relative z-10 mx-auto w-full max-w-md px-5 pb-14 pt-8">
      <div class="hp-rise flex items-center justify-between font-mono text-[8px] uppercase tracking-[0.24em] text-[#b08a6b]">
        <span class="inline-flex items-center gap-1.5"><span class="hp-fire"></span>Boiling Point</span>
        <span v-if="nfcId">NFC {{ nfcId }}</span>
      </div>

      <header class="hp-rise hp-d1 mt-6 text-center">
        <p class="font-mono text-[8px] tracking-[0.42em] text-[#e0603d]">九 宫 格 · 麻 辣 鲜 香</p>
        <h1 class="hp-title mt-3 text-[50px] font-black leading-[0.95] tracking-[-0.02em] text-[#ffedd6]">
          {{ payload?.shopName || '红汤九宫格' }}
        </h1>
        <p v-if="payload?.tagline" class="mx-auto mt-2 max-w-[18rem] text-[11px] leading-relaxed text-[#c19c7c]">{{ payload.tagline }}</p>
        <div class="mt-4 flex items-center justify-center gap-3 font-mono text-[8.5px] text-[#b08a6b]">
          <span v-if="payload?.hours">营业 {{ payload.hours }}</span>
          <span v-if="payload?.hours" class="h-3 w-px bg-[#5a3324]"></span>
          <span>辣度可调</span>
        </div>
      </header>

      <section v-if="groups.length" class="hp-rise hp-d2 mt-6">
        <div class="hp-pot">
          <div v-for="(group, gi) in groups.slice(0, 6)" :key="gi" class="hp-cell">
            <div class="flex items-center justify-between">
              <h2 class="text-[10px] font-bold tracking-[0.12em] text-[#ffd7ad]">{{ group.name }}</h2>
              <Flame class="h-3 w-3 text-[#e0603d]" />
            </div>
            <div class="hp-divider"></div>
            <ul class="mt-1 space-y-1">
              <li v-for="(item, i) in group.items.slice(0, 3)" :key="i" class="hp-dish">
                <span class="min-w-0 flex-1 truncate text-[10px] leading-tight text-[#f0ddc7]">{{ item?.name }}</span>
                <span v-if="hasPrice(item)" class="font-mono text-[8.5px] text-[#f0a06a]">{{ item.price }}</span>
              </li>
            </ul>
            <p v-if="group.items.length > 3" class="mt-1 font-mono text-[7px] text-[#9c7357]">+{{ group.items.length - 3 }} 道</p>
          </div>

          <div class="hp-cell hp-info">
            <span class="hp-info-k">辣度</span>
            <div class="flex items-center gap-0.5">
              <Flame v-for="n in 3" :key="n" class="h-3 w-3" :class="n <= 2 ? 'text-[#e0603d]' : 'text-[#5c3527]'" />
            </div>
            <p class="mt-1 text-[8px] text-[#c19c7c]">微辣 · 中辣 · 特辣</p>
          </div>
          <div class="hp-cell hp-info">
            <span class="hp-info-k">锅底</span>
            <p class="text-[10px] leading-tight text-[#ffd7ad]">牛油红汤<br />菌菇清汤</p>
            <p class="mt-1 text-[8px] text-[#c19c7c]">可选鸳鸯</p>
          </div>
          <div class="hp-cell hp-info">
            <span class="hp-info-k">蘸料</span>
            <p class="text-[9px] leading-[1.5] text-[#d9b493]">香油 · 蒜泥<br />蚝油 · 香菜</p>
            <p class="mt-1 text-[8px] text-[#c19c7c]">自助台</p>
          </div>
        </div>
      </section>

      <div v-else class="hp-rise mt-6 rounded border border-dashed border-[#5a3324] px-6 py-12 text-center">
        <Flame class="mx-auto h-6 w-6 text-[#7a422d]" />
        <p class="mt-3 text-[11px] text-[#b08a6b]">锅还没开，菜单正在下锅。</p>
      </div>

      <div class="hp-rise hp-d3 mt-5 rounded-lg border border-[#5a3324] bg-[#23110b] px-4 py-3">
        <div class="flex items-start gap-2">
          <Info class="mt-0.5 h-3.5 w-3.5 flex-shrink-0 text-[#e0603d]" />
          <p class="text-[10px] leading-relaxed text-[#c19c7c]">{{ payload?.notice || '所有菜品现点现配，红汤锅底每日现炒。' }}</p>
        </div>
      </div>

      <div class="hp-rise hp-d3 mt-5 flex items-center gap-2">
        <a v-if="payload?.phone" :href="'tel:' + payload.phone" class="hp-call">
          <Phone class="h-4 w-4" />
          <span class="text-[13px] font-bold tracking-[0.06em]">订一桌</span>
          <span class="ml-auto font-mono text-[9px]">{{ payload.phone }}</span>
        </a>
      </div>

      <div v-if="payload?.address" class="mt-4 flex items-start justify-center gap-1.5 text-[9px] leading-relaxed text-[#9c7357]">
        <MapPin class="mt-0.5 h-3 w-3 flex-shrink-0" />
        <span>{{ payload.address }}</span>
      </div>

      <div class="mt-7 flex items-center justify-between font-mono text-[7px] uppercase tracking-[0.22em] text-[#7d523b]">
        <span>Since 1998 · 老灶</span>
        <span>Open till 02:00</span>
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { Flame, Info, MapPin, Phone } from 'lucide-vue-next'

const props = defineProps(['payload', 'nfcId'])

const hasPrice = (item) => item?.price !== undefined && item?.price !== null && item?.price !== ''

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
.hp-heat {
  background: radial-gradient(70% 80% at 50% 0%, rgba(224, 96, 61, 0.28), transparent 70%),
    radial-gradient(40% 40% at 20% 20%, rgba(180, 40, 20, 0.18), transparent 70%);
}
.hp-grain {
  opacity: 0.35;
  background-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' width='130' height='130'%3E%3Cfilter id='n'%3E%3CfeTurbulence type='fractalNoise' baseFrequency='0.9' numOctaves='3'/%3E%3C/filter%3E%3Crect width='130' height='130' filter='url(%23n)' opacity='0.07'/%3E%3C/svg%3E");
}
.hp-title { font-family: 'STKaiti', 'Kaiti SC', 'Songti SC', serif; }
.hp-fire {
  display: inline-block;
  height: 6px;
  width: 6px;
  border-radius: 9999px;
  background: #e0603d;
  box-shadow: 0 0 9px 2px rgba(224, 96, 61, 0.5);
  animation: hp-pulse 2.2s ease-in-out infinite;
}
@keyframes hp-pulse {
  0%, 100% { transform: scale(1); opacity: 1; }
  50% { transform: scale(1.35); opacity: 0.7; }
}
.hp-pot {
  display: grid;
  grid-template-columns: repeat(3, minmax(0, 1fr));
  gap: 4px;
  border: 6px solid #3a1a10;
  border-radius: 14px;
  background: #3a1a10;
  box-shadow: inset 0 0 30px rgba(0, 0, 0, 0.6), 0 16px 40px -24px rgba(224, 96, 61, 0.5);
}
.hp-cell {
  min-height: 92px;
  border-radius: 5px;
  background: linear-gradient(160deg, #2c130c, #1e0d08);
  padding: 7px 7px 6px;
  box-shadow: inset 0 0 18px rgba(224, 96, 61, 0.12);
}
.hp-divider { margin-top: 5px; height: 1px; background: linear-gradient(90deg, rgba(224, 96, 61, 0.5), transparent); }
.hp-dish { display: flex; align-items: baseline; gap: 4px; }
.hp-info { display: flex; flex-direction: column; justify-content: center; background: #24100a; }
.hp-info-k { font-family: ui-monospace, monospace; font-size: 7px; letter-spacing: 0.2em; color: #e0603d; }
.hp-call {
  display: flex;
  align-items: center;
  gap: 9px;
  width: 100%;
  border-radius: 8px;
  background: linear-gradient(100deg, #c8381f, #8f1f10);
  padding: 12px 15px;
  color: #ffe9d2;
  box-shadow: 0 16px 32px -20px rgba(200, 56, 31, 0.9);
  transition: filter 0.2s ease, transform 0.15s ease;
}
.hp-call:hover { filter: brightness(1.08); }
.hp-call:active { transform: scale(0.985); }
.hp-rise { animation: hp-in 0.7s cubic-bezier(0.2, 0.9, 0.25, 1) both; }
.hp-d1 { animation-delay: 0.08s; }
.hp-d2 { animation-delay: 0.15s; }
.hp-d3 { animation-delay: 0.3s; }
@keyframes hp-in {
  from { opacity: 0; transform: translateY(14px); }
  to { opacity: 1; transform: none; }
}
@media (prefers-reduced-motion: reduce) {
  .hp-rise, .hp-fire, .hp-call { animation: none !important; transition: none !important; }
}
</style>
