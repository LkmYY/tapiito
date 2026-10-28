<template>
  <div class="ac-gb-root relative min-h-screen overflow-hidden px-5 py-9">
    <div class="ac-gb-bg"></div>
    <div class="relative z-10 mx-auto w-full max-w-md">
      <div class="ac-gb-shell">
        <div class="ac-gb-topline">
          <span>8-BIT CHARACTER SYSTEM</span>
          <span>MODEL AC-01</span>
        </div>

        <div class="ac-gb-bezel">
          <div class="ac-gb-bezel-label">
            <span>DOT MATRIX DATA LINK</span>
            <span class="ac-gb-power"><i></i>POWER</span>
          </div>
          <div class="ac-gb-screen">
            <div class="ac-gb-lcd">
              <div class="ac-gb-status">
                <span>LV.99</span>
                <span class="ac-gb-hp">HP <b>▮▮▮▮▮</b></span>
              </div>

              <p class="ac-gb-kicker">◆ CHARACTER DATA ◆</p>
              <h1 class="ac-gb-name">{{ p.name || '未命名角色' }}</h1>
              <p class="ac-gb-title">{{ p.title }}</p>

              <div class="ac-gb-rows">
                <p><span>作品</span>{{ p.series || '未知作品' }}</p>
                <p><span>属性</span>{{ p.element }}</p>
              </div>

              <div class="ac-gb-dialog">
                <p>{{ p.quote || '……还没有录入台词。' }}</p>
                <span class="ac-gb-cursor">▼</span>
              </div>

              <div class="ac-gb-foot">
                <span>@{{ p.social }}</span>
                <span>ID:{{ nfcId || '0000' }}</span>
              </div>
            </div>
            <div class="ac-gb-glass"></div>
          </div>
        </div>

        <p class="ac-gb-brand">ANIME <b>POCKET</b></p>

        <div class="ac-gb-controls">
          <div class="ac-gb-dpad">
            <i class="ac-gb-dpad-h"></i>
            <i class="ac-gb-dpad-v"></i>
            <span class="ac-gb-dpad-center"></span>
          </div>

          <div class="ac-gb-ab">
            <div class="ac-gb-btn"><span>B</span></div>
            <div class="ac-gb-btn"><span>A</span></div>
          </div>
        </div>

        <div class="ac-gb-selects">
          <span><i></i>SELECT</span>
          <span><i></i>START</span>
        </div>

        <div class="ac-gb-speaker">
          <i v-for="n in 6" :key="n"></i>
        </div>
      </div>

      <p class="ac-gb-caption">INSERT CARD · PRESS START · NFC {{ nfcId || '0000' }}</p>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps(['payload', 'nfcId'])
const p = computed(() => props.payload || {})
</script>

<style scoped>
.ac-gb-root {
  background: #131b16;
  font-family: 'Courier New', Consolas, monospace;
}
.ac-gb-bg {
  position: absolute;
  inset: 0;
  background-image:
    linear-gradient(rgba(155, 188, 15, 0.05) 1px, transparent 1px),
    linear-gradient(90deg, rgba(155, 188, 15, 0.05) 1px, transparent 1px);
  background-size: 18px 18px;
  mask-image: radial-gradient(ellipse at 50% 40%, #000 20%, transparent 78%);
}
.ac-gb-shell {
  position: relative;
  border-radius: 14px 14px 46px 14px;
  background: linear-gradient(155deg, #d6d3c0, #b9b6a2 58%, #a5a28f);
  border: 2px solid #8f8c7a;
  padding: 14px 14px 18px;
  box-shadow:
    4px 5px 0 #6f6c5e,
    inset 0 2px 0 rgba(255, 255, 255, 0.65),
    inset 0 -4px 12px rgba(0, 0, 0, 0.16),
    0 26px 50px -22px rgba(0, 0, 0, 0.75);
  animation: ac-gb-on 1s steps(5) both;
  transition: transform 0.45s ease, box-shadow 0.45s ease;
}
.ac-gb-shell:hover {
  transform: translate(-2px, -3px) rotate(-0.5deg);
  box-shadow:
    7px 8px 0 #6f6c5e,
    inset 0 2px 0 rgba(255, 255, 255, 0.65),
    inset 0 -4px 12px rgba(0, 0, 0, 0.16),
    0 32px 58px -22px rgba(0, 0, 0, 0.8);
}
.ac-gb-topline {
  display: flex;
  justify-content: space-between;
  padding: 0 3px 9px;
  font-size: 6.5px;
  letter-spacing: 0.14em;
  color: #6b6859;
}
.ac-gb-bezel {
  border-radius: 9px 9px 30px 9px;
  background: linear-gradient(180deg, #5c5c54, #45453f);
  padding: 12px 14px 18px;
  box-shadow: inset 0 2px 8px rgba(0, 0, 0, 0.45), 0 1px 0 rgba(255, 255, 255, 0.3);
}
.ac-gb-bezel-label {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 0 2px 7px;
  font-size: 5.5px;
  letter-spacing: 0.12em;
  color: #c9c6b4;
}
.ac-gb-power { display: inline-flex; align-items: center; gap: 5px; color: #a3a08e; }
.ac-gb-power i {
  width: 5px; height: 5px; border-radius: 50%;
  background: #ef4444;
  box-shadow: 0 0 6px #ef4444;
  animation: ac-gb-blink 2.4s steps(2) infinite;
}
.ac-gb-screen {
  position: relative;
  border-radius: 3px;
  border: 2px solid #2f3a22;
  background: #8bac0f;
  padding: 8px;
  box-shadow: inset 0 0 22px rgba(15, 56, 15, 0.55);
}
.ac-gb-screen::after {
  content: '';
  position: absolute;
  inset: 0;
  background: repeating-linear-gradient(180deg, transparent 0 2px, rgba(15, 56, 15, 0.13) 2px 3px);
  pointer-events: none;
}
.ac-gb-lcd {
  position: relative;
  z-index: 1;
  border: 2px solid #0f380f;
  background: #9bbc0f;
  padding: 8px 9px 7px;
  box-shadow: inset 0 0 0 2px rgba(155, 188, 15, 0.7);
}
.ac-gb-status {
  display: flex;
  justify-content: space-between;
  border-bottom: 1px solid #306230;
  padding-bottom: 4px;
  font-size: 7px;
  font-weight: 700;
  color: #0f380f;
}
.ac-gb-hp b { font-weight: 400; letter-spacing: -0.06em; }
.ac-gb-kicker {
  margin-top: 7px;
  text-align: center;
  font-size: 6.5px;
  letter-spacing: 0.16em;
  color: #306230;
}
.ac-gb-name {
  margin-top: 3px;
  text-align: center;
  font-family: 'Courier New', Consolas, monospace;
  font-size: 27px;
  font-weight: 700;
  line-height: 1.1;
  letter-spacing: 0.02em;
  color: #0f380f;
  text-shadow: 1px 1px 0 rgba(48, 98, 48, 0.35);
  overflow-wrap: anywhere;
}
.ac-gb-title {
  margin-top: 2px;
  text-align: center;
  font-size: 9px;
  letter-spacing: 0.1em;
  color: #306230;
}
.ac-gb-rows {
  margin-top: 8px;
  border-top: 1px solid #306230;
  border-bottom: 1px solid #306230;
  padding: 4px 0;
}
.ac-gb-rows p {
  display: flex;
  gap: 7px;
  padding: 2px 1px;
  font-size: 8px;
  line-height: 1.5;
  color: #0f380f;
  overflow-wrap: anywhere;
}
.ac-gb-rows span {
  flex: none;
  width: 26px;
  color: #306230;
}
.ac-gb-dialog {
  position: relative;
  margin-top: 7px;
  border: 2px solid #0f380f;
  background: #9bbc0f;
  padding: 7px 9px 11px;
  box-shadow: 2px 2px 0 rgba(48, 98, 48, 0.4);
}
.ac-gb-dialog p {
  font-family: 'PingFang SC', 'Microsoft YaHei', monospace;
  font-size: 9.5px;
  font-weight: 600;
  line-height: 1.7;
  color: #0f380f;
  overflow-wrap: anywhere;
}
.ac-gb-cursor {
  position: absolute;
  right: 4px;
  bottom: 1px;
  font-size: 7px;
  color: #0f380f;
  animation: ac-gb-blink 1s steps(2) infinite;
}
.ac-gb-foot {
  display: flex;
  justify-content: space-between;
  gap: 8px;
  margin-top: 6px;
  font-size: 6.5px;
  color: #306230;
}
.ac-gb-foot span { overflow-wrap: anywhere; }
.ac-gb-glass {
  position: absolute;
  inset: 0;
  background: linear-gradient(125deg, rgba(255, 255, 255, 0.22), transparent 36%);
  pointer-events: none;
}
.ac-gb-brand {
  margin-top: 13px;
  padding-left: 5px;
  font-family: 'Arial Black', Impact, sans-serif;
  font-size: 11px;
  font-style: italic;
  letter-spacing: 0.08em;
  color: #4a4a42;
}
.ac-gb-brand b { color: #7c3aed; }
.ac-gb-controls {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 12px 8px 0;
}
.ac-gb-dpad {
  position: relative;
  width: 70px;
  height: 70px;
}
.ac-gb-dpad-h {
  position: absolute;
  top: 24px;
  left: 0;
  width: 70px;
  height: 22px;
  border-radius: 3px;
  background: #2e2e2b;
  box-shadow: inset 0 2px 2px rgba(255, 255, 255, 0.08), 2px 3px 0 rgba(0, 0, 0, 0.25);
}
.ac-gb-dpad-v {
  position: absolute;
  top: 0;
  left: 24px;
  width: 22px;
  height: 70px;
  border-radius: 3px;
  background: #2e2e2b;
  box-shadow: inset 2px 0 2px rgba(255, 255, 255, 0.06), 2px 3px 0 rgba(0, 0, 0, 0.25);
}
.ac-gb-dpad-center {
  position: absolute;
  top: 31px;
  left: 31px;
  width: 8px;
  height: 8px;
  border-radius: 50%;
  background: #1c1c1a;
}
.ac-gb-ab {
  display: flex;
  gap: 11px;
  align-items: flex-end;
  transform: rotate(-16deg);
}
.ac-gb-btn {
  display: flex;
  width: 31px;
  height: 31px;
  align-items: center;
  justify-content: center;
  border-radius: 50%;
  background: radial-gradient(circle at 34% 28%, #d9468f, #9d174d 70%);
  box-shadow: 1px 3px 0 #6b1035, inset 0 1px 0 rgba(255, 255, 255, 0.35);
  color: #fce7f3;
  font-size: 8px;
  font-weight: 700;
}
.ac-gb-btn span { transform: rotate(16deg); }
.ac-gb-selects {
  display: flex;
  justify-content: center;
  gap: 22px;
  margin-top: 14px;
}
.ac-gb-selects span {
  display: inline-flex;
  align-items: center;
  gap: 5px;
  font-size: 5.5px;
  letter-spacing: 0.08em;
  color: #5d5a4d;
}
.ac-gb-selects i {
  width: 17px;
  height: 5px;
  border-radius: 3px;
  background: #6b6859;
  box-shadow: 1px 1px 0 rgba(255, 255, 255, 0.35);
}
.ac-gb-speaker {
  display: flex;
  justify-content: flex-end;
  gap: 5px;
  margin-top: 16px;
  padding-right: 10px;
}
.ac-gb-speaker i {
  width: 3px;
  height: 3px;
  border-radius: 50%;
  background: #7a7768;
  box-shadow: inset 0 1px 1px rgba(0, 0, 0, 0.3);
}
.ac-gb-speaker i:nth-child(even) { transform: translateY(4px); }
.ac-gb-caption {
  margin-top: 23px;
  text-align: center;
  font-size: 7px;
  letter-spacing: 0.18em;
  color: rgba(155, 188, 15, 0.4);
  animation: ac-gb-fade 1.4s 0.7s both;
}
@keyframes ac-gb-on {
  0% { opacity: 0; filter: brightness(2.5) contrast(0.2); transform: translateY(18px); }
  40% { opacity: 1; filter: brightness(1.6); }
  60% { filter: brightness(0.9); }
  100% { opacity: 1; filter: none; transform: none; }
}
@keyframes ac-gb-fade { from { opacity: 0; } to { opacity: 1; } }
@keyframes ac-gb-blink { 0%, 49% { opacity: 1; } 50%, 100% { opacity: 0.15; } }
@media (prefers-reduced-motion: reduce) {
  .ac-gb-root *,
  .ac-gb-root *::before,
  .ac-gb-root *::after {
    animation: none !important;
    transition: none !important;
  }
}
</style>
