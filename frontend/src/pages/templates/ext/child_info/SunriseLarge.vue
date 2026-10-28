<!-- 求助信息卡 · 晨光大字：暖米底与超大字号，为老人优化阅读 -->
<template>
  <div class="ci-sunrise min-h-screen">
    <div class="mx-auto w-full max-w-md px-6 pb-12 pt-9">
      <div class="ci-sun" aria-hidden="true"></div>

      <header class="ci-head">
        <span class="ci-tag"><Sun :size="14" /> 晨光求助卡</span>
        <p v-if="nfcId" class="ci-id">编号 {{ nfcId }}</p>
      </header>

      <h1 class="ci-title">请帮我<br />联系家人</h1>
      <p class="ci-sub">我可能记不清回家的路了。<br />请您拨打卡片上的电话，陪我稍等一会儿。</p>

      <div class="ci-info">
        <div class="ci-block">
          <span class="ci-key">基本情况</span>
          <p class="ci-name">{{ payload.baseInfo }}</p>
        </div>
        <div class="ci-block">
          <span class="ci-key">学校 / 单位</span>
          <p class="ci-value">{{ payload.school }}</p>
        </div>
        <div class="ci-block">
          <span class="ci-key">过敏史 / 病史</span>
          <p class="ci-value ci-value--alert">{{ payload.allergies || '无' }}</p>
        </div>
      </div>

      <a
        v-if="payload.emergencyPhone"
        :href="'tel:' + payload.emergencyPhone"
        class="ci-call"
      >
        <span class="ci-call-label"><Phone :size="18" /> 紧急联系电话</span>
        <span class="ci-call-num">{{ payload.emergencyPhone }}</span>
        <span class="ci-call-tap">轻触此处拨打</span>
      </a>

      <p class="ci-foot">谢谢您停下来，这份善意对我们很重要。</p>
    </div>
  </div>
</template>

<script setup>
import { Phone, Sun } from 'lucide-vue-next'

defineProps(['payload', 'nfcId'])
</script>

<style scoped>
.ci-sunrise {
  position: relative;
  overflow: hidden;
  background: #fdf6e8;
  color: #3c2a16;
  font-family: "PingFang SC", "Microsoft YaHei", "Helvetica Neue", sans-serif;
}
.ci-sun {
  position: absolute;
  left: 50%;
  top: -104px;
  height: 280px;
  width: 280px;
  transform: translateX(-50%);
  border-radius: 50%;
  background: radial-gradient(circle, #ffc766 0%, #ffdfa3 52%, rgba(255, 223, 163, 0) 72%);
  opacity: 0.75;
  animation: ci-sunrise-glow 1s ease both;
  pointer-events: none;
}
.ci-head {
  position: relative;
  display: flex;
  align-items: center;
  justify-content: space-between;
  animation: ci-sunrise-in 0.55s ease both;
}
.ci-tag {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  font-size: 11px;
  font-weight: 700;
  letter-spacing: 0.16em;
  color: #9a5b1b;
}
.ci-id {
  font-family: ui-monospace, Consolas, monospace;
  font-size: 10px;
  color: #a8793c;
}
.ci-title {
  position: relative;
  margin-top: 36px;
  font-size: 46px;
  font-weight: 800;
  line-height: 1.14;
  letter-spacing: -0.035em;
  color: #3b250f;
  animation: ci-sunrise-in 0.6s 0.06s ease both;
}
.ci-sub {
  position: relative;
  margin-top: 15px;
  font-size: 16px;
  line-height: 1.85;
  color: #75512b;
  animation: ci-sunrise-in 0.6s 0.12s ease both;
}
.ci-info {
  margin-top: 27px;
  border-top: 2px solid #3b250f;
  animation: ci-sunrise-in 0.6s 0.18s ease both;
}
.ci-block {
  border-bottom: 1px solid rgba(59, 37, 15, 0.2);
  padding: 15px 0;
}
.ci-key {
  font-size: 12px;
  font-weight: 700;
  letter-spacing: 0.12em;
  color: #9a6f3e;
}
.ci-name {
  margin-top: 5px;
  font-size: 23px;
  font-weight: 800;
  line-height: 1.55;
}
.ci-value {
  margin-top: 5px;
  font-size: 19px;
  font-weight: 650;
  line-height: 1.55;
}
.ci-value--alert {
  color: #a33a12;
}
.ci-call {
  display: block;
  margin-top: 24px;
  border-radius: 18px;
  background: #3b250f;
  padding: 19px 20px 17px;
  color: #ffedc2;
  text-decoration: none;
  box-shadow: 0 20px 34px -24px rgba(59, 37, 15, 0.85);
  transition: transform 0.18s ease, background 0.18s ease;
  animation: ci-sunrise-in 0.6s 0.24s ease both;
}
.ci-call:hover {
  background: #50331a;
}
.ci-call:active {
  transform: scale(0.985);
}
.ci-call-label {
  display: flex;
  align-items: center;
  gap: 7px;
  font-size: 12px;
  font-weight: 700;
  letter-spacing: 0.1em;
  color: #ffd894;
}
.ci-call-num {
  display: block;
  margin-top: 7px;
  font-size: 38px;
  font-weight: 800;
  line-height: 1.16;
  letter-spacing: 0.01em;
  word-break: break-all;
  font-variant-numeric: tabular-nums;
}
.ci-call-tap {
  display: block;
  margin-top: 4px;
  font-size: 12px;
  color: #d8b98c;
}
.ci-empty {
  margin-top: 24px;
  border: 2px dashed #b08a55;
  border-radius: 18px;
  padding: 18px;
  font-size: 15px;
  color: #8a6332;
}
.ci-foot {
  margin-top: 20px;
  font-size: 12px;
  line-height: 1.7;
  color: #96703f;
  animation: ci-sunrise-in 0.6s 0.3s ease both;
}
@keyframes ci-sunrise-in {
  from { opacity: 0; transform: translateY(14px); }
  to { opacity: 1; transform: none; }
}
@keyframes ci-sunrise-glow {
  from { opacity: 0; transform: translateX(-50%) scale(0.82); }
  to { opacity: 0.75; transform: translateX(-50%) scale(1); }
}
@media (prefers-reduced-motion: reduce) {
  .ci-head, .ci-title, .ci-sub, .ci-info, .ci-call, .ci-foot, .ci-sun { animation: none; }
  .ci-call { transition: none; }
}
</style>
