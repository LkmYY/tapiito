<!-- 求助信息卡 · 洁净诊室：白底青绿，像一张清晰的病历卡 -->
<template>
  <div class="ci-clinic min-h-screen">
    <div class="mx-auto w-full max-w-md px-6 pb-14 pt-9">
      <div class="ci-top">
        <span class="ci-brand"><Cross :size="14" /> 安心求助卡</span>
        <span v-if="nfcId" class="ci-meta">NO. {{ nfcId }}</span>
      </div>

      <header class="ci-head">
        <span class="ci-badge"><HeartPulse :size="19" /></span>
        <h1>求助信息卡</h1>
        <p>如果遇见需要帮助的人，请先拨打电话，陪他/她等候家人。</p>
      </header>

      <section class="ci-card">
        <p class="ci-label">基本情况</p>
        <p class="ci-name">{{ payload.baseInfo }}</p>

        <div class="ci-line"></div>

        <div class="ci-row">
          <span class="ci-key">学校 / 单位</span>
          <span class="ci-val">{{ payload.school }}</span>
        </div>
        <div class="ci-row">
          <span class="ci-key">过敏史 / 病史</span>
          <span class="ci-val ci-val--alert">{{ payload.allergies || '无' }}</span>
        </div>
      </section>

      <a
        v-if="payload.emergencyPhone"
        :href="'tel:' + payload.emergencyPhone"
        class="ci-call"
      >
        <span class="ci-call-copy">
          <span class="ci-call-label">紧急联系电话 · 点击拨打</span>
          <span class="ci-call-num">{{ payload.emergencyPhone }}</span>
        </span>
        <span class="ci-call-icon"><Phone :size="23" /></span>
      </a>

      <p class="ci-foot"><ShieldCheck :size="14" /> 您的每一次停留，都是最及时的帮助</p>
    </div>
  </div>
</template>

<script setup>
import { Cross, HeartPulse, Phone, ShieldCheck } from 'lucide-vue-next'

defineProps(['payload', 'nfcId'])
</script>

<style scoped>
.ci-clinic {
  background: #f7faf9;
  color: #12332e;
  font-family: "Helvetica Neue", "PingFang SC", "Microsoft YaHei", sans-serif;
}
.ci-top {
  display: flex;
  align-items: center;
  justify-content: space-between;
  color: #6a8f88;
}
.ci-brand {
  display: flex;
  align-items: center;
  gap: 6px;
  font-size: 11px;
  letter-spacing: 0.2em;
}
.ci-meta {
  font-family: ui-monospace, Consolas, monospace;
  font-size: 10px;
  letter-spacing: 0.12em;
}
.ci-head {
  margin-top: 36px;
  animation: ci-clinic-in 0.6s ease both;
}
.ci-badge {
  display: inline-flex;
  height: 42px;
  width: 42px;
  align-items: center;
  justify-content: center;
  border-radius: 12px;
  background: #e3f3ee;
  color: #0f766e;
  box-shadow: inset 0 0 0 1px rgba(15, 118, 110, 0.14);
}
.ci-head h1 {
  margin-top: 16px;
  font-size: 27px;
  font-weight: 700;
  letter-spacing: -0.02em;
}
.ci-head p {
  margin-top: 8px;
  max-width: 21em;
  font-size: 13px;
  line-height: 1.7;
  color: #60827c;
}
.ci-card {
  margin-top: 26px;
  border-radius: 20px;
  background: #ffffff;
  padding: 24px;
  box-shadow: 0 22px 44px -34px rgba(15, 78, 68, 0.55), inset 0 0 0 1px rgba(15, 118, 110, 0.08);
  animation: ci-clinic-in 0.6s 0.08s ease both;
}
.ci-label {
  font-size: 10px;
  font-weight: 600;
  letter-spacing: 0.22em;
  color: #7d9f99;
}
.ci-name {
  margin-top: 10px;
  font-size: 22px;
  font-weight: 650;
  line-height: 1.5;
  color: #0c2f2a;
}
.ci-line {
  margin: 20px 0 4px;
  height: 1px;
  background: linear-gradient(90deg, #cfe4df, transparent);
}
.ci-row {
  display: flex;
  align-items: flex-start;
  justify-content: space-between;
  gap: 16px;
  padding: 13px 0;
  border-bottom: 1px solid #edf4f2;
}
.ci-key {
  flex-shrink: 0;
  font-size: 12px;
  color: #77958f;
}
.ci-val {
  max-width: 62%;
  text-align: right;
  font-size: 14px;
  font-weight: 600;
  line-height: 1.5;
  color: #1d3b36;
}
.ci-val--alert {
  color: #b45309;
}
.ci-call {
  margin-top: 22px;
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 14px;
  border-radius: 18px;
  background: #0f766e;
  padding: 22px;
  color: #ffffff;
  text-decoration: none;
  box-shadow: 0 20px 32px -22px rgba(15, 118, 110, 0.8);
  transition: transform 0.2s ease, background 0.2s ease, box-shadow 0.2s ease;
  animation: ci-clinic-in 0.6s 0.16s ease both;
}
.ci-call:hover {
  background: #0b655e;
  box-shadow: 0 24px 38px -20px rgba(15, 118, 110, 0.9);
}
.ci-call:active {
  transform: scale(0.98);
}
.ci-call-label {
  display: block;
  font-size: 10px;
  letter-spacing: 0.14em;
  color: #b7e5dc;
}
.ci-call-num {
  display: block;
  margin-top: 7px;
  font-size: 29px;
  font-weight: 700;
  letter-spacing: 0.01em;
  line-height: 1.2;
  word-break: break-all;
  font-variant-numeric: tabular-nums;
}
.ci-call-icon {
  display: flex;
  height: 44px;
  width: 44px;
  flex-shrink: 0;
  align-items: center;
  justify-content: center;
  border-radius: 50%;
  background: rgba(255, 255, 255, 0.16);
}
.ci-empty {
  margin-top: 22px;
  border: 1px dashed #a9c9c2;
  border-radius: 18px;
  padding: 20px;
  font-size: 14px;
  color: #77958f;
}
.ci-foot {
  margin-top: 22px;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 7px;
  font-size: 11px;
  color: #7c9a94;
  animation: ci-clinic-in 0.6s 0.24s ease both;
}
@keyframes ci-clinic-in {
  from { opacity: 0; transform: translateY(13px); }
  to { opacity: 1; transform: none; }
}
@media (prefers-reduced-motion: reduce) {
  .ci-head, .ci-card, .ci-call, .ci-foot { animation: none; }
  .ci-call { transition: none; }
}
</style>
