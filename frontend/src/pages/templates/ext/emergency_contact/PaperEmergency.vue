<template>
  <div class="min-h-screen bg-[#f4eee1] text-[#261f19]">
    <main class="mx-auto w-full max-w-md px-6 pb-12 pt-7">
      <header class="pe-enter">
        <div class="flex items-center justify-between border-b border-[#261f19] pb-2 text-[10px] uppercase tracking-[0.22em]">
          <span>Emergency Paper</span>
          <span class="tabular-nums">No. 01 / 2026</span>
        </div>
        <div class="mt-5 flex items-start justify-between gap-4">
          <div>
            <p class="text-[11px] tracking-[0.18em] text-[#9f2d20]">请交给现场救援人员</p>
            <h1 class="pe-serif mt-2 text-[42px] font-bold leading-[0.98] tracking-tight">紧急<br />联系人</h1>
          </div>
          <div class="mt-1 flex h-16 w-16 flex-shrink-0 rotate-3 flex-col items-center justify-center rounded-full border border-[#9f2d20] text-center text-[#9f2d20]">
            <Siren class="h-5 w-5" />
            <span class="mt-1 text-[9px] font-bold tracking-[0.14em]">SOS</span>
          </div>
        </div>
        <p class="pe-serif mt-4 border-t border-[#c9bda9] pt-3 text-[14px] italic leading-relaxed text-[#5f5346]">
          若持卡人失去意识或无法应答，请按顺序拨号，并说明现场地址。
        </p>
      </header>

      <section class="mt-7">
        <div class="flex items-center gap-3">
          <span class="h-px flex-1 bg-[#261f19]"></span>
          <p class="text-[10px] font-bold uppercase tracking-[0.24em]">Contact Index</p>
          <span class="h-px flex-1 bg-[#261f19]"></span>
        </div>

        <template v-if="contacts.length">
          <article v-for="(c, i) in contacts" :key="i" class="pe-item pe-enter" :class="'pe-d' + Math.min(i + 1, 3)">
            <div class="flex items-baseline justify-between gap-3 border-b border-[#c9bda9] pb-2 pt-5">
              <div class="flex min-w-0 items-baseline gap-3">
                <span class="pe-serif text-[12px] italic tabular-nums text-[#9f2d20]">{{ String(i + 1).padStart(2, '0') }}</span>
                <p class="truncate text-[13px] font-semibold tracking-wide">{{ c.name || '联系人' }}</p>
              </div>
              <span v-if="c.relation" class="flex-shrink-0 text-[11px] text-[#7c6e5e]">{{ c.relation }}</span>
            </div>
            <a :href="'tel:' + c.phone" class="pe-phone group block pt-2.5">
              <span class="pe-serif block break-words text-[30px] leading-tight tabular-nums tracking-tight transition-colors group-hover:text-[#9f2d20]" :class="i === 0 ? 'font-bold' : 'font-semibold'">{{ c.phone }}</span>
              <span class="mt-1 flex items-center gap-1.5 text-[10px] uppercase tracking-[0.18em] text-[#8b7d6b]">
                <Phone class="h-3 w-3" /> 点击拨号
              </span>
            </a>
          </article>
        </template>
        <div v-else class="mt-5 border border-dashed border-[#b9ab95] px-5 py-8 text-center text-[13px] text-[#7c6e5e]">暂无联系人号码 · 请补充后重试</div>
      </section>

      <section class="pe-appendix pe-enter pe-d3 mt-8 border-t-2 border-[#261f19] pt-4">
        <div class="flex items-center justify-between">
          <h2 class="pe-serif text-[15px] font-bold italic">附：医疗备注</h2>
          <HeartPulse class="h-4 w-4 text-[#9f2d20]" />
        </div>
        <dl class="mt-3 divide-y divide-[#d8cdbb] border-y border-[#d8cdbb]">
          <div v-for="m in medical" :key="m.label" class="grid grid-cols-[64px_1fr] gap-3 py-2.5">
            <dt class="text-[10px] uppercase tracking-[0.12em] text-[#8b7d6b]">{{ m.label }}</dt>
            <dd class="text-[12px] font-medium leading-snug">{{ m.value }}</dd>
          </div>
        </dl>
      </section>

      <div v-if="hasPet" class="pe-enter pe-d3 mt-4 border border-[#261f19] px-4 py-3">
        <p class="text-[10px] uppercase tracking-[0.2em] text-[#9f2d20]">随行宠物</p>
        <p class="mt-1 text-[13px] font-semibold">{{ payload.petName }}<span v-if="payload.breed" class="font-normal text-[#6b5f51]"> · {{ payload.breed }}</span></p>
        <a v-if="payload.vetPhone" :href="'tel:' + payload.vetPhone" class="mt-1 inline-block border-b border-[#9f2d20]/40 text-[12px] text-[#9f2d20]">{{ payload.vetPhone }} · 宠物医院</a>
      </div>

      <footer class="pe-enter pe-d3 mt-8 flex items-center justify-between border-t border-[#261f19] pt-3 text-[9px] uppercase tracking-[0.2em] text-[#8b7d6b]">
        <span>Emergency Contact Sheet</span>
        <span v-if="nfcId" class="max-w-[45%] truncate text-right">NFC · {{ nfcId }}</span>
      </footer>
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
  { label: '过敏史', value: props.payload?.allergies },
  { label: '正在用药', value: props.payload?.medications },
  { label: '既往病史', value: props.payload?.conditions },
])
const hasPet = computed(() => !!(props.payload?.petName || props.payload?.chipId))
</script>

<style scoped>
.pe-serif {
  font-family: "Iowan Old Style", "Palatino Linotype", Georgia, "Songti SC", serif;
}
.pe-enter { animation: pe-fade 0.8s cubic-bezier(0.22, 1, 0.36, 1) both; }
.pe-d1 { animation-delay: 0.08s; }
.pe-d2 { animation-delay: 0.16s; }
.pe-d3 { animation-delay: 0.24s; }
@keyframes pe-fade {
  from { opacity: 0; transform: translateY(10px); }
  to { opacity: 1; transform: translateY(0); }
}
.pe-item { animation: pe-fade 0.7s cubic-bezier(0.22, 1, 0.36, 1) both; }
.pe-phone { text-decoration: none; }
.pe-phone:hover span:first-child { color: #9f2d20; }
.pe-phone:active span:first-child { opacity: 0.7; }
@media (prefers-reduced-motion: reduce) {
  .pe-enter, .pe-item { animation: none; }
  .pe-phone span { transition: none; }
}
</style>
