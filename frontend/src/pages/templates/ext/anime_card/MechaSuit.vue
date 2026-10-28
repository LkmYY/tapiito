<template>
  <div class="ac-mecha-root relative min-h-screen overflow-hidden px-4 py-8 text-[#dff6ff]">
    <div class="ac-mecha-grid"></div>
    <div class="ac-mecha-glow"></div>
    <div class="ac-mecha-hazard ac-mecha-hz-top"></div>
    <div class="ac-mecha-hazard ac-mecha-hz-bottom"></div>

    <div class="relative z-10 mx-auto w-full max-w-md">
      <header class="ac-mecha-bar">
        <span class="ac-mecha-live"><i></i>UNIT ONLINE</span>
        <span>FRAME://{{ (nfcId || '0007') }}</span>
      </header>

      <article class="ac-mecha-card">
        <div class="ac-mecha-frame-line"></div>
        <span class="ac-mecha-notch ac-mecha-n1"></span>
        <span class="ac-mecha-notch ac-mecha-n2"></span>

        <div class="ac-mecha-id">
          <span class="ac-mecha-serial">MS-0{{ nfcId || '7' }}</span>
          <span class="ac-mecha-class">STRIKER / PROTOTYPE</span>
        </div>

        <div class="ac-mecha-hero">
          <span class="ac-mecha-ghost">07</span>
          <div class="ac-mecha-reticle">
            <Crosshair :size="22" />
          </div>
          <p class="ac-mecha-role">PILOT PROFILE</p>
          <h1 class="ac-mecha-name">{{ p.name || '未命名角色' }}</h1>
          <p class="ac-mecha-title">{{ p.title }}</p>
          <div class="ac-mecha-tags">
            <span><b>FRAME</b>{{ p.series || '未知作品' }}</span>
            <span class="ac-mecha-elem"><b>LOADOUT</b>{{ p.element }}</span>
          </div>
        </div>

        <div class="ac-mecha-telemetry">
          <div class="ac-mecha-cell">
            <span>SYNC RATE</span>
            <b>98.6<em>%</em></b>
            <i><u style="width: 98%"></u></i>
          </div>
          <div class="ac-mecha-cell">
            <span>THREAT LV</span>
            <b>S<em> / 07</em></b>
            <i><u style="width: 74%"></u></i>
          </div>
          <div class="ac-mecha-cell">
            <span>COMMS</span>
            <b class="ac-mecha-comms">@{{ p.social }}</b>
            <i><u style="width: 61%"></u></i>
          </div>
        </div>

        <div class="ac-mecha-quote">
          <span class="ac-mecha-quote-tag">VOICE LOG // 01</span>
          <p>{{ p.quote || '……通讯静默，尚未收录台词。' }}</p>
          <span class="ac-mecha-wave"><i></i><i></i><i></i><i></i><i></i><i></i><i></i></span>
        </div>

        <footer class="ac-mecha-foot">
          <span>NFC LINK · STABLE</span>
          <span>NEURAL DIVE READY</span>
        </footer>
      </article>

      <p class="ac-mecha-caption">[ 本机数据已加密 · 仅限授权阅览 ]</p>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { Crosshair } from 'lucide-vue-next'

const props = defineProps(['payload', 'nfcId'])
const p = computed(() => props.payload || {})
</script>

<style scoped>
.ac-mecha-root {
  background: #050a10;
  font-family: 'Rajdhani', 'PingFang SC', 'Microsoft YaHei', sans-serif;
}
.ac-mecha-grid {
  position: absolute;
  inset: 0;
  background-image:
    linear-gradient(rgba(34, 211, 238, 0.07) 1px, transparent 1px),
    linear-gradient(90deg, rgba(34, 211, 238, 0.07) 1px, transparent 1px);
  background-size: 32px 32px;
  mask-image: radial-gradient(ellipse at 50% 36%, #000 10%, transparent 78%);
}
.ac-mecha-glow {
  position: absolute;
  top: 80px;
  left: 50%;
  width: 340px;
  height: 340px;
  transform: translateX(-50%);
  border-radius: 50%;
  background: radial-gradient(circle, rgba(34, 211, 238, 0.18), transparent 68%);
  filter: blur(20px);
  animation: ac-mecha-breathe 6s ease-in-out infinite;
}
.ac-mecha-hazard {
  position: absolute;
  left: 0;
  right: 0;
  height: 7px;
  background: repeating-linear-gradient(115deg, #f59e0b 0 9px, #111 9px 18px);
  opacity: 0.65;
}
.ac-mecha-hz-top { top: 0; }
.ac-mecha-hz-bottom { bottom: 0; }
.ac-mecha-bar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  font-family: Consolas, monospace;
  font-size: 8px;
  letter-spacing: 0.2em;
  color: rgba(103, 232, 249, 0.55);
  animation: ac-mecha-fade 0.7s both;
}
.ac-mecha-live { display: inline-flex; align-items: center; gap: 6px; }
.ac-mecha-live i {
  width: 5px; height: 5px; background: #22d3ee;
  box-shadow: 0 0 9px #22d3ee;
  animation: ac-mecha-pulse 1.8s ease-in-out infinite;
}
.ac-mecha-card {
  position: relative;
  margin-top: 30px;
  border: 1px solid rgba(34, 211, 238, 0.32);
  background:
    linear-gradient(150deg, rgba(10, 26, 38, 0.97), rgba(4, 9, 14, 0.99));
  clip-path: polygon(0 0, calc(100% - 22px) 0, 100% 22px, 100% 100%, 22px 100%, 0 calc(100% - 22px));
  padding: 18px 16px 15px;
  box-shadow: 0 28px 70px -30px rgba(34, 211, 238, 0.4);
  animation: ac-mecha-enter 0.9s cubic-bezier(0.2, 1, 0.25, 1) both;
  transition: transform 0.45s ease, border-color 0.45s ease;
}
.ac-mecha-card:hover {
  transform: translateY(-4px) rotateX(3deg) rotateY(-3deg);
  border-color: rgba(245, 158, 11, 0.6);
}
.ac-mecha-card::before {
  content: '';
  position: absolute;
  inset: 6px;
  border: 1px solid rgba(34, 211, 238, 0.12);
  clip-path: polygon(0 0, calc(100% - 18px) 0, 100% 18px, 100% 100%, 18px 100%, 0 calc(100% - 18px));
  pointer-events: none;
}
.ac-mecha-frame-line {
  position: absolute;
  top: 0;
  left: 16px;
  width: 84px;
  height: 3px;
  background: #f59e0b;
  box-shadow: 0 0 12px rgba(245, 158, 11, 0.6);
}
.ac-mecha-notch {
  position: absolute;
  width: 7px;
  height: 7px;
  border: 1px solid #22d3ee;
}
.ac-mecha-n1 { top: 7px; right: 34px; }
.ac-mecha-n2 { bottom: 7px; left: 34px; }
.ac-mecha-id {
  display: flex;
  justify-content: space-between;
  align-items: baseline;
  border-bottom: 1px solid rgba(34, 211, 238, 0.15);
  padding-bottom: 10px;
  font-family: Consolas, monospace;
  font-size: 7.5px;
  letter-spacing: 0.16em;
}
.ac-mecha-serial { color: #f59e0b; font-size: 9px; }
.ac-mecha-class { color: rgba(255, 255, 255, 0.3); }
.ac-mecha-hero { position: relative; padding-top: 26px; }
.ac-mecha-ghost {
  position: absolute;
  top: 4px;
  right: -4px;
  font-family: 'Arial Black', Impact, sans-serif;
  font-size: 104px;
  line-height: 1;
  color: transparent;
  -webkit-text-stroke: 1px rgba(34, 211, 238, 0.16);
  pointer-events: none;
}
.ac-mecha-reticle {
  display: flex;
  width: 40px;
  height: 40px;
  align-items: center;
  justify-content: center;
  border: 1px solid rgba(34, 211, 238, 0.4);
  color: #22d3ee;
  animation: ac-mecha-spin 14s linear infinite;
}
.ac-mecha-role {
  margin-top: 14px;
  font-family: Consolas, monospace;
  font-size: 8px;
  letter-spacing: 0.32em;
  color: rgba(34, 211, 238, 0.6);
}
.ac-mecha-name {
  margin-top: 4px;
  font-size: 38px;
  font-weight: 700;
  line-height: 1.05;
  letter-spacing: 0.03em;
  color: #f0fdff;
  text-shadow: 0 0 26px rgba(34, 211, 238, 0.4);
  overflow-wrap: anywhere;
}
.ac-mecha-title {
  margin-top: 7px;
  font-size: 12.5px;
  font-weight: 500;
  letter-spacing: 0.2em;
  color: #fbbf24;
}
.ac-mecha-tags { display: flex; flex-wrap: wrap; gap: 8px; margin-top: 18px; }
.ac-mecha-tags span {
  display: inline-flex;
  flex-direction: column;
  gap: 2px;
  border-left: 2px solid #22d3ee;
  background: rgba(34, 211, 238, 0.06);
  padding: 5px 10px;
  font-size: 10.5px;
  color: rgba(223, 246, 255, 0.82);
}
.ac-mecha-tags b {
  font-family: Consolas, monospace;
  font-size: 6.5px;
  font-weight: 400;
  letter-spacing: 0.2em;
  color: rgba(34, 211, 238, 0.55);
}
.ac-mecha-elem { border-left-color: #f59e0b !important; }
.ac-mecha-elem b { color: rgba(245, 158, 11, 0.65) !important; }
.ac-mecha-telemetry {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 7px;
  margin-top: 20px;
}
.ac-mecha-cell {
  border: 1px solid rgba(34, 211, 238, 0.14);
  background: rgba(255, 255, 255, 0.015);
  padding: 8px 9px;
}
.ac-mecha-cell:nth-child(3) { grid-column: span 2; }
.ac-mecha-cell > span {
  display: block;
  font-family: Consolas, monospace;
  font-size: 6.5px;
  letter-spacing: 0.18em;
  color: rgba(255, 255, 255, 0.32);
}
.ac-mecha-cell b {
  display: block;
  margin-top: 2px;
  font-size: 16px;
  font-weight: 600;
  letter-spacing: 0.04em;
  color: #e0faff;
  overflow-wrap: anywhere;
}
.ac-mecha-cell b em {
  font-style: normal;
  font-size: 8px;
  color: rgba(255, 255, 255, 0.34);
}
.ac-mecha-cell i {
  display: block;
  height: 2px;
  margin-top: 6px;
  background: rgba(34, 211, 238, 0.12);
}
.ac-mecha-cell i u {
  display: block;
  height: 100%;
  background: linear-gradient(90deg, #0891b2, #67e8f9);
  box-shadow: 0 0 8px rgba(34, 211, 238, 0.7);
}
.ac-mecha-comms { font-family: Consolas, monospace; font-size: 12px !important; color: #fbbf24 !important; }
.ac-mecha-quote {
  position: relative;
  margin-top: 16px;
  border: 1px solid rgba(245, 158, 11, 0.22);
  background: linear-gradient(90deg, rgba(245, 158, 11, 0.08), transparent 60%);
  padding: 11px 12px 13px;
}
.ac-mecha-quote-tag {
  font-family: Consolas, monospace;
  font-size: 6.5px;
  letter-spacing: 0.22em;
  color: rgba(245, 158, 11, 0.7);
}
.ac-mecha-quote p {
  margin-top: 5px;
  font-family: 'PingFang SC', 'Microsoft YaHei', sans-serif;
  font-size: 13px;
  line-height: 1.8;
  color: rgba(223, 246, 255, 0.9);
  overflow-wrap: anywhere;
}
.ac-mecha-wave {
  display: flex;
  align-items: flex-end;
  gap: 3px;
  height: 12px;
  margin-top: 9px;
}
.ac-mecha-wave i { width: 2px; background: #22d3ee; opacity: 0.55; }
.ac-mecha-wave i:nth-child(1) { height: 4px; }
.ac-mecha-wave i:nth-child(2) { height: 9px; }
.ac-mecha-wave i:nth-child(3) { height: 6px; }
.ac-mecha-wave i:nth-child(4) { height: 12px; }
.ac-mecha-wave i:nth-child(5) { height: 5px; }
.ac-mecha-wave i:nth-child(6) { height: 8px; }
.ac-mecha-wave i:nth-child(7) { height: 3px; }
.ac-mecha-foot {
  display: flex;
  justify-content: space-between;
  border-top: 1px solid rgba(34, 211, 238, 0.13);
  margin-top: 14px;
  padding-top: 9px;
  font-family: Consolas, monospace;
  font-size: 6.5px;
  letter-spacing: 0.14em;
  color: rgba(255, 255, 255, 0.24);
}
.ac-mecha-caption {
  margin-top: 20px;
  text-align: center;
  font-size: 9px;
  letter-spacing: 0.22em;
  color: rgba(103, 232, 249, 0.28);
  animation: ac-mecha-fade 1.2s 0.6s both;
}
@keyframes ac-mecha-enter {
  from { opacity: 0; transform: translateY(28px) scale(0.97); clip-path: polygon(0 0, 100% 0, 100% 0, 100% 0, 0 100%, 0 100%); }
  to { opacity: 1; transform: none; clip-path: polygon(0 0, calc(100% - 22px) 0, 100% 22px, 100% 100%, 22px 100%, 0 calc(100% - 22px)); }
}
@keyframes ac-mecha-fade { from { opacity: 0; } to { opacity: 1; } }
@keyframes ac-mecha-pulse { 0%, 100% { opacity: 1; } 50% { opacity: 0.25; } }
@keyframes ac-mecha-breathe { 0%, 100% { opacity: 0.6; transform: translateX(-50%) scale(1); } 50% { opacity: 1; transform: translateX(-50%) scale(1.08); } }
@keyframes ac-mecha-spin { to { transform: rotate(360deg); } }
@media (prefers-reduced-motion: reduce) {
  .ac-mecha-root *,
  .ac-mecha-root *::before,
  .ac-mecha-root *::after {
    animation: none !important;
    transition: none !important;
  }
}
</style>
