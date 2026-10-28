<!-- 求助信息卡 · 云朵安睡：浅蓝紫的云朵卡，柔软安静，适合小朋友 -->
<template>
  <div class="ci-cloud min-h-screen">
    <div class="mx-auto w-full max-w-md px-6 pb-14 pt-9">
      <div class="ci-sky" aria-hidden="true">
        <span class="ci-star ci-star--one">✦</span>
        <span class="ci-star ci-star--two">✦</span>
      </div>

      <header class="ci-head">
        <span class="ci-brand"><Cloud :size="16" /> 小小的我</span>
        <span v-if="nfcId" class="ci-id">NFC {{ nfcId }}</span>
      </header>

      <h1 class="ci-title">我走丢了，<br />可以帮帮我吗？</h1>
      <p class="ci-sub">请不要着急，我的家人就在电话那头。<br />轻轻点一下下面的号码就好。</p>

      <div class="ci-puff ci-puff--one">
        <span class="ci-cloud-label">基本情况</span>
        <p class="ci-name">{{ payload.baseInfo }}</p>
      </div>

      <div class="ci-puff ci-puff--two">
        <div class="ci-row">
          <span><School :size="14" /> 学校 / 单位</span>
          <strong>{{ payload.school }}</strong>
        </div>
        <div class="ci-row">
          <span><HeartPulse :size="14" /> 过敏史 / 病史</span>
          <strong class="ci-warn">{{ payload.allergies || '无' }}</strong>
        </div>
      </div>

      <a
        v-if="payload.emergencyPhone"
        :href="'tel:' + payload.emergencyPhone"
        class="ci-call"
      >
        <span class="ci-call-icon"><Phone :size="22" /></span>
        <span class="ci-call-copy">
          <span class="ci-call-label">打给家人</span>
          <span class="ci-call-num">{{ payload.emergencyPhone }}</span>
        </span>
        <span class="ci-call-spark" aria-hidden="true">✦</span>
      </a>

      <p class="ci-foot">谢谢你陪我，等家人来接我。</p>
    </div>
  </div>
</template>

<script setup>
import { Cloud, HeartPulse, Phone, School } from 'lucide-vue-next'

defineProps(['payload', 'nfcId'])
</script>

<style scoped>
.ci-cloud {
  position: relative;
  overflow: hidden;
  background: linear-gradient(180deg, #dbeafe 0%, #e9e6ff 55%, #f3ecff 100%);
  color: #2c3459;
  font-family: "PingFang SC", "Microsoft YaHei", "Helvetica Neue", sans-serif;
}
.ci-sky {
  position: absolute;
  inset: 0;
  pointer-events: none;
}
.ci-star {
  position: absolute;
  color: #ffffff;
  opacity: 0.85;
}
.ci-star--one { top: 38px; right: 44px; font-size: 17px; }
.ci-star--two { top: 152px; left: 24px; font-size: 10px; }
.ci-head {
  position: relative;
  display: flex;
  align-items: center;
  justify-content: space-between;
  animation: ci-cloud-in 0.6s ease both;
}
.ci-brand {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  font-size: 11px;
  letter-spacing: 0.14em;
  color: #6675a4;
}
.ci-id {
  font-family: ui-monospace, Consolas, monospace;
  font-size: 8px;
  color: #8894bb;
}
.ci-title {
  position: relative;
  margin-top: 33px;
  font-size: 33px;
  font-weight: 700;
  line-height: 1.35;
  letter-spacing: -0.02em;
  color: #303a67;
  animation: ci-cloud-in 0.65s 0.07s ease both;
}
.ci-sub {
  position: relative;
  margin-top: 10px;
  font-size: 13px;
  line-height: 1.9;
  color: #6d789f;
  animation: ci-cloud-in 0.65s 0.12s ease both;
}
.ci-puff {
  position: relative;
  margin-top: 18px;
  border-radius: 32px;
  background: rgba(255, 255, 255, 0.78);
  padding: 18px 21px;
  box-shadow: 0 18px 34px -26px rgba(70, 84, 136, 0.6), inset 0 1px 0 #ffffff;
  backdrop-filter: blur(8px);
}
.ci-puff::before,
.ci-puff::after {
  content: "";
  position: absolute;
  border-radius: 50%;
  background: rgba(255, 255, 255, 0.78);
  z-index: -1;
}
.ci-puff::before {
  left: 25px;
  top: -11px;
  height: 34px;
  width: 58px;
}
.ci-puff::after {
  right: 27px;
  top: -7px;
  height: 26px;
  width: 40px;
}
.ci-puff--one {
  margin-right: 21px;
  animation: ci-cloud-in 0.65s 0.17s ease both;
}
.ci-puff--two {
  margin-left: 21px;
  padding-top: 7px;
  padding-bottom: 7px;
  animation: ci-cloud-in 0.65s 0.22s ease both;
}
.ci-cloud-label {
  font-size: 9px;
  font-weight: 600;
  letter-spacing: 0.16em;
  color: #8c98be;
}
.ci-name {
  margin-top: 4px;
  font-size: 20px;
  font-weight: 650;
  line-height: 1.55;
  color: #2c3459;
}
.ci-row {
  display: flex;
  align-items: flex-start;
  justify-content: space-between;
  gap: 12px;
  border-bottom: 1px solid rgba(120, 136, 180, 0.15);
  padding: 10px 0;
}
.ci-row:last-child { border-bottom: none; }
.ci-row span {
  display: inline-flex;
  align-items: center;
  gap: 5px;
  flex-shrink: 0;
  font-size: 10px;
  color: #7a88ae;
}
.ci-row strong {
  max-width: 60%;
  text-align: right;
  font-size: 12px;
  font-weight: 600;
  line-height: 1.55;
  color: #354064;
}
.ci-row .ci-warn { color: #a45a33; }
.ci-call {
  display: flex;
  align-items: center;
  gap: 12px;
  margin-top: 17px;
  border-radius: 28px;
  background: #46538f;
  padding: 16px 18px;
  color: #ffffff;
  text-decoration: none;
  box-shadow: 0 20px 32px -22px rgba(70, 83, 143, 0.85);
  transition: transform 0.18s ease, background 0.18s ease;
  animation: ci-cloud-in 0.65s 0.28s ease both;
}
.ci-call:hover { background: #3b477d; }
.ci-call:active { transform: scale(0.982); }
.ci-call-icon {
  display: flex;
  height: 41px;
  width: 41px;
  flex-shrink: 0;
  align-items: center;
  justify-content: center;
  border-radius: 50%;
  background: rgba(255, 255, 255, 0.16);
}
.ci-call-copy { min-width: 0; flex: 1; }
.ci-call-label {
  display: block;
  font-size: 10px;
  color: #c9d1f4;
}
.ci-call-num {
  display: block;
  margin-top: 2px;
  font-size: 27px;
  font-weight: 700;
  line-height: 1.18;
  word-break: break-all;
  font-variant-numeric: tabular-nums;
}
.ci-call-spark { color: #cdd6ff; font-size: 15px; }
.ci-empty {
  margin-top: 17px;
  border: 1px dashed #9aa5c9;
  border-radius: 28px;
  padding: 16px;
  font-size: 13px;
  color: #7180a8;
  text-align: center;
}
.ci-foot {
  position: relative;
  margin-top: 16px;
  text-align: center;
  font-size: 11px;
  color: #7d89ae;
  animation: ci-cloud-in 0.65s 0.33s ease both;
}
@keyframes ci-cloud-in {
  from { opacity: 0; transform: translateY(14px); }
  to { opacity: 1; transform: none; }
}
@media (prefers-reduced-motion: reduce) {
  .ci-head, .ci-title, .ci-sub, .ci-puff, .ci-call, .ci-foot { animation: none; }
  .ci-call { transition: none; }
}
</style>
