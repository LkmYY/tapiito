<template>
  <div class="cu-mp-root min-h-screen">
    <div class="cu-mp-wrap mx-auto w-full max-w-md">
      <header class="cu-mp-top cu-mp-in">
        <span class="cu-mp-logo">{{ initial }}<i>/</i></span>
        <span class="cu-mp-issue">MANIFESTO — 2026</span>
      </header>

      <section class="cu-mp-hero">
        <p class="cu-mp-kicker cu-mp-in cu-mp-d1">MAKE IT LOUD<br />MAKE IT USEFUL</p>
        <h1 class="cu-mp-name cu-mp-in cu-mp-d1">{{ payload?.companyName || '企业主页' }}</h1>
        <div class="cu-mp-badge cu-mp-in cu-mp-d2">
          <span>SINCE</span>
          <strong>26</strong>
          <span>STUDIO</span>
        </div>
      </section>

      <p class="cu-mp-desc cu-mp-in cu-mp-d2">{{ payload?.description }}</p>

      <div class="cu-mp-marquee cu-mp-in cu-mp-d3" aria-hidden="true">
        <div class="cu-mp-track">
          <span v-for="n in 2" :key="n">SERVICE ✳ DESIGN ✳ DELIVERY ✳ SERVICE ✳ DESIGN ✳ DELIVERY ✳</span>
        </div>
      </div>

      <section class="cu-mp-list cu-mp-in cu-mp-d3">
        <div class="cu-mp-list-head">
          <span>WHAT WE MAKE</span>
          <span>{{ String(services.length).padStart(2, '0') }}</span>
        </div>
        <div v-if="services.length" class="cu-mp-rows">
          <div v-for="(service, i) in services" :key="i" class="cu-mp-row">
            <span class="cu-mp-row-no">{{ String(i + 1).padStart(2, '0') }}</span>
            <h2>{{ service.name }}</h2>
            <Plus :size="18" class="cu-mp-row-plus" />
          </div>
        </div>
      </section>

      <section class="cu-mp-contact cu-mp-in cu-mp-d4">
        <div>
          <p class="cu-mp-contact-label">LET'S TALK / 直接联系</p>
          <a v-if="payload?.phone" :href="'tel:' + payload.phone" class="cu-mp-phone">
            <Phone :size="17" />
            {{ payload.phone }}
          </a>
        </div>
        <ArrowUpRight :size="30" class="cu-mp-arrow" />
      </section>

      <footer class="cu-mp-foot cu-mp-in cu-mp-d4">
        <span>{{ payload?.companyName || '企业主页' }}</span>
        <span>{{ nfcId ? 'NFC ' + nfcId : 'NFC —' }}</span>
      </footer>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { ArrowUpRight, Phone, Plus } from 'lucide-vue-next'

const props = defineProps(['payload', 'nfcId'])

const services = computed(() => {
  const list = props.payload?.services
  return Array.isArray(list) ? list.filter((item) => item && item.name) : []
})
const initial = computed(() => (props.payload?.companyName || '未').trim().slice(0, 1))
</script>

<style scoped>
.cu-mp-root {
  background: #1b2df0;
  color: #f7f6ef;
  font-family: 'Helvetica Neue', 'Arial Black', 'PingFang SC', sans-serif;
  overflow-x: hidden;
}
.cu-mp-wrap {
  padding: 18px 18px 30px;
}
.cu-mp-top {
  display: flex;
  align-items: center;
  justify-content: space-between;
  border-bottom: 1px solid rgba(247, 246, 239, 0.4);
  padding-bottom: 10px;
}
.cu-mp-logo {
  height: 30px;
  width: 30px;
  border: 1.5px solid #d8ff36;
  border-radius: 50%;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  font-family: Georgia, serif;
  font-size: 14px;
  color: #d8ff36;
}
.cu-mp-logo i {
  font-style: normal;
  font-size: 10px;
  margin-left: 1px;
}
.cu-mp-issue {
  font-family: 'Courier New', monospace;
  font-size: 8px;
  letter-spacing: 0.2em;
  color: #bcc4ff;
}
.cu-mp-hero {
  position: relative;
  padding-top: 26px;
}
.cu-mp-kicker {
  font-family: 'Courier New', monospace;
  font-size: 9px;
  line-height: 1.5;
  letter-spacing: 0.2em;
  color: #d8ff36;
}
.cu-mp-name {
  margin-top: 12px;
  max-width: 5.2em;
  font-size: 56px;
  line-height: 0.92;
  font-weight: 900;
  letter-spacing: -0.06em;
  color: #f7f6ef;
  word-break: break-word;
}
.cu-mp-badge {
  position: absolute;
  right: 0;
  top: 30px;
  height: 71px;
  width: 71px;
  border-radius: 50%;
  background: #d8ff36;
  color: #1b2df0;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  transform: rotate(11deg);
}
.cu-mp-badge span {
  font-family: 'Courier New', monospace;
  font-size: 6px;
  letter-spacing: 0.16em;
}
.cu-mp-badge strong {
  font-size: 25px;
  line-height: 0.9;
  font-weight: 900;
}
.cu-mp-desc {
  margin-top: 24px;
  max-width: 34ch;
  font-family: 'PingFang SC', 'Microsoft YaHei', sans-serif;
  font-size: 13px;
  font-weight: 400;
  line-height: 1.75;
  color: #c6cbff;
}
.cu-mp-marquee {
  margin: 24px -18px 0;
  overflow: hidden;
  background: #d8ff36;
  color: #1b2df0;
  padding: 7px 0;
}
.cu-mp-track {
  display: flex;
  width: max-content;
  animation: cu-mp-scroll 16s linear infinite;
}
.cu-mp-track span {
  white-space: nowrap;
  font-family: 'Courier New', monospace;
  font-size: 9px;
  font-weight: 700;
  letter-spacing: 0.16em;
}
.cu-mp-list {
  margin-top: 26px;
}
.cu-mp-list-head {
  display: flex;
  justify-content: space-between;
  font-family: 'Courier New', monospace;
  font-size: 8px;
  letter-spacing: 0.22em;
  color: #bcc4ff;
}
.cu-mp-rows {
  margin-top: 8px;
  border-top: 1px solid rgba(247, 246, 239, 0.5);
}
.cu-mp-row {
  display: flex;
  align-items: center;
  gap: 12px;
  border-bottom: 1px solid rgba(247, 246, 239, 0.5);
  padding: 15px 0;
  transition: color 0.2s ease, padding-left 0.2s ease;
}
.cu-mp-row:hover {
  color: #d8ff36;
  padding-left: 5px;
}
.cu-mp-row-no {
  font-family: 'Courier New', monospace;
  font-size: 8px;
  color: #d8ff36;
}
.cu-mp-row h2 {
  flex: 1;
  font-size: 25px;
  font-weight: 800;
  letter-spacing: -0.035em;
  line-height: 1.1;
}
.cu-mp-row-plus {
  flex-shrink: 0;
}
.cu-mp-empty {
  margin-top: 10px;
  border: 1px dashed rgba(247, 246, 239, 0.4);
  padding: 14px;
  font-family: 'PingFang SC', sans-serif;
  font-size: 11px;
  color: #c6cbff;
}
.cu-mp-contact {
  margin-top: 28px;
  background: #d8ff36;
  color: #1b2df0;
  padding: 17px 16px;
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 10px;
}
.cu-mp-contact-label {
  font-family: 'Courier New', monospace;
  font-size: 7px;
  letter-spacing: 0.18em;
}
.cu-mp-phone {
  margin-top: 6px;
  display: inline-flex;
  align-items: center;
  gap: 7px;
  font-size: 22px;
  font-weight: 800;
  letter-spacing: -0.03em;
  color: #1b2df0;
  transition: opacity 0.2s ease;
}
.cu-mp-phone:hover {
  opacity: 0.72;
}
.cu-mp-contact .cu-mp-empty {
  border-color: rgba(27, 45, 240, 0.3);
  color: #1b2df0;
}
.cu-mp-arrow {
  flex-shrink: 0;
}
.cu-mp-foot {
  margin-top: 18px;
  display: flex;
  justify-content: space-between;
  font-family: 'Courier New', monospace;
  font-size: 7px;
  letter-spacing: 0.14em;
  color: #bcc4ff;
}
@keyframes cu-mp-scroll {
  from {
    transform: translateX(0);
  }
  to {
    transform: translateX(-50%);
  }
}
@keyframes cu-mp-in {
  from {
    opacity: 0;
    transform: translateY(18px);
  }
  to {
    opacity: 1;
    transform: none;
  }
}
.cu-mp-in {
  animation: cu-mp-in 0.6s cubic-bezier(0.16, 1, 0.3, 1) both;
}
.cu-mp-d1 {
  animation-delay: 0.07s;
}
.cu-mp-d2 {
  animation-delay: 0.15s;
}
.cu-mp-d3 {
  animation-delay: 0.24s;
}
.cu-mp-d4 {
  animation-delay: 0.33s;
}
@media (prefers-reduced-motion: reduce) {
  .cu-mp-in {
    animation: none;
  }
  .cu-mp-track {
    animation: none;
  }
  .cu-mp-row,
  .cu-mp-phone {
    transition: none;
  }
}
</style>
