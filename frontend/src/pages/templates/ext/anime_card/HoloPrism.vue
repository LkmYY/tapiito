<template>
  <div class="ac-holo-root relative min-h-screen overflow-hidden px-5 py-10 text-white">
    <div class="ac-holo-aurora ac-holo-aurora-a"></div>
    <div class="ac-holo-aurora ac-holo-aurora-b"></div>
    <div class="ac-holo-aurora ac-holo-aurora-c"></div>

    <div class="relative z-10 mx-auto w-full max-w-md">
      <div class="ac-holo-stage">
        <article class="ac-holo-card">
          <div class="ac-holo-foil"></div>
          <div class="ac-holo-sheen"></div>
          <div class="ac-holo-grain"></div>

          <span class="ac-holo-spark ac-holo-spark-a"></span>
          <span class="ac-holo-spark ac-holo-spark-b"></span>
          <span class="ac-holo-spark ac-holo-spark-c"></span>

          <header class="ac-holo-head">
            <span class="ac-holo-kicker">HOLOGRAPHIC ARCHIVE</span>
            <span class="ac-holo-rarity">SSR</span>
          </header>

          <div class="ac-holo-emblem">
            <span class="ac-holo-emblem-inner">{{ initial }}</span>
            <span class="ac-holo-orbit"></span>
          </div>

          <p class="ac-holo-index">NO. {{ nfcId || '0001' }} / PRISM SERIES</p>
          <h1 class="ac-holo-name">{{ p.name || '未命名角色' }}</h1>
          <p class="ac-holo-title">{{ p.title }}</p>

          <div class="ac-holo-tags">
            <span><Sparkles :size="12" />{{ p.series || '未知作品' }}</span>
            <span class="ac-holo-element">{{ p.element }}</span>
          </div>

          <blockquote class="ac-holo-quote">
            <span class="ac-holo-mark">“</span>
            <p>{{ p.quote || '尚未收录台词。' }}</p>
          </blockquote>

          <footer class="ac-holo-foot">
            <span class="ac-holo-social"><AtSign :size="13" />{{ p.social }}</span>
            <span class="ac-holo-chip">HOLO / 1 OF 1</span>
          </footer>
        </article>
      </div>

      <p class="ac-holo-caption">TAP · TILT · COLLECT</p>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { AtSign, Sparkles } from 'lucide-vue-next'

const props = defineProps(['payload', 'nfcId'])
const p = computed(() => props.payload || {})
const initial = computed(() => (p.value.name || '？').toString().trim().slice(0, 1) || '？')
</script>

<style scoped>
.ac-holo-root {
  background: #07050f;
  font-family: 'Avenir Next', 'PingFang SC', 'Microsoft YaHei', sans-serif;
}
.ac-holo-aurora {
  position: absolute;
  border-radius: 50%;
  filter: blur(80px);
  opacity: 0.42;
  pointer-events: none;
}
.ac-holo-aurora-a {
  width: 300px; height: 300px; top: -110px; left: -100px;
  background: #7c3aed;
  animation: ac-holo-drift 12s ease-in-out infinite;
}
.ac-holo-aurora-b {
  width: 260px; height: 260px; right: -120px; top: 34%;
  background: #db2777;
  animation: ac-holo-drift 15s ease-in-out infinite reverse;
}
.ac-holo-aurora-c {
  width: 240px; height: 240px; bottom: -100px; left: 18%;
  background: #0891b2;
  animation: ac-holo-drift 18s ease-in-out infinite;
}
.ac-holo-stage {
  perspective: 1200px;
  animation: ac-holo-enter 1s cubic-bezier(0.2, 0.8, 0.2, 1) both;
}
.ac-holo-card {
  position: relative;
  overflow: hidden;
  border-radius: 24px;
  border: 1px solid rgba(255, 255, 255, 0.24);
  padding: 22px 22px 20px;
  background:
    radial-gradient(120% 80% at 12% 0%, rgba(124, 58, 237, 0.34), transparent 52%),
    radial-gradient(110% 70% at 100% 100%, rgba(8, 145, 178, 0.3), transparent 55%),
    linear-gradient(155deg, #17102b 0%, #0a0715 52%, #140d26 100%);
  box-shadow:
    0 0 0 1px rgba(0, 0, 0, 0.6),
    inset 0 1px 0 rgba(255, 255, 255, 0.2),
    0 34px 90px -35px rgba(168, 85, 247, 0.75);
  transform-style: preserve-3d;
  transition: transform 0.55s cubic-bezier(0.2, 0.8, 0.2, 1), box-shadow 0.55s ease;
}
.ac-holo-card:hover {
  transform: rotateX(5deg) rotateY(-7deg) translateY(-5px);
  box-shadow:
    0 0 0 1px rgba(255, 255, 255, 0.08),
    inset 0 1px 0 rgba(255, 255, 255, 0.28),
    0 44px 110px -30px rgba(217, 70, 239, 0.75);
}
.ac-holo-card::before {
  content: '';
  position: absolute;
  inset: 8px;
  border: 1px solid rgba(255, 255, 255, 0.1);
  border-radius: 17px;
  pointer-events: none;
}
.ac-holo-foil {
  position: absolute;
  inset: -45%;
  background: conic-gradient(from 20deg, #ff5ec4, #ffd86b, #5ef2c4, #5ea8ff, #b05eff, #ff5ec4);
  opacity: 0.15;
  filter: blur(34px);
  mix-blend-mode: screen;
  animation: ac-holo-spin 22s linear infinite;
  pointer-events: none;
}
.ac-holo-sheen {
  position: absolute;
  inset: 0;
  background: linear-gradient(112deg, transparent 30%, rgba(255, 255, 255, 0.14) 44%, rgba(255, 255, 255, 0.32) 49%, transparent 60%);
  background-size: 260% 100%;
  background-position: 80% 0;
  mix-blend-mode: screen;
  transition: background-position 0.9s ease;
  pointer-events: none;
}
.ac-holo-card:hover .ac-holo-sheen { background-position: 0% 0; }
.ac-holo-grain {
  position: absolute;
  inset: 0;
  opacity: 0.34;
  background-image: radial-gradient(rgba(255, 255, 255, 0.16) 0.5px, transparent 0.6px);
  background-size: 4px 4px;
  mix-blend-mode: overlay;
  pointer-events: none;
}
.ac-holo-head {
  position: relative;
  display: flex;
  align-items: center;
  justify-content: space-between;
}
.ac-holo-kicker {
  font-family: Consolas, monospace;
  font-size: 8px;
  letter-spacing: 0.3em;
  color: rgba(255, 255, 255, 0.5);
}
.ac-holo-rarity {
  padding: 3px 9px;
  border: 1px solid rgba(255, 216, 107, 0.65);
  border-radius: 999px;
  color: #ffe29a;
  font-family: Consolas, monospace;
  font-size: 9px;
  letter-spacing: 0.24em;
  box-shadow: 0 0 16px rgba(255, 216, 107, 0.25);
}
.ac-holo-emblem {
  position: relative;
  display: flex;
  align-items: center;
  justify-content: center;
  width: 94px;
  height: 94px;
  margin: 32px auto 22px;
  border-radius: 50%;
  background: conic-gradient(from 0deg, #67e8f9, #e879f9, #fde68a, #67e8f9);
  padding: 2px;
  box-shadow: 0 0 34px rgba(232, 121, 249, 0.35);
}
.ac-holo-emblem-inner {
  display: flex;
  width: 100%;
  height: 100%;
  align-items: center;
  justify-content: center;
  border-radius: 50%;
  background: radial-gradient(circle at 40% 30%, #2c1a4d, #0b0716 72%);
  color: #fdf4ff;
  font-family: 'Songti SC', 'Noto Serif SC', serif;
  font-size: 38px;
  letter-spacing: 0.06em;
  text-shadow: 0 0 22px rgba(232, 121, 249, 0.8);
}
.ac-holo-orbit {
  position: absolute;
  inset: -9px;
  border: 1px dashed rgba(255, 255, 255, 0.22);
  border-radius: 50%;
  animation: ac-holo-spin 26s linear infinite;
}
.ac-holo-orbit::after {
  content: '';
  position: absolute;
  top: -3px;
  left: 50%;
  width: 5px;
  height: 5px;
  border-radius: 50%;
  background: #fde68a;
  box-shadow: 0 0 10px #fde68a;
}
.ac-holo-index {
  position: relative;
  text-align: center;
  font-family: Consolas, monospace;
  font-size: 8px;
  letter-spacing: 0.24em;
  color: rgba(255, 255, 255, 0.36);
}
.ac-holo-name {
  position: relative;
  margin-top: 9px;
  text-align: center;
  font-size: 40px;
  font-weight: 700;
  line-height: 1.1;
  letter-spacing: 0.03em;
  background: linear-gradient(180deg, #ffffff 18%, #f5d0fe 55%, #a5f3fc 100%);
  -webkit-background-clip: text;
  background-clip: text;
  color: transparent;
  filter: drop-shadow(0 2px 18px rgba(232, 121, 249, 0.35));
  overflow-wrap: anywhere;
}
.ac-holo-title {
  position: relative;
  margin-top: 8px;
  text-align: center;
  font-size: 12px;
  letter-spacing: 0.32em;
  color: #f0abfc;
}
.ac-holo-tags {
  position: relative;
  display: flex;
  justify-content: center;
  flex-wrap: wrap;
  gap: 8px;
  margin-top: 20px;
}
.ac-holo-tags span {
  display: inline-flex;
  align-items: center;
  gap: 5px;
  border: 1px solid rgba(255, 255, 255, 0.14);
  border-radius: 999px;
  background: rgba(255, 255, 255, 0.05);
  padding: 5px 11px;
  font-size: 10.5px;
  letter-spacing: 0.08em;
  color: rgba(255, 255, 255, 0.72);
}
.ac-holo-tags .ac-holo-element {
  border-color: rgba(103, 232, 249, 0.4);
  color: #a5f3fc;
  background: rgba(8, 145, 178, 0.14);
}
.ac-holo-quote {
  position: relative;
  margin-top: 22px;
  border-top: 1px solid rgba(255, 255, 255, 0.1);
  border-bottom: 1px solid rgba(255, 255, 255, 0.1);
  padding: 18px 12px 16px 30px;
  background: linear-gradient(90deg, rgba(232, 121, 249, 0.09), transparent);
}
.ac-holo-mark {
  position: absolute;
  top: 8px;
  left: 8px;
  font-family: Georgia, serif;
  font-size: 34px;
  line-height: 1;
  color: rgba(232, 121, 249, 0.5);
}
.ac-holo-quote p {
  font-family: 'Songti SC', 'Noto Serif SC', serif;
  font-size: 14px;
  line-height: 1.85;
  color: rgba(255, 255, 255, 0.84);
  overflow-wrap: anywhere;
}
.ac-holo-foot {
  position: relative;
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 10px;
  margin-top: 18px;
}
.ac-holo-social {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  font-family: Consolas, monospace;
  font-size: 11px;
  letter-spacing: 0.06em;
  color: #f0abfc;
  overflow-wrap: anywhere;
}
.ac-holo-chip {
  flex: none;
  font-family: Consolas, monospace;
  font-size: 7.5px;
  letter-spacing: 0.18em;
  color: rgba(255, 255, 255, 0.32);
}
.ac-holo-spark {
  position: absolute;
  width: 12px;
  height: 12px;
  background: #fff;
  clip-path: polygon(50% 0, 61% 37%, 100% 50%, 61% 63%, 50% 100%, 39% 63%, 0 50%, 39% 37%);
  opacity: 0.9;
  animation: ac-holo-twinkle 3.2s ease-in-out infinite;
}
.ac-holo-spark-a { top: 15%; right: 12%; width: 9px; height: 9px; }
.ac-holo-spark-b { top: 44%; left: 7%; width: 7px; height: 7px; animation-delay: 1s; }
.ac-holo-spark-c { bottom: 23%; right: 9%; width: 11px; height: 11px; animation-delay: 2s; }
.ac-holo-caption {
  margin-top: 26px;
  text-align: center;
  font-family: Consolas, monospace;
  font-size: 8px;
  letter-spacing: 0.44em;
  color: rgba(255, 255, 255, 0.28);
  animation: ac-holo-fade 1.4s 0.6s both;
}
@keyframes ac-holo-enter {
  from { opacity: 0; transform: translateY(34px) rotateX(12deg) scale(0.94); }
  to { opacity: 1; transform: none; }
}
@keyframes ac-holo-fade {
  from { opacity: 0; } to { opacity: 1; }
}
@keyframes ac-holo-drift {
  0%, 100% { transform: translate3d(0, 0, 0) scale(1); }
  50% { transform: translate3d(18px, -24px, 0) scale(1.1); }
}
@keyframes ac-holo-spin {
  from { transform: rotate(0deg); } to { transform: rotate(360deg); }
}
@keyframes ac-holo-twinkle {
  0%, 100% { opacity: 0.25; transform: scale(0.7) rotate(0deg); }
  50% { opacity: 1; transform: scale(1) rotate(45deg); }
}
@media (prefers-reduced-motion: reduce) {
  .ac-holo-root *,
  .ac-holo-root *::before,
  .ac-holo-root *::after {
    animation: none !important;
    transition: none !important;
  }
}
</style>
