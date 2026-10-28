<template>
  <div class="pdan-root min-h-screen bg-[#0b0a08] text-[#efe8d8]">
    <div class="pdan-glow pointer-events-none"></div>
    <div class="pdan-shell mx-auto w-full max-w-md px-8 pb-16 pt-11">
      <header class="pdan-in flex items-center justify-between font-mono text-[9px] uppercase tracking-[0.34em] text-[#c9a96a]/70">
        <span>Atelier Noir</span>
        <span>{{ nfcId ? 'No. ' + nfcId : 'Maison' }}</span>
      </header>

      <div class="pdan-in pdan-d1 pdan-frame mt-10">
        <div class="pdan-corner pdan-corner-tl"></div>
        <div class="pdan-corner pdan-corner-tr"></div>
        <div class="pdan-corner pdan-corner-bl"></div>
        <div class="pdan-corner pdan-corner-br"></div>
        <div class="pdan-stage">
          <span class="pdan-mark">AN</span>
          <div class="pdan-watch">
            <span class="pdan-strap pdan-strap-t"></span>
            <span class="pdan-dial">
              <span class="pdan-ticks"></span>
              <span class="pdan-hand pdan-hand-h"></span>
              <span class="pdan-hand pdan-hand-m"></span>
              <span class="pdan-pin"></span>
            </span>
            <span class="pdan-strap pdan-strap-b"></span>
          </div>
          <span class="pdan-edition">ÉDITION 1888 · 38MM</span>
        </div>
      </div>

      <div class="pdan-in pdan-d2 mt-9 text-center">
        <p class="font-mono text-[9px] uppercase tracking-[0.4em] text-[#c9a96a]/80">The Signature</p>
        <h1 class="mt-4 font-serif text-[36px] leading-[1.14] tracking-tight">{{ p.title }}</h1>
        <p class="mt-3 text-[12px] tracking-[0.12em] text-[#efe8d8]/50">{{ p.subtitle || '—' }}</p>
        <div class="pdan-rule mx-auto mt-6"><i></i></div>
      </div>

      <div class="pdan-in pdan-d3 mt-7 text-center">
        <div class="flex items-baseline justify-center gap-4">
          <span class="font-serif text-[42px] leading-none tabular-nums"><em class="mr-1.5 font-sans text-[18px] not-italic text-[#c9a96a]">¥</em>{{ p.price || '—' }}</span>
          <span v-if="p.originalPrice" class="text-[14px] text-[#efe8d8]/30 line-through tabular-nums">¥{{ p.originalPrice }}</span>
        </div>
        <p class="mt-3 font-mono text-[9px] uppercase tracking-[0.28em] text-[#efe8d8]/30">
          {{ discount ? 'Privilege −' + discount + '%' : 'Tax included · Complimentary delivery' }}
        </p>
      </div>

      <p v-if="p.description" class="pdan-in pdan-d3 mt-8 text-center text-[12.5px] leading-[2.05] text-[#efe8d8]/55">{{ p.description }}</p>

      <section class="pdan-in pdan-d4 mt-11">
        <header class="flex items-center gap-4">
          <span class="h-px flex-1 bg-[#c9a96a]/25"></span>
          <h2 class="font-mono text-[9px] uppercase tracking-[0.34em] text-[#c9a96a]/80">Highlights</h2>
          <span class="h-px flex-1 bg-[#c9a96a]/25"></span>
        </header>
        <ul class="mt-4">
          <li v-for="(f, i) in features" :key="i" class="flex items-baseline gap-4 border-b border-[#efe8d8]/8 py-4">
            <span class="font-mono text-[9px] tracking-[0.18em] text-[#c9a96a]/75">{{ String(i + 1).padStart(2, '0') }}</span>
            <span class="flex-1 text-[12.5px] leading-relaxed text-[#efe8d8]/75">{{ f }}</span>
          </li>
          <li v-if="!features.length" class="border-b border-[#efe8d8]/8 py-4 text-center text-[12px] text-[#efe8d8]/35"></li>
        </ul>
      </section>

      <section class="pdan-in pdan-d4 mt-9">
        <dl class="grid grid-cols-2 gap-px overflow-hidden bg-[#efe8d8]/10">
          <div v-for="[k, v] in specs" :key="k" class="bg-[#0b0a08] px-4 py-4 text-center">
            <dt class="font-mono text-[8.5px] uppercase tracking-[0.22em] text-[#efe8d8]/35">{{ k }}</dt>
            <dd class="mt-1.5 font-serif text-[15px] text-[#efe8d8]/90">{{ v }}</dd>
          </div>
          <div v-if="!specs.length" class="col-span-2 bg-[#0b0a08] px-4 py-4 text-center text-[12px] text-[#efe8d8]/35"></div>
        </dl>
      </section>

      <button class="pdan-in pdan-d5 pdan-cta mt-10 flex w-full items-center justify-center gap-3 py-4" @click="buy">
        <span class="font-mono text-[10px] uppercase tracking-[0.32em]">{{ p.ctaText || '立即购买' }}</span>
        <ArrowUpRight class="h-4 w-4" />
      </button>

      <footer class="mt-7 text-center font-mono text-[8px] uppercase tracking-[0.3em] text-[#efe8d8]/22">
        Geneva · Since 1888 <span v-if="nfcId">· NFC {{ nfcId }}</span>
      </footer>
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
.pdan-root { position: relative; overflow: hidden; }
.pdan-glow {
  position: absolute;
  inset: -20% -40% auto;
  height: 70%;
  background: radial-gradient(50% 50% at 50% 40%, rgba(201, 169, 106, 0.18), transparent 72%);
  pointer-events: none;
}
.pdan-shell { position: relative; z-index: 1; }
.pdan-frame {
  position: relative;
  padding: 10px;
  border: 1px solid rgba(201, 169, 106, 0.35);
  background: rgba(201, 169, 106, 0.03);
}
.pdan-corner {
  position: absolute;
  width: 10px;
  height: 10px;
  border-color: #c9a96a;
  border-style: solid;
  z-index: 2;
}
.pdan-corner-tl { left: -4px; top: -4px; border-width: 1px 0 0 1px; }
.pdan-corner-tr { right: -4px; top: -4px; border-width: 1px 1px 0 0; }
.pdan-corner-bl { left: -4px; bottom: -4px; border-width: 0 0 1px 1px; }
.pdan-corner-br { right: -4px; bottom: -4px; border-width: 0 1px 1px 0; }
.pdan-stage {
  position: relative;
  display: flex;
  align-items: center;
  justify-content: center;
  aspect-ratio: 1 / 1.14;
  overflow: hidden;
  background:
    radial-gradient(60% 46% at 50% 42%, rgba(201, 169, 106, 0.18), transparent 75%),
    linear-gradient(160deg, #15130f, #080706 75%);
}
.pdan-stage::after {
  content: '';
  position: absolute;
  inset: 14px;
  border: 1px solid rgba(201, 169, 106, 0.12);
  pointer-events: none;
}
.pdan-mark {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -52%);
  font-family: Georgia, 'Times New Roman', serif;
  font-size: 150px;
  line-height: 1;
  color: rgba(201, 169, 106, 0.055);
  letter-spacing: -0.06em;
}
.pdan-watch {
  position: relative;
  display: flex;
  flex-direction: column;
  align-items: center;
  animation: pdanRise 1.1s cubic-bezier(0.22, 1, 0.36, 1) both;
}
.pdan-strap {
  width: 54px;
  height: 92px;
  background: repeating-linear-gradient(90deg, #2b1f18 0 2px, #3a2a20 2px 4px);
  box-shadow: inset 0 0 0 1px rgba(201, 169, 106, 0.16);
}
.pdan-strap-t { border-radius: 10px 10px 0 0; border-bottom: 0; }
.pdan-strap-b { border-radius: 0 0 10px 10px; border-top: 0; }
.pdan-dial {
  position: relative;
  width: 128px;
  height: 128px;
  border-radius: 999px;
  background:
    radial-gradient(circle at 38% 30%, rgba(255, 255, 255, 0.16), transparent 46%),
    radial-gradient(circle at 50% 50%, #1a1712 58%, #0a0908 100%);
  border: 5px solid #c9a96a;
  box-shadow: 0 0 0 1px #6e5a36, 0 18px 34px -18px rgba(0, 0, 0, 0.9), inset 0 0 22px rgba(0, 0, 0, 0.8);
  z-index: 2;
}
.pdan-ticks {
  position: absolute;
  inset: 9px;
  border-radius: 999px;
  background: repeating-conic-gradient(from 0deg, rgba(201, 169, 106, 0.55) 0deg 0.8deg, transparent 0.8deg 30deg);
  mask: radial-gradient(circle, transparent 72%, black 73%);
}
.pdan-hand { position: absolute; left: 50%; transform-origin: 50% 100%; border-radius: 4px; background: #e9d8b4; }
.pdan-hand-h { top: 28%; height: 24%; width: 2.5px; transform: translateX(-50%) rotate(38deg); }
.pdan-hand-m { top: 18%; height: 34%; width: 1.8px; transform: translateX(-50%) rotate(-64deg); background: #c9a96a; }
.pdan-pin { position: absolute; left: 50%; top: 50%; width: 7px; height: 7px; transform: translate(-50%, -50%); border-radius: 999px; background: #c9a96a; }
.pdan-edition {
  position: absolute;
  bottom: 24px;
  font-family: ui-monospace, monospace;
  font-size: 8px;
  letter-spacing: 0.28em;
  color: rgba(201, 169, 106, 0.55);
}
.pdan-rule { display: flex; align-items: center; justify-content: center; width: 120px; height: 1px; background: linear-gradient(90deg, transparent, rgba(201, 169, 106, 0.5), transparent); }
.pdan-rule i { width: 4px; height: 4px; transform: rotate(45deg); background: #c9a96a; }
.pdan-cta {
  border: 1px solid #c9a96a;
  background: transparent;
  color: #e8d7b3;
  transition: background 0.3s ease, color 0.3s ease, letter-spacing 0.3s ease;
}
.pdan-cta:hover { background: #c9a96a; color: #0b0a08; }
.pdan-cta:active { transform: scale(0.985); }

.pdan-in { animation: pdanIn 0.8s cubic-bezier(0.22, 1, 0.36, 1) both; }
.pdan-d1 { animation-delay: 0.08s; }
.pdan-d2 { animation-delay: 0.16s; }
.pdan-d3 { animation-delay: 0.24s; }
.pdan-d4 { animation-delay: 0.32s; }
.pdan-d5 { animation-delay: 0.4s; }
@keyframes pdanIn { from { opacity: 0; transform: translateY(18px); } to { opacity: 1; transform: none; } }
@keyframes pdanRise { from { opacity: 0; transform: translateY(24px) scale(0.96); } to { opacity: 1; transform: none; } }
@media (prefers-reduced-motion: reduce) {
  .pdan-in, .pdan-watch { animation: none !important; }
  .pdan-cta { transition: none; }
}
</style>
