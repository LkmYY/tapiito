<template>
  <div class="au-root min-h-screen text-white">
    <main class="mx-auto w-full max-w-md px-5 pb-12 pt-6">
      <div class="au-enter relative z-10 flex items-center justify-between text-[10px] uppercase tracking-[0.24em] text-white/60">
        <span class="flex items-center gap-2"><Sparkles class="h-3.5 w-3.5" />Aurora Line</span>
        <span v-if="nfcId" class="max-w-[42%] truncate">NFC · {{ nfcId }}</span>
      </div>

      <header class="au-enter au-d1 relative z-10 mt-10 text-center">
        <div class="au-halo mx-auto flex h-20 w-20 items-center justify-center rounded-full border border-white/25">
          <div class="flex h-14 w-14 items-center justify-center rounded-full bg-white/10">
            <Siren class="h-6 w-6 text-[#bae6fd]" />
          </div>
        </div>
        <p class="mt-5 text-[10px] uppercase tracking-[0.3em] text-[#a5f3fc]">One tap · one call</p>
        <h1 class="mt-2 text-[34px] font-black leading-tight tracking-tight">紧急联络</h1>
        <p class="mx-auto mt-3 max-w-[260px] text-[12px] leading-relaxed text-white/60">拨通后先报出位置。若无法说话，请让旁人代为说明。</p>
      </header>

      <section class="au-enter au-d2 relative z-10 mt-8">
        <a v-if="contacts.length" :href="'tel:' + contacts[0].phone" class="au-call block px-6 pb-6 pt-5 text-center">
          <p class="text-[9px] uppercase tracking-[0.26em] text-white/50">Primary · 首要联系人</p>
          <p class="mt-3 break-words text-[36px] font-black leading-none tabular-nums tracking-tight">{{ contacts[0].phone }}</p>
          <p class="mt-3 text-[12px] font-semibold text-[#cffafe]">{{ contacts[0].name || '联系人' }}<span v-if="contacts[0].relation" class="font-normal text-white/50"> · {{ contacts[0].relation }}</span></p>
          <span class="au-dial mx-auto mt-4 flex h-11 items-center justify-center gap-2 rounded-full bg-white px-5 text-[11px] font-black uppercase tracking-[0.16em] text-[#170b2e]">
            <Phone class="h-4 w-4" /> 立即拨打
          </span>
        </a>
      </section>

      <section v-if="contacts.length > 1" class="au-enter au-d3 relative z-10 mt-4 grid grid-cols-2 gap-2.5">
        <a
          v-for="(c, i) in contacts.slice(1)"
          :key="i"
          :href="'tel:' + c.phone"
          class="au-chip rounded-2xl px-4 py-3.5"
        >
          <p class="text-[9px] uppercase tracking-[0.16em] text-white/40">{{ String(i + 2).padStart(2, '0') }} · {{ c.relation || '备用' }}</p>
          <p class="mt-1.5 break-words text-[17px] font-bold leading-tight tabular-nums">{{ c.phone }}</p>
          <p class="mt-1 truncate text-[10px] text-white/50">{{ c.name || '联系人' }}</p>
        </a>
      </section>

      <section class="au-enter au-d3 relative z-10 mt-4 rounded-2xl border border-white/10 bg-white/[0.055] px-4 py-3.5">
        <div class="flex items-center justify-between">
          <p class="text-[9px] uppercase tracking-[0.22em] text-[#a5f3fc]">Medical</p>
          <HeartPulse class="h-4 w-4 text-[#f0abfc]" />
        </div>
        <div class="mt-3 grid grid-cols-2 gap-x-4 gap-y-2.5">
          <div v-for="m in medical" :key="m.label">
            <p class="text-[9px] text-white/35">{{ m.label }}</p>
            <p class="mt-0.5 text-[11px] font-semibold leading-snug text-white/85">{{ m.value }}</p>
          </div>
        </div>
      </section>

      <div v-if="hasPet" class="au-enter au-d3 relative z-10 mt-3 flex items-center justify-between rounded-2xl border border-white/10 bg-white/[0.055] px-4 py-3 text-[12px]">
        <span class="text-white/75">{{ payload.petName }}<span v-if="payload.breed" class="text-white/40"> · {{ payload.breed }}</span></span>
        <a v-if="payload.vetPhone" :href="'tel:' + payload.vetPhone" class="text-[#a5f3fc]">宠物医院</a>
      </div>

      <p class="au-enter au-d3 relative z-10 mt-7 text-center text-[9px] uppercase tracking-[0.26em] text-white/30">You are not alone · 有人在赶来</p>
    </main>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { Sparkles, Siren, Phone, HeartPulse } from 'lucide-vue-next'

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
.au-root {
  position: relative;
  overflow: hidden;
  background:
    radial-gradient(65% 40% at 15% 10%, rgba(6, 182, 212, 0.42), transparent 65%),
    radial-gradient(70% 45% at 90% 25%, rgba(168, 85, 247, 0.4), transparent 65%),
    radial-gradient(55% 40% at 45% 85%, rgba(20, 184, 166, 0.3), transparent 70%),
    #170b2e;
}
.au-root::before {
  content: "";
  position: fixed;
  inset: -20%;
  background: conic-gradient(from 120deg at 50% 50%, transparent, rgba(34, 211, 238, 0.14), transparent 40%, rgba(217, 70, 239, 0.12), transparent 70%);
  animation: au-spin 26s linear infinite;
  pointer-events: none;
}
.au-halo { background: radial-gradient(circle, rgba(165, 243, 252, 0.18), transparent 68%); box-shadow: 0 0 50px -8px rgba(34, 211, 238, 0.4); }
.au-call {
  background: linear-gradient(135deg, rgba(255, 255, 255, 0.14), rgba(255, 255, 255, 0.04));
  border: 1px solid rgba(255, 255, 255, 0.22);
  border-radius: 999px;
  box-shadow: 0 30px 70px -30px rgba(34, 211, 238, 0.55), inset 0 1px 0 rgba(255, 255, 255, 0.2);
  backdrop-filter: blur(14px);
  transition: transform 0.25s ease, box-shadow 0.25s ease;
}
.au-call:hover { transform: translateY(-2px); box-shadow: 0 36px 80px -28px rgba(34, 211, 238, 0.7); }
.au-call:active { transform: scale(0.985); }
.au-dial { transition: transform 0.2s ease, background-color 0.2s ease; box-shadow: 0 6px 20px rgba(0, 0, 0, 0.25); }
.au-dial:hover { transform: scale(1.03); background: #ecfeff; }
.au-chip {
  border: 1px solid rgba(255, 255, 255, 0.12);
  background: rgba(255, 255, 255, 0.055);
  backdrop-filter: blur(12px);
  transition: background-color 0.2s ease, border-color 0.2s ease, transform 0.2s ease;
}
.au-chip:hover { background: rgba(255, 255, 255, 0.11); border-color: rgba(165, 243, 252, 0.4); transform: translateY(-2px); }
.au-chip:active { transform: scale(0.97); }
.au-enter { animation: au-rise 0.8s cubic-bezier(0.22, 1, 0.36, 1) both; }
.au-d1 { animation-delay: 0.08s; }
.au-d2 { animation-delay: 0.16s; }
.au-d3 { animation-delay: 0.24s; }
@keyframes au-rise {
  from { opacity: 0; transform: translateY(18px) scale(0.98); }
  to { opacity: 1; transform: translateY(0) scale(1); }
}
@keyframes au-spin { to { transform: rotate(360deg); } }
@media (prefers-reduced-motion: reduce) {
  .au-root::before { animation: none; }
  .au-enter { animation: none; }
  .au-call, .au-dial, .au-chip { transition: none; }
}
</style>
