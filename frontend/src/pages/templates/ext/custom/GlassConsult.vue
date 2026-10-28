<template>
  <div class="cu-gc-root min-h-screen">
    <div class="cu-gc-orb cu-gc-orb--a" aria-hidden="true"></div>
    <div class="cu-gc-orb cu-gc-orb--b" aria-hidden="true"></div>
    <div class="cu-gc-orb cu-gc-orb--c" aria-hidden="true"></div>

    <div class="cu-gc-wrap mx-auto w-full max-w-md">
      <header class="cu-gc-top cu-gc-in">
        <span class="cu-gc-logo">{{ initial }}</span>
        <span class="cu-gc-label">CONSULTANCY / 2026</span>
      </header>

      <section class="cu-gc-card cu-gc-in cu-gc-d1">
        <p class="cu-gc-eyebrow">STRATEGY &amp; GROWTH</p>
        <h1 class="cu-gc-name">{{ payload?.companyName || '企业主页' }}</h1>
        <p class="cu-gc-desc">{{ payload?.description }}</p>

        <div class="cu-gc-pills">
          <span>品牌策略</span>
          <span>体验设计</span>
          <span>增长咨询</span>
        </div>
      </section>

      <div class="cu-gc-stats cu-gc-in cu-gc-d2">
        <div class="cu-gc-stat">
          <strong>{{ String(services.length).padStart(2, '0') }}</strong>
          <span>核心服务</span>
        </div>
        <div class="cu-gc-stat">
          <strong>1:1</strong>
          <span>专属顾问</span>
        </div>
        <div class="cu-gc-stat">
          <strong>24h</strong>
          <span>响应时间</span>
        </div>
      </div>

      <section class="cu-gc-services cu-gc-in cu-gc-d3">
        <div class="cu-gc-section-head">
          <h2>Tapiito能做什么</h2>
          <span>SERVICES</span>
        </div>
        <div v-if="services.length" class="cu-gc-grid">
          <article v-for="(service, i) in services" :key="i" class="cu-gc-service">
            <span class="cu-gc-service-no">0{{ i + 1 }}</span>
            <span class="cu-gc-service-dot"></span>
            <h3>{{ service.name }}</h3>
          </article>
        </div>
      </section>

      <section class="cu-gc-contact cu-gc-in cu-gc-d4">
        <p class="cu-gc-contact-label">预约一次 30 分钟沟通</p>
        <a v-if="payload?.phone" :href="'tel:' + payload.phone" class="cu-gc-phone">
          <Phone :size="16" />
          <span>{{ payload.phone }}</span>
          <ArrowUpRight :size="15" />
        </a>
        <a v-if="payload?.email" :href="'mailto:' + payload.email" class="cu-gc-mail">{{ payload.email }}</a>
      </section>

      <footer class="cu-gc-foot cu-gc-in cu-gc-d4">
        <span>{{ payload?.companyName || '企业主页' }}</span>
        <span>{{ nfcId ? 'NFC ' + nfcId : 'NFC —' }}</span>
      </footer>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { ArrowUpRight, Phone } from 'lucide-vue-next'

const props = defineProps(['payload', 'nfcId'])

const services = computed(() => {
  const list = props.payload?.services
  return Array.isArray(list) ? list.filter((item) => item && item.name) : []
})
const initial = computed(() => (props.payload?.companyName || '未').trim().slice(0, 1))
</script>

<style scoped>
.cu-gc-root {
  position: relative;
  overflow: hidden;
  background: linear-gradient(160deg, #e9e4ff 0%, #d4dcff 48%, #f2e8f7 100%);
  color: #241f3d;
  font-family: 'PingFang SC', 'Microsoft YaHei', sans-serif;
}
.cu-gc-orb {
  position: absolute;
  border-radius: 50%;
  filter: blur(2px);
  pointer-events: none;
}
.cu-gc-orb--a {
  top: -70px;
  right: -50px;
  height: 220px;
  width: 220px;
  background: radial-gradient(circle at 35% 35%, rgba(255, 255, 255, 0.9), rgba(167, 139, 250, 0.45));
}
.cu-gc-orb--b {
  top: 300px;
  left: -80px;
  height: 190px;
  width: 190px;
  background: radial-gradient(circle at 40% 40%, rgba(186, 230, 253, 0.9), rgba(96, 165, 250, 0.3));
}
.cu-gc-orb--c {
  bottom: 40px;
  right: -40px;
  height: 160px;
  width: 160px;
  background: radial-gradient(circle at 40% 40%, rgba(251, 207, 232, 0.85), rgba(244, 114, 182, 0.25));
}
.cu-gc-wrap {
  position: relative;
  z-index: 2;
  padding: 24px 18px 32px;
}
.cu-gc-top {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 0 4px;
}
.cu-gc-logo {
  height: 31px;
  width: 31px;
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 10px;
  background: rgba(255, 255, 255, 0.55);
  border: 1px solid rgba(255, 255, 255, 0.8);
  box-shadow: 0 8px 18px -12px rgba(76, 29, 149, 0.55);
  font-size: 14px;
  font-weight: 600;
  color: #5b21b6;
}
.cu-gc-label {
  font-family: 'Courier New', monospace;
  font-size: 8px;
  letter-spacing: 0.2em;
  color: #6d6390;
}
.cu-gc-card {
  margin-top: 22px;
  border-radius: 26px;
  border: 1px solid rgba(255, 255, 255, 0.85);
  background: rgba(255, 255, 255, 0.42);
  backdrop-filter: blur(22px);
  -webkit-backdrop-filter: blur(22px);
  box-shadow: 0 28px 60px -34px rgba(67, 42, 126, 0.55);
  padding: 24px 22px;
}
.cu-gc-eyebrow {
  font-family: 'Courier New', monospace;
  font-size: 8px;
  letter-spacing: 0.24em;
  color: #7c3aed;
}
.cu-gc-name {
  margin-top: 11px;
  font-size: 38px;
  line-height: 1.1;
  font-weight: 700;
  letter-spacing: -0.035em;
  color: #211a39;
  word-break: break-word;
}
.cu-gc-desc {
  margin-top: 12px;
  font-size: 12px;
  line-height: 1.85;
  color: #5e5578;
}
.cu-gc-pills {
  display: flex;
  flex-wrap: wrap;
  gap: 6px;
  margin-top: 16px;
}
.cu-gc-pills span {
  border-radius: 999px;
  background: rgba(255, 255, 255, 0.6);
  border: 1px solid rgba(255, 255, 255, 0.9);
  padding: 5px 10px;
  font-size: 9px;
  letter-spacing: 0.06em;
  color: #5b4b86;
}
.cu-gc-stats {
  margin-top: 12px;
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 8px;
}
.cu-gc-stat {
  border-radius: 16px;
  border: 1px solid rgba(255, 255, 255, 0.8);
  background: rgba(255, 255, 255, 0.35);
  backdrop-filter: blur(12px);
  -webkit-backdrop-filter: blur(12px);
  padding: 11px 8px;
  text-align: center;
}
.cu-gc-stat strong {
  display: block;
  font-size: 19px;
  font-weight: 700;
  color: #4c1d95;
  letter-spacing: -0.02em;
}
.cu-gc-stat span {
  display: block;
  margin-top: 2px;
  font-size: 8px;
  letter-spacing: 0.08em;
  color: #6d6390;
}
.cu-gc-services {
  margin-top: 24px;
  border-radius: 24px;
  border: 1px solid rgba(255, 255, 255, 0.75);
  background: rgba(255, 255, 255, 0.3);
  backdrop-filter: blur(16px);
  -webkit-backdrop-filter: blur(16px);
  padding: 18px 16px;
}
.cu-gc-section-head {
  display: flex;
  align-items: baseline;
  justify-content: space-between;
}
.cu-gc-section-head h2 {
  font-size: 14px;
  font-weight: 600;
  letter-spacing: 0.04em;
}
.cu-gc-section-head span {
  font-family: 'Courier New', monospace;
  font-size: 7px;
  letter-spacing: 0.2em;
  color: #7c6fa3;
}
.cu-gc-grid {
  margin-top: 12px;
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 8px;
}
.cu-gc-service {
  position: relative;
  border-radius: 15px;
  background: rgba(255, 255, 255, 0.5);
  border: 1px solid rgba(255, 255, 255, 0.9);
  padding: 13px 12px;
  min-height: 82px;
  transition: transform 0.22s ease, background-color 0.22s ease;
}
.cu-gc-service:hover {
  transform: translateY(-3px);
  background: rgba(255, 255, 255, 0.72);
}
.cu-gc-service-no {
  font-family: 'Courier New', monospace;
  font-size: 8px;
  color: #a78bfa;
}
.cu-gc-service-dot {
  display: block;
  margin-top: 9px;
  height: 5px;
  width: 5px;
  border-radius: 50%;
  background: #7c3aed;
}
.cu-gc-service h3 {
  margin-top: 7px;
  font-size: 13px;
  font-weight: 600;
  line-height: 1.35;
  color: #2d2547;
}
.cu-gc-empty {
  margin-top: 10px;
  border: 1px dashed rgba(91, 33, 182, 0.3);
  border-radius: 12px;
  padding: 12px;
  font-size: 10px;
  color: #6d6390;
}
.cu-gc-contact {
  margin-top: 18px;
  border-radius: 24px;
  border: 1px solid rgba(255, 255, 255, 0.8);
  background: rgba(76, 29, 149, 0.72);
  backdrop-filter: blur(14px);
  -webkit-backdrop-filter: blur(14px);
  padding: 18px 17px;
  box-shadow: 0 22px 44px -30px rgba(49, 16, 100, 0.85);
}
.cu-gc-contact-label {
  font-size: 10px;
  letter-spacing: 0.1em;
  color: #d8ccf6;
}
.cu-gc-phone {
  margin-top: 10px;
  display: flex;
  align-items: center;
  gap: 8px;
  border-radius: 999px;
  background: rgba(255, 255, 255, 0.92);
  padding: 11px 14px;
  font-size: 17px;
  font-weight: 600;
  color: #3b1d6e;
  transition: background-color 0.2s ease;
}
.cu-gc-phone span {
  flex: 1;
}
.cu-gc-phone:hover {
  background: #ffffff;
}
.cu-gc-mail {
  display: block;
  margin-top: 9px;
  font-family: 'Courier New', monospace;
  font-size: 9px;
  color: #c4b3ea;
  word-break: break-all;
}
.cu-gc-contact .cu-gc-empty {
  color: #d8ccf6;
  border-color: rgba(255, 255, 255, 0.3);
}
.cu-gc-foot {
  margin-top: 18px;
  display: flex;
  justify-content: space-between;
  padding: 0 5px;
  font-family: 'Courier New', monospace;
  font-size: 7px;
  letter-spacing: 0.14em;
  color: #7a6d99;
}
@keyframes cu-gc-in {
  from {
    opacity: 0;
    transform: translateY(18px);
  }
  to {
    opacity: 1;
    transform: none;
  }
}
.cu-gc-in {
  animation: cu-gc-in 0.8s cubic-bezier(0.22, 1, 0.36, 1) both;
}
.cu-gc-d1 {
  animation-delay: 0.08s;
}
.cu-gc-d2 {
  animation-delay: 0.18s;
}
.cu-gc-d3 {
  animation-delay: 0.27s;
}
.cu-gc-d4 {
  animation-delay: 0.36s;
}
@media (prefers-reduced-motion: reduce) {
  .cu-gc-in {
    animation: none;
  }
  .cu-gc-service,
  .cu-gc-phone {
    transition: none;
  }
}
</style>
