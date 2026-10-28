<template>
  <div class="ac-pixel-root relative min-h-screen overflow-hidden px-4 py-8 text-[#e6e8ff]">
    <div class="ac-pixel-stars"></div>
    <div class="ac-pixel-ground"></div>

    <div class="relative z-10 mx-auto w-full max-w-md">
      <div class="ac-pixel-window">
        <div class="ac-pixel-titlebar">
          <span>◆ PARTY STATUS ◆</span>
          <span class="ac-pixel-batt">BATTLE {{ nfcId || '0000' }}</span>
        </div>

        <div class="ac-pixel-body">
          <div class="ac-pixel-top">
            <div class="ac-pixel-portrait">
              <span class="ac-pixel-face">{{ initial }}</span>
              <i class="ac-pixel-scan"></i>
              <span class="ac-pixel-lv">Lv.99</span>
            </div>

            <div class="ac-pixel-ident">
              <p class="ac-pixel-class">／ HERO CLASS ／</p>
              <h1 class="ac-pixel-name">{{ p.name || '未命名角色' }}</h1>
              <p class="ac-pixel-job">{{ p.title }}</p>
              <div class="ac-pixel-bars">
                <div class="ac-pixel-bar">
                  <span>HP</span><i><b style="width: 92%"></b></i><em>920</em>
                </div>
                <div class="ac-pixel-bar ac-pixel-mp">
                  <span>MP</span><i><b style="width: 78%"></b></i><em>460</em>
                </div>
              </div>
            </div>
          </div>

          <div class="ac-pixel-stats">
            <div class="ac-pixel-row">
              <span class="ac-pixel-key">ORIGIN</span>
              <span class="ac-pixel-val">{{ p.series || '未知作品' }}</span>
            </div>
            <div class="ac-pixel-row">
              <span class="ac-pixel-key">ELEMENT</span>
              <span class="ac-pixel-val ac-pixel-elem">{{ p.element }}</span>
            </div>
            <div class="ac-pixel-row">
              <span class="ac-pixel-key">SOCIAL</span>
              <span class="ac-pixel-val ac-pixel-social">@{{ p.social }}</span>
            </div>
          </div>

          <div class="ac-pixel-dialog">
            <span class="ac-pixel-speaker">{{ p.name || '角色' }}：</span>
            <p>{{ p.quote || '……还没有录入台词。' }}</p>
            <span class="ac-pixel-next">▼</span>
          </div>

          <div class="ac-pixel-actions">
            <span>◀ 战斗</span>
            <span>道具 ▶</span>
            <span>★ 收藏</span>
          </div>
        </div>
      </div>

      <p class="ac-pixel-hint">PRESS START · NFC LINK ESTABLISHED</p>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps(['payload', 'nfcId'])
const p = computed(() => props.payload || {})
const initial = computed(() => (p.value.name || '？').toString().trim().slice(0, 1) || '？')
</script>

<style scoped>
.ac-pixel-root {
  background: #0b0b22;
  font-family: 'Courier New', Consolas, monospace;
}
.ac-pixel-stars {
  position: absolute;
  inset: 0;
  background-image:
    radial-gradient(#a5b4fc 1px, transparent 1px),
    radial-gradient(#f0abfc 1px, transparent 1px);
  background-size: 44px 44px, 67px 67px;
  background-position: 0 0, 21px 33px;
  opacity: 0.35;
  animation: ac-pixel-twinkle 4.5s steps(2) infinite;
}
.ac-pixel-ground {
  position: absolute;
  left: 0;
  right: 0;
  bottom: 0;
  height: 90px;
  background:
    repeating-linear-gradient(90deg, transparent 0 28px, rgba(129, 140, 248, 0.1) 28px 30px),
    linear-gradient(180deg, transparent, #0a0a1e);
}
.ac-pixel-window {
  border: 4px solid #c7d2fe;
  background: #141438;
  box-shadow:
    0 0 0 4px #1e1b4b,
    0 0 0 8px #6366f1,
    12px 12px 0 rgba(0, 0, 0, 0.55);
  animation: ac-pixel-enter 0.8s steps(6) both;
}
.ac-pixel-titlebar {
  display: flex;
  align-items: center;
  justify-content: space-between;
  background: #6366f1;
  padding: 7px 9px;
  color: #0b0b22;
  font-size: 9px;
  font-weight: 700;
  letter-spacing: 0.14em;
}
.ac-pixel-batt { color: #1e1b4b; font-size: 8px; }
.ac-pixel-body { padding: 14px 13px 12px; }
.ac-pixel-top {
  display: flex;
  gap: 13px;
  border-bottom: 2px dashed #3f3f8f;
  padding-bottom: 14px;
}
.ac-pixel-portrait {
  position: relative;
  flex: none;
  width: 92px;
  height: 102px;
  border: 3px solid #a5b4fc;
  background:
    linear-gradient(180deg, #312e81, #1e1b4b 55%, #312e81);
  box-shadow: inset 0 0 0 3px #0b0b22;
  overflow: hidden;
}
.ac-pixel-face {
  position: absolute;
  inset: 0;
  display: flex;
  align-items: center;
  justify-content: center;
  font-family: 'Songti SC', serif;
  font-size: 42px;
  font-weight: 700;
  color: #e0e7ff;
  text-shadow: 3px 3px 0 #4338ca, 6px 6px 0 #1e1b4b;
}
.ac-pixel-scan {
  position: absolute;
  inset: 0;
  background: repeating-linear-gradient(180deg, transparent 0 3px, rgba(0, 0, 0, 0.24) 3px 5px);
}
.ac-pixel-lv {
  position: absolute;
  right: 2px;
  bottom: 2px;
  background: #fbbf24;
  padding: 1px 3px;
  color: #271a00;
  font-size: 8px;
  font-weight: 700;
  letter-spacing: 0.04em;
}
.ac-pixel-ident { min-width: 0; flex: 1; }
.ac-pixel-class { font-size: 7px; letter-spacing: 0.2em; color: #818cf8; }
.ac-pixel-name {
  margin-top: 5px;
  font-family: 'PingFang SC', 'Microsoft YaHei', sans-serif;
  font-size: 25px;
  font-weight: 800;
  line-height: 1.15;
  letter-spacing: 0.04em;
  color: #f5f3ff;
  text-shadow: 2px 2px 0 #4338ca;
  overflow-wrap: anywhere;
}
.ac-pixel-job {
  margin-top: 5px;
  font-size: 10px;
  letter-spacing: 0.1em;
  color: #c4b5fd;
}
.ac-pixel-bars { margin-top: 13px; display: grid; gap: 6px; }
.ac-pixel-bar { display: flex; align-items: center; gap: 5px; }
.ac-pixel-bar span { width: 18px; font-size: 7px; font-weight: 700; color: #fda4af; }
.ac-pixel-bar i {
  height: 9px;
  flex: 1;
  border: 2px solid #c7d2fe;
  background: #0b0b22;
  padding: 1px;
}
.ac-pixel-bar b { display: block; height: 100%; background: repeating-linear-gradient(90deg, #f87171 0 5px, #ef4444 5px 7px); }
.ac-pixel-bar em { font-style: normal; font-size: 7px; color: #a5b4fc; }
.ac-pixel-mp span { color: #67e8f9; }
.ac-pixel-mp b { background: repeating-linear-gradient(90deg, #22d3ee 0 5px, #0891b2 5px 7px); }
.ac-pixel-stats { margin-top: 13px; display: grid; gap: 5px; }
.ac-pixel-row {
  display: flex;
  align-items: baseline;
  gap: 9px;
  border-bottom: 1px dotted #4338ca;
  padding: 6px 2px;
}
.ac-pixel-key {
  flex: none;
  width: 66px;
  font-size: 7px;
  letter-spacing: 0.12em;
  color: #818cf8;
}
.ac-pixel-val {
  min-width: 0;
  flex: 1;
  text-align: right;
  font-family: 'PingFang SC', 'Microsoft YaHei', sans-serif;
  font-size: 11.5px;
  color: #e0e7ff;
  overflow-wrap: anywhere;
}
.ac-pixel-elem { color: #67e8f9; }
.ac-pixel-social { color: #f0abfc; }
.ac-pixel-dialog {
  position: relative;
  margin-top: 14px;
  border: 3px solid #c7d2fe;
  background: #0b0b22;
  padding: 11px 12px 16px;
  box-shadow: inset 0 0 0 2px #312e81;
}
.ac-pixel-speaker { font-size: 8px; letter-spacing: 0.1em; color: #fbbf24; }
.ac-pixel-dialog p {
  margin-top: 5px;
  font-family: 'PingFang SC', 'Microsoft YaHei', sans-serif;
  font-size: 12px;
  line-height: 1.8;
  color: #e0e7ff;
  overflow-wrap: anywhere;
}
.ac-pixel-next {
  position: absolute;
  right: 7px;
  bottom: 3px;
  color: #fbbf24;
  font-size: 9px;
  animation: ac-pixel-blink 1s steps(2) infinite;
}
.ac-pixel-actions {
  display: flex;
  justify-content: space-between;
  margin-top: 11px;
  font-size: 8px;
  letter-spacing: 0.08em;
  color: #a5b4fc;
}
.ac-pixel-hint {
  margin-top: 22px;
  text-align: center;
  font-size: 7px;
  letter-spacing: 0.24em;
  color: #4c4c9c;
  animation: ac-pixel-fade 1.4s 0.8s both;
}
@keyframes ac-pixel-enter {
  from { opacity: 0; transform: translateY(26px); clip-path: inset(0 0 100% 0); }
  to { opacity: 1; transform: none; clip-path: inset(0 0 0 0); }
}
@keyframes ac-pixel-fade { from { opacity: 0; } to { opacity: 1; } }
@keyframes ac-pixel-twinkle { 0%, 100% { opacity: 0.35; } 50% { opacity: 0.6; } }
@keyframes ac-pixel-blink { 0%, 49% { opacity: 1; } 50%, 100% { opacity: 0; } }
@media (prefers-reduced-motion: reduce) {
  .ac-pixel-root *,
  .ac-pixel-root *::before,
  .ac-pixel-root *::after {
    animation: none !important;
    transition: none !important;
  }
}
</style>
