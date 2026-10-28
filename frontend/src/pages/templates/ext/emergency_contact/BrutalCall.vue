<template>
  <div class="bc-root min-h-screen text-[#0a0a0a]">
    <div class="bc-hazard h-3 w-full border-b-2 border-black"></div>
    <main class="mx-auto w-full max-w-md px-4 pb-12 pt-5">
      <div class="bc-enter flex items-center justify-between border-2 border-black bg-white px-3 py-1.5 text-[9px] font-black uppercase tracking-[0.14em] shadow-[3px_3px_0_#0a0a0a]">
        <span class="flex items-center gap-1.5"><Siren class="h-3.5 w-3.5" />Emergency Contact</span>
        <span v-if="nfcId" class="max-w-[38%] truncate">{{ nfcId }}</span>
      </div>

      <header class="bc-enter bc-d1 mt-5">
        <div class="flex items-start justify-between gap-2">
          <h1 class="bc-display text-[50px] font-black leading-[0.86] tracking-[-0.05em]">先打<br />电话<span class="text-white [text-shadow:3px_3px_0_#0a0a0a]">.</span></h1>
          <div class="bc-sticker mt-1 rotate-6 border-2 border-black bg-white px-2.5 py-2 text-center shadow-[3px_3px_0_#0a0a0a]">
            <p class="text-[9px] font-black leading-tight">DO<br />IT<br />NOW</p>
          </div>
        </div>
        <p class="mt-4 max-w-[290px] border-l-4 border-black pl-3 text-[12px] font-bold leading-relaxed">不要围观。按下面的顺序拨号，清楚说出地点和伤者状况。</p>
      </header>

      <section class="mt-6">
        <template v-if="contacts.length">
          <a :href="'tel:' + contacts[0].phone" class="bc-primary bc-enter bc-d2 block border-4 border-black bg-white px-5 pb-5 pt-4">
            <div class="flex items-center justify-between">
              <span class="bc-tag bg-[#0a0a0a] px-2 py-1 text-[9px] font-black uppercase tracking-[0.14em] text-white">01 / 首要</span>
              <span class="flex h-9 w-9 items-center justify-center rounded-full border-2 border-black bg-[#f2d800]"><Phone class="h-4 w-4" /></span>
            </div>
            <p class="mt-3 break-words text-[35px] font-black leading-none tabular-nums tracking-[-0.03em]">{{ contacts[0].phone }}</p>
            <div class="mt-4 flex items-center justify-between border-t-2 border-black pt-3">
              <p class="text-[12px] font-black">{{ contacts[0].name || '联系人' }}<span v-if="contacts[0].relation" class="font-bold"> / {{ contacts[0].relation }}</span></p>
              <span class="bc-btn inline-flex items-center gap-1.5 border-2 border-black bg-[#f2d800] px-3 py-1.5 text-[10px] font-black">CALL <ArrowUpRight class="h-3.5 w-3.5" /></span>
            </div>
          </a>

          <div class="mt-4 space-y-3">
            <a
              v-for="(c, i) in contacts.slice(1)"
              :key="i"
              :href="'tel:' + c.phone"
              class="bc-row flex items-center justify-between gap-3 border-[3px] border-black px-4 py-3.5"
              :class="i % 2 === 0 ? 'bg-[#fffdf2]' : 'bg-[#ffe94d]'"
            >
              <div class="min-w-0">
                <p class="truncate text-[22px] font-black leading-tight tabular-nums">{{ c.phone }}</p>
                <p class="mt-0.5 truncate text-[10px] font-bold uppercase tracking-[0.08em]">{{ c.name || '联系人' }}<span v-if="c.relation"> / {{ c.relation }}</span></p>
              </div>
              <span class="flex h-8 w-8 shrink-0 items-center justify-center border-2 border-black bg-white"><Phone class="h-3.5 w-3.5" /></span>
            </a>
          </div>
        </template>
      </section>

      <section class="bc-med bc-enter bc-d3 mt-5 border-4 border-black bg-[#0a0a0a] px-4 pb-4 pt-3 text-white">
        <div class="flex items-center justify-between border-b-2 border-[#f2d800] pb-2">
          <p class="text-[10px] font-black uppercase tracking-[0.16em] text-[#f2d800]">Medical / 医疗</p>
          <HeartPulse class="h-4 w-4 text-[#f2d800]" />
        </div>
        <div class="mt-3 grid grid-cols-2 gap-x-4 gap-y-3">
          <div v-for="m in medical" :key="m.label">
            <p class="text-[9px] font-bold uppercase tracking-[0.1em] text-white/45">{{ m.label }}</p>
            <p class="mt-1 text-[11px] font-bold leading-snug">{{ m.value }}</p>
          </div>
        </div>
      </section>

      <div v-if="hasPet" class="bc-enter bc-d3 mt-3 flex items-center justify-between border-[3px] border-black bg-white px-4 py-2.5 text-[11px] font-bold">
        <span>PET / {{ payload.petName }}<span v-if="payload.breed"> · {{ payload.breed }}</span></span>
        <a v-if="payload.vetPhone" :href="'tel:' + payload.vetPhone" class="underline decoration-2 underline-offset-2">{{ payload.vetPhone }}</a>
      </div>

      <p class="bc-enter bc-d3 mt-5 text-center text-[10px] font-black uppercase tracking-[0.2em]">Stay until help arrives · 别离开</p>
    </main>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { Siren, Phone, ArrowUpRight, HeartPulse } from 'lucide-vue-next'

const props = defineProps(['payload', 'nfcId'])
const contacts = computed(() =>
  (Array.isArray(props.payload?.contacts) ? props.payload.contacts : []).filter((c) => c && c.phone)
)
const medical = computed(() => [
  { label: '血型', value: props.payload?.bloodType || '—' },
  { label: '过敏', value: props.payload?.allergies },
  { label: '用药', value: props.payload?.medications },
  { label: '病史', value: props.payload?.conditions },
])
const hasPet = computed(() => !!(props.payload?.petName || props.payload?.chipId))
</script>

<style scoped>
.bc-root { background: #f2d800; }
.bc-hazard { background: repeating-linear-gradient(-45deg, #0a0a0a 0 12px, #f2d800 12px 24px); }
.bc-display { font-family: "Arial Black", "Helvetica Neue", Arial, "PingFang SC", sans-serif; }
.bc-sticker { font-family: "Arial Black", Arial, sans-serif; }
.bc-primary {
  box-shadow: 7px 7px 0 #0a0a0a;
  transition: transform 0.2s cubic-bezier(0.34, 1.56, 0.64, 1), box-shadow 0.2s ease;
}
.bc-primary:hover { transform: translate(-3px, -3px); box-shadow: 10px 10px 0 #0a0a0a; }
.bc-primary:active { transform: translate(2px, 2px); box-shadow: 3px 3px 0 #0a0a0a; }
.bc-btn { box-shadow: 2px 2px 0 #0a0a0a; transition: transform 0.15s ease; }
.bc-primary:hover .bc-btn { transform: translate(-1px, -1px); }
.bc-row {
  box-shadow: 4px 4px 0 #0a0a0a;
  transition: transform 0.18s ease, box-shadow 0.18s ease;
}
.bc-row:hover { transform: translate(-2px, -2px); box-shadow: 6px 6px 0 #0a0a0a; }
.bc-row:active { transform: translate(1px, 1px); box-shadow: 2px 2px 0 #0a0a0a; }
.bc-enter { animation: bc-slam 0.55s cubic-bezier(0.34, 1.56, 0.64, 1) both; }
.bc-d1 { animation-delay: 0.08s; }
.bc-d2 { animation-delay: 0.16s; }
.bc-d3 { animation-delay: 0.24s; }
@keyframes bc-slam {
  from { opacity: 0; transform: translateY(22px) scale(0.97); }
  to { opacity: 1; transform: translateY(0) scale(1); }
}
@media (prefers-reduced-motion: reduce) {
  .bc-enter { animation: none; }
  .bc-primary, .bc-btn, .bc-row { transition: none; }
}
</style>
