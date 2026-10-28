<template>
  <div class="ac-back-root relative min-h-screen overflow-hidden px-4 pb-8 pt-2 text-[#101010]">
    <div class="ac-back-beams" aria-hidden="true"></div>
    <div class="ac-back-glow" aria-hidden="true"></div>
    <div class="ac-back-hazard" aria-hidden="true"></div>

    <div class="relative z-10 mx-auto w-full max-w-md">
      <div class="ac-back-lanyard">
        <div class="ac-back-strap" aria-hidden="true"><span>ALL ACCESS · BACKSTAGE</span></div>
        <div class="ac-back-clip" aria-hidden="true"><i></i></div>
      </div>

      <main class="ac-back-pass">
        <div class="ac-back-glint" aria-hidden="true"></div>

        <header class="ac-back-head">
          <span class="ac-back-brand">ARTIST<br /><b>PASS</b></span>
          <span class="ac-back-no">NO. {{ nfcId || 'PREVIEW' }}</span>
        </header>

        <div class="ac-back-stripes" aria-hidden="true"></div>

        <div class="ac-back-identity">
          <div class="ac-back-photo">
            <div class="ac-back-rays" aria-hidden="true"></div>
            <svg class="ac-back-star" viewBox="0 0 100 100" aria-hidden="true">
              <path d="M50 4 61 37 96 38 68 58 78 92 50 72 22 92 32 58 4 38 39 37Z" fill="#ffe600" stroke="#101010" stroke-width="2.5" stroke-linejoin="round" />
            </svg>
            <span class="ac-back-live" aria-hidden="true">LIVE</span>
            <span class="ac-back-photo-element">{{ p.element }}</span>
            <span class="ac-back-photo-stamp" aria-hidden="true">ALL<br />ACCESS</span>
          </div>

          <section class="ac-back-who">
            <p class="ac-back-role">{{ p.title }}</p>
            <h1>{{ p.name || '未命名角色' }}</h1>
            <p class="ac-back-series">FROM · {{ p.series || '未知作品' }}</p>
            <span class="ac-back-artist">ARTIST · 今晚出演</span>
          </section>
        </div>

        <blockquote class="ac-back-quote">
          <span class="ac-back-quote-label">MC NOTE / 后台留言</span>
          <p>“{{ p.quote || '灯亮之前，先为自己鼓掌。' }}”</p>
          <i class="ac-back-quote-star" aria-hidden="true">★</i>
        </blockquote>

        <div class="ac-back-info">
          <div><span>ACCESS</span><b>STAGE · WING A · B1</b></div>
          <div><span>SOCIAL</span><b class="ac-back-social">@{{ p.social }}</b></div>
        </div>

        <footer class="ac-back-stub">
          <div class="ac-back-barcode" aria-hidden="true"><i v-for="n in 28" :key="n"></i></div>
          <div class="ac-back-stub-meta">
            <span>PASS NO. {{ nfcId || 'PREVIEW' }}</span>
            <span>NO RE-ENTRY</span>
          </div>
        </footer>
      </main>

      <p class="ac-back-caption">★ TONIGHT ONLY · DON'T FOLD THE PASS ★</p>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps(['payload', 'nfcId'])
const p = computed(() => props.payload || {})
</script>

<style scoped>
.ac-back-root { background: radial-gradient(ellipse at 18% 18%, rgba(255, 46, 136, 0.16), transparent 42%), radial-gradient(ellipse at 84% 30%, rgba(255, 230, 0, 0.1), transparent 36%), linear-gradient(180deg, #080808 0%, #050505 62%, #0b0b0b 100%); font-family: 'PingFang SC', 'Microsoft YaHei', sans-serif; }
.ac-back-beams { position: absolute; inset: -10%; background: linear-gradient(112deg, transparent 18%, rgba(255, 46, 136, 0.15) 33%, transparent 47%), linear-gradient(248deg, transparent 20%, rgba(255, 230, 0, 0.09) 37%, transparent 52%); filter: blur(6px); animation: ac-back-breathe 6s ease-in-out infinite; pointer-events: none; }
.ac-back-glow { position: absolute; bottom: -110px; left: 50%; width: 420px; height: 300px; border-radius: 50%; background: radial-gradient(ellipse, rgba(255, 46, 136, 0.22), transparent 68%); filter: blur(28px); transform: translateX(-50%); pointer-events: none; }
.ac-back-hazard { position: absolute; left: 0; right: 0; bottom: 0; height: 12px; background: repeating-linear-gradient(135deg, #ffe600 0 12px, #0a0a0a 12px 24px); opacity: 0.35; pointer-events: none; }
.ac-back-lanyard { position: relative; z-index: 2; display: flex; flex-direction: column; align-items: center; height: 91px; }
.ac-back-strap { position: relative; display: flex; align-items: center; justify-content: center; width: 57px; height: 79px; border-left: 1px solid rgba(255, 255, 255, 0.18); border-right: 1px solid rgba(0, 0, 0, 0.5); background: repeating-linear-gradient(90deg, #ff2e88 0 9px, #101010 9px 19px); box-shadow: 0 7px 12px rgba(0, 0, 0, 0.35); }
.ac-back-strap::before { content: ''; position: absolute; inset: 0 17px; background: rgba(8, 8, 8, 0.7); }
.ac-back-strap span { position: relative; z-index: 1; writing-mode: vertical-rl; text-orientation: upright; font-family: 'Arial Black', 'PingFang SC', sans-serif; font-size: 6px; letter-spacing: 0.18em; color: #ffe600; white-space: nowrap; }
.ac-back-clip { position: relative; z-index: 3; display: grid; place-items: center; width: 49px; height: 27px; margin-top: -10px; border: 1px solid #777; border-radius: 6px; background: linear-gradient(180deg, #e8e8e8 0%, #8d8d8d 42%, #4a4a4a 100%); box-shadow: 0 3px 7px rgba(0, 0, 0, 0.55), inset 0 1px 0 rgba(255, 255, 255, 0.8); }
.ac-back-clip::before { content: ''; position: absolute; top: 5px; width: 25px; height: 7px; border: 1px solid #3e3e3e; border-radius: 999px; background: #1a1a1a; box-shadow: inset 0 1px 2px rgba(0, 0, 0, 0.7); }
.ac-back-clip i { position: absolute; bottom: -5px; width: 25px; height: 9px; border: 1px solid #666; border-radius: 0 0 4px 4px; background: #bdbdbd; }
.ac-back-pass { position: relative; z-index: 1; overflow: hidden; margin-top: -4px; border: 3px solid #101010; border-radius: 16px; background: #f7f1e6; box-shadow: 7px 7px 0 #ff2e88, -6px 6px 0 #ffe600, 0 28px 50px -20px rgba(0, 0, 0, 0.75); animation: ac-back-drop 1s cubic-bezier(0.16, 1, 0.3, 1) both; }
.ac-back-pass::before { content: ''; position: absolute; inset: 0; z-index: 3; border: 1px solid rgba(16, 16, 16, 0.22); border-radius: 12px; pointer-events: none; }
.ac-back-glint { position: absolute; top: -20%; left: -45%; z-index: 4; width: 36%; height: 140%; background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.32), transparent); transform: rotate(18deg); animation: ac-back-glint 8s ease-in-out infinite; pointer-events: none; }
.ac-back-head { display: flex; align-items: center; justify-content: space-between; gap: 8px; background: #0c0c0c; padding: 13px 15px 11px; color: #fff; }
.ac-back-brand { font-family: 'Arial Black', 'PingFang SC', sans-serif; font-size: 21px; font-weight: 900; line-height: 0.88; letter-spacing: -0.04em; color: #ff2e88; }
.ac-back-brand b { font-size: 27px; color: #fff; letter-spacing: -0.03em; }
.ac-back-no { border: 1px solid rgba(255, 230, 0, 0.5); padding: 4px 8px; font-family: Consolas, monospace; font-size: 6.5px; letter-spacing: 0.08em; color: #ffe600; white-space: nowrap; }
.ac-back-stripes { height: 10px; background: repeating-linear-gradient(135deg, #ffe600 0 14px, #ff2e88 14px 28px, #0c0c0c 28px 42px); }
.ac-back-identity { display: grid; grid-template-columns: 106px minmax(0, 1fr); gap: 13px; align-items: stretch; padding: 15px 15px 0; }
.ac-back-photo { position: relative; overflow: hidden; min-height: 128px; border: 2px solid #101010; background: #1a1a1a; box-shadow: 3px 3px 0 rgba(16, 16, 16, 0.16); }
.ac-back-rays { position: absolute; inset: -38%; background: repeating-conic-gradient(from 0deg, #ff2e88 0deg 9deg, transparent 9deg 26deg); opacity: 0.9; animation: ac-back-spin 34s linear infinite; }
.ac-back-rays::after { content: ''; position: absolute; inset: 22%; border-radius: 50%; background: #1a1a1a; box-shadow: 0 0 0 9px rgba(255, 230, 0, 0.16); }
.ac-back-star { position: absolute; top: 19px; left: 50%; z-index: 1; width: 59px; height: 59px; transform: translateX(-50%) rotate(-7deg); filter: drop-shadow(2px 3px 0 rgba(0, 0, 0, 0.35)); }
.ac-back-live { position: absolute; top: 8px; left: 8px; z-index: 2; font-family: 'Arial Black', sans-serif; font-size: 8px; letter-spacing: 0.12em; color: #fff; text-shadow: 1px 1px 0 #101010; }
.ac-back-photo-element { position: absolute; left: 0; right: 0; bottom: 0; z-index: 2; border-top: 2px solid #101010; background: #ffe600; padding: 4px 5px; text-align: center; font-size: 8px; font-weight: 700; line-height: 1.2; color: #101010; overflow-wrap: anywhere; }
.ac-back-photo-stamp { position: absolute; right: 5px; bottom: 29px; z-index: 2; display: grid; place-items: center; width: 34px; height: 34px; border: 1.5px solid #ffe600; border-radius: 50%; text-align: center; font-family: 'Arial Black', sans-serif; font-size: 4.5px; line-height: 1.2; color: #ffe600; transform: rotate(-12deg); }
.ac-back-who { min-width: 0; display: flex; flex-direction: column; }
.ac-back-role { font-size: 8px; font-weight: 700; letter-spacing: 0.1em; line-height: 1.4; color: #d6196d; overflow-wrap: anywhere; }
.ac-back-who h1 { margin-top: 5px; font-family: 'Arial Black', 'PingFang SC', 'Microsoft YaHei', sans-serif; font-size: clamp(25px, 7.1vw, 32px); font-weight: 900; line-height: 1.05; letter-spacing: -0.045em; color: #101010; overflow-wrap: anywhere; }
.ac-back-series { margin-top: 6px; font-size: 8px; letter-spacing: 0.03em; line-height: 1.45; color: #6e655c; overflow-wrap: anywhere; }
.ac-back-artist { align-self: flex-start; margin-top: auto; border: 1.5px solid #101010; background: #ff2e88; padding: 4px 7px; font-family: Consolas, 'PingFang SC', monospace; font-size: 6.5px; font-weight: 700; letter-spacing: 0.06em; color: #101010; transform: rotate(-1.5deg); }
.ac-back-quote { position: relative; margin: 16px 15px 0; border: 1.5px solid #101010; background: linear-gradient(120deg, rgba(255, 255, 255, 0.25), transparent 42%), #ffe600; padding: 16px 14px 13px; transform: rotate(-0.8deg); box-shadow: 3px 3px 0 rgba(16, 16, 16, 0.15); }
.ac-back-quote-label { position: absolute; top: -7px; left: 12px; background: #101010; padding: 2px 7px; font-family: Consolas, monospace; font-size: 5.5px; letter-spacing: 0.1em; color: #ffe600; }
.ac-back-quote p { font-family: 'Songti SC', 'Noto Serif SC', serif; font-size: 14.5px; font-weight: 600; line-height: 1.7; color: #111; overflow-wrap: anywhere; }
.ac-back-quote-star { position: absolute; right: 8px; bottom: 3px; font-family: Georgia, serif; font-size: 16px; font-style: normal; color: rgba(16, 16, 16, 0.28); }
.ac-back-info { margin: 14px 15px 0; border-top: 2px solid #101010; }
.ac-back-info > div { display: flex; align-items: baseline; justify-content: space-between; gap: 10px; border-bottom: 1px dotted rgba(16, 16, 16, 0.35); padding: 6px 0; }
.ac-back-info span { flex: none; font-family: Consolas, monospace; font-size: 6px; font-weight: 700; letter-spacing: 0.08em; color: #777; }
.ac-back-info b { min-width: 0; text-align: right; font-family: Consolas, 'PingFang SC', monospace; font-size: 8.5px; font-weight: 700; line-height: 1.4; color: #101010; overflow-wrap: anywhere; }
.ac-back-info b.ac-back-social { color: #d6196d; }
.ac-back-stub { margin-top: 15px; border-top: 2px solid #101010; background: #0c0c0c; padding: 10px 15px 12px; color: #fff; }
.ac-back-barcode { display: flex; align-items: stretch; gap: 1.5px; height: 27px; }
.ac-back-barcode i { flex: 1; background: #ffe600; }
.ac-back-barcode i:nth-child(2n) { flex: 0 0 1px; background: #ff2e88; }
.ac-back-barcode i:nth-child(3n) { flex: 0 0 3px; background: #fff; opacity: 0.8; }
.ac-back-barcode i:nth-child(5n) { flex: 0 0 1px; opacity: 0.45; }
.ac-back-stub-meta { display: flex; align-items: center; justify-content: space-between; gap: 8px; margin-top: 6px; }
.ac-back-stub-meta span { font-family: Consolas, monospace; font-size: 5.5px; letter-spacing: 0.08em; color: rgba(255, 255, 255, 0.7); overflow-wrap: anywhere; }
.ac-back-stub-meta span:last-child { color: #ff2e88; text-align: right; white-space: nowrap; }
.ac-back-caption { margin-top: 22px; text-align: center; font-family: Consolas, monospace; font-size: 6.5px; font-weight: 700; letter-spacing: 0.14em; color: rgba(255, 230, 0, 0.55); animation: ac-back-fade 1.2s 0.5s both; }
@keyframes ac-back-drop {
  from { opacity: 0; transform: translateY(-26px) rotate(-2deg); }
  60% { transform: translateY(3px) rotate(0.4deg); }
  to { opacity: 1; transform: none; }
}
@keyframes ac-back-fade {
  from { opacity: 0; }
  to { opacity: 1; }
}
@keyframes ac-back-spin {
  to { transform: rotate(360deg); }
}
@keyframes ac-back-breathe {
  0%, 100% { opacity: 0.75; }
  50% { opacity: 1; }
}
@keyframes ac-back-glint {
  0%, 64% { left: -45%; opacity: 0; }
  70% { opacity: 1; }
  88% { left: 120%; opacity: 0; }
  100% { left: 120%; opacity: 0; }
}
@media (prefers-reduced-motion: reduce) {
  .ac-back-root *, .ac-back-root *::before, .ac-back-root *::after { animation: none !important; transition: none !important; }
}
</style>
