<template>
  <div class="px-root min-h-screen text-[#e7ffd4]">
    <main class="mx-auto w-full max-w-md px-5 pb-12 pt-6 font-mono">
      <div class="px-enter flex items-center justify-between text-[10px] uppercase tracking-[0.12em] text-[#7ba85a]">
        <span class="flex items-center gap-1.5"><span class="px-heart"></span>PLAYER · NEED HELP</span>
        <span v-if="nfcId" class="max-w-[42%] truncate">ID:{{ nfcId }}</span>
      </div>

      <header class="px-enter px-d1 mt-5 border-[3px] border-[#a3e635] bg-[#10180c] px-5 py-5 shadow-[6px_6px_0_#2f4a1f]">
        <div class="flex items-center justify-between text-[9px] uppercase tracking-[0.18em] text-[#a3e635]">
          <span>EMERGENCY QUEST</span><span>[ LV.01 ]</span>
        </div>
        <h1 class="px-title mt-3 text-[32px] font-black leading-none tracking-tight text-[#f3ffe8]">紧急联络</h1>
        <p class="mt-3 text-[11px] leading-5 text-[#9bc47d]">&gt; 请选择联络对象<br />&gt; 报出位置，等待支援</p>
        <div class="mt-4 flex items-center gap-2 border-t-2 border-dashed border-[#3b5a28] pt-3 text-[9px] uppercase tracking-[0.14em] text-[#6f9a52]">
          <span class="px-key">A</span> CALL <span class="px-key">B</span> BACK
        </div>
      </header>

      <section class="mt-5">
        <p class="px-enter px-d2 text-[10px] uppercase tracking-[0.16em] text-[#6f9a52]">► ACTIVE CONTACTS / {{ String(contacts.length).padStart(2, '0') }}</p>

        <template v-if="contacts.length">
          <a
            :href="'tel:' + contacts[0].phone"
            class="px-primary px-enter px-d2 mt-2 block border-[3px] border-[#a3e635] bg-[#16220f] px-5 py-5 shadow-[5px_5px_0_#000]"
          >
            <div class="flex items-start justify-between text-[9px] uppercase tracking-[0.16em] text-[#a3e635]">
              <span>▶ 01 · PRIMARY</span><Phone class="h-4 w-4" />
            </div>
            <p class="mt-3 break-words text-[35px] font-black leading-none tabular-nums tracking-tight text-white">{{ contacts[0].phone }}</p>
            <div class="mt-4 flex items-center justify-between border-t-2 border-[#2f4a1f] pt-3">
              <span class="text-[11px] font-bold">{{ contacts[0].name || '联系人' }}<span v-if="contacts[0].relation" class="text-[#8db972]"> / {{ contacts[0].relation }}</span></span>
              <span class="px-btn bg-[#a3e635] px-3 py-1.5 text-[9px] font-black text-[#10180c]">CALL</span>
            </div>
          </a>

          <div class="mt-3 grid gap-2">
            <a
              v-for="(c, i) in contacts.slice(1)"
              :key="i"
              :href="'tel:' + c.phone"
              class="px-row px-4 py-3.5"
            >
              <div class="flex items-center gap-3">
                <span class="flex h-7 w-7 shrink-0 items-center justify-center border-2 border-[#5a8a3c] text-[9px] text-[#a3e635]">{{ String(i + 2).padStart(2, '0') }}</span>
                <div class="min-w-0 flex-1">
                  <p class="truncate text-[20px] font-bold leading-tight tabular-nums text-white">{{ c.phone }}</p>
                  <p class="mt-0.5 truncate text-[10px] text-[#7ba85a]">{{ c.name || '联系人' }}<span v-if="c.relation"> · {{ c.relation }}</span></p>
                </div>
                <span class="text-[9px] text-[#a3e635]">▶</span>
              </div>
            </a>
          </div>
        </template>
        <div v-else class="mt-2 border-[3px] border-dashed border-[#3b5a28] px-4 py-8 text-center text-[11px] text-[#7ba85a]">&gt; NO CONTACTS FOUND<br />请补充紧急号码</div>
      </section>

      <section class="px-enter px-d3 mt-5 border-[3px] border-[#3b5a28] bg-[#0d150a] p-4">
        <div class="flex items-center justify-between text-[10px] uppercase tracking-[0.16em] text-[#a3e635]">
          <span>► STATS · 医疗</span><Heart class="h-4 w-4" />
        </div>
        <div class="mt-3 grid grid-cols-2 gap-2">
          <div v-for="m in medical" :key="m.label" class="border-2 border-[#263d1b] bg-[#131e0d] px-3 py-2.5">
            <p class="text-[9px] uppercase tracking-[0.1em] text-[#6f9a52]">{{ m.label }}</p>
            <p class="mt-1 text-[12px] font-bold leading-snug text-[#dcffc2]">{{ m.value }}</p>
          </div>
        </div>
      </section>

      <div v-if="hasPet" class="px-enter px-d3 mt-3 border-2 border-[#6f8f2e] bg-[#1a230c] px-4 py-3 text-[11px] text-[#d7f5a8]">
        <span class="text-[#93b35d]">PET:</span> {{ payload.petName }}<span v-if="payload.breed"> / {{ payload.breed }}</span>
        <a v-if="payload.vetPhone" :href="'tel:' + payload.vetPhone" class="ml-2 text-[#d9f99d] underline underline-offset-2">VET</a>
      </div>

      <p class="px-enter px-d3 mt-6 text-center text-[9px] uppercase tracking-[0.2em] text-[#4f7238]">— SAVE &amp; CALL —</p>
    </main>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { Phone, Heart } from 'lucide-vue-next'

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
.px-root {
  background:
    linear-gradient(rgba(163, 230, 53, 0.045) 1px, transparent 1px),
    linear-gradient(90deg, rgba(163, 230, 53, 0.045) 1px, transparent 1px),
    #080d06;
  background-size: 22px 22px, 22px 22px, 100% 100%;
}
.px-title {
  font-family: "Courier New", monospace;
  text-shadow: 3px 3px 0 #3f6212;
}
.px-heart {
  width: 8px;
  height: 8px;
  background: #a3e635;
  box-shadow: 8px 0 0 #a3e635, 0 8px 0 #a3e635, 8px 8px 0 #a3e635;
  margin-right: 8px;
}
.px-key {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  width: 16px;
  height: 16px;
  border: 1px solid #6f9a52;
  color: #d9f99d;
  font-size: 8px;
}
.px-primary {
  transition: transform 0.12s steps(2, end), background-color 0.2s ease, box-shadow 0.12s steps(2, end);
}
.px-primary:hover { background: #1d2c13; box-shadow: 8px 8px 0 #000; transform: translate(-2px, -2px); }
.px-primary:active { transform: translate(2px, 2px); box-shadow: 2px 2px 0 #000; }
.px-btn { box-shadow: 3px 3px 0 #000; }
.px-row {
  border: 2px solid #2f4a1f;
  background: #0d150a;
  transition: border-color 0.2s ease, background-color 0.2s ease, transform 0.12s steps(2, end);
}
.px-row:hover { border-color: #a3e635; background: #15200d; transform: translateX(3px); }
.px-row:active { transform: translateX(0); }
.px-enter { animation: px-in 0.5s steps(5, end) both; }
.px-d1 { animation-delay: 0.08s; }
.px-d2 { animation-delay: 0.16s; }
.px-d3 { animation-delay: 0.24s; }
@keyframes px-in {
  from { opacity: 0; transform: translateY(12px); }
  to { opacity: 1; transform: translateY(0); }
}
@media (prefers-reduced-motion: reduce) {
  .px-enter { animation: none; }
  .px-primary, .px-row { transition: none; }
}
</style>
