<template>
  <div class="relative min-h-screen overflow-hidden bg-[#faf6ef] text-[#2a2119]">
    <!-- Warm light -->
    <div class="pointer-events-none absolute inset-x-0 top-0 h-64 bg-[radial-gradient(60%_100%_at_50%_0%,rgba(194,65,12,0.10),transparent_70%)]"></div>

    <div class="relative z-10 mx-auto w-full max-w-md px-6 pb-12 pt-10">
      <!-- Meta -->
      <div class="rise flex items-center justify-between font-mono text-[10px] uppercase tracking-[0.24em] text-[#b9a483]">
        <span>Menu Card</span>
        <span v-if="nfcId">NFC · {{ nfcId }}</span>
      </div>

      <!-- Identity -->
      <div class="rise rise-1 mt-8">
        <h1 class="text-balance font-serif text-[34px] leading-[1.1] tracking-tight text-[#7c2d12]">
          {{ payload?.shopName || '餐厅' }}
        </h1>
        <p v-if="payload?.cuisine" class="mt-3 flex items-center gap-2.5 text-[12px] tracking-[0.2em] text-[#c2410c]">
          <span class="h-px w-7 bg-[#c2410c]/40"></span>
          {{ payload.cuisine }}
        </p>
      </div>

      <!-- Hours sign -->
      <div v-if="payload?.hours || payload?.address" class="rsv-sign rise rise-2 mt-8">
        <div v-if="payload?.hours" class="flex items-center gap-3.5">
          <span class="flex h-10 w-10 flex-shrink-0 items-center justify-center rounded-lg bg-[#c2410c]/10 text-[#c2410c]">
            <Clock class="h-[18px] w-[18px]" />
          </span>
          <div class="min-w-0">
            <p class="text-[10px] font-medium uppercase tracking-[0.22em] text-[#b9a483]">营业时间</p>
            <p class="mt-0.5 text-[15px] font-semibold tracking-wide tabular-nums text-[#2a2119]">{{ payload.hours }}</p>
          </div>
        </div>
        <div v-if="payload?.address" class="mt-4 flex items-start gap-3.5 border-t border-[#f0e5d3] pt-4">
          <span class="flex h-10 w-10 flex-shrink-0 items-center justify-center rounded-lg bg-[#c2410c]/10 text-[#c2410c]">
            <MapPin class="h-[18px] w-[18px]" />
          </span>
          <div class="min-w-0 pt-0.5">
            <p class="text-[10px] font-medium uppercase tracking-[0.22em] text-[#b9a483]">门店地址</p>
            <p class="mt-0.5 text-[13px] leading-relaxed text-[#5f5240]">{{ payload.address }}</p>
          </div>
        </div>
      </div>

      <!-- Menu -->
      <div class="rise rise-3 mt-10">
        <div class="flex items-center gap-3">
          <h2 class="text-[13px] font-semibold tracking-wide text-[#7c2d12]">菜单推荐</h2>
          <span class="h-px flex-1 bg-[#e8d9c5]"></span>
          <span class="font-mono text-[10px] uppercase tracking-[0.24em] text-[#b9a483]">Menu</span>
        </div>

        <div v-if="menu.length" class="mt-2">
          <div v-for="(item, i) in menu" :key="i" class="rsv-item">
            <div v-if="item?.image" class="h-14 w-14 flex-shrink-0 overflow-hidden rounded-xl border border-[#eadcc8] bg-[#f3ead9]">
              <img :src="item.image" :alt="item.name || '菜品'" class="h-full w-full object-cover" />
            </div>
            <div class="min-w-0 flex-1">
              <div class="flex items-baseline gap-2">
                <p class="min-w-0 truncate text-[15px] font-medium text-[#2a2119]">{{ item?.name || '菜品' }}</p>
                <span class="rsv-dots"></span>
                <p v-if="item?.price || item?.price === 0" class="flex-shrink-0 text-[15px] font-semibold tabular-nums text-[#c2410c]">
                  ¥{{ item.price }}
                </p>
              </div>
              <p v-if="item?.desc" class="mt-1 text-xs leading-relaxed text-[#8a7860]">{{ item.desc }}</p>
            </div>
          </div>
        </div>

        <div v-else class="rsv-empty mt-4">
          <UtensilsCrossed class="h-5 w-5 text-[#c9b394]" />
          <p class="mt-2 text-xs text-[#a89275]">菜单更新中，欢迎电话咨询</p>
        </div>
      </div>

      <!-- CTA -->
      <div v-if="payload?.phone" class="rise rise-3 mt-9">
        <a :href="'tel:' + payload.phone" class="rsv-cta">
          <Phone class="h-4 w-4" />
          电话预约
        </a>
        <p class="mt-3 text-center font-mono text-xs tracking-[0.14em] tabular-nums text-[#a89275]">{{ payload.phone }}</p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { Clock, MapPin, Phone, UtensilsCrossed } from 'lucide-vue-next'

const props = defineProps(['payload', 'nfcId'])

const menu = computed(() => {
  const list = props.payload?.menu
  return Array.isArray(list) ? list.filter((item) => item && (item.name || item.price)) : []
})
</script>

<style scoped>
.rsv-sign {
  position: relative;
  border-radius: 16px;
  border: 1px solid #e8d9c5;
  background: #fffdf8;
  padding: 18px;
  box-shadow: 0 22px 44px -30px rgba(124, 45, 18, 0.5);
}
.rsv-sign::before,
.rsv-sign::after {
  content: '';
  position: absolute;
  top: 9px;
  height: 4px;
  width: 4px;
  border-radius: 9999px;
  background: #e3d3ba;
}
.rsv-sign::before {
  left: 11px;
}
.rsv-sign::after {
  right: 11px;
}

.rsv-item {
  display: flex;
  align-items: center;
  gap: 14px;
  padding: 15px 2px;
  border-bottom: 1px solid #efe4d2;
}
.rsv-item:last-child {
  border-bottom: none;
}
.rsv-dots {
  flex: 1;
  min-width: 18px;
  border-bottom: 2px dotted rgba(154, 52, 18, 0.28);
  transform: translateY(-3px);
}

.rsv-empty {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 32px 0;
  border: 1px dashed #e0d0b8;
  border-radius: 16px;
  background: #fffdf8;
}

.rsv-cta {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
  width: 100%;
  border-radius: 14px;
  background: #c2410c;
  padding: 14px 0;
  font-size: 14px;
  font-weight: 600;
  letter-spacing: 0.06em;
  color: #fff8f0;
  box-shadow: 0 18px 36px -18px rgba(194, 65, 12, 0.6);
  transition: background-color 0.2s ease, transform 0.15s ease;
}
.rsv-cta:hover {
  background: #9a3412;
}
.rsv-cta:active {
  transform: scale(0.98);
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
