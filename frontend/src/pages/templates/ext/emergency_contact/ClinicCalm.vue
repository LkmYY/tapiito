<template>
  <div class="min-h-screen bg-[#eef5f2] text-[#12332e]">
    <main class="mx-auto w-full max-w-md px-5 pb-12 pt-6">
      <div class="cc-top cc-enter flex items-center justify-between text-[10px] font-semibold uppercase tracking-[0.22em] text-[#3f6b62]">
        <span class="flex items-center gap-2"><span class="h-1.5 w-1.5 rounded-full bg-[#0f766e]"></span>Rapid Response</span>
        <span v-if="nfcId" class="max-w-[45%] truncate">NFC · {{ nfcId }}</span>
      </div>

      <header class="cc-enter cc-d1 mt-4 grid grid-cols-[1.12fr_0.88fr] overflow-hidden rounded-[26px] border border-[#bfd9d1] bg-white">
        <div class="cc-title p-5">
          <Stethoscope class="h-5 w-5 text-[#0f766e]" />
          <h1 class="mt-5 text-[26px] font-bold leading-[1.15] tracking-tight">紧急联系<br />与医疗摘要</h1>
          <p class="mt-3 text-[11px] leading-relaxed text-[#60807a]">请优先联系家属，并向急救人员说明过敏与用药。</p>
        </div>
        <div class="cc-summary flex flex-col justify-between p-4 text-white">
          <div>
            <p class="text-[10px] uppercase tracking-[0.2em] text-[#b8ebe2]">Blood</p>
            <p class="mt-1 text-[46px] font-black leading-none">{{ payload.bloodType || '—' }}</p>
          </div>
          <div class="mt-5 border-t border-white/25 pt-3">
            <p class="text-[10px] uppercase tracking-[0.16em] text-[#b8ebe2]">Allergy</p>
            <p class="mt-1 line-clamp-2 text-[12px] font-semibold leading-snug">{{ payload.allergies }}</p>
          </div>
        </div>
      </header>

      <section class="mt-6">
        <div class="flex items-end justify-between">
          <p class="text-[11px] font-bold uppercase tracking-[0.18em] text-[#3f6b62]">Call Sequence · 拨打顺序</p>
          <span class="text-[10px] tabular-nums text-[#7d9a94]">{{ String(contacts.length).padStart(2, '0') }} CONTACTS</span>
        </div>

        <template v-if="contacts.length">
          <a
            :href="'tel:' + contacts[0].phone"
            class="cc-primary cc-enter cc-d2 mt-3 block rounded-2xl border border-[#0f766e]/20 bg-white p-5 shadow-[0_18px_40px_-30px_rgba(15,118,110,0.55)]"
          >
            <div class="flex items-start justify-between gap-3">
              <div>
                <p class="text-[10px] font-bold uppercase tracking-[0.2em] text-[#0f766e]">Primary · 首要</p>
                <p class="mt-2 text-[13px] font-semibold text-[#254d46]">{{ contacts[0].name || '联系人' }}<span v-if="contacts[0].relation" class="font-normal text-[#78938d]"> / {{ contacts[0].relation }}</span></p>
              </div>
              <span class="flex h-10 w-10 items-center justify-center rounded-full bg-[#0f766e] text-white"><Phone class="h-4 w-4" /></span>
            </div>
            <p class="mt-3 break-words text-[34px] font-bold leading-tight tabular-nums tracking-tight text-[#0b3f38]">{{ contacts[0].phone }}</p>
          </a>

          <div class="mt-2.5 overflow-hidden rounded-2xl border border-[#bfd9d1] bg-white/70">
            <a
              v-for="(c, i) in contacts.slice(1)"
              :key="i"
              :href="'tel:' + c.phone"
              class="cc-row flex items-center gap-4 border-b border-[#d8e7e2] px-4 py-3.5 last:border-b-0"
            >
              <span class="w-5 text-[10px] font-bold tabular-nums text-[#8aa49e]">{{ String(i + 2).padStart(2, '0') }}</span>
              <div class="min-w-0 flex-1">
                <p class="truncate text-[20px] font-semibold leading-tight tabular-nums text-[#12332e]">{{ c.phone }}</p>
                <p class="mt-0.5 truncate text-[11px] text-[#6b8882]">{{ c.name || '联系人' }}<span v-if="c.relation"> · {{ c.relation }}</span></p>
              </div>
              <Phone class="h-4 w-4 flex-shrink-0 text-[#0f766e]" />
            </a>
          </div>
        </template>
        <div v-else class="mt-3 rounded-2xl border border-dashed border-[#a9c9c1] bg-white/60 px-5 py-8 text-center text-sm text-[#5c7d76]">暂无紧急联系人，请补充号码。</div>
      </section>

      <section class="cc-enter cc-d3 mt-6 rounded-2xl border border-[#bfd9d1] bg-white">
        <div class="flex items-center gap-2 border-b border-[#d8e7e2] px-4 py-3">
          <Activity class="h-4 w-4 text-[#0f766e]" />
          <p class="text-[11px] font-bold uppercase tracking-[0.18em] text-[#3f6b62]">Patient Notes · 医疗信息</p>
        </div>
        <div class="grid grid-cols-2 gap-px bg-[#d8e7e2]">
          <div v-for="m in medical" :key="m.label" class="bg-white px-4 py-3.5">
            <p class="text-[10px] font-semibold uppercase tracking-[0.14em] text-[#90a8a2]">{{ m.label }}</p>
            <p class="mt-1 text-[13px] font-semibold leading-snug" :class="m.warn ? 'text-[#b45309]' : 'text-[#254d46]'">{{ m.value }}</p>
          </div>
        </div>
      </section>

      <div v-if="hasPet" class="cc-enter cc-d3 mt-3 flex items-center justify-between rounded-2xl border border-[#c7d8b8] bg-[#f3f8ea] px-4 py-3 text-[12px] text-[#405c33]">
        <span><strong>{{ payload.petName }}</strong><span v-if="payload.breed"> · {{ payload.breed }}</span></span>
        <a v-if="payload.vetPhone" :href="'tel:' + payload.vetPhone" class="font-semibold text-[#0f766e] underline decoration-[#0f766e]/30 underline-offset-4">宠物医院</a>
      </div>

      <p class="cc-enter cc-d3 mt-6 text-center text-[10px] font-semibold uppercase tracking-[0.2em] text-[#8aa49e]">Stay calm · 保持通话畅通</p>
    </main>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { Stethoscope, Phone, Activity } from 'lucide-vue-next'

const props = defineProps(['payload', 'nfcId'])
const contacts = computed(() =>
  (Array.isArray(props.payload?.contacts) ? props.payload.contacts : []).filter((c) => c && c.phone)
)
const medical = computed(() => [
  { label: '血型', value: props.payload?.bloodType || '—', warn: false },
  { label: '过敏史', value: props.payload?.allergies, warn: true },
  { label: '正在用药', value: props.payload?.medications, warn: false },
  { label: '既往病史', value: props.payload?.conditions, warn: false },
])
const hasPet = computed(() => !!(props.payload?.petName || props.payload?.chipId))
</script>

<style scoped>
.cc-title {
  background:
    radial-gradient(80% 60% at 0% 0%, rgba(15, 118, 110, 0.09), transparent 60%),
    #ffffff;
}
.cc-summary {
  background: linear-gradient(170deg, #0f766e, #115e59);
}
.cc-enter { animation: cc-rise 0.7s cubic-bezier(0.22, 1, 0.36, 1) both; }
.cc-d1 { animation-delay: 0.08s; }
.cc-d2 { animation-delay: 0.16s; }
.cc-d3 { animation-delay: 0.24s; }
@keyframes cc-rise {
  from { opacity: 0; transform: translateY(12px); }
  to { opacity: 1; transform: translateY(0); }
}
.cc-primary { transition: border-color 0.2s ease, transform 0.2s ease, box-shadow 0.2s ease; }
.cc-primary:hover { border-color: #0f766e; transform: translateY(-2px); box-shadow: 0 24px 50px -28px rgba(15, 118, 110, 0.6); }
.cc-primary:active { transform: scale(0.985); }
.cc-row { transition: background-color 0.18s ease, padding-left 0.18s ease; }
.cc-row:hover { background: #eef7f4; padding-left: 1.15rem; }
.cc-row:active { background: #e2f1ed; }
@media (prefers-reduced-motion: reduce) {
  .cc-enter { animation: none; }
  .cc-primary, .cc-row { transition: none; }
}
</style>
