<template>
  <div class="ck-root min-h-screen text-[#3d2f26]">
    <main class="mx-auto w-full max-w-md px-5 pb-12 pt-6">
      <div class="ck-enter flex items-center justify-between text-[9px] font-bold uppercase tracking-[0.2em] text-[#8d6e5a]">
        <span>Kindly Keep · 请保留</span>
        <span v-if="nfcId" class="max-w-[42%] truncate">NFC · {{ nfcId }}</span>
      </div>

      <header class="ck-enter ck-d1 relative mt-5">
        <div class="ck-tape absolute -top-2 left-8 z-10 -rotate-3"></div>
        <div class="ck-paper relative rotate-[-0.8deg] px-5 pb-5 pt-6">
          <p class="ck-hand text-[13px] text-[#b45309]">如果你看到这张纸，请不要慌。</p>
          <h1 class="mt-2 text-[36px] font-black leading-[1.02] tracking-tight">
            先帮我<br /><span class="ck-marker">打个电话</span>
          </h1>
          <p class="mt-4 max-w-[280px] text-[12px] leading-relaxed text-[#7a6354]">告诉对方我们所在的位置，然后在旁边陪我等待救援。谢谢。</p>
          <div class="mt-4 flex items-center gap-2 text-[10px] font-bold text-[#a16207]">
            <span class="h-px w-8 bg-[#a16207]/50"></span>EMERGENCY CONTACT
          </div>
        </div>
      </header>

      <section class="relative mt-5">
        <template v-if="contacts.length">
          <div class="ck-enter ck-d2 relative">
            <div class="ck-tape absolute -top-2 right-9 z-10 rotate-2"></div>
            <a :href="'tel:' + contacts[0].phone" class="ck-primary block rotate-[0.6deg] px-5 pb-5 pt-5">
              <div class="flex items-center justify-between">
                <span class="rounded-full bg-[#9a3412] px-2.5 py-1 text-[9px] font-bold uppercase tracking-[0.16em] text-[#fff7ed]">First call</span>
                <Phone class="h-5 w-5 text-[#9a3412]" />
              </div>
              <p class="ck-phone mt-4 break-words text-[36px] font-black leading-none tabular-nums tracking-tight text-[#7c2d12]">{{ contacts[0].phone }}</p>
              <p class="mt-3 text-[13px] font-bold">{{ contacts[0].name || '联系人' }}<span v-if="contacts[0].relation" class="font-normal text-[#8d6e5a]"> · {{ contacts[0].relation }}</span></p>
              <p class="ck-hand mt-2 text-[12px] text-[#b45309]">点一下就能拨号 →</p>
            </a>
          </div>

          <div class="mt-4 space-y-3">
            <a
              v-for="(c, i) in contacts.slice(1)"
              :key="i"
              :href="'tel:' + c.phone"
              class="ck-note ck-enter ck-d3 flex items-center justify-between gap-3 px-4 py-3.5"
              :class="i % 2 === 0 ? 'rotate-[0.7deg]' : '-rotate-[0.5deg]'"
            >
              <div class="min-w-0">
                <p class="truncate text-[22px] font-bold leading-tight tabular-nums">{{ c.phone }}</p>
                <p class="mt-1 truncate text-[11px] text-[#8d6e5a]">{{ c.name || '联系人' }}<span v-if="c.relation"> · {{ c.relation }}</span></p>
              </div>
              <Phone class="h-4 w-4 flex-shrink-0 text-[#9a3412]" />
            </a>
          </div>
        </template>
        <div v-else class="ck-paper rotate-[-0.5deg] px-5 py-8 text-center text-[12px] text-[#8d6e5a]">这张纸上还没有写号码，请帮我们补充。</div>
      </section>

      <section class="ck-enter ck-d3 relative mt-5">
        <div class="ck-tape absolute -top-2 left-5 z-10 rotate-1"></div>
        <div class="ck-med px-4 pb-4 pt-5">
          <div class="flex items-center justify-between">
            <p class="text-[10px] font-black uppercase tracking-[0.18em] text-[#9a3412]">Medical · 身体备注</p>
            <HeartPulse class="h-4 w-4 text-[#9a3412]" />
          </div>
          <div class="mt-3 space-y-2">
            <div v-for="m in medical" :key="m.label" class="flex items-start gap-3 border-b border-dashed border-[#c4a88e] pb-2 last:border-b-0 last:pb-0">
              <span class="w-10 flex-shrink-0 text-[10px] font-bold text-[#8d6e5a]">{{ m.label }}</span>
              <span class="text-[12px] font-medium leading-snug">{{ m.value }}</span>
            </div>
          </div>
        </div>
      </section>

      <div v-if="hasPet" class="ck-enter ck-d3 mt-4 flex items-center justify-between rounded-sm bg-[#e7d7bd] px-4 py-3 text-[12px] shadow-[2px_3px_0_rgba(120,90,60,0.12)]">
        <span><strong>{{ payload.petName }}</strong><span v-if="payload.breed" class="text-[#8d6e5a]"> · {{ payload.breed }}</span></span>
        <a v-if="payload.vetPhone" :href="'tel:' + payload.vetPhone" class="font-bold text-[#7c2d12] underline decoration-dotted underline-offset-4">医院 {{ payload.vetPhone }}</a>
      </div>

      <p class="ck-enter ck-d3 ck-hand mt-6 text-center text-[12px] text-[#a18875]">谢谢你停下來帮忙。</p>
    </main>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { Phone, HeartPulse } from 'lucide-vue-next'

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
.ck-root {
  background:
    radial-gradient(rgba(120, 90, 60, 0.08) 0.8px, transparent 0.8px),
    linear-gradient(160deg, #f3e7d8, #e9d8c4);
  background-size: 6px 6px, 100% 100%;
}
.ck-hand { font-family: "Segoe Script", "Bradley Hand", "Comic Sans MS", cursive; }
.ck-paper {
  background: #fdf6e9;
  border: 1px solid #d9c2a8;
  box-shadow: 0 16px 38px -28px rgba(80, 50, 20, 0.6), 2px 3px 0 rgba(120, 90, 60, 0.08);
}
.ck-primary {
  background:
    linear-gradient(90deg, rgba(180, 83, 9, 0.06) 1px, transparent 1px),
    #fff8eb;
  background-size: 24px 100%;
  border: 1px solid #d9b896;
  box-shadow: 5px 6px 0 rgba(154, 52, 18, 0.14), 0 20px 40px -32px rgba(80, 50, 20, 0.6);
  transition: transform 0.25s ease, box-shadow 0.25s ease;
}
.ck-primary:hover { transform: rotate(0) translateY(-2px); box-shadow: 7px 9px 0 rgba(154, 52, 18, 0.18); }
.ck-primary:active { transform: scale(0.985); }
.ck-marker {
  background: linear-gradient(transparent 56%, rgba(245, 158, 11, 0.45) 56% 92%, transparent 92%);
  padding: 0 2px;
}
.ck-tape {
  width: 70px;
  height: 22px;
  background: rgba(245, 230, 200, 0.75);
  border-left: 1px solid rgba(190, 160, 120, 0.3);
  border-right: 1px solid rgba(190, 160, 120, 0.3);
  box-shadow: 0 1px 2px rgba(90, 60, 30, 0.08);
  transform: rotate(-3deg);
}
.ck-note {
  background: #f7ecd9;
  border: 1px solid #dcc6ab;
  box-shadow: 2px 3px 0 rgba(120, 90, 60, 0.1);
  transition: transform 0.2s ease, background-color 0.2s ease;
}
.ck-note:hover { transform: rotate(0) translateY(-2px); background: #fff6e5; }
.ck-note:active { transform: scale(0.985); }
.ck-med {
  background: #f7efdf;
  border: 1px dashed #b99a7b;
  box-shadow: 2px 3px 0 rgba(120, 90, 60, 0.08);
}
.ck-enter { animation: ck-in 0.75s cubic-bezier(0.22, 1, 0.36, 1) both; }
.ck-d1 { animation-delay: 0.08s; }
.ck-d2 { animation-delay: 0.16s; }
.ck-d3 { animation-delay: 0.24s; }
@keyframes ck-in {
  from { opacity: 0; transform: translateY(16px) rotate(-1.5deg); }
  to { opacity: 1; transform: translateY(0) rotate(0); }
}
@media (prefers-reduced-motion: reduce) {
  .ck-enter { animation: none; }
  .ck-primary, .ck-note { transition: none; }
}
</style>
