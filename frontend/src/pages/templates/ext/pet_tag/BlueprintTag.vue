<template>
  <div class="pt-blue-root min-h-screen px-5 py-8 text-sky-100">
    <div class="pt-blue-content mx-auto w-full max-w-md">
      <header class="pt-blue-header">
        <div>
          <p class="pt-blue-doc">SPECIFICATION SHEET</p>
          <h1>PET TAG / 宠物铭牌</h1>
        </div>
        <span class="pt-blue-rev">REV. C</span>
      </header>

      <article class="pt-blue-sheet">
        <span class="pt-blue-corner pt-blue-corner-tl" aria-hidden="true"></span>
        <span class="pt-blue-corner pt-blue-corner-tr" aria-hidden="true"></span>
        <span class="pt-blue-corner pt-blue-corner-bl" aria-hidden="true"></span>
        <span class="pt-blue-corner pt-blue-corner-br" aria-hidden="true"></span>

        <div class="pt-blue-meta">
          <span>DWG NO. {{ nfcId || 'PT-0417' }}</span>
          <span>SCALE 1:1</span>
          <span>UNIT mm</span>
        </div>

        <div class="pt-blue-drawing">
          <div class="pt-blue-dimension pt-blue-dimension-top">
            <span>32.0</span>
            <i></i>
          </div>
          <div class="pt-blue-dimension pt-blue-dimension-side">
            <span>48.0</span>
            <i></i>
          </div>
          <div class="pt-blue-tag-outline">
            <div class="pt-blue-hole-outline"></div>
            <div class="pt-blue-engrave">
              <span class="pt-blue-engrave-label">ENGRAVED FACE</span>
              <p>{{ payload.petName || '未命名' }}</p>
              <span class="pt-blue-engrave-sub">{{ payload.breed }} / {{ payload.age || '—' }}</span>
            </div>
          </div>
          <span class="pt-blue-note-arrow" aria-hidden="true">↖</span>
          <span class="pt-blue-note">A. 挂孔 Ø 4.0，兼容标准项圈环</span>
        </div>

        <div class="pt-blue-callout">
          <span class="pt-blue-callout-tag">A</span>
          <div class="pt-blue-callout-body">
            <span class="pt-blue-label">PRIMARY CONTACT / 第一联络</span>
            <a :href="'tel:' + payload.phone">{{ payload.phone }}</a>
            <span class="pt-blue-callout-note">该号码刻于铭牌正面，走失时优先拨打</span>
          </div>
        </div>

        <div class="pt-blue-table">
          <div class="pt-blue-row">
            <span>OWNER / 监护人</span>
            <p>{{ payload.ownerName }}</p>
          </div>
          <div class="pt-blue-row">
            <span>ADDRESS / 住址</span>
            <p>{{ payload.address }}</p>
          </div>
          <div class="pt-blue-row pt-blue-row-alert">
            <span>MEDICAL / 医疗备注</span>
            <p>{{ payload.medical || '无特殊记录' }}</p>
          </div>
          <div class="pt-blue-row">
            <span>NOTE / 行为提示</span>
            <p>{{ payload.note || '怕生，请勿追赶或强行抱起。' }}</p>
          </div>
        </div>

        <div class="pt-blue-dims">
          <div><span>MAT</span><strong>SUS 316L</strong></div>
          <div><span>FINISH</span><strong>BRUSHED</strong></div>
          <div><span>WEIGHT</span><strong>8.5 g</strong></div>
          <div><span>NFC</span><strong>NTAG 213</strong></div>
        </div>

        <footer class="pt-blue-titleblock">
          <div class="pt-blue-title-main">
            <span>DRAWN BY</span>
            <p>青禾工作室 / QINGHE STUDIO</p>
          </div>
          <div class="pt-blue-title-cell">
            <span>DATE</span>
            <p>2026.04</p>
          </div>
          <div class="pt-blue-title-cell">
            <span>SHEET</span>
            <p>1 / 1</p>
          </div>
        </footer>
      </article>

      <p class="pt-blue-footer">APPROVED FOR PRODUCTION · 请勿折叠</p>
    </div>
  </div>
</template>

<script setup>
defineProps(['payload', 'nfcId'])
</script>

<style scoped>
.pt-blue-root {
  background:
    radial-gradient(circle at 85% 6%, rgba(56, 189, 248, 0.14), transparent 33%),
    #062a4d;
  font-family: 'Courier New', 'PingFang SC', monospace;
}
.pt-blue-content { animation: pt-blue-scan 0.85s ease both; }
.pt-blue-header { display: flex; justify-content: space-between; align-items: flex-end; margin-bottom: 13px; padding: 0 3px; }
.pt-blue-doc { margin: 0 0 4px; font-size: 8px; letter-spacing: 0.23em; color: #7dd3fc; }
.pt-blue-header h1 { margin: 0; font-size: 19px; letter-spacing: 0.06em; font-weight: 700; color: #e0f2fe; }
.pt-blue-rev { padding: 3px 6px; border: 1px solid #38bdf8; color: #7dd3fc; font-size: 8px; letter-spacing: 0.1em; }
.pt-blue-sheet {
  position: relative;
  padding: 17px 16px 0;
  border: 1px solid rgba(125, 211, 252, 0.65);
  background:
    linear-gradient(rgba(56, 189, 248, 0.11) 1px, transparent 1px),
    linear-gradient(90deg, rgba(56, 189, 248, 0.11) 1px, transparent 1px),
    linear-gradient(rgba(56, 189, 248, 0.05) 1px, transparent 1px),
    linear-gradient(90deg, rgba(56, 189, 248, 0.05) 1px, transparent 1px),
    rgba(3, 30, 56, 0.82);
  background-size: 48px 48px, 48px 48px, 12px 12px, 12px 12px, 100% 100%;
  box-shadow: 0 22px 46px rgba(0, 0, 0, 0.35), inset 0 0 60px rgba(14, 116, 144, 0.12);
  transition: box-shadow 0.35s ease, transform 0.35s ease;
}
.pt-blue-sheet:hover { transform: translateY(-3px); box-shadow: 0 28px 56px rgba(0, 0, 0, 0.42), inset 0 0 60px rgba(14, 116, 144, 0.16); }
.pt-blue-corner { position: absolute; width: 13px; height: 13px; border-color: #7dd3fc; border-style: solid; }
.pt-blue-corner-tl { top: 4px; left: 4px; border-width: 1px 0 0 1px; }
.pt-blue-corner-tr { top: 4px; right: 4px; border-width: 1px 1px 0 0; }
.pt-blue-corner-bl { bottom: 4px; left: 4px; border-width: 0 0 1px 1px; }
.pt-blue-corner-br { bottom: 4px; right: 4px; border-width: 0 1px 1px 0; }
.pt-blue-meta { display: flex; justify-content: space-between; padding-bottom: 8px; border-bottom: 1px dashed rgba(125, 211, 252, 0.33); font-size: 7px; letter-spacing: 0.08em; color: #7dd3fc; }
.pt-blue-drawing { position: relative; height: 232px; margin: 8px 4px 10px; }
.pt-blue-dimension { position: absolute; color: #7dd3fc; font-size: 7px; }
.pt-blue-dimension i { display: block; background: rgba(125, 211, 252, 0.55); }
.pt-blue-dimension-top { top: 0; left: 26px; right: 26px; display: flex; align-items: center; gap: 5px; }
.pt-blue-dimension-top i { flex: 1; height: 1px; position: relative; }
.pt-blue-dimension-top i::before, .pt-blue-dimension-top i::after { content: ''; position: absolute; top: -3px; width: 1px; height: 7px; background: rgba(125, 211, 252, 0.55); }
.pt-blue-dimension-top i::before { left: 0; }
.pt-blue-dimension-top i::after { right: 0; }
.pt-blue-dimension-side { top: 31px; bottom: 34px; left: 2px; display: flex; align-items: center; gap: 5px; writing-mode: vertical-rl; }
.pt-blue-dimension-side i { flex: 1; width: 1px; background: rgba(125, 211, 252, 0.55); }
.pt-blue-tag-outline {
  position: absolute;
  top: 16px;
  left: 50%;
  width: 180px;
  height: 186px;
  margin-left: -90px;
  border: 1.5px solid #bae6fd;
  border-radius: 21px 21px 37px 37px;
  background: rgba(56, 189, 248, 0.04);
  box-shadow: inset 0 0 30px rgba(56, 189, 248, 0.08);
  animation: pt-blue-draw 0.9s 0.15s ease both;
}
.pt-blue-tag-outline::before { content: ''; position: absolute; inset: 12px; border: 1px dashed rgba(186, 230, 253, 0.36); border-radius: 14px 14px 29px 29px; }
.pt-blue-hole-outline { position: absolute; top: 16px; left: 50%; width: 17px; height: 17px; margin-left: -8.5px; border: 1.5px solid #bae6fd; border-radius: 50%; }
.pt-blue-hole-outline::after { content: 'Ø4.0'; position: absolute; top: -2px; left: 24px; font-size: 6px; color: #7dd3fc; }
.pt-blue-engrave { position: absolute; left: 17px; right: 17px; top: 57px; padding: 10px 8px 9px; border: 1px solid rgba(186, 230, 253, 0.5); text-align: center; }
.pt-blue-engrave-label { display: block; font-size: 6px; letter-spacing: 0.2em; color: #7dd3fc; }
.pt-blue-engrave p { margin: 5px 0 3px; font-family: 'Arial Narrow', 'PingFang SC', sans-serif; font-size: 29px; line-height: 1; font-weight: 700; letter-spacing: 0.03em; color: #e0f2fe; }
.pt-blue-engrave-sub { font-size: 7px; letter-spacing: 0.1em; color: #7dd3fc; }
.pt-blue-note-arrow { position: absolute; right: 5px; bottom: 25px; font-size: 17px; color: #7dd3fc; }
.pt-blue-note { position: absolute; right: 0; bottom: 8px; font-size: 7px; letter-spacing: 0.03em; color: #7dd3fc; }
.pt-blue-callout { display: flex; gap: 9px; margin: 2px 0 13px; padding: 9px 10px; border: 1px solid rgba(186, 230, 253, 0.55); background: rgba(8, 47, 73, 0.8); }
.pt-blue-callout-tag { flex: none; width: 18px; height: 18px; display: grid; place-items: center; border: 1px solid #bae6fd; color: #bae6fd; font-size: 8px; }
.pt-blue-callout-body { flex: 1; min-width: 0; }
.pt-blue-label { display: block; font-size: 7px; letter-spacing: 0.13em; color: #7dd3fc; }
.pt-blue-callout a { display: block; margin: 1px 0; font-family: 'Arial', 'PingFang SC', sans-serif; font-size: clamp(27px, 8vw, 35px); font-weight: 700; letter-spacing: 0.02em; color: #f0f9ff; text-decoration: none; transition: color 0.2s ease, transform 0.2s ease; }
.pt-blue-callout a:hover { color: #7dd3fc; transform: translateX(2px); }
.pt-blue-callout-note { font-size: 7px; letter-spacing: 0.03em; color: #7fb7d4; }
.pt-blue-table { border-top: 1px solid rgba(125, 211, 252, 0.4); }
.pt-blue-row { display: grid; grid-template-columns: 98px 1fr; gap: 8px; padding: 8px 0; border-bottom: 1px solid rgba(125, 211, 252, 0.23); font-size: 9px; line-height: 1.5; }
.pt-blue-row span { font-size: 7px; letter-spacing: 0.05em; color: #7dd3fc; padding-top: 1px; }
.pt-blue-row p { margin: 0; color: #dbeafe; }
.pt-blue-row-alert p { color: #fca5a5; font-weight: 700; }
.pt-blue-row-alert { background: rgba(239, 68, 68, 0.07); margin: 0 -6px; padding: 8px 6px; }
.pt-blue-dims { display: grid; grid-template-columns: repeat(4, 1fr); margin-top: 11px; border: 1px solid rgba(125, 211, 252, 0.35); }
.pt-blue-dims div { padding: 6px 5px; border-right: 1px solid rgba(125, 211, 252, 0.25); }
.pt-blue-dims div:last-child { border-right: 0; }
.pt-blue-dims span { display: block; font-size: 6px; letter-spacing: 0.1em; color: #7dd3fc; }
.pt-blue-dims strong { display: block; margin-top: 2px; font-size: 7px; letter-spacing: 0.02em; color: #dbeafe; }
.pt-blue-titleblock { display: grid; grid-template-columns: 1fr 62px 45px; margin: 12px -16px 0; border-top: 1px solid rgba(125, 211, 252, 0.5); }
.pt-blue-titleblock > div { padding: 7px 8px; border-right: 1px solid rgba(125, 211, 252, 0.3); }
.pt-blue-titleblock > div:last-child { border-right: 0; }
.pt-blue-titleblock span { display: block; font-size: 6px; letter-spacing: 0.1em; color: #7dd3fc; }
.pt-blue-titleblock p { margin: 2px 0 0; font-size: 7px; color: #dbeafe; }
.pt-blue-title-main p { font-size: 8px; }
.pt-blue-footer { margin: 14px 0 0; text-align: center; font-size: 7px; letter-spacing: 0.17em; color: #7fb7d4; }
@keyframes pt-blue-scan {
  from { opacity: 0; clip-path: inset(0 0 100% 0); }
  to { opacity: 1; clip-path: inset(0 0 0 0); }
}
@keyframes pt-blue-draw {
  from { opacity: 0; transform: scale(0.94); }
  to { opacity: 1; transform: scale(1); }
}
@media (prefers-reduced-motion: reduce) {
  .pt-blue-content, .pt-blue-tag-outline { animation: none; }
  .pt-blue-sheet, .pt-blue-callout a { transition: none; }
}
</style>
