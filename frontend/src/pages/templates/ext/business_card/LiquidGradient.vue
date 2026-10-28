<template>
  <div class="bc-liquid-root min-h-screen px-5 py-7 text-white">
    <div class="bc-liquid-flow bc-liquid-flow-a"></div>
    <div class="bc-liquid-flow bc-liquid-flow-b"></div>
    <div class="bc-liquid-flow bc-liquid-flow-c"></div>
    <div class="bc-liquid-grain"></div>

    <div class="bc-liquid-wrap mx-auto flex min-h-[calc(100vh-3.5rem)] w-full max-w-md flex-col">
      <header class="bc-liquid-head">
        <span class="bc-liquid-dot"></span>
        <span>LIQUID / ID</span>
        <span class="bc-liquid-id">{{ nfcId ? '#' + nfcId : 'FLUID FORM' }}</span>
      </header>

      <section class="bc-liquid-hero">
        <div class="bc-liquid-avatar">{{ initial }}</div>
        <h1>{{ payload.name || '未命名' }}</h1>
        <p class="bc-liquid-role">{{ payload.title }}</p>
        <p class="bc-liquid-company">{{ payload.company }}</p>
      </section>

      <section class="bc-liquid-panel">
        <a v-if="payload.phone" :href="'tel:' + payload.phone" class="bc-liquid-pill">
          <span class="bc-liquid-pill-icon"><Phone :size="16" /></span>
          <span class="bc-liquid-pill-text"><small>拨打电话</small><b>{{ payload.phone }}</b></span>
          <span class="bc-liquid-pill-go"><ArrowUpRight :size="16" /></span>
        </a>
        <div v-else class="bc-liquid-pill bc-liquid-pill-off">
          <span class="bc-liquid-pill-icon"><Phone :size="16" /></span>
          <span class="bc-liquid-pill-text"><small>拨打电话</small><b></b></span>
        </div>

        <a v-if="payload.email" :href="'mailto:' + payload.email" class="bc-liquid-pill">
          <span class="bc-liquid-pill-icon"><Mail :size="16" /></span>
          <span class="bc-liquid-pill-text"><small>电子邮箱</small><b>{{ payload.email }}</b></span>
          <span class="bc-liquid-pill-go"><ArrowUpRight :size="16" /></span>
        </a>
        <div v-else class="bc-liquid-pill bc-liquid-pill-off">
          <span class="bc-liquid-pill-icon"><Mail :size="16" /></span>
          <span class="bc-liquid-pill-text"><small>电子邮箱</small><b></b></span>
        </div>

        <a
          v-if="payload.website"
          :href="payload.website"
          target="_blank"
          rel="noopener noreferrer"
          class="bc-liquid-pill"
        >
          <span class="bc-liquid-pill-icon"><Globe :size="16" /></span>
          <span class="bc-liquid-pill-text"><small>个人网站</small><b>{{ payload.website }}</b></span>
          <span class="bc-liquid-pill-go"><ArrowUpRight :size="16" /></span>
        </a>
        <div v-else class="bc-liquid-pill bc-liquid-pill-off">
          <span class="bc-liquid-pill-icon"><Globe :size="16" /></span>
          <span class="bc-liquid-pill-text"><small>个人网站</small><b></b></span>
        </div>

        <div class="bc-liquid-pill bc-liquid-pill-off">
          <span class="bc-liquid-pill-icon"><MessageCircle :size="16" /></span>
          <span class="bc-liquid-pill-text"><small>微信</small><b>{{ payload.wechat }}</b></span>
        </div>
      </section>

      <footer class="bc-liquid-foot">
        <span>FLOW WITH ME</span>
        <span>2026 © {{ payload.company }}</span>
      </footer>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { Phone, Mail, Globe, MessageCircle, ArrowUpRight } from 'lucide-vue-next'

const props = defineProps(['payload', 'nfcId'])

const initial = computed(() => (props.payload?.name || '?').toString().trim()[0] || '?')
</script>

<style scoped>
.bc-liquid-root {
  position: relative;
  overflow: hidden;
  background: #140b2e;
  font-family: 'Inter', 'PingFang SC', 'Microsoft YaHei', sans-serif;
}
.bc-liquid-flow {
  position: absolute;
  border-radius: 46% 54% 60% 40% / 50% 42% 58% 50%;
  filter: blur(38px);
  opacity: 0.75;
}
.bc-liquid-flow-a {
  top: -90px;
  left: -70px;
  width: 300px;
  height: 300px;
  background: linear-gradient(135deg, #fb7185, #f97316);
  animation: bc-liquid-move 14s ease-in-out infinite;
}
.bc-liquid-flow-b {
  top: 90px;
  right: -110px;
  width: 320px;
  height: 320px;
  background: linear-gradient(135deg, #6366f1, #22d3ee);
  animation: bc-liquid-move 18s ease-in-out infinite reverse;
}
.bc-liquid-flow-c {
  bottom: -110px;
  left: 14%;
  width: 290px;
  height: 290px;
  background: linear-gradient(135deg, #a855f7, #ec4899);
  animation: bc-liquid-move 12s ease-in-out infinite 1s;
}
.bc-liquid-grain {
  position: absolute;
  inset: 0;
  background-image: radial-gradient(rgba(255, 255, 255, 0.22) 0.5px, transparent 0.5px);
  background-size: 4px 4px;
  opacity: 0.14;
  pointer-events: none;
}
.bc-liquid-wrap {
  position: relative;
  z-index: 2;
}
.bc-liquid-head {
  display: flex;
  align-items: center;
  gap: 8px;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 8px;
  letter-spacing: 0.24em;
  color: rgba(255, 255, 255, 0.65);
  animation: bc-liquid-fade 0.8s both;
}
.bc-liquid-dot {
  width: 7px;
  height: 7px;
  border-radius: 50%;
  background: #fda4af;
  box-shadow: 0 0 12px #fb7185;
  animation: bc-liquid-pulse 2.4s ease-in-out infinite;
}
.bc-liquid-id {
  margin-left: auto;
  color: rgba(255, 255, 255, 0.4);
}
.bc-liquid-hero {
  margin-top: 54px;
  animation: bc-liquid-rise 0.9s 0.1s cubic-bezier(0.22, 1, 0.36, 1) both;
}
.bc-liquid-avatar {
  display: flex;
  width: 70px;
  height: 70px;
  align-items: center;
  justify-content: center;
  border: 1px solid rgba(255, 255, 255, 0.35);
  border-radius: 50%;
  background: linear-gradient(145deg, rgba(255, 255, 255, 0.28), rgba(255, 255, 255, 0.06));
  box-shadow: inset 0 1px 0 rgba(255, 255, 255, 0.4), 0 12px 34px -14px rgba(0, 0, 0, 0.6);
  backdrop-filter: blur(12px);
  font-size: 28px;
  font-weight: 300;
}
.bc-liquid-hero h1 {
  margin-top: 22px;
  font-size: clamp(42px, 14vw, 58px);
  font-weight: 700;
  line-height: 1;
  letter-spacing: -0.04em;
  overflow-wrap: anywhere;
  text-shadow: 0 10px 36px rgba(20, 11, 46, 0.45);
}
.bc-liquid-role {
  margin-top: 14px;
  font-size: 13px;
  letter-spacing: 0.12em;
  color: rgba(255, 255, 255, 0.85);
}
.bc-liquid-company {
  margin-top: 3px;
  font-size: 11px;
  letter-spacing: 0.06em;
  color: rgba(255, 255, 255, 0.5);
}
.bc-liquid-panel {
  display: flex;
  flex-direction: column;
  gap: 9px;
  margin-top: 42px;
  padding-bottom: 8px;
  animation: bc-liquid-rise 1s 0.24s cubic-bezier(0.22, 1, 0.36, 1) both;
}
.bc-liquid-pill {
  display: flex;
  align-items: center;
  gap: 12px;
  border: 1px solid rgba(255, 255, 255, 0.24);
  border-radius: 999px;
  background: rgba(255, 255, 255, 0.1);
  padding: 9px 11px;
  text-decoration: none;
  color: inherit;
  backdrop-filter: blur(16px);
  transition: background 0.3s ease, transform 0.3s cubic-bezier(0.22, 1, 0.36, 1), border-color 0.3s ease;
}
.bc-liquid-pill:hover {
  border-color: rgba(255, 255, 255, 0.55);
  background: rgba(255, 255, 255, 0.2);
  transform: translateX(4px);
}
.bc-liquid-pill:active {
  transform: translateX(2px) scale(0.99);
}
.bc-liquid-pill-off {
  opacity: 0.55;
}
.bc-liquid-pill-icon {
  display: flex;
  width: 36px;
  height: 36px;
  flex: none;
  align-items: center;
  justify-content: center;
  border-radius: 50%;
  background: rgba(255, 255, 255, 0.15);
  color: #fecdd3;
}
.bc-liquid-pill-text {
  display: flex;
  min-width: 0;
  flex: 1;
  flex-direction: column;
  gap: 1px;
}
.bc-liquid-pill-text small {
  font-size: 8px;
  letter-spacing: 0.14em;
  color: rgba(255, 255, 255, 0.5);
}
.bc-liquid-pill-text b {
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  font-size: 12.5px;
  font-weight: 500;
  color: rgba(255, 255, 255, 0.95);
}
.bc-liquid-pill-go {
  display: flex;
  width: 30px;
  height: 30px;
  flex: none;
  align-items: center;
  justify-content: center;
  border-radius: 50%;
  background: #ffffff;
  color: #1e1b4b;
  transition: transform 0.3s ease;
}
.bc-liquid-pill:hover .bc-liquid-pill-go {
  transform: rotate(45deg);
}
.bc-liquid-foot {
  display: flex;
  justify-content: space-between;
  gap: 10px;
  margin-top: auto;
  padding-top: 30px;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 7px;
  letter-spacing: 0.16em;
  color: rgba(255, 255, 255, 0.4);
}
@keyframes bc-liquid-rise {
  from { opacity: 0; transform: translateY(22px); }
  to { opacity: 1; transform: none; }
}
@keyframes bc-liquid-fade {
  from { opacity: 0; }
  to { opacity: 1; }
}
@keyframes bc-liquid-move {
  0%, 100% { transform: translate3d(0, 0, 0) rotate(0) scale(1); }
  33% { transform: translate3d(20px, 26px, 0) rotate(18deg) scale(1.08); }
  66% { transform: translate3d(-14px, 12px, 0) rotate(-12deg) scale(0.96); }
}
@keyframes bc-liquid-pulse {
  0%, 100% { opacity: 1; transform: scale(1); }
  50% { opacity: 0.5; transform: scale(0.8); }
}
@media (prefers-reduced-motion: reduce) {
  .bc-liquid-root *,
  .bc-liquid-root *::before,
  .bc-liquid-root *::after {
    animation: none !important;
    transition: none !important;
  }
}
</style>
