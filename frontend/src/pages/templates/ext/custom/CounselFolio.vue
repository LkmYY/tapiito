<template>
  <div class="cu-cf-root min-h-screen">
    <div class="cu-cf-wrap mx-auto w-full max-w-md">
      <header class="cu-cf-top cu-cf-in">
        <span>CONFIDENTIAL BRIEF</span>
        <span>REF. {{ nfcId ? 'CF-' + nfcId : 'CF-000' }}</span>
      </header>

      <section class="cu-cf-doc cu-cf-in cu-cf-d1">
        <div class="cu-cf-doc-head">
          <div class="cu-cf-seal">{{ initial }}</div>
          <div class="cu-cf-head-copy">
            <p>COUNSEL &amp; ADVISORY</p>
            <span>二〇二六年 · 顾问备忘</span>
          </div>
        </div>

        <h1 class="cu-cf-name">{{ payload?.companyName || '企业主页' }}</h1>
        <div class="cu-cf-double"></div>

        <p class="cu-cf-desc cu-cf-in cu-cf-d2">
          <span class="cu-cf-dropcap">{{ dropcap }}</span>{{ payload?.description }}
        </p>

        <div class="cu-cf-meta cu-cf-in cu-cf-d2">
          <div><span>PREPARED FOR</span><p>来访客户</p></div>
          <div><span>ISSUED BY</span><p>{{ payload?.companyName || '企业主页' }}</p></div>
          <div><span>STATUS</span><p>有效 / VALID</p></div>
        </div>
      </section>

      <section class="cu-cf-clauses cu-cf-in cu-cf-d3">
        <div class="cu-cf-clause-head">
          <h2>§ 服务范围</h2>
          <span>SCOPE OF SERVICES</span>
        </div>
        <div v-if="services.length" class="cu-cf-clause-list">
          <article v-for="(service, i) in services" :key="i" class="cu-cf-clause">
            <span class="cu-cf-clause-no">{{ String(i + 1).padStart(2, '0') }}</span>
            <div>
              <h3>{{ service.name }}</h3>
              <p>由本所专业团队承办，具体安排以书面约定为准。</p>
            </div>
          </article>
        </div>
      </section>

      <section class="cu-cf-contact cu-cf-in cu-cf-d4">
        <p class="cu-cf-contact-label">联系方式 / CORRESPONDENCE</p>
        <a v-if="payload?.phone" :href="'tel:' + payload.phone" class="cu-cf-phone">
          <Phone :size="15" />
          <span>{{ payload.phone }}</span>
        </a>
        <a v-if="payload?.email" :href="'mailto:' + payload.email" class="cu-cf-mail">{{ payload.email }}</a>
        <div class="cu-cf-sign">
          <span></span>
          <p>{{ payload?.companyName || '企业主页' }}</p>
        </div>
      </section>

      <footer class="cu-cf-foot cu-cf-in cu-cf-d4">
        <span>此页为正式企业资料</span>
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
const initial = computed(() => (props.payload?.companyName || '未').trim().slice(0, 1))
const dropcap = computed(() => (props.payload?.description || '业').trim().slice(0, 1))
</script>

<style scoped>
.cu-cf-root {
  background: #f4efe2;
  color: #1c2b3a;
  font-family: Georgia, 'Songti SC', 'Noto Serif SC', serif;
  overflow-x: hidden;
}
.cu-cf-root::before {
  content: '';
  position: fixed;
  inset: 0;
  pointer-events: none;
  background-image: radial-gradient(rgba(28, 43, 58, 0.05) 0.6px, transparent 0.6px);
  background-size: 5px 5px;
}
.cu-cf-wrap {
  position: relative;
  padding: 20px 18px 30px;
}
.cu-cf-top {
  display: flex;
  justify-content: space-between;
  border-bottom: 1px solid rgba(28, 43, 58, 0.35);
  padding-bottom: 7px;
  font-family: 'Courier New', monospace;
  font-size: 7px;
  letter-spacing: 0.2em;
  color: #6e6a5c;
}
.cu-cf-doc {
  margin-top: 22px;
  border: 1px solid rgba(28, 43, 58, 0.3);
  background: #fbf7ec;
  padding: 20px 18px;
  box-shadow: 7px 8px 0 rgba(28, 43, 58, 0.06);
}
.cu-cf-doc-head {
  display: flex;
  align-items: center;
  gap: 12px;
}
.cu-cf-seal {
  height: 48px;
  width: 48px;
  flex-shrink: 0;
  display: flex;
  align-items: center;
  justify-content: center;
  border: 1px solid #1c2b3a;
  border-radius: 50%;
  font-size: 22px;
  color: #1c2b3a;
}
.cu-cf-head-copy p {
  font-family: 'Courier New', monospace;
  font-size: 8px;
  letter-spacing: 0.22em;
  color: #1c2b3a;
}
.cu-cf-head-copy span {
  display: block;
  margin-top: 3px;
  font-size: 10px;
  letter-spacing: 0.14em;
  color: #6e6a5c;
}
.cu-cf-name {
  margin-top: 22px;
  font-size: 40px;
  line-height: 1.1;
  font-weight: 400;
  letter-spacing: 0.02em;
  word-break: break-word;
}
.cu-cf-double {
  margin-top: 14px;
  border-top: 3px double rgba(28, 43, 58, 0.55);
}
.cu-cf-desc {
  margin-top: 16px;
  font-size: 13px;
  line-height: 1.95;
  color: #3b4754;
  text-align: justify;
}
.cu-cf-dropcap {
  float: left;
  margin: 3px 7px 0 0;
  font-size: 45px;
  line-height: 0.78;
  color: #8b2f2f;
  font-weight: 400;
}
.cu-cf-meta {
  margin-top: 18px;
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 11px;
  border-top: 1px solid rgba(28, 43, 58, 0.2);
  padding-top: 13px;
}
.cu-cf-meta div:last-child {
  grid-column: span 2;
}
.cu-cf-meta span {
  font-family: 'Courier New', monospace;
  font-size: 6px;
  letter-spacing: 0.16em;
  color: #8b877a;
}
.cu-cf-meta p {
  margin-top: 3px;
  font-size: 11px;
  color: #1c2b3a;
}
.cu-cf-clauses {
  margin-top: 24px;
}
.cu-cf-clause-head {
  display: flex;
  align-items: baseline;
  justify-content: space-between;
  border-bottom: 2px solid #1c2b3a;
  padding-bottom: 6px;
}
.cu-cf-clause-head h2 {
  font-size: 15px;
  font-weight: 600;
  letter-spacing: 0.04em;
}
.cu-cf-clause-head span {
  font-family: 'Courier New', monospace;
  font-size: 7px;
  letter-spacing: 0.14em;
  color: #6e6a5c;
}
.cu-cf-clause {
  display: flex;
  gap: 13px;
  padding: 15px 0;
  border-bottom: 1px solid rgba(28, 43, 58, 0.18);
}
.cu-cf-clause-no {
  flex-shrink: 0;
  font-family: 'Courier New', monospace;
  font-size: 9px;
  color: #8b2f2f;
  padding-top: 3px;
}
.cu-cf-clause h3 {
  font-size: 16px;
  font-weight: 600;
  letter-spacing: 0.02em;
}
.cu-cf-clause p {
  margin-top: 4px;
  font-size: 10px;
  line-height: 1.7;
  color: #7d786b;
}
.cu-cf-empty {
  padding: 16px 0;
  font-size: 12px;
  color: #7d786b;
}
.cu-cf-contact {
  margin-top: 24px;
  border-top: 1px solid rgba(28, 43, 58, 0.3);
  padding-top: 15px;
}
.cu-cf-contact-label {
  font-family: 'Courier New', monospace;
  font-size: 7px;
  letter-spacing: 0.18em;
  color: #6e6a5c;
}
.cu-cf-phone {
  margin-top: 8px;
  display: inline-flex;
  align-items: center;
  gap: 7px;
  font-size: 20px;
  letter-spacing: 0.03em;
  color: #1c2b3a;
  border-bottom: 1px solid rgba(139, 47, 47, 0.6);
  padding-bottom: 3px;
  transition: color 0.2s ease;
}
.cu-cf-phone:hover {
  color: #8b2f2f;
}
.cu-cf-mail,
.cu-cf-muted {
  display: block;
  margin-top: 8px;
  font-family: 'Courier New', monospace;
  font-size: 9px;
  color: #7d786b;
  word-break: break-all;
}
.cu-cf-sign {
  margin-top: 18px;
  display: flex;
  align-items: flex-end;
  justify-content: flex-end;
  gap: 8px;
}
.cu-cf-sign span {
  height: 1px;
  width: 56px;
  background: rgba(28, 43, 58, 0.4);
  margin-bottom: 5px;
}
.cu-cf-sign p {
  font-size: 12px;
  letter-spacing: 0.08em;
  color: #1c2b3a;
}
.cu-cf-foot {
  margin-top: 18px;
  display: flex;
  justify-content: space-between;
  font-family: 'Courier New', monospace;
  font-size: 7px;
  letter-spacing: 0.12em;
  color: #8b877a;
}
@keyframes cu-cf-in {
  from {
    opacity: 0;
    transform: translateY(14px);
  }
  to {
    opacity: 1;
    transform: none;
  }
}
.cu-cf-in {
  animation: cu-cf-in 0.85s cubic-bezier(0.22, 1, 0.36, 1) both;
}
.cu-cf-d1 {
  animation-delay: 0.08s;
}
.cu-cf-d2 {
  animation-delay: 0.17s;
}
.cu-cf-d3 {
  animation-delay: 0.26s;
}
.cu-cf-d4 {
  animation-delay: 0.35s;
}
@media (prefers-reduced-motion: reduce) {
  .cu-cf-in {
    animation: none;
  }
  .cu-cf-phone {
    transition: none;
  }
}
</style>
