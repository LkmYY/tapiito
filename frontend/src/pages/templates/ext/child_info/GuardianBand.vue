<!-- 求助信息卡 · 守护腕带：黑黄高对比，超大号码，老人也看得清 -->
<template>
  <div class="ci-band min-h-screen">
    <div class="mx-auto w-full max-w-md px-5 pb-12 pt-7">
      <div class="ci-stripes" aria-hidden="true"></div>

      <div class="ci-shell">
        <div class="ci-band-top">
          <span class="ci-chip"><Shield :size="15" /> SAFE RETURN</span>
          <span v-if="nfcId" class="ci-serial">#{{ nfcId }}</span>
        </div>

        <p class="ci-vertical" aria-hidden="true">守护腕带 · 请帮我回家</p>

        <h1 class="ci-title">我可能需要<br />您的帮助</h1>
        <p class="ci-sub">如果我独自一人，请帮我联系家人。谢谢您。</p>

        <div class="ci-owner">
          <span class="ci-owner-tag">基本情况</span>
          <p class="ci-owner-name">{{ payload.baseInfo }}</p>
          <div class="ci-owner-grid">
            <div>
              <span>学校 / 单位</span>
              <strong>{{ payload.school }}</strong>
            </div>
            <div>
              <span>过敏史 / 病史</span>
              <strong>{{ payload.allergies || '无' }}</strong>
            </div>
          </div>
        </div>

        <a
          v-if="payload.emergencyPhone"
          :href="'tel:' + payload.emergencyPhone"
          class="ci-phone"
        >
          <span class="ci-phone-head">
            <span class="ci-phone-label">紧急联系电话</span>
            <span class="ci-phone-cta"><Phone :size="16" /> 点击拨打</span>
          </span>
          <span class="ci-phone-num">{{ payload.emergencyPhone }}</span>
        </a>

        <p class="ci-note">请留在原地陪着我，家人正在赶来</p>
      </div>

      <div class="ci-stripes ci-stripes--bottom" aria-hidden="true"></div>
    </div>
  </div>
</template>

<script setup>
import { Phone, Shield } from 'lucide-vue-next'

defineProps(['payload', 'nfcId'])
</script>

<style scoped>
.ci-band {
  background: #090909;
  color: #f7f7f2;
  font-family: "Arial Narrow", "PingFang SC", "Microsoft YaHei", sans-serif;
}
.ci-stripes {
  height: 13px;
  border-radius: 3px;
  background: repeating-linear-gradient(115deg, #ffd400 0 16px, #111111 16px 32px);
  animation: ci-band-in 0.5s ease both;
}
.ci-stripes--bottom {
  margin-top: 14px;
}
.ci-shell {
  position: relative;
  margin-top: 14px;
  overflow: hidden;
  border-radius: 14px;
  background: #ffd400;
  padding: 22px 22px 24px;
  color: #0a0a0a;
  box-shadow: 0 0 0 1px #2b2b2b, 0 30px 60px -34px rgba(255, 212, 0, 0.5);
  animation: ci-band-in 0.62s 0.08s ease both;
}
.ci-band-top {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 10px;
  border-bottom: 2px solid #0a0a0a;
  padding-bottom: 11px;
}
.ci-chip {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  border-radius: 3px;
  background: #0a0a0a;
  padding: 5px 8px;
  color: #ffd400;
  font-size: 10px;
  font-weight: 700;
  letter-spacing: 0.16em;
}
.ci-serial {
  font-family: ui-monospace, Consolas, monospace;
  font-size: 11px;
  font-weight: 700;
}
.ci-vertical {
  position: absolute;
  right: 9px;
  top: 96px;
  font-size: 9px;
  font-weight: 700;
  letter-spacing: 0.32em;
  writing-mode: vertical-rl;
  color: rgba(10, 10, 10, 0.42);
}
.ci-title {
  margin-top: 20px;
  font-size: 36px;
  font-weight: 800;
  line-height: 1.08;
  letter-spacing: -0.03em;
}
.ci-sub {
  margin-top: 10px;
  max-width: 18em;
  font-size: 14px;
  line-height: 1.6;
  font-weight: 500;
  color: #332a00;
}
.ci-owner {
  margin-top: 20px;
  border-top: 2px solid #0a0a0a;
  padding-top: 14px;
}
.ci-owner-tag {
  font-size: 9px;
  font-weight: 700;
  letter-spacing: 0.22em;
}
.ci-owner-name {
  margin-top: 6px;
  font-size: 19px;
  font-weight: 800;
  line-height: 1.5;
}
.ci-owner-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 10px;
  margin-top: 12px;
}
.ci-owner-grid div {
  border-left: 3px solid #0a0a0a;
  padding-left: 9px;
}
.ci-owner-grid span {
  display: block;
  font-size: 9px;
  letter-spacing: 0.08em;
  color: #5c4e00;
}
.ci-owner-grid strong {
  display: block;
  margin-top: 3px;
  font-size: 13px;
  font-weight: 700;
  line-height: 1.45;
}
.ci-phone {
  display: block;
  margin-top: 16px;
  border-radius: 4px;
  background: #0a0a0a;
  padding: 16px 16px 17px;
  color: #ffd400;
  text-decoration: none;
  transition: transform 0.16s ease, background 0.16s ease;
  animation: ci-band-in 0.6s 0.18s ease both;
}
.ci-phone:hover {
  background: #1c1c1c;
}
.ci-phone:active {
  transform: scale(0.985);
}
.ci-phone-head {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 10px;
}
.ci-phone-label {
  font-size: 10px;
  font-weight: 700;
  letter-spacing: 0.18em;
  color: #fff3ad;
}
.ci-phone-cta {
  display: inline-flex;
  align-items: center;
  gap: 4px;
  border: 1px solid #ffd400;
  border-radius: 20px;
  padding: 3px 8px;
  font-size: 10px;
  font-weight: 700;
}
.ci-phone-num {
  display: block;
  margin-top: 9px;
  font-size: 37px;
  font-weight: 800;
  letter-spacing: 0.01em;
  line-height: 1.12;
  word-break: break-all;
  font-variant-numeric: tabular-nums;
}
.ci-empty {
  margin-top: 16px;
  border: 2px dashed #0a0a0a;
  padding: 16px;
  font-size: 14px;
  font-weight: 700;
}
.ci-note {
  margin-top: 13px;
  font-size: 11px;
  font-weight: 600;
  color: #4a3d00;
  text-align: center;
}
@keyframes ci-band-in {
  from { opacity: 0; transform: translateY(16px); }
  to { opacity: 1; transform: none; }
}
@media (prefers-reduced-motion: reduce) {
  .ci-stripes, .ci-shell, .ci-phone { animation: none; }
  .ci-phone { transition: none; }
}
</style>
