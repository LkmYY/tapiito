<template>
  <div class="tc-root min-h-screen text-[#252a28]">
    <main class="mx-auto w-full max-w-md px-5 pb-12 pt-6">
      <div class="tc-enter flex items-center justify-between text-[9px] font-bold uppercase tracking-[0.22em] text-[#77817d]">
        <span class="flex items-center gap-2"><span class="h-2 w-2 rounded-full bg-[#c0392b]"></span>Triage Contact Card</span>
        <span v-if="nfcId" class="max-w-[42%] truncate">NFC · {{ nfcId }}</span>
      </div>

      <header class="tc-enter tc-d1 relative mt-5">
        <div class="tc-string" aria-hidden="true"></div>
        <div class="tc-tag relative -rotate-[0.7deg] border border-[#c9c4b8] bg-white px-5 pb-5 pt-7">
          <div class="tc-hole absolute left-1/2 top-3 h-4 w-4 -translate-x-1/2 rounded-full border border-[#b7b2a6] bg-[#f4f1ea]"></div>
          <div class="flex items-start justify-between gap-3 border-b-2 border-[#252a28] pb-3">
            <div>
              <p class="text-[9px] font-bold uppercase tracking-[0.18em] text-[#c0392b]">Priority Red · 一级</p>
              <h1 class="mt-1 text-[31px] font-black leading-tight tracking-tight">紧急联系卡</h1>
            </div>
            <span class="tc-level flex h-12 w-12 shrink-0 flex-col items-center justify-center rounded-full bg-[#c0392b] text-center text-[8px] font-black uppercase leading-tight text-white">CALL<br />NOW</span>
          </div>
          <p class="mt-3 text-[11px] leading-relaxed text-[#626b67]">请按分诊顺序拨打。先确认现场安全，再向对方说明位置、意识与呼吸状况。</p>
          <div class="mt-3 flex items-center justify-between text-[9px] font-bold uppercase tracking-[0.14em] text-[#77817d]">
            <span>Patient unknown · 请协助</span><span>Form 03</span>
          </div>
        </div>
      </header>

      <section class="mt-5">
        <div class="tc-enter tc-d2 flex items-center justify-between border-b border-[#b7b2a6] pb-1.5">
          <p class="text-[9px] font-bold uppercase tracking-[0.2em] text-[#77817d]">Contact priority · 联络优先级</p>
          <span class="text-[9px] tabular-nums text-[#77817d]">{{ String(contacts.length).padStart(2, '0') }}</span>
        </div>

        <template v-if="contacts.length">
          <a :href="'tel:' + contacts[0].phone" class="tc-primary tc-enter tc-d2 mt-3 block bg-white px-5 py-4">
            <div class="flex items-center justify-between">
              <span class="rounded-sm bg-[#c0392b] px-2 py-1 text-[9px] font-black uppercase tracking-[0.12em] text-white">立即 · 01</span>
              <Phone class="h-4 w-4 text-[#c0392b]" />
            </div>
            <p class="mt-3 break-words text-[34px] font-black leading-none tabular-nums tracking-tight">{{ contacts[0].phone }}</p>
            <div class="mt-3 flex items-center justify-between border-t border-[#dcd7cc] pt-2.5">
              <p class="text-[12px] font-semibold">{{ contacts[0].name || '联系人' }}<span v-if="contacts[0].relation" class="font-normal text-[#77817d]"> · {{ contacts[0].relation }}</span></p>
              <span class="text-[9px] font-bold uppercase tracking-[0.12em] text-[#c0392b]">Tap to call</span>
            </div>
          </a>

          <div class="mt-2 space-y-2">
            <a
              v-for="(c, i) in contacts.slice(1)"
              :key="i"
              :href="'tel:' + c.phone"
              class="tc-row flex items-stretch bg-white"
              :class="i % 2 === 0 ? 'tc-amber' : 'tc-green'"
            >
              <div class="flex w-7 shrink-0 flex-col items-center justify-center border-r border-[#e6e1d6]">
                <span class="text-[8px] font-black uppercase tracking-[0.08em] text-[#77817d]">P{{ i + 2 }}</span>
              </div>
              <div class="min-w-0 flex-1 px-4 py-3.5">
                <p class="truncate text-[21px] font-bold leading-tight tabular-nums">{{ c.phone }}</p>
                <p class="mt-0.5 truncate text-[10px] text-[#77817d]">{{ c.name || '联系人' }}<span v-if="c.relation"> · {{ c.relation }}</span></p>
              </div>
              <div class="flex items-center pr-4">
                <Phone class="h-4 w-4 text-[#4e5a55]" />
              </div>
            </a>
          </div>
        </template>
        <div v-else class="mt-3 border border-dashed border-[#b7b2a6] bg-white/60 px-4 py-7 text-center text-[12px] text-[#77817d]">未登记联系人 · 请补充</div>
      </section>

      <section class="tc-enter tc-d3 mt-5 bg-white px-4 py-4">
        <div class="flex items-center justify-between border-b border-[#dcd7cc] pb-2">
          <p class="text-[9px] font-bold uppercase tracking-[0.2em] text-[#77817d]">Vitals &amp; notes · 医疗信息</p>
          <HeartPulse class="h-4 w-4 text-[#c0392b]" />
        </div>
        <div class="mt-3 grid grid-cols-2 gap-x-4 gap-y-3">
          <div v-for="m in medical" :key="m.label">
            <p class="text-[9px] font-bold uppercase tracking-[0.1em] text-[#98a19c]">{{ m.label }}</p>
            <p class="mt-1 text-[12px] font-semibold leading-snug">{{ m.value }}</p>
          </div>
        </div>
        <div class="mt-3 flex items-center gap-2 border-t border-dashed border-[#dcd7cc] pt-2.5 text-[9px] text-[#8b948f]">
          <Activity class="h-3.5 w-3.5" /> 若信息空白，请以现场急救人员判断为准。
        </div>
      </section>

      <div v-if="hasPet" class="tc-enter tc-d3 mt-3 flex items-center justify-between border-l-4 border-[#1e8449] bg-white px-4 py-2.5 text-[11px]">
        <span>{{ payload.petName }}<span v-if="payload.breed" class="text-[#77817d]"> · {{ payload.breed }}</span></span>
        <a v-if="payload.vetPhone" :href="'tel:' + payload.vetPhone" class="font-bold text-[#1e8449]">{{ payload.vetPhone }}</a>
      </div>

      <p class="tc-enter tc-d3 mt-5 text-center text-[8px] font-bold uppercase tracking-[0.22em] text-[#98a19c]">Triage card · Keep with patient</p>
    </main>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { Phone, HeartPulse, Activity } from 'lucide-vue-next'

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
.tc-root {
  background:
    linear-gradient(rgba(119, 129, 125, 0.055) 1px, transparent 1px),
    #f2efe7;
  background-size: 100% 28px;
}
.tc-tag {
  clip-path: polygon(0 0, 100% 0, 100% calc(100% - 20px), calc(100% - 20px) 100%, 0 100%);
  box-shadow: 0 18px 38px -30px rgba(37, 42, 40, 0.55);
}
.tc-string {
  position: absolute;
  left: 50%;
  top: -14px;
  width: 1px;
  height: 16px;
  background: #a9a397;
  transform: translateX(-50%);
}
.tc-hole { box-shadow: inset 0 2px 3px rgba(80, 70, 55, 0.18); }
.tc-primary {
  border-left: 4px solid #c0392b;
  box-shadow: 0 16px 32px -28px rgba(37, 42, 40, 0.55);
  transition: transform 0.2s ease, box-shadow 0.2s ease;
}
.tc-primary:hover { transform: translateY(-2px); box-shadow: 0 22px 38px -26px rgba(37, 42, 40, 0.6); }
.tc-primary:active { transform: scale(0.99); }
.tc-row {
  border-left: 4px solid transparent;
  box-shadow: 0 10px 24px -24px rgba(37, 42, 40, 0.45);
  transition: background-color 0.18s ease, transform 0.18s ease;
}
.tc-amber { border-left-color: #d68910; }
.tc-green { border-left-color: #1e8449; }
.tc-row:hover { background: #fbf8f0; transform: translateX(2px); }
.tc-row:active { background: #f4efe2; }
.tc-level { box-shadow: 0 0 0 4px rgba(192, 57, 43, 0.12); }
.tc-enter { animation: tc-drop 0.7s cubic-bezier(0.22, 1, 0.36, 1) both; }
.tc-d1 { animation-delay: 0.08s; }
.tc-d2 { animation-delay: 0.16s; }
.tc-d3 { animation-delay: 0.24s; }
@keyframes tc-drop {
  from { opacity: 0; transform: translateY(-10px) rotate(-1deg); }
  to { opacity: 1; transform: translateY(0) rotate(0); }
}
@media (prefers-reduced-motion: reduce) {
  .tc-enter { animation: none; }
  .tc-primary, .tc-row { transition: none; }
}
</style>
