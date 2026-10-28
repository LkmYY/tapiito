<template>
  <div class="et-root min-h-screen text-[#202326]">
    <main class="mx-auto w-full max-w-md px-4 pb-12 pt-5">
      <div class="et-enter flex items-center justify-between text-[9px] uppercase tracking-[0.22em] text-[#aeb4b8]">
        <span>Asset Tag · Emergency</span>
        <span v-if="nfcId" class="max-w-[42%] truncate">NFC {{ nfcId }}</span>
      </div>

      <article class="et-plate et-enter et-d1 relative mt-4 overflow-hidden rounded-lg px-5 pb-5 pt-6">
        <span class="et-rivet et-rivet-tl"></span><span class="et-rivet et-rivet-tr"></span><span class="et-rivet et-rivet-bl"></span><span class="et-rivet et-rivet-br"></span>

        <header class="relative border-b-2 border-[#202326]/60 pb-3">
          <div class="flex items-start justify-between gap-3">
            <div>
              <p class="text-[8px] font-black uppercase tracking-[0.24em] text-[#575d61]">Emergency Contact Plate</p>
              <h1 class="mt-1 text-[27px] font-black leading-tight tracking-[0.02em] text-[#181b1e]">紧急联络铭牌</h1>
            </div>
            <span class="flex h-11 w-11 shrink-0 items-center justify-center rounded-sm bg-[#202326] text-[#f26b21]"><Siren class="h-5 w-5" /></span>
          </div>
          <div class="mt-3 flex items-center justify-between text-[8px] font-bold uppercase tracking-[0.14em] text-[#4f5559]">
            <span>SPEC / EC-25 · REV.03</span><span>MADE FOR FIELD USE</span>
          </div>
        </header>

        <section class="relative mt-4">
          <template v-if="contacts.length">
            <a :href="'tel:' + contacts[0].phone" class="et-primary block px-4 py-4">
              <div class="flex items-center justify-between text-[8px] font-black uppercase tracking-[0.18em] text-[#f26b21]">
                <span>Primary call point · 01</span><Phone class="h-3.5 w-3.5" />
              </div>
              <p class="mt-2 break-words text-[33px] font-black leading-none tabular-nums tracking-[0.01em] text-[#f5f6f3]">{{ contacts[0].phone }}</p>
              <div class="mt-3 flex items-center justify-between border-t border-white/20 pt-2.5">
                <p class="text-[11px] font-bold text-[#dfe3e0]">{{ contacts[0].name || '联系人' }}<span v-if="contacts[0].relation" class="font-normal text-[#9aa3a0]"> / {{ contacts[0].relation }}</span></p>
                <span class="text-[8px] font-black uppercase tracking-[0.12em] text-[#f26b21]">Press to call</span>
              </div>
            </a>

            <div class="mt-3 border border-[#202326]/25">
              <div class="flex items-center justify-between border-b border-[#202326]/20 bg-[#202326]/[0.06] px-2.5 py-1.5 text-[8px] font-black uppercase tracking-[0.14em] text-[#555b5f]">
                <span>Service contacts / 备用</span><span>Line</span>
              </div>
              <a
                v-for="(c, i) in contacts.slice(1)"
                :key="i"
                :href="'tel:' + c.phone"
                class="et-row flex items-center justify-between gap-3 border-b border-dashed border-[#202326]/20 px-2.5 py-2.5 last:border-b-0"
              >
                <div class="flex min-w-0 items-center gap-2.5">
                  <span class="text-[9px] font-black tabular-nums text-[#8a9093]">{{ String(i + 2).padStart(2, '0') }}</span>
                  <div class="min-w-0">
                    <p class="truncate text-[19px] font-black leading-tight tabular-nums">{{ c.phone }}</p>
                    <p class="truncate text-[9px] text-[#63696d]">{{ c.name || '联系人' }}<span v-if="c.relation"> · {{ c.relation }}</span></p>
                  </div>
                </div>
                <Phone class="h-3.5 w-3.5 shrink-0 text-[#575d61]" />
              </a>
            </div>
          </template>
          <div v-else class="mt-2 border-2 border-dashed border-[#202326]/30 px-4 py-7 text-center text-[11px] font-bold text-[#63696d]">NO CONTACT REGISTERED / 未登记</div>
        </section>

        <section class="relative mt-4 border-t-2 border-[#202326]/50 pt-3">
          <div class="flex items-center justify-between">
            <p class="text-[8px] font-black uppercase tracking-[0.2em] text-[#555b5f]">Inspection checklist · 医疗项</p>
            <HeartPulse class="h-3.5 w-3.5 text-[#b4482b]" />
          </div>
          <div class="mt-2 grid grid-cols-2 gap-x-4 gap-y-1.5">
            <div v-for="m in medical" :key="m.label" class="flex items-start gap-1.5 border-b border-[#202326]/12 pb-1.5">
              <span class="mt-0.5 flex h-3 w-3 shrink-0 items-center justify-center border border-[#202326]/45 text-[7px] font-black" :class="m.value !== '未配置' && m.value !== '—' ? 'bg-[#202326] text-white' : 'text-transparent'">✓</span>
              <div class="min-w-0">
                <p class="text-[8px] font-bold uppercase tracking-[0.08em] text-[#777d80]">{{ m.label }}</p>
                <p class="mt-0.5 text-[10px] font-semibold leading-snug text-[#282c2f]">{{ m.value }}</p>
              </div>
            </div>
          </div>
        </section>

        <div v-if="hasPet" class="relative mt-3 flex items-center justify-between border-t border-[#202326]/20 pt-2.5 text-[10px]">
          <span class="font-semibold text-[#444a4e]">{{ payload.petName }}<span v-if="payload.breed" class="text-[#737a7e]"> / {{ payload.breed }}</span></span>
          <a v-if="payload.vetPhone" :href="'tel:' + payload.vetPhone" class="font-black text-[#b4482b]">{{ payload.vetPhone }}</a>
        </div>

        <div class="relative mt-4 flex items-end justify-between gap-3 border-t-2 border-[#202326]/50 pt-3">
          <div class="et-barcode h-7 flex-1 opacity-60" aria-hidden="true"></div>
          <p class="text-right font-mono text-[8px] leading-tight tracking-[0.08em] text-[#63696d]">EC-{{ (nfcId || '0000').toString().slice(-8) }}<br />EMERGENCY USE</p>
        </div>
      </article>

      <p class="et-enter et-d3 mt-4 text-center text-[8px] font-bold uppercase tracking-[0.2em] text-[#858c90]">Fasten to equipment bag · 随包携带</p>
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
  { label: '过敏', value: props.payload?.allergies },
  { label: '用药', value: props.payload?.medications },
  { label: '病史', value: props.payload?.conditions },
])
const hasPet = computed(() => !!(props.payload?.petName || props.payload?.chipId))
</script>

<style scoped>
.et-root {
  background:
    radial-gradient(rgba(255, 255, 255, 0.035) 0.8px, transparent 0.8px),
    #2a2d30;
  background-size: 4px 4px;
}
.et-plate {
  background:
    repeating-linear-gradient(96deg, rgba(255, 255, 255, 0.14) 0 1px, rgba(0, 0, 0, 0.025) 1px 3px),
    linear-gradient(145deg, #d6d9d7, #b8bdbd 48%, #cdd1cf);
  border: 1px solid #eef0ee;
  box-shadow: 0 26px 60px -34px #000, inset 0 1px 0 #fff, inset 0 -2px 4px rgba(0, 0, 0, 0.18);
}
.et-rivet {
  position: absolute;
  z-index: 2;
  width: 9px;
  height: 9px;
  border-radius: 999px;
  background: radial-gradient(circle at 35% 30%, #f4f5f3, #8f9596 55%, #555b5e);
  box-shadow: inset 0 -1px 1px rgba(0, 0, 0, 0.4), 0 1px 1px rgba(255, 255, 255, 0.5);
}
.et-rivet-tl { left: 9px; top: 9px; }
.et-rivet-tr { right: 9px; top: 9px; }
.et-rivet-bl { left: 9px; bottom: 9px; }
.et-rivet-br { right: 9px; bottom: 9px; }
.et-primary {
  background:
    repeating-linear-gradient(90deg, rgba(255, 255, 255, 0.025) 0 1px, transparent 1px 4px),
    #202326;
  border-radius: 4px;
  box-shadow: inset 0 2px 10px rgba(0, 0, 0, 0.55), 0 1px 0 rgba(255, 255, 255, 0.4);
  transition: background-color 0.2s ease, transform 0.18s ease;
}
.et-primary:hover { background-color: #2b2f32; transform: translateY(-1px); }
.et-primary:active { transform: scale(0.99); }
.et-row { transition: background-color 0.18s ease; }
.et-row:hover { background: rgba(32, 35, 38, 0.055); }
.et-barcode {
  background: repeating-linear-gradient(90deg, #202326 0 2px, transparent 2px 4px, #202326 4px 5px, transparent 5px 9px, #202326 9px 12px, transparent 12px 14px);
}
.et-enter { animation: et-clamp 0.7s cubic-bezier(0.22, 1, 0.36, 1) both; }
.et-d1 { animation-delay: 0.08s; }
.et-d3 { animation-delay: 0.24s; }
@keyframes et-clamp {
  from { opacity: 0; transform: translateY(-12px) scale(0.985); }
  to { opacity: 1; transform: translateY(0) scale(1); }
}
@media (prefers-reduced-motion: reduce) {
  .et-enter { animation: none; }
  .et-primary, .et-row { transition: none; }
}
</style>
