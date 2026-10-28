<template>
  <div class="ac-vapor-root relative min-h-screen overflow-hidden px-4 py-8 text-[#2b1240]">
    <div class="ac-vapor-sky"></div>
    <div class="ac-vapor-sun"></div>
    <div class="ac-vapor-floor"></div>
    <div class="ac-vapor-scan"></div>

    <div class="relative z-10 mx-auto w-full max-w-md">
      <div class="ac-vapor-window">
        <div class="ac-vapor-titlebar">
          <span class="ac-vapor-app">★ ANIME_CARD.EXE</span>
          <span class="ac-vapor-btns"><i>_</i><i>□</i><i>×</i></span>
        </div>
        <div class="ac-vapor-menu">
          <span>File</span><span>Edit</span><span>View</span><span>Collection</span><span>Help</span>
        </div>

        <div class="ac-vapor-content">
          <div class="ac-vapor-kana">星海紀元</div>
          <p class="ac-vapor-kicker">DIGITAL CHARACTER ARCHIVE · 1998</p>
          <h1 class="ac-vapor-name">{{ p.name || '未命名角色' }}</h1>
          <p class="ac-vapor-title">{{ p.title }}</p>

          <div class="ac-vapor-specs">
            <div><span>作品</span><b>{{ p.series || '未知作品' }}</b></div>
            <div><span>属性</span><b class="ac-vapor-elem">{{ p.element }}</b></div>
            <div><span>社交</span><b class="ac-vapor-handle">@{{ p.social }}</b></div>
          </div>

          <div class="ac-vapor-dialog">
            <span class="ac-vapor-dialog-title">quote.txt</span>
            <p>{{ p.quote || '……信号加载中，暂无台词。' }}</p>
            <span class="ac-vapor-dialog-ok">[ OK ]</span>
          </div>

          <div class="ac-vapor-player">
            <span class="ac-vapor-play"><Play :size="11" fill="currentColor" /></span>
            <div class="ac-vapor-track"><i></i><i></i><i></i><i></i><i></i><i></i><i></i><i></i><i></i><i></i></div>
            <span>NFC-{{ nfcId || '0000' }}</span>
          </div>
        </div>

        <div class="ac-vapor-status">
          <span>MEMORY: 640K OK</span>
          <span>ONLINE</span>
        </div>
      </div>

      <div class="ac-vapor-sticker">
        <Disc3 :size="17" />
        <span>HOLO<br />DISC</span>
      </div>

      <p class="ac-vapor-caption">Ａ Ｅ Ｓ Ｔ Ｈ Ｅ Ｔ Ｉ Ｃ  ·  2 4 H  R A D I O</p>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { Disc3, Play } from 'lucide-vue-next'

const props = defineProps(['payload', 'nfcId'])
const p = computed(() => props.payload || {})
</script>

<style scoped>
.ac-vapor-root {
  background: #1a0733;
  font-family: 'MS Gothic', 'PingFang SC', 'Microsoft YaHei', monospace;
}
.ac-vapor-sky {
  position: absolute;
  inset: 0;
  background: linear-gradient(180deg, #12042b 0%, #4c1d95 28%, #db2777 62%, #fb923c 86%, #fde68a 100%);
}
.ac-vapor-sun {
  position: absolute;
  top: 18%;
  left: 50%;
  width: 190px;
  height: 190px;
  transform: translateX(-50%);
  border-radius: 50%;
  background: linear-gradient(180deg, #fef08a, #fb7185 58%, #db2777);
  box-shadow: 0 0 70px rgba(251, 146, 60, 0.7);
  overflow: hidden;
}
.ac-vapor-sun::after {
  content: '';
  position: absolute;
  inset: 52% 0 0;
  background: repeating-linear-gradient(180deg, #4c1d95 0 4px, transparent 4px 11px);
  opacity: 0.75;
}
.ac-vapor-floor {
  position: absolute;
  left: -40%;
  right: -40%;
  bottom: -10%;
  height: 46%;
  background-image:
    linear-gradient(rgba(244, 114, 182, 0.65) 1px, transparent 1px),
    linear-gradient(90deg, rgba(244, 114, 182, 0.65) 1px, transparent 1px);
  background-size: 44px 30px;
  transform: perspective(280px) rotateX(62deg);
  transform-origin: top center;
  mask-image: linear-gradient(180deg, transparent, #000 28%);
  opacity: 0.7;
}
.ac-vapor-scan {
  position: absolute;
  inset: 0;
  background: repeating-linear-gradient(180deg, transparent 0 2px, rgba(0, 0, 0, 0.08) 2px 3px);
  pointer-events: none;
}
.ac-vapor-window {
  border: 2px solid #2b1240;
  background: #ece9d8;
  box-shadow: 5px 5px 0 rgba(43, 18, 64, 0.45), 0 22px 50px -20px rgba(0, 0, 0, 0.5);
  animation: ac-vapor-pop 0.8s cubic-bezier(0.18, 1.2, 0.3, 1) both;
  transition: transform 0.4s ease, box-shadow 0.4s ease;
}
.ac-vapor-window:hover {
  transform: translate(-2px, -2px) rotate(-0.3deg);
  box-shadow: 9px 9px 0 rgba(43, 18, 64, 0.45), 0 28px 60px -20px rgba(0, 0, 0, 0.55);
}
.ac-vapor-titlebar {
  display: flex;
  align-items: center;
  justify-content: space-between;
  background: linear-gradient(90deg, #7c3aed, #c026d3 58%, #db2777);
  padding: 4px 5px;
  color: #fff;
  font-size: 8.5px;
  letter-spacing: 0.06em;
}
.ac-vapor-app { font-weight: 700; }
.ac-vapor-btns { display: flex; gap: 2px; }
.ac-vapor-btns i {
  display: flex;
  width: 14px;
  height: 12px;
  align-items: center;
  justify-content: center;
  border: 1px solid rgba(43, 18, 64, 0.5);
  background: #ece9d8;
  color: #2b1240;
  font-style: normal;
  font-size: 8px;
}
.ac-vapor-menu {
  display: flex;
  gap: 13px;
  border-bottom: 1px solid #aca899;
  background: #f4f1e4;
  padding: 4px 8px;
  font-size: 7.5px;
  color: #4a3f35;
}
.ac-vapor-content {
  position: relative;
  overflow: hidden;
  margin: 5px;
  border: 2px solid;
  border-color: #7a6a5a #fff #fff #7a6a5a;
  background:
    linear-gradient(155deg, rgba(255, 255, 255, 0.55), rgba(237, 233, 216, 0.2)),
    linear-gradient(150deg, #fdf2f8, #ede9fe 48%, #cffafe);
  padding: 16px 14px 12px;
}
.ac-vapor-kana {
  position: absolute;
  top: 12px;
  right: 9px;
  writing-mode: vertical-rl;
  font-size: 8px;
  letter-spacing: 0.28em;
  color: rgba(124, 58, 237, 0.4);
}
.ac-vapor-kicker {
  font-family: Consolas, monospace;
  font-size: 6.5px;
  letter-spacing: 0.2em;
  color: #a855f7;
}
.ac-vapor-name {
  margin-top: 8px;
  font-family: 'Arial Black', Impact, 'PingFang SC', sans-serif;
  font-size: 40px;
  font-weight: 900;
  line-height: 1.02;
  letter-spacing: -0.02em;
  background: linear-gradient(180deg, #ffffff 0%, #c4b5fd 38%, #6d28d9 52%, #f0abfc 70%, #ffffff 100%);
  -webkit-background-clip: text;
  background-clip: text;
  color: transparent;
  filter: drop-shadow(1px 2px 0 rgba(91, 33, 182, 0.35));
  overflow-wrap: anywhere;
}
.ac-vapor-title {
  display: inline-block;
  margin-top: 7px;
  border: 1px solid #db2777;
  background: rgba(255, 255, 255, 0.5);
  padding: 2px 7px;
  font-size: 10.5px;
  letter-spacing: 0.14em;
  color: #be185d;
}
.ac-vapor-specs {
  display: grid;
  gap: 0;
  margin-top: 16px;
  border: 1px solid #a78bfa;
  background: rgba(255, 255, 255, 0.46);
}
.ac-vapor-specs > div {
  display: flex;
  align-items: baseline;
  gap: 10px;
  border-bottom: 1px dotted #c4b5fd;
  padding: 7px 9px;
}
.ac-vapor-specs > div:last-child { border-bottom: none; }
.ac-vapor-specs span {
  flex: none;
  width: 32px;
  font-size: 8px;
  letter-spacing: 0.12em;
  color: #8b5cf6;
}
.ac-vapor-specs b {
  min-width: 0;
  flex: 1;
  text-align: right;
  font-family: 'PingFang SC', 'Microsoft YaHei', sans-serif;
  font-size: 11.5px;
  font-weight: 600;
  color: #3b1f56;
  overflow-wrap: anywhere;
}
.ac-vapor-elem { color: #0e7490 !important; }
.ac-vapor-handle { font-family: Consolas, monospace; font-size: 11px !important; color: #be185d !important; }
.ac-vapor-dialog {
  position: relative;
  margin-top: 13px;
  border: 1px solid #7c3aed;
  background: #fff;
  padding: 20px 11px 11px;
  box-shadow: 3px 3px 0 rgba(124, 58, 237, 0.18);
}
.ac-vapor-dialog-title {
  position: absolute;
  top: 0;
  left: 0;
  padding: 2px 8px;
  background: #7c3aed;
  color: #fff;
  font-family: Consolas, monospace;
  font-size: 6.5px;
  letter-spacing: 0.1em;
}
.ac-vapor-dialog p {
  font-family: 'PingFang SC', 'Microsoft YaHei', sans-serif;
  font-size: 12.5px;
  line-height: 1.8;
  color: #3b1f56;
  overflow-wrap: anywhere;
}
.ac-vapor-dialog-ok {
  display: block;
  margin-top: 8px;
  text-align: right;
  font-family: Consolas, monospace;
  font-size: 7.5px;
  color: #a855f7;
}
.ac-vapor-player {
  display: flex;
  align-items: center;
  gap: 8px;
  margin-top: 12px;
  border: 1px solid #c4b5fd;
  background: rgba(255, 255, 255, 0.5);
  padding: 5px 7px;
}
.ac-vapor-play {
  display: flex;
  width: 20px;
  height: 20px;
  align-items: center;
  justify-content: center;
  border: 1px solid #7c3aed;
  color: #7c3aed;
}
.ac-vapor-track { display: flex; flex: 1; align-items: flex-end; gap: 2px; height: 14px; }
.ac-vapor-track i { flex: 1; background: #c026d3; opacity: 0.55; }
.ac-vapor-track i:nth-child(odd) { height: 60%; }
.ac-vapor-track i:nth-child(even) { height: 100%; }
.ac-vapor-track i:nth-child(3n) { height: 35%; }
.ac-vapor-player > span:last-child {
  font-family: Consolas, monospace;
  font-size: 7px;
  letter-spacing: 0.08em;
  color: #7c3aed;
}
.ac-vapor-status {
  display: flex;
  justify-content: space-between;
  border-top: 1px solid #aca899;
  background: #ece9d8;
  padding: 4px 7px;
  font-family: Consolas, monospace;
  font-size: 6.5px;
  color: #6b5b4d;
}
.ac-vapor-status span:last-child { color: #15803d; }
.ac-vapor-sticker {
  position: absolute;
  right: 2px;
  bottom: -18px;
  display: flex;
  width: 78px;
  height: 78px;
  align-items: center;
  justify-content: center;
  gap: 4px;
  border-radius: 50%;
  background: radial-gradient(circle at 35% 30%, #fff, #a5f3fc 35%, #c4b5fd 68%, #f0abfc);
  box-shadow: 0 6px 16px rgba(43, 18, 64, 0.35);
  color: #6d28d9;
  transform: rotate(12deg);
  animation: ac-vapor-spin 14s linear infinite;
}
.ac-vapor-sticker span { font-size: 6.5px; font-weight: 700; line-height: 1.15; letter-spacing: 0.08em; }
.ac-vapor-caption {
  margin-top: 36px;
  text-align: center;
  font-family: Consolas, monospace;
  font-size: 7px;
  letter-spacing: 0.18em;
  color: rgba(255, 255, 255, 0.6);
  animation: ac-vapor-fade 1.4s 0.6s both;
}
@keyframes ac-vapor-pop {
  from { opacity: 0; transform: scale(0.88) translateY(24px); }
  to { opacity: 1; transform: none; }
}
@keyframes ac-vapor-fade { from { opacity: 0; } to { opacity: 1; } }
@keyframes ac-vapor-spin { to { transform: rotate(372deg); } }
@media (prefers-reduced-motion: reduce) {
  .ac-vapor-root *,
  .ac-vapor-root *::before,
  .ac-vapor-root *::after {
    animation: none !important;
    transition: none !important;
  }
}
</style>
