<template>
  <div class="pdb-root min-h-screen bg-[#f5e900] text-[#101010]">
    <div class="mx-auto w-full max-w-md border-x-2 border-[#101010] bg-[#f5e900]">
      <div class="pdb-ticker overflow-hidden border-b-2 border-[#101010] bg-[#101010] py-1.5 text-[#f5e900]">
        <div class="pdb-ticker-track font-mono text-[9px] uppercase tracking-[0.24em]">
          <span v-for="n in 4" :key="n" class="px-4">In stock — ships in 24h — no app needed — tap to buy —</span>
        </div>
      </div>

      <header class="pdb-in flex items-center justify-between border-b-2 border-[#101010] px-5 py-3">
        <span class="text-[17px] font-black uppercase tracking-[-0.04em]">WARE/HOUSE</span>
        <span class="border-2 border-[#101010] px-2 py-0.5 font-mono text-[9px] uppercase">{{ nfcId || 'DROP 07' }}</span>
      </header>

      <section class="pdb-in pdb-d1 pdb-visual relative m-4 aspect-[1/1.04] border-2 border-[#101010] bg-[#ff5c00] shadow-[8px_8px_0_#101010]">
        <span class="pdb-kicker">NEW OBJECT</span>
        <div class="pdb-box">
          <div class="pdb-box-lid"></div>
          <div class="pdb-box-face">
            <span class="pdb-box-logo">W/H</span>
            <span class="pdb-box-text">HEAVY DUTY<br />OBJECT 07</span>
          </div>
          <span class="pdb-box-tape"></span>
        </div>
        <span class="pdb-sticker">{{ discount ? '−' + discount + '%' : 'DROP 07' }}</span>
        <span class="pdb-cross pdb-cross-a">+</span>
        <span class="pdb-cross pdb-cross-b">+</span>
      </section>

      <div class="px-4">
        <div class="pdb-in pdb-d2 flex items-start justify-between gap-3">
          <div>
            <h1 class="text-[42px] font-black uppercase leading-[0.88] tracking-[-0.05em]">{{ p.title }}</h1>
            <p class="mt-3 max-w-[240px] text-[12px] font-medium leading-snug">{{ p.subtitle || '—' }}</p>
          </div>
          <span class="border-2 border-[#101010] bg-white px-2 py-1 font-mono text-[9px] uppercase leading-tight">SKU<br />{{ nfcId || '0007' }}</span>
        </div>

        <div class="pdb-in pdb-d2 mt-5 flex items-stretch border-2 border-[#101010] shadow-[5px_5px_0_#101010]">
          <div class="flex-1 bg-[#101010] px-4 py-3 text-[#f5e900]">
            <p class="font-mono text-[9px] uppercase tracking-[0.2em]">Now</p>
            <p class="mt-1 text-[38px] font-black leading-none tabular-nums">¥{{ p.price || '—' }}</p>
          </div>
          <div class="flex flex-1 flex-col justify-center bg-white px-4 py-3">
            <p class="font-mono text-[9px] uppercase tracking-[0.16em] text-[#101010]/50">Before</p>
            <p class="mt-1 text-[18px] font-bold tabular-nums text-[#101010]/45 line-through">¥{{ p.originalPrice || '—' }}</p>
            <p v-if="discount" class="mt-1 text-[10px] font-black uppercase text-[#ff2d00]">You save {{ discount }}%</p>
          </div>
        </div>

        <section class="pdb-in pdb-d3 mt-5">
          <h2 class="flex items-center gap-2 text-[12px] font-black uppercase tracking-[0.06em]">
            <span class="grid h-5 w-5 place-items-center bg-[#ff2d00] text-[11px] text-white">✓</span> Why it slaps
          </h2>
          <ul class="mt-2 grid grid-cols-2 gap-2">
            <li v-for="(f, i) in features" :key="i" class="border-2 border-[#101010] bg-white p-2.5 text-[11px] font-semibold leading-snug">
              <span class="mb-1 block font-mono text-[8px] text-[#ff2d00]">{{ String(i + 1).padStart(2, '0') }}</span>{{ f }}
            </li>
            <li v-if="!features.length" class="col-span-2 border-2 border-[#101010] bg-white p-2.5 text-[11px]"></li>
          </ul>
        </section>

        <section class="pdb-in pdb-d3 mt-4">
          <h2 class="text-[12px] font-black uppercase tracking-[0.06em]">The boring details</h2>
          <dl class="mt-2 border-2 border-[#101010] bg-white">
            <div v-for="([k, v], i) in specs" :key="k" :class="['flex items-center justify-between px-3 py-2', i ? 'border-t-2 border-[#101010]' : '']">
              <dt class="font-mono text-[9px] uppercase tracking-[0.12em] text-[#101010]/50">{{ k }}</dt>
              <dd class="text-[12px] font-bold">{{ v }}</dd>
            </div>
            <div v-if="!specs.length" class="px-3 py-2 text-[12px]"></div>
          </dl>
        </section>

        <p v-if="p.description" class="pdb-in pdb-d3 mt-4 border-l-4 border-[#ff2d00] pl-3 text-[11.5px] font-medium leading-[1.8]">{{ p.description }}</p>

        <button class="pdb-in pdb-d4 pdb-buy mt-5 flex w-full items-center justify-between border-2 border-[#101010] px-4 py-4" @click="buy">
          <span class="text-[15px] font-black uppercase tracking-[-0.01em]">{{ p.ctaText || '立即购买' }}</span>
          <span class="grid h-8 w-8 place-items-center border-2 border-[#f5e900]"><ArrowRight class="h-4 w-4" /></span>
        </button>

        <p class="py-4 text-center font-mono text-[8px] uppercase tracking-[0.26em] text-[#101010]/45">WARE/HOUSE · UTILITY GOODS · EST. 2026</p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { ArrowRight } from 'lucide-vue-next'

const props = defineProps(['payload', 'nfcId'])
const p = computed(() => props.payload || {})
const features = computed(() => (Array.isArray(p.value.features) ? p.value.features.filter(Boolean) : []))
const specs = computed(() => Object.entries(p.value.specifications || {}))

const discount = computed(() => {
  const price = Number(p.value.price)
  const original = Number(p.value.originalPrice)
  if (!Number.isFinite(price) || !Number.isFinite(original) || original <= price) return null
  return Math.round(((original - price) / original) * 100)
})

function buy() {
  const action = p.value.ctaAction
  if (!action) return
  try {
    const url = new URL(action)
    if (['http:', 'https:'].includes(url.protocol)) window.location.href = url.href
  } catch {
    // 链接不合法时忽略
  }
}
</script>

<style scoped>
.pdb-ticker-track {
  display: flex;
  width: max-content;
  animation: pdbMarquee 18s linear infinite;
}
.pdb-visual { overflow: hidden; }
.pdb-kicker {
  position: absolute;
  left: 10px;
  top: 8px;
  z-index: 2;
  font-family: ui-monospace, monospace;
  font-size: 8px;
  letter-spacing: 0.2em;
  color: #101010;
}
.pdb-box {
  position: absolute;
  left: 50%;
  top: 52%;
  width: 58%;
  height: 48%;
  transform: translate(-50%, -50%) rotate(-4deg);
  filter: drop-shadow(10px 12px 0 rgba(16, 16, 16, 0.25));
  animation: pdbWobble 6s ease-in-out infinite;
}
.pdb-box-lid {
  position: absolute;
  left: -3%;
  top: -8%;
  width: 106%;
  height: 22%;
  background: #fafafa;
  border: 2px solid #101010;
  transform: skewX(-8deg);
}
.pdb-box-face {
  position: absolute;
  inset: 10% 0 0;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  gap: 8px;
  background: #fafafa;
  border: 2px solid #101010;
}
.pdb-box-logo { font-size: 30px; font-weight: 900; letter-spacing: -0.05em; }
.pdb-box-text { font-family: ui-monospace, monospace; font-size: 8px; line-height: 1.5; letter-spacing: 0.16em; text-align: center; }
.pdb-box-tape {
  position: absolute;
  left: 50%;
  top: 8%;
  width: 18%;
  height: 84%;
  transform: translateX(-50%);
  background: rgba(255, 92, 0, 0.55);
  border-inline: 1px dashed rgba(16, 16, 16, 0.35);
}
.pdb-sticker {
  position: absolute;
  right: -6px;
  bottom: 32px;
  display: grid;
  place-items: center;
  width: 68px;
  height: 68px;
  border-radius: 999px;
  border: 2px solid #101010;
  background: #f5e900;
  font-size: 14px;
  font-weight: 900;
  transform: rotate(12deg);
  box-shadow: 4px 4px 0 #101010;
}
.pdb-cross { position: absolute; font-size: 22px; font-weight: 300; line-height: 1; }
.pdb-cross-a { right: 12px; top: 12px; }
.pdb-cross-b { left: 12px; bottom: 12px; }
.pdb-buy {
  background: #ff2d00;
  color: #101010;
  box-shadow: 6px 6px 0 #101010;
  transition: transform 0.16s ease, box-shadow 0.16s ease, background 0.2s ease;
}
.pdb-buy:hover { background: #ff4a22; transform: translate(-2px, -2px); box-shadow: 9px 9px 0 #101010; }
.pdb-buy:active { transform: translate(3px, 3px); box-shadow: 2px 2px 0 #101010; }
.pdb-in { animation: pdbIn 0.6s cubic-bezier(0.22, 1, 0.36, 1) both; }
.pdb-d1 { animation-delay: 0.06s; }
.pdb-d2 { animation-delay: 0.12s; }
.pdb-d3 { animation-delay: 0.18s; }
.pdb-d4 { animation-delay: 0.24s; }
@keyframes pdbIn { from { opacity: 0; transform: translateY(14px); } to { opacity: 1; transform: none; } }
@keyframes pdbMarquee { to { transform: translateX(-25%); } }
@keyframes pdbWobble { 0%, 100% { transform: translate(-50%, -50%) rotate(-4deg); } 50% { transform: translate(-50%, -51%) rotate(-1deg); } }
@media (prefers-reduced-motion: reduce) {
  .pdb-in, .pdb-ticker-track, .pdb-box { animation: none !important; }
  .pdb-buy { transition: none; }
}
</style>
