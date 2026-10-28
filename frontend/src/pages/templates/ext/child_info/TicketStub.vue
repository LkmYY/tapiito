<!-- 求助信息卡 · 车票存根：票根造型与虚线齿孔，信息分区一目了然 -->
<template>
  <div class="ci-ticket min-h-screen">
    <div class="mx-auto w-full max-w-md px-5 pb-12 pt-10">
      <div class="ci-ticket-wrap">
        <div class="ci-ticket-main">
          <header class="ci-head">
            <span class="ci-brand"><Ticket :size="15" /> 平安返程票</span>
            <span class="ci-serial">NO. {{ nfcId || '—' }}</span>
          </header>

          <div class="ci-route">
            <div>
              <span class="ci-route-label">FROM / 当前位置</span>
              <strong>需要帮助</strong>
            </div>
            <span class="ci-route-arrow" aria-hidden="true">→</span>
            <div class="ci-route-right">
              <span class="ci-route-label">TO / 目的地</span>
              <strong>家人身边</strong>
            </div>
          </div>

          <div class="ci-perf" aria-hidden="true"></div>

          <div class="ci-body">
            <span class="ci-label">基本情况 / PASSENGER</span>
            <p class="ci-name">{{ payload.baseInfo }}</p>

            <div class="ci-grid">
              <div>
                <span>学校 / 单位</span>
                <p>{{ payload.school }}</p>
              </div>
              <div>
                <span>过敏史 / 病史</span>
                <p class="ci-alert">{{ payload.allergies || '无' }}</p>
              </div>
            </div>
          </div>
        </div>

        <div class="ci-stub">
          <span class="ci-stub-label">紧急联系电话<br />EMERGENCY CALL</span>
          <a
            v-if="payload.emergencyPhone"
            :href="'tel:' + payload.emergencyPhone"
            class="ci-call"
          >
            <span class="ci-call-num">{{ payload.emergencyPhone }}</span>
            <span class="ci-call-cta"><Phone :size="15" /> 点击拨打</span>
          </a>

          <div class="ci-barcode" aria-hidden="true"></div>
          <p class="ci-stub-foot">请陪伴至家人到达<br />THANK YOU FOR YOUR CARE</p>
        </div>
      </div>

      <p class="ci-foot">此票不作乘车使用 · 仅用于紧急联系</p>
    </div>
  </div>
</template>

<script setup>
import { Phone, Ticket } from 'lucide-vue-next'

defineProps(['payload', 'nfcId'])
</script>

<style scoped>
.ci-ticket {
  background: #e8e2d4;
  color: #201c16;
  font-family: "Bahnschrift", "PingFang SC", "Microsoft YaHei", sans-serif;
}
.ci-ticket-wrap {
  display: flex;
  overflow: hidden;
  border-radius: 4px;
  background: #fbf7ec;
  box-shadow: 0 24px 44px -26px rgba(52, 42, 24, 0.75);
  animation: ci-ticket-in 0.6s ease both;
}
.ci-ticket-main {
  min-width: 0;
  flex: 1;
  padding: 18px 17px 16px;
}
.ci-head {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 8px;
  border-bottom: 1px solid #d8cfba;
  padding-bottom: 8px;
}
.ci-brand {
  display: inline-flex;
  align-items: center;
  gap: 5px;
  font-size: 10px;
  font-weight: 700;
  letter-spacing: 0.08em;
}
.ci-serial {
  font-family: ui-monospace, Consolas, monospace;
  font-size: 8px;
  color: #877a61;
}
.ci-route {
  display: flex;
  align-items: flex-end;
  gap: 9px;
  padding: 14px 0 12px;
}
.ci-route-label {
  display: block;
  font-family: ui-monospace, Consolas, monospace;
  font-size: 7px;
  letter-spacing: 0.06em;
  color: #8d7f65;
}
.ci-route strong {
  display: block;
  margin-top: 3px;
  font-size: 15px;
  font-weight: 800;
  line-height: 1.3;
}
.ci-route-arrow {
  padding-bottom: 2px;
  font-size: 19px;
  color: #b45309;
}
.ci-route-right { text-align: right; }
.ci-perf {
  height: 1px;
  background: repeating-linear-gradient(90deg, #bfb299 0 4px, transparent 4px 8px);
}
.ci-body { padding-top: 13px; }
.ci-label {
  display: block;
  font-family: ui-monospace, Consolas, monospace;
  font-size: 7.5px;
  letter-spacing: 0.1em;
  color: #8d7f65;
}
.ci-name {
  margin-top: 5px;
  font-size: 17px;
  font-weight: 800;
  line-height: 1.5;
}
.ci-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 9px;
  margin-top: 12px;
  border-top: 1px solid #e4dccb;
  padding-top: 9px;
}
.ci-grid span {
  display: block;
  font-size: 8px;
  color: #8d7f65;
}
.ci-grid p {
  margin-top: 3px;
  font-size: 11px;
  font-weight: 700;
  line-height: 1.5;
}
.ci-grid .ci-alert { color: #a33a12; }
.ci-stub {
  display: flex;
  width: 45%;
  flex-shrink: 0;
  flex-direction: column;
  border-left: 1px dashed #bfb299;
  background: #f3ecdb;
  padding: 18px 13px 15px;
  animation: ci-ticket-in 0.6s 0.1s ease both;
}
.ci-stub-label {
  font-family: ui-monospace, Consolas, monospace;
  font-size: 7.5px;
  line-height: 1.6;
  letter-spacing: 0.06em;
  color: #8d7f65;
}
.ci-call {
  display: block;
  margin-top: 10px;
  border-radius: 3px;
  background: #201c16;
  padding: 11px 10px;
  color: #fbf7ec;
  text-decoration: none;
  transition: transform 0.18s ease, background 0.18s ease;
}
.ci-call:hover { background: #3a3329; }
.ci-call:active { transform: scale(0.98); }
.ci-call-num {
  display: block;
  font-size: 25px;
  font-weight: 800;
  line-height: 1.3;
  word-break: break-all;
  font-variant-numeric: tabular-nums;
}
.ci-call-cta {
  display: flex;
  align-items: center;
  gap: 4px;
  margin-top: 5px;
  font-size: 8px;
  color: #d8c8a8;
}
.ci-empty {
  margin-top: 10px;
  border: 1px dashed #a89678;
  padding: 10px;
  font-size: 10px;
  color: #7d6e55;
}
.ci-barcode {
  height: 26px;
  margin-top: auto;
  padding-top: 16px;
  background: repeating-linear-gradient(90deg, #201c16 0 1px, transparent 1px 3px, #201c16 3px 5px, transparent 5px 7px);
  opacity: 0.7;
}
.ci-stub-foot {
  margin-top: 8px;
  font-family: ui-monospace, Consolas, monospace;
  font-size: 6.5px;
  line-height: 1.7;
  letter-spacing: 0.03em;
  color: #8d7f65;
}
.ci-foot {
  margin-top: 13px;
  text-align: center;
  font-family: ui-monospace, Consolas, monospace;
  font-size: 8px;
  letter-spacing: 0.05em;
  color: #8a7c63;
  animation: ci-ticket-in 0.6s 0.2s ease both;
}
@media (max-width: 360px) {
  .ci-route { gap: 5px; }
  .ci-route strong { font-size: 11px; }
  .ci-route-arrow { font-size: 15px; }
  .ci-route-label { font-size: 6px; }
  .ci-call-num { font-size: 23px; }
}
@keyframes ci-ticket-in {
  from { opacity: 0; transform: translateY(14px); }
  to { opacity: 1; transform: none; }
}
@media (prefers-reduced-motion: reduce) {
  .ci-ticket-wrap, .ci-stub, .ci-foot { animation: none; }
  .ci-call { transition: none; }
}
</style>
