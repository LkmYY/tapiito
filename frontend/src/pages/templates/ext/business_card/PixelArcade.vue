<template>
  <div class="bc-pixel-root min-h-screen px-4 py-8">
    <div class="bc-pixel-sky"></div>
    <div class="bc-pixel-wrap mx-auto flex min-h-[calc(100vh-4rem)] w-full max-w-md flex-col">
      <div class="bc-pixel-hud">
        <span>PLAYER 1</span>
        <span class="bc-pixel-coin">◉ {{ nfcId || 'CREDIT 01' }}</span>
      </div>

      <div class="bc-pixel-console">
        <div class="bc-pixel-titlebar">
          <span class="bc-pixel-square"></span>
          <span>CARD.EXE</span>
          <span class="bc-pixel-blink">_</span>
        </div>

        <div class="bc-pixel-screen">
          <p class="bc-pixel-over">INSERT COIN TO CONNECT</p>
          <div class="bc-pixel-avatar">
            <span>{{ initial }}</span>
          </div>
          <h1 class="bc-pixel-name">{{ payload.name || '未命名' }}</h1>
          <p class="bc-pixel-role">CLASS: {{ payload.title }}</p>
          <p class="bc-pixel-company">GUILD: {{ payload.company }}</p>

          <div class="bc-pixel-stats">
            <div class="bc-pixel-stat">
              <span>LV.07</span>
              <span>NETWORKER</span>
            </div>
            <div class="bc-pixel-bar"><i></i></div>
            <div class="bc-pixel-hearts">♥ ♥ ♥ ♥ ♡</div>
          </div>

          <nav class="bc-pixel-menu">
            <a v-if="payload.phone" :href="'tel:' + payload.phone" class="bc-pixel-item">
              <span>&gt; CALL</span><b>{{ payload.phone }}</b>
            </a>
            <div v-else class="bc-pixel-item bc-pixel-item-off">
              <span>&gt; CALL</span><b></b>
            </div>
            <a v-if="payload.email" :href="'mailto:' + payload.email" class="bc-pixel-item">
              <span>&gt; MAIL</span><b>{{ payload.email }}</b>
            </a>
            <div v-else class="bc-pixel-item bc-pixel-item-off">
              <span>&gt; MAIL</span><b></b>
            </div>
            <a v-if="payload.website" :href="payload.website" target="_blank" rel="noopener noreferrer" class="bc-pixel-item">
              <span>&gt; WWW</span><b>{{ payload.website }}</b>
            </a>
            <div v-else class="bc-pixel-item bc-pixel-item-off">
              <span>&gt; WWW</span><b></b>
            </div>
            <div class="bc-pixel-item bc-pixel-item-off">
              <span>&gt; WX</span><b>{{ payload.wechat }}</b>
            </div>
          </nav>
        </div>

        <div class="bc-pixel-bottom">
          <span>PRESS START</span>
          <span>1UP</span>
        </div>
      </div>

      <p class="bc-pixel-foot">© 1987 PIXEL BUREAU · ALL RIGHTS RESERVED</p>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps(['payload', 'nfcId'])

const initial = computed(() => (props.payload?.name || '?').toString().trim()[0] || '?')
</script>

<style scoped>
.bc-pixel-root {
  position: relative;
  overflow: hidden;
  background: #080b1a;
  font-family: 'Courier New', 'PingFang SC', 'Microsoft YaHei', monospace;
  image-rendering: pixelated;
}
.bc-pixel-sky {
  position: absolute;
  inset: 0;
  background:
    radial-gradient(circle at 20% 15%, rgba(34, 197, 94, 0.14), transparent 30%),
    radial-gradient(circle at 80% 80%, rgba(59, 130, 246, 0.18), transparent 38%),
    linear-gradient(180deg, #0f172a, #080b1a);
}
.bc-pixel-sky::before {
  content: '';
  position: absolute;
  inset: 0;
  background-image: linear-gradient(rgba(34, 197, 94, 0.07) 1px, transparent 1px),
    linear-gradient(90deg, rgba(34, 197, 94, 0.07) 1px, transparent 1px);
  background-size: 16px 16px;
}
.bc-pixel-root::after {
  content: '';
  position: absolute;
  inset: 0;
  background: repeating-linear-gradient(180deg, transparent 0 2px, rgba(0, 0, 0, 0.28) 2px 4px);
  pointer-events: none;
  opacity: 0.5;
}
.bc-pixel-wrap {
  position: relative;
  z-index: 2;
}
.bc-pixel-hud {
  display: flex;
  justify-content: space-between;
  padding: 0 3px 8px;
  font-size: 8px;
  letter-spacing: 0.08em;
  color: #4ade80;
  animation: bc-pixel-fade 0.7s both;
}
.bc-pixel-coin {
  color: #fde047;
}
.bc-pixel-console {
  border: 3px solid #22c55e;
  background: #0a0f1f;
  box-shadow:
    0 0 0 3px #080b1a,
    0 0 0 6px #14532d,
    10px 10px 0 rgba(0, 0, 0, 0.55),
    0 0 34px rgba(34, 197, 94, 0.18);
  animation: bc-pixel-boot 0.75s steps(6, end) both;
}
.bc-pixel-titlebar {
  display: flex;
  align-items: center;
  gap: 8px;
  border-bottom: 3px solid #22c55e;
  padding: 7px 9px;
  font-size: 9px;
  letter-spacing: 0.14em;
  color: #86efac;
  background: #052e16;
}
.bc-pixel-square {
  width: 8px;
  height: 8px;
  background: #22c55e;
  box-shadow: 10px 0 0 #facc15, 20px 0 0 #ef4444;
  margin-right: 20px;
}
.bc-pixel-titlebar span:nth-child(2) {
  flex: 1;
}
.bc-pixel-blink {
  animation: bc-pixel-blink 1s steps(1, end) infinite;
}
.bc-pixel-screen {
  padding: 22px 18px 20px;
  background-image: repeating-linear-gradient(0deg, rgba(34, 197, 94, 0.035) 0 1px, transparent 1px 3px);
}
.bc-pixel-over {
  text-align: center;
  font-size: 7px;
  letter-spacing: 0.2em;
  color: #facc15;
  animation: bc-pixel-blink 2.2s steps(1, end) infinite;
}
.bc-pixel-avatar {
  display: flex;
  width: 72px;
  height: 72px;
  margin: 18px auto 0;
  align-items: center;
  justify-content: center;
  border: 3px solid #4ade80;
  background: #14532d;
  box-shadow:
    inset 0 0 0 3px #052e16,
    6px 6px 0 #052e16;
}
.bc-pixel-avatar span {
  font-family: 'Courier New', monospace;
  font-size: 34px;
  font-weight: 700;
  color: #dcfce7;
}
.bc-pixel-name {
  margin-top: 18px;
  text-align: center;
  font-family: 'Courier New', 'PingFang SC', monospace;
  font-size: clamp(27px, 9vw, 36px);
  font-weight: 700;
  line-height: 1.15;
  letter-spacing: 0.02em;
  color: #f0fdf4;
  text-shadow: 3px 3px 0 #166534;
  overflow-wrap: anywhere;
}
.bc-pixel-role {
  margin-top: 12px;
  text-align: center;
  font-size: 9px;
  letter-spacing: 0.08em;
  color: #4ade80;
}
.bc-pixel-company {
  margin-top: 4px;
  text-align: center;
  font-size: 8px;
  letter-spacing: 0.06em;
  color: #86efac;
  opacity: 0.7;
}
.bc-pixel-stats {
  margin-top: 20px;
  border: 2px solid #14532d;
  padding: 9px 10px;
  background: #03140a;
}
.bc-pixel-stat {
  display: flex;
  justify-content: space-between;
  font-size: 6.5px;
  letter-spacing: 0.1em;
  color: #86efac;
}
.bc-pixel-bar {
  height: 7px;
  margin-top: 7px;
  border: 1px solid #22c55e;
  padding: 1px;
}
.bc-pixel-bar i {
  display: block;
  width: 72%;
  height: 100%;
  background: repeating-linear-gradient(90deg, #4ade80 0 4px, #22c55e 4px 6px);
  animation: bc-pixel-load 1.2s 0.5s steps(6, end) both;
}
.bc-pixel-hearts {
  margin-top: 8px;
  font-size: 8px;
  letter-spacing: 0.2em;
  color: #f87171;
}
.bc-pixel-menu {
  margin-top: 18px;
}
.bc-pixel-item {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 7px 6px;
  border-bottom: 1px dashed rgba(74, 222, 128, 0.22);
  text-decoration: none;
  color: inherit;
  transition: background 0.12s steps(2, end), padding-left 0.12s steps(2, end);
}
.bc-pixel-item:hover {
  background: #14532d;
  padding-left: 10px;
}
.bc-pixel-item:active {
  background: #166534;
}
.bc-pixel-item span {
  flex: none;
  width: 58px;
  font-size: 7.5px;
  letter-spacing: 0.06em;
  color: #facc15;
}
.bc-pixel-item b {
  min-width: 0;
  flex: 1;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  font-size: 9px;
  font-weight: 400;
  text-align: right;
  color: #dcfce7;
}
.bc-pixel-item-off {
  opacity: 0.45;
}
.bc-pixel-bottom {
  display: flex;
  justify-content: space-between;
  border-top: 3px solid #22c55e;
  padding: 8px 9px;
  background: #052e16;
  font-size: 7px;
  letter-spacing: 0.16em;
  color: #86efac;
}
.bc-pixel-bottom span:first-child {
  animation: bc-pixel-blink 1.6s steps(1, end) infinite;
}
.bc-pixel-foot {
  margin-top: auto;
  padding-top: 18px;
  text-align: center;
  font-size: 6px;
  letter-spacing: 0.12em;
  color: rgba(74, 222, 128, 0.32);
}
@keyframes bc-pixel-boot {
  from { opacity: 0; transform: scaleY(0.7); }
  to { opacity: 1; transform: none; }
}
@keyframes bc-pixel-fade {
  from { opacity: 0; }
  to { opacity: 1; }
}
@keyframes bc-pixel-load {
  from { width: 0; }
  to { width: 72%; }
}
@keyframes bc-pixel-blink {
  0%, 49% { opacity: 1; }
  50%, 100% { opacity: 0; }
}
@media (prefers-reduced-motion: reduce) {
  .bc-pixel-root *,
  .bc-pixel-root *::before,
  .bc-pixel-root *::after {
    animation: none !important;
    transition: none !important;
  }
}
</style>
