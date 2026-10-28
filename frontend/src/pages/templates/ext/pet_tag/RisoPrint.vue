<template>
  <div class="pt-riso-root min-h-screen px-5 py-8 text-neutral-900">
    <div class="pt-riso-content mx-auto w-full max-w-md">
      <div class="pt-riso-top">
        <span>RISOGRAPH 2C · 青禾印社</span>
        <span>NO. {{ nfcId || '086' }}</span>
      </div>

      <article class="pt-riso-sheet">
        <div class="pt-riso-grid" aria-hidden="true"></div>
        <span class="pt-riso-ink pt-riso-ink-cyan" aria-hidden="true"></span>
        <span class="pt-riso-ink pt-riso-ink-pink" aria-hidden="true"></span>

        <header class="pt-riso-header">
          <p class="pt-riso-mono">LOST PET / EMERGENCY PRINT</p>
          <span class="pt-riso-edition">ED. 30</span>
        </header>

        <div class="pt-riso-title">
          <span class="pt-riso-big">FOUND</span>
          <h1>{{ payload.petName || '未命名' }}</h1>
          <div class="pt-riso-tags">
            <span>{{ payload.breed }}</span>
            <span>{{ payload.age || '—' }}</span>
            <span>MICROCHIP {{ nfcId || '086' }}</span>
          </div>
        </div>

        <div class="pt-riso-poster">
          <div class="pt-riso-half pt-riso-half-pink" aria-hidden="true">
            <div class="pt-riso-pet">
              <span class="pt-riso-ear pt-riso-ear-l"></span>
              <span class="pt-riso-ear pt-riso-ear-r"></span>
              <span class="pt-riso-face"></span>
              <span class="pt-riso-eye pt-riso-eye-l"></span>
              <span class="pt-riso-eye pt-riso-eye-r"></span>
              <span class="pt-riso-snout"></span>
            </div>
            <span class="pt-riso-halftone"></span>
          </div>
          <div class="pt-riso-half pt-riso-half-cyan">
            <p class="pt-riso-half-label">CALL IMMEDIATELY</p>
            <a :href="'tel:' + payload.phone">{{ payload.phone }}</a>
            <p class="pt-riso-half-sub">A REWARD AWAITS<br />必有重谢</p>
          </div>
        </div>

        <div class="pt-riso-body">
          <div class="pt-riso-col">
            <span class="pt-riso-label">OWNER</span>
            <p>{{ payload.ownerName }}</p>
            <span class="pt-riso-label">ADDRESS</span>
            <p>{{ payload.address }}</p>
          </div>
          <div class="pt-riso-col pt-riso-col-box">
            <span class="pt-riso-label">MEDICAL / 健康</span>
            <p class="pt-riso-med">{{ payload.medical || '无特殊记录' }}</p>
          </div>
        </div>

        <p class="pt-riso-note">{{ payload.note || '怕生，请勿强行抱抱。请先联系我家人。' }}</p>

        <div class="pt-riso-footer">
          <span class="pt-riso-mark">IF FOUND<br />PLEASE CALL</span>
          <div class="pt-riso-small">
            <span>PRINTED WITH CARE</span>
            <span>青禾社区 · 宠物互助</span>
          </div>
        </div>
      </article>

      <div class="pt-riso-reg" aria-hidden="true">
        <span class="pt-riso-reg-pink"></span>
        <span class="pt-riso-reg-cyan"></span>
        <span class="pt-riso-reg-mark">®</span>
      </div>

      <p class="pt-riso-below">TWO INKS · ONE HOME · 请帮助我回家</p>
    </div>
  </div>
</template>

<script setup>
defineProps(['payload', 'nfcId'])
</script>

<style scoped>
.pt-riso-root {
  background:
    repeating-linear-gradient(0deg, rgba(0, 0, 0, 0.04) 0 1px, transparent 1px 4px),
    #e7e1d4;
  font-family: 'Arial Black', 'PingFang SC', sans-serif;
}
.pt-riso-content { animation: pt-riso-settle 0.72s cubic-bezier(0.2, 0.8, 0.25, 1) both; }
.pt-riso-top { display: flex; justify-content: space-between; margin: 0 5px 9px; font-family: 'Courier New', monospace; font-size: 7px; font-weight: 700; letter-spacing: 0.1em; color: #5f5a50; }
.pt-riso-sheet { position: relative; padding: 17px 16px 14px; background: #f4efe3; border: 1px solid #bdb4a2; box-shadow: 8px 9px 0 rgba(40, 35, 28, 0.15), 0 19px 34px rgba(60, 52, 40, 0.12); overflow: hidden; transition: transform 0.3s ease; }
.pt-riso-sheet:hover { transform: translate(-2px, -2px); }
.pt-riso-grid { position: absolute; inset: 0; background-image: linear-gradient(rgba(20, 20, 20, 0.05) 1px, transparent 1px), linear-gradient(90deg, rgba(20, 20, 20, 0.05) 1px, transparent 1px); background-size: 24px 24px; }
.pt-riso-ink { position: absolute; border-radius: 50%; filter: blur(30px); pointer-events: none; opacity: 0.3; }
.pt-riso-ink-cyan { width: 170px; height: 170px; top: -70px; right: -50px; background: #22d3ee; }
.pt-riso-ink-pink { width: 145px; height: 145px; bottom: 15%; left: -70px; background: #ec4899; }
.pt-riso-header { position: relative; display: flex; justify-content: space-between; padding-bottom: 7px; border-bottom: 3px solid #1a1a1a; }
.pt-riso-mono { margin: 0; font-family: 'Courier New', monospace; font-size: 7px; font-weight: 700; letter-spacing: 0.14em; }
.pt-riso-edition { font-family: 'Courier New', monospace; font-size: 7px; padding: 2px 4px; background: #ec4899; color: #fff; }
.pt-riso-title { position: relative; padding: 11px 0 8px; }
.pt-riso-big { display: block; font-size: clamp(26px, 8.5vw, 35px); line-height: 0.9; letter-spacing: -0.03em; color: #06b6d4; -webkit-text-stroke: 1px #1a1a1a; text-shadow: 3px 2px 0 #ec4899; }
.pt-riso-title h1 { margin: -2px 0 0; font-size: clamp(58px, 18vw, 78px); line-height: 0.88; letter-spacing: -0.05em; color: #1a1a1a; text-shadow: 4px 4px 0 rgba(236, 72, 153, 0.55); }
.pt-riso-tags { display: flex; flex-wrap: wrap; gap: 4px; margin-top: 8px; }
.pt-riso-tags span { padding: 3px 6px; border: 1.5px solid #1a1a1a; font-family: 'Courier New', monospace; font-size: 7px; font-weight: 700; letter-spacing: 0.05em; background: #f4efe3; }
.pt-riso-tags span:nth-child(2) { background: #a5f3fc; }
.pt-riso-tags span:nth-child(3) { background: #fbcfe8; }
.pt-riso-poster { position: relative; display: grid; grid-template-columns: 1fr 1.18fr; margin-top: 7px; border: 2.5px solid #1a1a1a; min-height: 177px; }
.pt-riso-half { position: relative; overflow: hidden; }
.pt-riso-half-pink { background: #f9a8d4; }
.pt-riso-half-cyan { display: flex; flex-direction: column; justify-content: center; padding: 12px 11px; background: #67e8f9; border-left: 2.5px solid #1a1a1a; }
.pt-riso-pet { position: absolute; left: 50%; bottom: 0; width: 104px; height: 121px; margin-left: -52px; }
.pt-riso-face { position: absolute; left: 14px; top: 25px; width: 76px; height: 84px; background: #1a1a1a; border-radius: 45% 45% 42% 42%; }
.pt-riso-face::after { content: ''; position: absolute; left: 11px; top: 38px; width: 54px; height: 36px; border-radius: 50%; background: #f4efe3; opacity: 0.9; }
.pt-riso-ear { position: absolute; top: 5px; width: 29px; height: 42px; background: #1a1a1a; border-radius: 50% 50% 18% 18%; }
.pt-riso-ear-l { left: 5px; transform: rotate(-18deg); }
.pt-riso-ear-r { right: 5px; transform: rotate(18deg); }
.pt-riso-eye { position: absolute; z-index: 2; top: 55px; width: 7px; height: 9px; border-radius: 50%; background: #f4efe3; }
.pt-riso-eye-l { left: 32px; }
.pt-riso-eye-r { right: 32px; }
.pt-riso-snout { position: absolute; z-index: 2; left: 50%; top: 77px; width: 10px; height: 7px; margin-left: -5px; background: #1a1a1a; border-radius: 50%; }
.pt-riso-halftone { position: absolute; inset: 0; background-image: radial-gradient(rgba(26, 26, 26, 0.3) 0.9px, transparent 1px); background-size: 5px 5px; mix-blend-mode: multiply; }
.pt-riso-half-label { margin: 0; font-family: 'Courier New', monospace; font-size: 6.5px; font-weight: 700; letter-spacing: 0.1em; }
.pt-riso-half a { display: block; margin: 3px 0; font-size: clamp(21px, 6.3vw, 26px); line-height: 1.1; font-weight: 900; letter-spacing: -0.03em; color: #1a1a1a; text-decoration: none; word-break: break-word; transition: color 0.18s ease; }
.pt-riso-half a:hover { color: #db2777; }
.pt-riso-half-sub { margin: 0; font-family: 'Courier New', monospace; font-size: 7px; font-weight: 700; line-height: 1.5; }
.pt-riso-body { display: grid; grid-template-columns: 1.1fr 0.9fr; gap: 11px; margin-top: 12px; }
.pt-riso-col { padding-top: 2px; }
.pt-riso-col-box { padding: 7px 8px; border: 2px solid #1a1a1a; background: #fce7f3; }
.pt-riso-label { display: block; margin: 5px 0 1px; font-family: 'Courier New', monospace; font-size: 6.5px; font-weight: 700; letter-spacing: 0.11em; color: #6b6357; }
.pt-riso-col-box .pt-riso-label { margin-top: 0; }
.pt-riso-col p { margin: 0; font-family: 'PingFang SC', 'Microsoft YaHei', sans-serif; font-size: 10px; font-weight: 700; line-height: 1.5; color: #262626; }
.pt-riso-col .pt-riso-med { color: #a1154d; }
.pt-riso-note { margin: 11px 0 12px; padding: 8px 9px; border-top: 2px solid #1a1a1a; border-bottom: 2px solid #1a1a1a; font-family: 'PingFang SC', 'Microsoft YaHei', sans-serif; font-size: 10px; line-height: 1.6; font-weight: 600; color: #333; }
.pt-riso-footer { display: flex; align-items: center; gap: 12px; }
.pt-riso-mark { flex: none; padding: 7px 9px; background: #ec4899; color: #fff; font-size: 10px; line-height: 1.15; letter-spacing: 0.02em; transform: rotate(-2deg); box-shadow: 2px 2px 0 #1a1a1a; }
.pt-riso-small { display: flex; flex-direction: column; gap: 2px; font-family: 'Courier New', monospace; font-size: 6.5px; letter-spacing: 0.08em; color: #5f5a50; }
.pt-riso-reg { position: relative; height: 28px; margin: 9px 3px 0; }
.pt-riso-reg span { position: absolute; top: 6px; height: 12px; width: 52px; }
.pt-riso-reg-pink { left: 0; background: #ec4899; opacity: 0.7; }
.pt-riso-reg-cyan { left: 5px; background: #22d3ee; opacity: 0.7; }
.pt-riso-reg-mark { position: absolute; right: 0; top: 0; font-family: 'Courier New', monospace; font-size: 13px; color: #524b3f; }
.pt-riso-below { margin: 3px 0 0; text-align: center; font-family: 'Courier New', monospace; font-size: 8px; font-weight: 700; letter-spacing: 0.09em; color: #5f5a50; }
@keyframes pt-riso-settle {
  0% { opacity: 0; transform: translate(8px, 8px); filter: blur(1px); }
  100% { opacity: 1; transform: translate(0, 0); filter: blur(0); }
}
@media (prefers-reduced-motion: reduce) {
  .pt-riso-content { animation: none; }
  .pt-riso-sheet, .pt-riso-half a { transition: none; }
}
</style>
