<template>
  <div class="ac-sumi-root relative min-h-screen overflow-hidden px-5 py-8 text-[#211d18]">
    <div class="ac-sumi-wash ac-sumi-wash-a"></div>
    <div class="ac-sumi-wash ac-sumi-wash-b"></div>
    <div class="ac-sumi-fibers"></div>

    <div class="relative z-10 mx-auto w-full max-w-md">
      <div class="ac-sumi-scroll">
        <div class="ac-sumi-rod ac-sumi-rod-top">
          <span class="ac-sumi-cord ac-sumi-cord-l"></span>
          <span class="ac-sumi-cord ac-sumi-cord-r"></span>
        </div>

        <div class="ac-sumi-paper">
          <div class="ac-sumi-enso">
            <span class="ac-sumi-initial">{{ initial }}</span>
          </div>

          <p class="ac-sumi-kicker">CHARACTER SCROLL · 星海紀元</p>
          <h1 class="ac-sumi-name">{{ p.name || '未命名角色' }}</h1>
          <span class="ac-sumi-seal-name">印</span>

          <div class="ac-sumi-vertical">
            <p class="ac-sumi-title">{{ p.title }}</p>
            <span class="ac-sumi-dot">・</span>
            <p class="ac-sumi-series">{{ p.series || '未知作品' }}</p>
          </div>

          <div class="ac-sumi-element">
            <span>{{ p.element }}</span>
          </div>

          <blockquote class="ac-sumi-quote">
            <p>{{ p.quote || '……风过无声，星落无痕。' }}</p>
            <svg viewBox="0 0 200 10" aria-hidden="true">
              <path d="M2 7 C40 2, 70 9, 105 5 S170 3, 198 6" fill="none" stroke="#2b2620" stroke-width="2.2" stroke-linecap="round" opacity="0.55" />
            </svg>
          </blockquote>

          <footer class="ac-sumi-foot">
            <span class="ac-sumi-thread"></span>
            <div>
              <small>落款 · SNS</small>
              <b>@{{ p.social }}</b>
            </div>
            <span class="ac-sumi-hanko">澪</span>
          </footer>

          <p class="ac-sumi-id">NFC {{ nfcId || '一二' }} · 一紙一世界</p>
        </div>

        <div class="ac-sumi-rod ac-sumi-rod-bottom"></div>
      </div>

      <p class="ac-sumi-caption">墨 の 余 白  ·  M I O</p>
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
.ac-sumi-root {
  background: #ded3bd;
  font-family: 'Songti SC', 'Noto Serif SC', 'STSong', serif;
}
.ac-sumi-fibers {
  position: absolute;
  inset: 0;
  background-image:
    repeating-linear-gradient(92deg, rgba(120, 100, 70, 0.05) 0 1px, transparent 1px 7px),
    repeating-linear-gradient(2deg, rgba(120, 100, 70, 0.04) 0 1px, transparent 1px 11px);
  opacity: 0.8;
}
.ac-sumi-wash {
  position: absolute;
  border-radius: 50%;
  background: radial-gradient(circle, rgba(33, 29, 24, 0.3), rgba(33, 29, 24, 0.05) 58%, transparent 72%);
  filter: blur(18px);
}
.ac-sumi-wash-a { width: 300px; height: 220px; top: -70px; left: -100px; transform: rotate(-18deg); }
.ac-sumi-wash-b { width: 260px; height: 200px; bottom: -60px; right: -90px; transform: rotate(14deg); opacity: 0.75; }
.ac-sumi-scroll {
  animation: ac-sumi-unroll 1.2s cubic-bezier(0.22, 1, 0.36, 1) both;
  transform-origin: top center;
}
.ac-sumi-rod {
  position: relative;
  z-index: 2;
  height: 15px;
  border-radius: 3px;
  background: linear-gradient(180deg, #6b5140, #3d2b22 48%, #241713);
  box-shadow: 0 3px 7px rgba(40, 25, 18, 0.35);
}
.ac-sumi-rod-top::after,
.ac-sumi-rod-bottom::after {
  content: '';
  position: absolute;
  top: 3px;
  width: 100%;
  height: 1px;
  background: rgba(255, 236, 210, 0.22);
}
.ac-sumi-rod-top { margin: 0 -7px; width: calc(100% + 14px); }
.ac-sumi-rod-bottom { margin: 0 -7px; width: calc(100% + 14px); }
.ac-sumi-cord {
  position: absolute;
  top: -34px;
  width: 1px;
  height: 36px;
  background: linear-gradient(180deg, transparent, #6b5140);
}
.ac-sumi-cord-l { left: 22%; }
.ac-sumi-cord-r { right: 22%; }
.ac-sumi-paper {
  position: relative;
  background:
    linear-gradient(180deg, rgba(255, 250, 236, 0.5), transparent 24%),
    #f3ead6;
  padding: 30px 26px 18px;
  box-shadow:
    0 18px 42px -18px rgba(50, 35, 20, 0.5),
    inset 0 0 60px rgba(168, 140, 95, 0.16);
  border-left: 1px solid rgba(120, 95, 65, 0.14);
  border-right: 1px solid rgba(120, 95, 65, 0.14);
  transition: transform 0.6s ease;
}
.ac-sumi-paper:hover { transform: rotate(-0.3deg) translateY(-2px); }
.ac-sumi-paper::before {
  content: '';
  position: absolute;
  inset: 10px;
  border: 1px solid rgba(120, 95, 65, 0.14);
  pointer-events: none;
}
.ac-sumi-enso {
  position: absolute;
  top: 48px;
  right: 30px;
  display: flex;
  width: 112px;
  height: 112px;
  align-items: center;
  justify-content: center;
  border: 9px solid rgba(33, 29, 24, 0.8);
  border-right-color: rgba(33, 29, 24, 0.18);
  border-radius: 50%;
  transform: rotate(-26deg);
  filter: blur(0.3px);
  opacity: 0.82;
}
.ac-sumi-enso::after {
  content: '';
  position: absolute;
  inset: 7px;
  border: 1px solid rgba(33, 29, 24, 0.2);
  border-radius: 50%;
}
.ac-sumi-initial {
  transform: rotate(26deg);
  font-size: 38px;
  font-weight: 600;
  color: #211d18;
}
.ac-sumi-kicker {
  font-family: Consolas, monospace;
  font-size: 7px;
  letter-spacing: 0.28em;
  color: #8a7a62;
}
.ac-sumi-name {
  position: relative;
  display: inline-block;
  margin-top: 12px;
  font-family: 'Kaiti SC', 'STKaiti', 'Songti SC', serif;
  font-size: 52px;
  font-weight: 700;
  line-height: 1.08;
  letter-spacing: 0.08em;
  color: #1f1b16;
  text-shadow: 1px 1px 0 rgba(255, 248, 230, 0.8);
  overflow-wrap: anywhere;
}
.ac-sumi-seal-name {
  display: inline-flex;
  width: 22px;
  height: 22px;
  margin-left: 9px;
  align-items: center;
  justify-content: center;
  border: 1.5px solid #a52a1e;
  color: #a52a1e;
  font-size: 9px;
  transform: translateY(-12px) rotate(4deg);
  opacity: 0.85;
}
.ac-sumi-vertical {
  position: absolute;
  top: 92px;
  left: 12px;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 4px;
  writing-mode: vertical-rl;
  font-size: 10.5px;
  letter-spacing: 0.24em;
  color: #5d5142;
}
.ac-sumi-title { font-weight: 600; }
.ac-sumi-dot { font-size: 7px; color: #a52a1e; }
.ac-sumi-series { color: #8a7a62; }
.ac-sumi-element {
  display: inline-flex;
  margin-top: 16px;
  border: 1.5px solid #a52a1e;
  padding: 5px 12px;
  color: #a52a1e;
  font-size: 11px;
  letter-spacing: 0.16em;
  transform: rotate(-1.5deg);
  background: rgba(165, 42, 30, 0.03);
}
.ac-sumi-quote {
  margin-top: 32px;
  padding: 0 4px 4px;
}
.ac-sumi-quote p {
  font-family: 'Kaiti SC', 'STKaiti', 'Songti SC', serif;
  font-size: 19px;
  line-height: 1.9;
  letter-spacing: 0.06em;
  color: #2c261f;
  overflow-wrap: anywhere;
}
.ac-sumi-quote svg { width: 58%; height: 9px; margin-top: 3px; }
.ac-sumi-foot {
  display: flex;
  align-items: center;
  gap: 10px;
  margin-top: 28px;
  border-top: 1px solid rgba(120, 95, 65, 0.22);
  padding-top: 13px;
}
.ac-sumi-thread {
  width: 28px;
  height: 1px;
  background: #a52a1e;
  opacity: 0.6;
}
.ac-sumi-foot small {
  display: block;
  font-family: Consolas, monospace;
  font-size: 6px;
  letter-spacing: 0.2em;
  color: #9a8a72;
}
.ac-sumi-foot b {
  display: block;
  margin-top: 2px;
  font-family: Consolas, monospace;
  font-size: 11px;
  font-weight: 400;
  letter-spacing: 0.06em;
  color: #5d5142;
  overflow-wrap: anywhere;
}
.ac-sumi-hanko {
  display: flex;
  width: 30px;
  height: 30px;
  margin-left: auto;
  align-items: center;
  justify-content: center;
  border: 1.5px solid #a52a1e;
  color: #a52a1e;
  font-family: 'Kaiti SC', 'STKaiti', serif;
  font-size: 15px;
  font-weight: 700;
  transform: rotate(5deg);
  opacity: 0.8;
}
.ac-sumi-id {
  margin-top: 12px;
  text-align: right;
  font-size: 7.5px;
  letter-spacing: 0.16em;
  color: #9a8a72;
}
.ac-sumi-caption {
  margin-top: 22px;
  text-align: center;
  font-size: 9px;
  letter-spacing: 0.4em;
  color: rgba(60, 48, 34, 0.4);
  animation: ac-sumi-fade 1.4s 0.8s both;
}
@keyframes ac-sumi-unroll {
  from { opacity: 0; transform: scaleY(0.6) translateY(-14px); }
  to { opacity: 1; transform: none; }
}
@keyframes ac-sumi-fade { from { opacity: 0; } to { opacity: 1; } }
@media (prefers-reduced-motion: reduce) {
  .ac-sumi-root *,
  .ac-sumi-root *::before,
  .ac-sumi-root *::after {
    animation: none !important;
    transition: none !important;
  }
}
</style>
