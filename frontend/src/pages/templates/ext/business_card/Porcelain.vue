<template>
  <div class="bc-porc-root min-h-screen px-6 py-10 text-stone-900">
    <div class="bc-porc-content mx-auto flex min-h-[calc(100vh-5rem)] w-full max-w-md flex-col">
      <header class="bc-porc-head">
        <span>PORTRAIT — 01</span>
        <span>{{ nfcId ? 'NFC ' + nfcId : 'DIGITAL CARD' }}</span>
      </header>

      <section class="bc-porc-hero">
        <p class="bc-porc-eyebrow">A quiet introduction</p>
        <h1 class="bc-porc-name">{{ payload.name || '未命名' }}</h1>
        <div class="bc-porc-rule"></div>
        <p class="bc-porc-title">{{ payload.title }}</p>
        <p class="bc-porc-company">{{ payload.company }}</p>
      </section>

      <section class="bc-porc-details">
        <div class="bc-porc-label">
          <span>CONTACT</span>
          <span>联络方式</span>
        </div>

        <div class="bc-porc-list">
          <div class="bc-porc-row">
            <span class="bc-porc-key">TEL</span>
            <a v-if="payload.phone" :href="'tel:' + payload.phone" class="bc-porc-value bc-porc-link">
              {{ payload.phone }}
            </a>
          </div>
          <div class="bc-porc-row">
            <span class="bc-porc-key">MAIL</span>
            <a v-if="payload.email" :href="'mailto:' + payload.email" class="bc-porc-value bc-porc-link">
              {{ payload.email }}
            </a>
          </div>
          <div class="bc-porc-row">
            <span class="bc-porc-key">WEB</span>
            <a
              v-if="payload.website"
              :href="payload.website"
              target="_blank"
              rel="noopener noreferrer"
              class="bc-porc-value bc-porc-link"
            >
              {{ payload.website }}
            </a>
          </div>
          <div class="bc-porc-row">
            <span class="bc-porc-key">WECHAT</span>
            <span class="bc-porc-value" :class="{ 'bc-porc-empty': !payload.wechat }">
              {{ payload.wechat }}
            </span>
          </div>
        </div>
      </section>

      <footer class="bc-porc-foot">
        <p>以简洁，留下余韵。</p>
        <span class="bc-porc-mark">{{ initial }}</span>
      </footer>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps(['payload', 'nfcId'])

const initial = computed(() => (props.payload?.name || '?').toString().trim()[0] || '?')
</script>

<style scoped>
.bc-porc-root {
  position: relative;
  overflow: hidden;
  background: #fdfcf9;
  font-family: 'Helvetica Neue', 'PingFang SC', 'Microsoft YaHei', sans-serif;
}
.bc-porc-root::before {
  content: '';
  position: absolute;
  top: 0;
  right: 0;
  width: 55%;
  height: 38%;
  background: radial-gradient(ellipse at 100% 0, rgba(214, 211, 209, 0.4), transparent 68%);
  pointer-events: none;
}
.bc-porc-root::after {
  content: '';
  position: absolute;
  inset: 0;
  background-image: radial-gradient(rgba(41, 37, 36, 0.15) 0.5px, transparent 0.5px);
  background-size: 5px 5px;
  opacity: 0.18;
  pointer-events: none;
}
.bc-porc-content {
  position: relative;
  z-index: 1;
}
.bc-porc-head {
  display: flex;
  justify-content: space-between;
  padding-bottom: 12px;
  border-bottom: 1px solid #d6d3d1;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 8.5px;
  letter-spacing: 0.22em;
  color: #a8a29e;
  animation: bc-porc-in 0.8s both;
}
.bc-porc-hero {
  padding-top: 74px;
  animation: bc-porc-in 0.9s 0.08s cubic-bezier(0.22, 1, 0.36, 1) both;
}
.bc-porc-eyebrow {
  font-family: Georgia, 'Songti SC', serif;
  font-size: 13px;
  font-style: italic;
  letter-spacing: 0.04em;
  color: #a8a29e;
}
.bc-porc-name {
  margin-top: 18px;
  font-family: Georgia, 'Songti SC', 'SimSun', serif;
  font-size: clamp(46px, 15vw, 68px);
  font-weight: 400;
  line-height: 0.98;
  letter-spacing: -0.045em;
  overflow-wrap: anywhere;
}
.bc-porc-rule {
  width: 100%;
  height: 1px;
  margin: 30px 0 22px;
  background: #292524;
  transform-origin: left;
  animation: bc-porc-grow 1s 0.35s cubic-bezier(0.22, 1, 0.36, 1) both;
}
.bc-porc-title {
  font-size: 14px;
  font-weight: 500;
  letter-spacing: 0.06em;
}
.bc-porc-company {
  margin-top: 6px;
  font-size: 12px;
  letter-spacing: 0.12em;
  color: #a8a29e;
}
.bc-porc-details {
  margin-top: 66px;
  animation: bc-porc-in 0.9s 0.22s cubic-bezier(0.22, 1, 0.36, 1) both;
}
.bc-porc-label {
  display: flex;
  align-items: baseline;
  justify-content: space-between;
  padding-bottom: 12px;
  border-bottom: 1px solid #292524;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 8.5px;
  letter-spacing: 0.2em;
  color: #78716c;
}
.bc-porc-label span:last-child {
  font-family: 'PingFang SC', 'Microsoft YaHei', sans-serif;
  letter-spacing: 0.2em;
}
.bc-porc-list {
  display: flex;
  flex-direction: column;
}
.bc-porc-row {
  display: flex;
  align-items: baseline;
  gap: 18px;
  padding: 17px 0;
  border-bottom: 1px solid #e7e5e4;
}
.bc-porc-key {
  width: 52px;
  flex: none;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 8px;
  letter-spacing: 0.16em;
  color: #a8a29e;
}
.bc-porc-value {
  min-width: 0;
  flex: 1;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  font-size: 13px;
  letter-spacing: 0.02em;
  color: #292524;
  text-align: right;
}
.bc-porc-link {
  position: relative;
  text-decoration: none;
  transition: color 0.25s ease;
}
.bc-porc-link::after {
  content: '';
  position: absolute;
  left: 0;
  bottom: -2px;
  width: 100%;
  height: 1px;
  background: currentColor;
  transform: scaleX(0);
  transform-origin: right;
  transition: transform 0.3s cubic-bezier(0.22, 1, 0.36, 1);
}
.bc-porc-link:hover {
  color: #a16207;
}
.bc-porc-link:hover::after {
  transform: scaleX(1);
  transform-origin: left;
}
.bc-porc-link:active {
  opacity: 0.65;
}
.bc-porc-empty {
  color: #d6d3d1;
}
.bc-porc-foot {
  display: flex;
  align-items: flex-end;
  justify-content: space-between;
  margin-top: auto;
  padding-top: 46px;
  animation: bc-porc-in 1s 0.4s both;
}
.bc-porc-foot p {
  font-family: Georgia, 'Songti SC', serif;
  font-size: 11px;
  font-style: italic;
  letter-spacing: 0.08em;
  color: #a8a29e;
}
.bc-porc-mark {
  display: flex;
  width: 36px;
  height: 36px;
  align-items: center;
  justify-content: center;
  border: 1px solid #d6d3d1;
  border-radius: 50%;
  font-family: Georgia, serif;
  font-size: 15px;
  color: #78716c;
  transition: background 0.3s ease, color 0.3s ease, transform 0.3s ease;
}
.bc-porc-mark:hover {
  background: #292524;
  color: #fdfcf9;
  transform: rotate(-8deg);
}
@keyframes bc-porc-in {
  from { opacity: 0; transform: translateY(16px); }
  to { opacity: 1; transform: none; }
}
@keyframes bc-porc-grow {
  from { transform: scaleX(0); }
  to { transform: scaleX(1); }
}
@media (prefers-reduced-motion: reduce) {
  .bc-porc-root *,
  .bc-porc-root *::before,
  .bc-porc-root *::after {
    animation: none !important;
    transition: none !important;
  }
}
</style>
