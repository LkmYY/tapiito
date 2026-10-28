<template>
  <div class="pt-split-root min-h-screen px-4 py-8 text-stone-900">
    <div class="pt-split-content mx-auto w-full max-w-md">
      <div class="pt-split-masthead">
        <span>THE COMPANION</span>
        <span>ISSUE 04 · 2026</span>
        <span>¥0</span>
      </div>

      <article class="pt-split-card">
        <section class="pt-split-visual">
          <div class="pt-split-pattern" aria-hidden="true"></div>
          <span class="pt-split-index">01</span>
          <span class="pt-split-vertical">A PET, A HOME, A PROMISE</span>

          <div class="pt-split-portrait" aria-hidden="true">
            <span class="pt-split-ear pt-split-ear-l"></span>
            <span class="pt-split-ear pt-split-ear-r"></span>
            <span class="pt-split-face"></span>
            <span class="pt-split-eye pt-split-eye-l"></span>
            <span class="pt-split-eye pt-split-eye-r"></span>
            <span class="pt-split-muzzle"></span>
            <span class="pt-split-nose"></span>
          </div>

          <div class="pt-split-covertext">
            <p>FAMILY MEMBER OF</p>
            <h1>{{ payload.petName || '未命名' }}</h1>
            <span>{{ payload.breed }} / {{ payload.age || '—' }}</span>
          </div>
        </section>

        <section class="pt-split-info">
          <div class="pt-split-call">
            <span class="pt-split-kicker">IF FOUND, CALL FIRST</span>
            <a :href="'tel:' + payload.phone">
              <span>{{ payload.phone }}</span>
              <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.7" aria-hidden="true">
                <path d="M4 12h16M14 6l6 6-6 6" />
              </svg>
            </a>
            <p>走失时请先拨打电话，谢谢你的善意。</p>
          </div>

          <div class="pt-split-columns">
            <div class="pt-split-col">
              <span class="pt-split-label">HOME / 家人</span>
              <p class="pt-split-owner">{{ payload.ownerName }}</p>
              <p class="pt-split-address">{{ payload.address }}</p>
            </div>
            <div class="pt-split-col pt-split-col-med">
              <span class="pt-split-label">MEDICAL / 健康</span>
              <p>{{ payload.medical || '无特殊记录' }}</p>
            </div>
          </div>

          <blockquote class="pt-split-quote">
            <span class="pt-split-quote-mark" aria-hidden="true">“</span>
            <p>{{ payload.note || '怕生，请勿强行抱抱。让我先闻闻你。' }}</p>
          </blockquote>

          <footer class="pt-split-foot">
            <span>NFC {{ nfcId || '0086' }}</span>
            <span>KEEP ME SAFE</span>
          </footer>
        </section>
      </article>

      <p class="pt-split-below">A SMALL TAG · A LONG WAY HOME</p>
    </div>
  </div>
</template>

<script setup>
defineProps(['payload', 'nfcId'])
</script>

<style scoped>
.pt-split-root {
  background:
    radial-gradient(circle at 95% 4%, rgba(194, 92, 55, 0.12), transparent 28%),
    #ebe4d7;
  font-family: 'Helvetica Neue', 'PingFang SC', sans-serif;
}
.pt-split-content { animation: pt-split-reveal 0.8s cubic-bezier(0.2, 0.8, 0.25, 1) both; }
.pt-split-masthead { display: flex; justify-content: space-between; gap: 8px; margin: 0 5px 9px; padding-bottom: 6px; border-bottom: 1px solid #b9af9e; font-family: 'Courier New', monospace; font-size: 7px; letter-spacing: 0.1em; color: #726a5c; }
.pt-split-card { overflow: hidden; background: #fdfaf4; border: 1px solid #d3c9b8; box-shadow: 0 20px 38px rgba(75, 65, 50, 0.16); transition: transform 0.35s ease, box-shadow 0.35s ease; }
.pt-split-card:hover { transform: translateY(-4px); box-shadow: 0 28px 48px rgba(75, 65, 50, 0.2); }
.pt-split-visual { position: relative; height: 305px; overflow: hidden; color: #f9f1e2; background: linear-gradient(155deg, #2f4a3c 0%, #22392e 62%, #1b2f26 100%); }
.pt-split-pattern { position: absolute; inset: 0; background: repeating-linear-gradient(115deg, rgba(249, 241, 226, 0.055) 0 1px, transparent 1px 13px); }
.pt-split-visual::after { content: ''; position: absolute; left: 0; right: 0; bottom: 0; height: 62%; background: linear-gradient(0deg, rgba(17, 35, 27, 0.82), transparent); }
.pt-split-index { position: absolute; top: 12px; left: 15px; font-family: 'Times New Roman', serif; font-size: 13px; letter-spacing: 0.12em; color: #d8c8a7; }
.pt-split-vertical { position: absolute; top: 18px; right: 15px; writing-mode: vertical-rl; font-family: 'Courier New', monospace; font-size: 6px; letter-spacing: 0.25em; color: #c7b691; }
.pt-split-portrait { position: absolute; left: 50%; top: 50px; width: 150px; height: 160px; margin-left: -75px; animation: pt-split-portrait 1s 0.12s cubic-bezier(0.2, 0.9, 0.3, 1) both; }
.pt-split-face { position: absolute; left: 20px; top: 23px; width: 110px; height: 118px; border-radius: 48% 48% 44% 44%; background: #c99762; box-shadow: inset -13px -9px 0 rgba(78, 47, 22, 0.15); }
.pt-split-ear { position: absolute; top: 2px; width: 45px; height: 70px; background: #b07d4d; border-radius: 50% 50% 20% 20%; }
.pt-split-ear-l { left: 15px; transform: rotate(-17deg); }
.pt-split-ear-r { right: 15px; transform: rotate(17deg); }
.pt-split-eye { position: absolute; top: 74px; width: 10px; height: 13px; border-radius: 50%; background: #24180e; }
.pt-split-eye-l { left: 51px; }
.pt-split-eye-r { right: 51px; }
.pt-split-muzzle { position: absolute; left: 50%; top: 96px; width: 55px; height: 39px; margin-left: -27.5px; border-radius: 50%; background: #e6c9a4; }
.pt-split-nose { position: absolute; left: 50%; top: 108px; width: 15px; height: 11px; margin-left: -7.5px; background: #2d1d10; border-radius: 50% 50% 48% 48%; }
.pt-split-covertext { position: absolute; z-index: 2; left: 17px; right: 17px; bottom: 15px; }
.pt-split-covertext p { margin: 0; font-family: 'Courier New', monospace; font-size: 7px; letter-spacing: 0.22em; color: #cfbe9a; }
.pt-split-covertext h1 { margin: 1px 0 3px; font-family: 'Times New Roman', 'Songti SC', serif; font-size: clamp(48px, 14vw, 63px); line-height: 0.95; font-weight: 400; letter-spacing: -0.01em; color: #fbf3e4; }
.pt-split-covertext span { font-family: 'Courier New', monospace; font-size: 8px; letter-spacing: 0.12em; color: #d7c4a0; }
.pt-split-info { padding: 18px 18px 14px; }
.pt-split-call { padding-bottom: 14px; border-bottom: 2px solid #2f4a3c; }
.pt-split-kicker { display: block; font-family: 'Courier New', monospace; font-size: 7px; font-weight: 700; letter-spacing: 0.16em; color: #a15b3c; }
.pt-split-call a { display: flex; align-items: center; justify-content: space-between; margin: 1px 0 2px; text-decoration: none; color: #232d26; transition: color 0.2s ease; }
.pt-split-call a span { font-family: 'Times New Roman', 'PingFang SC', serif; font-size: clamp(31px, 9.5vw, 40px); line-height: 1.1; letter-spacing: -0.015em; }
.pt-split-call a svg { width: 27px; height: 27px; flex: none; color: #2f4a3c; transition: transform 0.2s ease; }
.pt-split-call a:hover { color: #a15b3c; }
.pt-split-call a:hover svg { transform: translateX(4px); color: #a15b3c; }
.pt-split-call > p { margin: 0; font-size: 9px; letter-spacing: 0.04em; color: #7f7668; }
.pt-split-columns { display: grid; grid-template-columns: 1.1fr 0.9fr; gap: 14px; padding: 14px 0 13px; }
.pt-split-col-med { padding-left: 13px; border-left: 1px solid #ddd3c3; }
.pt-split-label { display: block; margin-bottom: 4px; font-family: 'Courier New', monospace; font-size: 7px; font-weight: 700; letter-spacing: 0.11em; color: #958a79; }
.pt-split-col p { margin: 0; font-size: 11px; line-height: 1.55; color: #4f4a40; }
.pt-split-owner { font-family: 'Times New Roman', 'Songti SC', serif; font-size: 16px !important; color: #2f4a3c !important; }
.pt-split-address { margin-top: 2px !important; font-size: 9px !important; color: #8d8374 !important; }
.pt-split-quote { position: relative; margin: 0; padding: 10px 11px 10px 26px; background: #f2ecdf; border-left: 3px solid #c0a575; }
.pt-split-quote-mark { position: absolute; left: 7px; top: 2px; font-family: Georgia, serif; font-size: 28px; line-height: 1; color: #b8a17a; }
.pt-split-quote p { margin: 0; font-family: 'Times New Roman', 'Songti SC', serif; font-size: 12px; font-style: italic; line-height: 1.7; color: #5b5346; }
.pt-split-foot { display: flex; justify-content: space-between; margin-top: 13px; font-family: 'Courier New', monospace; font-size: 7px; letter-spacing: 0.11em; color: #a29a8c; }
.pt-split-below { margin: 17px 0 0; text-align: center; font-family: 'Courier New', monospace; font-size: 8px; letter-spacing: 0.2em; color: #8d8374; }
@keyframes pt-split-reveal {
  from { opacity: 0; transform: translateY(24px); }
  to { opacity: 1; transform: translateY(0); }
}
@keyframes pt-split-portrait {
  from { opacity: 0; transform: translateY(14px) scale(0.94); }
  to { opacity: 1; transform: translateY(0) scale(1); }
}
@media (prefers-reduced-motion: reduce) {
  .pt-split-content, .pt-split-portrait { animation: none; }
  .pt-split-card, .pt-split-call a, .pt-split-call a svg { transition: none; }
}
</style>
