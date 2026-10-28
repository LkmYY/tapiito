<template>
  <div class="ac-maid-root relative min-h-screen overflow-hidden px-4 py-7 text-[#5b3940]">
    <div class="ac-maid-pattern" aria-hidden="true"></div>
    <span class="ac-maid-blob ac-maid-blob-a" aria-hidden="true"></span>
    <span class="ac-maid-blob ac-maid-blob-b" aria-hidden="true"></span>
    <span class="ac-maid-spark ac-maid-spark-a" aria-hidden="true">✦</span>
    <span class="ac-maid-spark ac-maid-spark-b" aria-hidden="true">✧</span>

    <div class="relative z-10 mx-auto w-full max-w-md">
      <header class="ac-maid-top">
        <div>
          <p class="ac-maid-script">La Petite Rose</p>
          <p class="ac-maid-shop">甜梦女仆咖啡厅 · SWEET DREAM CAFÉ</p>
        </div>
        <span class="ac-maid-open"><i></i> OPEN</span>
      </header>

      <main class="ac-maid-card">
        <div class="ac-maid-frame" aria-hidden="true"></div>
        <div class="ac-maid-bow" aria-hidden="true"><i></i><i></i><b></b></div>

        <p class="ac-maid-kicker">TODAY'S SPECIAL · 本日限定</p>

        <div class="ac-maid-hero">
          <div class="ac-maid-cup" aria-hidden="true">
            <div class="ac-maid-steam"><i></i><i></i><i></i></div>
            <div class="ac-maid-cup-body"><span>♡</span></div>
            <div class="ac-maid-handle"></div>
            <div class="ac-maid-saucer"></div>
            <div class="ac-maid-plate"></div>
          </div>
          <span class="ac-maid-order">ORDER <b>#{{ nfcId || 'PREVIEW' }}</b></span>
        </div>

        <section class="ac-maid-profile">
          <p class="ac-maid-role">{{ p.title || '今日当班女仆' }}</p>
          <h1>{{ p.name || '未命名角色' }}</h1>
          <p class="ac-maid-series">出自作品 · {{ p.series || '未知作品' }}</p>
        </section>

        <div class="ac-maid-divider" aria-hidden="true"><i></i><span>❀</span><i></i></div>

        <div class="ac-maid-rows">
          <div>
            <span class="ac-maid-row-label">今日特调 / DRINK</span>
            <i aria-hidden="true"></i>
            <b>{{ p.element }}</b>
          </div>
          <div>
            <span class="ac-maid-row-label">甜度 / SWEETNESS</span>
            <i aria-hidden="true"></i>
            <b>全糖 · 100%</b>
          </div>
          <div>
            <span class="ac-maid-row-label">会员账号 / MEMBER</span>
            <i aria-hidden="true"></i>
            <b class="ac-maid-social">@{{ p.social }}</b>
          </div>
        </div>

        <blockquote class="ac-maid-quote">
          <span class="ac-maid-quote-label">店员手写 · A LITTLE NOTE</span>
          <p>{{ p.quote || '欢迎回来。今天的糖分和好心情，都为你准备好了。' }}</p>
          <i class="ac-maid-quote-heart" aria-hidden="true">♡</i>
        </blockquote>

        <footer class="ac-maid-foot">
          <div class="ac-maid-foot-copy">
            <span>THANK YOU FOR VISITING</span>
            <p>今日微笑已加入订单 · 不另收费</p>
          </div>
          <div class="ac-maid-receipt" aria-hidden="true">
            <div class="ac-maid-barcode"><i v-for="n in 19" :key="n"></i></div>
            <small>{{ nfcId || 'PREVIEW' }}</small>
          </div>
        </footer>

        <span class="ac-maid-stamp" aria-hidden="true">SWEET<br />DAY</span>
      </main>

      <p class="ac-maid-caption">☕ 请慢慢享用 · 本店禁止空腹离店 ☕</p>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps(['payload', 'nfcId'])
const p = computed(() => props.payload || {})
</script>

<style scoped>
.ac-maid-root { background: radial-gradient(ellipse at 18% 12%, rgba(255, 255, 255, 0.82), transparent 30%), radial-gradient(ellipse at 84% 74%, rgba(255, 193, 213, 0.7), transparent 38%), linear-gradient(155deg, #ffe9ef 0%, #ffdfe8 48%, #ffd5e2 100%); font-family: 'PingFang SC', 'Hiragino Sans GB', 'Microsoft YaHei', sans-serif; }
.ac-maid-pattern { position: absolute; inset: 0; background-image: radial-gradient(rgba(224, 117, 148, 0.26) 1px, transparent 1.2px), linear-gradient(90deg, transparent 49.6%, rgba(255, 255, 255, 0.35) 50%, transparent 50.4%); background-size: 21px 21px, 42px 42px; opacity: 0.6; pointer-events: none; }
.ac-maid-blob { position: absolute; border-radius: 50%; filter: blur(8px); pointer-events: none; }
.ac-maid-blob-a { top: 64px; right: -38px; width: 130px; height: 130px; background: rgba(255, 249, 241, 0.7); }
.ac-maid-blob-b { bottom: 90px; left: -44px; width: 110px; height: 110px; background: rgba(255, 177, 200, 0.35); }
.ac-maid-spark { position: absolute; color: rgba(217, 95, 131, 0.42); font-family: Georgia, serif; animation: ac-maid-twinkle 3.8s ease-in-out infinite; pointer-events: none; }
.ac-maid-spark-a { top: 24%; left: 7%; font-size: 14px; }
.ac-maid-spark-b { top: 54%; right: 6%; font-size: 18px; animation-delay: 1.4s; }
.ac-maid-top { display: flex; align-items: flex-start; justify-content: space-between; gap: 12px; animation: ac-maid-fade 0.9s both; }
.ac-maid-script { font-family: Georgia, 'Times New Roman', serif; font-size: 22px; font-style: italic; line-height: 1; letter-spacing: 0.01em; color: #b34267; text-shadow: 0 1px 0 rgba(255, 255, 255, 0.75); }
.ac-maid-shop { margin-top: 5px; font-size: 6.5px; letter-spacing: 0.14em; color: rgba(124, 69, 80, 0.55); }
.ac-maid-open { display: inline-flex; align-items: center; gap: 5px; margin-top: 2px; border: 1px solid rgba(179, 66, 103, 0.38); border-radius: 999px; background: rgba(255, 255, 255, 0.58); padding: 4px 8px; font-family: Consolas, monospace; font-size: 6.5px; letter-spacing: 0.1em; color: #a33a5e; white-space: nowrap; }
.ac-maid-open i { width: 5px; height: 5px; border-radius: 50%; background: #68a36f; box-shadow: 0 0 0 2px rgba(104, 163, 111, 0.15); animation: ac-maid-breathe 2.2s ease-in-out infinite; }
.ac-maid-card { position: relative; margin-top: 30px; border: 2px solid #fffdfa; border-radius: 30px; background: radial-gradient(circle at 18% 4%, rgba(255, 255, 255, 0.95), transparent 31%), linear-gradient(165deg, #fffdfa 0%, #fff8f4 52%, #fff0f4 100%); padding: 38px 20px 22px; box-shadow: 0 28px 54px -30px rgba(142, 62, 86, 0.7), 0 0 0 1px rgba(240, 178, 195, 0.28), inset 0 0 38px rgba(255, 226, 234, 0.5); animation: ac-maid-reveal 1s cubic-bezier(0.16, 1, 0.3, 1) both; }
.ac-maid-card::before, .ac-maid-card::after { content: ''; position: absolute; left: 13px; right: 13px; height: 14px; background: radial-gradient(circle at 7px 7px, #fffdfa 6.5px, transparent 6.8px) repeat-x; background-size: 14px 14px; pointer-events: none; }
.ac-maid-card::before { top: -9px; }
.ac-maid-card::after { bottom: -9px; }
.ac-maid-frame { position: absolute; inset: 11px; border: 1px dashed rgba(224, 143, 164, 0.42); border-radius: 23px; pointer-events: none; }
.ac-maid-bow { position: absolute; top: -25px; left: 50%; z-index: 3; width: 92px; height: 40px; transform: translateX(-50%); animation: ac-maid-sway 5.5s ease-in-out infinite; }
.ac-maid-bow i { position: absolute; top: 6px; width: 36px; height: 25px; border: 2px solid #fff8f5; border-radius: 60% 60% 46% 46%; background: linear-gradient(135deg, #ffb4c7, #ef7397); box-shadow: 0 3px 6px rgba(177, 70, 103, 0.2); }
.ac-maid-bow i:first-child { left: 3px; transform: rotate(-17deg); }
.ac-maid-bow i:nth-child(2) { right: 3px; transform: rotate(17deg); }
.ac-maid-bow b { position: absolute; top: 14px; left: 50%; width: 18px; height: 18px; border: 2px solid #fff8f5; border-radius: 5px; background: #f487a5; transform: translateX(-50%) rotate(45deg); box-shadow: 0 2px 4px rgba(177, 70, 103, 0.25); }
.ac-maid-bow::before, .ac-maid-bow::after { content: ''; position: absolute; top: 25px; width: 15px; height: 23px; background: #f78da9; clip-path: polygon(0 0, 100% 0, 100% 100%, 50% 76%, 0 100%); opacity: 0.9; }
.ac-maid-bow::before { left: 30px; transform: rotate(8deg); }
.ac-maid-bow::after { right: 30px; transform: rotate(-8deg); }
.ac-maid-kicker { text-align: center; font-family: Consolas, monospace; font-size: 6.5px; letter-spacing: 0.24em; color: #b06b7d; }
.ac-maid-hero { position: relative; display: flex; flex-direction: column; align-items: center; margin-top: 2px; }
.ac-maid-cup { position: relative; width: 130px; height: 103px; }
.ac-maid-steam { position: absolute; top: 0; left: 44px; width: 44px; height: 34px; }
.ac-maid-steam i { position: absolute; bottom: 0; width: 7px; height: 22px; border-left: 2px solid rgba(178, 104, 122, 0.4); border-radius: 50%; animation: ac-maid-steam 3.6s ease-in-out infinite; }
.ac-maid-steam i:nth-child(1) { left: 3px; animation-delay: 0.2s; }
.ac-maid-steam i:nth-child(2) { left: 18px; height: 28px; animation-delay: 1.1s; }
.ac-maid-steam i:nth-child(3) { left: 33px; animation-delay: 2s; }
.ac-maid-cup-body { position: absolute; top: 31px; left: 26px; z-index: 2; width: 64px; height: 50px; border: 2px solid #fff8f4; border-radius: 7px 7px 26px 26px; background: linear-gradient(100deg, rgba(255, 255, 255, 0.42), transparent 36%), linear-gradient(160deg, #ffb1c4, #ed7e9e 72%); box-shadow: inset 0 -8px 14px rgba(158, 67, 92, 0.13), 0 4px 8px rgba(160, 75, 99, 0.14); display: flex; align-items: center; justify-content: center; }
.ac-maid-cup-body span { font-family: Georgia, serif; font-size: 23px; line-height: 1; color: rgba(255, 251, 247, 0.92); transform: translateY(-2px); }
.ac-maid-handle { position: absolute; top: 42px; right: 24px; width: 25px; height: 23px; border: 5px solid #f08aa7; border-left: 0; border-radius: 0 15px 15px 0; box-shadow: 2px 0 0 rgba(255, 255, 255, 0.28); }
.ac-maid-saucer { position: absolute; bottom: 19px; left: 18px; z-index: 1; width: 86px; height: 13px; border-radius: 50%; background: linear-gradient(180deg, #fffdfa, #f4cbd6); box-shadow: 0 4px 5px rgba(154, 79, 100, 0.14); }
.ac-maid-plate { position: absolute; bottom: 9px; left: 29px; width: 66px; height: 9px; border-radius: 50%; background: rgba(228, 159, 178, 0.22); filter: blur(1px); }
.ac-maid-order { position: absolute; top: 56px; right: 0; border: 1px solid rgba(221, 130, 154, 0.4); border-radius: 999px; background: rgba(255, 255, 255, 0.68); padding: 3px 8px; font-family: Consolas, monospace; font-size: 5.5px; letter-spacing: 0.08em; color: #b36a7c; transform: rotate(3deg); }
.ac-maid-order b { font-weight: 600; color: #9d3b58; }
.ac-maid-profile { margin-top: 10px; text-align: center; }
.ac-maid-role { font-size: 9px; letter-spacing: 0.2em; color: #b4667b; overflow-wrap: anywhere; }
.ac-maid-profile h1 { margin-top: 4px; font-family: 'Songti SC', 'Noto Serif SC', Georgia, serif; font-size: clamp(32px, 9vw, 41px); font-weight: 700; line-height: 1.12; letter-spacing: 0.04em; color: #64343f; text-shadow: 0 1px 0 #fff; overflow-wrap: anywhere; }
.ac-maid-series { margin-top: 4px; font-size: 9.5px; letter-spacing: 0.08em; color: #9b6f78; overflow-wrap: anywhere; }
.ac-maid-divider { display: flex; align-items: center; gap: 9px; margin: 14px 0 11px; color: #d88ba0; }
.ac-maid-divider i { flex: 1; height: 1px; background: linear-gradient(90deg, transparent, rgba(213, 137, 157, 0.55)); }
.ac-maid-divider i:last-child { background: linear-gradient(90deg, rgba(213, 137, 157, 0.55), transparent); }
.ac-maid-divider span { font-size: 10px; }
.ac-maid-rows { display: grid; gap: 2px; }
.ac-maid-rows > div { display: flex; align-items: baseline; gap: 7px; min-width: 0; padding: 5px 1px; }
.ac-maid-row-label { flex: none; width: 96px; font-family: Consolas, 'PingFang SC', monospace; font-size: 6.8px; letter-spacing: 0.03em; color: #a8707f; }
.ac-maid-rows i { flex: 1; min-width: 8px; border-bottom: 1px dotted rgba(202, 136, 153, 0.6); transform: translateY(-2px); }
.ac-maid-rows b { max-width: 52%; text-align: right; font-size: 10px; font-weight: 600; line-height: 1.4; color: #6d3e47; overflow-wrap: anywhere; }
.ac-maid-rows b.ac-maid-social { color: #bd5472; }
.ac-maid-quote { position: relative; margin-top: 14px; border: 1px dashed rgba(221, 140, 160, 0.68); border-radius: 16px; background: radial-gradient(circle at 92% 8%, rgba(255, 255, 255, 0.85), transparent 26%), #fff0f3; padding: 17px 14px 14px; text-align: center; }
.ac-maid-quote-label { position: absolute; top: -6px; left: 50%; transform: translateX(-50%); background: #fff0f3; padding: 0 8px; font-family: Consolas, monospace; font-size: 5.5px; letter-spacing: 0.14em; color: #bd6b80; white-space: nowrap; }
.ac-maid-quote p { font-family: 'Kaiti SC', 'STKaiti', 'Songti SC', serif; font-size: 14px; line-height: 1.85; color: #6e3d48; overflow-wrap: anywhere; }
.ac-maid-quote-heart { position: absolute; right: 9px; bottom: 5px; font-family: Georgia, serif; font-size: 17px; font-style: normal; color: rgba(216, 123, 147, 0.45); }
.ac-maid-foot { display: flex; align-items: flex-end; justify-content: space-between; gap: 12px; margin-top: 17px; border-top: 1px solid rgba(228, 173, 187, 0.55); padding-top: 11px; }
.ac-maid-foot-copy span { font-family: Consolas, monospace; font-size: 5.5px; letter-spacing: 0.13em; color: #c08796; }
.ac-maid-foot-copy p { margin-top: 3px; font-size: 8px; letter-spacing: 0.04em; color: #a56e7c; }
.ac-maid-receipt { flex: none; text-align: right; }
.ac-maid-barcode { display: flex; justify-content: flex-end; align-items: stretch; gap: 1.5px; height: 18px; opacity: 0.56; }
.ac-maid-barcode i { flex: 0 0 2px; background: #b7657a; }
.ac-maid-barcode i:nth-child(2n) { flex-basis: 1px; }
.ac-maid-barcode i:nth-child(3n) { flex-basis: 3px; opacity: 0.6; }
.ac-maid-receipt small { display: block; margin-top: 3px; font-family: Consolas, monospace; font-size: 5px; letter-spacing: 0.08em; color: #b0868f; }
.ac-maid-stamp { position: absolute; right: 24px; bottom: 70px; display: flex; align-items: center; justify-content: center; width: 43px; height: 43px; border: 1.5px solid rgba(207, 94, 125, 0.55); border-radius: 50%; text-align: center; font-family: Consolas, monospace; font-size: 5.5px; line-height: 1.35; letter-spacing: 0.08em; color: rgba(189, 74, 106, 0.7); transform: rotate(-12deg); opacity: 0.8; }
.ac-maid-caption { margin-top: 20px; text-align: center; font-size: 8px; letter-spacing: 0.1em; color: rgba(122, 60, 76, 0.52); animation: ac-maid-fade 1.2s 0.5s both; }
@keyframes ac-maid-reveal {
  from { opacity: 0; transform: translateY(30px) scale(0.97); }
  to { opacity: 1; transform: none; }
}
@keyframes ac-maid-fade {
  from { opacity: 0; }
  to { opacity: 1; }
}
@keyframes ac-maid-steam {
  0% { opacity: 0; transform: translateY(8px) scaleX(0.9); }
  30% { opacity: 0.8; }
  100% { opacity: 0; transform: translateY(-13px) scaleX(1.3); }
}
@keyframes ac-maid-sway {
  0%, 100% { transform: translateX(-50%) rotate(-1.5deg); }
  50% { transform: translateX(-50%) rotate(1.5deg); }
}
@keyframes ac-maid-twinkle {
  0%, 100% { opacity: 0.35; transform: scale(0.85) rotate(0deg); }
  50% { opacity: 0.85; transform: scale(1.12) rotate(18deg); }
}
@keyframes ac-maid-breathe {
  0%, 100% { opacity: 1; }
  50% { opacity: 0.35; }
}
@media (prefers-reduced-motion: reduce) {
  .ac-maid-root *, .ac-maid-root *::before, .ac-maid-root *::after { animation: none !important; transition: none !important; }
}
</style>
