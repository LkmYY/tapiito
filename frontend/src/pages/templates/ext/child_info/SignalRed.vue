<!-- 求助信息卡 · 信号红带：白底红黑警示带，最大字号，远看也清楚 -->
<template>
  <div class="ci-signal min-h-screen">
    <div class="mx-auto w-full max-w-md pb-12">
      <div class="ci-hazard" aria-hidden="true"></div>

      <div class="ci-hero">
        <div class="ci-hero-top">
          <span class="ci-alert-chip"><TriangleAlert :size="14" /> 紧急求助</span>
          <span v-if="nfcId" class="ci-id">CARD {{ nfcId }}</span>
        </div>
        <h1>请帮这个<br />人回家</h1>
        <p>他需要您的帮助。请拨打电话联系家人，并在原地陪伴。</p>
      </div>

      <div class="ci-body">
        <section class="ci-name-block">
          <span class="ci-label">基本情况</span>
          <p class="ci-name">{{ payload.baseInfo }}</p>
        </section>

        <div class="ci-grid">
          <div>
            <span class="ci-label">学校 / 单位</span>
            <p>{{ payload.school }}</p>
          </div>
          <div>
            <span class="ci-label">过敏史 / 病史</span>
            <p class="ci-red">{{ payload.allergies || '无' }}</p>
          </div>
        </div>

        <a
          v-if="payload.emergencyPhone"
          :href="'tel:' + payload.emergencyPhone"
          class="ci-call"
        >
          <span class="ci-call-head">
            <span class="ci-call-label">紧急联系电话</span>
            <span class="ci-call-badge"><Phone :size="16" /> 立即拨打</span>
          </span>
          <span class="ci-call-num">{{ payload.emergencyPhone }}</span>
        </a>

        <p class="ci-foot"><HeartHandshake :size="15" /> 您的停留，就是他的安全感。</p>
      </div>

      <div class="ci-hazard ci-hazard--bottom" aria-hidden="true"></div>
    </div>
  </div>
</template>

<script setup>
import { HeartHandshake, Phone, TriangleAlert } from 'lucide-vue-next'

defineProps(['payload', 'nfcId'])
</script>

<style scoped>
.ci-signal {
  background: #ffffff;
  color: #171717;
  font-family: "Arial Black", "PingFang SC", "Microsoft YaHei", sans-serif;
}
.ci-hazard {
  height: 16px;
  background: repeating-linear-gradient(115deg, #e11d2e 0 14px, #ffffff 14px 28px);
  animation: ci-signal-in 0.4s ease both;
}
.ci-hazard--bottom { margin-top: 14px; }
.ci-hero {
  position: relative;
  background: #e11d2e;
  padding: 22px 22px 24px;
  color: #ffffff;
  animation: ci-signal-in 0.55s 0.04s ease both;
}
.ci-hero::after {
  content: "";
  position: absolute;
  right: -33px;
  top: -38px;
  height: 130px;
  width: 130px;
  border: 18px solid rgba(255, 255, 255, 0.13);
  border-radius: 50%;
}
.ci-hero-top {
  position: relative;
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 8px;
}
.ci-alert-chip {
  display: inline-flex;
  align-items: center;
  gap: 5px;
  border-radius: 3px;
  background: #171717;
  padding: 5px 8px;
  font-size: 9px;
  letter-spacing: 0.08em;
}
.ci-id {
  font-family: ui-monospace, Consolas, monospace;
  font-size: 8px;
  opacity: 0.8;
}
.ci-hero h1 {
  position: relative;
  margin-top: 20px;
  font-size: 46px;
  font-weight: 900;
  line-height: 1.08;
  letter-spacing: -0.045em;
}
.ci-hero p {
  position: relative;
  margin-top: 11px;
  max-width: 19em;
  font-family: "PingFang SC", "Microsoft YaHei", sans-serif;
  font-size: 14px;
  font-weight: 500;
  line-height: 1.7;
  color: #ffe1e4;
}
.ci-body {
  padding: 20px 22px 0;
}
.ci-name-block {
  border-bottom: 3px solid #171717;
  padding-bottom: 13px;
  animation: ci-signal-in 0.55s 0.1s ease both;
}
.ci-label {
  display: block;
  font-family: ui-monospace, Consolas, monospace;
  font-size: 8px;
  font-weight: 700;
  letter-spacing: 0.14em;
  color: #8c8c8c;
}
.ci-name {
  margin-top: 5px;
  font-size: 25px;
  font-weight: 900;
  line-height: 1.45;
  letter-spacing: -0.02em;
}
.ci-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 14px;
  border-bottom: 1px solid #d4d4d4;
  padding: 13px 0;
  animation: ci-signal-in 0.55s 0.15s ease both;
}
.ci-grid p {
  margin-top: 5px;
  font-size: 14px;
  font-weight: 700;
  line-height: 1.5;
}
.ci-grid .ci-red { color: #c41220; }
.ci-call {
  display: block;
  margin-top: 18px;
  background: #171717;
  padding: 16px 17px 15px;
  color: #ffffff;
  text-decoration: none;
  transition: transform 0.16s ease, background 0.16s ease;
  animation: ci-signal-in 0.55s 0.2s ease both;
}
.ci-call:hover { background: #2d2d2d; }
.ci-call:active { transform: scale(0.985); }
.ci-call-head {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 8px;
}
.ci-call-label {
  font-size: 10px;
  font-weight: 700;
  letter-spacing: 0.12em;
  color: #ff9a9f;
}
.ci-call-badge {
  display: inline-flex;
  align-items: center;
  gap: 4px;
  border-radius: 2px;
  background: #e11d2e;
  padding: 4px 7px;
  font-size: 9px;
  font-weight: 700;
}
.ci-call-num {
  display: block;
  margin-top: 8px;
  font-size: 36px;
  font-weight: 900;
  line-height: 1.12;
  letter-spacing: 0.005em;
  word-break: break-all;
  font-variant-numeric: tabular-nums;
  color: #ffffff;
}
.ci-empty {
  margin-top: 18px;
  border: 3px dashed #e11d2e;
  padding: 16px;
  font-size: 14px;
  font-weight: 700;
  color: #c41220;
}
.ci-foot {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 6px;
  margin-top: 15px;
  font-family: "PingFang SC", "Microsoft YaHei", sans-serif;
  font-size: 11px;
  font-weight: 500;
  color: #737373;
  animation: ci-signal-in 0.55s 0.26s ease both;
}
@keyframes ci-signal-in {
  from { opacity: 0; transform: translateY(13px); }
  to { opacity: 1; transform: none; }
}
@media (prefers-reduced-motion: reduce) {
  .ci-hazard, .ci-hero, .ci-name-block, .ci-grid, .ci-call, .ci-foot { animation: none; }
  .ci-call { transition: none; }
}
</style>
