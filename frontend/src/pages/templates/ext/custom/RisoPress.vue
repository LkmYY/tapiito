<template>
  <div class="cu-rp-root min-h-screen">
    <div class="cu-rp-wrap mx-auto w-full max-w-md">
      <header class="cu-rp-top cu-rp-in">
        <span>RISOGRAPH EDITION</span>
        <span>02 COLORS / {{ nfcId ? String(nfcId).slice(-3) : '026' }}</span>
      </header>

      <section class="cu-rp-hero">
        <div class="cu-rp-halftone cu-rp-in cu-rp-d1" aria-hidden="true"></div>
        <p class="cu-rp-overline cu-rp-in cu-rp-d1">PRINTED MATTER / 企业手册</p>
        <h1 class="cu-rp-name cu-rp-in cu-rp-d1" :data-text="payload?.companyName || '企业主页'">{{ payload?.companyName || '企业主页' }}</h1>
        <div class="cu-rp-band cu-rp-in cu-rp-d2">
          <span>ORANGE</span>
          <i>×</i>
          <span>INK BLUE</span>
        </div>
        <p class="cu-rp-desc cu-rp-in cu-rp-d2">{{ payload?.description }}</p>
      </section>

      <section v-if="images.length" class="cu-rp-photo cu-rp-in cu-rp-d3">
        <img :src="images[0]" :alt="payload?.companyName || '企业图片'" />
        <span class="cu-rp-photo-mark">OVERPRINT / 01</span>
      </section>

      <section class="cu-rp-services cu-rp-in cu-rp-d3">
        <div class="cu-rp-services-head">
          <h2>服务印刷清单</h2>
          <span>SPECIMEN</span>
        </div>
        <div v-if="services.length" class="cu-rp-list">
          <article v-for="(service, i) in services" :key="i" class="cu-rp-item">
            <span class="cu-rp-item-no">{{ String(i + 1).padStart(2, '0') }}</span>
            <h3>{{ service.name }}</h3>
            <span class="cu-rp-item-chip">INK {{ i + 1 }}</span>
          </article>
        </div>
      </section>

      <section class="cu-rp-contact cu-rp-in cu-rp-d4">
        <div class="cu-rp-contact-head">
          <span>CONTACT IMPRINT</span>
          <span>可直接致电</span>
        </div>
        <a v-if="payload?.phone" :href="'tel:' + payload.phone" class="cu-rp-phone">
          <Phone :size="16" />
          <span>{{ payload.phone }}</span>
        </a>
        <div class="cu-rp-contact-foot">
          <span>{{ payload?.companyName || '企业主页' }}</span>
          <span>{{ nfcId ? 'NFC ' + nfcId : 'NFC —' }}</span>
        </div>
      </section>
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
</script>

<style scoped>
.cu-rp-root {
  background: #f2ead8;
  color: #17265e;
  font-family: 'Helvetica Neue', 'PingFang SC', sans-serif;
  overflow-x: hidden;
}
.cu-rp-wrap {
  padding: 20px 18px 30px;
}
.cu-rp-top {
  display: flex;
  justify-content: space-between;
  border-top: 6px solid #ef5b2a;
  border-bottom: 1px solid rgba(23, 38, 94, 0.3);
  padding: 7px 0;
  font-family: 'Courier New', monospace;
  font-size: 8px;
  letter-spacing: 0.18em;
  color: #5b6390;
}
.cu-rp-hero {
  position: relative;
  padding-top: 28px;
}
.cu-rp-halftone {
  position: absolute;
  right: -10px;
  top: 4px;
  height: 150px;
  width: 150px;
  border-radius: 50%;
  background-image: radial-gradient(#ef5b2a 1.6px, transparent 1.8px);
  background-size: 9px 9px;
  opacity: 0.5;
  -webkit-mask-image: radial-gradient(circle, #000 35%, transparent 70%);
  mask-image: radial-gradient(circle, #000 35%, transparent 70%);
}
.cu-rp-overline {
  font-family: 'Courier New', monospace;
  font-size: 8px;
  letter-spacing: 0.24em;
  color: #ef5b2a;
}
.cu-rp-name {
  position: relative;
  margin-top: 13px;
  max-width: 6em;
  font-size: 52px;
  line-height: 0.95;
  font-weight: 800;
  letter-spacing: -0.055em;
  color: #17265e;
  word-break: break-word;
  text-shadow: 4px 4px 0 rgba(239, 91, 42, 0.55);
}
.cu-rp-band {
  margin-top: 18px;
  display: inline-flex;
  align-items: center;
  gap: 8px;
  background: #17265e;
  color: #f2ead8;
  padding: 5px 10px;
  font-family: 'Courier New', monospace;
  font-size: 7px;
  letter-spacing: 0.16em;
}
.cu-rp-band i {
  font-style: normal;
  color: #ef5b2a;
}
.cu-rp-desc {
  margin-top: 14px;
  max-width: 34ch;
  font-size: 12px;
  line-height: 1.8;
  color: #46507e;
}
.cu-rp-photo {
  position: relative;
  margin-top: 22px;
  border: 2px solid #17265e;
  background: #ef5b2a;
  padding: 7px;
}
.cu-rp-photo img {
  display: block;
  width: 100%;
  aspect-ratio: 5 / 3;
  object-fit: cover;
  filter: grayscale(1) contrast(1.15) sepia(0.25);
  mix-blend-mode: multiply;
}
.cu-rp-photo-mark {
  position: absolute;
  right: 8px;
  bottom: 8px;
  background: #f2ead8;
  padding: 2px 5px;
  font-family: 'Courier New', monospace;
  font-size: 6px;
  letter-spacing: 0.12em;
  color: #17265e;
}
.cu-rp-services {
  margin-top: 26px;
}
.cu-rp-services-head {
  display: flex;
  align-items: baseline;
  justify-content: space-between;
  border-bottom: 2px solid #17265e;
  padding-bottom: 6px;
}
.cu-rp-services-head h2 {
  font-size: 15px;
  font-weight: 800;
  letter-spacing: 0.01em;
}
.cu-rp-services-head span {
  font-family: 'Courier New', monospace;
  font-size: 7px;
  letter-spacing: 0.18em;
  color: #ef5b2a;
}
.cu-rp-item {
  display: flex;
  align-items: center;
  gap: 11px;
  border-bottom: 1px solid rgba(23, 38, 94, 0.22);
  padding: 14px 0;
  transition: padding-left 0.2s ease;
}
.cu-rp-item:hover {
  padding-left: 5px;
}
.cu-rp-item-no {
  font-family: 'Courier New', monospace;
  font-size: 9px;
  color: #ef5b2a;
  width: 20px;
}
.cu-rp-item h3 {
  flex: 1;
  font-size: 17px;
  font-weight: 700;
  letter-spacing: -0.02em;
}
.cu-rp-item-chip {
  flex-shrink: 0;
  border: 1px solid #ef5b2a;
  color: #ef5b2a;
  padding: 2px 5px;
  font-family: 'Courier New', monospace;
  font-size: 6px;
  letter-spacing: 0.1em;
}
.cu-rp-empty {
  padding: 14px 0;
  font-size: 11px;
  color: #5b6390;
}
.cu-rp-contact {
  margin-top: 24px;
  background: #17265e;
  color: #f2ead8;
  padding: 16px;
  box-shadow: 6px 6px 0 #ef5b2a;
}
.cu-rp-contact-head {
  display: flex;
  justify-content: space-between;
  border-bottom: 1px solid rgba(242, 234, 216, 0.3);
  padding-bottom: 7px;
  font-family: 'Courier New', monospace;
  font-size: 7px;
  letter-spacing: 0.14em;
  color: #b9c0e0;
}
.cu-rp-phone {
  margin-top: 13px;
  display: inline-flex;
  align-items: center;
  gap: 8px;
  font-size: 23px;
  font-weight: 700;
  letter-spacing: -0.02em;
  color: #fff5e0;
  transition: color 0.2s ease;
}
.cu-rp-phone:hover {
  color: #ffb493;
}
.cu-rp-contact .cu-rp-empty {
  color: #b9c0e0;
}
.cu-rp-contact-foot {
  margin-top: 13px;
  display: flex;
  justify-content: space-between;
  padding-top: 8px;
  border-top: 1px solid rgba(242, 234, 216, 0.3);
  font-family: 'Courier New', monospace;
  font-size: 7px;
  letter-spacing: 0.12em;
  color: #b9c0e0;
}
@keyframes cu-rp-in {
  from {
    opacity: 0;
    transform: translateY(16px);
  }
  to {
    opacity: 1;
    transform: none;
  }
}
.cu-rp-in {
  animation: cu-rp-in 0.65s cubic-bezier(0.2, 0.9, 0.2, 1) both;
}
.cu-rp-d1 {
  animation-delay: 0.07s;
}
.cu-rp-d2 {
  animation-delay: 0.16s;
}
.cu-rp-d3 {
  animation-delay: 0.25s;
}
.cu-rp-d4 {
  animation-delay: 0.34s;
}
@media (prefers-reduced-motion: reduce) {
  .cu-rp-in {
    animation: none;
  }
  .cu-rp-item,
  .cu-rp-phone {
    transition: none;
  }
}
</style>
