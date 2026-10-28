<template>
  <div class="bc-glass-root min-h-screen overflow-hidden text-white">
    <div class="bc-glass-aurora">
      <span class="bc-glass-orb bc-glass-orb-a"></span>
      <span class="bc-glass-orb bc-glass-orb-b"></span>
      <span class="bc-glass-orb bc-glass-orb-c"></span>
    </div>

    <div class="bc-glass-wrap mx-auto flex min-h-screen w-full max-w-md flex-col">
      <header class="bc-glass-top">
        <span class="bc-glass-logo">{{ initial }}</span>
        <span class="bc-glass-nfc">NFC · {{ nfcId || '0000' }}</span>
      </header>

      <section class="bc-glass-hero">
        <p class="bc-glass-eyebrow">AURORA IDENTITY</p>
        <h1 class="bc-glass-name">{{ payload.name || '未命名' }}</h1>
        <p class="bc-glass-role">{{ payload.title }}</p>
        <p class="bc-glass-company">{{ payload.company }}</p>
      </section>

      <section class="bc-glass-panel">
        <div class="bc-glass-panel-head">
          <span>CONNECT</span>
          <span class="bc-glass-status"><i></i>AVAILABLE</span>
        </div>

        <div class="bc-glass-grid">
          <a v-if="payload.phone" :href="'tel:' + payload.phone" class="bc-glass-cell bc-glass-wide">
            <span class="bc-glass-icon"><Phone :size="16" /></span>
            <span class="bc-glass-meta"><small>PHONE</small><b>{{ payload.phone }}</b></span>
            <ArrowUpRight :size="15" class="bc-glass-arrow" />
          </a>
          <div v-else class="bc-glass-cell bc-glass-wide bc-glass-dim">
            <span class="bc-glass-icon"><Phone :size="16" /></span>
            <span class="bc-glass-meta"><small>PHONE</small><b></b></span>
          </div>

          <a v-if="payload.email" :href="'mailto:' + payload.email" class="bc-glass-cell">
            <span class="bc-glass-icon"><Mail :size="16" /></span>
            <span class="bc-glass-meta"><small>EMAIL</small><b>{{ payload.email }}</b></span>
          </a>
          <div v-else class="bc-glass-cell bc-glass-dim">
            <span class="bc-glass-icon"><Mail :size="16" /></span>
            <span class="bc-glass-meta"><small>EMAIL</small><b></b></span>
          </div>

          <div class="bc-glass-cell bc-glass-dim">
            <span class="bc-glass-icon"><MessageCircle :size="16" /></span>
            <span class="bc-glass-meta"><small>WECHAT</small><b>{{ payload.wechat }}</b></span>
          </div>

          <a
            v-if="payload.website"
            :href="payload.website"
            target="_blank"
            rel="noopener noreferrer"
            class="bc-glass-cell bc-glass-wide"
          >
            <span class="bc-glass-icon"><Globe :size="16" /></span>
            <span class="bc-glass-meta"><small>WEBSITE</small><b>{{ payload.website }}</b></span>
            <ArrowUpRight :size="15" class="bc-glass-arrow" />
          </a>
          <div v-else class="bc-glass-cell bc-glass-wide bc-glass-dim">
            <span class="bc-glass-icon"><Globe :size="16" /></span>
            <span class="bc-glass-meta"><small>WEBSITE</small><b></b></span>
          </div>
        </div>

        <p class="bc-glass-hint">轻触玻璃面板，建立联系</p>
      </section>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { Phone, Mail, Globe, MessageCircle, ArrowUpRight } from 'lucide-vue-next'

const props = defineProps(['payload', 'nfcId'])

const initial = computed(() => (props.payload?.name || '?').toString().trim()[0] || '?')
</script>

<style scoped>
.bc-glass-root {
  position: relative;
  background: #0b1026;
  font-family: 'Inter', 'PingFang SC', 'Microsoft YaHei', sans-serif;
}
.bc-glass-aurora {
  position: absolute;
  inset: 0;
  overflow: hidden;
  background:
    radial-gradient(90% 60% at 15% 0%, rgba(67, 56, 202, 0.9), transparent 60%),
    radial-gradient(80% 55% at 90% 12%, rgba(14, 165, 233, 0.7), transparent 62%),
    radial-gradient(70% 50% at 50% 44%, rgba(168, 85, 247, 0.42), transparent 65%),
    linear-gradient(180deg, #1e1b4b 0%, #0b1026 58%, #060815 100%);
}
.bc-glass-orb {
  position: absolute;
  border-radius: 50%;
  filter: blur(46px);
  opacity: 0.55;
}
.bc-glass-orb-a {
  top: 6%;
  left: -30px;
  width: 170px;
  height: 170px;
  background: #818cf8;
  animation: bc-glass-drift 12s ease-in-out infinite;
}
.bc-glass-orb-b {
  top: 20%;
  right: -40px;
  width: 200px;
  height: 200px;
  background: #22d3ee;
  animation: bc-glass-drift 15s ease-in-out infinite reverse;
}
.bc-glass-orb-c {
  top: 42%;
  left: 32%;
  width: 130px;
  height: 130px;
  background: #f472b6;
  animation: bc-glass-drift 10s ease-in-out infinite 1s;
}
.bc-glass-wrap {
  position: relative;
  z-index: 2;
  padding: 26px 22px 30px;
}
.bc-glass-top {
  display: flex;
  align-items: center;
  justify-content: space-between;
  animation: bc-glass-fade 0.8s both;
}
.bc-glass-logo {
  display: flex;
  width: 38px;
  height: 38px;
  align-items: center;
  justify-content: center;
  border: 1px solid rgba(255, 255, 255, 0.35);
  border-radius: 12px;
  background: rgba(255, 255, 255, 0.08);
  backdrop-filter: blur(14px);
  font-size: 16px;
  font-weight: 300;
}
.bc-glass-nfc {
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 8px;
  letter-spacing: 0.24em;
  color: rgba(255, 255, 255, 0.5);
}
.bc-glass-hero {
  padding-top: 66px;
  animation: bc-glass-rise 0.9s 0.1s cubic-bezier(0.22, 1, 0.36, 1) both;
}
.bc-glass-eyebrow {
  font-size: 8px;
  letter-spacing: 0.36em;
  color: rgba(255, 255, 255, 0.5);
}
.bc-glass-name {
  margin-top: 16px;
  font-size: clamp(42px, 14vw, 58px);
  font-weight: 200;
  line-height: 1.02;
  letter-spacing: -0.02em;
  overflow-wrap: anywhere;
  text-shadow: 0 10px 40px rgba(15, 23, 42, 0.35);
}
.bc-glass-role {
  margin-top: 16px;
  font-size: 13px;
  font-weight: 500;
  letter-spacing: 0.1em;
  color: #e0e7ff;
}
.bc-glass-company {
  margin-top: 4px;
  font-size: 11px;
  letter-spacing: 0.06em;
  color: rgba(255, 255, 255, 0.55);
}
.bc-glass-panel {
  margin-top: auto;
  border: 1px solid rgba(255, 255, 255, 0.18);
  border-radius: 28px;
  background: linear-gradient(150deg, rgba(255, 255, 255, 0.16), rgba(255, 255, 255, 0.05));
  box-shadow: 0 30px 70px -30px rgba(2, 6, 23, 0.85), inset 0 1px 0 rgba(255, 255, 255, 0.28);
  padding: 22px 18px 16px;
  backdrop-filter: blur(26px) saturate(140%);
  animation: bc-glass-rise 1s 0.24s cubic-bezier(0.22, 1, 0.36, 1) both;
}
.bc-glass-panel-head {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 0 4px 14px;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 8px;
  letter-spacing: 0.26em;
  color: rgba(255, 255, 255, 0.56);
}
.bc-glass-status {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  color: #6ee7b7;
}
.bc-glass-status i {
  width: 5px;
  height: 5px;
  border-radius: 50%;
  background: #34d399;
  box-shadow: 0 0 8px #34d399;
  animation: bc-glass-blink 2.4s ease-in-out infinite;
}
.bc-glass-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 8px;
}
.bc-glass-cell {
  display: flex;
  align-items: center;
  gap: 10px;
  min-height: 66px;
  border: 1px solid rgba(255, 255, 255, 0.13);
  border-radius: 16px;
  background: rgba(255, 255, 255, 0.07);
  padding: 12px;
  text-decoration: none;
  color: inherit;
  transition: background 0.25s ease, border-color 0.25s ease, transform 0.25s ease;
}
.bc-glass-cell:hover {
  border-color: rgba(165, 243, 252, 0.5);
  background: rgba(255, 255, 255, 0.14);
  transform: translateY(-2px);
}
.bc-glass-cell:active {
  transform: translateY(0) scale(0.985);
}
.bc-glass-wide {
  grid-column: 1 / -1;
}
.bc-glass-icon {
  display: flex;
  width: 32px;
  height: 32px;
  flex: none;
  align-items: center;
  justify-content: center;
  border-radius: 10px;
  background: rgba(255, 255, 255, 0.12);
  color: #c7d2fe;
  transition: color 0.25s ease, background 0.25s ease;
}
.bc-glass-cell:hover .bc-glass-icon {
  background: rgba(255, 255, 255, 0.22);
  color: #ffffff;
}
.bc-glass-meta {
  display: flex;
  min-width: 0;
  flex: 1;
  flex-direction: column;
  gap: 2px;
  text-align: left;
}
.bc-glass-meta small {
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 7px;
  letter-spacing: 0.2em;
  color: rgba(255, 255, 255, 0.4);
}
.bc-glass-meta b {
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  font-size: 11.5px;
  font-weight: 400;
  color: rgba(255, 255, 255, 0.92);
}
.bc-glass-arrow {
  flex: none;
  color: rgba(255, 255, 255, 0.35);
  transition: transform 0.25s ease, color 0.25s ease;
}
.bc-glass-cell:hover .bc-glass-arrow {
  color: #ffffff;
  transform: translate(2px, -2px);
}
.bc-glass-dim {
  opacity: 0.55;
}
.bc-glass-hint {
  margin-top: 14px;
  text-align: center;
  font-size: 8px;
  letter-spacing: 0.18em;
  color: rgba(255, 255, 255, 0.3);
}
@keyframes bc-glass-rise {
  from { opacity: 0; transform: translateY(22px); }
  to { opacity: 1; transform: none; }
}
@keyframes bc-glass-fade {
  from { opacity: 0; }
  to { opacity: 1; }
}
@keyframes bc-glass-drift {
  0%, 100% { transform: translate3d(0, 0, 0) scale(1); }
  50% { transform: translate3d(18px, 24px, 0) scale(1.15); }
}
@keyframes bc-glass-blink {
  0%, 100% { opacity: 1; }
  50% { opacity: 0.35; }
}
@media (prefers-reduced-motion: reduce) {
  .bc-glass-root *,
  .bc-glass-root *::before,
  .bc-glass-root *::after {
    animation: none !important;
    transition: none !important;
  }
}
</style>
