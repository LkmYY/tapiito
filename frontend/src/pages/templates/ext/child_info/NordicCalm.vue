<!-- 求助信息卡 · 北欧静谧：灰绿与奶白，留白充足，安静而可信 -->
<template>
  <div class="ci-nordic min-h-screen">
    <div class="mx-auto w-full max-w-md px-7 pb-14 pt-11">
      <header class="ci-head">
        <span class="ci-brand">HJÆLP · 求助</span>
        <span v-if="nfcId" class="ci-id">{{ nfcId }}</span>
      </header>

      <div class="ci-lead">
        <span class="ci-rule"></span>
        <p>如果你遇见<br />需要帮助的人</p>
      </div>

      <p class="ci-intro">请停下来，拨通下面的电话。家人会尽快赶到，你也可以陪他一起等。</p>

      <div class="ci-list">
        <div class="ci-item">
          <span class="ci-item-no">01</span>
          <div>
            <span class="ci-key">基本情况</span>
            <p class="ci-name">{{ payload.baseInfo }}</p>
          </div>
        </div>
        <div class="ci-item">
          <span class="ci-item-no">02</span>
          <div>
            <span class="ci-key">学校 / 单位</span>
            <p class="ci-value">{{ payload.school }}</p>
          </div>
        </div>
        <div class="ci-item">
          <span class="ci-item-no">03</span>
          <div>
            <span class="ci-key">过敏史 / 病史</span>
            <p class="ci-value ci-value--alert">{{ payload.allergies || '无' }}</p>
          </div>
        </div>
      </div>

      <p class="ci-quiet">电话 · CALL</p>
      <a
        v-if="payload.emergencyPhone"
        :href="'tel:' + payload.emergencyPhone"
        class="ci-call"
      >
        <span class="ci-call-num">{{ payload.emergencyPhone }}</span>
        <span class="ci-call-arrow"><ArrowUpRight :size="21" /></span>
      </a>

      <footer class="ci-foot">
        <p>请陪伴至家人到达</p>
        <p class="ci-foot-en">TAK FOR DIN OMSORG</p>
      </footer>
    </div>
  </div>
</template>

<script setup>
import { ArrowUpRight } from 'lucide-vue-next'

defineProps(['payload', 'nfcId'])
</script>

<style scoped>
.ci-nordic {
  background: #e9e8e1;
  color: #26302a;
  font-family: "Avenir Next", "Helvetica Neue", "PingFang SC", sans-serif;
}
.ci-head {
  display: flex;
  align-items: center;
  justify-content: space-between;
  animation: ci-nordic-in 0.65s ease both;
}
.ci-brand {
  font-size: 9px;
  letter-spacing: 0.26em;
  color: #5e6b61;
}
.ci-id {
  font-family: ui-monospace, Consolas, monospace;
  font-size: 8px;
  color: #879187;
}
.ci-lead {
  margin-top: 42px;
  display: flex;
  align-items: flex-start;
  gap: 13px;
  animation: ci-nordic-in 0.65s 0.06s ease both;
}
.ci-rule {
  width: 25px;
  height: 1px;
  margin-top: 12px;
  background: #5c6b60;
}
.ci-lead p {
  font-size: 31px;
  font-weight: 300;
  line-height: 1.4;
  letter-spacing: -0.01em;
  color: #26302a;
}
.ci-intro {
  margin-top: 15px;
  margin-left: 38px;
  max-width: 19em;
  font-size: 12px;
  line-height: 1.9;
  color: #67736a;
  animation: ci-nordic-in 0.65s 0.11s ease both;
}
.ci-list {
  margin-top: 33px;
  animation: ci-nordic-in 0.65s 0.16s ease both;
}
.ci-item {
  display: flex;
  gap: 16px;
  border-top: 1px solid #c9ccc2;
  padding: 14px 0;
}
.ci-item:last-child { border-bottom: 1px solid #c9ccc2; }
.ci-item-no {
  flex-shrink: 0;
  padding-top: 2px;
  font-family: ui-monospace, Consolas, monospace;
  font-size: 8px;
  letter-spacing: 0.08em;
  color: #8a958a;
}
.ci-key {
  display: block;
  font-size: 8px;
  font-weight: 600;
  letter-spacing: 0.2em;
  color: #839084;
}
.ci-name {
  margin-top: 4px;
  font-size: 19px;
  font-weight: 500;
  line-height: 1.55;
  color: #243028;
}
.ci-value {
  margin-top: 4px;
  font-size: 14px;
  font-weight: 500;
  line-height: 1.55;
  color: #3a463e;
}
.ci-value--alert { color: #9a4e26; }
.ci-quiet {
  margin-top: 29px;
  font-size: 8px;
  font-weight: 600;
  letter-spacing: 0.24em;
  color: #839084;
  animation: ci-nordic-in 0.65s 0.22s ease both;
}
.ci-call {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 10px;
  margin-top: 7px;
  border-top: 1px solid #26302a;
  border-bottom: 1px solid #26302a;
  padding: 13px 1px;
  color: #1f2a24;
  text-decoration: none;
  transition: padding 0.22s ease, background 0.22s ease;
  animation: ci-nordic-in 0.65s 0.26s ease both;
}
.ci-call:hover {
  background: #f1f0ea;
  padding-left: 10px;
  padding-right: 10px;
}
.ci-call:active { transform: scale(0.992); }
.ci-call-num {
  font-size: 31px;
  font-weight: 400;
  line-height: 1.2;
  letter-spacing: 0.005em;
  word-break: break-all;
  font-variant-numeric: tabular-nums;
  color: #1b2620;
}
.ci-call-arrow {
  display: flex;
  height: 33px;
  width: 33px;
  flex-shrink: 0;
  align-items: center;
  justify-content: center;
  border-radius: 50%;
  background: #2e3b33;
  color: #e9e8e1;
}
.ci-empty {
  margin-top: 7px;
  border-top: 1px solid #26302a;
  border-bottom: 1px solid #26302a;
  padding: 13px 1px;
  font-size: 12px;
  color: #6b776d;
}
.ci-foot {
  display: flex;
  align-items: flex-end;
  justify-content: space-between;
  gap: 10px;
  margin-top: 20px;
  font-size: 10px;
  color: #788479;
  animation: ci-nordic-in 0.65s 0.31s ease both;
}
.ci-foot-en {
  font-size: 7px;
  letter-spacing: 0.14em;
  color: #97a096;
}
@keyframes ci-nordic-in {
  from { opacity: 0; transform: translateY(12px); }
  to { opacity: 1; transform: none; }
}
@media (prefers-reduced-motion: reduce) {
  .ci-head, .ci-lead, .ci-intro, .ci-list, .ci-quiet, .ci-call, .ci-foot { animation: none; }
  .ci-call { transition: none; }
}
</style>
