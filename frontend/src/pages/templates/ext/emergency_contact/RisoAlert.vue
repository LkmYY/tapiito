<template>
  <div class="ra-root min-h-screen text-[#1d2a3a]">
    <main class="mx-auto w-full max-w-md px-5 pb-12 pt-6">
      <div class="ra-enter flex items-center justify-between text-[9px] font-bold uppercase tracking-[0.22em]">
        <span>Riso Emergency Ticket</span>
        <span v-if="nfcId" class="max-w-[42%] truncate">NFC · {{ nfcId }}</span>
      </div>

      <header class="ra-enter ra-d1 relative mt-4">
        <div class="ra-shadow-blue absolute inset-0 translate-x-1.5 translate-y-1.5 rounded-[24px]"></div>
        <div class="ra-shadow-red absolute inset-0 -translate-x-1 rounded-[24px]"></div>
        <div class="ra-card relative rounded-[24px] border-2 border-[#1d2a3a] px-6 pb-6 pt-5">
          <div class="flex items-start justify-between gap-3">
            <div>
              <p class="text-[10px] font-bold uppercase tracking-[0.2em] text-[#e2452f]">Admit One · 紧急联络</p>
              <h1 class="ra-display mt-2 text-[40px] font-black leading-[0.95] tracking-tight">先打电话<br /><span class="text-[#e2452f]">再等待救援</span></h1>
            </div>
            <div class="ra-seal flex h-16 w-16 shrink-0 -rotate-6 flex-col items-center justify-center rounded-full border-2 border-[#1d2a3a] text-center text-[#e2452f]">
              <Siren class="h-5 w-5" />
              <span class="mt-0.5 text-[8px] font-bold tracking-[0.16em]">CALL</span>
            </div>
          </div>
          <p class="mt-4 max-w-[260px] text-[12px] leading-relaxed text-[#4c5667]">请选择最近的联系人拨号，说明持卡人姓名、所在位置与现场状况。</p>
          <div class="mt-4 flex items-center justify-between border-t border-dashed border-[#1d2a3a]/35 pt-3 text-[9px] font-bold uppercase tracking-[0.16em] text-[#6a7383]">
            <span>Issued for emergency use</span><span>No. 025</span>
          </div>
        </div>
      </header>

      <section class="mt-6">
        <div class="ra-enter ra-d2 flex items-end justify-between">
          <p class="text-[10px] font-black uppercase tracking-[0.2em] text-[#e2452f]">Call stubs · 拨打联</p>
          <span class="text-[10px] tabular-nums text-[#6a7383]">{{ String(contacts.length).padStart(2, '0') }}</span>
        </div>

        <template v-if="contacts.length">
          <article v-for="(c, i) in contacts" :key="i" class="ra-ticket ra-enter" :class="'ra-d' + Math.min(i + 1, 3)">
            <div class="flex items-stretch">
              <a :href="'tel:' + c.phone" class="ra-phone min-w-0 flex-1 px-5 py-4">
                <div class="flex items-center justify-between gap-2">
                  <p class="text-[11px] font-bold text-[#1d2a3a]">{{ c.name || '联系人' }}<span v-if="c.relation" class="ml-1 font-normal text-[#6a7383]">/ {{ c.relation }}</span></p>
                  <Phone class="h-3.5 w-3.5 shrink-0 text-[#e2452f]" />
                </div>
                <p class="mt-2 break-words text-[28px] font-black leading-tight tabular-nums tracking-tight" :class="i === 0 ? 'text-[#1d2a3a]' : 'text-[#1d2a3a]/85'">{{ c.phone }}</p>
                <p class="mt-1.5 text-[9px] font-bold uppercase tracking-[0.18em] text-[#e2452f]">Tap to call</p>
              </a>
              <div class="ra-stub flex w-[62px] shrink-0 flex-col items-center justify-center border-l-2 border-dashed border-[#1d2a3a]/35">
                <span class="ra-display text-[24px] font-black leading-none">{{ String(i + 1).padStart(2, '0') }}</span>
                <span class="mt-1 text-[8px] font-bold uppercase tracking-[0.14em] text-[#6a7383]">PRIORITY</span>
              </div>
            </div>
          </article>
        </template>
        <div v-else class="ra-ticket mt-2 border-2 border-dashed border-[#1d2a3a]/35 px-5 py-8 text-center text-[12px] text-[#6a7383]">暂无联系人 · 请补充号码</div>
      </section>

      <section class="ra-enter ra-d3 relative mt-6">
        <div class="rounded-2xl border-2 border-[#1d2a3a] bg-[#f8f0df] p-4 shadow-[5px_5px_0_rgba(38,75,115,0.85)]">
          <div class="flex items-center justify-between">
            <p class="text-[10px] font-black uppercase tracking-[0.2em]">Medical Notes · 医疗联</p>
            <HeartPulse class="h-4 w-4 text-[#e2452f]" />
          </div>
          <div class="mt-3 grid grid-cols-2 gap-x-4 gap-y-3">
            <div v-for="m in medical" :key="m.label" class="border-t border-[#1d2a3a]/20 pt-2">
              <p class="text-[9px] font-bold uppercase tracking-[0.12em] text-[#6a7383]">{{ m.label }}</p>
              <p class="mt-1 text-[12px] font-semibold leading-snug">{{ m.value }}</p>
            </div>
          </div>
        </div>
      </section>

      <div v-if="hasPet" class="ra-enter ra-d3 mt-4 border-2 border-[#1d2a3a] bg-[#e7f0d8] px-4 py-3 text-[12px]">
        <span class="font-bold">同行宠物：</span>{{ payload.petName }}<span v-if="payload.breed"> · {{ payload.breed }}</span>
        <a v-if="payload.vetPhone" :href="'tel:' + payload.vetPhone" class="ml-2 font-bold text-[#3b5d2a] underline underline-offset-2">{{ payload.vetPhone }}</a>
      </div>

      <p class="ra-enter ra-d3 mt-6 text-center text-[9px] font-bold uppercase tracking-[0.2em] text-[#6a7383]">Keep this ticket with the patient</p>
    </main>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { Siren, Phone, HeartPulse } from 'lucide-vue-next'

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
.ra-root {
  background:
    radial-gradient(rgba(29, 42, 58, 0.07) 0.7px, transparent 0.7px),
    #f3e8d3;
  background-size: 5px 5px;
}
.ra-display { font-family: "Arial Black", "Helvetica Neue", Arial, sans-serif; }
.ra-enter { animation: ra-rise 0.7s cubic-bezier(0.22, 1, 0.36, 1) both; }
.ra-d1 { animation-delay: 0.08s; }
.ra-d2 { animation-delay: 0.16s; }
.ra-d3 { animation-delay: 0.24s; }
@keyframes ra-rise {
  from { opacity: 0; transform: translateY(14px) rotate(-0.6deg); }
  to { opacity: 1; transform: translateY(0) rotate(0); }
}
.ra-shadow-red { background: #e2452f; opacity: 0.85; }
.ra-shadow-blue { background: #264b73; opacity: 0.8; }
.ra-card {
  background:
    radial-gradient(circle at 100% 0%, rgba(226, 69, 47, 0.1), transparent 40%),
    #f8f0df;
}
.ra-seal { background: rgba(248, 240, 223, 0.7); }
.ra-ticket {
  position: relative;
  margin-top: 10px;
  overflow: hidden;
  border: 2px solid #1d2a3a;
  border-radius: 14px;
  background: #f8f0df;
  box-shadow: 4px 4px 0 rgba(38, 75, 115, 0.85);
  animation: ra-rise 0.7s cubic-bezier(0.22, 1, 0.36, 1) both;
}
.ra-ticket::before,
.ra-ticket::after {
  content: "";
  position: absolute;
  left: -7px;
  width: 12px;
  height: 12px;
  border-radius: 999px;
  background: #f3e8d3;
  border: 2px solid #1d2a3a;
  z-index: 2;
}
.ra-ticket::before { top: -7px; }
.ra-ticket::after { bottom: -7px; }
.ra-phone { display: block; text-decoration: none; transition: background-color 0.2s ease; }
.ra-phone:hover { background: #fff6e4; }
.ra-phone:active { background: #f0dfc4; }
.ra-stub { background: rgba(226, 69, 47, 0.06); }
@media (prefers-reduced-motion: reduce) {
  .ra-enter, .ra-ticket { animation: none; }
  .ra-phone { transition: none; }
}
</style>
