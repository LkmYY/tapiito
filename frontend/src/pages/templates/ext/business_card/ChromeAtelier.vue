<template>
  <div class="bc-chrome-root min-h-screen overflow-hidden text-[#e7e5e4]">
    <div class="bc-chrome-noise"></div>
    <div class="bc-chrome-blob"></div>
    <div class="bc-chrome-ring"></div>

    <div class="bc-chrome-wrap mx-auto flex min-h-screen w-full max-w-md flex-col px-7 py-8">
      <header class="bc-chrome-top">
        <span>ATELIER No.7</span>
        <span>{{ nfcId ? 'NFC ' + nfcId : 'PRIVATE EDITION' }}</span>
      </header>

      <main class="bc-chrome-main">
        <p class="bc-chrome-kicker">CRAFTED IDENTITY</p>
        <h1 class="bc-chrome-name">{{ payload.name || '未命名' }}</h1>
        <div class="bc-chrome-rule"><i></i></div>
        <p class="bc-chrome-role">{{ payload.title }}</p>
        <p class="bc-chrome-company">{{ payload.company }}</p>

        <div class="bc-chrome-card">
          <span class="bc-chrome-card-edge"></span>
          <div class="bc-chrome-row">
            <span class="bc-chrome-label">TELEPHONE</span>
            <a v-if="payload.phone" :href="'tel:' + payload.phone" class="bc-chrome-value">{{ payload.phone }}</a>
          </div>
          <div class="bc-chrome-row">
            <span class="bc-chrome-label">CORRESPONDENCE</span>
            <a v-if="payload.email" :href="'mailto:' + payload.email" class="bc-chrome-value">{{ payload.email }}</a>
          </div>
          <div class="bc-chrome-row">
            <span class="bc-chrome-label">PORTFOLIO</span>
            <a
              v-if="payload.website"
              :href="payload.website"
              target="_blank"
              rel="noopener noreferrer"
              class="bc-chrome-value"
            >
              {{ payload.website }}
            </a>
          </div>
          <div class="bc-chrome-row">
            <span class="bc-chrome-label">WECHAT</span>
            <span class="bc-chrome-value" :class="{ 'bc-chrome-muted': !payload.wechat }">{{ payload.wechat }}</span>
          </div>
        </div>
      </main>

      <footer class="bc-chrome-foot">
        <span class="bc-chrome-monogram">{{ initial }}</span>
        <p>MMXXVI · HAND FINISHED<br />SILVER SERIES</p>
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
.bc-chrome-root {
  position: relative;
  background: #08080a;
  font-family: 'Helvetica Neue', 'PingFang SC', 'Microsoft YaHei', sans-serif;
}
.bc-chrome-noise {
  position: absolute;
  inset: 0;
  background-image: radial-gradient(rgba(255, 255, 255, 0.16) 0.4px, transparent 0.4px);
  background-size: 3px 3px;
  opacity: 0.14;
  pointer-events: none;
}
.bc-chrome-blob {
  position: absolute;
  top: 4%;
  left: 50%;
  width: 430px;
  height: 430px;
  margin-left: -215px;
  border-radius: 42% 58% 63% 37% / 46% 38% 62% 54%;
  background:
    radial-gradient(circle at 30% 22%, #ffffff 0%, #d6d3d1 14%, #78716c 30%, #18181b 52%, #a8a29e 68%, #52525b 82%, #09090b 100%);
  filter: blur(1px) contrast(1.15);
  opacity: 0.65;
  animation: bc-chrome-morph 16s ease-in-out infinite;
}
.bc-chrome-blob::after {
  content: '';
  position: absolute;
  inset: 8%;
  border-radius: inherit;
  background: radial-gradient(circle at 65% 70%, rgba(255, 255, 255, 0.45), transparent 32%);
  mix-blend-mode: screen;
}
.bc-chrome-ring {
  position: absolute;
  right: -130px;
  bottom: -140px;
  width: 340px;
  height: 340px;
  border: 1px solid rgba(212, 175, 55, 0.28);
  border-radius: 50%;
  pointer-events: none;
}
.bc-chrome-ring::before {
  content: '';
  position: absolute;
  inset: 26px;
  border: 1px solid rgba(212, 175, 55, 0.16);
  border-radius: 50%;
}
.bc-chrome-wrap {
  position: relative;
  z-index: 2;
}
.bc-chrome-top {
  display: flex;
  justify-content: space-between;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 7.5px;
  letter-spacing: 0.24em;
  color: rgba(231, 229, 228, 0.45);
  animation: bc-chrome-fade 0.9s both;
}
.bc-chrome-main {
  margin-top: auto;
  padding-top: 120px;
  animation: bc-chrome-rise 1s 0.1s cubic-bezier(0.22, 1, 0.36, 1) both;
}
.bc-chrome-kicker {
  font-size: 8px;
  letter-spacing: 0.42em;
  color: #d4af37;
}
.bc-chrome-name {
  margin-top: 18px;
  font-family: 'Didot', 'Times New Roman', Georgia, 'Songti SC', serif;
  font-size: clamp(50px, 16vw, 72px);
  font-weight: 400;
  line-height: 0.98;
  letter-spacing: -0.03em;
  overflow-wrap: anywhere;
  background: linear-gradient(115deg, #ffffff 0%, #a8a29e 30%, #f5f5f4 48%, #71717a 64%, #e7e5e4 82%, #ffffff 100%);
  background-size: 220% 100%;
  -webkit-background-clip: text;
  background-clip: text;
  color: transparent;
  animation: bc-chrome-sheen 7s ease-in-out infinite;
}
.bc-chrome-rule {
  display: flex;
  align-items: center;
  gap: 10px;
  margin: 24px 0 20px;
}
.bc-chrome-rule::before {
  content: '';
  width: 48px;
  height: 1px;
  background: linear-gradient(90deg, #d4af37, transparent);
}
.bc-chrome-rule i {
  width: 4px;
  height: 4px;
  background: #d4af37;
  transform: rotate(45deg);
}
.bc-chrome-role {
  font-size: 12px;
  font-weight: 500;
  letter-spacing: 0.2em;
  color: #e7e5e4;
}
.bc-chrome-company {
  margin-top: 6px;
  font-family: Georgia, 'Songti SC', serif;
  font-size: 12px;
  font-style: italic;
  letter-spacing: 0.04em;
  color: rgba(231, 229, 228, 0.5);
}
.bc-chrome-card {
  position: relative;
  margin-top: 44px;
  border-top: 1px solid rgba(212, 175, 55, 0.36);
  border-bottom: 1px solid rgba(212, 175, 55, 0.2);
  background: linear-gradient(160deg, rgba(255, 255, 255, 0.055), rgba(255, 255, 255, 0.012));
  backdrop-filter: blur(8px);
}
.bc-chrome-card-edge {
  position: absolute;
  top: -1px;
  left: 0;
  width: 34%;
  height: 1px;
  background: #d4af37;
  box-shadow: 0 0 12px rgba(212, 175, 55, 0.7);
}
.bc-chrome-row {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 16px;
  padding: 15px 4px;
  border-bottom: 1px solid rgba(255, 255, 255, 0.055);
}
.bc-chrome-row:last-child {
  border-bottom: 0;
}
.bc-chrome-label {
  flex: none;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6.5px;
  letter-spacing: 0.22em;
  color: rgba(212, 175, 55, 0.72);
}
.bc-chrome-value {
  min-width: 0;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  font-family: Georgia, 'Songti SC', serif;
  font-size: 12.5px;
  letter-spacing: 0.02em;
  text-align: right;
  color: rgba(245, 245, 244, 0.92);
  text-decoration: none;
  transition: color 0.25s ease, letter-spacing 0.25s ease;
}
a.bc-chrome-value:hover {
  color: #e5c76b;
  letter-spacing: 0.05em;
}
a.bc-chrome-value:active {
  opacity: 0.6;
}
.bc-chrome-muted {
  color: rgba(245, 245, 244, 0.25);
}
.bc-chrome-foot {
  display: flex;
  align-items: flex-end;
  justify-content: space-between;
  padding-top: 36px;
  animation: bc-chrome-fade 1s 0.35s both;
}
.bc-chrome-monogram {
  display: flex;
  width: 46px;
  height: 46px;
  align-items: center;
  justify-content: center;
  border: 1px solid rgba(212, 175, 55, 0.55);
  border-radius: 50%;
  font-family: Georgia, serif;
  font-size: 18px;
  color: #d4af37;
  transition: background 0.3s ease, color 0.3s ease, transform 0.3s ease;
}
.bc-chrome-monogram:hover {
  background: #d4af37;
  color: #09090b;
  transform: rotate(8deg);
}
.bc-chrome-foot p {
  text-align: right;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6.5px;
  line-height: 1.8;
  letter-spacing: 0.16em;
  color: rgba(231, 229, 228, 0.32);
}
@keyframes bc-chrome-rise {
  from { opacity: 0; transform: translateY(28px); }
  to { opacity: 1; transform: none; }
}
@keyframes bc-chrome-fade {
  from { opacity: 0; }
  to { opacity: 1; }
}
@keyframes bc-chrome-morph {
  0%, 100% { border-radius: 42% 58% 63% 37% / 46% 38% 62% 54%; transform: rotate(0) scale(1); }
  33% { border-radius: 58% 42% 38% 62% / 58% 62% 38% 42%; transform: rotate(8deg) scale(1.03); }
  66% { border-radius: 38% 62% 58% 42% / 42% 38% 62% 58%; transform: rotate(-6deg) scale(0.98); }
}
@keyframes bc-chrome-sheen {
  0%, 100% { background-position: 0% 50%; }
  50% { background-position: 100% 50%; }
}
@media (prefers-reduced-motion: reduce) {
  .bc-chrome-root *,
  .bc-chrome-root *::before,
  .bc-chrome-root *::after {
    animation: none !important;
    transition: none !important;
  }
}
</style>
