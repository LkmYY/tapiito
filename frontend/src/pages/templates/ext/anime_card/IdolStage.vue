<template>
  <div class="ac-idol-root relative min-h-screen overflow-hidden px-5 py-8 text-white">
    <div class="ac-idol-beam ac-idol-beam-a"></div>
    <div class="ac-idol-beam ac-idol-beam-b"></div>
    <div class="ac-idol-beam ac-idol-beam-c"></div>
    <div class="ac-idol-confetti"></div>
    <div class="ac-idol-floor"></div>

    <div class="relative z-10 mx-auto w-full max-w-md">
      <header class="ac-idol-top">
        <span class="ac-idol-live"><i></i>LIVE STAGE</span>
        <span>FANCAM · {{ nfcId || '0001' }}</span>
      </header>

      <div class="ac-idol-card">
        <div class="ac-idol-glare"></div>
        <span class="ac-idol-corner ac-idol-tl">✦</span>
        <span class="ac-idol-corner ac-idol-tr">✦</span>
        <span class="ac-idol-corner ac-idol-bl">✦</span>
        <span class="ac-idol-corner ac-idol-br">✦</span>

        <div class="ac-idol-card-head">
          <span>OFFICIAL PHOTOCARD</span>
          <span class="ac-idol-no">NO.01</span>
        </div>

        <div class="ac-idol-portrait">
          <span class="ac-idol-halo"></span>
          <span class="ac-idol-face">{{ initial }}</span>
          <span class="ac-idol-debut">DEBUT</span>
          <span class="ac-idol-sign">{{ p.social || 'mio_star' }}</span>
        </div>

        <div class="ac-idol-plate">
          <p class="ac-idol-group">{{ p.series || '未知作品' }}</p>
          <h1 class="ac-idol-name">{{ p.name || '未命名角色' }}</h1>
          <p class="ac-idol-title">{{ p.title }}</p>
          <span class="ac-idol-elem"><Sparkles :size="11" />{{ p.element }}</span>
        </div>
      </div>

      <div class="ac-idol-cheer">
        <Heart :size="13" fill="currentColor" />
        <p>{{ p.quote || '……请再为我应援一次吧。' }}</p>
        <Heart :size="13" fill="currentColor" />
      </div>

      <div class="ac-idol-banner">
        <Music :size="13" />
        <span>FANCAFE</span>
        <b>@{{ p.social }}</b>
        <span class="ac-idol-hash">#{{ p.name || '星海纪元' }}</span>
      </div>

      <p class="ac-idol-foot">OFFICIAL GOODS · NOT FOR SALE · {{ nfcId || '0000' }}</p>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { Heart, Music, Sparkles } from 'lucide-vue-next'

const props = defineProps(['payload', 'nfcId'])
const p = computed(() => props.payload || {})
const initial = computed(() => (p.value.name || '？').toString().trim().slice(0, 1) || '？')
</script>

<style scoped>
.ac-idol-root {
  background: #0b0614;
  font-family: 'PingFang SC', 'Microsoft YaHei', sans-serif;
}
.ac-idol-beam {
  position: absolute;
  top: -24%;
  width: 170px;
  height: 130%;
  background: linear-gradient(180deg, rgba(244, 114, 182, 0.3), transparent 72%);
  filter: blur(12px);
  transform-origin: top center;
  pointer-events: none;
}
.ac-idol-beam-a { left: 6%; transform: rotate(17deg); animation: ac-idol-sway 7s ease-in-out infinite; }
.ac-idol-beam-b { left: 50%; width: 130px; background: linear-gradient(180deg, rgba(196, 181, 253, 0.28), transparent 66%); transform: rotate(-6deg); animation: ac-idol-sway 9s ease-in-out infinite reverse; }
.ac-idol-beam-c { right: 4%; transform: rotate(-19deg); animation: ac-idol-sway 8s ease-in-out infinite; }
.ac-idol-confetti {
  position: absolute;
  inset: 0;
  background-image:
    radial-gradient(circle, #f9a8d4 1.3px, transparent 1.5px),
    radial-gradient(circle, #fde68a 1px, transparent 1.2px),
    radial-gradient(circle, #a5f3fc 0.9px, transparent 1.1px);
  background-size: 73px 91px, 103px 67px, 131px 113px;
  background-position: 0 0, 31px 47px, 71px 19px;
  opacity: 0.55;
  animation: ac-idol-confetti 18s linear infinite;
}
.ac-idol-floor {
  position: absolute;
  left: 0;
  right: 0;
  bottom: 0;
  height: 150px;
  background:
    linear-gradient(180deg, transparent, rgba(190, 24, 93, 0.2)),
    repeating-linear-gradient(90deg, rgba(255, 255, 255, 0.05) 0 1px, transparent 1px 49px);
  transform: perspective(200px) rotateX(48deg);
  transform-origin: bottom center;
}
.ac-idol-top {
  display: flex;
  justify-content: space-between;
  align-items: center;
  font-family: Consolas, monospace;
  font-size: 8px;
  letter-spacing: 0.2em;
  color: rgba(255, 255, 255, 0.42);
  animation: ac-idol-fade 0.8s both;
}
.ac-idol-live { display: inline-flex; align-items: center; gap: 6px; color: #f9a8d4; }
.ac-idol-live i {
  width: 5px; height: 5px; border-radius: 50%; background: #f43f5e;
  box-shadow: 0 0 9px #f43f5e;
  animation: ac-idol-pulse 1.6s ease-in-out infinite;
}
.ac-idol-card {
  position: relative;
  margin-top: 26px;
  overflow: hidden;
  border-radius: 4px;
  border: 1px solid rgba(255, 255, 255, 0.32);
  background: linear-gradient(160deg, #2a1038, #170a22 58%, #2b0d2c);
  padding: 12px 12px 0;
  box-shadow: 0 28px 70px -28px rgba(244, 114, 182, 0.55);
  animation: ac-idol-enter 0.95s cubic-bezier(0.18, 1, 0.3, 1) both;
  transition: transform 0.5s ease, box-shadow 0.5s ease;
}
.ac-idol-card:hover {
  transform: translateY(-5px) rotateX(3deg) rotateY(-4deg);
  box-shadow: 0 36px 88px -26px rgba(244, 114, 182, 0.7);
}
.ac-idol-glare {
  position: absolute;
  inset: -30% -40%;
  background: linear-gradient(112deg, transparent 36%, rgba(255, 255, 255, 0.16) 46%, rgba(255, 255, 255, 0.05) 51%, transparent 58%);
  background-size: 220% 100%;
  background-position: 88% 0;
  transition: background-position 0.9s ease;
  pointer-events: none;
}
.ac-idol-card:hover .ac-idol-glare { background-position: 8% 0; }
.ac-idol-corner { position: absolute; z-index: 2; color: rgba(253, 230, 138, 0.6); font-size: 9px; }
.ac-idol-tl { top: 7px; left: 8px; }
.ac-idol-tr { top: 7px; right: 8px; }
.ac-idol-bl { bottom: 7px; left: 8px; }
.ac-idol-br { bottom: 7px; right: 8px; }
.ac-idol-card-head {
  position: relative;
  z-index: 2;
  display: flex;
  justify-content: space-between;
  padding: 3px 6px 9px;
  font-family: Consolas, monospace;
  font-size: 6.5px;
  letter-spacing: 0.18em;
  color: rgba(255, 255, 255, 0.45);
}
.ac-idol-no { color: #f9a8d4; }
.ac-idol-portrait {
  position: relative;
  display: flex;
  height: 214px;
  align-items: center;
  justify-content: center;
  overflow: hidden;
  border: 1px solid rgba(255, 255, 255, 0.16);
  background:
    radial-gradient(circle at 50% 44%, rgba(244, 114, 182, 0.55), transparent 46%),
    linear-gradient(160deg, #4c1d95, #831843 62%, #1e1b4b);
}
.ac-idol-portrait::before {
  content: '';
  position: absolute;
  inset: 0;
  background: repeating-conic-gradient(from 0deg at 50% 62%, rgba(255, 255, 255, 0.13) 0deg 0.7deg, transparent 0.7deg 9deg);
  mask-image: radial-gradient(circle at 50% 60%, transparent 12%, #000 60%);
}
.ac-idol-halo {
  position: absolute;
  top: 24px;
  width: 150px;
  height: 150px;
  border: 1px solid rgba(253, 230, 138, 0.42);
  border-radius: 50%;
  box-shadow: 0 0 40px rgba(253, 230, 138, 0.18), inset 0 0 34px rgba(253, 230, 138, 0.12);
}
.ac-idol-halo::before,
.ac-idol-halo::after {
  content: '';
  position: absolute;
  inset: 12px;
  border: 1px dashed rgba(255, 255, 255, 0.2);
  border-radius: 50%;
}
.ac-idol-halo::after { inset: -12px; border-style: solid; opacity: 0.5; }
.ac-idol-face {
  position: relative;
  z-index: 1;
  display: flex;
  width: 112px;
  height: 112px;
  align-items: center;
  justify-content: center;
  border-radius: 50%;
  background: radial-gradient(circle at 38% 28%, rgba(255, 255, 255, 0.32), rgba(131, 24, 67, 0.65) 62%);
  color: #fff7fb;
  font-family: 'Songti SC', 'Noto Serif SC', serif;
  font-size: 50px;
  text-shadow: 0 0 26px rgba(255, 255, 255, 0.65);
}
.ac-idol-debut {
  position: absolute;
  top: 12px;
  right: 10px;
  border: 1px solid rgba(253, 230, 138, 0.6);
  padding: 2px 7px;
  font-family: Consolas, monospace;
  font-size: 6.5px;
  letter-spacing: 0.2em;
  color: #fde68a;
}
.ac-idol-sign {
  position: absolute;
  right: 15px;
  bottom: 13px;
  transform: rotate(-7deg);
  font-family: 'Segoe Script', 'Brush Script MT', cursive;
  font-size: 20px;
  color: #fdf2f8;
  text-shadow: 0 1px 8px rgba(0, 0, 0, 0.4);
  opacity: 0.9;
}
.ac-idol-plate {
  position: relative;
  z-index: 2;
  padding: 18px 8px 20px;
  text-align: center;
}
.ac-idol-group {
  font-family: Consolas, monospace;
  font-size: 7px;
  letter-spacing: 0.3em;
  color: #f9a8d4;
}
.ac-idol-name {
  margin-top: 6px;
  font-size: 39px;
  font-weight: 800;
  line-height: 1.06;
  letter-spacing: 0.05em;
  background: linear-gradient(180deg, #ffffff 20%, #fbcfe8 60%, #c4b5fd 100%);
  -webkit-background-clip: text;
  background-clip: text;
  color: transparent;
  overflow-wrap: anywhere;
}
.ac-idol-title {
  margin-top: 6px;
  font-size: 11.5px;
  letter-spacing: 0.24em;
  color: rgba(255, 255, 255, 0.7);
}
.ac-idol-elem {
  display: inline-flex;
  align-items: center;
  gap: 4px;
  margin-top: 11px;
  border: 1px solid rgba(249, 168, 212, 0.5);
  border-radius: 999px;
  background: rgba(244, 114, 182, 0.12);
  padding: 3px 10px;
  font-size: 9px;
  letter-spacing: 0.1em;
  color: #fbcfe8;
}
.ac-idol-cheer {
  display: flex;
  align-items: center;
  gap: 10px;
  margin-top: 16px;
  border-top: 1px solid rgba(244, 114, 182, 0.22);
  border-bottom: 1px solid rgba(244, 114, 182, 0.22);
  padding: 13px 5px;
  color: #f472b6;
}
.ac-idol-cheer p {
  flex: 1;
  text-align: center;
  font-family: 'Songti SC', 'Noto Serif SC', serif;
  font-size: 14px;
  line-height: 1.8;
  color: rgba(255, 255, 255, 0.86);
  overflow-wrap: anywhere;
}
.ac-idol-banner {
  display: flex;
  align-items: center;
  gap: 7px;
  margin-top: 15px;
  border: 1px solid rgba(255, 255, 255, 0.14);
  background: rgba(255, 255, 255, 0.045);
  padding: 8px 10px;
  color: #c4b5fd;
  font-family: Consolas, monospace;
  font-size: 7.5px;
  letter-spacing: 0.12em;
}
.ac-idol-banner b {
  color: #f9a8d4;
  font-weight: 400;
  overflow-wrap: anywhere;
}
.ac-idol-hash {
  margin-left: auto;
  color: rgba(255, 255, 255, 0.3);
  overflow-wrap: anywhere;
}
.ac-idol-foot {
  margin-top: 17px;
  text-align: center;
  font-family: Consolas, monospace;
  font-size: 6.5px;
  letter-spacing: 0.2em;
  color: rgba(255, 255, 255, 0.24);
  animation: ac-idol-fade 1.2s 0.6s both;
}
@keyframes ac-idol-enter {
  from { opacity: 0; transform: translateY(34px) scale(0.94) rotate(-1.5deg); }
  to { opacity: 1; transform: none; }
}
@keyframes ac-idol-fade { from { opacity: 0; } to { opacity: 1; } }
@keyframes ac-idol-pulse { 0%, 100% { opacity: 1; } 50% { opacity: 0.25; } }
@keyframes ac-idol-sway { 0%, 100% { opacity: 0.55; transform: rotate(17deg) translateX(-8px); } 50% { opacity: 0.9; transform: rotate(11deg) translateX(10px); } }
@keyframes ac-idol-confetti { to { background-position: 73px 91px, -72px 114px, 202px 132px; } }
@media (prefers-reduced-motion: reduce) {
  .ac-idol-root *,
  .ac-idol-root *::before,
  .ac-idol-root *::after {
    animation: none !important;
    transition: none !important;
  }
}
</style>
