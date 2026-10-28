<!-- 求助信息卡 · 分屏守护：上半深绿放电话，下半米白放资料，重点极端明确 -->
<template>
  <div class="ci-split min-h-screen">
    <div class="mx-auto flex min-h-screen w-full max-w-md flex-col">
      <section class="ci-upper">
        <div class="ci-upper-glow" aria-hidden="true"></div>
        <div class="ci-upper-inner">
          <header class="ci-head">
            <span class="ci-brand"><ShieldCheck :size="15" /> 安心守护</span>
            <span v-if="nfcId" class="ci-id">NO. {{ nfcId }}</span>
          </header>

          <span class="ci-icon"><HandHeart :size="24" /></span>
          <h1>请帮我<br />拨打这个电话</h1>
          <p class="ci-sub">我可能迷路了。电话那头是我的家人，他们很快就会来接我。</p>

          <a
            v-if="payload.emergencyPhone"
            :href="'tel:' + payload.emergencyPhone"
            class="ci-call"
          >
            <span class="ci-call-label">紧急联系电话</span>
            <span class="ci-call-num">{{ payload.emergencyPhone }}</span>
            <span class="ci-call-btn"><Phone :size="16" /> 点此拨打</span>
          </a>
        </div>
        <span class="ci-wave" aria-hidden="true"></span>
      </section>

      <section class="ci-lower">
        <div class="ci-lower-inner">
          <div class="ci-block ci-block--name">
            <span class="ci-key">基本情况</span>
            <p class="ci-name">{{ payload.baseInfo }}</p>
          </div>
          <div class="ci-block">
            <span class="ci-key">学校 / 单位</span>
            <p class="ci-value">{{ payload.school }}</p>
          </div>
          <div class="ci-block">
            <span class="ci-key">过敏史 / 病史</span>
            <p class="ci-value ci-value--warn">{{ payload.allergies || '无' }}</p>
          </div>
          <p class="ci-foot">请陪伴在旁，等家人到达后再离开。谢谢您。</p>
        </div>
      </section>
    </div>
  </div>
</template>

<script setup>
import { HandHeart, Phone, ShieldCheck } from 'lucide-vue-next'

defineProps(['payload', 'nfcId'])
</script>

<style scoped>
.ci-split {
  background: #f7f2e7;
  color: #22312c;
  font-family: "PingFang SC", "Microsoft YaHei", sans-serif;
}
.ci-upper {
  position: relative;
  overflow: hidden;
  background: linear-gradient(160deg, #17483f 0%, #0e332d 100%);
  padding: 22px 24px 0;
  color: #f5f1e4;
  animation: ci-split-in 0.65s ease both;
}
.ci-upper-glow {
  position: absolute;
  right: -70px;
  top: -90px;
  height: 260px;
  width: 260px;
  border-radius: 50%;
  background: radial-gradient(circle, rgba(126, 196, 169, 0.3), transparent 68%);
  pointer-events: none;
}
.ci-upper-inner {
  position: relative;
  z-index: 1;
  padding-bottom: 41px;
}
.ci-head {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 8px;
  border-bottom: 1px solid rgba(233, 241, 232, 0.18);
  padding-bottom: 10px;
}
.ci-brand {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  font-size: 10px;
  letter-spacing: 0.16em;
  color: #b9d6c8;
}
.ci-id {
  font-family: ui-monospace, Consolas, monospace;
  font-size: 8px;
  color: #8fb7a6;
}
.ci-icon {
  display: flex;
  height: 45px;
  width: 45px;
  align-items: center;
  justify-content: center;
  border-radius: 50%;
  background: rgba(245, 241, 228, 0.1);
  border: 1px solid rgba(245, 241, 228, 0.2);
  margin-top: 20px;
  color: #e9d9ae;
}
.ci-upper h1 {
  margin-top: 12px;
  font-size: 34px;
  font-weight: 700;
  line-height: 1.3;
  letter-spacing: -0.02em;
  color: #fbf7ea;
  animation: ci-split-in 0.6s 0.08s ease both;
}
.ci-sub {
  margin-top: 9px;
  max-width: 21em;
  font-size: 12px;
  line-height: 1.8;
  color: #a9c7ba;
  animation: ci-split-in 0.6s 0.12s ease both;
}
.ci-call {
  display: block;
  margin-top: 19px;
  border-radius: 15px;
  background: #f0dfb3;
  padding: 15px 17px 14px;
  color: #22312c;
  text-decoration: none;
  box-shadow: 0 18px 30px -22px rgba(0, 0, 0, 0.8);
  transition: transform 0.18s ease, background 0.18s ease;
  animation: ci-split-in 0.6s 0.17s ease both;
}
.ci-call:hover { background: #f6e9c8; }
.ci-call:active { transform: scale(0.983); }
.ci-call-label {
  display: block;
  font-size: 10px;
  font-weight: 700;
  letter-spacing: 0.12em;
  color: #6f5a28;
}
.ci-call-num {
  display: block;
  margin-top: 4px;
  font-size: 32px;
  font-weight: 800;
  line-height: 1.16;
  letter-spacing: 0.01em;
  word-break: break-all;
  font-variant-numeric: tabular-nums;
  color: #182723;
}
.ci-call-btn {
  display: inline-flex;
  align-items: center;
  gap: 5px;
  margin-top: 8px;
  border-radius: 999px;
  background: #17483f;
  padding: 5px 11px;
  font-size: 10px;
  font-weight: 700;
  color: #f0dfb3;
}
.ci-empty {
  margin-top: 19px;
  border: 1px dashed rgba(240, 223, 179, 0.5);
  border-radius: 15px;
  padding: 15px;
  font-size: 12px;
  color: #d8c491;
}
.ci-wave {
  position: absolute;
  left: 0;
  right: 0;
  bottom: -1px;
  height: 28px;
  background: #f7f2e7;
  clip-path: polygon(0 68%, 9% 58%, 21% 72%, 34% 55%, 48% 70%, 62% 52%, 77% 69%, 90% 57%, 100% 66%, 100% 100%, 0 100%);
}
.ci-lower {
  flex: 1;
  padding: 9px 24px 25px;
  animation: ci-split-in 0.65s 0.1s ease both;
}
.ci-lower-inner {
  max-width: 32rem;
}
.ci-block {
  border-bottom: 1px solid #ded6c4;
  padding: 12px 0;
}
.ci-block--name { padding-top: 4px; }
.ci-block:last-of-type { border-bottom: none; }
.ci-key {
  font-size: 9px;
  font-weight: 700;
  letter-spacing: 0.14em;
  color: #8c8a7a;
}
.ci-name {
  margin-top: 4px;
  font-size: 19px;
  font-weight: 700;
  line-height: 1.5;
  color: #24322d;
}
.ci-value {
  margin-top: 4px;
  font-size: 14px;
  font-weight: 550;
  line-height: 1.55;
  color: #40524b;
}
.ci-value--warn { color: #9c4f20; }
.ci-foot {
  margin-top: 11px;
  font-size: 10px;
  line-height: 1.7;
  color: #8c8a7a;
}
@keyframes ci-split-in {
  from { opacity: 0; transform: translateY(14px); }
  to { opacity: 1; transform: none; }
}
@media (prefers-reduced-motion: reduce) {
  .ci-upper, .ci-upper h1, .ci-sub, .ci-call, .ci-lower { animation: none; }
  .ci-call { transition: none; }
}
</style>
