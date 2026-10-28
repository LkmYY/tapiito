<template>
  <div class="cu-bt-root min-h-screen">
    <div class="cu-bt-wrap mx-auto w-full max-w-md">
      <section class="cu-bt-stage cu-bt-in">
        <img v-if="images.length" :src="images[0]" :alt="payload?.companyName || '空间照片'" class="cu-bt-stage-img" />
        <div v-else class="cu-bt-stage-art" aria-hidden="true">
          <span class="cu-bt-sun"></span>
          <span class="cu-bt-hill cu-bt-hill--back"></span>
          <span class="cu-bt-hill cu-bt-hill--front"></span>
        </div>
        <div class="cu-bt-stage-label">
          <span>BOUTIQUE STAY</span>
          <span>{{ nfcId ? 'ROOM ' + String(nfcId).slice(-2) : 'ROOM 07' }}</span>
        </div>
      </section>

      <div class="cu-bt-key cu-bt-in cu-bt-d1">
        <div class="cu-bt-key-head">
          <span class="cu-bt-key-dot"></span>
          <span>ACCESS CARD</span>
          <ArrowUpRight :size="14" />
        </div>
        <h1>{{ payload?.companyName || '企业主页' }}</h1>
        <p>{{ payload?.description }}</p>
        <div class="cu-bt-key-foot">
          <span>CHECK-IN / 随时</span>
          <span>{{ nfcId || '0000' }}</span>
        </div>
      </div>

      <section class="cu-bt-amenities cu-bt-in cu-bt-d2">
        <div class="cu-bt-section-head">
          <h2>AMENITIES</h2>
          <span>服务与体验</span>
        </div>
        <div v-if="services.length" class="cu-bt-list">
          <div v-for="(service, i) in services" :key="i" class="cu-bt-item">
            <span class="cu-bt-check"><Check :size="12" /></span>
            <span class="cu-bt-item-name">{{ service.name }}</span>
            <span class="cu-bt-item-tag">INCLUDED</span>
          </div>
        </div>
      </section>

      <section class="cu-bt-contact cu-bt-in cu-bt-d3">
        <div class="cu-bt-desk">
          <p>FRONT DESK</p>
          <span>前台值班 · 全天候响应</span>
        </div>
        <a v-if="payload?.phone" :href="'tel:' + payload.phone" class="cu-bt-phone">
          <Phone :size="15" />
          {{ payload.phone }}
        </a>
      </section>

      <footer class="cu-bt-foot cu-bt-in cu-bt-d3">
        <span>{{ payload?.companyName || '企业主页' }}</span>
        <span>STAY / 2026</span>
      </footer>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { ArrowUpRight, Check, Phone } from 'lucide-vue-next'

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
.cu-bt-root {
  background: #e9dfcc;
  color: #1f3a2e;
  font-family: 'PingFang SC', 'Microsoft YaHei', sans-serif;
  overflow-x: hidden;
}
.cu-bt-wrap {
  padding: 14px 14px 28px;
}
.cu-bt-stage {
  position: relative;
  height: 330px;
  overflow: hidden;
  border-radius: 4px;
  background: #24422f;
}
.cu-bt-stage-img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  filter: saturate(0.85) contrast(1.02);
}
.cu-bt-stage-art {
  position: absolute;
  inset: 0;
  background: linear-gradient(180deg, #35604a 0%, #20402f 62%, #172e22 100%);
}
.cu-bt-sun {
  position: absolute;
  left: 50%;
  top: 62px;
  height: 62px;
  width: 62px;
  margin-left: -31px;
  border-radius: 50%;
  background: #e9c98d;
  opacity: 0.85;
}
.cu-bt-hill {
  position: absolute;
  left: -10%;
  right: -10%;
  border-radius: 50% 50% 0 0;
}
.cu-bt-hill--back {
  bottom: 0;
  height: 42%;
  background: #1d3628;
}
.cu-bt-hill--front {
  bottom: -6%;
  height: 32%;
  background: #12271c;
  left: -26%;
  right: 6%;
}
.cu-bt-stage-label {
  position: absolute;
  left: 14px;
  right: 14px;
  bottom: 12px;
  display: flex;
  justify-content: space-between;
  font-family: 'Courier New', monospace;
  font-size: 8px;
  letter-spacing: 0.22em;
  color: rgba(233, 223, 204, 0.85);
}
.cu-bt-key {
  position: relative;
  margin: -46px 12px 0;
  background: #f6f1e4;
  border-radius: 4px;
  padding: 18px 17px 14px;
  box-shadow: 0 24px 46px -30px rgba(15, 35, 25, 0.9);
}
.cu-bt-key::after {
  content: '';
  position: absolute;
  right: 16px;
  top: 16px;
  height: 22px;
  width: 22px;
  border-radius: 3px;
  background: linear-gradient(135deg, #d9cba9, #b7a67f);
  box-shadow: inset 0 0 0 1px rgba(31, 58, 46, 0.25);
}
.cu-bt-key-head {
  display: flex;
  align-items: center;
  gap: 7px;
  font-family: 'Courier New', monospace;
  font-size: 7px;
  letter-spacing: 0.2em;
  color: #6d7f6f;
}
.cu-bt-key-dot {
  height: 5px;
  width: 5px;
  border-radius: 50%;
  background: #c99a4e;
}
.cu-bt-key h1 {
  margin-top: 14px;
  font-family: 'Songti SC', 'Noto Serif SC', Georgia, serif;
  font-size: 32px;
  line-height: 1.14;
  font-weight: 600;
  letter-spacing: 0.01em;
  color: #1f3a2e;
  word-break: break-word;
}
.cu-bt-key p {
  margin-top: 9px;
  max-width: 30ch;
  font-size: 11px;
  line-height: 1.8;
  color: #64756a;
}
.cu-bt-key-foot {
  margin-top: 15px;
  display: flex;
  justify-content: space-between;
  border-top: 1px dashed rgba(31, 58, 46, 0.28);
  padding-top: 9px;
  font-family: 'Courier New', monospace;
  font-size: 7px;
  letter-spacing: 0.14em;
  color: #849082;
}
.cu-bt-amenities {
  margin-top: 28px;
  padding: 0 4px;
}
.cu-bt-section-head {
  display: flex;
  align-items: baseline;
  justify-content: space-between;
  border-bottom: 1px solid rgba(31, 58, 46, 0.3);
  padding-bottom: 7px;
}
.cu-bt-section-head h2 {
  font-family: 'Courier New', monospace;
  font-size: 10px;
  letter-spacing: 0.22em;
}
.cu-bt-section-head span {
  font-size: 9px;
  letter-spacing: 0.08em;
  color: #6d7f6f;
}
.cu-bt-item {
  display: flex;
  align-items: center;
  gap: 11px;
  padding: 14px 0;
  border-bottom: 1px solid rgba(31, 58, 46, 0.12);
  transition: padding-left 0.22s ease;
}
.cu-bt-item:hover {
  padding-left: 5px;
}
.cu-bt-check {
  display: inline-flex;
  height: 22px;
  width: 22px;
  align-items: center;
  justify-content: center;
  border-radius: 50%;
  background: #1f3a2e;
  color: #e9dfcc;
  flex-shrink: 0;
}
.cu-bt-item-name {
  flex: 1;
  font-size: 14px;
  font-weight: 600;
}
.cu-bt-item-tag {
  font-family: 'Courier New', monospace;
  font-size: 6px;
  letter-spacing: 0.14em;
  color: #9a8a5c;
  border: 1px solid rgba(201, 154, 78, 0.55);
  padding: 2px 5px;
}
.cu-bt-empty {
  margin-top: 10px;
  border: 1px dashed rgba(31, 58, 46, 0.3);
  padding: 13px;
  font-size: 11px;
  color: #6d7f6f;
}
.cu-bt-contact {
  margin-top: 24px;
  display: flex;
  align-items: center;
  gap: 12px;
  background: #1f3a2e;
  border-radius: 6px;
  padding: 15px 15px;
  color: #e9dfcc;
}
.cu-bt-desk {
  flex: 1;
  min-width: 0;
}
.cu-bt-desk p {
  font-family: 'Courier New', monospace;
  font-size: 7px;
  letter-spacing: 0.22em;
  color: #c9b98f;
}
.cu-bt-desk span {
  display: block;
  margin-top: 3px;
  font-size: 10px;
  color: #b9c6b4;
}
.cu-bt-phone {
  display: inline-flex;
  align-items: center;
  gap: 7px;
  border: 1px solid rgba(233, 223, 204, 0.55);
  border-radius: 999px;
  padding: 9px 13px;
  font-size: 14px;
  color: #f4eddc;
  white-space: nowrap;
  transition: background-color 0.2s ease;
}
.cu-bt-phone:hover {
  background: rgba(233, 223, 204, 0.14);
}
.cu-bt-contact .cu-bt-empty {
  color: #b9c6b4;
  border-color: rgba(233, 223, 204, 0.3);
}
.cu-bt-foot {
  margin-top: 18px;
  display: flex;
  justify-content: space-between;
  padding: 0 4px;
  font-family: 'Courier New', monospace;
  font-size: 7px;
  letter-spacing: 0.14em;
  color: #7d8b7c;
}
@keyframes cu-bt-in {
  from {
    opacity: 0;
    transform: translateY(18px);
  }
  to {
    opacity: 1;
    transform: none;
  }
}
.cu-bt-in {
  animation: cu-bt-in 0.85s cubic-bezier(0.22, 1, 0.36, 1) both;
}
.cu-bt-d1 {
  animation-delay: 0.1s;
}
.cu-bt-d2 {
  animation-delay: 0.2s;
}
.cu-bt-d3 {
  animation-delay: 0.3s;
}
@media (prefers-reduced-motion: reduce) {
  .cu-bt-in {
    animation: none;
  }
  .cu-bt-item,
  .cu-bt-phone {
    transition: none;
  }
}
</style>
