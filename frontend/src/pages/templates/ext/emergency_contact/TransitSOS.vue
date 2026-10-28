<template>
  <div class="ts-root min-h-screen text-[#f5f6f2]">
    <main class="mx-auto w-full max-w-md px-5 pb-12 pt-6">
      <div class="ts-enter flex items-center justify-between text-[9px] font-bold uppercase tracking-[0.2em] text-[#f5d400]">
        <span class="flex items-center gap-2"><TrainFront class="h-3.5 w-3.5" />Emergency Line</span>
        <span v-if="nfcId" class="max-w-[42%] truncate text-white/45">NFC · {{ nfcId }}</span>
      </div>

      <header class="ts-enter ts-d1 mt-4 border-y-4 border-[#f5d400] py-4">
        <div class="flex items-start justify-between gap-4">
          <div>
            <p class="text-[10px] font-bold uppercase tracking-[0.24em] text-[#f5d400]">E-01 · TO SAFETY</p>
            <h1 class="mt-1 text-[36px] font-black leading-[0.95] tracking-tight">紧急联络线</h1>
            <p class="mt-2 text-[11px] leading-relaxed text-white/55">沿路线依次拨号，不要跳过第一条。报清站点位置。</p>
          </div>
          <span class="flex h-14 w-14 shrink-0 items-center justify-center rounded-full bg-[#f5d400] text-[13px] font-black text-[#0b0f12]">SOS</span>
        </div>
      </header>

      <section class="ts-enter ts-d2 mt-5">
        <div class="flex items-center justify-between text-[9px] font-bold uppercase tracking-[0.2em] text-white/40">
          <span>Route · 联络顺序</span><span>{{ String(contacts.length).padStart(2, '0') }} STOPS</span>
        </div>

        <div class="relative mt-3 pl-9">
          <div class="ts-line absolute bottom-3 left-[13px] top-3 w-1 bg-[#f5d400]/35"></div>
          <template v-if="contacts.length">
            <a :href="'tel:' + contacts[0].phone" class="ts-primary relative mb-3 block bg-[#f5d400] px-5 py-4 text-[#0b0f12]">
              <span class="ts-stop absolute -left-9 top-5 flex h-7 w-7 items-center justify-center rounded-full border-4 border-[#0b0f12] bg-[#f5d400] text-[9px] font-black">01</span>
              <div class="flex items-center justify-between text-[9px] font-black uppercase tracking-[0.16em]">
                <span>Current stop · 立即</span><Phone class="h-4 w-4" />
              </div>
              <p class="mt-2 break-words text-[34px] font-black leading-none tabular-nums tracking-tight">{{ contacts[0].phone }}</p>
              <div class="mt-3 flex items-center justify-between border-t border-[#0b0f12]/20 pt-2.5">
                <p class="text-[12px] font-bold">{{ contacts[0].name || '联系人' }}<span v-if="contacts[0].relation"> · {{ contacts[0].relation }}</span></p>
                <span class="text-[9px] font-black uppercase tracking-[0.12em]">Tap to call</span>
              </div>
            </a>

            <a
              v-for="(c, i) in contacts.slice(1)"
              :key="i"
              :href="'tel:' + c.phone"
              class="ts-row relative mb-2.5 flex items-center justify-between gap-3 border border-white/12 bg-white/[0.04] px-4 py-3.5 last:mb-0"
            >
              <span class="ts-stop absolute -left-9 top-1/2 flex h-6 w-6 -translate-y-1/2 items-center justify-center rounded-full border-2 border-[#f5d400]/60 bg-[#0b0f12] text-[8px] font-black text-[#f5d400]">{{ String(i + 2).padStart(2, '0') }}</span>
              <div class="min-w-0">
                <p class="truncate text-[21px] font-bold leading-tight tabular-nums text-white">{{ c.phone }}</p>
                <p class="mt-0.5 truncate text-[10px] text-white/45">{{ c.name || '联系人' }}<span v-if="c.relation"> · {{ c.relation }}</span></p>
              </div>
              <span class="flex h-8 w-8 shrink-0 items-center justify-center rounded-full border border-[#f5d400]/40 text-[#f5d400]"><Phone class="h-3.5 w-3.5" /></span>
            </a>
          </template>
          <div v-else class="relative border border-dashed border-white/25 px-4 py-8 text-center text-[12px] text-white/45">
            <span class="ts-stop absolute -left-9 top-1/2 flex h-7 w-7 -translate-y-1/2 items-center justify-center rounded-full border-2 border-white/30 bg-[#0b0f12] text-[10px] text-white/40">—</span>
            线路未配置 · 暂无联络站
          </div>
        </div>
      </section>

      <section class="ts-enter ts-d3 mt-5 border border-white/15 bg-[#11181d]">
        <div class="flex items-center justify-between bg-[#1a2329] px-4 py-2.5">
          <p class="text-[9px] font-bold uppercase tracking-[0.2em] text-[#f5d400]">Service board · 医疗</p>
          <Info class="h-4 w-4 text-[#f5d400]" />
        </div>
        <div class="grid grid-cols-2 gap-px bg-white/10">
          <div v-for="m in medical" :key="m.label" class="bg-[#11181d] px-4 py-3">
            <p class="text-[9px] uppercase tracking-[0.12em] text-white/35">{{ m.label }}</p>
            <p class="mt-1 text-[11px] font-semibold leading-snug text-white/85">{{ m.value }}</p>
          </div>
        </div>
      </section>

      <div v-if="hasPet" class="ts-enter ts-d3 mt-3 flex items-center justify-between border-l-4 border-[#f5d400] bg-white/[0.04] px-4 py-2.5 text-[11px]">
        <span class="text-white/70">{{ payload.petName }}<span v-if="payload.breed"> · {{ payload.breed }}</span></span>
        <a v-if="payload.vetPhone" :href="'tel:' + payload.vetPhone" class="font-bold text-[#f5d400]">宠物医院 {{ payload.vetPhone }}</a>
      </div>

      <p class="ts-enter ts-d3 mt-5 flex items-center justify-between text-[8px] font-bold uppercase tracking-[0.2em] text-white/30">
        <span>Approach carefully · 注意安全</span><span>E-01</span>
      </p>
    </main>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { TrainFront, Phone, Info } from 'lucide-vue-next'

const props = defineProps(['payload', 'nfcId'])
const contacts = computed(() =>
  (Array.isArray(props.payload?.contacts) ? props.payload.contacts : []).filter((c) => c && c.phone)
)
const medical = computed(() => [
  { label: '血型', value: props.payload?.bloodType || '—' },
  { label: '过敏史', value: props.payload?.allergies },
  { label: '正在用药', value: props.payload?.medications },
  { label: '既往病史', value: props.payload?.conditions },
])
const hasPet = computed(() => !!(props.payload?.petName || props.payload?.chipId))
</script>

<style scoped>
.ts-root { background: #0b0f12; }
.ts-primary {
  box-shadow: 0 18px 40px -28px rgba(245, 212, 0, 0.55);
  transition: transform 0.2s ease, background-color 0.2s ease;
}
.ts-primary:hover { background: #ffe83d; transform: translateX(3px); }
.ts-primary:active { transform: scale(0.99); }
.ts-stop { box-shadow: 0 0 0 3px #0b0f12; }
.ts-row { transition: background-color 0.18s ease, border-color 0.18s ease, transform 0.18s ease; }
.ts-row:hover { background: rgba(245, 212, 0, 0.08); border-color: rgba(245, 212, 0, 0.5); transform: translateX(3px); }
.ts-row:active { transform: translateX(0); }
.ts-line { transform-origin: top; animation: ts-draw 1.1s cubic-bezier(0.22, 1, 0.36, 1) both; }
.ts-enter { animation: ts-in 0.7s cubic-bezier(0.22, 1, 0.36, 1) both; }
.ts-d1 { animation-delay: 0.08s; }
.ts-d2 { animation-delay: 0.16s; }
.ts-d3 { animation-delay: 0.24s; }
@keyframes ts-in {
  from { opacity: 0; transform: translateY(14px); }
  to { opacity: 1; transform: translateY(0); }
}
@keyframes ts-draw { from { transform: scaleY(0); } to { transform: scaleY(1); } }
@media (prefers-reduced-motion: reduce) {
  .ts-enter, .ts-line { animation: none; }
  .ts-primary, .ts-row { transition: none; }
}
</style>
