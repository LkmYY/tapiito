<template>
  <div class="gg-root min-h-screen text-white">
    <main class="mx-auto w-full max-w-md px-5 pb-12 pt-7">
      <div class="gg-enter flex items-center justify-between text-[10px] uppercase tracking-[0.24em] text-white/55">
        <span class="flex items-center gap-2"><ShieldCheck class="h-3.5 w-3.5 text-[#93c5fd]" />Guardian Link</span>
        <span v-if="nfcId" class="max-w-[42%] truncate">NFC · {{ nfcId }}</span>
      </div>

      <header class="gg-glass gg-enter gg-d1 relative mt-5 overflow-hidden rounded-[30px] px-6 pb-7 pt-6">
        <div class="gg-orb" aria-hidden="true"></div>
        <p class="text-[10px] font-semibold uppercase tracking-[0.26em] text-[#93c5fd]">Emergency Contact</p>
        <h1 class="mt-3 text-[32px] font-black leading-[1.05] tracking-tight">紧急联系人<br /><span class="text-white/50">请代为拨号</span></h1>
        <p class="mt-4 max-w-[250px] text-[12px] leading-relaxed text-white/60">持卡人可能无法自行说明情况。拨通后请告知地点，并留在现场等待救援。</p>
        <div class="mt-6 flex items-center gap-2 text-[10px] uppercase tracking-[0.2em] text-white/40">
          <span class="h-1.5 w-1.5 animate-pulse rounded-full bg-[#67e8f9]"></span> Live assistance
        </div>
      </header>

      <section class="mt-4">
        <a
          v-if="contacts.length"
          :href="'tel:' + contacts[0].phone"
          class="gg-call gg-enter gg-d2 block rounded-[26px] px-6 py-6"
        >
          <div class="flex items-center justify-between">
            <span class="rounded-full border border-white/20 bg-white/10 px-3 py-1 text-[10px] uppercase tracking-[0.18em] text-[#bae6fd]">Primary</span>
            <span class="flex h-11 w-11 items-center justify-center rounded-full bg-[#67e8f9] text-[#081120]"><Phone class="h-5 w-5" /></span>
          </div>
          <p class="mt-5 break-words text-[38px] font-black leading-none tabular-nums tracking-tight text-white">{{ contacts[0].phone }}</p>
          <p class="mt-4 text-[13px] font-semibold text-white/85">{{ contacts[0].name || '联系人' }}<span v-if="contacts[0].relation" class="font-normal text-white/45"> · {{ contacts[0].relation }}</span></p>
        </a>
        <div v-else class="gg-glass gg-enter gg-d2 rounded-[26px] px-6 py-8 text-center text-sm text-white/55">暂无紧急联系人号码</div>

        <div v-if="contacts.length > 1" class="mt-3 grid gap-2.5">
          <a
            v-for="(c, i) in contacts.slice(1)"
            :key="i"
            :href="'tel:' + c.phone"
            class="gg-row gg-enter gg-d3 flex items-center justify-between gap-4 rounded-2xl px-5 py-4"
          >
            <div class="min-w-0">
              <p class="truncate text-[21px] font-bold leading-tight tabular-nums">{{ c.phone }}</p>
              <p class="mt-1 truncate text-[11px] text-white/45">{{ c.name || '联系人' }}<span v-if="c.relation"> · {{ c.relation }}</span></p>
            </div>
            <Phone class="h-4 w-4 flex-shrink-0 text-[#a5f3fc]" />
          </a>
        </div>
      </section>

      <section class="gg-glass gg-enter gg-d3 mt-4 rounded-2xl p-4">
        <div class="flex items-center justify-between">
          <p class="text-[10px] font-semibold uppercase tracking-[0.22em] text-white/50">Medical Snapshot</p>
          <HeartPulse class="h-4 w-4 text-[#f9a8d4]" />
        </div>
        <div class="mt-3 grid grid-cols-2 gap-2">
          <div v-for="m in medical" :key="m.label" class="rounded-xl border border-white/10 bg-white/[0.045] px-3 py-2.5">
            <p class="text-[9px] uppercase tracking-[0.16em] text-white/35">{{ m.label }}</p>
            <p class="mt-1 text-[12px] font-semibold leading-snug text-white/90">{{ m.value }}</p>
          </div>
        </div>
      </section>

      <div v-if="hasPet" class="gg-glass gg-enter gg-d3 mt-3 flex items-center justify-between rounded-2xl px-4 py-3 text-[12px]">
        <span class="text-white/75">{{ payload.petName }}<span v-if="payload.breed" class="text-white/40"> · {{ payload.breed }}</span></span>
        <a v-if="payload.vetPhone" :href="'tel:' + payload.vetPhone" class="font-semibold text-[#a5f3fc]">宠物医院</a>
      </div>

      <div class="gg-enter gg-d3 mt-7 flex items-center justify-between border-t border-white/10 pt-4 text-[9px] uppercase tracking-[0.2em] text-white/30">
        <span>Stay on the line</span>
        <span v-if="nfcId" class="max-w-[40%] truncate">NFC · {{ nfcId }}</span>
      </div>
    </main>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { ShieldCheck, Phone, HeartPulse } from 'lucide-vue-next'

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
.gg-root {
  position: relative;
  isolation: isolate;
  background:
    radial-gradient(80% 50% at 8% 5%, rgba(37, 99, 235, 0.28), transparent 62%),
    radial-gradient(70% 45% at 100% 36%, rgba(124, 58, 237, 0.3), transparent 60%),
    radial-gradient(55% 35% at 60% 100%, rgba(6, 182, 212, 0.18), transparent 65%),
    #081120;
}
.gg-glass {
  border: 1px solid rgba(255, 255, 255, 0.13);
  background: linear-gradient(150deg, rgba(255, 255, 255, 0.11), rgba(255, 255, 255, 0.035));
  box-shadow: 0 24px 60px -35px rgba(0, 0, 0, 0.9), inset 0 1px 0 rgba(255, 255, 255, 0.12);
  backdrop-filter: blur(20px);
  -webkit-backdrop-filter: blur(20px);
}
.gg-orb {
  position: absolute;
  right: -60px;
  top: -70px;
  width: 190px;
  height: 190px;
  border-radius: 999px;
  background: radial-gradient(circle, rgba(147, 197, 253, 0.55), rgba(124, 58, 237, 0.12) 55%, transparent 70%);
  filter: blur(6px);
  animation: gg-drift 8s ease-in-out infinite alternate;
}
.gg-call {
  background: linear-gradient(135deg, rgba(37, 99, 235, 0.48), rgba(124, 58, 237, 0.34));
  border: 1px solid rgba(191, 219, 254, 0.36);
  box-shadow: 0 26px 70px -30px rgba(59, 130, 246, 0.65), inset 0 1px 0 rgba(255, 255, 255, 0.2);
  backdrop-filter: blur(18px);
  transition: transform 0.25s ease, background 0.25s ease, box-shadow 0.25s ease;
}
.gg-call:hover { transform: translateY(-2px); background: linear-gradient(135deg, rgba(37, 99, 235, 0.62), rgba(124, 58, 237, 0.45)); box-shadow: 0 30px 80px -30px rgba(59, 130, 246, 0.8); }
.gg-call:active { transform: scale(0.985); }
.gg-row {
  border: 1px solid rgba(255, 255, 255, 0.1);
  background: rgba(255, 255, 255, 0.045);
  transition: background-color 0.2s ease, border-color 0.2s ease, transform 0.2s ease;
}
.gg-row:hover { background: rgba(255, 255, 255, 0.09); border-color: rgba(147, 197, 253, 0.35); transform: translateX(2px); }
.gg-row:active { transform: scale(0.985); }
.gg-enter { animation: gg-rise 0.8s cubic-bezier(0.22, 1, 0.36, 1) both; }
.gg-d1 { animation-delay: 0.08s; }
.gg-d2 { animation-delay: 0.16s; }
.gg-d3 { animation-delay: 0.24s; }
@keyframes gg-rise {
  from { opacity: 0; transform: translateY(16px) scale(0.985); filter: blur(4px); }
  to { opacity: 1; transform: translateY(0) scale(1); filter: blur(0); }
}
@keyframes gg-drift {
  from { transform: translate3d(0, 0, 0) scale(1); }
  to { transform: translate3d(-18px, 22px, 0) scale(1.12); }
}
@media (prefers-reduced-motion: reduce) {
  .gg-enter { animation: none; }
  .gg-orb { animation: none; }
  .gg-call, .gg-row { transition: none; }
}
</style>
