<!-- 求助信息卡 · 玻璃关怀：柔光渐变与磨砂玻璃，轻盈但不失清晰 -->
<template>
  <div class="ci-glass min-h-screen">
    <div class="ci-orb ci-orb--one" aria-hidden="true"></div>
    <div class="ci-orb ci-orb--two" aria-hidden="true"></div>

    <div class="mx-auto w-full max-w-md px-6 pb-12 pt-9">
      <header class="ci-head">
        <span class="ci-brand"><HandHeart :size="15" /> 关怀卡</span>
        <span v-if="nfcId" class="ci-id">ID {{ nfcId }}</span>
      </header>

      <h1 class="ci-title">需要帮助时，<br />请拨这个电话</h1>
      <p class="ci-sub">如果您遇见需要协助的人，请先联系家人。我们会尽快赶到。</p>

      <div class="ci-stack">
        <section class="ci-panel ci-panel--profile">
          <div class="ci-panel-head">
            <span class="ci-dot"></span>
            <span>基本情况</span>
          </div>
          <p class="ci-name">{{ payload.baseInfo }}</p>
        </section>

        <section class="ci-panel ci-panel--rows">
          <div class="ci-row">
            <span>学校 / 单位</span>
            <strong>{{ payload.school }}</strong>
          </div>
          <div class="ci-row">
            <span>过敏史 / 病史</span>
            <strong>{{ payload.allergies || '无' }}</strong>
          </div>
        </section>
      </div>

      <a
        v-if="payload.emergencyPhone"
        :href="'tel:' + payload.emergencyPhone"
        class="ci-call"
      >
        <span class="ci-call-copy">
          <span class="ci-call-label"><Phone :size="16" /> 紧急联系电话</span>
          <span class="ci-call-num">{{ payload.emergencyPhone }}</span>
        </span>
        <span class="ci-call-btn">拨打</span>
      </a>

      <p class="ci-foot"><ShieldCheck :size="13" /> 信息仅用于本次求助 · 感谢您的善意</p>
    </div>
  </div>
</template>

<script setup>
import { HandHeart, Phone, ShieldCheck } from 'lucide-vue-next'

defineProps(['payload', 'nfcId'])
</script>

<style scoped>
.ci-glass {
  position: relative;
  overflow: hidden;
  background: linear-gradient(150deg, #e7f0ff 0%, #eef3fb 43%, #f8ecf4 100%);
  color: #24364f;
  font-family: "PingFang SC", "Microsoft YaHei", "Helvetica Neue", sans-serif;
}
.ci-orb {
  position: absolute;
  border-radius: 50%;
  filter: blur(2px);
  pointer-events: none;
}
.ci-orb--one {
  right: -60px;
  top: 34px;
  height: 220px;
  width: 220px;
  background: radial-gradient(circle at 32% 32%, rgba(125, 211, 252, 0.7), rgba(125, 211, 252, 0) 68%);
}
.ci-orb--two {
  left: -72px;
  top: 348px;
  height: 230px;
  width: 230px;
  background: radial-gradient(circle at 60% 40%, rgba(249, 168, 212, 0.45), rgba(249, 168, 212, 0) 70%);
}
.ci-head {
  position: relative;
  display: flex;
  align-items: center;
  justify-content: space-between;
  animation: ci-glass-in 0.6s ease both;
}
.ci-brand {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  font-size: 11px;
  letter-spacing: 0.16em;
  color: #5d7798;
}
.ci-id {
  font-family: ui-monospace, Consolas, monospace;
  font-size: 9px;
  color: #8ba2bf;
}
.ci-title {
  position: relative;
  margin-top: 39px;
  font-size: 33px;
  font-weight: 700;
  line-height: 1.3;
  letter-spacing: -0.02em;
  color: #22334c;
  animation: ci-glass-in 0.65s 0.07s ease both;
}
.ci-sub {
  position: relative;
  margin-top: 11px;
  max-width: 20em;
  font-size: 13px;
  line-height: 1.8;
  color: #647b99;
  animation: ci-glass-in 0.65s 0.13s ease both;
}
.ci-stack {
  position: relative;
  margin-top: 23px;
  display: flex;
  flex-direction: column;
  gap: 10px;
  animation: ci-glass-in 0.65s 0.19s ease both;
}
.ci-panel {
  border: 1px solid rgba(255, 255, 255, 0.85);
  border-radius: 20px;
  background: rgba(255, 255, 255, 0.56);
  padding: 18px 19px;
  box-shadow: 0 20px 44px -30px rgba(59, 88, 126, 0.55), inset 0 1px 0 rgba(255, 255, 255, 0.9);
  backdrop-filter: blur(14px);
}
.ci-panel--rows {
  margin-left: 20px;
  padding: 7px 19px;
}
.ci-panel-head {
  display: flex;
  align-items: center;
  gap: 7px;
  font-size: 10px;
  font-weight: 600;
  letter-spacing: 0.16em;
  color: #7890ad;
}
.ci-dot {
  height: 5px;
  width: 5px;
  border-radius: 50%;
  background: #6ea8d8;
}
.ci-name {
  margin-top: 7px;
  font-size: 20px;
  font-weight: 650;
  line-height: 1.55;
  color: #1e3049;
}
.ci-row {
  display: flex;
  align-items: flex-start;
  justify-content: space-between;
  gap: 13px;
  border-bottom: 1px solid rgba(90, 116, 148, 0.13);
  padding: 11px 0;
}
.ci-row:last-child { border-bottom: none; }
.ci-row span {
  flex-shrink: 0;
  font-size: 11px;
  color: #788dab;
}
.ci-row strong {
  max-width: 61%;
  text-align: right;
  font-size: 13px;
  font-weight: 600;
  line-height: 1.5;
  color: #263a55;
}
.ci-call {
  position: relative;
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 12px;
  margin-top: 16px;
  overflow: hidden;
  border: 1px solid rgba(255, 255, 255, 0.9);
  border-radius: 20px;
  background: linear-gradient(120deg, rgba(30, 64, 175, 0.92), rgba(37, 99, 235, 0.82));
  padding: 18px 19px;
  color: #ffffff;
  text-decoration: none;
  box-shadow: 0 24px 42px -24px rgba(30, 64, 175, 0.7);
  backdrop-filter: blur(10px);
  transition: transform 0.18s ease, box-shadow 0.18s ease;
  animation: ci-glass-in 0.65s 0.25s ease both;
}
.ci-call:hover {
  box-shadow: 0 28px 48px -22px rgba(30, 64, 175, 0.85);
}
.ci-call:active { transform: scale(0.982); }
.ci-call-copy { min-width: 0; }
.ci-call-label {
  display: flex;
  align-items: center;
  gap: 6px;
  font-size: 10px;
  letter-spacing: 0.1em;
  color: #cfdffb;
}
.ci-call-num {
  display: block;
  margin-top: 6px;
  font-size: 29px;
  font-weight: 750;
  line-height: 1.15;
  word-break: break-all;
  font-variant-numeric: tabular-nums;
}
.ci-call-btn {
  flex-shrink: 0;
  border-radius: 999px;
  background: rgba(255, 255, 255, 0.94);
  padding: 9px 15px;
  color: #1e40af;
  font-size: 12px;
  font-weight: 700;
}
.ci-empty {
  margin-top: 16px;
  border: 1px dashed rgba(90, 116, 148, 0.4);
  border-radius: 20px;
  padding: 18px;
  font-size: 14px;
  color: #7189a7;
}
.ci-foot {
  position: relative;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 6px;
  margin-top: 18px;
  font-size: 10px;
  color: #7d92ae;
  animation: ci-glass-in 0.65s 0.31s ease both;
}
@keyframes ci-glass-in {
  from { opacity: 0; transform: translateY(14px); }
  to { opacity: 1; transform: none; }
}
@media (prefers-reduced-motion: reduce) {
  .ci-head, .ci-title, .ci-sub, .ci-stack, .ci-call, .ci-foot { animation: none; }
  .ci-call { transition: none; }
}
</style>
