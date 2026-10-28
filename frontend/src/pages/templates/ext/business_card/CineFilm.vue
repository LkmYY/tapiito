<template>
  <div class="bc-film-root min-h-screen px-4 py-8">
    <div class="bc-film-glow"></div>
    <div class="bc-film-wrap mx-auto flex min-h-[calc(100vh-4rem)] w-full max-w-md flex-col">
      <header class="bc-film-top">
        <span>REEL 01 / TAKE 07</span>
        <span class="bc-film-rec"><i></i>REC</span>
      </header>

      <div class="bc-film-strip">
        <div class="bc-film-holes bc-film-holes-l"></div>
        <div class="bc-film-holes bc-film-holes-r"></div>

        <div class="bc-film-frame">
          <div class="bc-film-scan"></div>
          <p class="bc-film-frame-meta">35MM · COLOR · 2026</p>
          <div class="bc-film-portrait">
            <span>{{ initial }}</span>
            <div class="bc-film-light"></div>
          </div>
          <h1 class="bc-film-name">{{ payload.name || '未命名' }}</h1>
          <p class="bc-film-role">{{ payload.title }}</p>
          <p class="bc-film-company">{{ payload.company }}</p>
        </div>

        <div class="bc-film-credits">
          <p class="bc-film-credits-title">CONTACT / CREDITS</p>
          <div class="bc-film-credit">
            <span>STARRING</span>
            <p>{{ payload.name || '未命名' }}</p>
          </div>
          <div class="bc-film-credit">
            <span>DIRECT LINE</span>
            <a v-if="payload.phone" :href="'tel:' + payload.phone">{{ payload.phone }}</a>
          </div>
          <div class="bc-film-credit">
            <span>WRITE TO</span>
            <a v-if="payload.email" :href="'mailto:' + payload.email">{{ payload.email }}</a>
          </div>
          <div class="bc-film-credit">
            <span>SCREENING</span>
            <a v-if="payload.website" :href="payload.website" target="_blank" rel="noopener noreferrer">{{ payload.website }}</a>
          </div>
          <div class="bc-film-credit">
            <span>BACKSTAGE</span>
            <p :class="{ 'bc-film-empty': !payload.wechat }">{{ payload.wechat }}</p>
          </div>
        </div>
      </div>

      <footer class="bc-film-foot">
        <span>A YUNTU PICTURE</span>
        <span class="bc-film-foot-num">{{ nfcId || '000' }}</span>
        <span>FIN</span>
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
.bc-film-root {
  position: relative;
  overflow: hidden;
  background: #0c0a09;
  font-family: 'Helvetica Neue', 'PingFang SC', 'Microsoft YaHei', sans-serif;
  color: #e7e5e4;
}
.bc-film-root::before {
  content: '';
  position: absolute;
  inset: 0;
  background-image: radial-gradient(rgba(245, 158, 11, 0.14) 0.6px, transparent 0.6px);
  background-size: 5px 5px;
  opacity: 0.5;
}
.bc-film-glow {
  position: absolute;
  top: 12%;
  left: 50%;
  width: 360px;
  height: 360px;
  margin-left: -180px;
  border-radius: 50%;
  background: radial-gradient(circle, rgba(245, 158, 11, 0.2), transparent 65%);
  filter: blur(24px);
  animation: bc-film-flicker 6s ease-in-out infinite;
}
.bc-film-wrap {
  position: relative;
  z-index: 1;
}
.bc-film-top {
  display: flex;
  justify-content: space-between;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 7px;
  letter-spacing: 0.2em;
  color: rgba(231, 229, 228, 0.45);
  animation: bc-film-fade 0.8s both;
}
.bc-film-rec {
  display: inline-flex;
  align-items: center;
  gap: 5px;
  color: #f87171;
}
.bc-film-rec i {
  width: 5px;
  height: 5px;
  border-radius: 50%;
  background: #ef4444;
  box-shadow: 0 0 8px #ef4444;
  animation: bc-film-blink 1.6s steps(1, end) infinite;
}
.bc-film-strip {
  position: relative;
  margin-top: 18px;
  border-top: 1px solid rgba(231, 229, 228, 0.2);
  border-bottom: 1px solid rgba(231, 229, 228, 0.2);
  padding: 0 18px;
  background: #171412;
  box-shadow: 0 30px 70px -40px rgba(0, 0, 0, 0.9);
  animation: bc-film-roll 0.9s cubic-bezier(0.22, 1, 0.36, 1) both;
}
.bc-film-holes {
  position: absolute;
  top: 0;
  bottom: 0;
  width: 18px;
  background: repeating-linear-gradient(180deg, transparent 0 9px, #e7e5e4 9px 15px, transparent 15px 26px);
  opacity: 0.9;
}
.bc-film-holes-l {
  left: 0;
  border-right: 1px solid rgba(231, 229, 228, 0.14);
}
.bc-film-holes-r {
  right: 0;
  border-left: 1px solid rgba(231, 229, 228, 0.14);
}
.bc-film-frame {
  position: relative;
  margin: 0 4px;
  padding: 16px 14px 22px;
  border-bottom: 1px dashed rgba(231, 229, 228, 0.25);
  background: linear-gradient(180deg, rgba(245, 158, 11, 0.06), transparent 26%);
}
.bc-film-scan {
  position: absolute;
  inset: 0;
  background: repeating-linear-gradient(180deg, transparent 0 2px, rgba(0, 0, 0, 0.16) 2px 4px);
  pointer-events: none;
}
.bc-film-frame-meta {
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 5.5px;
  letter-spacing: 0.18em;
  color: rgba(245, 158, 11, 0.65);
}
.bc-film-portrait {
  position: relative;
  display: flex;
  height: 188px;
  margin-top: 12px;
  align-items: center;
  justify-content: center;
  overflow: hidden;
  border: 1px solid rgba(231, 229, 228, 0.16);
  background:
    radial-gradient(circle at 62% 28%, rgba(245, 158, 11, 0.55), transparent 32%),
    radial-gradient(circle at 30% 70%, rgba(120, 53, 15, 0.7), transparent 55%),
    linear-gradient(145deg, #292524, #0c0a09);
  transition: filter 0.4s ease;
}
.bc-film-portrait:hover {
  filter: brightness(1.15);
}
.bc-film-portrait span {
  position: relative;
  z-index: 1;
  font-family: 'Times New Roman', Georgia, serif;
  font-size: 104px;
  line-height: 1;
  color: rgba(255, 247, 237, 0.92);
  text-shadow: 0 0 30px rgba(245, 158, 11, 0.45);
}
.bc-film-light {
  position: absolute;
  top: -30%;
  left: -20%;
  width: 80%;
  height: 80%;
  border-radius: 50%;
  background: radial-gradient(circle, rgba(255, 237, 213, 0.22), transparent 60%);
  animation: bc-film-sweep 8s ease-in-out infinite;
}
.bc-film-name {
  margin-top: 16px;
  font-size: clamp(30px, 10vw, 40px);
  font-weight: 300;
  line-height: 1.05;
  letter-spacing: 0.02em;
  text-transform: uppercase;
  overflow-wrap: anywhere;
}
.bc-film-role {
  margin-top: 8px;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 8px;
  letter-spacing: 0.16em;
  color: #f59e0b;
}
.bc-film-company {
  margin-top: 3px;
  font-size: 9px;
  letter-spacing: 0.08em;
  color: rgba(231, 229, 228, 0.45);
}
.bc-film-credits {
  margin: 0 4px;
  padding: 14px 4px 18px;
}
.bc-film-credits-title {
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 5.5px;
  letter-spacing: 0.24em;
  color: rgba(231, 229, 228, 0.35);
}
.bc-film-credit {
  display: flex;
  align-items: baseline;
  gap: 12px;
  padding: 7px 0;
  border-bottom: 1px solid rgba(231, 229, 228, 0.08);
}
.bc-film-credit span {
  flex: none;
  width: 78px;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6px;
  letter-spacing: 0.12em;
  color: rgba(245, 158, 11, 0.6);
}
.bc-film-credit p,
.bc-film-credit a,
.bc-film-credit em {
  min-width: 0;
  flex: 1;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  font-size: 9.5px;
  font-style: normal;
  color: rgba(245, 245, 244, 0.85);
  text-decoration: none;
  transition: color 0.2s ease;
}
.bc-film-credit a:hover {
  color: #fbbf24;
  text-decoration: underline;
}
.bc-film-credit a:active {
  opacity: 0.6;
}
.bc-film-credit em,
.bc-film-empty {
  color: rgba(245, 245, 244, 0.28);
}
.bc-film-foot {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-top: auto;
  padding-top: 20px;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6px;
  letter-spacing: 0.2em;
  color: rgba(231, 229, 228, 0.35);
}
.bc-film-foot-num {
  color: rgba(245, 158, 11, 0.6);
}
@keyframes bc-film-roll {
  from { opacity: 0; transform: translateY(26px) scaleY(0.96); }
  to { opacity: 1; transform: none; }
}
@keyframes bc-film-fade {
  from { opacity: 0; }
  to { opacity: 1; }
}
@keyframes bc-film-flicker {
  0%, 100% { opacity: 0.75; }
  45% { opacity: 1; }
  50% { opacity: 0.55; }
  55% { opacity: 0.95; }
}
@keyframes bc-film-sweep {
  0%, 100% { transform: translate3d(0, 0, 0); opacity: 0.6; }
  50% { transform: translate3d(40px, 24px, 0); opacity: 1; }
}
@keyframes bc-film-blink {
  0%, 49% { opacity: 1; }
  50%, 100% { opacity: 0; }
}
@media (prefers-reduced-motion: reduce) {
  .bc-film-root *,
  .bc-film-root *::before,
  .bc-film-root *::after {
    animation: none !important;
    transition: none !important;
  }
}
</style>
