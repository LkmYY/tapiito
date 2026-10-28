<template>
  <div class="pt-pass-root min-h-screen px-5 py-8 text-sky-50">
    <div class="pt-pass-content mx-auto w-full max-w-md">
      <div class="pt-pass-book">
        <div class="pt-pass-spine" aria-hidden="true"><span v-for="n in 12" :key="n"></span></div>

        <article class="pt-pass-page">
          <header class="pt-pass-head">
            <div class="pt-pass-emblem" aria-hidden="true">
              <svg viewBox="0 0 48 48" fill="none">
                <circle cx="24" cy="24" r="21" stroke="currentColor" stroke-width="1.2" />
                <ellipse cx="24" cy="29" rx="8" ry="6.5" fill="currentColor" />
                <ellipse cx="13" cy="21" rx="3.2" ry="4" fill="currentColor" />
                <ellipse cx="21" cy="16" rx="3.2" ry="4.3" fill="currentColor" />
                <ellipse cx="30" cy="16.5" rx="3.2" ry="4" fill="currentColor" />
                <ellipse cx="37" cy="22" rx="2.8" ry="3.6" fill="currentColor" />
              </svg>
            </div>
            <div>
              <p class="pt-pass-country">PET REPUBLIC · 宠物共和国</p>
              <h1>PET<br />PASSPORT</h1>
            </div>
            <span class="pt-pass-type">P<br /><small>P</small></span>
          </header>

          <div class="pt-pass-rule"><span>护照 / PASSPORT</span><span>NO. {{ nfcId || 'P-0086' }}</span></div>

          <div class="pt-pass-identity">
            <div class="pt-pass-portrait">
              <div class="pt-pass-portrait-art" aria-hidden="true">
                <span class="pt-pass-ear pt-pass-ear-l"></span>
                <span class="pt-pass-ear pt-pass-ear-r"></span>
                <span class="pt-pass-face"></span>
                <span class="pt-pass-eye pt-pass-eye-l"></span>
                <span class="pt-pass-eye pt-pass-eye-r"></span>
                <span class="pt-pass-snout"></span>
              </div>
              <span class="pt-pass-portrait-label">PHOTO</span>
            </div>
            <div class="pt-pass-fields">
              <div class="pt-pass-field pt-pass-field-wide">
                <span>姓名 / NAME</span>
                <strong>{{ payload.petName || '未命名' }}</strong>
              </div>
              <div class="pt-pass-field">
                <span>品种 / BREED</span>
                <strong>{{ payload.breed }}</strong>
              </div>
              <div class="pt-pass-field">
                <span>年龄 / AGE</span>
                <strong>{{ payload.age || '—' }}</strong>
              </div>
              <div class="pt-pass-field">
                <span>监护人 / OWNER</span>
                <strong>{{ payload.ownerName }}</strong>
              </div>
              <div class="pt-pass-field">
                <span>住址 / ADDRESS</span>
                <strong>{{ payload.address }}</strong>
              </div>
            </div>
          </div>

          <div class="pt-pass-medical">
            <div class="pt-pass-medical-head">
              <span>医疗备注 / MEDICAL</span>
              <span class="pt-pass-cross" aria-hidden="true">✚</span>
            </div>
            <p>{{ payload.medical || '无特殊记录' }}</p>
          </div>

          <div class="pt-pass-note">
            <span>提示 / NOTE</span>
            <p>{{ payload.note || '请温柔对待我，并帮助我回到家人身边。' }}</p>
          </div>

          <a class="pt-pass-phone" :href="'tel:' + payload.phone">
            <div class="pt-pass-phone-top">
              <span>紧急联络 / EMERGENCY</span>
              <span class="pt-pass-phone-icon" aria-hidden="true">
                <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M5 4h4l2 5-2.5 1.5a11 11 0 0 0 5 5L15 13l5 2v4a2 2 0 0 1-2 2A16 16 0 0 1 3 6a2 2 0 0 1 2-2z" /></svg>
              </span>
            </div>
            <strong>{{ payload.phone }}</strong>
          </a>

          <div class="pt-pass-mrz">
            <p>{{ mrzName }}</p>
            <p>{{ mrzCode }}</p>
          </div>

          <div class="pt-pass-visa">
            <span class="pt-pass-visa-mark">GOOD<br />PET</span>
            <div>
              <p>准予进入家庭 · 永久有效</p>
              <small>ISSUED WITH LOVE · 2026</small>
            </div>
          </div>
        </article>
      </div>

      <p class="pt-pass-hint">轻触护照 · 拨打监护人电话</p>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps(['payload', 'nfcId'])

const mrzName = computed(() => {
  const name = String(props.payload?.petName || 'UNNAMED').toUpperCase().replace(/\s/g, '')
  return `P<PETREPUBLIC<${name}${'<'.repeat(Math.max(3, 18 - name.length))}`
})

const mrzCode = computed(() => {
  const id = String(props.nfcId || 'P00000086').padEnd(18, '<')
  const phone = String(props.payload?.phone || '').replace(/\s/g, '')
  return `${id}<<${phone}<<<`
})
</script>

<style scoped>
.pt-pass-root {
  background:
    radial-gradient(circle at 50% 0, rgba(212, 175, 55, 0.13), transparent 38%),
    linear-gradient(160deg, #0b1f38, #132f52 58%, #081728);
  font-family: 'Baskerville', 'Songti SC', Georgia, serif;
}
.pt-pass-content { perspective: 1200px; animation: pt-pass-open 0.9s cubic-bezier(0.2, 0.8, 0.25, 1) both; }
.pt-pass-book { position: relative; display: flex; filter: drop-shadow(0 26px 38px rgba(0, 0, 0, 0.5)); }
.pt-pass-spine { width: 17px; flex: none; display: flex; flex-direction: column; justify-content: space-around; align-items: center; padding: 26px 0; background: linear-gradient(90deg, #071426, #1b3b5f); border-right: 1px solid rgba(212, 175, 55, 0.3); border-radius: 9px 0 0 9px; }
.pt-pass-spine span { width: 7px; height: 2px; background: rgba(212, 175, 55, 0.55); border-radius: 2px; }
.pt-pass-page {
  position: relative;
  flex: 1;
  overflow: hidden;
  padding: 22px 20px 18px;
  color: #14263c;
  background:
    repeating-linear-gradient(0deg, rgba(30, 58, 88, 0.035) 0 1px, transparent 1px 4px),
    radial-gradient(circle at 90% 8%, rgba(212, 175, 55, 0.12), transparent 34%),
    #f3eedf;
  border-radius: 0 8px 8px 0;
  border: 1px solid rgba(212, 175, 55, 0.5);
  border-left: 0;
  transition: transform 0.4s ease;
}
.pt-pass-page:hover { transform: translateY(-3px); }
.pt-pass-page::after {
  content: '';
  position: absolute;
  inset: 8px;
  border: 1px solid rgba(122, 96, 35, 0.35);
  pointer-events: none;
}
.pt-pass-head { position: relative; z-index: 1; display: flex; align-items: flex-start; gap: 10px; }
.pt-pass-emblem { width: 46px; flex: none; color: #96742a; }
.pt-pass-emblem svg { display: block; width: 100%; height: auto; }
.pt-pass-country { margin: 1px 0 2px; font-family: 'Courier New', monospace; font-size: 6.5px; font-weight: 700; letter-spacing: 0.13em; color: #8b6f32; }
.pt-pass-head h1 { margin: 0; font-family: 'Arial Narrow', 'PingFang SC', sans-serif; font-size: 29px; line-height: 0.9; font-weight: 800; letter-spacing: 0.03em; color: #14263c; }
.pt-pass-type { margin-left: auto; align-self: flex-start; display: flex; flex-direction: column; align-items: center; padding: 3px 7px; border: 1.5px solid #96742a; color: #7d5e20; font-family: 'Courier New', monospace; font-size: 17px; line-height: 1; }
.pt-pass-type small { font-size: 7px; margin-top: 2px; }
.pt-pass-rule { display: flex; justify-content: space-between; margin: 14px 0 11px; padding: 5px 0; border-top: 1px solid #a98a45; border-bottom: 1px solid #a98a45; font-family: 'Courier New', monospace; font-size: 6.5px; letter-spacing: 0.1em; color: #765c27; }
.pt-pass-identity { display: grid; grid-template-columns: 82px 1fr; gap: 12px; }
.pt-pass-portrait { position: relative; padding: 4px; background: #fffdf6; border: 1px solid #bda269; }
.pt-pass-portrait-art { position: relative; height: 118px; overflow: hidden; background: linear-gradient(165deg, #cfd9d2, #9eb4b0); }
.pt-pass-ear { position: absolute; top: 18px; width: 20px; height: 31px; background: #7b5638; border-radius: 50% 50% 30% 30%; }
.pt-pass-ear-l { left: 14px; transform: rotate(-20deg); }
.pt-pass-ear-r { right: 14px; transform: rotate(20deg); }
.pt-pass-face { position: absolute; left: 50%; top: 32px; width: 56px; height: 57px; margin-left: -28px; border-radius: 48%; background: #c49363; }
.pt-pass-eye { position: absolute; top: 53px; width: 6px; height: 8px; border-radius: 50%; background: #2c2115; }
.pt-pass-eye-l { left: 27px; }
.pt-pass-eye-r { right: 27px; }
.pt-pass-snout { position: absolute; left: 50%; top: 72px; width: 22px; height: 14px; margin-left: -11px; border-radius: 50%; background: #e2c19e; }
.pt-pass-portrait-label { display: block; padding: 3px 0 1px; text-align: center; font-family: 'Courier New', monospace; font-size: 6px; letter-spacing: 0.2em; color: #8f7435; }
.pt-pass-fields { display: grid; grid-template-columns: 1fr 1fr; gap: 8px 10px; align-content: start; }
.pt-pass-field-wide { grid-column: 1 / -1; }
.pt-pass-field span { display: block; font-family: 'Courier New', monospace; font-size: 6px; font-weight: 700; letter-spacing: 0.08em; color: #8a713c; margin-bottom: 1px; }
.pt-pass-field strong { display: block; font-size: 11px; line-height: 1.35; font-weight: 600; color: #1c3550; word-break: break-word; }
.pt-pass-field-wide strong { font-family: 'Songti SC', 'Georgia', serif; font-size: 19px; letter-spacing: 0.05em; color: #14263c; }
.pt-pass-medical { position: relative; z-index: 1; margin-top: 13px; padding: 8px 10px; background: rgba(176, 74, 58, 0.08); border: 1px solid rgba(176, 74, 58, 0.45); border-left: 3px solid #a84535; }
.pt-pass-medical-head { display: flex; justify-content: space-between; font-family: 'Courier New', monospace; font-size: 6.5px; font-weight: 700; letter-spacing: 0.1em; color: #8f3c2d; }
.pt-pass-cross { font-size: 10px; line-height: 1; }
.pt-pass-medical p { margin: 3px 0 0; font-family: 'PingFang SC', 'Microsoft YaHei', sans-serif; font-size: 11px; line-height: 1.5; color: #7a3327; }
.pt-pass-note { position: relative; z-index: 1; display: grid; grid-template-columns: 46px 1fr; gap: 8px; margin-top: 10px; }
.pt-pass-note span { font-family: 'Courier New', monospace; font-size: 6.5px; font-weight: 700; letter-spacing: 0.06em; color: #8a713c; padding-top: 2px; }
.pt-pass-note p { margin: 0; font-family: 'PingFang SC', 'Microsoft YaHei', sans-serif; font-size: 11px; line-height: 1.6; color: #3c4f60; }
.pt-pass-phone { position: relative; z-index: 1; display: block; margin-top: 13px; padding: 9px 12px 10px; background: #102a47; color: #f5edd8; text-decoration: none; border-left: 4px solid #c6a14e; transition: background 0.22s ease, transform 0.18s ease; }
.pt-pass-phone:hover { background: #17395f; }
.pt-pass-phone:active { transform: scale(0.99); }
.pt-pass-phone-top { display: flex; justify-content: space-between; align-items: center; font-family: 'Courier New', monospace; font-size: 6.5px; letter-spacing: 0.12em; color: #cbb98d; }
.pt-pass-phone-icon svg { width: 12px; height: 12px; color: #d8b45f; }
.pt-pass-phone strong { display: block; margin-top: 1px; font-family: 'Arial', 'PingFang SC', sans-serif; font-size: clamp(24px, 7.5vw, 31px); font-weight: 700; letter-spacing: 0.03em; color: #fff9ea; }
.pt-pass-mrz { position: relative; z-index: 1; margin-top: 10px; padding: 5px 6px; background: rgba(20, 38, 60, 0.06); border-top: 1px dashed #a68b52; border-bottom: 1px dashed #a68b52; }
.pt-pass-mrz p { margin: 0; overflow: hidden; font-family: 'Courier New', monospace; font-size: 6.6px; line-height: 1.5; letter-spacing: 0.12em; color: #4f4a36; white-space: nowrap; }
.pt-pass-visa { position: relative; z-index: 1; display: flex; gap: 10px; align-items: center; margin-top: 11px; transform: rotate(-2deg); opacity: 0.85; }
.pt-pass-visa-mark { width: 46px; height: 46px; flex: none; display: grid; place-items: center; border: 2px solid #a84535; border-radius: 50%; color: #a84535; font-family: 'Arial Black', sans-serif; font-size: 8px; line-height: 1.05; text-align: center; transform: rotate(-8deg); }
.pt-pass-visa p { margin: 0; font-family: 'PingFang SC', 'Microsoft YaHei', sans-serif; font-size: 9px; font-weight: 600; color: #8d3a2c; letter-spacing: 0.06em; }
.pt-pass-visa small { font-family: 'Courier New', monospace; font-size: 6px; letter-spacing: 0.12em; color: #96743a; }
.pt-pass-hint { margin: 17px 0 0; text-align: center; font-family: 'Courier New', monospace; font-size: 8px; letter-spacing: 0.16em; color: rgba(243, 238, 223, 0.55); }
@keyframes pt-pass-open {
  from { opacity: 0; transform: perspective(900px) rotateY(-13deg) translateX(-18px); }
  to { opacity: 1; transform: perspective(900px) rotateY(0) translateX(0); }
}
@media (prefers-reduced-motion: reduce) {
  .pt-pass-content { animation: none; }
  .pt-pass-page, .pt-pass-phone { transition: none; }
}
</style>
