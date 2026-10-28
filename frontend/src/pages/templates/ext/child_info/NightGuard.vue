<!-- 求助信息卡 · 夜间守护：深蓝夜空与琥珀色电话，夜里也清晰可辨 -->
<template>
  <div class="ci-night min-h-screen">
    <div class="ci-stars" aria-hidden="true"></div>
    <div class="mx-auto w-full max-w-md px-6 pb-12 pt-10">
      <header class="ci-head">
        <span class="ci-brand"><MoonStar :size="16" /> 夜间守护卡</span>
        <span v-if="nfcId" class="ci-id">NFC {{ nfcId }}</span>
      </header>

      <div class="ci-moon" aria-hidden="true"></div>

      <h1 class="ci-title">夜深了，<br />请帮我回家</h1>
      <p class="ci-sub">如果您遇见独自在外的我，请拨打下面的电话。家人会立刻赶来。</p>

      <section class="ci-card">
        <span class="ci-card-label">基本情况</span>
        <p class="ci-name">{{ payload.baseInfo }}</p>
        <div class="ci-rows">
          <div class="ci-row">
            <span>学校 / 单位</span>
            <strong>{{ payload.school }}</strong>
          </div>
          <div class="ci-row">
            <span>过敏史 / 病史</span>
            <strong class="ci-warn">{{ payload.allergies || '无' }}</strong>
          </div>
        </div>
      </section>

      <a
        v-if="payload.emergencyPhone"
        :href="'tel:' + payload.emergencyPhone"
        class="ci-call"
      >
        <span class="ci-call-label"><Phone :size="17" /> 紧急联系电话 · 点击拨打</span>
        <span class="ci-call-num">{{ payload.emergencyPhone }}</span>
        <span class="ci-call-ring" aria-hidden="true"></span>
      </a>

      <p class="ci-foot">愿每一次求助，都能被温柔接住。</p>
    </div>
  </div>
</template>

<script setup>
import { MoonStar, Phone } from 'lucide-vue-next'

defineProps(['payload', 'nfcId'])
</script>

<style scoped>
.ci-night {
  position: relative;
  overflow: hidden;
  background: radial-gradient(circle at 78% 8%, #172a4d 0%, #0a1122 46%, #070b16 100%);
  color: #edf2ff;
  font-family: "PingFang SC", "Microsoft YaHei", sans-serif;
}
.ci-stars {
  position: absolute;
  inset: 0;
  background-image:
    radial-gradient(1px 1px at 12% 18%, rgba(255, 255, 255, 0.75), transparent),
    radial-gradient(1.4px 1.4px at 72% 12%, rgba(255, 214, 138, 0.8), transparent),
    radial-gradient(1px 1px at 38% 34%, rgba(255, 255, 255, 0.5), transparent),
    radial-gradient(1.6px 1.6px at 88% 42%, rgba(255, 255, 255, 0.65), transparent),
    radial-gradient(1px 1px at 22% 68%, rgba(255, 255, 255, 0.4), transparent);
  pointer-events: none;
  animation: ci-night-fade 1.2s ease both;
}
.ci-head {
  position: relative;
  display: flex;
  align-items: center;
  justify-content: space-between;
  animation: ci-night-in 0.6s ease both;
}
.ci-brand {
  display: inline-flex;
  align-items: center;
  gap: 7px;
  font-size: 11px;
  letter-spacing: 0.16em;
  color: #a8bbdd;
}
.ci-id {
  font-family: ui-monospace, Consolas, monospace;
  font-size: 9px;
  color: #6d80a5;
}
.ci-moon {
  position: absolute;
  right: 27px;
  top: 99px;
  height: 53px;
  width: 53px;
  border-radius: 50%;
  background: #ffd98a;
  box-shadow: 0 0 38px rgba(255, 217, 138, 0.3);
  animation: ci-night-in 0.8s 0.1s ease both;
}
.ci-moon::after {
  content: "";
  position: absolute;
  left: 17px;
  top: -6px;
  height: 53px;
  width: 53px;
  border-radius: 50%;
  background: #0b1428;
}
.ci-title {
  position: relative;
  margin-top: 61px;
  font-size: 37px;
  font-weight: 700;
  line-height: 1.25;
  letter-spacing: -0.02em;
  color: #f4f7ff;
  animation: ci-night-in 0.65s 0.07s ease both;
}
.ci-sub {
  position: relative;
  margin-top: 13px;
  max-width: 20em;
  font-size: 13px;
  line-height: 1.85;
  color: #9fb1d2;
  animation: ci-night-in 0.65s 0.13s ease both;
}
.ci-card {
  margin-top: 24px;
  border: 1px solid rgba(160, 185, 225, 0.17);
  border-radius: 18px;
  background: rgba(255, 255, 255, 0.045);
  padding: 20px;
  animation: ci-night-in 0.65s 0.19s ease both;
}
.ci-card-label {
  font-size: 10px;
  letter-spacing: 0.2em;
  color: #8ea4c9;
}
.ci-name {
  margin-top: 7px;
  font-size: 20px;
  font-weight: 650;
  line-height: 1.55;
  color: #ffffff;
}
.ci-rows {
  margin-top: 13px;
  border-top: 1px solid rgba(160, 185, 225, 0.14);
}
.ci-row {
  display: flex;
  align-items: flex-start;
  justify-content: space-between;
  gap: 12px;
  border-bottom: 1px solid rgba(160, 185, 225, 0.1);
  padding: 10px 0;
}
.ci-row:last-child { border-bottom: none; }
.ci-row span {
  flex-shrink: 0;
  font-size: 11px;
  color: #8fa3c5;
}
.ci-row strong {
  max-width: 61%;
  text-align: right;
  font-size: 13px;
  font-weight: 600;
  line-height: 1.5;
  color: #dfe8fa;
}
.ci-row .ci-warn { color: #ffc27a; }
.ci-call {
  position: relative;
  display: block;
  overflow: hidden;
  margin-top: 18px;
  border-radius: 18px;
  background: #ffc857;
  padding: 19px 20px 18px;
  color: #1a1607;
  text-decoration: none;
  box-shadow: 0 0 0 1px rgba(255, 231, 173, 0.5), 0 22px 40px -22px rgba(255, 200, 87, 0.6);
  transition: transform 0.18s ease, box-shadow 0.18s ease;
  animation: ci-night-in 0.65s 0.25s ease both;
}
.ci-call:hover {
  box-shadow: 0 0 0 1px #ffe7ad, 0 26px 46px -20px rgba(255, 200, 87, 0.8);
}
.ci-call:active { transform: scale(0.982); }
.ci-call-label {
  display: flex;
  align-items: center;
  gap: 6px;
  font-size: 11px;
  font-weight: 700;
  letter-spacing: 0.05em;
  color: #4d3a06;
}
.ci-call-num {
  display: block;
  margin-top: 6px;
  font-size: 34px;
  font-weight: 800;
  line-height: 1.15;
  letter-spacing: 0.01em;
  word-break: break-all;
  font-variant-numeric: tabular-nums;
}
.ci-call-ring {
  position: absolute;
  right: -26px;
  bottom: -34px;
  height: 122px;
  width: 122px;
  border: 1px solid rgba(77, 58, 6, 0.18);
  border-radius: 50%;
}
.ci-empty {
  margin-top: 18px;
  border: 1px dashed rgba(255, 200, 87, 0.45);
  border-radius: 18px;
  padding: 18px;
  font-size: 14px;
  color: #ffc857;
}
.ci-foot {
  margin-top: 20px;
  text-align: center;
  font-size: 11px;
  letter-spacing: 0.02em;
  color: #8195b8;
  animation: ci-night-in 0.65s 0.31s ease both;
}
@keyframes ci-night-in {
  from { opacity: 0; transform: translateY(15px); }
  to { opacity: 1; transform: none; }
}
@keyframes ci-night-fade {
  from { opacity: 0; }
  to { opacity: 1; }
}
@media (prefers-reduced-motion: reduce) {
  .ci-stars, .ci-head, .ci-moon, .ci-title, .ci-sub, .ci-card, .ci-call, .ci-foot { animation: none; }
  .ci-call { transition: none; }
}
</style>
