<template>
  <div class="cu-rw-root min-h-screen">
    <div class="cu-rw-wrap mx-auto w-full max-w-md">
      <header class="cu-rw-top cu-rw-in">
        <span class="cu-rw-brand">SLOW ROAST</span>
        <span class="cu-rw-date">2026 · 秋</span>
      </header>

      <section class="cu-rw-hero">
        <div class="cu-rw-copy cu-rw-in cu-rw-d1">
          <p class="cu-rw-eyebrow">NEIGHBOURHOOD ROASTERY</p>
          <h1>{{ payload?.companyName || '企业主页' }}</h1>
          <p class="cu-rw-desc">{{ payload?.description }}</p>
        </div>
        <div class="cu-rw-stamp cu-rw-in cu-rw-d2">
          <span class="cu-rw-stamp-top">FRESH</span>
          <span class="cu-rw-stamp-initial">{{ initial }}</span>
          <span class="cu-rw-stamp-bottom">SINCE 2026</span>
        </div>
      </section>

      <div class="cu-rw-stain" aria-hidden="true"></div>

      <section class="cu-rw-board cu-rw-in cu-rw-d3">
        <div class="cu-rw-board-head">
          <span>TODAY'S OFFER</span>
          <span>今日供应</span>
        </div>
        <div v-if="services.length" class="cu-rw-menu">
          <div v-for="(service, i) in services" :key="i" class="cu-rw-line">
            <span class="cu-rw-line-no">{{ String(i + 1).padStart(2, '0') }}</span>
            <span class="cu-rw-line-name">{{ service.name }}</span>
            <span class="cu-rw-dots"></span>
            <span class="cu-rw-line-mark">●</span>
          </div>
        </div>
        <p class="cu-rw-note">当日现做 · 售完即止 · 支持自取</p>
      </section>

      <div v-if="images.length" class="cu-rw-photo cu-rw-in cu-rw-d3">
        <img :src="images[0]" :alt="payload?.companyName || '店内照片'" />
        <span>OUR COUNTER / 店内一角</span>
      </div>

      <section class="cu-rw-contact cu-rw-in cu-rw-d4">
        <div class="cu-rw-ticket">
          <span class="cu-rw-ticket-label">CALL THE BARISTA</span>
          <a v-if="payload?.phone" :href="'tel:' + payload.phone" class="cu-rw-phone">
            <Phone :size="16" />
            {{ payload.phone }}
          </a>
          <a v-if="payload?.email" :href="'mailto:' + payload.email" class="cu-rw-mail">{{ payload.email }}</a>
        </div>
      </section>

      <footer class="cu-rw-foot cu-rw-in cu-rw-d4">
        <span>{{ payload?.companyName || '企业主页' }}</span>
        <span>{{ nfcId ? 'TAG ' + nfcId : 'TAG —' }}</span>
      </footer>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { Phone } from 'lucide-vue-next'

const props = defineProps(['payload', 'nfcId'])

const services = computed(() => {
  const list = props.payload?.services
  return Array.isArray(list) ? list.filter((item) => item && item.name) : []
})
const images = computed(() => {
  const list = props.payload?.images
  return Array.isArray(list) ? list.filter((src) => typeof src === 'string' && src.trim()) : []
})
const initial = computed(() => (props.payload?.companyName || '未').trim().slice(0, 1))
</script>

<style scoped>
.cu-rw-root {
  background: #e9dfcc;
  color: #32241a;
  font-family: 'PingFang SC', 'Microsoft YaHei', sans-serif;
  overflow-x: hidden;
}
.cu-rw-wrap {
  position: relative;
  padding: 22px 20px 34px;
}
.cu-rw-top {
  display: flex;
  justify-content: space-between;
  border-bottom: 1px solid rgba(50, 36, 26, 0.3);
  padding-bottom: 9px;
  font-family: 'Courier New', monospace;
  font-size: 8px;
  letter-spacing: 0.24em;
}
.cu-rw-brand {
  font-weight: 700;
  letter-spacing: 0.3em;
}
.cu-rw-date {
  color: #7d6a56;
}
.cu-rw-hero {
  display: flex;
  align-items: flex-start;
  gap: 8px;
  margin-top: 26px;
}
.cu-rw-copy {
  flex: 1;
  min-width: 0;
}
.cu-rw-eyebrow {
  font-family: 'Courier New', monospace;
  font-size: 8px;
  letter-spacing: 0.2em;
  color: #9b5b34;
}
.cu-rw-copy h1 {
  margin-top: 9px;
  font-family: 'Songti SC', 'Noto Serif SC', Georgia, serif;
  font-size: 40px;
  line-height: 1.05;
  font-weight: 700;
  letter-spacing: -0.02em;
  word-break: break-word;
}
.cu-rw-desc {
  margin-top: 12px;
  max-width: 26ch;
  font-size: 12px;
  line-height: 1.8;
  color: #6f5a45;
}
.cu-rw-stamp {
  flex-shrink: 0;
  height: 92px;
  width: 92px;
  margin-top: 4px;
  border: 1.5px dashed #9b5b34;
  border-radius: 50%;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  transform: rotate(9deg);
  color: #9b5b34;
  background: rgba(233, 223, 204, 0.5);
}
.cu-rw-stamp-top,
.cu-rw-stamp-bottom {
  font-family: 'Courier New', monospace;
  font-size: 6px;
  letter-spacing: 0.16em;
}
.cu-rw-stamp-initial {
  font-family: Georgia, 'Songti SC', serif;
  font-size: 31px;
  line-height: 1.1;
  margin: 1px 0;
}
.cu-rw-stain {
  position: relative;
  height: 1px;
  margin-top: 20px;
}
.cu-rw-stain::before {
  content: '';
  position: absolute;
  right: -6px;
  top: -12px;
  height: 64px;
  width: 64px;
  border-radius: 50%;
  border: 7px solid rgba(151, 96, 54, 0.08);
  box-shadow: inset 0 0 0 3px rgba(151, 96, 54, 0.05);
}
.cu-rw-board {
  position: relative;
  margin-top: 18px;
  border: 1px solid rgba(50, 36, 26, 0.38);
  background: #f4ecdd;
  padding: 16px 16px 12px;
  box-shadow: 5px 6px 0 rgba(50, 36, 26, 0.1);
}
.cu-rw-board::before {
  content: '';
  position: absolute;
  top: -6px;
  left: 24px;
  height: 11px;
  width: 54px;
  background: rgba(155, 91, 52, 0.55);
  transform: rotate(-2deg);
}
.cu-rw-board-head {
  display: flex;
  justify-content: space-between;
  border-bottom: 2px solid #32241a;
  padding-bottom: 7px;
  font-family: 'Courier New', monospace;
  font-size: 8px;
  letter-spacing: 0.2em;
}
.cu-rw-board-head span:last-child {
  font-family: 'PingFang SC', sans-serif;
  letter-spacing: 0.12em;
}
.cu-rw-menu {
  padding-top: 4px;
}
.cu-rw-line {
  display: flex;
  align-items: baseline;
  gap: 8px;
  padding: 13px 0;
  border-bottom: 1px dashed rgba(50, 36, 26, 0.26);
}
.cu-rw-line-no {
  font-family: 'Courier New', monospace;
  font-size: 8px;
  color: #9b5b34;
}
.cu-rw-line-name {
  font-size: 15px;
  font-weight: 600;
  letter-spacing: 0.01em;
}
.cu-rw-dots {
  flex: 1;
  min-width: 12px;
  border-bottom: 1px dotted rgba(50, 36, 26, 0.35);
  transform: translateY(-3px);
}
.cu-rw-line-mark {
  font-size: 6px;
  color: #9b5b34;
}
.cu-rw-note {
  margin-top: 10px;
  text-align: center;
  font-family: 'Courier New', monospace;
  font-size: 7px;
  letter-spacing: 0.14em;
  color: #8a7259;
}
.cu-rw-photo {
  margin-top: 18px;
  position: relative;
  border: 1px solid rgba(50, 36, 26, 0.3);
  padding: 6px;
  background: #f4ecdd;
}
.cu-rw-photo img {
  width: 100%;
  aspect-ratio: 4 / 3;
  object-fit: cover;
  display: block;
  filter: sepia(0.18) contrast(1.03);
}
.cu-rw-photo span {
  display: block;
  padding-top: 6px;
  font-family: 'Courier New', monospace;
  font-size: 7px;
  letter-spacing: 0.16em;
  color: #7d6a56;
}
.cu-rw-contact {
  margin-top: 20px;
}
.cu-rw-ticket {
  position: relative;
  border: 1px solid rgba(50, 36, 26, 0.4);
  background: #9b5b34;
  color: #f8f1e3;
  padding: 16px;
}
.cu-rw-ticket::before,
.cu-rw-ticket::after {
  content: '';
  position: absolute;
  top: 50%;
  height: 14px;
  width: 14px;
  border-radius: 50%;
  background: #e9dfcc;
  transform: translateY(-50%);
}
.cu-rw-ticket::before {
  left: -7px;
}
.cu-rw-ticket::after {
  right: -7px;
}
.cu-rw-ticket-label {
  font-family: 'Courier New', monospace;
  font-size: 8px;
  letter-spacing: 0.22em;
  opacity: 0.8;
}
.cu-rw-phone {
  margin-top: 8px;
  display: inline-flex;
  align-items: center;
  gap: 8px;
  font-size: 21px;
  font-weight: 600;
  letter-spacing: 0.02em;
  color: #fffaf0;
  border-bottom: 1px solid rgba(255, 250, 240, 0.5);
  padding-bottom: 3px;
  transition: opacity 0.2s ease;
}
.cu-rw-phone:hover {
  opacity: 0.78;
}
.cu-rw-mail {
  display: block;
  margin-top: 9px;
  font-family: 'Courier New', monospace;
  font-size: 10px;
  color: #f0dcc7;
  word-break: break-all;
}
.cu-rw-empty {
  padding: 10px 0;
  font-size: 11px;
  color: #8a7259;
}
.cu-rw-foot {
  margin-top: 20px;
  display: flex;
  justify-content: space-between;
  font-family: 'Courier New', monospace;
  font-size: 8px;
  letter-spacing: 0.14em;
  color: #7d6a56;
}
@keyframes cu-rw-in {
  from {
    opacity: 0;
    transform: translateY(16px);
  }
  to {
    opacity: 1;
    transform: none;
  }
}
.cu-rw-in {
  animation: cu-rw-in 0.7s cubic-bezier(0.22, 1, 0.36, 1) both;
}
.cu-rw-d1 {
  animation-delay: 0.08s;
}
.cu-rw-d2 {
  animation-delay: 0.18s;
}
.cu-rw-d3 {
  animation-delay: 0.26s;
}
.cu-rw-d4 {
  animation-delay: 0.34s;
}
@media (prefers-reduced-motion: reduce) {
  .cu-rw-in {
    animation: none;
  }
  .cu-rw-phone {
    transition: none;
  }
}
</style>
