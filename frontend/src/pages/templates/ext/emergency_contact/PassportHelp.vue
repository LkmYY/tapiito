<template>
  <div class="ph-root min-h-screen text-[#182335]">
    <main class="mx-auto w-full max-w-md px-4 pb-12 pt-6">
      <div class="ph-enter flex items-center justify-between text-[9px] uppercase tracking-[0.22em] text-[#c9b07a]">
        <span class="flex items-center gap-2"><Globe2 class="h-3.5 w-3.5" />Emergency Document</span>
        <span class="text-white/35">No. EC-2026</span>
      </div>

      <article class="ph-page ph-enter ph-d1 relative mt-4 overflow-hidden rounded-sm px-5 pb-0 pt-6">
        <div class="ph-guilloche" aria-hidden="true"></div>
        <header class="relative flex items-start justify-between gap-3 border-b border-[#1d3556]/20 pb-4">
          <div>
            <p class="text-[8px] font-bold uppercase tracking-[0.26em] text-[#9a7b43]">Emergency Contact Page</p>
            <h1 class="ph-serif mt-1 text-[29px] font-semibold leading-tight tracking-wide text-[#13253f]">紧急联络护照页</h1>
            <p class="mt-1 text-[9px] uppercase tracking-[0.18em] text-[#6f7c8d]">Please assist the bearer</p>
          </div>
          <div class="flex h-12 w-12 shrink-0 items-center justify-center rounded-full border border-[#9a7b43]/60 text-[#9a7b43]">
            <ShieldCheck class="h-6 w-6" />
          </div>
        </header>

        <section class="relative mt-4">
          <p class="text-[8px] font-bold uppercase tracking-[0.22em] text-[#9a7b43]">Primary contact / 首要联络人</p>
          <template v-if="contacts.length">
            <a :href="'tel:' + contacts[0].phone" class="ph-primary mt-2 block border border-[#1d3556]/25 bg-[#fbf8ee] px-4 py-4">
              <p class="break-words text-[31px] font-bold leading-none tabular-nums tracking-wide text-[#13253f]">{{ contacts[0].phone }}</p>
              <div class="mt-2 flex items-center justify-between">
                <p class="text-[11px] font-semibold">{{ contacts[0].name || '联系人' }}<span v-if="contacts[0].relation" class="font-normal text-[#6f7c8d]"> / {{ contacts[0].relation }}</span></p>
                <span class="text-[8px] font-bold uppercase tracking-[0.14em] text-[#9a7b43]">Dial</span>
              </div>
            </a>

            <div class="mt-3 space-y-0">
              <a
                v-for="(c, i) in contacts.slice(1)"
                :key="i"
                :href="'tel:' + c.phone"
                class="ph-row grid grid-cols-[1fr_auto] items-center gap-3 border-b border-[#1d3556]/15 py-2.5"
              >
                <div class="min-w-0">
                  <p class="text-[8px] uppercase tracking-[0.12em] text-[#8a94a2]">Contact {{ String(i + 2).padStart(2, '0') }} · {{ c.relation || '备用' }}</p>
                  <p class="mt-0.5 truncate text-[19px] font-bold leading-tight tabular-nums text-[#13253f]">{{ c.phone }}</p>
                  <p class="truncate text-[9px] text-[#6f7c8d]">{{ c.name || '联系人' }}</p>
                </div>
                <span class="flex h-7 w-7 items-center justify-center rounded-full border border-[#9a7b43]/50 text-[#9a7b43]"><Phone class="h-3 w-3" /></span>
              </a>
            </div>
          </template>
          <div v-else class="mt-2 border border-dashed border-[#1d3556]/30 px-4 py-6 text-center text-[11px] text-[#6f7c8d]">本页未登记联系人</div>
        </section>

        <section class="relative mt-4 border-t border-[#1d3556]/20 pt-3">
          <div class="flex items-center justify-between">
            <p class="text-[8px] font-bold uppercase tracking-[0.22em] text-[#9a7b43]">Health endorsements / 医疗备注</p>
            <HeartPulse class="h-3.5 w-3.5 text-[#9a7b43]" />
          </div>
          <div class="mt-2 grid grid-cols-2 gap-x-4 gap-y-2.5">
            <div v-for="m in medical" :key="m.label" class="border-b border-[#1d3556]/10 pb-1.5">
              <p class="text-[8px] uppercase tracking-[0.1em] text-[#8a94a2]">{{ m.label }}</p>
              <p class="mt-0.5 text-[10px] font-medium leading-snug text-[#253a54]">{{ m.value }}</p>
            </div>
          </div>
        </section>

        <div v-if="hasPet" class="relative mt-3 border-t border-[#1d3556]/15 pt-2.5 text-[10px] text-[#4f5f73]">
          <span class="font-bold uppercase tracking-[0.12em] text-[#9a7b43]">Pet /</span> {{ payload.petName }}<span v-if="payload.breed"> · {{ payload.breed }}</span>
          <a v-if="payload.vetPhone" :href="'tel:' + payload.vetPhone" class="ml-1 font-semibold text-[#13253f] underline underline-offset-2">{{ payload.vetPhone }}</a>
        </div>

        <div class="ph-stamp absolute bottom-16 right-4 flex h-16 w-16 -rotate-12 flex-col items-center justify-center rounded-full border-2 border-[#a5342a]/70 text-center text-[7px] font-bold uppercase leading-tight tracking-[0.1em] text-[#a5342a]/80">
          Help<br />Needed<br />2026
        </div>

        <div class="ph-mrz relative -mx-5 mt-5 bg-[#13253f] px-4 py-2.5 font-mono text-[9px] leading-5 tracking-[0.1em] text-[#d9c9a3]">
          <p>NFC&lt;{{ (nfcId || 'UNBOUND').toString().toUpperCase() }}</p>
          <p>EMERGENCY&lt;CONTACT&lt;&lt;KEEP&lt;THIS&lt;PAGE</p>
        </div>
      </article>

      <p class="ph-enter ph-d3 mt-4 text-center text-[8px] uppercase tracking-[0.24em] text-[#8ea0b8]">Official assistance record · 请出示给救援人员</p>
    </main>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { Globe2, ShieldCheck, Phone, HeartPulse } from 'lucide-vue-next'

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
.ph-root {
  background:
    radial-gradient(rgba(201, 176, 122, 0.12) 0.7px, transparent 0.7px),
    #0d1b2f;
  background-size: 5px 5px;
}
.ph-page {
  background: linear-gradient(160deg, #faf6ea, #f1ead8);
  box-shadow: 0 30px 70px -40px #000, inset 0 0 0 1px rgba(154, 123, 67, 0.35);
}
.ph-guilloche {
  position: absolute;
  inset: 0;
  background:
    repeating-radial-gradient(ellipse at 20% 15%, rgba(29, 53, 86, 0.04) 0 1px, transparent 1px 7px),
    repeating-radial-gradient(ellipse at 85% 80%, rgba(154, 123, 67, 0.045) 0 1px, transparent 1px 9px);
  pointer-events: none;
}
.ph-serif { font-family: "Iowan Old Style", "Palatino Linotype", Georgia, "Songti SC", serif; }
.ph-primary {
  transition: border-color 0.2s ease, background-color 0.2s ease, transform 0.2s ease;
  box-shadow: 2px 3px 0 rgba(19, 37, 63, 0.05);
}
.ph-primary:hover { border-color: #9a7b43; background: #fffdf7; transform: translateY(-1px); }
.ph-primary:active { transform: scale(0.99); }
.ph-row { transition: background-color 0.18s ease; }
.ph-row:hover { background: rgba(19, 37, 63, 0.035); }
.ph-stamp { transform: rotate(-12deg); background: rgba(250, 246, 234, 0.6); }
.ph-enter { animation: ph-in 0.8s cubic-bezier(0.22, 1, 0.36, 1) both; }
.ph-d1 { animation-delay: 0.08s; }
.ph-d3 { animation-delay: 0.24s; }
@keyframes ph-in {
  from { opacity: 0; transform: translateY(14px); }
  to { opacity: 1; transform: translateY(0); }
}
@media (prefers-reduced-motion: reduce) {
  .ph-enter { animation: none; }
  .ph-primary, .ph-row { transition: none; }
}
</style>
