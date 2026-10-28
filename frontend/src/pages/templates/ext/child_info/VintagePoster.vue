<!-- 求助信息卡 · 复古告示：旧式印刷海报，红黑双色与网点质感 -->
<template>
  <div class="ci-poster min-h-screen">
    <div class="mx-auto w-full max-w-md px-5 pb-12 pt-7">
      <div class="ci-print">
        <div class="ci-marks" aria-hidden="true">
          <span class="ci-cross ci-cross--tl">+</span>
          <span class="ci-cross ci-cross--tr">+</span>
          <span class="ci-cross ci-cross--bl">+</span>
          <span class="ci-cross ci-cross--br">+</span>
        </div>

        <header class="ci-masthead">
          <span>COMMUNITY NOTICE</span>
          <span>NO. 072</span>
        </header>

        <div class="ci-headline">
          <p class="ci-over">请留意 · PLEASE NOTICE</p>
          <h1>帮我<br />回家</h1>
          <span class="ci-seal">求助</span>
        </div>

        <div class="ci-band">
          <span>如遇此人，请先拨打电话</span>
          <span class="ci-band-en">CALL FIRST</span>
        </div>

        <section class="ci-info">
          <div class="ci-info-main">
            <span class="ci-label">基本情况</span>
            <p class="ci-name">{{ payload.baseInfo }}</p>
          </div>
          <div class="ci-info-grid">
            <div>
              <span class="ci-label">学校 / 单位</span>
              <p>{{ payload.school }}</p>
            </div>
            <div>
              <span class="ci-label">过敏史 / 病史</span>
              <p class="ci-red">{{ payload.allergies || '无' }}</p>
            </div>
          </div>
        </section>

        <a
          v-if="payload.emergencyPhone"
          :href="'tel:' + payload.emergencyPhone"
          class="ci-call"
        >
          <span class="ci-call-label">紧急联系电话 · EMERGENCY</span>
          <span class="ci-call-num">{{ payload.emergencyPhone }}</span>
          <span class="ci-call-cta"><Phone :size="15" /> 点击号码拨打</span>
        </a>

        <footer class="ci-footer">
          <span>请陪伴在旁 · 等家人到达</span>
          <span>THANK YOU</span>
        </footer>
      </div>
    </div>
  </div>
</template>

<script setup>
import { Phone } from 'lucide-vue-next'

defineProps(['payload', 'nfcId'])
</script>

<style scoped>
.ci-poster {
  background: #d9c9ab;
  color: #1c1917;
  font-family: "Arial Black", "PingFang SC", "Microsoft YaHei", sans-serif;
}
.ci-print {
  position: relative;
  overflow: hidden;
  border: 1px solid #b9a685;
  background-color: #f1e5cd;
  background-image:
    radial-gradient(rgba(28, 25, 23, 0.14) 0.7px, transparent 0.8px),
    linear-gradient(160deg, rgba(255, 255, 255, 0.28), transparent 55%);
  background-size: 5px 5px, 100% 100%;
  padding: 17px 19px 19px;
  box-shadow: 0 28px 50px -30px rgba(60, 45, 24, 0.7);
  animation: ci-poster-in 0.6s ease both;
}
.ci-marks span {
  position: absolute;
  font-family: ui-monospace, Consolas, monospace;
  font-size: 13px;
  color: #a63a1e;
  opacity: 0.7;
}
.ci-cross--tl { left: 7px; top: 5px; }
.ci-cross--tr { right: 7px; top: 5px; }
.ci-cross--bl { left: 7px; bottom: 5px; }
.ci-cross--br { right: 7px; bottom: 5px; }
.ci-masthead {
  display: flex;
  justify-content: space-between;
  border-top: 2px solid #1c1917;
  border-bottom: 1px solid #1c1917;
  padding: 6px 1px;
  font-family: ui-monospace, Consolas, monospace;
  font-size: 8px;
  font-weight: 700;
  letter-spacing: 0.14em;
}
.ci-headline {
  position: relative;
  padding-top: 15px;
}
.ci-over {
  font-size: 10px;
  font-weight: 700;
  letter-spacing: 0.16em;
  color: #a63a1e;
}
.ci-headline h1 {
  margin-top: 6px;
  font-size: 58px;
  font-weight: 900;
  line-height: 0.96;
  letter-spacing: -0.05em;
  color: #1c1917;
}
.ci-seal {
  position: absolute;
  right: 1px;
  bottom: 5px;
  display: flex;
  height: 57px;
  width: 57px;
  align-items: center;
  justify-content: center;
  border: 2px solid #a63a1e;
  color: #a63a1e;
  font-size: 15px;
  font-weight: 800;
  letter-spacing: 0.08em;
  transform: rotate(-9deg);
  animation: ci-stamp-in 0.45s 0.35s ease both;
}
.ci-band {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 8px;
  margin-top: 13px;
  background: #1c1917;
  padding: 8px 10px;
  color: #f1e5cd;
}
.ci-band span:first-child { font-size: 11px; font-weight: 700; }
.ci-band-en {
  font-family: ui-monospace, Consolas, monospace;
  font-size: 8px;
  letter-spacing: 0.16em;
  color: #d9a27a;
}
.ci-info {
  border-bottom: 2px solid #1c1917;
  padding: 13px 0;
  animation: ci-poster-in 0.55s 0.11s ease both;
}
.ci-label {
  display: block;
  font-family: ui-monospace, Consolas, monospace;
  font-size: 8px;
  font-weight: 700;
  letter-spacing: 0.14em;
  color: #78664d;
}
.ci-name {
  margin-top: 4px;
  font-size: 20px;
  font-weight: 800;
  line-height: 1.45;
}
.ci-info-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 11px;
  margin-top: 12px;
  border-top: 1px solid #c3b190;
  padding-top: 10px;
}
.ci-info-grid p {
  margin-top: 3px;
  font-size: 12px;
  font-weight: 700;
  line-height: 1.5;
}
.ci-info-grid .ci-red { color: #a63a1e; }
.ci-call {
  display: block;
  margin-top: 15px;
  background: #1c1917;
  padding: 14px 14px 13px;
  color: #f5ead6;
  text-decoration: none;
  transition: transform 0.18s ease, background 0.18s ease;
  animation: ci-poster-in 0.55s 0.18s ease both;
}
.ci-call:hover { background: #352f2a; }
.ci-call:active { transform: scale(0.985); }
.ci-call-label {
  display: block;
  font-family: ui-monospace, Consolas, monospace;
  font-size: 8px;
  letter-spacing: 0.14em;
  color: #e3b28c;
}
.ci-call-num {
  display: block;
  margin-top: 5px;
  font-size: 34px;
  font-weight: 900;
  line-height: 1.14;
  letter-spacing: 0.01em;
  word-break: break-all;
  font-variant-numeric: tabular-nums;
  color: #ffffff;
}
.ci-call-cta {
  display: flex;
  align-items: center;
  gap: 5px;
  margin-top: 5px;
  font-family: ui-monospace, Consolas, monospace;
  font-size: 9px;
  color: #c9a98a;
}
.ci-empty {
  margin-top: 15px;
  border: 2px dashed #1c1917;
  padding: 15px;
  font-size: 13px;
  font-weight: 700;
}
.ci-footer {
  display: flex;
  justify-content: space-between;
  gap: 9px;
  margin-top: 13px;
  font-family: ui-monospace, Consolas, monospace;
  font-size: 8px;
  letter-spacing: 0.1em;
  color: #6f5e46;
}
@keyframes ci-poster-in {
  from { opacity: 0; transform: translateY(15px); }
  to { opacity: 1; transform: none; }
}
@keyframes ci-stamp-in {
  from { opacity: 0; transform: rotate(-9deg) scale(1.6); }
  to { opacity: 1; transform: rotate(-9deg) scale(1); }
}
@media (prefers-reduced-motion: reduce) {
  .ci-print, .ci-info, .ci-call, .ci-seal { animation: none; }
  .ci-call { transition: none; }
}
</style>
