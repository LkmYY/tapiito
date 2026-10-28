<template>
  <div class="bc-riso-root min-h-screen px-5 py-8">
    <div class="bc-riso-grain"></div>
    <div class="bc-riso-wrap mx-auto flex min-h-[calc(100vh-4rem)] w-full max-w-md flex-col">
      <header class="bc-riso-head">
        <span class="bc-riso-studio">RISO CLUB</span>
        <span class="bc-riso-edition">PRINT NO. {{ nfcId || '07' }} / 2026</span>
      </header>

      <section class="bc-riso-poster">
        <div class="bc-riso-blob bc-riso-blob-pink"></div>
        <div class="bc-riso-blob bc-riso-blob-blue"></div>

        <p class="bc-riso-kicker">AN INTRODUCTION IN TWO INKS</p>
        <div class="bc-riso-name-wrap">
          <h1 class="bc-riso-name">{{ payload.name || '未命名' }}</h1>
          <span class="bc-riso-name-ghost" aria-hidden="true">{{ payload.name || '未命名' }}</span>
        </div>

        <div class="bc-riso-role-block">
          <span class="bc-riso-role-mark">✳</span>
          <p>{{ payload.title }}</p>
          <p>{{ payload.company }}</p>
        </div>

        <div class="bc-riso-dots"></div>
      </section>

      <section class="bc-riso-info">
        <div class="bc-riso-info-head">
          <span>CONTACT REGISTRATION</span>
          <span>KEEP THIS PRINT</span>
        </div>
        <div class="bc-riso-row">
          <span class="bc-riso-num">01</span>
          <span class="bc-riso-key">PHONE</span>
          <a v-if="payload.phone" :href="'tel:' + payload.phone">{{ payload.phone }}</a>
        </div>
        <div class="bc-riso-row">
          <span class="bc-riso-num">02</span>
          <span class="bc-riso-key">EMAIL</span>
          <a v-if="payload.email" :href="'mailto:' + payload.email">{{ payload.email }}</a>
        </div>
        <div class="bc-riso-row">
          <span class="bc-riso-num">03</span>
          <span class="bc-riso-key">WEBSITE</span>
          <a v-if="payload.website" :href="payload.website" target="_blank" rel="noopener noreferrer">{{ payload.website }}</a>
        </div>
        <div class="bc-riso-row">
          <span class="bc-riso-num">04</span>
          <span class="bc-riso-key">WECHAT</span>
          <span :class="{ 'bc-riso-empty': !payload.wechat }">{{ payload.wechat }}</span>
        </div>
      </section>

      <footer class="bc-riso-foot">
        <span class="bc-riso-stamp">PRINTED<br />WITH<br />CARE</span>
        <p>FLUORESCENT PINK 801<br />MEDIUM BLUE 300<br />100% RECYCLED STOCK</p>
        <span class="bc-riso-initial">{{ initial }}</span>
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
.bc-riso-root {
  position: relative;
  overflow: hidden;
  background: #f3f0e8;
  font-family: 'Helvetica Neue', 'PingFang SC', 'Microsoft YaHei', sans-serif;
  color: #17130f;
}
.bc-riso-root::before {
  content: '';
  position: absolute;
  top: -80px;
  right: -90px;
  width: 280px;
  height: 280px;
  border-radius: 50%;
  background: rgba(236, 72, 153, 0.14);
  filter: blur(6px);
}
.bc-riso-grain {
  position: absolute;
  inset: 0;
  background-image: radial-gradient(rgba(23, 19, 15, 0.28) 0.45px, transparent 0.45px);
  background-size: 3px 3px;
  opacity: 0.24;
  pointer-events: none;
  z-index: 4;
}
.bc-riso-wrap {
  position: relative;
  z-index: 1;
}
.bc-riso-head {
  display: flex;
  justify-content: space-between;
  padding-bottom: 14px;
  border-bottom: 2px solid #17130f;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 7px;
  letter-spacing: 0.18em;
  animation: bc-riso-fade 0.8s both;
}
.bc-riso-studio {
  font-weight: 700;
  color: #db2777;
}
.bc-riso-edition {
  color: #2563eb;
}
.bc-riso-poster {
  position: relative;
  overflow: hidden;
  padding: 30px 4px 20px;
  border-bottom: 2px solid #17130f;
  animation: bc-riso-in 0.9s 0.1s cubic-bezier(0.22, 1, 0.36, 1) both;
}
.bc-riso-blob {
  position: absolute;
  border-radius: 50%;
  mix-blend-mode: multiply;
  opacity: 0.7;
}
.bc-riso-blob-pink {
  top: 44px;
  right: -30px;
  width: 220px;
  height: 220px;
  background: radial-gradient(circle at 45% 45%, #f472b6, rgba(244, 114, 182, 0.15) 68%);
}
.bc-riso-blob-blue {
  left: -60px;
  bottom: -30px;
  width: 200px;
  height: 200px;
  background: radial-gradient(circle at 50% 50%, #38bdf8, rgba(56, 189, 248, 0.12) 70%);
}
.bc-riso-kicker {
  position: relative;
  z-index: 1;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6.5px;
  letter-spacing: 0.24em;
  color: #2563eb;
}
.bc-riso-name-wrap {
  position: relative;
  z-index: 1;
  margin-top: 16px;
}
.bc-riso-name,
.bc-riso-name-ghost {
  font-size: clamp(46px, 15vw, 64px);
  font-weight: 800;
  line-height: 0.92;
  letter-spacing: -0.06em;
  overflow-wrap: anywhere;
  text-transform: uppercase;
}
.bc-riso-name {
  position: relative;
  z-index: 2;
  color: #17130f;
  text-shadow: 3px 2px 0 rgba(236, 72, 153, 0.85), -3px -2px 0 rgba(37, 99, 235, 0.7);
  transition: text-shadow 0.3s ease, transform 0.3s ease;
}
.bc-riso-name:hover {
  text-shadow: 5px 4px 0 rgba(236, 72, 153, 0.9), -5px -4px 0 rgba(37, 99, 235, 0.8);
  transform: translateX(1px);
}
.bc-riso-name-ghost {
  position: absolute;
  top: 0;
  left: 3px;
  z-index: 0;
  color: rgba(37, 99, 235, 0.18);
  pointer-events: none;
}
.bc-riso-role-block {
  position: relative;
  z-index: 1;
  display: flex;
  align-items: flex-start;
  gap: 8px;
  margin-top: 24px;
}
.bc-riso-role-mark {
  color: #db2777;
  font-size: 13px;
  line-height: 1.3;
}
.bc-riso-role-block p {
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 9px;
  line-height: 1.6;
  letter-spacing: 0.04em;
  color: #17130f;
}
.bc-riso-role-block p + p {
  margin-left: 8px;
  color: rgba(23, 19, 15, 0.5);
}
.bc-riso-dots {
  position: relative;
  z-index: 1;
  height: 14px;
  margin-top: 24px;
  background-image: radial-gradient(#17130f 1px, transparent 1px);
  background-size: 7px 7px;
  opacity: 0.7;
}
.bc-riso-info {
  padding-top: 14px;
  animation: bc-riso-fade 1s 0.3s both;
}
.bc-riso-info-head {
  display: flex;
  justify-content: space-between;
  padding-bottom: 6px;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 5.5px;
  letter-spacing: 0.16em;
  color: rgba(23, 19, 15, 0.5);
}
.bc-riso-row {
  display: grid;
  grid-template-columns: 26px 62px 1fr;
  align-items: baseline;
  gap: 4px;
  padding: 10px 4px;
  border-top: 1px solid rgba(23, 19, 15, 0.28);
  transition: background 0.2s ease, padding-left 0.2s ease;
}
.bc-riso-row:hover {
  background: rgba(244, 114, 182, 0.14);
  padding-left: 8px;
}
.bc-riso-num {
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 5.5px;
  color: #db2777;
}
.bc-riso-key {
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6px;
  letter-spacing: 0.12em;
  color: #2563eb;
}
.bc-riso-row a,
.bc-riso-row span:last-child,
.bc-riso-row em {
  min-width: 0;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  font-size: 10px;
  font-style: normal;
  letter-spacing: 0.01em;
  color: #17130f;
  text-decoration: none;
  transition: color 0.2s ease;
}
.bc-riso-row a {
  font-weight: 600;
}
.bc-riso-row a:hover {
  color: #db2777;
}
.bc-riso-row a:active {
  opacity: 0.6;
}
.bc-riso-row em,
.bc-riso-empty {
  color: rgba(23, 19, 15, 0.3);
  font-weight: 400;
}
.bc-riso-foot {
  display: flex;
  align-items: flex-end;
  gap: 14px;
  margin-top: auto;
  padding-top: 26px;
}
.bc-riso-stamp {
  border: 2px solid #db2777;
  padding: 5px 7px;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6px;
  line-height: 1.4;
  letter-spacing: 0.1em;
  color: #db2777;
  transform: rotate(-5deg);
  transition: transform 0.3s ease, background 0.3s ease, color 0.3s ease;
}
.bc-riso-stamp:hover {
  transform: rotate(2deg);
  background: #db2777;
  color: #f3f0e8;
}
.bc-riso-foot p {
  flex: 1;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 5.5px;
  line-height: 1.8;
  letter-spacing: 0.08em;
  color: rgba(23, 19, 15, 0.45);
}
.bc-riso-initial {
  display: flex;
  width: 32px;
  height: 32px;
  align-items: center;
  justify-content: center;
  border-radius: 50%;
  border: 1px solid #2563eb;
  font-family: Georgia, serif;
  font-size: 13px;
  color: #2563eb;
}
@keyframes bc-riso-in {
  from { opacity: 0; transform: translateY(20px); }
  to { opacity: 1; transform: none; }
}
@keyframes bc-riso-fade {
  from { opacity: 0; }
  to { opacity: 1; }
}
@media (prefers-reduced-motion: reduce) {
  .bc-riso-root *,
  .bc-riso-root *::before,
  .bc-riso-root *::after {
    animation: none !important;
    transition: none !important;
  }
}
</style>
