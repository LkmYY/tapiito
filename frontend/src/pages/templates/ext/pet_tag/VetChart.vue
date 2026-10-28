<template>
  <div class="pt-vet-root min-h-screen px-5 py-8 text-slate-800">
    <div class="pt-vet-content mx-auto w-full max-w-md">
      <div class="pt-vet-brandrow">
        <div class="pt-vet-logo">
          <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8" aria-hidden="true">
            <path d="M4 10v5a2 2 0 0 0 2 2h1v3h3v-3h4v3h3v-3h1a2 2 0 0 0 2-2v-5" />
            <path d="M6 10V8a6 6 0 0 1 12 0v2" />
            <path d="M12 7v3M10 8.5h4" />
          </svg>
        </div>
        <div>
          <p class="pt-vet-hospital">青禾动物医院</p>
          <p class="pt-vet-dept">COMPANION ANIMAL CLINIC · 病历档案</p>
        </div>
        <span class="pt-vet-file">NO. {{ nfcId || 'V-2306' }}</span>
      </div>

      <article class="pt-vet-clipboard">
        <div class="pt-vet-clip" aria-hidden="true"><span></span></div>

        <div class="pt-vet-paper">
          <div class="pt-vet-paperhead">
            <div>
              <p class="pt-vet-overline">PATIENT MEDICAL CARD</p>
              <h1>{{ payload.petName || '未命名' }}</h1>
              <p class="pt-vet-sub">{{ payload.breed }} · {{ payload.age || '—' }}</p>
            </div>
            <div class="pt-vet-chip">
              <span>CHIP</span>
              <strong>{{ nfcId || '985 141 002' }}</strong>
            </div>
          </div>

          <div class="pt-vet-grid">
            <div class="pt-vet-cell">
              <span class="pt-vet-key">监护人 / OWNER</span>
              <p>{{ payload.ownerName }}</p>
            </div>
            <div class="pt-vet-cell">
              <span class="pt-vet-key">住址 / ADDRESS</span>
              <p>{{ payload.address }}</p>
            </div>
          </div>

          <div class="pt-vet-section-title">
            <span>MEDICAL NOTES</span>
            <span>过敏 / 用药</span>
          </div>
          <div class="pt-vet-medical">
            <span class="pt-vet-cross" aria-hidden="true">＋</span>
            <p>{{ payload.medical || '无已知过敏与长期用药记录。' }}</p>
          </div>

          <div class="pt-vet-section-title">
            <span>BEHAVIOUR</span>
            <span>行为提示</span>
          </div>
          <p class="pt-vet-note">{{ payload.note || '性格温顺，遇生人请先轻声安抚。' }}</p>

          <div class="pt-vet-emergency">
            <div class="pt-vet-emergency-head">
              <span class="pt-vet-pulse" aria-hidden="true"></span>
              <span>EMERGENCY CONTACT · 紧急联络</span>
            </div>
            <a :href="'tel:' + payload.phone">
              <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" aria-hidden="true">
                <path d="M22 16.9v3a2 2 0 0 1-2.2 2 19.8 19.8 0 0 1-8.6-3.1 19.5 19.5 0 0 1-6-6A19.8 19.8 0 0 1 2.1 4.2 2 2 0 0 1 4.1 2h3a2 2 0 0 1 2 1.7c.1 1 .4 2 .7 2.9a2 2 0 0 1-.5 2.1L8.1 9.9a16 16 0 0 0 6 6l1.2-1.2a2 2 0 0 1 2.1-.5c.9.3 1.9.6 2.9.7a2 2 0 0 1 1.7 2z" />
              </svg>
              {{ payload.phone }}
            </a>
            <p class="pt-vet-emergency-sub">请优先联系监护人 · 医院 24h 值班</p>
          </div>

          <div class="pt-vet-signrow">
            <div>
              <span class="pt-vet-sign-label">主治医师</span>
              <p class="pt-vet-sign">Dr. Lin</p>
            </div>
            <div class="pt-vet-stamp">已建档<br /><small>ARCHIVED</small></div>
          </div>
        </div>
      </article>

      <p class="pt-vet-foot">本卡片由 NFC 铭牌读取 · 信息仅用于宠物救助</p>
    </div>
  </div>
</template>

<script setup>
defineProps(['payload', 'nfcId'])
</script>

<style scoped>
.pt-vet-root {
  background:
    linear-gradient(rgba(15, 118, 110, 0.045) 1px, transparent 1px),
    linear-gradient(90deg, rgba(15, 118, 110, 0.045) 1px, transparent 1px),
    #edf5f1;
  background-size: 22px 22px;
  font-family: 'Avenir Next', 'PingFang SC', 'Microsoft YaHei', sans-serif;
}
.pt-vet-content { animation: pt-vet-rise 0.7s cubic-bezier(0.25, 0.85, 0.3, 1) both; }
.pt-vet-brandrow { display: flex; align-items: center; gap: 10px; padding: 0 4px 14px; }
.pt-vet-logo { width: 38px; height: 38px; flex: none; display: grid; place-items: center; color: #0f766e; border: 1.5px solid #0f766e; border-radius: 7px; background: #fff; }
.pt-vet-logo svg { width: 23px; height: 23px; }
.pt-vet-hospital { margin: 0; font-size: 14px; font-weight: 700; letter-spacing: 0.04em; color: #134e4a; }
.pt-vet-dept { margin: 2px 0 0; font-family: 'Courier New', monospace; font-size: 7px; letter-spacing: 0.12em; color: #5f7f7a; }
.pt-vet-file { margin-left: auto; font-family: 'Courier New', monospace; font-size: 8px; letter-spacing: 0.08em; color: #0f766e; border-bottom: 1px solid #0f766e; padding-bottom: 2px; }
.pt-vet-clipboard { position: relative; padding-top: 16px; }
.pt-vet-clip {
  position: absolute;
  z-index: 3;
  top: 0;
  left: 50%;
  width: 98px;
  height: 32px;
  transform: translateX(-50%);
  border-radius: 7px;
  background: linear-gradient(180deg, #d8e4e1, #9fb5b0);
  box-shadow: 0 4px 7px rgba(30, 64, 60, 0.22), inset 0 1px 0 #f4faf8;
  animation: pt-vet-clip 0.5s 0.3s cubic-bezier(0.3, 0.9, 0.4, 1.3) both;
}
.pt-vet-clip span { position: absolute; inset: 9px 22px; border-radius: 3px; background: #f8fbfa; border: 1px solid #a6b8b4; }
.pt-vet-clip::before, .pt-vet-clip::after { content: ''; position: absolute; top: 7px; width: 5px; height: 5px; border-radius: 50%; background: #6c8580; }
.pt-vet-clip::before { left: 18px; }
.pt-vet-clip::after { right: 18px; }
.pt-vet-paper {
  position: relative;
  padding: 32px 22px 22px;
  background: #fffefb;
  border: 1px solid #d7e2de;
  border-radius: 4px;
  box-shadow: 0 18px 32px rgba(32, 67, 61, 0.14);
  background-image: repeating-linear-gradient(0deg, transparent 0 31px, rgba(15, 118, 110, 0.055) 31px 32px);
}
.pt-vet-paper::before { content: ''; position: absolute; left: 34px; top: 0; bottom: 0; width: 1px; background: rgba(225, 105, 105, 0.18); }
.pt-vet-paperhead { display: flex; justify-content: space-between; gap: 10px; padding-bottom: 14px; border-bottom: 2px solid #134e4a; }
.pt-vet-overline { margin: 0 0 5px; font-family: 'Courier New', monospace; font-size: 7px; letter-spacing: 0.2em; color: #6d8d87; }
.pt-vet-paperhead h1 { margin: 0; font-size: clamp(34px, 10vw, 44px); line-height: 1; letter-spacing: -0.02em; color: #123f3b; }
.pt-vet-sub { margin: 6px 0 0; font-size: 11px; font-weight: 600; color: #4d726c; letter-spacing: 0.08em; }
.pt-vet-chip { flex: none; align-self: flex-start; display: flex; flex-direction: column; gap: 2px; padding: 6px 7px; background: #ecf7f3; border: 1px solid #b7d8cf; }
.pt-vet-chip span { font-family: 'Courier New', monospace; font-size: 6px; letter-spacing: 0.18em; color: #62877f; }
.pt-vet-chip strong { font-family: 'Courier New', monospace; font-size: 8px; color: #1d5a53; }
.pt-vet-grid { display: grid; grid-template-columns: 0.85fr 1.15fr; gap: 12px; padding: 14px 0 12px; border-bottom: 1px dashed #c2d5cf; }
.pt-vet-cell p { margin: 3px 0 0; font-size: 12px; font-weight: 600; line-height: 1.5; color: #264a45; }
.pt-vet-key { font-family: 'Courier New', monospace; font-size: 7px; font-weight: 700; letter-spacing: 0.09em; color: #77948f; }
.pt-vet-section-title { display: flex; justify-content: space-between; margin-top: 15px; font-family: 'Courier New', monospace; font-size: 8px; font-weight: 700; letter-spacing: 0.13em; color: #0f766e; }
.pt-vet-section-title span:last-child { font-family: 'PingFang SC', 'Microsoft YaHei', sans-serif; letter-spacing: 0.08em; color: #6f8e89; }
.pt-vet-medical { display: flex; gap: 8px; margin-top: 7px; padding: 9px 10px; background: #fff2ee; border-left: 3px solid #d9553f; transition: background 0.2s ease; }
.pt-vet-medical:hover { background: #ffe9e3; }
.pt-vet-cross { flex: none; color: #c24430; font-weight: 700; line-height: 1.4; }
.pt-vet-medical p { margin: 0; font-size: 12px; line-height: 1.55; font-weight: 600; color: #8d3628; }
.pt-vet-note { margin: 6px 0 0; font-size: 12px; line-height: 1.7; color: #36534e; }
.pt-vet-emergency { margin-top: 18px; padding: 12px 13px 11px; background: #123f3b; color: #e8f5f1; border-radius: 3px; }
.pt-vet-emergency-head { display: flex; align-items: center; gap: 6px; font-family: 'Courier New', monospace; font-size: 7px; letter-spacing: 0.14em; color: #9dd3c8; }
.pt-vet-pulse { width: 6px; height: 6px; border-radius: 50%; background: #6ee7b7; box-shadow: 0 0 0 0 rgba(110, 231, 183, 0.5); animation: pt-vet-pulse 2s infinite; }
.pt-vet-emergency a { display: flex; align-items: center; gap: 8px; margin: 3px 0 1px; color: #fff; font-size: clamp(27px, 8vw, 34px); font-weight: 700; letter-spacing: 0.02em; text-decoration: none; transition: color 0.2s ease; }
.pt-vet-emergency a svg { width: 21px; height: 21px; color: #6ee7b7; flex: none; }
.pt-vet-emergency a:hover { color: #6ee7b7; }
.pt-vet-emergency-sub { margin: 0; font-size: 8px; letter-spacing: 0.08em; color: #9cbab4; }
.pt-vet-signrow { display: flex; justify-content: space-between; align-items: flex-end; margin-top: 16px; }
.pt-vet-sign-label { font-size: 8px; letter-spacing: 0.1em; color: #7d9994; }
.pt-vet-sign { margin: 1px 0 0; font-family: 'Segoe Script', 'Bradley Hand', cursive; font-size: 18px; color: #1d5a53; border-bottom: 1px solid #9ebdb6; padding: 0 6px 2px 0; }
.pt-vet-stamp { padding: 5px 8px; color: #b04b3a; border: 2px double #b04b3a; font-size: 10px; font-weight: 700; line-height: 1.2; text-align: center; transform: rotate(-6deg); opacity: 0.75; }
.pt-vet-stamp small { font-family: 'Courier New', monospace; font-size: 6px; letter-spacing: 0.12em; }
.pt-vet-foot { margin: 15px 0 0; text-align: center; font-size: 8px; letter-spacing: 0.1em; color: #78938e; }
@keyframes pt-vet-rise {
  from { opacity: 0; transform: translateY(30px); }
  to { opacity: 1; transform: translateY(0); }
}
@keyframes pt-vet-clip {
  from { opacity: 0; transform: translate(-50%, -20px); }
  to { opacity: 1; transform: translate(-50%, 0); }
}
@keyframes pt-vet-pulse {
  0% { box-shadow: 0 0 0 0 rgba(110, 231, 183, 0.45); }
  70% { box-shadow: 0 0 0 7px rgba(110, 231, 183, 0); }
  100% { box-shadow: 0 0 0 0 rgba(110, 231, 183, 0); }
}
@media (prefers-reduced-motion: reduce) {
  .pt-vet-content, .pt-vet-clip, .pt-vet-pulse { animation: none; }
}
</style>
