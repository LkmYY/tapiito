<template>
  <div class="cu-vs-root min-h-screen">
    <div class="cu-vs-wrap mx-auto w-full max-w-md">
      <header class="cu-vs-top cu-vs-in">
        <span>ANALOG SESSIONS</span>
        <span>33⅓ RPM / STEREO</span>
      </header>

      <section class="cu-vs-cover">
        <div class="cu-vs-sleeve cu-vs-in cu-vs-d1">
          <img v-if="images.length" :src="images[0]" :alt="payload?.companyName || '封面'" class="cu-vs-sleeve-img" />
          <div v-else class="cu-vs-disc" aria-hidden="true">
            <span class="cu-vs-grooves"></span>
            <span class="cu-vs-label">
              <i>拾光</i>
              <b>{{ initial }}</b>
            </span>
            <span class="cu-vs-hole"></span>
          </div>
          <span class="cu-vs-sleeve-cat">CAT. NO. {{ nfcId ? String(nfcId).slice(-4) : '0026' }}</span>
        </div>

        <div class="cu-vs-title cu-vs-in cu-vs-d2">
          <p>THE COMPANY LP</p>
          <h1>{{ payload?.companyName || '企业主页' }}</h1>
          <p class="cu-vs-desc">{{ payload?.description }}</p>
        </div>
      </section>

      <section class="cu-vs-tracklist cu-vs-in cu-vs-d3">
        <div class="cu-vs-side">
          <span>SIDE A</span>
          <span>SERVICES</span>
        </div>
        <div v-if="services.length" class="cu-vs-tracks">
          <div v-for="(service, i) in services" :key="i" class="cu-vs-track">
            <span class="cu-vs-track-no">A{{ i + 1 }}</span>
            <span class="cu-vs-track-name">{{ service.name }}</span>
            <span class="cu-vs-track-time">3:{{ String(21 + i * 14).slice(-2) }}</span>
          </div>
        </div>
        <div class="cu-vs-side cu-vs-side--bottom">
          <span>SIDE B</span>
          <span>ENQUIRIES</span>
        </div>
        <a v-if="payload?.phone" :href="'tel:' + payload.phone" class="cu-vs-contact">
          <span class="cu-vs-contact-label">BOOK A SESSION</span>
          <span class="cu-vs-contact-phone">{{ payload.phone }}</span>
          <Phone :size="15" />
        </a>
      </section>

      <footer class="cu-vs-foot cu-vs-in cu-vs-d4">
        <span>RECORDED IN 2026</span>
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
.cu-vs-root {
  background: #efebe1;
  color: #191817;
  font-family: 'Helvetica Neue', 'PingFang SC', sans-serif;
  overflow-x: hidden;
}
.cu-vs-wrap {
  padding: 20px 18px 30px;
}
.cu-vs-top {
  display: flex;
  justify-content: space-between;
  border-bottom: 2px solid #191817;
  padding-bottom: 8px;
  font-family: 'Courier New', monospace;
  font-size: 8px;
  letter-spacing: 0.2em;
}
.cu-vs-cover {
  margin-top: 22px;
}
.cu-vs-sleeve {
  position: relative;
  aspect-ratio: 1 / 1;
  background: #191817;
  overflow: hidden;
  box-shadow: 12px 14px 0 rgba(25, 24, 23, 0.12);
}
.cu-vs-sleeve::before {
  content: '';
  position: absolute;
  inset: 10px;
  border: 1px solid rgba(239, 235, 225, 0.28);
  z-index: 2;
  pointer-events: none;
}
.cu-vs-sleeve-img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  filter: grayscale(1) contrast(1.1);
}
.cu-vs-disc {
  position: absolute;
  inset: 8%;
  border-radius: 50%;
  background:
    repeating-radial-gradient(circle at 50% 50%, #111 0 2px, #1d1c1a 2px 4px, #131211 4px 6px);
  display: flex;
  align-items: center;
  justify-content: center;
  animation: cu-vs-spin 26s linear infinite;
}
.cu-vs-grooves {
  position: absolute;
  inset: 6%;
  border-radius: 50%;
  border: 1px solid rgba(239, 235, 225, 0.08);
  box-shadow:
    inset 0 0 0 9px rgba(255, 255, 255, 0.018),
    inset 0 0 0 22px rgba(255, 255, 255, 0.014),
    inset 0 0 0 38px rgba(255, 255, 255, 0.012);
}
.cu-vs-label {
  position: relative;
  z-index: 2;
  height: 36%;
  width: 36%;
  border-radius: 50%;
  background: #b3322c;
  color: #f4eddc;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
}
.cu-vs-label i {
  font-family: 'Songti SC', serif;
  font-size: 10px;
  font-style: normal;
  letter-spacing: 0.2em;
}
.cu-vs-label b {
  font-family: Georgia, serif;
  font-size: 43px;
  line-height: 0.9;
  font-weight: 400;
}
.cu-vs-hole {
  position: absolute;
  z-index: 3;
  height: 9px;
  width: 9px;
  border-radius: 50%;
  background: #efebe1;
  box-shadow: 0 0 0 3px #191817;
}
.cu-vs-sleeve-cat {
  position: absolute;
  left: 20px;
  bottom: 18px;
  z-index: 3;
  font-family: 'Courier New', monospace;
  font-size: 7px;
  letter-spacing: 0.18em;
  color: rgba(239, 235, 225, 0.8);
}
.cu-vs-title {
  margin-top: 24px;
}
.cu-vs-title > p:first-child {
  font-family: 'Courier New', monospace;
  font-size: 8px;
  letter-spacing: 0.26em;
  color: #9b3a33;
}
.cu-vs-title h1 {
  margin-top: 8px;
  font-family: 'Songti SC', 'Noto Serif SC', Georgia, serif;
  font-size: 40px;
  line-height: 1.05;
  font-weight: 600;
  letter-spacing: -0.01em;
  word-break: break-word;
}
.cu-vs-desc {
  margin-top: 10px;
  max-width: 34ch;
  font-size: 12px;
  line-height: 1.8;
  color: #6d675c;
}
.cu-vs-tracklist {
  margin-top: 26px;
  border-top: 2px solid #191817;
}
.cu-vs-side {
  display: flex;
  justify-content: space-between;
  border-bottom: 1px solid #191817;
  padding: 6px 0;
  font-family: 'Courier New', monospace;
  font-size: 7px;
  letter-spacing: 0.2em;
  color: #6d675c;
}
.cu-vs-side--bottom {
  margin-top: 20px;
  border-top: 1px solid #191817;
}
.cu-vs-track {
  display: flex;
  align-items: baseline;
  gap: 11px;
  border-bottom: 1px solid rgba(25, 24, 23, 0.16);
  padding: 13px 0;
  transition: background-color 0.2s ease;
}
.cu-vs-track:hover {
  background: rgba(25, 24, 23, 0.04);
}
.cu-vs-track-no {
  font-family: 'Courier New', monospace;
  font-size: 9px;
  color: #9b3a33;
  width: 22px;
}
.cu-vs-track-name {
  flex: 1;
  font-family: 'Songti SC', serif;
  font-size: 17px;
  font-weight: 600;
}
.cu-vs-track-time {
  font-family: 'Courier New', monospace;
  font-size: 8px;
  color: #857e70;
}
.cu-vs-contact {
  display: flex;
  align-items: center;
  gap: 9px;
  margin-top: 12px;
  background: #191817;
  color: #f4eddc;
  padding: 14px 15px;
  transition: background-color 0.2s ease;
}
.cu-vs-contact:hover {
  background: #b3322c;
}
.cu-vs-contact-label {
  font-family: 'Courier New', monospace;
  font-size: 7px;
  letter-spacing: 0.16em;
  color: #c7bca6;
}
.cu-vs-contact-phone {
  flex: 1;
  text-align: right;
  font-size: 18px;
  letter-spacing: 0.02em;
}
.cu-vs-empty {
  padding: 13px 0;
  font-size: 11px;
  color: #857e70;
}
.cu-vs-foot {
  margin-top: 18px;
  display: flex;
  justify-content: space-between;
  font-family: 'Courier New', monospace;
  font-size: 7px;
  letter-spacing: 0.14em;
  color: #857e70;
}
@keyframes cu-vs-spin {
  to {
    transform: rotate(360deg);
  }
}
@keyframes cu-vs-in {
  from {
    opacity: 0;
    transform: translateY(16px);
  }
  to {
    opacity: 1;
    transform: none;
  }
}
.cu-vs-in {
  animation: cu-vs-in 0.8s cubic-bezier(0.22, 1, 0.36, 1) both;
}
.cu-vs-d1 {
  animation-delay: 0.08s;
}
.cu-vs-d2 {
  animation-delay: 0.18s;
}
.cu-vs-d3 {
  animation-delay: 0.27s;
}
.cu-vs-d4 {
  animation-delay: 0.36s;
}
@media (prefers-reduced-motion: reduce) {
  .cu-vs-in {
    animation: none;
  }
  .cu-vs-disc {
    animation: none;
  }
  .cu-vs-track,
  .cu-vs-contact {
    transition: none;
  }
}
</style>
