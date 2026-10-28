<template>
  <div class="ac-glass-root relative min-h-screen overflow-hidden px-5 py-9 text-[#3d3454]">
    <div class="ac-glass-blob ac-glass-b1"></div>
    <div class="ac-glass-blob ac-glass-b2"></div>
    <div class="ac-glass-blob ac-glass-b3"></div>
    <div class="ac-glass-bubbles">
      <i v-for="n in 8" :key="n"></i>
    </div>

    <div class="relative z-10 mx-auto flex min-h-[calc(100vh-4.5rem)] w-full max-w-md flex-col justify-center">
      <article class="ac-glass-card">
        <span class="ac-glass-shine"></span>
        <header class="ac-glass-head">
          <span class="ac-glass-chip"><Sparkles :size="10" />MOE ARCHIVE</span>
          <span class="ac-glass-rarity">★ SSR</span>
        </header>

        <div class="ac-glass-avatar">
          <span>{{ initial }}</span>
          <i class="ac-glass-status"></i>
        </div>

        <div class="ac-glass-title">
          <p class="ac-glass-kicker">{{ p.series || '未知作品' }}</p>
          <h1 class="ac-glass-name">{{ p.name || '未命名角色' }}</h1>
          <p class="ac-glass-sub">{{ p.title }}</p>
        </div>

        <div class="ac-glass-pills">
          <span>{{ p.element }}</span>
          <span>LIMITED</span>
          <span>1 / 1</span>
        </div>

        <blockquote class="ac-glass-quote">
          <p>{{ p.quote || '……像气泡一样，轻轻飘走了。' }}</p>
        </blockquote>

        <footer class="ac-glass-foot">
          <span class="ac-glass-label">SOCIAL</span>
          <b><AtSign :size="12" />{{ p.social }}</b>
          <span class="ac-glass-id">NFC {{ nfcId || '0000' }}</span>
        </footer>
      </article>

      <p class="ac-glass-caption">SOFT · LIGHT · COLLECTIBLE</p>
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
.ac-glass-root {
  background: linear-gradient(155deg, #e0e7ff 0%, #fae8ff 42%, #cffafe 100%);
  font-family: 'Avenir Next', 'PingFang SC', 'Microsoft YaHei', sans-serif;
}
.ac-glass-blob {
  position: absolute;
  border-radius: 50%;
  filter: blur(48px);
  opacity: 0.72;
  pointer-events: none;
}
.ac-glass-b1 {
  width: 280px; height: 280px; top: -70px; left: -80px;
  background: #c4b5fd;
  animation: ac-glass-float 11s ease-in-out infinite;
}
.ac-glass-b2 {
  width: 240px; height: 240px; right: -70px; top: 36%;
  background: #f9a8d4;
  animation: ac-glass-float 14s ease-in-out infinite reverse;
}
.ac-glass-b3 {
  width: 230px; height: 230px; bottom: -60px; left: 22%;
  background: #a5f3fc;
  animation: ac-glass-float 17s ease-in-out infinite;
}
.ac-glass-bubbles {
  position: absolute;
  inset: 0;
  pointer-events: none;
}
.ac-glass-bubbles i {
  position: absolute;
  border: 1px solid rgba(255, 255, 255, 0.75);
  border-radius: 50%;
  background: rgba(255, 255, 255, 0.18);
  box-shadow: inset 0 1px 6px rgba(255, 255, 255, 0.7);
  animation: ac-glass-rise 12s ease-in-out infinite;
}
.ac-glass-bubbles i:nth-child(1) { width: 22px; height: 22px; left: 9%; top: 18%; }
.ac-glass-bubbles i:nth-child(2) { width: 12px; height: 12px; left: 82%; top: 12%; animation-delay: 1.4s; }
.ac-glass-bubbles i:nth-child(3) { width: 30px; height: 30px; left: 78%; top: 58%; animation-delay: 2.6s; }
.ac-glass-bubbles i:nth-child(4) { width: 16px; height: 16px; left: 13%; top: 70%; animation-delay: 3.5s; }
.ac-glass-bubbles i:nth-child(5) { width: 9px; height: 9px; left: 46%; top: 8%; animation-delay: 5s; }
.ac-glass-bubbles i:nth-child(6) { width: 18px; height: 18px; left: 90%; top: 38%; animation-delay: 6s; }
.ac-glass-bubbles i:nth-child(7) { width: 7px; height: 7px; left: 23%; top: 40%; animation-delay: 7s; }
.ac-glass-bubbles i:nth-child(8) { width: 13px; height: 13px; left: 62%; top: 84%; animation-delay: 8.2s; }
.ac-glass-card {
  position: relative;
  overflow: hidden;
  border: 1px solid rgba(255, 255, 255, 0.85);
  border-radius: 30px;
  background: linear-gradient(150deg, rgba(255, 255, 255, 0.62), rgba(255, 255, 255, 0.24));
  padding: 22px 22px 20px;
  box-shadow:
    0 30px 60px -26px rgba(124, 58, 237, 0.36),
    inset 0 1px 0 rgba(255, 255, 255, 0.95),
    inset 0 -1px 0 rgba(255, 255, 255, 0.4);
  backdrop-filter: blur(24px) saturate(150%);
  animation: ac-glass-enter 1s cubic-bezier(0.2, 0.9, 0.2, 1) both;
  transition: transform 0.5s ease, box-shadow 0.5s ease;
}
.ac-glass-card:hover {
  transform: translateY(-6px) rotateX(3deg) rotateY(-3deg);
  box-shadow:
    0 40px 76px -26px rgba(124, 58, 237, 0.46),
    inset 0 1px 0 rgba(255, 255, 255, 1),
    inset 0 -1px 0 rgba(255, 255, 255, 0.4);
}
.ac-glass-shine {
  position: absolute;
  top: -60%;
  left: -30%;
  width: 90%;
  height: 160%;
  transform: rotate(24deg);
  background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.55), transparent);
  filter: blur(12px);
  animation: ac-glass-shine 8s ease-in-out infinite;
  pointer-events: none;
}
.ac-glass-head { position: relative; display: flex; align-items: center; justify-content: space-between; }
.ac-glass-chip {
  display: inline-flex;
  align-items: center;
  gap: 5px;
  border-radius: 999px;
  background: rgba(255, 255, 255, 0.72);
  border: 1px solid rgba(255, 255, 255, 0.9);
  padding: 4px 9px;
  font-family: Consolas, monospace;
  font-size: 6.5px;
  letter-spacing: 0.16em;
  color: #7c3aed;
}
.ac-glass-rarity {
  border-radius: 999px;
  background: linear-gradient(90deg, #c084fc, #f472b6);
  padding: 4px 10px;
  color: #fff;
  font-family: Consolas, monospace;
  font-size: 7px;
  letter-spacing: 0.14em;
  box-shadow: 0 3px 10px rgba(192, 132, 252, 0.4);
}
.ac-glass-avatar {
  position: relative;
  display: flex;
  width: 104px;
  height: 104px;
  margin: 30px auto 20px;
  align-items: center;
  justify-content: center;
  border: 1px solid rgba(255, 255, 255, 0.95);
  border-radius: 50%;
  background:
    radial-gradient(circle at 34% 28%, rgba(255, 255, 255, 0.95), rgba(221, 214, 254, 0.7) 45%, rgba(251, 207, 232, 0.55));
  box-shadow: 0 14px 30px -12px rgba(124, 58, 237, 0.45), inset 0 0 26px rgba(255, 255, 255, 0.8);
  backdrop-filter: blur(8px);
}
.ac-glass-avatar span {
  font-family: 'Songti SC', serif;
  font-size: 42px;
  color: #6d28d9;
  text-shadow: 0 2px 12px rgba(255, 255, 255, 0.9);
}
.ac-glass-status {
  position: absolute;
  right: 7px;
  bottom: 9px;
  width: 12px;
  height: 12px;
  border: 2px solid #fff;
  border-radius: 50%;
  background: #34d399;
  box-shadow: 0 0 8px rgba(52, 211, 153, 0.8);
}
.ac-glass-title { text-align: center; }
.ac-glass-kicker {
  font-family: Consolas, monospace;
  font-size: 7px;
  letter-spacing: 0.3em;
  color: rgba(109, 40, 217, 0.55);
}
.ac-glass-name {
  margin-top: 6px;
  font-size: 39px;
  font-weight: 700;
  line-height: 1.1;
  letter-spacing: 0.03em;
  color: #3b1f6e;
  overflow-wrap: anywhere;
}
.ac-glass-sub { margin-top: 6px; font-size: 12px; letter-spacing: 0.22em; color: #a855f7; }
.ac-glass-pills {
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
  gap: 7px;
  margin-top: 17px;
}
.ac-glass-pills span {
  border: 1px solid rgba(255, 255, 255, 0.9);
  border-radius: 999px;
  background: rgba(255, 255, 255, 0.5);
  padding: 4px 11px;
  font-size: 9px;
  letter-spacing: 0.1em;
  color: #6d28d9;
}
.ac-glass-pills span:first-child {
  background: linear-gradient(90deg, rgba(196, 181, 253, 0.55), rgba(249, 168, 212, 0.55));
  color: #5b21b6;
}
.ac-glass-quote {
  margin-top: 20px;
  border: 1px solid rgba(255, 255, 255, 0.85);
  border-radius: 18px;
  background: rgba(255, 255, 255, 0.36);
  padding: 15px 16px;
  box-shadow: inset 0 1px 8px rgba(255, 255, 255, 0.55);
}
.ac-glass-quote p {
  text-align: center;
  font-family: 'Songti SC', 'Noto Serif SC', serif;
  font-size: 14.5px;
  line-height: 1.9;
  color: #4c2a72;
  overflow-wrap: anywhere;
}
.ac-glass-foot {
  display: flex;
  align-items: center;
  gap: 8px;
  margin-top: 18px;
  border-top: 1px solid rgba(255, 255, 255, 0.7);
  padding-top: 13px;
}
.ac-glass-label {
  font-family: Consolas, monospace;
  font-size: 6.5px;
  letter-spacing: 0.2em;
  color: rgba(109, 40, 217, 0.45);
}
.ac-glass-foot b {
  display: inline-flex;
  align-items: center;
  gap: 4px;
  font-family: Consolas, monospace;
  font-size: 11px;
  font-weight: 500;
  color: #7c3aed;
  overflow-wrap: anywhere;
}
.ac-glass-id {
  margin-left: auto;
  flex: none;
  font-family: Consolas, monospace;
  font-size: 6.5px;
  letter-spacing: 0.1em;
  color: rgba(61, 52, 84, 0.4);
}
.ac-glass-caption {
  margin-top: 24px;
  text-align: center;
  font-family: Consolas, monospace;
  font-size: 7px;
  letter-spacing: 0.36em;
  color: rgba(109, 40, 217, 0.32);
  animation: ac-glass-fade 1.4s 0.6s both;
}
@keyframes ac-glass-enter {
  from { opacity: 0; transform: translateY(34px) scale(0.95); filter: blur(8px); }
  to { opacity: 1; transform: none; filter: none; }
}
@keyframes ac-glass-fade { from { opacity: 0; } to { opacity: 1; } }
@keyframes ac-glass-float {
  0%, 100% { transform: translate3d(0, 0, 0) scale(1); }
  50% { transform: translate3d(18px, -22px, 0) scale(1.08); }
}
@keyframes ac-glass-rise {
  0%, 100% { transform: translateY(0); opacity: 0.5; }
  50% { transform: translateY(-22px); opacity: 0.95; }
}
@keyframes ac-glass-shine {
  0%, 100% { transform: translateX(-30%) rotate(24deg); opacity: 0; }
  45% { opacity: 0.8; }
  70% { transform: translateX(180%) rotate(24deg); opacity: 0; }
}
@media (prefers-reduced-motion: reduce) {
  .ac-glass-root *,
  .ac-glass-root *::before,
  .ac-glass-root *::after {
    animation: none !important;
    transition: none !important;
  }
}
</style>
