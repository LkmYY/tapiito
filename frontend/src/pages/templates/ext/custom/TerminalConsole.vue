<template>
  <div class="cu-tc-root min-h-screen">
    <div class="cu-tc-scanlines" aria-hidden="true"></div>
    <div class="cu-tc-wrap mx-auto w-full max-w-md">
      <div class="cu-tc-window cu-tc-in">
        <div class="cu-tc-bar">
          <span class="cu-tc-dots"><i></i><i></i><i></i></span>
          <span class="cu-tc-path">~/company/profile</span>
          <span class="cu-tc-status">ONLINE</span>
        </div>

        <div class="cu-tc-body">
          <p class="cu-tc-prompt cu-tc-in cu-tc-d1">$ whoami<span class="cu-tc-cursor"></span></p>
          <h1 class="cu-tc-name cu-tc-in cu-tc-d1">{{ payload?.companyName || '企业主页' }}</h1>

          <p class="cu-tc-rule cu-tc-in cu-tc-d2">────────────────────────</p>
          <p class="cu-tc-line cu-tc-in cu-tc-d2"><span class="cu-tc-key">// intro</span> {{ payload?.description }}</p>

          <div class="cu-tc-services cu-tc-in cu-tc-d3">
            <p class="cu-tc-prompt">$ services --list</p>
            <div v-if="services.length" class="cu-tc-output">
              <div v-for="(service, i) in services" :key="i" class="cu-tc-service">
                <span class="cu-tc-bullet">▸</span>
                <span class="cu-tc-service-name">{{ service.name }}</span>
                <span class="cu-tc-status-ok">[ READY ]</span>
              </div>
            </div>
            <p v-else class="cu-tc-muted">no services configured</p>
          </div>

          <div class="cu-tc-contact cu-tc-in cu-tc-d4">
            <p class="cu-tc-prompt">$ contact --call</p>
            <a v-if="payload?.phone" :href="'tel:' + payload.phone" class="cu-tc-phone">
              <Phone :size="15" />
              {{ payload.phone }}
            </a>
            <p v-else class="cu-tc-muted">phone not configured</p>
            <a v-if="payload?.email" :href="'mailto:' + payload.email" class="cu-tc-mail">mailto:{{ payload.email }}</a>
          </div>

          <div class="cu-tc-foot cu-tc-in cu-tc-d4">
            <span>SESSION {{ nfcId || '0000' }}</span>
            <span>UTF-8 / 2026</span>
          </div>
        </div>
      </div>

      <p class="cu-tc-outside cu-tc-in cu-tc-d4">[ 连接稳定 · 点击电话可直接呼叫 ]</p>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { Phone } from 'lucide-vue-next'

const props = defineProps(['payload', 'nfcId'])

const services = computed(() => {
  const list = props.payload?.services
  return Array.isArray(list) ? list.filter((item) => item && item.name) : []
})
</script>

<style scoped>
.cu-tc-root {
  position: relative;
  background: #040705;
  color: #9dffa8;
  font-family: 'Courier New', 'Consolas', monospace;
  overflow-x: hidden;
}
.cu-tc-scanlines {
  position: fixed;
  inset: 0;
  pointer-events: none;
  background: repeating-linear-gradient(0deg, rgba(125, 255, 156, 0.045) 0 1px, transparent 1px 3px);
  z-index: 5;
}
.cu-tc-wrap {
  position: relative;
  padding: 16px 12px 28px;
}
.cu-tc-window {
  border: 1px solid #2d6a3a;
  background: #071009;
  box-shadow: 0 0 36px rgba(85, 255, 122, 0.1);
}
.cu-tc-bar {
  display: flex;
  align-items: center;
  gap: 8px;
  border-bottom: 1px solid #234a2b;
  background: #0c1a0f;
  padding: 8px 10px;
  font-size: 7px;
  letter-spacing: 0.08em;
}
.cu-tc-dots {
  display: flex;
  gap: 4px;
}
.cu-tc-dots i {
  height: 6px;
  width: 6px;
  border-radius: 50%;
  background: #2d6a3a;
}
.cu-tc-dots i:first-child {
  background: #ff5c5c;
}
.cu-tc-dots i:nth-child(2) {
  background: #ffd166;
}
.cu-tc-dots i:nth-child(3) {
  background: #60ff8f;
}
.cu-tc-path {
  flex: 1;
  color: #6fae7c;
}
.cu-tc-status {
  color: #60ff8f;
}
.cu-tc-body {
  padding: 20px 16px 16px;
}
.cu-tc-prompt {
  font-size: 11px;
  color: #6fae7c;
  letter-spacing: 0.02em;
}
.cu-tc-cursor {
  display: inline-block;
  margin-left: 3px;
  height: 11px;
  width: 6px;
  background: #7dff9c;
  vertical-align: -1px;
  animation: cu-tc-blink 1.1s steps(2, end) infinite;
}
.cu-tc-name {
  margin-top: 9px;
  font-family: 'Courier New', 'PingFang SC', monospace;
  font-size: 37px;
  line-height: 1.02;
  font-weight: 700;
  letter-spacing: -0.05em;
  color: #dcffe2;
  word-break: break-word;
  text-shadow: 0 0 18px rgba(125, 255, 156, 0.24);
}
.cu-tc-rule {
  margin-top: 17px;
  font-size: 9px;
  color: #2d6a3a;
  overflow: hidden;
  white-space: nowrap;
}
.cu-tc-line {
  margin-top: 11px;
  font-family: 'PingFang SC', 'Microsoft YaHei', monospace;
  font-size: 12px;
  line-height: 1.85;
  color: #9acfa3;
}
.cu-tc-key {
  display: inline-block;
  margin-bottom: 3px;
  color: #4f9660;
  font-family: 'Courier New', monospace;
  font-size: 9px;
}
.cu-tc-services {
  margin-top: 24px;
}
.cu-tc-output {
  margin-top: 9px;
}
.cu-tc-service {
  display: flex;
  align-items: center;
  gap: 8px;
  border-bottom: 1px dashed #1d3a24;
  padding: 10px 0;
  transition: background-color 0.2s ease;
}
.cu-tc-service:hover {
  background: rgba(125, 255, 156, 0.05);
}
.cu-tc-bullet {
  color: #60ff8f;
  font-size: 10px;
}
.cu-tc-service-name {
  flex: 1;
  font-family: 'PingFang SC', sans-serif;
  font-size: 13px;
  color: #c8f7d1;
}
.cu-tc-status-ok {
  font-size: 7px;
  letter-spacing: 0.08em;
  color: #4f9660;
}
.cu-tc-contact {
  margin-top: 24px;
}
.cu-tc-phone {
  margin-top: 9px;
  display: inline-flex;
  align-items: center;
  gap: 8px;
  border: 1px solid #3f8c50;
  background: rgba(125, 255, 156, 0.06);
  padding: 10px 13px;
  font-size: 18px;
  letter-spacing: 0.03em;
  color: #e2ffe8;
  transition: background-color 0.2s ease, border-color 0.2s ease;
}
.cu-tc-phone:hover {
  background: rgba(125, 255, 156, 0.16);
  border-color: #7dff9c;
}
.cu-tc-mail {
  display: block;
  margin-top: 9px;
  font-size: 9px;
  color: #6fae7c;
  word-break: break-all;
}
.cu-tc-muted {
  margin-top: 8px;
  font-size: 10px;
  color: #4f7a58;
}
.cu-tc-foot {
  margin-top: 24px;
  display: flex;
  justify-content: space-between;
  border-top: 1px solid #234a2b;
  padding-top: 8px;
  font-size: 7px;
  letter-spacing: 0.1em;
  color: #4f7a58;
}
.cu-tc-outside {
  margin-top: 12px;
  text-align: center;
  font-size: 8px;
  letter-spacing: 0.12em;
  color: #406b49;
}
@keyframes cu-tc-blink {
  50% {
    opacity: 0;
  }
}
@keyframes cu-tc-in {
  from {
    opacity: 0;
    transform: translateY(8px);
  }
  to {
    opacity: 1;
    transform: none;
  }
}
.cu-tc-in {
  animation: cu-tc-in 0.55s steps(4, end) both;
}
.cu-tc-d1 {
  animation-delay: 0.12s;
}
.cu-tc-d2 {
  animation-delay: 0.22s;
}
.cu-tc-d3 {
  animation-delay: 0.34s;
}
.cu-tc-d4 {
  animation-delay: 0.46s;
}
@media (prefers-reduced-motion: reduce) {
  .cu-tc-in {
    animation: none;
  }
  .cu-tc-cursor {
    animation: none;
  }
  .cu-tc-service,
  .cu-tc-phone {
    transition: none;
  }
  .cu-tc-scanlines {
    display: none;
  }
}
</style>
