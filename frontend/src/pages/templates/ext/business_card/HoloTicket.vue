<template>
  <div class="bc-holo-root min-h-screen px-4 py-8">
    <div class="bc-holo-bg"></div>
    <div class="bc-holo-wrap mx-auto flex min-h-[calc(100vh-4rem)] w-full max-w-md flex-col">
      <header class="bc-holo-head">
        <span>HOLOGRAPHIC ACCESS</span>
        <span>{{ nfcId ? 'SERIES ' + nfcId : 'VIP / 01' }}</span>
      </header>

      <div class="bc-holo-ticket">
        <div class="bc-holo-sheen"></div>
        <div class="bc-holo-notch bc-holo-notch-l"></div>
        <div class="bc-holo-notch bc-holo-notch-r"></div>

        <div class="bc-holo-top">
          <div>
            <p class="bc-holo-brand">YUNTU / LIVE</p>
            <p class="bc-holo-sub">ONE NIGHT · ONE CONNECTION</p>
          </div>
          <span class="bc-holo-admit">ADMIT<br />ONE</span>
        </div>

        <div class="bc-holo-hero">
          <span class="bc-holo-chip">{{ initial }}</span>
          <p class="bc-holo-kicker">HEADLINER</p>
          <h1 class="bc-holo-name">{{ payload.name || '未命名' }}</h1>
          <p class="bc-holo-role">{{ payload.title }} · {{ payload.company }}</p>
        </div>

        <div class="bc-holo-tear">
          <span>✂</span>
          <i></i>
          <span>TEAR HERE</span>
          <i></i>
          <span>✂</span>
        </div>

        <div class="bc-holo-details">
          <div class="bc-holo-detail">
            <span>SECTION</span>
            <p>TEL</p>
          </div>
          <div class="bc-holo-detail bc-holo-detail-wide">
            <span>ENTRY</span>
            <a v-if="payload.phone" :href="'tel:' + payload.phone">{{ payload.phone }}</a>
          </div>
          <div class="bc-holo-detail">
            <span>ROW</span>
            <p>MAIL</p>
          </div>
          <div class="bc-holo-detail bc-holo-detail-wide">
            <span>SEAT</span>
            <a v-if="payload.email" :href="'mailto:' + payload.email">{{ payload.email }}</a>
          </div>
          <div class="bc-holo-detail">
            <span>GATE</span>
            <p>WEB</p>
          </div>
          <div class="bc-holo-detail bc-holo-detail-wide">
            <span>ACCESS</span>
            <a v-if="payload.website" :href="payload.website" target="_blank" rel="noopener noreferrer">{{ payload.website }}</a>
          </div>
          <div class="bc-holo-detail">
            <span>BACKSTAGE</span>
            <p>WECHAT</p>
          </div>
          <div class="bc-holo-detail bc-holo-detail-wide">
            <span>PASS</span>
            <p :class="{ 'bc-holo-empty': !payload.wechat }">{{ payload.wechat }}</p>
          </div>
        </div>

        <div class="bc-holo-bottom">
          <div class="bc-holo-barcode"></div>
          <div class="bc-holo-serial">
            <span>NO. {{ nfcId || '000000' }}</span>
            <span>NON-TRANSFERABLE</span>
          </div>
        </div>
      </div>

      <p class="bc-holo-foot">HOLD NEAR READER · KEEP THIS MOMENT</p>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps(['payload', 'nfcId'])

const initial = computed(() => (props.payload?.name || '?').toString().trim()[0] || '?')
</script>

<style scoped>
.bc-holo-root {
  position: relative;
  overflow: hidden;
  background: #12081f;
  font-family: 'Helvetica Neue', 'PingFang SC', 'Microsoft YaHei', sans-serif;
  color: #f5f3ff;
}
.bc-holo-bg {
  position: absolute;
  inset: 0;
  background:
    radial-gradient(circle at 14% 8%, rgba(192, 132, 252, 0.5), transparent 38%),
    radial-gradient(circle at 88% 20%, rgba(34, 211, 238, 0.38), transparent 35%),
    radial-gradient(circle at 50% 100%, rgba(244, 114, 182, 0.32), transparent 48%);
  filter: blur(8px);
}
.bc-holo-wrap {
  position: relative;
  z-index: 1;
}
.bc-holo-head {
  display: flex;
  justify-content: space-between;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6.5px;
  letter-spacing: 0.22em;
  color: rgba(245, 243, 255, 0.55);
  animation: bc-holo-fade 0.8s both;
}
.bc-holo-ticket {
  position: relative;
  margin-top: 18px;
  border: 1px solid rgba(255, 255, 255, 0.5);
  border-radius: 18px;
  background:
    conic-gradient(from 210deg at 60% 20%, rgba(192, 132, 252, 0.4), rgba(34, 211, 238, 0.32), rgba(251, 207, 232, 0.36), rgba(129, 140, 248, 0.3), rgba(192, 132, 252, 0.4)),
    rgba(20, 10, 35, 0.9);
  box-shadow: 0 30px 80px -30px rgba(192, 132, 252, 0.5), inset 0 1px 0 rgba(255, 255, 255, 0.4);
  overflow: hidden;
  backdrop-filter: blur(18px);
  animation: bc-holo-in 0.95s cubic-bezier(0.22, 1, 0.36, 1) both;
  transition: transform 0.45s cubic-bezier(0.22, 1, 0.36, 1), box-shadow 0.45s ease;
}
.bc-holo-ticket:hover {
  transform: perspective(900px) rotateX(1.5deg) rotateY(-2deg) translateY(-3px);
  box-shadow: 0 38px 90px -32px rgba(192, 132, 252, 0.65), inset 0 1px 0 rgba(255, 255, 255, 0.5);
}
.bc-holo-sheen {
  position: absolute;
  top: -40%;
  left: -70%;
  width: 70%;
  height: 180%;
  background: linear-gradient(105deg, transparent, rgba(255, 255, 255, 0.36), transparent);
  transform: rotate(12deg);
  animation: bc-holo-shine 6s ease-in-out infinite;
  pointer-events: none;
}
.bc-holo-notch {
  position: absolute;
  top: 62%;
  width: 22px;
  height: 22px;
  border-radius: 50%;
  background: #12081f;
  box-shadow: inset 0 0 0 1px rgba(255, 255, 255, 0.25);
}
.bc-holo-notch-l { left: -12px; }
.bc-holo-notch-r { right: -12px; }
.bc-holo-top {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  padding: 20px 20px 0;
}
.bc-holo-brand {
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 9px;
  font-weight: 700;
  letter-spacing: 0.16em;
}
.bc-holo-sub {
  margin-top: 4px;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 5.5px;
  letter-spacing: 0.18em;
  color: rgba(245, 243, 255, 0.5);
}
.bc-holo-admit {
  border: 1px solid rgba(255, 255, 255, 0.55);
  padding: 4px 7px;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6px;
  line-height: 1.3;
  letter-spacing: 0.14em;
  text-align: center;
  transform: rotate(3deg);
}
.bc-holo-hero {
  position: relative;
  padding: 30px 20px 24px;
}
.bc-holo-chip {
  display: flex;
  width: 42px;
  height: 32px;
  align-items: center;
  justify-content: center;
  border: 1px solid rgba(255, 255, 255, 0.45);
  border-radius: 6px;
  background:
    linear-gradient(90deg, transparent 47%, rgba(255, 255, 255, 0.4) 48% 52%, transparent 53%),
    linear-gradient(0deg, transparent 47%, rgba(255, 255, 255, 0.4) 48% 52%, transparent 53%),
    rgba(255, 255, 255, 0.12);
  font-family: Georgia, serif;
  font-size: 15px;
}
.bc-holo-kicker {
  margin-top: 20px;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6px;
  letter-spacing: 0.3em;
  color: rgba(245, 243, 255, 0.55);
}
.bc-holo-name {
  margin-top: 5px;
  font-size: clamp(36px, 12vw, 48px);
  font-weight: 200;
  line-height: 1.02;
  letter-spacing: -0.01em;
  text-transform: uppercase;
  overflow-wrap: anywhere;
  text-shadow: 0 2px 20px rgba(255, 255, 255, 0.3);
}
.bc-holo-role {
  margin-top: 10px;
  font-size: 9px;
  letter-spacing: 0.08em;
  color: rgba(245, 243, 255, 0.65);
}
.bc-holo-tear {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 0 18px;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 5px;
  letter-spacing: 0.14em;
  color: rgba(245, 243, 255, 0.45);
}
.bc-holo-tear i {
  flex: 1;
  height: 1px;
  background: repeating-linear-gradient(90deg, rgba(245, 243, 255, 0.4) 0 4px, transparent 4px 8px);
}
.bc-holo-details {
  display: grid;
  grid-template-columns: 58px 1fr 58px 1fr;
  gap: 0 10px;
  padding: 16px 20px 12px;
}
.bc-holo-detail {
  padding: 6px 0;
  border-bottom: 1px solid rgba(255, 255, 255, 0.14);
}
.bc-holo-detail span {
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 5px;
  letter-spacing: 0.16em;
  color: rgba(245, 243, 255, 0.4);
}
.bc-holo-detail p {
  font-size: 7px;
  letter-spacing: 0.08em;
  color: rgba(245, 243, 255, 0.75);
}
.bc-holo-detail a,
.bc-holo-detail em {
  display: block;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  font-size: 8.5px;
  font-style: normal;
  color: #f5f3ff;
  text-decoration: none;
  transition: color 0.2s ease;
}
.bc-holo-detail a:hover {
  color: #67e8f9;
}
.bc-holo-detail a:active {
  opacity: 0.6;
}
.bc-holo-detail em,
.bc-holo-empty {
  color: rgba(245, 243, 255, 0.3);
}
.bc-holo-bottom {
  display: flex;
  align-items: flex-end;
  gap: 14px;
  padding: 10px 20px 18px;
}
.bc-holo-barcode {
  width: 108px;
  height: 34px;
  background: repeating-linear-gradient(90deg, #f5f3ff 0 1px, transparent 1px 3px, #f5f3ff 3px 5px, transparent 5px 7px, #f5f3ff 7px 10px, transparent 10px 12px);
  opacity: 0.85;
}
.bc-holo-serial {
  display: flex;
  flex: 1;
  flex-direction: column;
  gap: 3px;
  text-align: right;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 5.5px;
  letter-spacing: 0.12em;
  color: rgba(245, 243, 255, 0.5);
}
.bc-holo-foot {
  margin-top: auto;
  padding-top: 18px;
  text-align: center;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6px;
  letter-spacing: 0.2em;
  color: rgba(245, 243, 255, 0.35);
}
@keyframes bc-holo-in {
  from { opacity: 0; transform: translateY(24px) rotateX(6deg); }
  to { opacity: 1; transform: none; }
}
@keyframes bc-holo-fade {
  from { opacity: 0; }
  to { opacity: 1; }
}
@keyframes bc-holo-shine {
  0%, 55% { left: -70%; }
  85%, 100% { left: 140%; }
}
@media (prefers-reduced-motion: reduce) {
  .bc-holo-root *,
  .bc-holo-root *::before,
  .bc-holo-root *::after {
    animation: none !important;
    transition: none !important;
  }
}
</style>
