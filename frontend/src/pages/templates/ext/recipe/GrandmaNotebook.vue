<template>
  <div class="relative min-h-screen overflow-hidden bg-[#efe2c8] text-[#3a2b1d]">
    <div class="gn-dots pointer-events-none absolute inset-0"></div>
    <div class="relative z-10 mx-auto w-full max-w-md px-5 pb-14 pt-8">
      <div class="gn-sheet relative px-6 pb-10 pt-9">
        <span class="gn-tape gn-tape-l"></span>
        <span class="gn-tape gn-tape-r"></span>
        <div class="gn-holes">
          <span v-for="n in 6" :key="n"></span>
        </div>
        <span class="gn-margin pointer-events-none absolute inset-y-8 left-11 w-px"></span>

        <div class="gn-rise relative pl-8">
          <div class="flex items-center justify-between font-mono text-[9px] uppercase tracking-[0.2em] text-[#a08a6b]">
            <span>家庭食谱 · 第 01 册</span>
            <span v-if="nfcId">NFC {{ nfcId }}</span>
          </div>

          <header class="mt-6">
            <p class="gn-script text-[13px] text-[#9c7c56]">今天想吃点什么？</p>
            <h1 class="gn-hand mt-1 text-[36px] font-bold leading-[1.15] text-[#3a2b1d]">
              {{ payload?.shopName || '外婆的厨房' }}
            </h1>
            <p v-if="payload?.tagline" class="gn-script mt-2 max-w-[17rem] text-[15px] leading-[1.7] text-[#6f5a40]">
              {{ payload.tagline }}
            </p>
            <div class="mt-4 flex flex-wrap gap-x-5 gap-y-1.5 font-mono text-[10px] text-[#8b7455]">
              <span v-if="payload?.hours" class="inline-flex items-center gap-1.5"><Clock class="h-3 w-3" />{{ payload.hours }}</span>
              <span v-if="payload?.address" class="inline-flex items-center gap-1.5"><MapPin class="h-3 w-3" />{{ payload.address }}</span>
            </div>
          </header>

          <div class="gn-divider mt-5"></div>

          <section v-if="groups.length" class="mt-6 space-y-7">
            <div v-for="(group, gi) in groups" :key="gi" class="gn-rise" :style="{ animationDelay: 0.12 + gi * 0.07 + 's' }">
              <div class="flex items-baseline justify-between gap-3">
                <h2 class="gn-hand text-[22px] font-bold text-[#7c3f2c]">{{ group.name }}</h2>
                <span class="gn-script text-[12px] text-[#ad9371]">{{ group.items.length }} 样</span>
              </div>
              <ul class="mt-1">
                <li v-for="(item, i) in group.items" :key="i" class="gn-row">
                  <div class="min-w-0 flex-1">
                    <div class="flex flex-wrap items-center gap-x-2 gap-y-0.5">
                      <span class="gn-hand text-[19px] leading-tight text-[#332619]">{{ item?.name }}</span>
                      <span v-if="signature(item)" class="gn-stamp">招牌</span>
                      <span v-if="veg(item)" class="text-[11px] text-[#5f7d43]">素</span>
                      <span v-if="spicy(item)" class="inline-flex text-[#b24c2e]"><Flame v-for="n in spicy(item)" :key="n" class="h-3 w-3" /></span>
                    </div>
                    <p v-if="item?.desc" class="gn-script mt-0.5 text-[13px] leading-[1.55] text-[#816b4f]">{{ item.desc }}</p>
                  </div>
                  <span v-if="hasPrice(item)" class="gn-price">{{ item.price }}</span>
                </li>
              </ul>
            </div>
          </section>

          <div v-else class="gn-empty">
            <UtensilsCrossed class="h-5 w-5" />
            <p>这页还空着，等一道好菜。</p>
          </div>

          <div v-if="payload?.notice" class="gn-note mt-6">
            <span class="gn-pin"></span>
            <p class="gn-script text-[13px] leading-relaxed text-[#6f5a40]">{{ payload.notice }}</p>
          </div>

          <a v-if="payload?.phone" :href="'tel:' + payload.phone" class="gn-call mt-7">
            <Phone class="h-4 w-4" />
            <span class="gn-hand text-[16px]">打电话留一份</span>
            <span class="ml-auto font-mono text-[10px] tracking-wide text-[#8a6a48]">{{ payload.phone }}</span>
          </a>

          <p class="gn-script mt-8 text-center text-[12px] text-[#b29a79]">— 记于灶台边 · 盐少许，爱多一点 —</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { Clock, MapPin, Phone, Flame, UtensilsCrossed } from 'lucide-vue-next'

const props = defineProps(['payload', 'nfcId'])

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
</script>

<style scoped>
.gn-dots {
  background-image: radial-gradient(rgba(120, 88, 48, 0.16) 1px, transparent 1px);
  background-size: 18px 18px;
}
.gn-sheet {
  border-radius: 3px 6px 5px 3px;
  background-color: #fcf6e4;
  background-image: repeating-linear-gradient(transparent 0 31px, rgba(84, 126, 156, 0.16) 31px 32px);
  box-shadow: 0 2px 0 #e6d6b6, 0 30px 60px -34px rgba(76, 52, 26, 0.55);
}
.gn-tape {
  position: absolute;
  top: -12px;
  z-index: 2;
  height: 28px;
  width: 86px;
  background: rgba(201, 178, 130, 0.65);
  box-shadow: inset 0 0 0 1px rgba(255, 255, 255, 0.2);
}
.gn-tape-l { left: 24px; transform: rotate(-6deg); }
.gn-tape-r { right: 20px; transform: rotate(5deg); background: rgba(189, 169, 122, 0.55); }
.gn-holes {
  position: absolute;
  left: 9px;
  top: 76px;
  bottom: 44px;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
}
.gn-holes span {
  height: 9px;
  width: 9px;
  border-radius: 9999px;
  background: #e8d9bd;
  box-shadow: inset 0 1px 2px rgba(82, 58, 32, 0.35);
}
.gn-margin {
  background: rgba(190, 92, 70, 0.32);
}
.gn-hand { font-family: 'Kaiti SC', 'STKaiti', 'KaiTi', 'Songti SC', serif; }
.gn-script { font-family: 'Kaiti SC', 'STKaiti', 'KaiTi', 'Songti SC', serif; }
.gn-divider {
  height: 7px;
  background: repeating-linear-gradient(-45deg, rgba(124, 63, 44, 0.35) 0 2px, transparent 2px 7px);
  opacity: 0.6;
}
.gn-row {
  display: flex;
  align-items: baseline;
  gap: 10px;
  padding: 7px 0;
  border-bottom: 1px dashed rgba(124, 98, 60, 0.22);
}
.gn-row:last-child { border-bottom: none; }
.gn-price {
  flex-shrink: 0;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  min-width: 40px;
  padding: 2px 7px;
  border: 1.5px solid rgba(178, 76, 46, 0.55);
  border-radius: 50%;
  font-family: 'Kaiti SC', 'STKaiti', serif;
  font-size: 15px;
  color: #a8452c;
  transform: rotate(-4deg);
}
.gn-price::before { content: '¥'; font-size: 9px; margin-right: 1px; opacity: 0.7; }
.gn-stamp {
  border: 1px solid rgba(178, 76, 46, 0.5);
  border-radius: 2px;
  padding: 0 4px;
  font-size: 9px;
  letter-spacing: 0.1em;
  color: #a8452c;
  transform: rotate(-2deg);
}
.gn-note {
  position: relative;
  padding: 13px 15px 13px 18px;
  background: #f5e8c8;
  box-shadow: 1px 2px 0 rgba(150, 120, 78, 0.15);
  transform: rotate(-0.6deg);
}
.gn-pin {
  position: absolute;
  top: -6px;
  left: 50%;
  height: 12px;
  width: 12px;
  border-radius: 9999px;
  background: radial-gradient(circle at 35% 30%, #d98f74, #9d4b34);
  box-shadow: 0 2px 3px rgba(80, 45, 30, 0.35);
}
.gn-empty {
  margin-top: 28px;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 8px;
  border: 1px dashed rgba(124, 98, 60, 0.35);
  padding: 32px 16px;
  color: #9a8568;
  font-family: 'Kaiti SC', 'STKaiti', serif;
  font-size: 14px;
}
.gn-call {
  display: flex;
  align-items: center;
  gap: 9px;
  border: 1.5px solid #7c3f2c;
  border-radius: 3px;
  background: #f6ecd2;
  padding: 10px 13px;
  color: #6d3624;
  transition: transform 0.16s ease, background-color 0.2s ease;
}
.gn-call:hover { background: #f1e1bc; transform: rotate(-0.5deg); }
.gn-call:active { transform: scale(0.985); }
.gn-rise { animation: gn-in 0.7s cubic-bezier(0.2, 0.9, 0.25, 1) both; }
@keyframes gn-in {
  from { opacity: 0; transform: translateY(14px) rotate(0.4deg); }
  to { opacity: 1; transform: none; }
}
@media (prefers-reduced-motion: reduce) {
  .gn-rise, .gn-call { animation: none !important; transition: none !important; }
  .gn-call:hover { transform: none; }
}
</style>
