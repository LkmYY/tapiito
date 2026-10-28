<template>
  <div class="nc-ci-root relative min-h-screen overflow-hidden px-5 py-8 text-[#111110]">
    <div class="nc-ci-paper"></div>
    <span class="nc-ci-rail" aria-hidden="true">NIGHTCLUB DIRECTORY · 2026</span>

    <div class="relative z-10 mx-auto flex min-h-[calc(100vh-4rem)] w-full max-w-md flex-col">
      <header class="nc-ci-head">
        <span class="nc-ci-brand">INDEX</span>
        <span class="nc-ci-head-mid">NC—24 / SHANGHAI</span>
        <span class="nc-ci-head-num">{{ nfcId || '0000' }}</span>
      </header>

      <main class="nc-ci-main">
        <p class="nc-ci-kicker">A DIRECTORY OF PEOPLE<br />WHO OWN THE NIGHT</p>

        <div class="nc-ci-name-row">
          <h1 class="nc-ci-name">{{ payload?.name }}</h1>
          <span class="nc-ci-slash">/</span>
          <span class="nc-ci-alias">{{ payload?.alias || 'NO ALIAS' }}</span>
        </div>
        <p class="nc-ci-role">{{ payload?.role }}</p>

        <div class="nc-ci-intro">
          <span class="nc-ci-intro-no">00</span>
          <p>Resident at <b>{{ payload?.venue }}</b>, {{ payload?.city || '—' }}. Listed here for those who know the address.</p>
        </div>

        <section class="nc-ci-list">
          <div class="nc-ci-row">
            <span class="nc-ci-no">01</span>
            <span class="nc-ci-key">VENUE</span>
            <span class="nc-ci-value">{{ payload?.venue }}</span>
            <span class="nc-ci-meta">{{ payload?.city || '—' }}</span>
          </div>
          <div class="nc-ci-row">
            <span class="nc-ci-no">02</span>
            <span class="nc-ci-key">WECHAT</span>
            <span class="nc-ci-value">{{ payload?.wechat }}</span>
            <span class="nc-ci-meta">ADD</span>
          </div>
          <a v-if="payload?.phone" :href="'tel:' + payload.phone" class="nc-ci-row nc-ci-row-link">
            <span class="nc-ci-no">03</span>
            <span class="nc-ci-key">PHONE</span>
            <span class="nc-ci-value">{{ payload.phone }}</span>
            <span class="nc-ci-meta">CALL ↗</span>
          </a>
          <div v-else class="nc-ci-row">
            <span class="nc-ci-no">03</span>
            <span class="nc-ci-key">PHONE</span>
            <span class="nc-ci-value nc-ci-empty"></span>
            <span class="nc-ci-meta">—</span>
          </div>
          <div class="nc-ci-row">
            <span class="nc-ci-no">04</span>
            <span class="nc-ci-key">HOURS</span>
            <span class="nc-ci-value">23:00 — LATE</span>
            <span class="nc-ci-meta">DAILY</span>
          </div>
        </section>

        <div class="nc-ci-footnote">
          <span class="nc-ci-mark">✳</span>
          <p>No guest list. No photographs. The index is updated after every full moon.</p>
        </div>
      </main>

      <footer class="nc-ci-foot">
        <span>INDEX / 30</span>
        <span>PRINTED MATTER</span>
        <span>END</span>
      </footer>
    </div>
  </div>
</template>

<script setup>
defineProps(['payload', 'nfcId'])
</script>

<style scoped>
.nc-ci-root {
  background: #efede6;
  font-family: 'Helvetica Neue', 'PingFang SC', 'Microsoft YaHei', sans-serif;
}
.nc-ci-paper {
  position: absolute;
  inset: 0;
  background:
    radial-gradient(ellipse at 84% 6%, rgba(255, 255, 255, 0.8), transparent 40%),
    repeating-linear-gradient(90deg, rgba(17, 17, 16, 0.025) 0 1px, transparent 1px 4px);
  pointer-events: none;
}
.nc-ci-rail {
  position: absolute;
  top: 50%;
  right: 9px;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 5.5px;
  letter-spacing: 0.3em;
  color: rgba(17, 17, 16, 0.28);
  writing-mode: vertical-rl;
  transform: translateY(-50%);
}
.nc-ci-head {
  display: flex;
  align-items: flex-end;
  justify-content: space-between;
  gap: 8px;
  padding-bottom: 8px;
  border-bottom: 2px solid #111110;
  animation: nc-ci-in 0.8s both;
}
.nc-ci-brand {
  font-size: 26px;
  font-weight: 900;
  line-height: 0.8;
  letter-spacing: -0.05em;
}
.nc-ci-head-mid {
  flex: 1;
  padding-bottom: 1px;
  text-align: center;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 5.5px;
  letter-spacing: 0.14em;
  color: #6e6b63;
}
.nc-ci-head-num {
  padding-bottom: 1px;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6px;
  letter-spacing: 0.1em;
  color: #6e6b63;
}
.nc-ci-main {
  flex: 1;
  padding-top: 30px;
}
.nc-ci-kicker {
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6.5px;
  line-height: 1.7;
  letter-spacing: 0.18em;
  color: #6e6b63;
  animation: nc-ci-in 0.8s 0.06s both;
}
.nc-ci-name-row {
  display: flex;
  flex-wrap: wrap;
  align-items: baseline;
  gap: 8px;
  margin-top: 22px;
  animation: nc-ci-in 0.95s 0.1s cubic-bezier(0.22, 1, 0.36, 1) both;
}
.nc-ci-name {
  font-size: clamp(52px, 16.5vw, 72px);
  font-weight: 900;
  line-height: 0.84;
  letter-spacing: -0.065em;
  text-transform: uppercase;
  color: #111110;
  overflow-wrap: anywhere;
}
.nc-ci-slash {
  font-size: 30px;
  font-weight: 300;
  color: #e63322;
}
.nc-ci-alias {
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 8px;
  letter-spacing: 0.3em;
  color: #6e6b63;
  text-transform: uppercase;
}
.nc-ci-role {
  margin-top: 8px;
  font-size: 10px;
  font-weight: 700;
  letter-spacing: 0.2em;
  color: #e63322;
  text-transform: uppercase;
  animation: nc-ci-in 0.8s 0.2s both;
}
.nc-ci-intro {
  display: flex;
  gap: 12px;
  margin-top: 30px;
  padding-top: 12px;
  border-top: 1px solid rgba(17, 17, 16, 0.3);
  animation: nc-ci-in 0.9s 0.26s both;
}
.nc-ci-intro-no {
  flex: none;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 7px;
  color: #e63322;
}
.nc-ci-intro p {
  max-width: 300px;
  font-size: 10.5px;
  line-height: 1.6;
  letter-spacing: 0.01em;
  color: #44423d;
}
.nc-ci-intro b {
  font-weight: 700;
  color: #111110;
}
.nc-ci-list {
  margin-top: 28px;
  border-top: 2px solid #111110;
  animation: nc-ci-in 1s 0.32s both;
}
.nc-ci-row {
  display: flex;
  align-items: baseline;
  gap: 9px;
  padding: 11px 0;
  border-bottom: 1px solid rgba(17, 17, 16, 0.2);
  text-decoration: none;
  color: inherit;
}
.nc-ci-no {
  width: 18px;
  flex: none;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6px;
  color: #a4a097;
}
.nc-ci-key {
  width: 50px;
  flex: none;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 6px;
  letter-spacing: 0.12em;
  color: #6e6b63;
}
.nc-ci-value {
  min-width: 0;
  flex: 1;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  font-size: 11.5px;
  font-weight: 600;
  letter-spacing: 0.03em;
  color: #111110;
}
.nc-ci-meta {
  flex: none;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 5.5px;
  letter-spacing: 0.08em;
  color: #a4a097;
}
.nc-ci-row-link {
  margin: 0 -7px;
  padding: 11px 7px;
  transition: background 0.25s ease, padding 0.25s ease, color 0.25s ease;
}
.nc-ci-row-link:hover {
  background: #e63322;
  padding-left: 14px;
}
.nc-ci-row-link:hover .nc-ci-no,
.nc-ci-row-link:hover .nc-ci-key,
.nc-ci-row-link:hover .nc-ci-meta,
.nc-ci-row-link:hover .nc-ci-value {
  color: #f7f3ea;
}
.nc-ci-row-link:active {
  background: #c42718;
}
.nc-ci-empty {
  color: #b3afa6;
}
.nc-ci-footnote {
  display: flex;
  gap: 10px;
  margin-top: 28px;
  animation: nc-ci-in 1s 0.46s both;
}
.nc-ci-mark {
  font-size: 15px;
  line-height: 1;
  color: #e63322;
}
.nc-ci-footnote p {
  font-family: Georgia, 'Songti SC', serif;
  font-size: 9.5px;
  font-style: italic;
  line-height: 1.6;
  color: #6e6b63;
}
.nc-ci-foot {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding-top: 12px;
  border-top: 2px solid #111110;
  font-family: 'SFMono-Regular', Consolas, monospace;
  font-size: 5.5px;
  letter-spacing: 0.16em;
  color: #6e6b63;
  animation: nc-ci-in 1s 0.56s both;
}
@keyframes nc-ci-in {
  from { opacity: 0; transform: translateY(14px); }
  to { opacity: 1; transform: none; }
}
@media (prefers-reduced-motion: reduce) {
  .nc-ci-root *,
  .nc-ci-root *::before,
  .nc-ci-root *::after {
    animation: none !important;
    transition: none !important;
  }
}
</style>
