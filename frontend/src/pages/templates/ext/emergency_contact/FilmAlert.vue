<template>
  <div class="fa-root min-h-screen text-[#f5ead7]">
    <main class="mx-auto w-full max-w-md px-4 pb-12 pt-6">
      <div class="fa-enter flex items-center justify-between px-2 text-[9px] uppercase tracking-[0.22em] text-[#f5b544]/70">
        <span class="flex items-center gap-2"><span class="fa-rec"></span>REC · Emergency Reel</span>
        <span v-if="nfcId" class="max-w-[42%] truncate">NFC {{ nfcId }}</span>
      </div>

      <div class="fa-strip relative mt-4 px-6 pb-7 pt-6">
        <header class="fa-enter fa-d1">
          <div class="flex items-center justify-between border-y border-white/15 py-2 text-[9px] uppercase tracking-[0.2em] text-[#c9b18e]">
            <span>ROLL 01 · 紧急联络</span><span>TAKE 02</span>
          </div>
          <h1 class="mt-5 text-[34px] font-black leading-[1.02] tracking-tight">找到我时<br /><span class="text-[#f5b544]">请先拨这通电话</span></h1>
          <p class="mt-3 max-w-[280px] text-[11px] leading-relaxed text-[#b8a58a]">保持通话，把镜头里的地点描述给对方。下面的号码按优先顺序排列。</p>
        </header>

        <section class="mt-6">
          <template v-if="contacts.length">
            <a
              :href="'tel:' + contacts[0].phone"
              class="fa-frame fa-primary fa-enter fa-d2 block p-4"
            >
              <div class="flex items-center justify-between text-[9px] uppercase tracking-[0.2em] text-[#f5b544]">
                <span>FRAME 01 / PRIMARY</span><span class="flex items-center gap-1"><span class="h-1.5 w-1.5 rounded-full bg-[#f5b544]"></span>LIVE</span>
              </div>
              <p class="mt-3 break-words text-[34px] font-black leading-none tabular-nums tracking-tight text-white">{{ contacts[0].phone }}</p>
              <div class="mt-3 flex items-end justify-between border-t border-white/10 pt-2.5">
                <p class="text-[12px] font-semibold text-[#eadbc1]">{{ contacts[0].name || '联系人' }}<span v-if="contacts[0].relation" class="font-normal text-[#a89578]"> · {{ contacts[0].relation }}</span></p>
                <span class="flex items-center gap-1.5 text-[9px] uppercase tracking-[0.16em] text-[#f5b544]"><Phone class="h-3.5 w-3.5" />CALL</span>
              </div>
            </a>

            <div class="mt-3 space-y-3">
              <a
                v-for="(c, i) in contacts.slice(1)"
                :key="i"
                :href="'tel:' + c.phone"
                class="fa-frame fa-enter fa-d3 flex items-stretch gap-3 p-3.5"
              >
                <div class="flex w-11 shrink-0 flex-col items-center justify-between border-r border-white/10 pr-3">
                  <span class="text-[8px] uppercase tracking-[0.1em] text-[#8f7d63]">FRAME</span>
                  <span class="text-[19px] font-black tabular-nums text-[#f5b544]">{{ String(i + 2).padStart(2, '0') }}</span>
                </div>
                <div class="min-w-0 flex-1">
                  <p class="truncate text-[21px] font-bold leading-tight tabular-nums text-white">{{ c.phone }}</p>
                  <p class="mt-1 truncate text-[10px] text-[#a89578]">{{ c.name || '联系人' }}<span v-if="c.relation"> · {{ c.relation }}</span></p>
                </div>
                <Phone class="h-4 w-4 shrink-0 self-center text-[#f5b544]" />
              </a>
            </div>
          </template>
        </section>

        <section class="fa-enter fa-d3 mt-5 border-t border-white/15 pt-4">
          <div class="flex items-center justify-between">
            <p class="text-[9px] uppercase tracking-[0.22em] text-[#f5b544]">DOP NOTES · 医疗记录</p>
            <HeartPulse class="h-4 w-4 text-[#f5b544]" />
          </div>
          <div class="mt-3 grid grid-cols-2 gap-x-4 gap-y-3">
            <div v-for="m in medical" :key="m.label" class="border-l border-[#f5b544]/30 pl-2.5">
              <p class="text-[9px] uppercase tracking-[0.12em] text-[#8f7d63]">{{ m.label }}</p>
              <p class="mt-1 text-[11px] font-medium leading-snug text-[#e6d6bb]">{{ m.value }}</p>
            </div>
          </div>
        </section>

        <div v-if="hasPet" class="fa-enter fa-d3 mt-4 border-y border-white/15 py-2.5 text-[11px] text-[#c9b18e]">
          <span class="text-[#f5b544]">PET /</span> {{ payload.petName }}<span v-if="payload.breed"> · {{ payload.breed }}</span>
          <a v-if="payload.vetPhone" :href="'tel:' + payload.vetPhone" class="ml-2 text-[#f5b544] underline underline-offset-4">{{ payload.vetPhone }}</a>
        </div>

        <div class="fa-enter fa-d3 mt-5 flex items-center justify-between text-[8px] uppercase tracking-[0.2em] text-[#7f6f59]">
          <span>Emergency Contact · 30fps</span><span>END OF REEL</span>
        </div>
      </div>
    </main>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { Phone, HeartPulse } from 'lucide-vue-next'

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
.fa-root {
  background:
    radial-gradient(80% 55% at 50% 0%, rgba(245, 181, 68, 0.08), transparent 60%),
    #0d0c0a;
}
.fa-strip {
  border: 1px solid rgba(255, 255, 255, 0.1);
  background:
    linear-gradient(90deg, transparent 0 14px, rgba(255, 255, 255, 0.025) 14px 100%),
    #14110c;
  box-shadow: 0 30px 70px -40px #000;
}
.fa-strip::before,
.fa-strip::after {
  content: "";
  position: absolute;
  top: 12px;
  bottom: 12px;
  width: 10px;
  background: repeating-linear-gradient(180deg, rgba(245, 181, 68, 0.55) 0 10px, transparent 10px 26px);
  border-radius: 3px;
  opacity: 0.55;
}
.fa-strip::before { left: 6px; }
.fa-strip::after { right: 6px; }
.fa-rec { width: 7px; height: 7px; border-radius: 999px; background: #ef4444; animation: fa-blink 1.5s ease-in-out infinite; }
.fa-frame {
  border: 1px solid rgba(255, 255, 255, 0.15);
  background: linear-gradient(155deg, rgba(255, 255, 255, 0.045), rgba(255, 255, 255, 0.015));
  border-radius: 4px;
  transition: border-color 0.2s ease, background-color 0.2s ease, transform 0.2s ease;
}
.fa-frame:hover { border-color: rgba(245, 181, 68, 0.55); background: rgba(245, 181, 68, 0.055); transform: translateX(2px); }
.fa-frame:active { transform: scale(0.99); }
.fa-primary { border-color: rgba(245, 181, 68, 0.55); background: linear-gradient(150deg, rgba(245, 181, 68, 0.12), rgba(245, 181, 68, 0.025)); }
.fa-primary:hover { border-color: #f5b544; background: rgba(245, 181, 68, 0.16); }
.fa-enter { animation: fa-slide 0.65s cubic-bezier(0.22, 1, 0.36, 1) both; }
.fa-d1 { animation-delay: 0.08s; }
.fa-d2 { animation-delay: 0.16s; }
.fa-d3 { animation-delay: 0.24s; }
@keyframes fa-slide {
  from { opacity: 0; transform: translateX(-12px); }
  to { opacity: 1; transform: translateX(0); }
}
@keyframes fa-blink { 50% { opacity: 0.25; } }
@media (prefers-reduced-motion: reduce) {
  .fa-enter, .fa-rec { animation: none; }
  .fa-frame { transition: none; }
}
</style>
