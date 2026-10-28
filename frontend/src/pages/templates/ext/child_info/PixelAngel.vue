<!-- 求助信息卡 · 像素守护：像素游戏对话框，用熟悉的界面降低孩子的紧张 -->
<template>
  <div class="ci-pixel min-h-screen">
    <div class="mx-auto w-full max-w-md px-6 pb-14 pt-10">
      <div class="ci-hud">
        <span class="ci-hud-item"><Heart :size="13" /> CARE POINT</span>
        <span v-if="nfcId" class="ci-hud-item">ID {{ nfcId }}</span>
      </div>

      <div class="ci-dialog">
        <div class="ci-namebox">
          <span class="ci-avatar" aria-hidden="true"><span class="ci-avatar-face">:)</span></span>
          <div>
            <p class="ci-speaker">守护精灵</p>
            <p class="ci-speech">我迷路了，可以帮我打电话给家人吗？</p>
          </div>
        </div>

        <div class="ci-stats">
          <div class="ci-stat">
            <span class="ci-stat-key">基本情况</span>
            <p class="ci-stat-val">{{ payload.baseInfo }}</p>
          </div>
          <div class="ci-stat">
            <span class="ci-stat-key">学校 / 单位</span>
            <p class="ci-stat-val">{{ payload.school }}</p>
          </div>
          <div class="ci-stat">
            <span class="ci-stat-key">过敏史 / 病史</span>
            <p class="ci-stat-val ci-stat-val--warn">{{ payload.allergies || '无' }}</p>
          </div>
        </div>

        <a
          v-if="payload.emergencyPhone"
          :href="'tel:' + payload.emergencyPhone"
          class="ci-call"
        >
          <span class="ci-call-top">
            <span class="ci-call-label">▶ 紧急联系电话</span>
            <Phone :size="17" />
          </span>
          <span class="ci-call-num">{{ payload.emergencyPhone }}</span>
          <span class="ci-call-tap">按此拨号 · CALL</span>
        </a>

        <p class="ci-footer">▶ 请陪我等到家人出现。谢谢玩家！</p>
      </div>

      <p class="ci-credit">SAVE THE DAY · 一次善意的存档</p>
    </div>
  </div>
</template>

<script setup>
import { Heart, Phone } from 'lucide-vue-next'

defineProps(['payload', 'nfcId'])
</script>

<style scoped>
.ci-pixel {
  background: #ece8ff;
  color: #241d3f;
  font-family: "Cascadia Mono", "Courier New", monospace;
}
.ci-hud {
  display: flex;
  justify-content: space-between;
  gap: 8px;
  animation: ci-pixel-in 0.45s steps(4) both;
}
.ci-hud-item {
  display: inline-flex;
  align-items: center;
  gap: 4px;
  border: 2px solid #241d3f;
  background: #ffffff;
  padding: 5px 7px;
  font-size: 8px;
  font-weight: 700;
  letter-spacing: 0.05em;
  box-shadow: 3px 3px 0 #a79ce0;
}
.ci-dialog {
  margin-top: 17px;
  border: 3px solid #241d3f;
  background: #fbf9ff;
  padding: 16px;
  box-shadow: 7px 7px 0 #a79ce0;
  animation: ci-pixel-in 0.5s 0.05s steps(4) both;
}
.ci-namebox {
  display: flex;
  align-items: flex-start;
  gap: 10px;
  border: 2px solid #241d3f;
  background: #ffffff;
  padding: 10px;
}
.ci-avatar {
  display: flex;
  height: 43px;
  width: 43px;
  flex-shrink: 0;
  align-items: center;
  justify-content: center;
  border: 2px solid #241d3f;
  background: #ffd98a;
  box-shadow: 3px 3px 0 #241d3f;
}
.ci-avatar-face {
  font-size: 17px;
  font-weight: 700;
  letter-spacing: -1px;
}
.ci-speaker {
  font-size: 9px;
  font-weight: 700;
  letter-spacing: 0.08em;
  color: #6d5bd0;
}
.ci-speech {
  margin-top: 3px;
  font-size: 12px;
  font-weight: 700;
  line-height: 1.65;
  color: #241d3f;
}
.ci-stats {
  margin-top: 13px;
  display: flex;
  flex-direction: column;
  gap: 7px;
}
.ci-stat {
  border: 2px solid #241d3f;
  background: #ffffff;
  padding: 8px 9px;
}
.ci-stat-key {
  display: block;
  font-size: 7.5px;
  font-weight: 700;
  letter-spacing: 0.1em;
  color: #6f66a0;
}
.ci-stat-val {
  margin-top: 3px;
  font-size: 12px;
  font-weight: 700;
  line-height: 1.55;
  color: #241d3f;
}
.ci-stat-val--warn { color: #b03a1e; }
.ci-call {
  display: block;
  margin-top: 14px;
  border: 3px solid #1c1533;
  background: #2b2154;
  padding: 12px 13px;
  color: #f5f1ff;
  text-decoration: none;
  box-shadow: 5px 5px 0 #a79ce0;
  transition: transform 0.08s steps(2), box-shadow 0.08s steps(2);
  animation: ci-pixel-in 0.5s 0.12s steps(4) both;
}
.ci-call:hover {
  background: #392b6e;
}
.ci-call:active {
  transform: translate(4px, 4px);
  box-shadow: 1px 1px 0 #a79ce0;
}
.ci-call-top {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 8px;
}
.ci-call-label {
  font-size: 8.5px;
  font-weight: 700;
  letter-spacing: 0.08em;
  color: #cdc3ff;
}
.ci-call-num {
  display: block;
  margin-top: 6px;
  font-size: 28px;
  font-weight: 700;
  line-height: 1.18;
  letter-spacing: 0.02em;
  word-break: break-all;
  font-variant-numeric: tabular-nums;
  color: #ffffff;
}
.ci-call-tap {
  display: block;
  margin-top: 5px;
  font-size: 8px;
  color: #a79ce0;
}
.ci-empty {
  margin-top: 14px;
  border: 2px dashed #b03a1e;
  background: #fff0eb;
  padding: 11px;
  font-size: 11px;
  font-weight: 700;
  color: #b03a1e;
}
.ci-footer {
  margin-top: 13px;
  border-top: 2px dashed #cfc6ef;
  padding-top: 9px;
  font-size: 8.5px;
  font-weight: 700;
  line-height: 1.7;
  color: #5b4f8c;
  animation: ci-pixel-in 0.5s 0.18s steps(4) both;
}
.ci-credit {
  margin-top: 13px;
  text-align: center;
  font-size: 7px;
  font-weight: 700;
  letter-spacing: 0.12em;
  color: #8d81bd;
  animation: ci-pixel-in 0.5s 0.24s steps(4) both;
}
@keyframes ci-pixel-in {
  from { opacity: 0; transform: translateY(9px); }
  to { opacity: 1; transform: none; }
}
@media (prefers-reduced-motion: reduce) {
  .ci-hud, .ci-dialog, .ci-call, .ci-footer, .ci-credit { animation: none; }
  .ci-call { transition: none; }
}
</style>
