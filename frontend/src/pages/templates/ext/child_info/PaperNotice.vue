<!-- 求助信息卡 · 社区告示：米色纸张与红色印章，像贴在楼下的寻人启事 -->
<template>
  <div class="ci-paper min-h-screen">
    <div class="mx-auto w-full max-w-md px-5 pb-14 pt-8">
      <div class="ci-sheet">
        <div class="ci-stamp" aria-hidden="true">
          <span>急</span>
          <span>寻</span>
        </div>

        <p class="ci-eyebrow">社区互助 · 请留意</p>
        <h1 class="ci-title">求助<br />信息卡</h1>
        <p class="ci-lead">如果您看到这张卡片上的老人，请先停下来，帮他拨打下面这个电话。</p>

        <div class="ci-rule"><span>基本情况</span></div>
        <p class="ci-name">{{ payload.baseInfo }}</p>

        <div class="ci-details">
          <div class="ci-detail">
            <span class="ci-detail-key">学校 / 单位</span>
            <span class="ci-detail-dots" aria-hidden="true"></span>
            <span class="ci-detail-val">{{ payload.school }}</span>
          </div>
          <div class="ci-detail">
            <span class="ci-detail-key">过敏史 / 病史</span>
            <span class="ci-detail-dots" aria-hidden="true"></span>
            <span class="ci-detail-val">{{ payload.allergies || '无' }}</span>
          </div>
        </div>

        <a
          v-if="payload.emergencyPhone"
          :href="'tel:' + payload.emergencyPhone"
          class="ci-call"
        >
          <span class="ci-call-top">
            <Phone :size="17" />
            <span>紧急联系电话</span>
          </span>
          <span class="ci-call-num">{{ payload.emergencyPhone }}</span>
          <span class="ci-call-hint">点击号码即可拨打</span>
        </a>

        <p class="ci-sign">谢谢您愿意多看一眼。<br />—— 家人敬上</p>
      </div>

      <p class="ci-foot">请陪伴在旁，等家人赶到再离开</p>
    </div>
  </div>
</template>

<script setup>
import { Phone } from 'lucide-vue-next'

defineProps(['payload', 'nfcId'])
</script>

<style scoped>
.ci-paper {
  background: #e7ddc7;
  color: #2b2117;
  font-family: Georgia, "Songti SC", "SimSun", serif;
}
.ci-sheet {
  position: relative;
  overflow: hidden;
  border-radius: 3px;
  background-color: #f7f0df;
  background-image:
    repeating-linear-gradient(0deg, rgba(107, 84, 52, 0.055) 0 1px, transparent 1px 26px),
    radial-gradient(circle at 82% 8%, rgba(181, 148, 92, 0.16), transparent 44%);
  padding: 30px 26px 28px;
  box-shadow: 0 3px 0 #d6c7a9, 0 26px 50px -30px rgba(63, 45, 20, 0.65);
  animation: ci-paper-in 0.6s ease both;
}
.ci-sheet::before {
  content: "";
  position: absolute;
  inset: 9px;
  border: 1px solid rgba(107, 84, 52, 0.24);
  pointer-events: none;
}
.ci-stamp {
  position: absolute;
  right: 21px;
  top: 22px;
  display: flex;
  height: 64px;
  width: 64px;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  border: 2px solid #a92c20;
  border-radius: 50%;
  color: #a92c20;
  font-size: 16px;
  font-weight: 700;
  line-height: 1.05;
  transform: rotate(11deg);
  opacity: 0.85;
  animation: ci-paper-stamp 0.5s 0.4s ease both;
}
.ci-eyebrow {
  font-family: ui-monospace, Consolas, monospace;
  font-size: 10px;
  letter-spacing: 0.3em;
  color: #8b7351;
}
.ci-title {
  margin-top: 16px;
  font-size: 39px;
  font-weight: 700;
  line-height: 1.05;
  letter-spacing: 0.04em;
  color: #2b2117;
}
.ci-lead {
  margin-top: 14px;
  max-width: 19em;
  font-size: 14px;
  line-height: 1.9;
  color: #5e4c36;
}
.ci-rule {
  margin-top: 24px;
  display: flex;
  align-items: center;
  gap: 12px;
  border-top: 2px solid #2b2117;
  padding-top: 10px;
}
.ci-rule span {
  font-family: ui-monospace, Consolas, monospace;
  font-size: 10px;
  letter-spacing: 0.24em;
  color: #8b7351;
}
.ci-name {
  margin-top: 8px;
  font-size: 21px;
  font-weight: 700;
  line-height: 1.6;
}
.ci-details {
  margin-top: 16px;
}
.ci-detail {
  display: flex;
  align-items: baseline;
  gap: 8px;
  padding: 9px 0;
  border-bottom: 1px dotted rgba(107, 84, 52, 0.38);
}
.ci-detail-key {
  flex-shrink: 0;
  font-size: 12px;
  color: #7a6448;
}
.ci-detail-dots {
  flex: 1;
  border-bottom: 1px dotted rgba(107, 84, 52, 0.25);
  transform: translateY(-3px);
}
.ci-detail-val {
  max-width: 55%;
  text-align: right;
  font-size: 13px;
  font-weight: 700;
}
.ci-call {
  display: block;
  margin-top: 24px;
  border: 2px solid #a92c20;
  background: #fffaf0;
  padding: 16px 17px 14px;
  color: #a92c20;
  text-decoration: none;
  transition: background 0.18s ease, transform 0.18s ease;
  animation: ci-paper-in 0.55s 0.16s ease both;
}
.ci-call:hover {
  background: #fbe9dd;
}
.ci-call:active {
  transform: scale(0.985);
}
.ci-call-top {
  display: flex;
  align-items: center;
  gap: 6px;
  font-family: ui-monospace, Consolas, monospace;
  font-size: 10px;
  letter-spacing: 0.2em;
}
.ci-call-num {
  display: block;
  margin-top: 7px;
  font-size: 33px;
  font-weight: 700;
  line-height: 1.16;
  letter-spacing: 0.01em;
  word-break: break-all;
  font-variant-numeric: tabular-nums;
}
.ci-call-hint {
  display: block;
  margin-top: 5px;
  font-size: 11px;
  color: #8b5a45;
}
.ci-empty {
  margin-top: 24px;
  border: 2px dashed #a92c20;
  padding: 16px;
  font-size: 14px;
  color: #a92c20;
}
.ci-sign {
  margin-top: 21px;
  text-align: right;
  font-size: 13px;
  line-height: 1.8;
  color: #6b5638;
}
.ci-foot {
  margin-top: 14px;
  text-align: center;
  font-family: ui-monospace, Consolas, monospace;
  font-size: 10px;
  letter-spacing: 0.12em;
  color: #6f5d40;
  animation: ci-paper-in 0.6s 0.28s ease both;
}
@keyframes ci-paper-in {
  from { opacity: 0; transform: translateY(14px); }
  to { opacity: 1; transform: none; }
}
@keyframes ci-paper-stamp {
  from { opacity: 0; transform: rotate(11deg) scale(1.5); }
  to { opacity: 0.85; transform: rotate(11deg) scale(1); }
}
@media (prefers-reduced-motion: reduce) {
  .ci-sheet, .ci-call, .ci-foot, .ci-stamp { animation: none; }
  .ci-call { transition: none; }
}
</style>
