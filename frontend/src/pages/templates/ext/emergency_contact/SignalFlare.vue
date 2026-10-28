<template>
  <div class="min-h-screen bg-[#120505] text-[#fff7f5]">
    <main class="mx-auto w-full max-w-md px-5 pb-12 pt-6">
      <header class="sf-hero sf-enter relative overflow-hidden rounded-[28px] border border-red-500/25 px-6 pb-7 pt-6">
        <div class="sf-beacon" aria-hidden="true">
          <span></span><span></span><span></span>
        </div>
        <div class="relative flex items-center justify-between font-mono text-[10px] uppercase tracking-[0.24em] text-red-200/70">
          <span>Emergency Contact</span>
          <span v-if="nfcId" class="max-w-[42%] truncate">NFC · {{ nfcId }}</span>
        </div>
        <div class="relative mt-8 flex items-center gap-3.5">
          <span class="flex h-11 w-11 items-center justify-center rounded-2xl bg-red-600 text-white shadow-[0_0_28px_rgba(220,38,38,0.55)]">
            <Siren class="h-6 w-6" />
          </span>
          <div>
            <p class="font-mono text-[10px] uppercase tracking-[0.28em] text-red-300/80">Signal Active</p>
            <h1 class="mt-1 text-[26px] font-black leading-none tracking-tight">紧急联系人</h1>
          </div>
        </div>
        <p class="relative mt-5 max-w-[250px] text-[13px] leading-relaxed text-red-50/65">
          发现持卡人无法回应时，请先拨打电话，再告知所在位置与现场情况。
        </p>
      </header>

      <section class="mt-4">
        <a
          v-if="contacts.length"
          :href="'tel:' + contacts[0].phone"
          class="sf-call sf-enter sf-d1 block rounded-[24px] bg-red-600 px-6 pb-6 pt-5 text-white shadow-[0_26px_60px_-24px_rgba(220,38,38,0.8)]"
        >
          <div class="flex items-center justify-between font-mono text-[10px] uppercase tracking-[0.22em] text-red-100/85">
            <span>01 · Primary</span>
            <Phone class="h-4 w-4" />
          </div>
          <p class="mt-4 break-words text-[38px] font-black leading-[1.05] tabular-nums tracking-tight">{{ contacts[0].phone }}</p>
          <div class="mt-4 flex items-center justify-between border-t border-white/20 pt-3.5">
            <p class="text-[13px] font-semibold">
              {{ contacts[0].name || '联系人' }}<span v-if="contacts[0].relation" class="font-normal text-red-100/80"> · {{ contacts[0].relation }}</span>
            </p>
            <span class="font-mono text-[10px] uppercase tracking-[0.2em]">点击拨打</span>
          </div>
        </a>
        <div v-else class="sf-enter sf-d1 rounded-[24px] border border-dashed border-red-500/40 px-6 py-8 text-center text-sm text-red-100/55">
          暂无紧急联系人 · 请补充号码
        </div>

        <div v-if="contacts.length > 1" class="mt-3 space-y-2.5">
          <a
            v-for="(c, i) in contacts.slice(1)"
            :key="i"
            :href="'tel:' + c.phone"
            class="sf-row sf-enter sf-d2 flex items-center justify-between gap-4 rounded-2xl border border-white/10 bg-white/[0.035] px-5 py-4"
          >
            <div class="min-w-0">
              <p class="truncate text-[22px] font-bold leading-tight tabular-nums tracking-wide">{{ c.phone }}</p>
              <p class="mt-1 truncate text-[11px] text-white/45">
                {{ c.name || '联系人' }}<template v-if="c.relation"> · {{ c.relation }}</template>
              </p>
            </div>
            <span class="flex h-10 w-10 flex-shrink-0 items-center justify-center rounded-full border border-red-500/40 text-red-300">
              <Phone class="h-4 w-4" />
            </span>
          </a>
        </div>
      </section>

      <section class="sf-enter sf-d3 mt-4 rounded-2xl border border-white/10 bg-[#1a0909] p-4">
        <div class="flex items-center justify-between">
          <p class="font-mono text-[10px] uppercase tracking-[0.24em] text-red-300/60">Medical Brief</p>
          <HeartPulse class="h-4 w-4 text-red-400" />
        </div>
        <div v-if="hasMedical" class="mt-3 grid grid-cols-2 gap-2">
          <div v-for="m in medical" :key="m.label" class="rounded-xl bg-white/[0.04] px-3 py-2.5">
            <p class="text-[10px] text-white/35">{{ m.label }}</p>
            <p class="mt-1 text-[13px] font-semibold leading-snug text-white/90">{{ m.value }}</p>
          </div>
        </div>
      </section>

      <div v-if="hasPet" class="sf-enter sf-d3 mt-3 rounded-2xl border border-emerald-400/20 bg-emerald-400/[0.06] px-4 py-3 text-[12px] text-emerald-100/80">
        <span class="font-mono text-[10px] uppercase tracking-[0.2em] text-emerald-300/70">Pet</span>
        <span class="ml-2 font-semibold">{{ payload.petName }}</span>
        <span v-if="payload.breed"> · {{ payload.breed }}</span>
      </div>

      <p class="sf-enter sf-d3 mt-6 text-center font-mono text-[10px] uppercase tracking-[0.26em] text-white/25">
        Tap to call · 保持冷静，先拨号
      </p>
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
const hasMedical = computed(() =>
  !!(props.payload?.bloodType || props.payload?.allergies || props.payload?.medications || props.payload?.conditions)
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
.sf-hero {
  background:
    radial-gradient(120% 90% at 100% 0%, rgba(220, 38, 38, 0.24), transparent 58%),
    linear-gradient(160deg, #220909, #140505 70%);
}
.sf-beacon {
  position: absolute;
  right: -54px;
  top: -64px;
  width: 230px;
  height: 230px;
}
.sf-beacon span {
  position: absolute;
  inset: 0;
  border: 1px solid rgba(248, 113, 113, 0.25);
  border-radius: 999px;
  animation: sf-ping 4.8s ease-out infinite;
}
.sf-beacon span:nth-child(2) { animation-delay: 1.6s; }
.sf-beacon span:nth-child(3) { animation-delay: 3.2s; }
@keyframes sf-ping {
  0% { transform: scale(0.35); opacity: 0; }
  18% { opacity: 0.8; }
  100% { transform: scale(1); opacity: 0; }
}
.sf-enter { animation: sf-rise 0.7s cubic-bezier(0.22, 1, 0.36, 1) both; }
.sf-d1 { animation-delay: 0.08s; }
.sf-d2 { animation-delay: 0.16s; }
.sf-d3 { animation-delay: 0.24s; }
@keyframes sf-rise {
  from { opacity: 0; transform: translateY(16px); }
  to { opacity: 1; transform: translateY(0); }
}
.sf-call { transition: transform 0.25s ease, background-color 0.25s ease, box-shadow 0.25s ease; }
.sf-call:hover { background: #ef4444; transform: translateY(-2px); box-shadow: 0 32px 70px -24px rgba(239, 68, 68, 0.9); }
.sf-call:active { transform: scale(0.985); }
.sf-row { transition: border-color 0.2s ease, background-color 0.2s ease, transform 0.2s ease; }
.sf-row:hover { border-color: rgba(239, 68, 68, 0.45); background: rgba(239, 68, 68, 0.08); transform: translateX(2px); }
.sf-row:active { transform: scale(0.985); }
@media (prefers-reduced-motion: reduce) {
  .sf-beacon span { animation: none; opacity: 0.25; transform: scale(0.7); }
  .sf-enter { animation: none; }
  .sf-call, .sf-row { transition: none; }
}
</style>
