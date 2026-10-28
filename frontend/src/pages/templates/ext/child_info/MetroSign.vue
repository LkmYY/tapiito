<!-- 求助信息卡 · 地铁导视：蓝白标识系统，大字号与箭头，一眼找到电话 -->
<template>
  <div class="ci-metro min-h-screen">
    <div class="mx-auto w-full max-w-md pb-12">
      <header class="ci-top">
        <span class="ci-line-mark">CI · 01</span>
        <span class="ci-top-title">求助信息 · HELP POINT</span>
        <span v-if="nfcId" class="ci-top-id">{{ nfcId }}</span>
      </header>

      <div class="ci-body">
        <div class="ci-rail" aria-hidden="true">
          <span class="ci-rail-dot"></span>
          <span class="ci-rail-line"></span>
          <span class="ci-rail-dot ci-rail-dot--hollow"></span>
        </div>

        <div class="ci-content">
          <p class="ci-kicker">如需帮助 · NEED HELP?</p>
          <h1 class="ci-title">请带我<br />联系家人</h1>
          <p class="ci-intro">我可能迷路了。请拨打下方电话，并陪我在这里等候。</p>

          <section class="ci-panel">
            <div class="ci-panel-head">
              <span class="ci-badge"><UserRound :size="17" /></span>
              <span>基本情况</span>
            </div>
            <p class="ci-name">{{ payload.baseInfo }}</p>

            <div class="ci-grid">
              <div class="ci-cell">
                <span class="ci-cell-key">学校 / 单位</span>
                <span class="ci-cell-val">{{ payload.school }}</span>
              </div>
              <div class="ci-cell">
                <span class="ci-cell-key">过敏史 / 病史</span>
                <span class="ci-cell-val ci-cell-val--warn">{{ payload.allergies || '无' }}</span>
              </div>
            </div>
          </section>

          <a
            v-if="payload.emergencyPhone"
            :href="'tel:' + payload.emergencyPhone"
            class="ci-call"
          >
            <span class="ci-call-side">
              <Phone :size="20" />
              <span>拨打</span>
            </span>
            <span class="ci-call-main">
              <span class="ci-call-label">紧急联系电话</span>
              <span class="ci-call-num">{{ payload.emergencyPhone }}</span>
            </span>
            <span class="ci-arrow" aria-hidden="true">→</span>
          </a>

          <div class="ci-bottom">
            <span class="ci-exit">出口</span>
            <span class="ci-bottom-text">请陪伴至家人到达 · 感谢您的帮助</span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { Phone, UserRound } from 'lucide-vue-next'

defineProps(['payload', 'nfcId'])
</script>

<style scoped>
.ci-metro {
  background: #f2f5f8;
  color: #0b1f3a;
  font-family: "Arial Narrow", "PingFang SC", "Microsoft YaHei", sans-serif;
}
.ci-top {
  display: flex;
  align-items: center;
  gap: 9px;
  background: #0b4da2;
  padding: 13px 18px;
  color: #ffffff;
  animation: ci-metro-in 0.5s ease both;
}
.ci-line-mark {
  border: 1px solid rgba(255, 255, 255, 0.55);
  border-radius: 2px;
  padding: 3px 6px;
  font-family: ui-monospace, Consolas, monospace;
  font-size: 9px;
  letter-spacing: 0.08em;
}
.ci-top-title {
  font-size: 11px;
  font-weight: 700;
  letter-spacing: 0.08em;
}
.ci-top-id {
  margin-left: auto;
  font-family: ui-monospace, Consolas, monospace;
  font-size: 9px;
  opacity: 0.7;
}
.ci-body {
  display: flex;
  gap: 15px;
  padding: 27px 18px 0 13px;
}
.ci-rail {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding-top: 5px;
}
.ci-rail-dot {
  height: 11px;
  width: 11px;
  border-radius: 50%;
  background: #0b4da2;
  box-shadow: 0 0 0 4px #f2f5f8, 0 0 0 5px #b9cbe2;
}
.ci-rail-dot--hollow {
  background: #ffffff;
  box-shadow: 0 0 0 4px #f2f5f8, 0 0 0 5px #b9cbe2;
}
.ci-rail-line {
  width: 2px;
  flex: 1;
  min-height: 280px;
  background: repeating-linear-gradient(180deg, #b9cbe2 0 7px, transparent 7px 12px);
}
.ci-content {
  min-width: 0;
  flex: 1;
  animation: ci-metro-in 0.6s 0.07s ease both;
}
.ci-kicker {
  font-size: 10px;
  font-weight: 700;
  letter-spacing: 0.16em;
  color: #5c7897;
}
.ci-title {
  margin-top: 9px;
  font-size: 38px;
  font-weight: 800;
  line-height: 1.12;
  letter-spacing: -0.03em;
  color: #0b1f3a;
}
.ci-intro {
  margin-top: 11px;
  max-width: 20em;
  font-size: 13px;
  line-height: 1.75;
  color: #4d6480;
}
.ci-panel {
  margin-top: 21px;
  border-radius: 10px;
  background: #ffffff;
  padding: 17px;
  box-shadow: 0 10px 26px -22px rgba(11, 31, 58, 0.55), inset 0 0 0 1px #dbe5ef;
}
.ci-panel-head {
  display: flex;
  align-items: center;
  gap: 7px;
  font-size: 10px;
  font-weight: 700;
  letter-spacing: 0.14em;
  color: #5c7897;
}
.ci-badge {
  display: flex;
  height: 28px;
  width: 28px;
  align-items: center;
  justify-content: center;
  border-radius: 7px;
  background: #e7f0fa;
  color: #0b4da2;
}
.ci-name {
  margin-top: 9px;
  font-size: 20px;
  font-weight: 700;
  line-height: 1.55;
}
.ci-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 9px;
  margin-top: 13px;
}
.ci-cell {
  border: 1px solid #e0e9f2;
  border-radius: 7px;
  padding: 9px 10px;
}
.ci-cell-key {
  display: block;
  font-size: 9px;
  letter-spacing: 0.04em;
  color: #6b829e;
}
.ci-cell-val {
  display: block;
  margin-top: 4px;
  font-size: 12px;
  font-weight: 700;
  line-height: 1.5;
}
.ci-cell-val--warn {
  color: #a84a00;
}
.ci-call {
  display: flex;
  align-items: stretch;
  margin-top: 17px;
  overflow: hidden;
  border-radius: 10px;
  background: #0b4da2;
  color: #ffffff;
  text-decoration: none;
  box-shadow: 0 18px 28px -22px rgba(11, 77, 162, 0.9);
  transition: transform 0.16s ease, background 0.16s ease;
  animation: ci-metro-in 0.55s 0.18s ease both;
}
.ci-call:hover {
  background: #093f86;
}
.ci-call:active {
  transform: scale(0.985);
}
.ci-call-side {
  display: flex;
  width: 57px;
  flex-shrink: 0;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  gap: 4px;
  border-right: 1px solid rgba(255, 255, 255, 0.2);
  background: rgba(0, 0, 0, 0.1);
  font-size: 10px;
  font-weight: 700;
}
.ci-call-main {
  min-width: 0;
  flex: 1;
  padding: 14px 12px;
}
.ci-call-label {
  display: block;
  font-size: 9px;
  letter-spacing: 0.15em;
  color: #c3daf5;
}
.ci-call-num {
  display: block;
  margin-top: 4px;
  font-size: 29px;
  font-weight: 800;
  line-height: 1.15;
  letter-spacing: 0.01em;
  word-break: break-all;
  font-variant-numeric: tabular-nums;
}
.ci-arrow {
  display: flex;
  align-items: center;
  padding-right: 13px;
  font-size: 25px;
  font-weight: 400;
  color: #c3daf5;
}
.ci-empty {
  margin-top: 17px;
  border: 1px dashed #9db4cc;
  border-radius: 10px;
  padding: 16px;
  font-size: 14px;
  color: #5c7897;
}
.ci-bottom {
  display: flex;
  align-items: center;
  gap: 8px;
  margin-top: 19px;
  border-top: 1px solid #d5e0eb;
  padding-top: 13px;
}
.ci-exit {
  border-radius: 2px;
  background: #0b1f3a;
  padding: 3px 6px;
  color: #ffffff;
  font-size: 9px;
  font-weight: 700;
}
.ci-bottom-text {
  font-size: 10px;
  color: #5c7897;
}
@keyframes ci-metro-in {
  from { opacity: 0; transform: translateY(12px); }
  to { opacity: 1; transform: none; }
}
@media (prefers-reduced-motion: reduce) {
  .ci-top, .ci-content, .ci-call { animation: none; }
  .ci-call { transition: none; }
}
</style>
