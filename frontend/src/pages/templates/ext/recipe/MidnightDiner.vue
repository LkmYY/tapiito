<template>
  <div class="relative min-h-screen overflow-hidden bg-[#0b1020] text-[#f2ead9]">
    <div class="md-glow pointer-events-none absolute inset-x-0 top-0 h-[420px]"></div>
    <div class="md-grain pointer-events-none absolute inset-0"></div>
    <div class="pointer-events-none absolute left-0 top-0 h-full w-9 border-r border-[#f2ead9]/10 bg-[#080c18]/60"></div>
    <p class="pointer-events-none absolute left-2 top-10 font-serif text-[10px] tracking-[0.3em] text-[#c7a26a]/60" style="writing-mode: vertical-rl">深夜食堂 · 营业中</p>

    <div class="relative z-10 mx-auto w-full max-w-md px-6 pb-14 pt-10 pl-14">
      <div class="md-rise flex items-center justify-between font-mono text-[8.5px] uppercase tracking-[0.26em] text-[#8d806c]">
        <span class="inline-flex items-center gap-2"><span class="md-lantern"></span>Open after dark</span>
        <span v-if="nfcId">NFC {{ nfcId }}</span>
      </div>

      <header class="md-rise md-d1 mt-8">
        <p class="font-serif text-[11px] italic tracking-[0.18em] text-[#c7a26a]">A small kitchen, a long night.</p>
        <h1 class="md-serif mt-1 text-[42px] font-semibold leading-[1.05] tracking-[0.02em] text-[#faf3e3]">
          {{ payload?.shopName || '深夜食堂' }}
        </h1>
        <p v-if="payload?.tagline" class="mt-3 max-w-[18rem] text-[11.5px] leading-[1.8] text-[#a99a85]">{{ payload.tagline }}</p>

        <div class="mt-5 flex flex-wrap items-center gap-x-5 gap-y-2 border-y border-[#f2ead9]/12 py-2.5 font-mono text-[9px] tracking-[0.08em] text-[#a99a85]">
          <span v-if="payload?.hours" class="inline-flex items-center gap-1.5"><Clock class="h-3 w-3 text-[#c7a26a]" />{{ payload.hours }}</span>
          <span v-if="payload?.phone" class="inline-flex items-center gap-1.5"><Phone class="h-3 w-3 text-[#c7a26a]" />{{ payload.phone }}</span>
        </div>
      </header>

      <div class="md-rise md-d2 mt-7 flex items-center gap-3">
        <span class="h-px w-6 bg-[#b4472b]"></span>
        <h2 class="md-serif text-[13px] tracking-[0.28em] text-[#e6d2ac]">今 夜 菜 单</h2>
        <span class="h-px flex-1 bg-[#f2ead9]/12"></span>
      </div>

      <section v-if="groups.length" class="mt-4">
        <div v-for="(group, gi) in groups" :key="gi" class="md-rise md-group" :style="{ animationDelay: 0.14 + gi * 0.07 + 's' }">
          <div class="flex items-baseline gap-3">
            <span class="font-mono text-[8px] text-[#b4472b]">{{ pad(gi + 1) }}</span>
            <h3 class="md-serif text-[16px] tracking-[0.12em] text-[#e9d5af]">{{ group.name }}</h3>
            <span class="h-px flex-1 bg-[#f2ead9]/8"></span>
          </div>
          <ul class="mt-1">
            <li v-for="(item, i) in group.items" :key="i" class="md-item">
              <div class="min-w-0 flex-1">
                <div class="flex flex-wrap items-baseline gap-x-2 gap-y-0.5">
                  <span class="md-serif text-[16.5px] text-[#f2ead9]">{{ item?.name }}</span>
                  <span v-if="signature(item)" class="md-pick">今夜必点</span>
                  <Flame v-for="n in spicy(item)" :key="n" class="h-3 w-3 text-[#c05a3a]" />
                </div>
                <p v-if="item?.desc" class="mt-1 text-[10.5px] leading-[1.65] text-[#958875]">{{ item.desc }}</p>
              </div>
              <span class="md-dots"></span>
              <span v-if="hasPrice(item)" class="flex-shrink-0 font-mono text-[13px] tabular-nums text-[#d9b97c]">¥{{ item.price }}</span>
              <span v-else class="font-mono text-[10px] text-[#6f6556]">—</span>
            </li>
          </ul>
        </div>
      </section>

      <div v-else class="md-rise mt-5 border border-[#f2ead9]/12 px-6 py-12 text-center">
        <Moon class="mx-auto h-6 w-6 text-[#8d806c]" />
        <p class="md-serif mt-3 text-[13px] text-[#a99a85]">灶火还没点，稍坐片刻。</p>
      </div>

      <div v-if="payload?.notice" class="md-rise mt-5 border-l border-[#b4472b]/60 pl-3">
        <p class="text-[10px] leading-[1.7] text-[#958875]">{{ payload.notice }}</p>
      </div>

      <div class="md-rise md-d3 mt-7 flex items-center justify-between gap-3">
        <a v-if="payload?.phone" :href="'tel:' + payload.phone" class="md-call">
          <span class="md-call-dot"></span>
          <span class="md-serif text-[14px] tracking-[0.08em]">给店里打个电话</span>
          <span class="ml-auto font-mono text-[9px] tracking-[0.06em] text-[#cbb28a]">{{ payload.phone }}</span>
        </a>
      </div>

      <div v-if="payload?.address" class="mt-4 flex items-start gap-1.5 text-[9.5px] leading-relaxed text-[#7f7464]">
        <MapPin class="mt-0.5 h-3 w-3 flex-shrink-0" />
        <span>{{ payload.address }}</span>
      </div>

      <p class="mt-8 text-center font-mono text-[7px] uppercase tracking-[0.36em] text-[#5f564a]">Last order 01:30</p>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { Clock, MapPin, Phone, Flame, Moon } from 'lucide-vue-next'

const props = defineProps(['payload', 'nfcId'])

const pad = (n) => String(n).padStart(2, '0')
const hasPrice = (item) => item?.price !== undefined && item?.price !== null && item?.price !== ''
const signature = (item) => item?.signature === true || (Array.isArray(item?.tags) && item.tags.some((t) => /招牌|必点|signature/i.test(String(t))))
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
.md-glow {
  background: radial-gradient(60% 70% at 50% 0%, rgba(180, 71, 43, 0.24), transparent 70%),
    radial-gradient(40% 40% at 80% 10%, rgba(199, 162, 106, 0.10), transparent 70%);
}
.md-grain {
  opacity: 0.4;
  background-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' width='150' height='150'%3E%3Cfilter id='n'%3E%3CfeTurbulence type='fractalNoise' baseFrequency='0.9' numOctaves='3'/%3E%3C/filter%3E%3Crect width='150' height='150' filter='url(%23n)' opacity='0.08'/%3E%3C/svg%3E");
}
.md-serif { font-family: 'Songti SC', 'STSong', 'Yu Mincho', 'Hiragino Mincho ProN', serif; }
.md-lantern {
  display: inline-block;
  height: 6px;
  width: 6px;
  border-radius: 9999px;
  background: #d0603c;
  box-shadow: 0 0 10px 3px rgba(208, 96, 60, 0.45);
  animation: md-flicker 3.4s ease-in-out infinite;
}
@keyframes md-flicker {
  0%, 100% { opacity: 1; }
  42% { opacity: 0.6; }
  46% { opacity: 1; }
  70% { opacity: 0.75; }
}
.md-group { padding-top: 18px; }
.md-item {
  display: flex;
  align-items: baseline;
  gap: 9px;
  padding: 10px 0;
  border-bottom: 1px solid rgba(242, 234, 217, 0.06);
}
.md-item:last-child { border-bottom: none; }
.md-dots { flex: 1; min-width: 12px; border-bottom: 1px dotted rgba(242, 234, 217, 0.2); transform: translateY(-2px); }
.md-pick { border: 1px solid rgba(208, 96, 60, 0.6); padding: 0 4px; font-size: 8px; letter-spacing: 0.08em; color: #d98163; }
.md-call {
  display: flex;
  align-items: center;
  gap: 10px;
  width: 100%;
  border: 1px solid rgba(199, 162, 106, 0.35);
  border-radius: 2px;
  background: linear-gradient(100deg, rgba(180, 71, 43, 0.16), rgba(11, 16, 32, 0.2));
  padding: 13px 15px;
  color: #ecd9b8;
  transition: border-color 0.2s ease, background-color 0.2s ease, transform 0.15s ease;
}
.md-call:hover { border-color: rgba(199, 162, 106, 0.8); background-color: rgba(180, 71, 43, 0.2); }
.md-call:active { transform: scale(0.985); }
.md-call-dot { height: 7px; width: 7px; border-radius: 9999px; background: #c7a26a; box-shadow: 0 0 0 3px rgba(199, 162, 106, 0.15); }
.md-rise { animation: md-in 0.8s cubic-bezier(0.2, 0.9, 0.25, 1) both; }
.md-d1 { animation-delay: 0.1s; }
.md-d2 { animation-delay: 0.18s; }
.md-d3 { animation-delay: 0.34s; }
@keyframes md-in {
  from { opacity: 0; transform: translateY(16px); }
  to { opacity: 1; transform: none; }
}
@media (prefers-reduced-motion: reduce) {
  .md-rise, .md-call, .md-lantern { animation: none !important; transition: none !important; }
}
</style>
