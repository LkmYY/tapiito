<template>
  <div class="relative flex min-h-full flex-1 flex-col overflow-hidden bg-[#fbfaf7] font-sans text-[#14181d]">
    <!-- 背景光斑 -->
    <div class="pointer-events-none absolute inset-0">
      <div class="absolute -left-16 -top-24 h-[260px] w-[260px] rounded-full bg-[#ff6b35]/15 blur-[80px]" />
      <div class="absolute -right-16 top-[18%] h-[240px] w-[240px] rounded-full bg-[#ffd166]/30 blur-[80px]" />
      <div class="absolute bottom-[-120px] left-1/4 h-[260px] w-[260px] rounded-full bg-[#7bdff2]/25 blur-[90px]" />
    </div>

    <div class="relative z-10 flex flex-1 flex-col px-6 pb-8 pt-10">
      <!-- 品牌 -->
      <div class="flex items-center gap-2.5">
        <span class="grid h-9 w-9 place-items-center rounded-[11px]  bg-[#15151B]">
          <img src="/tapiito-icon-white.png" alt="Tapiito" class="h-[18px] w-[18px]" draggable="false" />
        </span>
        <span class="text-[18px] font-semibold tracking-tight">Tapiito</span>
      </div>

      <!-- 状态说明 -->
      <div class="mt-12">
        <span class="inline-flex items-center gap-1.5 rounded-full border border-[#14181d]/10 bg-white/70 px-3 py-1.5 text-[11.5px] font-medium text-[#14181d]/60 backdrop-blur">
          <span class="h-1.5 w-1.5 rounded-full bg-[#ff6b35]" />
          {{ meta.badge }}
        </span>

        <h1 class="mt-5 text-[30px] font-semibold leading-[1.18] tracking-[-0.03em]">
          {{ meta.title }}
        </h1>
        <p class="mt-3 text-[14px] leading-relaxed text-[#14181d]/55">
          {{ meta.desc }}
        </p>
      </div>

      <!-- 品牌主张 -->
      <div class="mt-9 rounded-2xl border border-[#14181d]/[0.07] bg-white/80 p-5 backdrop-blur">
        <p class="text-[11px] font-semibold uppercase tracking-[0.18em] text-[#ff6b35]">What we do</p>
        <p class="mt-2.5 text-[15px] font-medium leading-relaxed">
          一张薄薄的标签，碰一下，内容自己跳出来。
        </p>
        <p class="mt-2 text-[13px] leading-relaxed text-[#14181d]/50">
          名片、菜单、门店、宠物牌、紧急联系人 —— 贴上去就能用；购买即完成首次写入，含 1 次免费内容调整。
        </p>

        <ul class="mt-4 grid gap-2">
          <li v-for="item in highlights" :key="item.text" class="flex items-start gap-2 text-[12.5px] text-[#14181d]/60">
            <Check class="mt-[2px] h-3.5 w-3.5 flex-shrink-0 text-[#16a34a]" :stroke-width="2.6" />
            <span class="min-w-0">
              {{ item.text }}
              <span v-if="item.note" class="mt-0.5 block text-[11.5px] leading-relaxed text-[#14181d]/40">{{ item.note }}</span>
            </span>
          </li>
        </ul>
      </div>

      <!-- 行动 -->
      <div class="mt-7 flex flex-col gap-2.5">
        <router-link to="/showroom" class="cf-cta">
          看看都有哪些玩法
          <ArrowRight class="h-4 w-4" />
        </router-link>
        <router-link to="/" class="cf-ghost">了解 Tapiito</router-link>
      </div>

      <p class="mt-auto pt-8 text-center text-[11px] text-[#14181d]/30">
        Tap it alive. · Tapiito
      </p>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { Nfc, ArrowRight, Check } from 'lucide-vue-next'

const props = defineProps({
  /** disabled = 标签被停用；notfound = 标签不存在；error = 加载失败 */
  reason: { type: String, default: 'disabled' },
})

const CONTENT = {
  disabled: {
    badge: '这枚标签暂停中',
    title: '标签主人暂时关掉了它',
    desc: '内容被停用了，不是出错。滑到底下看看 Tapiito 能做什么。',
  },
  notfound: {
    badge: '这枚标签还没激活',
    title: '没找到对应的内容',
    desc: '可能还没绑定内容，或者链接不太对。先看看我们做什么吧。',
  },
  error: {
    badge: '页面暂时打不开',
    title: '这会儿有点连不上',
    desc: '网络或服务出了点小状况，稍后再试一次就好。',
  },
}

const meta = computed(() => CONTENT[props.reason] || CONTENT.disabled)

const highlights = [
  { text: '10 种标签类型 × 30 种风格，随便挑' },
  { text: '手机一碰就打开，不用装 App' },
  { text: '购买即完成首次写入', note: '含 1 次免费内容调整，后续调整另收制作费' },
]
</script>

<style scoped>
.cf-cta {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  gap: 7px;
  height: 46px;
  border-radius: 13px;
  background: #14181d;
  color: #fff;
  font-size: 14px;
  font-weight: 600;
  transition: all 0.22s ease;
}
.cf-cta:hover {
  background: #ff6b35;
  transform: translateY(-1px);
  box-shadow: 0 14px 28px -16px rgba(255, 107, 53, 0.9);
}

.cf-ghost {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  height: 44px;
  border-radius: 13px;
  border: 1px solid rgba(20, 24, 29, 0.12);
  background: rgba(255, 255, 255, 0.7);
  color: rgba(20, 24, 29, 0.7);
  font-size: 13.5px;
  font-weight: 500;
  transition: all 0.2s ease;
}
.cf-ghost:hover {
  border-color: rgba(20, 24, 29, 0.3);
  color: #14181d;
}
</style>
