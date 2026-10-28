<template>
  <div class="pde-root min-h-screen bg-[#f6f1e6] text-[#1f1b16]">
    <div class="mx-auto w-full max-w-md pb-14">
      <header class="pde-in flex items-center justify-between border-b border-[#1f1b16] px-6 py-4 font-mono text-[9px] uppercase tracking-[0.24em]">
        <span>Form Journal</span>
        <span>Issue 04 — {{ nfcId || '—' }}</span>
      </header>

      <div class="px-6">
        <div class="pde-in pde-d1 mt-6 flex items-end justify-between">
          <span class="font-mono text-[10px] uppercase tracking-[0.18em] text-[#9a3412]">Objects of desire</span>
          <span class="font-mono text-[10px] text-[#1f1b16]/35">P. 018</span>
        </div>

        <article class="pde-in pde-d2 mt-3 grid grid-cols-12 gap-x-4">
          <h1 class="col-span-12 font-serif text-[52px] leading-[0.94] tracking-[-0.04em]">
            {{ p.title }}
          </h1>
          <p class="col-span-7 mt-4 border-t border-[#1f1b16] pt-3 font-serif text-[16px] italic leading-snug text-[#1f1b16]/70">{{ p.subtitle || '—' }}</p>
          <div class="col-span-5 mt-4 border-t border-[#1f1b16] pt-3 text-right">
            <p class="font-mono text-[9px] uppercase tracking-[0.16em] text-[#1f1b16]/45">Price</p>
            <p class="mt-1 font-serif text-[27px] leading-none tabular-nums">¥{{ p.price || '—' }}</p>
            <p v-if="p.originalPrice" class="mt-1.5 font-mono text-[10px] text-[#1f1b16]/40 line-through">¥{{ p.originalPrice }}</p>
          </div>
        </article>

        <section class="pde-in pde-d2 pde-visual mt-7">
          <div class="pde-sun"></div>
          <div class="pde-object">
            <span class="pde-neck"></span>
            <span class="pde-body"></span>
            <span class="pde-cap"></span>
            <span class="pde-label">FORM<br />NO. 04</span>
          </div>
          <div class="pde-caption">
            <span>FIG. 01</span>
            <span>Ceramic / matte glaze</span>
          </div>
          <span class="pde-stamp">{{ discount ? '-' + discount + '%' : 'NEW' }}</span>
        </section>

        <section class="pde-in pde-d3 mt-8 grid grid-cols-12 gap-5">
          <div class="col-span-4">
            <h2 class="border-t-2 border-[#9a3412] pt-2 font-mono text-[9px] uppercase tracking-[0.2em] text-[#9a3412]">The brief</h2>
            <p class="mt-3 font-serif text-[13px] leading-[1.9] text-[#1f1b16]/70">{{ p.description }}</p>
          </div>
          <div class="col-span-8">
            <h2 class="border-t-2 border-[#1f1b16] pt-2 font-mono text-[9px] uppercase tracking-[0.2em]">Why we love it</h2>
            <ul class="mt-1">
              <li v-for="(f, i) in features" :key="i" class="flex gap-3 border-b border-[#1f1b16]/15 py-3">
                <span class="font-serif text-[13px] italic text-[#9a3412]">{{ String.fromCharCode(97 + i) }}.</span>
                <span class="text-[12.5px] leading-snug">{{ f }}</span>
              </li>
              <li v-if="!features.length" class="border-b border-[#1f1b16]/15 py-3 text-[12px] text-[#1f1b16]/40"></li>
            </ul>
          </div>
        </section>

        <section class="pde-in pde-d3 mt-8">
          <div class="flex items-baseline justify-between border-t-2 border-[#1f1b16] pt-2">
            <h2 class="font-mono text-[9px] uppercase tracking-[0.2em]">The details</h2>
            <span class="font-mono text-[9px] text-[#1f1b16]/40">Specification</span>
          </div>
          <dl class="mt-2 grid grid-cols-2 gap-x-6">
            <div v-for="[k, v] in specs" :key="k" class="flex items-baseline justify-between border-b border-[#1f1b16]/15 py-3">
              <dt class="font-mono text-[9px] uppercase tracking-[0.1em] text-[#1f1b16]/45">{{ k }}</dt>
              <dd class="font-serif text-[14px]">{{ v }}</dd>
            </div>
            <div v-if="!specs.length" class="col-span-2 border-b border-[#1f1b16]/15 py-3 font-serif text-[12px] text-[#1f1b16]/40"></div>
          </dl>
        </section>

        <button class="pde-in pde-d4 pde-buy mt-7 flex w-full items-center justify-between border border-[#1f1b16] px-5 py-4" @click="buy">
          <span class="font-serif text-[16px] italic">{{ p.ctaText || '立即购买' }}</span>
          <span class="flex items-center gap-2 font-mono text-[9px] uppercase tracking-[0.18em]">Order now <ArrowUpRight class="h-3.5 w-3.5" /></span>
        </button>

        <p class="mt-5 text-center font-mono text-[8px] uppercase tracking-[0.28em] text-[#1f1b16]/30">Printed in small batches · Form Journal 2026</p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { ArrowUpRight } from 'lucide-vue-next'

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
.pde-visual {
  position: relative;
  aspect-ratio: 4 / 5;
  overflow: hidden;
  background: linear-gradient(165deg, #efe4ce, #ddcdae 62%, #cbb894);
}
.pde-sun {
  position: absolute;
  left: 50%;
  top: 44%;
  width: 74%;
  aspect-ratio: 1;
  transform: translate(-50%, -50%);
  border-radius: 999px;
  background: #c2410c;
  opacity: 0.92;
}
.pde-object {
  position: absolute;
  left: 50%;
  top: 52%;
  width: 34%;
  height: 62%;
  transform: translate(-50%, -50%);
  filter: drop-shadow(16px 22px 18px rgba(60, 35, 15, 0.3));
  animation: pdeSway 7s ease-in-out infinite;
}
.pde-body {
  position: absolute;
  inset: 22% 0 0;
  border-radius: 22% 22% 30% 30% / 12% 12% 34% 34%;
  background: linear-gradient(100deg, #f3ead8 0%, #e5d8bd 38%, #c9b48f 100%);
}
.pde-neck {
  position: absolute;
  left: 34%;
  top: 4%;
  width: 32%;
  height: 24%;
  border-radius: 6px 6px 0 0;
  background: linear-gradient(90deg, #d8c8a9, #f1e8d6 50%, #bfa780);
}
.pde-cap {
  position: absolute;
  left: 30%;
  top: 0;
  width: 40%;
  height: 7%;
  border-radius: 999px;
  background: #2d2820;
}
.pde-label {
  position: absolute;
  left: 50%;
  top: 54%;
  transform: translateX(-50%);
  font-family: ui-monospace, monospace;
  font-size: 8px;
  line-height: 1.5;
  letter-spacing: 0.18em;
  text-align: center;
  color: #4b4234;
}
.pde-caption {
  position: absolute;
  inset-inline: 12px;
  bottom: 10px;
  display: flex;
  justify-content: space-between;
  font-family: ui-monospace, monospace;
  font-size: 8px;
  letter-spacing: 0.14em;
  color: rgba(41, 32, 22, 0.6);
  border-top: 1px solid rgba(41, 32, 22, 0.25);
  padding-top: 6px;
}
.pde-stamp {
  position: absolute;
  right: 16px;
  top: 16px;
  display: grid;
  place-items: center;
  width: 52px;
  height: 52px;
  border: 1px solid #1f1b16;
  border-radius: 999px;
  font-family: ui-monospace, monospace;
  font-size: 10px;
  letter-spacing: 0.06em;
  color: #1f1b16;
  transform: rotate(9deg);
  background: rgba(246, 241, 230, 0.7);
}
.pde-buy { position: relative; overflow: hidden; background: transparent; transition: color 0.3s ease; }
.pde-buy::before {
  content: '';
  position: absolute;
  inset: 0;
  background: #1f1b16;
  transform: translateY(101%);
  transition: transform 0.35s cubic-bezier(0.22, 1, 0.36, 1);
}
.pde-buy:hover { color: #f6f1e6; }
.pde-buy:hover::before { transform: none; }
.pde-buy:active { transform: scale(0.99); }
.pde-buy > * { position: relative; z-index: 1; }
.pde-in { animation: pdeIn 0.75s cubic-bezier(0.22, 1, 0.36, 1) both; }
.pde-d1 { animation-delay: 0.05s; }
.pde-d2 { animation-delay: 0.12s; }
.pde-d3 { animation-delay: 0.2s; }
.pde-d4 { animation-delay: 0.28s; }
@keyframes pdeIn { from { opacity: 0; transform: translateY(16px); } to { opacity: 1; transform: none; } }
@keyframes pdeSway { 0%, 100% { transform: translate(-50%, -50%) rotate(-3deg); } 50% { transform: translate(-50%, -51%) rotate(3deg); } }
@media (prefers-reduced-motion: reduce) {
  .pde-in, .pde-object { animation: none !important; }
  .pde-buy, .pde-buy::before { transition: none; }
}
</style>
