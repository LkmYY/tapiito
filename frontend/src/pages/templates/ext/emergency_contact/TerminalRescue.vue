<template>
  <div class="tr-root min-h-screen text-[#b8ffd0]">
    <main class="mx-auto w-full max-w-md px-4 pb-12 pt-6 font-mono">
      <div class="tr-enter flex items-center justify-between rounded-t-xl border border-[#1f6b3d] bg-[#082014] px-4 py-2.5 text-[10px] uppercase tracking-[0.16em] text-[#6ee7a0]">
        <span class="flex items-center gap-2"><span class="h-2 w-2 rounded-full bg-[#4ade80] tr-blink"></span>rescue@nfc:~</span>
        <span v-if="nfcId" class="max-w-[42%] truncate text-[#3f8a5e]">{{ nfcId }}</span>
      </div>

      <header class="tr-enter tr-d1 relative overflow-hidden border-x border-b border-[#1f6b3d] bg-[#04120a]/95 px-5 pb-6 pt-5">
        <div class="tr-scan" aria-hidden="true"></div>
        <p class="text-[10px] tracking-[0.18em] text-[#3f8a5e]">// EMERGENCY_CONTACT_PROTOCOL</p>
        <div class="relative mt-3">
          <h1 class="text-[28px] font-bold leading-tight tracking-tight text-[#d9ffe6]">
            <span class="text-[#4ade80]">&gt;</span> 紧急联系人<span class="tr-caret">_</span>
          </h1>
          <p class="mt-3 text-[11px] leading-5 text-[#72b78b]">$ status: standby · 请按优先级拨号<br />$ tip: 告知位置后保持通话</p>
        </div>
        <div class="mt-4 flex items-center gap-3 border-t border-dashed border-[#1f6b3d] pt-3 text-[9px] uppercase tracking-[0.14em] text-[#4b9565]">
          <span>links {{ contacts.length }}</span><span>·</span><span>med {{ hasMedical ? 'loaded' : 'none' }}</span><span>·</span><span>secure</span>
        </div>
      </header>

      <section class="mt-4">
        <p class="tr-enter tr-d2 px-1 text-[10px] uppercase tracking-[0.18em] text-[#4b9565]">$ call --priority</p>
        <template v-if="contacts.length">
          <a
            :href="'tel:' + contacts[0].phone"
            class="tr-primary tr-enter tr-d2 mt-2 block border border-[#4ade80] bg-[#0a2d19] px-5 py-5 shadow-[0_0_30px_-10px_rgba(74,222,128,0.35)]"
          >
            <div class="flex items-center justify-between text-[9px] uppercase tracking-[0.2em] text-[#4ade80]">
              <span>[01] primary_target</span><Phone class="h-4 w-4" />
            </div>
            <p class="mt-3 break-words text-[36px] font-bold leading-none tabular-nums tracking-tight text-[#e9fff0]">{{ contacts[0].phone }}</p>
            <p class="mt-3 text-[11px] text-[#8fd6a6]">{{ contacts[0].name || '联系人' }}<span v-if="contacts[0].relation"> / {{ contacts[0].relation }}</span> <span class="text-[#4ade80]">[ dial ]</span></p>
          </a>

          <div class="mt-2 border border-[#1f6b3d] bg-[#04120a]">
            <a
              v-for="(c, i) in contacts.slice(1)"
              :key="i"
              :href="'tel:' + c.phone"
              class="tr-row flex items-center justify-between gap-3 border-b border-[#123c25] px-4 py-3.5 last:border-b-0"
            >
              <div class="flex min-w-0 items-center gap-3">
                <span class="text-[10px] text-[#3f8a5e]">{{ String(i + 2).padStart(2, '0') }}</span>
                <div class="min-w-0">
                  <p class="truncate text-[21px] font-semibold leading-tight tabular-nums text-[#d9ffe6]">{{ c.phone }}</p>
                  <p class="mt-0.5 truncate text-[10px] text-[#5f9e75]">{{ c.name || '联系人' }}<span v-if="c.relation"> · {{ c.relation }}</span></p>
                </div>
              </div>
              <span class="text-[9px] text-[#4ade80]">CALL</span>
            </a>
          </div>
        </template>
      </section>

      <section class="tr-enter tr-d3 mt-4 border border-[#1f6b3d] bg-[#04120a] px-4 py-4">
        <p class="text-[10px] uppercase tracking-[0.18em] text-[#4ade80]">$ cat medical.log</p>
        <div class="mt-3 space-y-2 text-[11px]">
          <div v-for="m in medical" :key="m.label" class="grid grid-cols-[52px_1fr] gap-3 border-b border-[#123c25] pb-2 last:border-b-0 last:pb-0">
            <span class="text-[#4b9565]">{{ m.label }}</span>
            <span class="leading-snug text-[#b8ffd0]">{{ m.value }}</span>
          </div>
        </div>
      </section>

      <div v-if="hasPet" class="tr-enter tr-d3 mt-3 border border-[#4a7c2a] bg-[#132009] px-4 py-3 text-[11px] text-[#d7f5a8]">
        <span class="text-[#86a95f]">pet:</span> {{ payload.petName }}<span v-if="payload.breed"> / {{ payload.breed }}</span>
        <a v-if="payload.vetPhone" :href="'tel:' + payload.vetPhone" class="ml-2 text-[#bef264] underline underline-offset-2">vet {{ payload.vetPhone }}</a>
      </div>

      <p class="tr-enter tr-d3 mt-5 text-center text-[9px] uppercase tracking-[0.2em] text-[#3f8a5e]">end of record · tap number to call</p>
    </main>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { Phone } from 'lucide-vue-next'

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
.tr-root {
  background:
    linear-gradient(rgba(74, 222, 128, 0.035) 1px, transparent 1px),
    #020a05;
  background-size: 100% 28px;
}
.tr-enter { animation: tr-in 0.55s steps(6, end) both; }
.tr-d1 { animation-delay: 0.1s; }
.tr-d2 { animation-delay: 0.2s; }
.tr-d3 { animation-delay: 0.3s; }
@keyframes tr-in {
  from { opacity: 0; transform: translateY(8px); clip-path: inset(0 0 100% 0); }
  to { opacity: 1; transform: translateY(0); clip-path: inset(0 0 0 0); }
}
.tr-primary { transition: background-color 0.2s ease, box-shadow 0.2s ease; }
.tr-primary:hover { background: #0f4226; box-shadow: 0 0 42px -8px rgba(74, 222, 128, 0.5); }
.tr-primary:active { transform: scale(0.99); }
.tr-row { transition: background-color 0.18s ease; }
.tr-row:hover { background: #0a2d19; }
.tr-scan {
  position: absolute;
  inset: 0;
  background: repeating-linear-gradient(0deg, transparent 0 3px, rgba(74, 222, 128, 0.035) 3px 4px);
  pointer-events: none;
}
.tr-blink { animation: tr-blink 1.4s steps(2, start) infinite; }
.tr-caret { animation: tr-blink 1.1s steps(2, start) infinite; }
@keyframes tr-blink { 50% { opacity: 0.15; } }
@media (prefers-reduced-motion: reduce) {
  .tr-enter, .tr-blink, .tr-caret { animation: none; }
  .tr-primary, .tr-row { transition: none; }
}
</style>
