<template>
  <div class="bc-washi-root min-h-screen px-6 py-8">
    <div class="bc-washi-wrap mx-auto w-full max-w-md">
      <header class="bc-washi-head">
        <span>京 の 名 帖</span>
        <span class="bc-washi-head-en">KYOTO / WASHI</span>
        <span class="bc-washi-head-num">{{ nfcId || '〇七' }}</span>
      </header>

      <div class="bc-washi-sheet">
        <div class="bc-washi-vertical">
          <p class="bc-washi-label">ご挨拶</p>
          <h1 class="bc-washi-name">{{ payload.name || '未命名' }}</h1>
          <p class="bc-washi-role">{{ payload.title }}</p>
          <p class="bc-washi-company">{{ payload.company }}</p>
        </div>

        <div class="bc-washi-side">
          <div class="bc-washi-seal">
            <span>{{ initial }}</span>
          </div>
          <p class="bc-washi-side-text">一筆啓上<br />よろしくお願い申し上げます</p>
          <div class="bc-washi-thread"></div>
        </div>
      </div>

      <section class="bc-washi-contacts">
        <div class="bc-washi-row">
          <span class="bc-washi-key">電 話</span>
          <a v-if="payload.phone" :href="'tel:' + payload.phone">{{ payload.phone }}</a>
        </div>
        <div class="bc-washi-row">
          <span class="bc-washi-key">メール</span>
          <a v-if="payload.email" :href="'mailto:' + payload.email">{{ payload.email }}</a>
        </div>
        <div class="bc-washi-row">
          <span class="bc-washi-key">網 址</span>
          <a v-if="payload.website" :href="payload.website" target="_blank" rel="noopener noreferrer">{{ payload.website }}</a>
        </div>
        <div class="bc-washi-row">
          <span class="bc-washi-key">微 信</span>
          <span :class="{ 'bc-washi-empty': !payload.wechat }">{{ payload.wechat }}</span>
        </div>
      </section>

      <footer class="bc-washi-foot">
        <span>雲図 · 二〇二六</span>
        <span class="bc-washi-line"></span>
        <span>京都 御所南</span>
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
.bc-washi-root {
  position: relative;
  overflow: hidden;
  background: #ece2d0;
  font-family: 'Yu Mincho', 'Hiragino Mincho ProN', 'Songti SC', serif;
  color: #2c241b;
}
.bc-washi-root::before {
  content: '';
  position: absolute;
  inset: 0;
  background-image:
    repeating-linear-gradient(92deg, rgba(120, 90, 60, 0.055) 0 1px, transparent 1px 7px),
    radial-gradient(rgba(120, 90, 60, 0.18) 0.5px, transparent 0.5px);
  background-size: auto, 4px 4px;
  opacity: 0.8;
}
.bc-washi-wrap {
  position: relative;
  z-index: 1;
  animation: bc-washi-in 1s cubic-bezier(0.22, 1, 0.36, 1) both;
}
.bc-washi-head {
  display: flex;
  align-items: baseline;
  gap: 12px;
  padding-bottom: 14px;
  border-bottom: 1px solid rgba(44, 36, 27, 0.25);
}
.bc-washi-head > span:first-child {
  font-size: 13px;
  letter-spacing: 0.3em;
}
.bc-washi-head-en {
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6px;
  letter-spacing: 0.18em;
  color: rgba(44, 36, 27, 0.45);
}
.bc-washi-head-num {
  margin-left: auto;
  font-size: 10px;
  color: #9f1239;
}
.bc-washi-sheet {
  display: grid;
  grid-template-columns: 1fr 84px;
  gap: 12px;
  padding: 28px 0 20px;
  border-bottom: 2px solid rgba(44, 36, 27, 0.6);
}
.bc-washi-vertical {
  display: flex;
  flex-direction: row-reverse;
  align-items: flex-start;
  justify-content: flex-start;
  gap: 8px;
  min-height: 268px;
}
.bc-washi-label {
  writing-mode: vertical-rl;
  font-size: 9px;
  letter-spacing: 0.4em;
  color: rgba(159, 18, 57, 0.75);
}
.bc-washi-name {
  writing-mode: vertical-rl;
  font-size: clamp(38px, 12vw, 50px);
  font-weight: 500;
  line-height: 1.25;
  letter-spacing: 0.18em;
  color: #1c1917;
  max-height: 260px;
  overflow: hidden;
}
.bc-washi-role {
  writing-mode: vertical-rl;
  margin-top: 14px;
  font-size: 10px;
  letter-spacing: 0.24em;
  color: rgba(44, 36, 27, 0.7);
}
.bc-washi-company {
  writing-mode: vertical-rl;
  margin-top: 10px;
  font-size: 9px;
  letter-spacing: 0.2em;
  color: rgba(44, 36, 27, 0.45);
}
.bc-washi-side {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 16px;
  padding-top: 4px;
}
.bc-washi-seal {
  display: flex;
  width: 58px;
  height: 58px;
  align-items: center;
  justify-content: center;
  border: 2px solid #9f1239;
  border-radius: 3px;
  background: rgba(159, 18, 57, 0.06);
  transform: rotate(3deg);
  transition: transform 0.4s ease, background 0.4s ease;
}
.bc-washi-seal::before {
  content: '';
  position: absolute;
  width: 50px;
  height: 50px;
  border: 1px solid rgba(159, 18, 57, 0.45);
}
.bc-washi-seal span {
  font-size: 27px;
  color: #9f1239;
  transform: rotate(-3deg);
}
.bc-washi-seal:hover {
  transform: rotate(-3deg) scale(1.05);
  background: rgba(159, 18, 57, 0.13);
}
.bc-washi-side-text {
  writing-mode: vertical-rl;
  font-size: 9px;
  line-height: 1.9;
  letter-spacing: 0.16em;
  color: rgba(44, 36, 27, 0.55);
}
.bc-washi-thread {
  flex: 1;
  width: 1px;
  margin-top: 2px;
  background: linear-gradient(180deg, rgba(159, 18, 57, 0.4), transparent);
}
.bc-washi-contacts {
  padding: 10px 0 4px;
}
.bc-washi-row {
  display: flex;
  align-items: baseline;
  gap: 18px;
  padding: 13px 2px;
  border-bottom: 1px dotted rgba(44, 36, 27, 0.25);
}
.bc-washi-key {
  flex: none;
  width: 44px;
  font-size: 8px;
  letter-spacing: 0.16em;
  color: rgba(44, 36, 27, 0.5);
}
.bc-washi-row a,
.bc-washi-row span,
.bc-washi-row em {
  min-width: 0;
  flex: 1;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 11px;
  font-style: normal;
  letter-spacing: 0.02em;
  color: #2c241b;
  text-decoration: none;
  transition: color 0.25s ease, padding-left 0.25s ease;
}
.bc-washi-row a:hover {
  color: #9f1239;
  padding-left: 4px;
}
.bc-washi-row a:active {
  opacity: 0.6;
}
.bc-washi-row em,
.bc-washi-empty {
  color: rgba(44, 36, 27, 0.3);
}
.bc-washi-foot {
  display: flex;
  align-items: center;
  gap: 12px;
  padding-top: 28px;
  font-size: 7px;
  letter-spacing: 0.2em;
  color: rgba(44, 36, 27, 0.45);
}
.bc-washi-line {
  flex: 1;
  height: 1px;
  background: rgba(44, 36, 27, 0.2);
}
@keyframes bc-washi-in {
  from { opacity: 0; transform: translateY(20px); }
  to { opacity: 1; transform: none; }
}
@media (prefers-reduced-motion: reduce) {
  .bc-washi-root *,
  .bc-washi-root *::before,
  .bc-washi-root *::after {
    animation: none !important;
    transition: none !important;
  }
}
</style>
