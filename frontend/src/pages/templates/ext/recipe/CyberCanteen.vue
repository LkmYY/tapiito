<template>
  <div class="relative min-h-screen overflow-hidden bg-[#03100a] text-[#9bf0b6]">
    <div class="cy-scan pointer-events-none absolute inset-0"></div>
    <div class="cy-glow pointer-events-none absolute inset-x-0 top-0 h-80"></div>

    <div class="relative z-10 mx-auto w-full max-w-md px-5 pb-14 pt-8 font-mono">
      <div class="cy-rise flex items-center justify-between text-[8px] uppercase tracking-[0.16em] text-[#4f9a68]">
        <span>guest@canteen:~$ cat menu.txt</span>
        <span v-if="nfcId">TTY{{ nfcId }}</span>
      </div>

      <header class="cy-rise cy-d1 mt-5">
        <div class="cy-boot">
          <p><span class="cy-dim">[ OK ]</span> mounting /dev/kitchen</p>
          <p><span class="cy-dim">[ OK ]</span> loading seasonal.mod</p>
          <p><span class="cy-ok">[ OK ]</span> menu ready</p>
        </div>
        <p class="mt-4 text-[8px] uppercase tracking-[0.24em] text-[#4f9a68]">// CURRENT INSTANCE</p>
        <h1 class="cy-title mt-2 text-[30px] font-bold leading-[1.05] text-[#c8ffd8]">
          {{ payload?.shopName || 'CYBER_CANTEEN' }}<span class="cy-caret">_</span>
        </h1>
        <p v-if="payload?.tagline" class="mt-2 max-w-[19rem] text-[9.5px] leading-[1.7] text-[#6fbf88]">
          <span class="cy-dim">&gt;</span> {{ payload.tagline }}
        </p>
        <div class="mt-3 flex flex-wrap gap-x-4 gap-y-1 text-[7.5px] uppercase tracking-[0.1em] text-[#4f9a68]">
          <span v-if="payload?.hours">UPTIME {{ payload.hours }}</span>
          <span v-if="payload?.address">LOC {{ payload.address }}</span>
        </div>
      </header>

      <div class="cy-rise cy-d2 mt-5 flex items-center gap-2 text-[7.5px] text-[#2f6f45]">
        <span class="cy-line flex-1"></span>
        <span>MENU.TXT — {{ pad(total) }} LINES</span>
        <span class="cy-line flex-1"></span>
      </div>

      <section v-if="groups.length" class="mt-2">
        <div v-for="(group, gi) in groups" :key="gi" class="cy-rise cy-group" :style="{ animationDelay: 0.14 + gi * 0.06 + 's' }">
          <div class="flex items-center gap-2 text-[9px]">
            <span class="cy-dim">[{{ pad(gi + 1) }}]</span>
            <h2 class="font-bold uppercase tracking-[0.12em] text-[#79e69a]">{{ group.name }}</h2>
            <span class="cy-line flex-1"></span>
            <span class="text-[7px] text-[#2f6f45]">{{ pad(group.items.length) }}</span>
          </div>
          <ul class="mt-1">
            <li v-for="(item, i) in group.items" :key="i" class="cy-row">
              <span class="cy-dim flex-shrink-0 text-[8px]">├─</span>
              <div class="min-w-0 flex-1">
                <div class="flex flex-wrap items-baseline gap-x-1.5">
                  <span class="text-[11.5px] text-[#bdf8cd]">{{ item?.name || 'UNNAMED' }}</span>
                  <span v-if="signature(item)" class="cy-tag cy-tag-hot">★SIG</span>
                  <span v-if="veg(item)" class="cy-tag">VEG</span>
                  <span v-if="spicy(item)" class="cy-tag cy-tag-hot">S{{ spicy(item) }}</span>
                </div>
                <p v-if="item?.desc" class="mt-0.5 text-[8px] leading-[1.5] text-[#5da576]"><span class="cy-dim">#</span> {{ item.desc }}</p>
              </div>
              <span class="cy-dots"></span>
              <span v-if="hasPrice(item)" class="flex-shrink-0 text-[10px] text-[#e8f7a0]">¥{{ item.price }}</span>
              <span v-else class="text-[8px] text-[#2f6f45]">--</span>
            </li>
          </ul>
        </div>
      </section>

      <div v-if="!groups.length" class="cy-rise mt-4 border border-[#1f4a31] px-4 py-10 text-center">
        <Terminal class="mx-auto h-5 w-5 text-[#4f9a68]" />
        <p class="mt-2 text-[9px] uppercase tracking-[0.14em] text-[#4f9a68]">menu.txt is empty</p>
        <p class="mt-1 text-[8px] text-[#2f6f45]">$ touch new_menu &amp;&amp; restart kitchen</p>
      </div>

      <div v-if="payload?.notice" class="cy-rise mt-4 border-l-2 border-[#79e69a] bg-[#0a1f14] px-3 py-2.5">
        <p class="text-[8.5px] leading-[1.7] text-[#6fbf88]"><span class="text-[#e8f7a0]">NOTICE:</span> {{ payload.notice }}</p>
      </div>

      <div class="cy-rise cy-d3 mt-4 border border-[#1f4a31] p-3">
        <p class="text-[8px] text-[#2f6f45]">guest@canteen:~$ ./contact --now</p>
        <a v-if="payload?.phone" :href="'tel:' + payload.phone" class="cy-call mt-2">
          <Phone class="h-3.5 w-3.5" />
          <span class="text-[9px] uppercase tracking-[0.1em]">CALL</span>
          <span class="ml-auto text-[10px] text-[#c8ffd8]">{{ payload.phone }}</span>
        </a>
        <p v-else class="mt-1 text-[8px] text-[#2f6f45]">phone: null</p>
        <p class="mt-2 text-[7px] text-[#2f6f45]">exit code 0 · kitchen is listening<span class="cy-caret">_</span></p>
      </div>

      <p class="mt-6 text-center text-[6.5px] uppercase tracking-[0.24em] text-[#2f6f45]">root access · flavor enabled</p>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { Phone, Terminal } from 'lucide-vue-next'

const props = defineProps(['payload', 'nfcId'])

const pad = (n) => String(n).padStart(2, '0')
const hasPrice = (item) => item?.price !== undefined && item?.price !== null && item?.price !== ''
const signature = (item) => item?.signature === true || (Array.isArray(item?.tags) && item.tags.some((t) => /招牌|必点|signature/i.test(String(t))))
const veg = (item) => item?.vegetarian === true || (Array.isArray(item?.tags) && item.tags.some((t) => /素|vegan|vegetarian/i.test(String(t))))
const spicy = (item) => { const n = Number(item?.spicy); return Number.isFinite(n) ? Math.max(0, Math.min(3, Math.round(n))) : 0 }

const groups = computed(() => {
  const p = props.payload || {}
  const list = Array.isArray(p.categories) ? p.categories : []
  return list.map((c) => ({
    name: (c && c.name) || 'UNNAMED_SECTION',
    items: (Array.isArray(c?.items) ? c.items : []).filter((i) => i && (i.name || i.price)),
  })).filter((g) => g.items.length)
})
const total = computed(() => groups.value.reduce((n, g) => n + g.items.length, 0))
</script>

<style scoped>
.cy-scan {
  background: repeating-linear-gradient(180deg, rgba(155, 240, 182, 0.045) 0 1px, transparent 1px 3px);
  pointer-events: none;
}
.cy-glow { background: radial-gradient(60% 90% at 50% 0%, rgba(75, 220, 120, 0.14), transparent 70%); }
.cy-boot { border: 1px solid #1f4a31; background: #061a0e; padding: 6px 9px; font-size: 8px; line-height: 1.7; color: #4f9a68; }
.cy-dim { color: #2f6f45; }
.cy-ok { color: #e8f7a0; }
.cy-title { font-family: ui-monospace, 'Cascadia Mono', 'Courier New', monospace; letter-spacing: 0.02em; text-shadow: 0 0 14px rgba(120, 240, 150, 0.35); }
.cy-caret { display: inline-block; margin-left: 2px; color: #e8f7a0; animation: cy-blink 1s steps(2, end) infinite; }
.cy-line { height: 1px; background: repeating-linear-gradient(90deg, #2f6f45 0 4px, transparent 4px 8px); }
.cy-group { padding-top: 16px; }
.cy-row { display: flex; align-items: baseline; gap: 6px; border-bottom: 1px solid #0f2b1c; padding: 7px 0; }
.cy-row:last-child { border-bottom: none; }
.cy-dots { flex: 1; min-width: 8px; border-bottom: 1px dotted #1f4a31; transform: translateY(-2px); }
.cy-tag { border: 1px solid #2f6f45; padding: 0 2px; font-size: 5.5px; letter-spacing: 0.06em; color: #79e69a; }
.cy-tag-hot { border-color: #86a13c; color: #e8f7a0; }
.cy-call {
  display: flex;
  align-items: center;
  gap: 8px;
  border: 1px solid #2f6f45;
  background: #0a2213;
  padding: 7px 9px;
  color: #9bf0b6;
  transition: border-color 0.2s ease, background-color 0.2s ease;
}
.cy-call:hover { border-color: #79e69a; background: #0e2e1a; }
.cy-rise { animation: cy-in 0.55s ease-out both; }
.cy-d1 { animation-delay: 0.07s; }
.cy-d2 { animation-delay: 0.14s; }
.cy-d3 { animation-delay: 0.28s; }
@keyframes cy-in {
  from { opacity: 0; transform: translateY(6px); }
  to { opacity: 1; transform: none; }
}
@keyframes cy-blink { 0%, 45% { opacity: 1; } 50%, 100% { opacity: 0; } }
@media (prefers-reduced-motion: reduce) {
  .cy-rise, .cy-caret, .cy-call { animation: none !important; transition: none !important; }
}
</style>
