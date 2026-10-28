<template>
  <div class="vh-root min-h-screen text-[#f5e9ff]">
    <main class="mx-auto w-full max-w-md px-5 pb-12 pt-6">
      <div class="vh-enter flex items-center justify-between text-[9px] uppercase tracking-[0.2em] text-[#ff8fc4]">
        <span class="flex items-center gap-2"><Play class="h-3.5 w-3.5 fill-current" />PLAY · SP</span>
        <span class="tabular-nums text-[#8fe8ff]">00:13:24</span>
      </div>

      <header class="vh-enter vh-d1 relative mt-4 border-2 border-[#ff4fa3]/60 bg-[#1f0e2a] px-5 pb-6 pt-5">
        <div class="flex items-center justify-between text-[9px] uppercase tracking-[0.16em] text-[#c79bdd]">
          <span>CH 03 / EMERGENCY</span>
          <span v-if="nfcId" class="max-w-[38%] truncate">{{ nfcId }}</span>
        </div>
        <p class="vh-glitch mt-4 text-[31px] font-black leading-none tracking-[0.02em]">紧急联络</p>
        <p class="mt-2 text-[10px] font-bold uppercase tracking-[0.3em] text-[#ff4fa3]">Please Rewind · 请先拨号</p>
        <p class="mt-4 max-w-[270px] text-[11px] leading-relaxed text-[#c9b2d6]">带着这张卡的人需要帮助。请拨通下方电话，并描述画面中你看到的位置。</p>
        <div class="mt-4 flex items-center gap-3 border-t border-[#ff4fa3]/25 pt-3 text-[8px] uppercase tracking-[0.14em] text-[#8f739c]">
          <span class="h-2 w-2 rounded-full bg-[#ff4fa3] vh-rec"></span><span>REC</span><span>·</span><span>STEREO</span><span>·</span><span class="text-[#8fe8ff]">TRACKING OK</span>
        </div>
      </header>

      <section class="mt-5">
        <template v-if="contacts.length">
          <a :href="'tel:' + contacts[0].phone" class="vh-primary vh-enter vh-d2 block px-5 py-5">
            <div class="flex items-center justify-between text-[9px] font-bold uppercase tracking-[0.18em] text-[#ff8fc4]">
              <span>▶ PRIMARY / 01</span><Phone class="h-4 w-4" />
            </div>
            <p class="mt-3 break-words text-[36px] font-black leading-none tabular-nums tracking-[0.01em] text-white">{{ contacts[0].phone }}</p>
            <div class="mt-4 flex items-center justify-between border-t border-white/15 pt-3">
              <p class="text-[12px] font-semibold">{{ contacts[0].name || '联系人' }}<span v-if="contacts[0].relation" class="text-[#c79bdd]"> · {{ contacts[0].relation }}</span></p>
              <span class="vh-btn border border-[#ff4fa3] bg-[#ff4fa3] px-3 py-1 text-[9px] font-black uppercase tracking-[0.14em] text-[#1a0b22]">CALL</span>
            </div>
          </a>

          <div class="mt-3 space-y-2">
            <a
              v-for="(c, i) in contacts.slice(1)"
              :key="i"
              :href="'tel:' + c.phone"
              class="vh-row flex items-center justify-between gap-3 px-4 py-3.5"
            >
              <div class="flex min-w-0 items-center gap-3">
                <span class="text-[9px] tabular-nums text-[#8fe8ff]">{{ String(i + 2).padStart(2, '0') }}</span>
                <div class="min-w-0">
                  <p class="truncate text-[21px] font-bold leading-tight tabular-nums text-white">{{ c.phone }}</p>
                  <p class="mt-0.5 truncate text-[10px] text-[#b99ac6]">{{ c.name || '联系人' }}<span v-if="c.relation"> · {{ c.relation }}</span></p>
                </div>
              </div>
              <Phone class="h-4 w-4 shrink-0 text-[#ff4fa3]" />
            </a>
          </div>
        </template>
        <div v-else class="vh-primary px-4 py-8 text-center text-[12px] text-[#c9b2d6]">NO SIGNAL · 暂无联系人</div>
      </section>

      <section class="vh-enter vh-d3 mt-4 border border-[#66e0ff]/30 bg-[#160b1f] p-4">
        <div class="flex items-center justify-between text-[9px] font-bold uppercase tracking-[0.2em] text-[#8fe8ff]">
          <span>ON-SCREEN DATA · 医疗</span><Activity class="h-4 w-4" />
        </div>
        <div class="mt-3 grid grid-cols-2 gap-x-4 gap-y-3">
          <div v-for="m in medical" :key="m.label" class="border-l border-[#ff4fa3]/40 pl-2.5">
            <p class="text-[8px] uppercase tracking-[0.12em] text-[#8f739c]">{{ m.label }}</p>
            <p class="mt-1 text-[11px] font-semibold leading-snug text-[#eadcf2]">{{ m.value }}</p>
          </div>
        </div>
      </section>

      <div v-if="hasPet" class="vh-enter vh-d3 mt-3 flex items-center justify-between border border-[#ff4fa3]/30 px-4 py-2.5 text-[11px]">
        <span class="text-[#c9b2d6]">{{ payload.petName }}<span v-if="payload.breed"> · {{ payload.breed }}</span></span>
        <a v-if="payload.vetPhone" :href="'tel:' + payload.vetPhone" class="text-[#8fe8ff]">VET {{ payload.vetPhone }}</a>
      </div>

      <div class="vh-enter vh-d3 mt-5 flex items-center justify-between border-t border-[#ff4fa3]/20 pt-3 text-[8px] uppercase tracking-[0.18em] text-[#8f739c]">
        <span>Emergency Copy · Do Not Erase</span><span class="text-[#ff8fc4]">▶ END</span>
      </div>
    </main>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { Play, Phone, Activity } from 'lucide-vue-next'

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
.vh-root {
  position: relative;
  background:
    radial-gradient(80% 55% at 50% 0%, rgba(255, 79, 163, 0.13), transparent 60%),
    #140b1c;
}
.vh-root::before {
  content: "";
  position: fixed;
  inset: 0;
  pointer-events: none;
  z-index: 20;
  background: repeating-linear-gradient(0deg, rgba(0, 0, 0, 0.18) 0 1px, transparent 1px 3px);
  opacity: 0.5;
}
.vh-glitch { text-shadow: 2px 0 #ff4fa3, -2px 0 #66e0ff; }
.vh-primary {
  border: 2px solid #ff4fa3;
  background: linear-gradient(155deg, rgba(255, 79, 163, 0.1), rgba(102, 224, 255, 0.035));
  box-shadow: 0 0 0 4px rgba(255, 79, 163, 0.08), 0 24px 50px -34px rgba(255, 79, 163, 0.7);
  transition: box-shadow 0.2s ease, transform 0.2s ease;
}
.vh-primary:hover { transform: translateY(-2px); box-shadow: 0 0 0 5px rgba(255, 79, 163, 0.14), 0 28px 60px -30px rgba(255, 79, 163, 0.8); }
.vh-primary:active { transform: scale(0.99); }
.vh-btn { box-shadow: 2px 2px 0 #66e0ff; }
.vh-row {
  border: 1px solid rgba(255, 255, 255, 0.1);
  background: rgba(255, 255, 255, 0.035);
  transition: border-color 0.18s ease, background-color 0.18s ease, transform 0.18s ease;
}
.vh-row:hover { border-color: rgba(255, 79, 163, 0.55); background: rgba(255, 79, 163, 0.08); transform: translateX(2px); }
.vh-row:active { transform: scale(0.99); }
.vh-enter { animation: vh-in 0.65s cubic-bezier(0.22, 1, 0.36, 1) both; }
.vh-d1 { animation-delay: 0.08s; }
.vh-d2 { animation-delay: 0.16s; }
.vh-d3 { animation-delay: 0.24s; }
@keyframes vh-in {
  0% { opacity: 0; transform: translateY(12px); filter: blur(2px); }
  60% { opacity: 0.85; filter: blur(0); }
  100% { opacity: 1; transform: translateY(0); filter: blur(0); }
}
.vh-rec { animation: vh-blink 1.6s ease-in-out infinite; }
@keyframes vh-blink { 50% { opacity: 0.2; } }
@media (prefers-reduced-motion: reduce) {
  .vh-enter, .vh-rec { animation: none; }
  .vh-primary, .vh-row { transition: none; }
}
</style>
