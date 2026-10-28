<template>
  <div class="su-root min-h-screen text-[#211e19]">
    <main class="su-wrap relative mx-auto w-full max-w-md pb-12 pl-14 pr-5 pt-7">
      <div class="su-rail" aria-hidden="true">
        <span class="su-vertical">紧急連絡先</span>
        <span class="su-rail-line"></span>
        <span class="su-rail-no">零一</span>
      </div>

      <div class="su-enter flex items-center justify-between text-[9px] uppercase tracking-[0.22em] text-[#8a8173]">
        <span>Sumi Contact Note</span>
        <span v-if="nfcId" class="max-w-[42%] truncate">NFC · {{ nfcId }}</span>
      </div>

      <header class="su-enter su-d1 relative mt-8">
        <div class="su-enso" aria-hidden="true"></div>
        <p class="text-[11px] tracking-[0.2em] text-[#a5342a]">如遇紧急情况</p>
        <h1 class="su-serif mt-2 text-[36px] font-semibold leading-[1.12] tracking-wide">请与我家人<br />取得联络</h1>
        <p class="su-serif mt-4 max-w-[255px] text-[12px] leading-6 text-[#716858]">拨通后，请告知持卡人所在位置。若对方未接听，请继续拨打下一位。</p>
        <div class="mt-5 h-px w-full bg-gradient-to-r from-[#211e19] via-[#211e19]/35 to-transparent"></div>
      </header>

      <section class="mt-6">
        <div class="flex items-center justify-between">
          <p class="text-[9px] uppercase tracking-[0.22em] text-[#8a8173]">Call in order · 依次拨打</p>
          <span class="text-[9px] tabular-nums text-[#8a8173]">{{ String(contacts.length).padStart(2, '0') }}</span>
        </div>

        <template v-if="contacts.length">
          <a
            :href="'tel:' + contacts[0].phone"
            class="su-call su-enter su-d2 mt-3 block border-b border-[#211e19]/20 pb-5 pt-2"
          >
            <div class="flex items-center gap-2 text-[11px] text-[#716858]">
              <span class="su-serif text-[15px] italic text-[#a5342a]">一</span>
              <span>{{ contacts[0].name || '联系人' }}</span>
              <span v-if="contacts[0].relation">· {{ contacts[0].relation }}</span>
            </div>
            <p class="su-serif mt-2 break-words text-[38px] font-semibold leading-tight tabular-nums tracking-wide text-[#15130f]">{{ contacts[0].phone }}</p>
            <div class="mt-2 flex items-center justify-between">
              <span class="text-[9px] uppercase tracking-[0.2em] text-[#a5342a]">首要联络 · Tap to call</span>
              <Phone class="h-4 w-4 text-[#a5342a]" />
            </div>
          </a>

          <a
            v-for="(c, i) in contacts.slice(1)"
            :key="i"
            :href="'tel:' + c.phone"
            class="su-row su-enter su-d3 flex items-center justify-between gap-4 border-b border-[#211e19]/15 py-4"
          >
            <div class="min-w-0">
              <p class="truncate text-[22px] font-semibold leading-tight tabular-nums tracking-wide">{{ c.phone }}</p>
              <p class="mt-1 truncate text-[11px] text-[#716858]">{{ c.name || '联系人' }}<span v-if="c.relation"> · {{ c.relation }}</span></p>
            </div>
            <span class="su-serif text-[15px] text-[#a5342a]">{{ ['二', '三', '四', '五'][i] || '辅' }}</span>
          </a>
        </template>
        <div v-else class="mt-3 border-y border-dashed border-[#bdb3a2] py-8 text-center text-[12px] text-[#716858]">暂无紧急联络人</div>
      </section>

      <section class="su-enter su-d3 relative mt-7">
        <div class="flex items-center gap-3">
          <span class="h-px flex-1 bg-[#211e19]/20"></span>
          <p class="text-[9px] uppercase tracking-[0.24em] text-[#8a8173]">Medical Memo</p>
          <span class="h-px flex-1 bg-[#211e19]/20"></span>
        </div>
        <div class="mt-4 space-y-2.5">
          <div v-for="m in medical" :key="m.label" class="flex items-start gap-4">
            <span class="w-9 flex-shrink-0 text-[10px] tracking-[0.14em] text-[#8a8173]">{{ m.label }}</span>
            <span class="su-serif flex-1 border-b border-[#211e19]/10 pb-1.5 text-[12px] leading-snug">{{ m.value }}</span>
          </div>
        </div>
        <div class="su-seal mt-5 ml-auto flex h-12 w-12 -rotate-6 items-center justify-center border border-[#a5342a] text-center text-[9px] leading-tight text-[#a5342a]">随身<br />携带</div>
      </section>

      <div v-if="hasPet" class="su-enter su-d3 mt-5 border-l-2 border-[#a5342a] pl-3 text-[11px] text-[#5f574a]">
        同行宠物：{{ payload.petName }}<span v-if="payload.breed"> · {{ payload.breed }}</span>
        <a v-if="payload.vetPhone" :href="'tel:' + payload.vetPhone" class="ml-1 underline decoration-[#a5342a]/40 underline-offset-4">{{ payload.vetPhone }}</a>
      </div>

      <p class="su-enter su-d3 mt-7 text-[9px] uppercase tracking-[0.26em] text-[#a59b8c]">Please stay until help arrives</p>
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
const medical = computed(() => [
  { label: '血型', value: props.payload?.bloodType || '—' },
  { label: '过敏', value: props.payload?.allergies },
  { label: '用药', value: props.payload?.medications },
  { label: '病史', value: props.payload?.conditions },
])
const hasPet = computed(() => !!(props.payload?.petName || props.payload?.chipId))
</script>

<style scoped>
.su-root {
  background:
    radial-gradient(70% 50% at 92% 12%, rgba(165, 52, 42, 0.07), transparent 60%),
    linear-gradient(170deg, #f6f2e9, #ece5d8);
}
.su-serif { font-family: "Songti SC", "Noto Serif SC", "Iowan Old Style", Georgia, serif; }
.su-wrap::before {
  content: "";
  position: absolute;
  inset: 0 0 0 45px;
  border-right: 1px solid rgba(33, 30, 25, 0.08);
  pointer-events: none;
}
.su-rail {
  position: absolute;
  left: 14px;
  top: 28px;
  bottom: 28px;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 12px;
  color: #8a8173;
}
.su-vertical {
  writing-mode: vertical-rl;
  font-family: "Songti SC", "Noto Serif SC", serif;
  font-size: 11px;
  letter-spacing: 0.35em;
}
.su-rail-line { flex: 1; width: 1px; background: linear-gradient(#211e19, rgba(33, 30, 25, 0.1)); }
.su-rail-no { writing-mode: vertical-rl; font-family: "Songti SC", serif; font-size: 9px; letter-spacing: 0.2em; color: #a5342a; }
.su-enso {
  position: absolute;
  right: -10px;
  top: -20px;
  width: 110px;
  height: 110px;
  border-radius: 999px;
  border: 9px solid rgba(33, 30, 25, 0.05);
  border-right-color: transparent;
  transform: rotate(-24deg);
}
.su-call { transition: border-color 0.2s ease; }
.su-call:hover { border-color: rgba(165, 52, 42, 0.55); }
.su-call:hover p { color: #a5342a; }
.su-call:active { opacity: 0.75; }
.su-call p { transition: color 0.2s ease; }
.su-row { transition: border-color 0.2s ease, background-color 0.2s ease, padding-left 0.2s ease; }
.su-row:hover { border-color: rgba(165, 52, 42, 0.5); background: rgba(165, 52, 42, 0.035); padding-left: 4px; }
.su-seal { background: rgba(165, 52, 42, 0.04); }
.su-enter { animation: su-ink 0.9s cubic-bezier(0.22, 1, 0.36, 1) both; }
.su-d1 { animation-delay: 0.1s; }
.su-d2 { animation-delay: 0.2s; }
.su-d3 { animation-delay: 0.3s; }
@keyframes su-ink {
  from { opacity: 0; clip-path: inset(0 0 20% 0); transform: translateY(8px); filter: blur(1.5px); }
  to { opacity: 1; clip-path: inset(0 0 0 0); transform: translateY(0); filter: blur(0); }
}
@media (prefers-reduced-motion: reduce) {
  .su-enter { animation: none; }
  .su-call, .su-call p, .su-row { transition: none; }
}
</style>
