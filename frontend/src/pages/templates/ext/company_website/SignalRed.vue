<template>
  <div class="cw-sig min-h-screen bg-[#f2f1ed] text-[#111111]">
    <div class="cw-sig-top bg-[#111111] px-5 py-2.5 text-[#f2f1ed]">
      <div class="mx-auto flex w-full max-w-md items-center justify-between font-mono text-[8px] uppercase tracking-[0.2em]">
        <span class="flex items-center gap-2"><span class="h-1.5 w-1.5 rounded-full bg-[#e0271c]"></span>Company wayfinding</span>
        <span class="text-[#f2f1ed]/45">NFC · {{ nfcId || '01' }}</span>
      </div>
    </div>

    <div class="mx-auto w-full max-w-md px-5 pb-9 pt-6">
      <header class="cw-sig-in flex items-center justify-between border-b border-[#111]/25 pb-2 font-mono text-[8px] uppercase tracking-[0.16em] text-[#111]/50">
        <span>Profile / Official</span>
        <span>Section A · {{ pad2(features.length) }}</span>
      </header>

      <section class="cw-sig-in cw-sig-d1 mt-6">
        <div class="flex items-start justify-between gap-3">
          <div class="min-w-0 flex-1">
            <p class="font-mono text-[8.5px] uppercase tracking-[0.24em] text-[#e0271c]">You are here</p>
            <h1 class="mt-3 break-words text-[52px] font-black uppercase leading-[0.84] tracking-[-0.06em]">{{ name }}</h1>
          </div>
          <span class="cw-sig-round flex h-16 w-16 flex-shrink-0 items-center justify-center rounded-full bg-[#e0271c] text-[#fff5ef]">
            <ArrowDownRight class="h-7 w-7" />
          </span>
        </div>
        <p class="mt-5 max-w-[21rem] border-l-4 border-[#111] pl-3 text-[13px] leading-[1.75] text-[#111]/70">{{ description }}</p>
      </section>

      <div class="cw-sig-in cw-sig-d1 cw-sig-hazard mt-6"></div>

      <section class="cw-sig-in cw-sig-d2 mt-6">
        <div class="flex items-baseline justify-between">
          <h2 class="bg-[#111] px-2 py-1 font-mono text-[9px] font-bold uppercase tracking-[0.16em] text-[#f2f1ed]">Directions / 服务</h2>
          <span class="font-mono text-[8px] uppercase tracking-[0.12em] text-[#111]/40">Follow the signs</span>
        </div>
        <div v-if="features.length" class="mt-2 space-y-1.5">
          <div v-for="(f, i) in features" :key="i" class="cw-sig-sign group">
            <span class="cw-sig-num">{{ pad2(i + 1) }}</span>
            <span class="flex-1 text-[14px] font-bold uppercase leading-snug tracking-[0.01em]">{{ f }}</span>
            <ArrowRight class="h-4 w-4 text-[#111]/30 transition-all duration-200 group-hover:translate-x-1 group-hover:text-[#e0271c]" />
          </div>
        </div>
      </section>

      <section class="cw-sig-in cw-sig-d2 mt-6 grid grid-cols-3 bg-[#e0271c] text-[#fff5ef]">
        <div class="p-3">
          <p class="text-[22px] font-black leading-none">{{ pad2(features.length) }}</p>
          <p class="mt-1.5 font-mono text-[7px] uppercase tracking-[0.1em] text-[#fff5ef]/65">Stops</p>
        </div>
        <div class="border-l border-[#fff5ef]/25 p-3">
          <p class="text-[22px] font-black leading-none">24H</p>
          <p class="mt-1.5 font-mono text-[7px] uppercase tracking-[0.1em] text-[#fff5ef]/65">Reply</p>
        </div>
        <div class="border-l border-[#fff5ef]/25 p-3">
          <p class="text-[22px] font-black leading-none">1:1</p>
          <p class="mt-1.5 font-mono text-[7px] uppercase tracking-[0.1em] text-[#fff5ef]/65">Guide</p>
        </div>
      </section>

      <section class="cw-sig-in cw-sig-d3 mt-7">
        <div class="flex items-center gap-2">
          <h2 class="font-mono text-[9px] font-bold uppercase tracking-[0.16em]">Information desk</h2>
          <span class="h-0.5 flex-1 bg-[#111]"></span>
        </div>
        <div class="mt-2 border-2 border-[#111]">
          <a v-if="contact.phone" :href="'tel:' + contact.phone" class="cw-sig-contact">
            <span class="cw-sig-k"><Phone class="h-3 w-3" /></span>
            <span class="cw-sig-label">电话</span>
            <span class="cw-sig-v tabular-nums">{{ contact.phone }}</span>
            <ArrowUpRight class="h-3.5 w-3.5 text-[#111]/30" />
          </a>
          <a v-if="contact.email" :href="'mailto:' + contact.email" class="cw-sig-contact">
            <span class="cw-sig-k"><Mail class="h-3 w-3" /></span>
            <span class="cw-sig-label">邮箱</span>
            <span class="cw-sig-v">{{ contact.email }}</span>
            <ArrowUpRight class="h-3.5 w-3.5 text-[#111]/30" />
          </a>
          <div v-if="contact.address" class="cw-sig-contact items-start">
            <span class="cw-sig-k"><MapPin class="h-3 w-3" /></span>
            <span class="cw-sig-label">地址</span>
            <span class="cw-sig-v whitespace-normal">{{ contact.address }}</span>
          </div>
          <p v-if="!hasContact" class="py-5 text-center font-mono text-[9px] text-[#111]/40">INFORMATION NOT CONFIGURED</p>
        </div>
      </section>

      <div class="cw-sig-in cw-sig-d3 mt-6">
        <a v-if="isSafeUrl" :href="payload.url" target="_blank" rel="noopener noreferrer" class="cw-sig-cta">
          <span class="flex-1 text-center">访问官方网站</span>
          <span class="border-l border-[#fff5ef]/30 pl-4"><ArrowUpRight class="h-4 w-4" /></span>
        </a>
      </div>

      <div class="cw-sig-in cw-sig-d3 cw-sig-hazard mt-6"></div>

      <footer class="cw-sig-in cw-sig-d3 mt-3 flex items-center justify-between font-mono text-[7.5px] uppercase tracking-[0.14em] text-[#111]/40">
        <span>You have arrived · {{ name }}</span>
        <span>End of route</span>
      </footer>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { ArrowDownRight, ArrowRight, ArrowUpRight, Phone, Mail, MapPin } from 'lucide-vue-next'

const props = defineProps(['payload', 'nfcId'])

const name = computed(() => props.payload?.name || '企业官网')
const description = computed(() => props.payload?.description)
const features = computed(() => (Array.isArray(props.payload?.features) ? props.payload.features.filter(Boolean) : []))
const contact = computed(() => props.payload?.contactInfo || {})
const hasContact = computed(() => Boolean(contact.value.phone || contact.value.email || contact.value.address))
const isSafeUrl = computed(() => /^https?:\/\//i.test(props.payload?.url || ''))
const pad2 = (n) => String(n).padStart(2, '0')
</script>

<style scoped>
.cw-sig-hazard {
  height: 6px;
  background: repeating-linear-gradient(135deg, #111 0 8px, #e0271c 8px 16px);
}
.cw-sig-round {
  box-shadow: 0 0 0 4px #f2f1ed, 0 0 0 5px rgba(17, 17, 17, 0.18);
  transition: transform 0.25s ease;
}
.cw-sig-round:hover {
  transform: rotate(8deg);
}
.cw-sig-sign {
  display: flex;
  align-items: center;
  gap: 11px;
  border: 2px solid #111;
  background: #faf9f5;
  padding: 9px 10px;
  box-shadow: 3px 3px 0 rgba(17, 17, 17, 0.9);
  transition: transform 0.18s ease, box-shadow 0.18s ease, background-color 0.18s ease, color 0.18s ease;
}
.cw-sig-sign:hover {
  background: #111;
  color: #f2f1ed;
  transform: translate(-1px, -1px);
  box-shadow: 5px 5px 0 rgba(224, 39, 28, 0.95);
}
.cw-sig-sign:active {
  transform: translate(2px, 2px);
  box-shadow: 1px 1px 0 rgba(17, 17, 17, 0.9);
}
.cw-sig-num {
  display: flex;
  height: 25px;
  width: 25px;
  flex-shrink: 0;
  align-items: center;
  justify-content: center;
  background: #e0271c;
  font-family: ui-monospace, SFMono-Regular, Menlo, monospace;
  font-size: 9px;
  font-weight: 700;
  color: #fff5ef;
}
.cw-sig-contact {
  display: flex;
  align-items: center;
  gap: 10px;
  border-bottom: 1px solid rgba(17, 17, 17, 0.18);
  padding: 10px 11px;
  transition: background-color 0.18s ease;
}
.cw-sig-contact:last-child {
  border-bottom: 0;
}
a.cw-sig-contact:hover {
  background: rgba(224, 39, 28, 0.07);
}
.cw-sig-k {
  display: flex;
  height: 24px;
  width: 24px;
  flex-shrink: 0;
  align-items: center;
  justify-content: center;
  border: 1.5px solid #111;
  color: #e0271c;
}
.cw-sig-label {
  width: 30px;
  flex-shrink: 0;
  font-family: ui-monospace, SFMono-Regular, Menlo, monospace;
  font-size: 8px;
  font-weight: 700;
  color: rgba(17, 17, 17, 0.45);
}
.cw-sig-v {
  flex: 1;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  font-size: 12px;
  font-weight: 600;
}
a.cw-sig-contact:hover .cw-sig-v {
  color: #e0271c;
}
.cw-sig-cta {
  display: flex;
  align-items: center;
  gap: 10px;
  background: #e0271c;
  padding: 12px 14px;
  font-size: 14px;
  font-weight: 900;
  letter-spacing: 0.04em;
  color: #fff5ef;
  transition: background-color 0.18s ease, transform 0.18s ease;
}
.cw-sig-cta:hover {
  background: #111111;
}
.cw-sig-cta:active {
  transform: scale(0.99);
}
.cw-sig-in {
  animation: cw-sig-in 0.55s cubic-bezier(0.16, 0.9, 0.2, 1) both;
}
.cw-sig-d1 { animation-delay: 0.05s; }
.cw-sig-d2 { animation-delay: 0.13s; }
.cw-sig-d3 { animation-delay: 0.21s; }
@keyframes cw-sig-in {
  from { opacity: 0; transform: translateX(-14px); }
  to { opacity: 1; transform: none; }
}
@media (prefers-reduced-motion: reduce) {
  .cw-sig-in {
    animation: none !important;
  }
  .cw-sig-sign,
  .cw-sig-contact,
  .cw-sig-logo,
  .cw-sig-round,
  .cw-sig-cta {
    transition: none !important;
  }
}
</style>
