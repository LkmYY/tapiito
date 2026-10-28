<template>
  <div class="cu-wc-root min-h-screen">
    <div class="cu-wc-wrap mx-auto w-full max-w-md">
      <header class="cu-wc-top cu-wc-in">
        <span>WHITE CUBE</span>
        <span>EXHIBITION / {{ nfcId ? String(nfcId).slice(-3) : '026' }}</span>
      </header>

      <section class="cu-wc-room cu-wc-in cu-wc-d1">
        <div class="cu-wc-frame">
          <img v-if="images.length" :src="images[0]" :alt="payload?.companyName || '展览作品'" />
          <div v-else class="cu-wc-art" aria-hidden="true">
            <span class="cu-wc-art-block cu-wc-art-block--a"></span>
            <span class="cu-wc-art-block cu-wc-art-block--b"></span>
            <span class="cu-wc-art-block cu-wc-art-block--c"></span>
            <span class="cu-wc-art-line"></span>
            <span class="cu-wc-art-caption">{{ initial }}</span>
          </div>
        </div>
        <div class="cu-wc-spotlight" aria-hidden="true"></div>
      </section>

      <section class="cu-wc-label cu-wc-in cu-wc-d2">
        <div class="cu-wc-label-head">
          <span>WALL LABEL / 01</span>
          <span>2026</span>
        </div>
        <h1>{{ payload?.companyName || '企业主页' }}</h1>
        <p class="cu-wc-medium">企业实践 · 综合媒介<br />Mixed media, ongoing</p>
        <p class="cu-wc-desc">{{ payload?.description }}</p>
      </section>

      <section class="cu-wc-exhibition cu-wc-in cu-wc-d3">
        <div class="cu-wc-section-head">
          <h2>ON VIEW</h2>
          <span>服务项目 / {{ String(services.length).padStart(2, '0') }}</span>
        </div>
        <div v-if="services.length" class="cu-wc-list">
          <div v-for="(service, i) in services" :key="i" class="cu-wc-item">
            <span class="cu-wc-item-no">{{ String(i + 1).padStart(2, '0') }}</span>
            <span class="cu-wc-item-name">{{ service.name }}</span>
            <span class="cu-wc-item-dot"></span>
          </div>
        </div>
      </section>

      <section class="cu-wc-contact cu-wc-in cu-wc-d4">
        <div>
          <p class="cu-wc-contact-label">GALLERY DESK</p>
          <a v-if="payload?.phone" :href="'tel:' + payload.phone" class="cu-wc-phone">
            <Phone :size="15" />
            <span>{{ payload.phone }}</span>
          </a>
        </div>
        <p class="cu-wc-hours">预约参观<br />MON — SAT</p>
      </section>

      <footer class="cu-wc-foot cu-wc-in cu-wc-d4">
        <span>{{ payload?.companyName || '企业主页' }}</span>
        <span>{{ nfcId ? 'NFC ' + nfcId : 'NFC —' }}</span>
      </footer>
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
const images = computed(() => {
  const list = props.payload?.images
  return Array.isArray(list) ? list.filter((src) => typeof src === 'string' && src.trim()) : []
})
const initial = computed(() => (props.payload?.companyName || '未').trim().slice(0, 1))
</script>

<style scoped>
.cu-wc-root {
  background: #fbfbf9;
  color: #161616;
  font-family: 'Helvetica Neue', 'PingFang SC', sans-serif;
  overflow-x: hidden;
}
.cu-wc-wrap {
  padding: 22px 22px 30px;
}
.cu-wc-top {
  display: flex;
  justify-content: space-between;
  border-bottom: 1px solid rgba(22, 22, 22, 0.2);
  padding-bottom: 8px;
  font-family: 'Courier New', monospace;
  font-size: 8px;
  letter-spacing: 0.22em;
  color: #77736c;
}
.cu-wc-room {
  position: relative;
  margin-top: 26px;
  padding: 26px 22px 0;
}
.cu-wc-spotlight {
  position: absolute;
  left: 50%;
  top: -26px;
  height: 220px;
  width: 280px;
  transform: translateX(-50%);
  background: radial-gradient(ellipse at 50% 20%, rgba(255, 239, 196, 0.5), transparent 66%);
  pointer-events: none;
}
.cu-wc-frame {
  position: relative;
  padding: 11px;
  background: #ffffff;
  border: 1px solid rgba(22, 22, 22, 0.18);
  box-shadow:
    0 2px 0 rgba(22, 22, 22, 0.04),
    0 34px 48px -36px rgba(22, 22, 22, 0.45);
}
.cu-wc-frame img {
  display: block;
  width: 100%;
  aspect-ratio: 4 / 5;
  object-fit: cover;
  filter: saturate(0.86);
}
.cu-wc-art {
  position: relative;
  aspect-ratio: 4 / 5;
  background: #f0eee7;
  overflow: hidden;
}
.cu-wc-art-block {
  position: absolute;
}
.cu-wc-art-block--a {
  left: 14%;
  top: 18%;
  width: 46%;
  height: 54%;
  background: #1d2b45;
  transform: rotate(-3deg);
}
.cu-wc-art-block--b {
  right: 12%;
  top: 31%;
  width: 30%;
  height: 42%;
  background: #c65d3f;
  mix-blend-mode: multiply;
}
.cu-wc-art-block--c {
  left: 21%;
  bottom: 12%;
  width: 58%;
  height: 11%;
  background: #d8c9a6;
}
.cu-wc-art-line {
  position: absolute;
  left: 8%;
  right: 8%;
  top: 50%;
  height: 1px;
  background: rgba(22, 22, 22, 0.3);
  transform: rotate(8deg);
}
.cu-wc-art-caption {
  position: absolute;
  right: 13px;
  bottom: 11px;
  font-family: Georgia, serif;
  font-size: 34px;
  color: rgba(22, 22, 22, 0.14);
}
.cu-wc-label {
  margin-top: 26px;
  border-top: 1px solid #161616;
  padding-top: 11px;
}
.cu-wc-label-head {
  display: flex;
  justify-content: space-between;
  font-family: 'Courier New', monospace;
  font-size: 7px;
  letter-spacing: 0.18em;
  color: #8a857c;
}
.cu-wc-label h1 {
  margin-top: 10px;
  font-family: 'Songti SC', 'Noto Serif SC', Georgia, serif;
  font-size: 35px;
  line-height: 1.16;
  font-weight: 500;
  letter-spacing: -0.01em;
  word-break: break-word;
}
.cu-wc-medium {
  margin-top: 9px;
  font-size: 10px;
  line-height: 1.65;
  color: #6e6a62;
}
.cu-wc-desc {
  margin-top: 11px;
  max-width: 34ch;
  font-size: 12px;
  line-height: 1.85;
  color: #4f4c46;
}
.cu-wc-exhibition {
  margin-top: 28px;
}
.cu-wc-section-head {
  display: flex;
  align-items: baseline;
  justify-content: space-between;
  border-bottom: 1px solid #161616;
  padding-bottom: 6px;
}
.cu-wc-section-head h2 {
  font-family: 'Courier New', monospace;
  font-size: 9px;
  letter-spacing: 0.24em;
}
.cu-wc-section-head span {
  font-size: 8px;
  letter-spacing: 0.08em;
  color: #8a857c;
}
.cu-wc-item {
  display: flex;
  align-items: center;
  gap: 12px;
  border-bottom: 1px solid rgba(22, 22, 22, 0.13);
  padding: 15px 0;
  transition: padding-left 0.22s ease;
}
.cu-wc-item:hover {
  padding-left: 5px;
}
.cu-wc-item-no {
  font-family: 'Courier New', monospace;
  font-size: 8px;
  color: #a09a90;
  width: 20px;
}
.cu-wc-item-name {
  flex: 1;
  font-family: 'Songti SC', serif;
  font-size: 18px;
}
.cu-wc-item-dot {
  height: 6px;
  width: 6px;
  border-radius: 50%;
  border: 1px solid #161616;
}
.cu-wc-empty {
  padding: 13px 0;
  font-size: 11px;
  color: #8a857c;
}
.cu-wc-contact {
  margin-top: 26px;
  border-top: 1px solid #161616;
  border-bottom: 1px solid #161616;
  padding: 14px 0;
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 12px;
}
.cu-wc-contact-label {
  font-family: 'Courier New', monospace;
  font-size: 7px;
  letter-spacing: 0.2em;
  color: #8a857c;
}
.cu-wc-phone {
  margin-top: 6px;
  display: inline-flex;
  align-items: center;
  gap: 7px;
  font-size: 18px;
  color: #161616;
  border-bottom: 1px solid rgba(198, 93, 63, 0.6);
  padding-bottom: 2px;
  transition: color 0.2s ease;
}
.cu-wc-phone:hover {
  color: #c65d3f;
}
.cu-wc-hours {
  text-align: right;
  font-family: 'Courier New', monospace;
  font-size: 7px;
  line-height: 1.7;
  letter-spacing: 0.1em;
  color: #8a857c;
}
.cu-wc-foot {
  margin-top: 16px;
  display: flex;
  justify-content: space-between;
  font-family: 'Courier New', monospace;
  font-size: 7px;
  letter-spacing: 0.14em;
  color: #8a857c;
}
@keyframes cu-wc-in {
  from {
    opacity: 0;
    transform: translateY(16px);
  }
  to {
    opacity: 1;
    transform: none;
  }
}
.cu-wc-in {
  animation: cu-wc-in 0.85s cubic-bezier(0.22, 1, 0.36, 1) both;
}
.cu-wc-d1 {
  animation-delay: 0.1s;
}
.cu-wc-d2 {
  animation-delay: 0.2s;
}
.cu-wc-d3 {
  animation-delay: 0.3s;
}
.cu-wc-d4 {
  animation-delay: 0.4s;
}
@media (prefers-reduced-motion: reduce) {
  .cu-wc-in {
    animation: none;
  }
  .cu-wc-item,
  .cu-wc-phone {
    transition: none;
  }
}
</style>
