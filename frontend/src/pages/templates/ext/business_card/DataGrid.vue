<template>
  <div class="bc-data-root min-h-screen px-4 py-7">
    <div class="bc-data-gridbg"></div>
    <div class="bc-data-wrap mx-auto w-full max-w-md">
      <header class="bc-data-head">
        <span class="bc-data-brand">PROFILE<span>.EXE</span></span>
        <span class="bc-data-live"><i></i>SYNCED</span>
        <span class="bc-data-clock">2026 / 10 / 20</span>
      </header>

      <section class="bc-data-hero">
        <div class="bc-data-identity">
          <p class="bc-data-file">FILE_ID / {{ nfcId || 'BC-0007' }}</p>
          <h1 class="bc-data-name">{{ payload.name || '未命名' }}</h1>
          <p class="bc-data-role">{{ payload.title }}</p>
          <p class="bc-data-company">{{ payload.company }}</p>
        </div>
        <div class="bc-data-ring">
          <div class="bc-data-ring-inner">{{ initial }}</div>
        </div>
      </section>

      <section class="bc-data-stats">
        <div class="bc-data-stat">
          <span>NETWORK</span>
          <strong>98<span>%</span></strong>
          <div class="bc-data-spark bc-data-spark-a"><i></i><i></i><i></i><i></i><i></i><i></i><i></i></div>
        </div>
        <div class="bc-data-stat">
          <span>RESPONSE</span>
          <strong>2<span>h</span></strong>
          <div class="bc-data-spark bc-data-spark-b"><i></i><i></i><i></i><i></i><i></i><i></i><i></i></div>
        </div>
      </section>

      <section class="bc-data-modules">
        <div class="bc-data-modules-head">
          <span>CONTACT MODULES</span>
          <span>4 ACTIVE</span>
        </div>
        <div class="bc-data-tiles">
          <a v-if="payload.phone" :href="'tel:' + payload.phone" class="bc-data-tile">
            <span class="bc-data-tile-id">M-01</span>
            <span class="bc-data-tile-label">VOICE</span>
            <b>{{ payload.phone }}</b>
            <i class="bc-data-tile-dot"></i>
          </a>
          <div v-else class="bc-data-tile bc-data-off">
            <span class="bc-data-tile-id">M-01</span>
            <span class="bc-data-tile-label">VOICE</span>
            <b></b>
          </div>

          <a v-if="payload.email" :href="'mailto:' + payload.email" class="bc-data-tile">
            <span class="bc-data-tile-id">M-02</span>
            <span class="bc-data-tile-label">MAIL</span>
            <b>{{ payload.email }}</b>
            <i class="bc-data-tile-dot"></i>
          </a>
          <div v-else class="bc-data-tile bc-data-off">
            <span class="bc-data-tile-id">M-02</span>
            <span class="bc-data-tile-label">MAIL</span>
            <b></b>
          </div>

          <a v-if="payload.website" :href="payload.website" target="_blank" rel="noopener noreferrer" class="bc-data-tile bc-data-wide">
            <span class="bc-data-tile-id">M-03</span>
            <span class="bc-data-tile-label">WEB / PORTFOLIO</span>
            <b>{{ payload.website }}</b>
            <i class="bc-data-tile-dot"></i>
          </a>
          <div v-else class="bc-data-tile bc-data-wide bc-data-off">
            <span class="bc-data-tile-id">M-03</span>
            <span class="bc-data-tile-label">WEB / PORTFOLIO</span>
            <b></b>
          </div>

          <div class="bc-data-tile bc-data-wide bc-data-off">
            <span class="bc-data-tile-id">M-04</span>
            <span class="bc-data-tile-label">WECHAT / RELAY</span>
            <b>{{ payload.wechat }}</b>
            <i class="bc-data-tile-dot bc-data-dim"></i>
          </div>
        </div>
      </section>

      <footer class="bc-data-foot">
        <span>LAST UPDATE 00:07</span>
        <div class="bc-data-bars"><i></i><i></i><i></i><i></i><i></i></div>
        <span>ENCRYPTED</span>
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
.bc-data-root {
  position: relative;
  overflow: hidden;
  background: #07111f;
  font-family: 'Inter', 'PingFang SC', 'Microsoft YaHei', sans-serif;
  color: #e2e8f0;
}
.bc-data-gridbg {
  position: absolute;
  inset: 0;
  background-image: linear-gradient(rgba(56, 189, 248, 0.06) 1px, transparent 1px),
    linear-gradient(90deg, rgba(56, 189, 248, 0.06) 1px, transparent 1px);
  background-size: 28px 28px;
}
.bc-data-root::before {
  content: '';
  position: absolute;
  top: -100px;
  right: -80px;
  width: 280px;
  height: 280px;
  border-radius: 50%;
  background: radial-gradient(circle, rgba(14, 165, 233, 0.28), transparent 65%);
  pointer-events: none;
}
.bc-data-wrap {
  position: relative;
  z-index: 1;
}
.bc-data-head {
  display: flex;
  align-items: center;
  gap: 10px;
  padding-bottom: 14px;
  border-bottom: 1px solid rgba(56, 189, 248, 0.2);
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6.5px;
  letter-spacing: 0.16em;
  animation: bc-data-fade 0.8s both;
}
.bc-data-brand {
  font-size: 11px;
  font-weight: 700;
  letter-spacing: 0.02em;
  color: #f8fafc;
}
.bc-data-brand span {
  color: #38bdf8;
}
.bc-data-live {
  display: inline-flex;
  align-items: center;
  gap: 4px;
  color: #34d399;
}
.bc-data-live i {
  width: 4px;
  height: 4px;
  border-radius: 50%;
  background: #34d399;
  box-shadow: 0 0 8px #34d399;
  animation: bc-data-pulse 2s ease-in-out infinite;
}
.bc-data-clock {
  margin-left: auto;
  color: rgba(226, 232, 240, 0.35);
}
.bc-data-hero {
  display: flex;
  align-items: center;
  gap: 14px;
  padding: 30px 0 24px;
  animation: bc-data-rise 0.9s 0.1s cubic-bezier(0.22, 1, 0.36, 1) both;
}
.bc-data-identity {
  min-width: 0;
  flex: 1;
}
.bc-data-file {
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 5.5px;
  letter-spacing: 0.2em;
  color: #38bdf8;
}
.bc-data-name {
  margin-top: 8px;
  font-size: clamp(30px, 9.6vw, 40px);
  font-weight: 600;
  line-height: 1.05;
  letter-spacing: -0.035em;
  overflow-wrap: anywhere;
  color: #f8fafc;
}
.bc-data-role {
  margin-top: 9px;
  font-size: 10.5px;
  letter-spacing: 0.08em;
  color: #7dd3fc;
}
.bc-data-company {
  margin-top: 3px;
  font-size: 9px;
  color: rgba(226, 232, 240, 0.4);
}
.bc-data-ring {
  position: relative;
  display: flex;
  width: 82px;
  height: 82px;
  flex: none;
  align-items: center;
  justify-content: center;
  border-radius: 50%;
  background: conic-gradient(#38bdf8 0 78%, rgba(56, 189, 248, 0.12) 78% 100%);
  box-shadow: 0 0 24px rgba(56, 189, 248, 0.2);
}
.bc-data-ring::before {
  content: '';
  position: absolute;
  inset: 5px;
  border-radius: 50%;
  background: #07111f;
}
.bc-data-ring-inner {
  position: relative;
  z-index: 1;
  display: flex;
  width: 60px;
  height: 60px;
  align-items: center;
  justify-content: center;
  border: 1px solid rgba(56, 189, 248, 0.28);
  border-radius: 50%;
  font-size: 24px;
  font-weight: 300;
  color: #bae6fd;
}
.bc-data-stats {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 8px;
  animation: bc-data-rise 0.9s 0.2s cubic-bezier(0.22, 1, 0.36, 1) both;
}
.bc-data-stat {
  border: 1px solid rgba(56, 189, 248, 0.16);
  background: rgba(15, 23, 42, 0.65);
  padding: 10px 11px;
}
.bc-data-stat > span {
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 5.5px;
  letter-spacing: 0.16em;
  color: rgba(125, 211, 252, 0.5);
}
.bc-data-stat strong {
  display: block;
  margin-top: 4px;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 22px;
  font-weight: 400;
  color: #e0f2fe;
}
.bc-data-stat strong span {
  margin-left: 2px;
  font-size: 9px;
  color: #38bdf8;
}
.bc-data-spark {
  display: flex;
  align-items: flex-end;
  gap: 3px;
  height: 16px;
  margin-top: 6px;
}
.bc-data-spark i {
  flex: 1;
  background: linear-gradient(180deg, #38bdf8, rgba(56, 189, 248, 0.2));
}
.bc-data-spark-a i:nth-child(1) { height: 30%; }
.bc-data-spark-a i:nth-child(2) { height: 52%; }
.bc-data-spark-a i:nth-child(3) { height: 44%; }
.bc-data-spark-a i:nth-child(4) { height: 72%; }
.bc-data-spark-a i:nth-child(5) { height: 60%; }
.bc-data-spark-a i:nth-child(6) { height: 88%; }
.bc-data-spark-a i:nth-child(7) { height: 100%; }
.bc-data-spark-b i:nth-child(1) { height: 80%; }
.bc-data-spark-b i:nth-child(2) { height: 62%; }
.bc-data-spark-b i:nth-child(3) { height: 70%; }
.bc-data-spark-b i:nth-child(4) { height: 42%; }
.bc-data-spark-b i:nth-child(5) { height: 52%; }
.bc-data-spark-b i:nth-child(6) { height: 30%; }
.bc-data-spark-b i:nth-child(7) { height: 46%; }
.bc-data-modules {
  margin-top: 22px;
  animation: bc-data-rise 1s 0.3s cubic-bezier(0.22, 1, 0.36, 1) both;
}
.bc-data-modules-head {
  display: flex;
  justify-content: space-between;
  padding-bottom: 7px;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 5.5px;
  letter-spacing: 0.16em;
  color: rgba(125, 211, 252, 0.45);
}
.bc-data-tiles {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 8px;
}
.bc-data-tile {
  position: relative;
  display: flex;
  min-height: 86px;
  flex-direction: column;
  justify-content: flex-end;
  gap: 3px;
  border: 1px solid rgba(56, 189, 248, 0.2);
  border-radius: 4px;
  background: linear-gradient(160deg, rgba(15, 23, 42, 0.9), rgba(7, 17, 31, 0.8));
  padding: 10px;
  text-decoration: none;
  color: inherit;
  transition: border-color 0.25s ease, background 0.25s ease, transform 0.25s ease;
}
.bc-data-tile:hover {
  border-color: rgba(56, 189, 248, 0.7);
  background: rgba(8, 47, 73, 0.55);
  transform: translateY(-2px);
}
.bc-data-tile:active {
  transform: translateY(0);
}
.bc-data-wide {
  grid-column: 1 / -1;
  min-height: 72px;
}
.bc-data-tile-id {
  position: absolute;
  top: 8px;
  left: 9px;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 5px;
  letter-spacing: 0.12em;
  color: rgba(125, 211, 252, 0.4);
}
.bc-data-tile-label {
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 5.5px;
  letter-spacing: 0.14em;
  color: #38bdf8;
}
.bc-data-tile b {
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  font-size: 10.5px;
  font-weight: 500;
  color: rgba(240, 249, 255, 0.9);
}
.bc-data-tile-dot {
  position: absolute;
  top: 10px;
  right: 10px;
  width: 4px;
  height: 4px;
  border-radius: 50%;
  background: #34d399;
  box-shadow: 0 0 6px #34d399;
}
.bc-data-dim {
  background: #64748b;
  box-shadow: none;
}
.bc-data-off {
  opacity: 0.55;
}
.bc-data-foot {
  display: flex;
  align-items: center;
  gap: 10px;
  margin-top: 22px;
  padding-top: 12px;
  border-top: 1px solid rgba(56, 189, 248, 0.14);
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 5.5px;
  letter-spacing: 0.14em;
  color: rgba(125, 211, 252, 0.4);
}
.bc-data-bars {
  display: flex;
  flex: 1;
  gap: 3px;
}
.bc-data-bars i {
  flex: 1;
  height: 2px;
  background: rgba(56, 189, 248, 0.25);
}
.bc-data-bars i:nth-child(-n + 3) {
  background: #38bdf8;
}
@keyframes bc-data-rise {
  from { opacity: 0; transform: translateY(16px); }
  to { opacity: 1; transform: none; }
}
@keyframes bc-data-fade {
  from { opacity: 0; }
  to { opacity: 1; }
}
@keyframes bc-data-pulse {
  0%, 100% { opacity: 1; }
  50% { opacity: 0.3; }
}
@media (prefers-reduced-motion: reduce) {
  .bc-data-root *,
  .bc-data-root *::before,
  .bc-data-root *::after {
    animation: none !important;
    transition: none !important;
  }
}
</style>
