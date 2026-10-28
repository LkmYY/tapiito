<template>
  <div class="cu-sw-root min-h-screen">
    <div class="cu-sw-wrap mx-auto w-full max-w-md">
      <header class="cu-sw-top cu-sw-in">
        <span class="cu-sw-mark"><span></span></span>
        <span class="cu-sw-code">{{ nfcId ? 'CARD/' + nfcId : 'CARD/—' }}</span>
      </header>

      <section class="cu-sw-headline">
        <p class="cu-sw-eyebrow cu-sw-in cu-sw-d1">Unternehmen · 企业服务</p>
        <h1 class="cu-sw-title cu-sw-in cu-sw-d2">{{ payload?.companyName || '企业主页' }}</h1>
        <p class="cu-sw-lead cu-sw-in cu-sw-d3">{{ payload?.description }}</p>
      </section>

      <div class="cu-sw-meta cu-sw-in cu-sw-d3">
        <div>
          <p class="cu-sw-label">Scope</p>
          <p class="cu-sw-value">Brand / Space / Service</p>
        </div>
        <div>
          <p class="cu-sw-label">Based</p>
          <p class="cu-sw-value">Local Practice</p>
        </div>
      </div>

      <section class="cu-sw-services cu-sw-in cu-sw-d4">
        <div class="cu-sw-section-bar">
          <span>Leistungen</span>
          <span>服务项目 · {{ String(services.length).padStart(2, '0') }}</span>
        </div>
        <div v-if="services.length" class="cu-sw-grid">
          <article v-for="(service, i) in services" :key="i" class="cu-sw-cell">
            <span class="cu-sw-cell-no">{{ String(i + 1).padStart(2, '0') }}</span>
            <span class="cu-sw-cell-dot"></span>
            <h2>{{ service.name }}</h2>
            <span class="cu-sw-cell-rule"></span>
          </article>
        </div>
      </section>

      <section class="cu-sw-contact cu-sw-in cu-sw-d5">
        <div class="cu-sw-contact-copy">
          <p class="cu-sw-label">Kontakt</p>
          <p class="cu-sw-contact-cn">项目洽谈与预约</p>
        </div>
        <a v-if="payload?.phone" :href="'tel:' + payload.phone" class="cu-sw-tel">
          <Phone :size="16" />
          <span>{{ payload.phone }}</span>
          <ArrowUpRight :size="15" />
        </a>
      </section>

      <footer class="cu-sw-foot cu-sw-in cu-sw-d5">
        <span>{{ payload?.companyName || '企业主页' }}</span>
        <span>SWISS FORM / 2026</span>
      </footer>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { ArrowUpRight, Phone } from 'lucide-vue-next'

const props = defineProps(['payload', 'nfcId'])

const services = computed(() => {
  const list = props.payload?.services
  return Array.isArray(list) ? list.filter((item) => item && item.name) : []
})
</script>

<style scoped>
.cu-sw-root {
  position: relative;
  overflow: hidden;
  background: #f6f5f1;
  color: #111111;
  font-family: 'Helvetica Neue', Arial, 'PingFang SC', sans-serif;
}
.cu-sw-root::before {
  content: '';
  position: absolute;
  inset: 0;
  pointer-events: none;
  background-image: linear-gradient(to right, rgba(17, 17, 17, 0.045) 1px, transparent 1px);
  background-size: 25% 100%;
}
.cu-sw-wrap {
  position: relative;
  padding: 24px 22px 34px;
}
.cu-sw-top {
  display: flex;
  align-items: center;
  justify-content: space-between;
  border-bottom: 2px solid #111;
  padding-bottom: 12px;
}
.cu-sw-mark {
  display: inline-flex;
  height: 21px;
  width: 21px;
  align-items: center;
  justify-content: center;
  background: #e2231a;
}
.cu-sw-mark span {
  height: 7px;
  width: 7px;
  background: #f6f5f1;
}
.cu-sw-code {
  font-family: 'Courier New', monospace;
  font-size: 9px;
  letter-spacing: 0.2em;
  text-transform: uppercase;
}
.cu-sw-headline {
  padding-top: 26px;
}
.cu-sw-eyebrow {
  font-size: 9px;
  font-weight: 600;
  letter-spacing: 0.26em;
  text-transform: uppercase;
  color: #e2231a;
}
.cu-sw-title {
  margin-top: 14px;
  font-size: 47px;
  line-height: 0.98;
  font-weight: 700;
  letter-spacing: -0.055em;
  word-break: break-word;
}
.cu-sw-lead {
  margin-top: 18px;
  max-width: 34ch;
  font-size: 13px;
  line-height: 1.7;
  color: #555555;
}
.cu-sw-meta {
  margin-top: 24px;
  display: grid;
  grid-template-columns: 1fr 1fr;
  border-top: 1px solid #111;
  border-bottom: 1px solid #111;
  padding: 12px 0;
  gap: 10px;
}
.cu-sw-meta > div + div {
  border-left: 1px solid #cfcdc6;
  padding-left: 12px;
}
.cu-sw-label {
  font-family: 'Courier New', monospace;
  font-size: 8px;
  letter-spacing: 0.2em;
  text-transform: uppercase;
  color: #8a8880;
}
.cu-sw-value {
  margin-top: 5px;
  font-size: 11px;
  font-weight: 600;
  letter-spacing: 0.02em;
}
.cu-sw-services {
  margin-top: 30px;
}
.cu-sw-section-bar {
  display: flex;
  align-items: baseline;
  justify-content: space-between;
  font-size: 10px;
  font-weight: 700;
  letter-spacing: 0.12em;
  text-transform: uppercase;
}
.cu-sw-section-bar span:last-child {
  font-family: 'Courier New', monospace;
  font-size: 8px;
  font-weight: 400;
  letter-spacing: 0.16em;
  color: #8a8880;
}
.cu-sw-grid {
  margin-top: 10px;
  display: grid;
  grid-template-columns: 1fr 1fr;
  border-left: 1px solid #111;
  border-top: 1px solid #111;
}
.cu-sw-cell {
  position: relative;
  min-height: 116px;
  padding: 14px 13px 13px;
  border-right: 1px solid #111;
  border-bottom: 1px solid #111;
  background: #f6f5f1;
  display: flex;
  flex-direction: column;
  transition: background-color 0.22s ease, color 0.22s ease;
}
.cu-sw-cell:hover {
  background: #111;
  color: #f6f5f1;
}
.cu-sw-cell-no {
  font-family: 'Courier New', monospace;
  font-size: 9px;
  color: #e2231a;
}
.cu-sw-cell-dot {
  margin-top: 20px;
  height: 7px;
  width: 7px;
  border-radius: 50%;
  background: #e2231a;
}
.cu-sw-cell h2 {
  margin-top: 10px;
  font-size: 16px;
  font-weight: 600;
  line-height: 1.3;
  letter-spacing: -0.01em;
}
.cu-sw-cell-rule {
  margin-top: auto;
  height: 1px;
  width: 100%;
  background: currentColor;
  opacity: 0.22;
}
.cu-sw-empty {
  padding: 18px;
  border: 1px solid #d6d4cd;
  font-size: 12px;
  color: #77756e;
}
.cu-sw-contact {
  margin-top: 28px;
  background: #e2231a;
  color: #ffffff;
  padding: 18px;
}
.cu-sw-contact-copy {
  display: flex;
  align-items: baseline;
  justify-content: space-between;
}
.cu-sw-contact .cu-sw-label {
  color: rgba(255, 255, 255, 0.72);
}
.cu-sw-contact-cn {
  font-size: 10px;
  letter-spacing: 0.1em;
}
.cu-sw-tel {
  margin-top: 16px;
  display: flex;
  align-items: center;
  gap: 10px;
  font-size: 22px;
  font-weight: 600;
  letter-spacing: -0.02em;
  color: #ffffff;
  transition: opacity 0.2s ease;
}
.cu-sw-tel span {
  flex: 1;
}
.cu-sw-tel:hover {
  opacity: 0.8;
}
.cu-sw-empty--light {
  border-color: rgba(255, 255, 255, 0.4);
  color: rgba(255, 255, 255, 0.8);
}
.cu-sw-foot {
  margin-top: 22px;
  display: flex;
  justify-content: space-between;
  font-family: 'Courier New', monospace;
  font-size: 8px;
  letter-spacing: 0.16em;
  text-transform: uppercase;
  color: #85837c;
}
@keyframes cu-sw-in {
  from {
    opacity: 0;
    transform: translateY(16px);
  }
  to {
    opacity: 1;
    transform: none;
  }
}
.cu-sw-in {
  animation: cu-sw-in 0.6s cubic-bezier(0.2, 0.8, 0.2, 1) both;
}
.cu-sw-d1 {
  animation-delay: 0.06s;
}
.cu-sw-d2 {
  animation-delay: 0.12s;
}
.cu-sw-d3 {
  animation-delay: 0.2s;
}
.cu-sw-d4 {
  animation-delay: 0.28s;
}
.cu-sw-d5 {
  animation-delay: 0.36s;
}
@media (prefers-reduced-motion: reduce) {
  .cu-sw-in {
    animation: none;
  }
  .cu-sw-cell,
  .cu-sw-tel {
    transition: none;
  }
}
</style>
