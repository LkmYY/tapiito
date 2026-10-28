<template>
  <div class="da-root min-h-screen text-[#f2e7cd]">
    <main class="mx-auto w-full max-w-md px-5 pb-12 pt-6">
      <div class="da-enter flex items-center justify-between border-b border-[#c9a24b]/35 pb-2 text-[9px] uppercase tracking-[0.24em] text-[#c9a24b]">
        <span>Maison Assistance</span>
        <span v-if="nfcId" class="max-w-[42%] truncate">NFC · {{ nfcId }}</span>
      </div>

      <header class="da-enter da-d1 relative mt-5 overflow-hidden border border-[#c9a24b]/35 px-5 pb-6 pt-7 text-center">
        <div class="da-fan" aria-hidden="true"></div>
        <p class="relative text-[9px] uppercase tracking-[0.36em] text-[#c9a24b]">In Case of Emergency</p>
        <h1 class="da-serif relative mt-3 text-[37px] font-medium leading-[1.05] tracking-[0.04em]">紧急联络</h1>
        <div class="relative mx-auto mt-3 flex w-28 items-center gap-2 text-[#c9a24b]">
          <span class="h-px flex-1 bg-current"></span><span class="h-1.5 w-1.5 rotate-45 border border-current"></span><span class="h-px flex-1 bg-current"></span>
        </div>
        <p class="relative mx-auto mt-4 max-w-[270px] text-[11px] leading-relaxed text-[#d7c8a8]/65">请告知接线员持卡人的位置与状况。若第一位未接听，请继续拨打下一位。</p>
      </header>

      <section class="mt-5">
        <div class="da-enter da-d2 flex items-center justify-between">
          <p class="text-[9px] uppercase tracking-[0.26em] text-[#c9a24b]">Telephone Register</p>
          <Phone class="h-4 w-4 text-[#c9a24b]" />
        </div>

        <template v-if="contacts.length">
          <a :href="'tel:' + contacts[0].phone" class="da-primary da-enter da-d2 mt-2 block px-5 py-5">
            <div class="flex items-center justify-between text-[9px] uppercase tracking-[0.2em] text-[#c9a24b]">
              <span>No. 01 · Principal</span><span class="h-px w-10 bg-[#c9a24b]/45"></span>
            </div>
            <p class="da-serif mt-3 break-words text-[36px] font-medium leading-none tabular-nums tracking-[0.02em] text-[#fbf1da]">{{ contacts[0].phone }}</p>
            <div class="mt-4 flex items-center justify-between border-t border-[#c9a24b]/30 pt-3">
              <p class="text-[12px]">{{ contacts[0].name || '联系人' }}<span v-if="contacts[0].relation" class="text-[#c9a24b]"> · {{ contacts[0].relation }}</span></p>
              <span class="text-[9px] uppercase tracking-[0.18em] text-[#c9a24b]">Tap to call</span>
            </div>
          </a>

          <div class="mt-2 divide-y divide-[#c9a24b]/20 border border-[#c9a24b]/25">
            <a
              v-for="(c, i) in contacts.slice(1)"
              :key="i"
              :href="'tel:' + c.phone"
              class="da-row flex items-center gap-4 px-4 py-3.5"
            >
              <span class="da-serif w-6 shrink-0 text-[13px] text-[#c9a24b]">{{ ['I', 'II', 'III', 'IV'][i] || '·' }}</span>
              <div class="min-w-0 flex-1">
                <p class="truncate text-[20px] font-medium leading-tight tabular-nums text-[#f2e7cd]">{{ c.phone }}</p>
                <p class="mt-0.5 truncate text-[10px] uppercase tracking-[0.1em] text-[#c9a24b]/75">{{ c.name || '联系人' }}<span v-if="c.relation"> · {{ c.relation }}</span></p>
              </div>
              <ArrowUpRight class="h-4 w-4 shrink-0 text-[#c9a24b]" />
            </a>
          </div>
        </template>
        <div v-else class="mt-2 border border-dashed border-[#c9a24b]/40 px-4 py-7 text-center text-[12px] text-[#d7c8a8]/60">暂未登记联络电话</div>
      </section>

      <section class="da-enter da-d3 mt-5 border border-[#c9a24b]/30 px-4 py-4">
        <div class="flex items-center justify-between">
          <p class="text-[9px] uppercase tracking-[0.24em] text-[#c9a24b]">Physician's Note · 医疗</p>
          <HeartPulse class="h-4 w-4 text-[#c9a24b]" />
        </div>
        <div class="mt-3 grid grid-cols-2 gap-x-5 gap-y-3">
          <div v-for="m in medical" :key="m.label" class="border-t border-[#c9a24b]/20 pt-2">
            <p class="text-[9px] uppercase tracking-[0.1em] text-[#c9a24b]/70">{{ m.label }}</p>
            <p class="mt-1 text-[11px] leading-snug text-[#e7d9bc]">{{ m.value }}</p>
          </div>
        </div>
      </section>

      <div v-if="hasPet" class="da-enter da-d3 mt-3 flex items-center justify-between border border-[#c9a24b]/30 px-4 py-2.5 text-[11px] text-[#e7d9bc]">
        <span>{{ payload.petName }}<span v-if="payload.breed" class="text-[#c9a24b]/70"> · {{ payload.breed }}</span></span>
        <a v-if="payload.vetPhone" :href="'tel:' + payload.vetPhone" class="text-[#c9a24b] underline decoration-[#c9a24b]/40 underline-offset-4">VET</a>
      </div>

      <p class="da-enter da-d3 mt-5 text-center text-[8px] uppercase tracking-[0.32em] text-[#c9a24b]/60">Courtesy · Calm · Assistance</p>
    </main>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { Phone, ArrowUpRight, HeartPulse } from 'lucide-vue-next'

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
.da-root {
  background:
    radial-gradient(80% 40% at 50% 0%, rgba(201, 162, 75, 0.12), transparent 65%),
    #05261f;
}
.da-serif { font-family: "Iowan Old Style", "Palatino Linotype", Georgia, "Songti SC", serif; }
.da-fan {
  position: absolute;
  left: 50%;
  top: -190px;
  width: 400px;
  height: 400px;
  transform: translateX(-50%);
  border-radius: 999px;
  background: repeating-conic-gradient(from 250deg, rgba(201, 162, 75, 0.15) 0 1deg, transparent 1deg 9deg);
  -webkit-mask-image: radial-gradient(circle, transparent 0 30%, #000 52% 72%, transparent 88%);
  mask-image: radial-gradient(circle, transparent 0 30%, #000 52% 72%, transparent 88%);
  opacity: 0.7;
}
.da-primary {
  border: 1px solid rgba(201, 162, 75, 0.45);
  background: linear-gradient(160deg, rgba(201, 162, 75, 0.09), rgba(0, 0, 0, 0.12));
  transition: background 0.22s ease, border-color 0.22s ease, transform 0.22s ease;
}
.da-primary:hover { border-color: #d8b768; background: rgba(201, 162, 75, 0.14); transform: translateY(-2px); }
.da-primary:active { transform: scale(0.99); }
.da-row { transition: background-color 0.18s ease, padding-left 0.18s ease; }
.da-row:hover { background: rgba(201, 162, 75, 0.08); padding-left: 1.1rem; }
.da-enter { animation: da-in 0.8s cubic-bezier(0.22, 1, 0.36, 1) both; }
.da-d1 { animation-delay: 0.08s; }
.da-d2 { animation-delay: 0.16s; }
.da-d3 { animation-delay: 0.24s; }
@keyframes da-in {
  from { opacity: 0; transform: translateY(16px); }
  to { opacity: 1; transform: translateY(0); }
}
@media (prefers-reduced-motion: reduce) {
  .da-enter { animation: none; }
  .da-primary, .da-row { transition: none; }
}
</style>
