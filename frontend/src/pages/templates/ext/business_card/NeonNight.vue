<template>
  <div class="bc-neon-root relative min-h-screen overflow-hidden px-5 py-8 text-white">
    <div class="bc-neon-glow bc-neon-glow-a"></div>
    <div class="bc-neon-glow bc-neon-glow-b"></div>
    <div class="bc-neon-grid"></div>
    <div class="bc-neon-scan"></div>

    <div class="relative z-10 mx-auto flex min-h-[calc(100vh-4rem)] w-full max-w-md flex-col">
      <header class="bc-neon-top">
        <span class="bc-neon-live"><i></i>ONLINE</span>
        <span>NFC / {{ nfcId || '0000' }}</span>
      </header>

      <main class="bc-neon-card">
        <span class="bc-neon-corner bc-neon-corner-tl"></span>
        <span class="bc-neon-corner bc-neon-corner-tr"></span>
        <span class="bc-neon-corner bc-neon-corner-bl"></span>
        <span class="bc-neon-corner bc-neon-corner-br"></span>

        <div class="bc-neon-badge">{{ initial }}</div>
        <p class="bc-neon-kicker">DIGITAL BUSINESS CARD</p>
        <h1 class="bc-neon-name">{{ payload.name || '未命名' }}</h1>
        <p class="bc-neon-role">{{ payload.title }}</p>
        <p class="bc-neon-company">{{ payload.company }}</p>

        <div class="bc-neon-divider"><span></span></div>

        <ul class="bc-neon-contacts">
          <li>
            <a v-if="payload.phone" :href="'tel:' + payload.phone" class="bc-neon-row">
              <span class="bc-neon-icon"><Phone :size="15" /></span>
              <span class="bc-neon-copy"><small>CALL</small><b>{{ payload.phone }}</b></span>
              <ArrowUpRight :size="15" class="bc-neon-arrow" />
            </a>
            <div v-else class="bc-neon-row bc-neon-off">
              <span class="bc-neon-icon"><Phone :size="15" /></span>
              <span class="bc-neon-copy"><small>CALL</small><b></b></span>
            </div>
          </li>
          <li>
            <a v-if="payload.email" :href="'mailto:' + payload.email" class="bc-neon-row">
              <span class="bc-neon-icon"><Mail :size="15" /></span>
              <span class="bc-neon-copy"><small>MAIL</small><b>{{ payload.email }}</b></span>
              <ArrowUpRight :size="15" class="bc-neon-arrow" />
            </a>
            <div v-else class="bc-neon-row bc-neon-off">
              <span class="bc-neon-icon"><Mail :size="15" /></span>
              <span class="bc-neon-copy"><small>MAIL</small><b></b></span>
            </div>
          </li>
          <li>
            <a v-if="payload.website" :href="payload.website" target="_blank" rel="noopener noreferrer" class="bc-neon-row">
              <span class="bc-neon-icon"><Globe :size="15" /></span>
              <span class="bc-neon-copy"><small>WEB</small><b>{{ payload.website }}</b></span>
              <ArrowUpRight :size="15" class="bc-neon-arrow" />
            </a>
            <div v-else class="bc-neon-row bc-neon-off">
              <span class="bc-neon-icon"><Globe :size="15" /></span>
              <span class="bc-neon-copy"><small>WEB</small><b></b></span>
            </div>
          </li>
          <li>
            <div class="bc-neon-row">
              <span class="bc-neon-icon"><MessageCircle :size="15" /></span>
              <span class="bc-neon-copy"><small>WECHAT</small><b>{{ payload.wechat }}</b></span>
              <span class="bc-neon-tag">WX</span>
            </div>
          </li>
        </ul>
      </main>

      <footer class="bc-neon-foot">
        <span>TAP TO CONNECT</span>
        <span class="bc-neon-dot"></span>
        <span>ALWAYS ON</span>
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
.bc-neon-root {
  background: #070510;
  font-family: 'Inter', 'PingFang SC', 'Microsoft YaHei', sans-serif;
}
.bc-neon-glow {
  position: absolute;
  border-radius: 999px;
  filter: blur(70px);
  opacity: 0.5;
  pointer-events: none;
}
.bc-neon-glow-a {
  width: 280px;
  height: 280px;
  top: -90px;
  left: -80px;
  background: #7c3aed;
  animation: bc-neon-float 9s ease-in-out infinite;
}
.bc-neon-glow-b {
  width: 240px;
  height: 240px;
  right: -100px;
  bottom: 40px;
  background: #db2777;
  animation: bc-neon-float 11s ease-in-out infinite reverse;
}
.bc-neon-grid {
  position: absolute;
  inset: 0;
  background-image: linear-gradient(rgba(168, 85, 247, 0.09) 1px, transparent 1px),
    linear-gradient(90deg, rgba(168, 85, 247, 0.09) 1px, transparent 1px);
  background-size: 34px 34px;
  mask-image: radial-gradient(ellipse at 50% 30%, #000 20%, transparent 75%);
  pointer-events: none;
}
.bc-neon-scan {
  position: absolute;
  inset: 0;
  background: repeating-linear-gradient(180deg, transparent 0 3px, rgba(255, 255, 255, 0.018) 3px 4px);
  pointer-events: none;
}
.bc-neon-top {
  display: flex;
  align-items: center;
  justify-content: space-between;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 9px;
  letter-spacing: 0.26em;
  color: rgba(255, 255, 255, 0.38);
  animation: bc-neon-fade 0.7s both;
}
.bc-neon-live {
  display: inline-flex;
  align-items: center;
  gap: 7px;
  color: #67e8f9;
}
.bc-neon-live i {
  width: 5px;
  height: 5px;
  border-radius: 50%;
  background: #22d3ee;
  box-shadow: 0 0 10px #22d3ee;
  animation: bc-neon-pulse 1.8s ease-in-out infinite;
}
.bc-neon-card {
  position: relative;
  margin-top: 34px;
  border: 1px solid rgba(192, 132, 252, 0.32);
  border-radius: 4px;
  background: linear-gradient(160deg, rgba(24, 16, 44, 0.92), rgba(9, 7, 19, 0.96));
  box-shadow: 0 0 0 1px rgba(0, 0, 0, 0.5), 0 30px 80px -30px rgba(168, 85, 247, 0.55);
  padding: 34px 26px 26px;
  backdrop-filter: blur(14px);
  animation: bc-neon-card-in 0.85s cubic-bezier(0.22, 1, 0.36, 1) both;
}
.bc-neon-card::before {
  content: '';
  position: absolute;
  inset: 6px;
  border: 1px dashed rgba(255, 255, 255, 0.07);
  pointer-events: none;
}
.bc-neon-corner {
  position: absolute;
  width: 16px;
  height: 16px;
  border-color: #e879f9;
  border-style: solid;
  filter: drop-shadow(0 0 6px rgba(232, 121, 249, 0.7));
}
.bc-neon-corner-tl { top: -1px; left: -1px; border-width: 2px 0 0 2px; }
.bc-neon-corner-tr { top: -1px; right: -1px; border-width: 2px 2px 0 0; }
.bc-neon-corner-bl { bottom: -1px; left: -1px; border-width: 0 0 2px 2px; }
.bc-neon-corner-br { bottom: -1px; right: -1px; border-width: 0 2px 2px 0; }
.bc-neon-badge {
  display: flex;
  width: 58px;
  height: 58px;
  align-items: center;
  justify-content: center;
  border: 1px solid rgba(103, 232, 249, 0.5);
  border-radius: 50%;
  background: radial-gradient(circle at 30% 25%, rgba(34, 211, 238, 0.28), rgba(126, 34, 206, 0.24));
  color: #a5f3fc;
  font-size: 24px;
  font-weight: 300;
  letter-spacing: 0.04em;
  box-shadow: 0 0 24px rgba(34, 211, 238, 0.22), inset 0 0 20px rgba(34, 211, 238, 0.12);
}
.bc-neon-kicker {
  margin-top: 24px;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 9px;
  letter-spacing: 0.34em;
  color: #e879f9;
}
.bc-neon-name {
  margin-top: 10px;
  font-size: 38px;
  font-weight: 200;
  line-height: 1.12;
  letter-spacing: 0.02em;
  text-shadow: 0 0 26px rgba(232, 121, 249, 0.4);
  overflow-wrap: anywhere;
}
.bc-neon-role {
  margin-top: 12px;
  font-size: 13px;
  font-weight: 500;
  letter-spacing: 0.16em;
  color: #67e8f9;
}
.bc-neon-company {
  margin-top: 5px;
  font-size: 12px;
  letter-spacing: 0.08em;
  color: rgba(255, 255, 255, 0.42);
}
.bc-neon-divider {
  position: relative;
  height: 1px;
  margin: 26px 0 8px;
  background: linear-gradient(90deg, rgba(232, 121, 249, 0.65), rgba(103, 232, 249, 0.32), transparent);
}
.bc-neon-divider span {
  position: absolute;
  top: -2px;
  left: 0;
  width: 42%;
  height: 5px;
  background: linear-gradient(90deg, transparent, rgba(232, 121, 249, 0.7), transparent);
  filter: blur(2px);
  animation: bc-neon-travel 4s ease-in-out infinite;
}
.bc-neon-contacts {
  list-style: none;
  margin: 0;
  padding: 0;
}
.bc-neon-row {
  position: relative;
  display: flex;
  align-items: center;
  gap: 13px;
  min-height: 58px;
  padding: 11px 8px;
  border-bottom: 1px solid rgba(255, 255, 255, 0.06);
  text-decoration: none;
  color: inherit;
  transition: background 0.25s ease, transform 0.25s ease, box-shadow 0.25s ease;
}
a.bc-neon-row:hover {
  background: linear-gradient(90deg, rgba(232, 121, 249, 0.12), transparent);
  box-shadow: inset 2px 0 0 #e879f9;
  transform: translateX(3px);
}
a.bc-neon-row:active {
  transform: translateX(1px) scale(0.995);
}
.bc-neon-icon {
  display: flex;
  width: 32px;
  height: 32px;
  flex: none;
  align-items: center;
  justify-content: center;
  border: 1px solid rgba(103, 232, 249, 0.22);
  border-radius: 3px;
  color: #67e8f9;
  background: rgba(34, 211, 238, 0.05);
  transition: box-shadow 0.25s ease, border-color 0.25s ease;
}
a.bc-neon-row:hover .bc-neon-icon {
  border-color: rgba(232, 121, 249, 0.65);
  box-shadow: 0 0 14px rgba(232, 121, 249, 0.28);
  color: #f0abfc;
}
.bc-neon-copy {
  display: flex;
  min-width: 0;
  flex: 1;
  flex-direction: column;
  gap: 3px;
}
.bc-neon-copy small {
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 8px;
  letter-spacing: 0.28em;
  color: rgba(255, 255, 255, 0.28);
}
.bc-neon-copy b {
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  font-size: 13.5px;
  font-weight: 400;
  letter-spacing: 0.04em;
  color: rgba(255, 255, 255, 0.88);
}
.bc-neon-off .bc-neon-copy b {
  color: rgba(255, 255, 255, 0.28);
}
.bc-neon-arrow {
  flex: none;
  color: rgba(255, 255, 255, 0.22);
  transition: color 0.25s ease, transform 0.25s ease;
}
a.bc-neon-row:hover .bc-neon-arrow {
  color: #e879f9;
  transform: translate(2px, -2px);
}
.bc-neon-tag {
  border: 1px solid rgba(255, 255, 255, 0.16);
  padding: 2px 6px;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 8px;
  letter-spacing: 0.16em;
  color: rgba(255, 255, 255, 0.4);
}
.bc-neon-foot {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 10px;
  padding: 26px 0 10px;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 8px;
  letter-spacing: 0.28em;
  color: rgba(255, 255, 255, 0.24);
  animation: bc-neon-fade 1s 0.5s both;
}
.bc-neon-dot {
  width: 3px;
  height: 3px;
  border-radius: 50%;
  background: #e879f9;
}
@keyframes bc-neon-card-in {
  from { opacity: 0; transform: translateY(26px) scale(0.98); }
  to { opacity: 1; transform: none; }
}
@keyframes bc-neon-fade {
  from { opacity: 0; }
  to { opacity: 1; }
}
@keyframes bc-neon-float {
  0%, 100% { transform: translate3d(0, 0, 0) scale(1); }
  50% { transform: translate3d(14px, -18px, 0) scale(1.08); }
}
@keyframes bc-neon-pulse {
  0%, 100% { opacity: 1; }
  50% { opacity: 0.3; }
}
@keyframes bc-neon-travel {
  0%, 100% { transform: translateX(0); opacity: 0.5; }
  50% { transform: translateX(120%); opacity: 1; }
}
@media (prefers-reduced-motion: reduce) {
  .bc-neon-root *,
  .bc-neon-root *::before,
  .bc-neon-root *::after {
    animation: none !important;
    transition: none !important;
  }
}
</style>
