<template>
  <div class="hh-root min-h-screen text-[#28382b]">
    <main class="mx-auto w-full max-w-md px-5 pb-12 pt-6">
      <div class="hh-enter flex items-center justify-between text-[9px] uppercase tracking-[0.22em] text-[#6f8065]">
        <span class="flex items-center gap-2"><Leaf class="h-3.5 w-3.5" />Herb &amp; Home</span>
        <span v-if="nfcId" class="max-w-[42%] truncate">NFC · {{ nfcId }}</span>
      </div>

      <header class="hh-enter hh-d1 relative mt-6">
        <svg class="hh-sprig" viewBox="0 0 120 180" fill="none" aria-hidden="true">
          <path d="M78 176C61 131 56 83 70 18" stroke="currentColor" stroke-width="1.2" />
          <path d="M68 62c-22-3-31-16-32-35 20 1 31 13 32 35Z" stroke="currentColor" stroke-width="1.2" />
          <path d="M70 94c21-4 30-16 31-34-19 1-29 13-31 34Z" stroke="currentColor" stroke-width="1.2" />
          <path d="M72 126c-20-2-30-13-31-31 19 1 29 12 31 31Z" stroke="currentColor" stroke-width="1.2" />
          <path d="M76 150c18-3 27-13 28-29-17 1-26 11-28 29Z" stroke="currentColor" stroke-width="1.2" />
        </svg>
        <p class="text-[10px] uppercase tracking-[0.24em] text-[#8a6a4a]">A quiet note for a hard moment</p>
        <h1 class="hh-serif mt-2 text-[35px] font-medium leading-[1.14]">需要帮助时，<br />请联络我的家人</h1>
        <p class="mt-4 max-w-[260px] text-[12px] leading-6 text-[#67765f]">先拨第一通电话。若无人接听，请沿着名单继续。告诉对方我们在哪里。</p>
        <div class="mt-5 flex items-center gap-3">
          <span class="h-px flex-1 bg-[#b5c1a6]"></span><Leaf class="h-3.5 w-3.5 text-[#6c8a58]" /><span class="h-px flex-1 bg-[#b5c1a6]"></span>
        </div>
      </header>

      <section class="mt-5">
        <div class="hh-enter hh-d2 flex items-end justify-between">
          <p class="text-[9px] uppercase tracking-[0.22em] text-[#6f8065]">Garden of contacts · 联络名单</p>
          <span class="text-[9px] tabular-nums text-[#8a9a80]">{{ String(contacts.length).padStart(2, '0') }}</span>
        </div>

        <template v-if="contacts.length">
          <a :href="'tel:' + contacts[0].phone" class="hh-primary hh-enter hh-d2 mt-3 block px-5 py-5">
            <div class="flex items-center justify-between">
              <span class="inline-flex items-center gap-1.5 rounded-full bg-[#e3ecd8] px-2.5 py-1 text-[9px] font-bold uppercase tracking-[0.12em] text-[#4f6b43]"><Sprout class="h-3 w-3" />First call</span>
              <Phone class="h-4 w-4 text-[#4f6b43]" />
            </div>
            <p class="hh-serif mt-3 break-words text-[34px] font-semibold leading-none tabular-nums tracking-wide text-[#213527]">{{ contacts[0].phone }}</p>
            <div class="mt-4 flex items-center justify-between border-t border-[#c8d4bc] pt-3">
              <p class="text-[12px] font-semibold">{{ contacts[0].name || '联系人' }}<span v-if="contacts[0].relation" class="font-normal text-[#7b8b71]"> · {{ contacts[0].relation }}</span></p>
              <span class="text-[9px] uppercase tracking-[0.14em] text-[#6c8a58]">Tap to call</span>
            </div>
          </a>

          <div class="mt-2">
            <a
              v-for="(c, i) in contacts.slice(1)"
              :key="i"
              :href="'tel:' + c.phone"
              class="hh-row flex items-center gap-3 border-b border-[#c8d4bc] py-3.5"
            >
              <span class="hh-serif w-5 text-[12px] italic text-[#8a9a80]">{{ ['ii', 'iii', 'iv', 'v'][i] || '·' }}</span>
              <div class="min-w-0 flex-1">
                <p class="truncate text-[20px] font-semibold leading-tight tabular-nums text-[#28382b]">{{ c.phone }}</p>
                <p class="mt-0.5 truncate text-[10px] text-[#7b8b71]">{{ c.name || '联系人' }}<span v-if="c.relation"> · {{ c.relation }}</span></p>
              </div>
              <Phone class="h-4 w-4 shrink-0 text-[#6c8a58]" />
            </a>
          </div>
        </template>
        <div v-else class="mt-3 rounded-2xl border border-dashed border-[#b5c1a6] px-4 py-8 text-center text-[12px] text-[#7b8b71]">名单还是空的，请补充联络人。</div>
      </section>

      <section class="hh-enter hh-d3 mt-5 rounded-2xl border border-[#c8d4bc] bg-[#f7f5ec] px-4 py-4">
        <div class="flex items-center justify-between">
          <p class="text-[9px] uppercase tracking-[0.2em] text-[#6f8065]">Care notes · 照护备注</p>
          <HeartPulse class="h-4 w-4 text-[#8a6a4a]" />
        </div>
        <div class="mt-3 grid grid-cols-2 gap-x-4 gap-y-3">
          <div v-for="m in medical" :key="m.label" class="border-t border-[#d8dfce] pt-2">
            <p class="text-[9px] uppercase tracking-[0.1em] text-[#98a38d]">{{ m.label }}</p>
            <p class="mt-1 text-[11px] font-medium leading-snug text-[#3e5142]">{{ m.value }}</p>
          </div>
        </div>
      </section>

      <div v-if="hasPet" class="hh-enter hh-d3 mt-3 flex items-center justify-between rounded-2xl bg-[#e6ecda] px-4 py-3 text-[11px] text-[#4e6245]">
        <span class="flex items-center gap-2"><PawPrint class="h-3.5 w-3.5" />{{ payload.petName }}<span v-if="payload.breed" class="text-[#7b8b71]"> · {{ payload.breed }}</span></span>
        <a v-if="payload.vetPhone" :href="'tel:' + payload.vetPhone" class="font-semibold text-[#3f5b35] underline decoration-[#3f5b35]/30 underline-offset-4">宠物医院</a>
      </div>

      <p class="hh-enter hh-d3 mt-6 text-center text-[9px] uppercase tracking-[0.24em] text-[#98a38d]">Breathe · Call · Stay</p>
    </main>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { Leaf, Sprout, Phone, HeartPulse, PawPrint } from 'lucide-vue-next'

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
.hh-root {
  background:
    radial-gradient(60% 40% at 100% 0%, rgba(108, 138, 88, 0.1), transparent 65%),
    linear-gradient(170deg, #f1f3e8, #e7ecdd);
}
.hh-serif { font-family: "Iowan Old Style", "Palatino Linotype", Georgia, "Songti SC", serif; }
.hh-sprig {
  position: absolute;
  right: -8px;
  top: -16px;
  width: 92px;
  height: 140px;
  color: rgba(108, 138, 88, 0.32);
}
.hh-primary {
  border: 1px solid #b5c9a5;
  background: linear-gradient(155deg, #fbfaf2, #f2f5e8);
  box-shadow: 0 18px 36px -30px rgba(50, 75, 45, 0.5);
  transition: transform 0.22s ease, box-shadow 0.22s ease, border-color 0.22s ease;
}
.hh-primary:hover { transform: translateY(-2px); border-color: #7f9c6c; box-shadow: 0 24px 44px -28px rgba(50, 75, 45, 0.55); }
.hh-primary:active { transform: scale(0.99); }
.hh-row { transition: background-color 0.18s ease, padding-left 0.18s ease, border-color 0.18s ease; }
.hh-row:hover { background: rgba(108, 138, 88, 0.06); border-color: #8a9f79; padding-left: 4px; }
.hh-enter { animation: hh-in 0.8s cubic-bezier(0.22, 1, 0.36, 1) both; }
.hh-d1 { animation-delay: 0.08s; }
.hh-d2 { animation-delay: 0.16s; }
.hh-d3 { animation-delay: 0.24s; }
@keyframes hh-in {
  from { opacity: 0; transform: translateY(14px); filter: blur(1px); }
  to { opacity: 1; transform: translateY(0); filter: blur(0); }
}
@media (prefers-reduced-motion: reduce) {
  .hh-enter { animation: none; }
  .hh-primary, .hh-row { transition: none; }
}
</style>
