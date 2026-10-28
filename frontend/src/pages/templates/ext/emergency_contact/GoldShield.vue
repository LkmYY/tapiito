<template>
  <div class="gs-root min-h-screen text-[#f4ead2]">
    <main class="mx-auto w-full max-w-md px-5 pb-12 pt-7">
      <div class="gs-shell gs-enter relative overflow-hidden rounded-[26px] border border-[#b9974d]/50 px-5 pb-7 pt-6">
        <div class="gs-rays" aria-hidden="true"></div>
        <div class="gs-frame" aria-hidden="true"></div>

        <div class="relative flex items-center justify-between text-[9px] uppercase tracking-[0.26em] text-[#c9b27c]">
          <span>Guardian Registry</span>
          <span v-if="nfcId" class="max-w-[42%] truncate">NFC · {{ nfcId }}</span>
        </div>

        <header class="relative mt-8 text-center">
          <div class="gs-crest mx-auto flex h-16 w-16 items-center justify-center rounded-full border border-[#d8b768]/70">
            <Shield class="h-7 w-7 text-[#e3c47b]" />
          </div>
          <p class="mt-5 text-[10px] uppercase tracking-[0.34em] text-[#b9974d]">In Case of Emergency</p>
          <h1 class="gs-serif mt-2 text-[34px] font-semibold leading-tight tracking-wide text-[#fbf3df]">紧急联络</h1>
          <div class="mx-auto mt-3 flex w-32 items-center gap-2">
            <span class="h-px flex-1 bg-[#b9974d]/60"></span>
            <span class="h-1 w-1 rotate-45 bg-[#d8b768]"></span>
            <span class="h-px flex-1 bg-[#b9974d]/60"></span>
          </div>
          <p class="mx-auto mt-4 max-w-[260px] text-[12px] leading-relaxed text-[#d8c8a6]/70">请依下列顺序联络家属，并向救援人员出示本页医疗信息。</p>
        </header>

        <section class="relative mt-7">
          <p class="text-center text-[9px] uppercase tracking-[0.28em] text-[#b9974d]">Primary Contact</p>
          <template v-if="contacts.length">
            <a :href="'tel:' + contacts[0].phone" class="gs-call gs-enter gs-d1 mt-3 block rounded-xl border border-[#d8b768]/45 bg-[#e3c47b]/[0.06] px-4 py-5 text-center">
              <p class="break-words text-[36px] font-semibold leading-none tabular-nums tracking-[0.03em] text-[#fbf3df]">{{ contacts[0].phone }}</p>
              <p class="mt-3 text-[12px] text-[#d8c8a6]">{{ contacts[0].name || '联系人' }}<span v-if="contacts[0].relation" class="text-[#b9974d]"> · {{ contacts[0].relation }}</span></p>
              <p class="mt-3 text-[9px] uppercase tracking-[0.24em] text-[#e3c47b]">Tap to call</p>
            </a>

            <div v-if="contacts.length > 1" class="mt-3 space-y-2">
              <a
                v-for="(c, i) in contacts.slice(1)"
                :key="i"
                :href="'tel:' + c.phone"
                class="gs-row flex items-center justify-between gap-3 rounded-lg border border-[#b9974d]/25 px-4 py-3"
              >
                <div class="min-w-0">
                  <p class="truncate text-[20px] font-medium leading-tight tabular-nums text-[#f4ead2]">{{ c.phone }}</p>
                  <p class="mt-0.5 truncate text-[10px] uppercase tracking-[0.12em] text-[#b9974d]">{{ c.name || '联系人' }}<span v-if="c.relation"> · {{ c.relation }}</span></p>
                </div>
                <Phone class="h-4 w-4 shrink-0 text-[#d8b768]" />
              </a>
            </div>
          </template>
          <div v-else class="mt-3 rounded-xl border border-dashed border-[#b9974d]/40 px-4 py-7 text-center text-[12px] text-[#c9b27c]/70">暂未登记紧急联系人</div>
        </section>

        <section class="relative mt-6 border-t border-[#b9974d]/30 pt-5">
          <div class="flex items-center justify-between">
            <p class="text-[9px] uppercase tracking-[0.24em] text-[#b9974d]">Medical Record</p>
            <HeartPulse class="h-4 w-4 text-[#d8b768]" />
          </div>
          <div class="mt-4 grid grid-cols-2 gap-x-5 gap-y-4">
            <div v-for="m in medical" :key="m.label">
              <p class="text-[9px] uppercase tracking-[0.16em] text-[#967c4b]">{{ m.label }}</p>
              <p class="mt-1 text-[12px] leading-snug text-[#e8dcc2]">{{ m.value }}</p>
            </div>
          </div>
        </section>

        <div v-if="hasPet" class="relative mt-5 border-t border-[#b9974d]/30 pt-4 text-center text-[11px] text-[#d8c8a6]/80">
          {{ payload.petName }}<span v-if="payload.breed"> · {{ payload.breed }}</span>
          <a v-if="payload.vetPhone" :href="'tel:' + payload.vetPhone" class="ml-2 text-[#e3c47b] underline decoration-[#e3c47b]/40 underline-offset-4">宠物医院 {{ payload.vetPhone }}</a>
        </div>

        <p class="relative mt-6 text-center text-[9px] uppercase tracking-[0.28em] text-[#967c4b]">Registered · Keep this page visible</p>
      </div>
    </main>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { Shield, Phone, HeartPulse } from 'lucide-vue-next'

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
.gs-root {
  background:
    radial-gradient(90% 45% at 50% 0%, rgba(185, 151, 77, 0.16), transparent 62%),
    #0f0c08;
}
.gs-shell {
  background:
    radial-gradient(80% 40% at 50% 0%, rgba(227, 196, 123, 0.09), transparent 60%),
    linear-gradient(170deg, #171208, #0d0a06 70%);
  box-shadow: 0 30px 80px -40px rgba(0, 0, 0, 0.95), inset 0 0 0 3px rgba(15, 12, 8, 0.9);
}
.gs-frame {
  position: absolute;
  inset: 8px;
  border: 1px solid rgba(185, 151, 77, 0.26);
  border-radius: 20px;
  pointer-events: none;
}
.gs-frame::before,
.gs-frame::after {
  content: "";
  position: absolute;
  width: 22px;
  height: 22px;
  border: 1px solid rgba(216, 183, 104, 0.7);
}
.gs-frame::before { left: -1px; top: -1px; border-right: 0; border-bottom: 0; }
.gs-frame::after { right: -1px; bottom: -1px; border-left: 0; border-top: 0; }
.gs-rays {
  position: absolute;
  left: 50%;
  top: -180px;
  width: 380px;
  height: 380px;
  transform: translateX(-50%);
  background: repeating-conic-gradient(from 0deg, rgba(216, 183, 104, 0.16) 0deg 1deg, transparent 1deg 12deg);
  -webkit-mask-image: radial-gradient(circle, #000 0 40%, transparent 68%);
  mask-image: radial-gradient(circle, #000 0 40%, transparent 68%);
  opacity: 0.55;
}
.gs-crest { background: radial-gradient(circle, rgba(227, 196, 123, 0.15), transparent 70%); box-shadow: 0 0 30px -8px rgba(227, 196, 123, 0.35); }
.gs-serif { font-family: "Iowan Old Style", "Palatino Linotype", Georgia, "Songti SC", serif; }
.gs-enter { animation: gs-in 0.8s cubic-bezier(0.22, 1, 0.36, 1) both; }
.gs-d1 { animation-delay: 0.1s; }
@keyframes gs-in {
  from { opacity: 0; transform: translateY(18px) scale(0.985); }
  to { opacity: 1; transform: translateY(0) scale(1); }
}
.gs-call { transition: background-color 0.25s ease, border-color 0.25s ease, transform 0.25s ease; }
.gs-call:hover { background: rgba(227, 196, 123, 0.12); border-color: rgba(227, 196, 123, 0.8); transform: translateY(-2px); }
.gs-call:active { transform: scale(0.985); }
.gs-row { transition: border-color 0.2s ease, background-color 0.2s ease; }
.gs-row:hover { border-color: rgba(216, 183, 104, 0.6); background: rgba(227, 196, 123, 0.05); }
@media (prefers-reduced-motion: reduce) {
  .gs-enter { animation: none; }
  .gs-call, .gs-row { transition: none; }
}
</style>
