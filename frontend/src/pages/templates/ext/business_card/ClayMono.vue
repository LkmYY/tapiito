<template>
  <div class="bc-clay-root min-h-screen px-6 py-9">
    <div class="bc-clay-blob bc-clay-blob-a"></div>
    <div class="bc-clay-blob bc-clay-blob-b"></div>

    <div class="bc-clay-wrap mx-auto flex min-h-[calc(100vh-4.5rem)] w-full max-w-md flex-col">
      <header class="bc-clay-head">
        <span class="bc-clay-brand">soft form</span>
        <span class="bc-clay-nfc">{{ nfcId ? 'NFC ' + nfcId : 'STUDIO CLAY' }}</span>
      </header>

      <section class="bc-clay-hero">
        <div class="bc-clay-avatar">{{ initial }}</div>
        <h1 class="bc-clay-name">{{ payload.name || '未命名' }}</h1>
        <p class="bc-clay-role">{{ payload.title }}</p>
        <p class="bc-clay-company">{{ payload.company }}</p>
      </section>

      <section class="bc-clay-cards">
        <a v-if="payload.phone" :href="'tel:' + payload.phone" class="bc-clay-card">
          <span class="bc-clay-icon"><Phone :size="17" /></span>
          <span class="bc-clay-text"><small>电话</small><b>{{ payload.phone }}</b></span>
          <span class="bc-clay-press">CALL</span>
        </a>
        <div v-else class="bc-clay-card bc-clay-card-off">
          <span class="bc-clay-icon"><Phone :size="17" /></span>
          <span class="bc-clay-text"><small>电话</small><b></b></span>
        </div>

        <a v-if="payload.email" :href="'mailto:' + payload.email" class="bc-clay-card">
          <span class="bc-clay-icon"><Mail :size="17" /></span>
          <span class="bc-clay-text"><small>邮箱</small><b>{{ payload.email }}</b></span>
          <span class="bc-clay-press">MAIL</span>
        </a>
        <div v-else class="bc-clay-card bc-clay-card-off">
          <span class="bc-clay-icon"><Mail :size="17" /></span>
          <span class="bc-clay-text"><small>邮箱</small><b></b></span>
        </div>

        <a v-if="payload.website" :href="payload.website" target="_blank" rel="noopener noreferrer" class="bc-clay-card">
          <span class="bc-clay-icon"><Globe :size="17" /></span>
          <span class="bc-clay-text"><small>网站</small><b>{{ payload.website }}</b></span>
          <span class="bc-clay-press">OPEN</span>
        </a>
        <div v-else class="bc-clay-card bc-clay-card-off">
          <span class="bc-clay-icon"><Globe :size="17" /></span>
          <span class="bc-clay-text"><small>网站</small><b></b></span>
        </div>

        <div class="bc-clay-card bc-clay-card-off">
          <span class="bc-clay-icon"><MessageCircle :size="17" /></span>
          <span class="bc-clay-text"><small>微信</small><b>{{ payload.wechat }}</b></span>
          <span class="bc-clay-press">WX</span>
        </div>
      </section>

      <footer class="bc-clay-foot">
        <span class="bc-clay-pill">{{ initial }}</span>
        <p>MADE SOFTLY<br />IN SHANGHAI</p>
        <span class="bc-clay-year">2026</span>
      </footer>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { Phone, Mail, Globe, MessageCircle } from 'lucide-vue-next'

const props = defineProps(['payload', 'nfcId'])

const initial = computed(() => (props.payload?.name || '?').toString().trim()[0] || '?')
</script>

<style scoped>
.bc-clay-root {
  position: relative;
  overflow: hidden;
  background: #e6e1da;
  font-family: 'Nunito', 'Avenir Next Rounded', 'PingFang SC', 'Microsoft YaHei', sans-serif;
  color: #4a4540;
}
.bc-clay-root::before {
  content: '';
  position: absolute;
  inset: 0;
  background-image: radial-gradient(rgba(74, 69, 64, 0.14) 0.5px, transparent 0.5px);
  background-size: 5px 5px;
  opacity: 0.2;
}
.bc-clay-blob {
  position: absolute;
  border-radius: 50%;
  filter: blur(2px);
  opacity: 0.65;
}
.bc-clay-blob-a {
  top: -60px;
  right: -50px;
  width: 180px;
  height: 180px;
  background: linear-gradient(145deg, #f5ede4, #d6cec4);
  box-shadow: inset -12px -14px 28px rgba(120, 110, 98, 0.2), inset 10px 10px 22px rgba(255, 255, 255, 0.7);
}
.bc-clay-blob-b {
  left: -70px;
  bottom: 60px;
  width: 150px;
  height: 150px;
  background: linear-gradient(145deg, #e8ded2, #cfc5b8);
  box-shadow: inset -10px -10px 24px rgba(120, 110, 98, 0.18), inset 8px 8px 20px rgba(255, 255, 255, 0.65);
}
.bc-clay-wrap {
  position: relative;
  z-index: 1;
}
.bc-clay-head {
  display: flex;
  justify-content: space-between;
  align-items: baseline;
  padding-bottom: 16px;
  animation: bc-clay-fade 0.9s both;
}
.bc-clay-brand {
  font-size: 17px;
  font-weight: 700;
  letter-spacing: -0.03em;
  color: #5c554d;
}
.bc-clay-nfc {
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6.5px;
  letter-spacing: 0.18em;
  color: rgba(74, 69, 64, 0.45);
}
.bc-clay-hero {
  padding: 32px 8px 26px;
  text-align: center;
  animation: bc-clay-in 0.9s 0.08s cubic-bezier(0.22, 1, 0.36, 1) both;
}
.bc-clay-avatar {
  display: flex;
  width: 86px;
  height: 86px;
  margin: 0 auto;
  align-items: center;
  justify-content: center;
  border-radius: 28px;
  background: #ece6de;
  box-shadow: 8px 8px 18px rgba(120, 110, 98, 0.24), -8px -8px 18px rgba(255, 255, 255, 0.75);
  font-size: 30px;
  font-weight: 300;
  color: #70675e;
  transition: box-shadow 0.35s ease, transform 0.35s ease;
}
.bc-clay-avatar:hover {
  transform: translateY(-3px);
  box-shadow: 10px 12px 24px rgba(120, 110, 98, 0.28), -8px -8px 20px rgba(255, 255, 255, 0.8);
}
.bc-clay-name {
  margin-top: 24px;
  font-size: clamp(34px, 11vw, 44px);
  font-weight: 700;
  line-height: 1.05;
  letter-spacing: -0.035em;
  color: #403a35;
  overflow-wrap: anywhere;
}
.bc-clay-role {
  margin-top: 10px;
  font-size: 12px;
  font-weight: 600;
  letter-spacing: 0.04em;
  color: #7a7167;
}
.bc-clay-company {
  margin-top: 3px;
  font-size: 10px;
  color: rgba(74, 69, 64, 0.45);
}
.bc-clay-cards {
  display: flex;
  flex-direction: column;
  gap: 11px;
  animation: bc-clay-in 1s 0.22s cubic-bezier(0.22, 1, 0.36, 1) both;
}
.bc-clay-card {
  display: flex;
  align-items: center;
  gap: 13px;
  min-height: 64px;
  border-radius: 22px;
  background: #eae4dc;
  padding: 10px 14px;
  text-decoration: none;
  color: inherit;
  box-shadow: 7px 7px 16px rgba(120, 110, 98, 0.2), -6px -6px 14px rgba(255, 255, 255, 0.68);
  transition: box-shadow 0.25s ease, transform 0.25s ease;
}
.bc-clay-card:hover {
  transform: translateY(-2px);
  box-shadow: 9px 11px 20px rgba(120, 110, 98, 0.26), -6px -6px 15px rgba(255, 255, 255, 0.75);
}
.bc-clay-card:active {
  transform: translateY(1px);
  box-shadow: inset 5px 5px 12px rgba(120, 110, 98, 0.24), inset -5px -5px 12px rgba(255, 255, 255, 0.7);
}
.bc-clay-icon {
  display: flex;
  width: 42px;
  height: 42px;
  flex: none;
  align-items: center;
  justify-content: center;
  border-radius: 15px;
  background: #e4ddd4;
  box-shadow: inset 3px 3px 7px rgba(120, 110, 98, 0.2), inset -3px -3px 7px rgba(255, 255, 255, 0.72);
  color: #8a7e70;
  transition: color 0.25s ease;
}
.bc-clay-card:hover .bc-clay-icon {
  color: #b45309;
}
.bc-clay-text {
  display: flex;
  min-width: 0;
  flex: 1;
  flex-direction: column;
  gap: 2px;
}
.bc-clay-text small {
  font-size: 8px;
  letter-spacing: 0.08em;
  color: rgba(74, 69, 64, 0.42);
}
.bc-clay-text b {
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  font-size: 12px;
  font-weight: 600;
  letter-spacing: 0.01em;
  color: #4a4540;
}
.bc-clay-press {
  flex: none;
  border-radius: 999px;
  background: #dcd4ca;
  padding: 4px 8px;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 5.5px;
  letter-spacing: 0.08em;
  color: rgba(74, 69, 64, 0.5);
  transition: background 0.25s ease, color 0.25s ease;
}
.bc-clay-card:hover .bc-clay-press {
  background: #4a4540;
  color: #efe9e1;
}
.bc-clay-card-off {
  opacity: 0.55;
}
.bc-clay-foot {
  display: flex;
  align-items: center;
  gap: 12px;
  margin-top: auto;
  padding-top: 28px;
  animation: bc-clay-fade 1s 0.4s both;
}
.bc-clay-pill {
  display: flex;
  width: 34px;
  height: 34px;
  align-items: center;
  justify-content: center;
  border-radius: 50%;
  background: #eae4dc;
  box-shadow: 4px 4px 9px rgba(120, 110, 98, 0.2), -3px -3px 8px rgba(255, 255, 255, 0.7);
  font-size: 12px;
  color: #7a7167;
}
.bc-clay-foot p {
  flex: 1;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6px;
  line-height: 1.7;
  letter-spacing: 0.12em;
  color: rgba(74, 69, 64, 0.4);
}
.bc-clay-year {
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 7px;
  color: rgba(74, 69, 64, 0.35);
}
@keyframes bc-clay-in {
  from { opacity: 0; transform: translateY(18px); }
  to { opacity: 1; transform: none; }
}
@keyframes bc-clay-fade {
  from { opacity: 0; }
  to { opacity: 1; }
}
@media (prefers-reduced-motion: reduce) {
  .bc-clay-root *,
  .bc-clay-root *::before,
  .bc-clay-root *::after {
    animation: none !important;
    transition: none !important;
  }
}
</style>
