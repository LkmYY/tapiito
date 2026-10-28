<template>
  <div class="pdai-root min-h-screen bg-[#e8e1d3] text-[#2a2b28]">
    <div class="mx-auto w-full max-w-md px-6 pb-14 pt-7">
      <header class="pdai-in flex items-center justify-between border-b border-[#2a2b28]/30 pb-2 font-mono text-[9px] uppercase tracking-[0.2em]">
        <span>Municipal Archive</span>
        <span>Box 07 / File {{ nfcId || '128' }}</span>
      </header>

      <div class="pdai-in pdai-d1 mt-4 flex items-start justify-between">
        <div>
          <p class="font-mono text-[8px] uppercase tracking-[0.2em] text-[#9a3412]">Object record · public copy</p>
          <h1 class="mt-2 max-w-[280px] text-[29px] font-semibold leading-[1.1] tracking-[-0.02em]">{{ p.title }}</h1>
          <p class="mt-1.5 font-mono text-[10px] text-[#2a2b28]/50">{{ p.subtitle || '—' }}</p>
        </div>
        <div class="pdai-stamp">FILED<br />2026</div>
      </div>

      <section class="pdai-in pdai-d1 pdai-photo mt-5">
        <div class="pdai-photo-corner pdai-photo-corner-a"></div>
        <div class="pdai-photo-corner pdai-photo-corner-b"></div>
        <div class="pdai-art">
          <div class="pdai-oval"></div>
          <div class="pdai-record">
            <span class="pdai-record-groove"></span>
            <span class="pdai-record-label">ARCH<br />07</span>
          </div>
          <span class="pdai-anno pdai-anno-a">FIG. A</span>
          <span class="pdai-anno pdai-anno-b">SCALE — 1:2</span>
        </div>
        <div class="flex items-center justify-between px-2 py-1.5 font-mono text-[8px] uppercase tracking-[0.18em] text-[#2a2b28]/45">
          <span>Negative no. 26–128</span>
          <span>Do not remove</span>
        </div>
      </section>

      <div class="pdai-in pdai-d2 mt-4 flex items-end justify-between border-y-2 border-[#2a2b28] py-3">
        <div>
          <p class="font-mono text-[8px] uppercase tracking-[0.16em] text-[#2a2b28]/45">Appraised value</p>
          <p class="mt-1 text-[38px] font-semibold leading-none tabular-nums">¥{{ p.price || '—' }}</p>
        </div>
        <div class="text-right font-mono text-[9px] leading-relaxed">
          <p v-if="p.originalPrice" class="text-[#2a2b28]/40">LIST <span class="line-through">¥{{ p.originalPrice }}</span></p>
          <p v-if="discount" class="mt-1 bg-[#9a3412] px-2 py-0.5 text-[#e8e1d3]">REVISED −{{ discount }}%</p>
          <p class="mt-1 text-[#2a2b28]/35">REF. 2026/07</p>
        </div>
      </div>

      <section class="pdai-in pdai-d3 mt-4">
        <div class="flex items-center justify-between">
          <h2 class="font-mono text-[9px] uppercase tracking-[0.2em]">I. Condition report</h2>
          <span class="font-mono text-[8px] text-[#2a2b28]/35">PP. 01–02</span>
        </div>
        <ul class="mt-2 border border-[#2a2b28]/20 bg-[#f3eee2]">
          <li v-for="(f, i) in features" :key="i" class="flex items-baseline gap-3 border-b border-dashed border-[#2a2b28]/20 px-3 py-2.5 last:border-0">
            <span class="font-mono text-[8px] text-[#9a3412]">{{ String(i + 1).padStart(2, '0') }}</span>
            <span class="flex-1 text-[11.5px] leading-snug">{{ f }}</span>
            <span class="h-2 w-2 rounded-full bg-[#3f6212]"></span>
          </li>
          <li v-if="!features.length" class="px-3 py-2.5 text-[11.5px] text-[#2a2b28]/40"></li>
        </ul>
      </section>

      <section class="pdai-in pdai-d3 mt-4">
        <h2 class="font-mono text-[9px] uppercase tracking-[0.2em]">II. Catalogue data</h2>
        <dl class="mt-2 grid grid-cols-2 gap-px border border-[#2a2b28]/20 bg-[#2a2b28]/20">
          <div v-for="[k, v] in specs" :key="k" class="bg-[#f3eee2] px-3 py-2.5">
            <dt class="font-mono text-[8px] uppercase tracking-[0.1em] text-[#2a2b28]/45">{{ k }}</dt>
            <dd class="mt-1 text-[11.5px] font-medium">{{ v }}</dd>
          </div>
          <div v-if="!specs.length" class="col-span-2 bg-[#f3eee2] px-3 py-2.5 text-[11.5px] text-[#2a2b28]/40"></div>
        </dl>
      </section>

      <section v-if="p.description" class="pdai-in pdai-d3 mt-4">
        <h2 class="font-mono text-[9px] uppercase tracking-[0.2em]">III. Curator's note</h2>
        <p class="mt-1.5 border-l-2 border-[#9a3412] pl-3 text-[11.5px] leading-[1.85] text-[#2a2b28]/65">{{ p.description }}</p>
      </section>

      <button class="pdai-in pdai-d4 pdai-buy mt-5 flex w-full items-center justify-between border border-[#2a2b28] px-4 py-3.5" @click="buy">
        <span class="font-mono text-[10px] uppercase tracking-[0.16em]">{{ p.ctaText || '立即购买' }}</span>
        <span class="flex items-center gap-2 font-mono text-[9px] uppercase">Request object <ArrowUpRight class="h-4 w-4" /></span>
      </button>

      <div class="mt-4 flex items-center justify-between font-mono text-[7.5px] uppercase tracking-[0.16em] text-[#2a2b28]/35">
        <span>Archive dept. · shelf B</span>
        <span>NFC {{ nfcId || '—' }}</span>
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
.pdai-stamp {
  display: grid;
  place-items: center;
  width: 56px;
  height: 56px;
  flex: none;
  border: 2px solid rgba(154, 52, 18, 0.75);
  color: rgba(154, 52, 18, 0.8);
  font-family: ui-monospace, monospace;
  font-size: 8px;
  line-height: 1.4;
  letter-spacing: 0.12em;
  text-align: center;
  transform: rotate(7deg);
}
.pdai-photo { position: relative; border: 7px solid #f3eee2; background: #f3eee2; box-shadow: 0 14px 28px -22px rgba(42, 43, 40, 0.6); }
.pdai-photo::after { content: ''; position: absolute; inset: 0; box-shadow: inset 0 0 0 1px rgba(42, 43, 40, 0.2); pointer-events: none; }
.pdai-photo-corner { position: absolute; width: 13px; height: 13px; background: rgba(42, 43, 40, 0.35); z-index: 3; }
.pdai-photo-corner-a { left: -4px; top: -4px; clip-path: polygon(0 0, 100% 0, 0 100%); }
.pdai-photo-corner-b { right: -4px; bottom: 22px; clip-path: polygon(100% 0, 100% 100%, 0 100%); }
.pdai-art { position: relative; aspect-ratio: 4 / 3.3; overflow: hidden; background: #d8d2c4; border: 1px solid rgba(42, 43, 40, 0.2); }
.pdai-art::before {
  content: '';
  position: absolute;
  inset: 0;
  background-image: linear-gradient(rgba(42, 43, 40, 0.06) 1px, transparent 1px), linear-gradient(90deg, rgba(42, 43, 40, 0.06) 1px, transparent 1px);
  background-size: 24px 24px;
}
.pdai-oval { position: absolute; left: 50%; top: 50%; width: 66%; height: 78%; transform: translate(-50%, -50%); border-radius: 50%; border: 1px solid rgba(42, 43, 40, 0.16); background: rgba(243, 238, 226, 0.5); }
.pdai-record { position: absolute; left: 50%; top: 50%; width: 38%; aspect-ratio: 1; transform: translate(-50%, -50%); border-radius: 999px; background: repeating-radial-gradient(circle at 50% 50%, #2a2b28 0 2px, #1b1c19 2px 4px); box-shadow: 0 14px 24px -14px rgba(42, 43, 40, 0.7), inset 0 0 0 1px rgba(255, 255, 255, 0.12); animation: pdaiSpin 24s linear infinite; }
.pdai-record-groove { position: absolute; inset: 9%; border-radius: 999px; border: 1px solid rgba(243, 238, 226, 0.12); }
.pdai-record-label { position: absolute; left: 50%; top: 50%; width: 36%; aspect-ratio: 1; transform: translate(-50%, -50%); border-radius: 999px; background: #9a3412; color: #f3eee2; display: grid; place-items: center; font-family: ui-monospace, monospace; font-size: 7px; line-height: 1.4; letter-spacing: 0.1em; text-align: center; }
.pdai-anno { position: absolute; font-family: ui-monospace, monospace; font-size: 7px; letter-spacing: 0.14em; color: rgba(42, 43, 40, 0.55); border-bottom: 1px solid rgba(42, 43, 40, 0.3); padding-bottom: 1px; }
.pdai-anno-a { left: 10px; top: 10px; }
.pdai-anno-b { right: 10px; bottom: 10px; }
.pdai-buy { background: transparent; transition: background 0.25s ease, color 0.25s ease; }
.pdai-buy:hover { background: #2a2b28; color: #e8e1d3; }
.pdai-buy:active { transform: translateY(1px); }
.pdai-in { animation: pdaiIn 0.75s cubic-bezier(0.22, 1, 0.36, 1) both; }
.pdai-d1 { animation-delay: 0.06s; }
.pdai-d2 { animation-delay: 0.13s; }
.pdai-d3 { animation-delay: 0.2s; }
.pdai-d4 { animation-delay: 0.27s; }
@keyframes pdaiIn { from { opacity: 0; transform: translateY(16px); } to { opacity: 1; transform: none; } }
@keyframes pdaiSpin { to { transform: translate(-50%, -50%) rotate(360deg); } }
@media (prefers-reduced-motion: reduce) {
  .pdai-in, .pdai-record { animation: none !important; }
  .pdai-buy { transition: none; }
}
</style>
