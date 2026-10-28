<template>
  <div class="ss-root min-h-screen text-[#111111]">
    <main class="mx-auto w-full max-w-md px-5 pb-12 pt-6">
      <div class="ss-enter flex items-center justify-between border-b-2 border-[#111] pb-2 text-[9px] font-bold uppercase tracking-[0.2em]">
        <span class="flex items-center gap-2"><span class="h-2 w-2 bg-[#e30613]"></span>Emergency Contact</span>
        <span v-if="nfcId" class="max-w-[40%] truncate">NFC {{ nfcId }}</span>
      </div>

      <header class="ss-enter ss-d1 relative mt-5">
        <p class="text-[10px] font-bold uppercase tracking-[0.28em] text-[#e30613]">Notfall · 紧急情况</p>
        <h1 class="ss-title mt-2 text-[54px] font-black leading-[0.88] tracking-[-0.04em]">紧急<br />联络<span class="text-[#e30613]">.</span></h1>
        <div class="mt-5 grid grid-cols-[1fr_auto] items-end gap-4 border-t border-[#111] pt-3">
          <p class="max-w-[250px] text-[11px] leading-relaxed text-[#555]">按顺序拨号。请说明持卡人位置，并留在现场等待救援人员。</p>
          <Siren class="h-6 w-6 text-[#e30613]" />
        </div>
      </header>

      <section class="mt-7">
        <div class="ss-enter ss-d2 flex items-baseline justify-between border-b border-[#111] pb-1.5">
          <p class="text-[9px] font-bold uppercase tracking-[0.22em]">Call register / 联络顺序</p>
          <span class="text-[9px] tabular-nums">{{ String(contacts.length).padStart(2, '0') }}</span>
        </div>

        <template v-if="contacts.length">
          <a :href="'tel:' + contacts[0].phone" class="ss-primary ss-enter ss-d2 mt-3 block bg-[#111] px-5 pb-5 pt-4 text-white">
            <div class="flex items-center justify-between text-[9px] font-bold uppercase tracking-[0.2em] text-white/55">
              <span>01 / Primary</span><span class="h-2 w-2 bg-[#e30613]"></span>
            </div>
            <p class="mt-3 break-words text-[36px] font-black leading-none tabular-nums tracking-[-0.02em]">{{ contacts[0].phone }}</p>
            <div class="mt-4 flex items-center justify-between border-t border-white/20 pt-3">
              <p class="text-[11px] font-medium">{{ contacts[0].name || '联系人' }}<span v-if="contacts[0].relation" class="text-white/45"> / {{ contacts[0].relation }}</span></p>
              <span class="flex items-center gap-1.5 text-[9px] font-bold uppercase tracking-[0.16em] text-[#ff5a5f]"><Phone class="h-3.5 w-3.5" />Dial</span>
            </div>
          </a>

          <div class="mt-2 border-b border-[#111]">
            <a
              v-for="(c, i) in contacts.slice(1)"
              :key="i"
              :href="'tel:' + c.phone"
              class="ss-row grid grid-cols-[32px_1fr_auto] items-center gap-3 border-t border-[#ccc] py-3.5"
            >
              <span class="text-[9px] font-bold tabular-nums text-[#e30613]">{{ String(i + 2).padStart(2, '0') }}</span>
              <div class="min-w-0">
                <p class="truncate text-[21px] font-bold leading-tight tabular-nums">{{ c.phone }}</p>
                <p class="mt-0.5 truncate text-[10px] text-[#666]">{{ c.name || '联系人' }}<span v-if="c.relation"> · {{ c.relation }}</span></p>
              </div>
              <ArrowUpRight class="h-4 w-4 text-[#111]" />
            </a>
          </div>
        </template>
      </section>

      <section class="ss-enter ss-d3 mt-7">
        <div class="flex items-center justify-between border-b-2 border-[#111] pb-1.5">
          <p class="text-[9px] font-bold uppercase tracking-[0.22em]">Medical / 医疗信息</p>
          <HeartPulse class="h-4 w-4 text-[#e30613]" />
        </div>
        <div class="grid grid-cols-2 gap-x-4 gap-y-4 pt-3">
          <div v-for="m in medical" :key="m.label" class="border-t border-[#d5d5d0] pt-2">
            <p class="text-[9px] font-bold uppercase tracking-[0.1em] text-[#888]">{{ m.label }}</p>
            <p class="mt-1 text-[12px] font-semibold leading-snug">{{ m.value }}</p>
          </div>
        </div>
      </section>

      <div v-if="hasPet" class="ss-enter ss-d3 mt-5 flex items-center justify-between border-y border-[#111] py-2.5 text-[11px]">
        <span><strong>{{ payload.petName }}</strong><span v-if="payload.breed" class="text-[#666]"> / {{ payload.breed }}</span></span>
        <a v-if="payload.vetPhone" :href="'tel:' + payload.vetPhone" class="font-bold text-[#e30613]">{{ payload.vetPhone }}</a>
      </div>

      <footer class="ss-enter ss-d3 mt-6 flex items-end justify-between">
        <p class="text-[9px] font-bold uppercase leading-relaxed tracking-[0.2em] text-[#999]">Keep calm.<br />Make the call.</p>
        <p class="text-[42px] font-black leading-none tracking-[-0.06em] text-[#e5e5e0]">SOS</p>
      </footer>
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
  { label: '过敏史', value: props.payload?.allergies },
  { label: '正在用药', value: props.payload?.medications },
  { label: '既往病史', value: props.payload?.conditions },
])
const hasPet = computed(() => !!(props.payload?.petName || props.payload?.chipId))
</script>

<style scoped>
.ss-root { background: #f7f7f3; }
.ss-title { font-family: "Helvetica Neue", "Arial Black", Arial, "PingFang SC", sans-serif; }
.ss-enter { animation: ss-up 0.65s cubic-bezier(0.22, 1, 0.36, 1) both; }
.ss-d1 { animation-delay: 0.08s; }
.ss-d2 { animation-delay: 0.16s; }
.ss-d3 { animation-delay: 0.24s; }
@keyframes ss-up {
  from { opacity: 0; transform: translateY(14px); clip-path: inset(0 0 100% 0); }
  to { opacity: 1; transform: translateY(0); clip-path: inset(0 0 0 0); }
}
.ss-primary { transition: background-color 0.2s ease, transform 0.2s ease; }
.ss-primary:hover { background: #242424; transform: translateY(-2px); }
.ss-primary:active { transform: scale(0.99); }
.ss-row { transition: background-color 0.18s ease, padding-left 0.18s ease; }
.ss-row:hover { background: #fff; padding-left: 5px; }
.ss-row:active { background: #efefea; }
@media (prefers-reduced-motion: reduce) {
  .ss-enter { animation: none; }
  .ss-primary, .ss-row { transition: none; }
}
</style>
