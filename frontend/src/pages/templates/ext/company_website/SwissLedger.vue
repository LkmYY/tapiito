<template>
  <div class="cw-sw min-h-screen bg-[#f2efe6] text-[#151515]">
    <div class="mx-auto w-full max-w-md px-6 pb-10 pt-7">
      <header class="cw-sw-in flex items-center justify-between border-b-2 border-[#151515] pb-3">
        <span class="text-[11px] font-semibold uppercase tracking-[0.14em]">Corporate Ledger</span>
        <span class="font-mono text-[10px] tabular-nums text-[#151515]/50">No. {{ nfcId || '000' }}</span>
      </header>

      <section class="cw-sw-in cw-sw-d1 pt-9">
        <div class="flex items-start justify-between gap-4">
          <p class="font-mono text-[10px] uppercase tracking-[0.2em] text-[#151515]/45">Official company record</p>
          <span class="mt-1 h-2.5 w-2.5 flex-shrink-0 rounded-full bg-[#e03131]"></span>
        </div>
        <h1 class="mt-8 text-[50px] font-semibold leading-[0.92] tracking-[-0.055em]">{{ name }}</h1>
        <p class="mt-7 max-w-[21rem] text-[14px] leading-[1.8] text-[#151515]/65">{{ description }}</p>
      </section>

      <section class="cw-sw-in cw-sw-d2 mt-10 border-t-2 border-[#151515]">
        <div class="flex items-center justify-between py-3">
          <h2 class="text-[11px] font-semibold uppercase tracking-[0.14em]">Services</h2>
          <span class="font-mono text-[10px] tabular-nums text-[#151515]/45">{{ pad2(features.length) }} / 04</span>
        </div>
        <div v-if="features.length">
          <div v-for="(f, i) in features" :key="i" class="cw-sw-row group">
            <span class="w-10 font-mono text-[11px] tabular-nums text-[#e03131]">{{ pad2(i + 1) }}</span>
            <span class="flex-1 text-[15px] font-medium leading-snug transition-transform duration-200 group-hover:translate-x-1">{{ f }}</span>
            <span class="font-mono text-[16px] font-light text-[#151515]/30 transition-colors group-hover:text-[#e03131]">↗</span>
          </div>
        </div>
      </section>

      <section class="cw-sw-in cw-sw-d3 mt-10">
        <div class="grid grid-cols-3 border-y border-[#151515]/20">
          <div class="py-4">
            <p class="font-mono text-[9px] uppercase tracking-[0.16em] text-[#151515]/40">Scope</p>
            <p class="mt-1 text-[13px] font-semibold">全域服务</p>
          </div>
          <div class="border-x border-[#151515]/20 px-4 py-4">
            <p class="font-mono text-[9px] uppercase tracking-[0.16em] text-[#151515]/40">Reply</p>
            <p class="mt-1 text-[13px] font-semibold">24 小时内</p>
          </div>
          <div class="py-4 pl-4">
            <p class="font-mono text-[9px] uppercase tracking-[0.16em] text-[#151515]/40">Mode</p>
            <p class="mt-1 text-[13px] font-semibold">专属对接</p>
          </div>
        </div>
      </section>

      <section class="cw-sw-in cw-sw-d3 mt-9">
        <div class="flex items-center justify-between border-b border-[#151515] pb-2">
          <h2 class="text-[11px] font-semibold uppercase tracking-[0.14em]">Contact Ledger</h2>
          <span class="font-mono text-[10px] text-[#151515]/40">CN / 01</span>
        </div>
        <dl class="mt-1">
          <a v-if="contact.phone" :href="'tel:' + contact.phone" class="cw-sw-line group">
            <dt>电话</dt>
            <dd class="tabular-nums">{{ contact.phone }}</dd>
            <span class="cw-sw-arrow">→</span>
          </a>
          <a v-if="contact.email" :href="'mailto:' + contact.email" class="cw-sw-line group">
            <dt>邮箱</dt>
            <dd>{{ contact.email }}</dd>
            <span class="cw-sw-arrow">→</span>
          </a>
          <div v-if="contact.address" class="cw-sw-line">
            <dt>地址</dt>
            <dd class="whitespace-normal text-right">{{ contact.address }}</dd>
          </div>
          <p v-if="!hasContact" class="py-6 text-[12px] text-[#151515]/40">联系方式未配置</p>
        </dl>
      </section>

      <div class="cw-sw-in cw-sw-d3 mt-8">
        <a v-if="isSafeUrl" :href="payload.url" target="_blank" rel="noopener noreferrer" class="cw-sw-cta">
          <span class="font-mono text-[11px] tracking-[0.1em]">WWW</span>
          <span class="flex-1">访问官方网站</span>
          <span class="text-[16px]">↗</span>
        </a>
      </div>

      <footer class="mt-8 flex items-center justify-between border-t border-[#151515]/15 pt-3 font-mono text-[9px] uppercase tracking-[0.16em] text-[#151515]/35">
        <span>Swiss Standard · 2026</span>
        <span>Page 01</span>
      </footer>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'

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
.cw-sw-row {
  display: flex;
  align-items: center;
  gap: 12px;
  border-bottom: 1px solid rgba(21, 21, 21, 0.18);
  padding: 17px 0;
  transition: background-color 0.2s ease;
}
.cw-sw-row:hover {
  background: rgba(224, 49, 49, 0.045);
}
.cw-sw-line {
  display: flex;
  align-items: baseline;
  gap: 12px;
  border-bottom: 1px solid rgba(21, 21, 21, 0.12);
  padding: 12px 0;
}
.cw-sw-line dt {
  width: 38px;
  flex-shrink: 0;
  font-size: 10px;
  letter-spacing: 0.1em;
  color: rgba(21, 21, 21, 0.42);
}
.cw-sw-line dd {
  flex: 1;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  font-size: 13px;
  font-weight: 500;
}
.cw-sw-line .cw-sw-arrow {
  color: rgba(21, 21, 21, 0.25);
  transition: color 0.2s ease, transform 0.2s ease;
}
a.cw-sw-line:hover dd {
  color: #e03131;
}
a.cw-sw-line:hover .cw-sw-arrow {
  color: #e03131;
  transform: translateX(3px);
}
.cw-sw-cta {
  display: flex;
  align-items: center;
  gap: 12px;
  background: #151515;
  padding: 14px 16px;
  color: #f2efe6;
  transition: background-color 0.2s ease, transform 0.2s ease;
}
.cw-sw-cta:hover {
  background: #e03131;
}
.cw-sw-cta:active {
  transform: scale(0.985);
}
.cw-sw-in {
  animation: cw-sw-in 0.62s cubic-bezier(0.2, 0.7, 0.2, 1) both;
}
.cw-sw-d1 { animation-delay: 0.08s; }
.cw-sw-d2 { animation-delay: 0.16s; }
.cw-sw-d3 { animation-delay: 0.24s; }
@keyframes cw-sw-in {
  from { opacity: 0; transform: translateY(18px); }
  to { opacity: 1; transform: none; }
}
@media (prefers-reduced-motion: reduce) {
  .cw-sw-in {
    animation: none !important;
  }
  .cw-sw-row,
  .cw-sw-line dd,
  .cw-sw-arrow,
  .cw-sw-cta {
    transition: none !important;
  }
}
</style>
