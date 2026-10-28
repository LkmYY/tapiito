<template>
  <div class="pt-garden-root min-h-screen px-5 py-9 text-emerald-950">
    <div class="pt-garden-content mx-auto w-full max-w-md">
      <header class="pt-garden-head">
        <span>FIELD NOTES · 家庭观察册</span>
        <span>NO. {{ nfcId || 'G-086' }}</span>
      </header>

      <article class="pt-garden-sheet">
        <span class="pt-garden-corner" aria-hidden="true"></span>
        <span class="pt-garden-corner pt-garden-corner-b" aria-hidden="true"></span>

        <div class="pt-garden-plant" aria-hidden="true">
          <svg viewBox="0 0 220 250" fill="none" stroke="currentColor" stroke-width="1.4" stroke-linecap="round">
            <path d="M110 238 C 108 180, 106 130, 108 72" />
            <path d="M108 198 C 80 184, 58 158, 52 127 C 82 133, 101 157, 108 180" />
            <path d="M109 168 C 138 154, 159 128, 166 98 C 136 105, 117 128, 109 151" />
            <path d="M108 137 C 82 124, 65 101, 61 73 C 87 79, 103 99, 108 121" />
            <path d="M108 108 C 132 96, 149 75, 154 50 C 130 57, 114 76, 108 96" />
            <path d="M109 84 C 100 63, 100 42, 108 24 C 118 42, 119 63, 112 83" />
            <path d="M108 238 C 90 230, 78 217, 72 200" />
            <path d="M108 238 C 127 229, 140 216, 146 198" />
          </svg>
          <span class="pt-garden-flower pt-garden-flower-a"></span>
          <span class="pt-garden-flower pt-garden-flower-b"></span>
          <span class="pt-garden-flower pt-garden-flower-c"></span>
        </div>

        <div class="pt-garden-title">
          <span class="pt-garden-latin">COMPANION SPECIMEN</span>
          <h1>{{ payload.petName || '未命名' }}</h1>
          <p>{{ payload.breed }}<span>·</span>{{ payload.age || '—' }}</p>
        </div>

        <div class="pt-garden-note-line">
          <span>采集于</span>
          <i></i>
          <span>{{ payload.address }}</span>
        </div>

        <div class="pt-garden-label">
          <span class="pt-garden-label-head">HERBARIUM LABEL / 标本标签</span>
          <div class="pt-garden-label-grid">
            <div>
              <span>监护人</span>
              <p>{{ payload.ownerName }}</p>
            </div>
            <div>
              <span>住址</span>
              <p>{{ payload.address }}</p>
            </div>
          </div>
          <div class="pt-garden-label-med">
            <span>健康备注 / MEDICAL</span>
            <p>{{ payload.medical || '无特殊记录' }}</p>
          </div>
        </div>

        <div class="pt-garden-quote">
          <span aria-hidden="true">❀</span>
          <p>{{ payload.note || '怕生，请勿强行抱抱。像对待一株小草那样温柔。' }}</p>
        </div>

        <a class="pt-garden-phone" :href="'tel:' + payload.phone">
          <div class="pt-garden-seed">
            <span class="pt-garden-seed-top">SEED PACKET · 紧急联络</span>
            <strong>{{ payload.phone }}</strong>
            <span class="pt-garden-seed-bottom">若在花园外遇见我，请拨这个号码</span>
          </div>
          <span class="pt-garden-seed-icon" aria-hidden="true">
            <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5">
              <path d="M5 4h4l2 5-2.5 1.5a11 11 0 0 0 5 5L15 13l5 2v4a2 2 0 0 1-2 2A16 16 0 0 1 3 6a2 2 0 0 1 2-2z" />
            </svg>
          </span>
        </a>

        <footer class="pt-garden-foot">
          <span>PRESSED WITH CARE · 2026</span>
          <span class="pt-garden-sign">青禾植物社</span>
        </footer>
      </article>

      <p class="pt-garden-below">grow gently · stay close · come home</p>
    </div>
  </div>
</template>

<script setup>
defineProps(['payload', 'nfcId'])
</script>

<style scoped>
.pt-garden-root {
  background:
    radial-gradient(circle at 12% 8%, rgba(120, 152, 113, 0.2), transparent 27%),
    radial-gradient(circle at 88% 92%, rgba(192, 154, 107, 0.16), transparent 30%),
    #f1eee1;
  font-family: 'Baskerville', 'Songti SC', Georgia, serif;
}
.pt-garden-content { animation: pt-garden-unfold 0.8s cubic-bezier(0.2, 0.8, 0.25, 1) both; }
.pt-garden-head { display: flex; justify-content: space-between; margin: 0 5px 11px; font-family: 'Courier New', monospace; font-size: 7px; letter-spacing: 0.12em; color: #7f9077; }
.pt-garden-sheet {
  position: relative;
  padding: 22px 19px 15px;
  background:
    radial-gradient(circle at 22% 20%, rgba(150, 130, 83, 0.08), transparent 28%),
    repeating-linear-gradient(0deg, transparent 0 29px, rgba(116, 138, 103, 0.055) 29px 30px),
    #faf7ec;
  border: 1px solid #d3cdb9;
  box-shadow: 0 20px 36px rgba(80, 79, 55, 0.14);
  transition: transform 0.35s ease, box-shadow 0.35s ease;
}
.pt-garden-sheet:hover { transform: translateY(-3px) rotate(0.25deg); box-shadow: 0 27px 45px rgba(80, 79, 55, 0.18); }
.pt-garden-corner { position: absolute; top: 7px; left: 7px; width: 16px; height: 16px; border-top: 1.5px solid #a8b29a; border-left: 1.5px solid #a8b29a; }
.pt-garden-corner-b { top: auto; left: auto; bottom: 7px; right: 7px; border-top: 0; border-left: 0; border-right: 1.5px solid #a8b29a; border-bottom: 1.5px solid #a8b29a; }
.pt-garden-plant { position: absolute; top: 54px; right: 7px; width: 153px; color: #94a886; opacity: 0.62; pointer-events: none; }
.pt-garden-plant svg { display: block; width: 100%; height: auto; }
.pt-garden-flower { position: absolute; width: 13px; height: 13px; border-radius: 50% 40% 50% 40%; background: #c9896d; opacity: 0.7; }
.pt-garden-flower-a { top: 22px; left: 86px; transform: rotate(24deg); }
.pt-garden-flower-b { top: 56px; right: 17px; background: #d6b47e; }
.pt-garden-flower-c { top: 105px; left: 21px; width: 10px; height: 10px; background: #b9a27c; }
.pt-garden-title { position: relative; z-index: 1; padding: 2px 0 14px; max-width: 65%; }
.pt-garden-latin { display: block; font-family: 'Courier New', monospace; font-size: 6.5px; letter-spacing: 0.2em; color: #849077; }
.pt-garden-title h1 { margin: 4px 0 0; font-size: clamp(43px, 12.5vw, 57px); line-height: 1.02; font-weight: 400; letter-spacing: 0.08em; color: #394b39; }
.pt-garden-title p { margin: 8px 0 0; font-family: 'PingFang SC', 'Microsoft YaHei', sans-serif; font-size: 10px; letter-spacing: 0.13em; color: #859078; }
.pt-garden-title p span { margin: 0 6px; color: #b6bca5; }
.pt-garden-note-line { display: flex; align-items: baseline; gap: 7px; padding: 8px 0 12px; border-top: 1px solid #d7d3c1; font-family: 'Kaiti SC', 'STKaiti', serif; font-size: 10px; color: #7b866e; }
.pt-garden-note-line span:first-child { flex: none; color: #a2aa95; }
.pt-garden-note-line i { flex: 1; border-bottom: 1px dotted #b8bda8; transform: translateY(-3px); }
.pt-garden-label { position: relative; z-index: 1; padding: 11px 12px; border: 1px solid #adba9e; background: rgba(242, 243, 228, 0.72); }
.pt-garden-label-head { display: block; margin-bottom: 7px; padding-bottom: 5px; border-bottom: 1px solid #c5cdb5; font-family: 'Courier New', monospace; font-size: 6.5px; font-weight: 700; letter-spacing: 0.13em; color: #7f8f73; }
.pt-garden-label-grid { display: grid; grid-template-columns: 0.85fr 1.15fr; gap: 10px; }
.pt-garden-label-grid span { display: block; font-family: 'PingFang SC', 'Microsoft YaHei', sans-serif; font-size: 7px; letter-spacing: 0.08em; color: #9aa68f; }
.pt-garden-label-grid p { margin: 2px 0 0; font-family: 'PingFang SC', 'Microsoft YaHei', sans-serif; font-size: 10px; line-height: 1.5; color: #506049; }
.pt-garden-label-med { margin-top: 8px; padding-top: 7px; border-top: 1px dashed #c3cbb4; }
.pt-garden-label-med span { font-family: 'Courier New', monospace; font-size: 6.5px; letter-spacing: 0.09em; color: #a08b6e; }
.pt-garden-label-med p { margin: 2px 0 0; font-family: 'PingFang SC', 'Microsoft YaHei', sans-serif; font-size: 10px; line-height: 1.5; color: #8a5a3f; }
.pt-garden-quote { display: flex; gap: 8px; align-items: flex-start; margin: 11px 2px 13px; }
.pt-garden-quote span { flex: none; color: #b38a66; font-size: 13px; line-height: 1.3; }
.pt-garden-quote p { margin: 0; font-family: 'Kaiti SC', 'STKaiti', serif; font-size: 12px; line-height: 1.7; color: #58664d; }
.pt-garden-phone { position: relative; display: block; padding: 1px; text-decoration: none; }
.pt-garden-seed { position: relative; padding: 10px 45px 11px 13px; background: #e9e3ce; border: 1px solid #b9af91; box-shadow: 3px 3px 0 rgba(151, 140, 108, 0.22); transition: transform 0.2s ease, box-shadow 0.2s ease; }
.pt-garden-phone:hover .pt-garden-seed { transform: translate(-1px, -1px); box-shadow: 5px 5px 0 rgba(151, 140, 108, 0.25); }
.pt-garden-phone:active .pt-garden-seed { transform: translate(2px, 2px); box-shadow: 1px 1px 0 rgba(151, 140, 108, 0.22); }
.pt-garden-seed-top { display: block; font-family: 'Courier New', monospace; font-size: 6.5px; font-weight: 700; letter-spacing: 0.16em; color: #8a7a5b; }
.pt-garden-seed strong { display: block; margin: 1px 0; font-family: 'Avenir Next', 'PingFang SC', sans-serif; font-size: clamp(26px, 8vw, 33px); font-weight: 700; letter-spacing: 0.015em; color: #34492f; }
.pt-garden-seed-bottom { font-family: 'PingFang SC', 'Microsoft YaHei', sans-serif; font-size: 8px; letter-spacing: 0.04em; color: #8c8369; }
.pt-garden-seed-icon { position: absolute; right: 14px; top: 50%; width: 26px; height: 26px; margin-top: -13px; color: #5b7052; }
.pt-garden-seed-icon svg { width: 100%; height: 100%; }
.pt-garden-foot { display: flex; justify-content: space-between; align-items: baseline; margin-top: 14px; padding-top: 9px; border-top: 1px solid #d7d3c1; font-family: 'Courier New', monospace; font-size: 6.5px; letter-spacing: 0.08em; color: #a2a994; }
.pt-garden-sign { font-family: 'Kaiti SC', 'KaiTi', serif; font-size: 11px; color: #8b987c; }
.pt-garden-below { margin: 17px 0 0; text-align: center; font-family: 'Georgia', serif; font-size: 10px; font-style: italic; letter-spacing: 0.08em; color: #93a084; }
@keyframes pt-garden-unfold {
  from { opacity: 0; transform: perspective(800px) rotateX(-8deg) translateY(18px); }
  to { opacity: 1; transform: perspective(800px) rotateX(0) translateY(0); }
}
@media (prefers-reduced-motion: reduce) {
  .pt-garden-content { animation: none; }
  .pt-garden-sheet, .pt-garden-seed { transition: none; }
}
</style>
