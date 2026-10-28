<template>
  <div class="ac-cass-root relative min-h-screen overflow-hidden px-5 py-9 text-[#f6ead2]">
    <div class="ac-cass-glow"></div>
    <div class="ac-cass-rings"></div>
    <div class="relative z-10 mx-auto w-full max-w-md">
      <div class="ac-cass-shell">
        <div class="ac-cass-screws">
          <i v-for="n in 4" :key="n"></i>
        </div>

        <div class="ac-cass-label">
          <div class="ac-cass-label-head">
            <span>STAR OCEAN RECORDS</span>
            <span>STEREO · TYPE II</span>
          </div>

          <div class="ac-cass-label-main">
            <div class="ac-cass-side">SIDE<br /><b>A</b></div>
            <div class="ac-cass-title">
              <p class="ac-cass-kicker">ANIME CARD SESSIONS · VOL. {{ nfcId || '01' }}</p>
              <h1 class="ac-cass-name">{{ p.name || '未命名角色' }}</h1>
              <p class="ac-cass-sub">{{ p.title }}</p>
            </div>
            <div class="ac-cass-stamp">LIMITED</div>
          </div>

          <div class="ac-cass-tracks">
            <p><span>A1</span><b>{{ p.series || '未知作品' }}</b><em>4:12</em></p>
            <p><span>A2</span><b>{{ p.element }}</b><em>3:48</em></p>
            <p><span>A3</span><b>@{{ p.social }}</b><em>∞</em></p>
          </div>

          <p class="ac-cass-note">{{ p.quote || '……录下这一句，循环一整夜。' }}</p>
        </div>

        <div class="ac-cass-window">
          <div class="ac-cass-reel ac-cass-reel-l">
            <span class="ac-cass-hub"></span>
          </div>
          <div class="ac-cass-tape-line"></div>
          <div class="ac-cass-reel ac-cass-reel-r">
            <span class="ac-cass-hub"></span>
          </div>
          <span class="ac-cass-window-text">HIGH BIAS · 120μs</span>
        </div>

        <div class="ac-cass-bottom">
          <span>◀ REC</span>
          <span>▷ PLAY</span>
          <span>NFC {{ nfcId || '0000' }}</span>
        </div>
      </div>

      <p class="ac-cass-caption">PRESS PLAY · REWIND THE MEMORY</p>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps(['payload', 'nfcId'])
const p = computed(() => props.payload || {})
</script>

<style scoped>
.ac-cass-root {
  background: #2b1710;
  font-family: 'Futura', 'PingFang SC', 'Microsoft YaHei', sans-serif;
}
.ac-cass-glow {
  position: absolute;
  top: -20%;
  left: 50%;
  width: 560px;
  height: 560px;
  transform: translateX(-50%);
  background: radial-gradient(circle, rgba(251, 146, 60, 0.3), transparent 64%);
  filter: blur(24px);
}
.ac-cass-rings {
  position: absolute;
  inset: 0;
  background-image: repeating-radial-gradient(circle at 50% 50%, transparent 0 38px, rgba(251, 146, 60, 0.07) 38px 39px);
  mask-image: radial-gradient(ellipse at 50% 45%, #000 10%, transparent 75%);
}
.ac-cass-shell {
  position: relative;
  border-radius: 16px;
  border: 1px solid #5a4437;
  background: linear-gradient(160deg, #3a2b24, #221712 58%, #30211a);
  padding: 16px 15px 14px;
  box-shadow:
    0 30px 60px -24px rgba(0, 0, 0, 0.8),
    inset 0 2px 0 rgba(255, 255, 255, 0.09),
    inset 0 -3px 10px rgba(0, 0, 0, 0.6);
  animation: ac-cass-enter 0.95s cubic-bezier(0.2, 1, 0.3, 1) both;
  transition: transform 0.5s ease, box-shadow 0.5s ease;
}
.ac-cass-shell:hover {
  transform: translateY(-5px) rotate(-1deg);
  box-shadow:
    0 40px 74px -24px rgba(0, 0, 0, 0.85),
    inset 0 2px 0 rgba(255, 255, 255, 0.12),
    inset 0 -3px 10px rgba(0, 0, 0, 0.6);
}
.ac-cass-screws i {
  position: absolute;
  width: 7px;
  height: 7px;
  border-radius: 50%;
  background: radial-gradient(circle at 35% 30%, #8a7262, #3b2a20);
  box-shadow: inset 0 0 0 1px rgba(0, 0, 0, 0.6);
}
.ac-cass-screws i:nth-child(1) { top: 7px; left: 7px; }
.ac-cass-screws i:nth-child(2) { top: 7px; right: 7px; }
.ac-cass-screws i:nth-child(3) { bottom: 7px; left: 7px; }
.ac-cass-screws i:nth-child(4) { bottom: 7px; right: 7px; }
.ac-cass-label {
  border-radius: 5px;
  background: #f6ead2;
  padding: 12px 14px 13px;
  color: #332016;
  box-shadow: inset 0 0 0 1px rgba(120, 80, 50, 0.16);
}
.ac-cass-label-head {
  display: flex;
  justify-content: space-between;
  border-bottom: 1px solid #d8c3a2;
  padding-bottom: 6px;
  font-family: Consolas, monospace;
  font-size: 6.5px;
  letter-spacing: 0.13em;
  color: #9b7449;
}
.ac-cass-label-main {
  display: flex;
  align-items: flex-start;
  gap: 10px;
  margin-top: 11px;
}
.ac-cass-side {
  flex: none;
  border: 1.5px solid #b45309;
  padding: 3px 5px;
  font-family: Consolas, monospace;
  font-size: 5.5px;
  line-height: 1.25;
  letter-spacing: 0.08em;
  color: #b45309;
  text-align: center;
}
.ac-cass-side b { font-size: 14px; }
.ac-cass-title { min-width: 0; flex: 1; }
.ac-cass-kicker {
  font-family: Consolas, monospace;
  font-size: 6px;
  letter-spacing: 0.14em;
  color: #b08a5c;
}
.ac-cass-name {
  margin-top: 2px;
  font-family: 'Arial Black', Impact, 'PingFang SC', sans-serif;
  font-size: 31px;
  font-weight: 900;
  line-height: 1.02;
  letter-spacing: -0.01em;
  color: #3b2418;
  overflow-wrap: anywhere;
}
.ac-cass-sub { margin-top: 4px; font-size: 10px; letter-spacing: 0.16em; color: #a16207; }
.ac-cass-stamp {
  flex: none;
  border: 1.5px solid #c2410c;
  border-radius: 50%;
  padding: 6px 3px;
  font-family: Consolas, monospace;
  font-size: 5px;
  letter-spacing: 0.08em;
  color: #c2410c;
  transform: rotate(12deg);
  opacity: 0.8;
}
.ac-cass-tracks {
  margin-top: 12px;
  border-top: 1px solid #d8c3a2;
  border-bottom: 1px solid #d8c3a2;
  padding: 5px 0;
}
.ac-cass-tracks p {
  display: flex;
  align-items: baseline;
  gap: 7px;
  padding: 2px 0;
  font-size: 9.5px;
  line-height: 1.5;
}
.ac-cass-tracks span {
  flex: none;
  width: 18px;
  font-family: Consolas, monospace;
  font-size: 7px;
  color: #b45309;
}
.ac-cass-tracks b {
  min-width: 0;
  flex: 1;
  font-weight: 600;
  color: #4a2f1e;
  overflow-wrap: anywhere;
}
.ac-cass-tracks em {
  flex: none;
  font-family: Consolas, monospace;
  font-style: normal;
  font-size: 7px;
  color: #b08a5c;
}
.ac-cass-note {
  margin-top: 9px;
  font-family: 'Songti SC', 'Noto Serif SC', serif;
  font-size: 12px;
  font-style: italic;
  line-height: 1.7;
  color: #6b4a32;
  overflow-wrap: anywhere;
}
.ac-cass-window {
  position: relative;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 22px;
  height: 80px;
  margin-top: 13px;
  border-radius: 6px;
  border: 1px solid #4a382c;
  background: linear-gradient(180deg, #120c09, #241a14);
  box-shadow: inset 0 3px 10px rgba(0, 0, 0, 0.7), inset 0 -1px 0 rgba(255, 255, 255, 0.08);
}
.ac-cass-reel {
  position: relative;
  display: flex;
  width: 48px;
  height: 48px;
  align-items: center;
  justify-content: center;
  border-radius: 50%;
  background:
    repeating-conic-gradient(#e7d6b8 0deg 12deg, #6b5642 12deg 24deg);
  border: 3px solid #d9c8aa;
  box-shadow: 0 0 0 2px #1a120e, inset 0 0 8px rgba(0, 0, 0, 0.5);
  animation: ac-cass-spin 6s linear infinite;
}
.ac-cass-reel-r { animation-direction: reverse; animation-duration: 7.5s; }
.ac-cass-hub {
  width: 16px;
  height: 16px;
  border-radius: 50%;
  background: #2a1e16;
  border: 2px solid #d9c8aa;
}
.ac-cass-tape-line {
  position: absolute;
  left: 50%;
  bottom: 20px;
  width: 84px;
  height: 1px;
  transform: translateX(-50%);
  background: rgba(214, 190, 158, 0.35);
}
.ac-cass-window-text {
  position: absolute;
  bottom: 5px;
  left: 50%;
  transform: translateX(-50%);
  font-family: Consolas, monospace;
  font-size: 5.5px;
  letter-spacing: 0.2em;
  color: rgba(231, 214, 184, 0.4);
  white-space: nowrap;
}
.ac-cass-bottom {
  display: flex;
  justify-content: space-between;
  margin-top: 10px;
  padding: 0 4px;
  font-family: Consolas, monospace;
  font-size: 6.5px;
  letter-spacing: 0.14em;
  color: rgba(246, 234, 210, 0.38);
}
.ac-cass-caption {
  margin-top: 24px;
  text-align: center;
  font-family: Consolas, monospace;
  font-size: 7px;
  letter-spacing: 0.28em;
  color: rgba(251, 146, 60, 0.4);
  animation: ac-cass-fade 1.3s 0.6s both;
}
@keyframes ac-cass-enter {
  from { opacity: 0; transform: translateY(36px) rotate(1.5deg) scale(0.96); }
  to { opacity: 1; transform: none; }
}
@keyframes ac-cass-fade { from { opacity: 0; } to { opacity: 1; } }
@keyframes ac-cass-spin { to { transform: rotate(360deg); } }
@media (prefers-reduced-motion: reduce) {
  .ac-cass-root *,
  .ac-cass-root *::before,
  .ac-cass-root *::after {
    animation: none !important;
    transition: none !important;
  }
}
</style>
