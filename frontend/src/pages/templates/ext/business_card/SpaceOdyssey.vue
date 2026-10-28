<template>
  <div class="bc-space-root min-h-screen px-5 py-8">
    <div class="bc-space-stars bc-space-stars-a"></div>
    <div class="bc-space-stars bc-space-stars-b"></div>
    <div class="bc-space-nebula"></div>

    <div class="bc-space-wrap mx-auto flex min-h-[calc(100vh-4rem)] w-full max-w-md flex-col">
      <header class="bc-space-head">
        <span>ORBITAL IDENTITY</span>
        <span class="bc-space-coord">31.2304° N / 121.4737° E</span>
      </header>

      <section class="bc-space-orbit">
        <div class="bc-space-ring bc-space-ring-a"></div>
        <div class="bc-space-ring bc-space-ring-b"></div>
        <div class="bc-space-ring bc-space-ring-c"></div>
        <span class="bc-space-sat bc-space-sat-a"></span>
        <span class="bc-space-sat bc-space-sat-b"></span>
        <div class="bc-space-planet">
          <span class="bc-space-initial">{{ initial }}</span>
          <div class="bc-space-planet-shadow"></div>
        </div>
        <div class="bc-space-patch">
          <span>MISSION</span>
          <strong>07</strong>
          <span>YUNTU</span>
        </div>
      </section>

      <section class="bc-space-hero">
        <p class="bc-space-kicker">COMMANDER OF CONNECTION</p>
        <h1 class="bc-space-name">{{ payload.name || '未命名' }}</h1>
        <p class="bc-space-role">{{ payload.title }}</p>
        <p class="bc-space-company">{{ payload.company }}</p>
      </section>

      <section class="bc-space-signals">
        <div class="bc-space-signal-head">
          <span>TRANSMISSION CHANNELS</span>
          <span>{{ nfcId ? 'SIGNAL ' + nfcId : 'LIVE' }}</span>
        </div>
        <a v-if="payload.phone" :href="'tel:' + payload.phone" class="bc-space-signal">
          <span class="bc-space-sig-no">01</span>
          <span class="bc-space-sig-label">VOICE</span>
          <span class="bc-space-sig-value">{{ payload.phone }}</span>
          <span class="bc-space-sig-arrow">↗</span>
        </a>
        <div v-else class="bc-space-signal bc-space-off">
          <span class="bc-space-sig-no">01</span>
          <span class="bc-space-sig-label">VOICE</span>
          <span class="bc-space-sig-value"></span>
        </div>
        <a v-if="payload.email" :href="'mailto:' + payload.email" class="bc-space-signal">
          <span class="bc-space-sig-no">02</span>
          <span class="bc-space-sig-label">DATA</span>
          <span class="bc-space-sig-value">{{ payload.email }}</span>
          <span class="bc-space-sig-arrow">↗</span>
        </a>
        <div v-else class="bc-space-signal bc-space-off">
          <span class="bc-space-sig-no">02</span>
          <span class="bc-space-sig-label">DATA</span>
          <span class="bc-space-sig-value"></span>
        </div>
        <a v-if="payload.website" :href="payload.website" target="_blank" rel="noopener noreferrer" class="bc-space-signal">
          <span class="bc-space-sig-no">03</span>
          <span class="bc-space-sig-label">PORTAL</span>
          <span class="bc-space-sig-value">{{ payload.website }}</span>
          <span class="bc-space-sig-arrow">↗</span>
        </a>
        <div v-else class="bc-space-signal bc-space-off">
          <span class="bc-space-sig-no">03</span>
          <span class="bc-space-sig-label">PORTAL</span>
          <span class="bc-space-sig-value"></span>
        </div>
        <div class="bc-space-signal bc-space-off">
          <span class="bc-space-sig-no">04</span>
          <span class="bc-space-sig-label">RELAY</span>
          <span class="bc-space-sig-value">{{ payload.wechat }}</span>
        </div>
      </section>

      <footer class="bc-space-foot">
        <span class="bc-space-line"></span>
        <span>SEE YOU AMONG THE STARS</span>
        <span class="bc-space-line"></span>
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
.bc-space-root {
  position: relative;
  overflow: hidden;
  background: #050716;
  font-family: 'Helvetica Neue', 'PingFang SC', 'Microsoft YaHei', sans-serif;
  color: #e0e7ff;
}
.bc-space-stars {
  position: absolute;
  inset: 0;
  background-repeat: repeat;
  pointer-events: none;
}
.bc-space-stars-a {
  background-image:
    radial-gradient(1px 1px at 20px 30px, #ffffff, transparent),
    radial-gradient(1px 1px at 120px 80px, rgba(255, 255, 255, 0.8), transparent),
    radial-gradient(1.4px 1.4px at 80px 150px, #a5b4fc, transparent),
    radial-gradient(1px 1px at 180px 40px, rgba(255, 255, 255, 0.6), transparent);
  background-size: 220px 200px;
  animation: bc-space-twinkle 7s ease-in-out infinite;
}
.bc-space-stars-b {
  background-image:
    radial-gradient(1px 1px at 40px 100px, rgba(255, 255, 255, 0.9), transparent),
    radial-gradient(1.6px 1.6px at 160px 20px, #f0abfc, transparent),
    radial-gradient(1px 1px at 100px 170px, rgba(255, 255, 255, 0.5), transparent);
  background-size: 260px 240px;
  opacity: 0.7;
  animation: bc-space-twinkle 9s 1s ease-in-out infinite reverse;
}
.bc-space-nebula {
  position: absolute;
  top: -10%;
  right: -30%;
  width: 420px;
  height: 420px;
  border-radius: 50%;
  background: radial-gradient(circle, rgba(109, 40, 217, 0.35), rgba(30, 27, 75, 0.15) 45%, transparent 68%);
  filter: blur(20px);
  pointer-events: none;
}
.bc-space-wrap {
  position: relative;
  z-index: 1;
}
.bc-space-head {
  display: flex;
  justify-content: space-between;
  gap: 10px;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6px;
  letter-spacing: 0.2em;
  color: rgba(224, 231, 255, 0.45);
  animation: bc-space-fade 0.9s both;
}
.bc-space-orbit {
  position: relative;
  height: 286px;
  margin-top: 6px;
  animation: bc-space-orbit-in 1.2s 0.1s cubic-bezier(0.22, 1, 0.36, 1) both;
}
.bc-space-ring {
  position: absolute;
  top: 50%;
  left: 50%;
  border: 1px solid rgba(165, 180, 252, 0.22);
  border-radius: 50%;
  transform: translate(-50%, -50%);
}
.bc-space-ring-a {
  width: 238px;
  height: 238px;
  animation: bc-space-rotate 28s linear infinite;
}
.bc-space-ring-b {
  width: 186px;
  height: 186px;
  border-style: dashed;
  border-color: rgba(165, 180, 252, 0.28);
  animation: bc-space-rotate 20s linear infinite reverse;
}
.bc-space-ring-c {
  width: 284px;
  height: 130px;
  transform: translate(-50%, -50%) rotate(-22deg);
  border-color: rgba(165, 180, 252, 0.14);
}
.bc-space-sat {
  position: absolute;
  width: 6px;
  height: 6px;
  border-radius: 50%;
  background: #a5b4fc;
  box-shadow: 0 0 10px #a5b4fc;
}
.bc-space-sat-a {
  top: 24px;
  left: 50%;
  animation: bc-space-blink 3s ease-in-out infinite;
}
.bc-space-sat-b {
  right: 52px;
  bottom: 54px;
  width: 4px;
  height: 4px;
  background: #f0abfc;
  box-shadow: 0 0 10px #f0abfc;
  animation: bc-space-blink 4s 0.8s ease-in-out infinite;
}
.bc-space-planet {
  position: absolute;
  top: 50%;
  left: 50%;
  display: flex;
  width: 132px;
  height: 132px;
  align-items: center;
  justify-content: center;
  overflow: hidden;
  border-radius: 50%;
  background:
    radial-gradient(circle at 32% 28%, #c7d2fe 0%, #818cf8 18%, #4f46e5 42%, #1e1b4b 78%);
  box-shadow: inset -18px -14px 32px rgba(2, 6, 23, 0.7), 0 0 60px rgba(99, 102, 241, 0.45);
  transform: translate(-50%, -50%);
  transition: box-shadow 0.5s ease, transform 0.5s ease;
}
.bc-space-planet:hover {
  transform: translate(-50%, -50%) scale(1.04);
  box-shadow: inset -18px -14px 32px rgba(2, 6, 23, 0.7), 0 0 90px rgba(129, 140, 248, 0.65);
}
.bc-space-initial {
  position: relative;
  z-index: 1;
  font-family: Georgia, serif;
  font-size: 52px;
  color: rgba(238, 242, 255, 0.9);
  text-shadow: 0 4px 24px rgba(2, 6, 23, 0.6);
}
.bc-space-planet-shadow {
  position: absolute;
  inset: 0;
  border-radius: 50%;
  background: repeating-linear-gradient(12deg, transparent 0 12px, rgba(15, 23, 42, 0.12) 12px 20px);
}
.bc-space-patch {
  position: absolute;
  right: 4px;
  bottom: 14px;
  display: flex;
  width: 60px;
  height: 60px;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  border: 1px solid rgba(224, 231, 255, 0.4);
  border-radius: 50%;
  background: rgba(15, 23, 42, 0.5);
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 4.5px;
  letter-spacing: 0.14em;
  color: rgba(224, 231, 255, 0.65);
  transform: rotate(8deg);
  backdrop-filter: blur(6px);
}
.bc-space-patch strong {
  font-size: 17px;
  font-weight: 400;
  letter-spacing: 0.02em;
  color: #f0abfc;
}
.bc-space-hero {
  margin-top: 6px;
  text-align: center;
  animation: bc-space-rise 0.9s 0.28s cubic-bezier(0.22, 1, 0.36, 1) both;
}
.bc-space-kicker {
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6px;
  letter-spacing: 0.34em;
  color: #a5b4fc;
}
.bc-space-name {
  margin-top: 10px;
  font-size: clamp(35px, 11vw, 46px);
  font-weight: 200;
  line-height: 1.05;
  letter-spacing: 0.1em;
  text-transform: uppercase;
  overflow-wrap: anywhere;
  color: #eef2ff;
  text-shadow: 0 0 34px rgba(165, 180, 252, 0.45);
}
.bc-space-role {
  margin-top: 12px;
  font-size: 10px;
  letter-spacing: 0.18em;
  color: rgba(224, 231, 255, 0.68);
}
.bc-space-company {
  margin-top: 4px;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 7px;
  letter-spacing: 0.14em;
  color: rgba(224, 231, 255, 0.38);
}
.bc-space-signals {
  margin-top: 28px;
  border-top: 1px solid rgba(165, 180, 252, 0.25);
  animation: bc-space-fade 1s 0.45s both;
}
.bc-space-signal-head {
  display: flex;
  justify-content: space-between;
  padding: 9px 2px;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 5.5px;
  letter-spacing: 0.18em;
  color: rgba(224, 231, 255, 0.35);
}
.bc-space-signal {
  display: grid;
  grid-template-columns: 22px 46px 1fr 14px;
  align-items: center;
  gap: 6px;
  padding: 10px 2px;
  border-top: 1px solid rgba(165, 180, 252, 0.15);
  text-decoration: none;
  color: inherit;
  transition: background 0.25s ease, padding-left 0.25s ease;
}
a.bc-space-signal:hover {
  background: linear-gradient(90deg, rgba(129, 140, 248, 0.14), transparent);
  padding-left: 7px;
}
a.bc-space-signal:hover .bc-space-sig-value {
  color: #ffffff;
}
.bc-space-sig-no {
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 5.5px;
  color: rgba(165, 180, 252, 0.4);
}
.bc-space-sig-label {
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 5.5px;
  letter-spacing: 0.14em;
  color: #a5b4fc;
}
.bc-space-sig-value {
  min-width: 0;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  font-size: 9.5px;
  letter-spacing: 0.02em;
  color: rgba(224, 231, 255, 0.82);
  transition: color 0.2s ease;
}
.bc-space-sig-arrow {
  font-size: 9px;
  color: rgba(224, 231, 255, 0.3);
}
.bc-space-off {
  opacity: 0.42;
}
.bc-space-foot {
  display: flex;
  align-items: center;
  gap: 10px;
  margin-top: auto;
  padding-top: 22px;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 5.5px;
  letter-spacing: 0.18em;
  color: rgba(224, 231, 255, 0.35);
}
.bc-space-line {
  flex: 1;
  height: 1px;
  background: rgba(165, 180, 252, 0.2);
}
@keyframes bc-space-orbit-in {
  from { opacity: 0; transform: scale(0.85) rotate(-8deg); }
  to { opacity: 1; transform: none; }
}
@keyframes bc-space-rise {
  from { opacity: 0; transform: translateY(18px); }
  to { opacity: 1; transform: none; }
}
@keyframes bc-space-fade {
  from { opacity: 0; }
  to { opacity: 1; }
}
@keyframes bc-space-rotate {
  from { transform: translate(-50%, -50%) rotate(0); }
  to { transform: translate(-50%, -50%) rotate(360deg); }
}
@keyframes bc-space-twinkle {
  0%, 100% { opacity: 0.55; }
  50% { opacity: 1; }
}
@keyframes bc-space-blink {
  0%, 100% { opacity: 0.35; }
  50% { opacity: 1; }
}
@media (prefers-reduced-motion: reduce) {
  .bc-space-root *,
  .bc-space-root *::before,
  .bc-space-root *::after {
    animation: none !important;
    transition: none !important;
  }
}
</style>
