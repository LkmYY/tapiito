<template>
  <div class="bp-root min-h-screen text-[#d7f3ff]">
    <main class="mx-auto w-full max-w-md px-5 pb-12 pt-6">
      <div class="bp-enter flex items-center justify-between border border-[#67c7e8]/35 px-3 py-2 text-[9px] uppercase tracking-[0.18em] text-[#7fd4ec]">
        <span>Drawing · EC-025</span>
        <span v-if="nfcId" class="max-w-[44%] truncate">NFC {{ nfcId }}</span>
      </div>

      <header class="bp-enter bp-d1 relative mt-4 border border-[#67c7e8]/45 p-5">
        <span class="bp-corner bp-tl"></span><span class="bp-corner bp-tr"></span><span class="bp-corner bp-bl"></span><span class="bp-corner bp-br"></span>
        <div class="flex items-start justify-between gap-3">
          <div>
            <p class="text-[9px] uppercase tracking-[0.24em] text-[#6fb9d3]">Emergency Response Plan</p>
            <h1 class="mt-2 text-[29px] font-bold leading-tight tracking-tight text-white">紧急联络 / 医疗预案</h1>
          </div>
          <Crosshair class="h-7 w-7 shrink-0 text-[#67c7e8]" />
        </div>
        <div class="mt-4 grid grid-cols-3 gap-px border border-[#67c7e8]/25 bg-[#67c7e8]/20 text-[9px]">
          <div class="bg-[#06243d] px-2.5 py-2"><p class="text-[#6fb9d3]">SCALE</p><p class="mt-0.5 text-white">1:1</p></div>
          <div class="bg-[#06243d] px-2.5 py-2"><p class="text-[#6fb9d3]">REV</p><p class="mt-0.5 text-white">A</p></div>
          <div class="bg-[#06243d] px-2.5 py-2"><p class="text-[#6fb9d3]">SHEET</p><p class="mt-0.5 text-white">01 / 01</p></div>
        </div>
      </header>

      <section class="mt-5">
        <div class="flex items-center justify-between text-[9px] uppercase tracking-[0.18em] text-[#6fb9d3]">
          <span>Call-out schedule</span><span>优先顺序 ↓</span>
        </div>

        <template v-if="contacts.length">
          <a :href="'tel:' + contacts[0].phone" class="bp-primary bp-enter bp-d2 relative mt-2 block border border-dashed border-[#67c7e8]/70 px-5 py-5">
            <span class="absolute -left-px top-1/2 h-px w-4 -translate-x-full bg-[#67c7e8]/60"></span>
            <p class="text-[9px] uppercase tracking-[0.2em] text-[#67c7e8]">CALL-OUT 01 · PRIMARY</p>
            <p class="mt-2 break-words text-[35px] font-bold leading-none tabular-nums tracking-tight text-white">{{ contacts[0].phone }}</p>
            <div class="mt-3 flex items-center justify-between border-t border-[#67c7e8]/20 pt-2.5">
              <p class="text-[11px] text-[#b9e8f7]">{{ contacts[0].name || '联系人' }}<span v-if="contacts[0].relation" class="text-[#6fb9d3]"> / {{ contacts[0].relation }}</span></p>
              <span class="flex items-center gap-1.5 text-[9px] uppercase tracking-[0.14em] text-[#67c7e8]"><Phone class="h-3.5 w-3.5" />DIAL</span>
            </div>
          </a>

          <div class="mt-3 border border-[#67c7e8]/30">
            <div class="grid grid-cols-[30px_1fr_auto] items-center gap-3 border-b border-[#67c7e8]/30 bg-[#0a3352] px-3 py-2 text-[8px] uppercase tracking-[0.14em] text-[#7fd4ec]">
              <span>NO.</span><span>CONTACT / RELATION</span><span>NUMBER</span>
            </div>
            <a
              v-for="(c, i) in contacts.slice(1)"
              :key="i"
              :href="'tel:' + c.phone"
              class="bp-row grid grid-cols-[30px_1fr_auto] items-center gap-3 border-b border-[#67c7e8]/15 px-3 py-3 last:border-b-0"
            >
              <span class="text-[9px] tabular-nums text-[#6fb9d3]">{{ String(i + 2).padStart(2, '0') }}</span>
              <div class="min-w-0">
                <p class="text-[11px] text-[#c9eefb]">{{ c.name || '联系人' }}</p>
                <p class="text-[9px] text-[#6fb9d3]">{{ c.relation || '未注明' }}</p>
              </div>
              <span class="text-right text-[16px] font-semibold tabular-nums text-white">{{ c.phone }}</span>
            </a>
          </div>
        </template>
      </section>

      <section class="bp-enter bp-d3 mt-5 border border-[#67c7e8]/30 p-4">
        <div class="flex items-center justify-between">
          <p class="text-[9px] uppercase tracking-[0.2em] text-[#67c7e8]">General notes · 医疗信息</p>
          <Activity class="h-4 w-4 text-[#67c7e8]" />
        </div>
        <div class="mt-3 space-y-2">
          <div v-for="m in medical" :key="m.label" class="grid grid-cols-[58px_1fr] gap-2 border-b border-[#67c7e8]/15 pb-2 last:border-b-0 last:pb-0">
            <span class="text-[9px] uppercase tracking-[0.1em] text-[#6fb9d3]">{{ m.label }}</span>
            <span class="text-[11px] leading-snug text-[#d7f3ff]">{{ m.value }}</span>
          </div>
        </div>
      </section>

      <div v-if="hasPet" class="bp-enter bp-d3 mt-3 flex items-center justify-between border border-[#67c7e8]/30 px-4 py-2.5 text-[10px] text-[#b9e8f7]">
        <span>PET ASSET / {{ payload.petName }}<span v-if="payload.breed"> · {{ payload.breed }}</span></span>
        <a v-if="payload.vetPhone" :href="'tel:' + payload.vetPhone" class="text-[#67c7e8]">VET {{ payload.vetPhone }}</a>
      </div>

      <div class="bp-enter bp-d3 mt-5 flex items-center justify-between text-[8px] uppercase tracking-[0.2em] text-[#4f92ac]">
        <span>Emergency Contact · Field Copy</span><span>Keep dry · 妥善保存</span>
      </div>
    </main>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { Crosshair, Phone, Activity } from 'lucide-vue-next'

const props = defineProps(['payload', 'nfcId'])
const contacts = computed(() =>
  (Array.isArray(props.payload?.contacts) ? props.payload.contacts : []).filter((c) => c && c.phone)
)
const medical = computed(() => [
  { label: '血型', value: props.payload?.bloodType || '—' },
  { label: '过敏史', value: props.payload?.allergies },
  { label: '用药', value: props.payload?.medications },
  { label: '病史', value: props.payload?.conditions },
])
const hasPet = computed(() => !!(props.payload?.petName || props.payload?.chipId))
</script>

<style scoped>
.bp-root {
  background-color: #06243d;
  background-image:
    linear-gradient(rgba(125, 211, 252, 0.075) 1px, transparent 1px),
    linear-gradient(90deg, rgba(125, 211, 252, 0.075) 1px, transparent 1px),
    linear-gradient(rgba(125, 211, 252, 0.03) 1px, transparent 1px),
    linear-gradient(90deg, rgba(125, 211, 252, 0.03) 1px, transparent 1px);
  background-size: 40px 40px, 40px 40px, 8px 8px, 8px 8px;
}
.bp-corner { position: absolute; width: 10px; height: 10px; border-color: #9fe8ff; border-style: solid; }
.bp-tl { left: -1px; top: -1px; border-width: 2px 0 0 2px; }
.bp-tr { right: -1px; top: -1px; border-width: 2px 2px 0 0; }
.bp-bl { left: -1px; bottom: -1px; border-width: 0 0 2px 2px; }
.bp-br { right: -1px; bottom: -1px; border-width: 0 2px 2px 0; }
.bp-primary {
  background: linear-gradient(135deg, rgba(103, 199, 232, 0.1), rgba(103, 199, 232, 0.02));
  transition: background-color 0.2s ease, border-color 0.2s ease, transform 0.2s ease;
}
.bp-primary:hover { background: rgba(103, 199, 232, 0.16); border-color: #9fe8ff; transform: translateY(-2px); }
.bp-primary:active { transform: scale(0.99); }
.bp-row { transition: background-color 0.18s ease; }
.bp-row:hover { background: rgba(103, 199, 232, 0.1); }
.bp-enter { animation: bp-in 0.7s cubic-bezier(0.22, 1, 0.36, 1) both; }
.bp-d1 { animation-delay: 0.08s; }
.bp-d2 { animation-delay: 0.16s; }
.bp-d3 { animation-delay: 0.24s; }
@keyframes bp-in {
  from { opacity: 0; transform: translateY(10px); clip-path: inset(0 0 30% 0); }
  to { opacity: 1; transform: translateY(0); clip-path: inset(0 0 0 0); }
}
@media (prefers-reduced-motion: reduce) {
  .bp-enter { animation: none; }
  .bp-primary, .bp-row { transition: none; }
}
</style>
