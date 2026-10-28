<template>
  <div class="cu-gm-root min-h-screen">
    <div class="cu-gm-wrap mx-auto w-full max-w-md">
      <div class="cu-gm-frame">
        <div class="cu-gm-corner cu-gm-corner--tl"></div>
        <div class="cu-gm-corner cu-gm-corner--tr"></div>
        <div class="cu-gm-corner cu-gm-corner--bl"></div>
        <div class="cu-gm-corner cu-gm-corner--br"></div>

        <header class="cu-gm-top cu-gm-in">
          <span>MAISON</span>
          <span>{{ nfcId ? 'N° ' + nfcId : 'N° 000' }}</span>
        </header>

        <div class="cu-gm-emblem cu-gm-in cu-gm-d1">
          <span class="cu-gm-emblem-line"></span>
          <span class="cu-gm-monogram">{{ initials }}</span>
          <span class="cu-gm-emblem-line"></span>
        </div>

        <section class="cu-gm-title cu-gm-in cu-gm-d2">
          <p class="cu-gm-overline">ATELIER &amp; ADVISORY</p>
          <h1>{{ payload?.companyName || '企业主页' }}</h1>
          <p class="cu-gm-desc">{{ payload?.description }}</p>
        </section>

        <div class="cu-gm-flourish cu-gm-in cu-gm-d2">
          <span></span><i>✦</i><span></span>
        </div>

        <section class="cu-gm-services cu-gm-in cu-gm-d3">
          <p class="cu-gm-section-label">SAVOIR-FAIRE</p>
          <div v-if="services.length" class="cu-gm-list">
            <div v-for="(service, i) in services" :key="i" class="cu-gm-service">
              <span class="cu-gm-service-no">{{ roman[i % roman.length] }}</span>
              <h2>{{ service.name }}</h2>
            </div>
          </div>
        </section>

        <section class="cu-gm-contact cu-gm-in cu-gm-d4">
          <p class="cu-gm-section-label">PRIVATE ENQUIRIES</p>
          <a v-if="payload?.phone" :href="'tel:' + payload.phone" class="cu-gm-phone">
            <Phone :size="14" />
            {{ payload.phone }}
            <ArrowUpRight :size="14" />
          </a>
          <a v-if="payload?.email" :href="'mailto:' + payload.email" class="cu-gm-mail">{{ payload.email }}</a>
        </section>

        <footer class="cu-gm-foot cu-gm-in cu-gm-d4">
          <span>PARIS · SHANGHAI</span>
          <span>{{ payload?.companyName || '企业主页' }}</span>
        </footer>
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { ArrowUpRight, Phone } from 'lucide-vue-next'

const props = defineProps(['payload', 'nfcId'])
const roman = ['I', 'II', 'III', 'IV', 'V']

const services = computed(() => {
  const list = props.payload?.services
  return Array.isArray(list) ? list.filter((item) => item && item.name) : []
})
const initials = computed(() => {
  const name = (props.payload?.companyName || '未').trim()
  return name.slice(0, 2)
})
</script>

<style scoped>
.cu-gm-root {
  background: #0b0a08;
  color: #e8ddc4;
  font-family: 'Times New Roman', Georgia, 'Songti SC', serif;
  overflow-x: hidden;
}
.cu-gm-root::before {
  content: '';
  position: fixed;
  inset: 0;
  pointer-events: none;
  background:
    radial-gradient(ellipse at 50% 0%, rgba(201, 162, 39, 0.14), transparent 55%),
    radial-gradient(ellipse at 50% 100%, rgba(201, 162, 39, 0.07), transparent 50%);
}
.cu-gm-wrap {
  position: relative;
  padding: 16px 13px 28px;
}
.cu-gm-frame {
  position: relative;
  border: 1px solid rgba(201, 162, 39, 0.55);
  padding: 14px 17px 20px;
  box-shadow: inset 0 0 0 3px #0b0a08, inset 0 0 0 4px rgba(201, 162, 39, 0.22);
}
.cu-gm-corner {
  position: absolute;
  height: 15px;
  width: 15px;
  border-color: #c9a227;
  border-style: solid;
}
.cu-gm-corner--tl {
  left: 8px;
  top: 8px;
  border-width: 1px 0 0 1px;
}
.cu-gm-corner--tr {
  right: 8px;
  top: 8px;
  border-width: 1px 1px 0 0;
}
.cu-gm-corner--bl {
  left: 8px;
  bottom: 8px;
  border-width: 0 0 1px 1px;
}
.cu-gm-corner--br {
  right: 8px;
  bottom: 8px;
  border-width: 0 1px 1px 0;
}
.cu-gm-top {
  display: flex;
  justify-content: space-between;
  font-family: 'Courier New', monospace;
  font-size: 7px;
  letter-spacing: 0.3em;
  color: #a68a4c;
}
.cu-gm-emblem {
  margin-top: 22px;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 12px;
}
.cu-gm-emblem-line {
  height: 1px;
  width: 38px;
  background: linear-gradient(90deg, transparent, rgba(201, 162, 39, 0.7), transparent);
}
.cu-gm-monogram {
  display: flex;
  height: 58px;
  width: 58px;
  align-items: center;
  justify-content: center;
  border: 1px solid rgba(201, 162, 39, 0.75);
  transform: rotate(45deg);
  font-size: 18px;
  letter-spacing: 0.1em;
  color: #d9b95f;
}
.cu-gm-monogram::first-letter {
  letter-spacing: 0;
}
.cu-gm-title {
  margin-top: 24px;
  text-align: center;
}
.cu-gm-overline {
  font-family: 'Courier New', monospace;
  font-size: 7px;
  letter-spacing: 0.34em;
  color: #a68a4c;
}
.cu-gm-title h1 {
  margin-top: 10px;
  font-size: 33px;
  line-height: 1.16;
  font-weight: 400;
  letter-spacing: 0.1em;
  color: #f1e6cd;
  word-break: break-word;
}
.cu-gm-desc {
  margin: 13px auto 0;
  max-width: 30ch;
  font-size: 12px;
  line-height: 1.9;
  font-style: italic;
  color: #a99c83;
}
.cu-gm-flourish {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 10px;
  margin: 18px 0 4px;
  color: #c9a227;
}
.cu-gm-flourish span {
  height: 1px;
  width: 46px;
  background: rgba(201, 162, 39, 0.38);
}
.cu-gm-flourish i {
  font-style: normal;
  font-size: 8px;
}
.cu-gm-services {
  margin-top: 16px;
}
.cu-gm-section-label {
  text-align: center;
  font-family: 'Courier New', monospace;
  font-size: 7px;
  letter-spacing: 0.32em;
  color: #a68a4c;
}
.cu-gm-list {
  margin-top: 10px;
  display: flex;
  flex-direction: column;
  gap: 7px;
}
.cu-gm-service {
  display: flex;
  align-items: center;
  gap: 12px;
  border: 1px solid rgba(201, 162, 39, 0.24);
  background: rgba(201, 162, 39, 0.03);
  padding: 11px 13px;
  transition: border-color 0.25s ease, background-color 0.25s ease;
}
.cu-gm-service:hover {
  border-color: rgba(201, 162, 39, 0.65);
  background: rgba(201, 162, 39, 0.08);
}
.cu-gm-service-no {
  font-size: 10px;
  letter-spacing: 0.14em;
  color: #c9a227;
  width: 24px;
  flex-shrink: 0;
}
.cu-gm-service h2 {
  font-family: 'Songti SC', Georgia, serif;
  font-size: 15px;
  font-weight: 400;
  letter-spacing: 0.08em;
  color: #eadfc7;
}
.cu-gm-contact {
  margin-top: 22px;
  border-top: 1px solid rgba(201, 162, 39, 0.3);
  padding-top: 15px;
  text-align: center;
}
.cu-gm-phone {
  margin-top: 10px;
  display: inline-flex;
  align-items: center;
  gap: 8px;
  border: 1px solid rgba(201, 162, 39, 0.65);
  padding: 10px 16px;
  font-size: 16px;
  letter-spacing: 0.08em;
  color: #e2c879;
  transition: background-color 0.25s ease, color 0.25s ease;
}
.cu-gm-phone:hover {
  background: #c9a227;
  color: #0b0a08;
}
.cu-gm-mail {
  display: block;
  margin-top: 9px;
  font-family: 'Courier New', monospace;
  font-size: 8px;
  letter-spacing: 0.12em;
  color: #a68a4c;
  word-break: break-all;
}
.cu-gm-muted {
  margin-top: 9px;
  font-size: 10px;
  color: #92866f;
  font-style: italic;
}
.cu-gm-foot {
  margin-top: 19px;
  display: flex;
  justify-content: space-between;
  font-family: 'Courier New', monospace;
  font-size: 6px;
  letter-spacing: 0.16em;
  color: #8a764a;
}
@keyframes cu-gm-in {
  from {
    opacity: 0;
    transform: translateY(12px);
  }
  to {
    opacity: 1;
    transform: none;
  }
}
.cu-gm-in {
  animation: cu-gm-in 0.9s cubic-bezier(0.22, 1, 0.36, 1) both;
}
.cu-gm-d1 {
  animation-delay: 0.1s;
}
.cu-gm-d2 {
  animation-delay: 0.2s;
}
.cu-gm-d3 {
  animation-delay: 0.3s;
}
.cu-gm-d4 {
  animation-delay: 0.4s;
}
@media (prefers-reduced-motion: reduce) {
  .cu-gm-in {
    animation: none;
  }
  .cu-gm-service,
  .cu-gm-phone {
    transition: none;
  }
}
</style>
