<template>
  <div class="bc-brutal-root min-h-screen px-4 py-6">
    <div class="bc-brutal-wrap mx-auto w-full max-w-md">
      <header class="bc-brutal-top">
        <span class="bc-brutal-tag">BC—25</span>
        <span class="bc-brutal-top-text">LOUD &amp; CLEAR</span>
      </header>

      <section class="bc-brutal-hero">
        <div class="bc-brutal-sticker">HEY!</div>
        <p class="bc-brutal-kicker">BUSINESS / CARD</p>
        <h1 class="bc-brutal-name">{{ payload.name || '未命名' }}</h1>
        <div class="bc-brutal-inset">
          <p>{{ payload.title }}</p>
          <span>{{ payload.company }}</span>
        </div>
        <div class="bc-brutal-badge">{{ initial }}</div>
      </section>

      <div class="bc-brutal-marquee">
        <span>LET'S TALK · LET'S BUILD · LET'S TALK · LET'S BUILD ·</span>
      </div>

      <section class="bc-brutal-grid">
        <a v-if="payload.phone" :href="'tel:' + payload.phone" class="bc-brutal-cell bc-brutal-blue">
          <span class="bc-brutal-cell-label">PHONE / 电话</span>
          <b>{{ payload.phone }}</b>
          <i>↗</i>
        </a>
        <div v-else class="bc-brutal-cell bc-brutal-blue bc-brutal-off">
          <span class="bc-brutal-cell-label">PHONE / 电话</span>
          <b></b>
        </div>

        <a v-if="payload.email" :href="'mailto:' + payload.email" class="bc-brutal-cell">
          <span class="bc-brutal-cell-label">EMAIL / 邮箱</span>
          <b>{{ payload.email }}</b>
          <i>↗</i>
        </a>
        <div v-else class="bc-brutal-cell bc-brutal-off">
          <span class="bc-brutal-cell-label">EMAIL / 邮箱</span>
          <b></b>
        </div>

        <a v-if="payload.website" :href="payload.website" target="_blank" rel="noopener noreferrer" class="bc-brutal-cell">
          <span class="bc-brutal-cell-label">WEB / 网站</span>
          <b>{{ payload.website }}</b>
          <i>↗</i>
        </a>
        <div v-else class="bc-brutal-cell bc-brutal-off">
          <span class="bc-brutal-cell-label">WEB / 网站</span>
          <b></b>
        </div>

        <div class="bc-brutal-cell bc-brutal-pink">
          <span class="bc-brutal-cell-label">WECHAT / 微信</span>
          <b>{{ payload.wechat }}</b>
          <i>✳</i>
        </div>
      </section>

      <footer class="bc-brutal-foot">
        <span class="bc-brutal-foot-block">SAY HI</span>
        <p>NO SMALL TALK.<br />JUST GOOD WORK.</p>
        <span class="bc-brutal-foot-num">{{ nfcId || '00' }}</span>
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
.bc-brutal-root {
  position: relative;
  overflow: hidden;
  background: #facc15;
  font-family: 'Arial Black', 'Helvetica Neue', 'PingFang SC', sans-serif;
  color: #0a0a0a;
}
.bc-brutal-root::before {
  content: '';
  position: absolute;
  top: 0;
  right: 0;
  width: 38%;
  height: 100%;
  background: #2563eb;
  clip-path: polygon(38% 0, 100% 0, 100% 100%, 0 100%);
  opacity: 0.95;
}
.bc-brutal-wrap {
  position: relative;
  z-index: 1;
  animation: bc-brutal-in 0.8s cubic-bezier(0.22, 1, 0.36, 1) both;
}
.bc-brutal-top {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 0 2px 14px;
}
.bc-brutal-tag {
  border: 3px solid #0a0a0a;
  background: #ffffff;
  padding: 4px 8px;
  font-family: 'Courier New', monospace;
  font-size: 8px;
  font-weight: 700;
  letter-spacing: 0.1em;
  box-shadow: 3px 3px 0 #0a0a0a;
}
.bc-brutal-top-text {
  font-size: 9px;
  letter-spacing: 0.16em;
  color: #ffffff;
  text-shadow: 1px 1px 0 #0a0a0a;
}
.bc-brutal-hero {
  position: relative;
  margin-top: 8px;
  border: 4px solid #0a0a0a;
  background: #ffffff;
  padding: 22px 18px 20px;
  box-shadow: 8px 8px 0 #0a0a0a;
}
.bc-brutal-sticker {
  position: absolute;
  top: -20px;
  right: 16px;
  display: flex;
  width: 64px;
  height: 64px;
  align-items: center;
  justify-content: center;
  border: 3px solid #0a0a0a;
  border-radius: 50%;
  background: #f472b6;
  font-size: 12px;
  letter-spacing: 0.02em;
  transform: rotate(10deg);
  box-shadow: 3px 3px 0 #0a0a0a;
  animation: bc-brutal-wobble 4s ease-in-out infinite;
}
.bc-brutal-kicker {
  font-family: 'Courier New', monospace;
  font-size: 8px;
  font-weight: 700;
  letter-spacing: 0.18em;
}
.bc-brutal-name {
  margin-top: 8px;
  font-size: clamp(40px, 14vw, 58px);
  font-weight: 900;
  line-height: 0.9;
  letter-spacing: -0.06em;
  text-transform: uppercase;
  overflow-wrap: anywhere;
}
.bc-brutal-inset {
  display: inline-block;
  margin-top: 18px;
  border: 3px solid #0a0a0a;
  background: #facc15;
  padding: 6px 10px;
  transform: rotate(-1.5deg);
}
.bc-brutal-inset p {
  font-size: 11px;
  font-weight: 900;
  letter-spacing: 0.02em;
}
.bc-brutal-inset span {
  display: block;
  margin-top: 2px;
  font-family: 'Courier New', monospace;
  font-size: 8px;
  font-weight: 400;
  letter-spacing: 0.06em;
}
.bc-brutal-badge {
  position: absolute;
  right: 12px;
  bottom: 12px;
  display: flex;
  width: 54px;
  height: 54px;
  align-items: center;
  justify-content: center;
  border: 3px solid #0a0a0a;
  background: #2563eb;
  color: #ffffff;
  font-size: 24px;
  box-shadow: 4px 4px 0 #0a0a0a;
  transition: transform 0.25s ease, box-shadow 0.25s ease;
}
.bc-brutal-badge:hover {
  transform: translate(2px, 2px);
  box-shadow: 2px 2px 0 #0a0a0a;
}
.bc-brutal-marquee {
  overflow: hidden;
  margin: 14px 0;
  border-top: 3px solid #0a0a0a;
  border-bottom: 3px solid #0a0a0a;
  background: #0a0a0a;
  padding: 5px 0;
  white-space: nowrap;
}
.bc-brutal-marquee span {
  display: inline-block;
  color: #facc15;
  font-family: 'Courier New', monospace;
  font-size: 9px;
  font-weight: 700;
  letter-spacing: 0.14em;
  animation: bc-brutal-scroll 14s linear infinite;
}
.bc-brutal-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 10px;
}
.bc-brutal-cell {
  position: relative;
  display: flex;
  min-height: 98px;
  flex-direction: column;
  justify-content: space-between;
  border: 3px solid #0a0a0a;
  background: #ffffff;
  padding: 11px 10px;
  text-decoration: none;
  color: #0a0a0a;
  box-shadow: 5px 5px 0 #0a0a0a;
  transition: transform 0.2s ease, box-shadow 0.2s ease, background 0.2s ease;
}
.bc-brutal-cell:hover {
  transform: translate(3px, 3px);
  box-shadow: 2px 2px 0 #0a0a0a;
  background: #fef9c3;
}
.bc-brutal-cell:active {
  transform: translate(4px, 4px);
  box-shadow: 1px 1px 0 #0a0a0a;
}
.bc-brutal-blue {
  background: #2563eb;
  color: #ffffff;
}
.bc-brutal-blue:hover {
  background: #1d4ed8;
}
.bc-brutal-pink {
  background: #f472b6;
}
.bc-brutal-pink:hover {
  background: #f9a8d4;
}
.bc-brutal-cell-label {
  font-family: 'Courier New', monospace;
  font-size: 6.5px;
  font-weight: 700;
  letter-spacing: 0.08em;
  opacity: 0.7;
}
.bc-brutal-cell b {
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  font-size: 10.5px;
  font-weight: 900;
  letter-spacing: -0.01em;
}
.bc-brutal-cell i {
  position: absolute;
  right: 8px;
  bottom: 7px;
  font-style: normal;
  font-size: 15px;
}
.bc-brutal-off {
  opacity: 0.55;
}
.bc-brutal-foot {
  display: grid;
  grid-template-columns: auto 1fr auto;
  gap: 12px;
  align-items: center;
  margin-top: 22px;
}
.bc-brutal-foot-block {
  border: 3px solid #0a0a0a;
  background: #f472b6;
  padding: 7px 10px;
  font-size: 11px;
  letter-spacing: 0.04em;
  box-shadow: 3px 3px 0 #0a0a0a;
  transition: background 0.2s ease, transform 0.2s ease;
}
.bc-brutal-foot-block:hover {
  background: #ffffff;
  transform: rotate(-2deg);
}
.bc-brutal-foot p {
  font-family: 'Courier New', monospace;
  font-size: 7px;
  font-weight: 700;
  line-height: 1.5;
  letter-spacing: 0.06em;
}
.bc-brutal-foot-num {
  font-family: 'Courier New', monospace;
  font-size: 9px;
  font-weight: 700;
}
@keyframes bc-brutal-in {
  from { opacity: 0; transform: translateY(22px) rotate(-1deg); }
  to { opacity: 1; transform: none; }
}
@keyframes bc-brutal-wobble {
  0%, 100% { transform: rotate(10deg); }
  50% { transform: rotate(4deg) scale(1.04); }
}
@keyframes bc-brutal-scroll {
  from { transform: translateX(0); }
  to { transform: translateX(-50%); }
}
@media (prefers-reduced-motion: reduce) {
  .bc-brutal-root *,
  .bc-brutal-root *::before,
  .bc-brutal-root *::after {
    animation: none !important;
    transition: none !important;
  }
}
</style>
