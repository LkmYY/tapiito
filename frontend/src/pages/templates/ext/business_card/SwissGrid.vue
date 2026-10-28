<template>
  <div class="bc-swiss-root min-h-screen px-6 py-8 text-[#111111]">
    <div class="bc-swiss-wrap mx-auto w-full max-w-md">
      <header class="bc-swiss-head">
        <span class="bc-swiss-logo">ID.</span>
        <span class="bc-swiss-meta">IDENTITY SYSTEM<br />ZÜRICH / SHANGHAI</span>
        <span class="bc-swiss-year">2026</span>
      </header>

      <div class="bc-swiss-bar"></div>

      <section class="bc-swiss-hero">
        <div class="bc-swiss-index">
          <span>01</span>
          <span>NAME</span>
        </div>
        <h1 class="bc-swiss-name">{{ payload.name || '未命名' }}</h1>
        <div class="bc-swiss-mark">✳</div>
      </section>

      <section class="bc-swiss-meta-grid">
        <div class="bc-swiss-cell">
          <span>ROLE</span>
          <p>{{ payload.title }}</p>
        </div>
        <div class="bc-swiss-cell">
          <span>COMPANY</span>
          <p>{{ payload.company }}</p>
        </div>
      </section>

      <section class="bc-swiss-list">
        <div class="bc-swiss-list-head">
          <span>NO.</span>
          <span>CHANNEL</span>
          <span>DETAIL</span>
        </div>
        <div class="bc-swiss-row">
          <span class="bc-swiss-no">02</span>
          <span class="bc-swiss-ch">TEL</span>
          <a v-if="payload.phone" :href="'tel:' + payload.phone">{{ payload.phone }}</a>
        </div>
        <div class="bc-swiss-row">
          <span class="bc-swiss-no">03</span>
          <span class="bc-swiss-ch">MAIL</span>
          <a v-if="payload.email" :href="'mailto:' + payload.email">{{ payload.email }}</a>
        </div>
        <div class="bc-swiss-row">
          <span class="bc-swiss-no">04</span>
          <span class="bc-swiss-ch">WEB</span>
          <a v-if="payload.website" :href="payload.website" target="_blank" rel="noopener noreferrer">{{ payload.website }}</a>
        </div>
        <div class="bc-swiss-row">
          <span class="bc-swiss-no">05</span>
          <span class="bc-swiss-ch">WECHAT</span>
          <span :class="{ 'bc-swiss-empty': !payload.wechat }">{{ payload.wechat }}</span>
        </div>
      </section>

      <section class="bc-swiss-footer">
        <div class="bc-swiss-redbox">
          <span>{{ initial }}</span>
          <small>NFC {{ nfcId || '000' }}</small>
        </div>
        <p>GOOD DESIGN IS<br />CLEAR COMMUNICATION.</p>
        <span class="bc-swiss-foot-num">06</span>
      </section>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps(['payload', 'nfcId'])

const initial = computed(() => (props.payload?.name || '?').toString().trim()[0] || '?')
</script>

<style scoped>
.bc-swiss-root {
  background: #f4f4f2;
  font-family: 'Helvetica Neue', Helvetica, 'PingFang SC', 'Microsoft YaHei', sans-serif;
}
.bc-swiss-wrap {
  animation: bc-swiss-in 0.9s cubic-bezier(0.22, 1, 0.36, 1) both;
}
.bc-swiss-head {
  display: grid;
  grid-template-columns: 44px 1fr auto;
  gap: 12px;
  align-items: start;
  padding-bottom: 14px;
}
.bc-swiss-logo {
  font-size: 30px;
  font-weight: 800;
  line-height: 0.8;
  letter-spacing: -0.08em;
  color: #dc2626;
}
.bc-swiss-meta {
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6.5px;
  line-height: 1.6;
  letter-spacing: 0.14em;
  color: #6b7280;
}
.bc-swiss-year {
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 8px;
  letter-spacing: 0.08em;
}
.bc-swiss-bar {
  height: 8px;
  background: #111111;
  transform-origin: left;
  animation: bc-swiss-bar 0.8s 0.15s cubic-bezier(0.22, 1, 0.36, 1) both;
}
.bc-swiss-hero {
  position: relative;
  padding: 28px 0 20px;
  border-bottom: 1px solid #111111;
}
.bc-swiss-index {
  display: flex;
  align-items: center;
  gap: 12px;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6.5px;
  letter-spacing: 0.14em;
  color: #9ca3af;
}
.bc-swiss-index span:first-child {
  display: inline-flex;
  width: 17px;
  height: 17px;
  align-items: center;
  justify-content: center;
  background: #111111;
  color: #ffffff;
  font-size: 6px;
}
.bc-swiss-name {
  margin-top: 14px;
  font-size: clamp(40px, 13.2vw, 56px);
  font-weight: 700;
  line-height: 0.98;
  letter-spacing: -0.055em;
  overflow-wrap: anywhere;
  text-transform: uppercase;
}
.bc-swiss-mark {
  position: absolute;
  right: 0;
  bottom: 24px;
  font-size: 26px;
  line-height: 1;
  color: #dc2626;
  animation: bc-swiss-spin 14s linear infinite;
}
.bc-swiss-meta-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  border-bottom: 1px solid #111111;
}
.bc-swiss-cell {
  padding: 14px 0;
}
.bc-swiss-cell + .bc-swiss-cell {
  border-left: 1px solid #d1d5db;
  padding-left: 14px;
}
.bc-swiss-cell span,
.bc-swiss-list-head,
.bc-swiss-ch,
.bc-swiss-no {
  font-family: 'SFMono-Regular', Consolas, monospace;
  letter-spacing: 0.12em;
}
.bc-swiss-cell span {
  font-size: 6px;
  color: #9ca3af;
}
.bc-swiss-cell p {
  margin-top: 4px;
  font-size: 12.5px;
  font-weight: 500;
  overflow-wrap: anywhere;
}
.bc-swiss-list {
  padding: 12px 0 4px;
}
.bc-swiss-list-head {
  display: grid;
  grid-template-columns: 28px 62px 1fr;
  padding: 0 0 6px;
  font-size: 5.5px;
  color: #9ca3af;
}
.bc-swiss-row {
  position: relative;
  display: grid;
  grid-template-columns: 28px 62px 1fr;
  align-items: baseline;
  gap: 4px;
  padding: 12px 6px;
  border-top: 1px solid #d1d5db;
  text-decoration: none;
  color: inherit;
  transition: background 0.22s ease, color 0.22s ease, padding-left 0.22s ease;
}
.bc-swiss-row:hover {
  background: #dc2626;
  color: #ffffff;
  padding-left: 10px;
}
.bc-swiss-row:active {
  background: #b91c1c;
}
.bc-swiss-no {
  font-size: 6px;
  color: #9ca3af;
  transition: color 0.22s ease;
}
.bc-swiss-row:hover .bc-swiss-no {
  color: rgba(255, 255, 255, 0.6);
}
.bc-swiss-ch {
  font-size: 6.5px;
  font-weight: 700;
}
.bc-swiss-row > a,
.bc-swiss-row > span:last-child {
  min-width: 0;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  font-size: 11px;
  letter-spacing: 0.01em;
}
.bc-swiss-row > a {
  font-weight: 500;
}
.bc-swiss-empty {
  color: #c4c4c0;
}
.bc-swiss-row:hover .bc-swiss-empty {
  color: rgba(255, 255, 255, 0.55);
}
.bc-swiss-footer {
  display: grid;
  grid-template-columns: 74px 1fr auto;
  gap: 14px;
  align-items: end;
  margin-top: 34px;
  padding-top: 16px;
  border-top: 8px solid #111111;
}
.bc-swiss-redbox {
  display: flex;
  height: 74px;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  gap: 2px;
  background: #dc2626;
  color: #ffffff;
  transition: transform 0.3s ease, background 0.3s ease;
}
.bc-swiss-redbox:hover {
  transform: translateY(-3px);
  background: #b91c1c;
}
.bc-swiss-redbox span {
  font-size: 26px;
  font-weight: 700;
  line-height: 1;
}
.bc-swiss-redbox small {
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 5px;
  letter-spacing: 0.1em;
}
.bc-swiss-footer p {
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6.5px;
  line-height: 1.7;
  letter-spacing: 0.1em;
  color: #4b5563;
}
.bc-swiss-foot-num {
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 7px;
  color: #9ca3af;
}
@keyframes bc-swiss-in {
  from { opacity: 0; transform: translateY(18px); }
  to { opacity: 1; transform: none; }
}
@keyframes bc-swiss-bar {
  from { transform: scaleX(0); }
  to { transform: scaleX(1); }
}
@keyframes bc-swiss-spin {
  from { transform: rotate(0); }
  to { transform: rotate(360deg); }
}
@media (prefers-reduced-motion: reduce) {
  .bc-swiss-root *,
  .bc-swiss-root *::before,
  .bc-swiss-root *::after {
    animation: none !important;
    transition: none !important;
  }
}
</style>
