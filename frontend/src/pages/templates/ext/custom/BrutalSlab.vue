<template>
  <div class="cu-bs-root min-h-screen">
    <div class="cu-bs-wrap mx-auto w-full max-w-md">
      <div class="cu-bs-status cu-bs-in">
        <span class="cu-bs-live"><i></i>OPEN FOR WORK</span>
        <span class="cu-bs-id">NFC / {{ nfcId || '—' }}</span>
      </div>

      <header class="cu-bs-hero cu-bs-in cu-bs-d1">
        <p class="cu-bs-tag">NO NONSENSE<br />COMPANY PROFILE</p>
        <h1 class="cu-bs-name">{{ payload?.companyName || '企业主页' }}</h1>
        <span class="cu-bs-star">✳</span>
      </header>

      <div class="cu-bs-desc cu-bs-in cu-bs-d2">
        <span class="cu-bs-desc-mark">↳</span>
        <p>{{ payload?.description }}</p>
      </div>

      <div class="cu-bs-visual cu-bs-in cu-bs-d2">
        <img v-if="images.length" :src="images[0]" :alt="payload?.companyName || '企业图片'" />
        <div v-else class="cu-bs-visual-fallback">
          <span>{{ initial }}</span>
          <small>NO IMAGE / TYPE ONLY</small>
        </div>
        <div class="cu-bs-visual-strip">
          <span>SERVICE</span><span>DESIGN</span><span>DELIVERY</span>
        </div>
      </div>

      <section class="cu-bs-services cu-bs-in cu-bs-d3">
        <div class="cu-bs-head">
          <h2>WHAT WE DO</h2>
          <span>{{ String(services.length).padStart(2, '0') }} ITEMS</span>
        </div>
        <div v-if="services.length" class="cu-bs-stack">
          <article v-for="(service, i) in services" :key="i" class="cu-bs-slab">
            <span class="cu-bs-slab-no">{{ String(i + 1).padStart(2, '0') }}</span>
            <h3>{{ service.name }}</h3>
            <ArrowDownRight :size="22" class="cu-bs-slab-arrow" />
          </article>
        </div>
      </section>

      <a v-if="payload?.phone" :href="'tel:' + payload.phone" class="cu-bs-cta cu-bs-in cu-bs-d4">
        <span>
          <small>CALL NOW / 直接致电</small>
          <strong>{{ payload.phone }}</strong>
        </span>
        <PhoneCall :size="26" />
      </a>

      <footer class="cu-bs-foot cu-bs-in cu-bs-d4">
        <span>{{ payload?.companyName || '企业主页' }}</span>
        <span>EST. 2026 / NO TEMPLATES</span>
      </footer>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { ArrowDownRight, PhoneCall } from 'lucide-vue-next'

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
.cu-bs-root {
  background: #e9ff3f;
  color: #0a0a0a;
  font-family: 'Arial Black', 'Helvetica Neue', 'PingFang SC', sans-serif;
  overflow-x: hidden;
}
.cu-bs-wrap {
  padding: 18px 16px 32px;
}
.cu-bs-status {
  display: flex;
  align-items: center;
  justify-content: space-between;
  border: 3px solid #0a0a0a;
  background: #0a0a0a;
  color: #e9ff3f;
  padding: 7px 10px;
  font-family: 'Courier New', monospace;
  font-size: 9px;
  letter-spacing: 0.12em;
}
.cu-bs-live {
  display: inline-flex;
  align-items: center;
  gap: 6px;
}
.cu-bs-live i {
  height: 6px;
  width: 6px;
  border-radius: 50%;
  background: #e9ff3f;
  animation: cu-bs-blink 1.8s steps(2, end) infinite;
}
.cu-bs-hero {
  position: relative;
  margin-top: 14px;
  border: 3px solid #0a0a0a;
  background: #fffdf5;
  padding: 18px 16px 22px;
  box-shadow: 7px 7px 0 #0a0a0a;
  overflow: hidden;
}
.cu-bs-tag {
  font-family: 'Courier New', monospace;
  font-size: 9px;
  font-weight: 700;
  line-height: 1.5;
  letter-spacing: 0.16em;
}
.cu-bs-name {
  position: relative;
  margin-top: 12px;
  font-size: 50px;
  line-height: 0.9;
  font-weight: 900;
  letter-spacing: -0.055em;
  text-transform: none;
  word-break: break-word;
  transform: rotate(-1.2deg);
}
.cu-bs-star {
  position: absolute;
  right: 12px;
  bottom: 8px;
  font-size: 42px;
  line-height: 1;
  color: #ff2d55;
  transform: rotate(12deg);
}
.cu-bs-desc {
  margin-top: 16px;
  display: flex;
  gap: 10px;
  border: 3px solid #0a0a0a;
  background: #e9ff3f;
  padding: 12px 14px;
}
.cu-bs-desc-mark {
  font-size: 22px;
  line-height: 1;
  font-weight: 900;
}
.cu-bs-desc p {
  font-family: 'PingFang SC', 'Microsoft YaHei', sans-serif;
  font-size: 13px;
  font-weight: 500;
  line-height: 1.65;
}
.cu-bs-visual {
  position: relative;
  margin-top: 16px;
  border: 3px solid #0a0a0a;
  background: #0a0a0a;
  overflow: hidden;
}
.cu-bs-visual img {
  width: 100%;
  aspect-ratio: 16 / 10;
  object-fit: cover;
  display: block;
  filter: grayscale(1) contrast(1.15);
}
.cu-bs-visual-fallback {
  position: relative;
  aspect-ratio: 16 / 9;
  display: flex;
  align-items: center;
  justify-content: center;
  background:
    repeating-linear-gradient(45deg, rgba(233, 255, 63, 0.18) 0 8px, transparent 8px 16px),
    #0a0a0a;
}
.cu-bs-visual-fallback span {
  font-size: 130px;
  line-height: 1;
  color: #e9ff3f;
  letter-spacing: -0.08em;
}
.cu-bs-visual-fallback small {
  position: absolute;
  bottom: 9px;
  right: 11px;
  font-family: 'Courier New', monospace;
  font-size: 8px;
  letter-spacing: 0.14em;
  color: #e9ff3f;
}
.cu-bs-visual-strip {
  display: flex;
  justify-content: space-between;
  padding: 7px 10px;
  background: #0a0a0a;
  color: #e9ff3f;
  font-family: 'Courier New', monospace;
  font-size: 8px;
  letter-spacing: 0.2em;
  border-top: 1px solid #e9ff3f;
}
.cu-bs-services {
  margin-top: 26px;
}
.cu-bs-head {
  display: flex;
  align-items: baseline;
  justify-content: space-between;
  border-bottom: 3px solid #0a0a0a;
  padding-bottom: 7px;
}
.cu-bs-head h2 {
  font-size: 18px;
  font-weight: 900;
  letter-spacing: 0.02em;
}
.cu-bs-head span {
  font-family: 'Courier New', monospace;
  font-size: 9px;
  letter-spacing: 0.12em;
}
.cu-bs-stack {
  margin-top: 10px;
  display: flex;
  flex-direction: column;
  gap: 7px;
}
.cu-bs-slab {
  display: flex;
  align-items: center;
  gap: 12px;
  border: 3px solid #0a0a0a;
  background: #fffdf5;
  padding: 12px 13px;
  transition: transform 0.18s ease, box-shadow 0.18s ease, background-color 0.18s ease;
}
.cu-bs-slab:nth-child(even) {
  background: #ff2d55;
  color: #fffdf5;
}
.cu-bs-slab:hover {
  transform: translate(-3px, -3px);
  box-shadow: 4px 4px 0 #0a0a0a;
}
.cu-bs-slab-no {
  font-family: 'Courier New', monospace;
  font-size: 10px;
  font-weight: 700;
}
.cu-bs-slab h3 {
  flex: 1;
  font-size: 18px;
  font-weight: 900;
  letter-spacing: -0.02em;
}
.cu-bs-slab-arrow {
  flex-shrink: 0;
}
.cu-bs-cta {
  margin-top: 22px;
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 12px;
  border: 3px solid #0a0a0a;
  background: #0a0a0a;
  color: #e9ff3f;
  padding: 16px;
  box-shadow: 7px 7px 0 #ff2d55;
  transition: transform 0.15s ease, box-shadow 0.15s ease;
}
.cu-bs-cta small {
  display: block;
  font-family: 'Courier New', monospace;
  font-size: 8px;
  letter-spacing: 0.18em;
  color: #e9ff3f;
  opacity: 0.7;
}
.cu-bs-cta strong {
  display: block;
  margin-top: 4px;
  font-size: 24px;
  font-weight: 900;
  letter-spacing: -0.02em;
}
.cu-bs-cta:active {
  transform: translate(4px, 4px);
  box-shadow: 3px 3px 0 #ff2d55;
}
.cu-bs-empty {
  margin-top: 12px;
  border: 3px dashed #0a0a0a;
  padding: 16px;
  font-family: 'PingFang SC', sans-serif;
  font-size: 12px;
}
.cu-bs-foot {
  margin-top: 24px;
  display: flex;
  justify-content: space-between;
  gap: 10px;
  font-family: 'Courier New', monospace;
  font-size: 8px;
  letter-spacing: 0.12em;
}
@keyframes cu-bs-in {
  0% {
    opacity: 0;
    transform: translateY(22px) rotate(1.5deg);
  }
  100% {
    opacity: 1;
    transform: none;
  }
}
.cu-bs-in {
  animation: cu-bs-in 0.5s cubic-bezier(0.2, 0.9, 0.2, 1) both;
}
.cu-bs-d1 {
  animation-delay: 0.07s;
}
.cu-bs-d2 {
  animation-delay: 0.15s;
}
.cu-bs-d3 {
  animation-delay: 0.23s;
}
.cu-bs-d4 {
  animation-delay: 0.31s;
}
@keyframes cu-bs-blink {
  50% {
    opacity: 0.25;
  }
}
@media (prefers-reduced-motion: reduce) {
  .cu-bs-in {
    animation: none;
  }
  .cu-bs-live i {
    animation: none;
  }
  .cu-bs-slab,
  .cu-bs-cta {
    transition: none;
  }
}
</style>
