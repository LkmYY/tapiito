<template>
  <div class="op-root min-h-screen text-[#dce7f3]">
    <main class="mx-auto w-full max-w-md px-4 pb-12 pt-5">
      <div class="op-enter flex items-center justify-between rounded-lg border border-white/10 bg-[#0e1621] px-3 py-2 text-[9px] uppercase tracking-[0.14em] text-[#7f93a8]">
        <span class="flex items-center gap-2"><span class="h-1.5 w-1.5 rounded-full bg-[#34d399] op-live"></span>OPS · Emergency</span>
        <span class="flex items-center gap-2"><span class="text-[#34d399]">READY</span><span v-if="nfcId" class="max-w-[36%] truncate text-[#53687c]">{{ nfcId }}</span></span>
      </div>

      <header class="op-enter op-d1 mt-4 flex items-start justify-between gap-3">
        <div>
          <p class="text-[9px] font-bold uppercase tracking-[0.22em] text-[#38bdf8]">On-call channel · 01</p>
          <h1 class="mt-1 text-[30px] font-black leading-tight tracking-tight text-white">值班紧急联络</h1>
          <p class="mt-2 text-[11px] leading-relaxed text-[#7f93a8]">选择可用线路，说明现场位置与人员状态。</p>
        </div>
        <Activity class="mt-1 h-7 w-7 shrink-0 text-[#38bdf8]" />
      </header>

      <section class="op-enter op-d2 mt-4 grid grid-cols-3 gap-2">
        <div class="op-stat">
          <p class="text-[8px] uppercase tracking-[0.12em] text-[#6c8196]">Contacts</p>
          <p class="mt-1 text-[20px] font-black leading-none tabular-nums text-white">{{ String(contacts.length).padStart(2, '0') }}</p>
          <p class="mt-1 text-[8px] text-[#34d399]">LINKED</p>
        </div>
        <div class="op-stat">
          <p class="text-[8px] uppercase tracking-[0.12em] text-[#6c8196]">Med flags</p>
          <p class="mt-1 text-[20px] font-black leading-none tabular-nums text-white">{{ String(medicalCount).padStart(2, '0') }}</p>
          <p class="mt-1 text-[8px]" :class="medicalCount ? 'text-[#fbbf24]' : 'text-[#6c8196]'">{{ medicalCount ? 'REVIEW' : 'CLEAR' }}</p>
        </div>
        <div class="op-stat">
          <p class="text-[8px] uppercase tracking-[0.12em] text-[#6c8196]">NFC tag</p>
          <p class="mt-1 text-[20px] font-black leading-none text-white">{{ nfcId ? 'ON' : '—' }}</p>
          <p class="mt-1 text-[8px] text-[#38bdf8]">BOUND</p>
        </div>
      </section>

      <section class="mt-4">
        <template v-if="contacts.length">
          <a :href="'tel:' + contacts[0].phone" class="op-primary op-enter op-d2 block px-5 py-5">
            <div class="flex items-center justify-between">
              <span class="inline-flex items-center gap-2 rounded bg-[#38bdf8]/15 px-2 py-1 text-[8px] font-bold uppercase tracking-[0.14em] text-[#7dd3fc]"><span class="h-1.5 w-1.5 rounded-full bg-[#38bdf8]"></span>Priority line</span>
              <Phone class="h-4 w-4 text-[#38bdf8]" />
            </div>
            <p class="mt-3 break-words text-[34px] font-black leading-none tabular-nums tracking-tight text-white">{{ contacts[0].phone }}</p>
            <div class="mt-4 flex items-center justify-between border-t border-white/10 pt-3">
              <p class="text-[12px] font-semibold text-[#c3d4e4]">{{ contacts[0].name || '联系人' }}<span v-if="contacts[0].relation" class="font-normal text-[#6c8196]"> / {{ contacts[0].relation }}</span></p>
              <span class="op-connect inline-flex items-center gap-1.5 rounded bg-[#38bdf8] px-3 py-1.5 text-[9px] font-black uppercase tracking-[0.12em] text-[#06121d]"><Phone class="h-3 w-3" />Connect</span>
            </div>
          </a>

          <div class="mt-3 overflow-hidden rounded-lg border border-white/10">
            <div class="grid grid-cols-[1fr_auto] items-center gap-2 border-b border-white/10 bg-[#0e1621] px-3 py-2 text-[8px] uppercase tracking-[0.14em] text-[#6c8196]">
              <span>Standby lines · 备用线路</span><span>Status</span>
            </div>
            <a
              v-for="(c, i) in contacts.slice(1)"
              :key="i"
              :href="'tel:' + c.phone"
              class="op-row flex items-center justify-between gap-3 border-b border-white/5 px-3 py-3 last:border-b-0"
            >
              <div class="flex min-w-0 items-center gap-3">
                <span class="flex h-6 w-6 shrink-0 items-center justify-center rounded bg-white/5 text-[8px] tabular-nums text-[#7f93a8]">{{ String(i + 2).padStart(2, '0') }}</span>
                <div class="min-w-0">
                  <p class="truncate text-[19px] font-bold leading-tight tabular-nums text-white">{{ c.phone }}</p>
                  <p class="truncate text-[9px] text-[#6c8196]">{{ c.name || '联系人' }}<span v-if="c.relation"> · {{ c.relation }}</span></p>
                </div>
              </div>
              <span class="text-[8px] font-bold uppercase tracking-[0.1em] text-[#34d399]">Ready</span>
            </a>
          </div>
        </template>
      </section>

      <section class="op-enter op-d3 mt-4 rounded-lg border border-white/10 bg-[#0b141e] p-4">
        <div class="flex items-center justify-between">
          <p class="text-[9px] font-bold uppercase tracking-[0.18em] text-[#7dd3fc]">Medical flags · 医疗标记</p>
          <HeartPulse class="h-4 w-4" :class="medicalCount ? 'text-[#fbbf24]' : 'text-[#38bdf8]'" />
        </div>
        <div class="mt-3 grid grid-cols-2 gap-2">
          <div v-for="m in medical" :key="m.label" class="rounded border border-white/8 bg-white/[0.025] px-3 py-2.5">
            <p class="text-[8px] uppercase tracking-[0.1em] text-[#6c8196]">{{ m.label }}</p>
            <p class="mt-1 text-[11px] font-semibold leading-snug text-[#dce7f3]">{{ m.value }}</p>
          </div>
        </div>
        <div v-if="payload.allergies" class="mt-2 flex items-center gap-2 rounded border border-[#fbbf24]/25 bg-[#fbbf24]/[0.07] px-3 py-2 text-[9px] text-[#fcd34d]">
          <AlertTriangle class="h-3.5 w-3.5 shrink-0" /> 过敏提示：{{ payload.allergies }}
        </div>
      </section>

      <div v-if="hasPet" class="op-enter op-d3 mt-3 flex items-center justify-between rounded-lg border border-white/10 bg-white/[0.03] px-3 py-2.5 text-[10px]">
        <span class="text-[#9fb2c4]">{{ payload.petName }}<span v-if="payload.breed"> / {{ payload.breed }}</span></span>
        <a v-if="payload.vetPhone" :href="'tel:' + payload.vetPhone" class="font-bold text-[#38bdf8]">VET LINE</a>
      </div>

      <div class="op-enter op-d3 mt-4 flex items-center justify-between text-[8px] uppercase tracking-[0.16em] text-[#42576b]">
        <span>All channels monitored · 别挂断</span><span>OPS v2.6</span>
      </div>
    </main>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { Activity, Phone, HeartPulse, AlertTriangle } from 'lucide-vue-next'

const props = defineProps(['payload', 'nfcId'])
const contacts = computed(() =>
  (Array.isArray(props.payload?.contacts) ? props.payload.contacts : []).filter((c) => c && c.phone)
)
const medicalCount = computed(() =>
  [props.payload?.bloodType, props.payload?.allergies, props.payload?.medications, props.payload?.conditions].filter(Boolean).length
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
.op-root {
  background:
    radial-gradient(90% 45% at 50% 0%, rgba(56, 189, 248, 0.09), transparent 65%),
    #070c12;
}
.op-stat {
  border: 1px solid rgba(255, 255, 255, 0.08);
  background: #0b141e;
  border-radius: 7px;
  padding: 10px 11px;
}
.op-primary {
  border: 1px solid rgba(56, 189, 248, 0.4);
  border-radius: 9px;
  background: linear-gradient(155deg, rgba(56, 189, 248, 0.1), rgba(56, 189, 248, 0.015));
  box-shadow: 0 20px 45px -32px rgba(56, 189, 248, 0.6);
  transition: border-color 0.2s ease, background 0.2s ease, transform 0.2s ease;
}
.op-primary:hover { border-color: #7dd3fc; background: rgba(56, 189, 248, 0.14); transform: translateY(-2px); }
.op-primary:active { transform: scale(0.99); }
.op-connect { box-shadow: 0 0 18px -4px rgba(56, 189, 248, 0.6); transition: background-color 0.18s ease, transform 0.18s ease; }
.op-primary:hover .op-connect { background: #7dd3fc; transform: translateX(1px); }
.op-row { transition: background-color 0.18s ease; }
.op-row:hover { background: rgba(56, 189, 248, 0.06); }
.op-live { animation: op-pulse 2s ease-in-out infinite; }
.op-enter { animation: op-boot 0.65s cubic-bezier(0.22, 1, 0.36, 1) both; }
.op-d1 { animation-delay: 0.08s; }
.op-d2 { animation-delay: 0.16s; }
.op-d3 { animation-delay: 0.24s; }
@keyframes op-boot {
  from { opacity: 0; transform: translateY(10px); }
  to { opacity: 1; transform: translateY(0); }
}
@keyframes op-pulse { 50% { opacity: 0.3; } }
@media (prefers-reduced-motion: reduce) {
  .op-enter, .op-live { animation: none; }
  .op-primary, .op-connect, .op-row { transition: none; }
}
</style>
