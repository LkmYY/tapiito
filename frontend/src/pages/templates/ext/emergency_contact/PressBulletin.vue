<template>
  <div class="pb-root min-h-screen text-[#181818]">
    <main class="mx-auto w-full max-w-md px-5 pb-12 pt-5">
      <div class="pb-enter flex items-center justify-between border-b border-[#181818] pb-1.5 text-[8px] font-bold uppercase tracking-[0.18em]">
        <span>紧急号外 · Special Edition</span>
        <span v-if="nfcId" class="max-w-[42%] truncate">NFC {{ nfcId }}</span>
      </div>

      <header class="pb-enter pb-d1 mt-2 text-center">
        <h1 class="pb-masthead text-[37px] font-black leading-none tracking-[0.06em]">紧急联络公报</h1>
        <div class="mt-1 border-y-2 border-[#181818] py-1 text-[8px] font-bold uppercase tracking-[0.26em]">The Emergency Contact Bulletin</div>
        <div class="mt-1 flex items-center justify-between border-b border-[#181818] pb-1 text-[8px] uppercase tracking-[0.12em] text-[#555]">
          <span>2026 年 · 紧急版</span><span>请传阅 · 请协助</span>
        </div>
      </header>

      <section class="pb-enter pb-d2 relative mt-4">
        <div class="pb-stamp absolute -right-1 top-0 z-10 flex h-14 w-14 rotate-6 items-center justify-center rounded-full border-[3px] border-[#a1281e] text-center text-[11px] font-black leading-tight text-[#a1281e]">号<br />外</div>
        <p class="text-[9px] font-bold uppercase tracking-[0.2em] text-[#a1281e]">头条 · 现场求助</p>
        <h2 class="pb-serif mt-1 max-w-[280px] text-[36px] font-black leading-[1.02] tracking-tight">发现此人，<br />请先拨打电话</h2>
        <div class="mt-3 grid grid-cols-[1fr_auto] items-start gap-3 border-t border-[#181818] pt-2">
          <p class="pb-serif text-[11px] leading-[1.7] text-[#3b3b3b]">持卡人可能无法清楚表达。请按下列名单联络家属，并向接线人说明地点、意识与呼吸情况。不要让其独处。</p>
          <Siren class="mt-1 h-5 w-5 text-[#a1281e]" />
        </div>
      </section>

      <section class="mt-5">
        <div class="pb-enter pb-d2 flex items-center gap-2 border-y-2 border-[#181818] py-1">
          <span class="bg-[#181818] px-1.5 py-0.5 text-[8px] font-bold uppercase tracking-[0.12em] text-[#f4efe2]">CALL LIST</span>
          <span class="text-[8px] uppercase tracking-[0.16em]">联络名单 · 按序拨打</span>
        </div>

        <template v-if="contacts.length">
          <a :href="'tel:' + contacts[0].phone" class="pb-primary pb-enter pb-d2 mt-3 block border-b-2 border-[#181818] pb-3">
            <div class="flex items-start justify-between gap-3">
              <p class="text-[9px] font-bold uppercase tracking-[0.16em] text-[#a1281e]">01 / 首要联络</p>
              <span class="border border-[#181818] px-1.5 py-0.5 text-[8px] font-bold">请立即</span>
            </div>
            <p class="pb-serif mt-1 break-words text-[35px] font-black leading-tight tabular-nums tracking-tight">{{ contacts[0].phone }}</p>
            <div class="mt-1 flex items-center justify-between">
              <p class="text-[11px] font-bold">{{ contacts[0].name || '联系人' }}<span v-if="contacts[0].relation" class="font-normal text-[#555]"> · {{ contacts[0].relation }}</span></p>
              <span class="flex items-center gap-1 text-[8px] font-bold uppercase tracking-[0.12em]"><Phone class="h-3 w-3" />拨号</span>
            </div>
          </a>

          <div class="mt-2">
            <a
              v-for="(c, i) in contacts.slice(1)"
              :key="i"
              :href="'tel:' + c.phone"
              class="pb-row grid grid-cols-[24px_1fr_auto] items-center gap-2 border-b border-[#b9b2a1] py-3"
            >
              <span class="text-[9px] font-bold tabular-nums text-[#a1281e]">{{ String(i + 2).padStart(2, '0') }}</span>
              <div class="min-w-0">
                <p class="truncate text-[20px] font-black leading-tight tabular-nums">{{ c.phone }}</p>
                <p class="mt-0.5 truncate text-[9px] text-[#555]">{{ c.name || '联系人' }}<span v-if="c.relation"> · {{ c.relation }}</span></p>
              </div>
              <Phone class="h-3.5 w-3.5" />
            </a>
          </div>
        </template>
        <div v-else class="mt-3 border-b border-dashed border-[#999] py-7 text-center text-[11px] text-[#555]">名单空缺 · 请协助补充</div>
      </section>

      <section class="pb-enter pb-d3 mt-1 columns-2 gap-4 border-b border-[#181818] py-3">
        <h3 class="mb-1 break-inside-avoid border-b border-[#181818] pb-0.5 text-[10px] font-black uppercase tracking-[0.12em]">医疗简讯 / Medical</h3>
        <div v-for="m in medical" :key="m.label" class="mb-2 break-inside-avoid">
          <p class="text-[8px] font-bold uppercase tracking-[0.08em] text-[#777]">{{ m.label }}</p>
          <p class="pb-serif text-[10px] leading-snug">{{ m.value }}</p>
        </div>
      </section>

      <div v-if="hasPet" class="pb-enter pb-d3 mt-2 border-b border-[#181818] py-2 text-[10px]">
        <span class="font-black uppercase tracking-[0.1em]">寻宠附注 /</span> {{ payload.petName }}<span v-if="payload.breed"> · {{ payload.breed }}</span>
        <a v-if="payload.vetPhone" :href="'tel:' + payload.vetPhone" class="ml-1 font-bold text-[#a1281e] underline underline-offset-2">{{ payload.vetPhone }}</a>
      </div>

      <footer class="pb-enter pb-d3 mt-3 flex items-center justify-between text-[8px] uppercase tracking-[0.14em] text-[#666]">
        <span>Printed for emergency use · 请勿丢弃</span><span>Page 1 / 1</span>
      </footer>
    </main>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { Siren, Phone } from 'lucide-vue-next'

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
.pb-root {
  background:
    radial-gradient(rgba(80, 70, 50, 0.055) 0.7px, transparent 0.7px),
    #f1ebdc;
  background-size: 4px 4px;
}
.pb-serif { font-family: "Times New Roman", "Songti SC", "Noto Serif SC", serif; }
.pb-masthead { font-family: "Songti SC", "Noto Serif SC", "Times New Roman", serif; }
.pb-primary { transition: background-color 0.2s ease; }
.pb-primary:hover { background: rgba(161, 40, 30, 0.045); }
.pb-primary:active { background: rgba(161, 40, 30, 0.08); }
.pb-row { transition: background-color 0.18s ease, padding-left 0.18s ease; }
.pb-row:hover { background: rgba(24, 24, 24, 0.04); padding-left: 4px; }
.pb-stamp { background: rgba(241, 235, 220, 0.9); }
.pb-enter { animation: pb-in 0.75s cubic-bezier(0.22, 1, 0.36, 1) both; }
.pb-d1 { animation-delay: 0.08s; }
.pb-d2 { animation-delay: 0.16s; }
.pb-d3 { animation-delay: 0.24s; }
@keyframes pb-in {
  from { opacity: 0; transform: translateY(12px); }
  to { opacity: 1; transform: translateY(0); }
}
@media (prefers-reduced-motion: reduce) {
  .pb-enter { animation: none; }
  .pb-primary, .pb-row { transition: none; }
}
</style>
