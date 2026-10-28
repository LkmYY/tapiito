<template>
  <div class="cu-fc-root min-h-screen">
    <div class="cu-fc-wrap mx-auto w-full max-w-md">
      <div class="cu-fc-sprocket cu-fc-in">
        <span v-for="n in 14" :key="n"></span>
      </div>

      <header class="cu-fc-top cu-fc-in">
        <span>KODAK / 400TX</span>
        <span>ROLL {{ nfcId ? String(nfcId).slice(-2) : '07' }} · 2026</span>
      </header>

      <section class="cu-fc-title">
        <p class="cu-fc-overline cu-fc-in cu-fc-d1">STUDIO CONTACT SHEET</p>
        <h1 class="cu-fc-name cu-fc-in cu-fc-d1">{{ payload?.companyName || '企业主页' }}</h1>
        <p class="cu-fc-desc cu-fc-in cu-fc-d2">{{ payload?.description }}</p>
      </section>

      <section class="cu-fc-sheet cu-fc-in cu-fc-d2">
        <div class="cu-fc-sheet-head">
          <span>FRAME 01—04</span>
          <span>ISO 400 / 35MM</span>
        </div>
        <div class="cu-fc-frames">
          <div v-for="i in 4" :key="i" class="cu-fc-frame" :class="{ 'cu-fc-frame--filled': images[i - 1] }">
            <img v-if="images[i - 1]" :src="images[i - 1]" :alt="payload?.companyName || '作品 ' + i" />
            <div v-else class="cu-fc-frame-ghost">
              <span>{{ i === 1 ? initial : '·' }}</span>
              <i>{{ String(i).padStart(2, '0') }}A</i>
            </div>
            <span class="cu-fc-frame-no">{{ String(i).padStart(2, '0') }}A</span>
            <span class="cu-fc-frame-notch"></span>
          </div>
        </div>
        <div class="cu-fc-sheet-foot">
          <span>▷ 12 EXPOSURES</span>
          <span class="cu-fc-rec"><i></i>REC</span>
        </div>
      </section>

      <section class="cu-fc-shots cu-fc-in cu-fc-d3">
        <div class="cu-fc-section-head">
          <h2>SHOT LIST</h2>
          <span>服务项目 / {{ String(services.length).padStart(2, '0') }}</span>
        </div>
        <div v-if="services.length" class="cu-fc-shot-list">
          <div v-for="(service, i) in services" :key="i" class="cu-fc-shot">
            <span class="cu-fc-shot-tc">00:{{ String(i + 1).padStart(2, '0') }}:12</span>
            <span class="cu-fc-shot-name">{{ service.name }}</span>
            <span class="cu-fc-shot-take">TAKE {{ i + 1 }}</span>
          </div>
        </div>
      </section>

      <section class="cu-fc-contact cu-fc-in cu-fc-d4">
        <div class="cu-fc-contact-label">BOOK THE STUDIO</div>
        <a v-if="payload?.phone" :href="'tel:' + payload.phone" class="cu-fc-phone">
          <Phone :size="16" />
          <span>{{ payload.phone }}</span>
          <ArrowUpRight :size="15" />
        </a>
        <a v-if="payload?.email" :href="'mailto:' + payload.email" class="cu-fc-mail">{{ payload.email }}</a>
      </section>

      <footer class="cu-fc-foot cu-fc-in cu-fc-d4">
        <span>{{ payload?.companyName || '企业主页' }}</span>
        <span>{{ nfcId ? 'NFC ' + nfcId : 'NFC —' }}</span>
      </footer>

      <div class="cu-fc-sprocket cu-fc-sprocket--bottom cu-fc-in">
        <span v-for="n in 14" :key="n"></span>
      </div>
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
const images = computed(() => {
  const list = props.payload?.images
  return Array.isArray(list) ? list.filter((src) => typeof src === 'string' && src.trim()) : []
})
const initial = computed(() => (props.payload?.companyName || '未').trim().slice(0, 1))
</script>

<style scoped>
.cu-fc-root {
  background: #080808;
  color: #e8e6df;
  font-family: 'Courier New', 'PingFang SC', monospace;
  overflow-x: hidden;
}
.cu-fc-wrap {
  padding: 10px 14px 18px;
}
.cu-fc-sprocket {
  display: flex;
  justify-content: space-between;
  padding: 0 2px;
}
.cu-fc-sprocket span {
  height: 9px;
  width: 12px;
  background: #e8e6df;
  opacity: 0.16;
  border-radius: 1px;
}
.cu-fc-sprocket--bottom {
  margin-top: 22px;
}
.cu-fc-top {
  display: flex;
  justify-content: space-between;
  margin-top: 9px;
  padding: 6px 8px;
  background: #cf3a2a;
  color: #fff8ef;
  font-size: 8px;
  letter-spacing: 0.18em;
}
.cu-fc-title {
  padding: 26px 4px 0;
}
.cu-fc-overline {
  font-size: 8px;
  letter-spacing: 0.28em;
  color: #8b8a84;
}
.cu-fc-name {
  margin-top: 9px;
  font-family: 'PingFang SC', 'Microsoft YaHei', sans-serif;
  font-size: 42px;
  font-weight: 800;
  line-height: 1;
  letter-spacing: -0.05em;
  color: #f3f1e9;
  word-break: break-word;
}
.cu-fc-desc {
  margin-top: 13px;
  max-width: 36ch;
  font-family: 'PingFang SC', sans-serif;
  font-size: 12px;
  line-height: 1.8;
  color: #a09e96;
}
.cu-fc-sheet {
  margin-top: 24px;
  border: 1px solid #3a3a36;
  padding: 8px;
  background: #101010;
}
.cu-fc-sheet-head,
.cu-fc-sheet-foot {
  display: flex;
  justify-content: space-between;
  font-size: 7px;
  letter-spacing: 0.16em;
  color: #8b8a84;
  padding: 3px 1px;
}
.cu-fc-frames {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 6px;
  margin: 5px 0;
}
.cu-fc-frame {
  position: relative;
  aspect-ratio: 4 / 3;
  overflow: hidden;
  background: #191919;
  border: 1px solid #2c2c29;
}
.cu-fc-frame img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  display: block;
  filter: grayscale(1) contrast(1.12);
  transition: transform 0.5s ease;
}
.cu-fc-frame:hover img {
  transform: scale(1.04);
}
.cu-fc-frame-ghost {
  position: absolute;
  inset: 0;
  display: flex;
  align-items: center;
  justify-content: center;
  background:
    linear-gradient(120deg, transparent 42%, rgba(255, 255, 255, 0.05) 42% 44%, transparent 44%),
    #171717;
}
.cu-fc-frame-ghost span {
  font-family: Georgia, serif;
  font-size: 52px;
  color: #2f2f2b;
  line-height: 1;
}
.cu-fc-frame-ghost i {
  position: absolute;
  right: 6px;
  top: 5px;
  font-style: normal;
  font-size: 6px;
  letter-spacing: 0.12em;
  color: #565650;
}
.cu-fc-frame-no {
  position: absolute;
  left: 5px;
  bottom: 4px;
  font-size: 6px;
  letter-spacing: 0.12em;
  color: #cf3a2a;
  background: rgba(8, 8, 8, 0.72);
  padding: 1px 3px;
}
.cu-fc-frame-notch {
  position: absolute;
  right: 4px;
  bottom: 4px;
  height: 3px;
  width: 16px;
  background: #cf3a2a;
  opacity: 0.8;
}
.cu-fc-rec {
  display: inline-flex;
  align-items: center;
  gap: 4px;
  color: #cf3a2a;
}
.cu-fc-rec i {
  height: 5px;
  width: 5px;
  border-radius: 50%;
  background: #cf3a2a;
  animation: cu-fc-blink 1.2s steps(2, end) infinite;
}
.cu-fc-shots {
  margin-top: 26px;
}
.cu-fc-section-head {
  display: flex;
  align-items: baseline;
  justify-content: space-between;
  border-bottom: 1px solid #3a3a36;
  padding-bottom: 6px;
}
.cu-fc-section-head h2 {
  font-size: 11px;
  letter-spacing: 0.2em;
  color: #e8e6df;
}
.cu-fc-section-head span {
  font-size: 7px;
  letter-spacing: 0.1em;
  color: #8b8a84;
}
.cu-fc-shot {
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 12px 2px;
  border-bottom: 1px solid #242421;
  transition: background-color 0.2s ease;
}
.cu-fc-shot:hover {
  background: #121212;
}
.cu-fc-shot-tc {
  font-size: 7px;
  letter-spacing: 0.08em;
  color: #cf3a2a;
}
.cu-fc-shot-name {
  flex: 1;
  font-family: 'PingFang SC', sans-serif;
  font-size: 14px;
  color: #e8e6df;
}
.cu-fc-shot-take {
  font-size: 7px;
  letter-spacing: 0.12em;
  color: #6f6e68;
}
.cu-fc-empty {
  margin-top: 12px;
  border: 1px dashed #3a3a36;
  padding: 14px;
  font-family: 'PingFang SC', sans-serif;
  font-size: 11px;
  color: #8b8a84;
}
.cu-fc-contact {
  margin-top: 24px;
  border-left: 3px solid #cf3a2a;
  padding: 4px 0 4px 13px;
}
.cu-fc-contact-label {
  font-size: 7px;
  letter-spacing: 0.22em;
  color: #8b8a84;
}
.cu-fc-phone {
  margin-top: 8px;
  display: inline-flex;
  align-items: center;
  gap: 8px;
  font-size: 21px;
  letter-spacing: 0.02em;
  color: #f3f1e9;
  transition: color 0.2s ease;
}
.cu-fc-phone:hover {
  color: #ff6a55;
}
.cu-fc-mail {
  display: block;
  margin-top: 8px;
  font-size: 9px;
  letter-spacing: 0.08em;
  color: #8b8a84;
  word-break: break-all;
}
.cu-fc-foot {
  margin-top: 20px;
  display: flex;
  justify-content: space-between;
  font-size: 7px;
  letter-spacing: 0.14em;
  color: #6f6e68;
}
@keyframes cu-fc-in {
  from {
    opacity: 0;
    transform: translateY(10px);
  }
  to {
    opacity: 1;
    transform: none;
  }
}
.cu-fc-in {
  animation: cu-fc-in 0.65s ease-out both;
}
.cu-fc-d1 {
  animation-delay: 0.08s;
}
.cu-fc-d2 {
  animation-delay: 0.17s;
}
.cu-fc-d3 {
  animation-delay: 0.26s;
}
.cu-fc-d4 {
  animation-delay: 0.35s;
}
@keyframes cu-fc-blink {
  50% {
    opacity: 0.2;
  }
}
@media (prefers-reduced-motion: reduce) {
  .cu-fc-in {
    animation: none;
  }
  .cu-fc-rec i {
    animation: none;
  }
  .cu-fc-frame img,
  .cu-fc-shot,
  .cu-fc-phone {
    transition: none;
  }
}
</style>
