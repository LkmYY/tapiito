<template>
  <div class="bc-letter-root min-h-screen px-4 py-8 text-[#3b2417]">
    <div class="bc-letter-stage mx-auto w-full max-w-md">
      <div class="bc-letter-ticket">
        <div class="bc-letter-topline">
          <span>YUNTU LETTERPRESS CO.</span>
          <span>EST. 1987</span>
        </div>

        <div class="bc-letter-body">
          <div class="bc-letter-main">
            <p class="bc-letter-admit">ADMIT ONE</p>
            <p class="bc-letter-serial">SERIAL {{ nfcId || 'NO. 0001' }}</p>

            <h1 class="bc-letter-name">{{ payload.name || '未命名' }}</h1>
            <p class="bc-letter-role">{{ payload.title }}</p>
            <p class="bc-letter-company">{{ payload.company }}</p>

            <div class="bc-letter-dots"></div>

            <div class="bc-letter-rows">
              <div class="bc-letter-row">
                <span>PHONE</span>
                <a v-if="payload.phone" :href="'tel:' + payload.phone">{{ payload.phone }}</a>
              </div>
              <div class="bc-letter-row">
                <span>EMAIL</span>
                <a v-if="payload.email" :href="'mailto:' + payload.email">{{ payload.email }}</a>
              </div>
              <div class="bc-letter-row">
                <span>WEB</span>
                <a v-if="payload.website" :href="payload.website" target="_blank" rel="noopener noreferrer">{{ payload.website }}</a>
              </div>
            </div>
          </div>

          <div class="bc-letter-stub">
            <span class="bc-letter-stub-label">KEEP THIS STUB</span>
            <span class="bc-letter-stub-name">{{ initial }}</span>
            <span class="bc-letter-stub-role">{{ payload.title }}</span>
            <div class="bc-letter-barcode"></div>
            <span class="bc-letter-stub-num">{{ payload.wechat }}</span>
          </div>
        </div>

        <div class="bc-letter-bottom">
          <span>GOOD FOR ONE INTRODUCTION</span>
          <span>★ ★ ★</span>
        </div>
      </div>

      <p class="bc-letter-note">NO REFUNDS · GOOD CONVERSATION ONLY</p>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps(['payload', 'nfcId'])

const initial = computed(() => (props.payload?.name || '?').toString().trim()[0] || '?')
</script>

<style scoped>
.bc-letter-root {
  position: relative;
  overflow: hidden;
  background: #e9dcc3;
  font-family: 'Courier New', 'PingFang SC', 'Microsoft YaHei', monospace;
}
.bc-letter-root::before {
  content: '';
  position: absolute;
  inset: 0;
  background-image: radial-gradient(rgba(92, 51, 25, 0.24) 0.7px, transparent 0.7px);
  background-size: 4px 4px;
  opacity: 0.3;
  pointer-events: none;
}
.bc-letter-stage {
  position: relative;
  z-index: 1;
  animation: bc-letter-in 0.9s cubic-bezier(0.22, 1, 0.36, 1) both;
}
.bc-letter-ticket {
  position: relative;
  background: #f4e8d2;
  border: 1px solid #8a5a3b;
  box-shadow: 0 24px 50px -30px rgba(59, 36, 23, 0.6), inset 0 0 60px rgba(138, 90, 59, 0.08);
  transform: rotate(-1.2deg);
}
.bc-letter-ticket::before {
  content: '';
  position: absolute;
  left: 0;
  right: 0;
  top: 14px;
  height: 1px;
  background: repeating-linear-gradient(90deg, #8a5a3b 0 6px, transparent 6px 11px);
  opacity: 0.7;
}
.bc-letter-topline {
  display: flex;
  justify-content: space-between;
  padding: 27px 20px 12px;
  font-size: 7.5px;
  font-weight: 700;
  letter-spacing: 0.2em;
  color: #7c4a2d;
}
.bc-letter-body {
  display: grid;
  grid-template-columns: 1fr 96px;
  border-top: 3px double #8a5a3b;
  border-bottom: 3px double #8a5a3b;
}
.bc-letter-main {
  min-width: 0;
  padding: 20px 18px 22px;
  border-right: 1px dashed #a97a58;
}
.bc-letter-admit {
  display: inline-block;
  border: 1px solid #8a5a3b;
  padding: 3px 8px;
  font-size: 8px;
  font-weight: 700;
  letter-spacing: 0.24em;
  color: #8a5a3b;
  transform: rotate(-2deg);
}
.bc-letter-serial {
  margin-top: 10px;
  font-size: 7px;
  letter-spacing: 0.14em;
  color: #a97a58;
}
.bc-letter-name {
  margin-top: 20px;
  font-family: Georgia, 'Songti SC', serif;
  font-size: 36px;
  font-weight: 700;
  line-height: 1.02;
  letter-spacing: -0.04em;
  overflow-wrap: anywhere;
  color: #3b2417;
  text-shadow: 1px 1px 0 rgba(138, 90, 59, 0.28);
}
.bc-letter-role {
  margin-top: 10px;
  font-family: 'Courier New', monospace;
  font-size: 10px;
  font-weight: 700;
  letter-spacing: 0.16em;
  color: #9a3412;
}
.bc-letter-company {
  margin-top: 3px;
  font-size: 9px;
  letter-spacing: 0.12em;
  color: #7c4a2d;
}
.bc-letter-dots {
  height: 1px;
  margin: 20px 0 8px;
  background: repeating-linear-gradient(90deg, #b08862 0 3px, transparent 3px 7px);
}
.bc-letter-rows {
  display: flex;
  flex-direction: column;
}
.bc-letter-row {
  display: flex;
  align-items: baseline;
  justify-content: space-between;
  gap: 10px;
  padding: 7px 0;
  border-bottom: 1px dotted rgba(138, 90, 59, 0.4);
}
.bc-letter-row span {
  flex: none;
  font-size: 6.5px;
  letter-spacing: 0.16em;
  color: #a97a58;
}
.bc-letter-row a,
.bc-letter-row em {
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  font-size: 9.5px;
  font-style: normal;
  color: #3b2417;
  text-decoration: none;
  transition: color 0.2s ease;
}
.bc-letter-row a:hover {
  color: #b91c1c;
  text-decoration: underline;
}
.bc-letter-row a:active {
  opacity: 0.6;
}
.bc-letter-row em {
  color: #b9a58c;
}
.bc-letter-stub {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: space-between;
  padding: 18px 8px;
  background: rgba(138, 90, 59, 0.07);
}
.bc-letter-stub-label {
  font-size: 5.5px;
  letter-spacing: 0.1em;
  color: #a97a58;
  writing-mode: vertical-rl;
  height: 76px;
}
.bc-letter-stub-name {
  display: flex;
  width: 48px;
  height: 48px;
  align-items: center;
  justify-content: center;
  border: 2px solid #9a3412;
  border-radius: 50%;
  font-family: Georgia, serif;
  font-size: 24px;
  color: #9a3412;
  transition: transform 0.4s ease;
}
.bc-letter-stub:hover .bc-letter-stub-name {
  transform: rotate(-12deg) scale(1.06);
}
.bc-letter-stub-role {
  max-width: 70px;
  text-align: center;
  font-size: 6px;
  line-height: 1.4;
  letter-spacing: 0.04em;
  color: #7c4a2d;
}
.bc-letter-barcode {
  width: 54px;
  height: 30px;
  background: repeating-linear-gradient(90deg, #3b2417 0 1px, transparent 1px 3px, #3b2417 3px 5px, transparent 5px 8px);
}
.bc-letter-stub-num {
  max-width: 74px;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  font-size: 5.5px;
  letter-spacing: 0.04em;
  color: #7c4a2d;
}
.bc-letter-bottom {
  display: flex;
  justify-content: space-between;
  padding: 12px 20px 14px;
  font-size: 6.5px;
  font-weight: 700;
  letter-spacing: 0.16em;
  color: #8a5a3b;
}
.bc-letter-note {
  margin-top: 16px;
  text-align: center;
  font-size: 6.5px;
  letter-spacing: 0.2em;
  color: rgba(59, 36, 23, 0.4);
}
@keyframes bc-letter-in {
  from { opacity: 0; transform: translateY(26px) rotate(1.5deg); }
  to { opacity: 1; transform: rotate(-1.2deg); }
}
@media (prefers-reduced-motion: reduce) {
  .bc-letter-root *,
  .bc-letter-root *::before,
  .bc-letter-root *::after {
    animation: none !important;
    transition: none !important;
  }
  .bc-letter-ticket {
    transform: none;
  }
}
</style>
