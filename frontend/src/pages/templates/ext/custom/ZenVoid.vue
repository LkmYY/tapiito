<template>
  <div class="cu-zv-root min-h-screen">
    <div class="cu-zv-wrap mx-auto w-full max-w-md">
      <header class="cu-zv-top cu-zv-in">
        <span class="cu-zv-mark">無</span>
        <span class="cu-zv-id">{{ nfcId ? 'NFC / ' + nfcId : 'NFC / —' }}</span>
      </header>

      <svg class="cu-zv-enso cu-zv-in cu-zv-d1" viewBox="0 0 120 120" fill="none" aria-hidden="true">
        <path
          d="M92 24c13 13 20 30 17 47-4 21-24 36-46 34-21-2-37-19-37-40 0-22 17-40 39-41 10-1 20 3 27 10"
          stroke="currentColor"
          stroke-width="2.2"
          stroke-linecap="round"
        />
      </svg>

      <section class="cu-zv-title">
        <p class="cu-zv-eyebrow cu-zv-in cu-zv-d1">企 业 简 介</p>
        <h1 class="cu-zv-name cu-zv-in cu-zv-d2">{{ payload?.companyName || '企业主页' }}</h1>
        <p class="cu-zv-desc cu-zv-in cu-zv-d3">{{ payload?.description }}</p>
      </section>

      <div class="cu-zv-thread cu-zv-in cu-zv-d3">
        <span></span>
        <i>— 静 候 咨 询 —</i>
      </div>

      <section class="cu-zv-services cu-zv-in cu-zv-d4">
        <p class="cu-zv-label">服务 / SERVICES</p>
        <div v-if="services.length" class="cu-zv-list">
          <div v-for="(service, i) in services" :key="i" class="cu-zv-item">
            <span class="cu-zv-item-dot"></span>
            <span class="cu-zv-item-name">{{ service.name }}</span>
            <span class="cu-zv-item-no">{{ String(i + 1).padStart(2, '0') }}</span>
          </div>
        </div>
      </section>

      <section class="cu-zv-contact cu-zv-in cu-zv-d5">
        <p class="cu-zv-label">联络 / CONTACT</p>
        <a v-if="payload?.phone" :href="'tel:' + payload.phone" class="cu-zv-phone">
          <span>{{ payload.phone }}</span>
          <ArrowUpRight :size="14" />
        </a>
        <a v-if="payload?.email" :href="'mailto:' + payload.email" class="cu-zv-mail">{{ payload.email }}</a>
      </section>

      <footer class="cu-zv-foot cu-zv-in cu-zv-d5">
        <span class="cu-zv-foot-line"></span>
        <span>{{ payload?.companyName || '企业主页' }}</span>
        <span class="cu-zv-foot-dot">·</span>
        <span>{{ nfcId ? String(nfcId) : '—' }}</span>
      </footer>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { ArrowUpRight } from 'lucide-vue-next'

const props = defineProps(['payload', 'nfcId'])

const services = computed(() => {
  const list = props.payload?.services
  return Array.isArray(list) ? list.filter((item) => item && item.name) : []
})
</script>

<style scoped>
.cu-zv-root {
  background: #f7f5ef;
  color: #26241f;
  font-family: 'Songti SC', 'Noto Serif SC', Georgia, serif;
  overflow-x: hidden;
}
.cu-zv-wrap {
  position: relative;
  padding: 30px 30px 40px;
  min-height: 100vh;
}
.cu-zv-top {
  display: flex;
  align-items: center;
  justify-content: space-between;
}
.cu-zv-mark {
  height: 27px;
  width: 27px;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  border: 1px solid rgba(38, 36, 31, 0.5);
  font-size: 13px;
  color: #8b2f2f;
}
.cu-zv-id {
  font-family: 'Courier New', monospace;
  font-size: 7px;
  letter-spacing: 0.22em;
  color: #a09a8d;
}
.cu-zv-enso {
  position: absolute;
  right: 22px;
  top: 82px;
  width: 126px;
  color: rgba(139, 47, 47, 0.28);
  pointer-events: none;
}
.cu-zv-title {
  margin-top: 78px;
}
.cu-zv-eyebrow {
  font-size: 10px;
  letter-spacing: 0.5em;
  color: #9c9484;
}
.cu-zv-name {
  position: relative;
  margin-top: 16px;
  max-width: 6em;
  font-size: 40px;
  line-height: 1.24;
  font-weight: 400;
  letter-spacing: 0.08em;
  color: #1f1d19;
  word-break: break-word;
}
.cu-zv-desc {
  margin-top: 26px;
  max-width: 23ch;
  font-size: 12px;
  line-height: 2.2;
  letter-spacing: 0.04em;
  color: #6f6a5f;
}
.cu-zv-thread {
  margin-top: 34px;
  display: flex;
  align-items: center;
  gap: 12px;
}
.cu-zv-thread span {
  height: 1px;
  width: 42px;
  background: rgba(38, 36, 31, 0.25);
}
.cu-zv-thread i {
  font-style: normal;
  font-size: 9px;
  letter-spacing: 0.24em;
  color: #a09a8d;
}
.cu-zv-services {
  margin-top: 54px;
}
.cu-zv-label {
  font-family: 'Courier New', monospace;
  font-size: 7px;
  letter-spacing: 0.24em;
  color: #a09a8d;
}
.cu-zv-list {
  margin-top: 13px;
}
.cu-zv-item {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 17px 0;
  border-bottom: 1px solid rgba(38, 36, 31, 0.1);
  transition: padding-left 0.3s ease;
}
.cu-zv-item:hover {
  padding-left: 5px;
}
.cu-zv-item-dot {
  height: 3px;
  width: 3px;
  border-radius: 50%;
  background: #8b2f2f;
  flex-shrink: 0;
}
.cu-zv-item-name {
  flex: 1;
  font-size: 15px;
  letter-spacing: 0.09em;
}
.cu-zv-item-no {
  font-family: 'Courier New', monospace;
  font-size: 7px;
  letter-spacing: 0.12em;
  color: #b0aa9d;
}
.cu-zv-contact {
  margin-top: 46px;
}
.cu-zv-phone {
  margin-top: 12px;
  display: inline-flex;
  align-items: center;
  gap: 9px;
  font-size: 19px;
  letter-spacing: 0.08em;
  color: #26241f;
  border-bottom: 1px solid rgba(139, 47, 47, 0.45);
  padding-bottom: 5px;
  transition: color 0.25s ease;
}
.cu-zv-phone:hover {
  color: #8b2f2f;
}
.cu-zv-mail {
  display: block;
  margin-top: 10px;
  font-family: 'Courier New', monospace;
  font-size: 9px;
  letter-spacing: 0.08em;
  color: #8c8679;
  word-break: break-all;
}
.cu-zv-muted {
  margin-top: 12px;
  font-size: 11px;
  letter-spacing: 0.08em;
  color: #a09a8d;
}
.cu-zv-foot {
  margin-top: 66px;
  display: flex;
  align-items: center;
  gap: 8px;
  font-family: 'Courier New', monospace;
  font-size: 7px;
  letter-spacing: 0.14em;
  color: #a09a8d;
}
.cu-zv-foot-line {
  height: 1px;
  width: 20px;
  background: rgba(38, 36, 31, 0.25);
}
.cu-zv-foot-dot {
  color: #c0baad;
}
@keyframes cu-zv-in {
  from {
    opacity: 0;
    transform: translateY(10px);
  }
  to {
    opacity: 1;
    transform: none;
  }
}
.cu-zv-in {
  animation: cu-zv-in 1s cubic-bezier(0.22, 1, 0.36, 1) both;
}
.cu-zv-d1 {
  animation-delay: 0.12s;
}
.cu-zv-d2 {
  animation-delay: 0.22s;
}
.cu-zv-d3 {
  animation-delay: 0.32s;
}
.cu-zv-d4 {
  animation-delay: 0.44s;
}
.cu-zv-d5 {
  animation-delay: 0.56s;
}
@media (prefers-reduced-motion: reduce) {
  .cu-zv-in {
    animation: none;
  }
  .cu-zv-item,
  .cu-zv-phone {
    transition: none;
  }
}
</style>
