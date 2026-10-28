<template>
  <div class="legal-page">
    <header class="legal-nav-wrap">
      <nav class="legal-shell legal-nav" aria-label="主导航">
        <router-link to="/" class="legal-brand">
          <span class="legal-brand-mark">
            <img src="/tapiito-icon-white.png" alt="Tapiito" class="h-5 w-5" draggable="false" />
          </span>
          <span>Tapiito</span>
        </router-link>

        <div class="legal-nav-links">
          <router-link to="/showroom" class="legal-nav-link">展厅</router-link>
          <router-link to="/about" class="legal-nav-link">关于我们</router-link>
          <router-link to="/consult" class="legal-nav-cta">
            预约咨询
            <ArrowUpRight class="h-3.5 w-3.5" />
          </router-link>
        </div>
      </nav>
    </header>

    <main>
      <section class="legal-hero">
        <div class="legal-hero-glow legal-hero-glow-a" aria-hidden="true"></div>
        <div class="legal-hero-glow legal-hero-glow-b" aria-hidden="true"></div>
        <div class="legal-hero-grid" aria-hidden="true"></div>

        <div class="legal-shell legal-hero-inner">
          <div>
            <p class="legal-kicker">{{ page.eyebrow }}</p>
            <h1 class="legal-title">
              {{ page.title }}<span class="legal-title-dot">.</span>
            </h1>
            <p class="legal-lead">{{ page.lead }}</p>

            <div class="legal-meta">
              <span class="legal-meta-pill">
                <component :is="page.icon" class="h-3.5 w-3.5" />
                {{ page.meta }}
              </span>
              <span>更新于 {{ page.updated }}</span>
            </div>
          </div>

          <div class="legal-hero-seal" aria-hidden="true">
            <div class="legal-seal-ring"></div>
            <div class="legal-seal-ring legal-seal-ring-inner"></div>
            <div class="legal-seal-core">
              <component :is="page.icon" class="h-7 w-7" />
              <span>TAPIITO<br />LEGAL</span>
            </div>
          </div>
        </div>
      </section>

      <div class="legal-shell legal-layout">
        <aside class="legal-aside">
          <div class="legal-aside-inner">
            <p class="legal-aside-title">本页目录</p>
            <a
              v-for="(section, index) in page.sections"
              :key="section.title"
              :href="'#legal-' + (index + 1)"
              class="legal-toc-link"
            >
              <span>{{ pad(index + 1) }}</span>
              {{ section.title }}
            </a>

            <div class="legal-aside-note">
              <p>需要帮助？</p>
              <span>可通过平台内订单页或客服入口联系我们。</span>
            </div>
          </div>
        </aside>

        <article class="legal-article">
          <p class="legal-intro">{{ page.intro }}</p>

          <section
            v-for="(section, index) in page.sections"
            :id="'legal-' + (index + 1)"
            :key="section.title"
            class="legal-section"
          >
            <div class="legal-section-head">
              <span class="legal-section-num">{{ pad(index + 1) }}</span>
              <h2>{{ section.title }}</h2>
            </div>

            <p v-for="paragraph in section.paragraphs" :key="paragraph" class="legal-p">
              {{ paragraph }}
            </p>

            <ul v-if="section.list" class="legal-list">
              <li v-for="item in section.list" :key="item">{{ item }}</li>
            </ul>
          </section>

          <div class="legal-contact">
            <div>
              <p class="legal-contact-kicker">CONTACT · 联系</p>
              <h2>还有疑问？</h2>
              <p>
                你可以在平台内「预约咨询」订单页或客服入口联系我们。提交隐私相关请求时，
                请注明「隐私请求」，以便我们更快处理。
              </p>
            </div>
            <router-link to="/consult" class="legal-contact-link">
              前往订单页
              <ArrowUpRight class="h-4 w-4" />
            </router-link>
          </div>

          <div class="legal-related">
            <p>继续了解</p>
            <router-link :to="page.other.to" class="legal-related-card">
              <span>
                <strong>{{ page.other.label }}</strong>
                <small>{{ page.other.desc }}</small>
              </span>
              <ArrowUpRight class="h-4 w-4" />
            </router-link>
            <router-link to="/about" class="legal-related-card">
              <span>
                <strong>关于我们</strong>
                <small>了解 Tapiito 在做什么，以及一张标签能走多远。</small>
              </span>
              <ArrowUpRight class="h-4 w-4" />
            </router-link>
            <router-link to="/" class="legal-back">
              <ArrowLeft class="h-3.5 w-3.5" />
              返回首页
            </router-link>
          </div>
        </article>
      </div>
    </main>

    <footer class="legal-footer">
      <div class="legal-shell legal-footer-inner">
        <div class="legal-footer-brand">
          <span class="legal-brand-mark legal-brand-mark-sm">
            <img src="/tapiito-icon-white.png" alt="Tapiito" class="h-[15px] w-[15px]" draggable="false" />
          </span>
          <span>Tapiito</span>
          <span class="legal-footer-tagline">Tap it alive.</span>
        </div>

        <nav class="legal-footer-links" aria-label="页脚导航">
          <router-link to="/showroom">展厅</router-link>
          <router-link to="/about">关于我们</router-link>
          <router-link to="/feedback">我有意见</router-link>
          <router-link to="/privacy">隐私政策</router-link>
          <router-link to="/disclaimer">免责声明</router-link>
        </nav>

        <p class="legal-footer-copy">© {{ new Date().getFullYear() }} Tapiito</p>
      </div>
    </footer>
  </div>
</template>

<script setup>
import { computed, onBeforeUnmount, watch } from 'vue'
import { useRoute } from 'vue-router'
import { ArrowLeft, ArrowUpRight, Nfc, Scale, ShieldCheck } from 'lucide-vue-next'

const route = useRoute()

const PAGES = {
  '/privacy': {
    eyebrow: 'LEGAL · PRIVACY',
    title: '隐私政策',
    lead: '我们如何收集、使用、保存与保护你的信息，以及你可以如何行使自己的权利。',
    meta: '个人信息保护说明',
    updated: '2026 年 10 月 21 日',
    icon: ShieldCheck,
    intro:
      'Tapiito 提供 NFC 标签与页面内容服务。我们重视你的个人信息，并尽量只收集提供服务所必需的信息。本政策适用于 Tapiito 网站、NFC 标签页面、订单页面及相关管理功能。',
    sections: [
      {
        title: '适用范围与运营者',
        paragraphs: [
          '本政策适用于 Tapiito 平台及其提供的 NFC 标签、展示页面、订单与管理服务。通过 NFC 标签打开的第三方网站或第三方服务，适用其自己的隐私政策。',
          '本政策中的「Tapiito」指 Tapiito 平台运营团队。若你对本政策有疑问，可通过平台内订单页或客服入口联系我们。'
        ]
      },
      {
        title: '我们收集的信息',
        paragraphs: [
          '我们会在你登录、创建标签、下单或使用平台功能时收集必要信息。不同功能所需的信息可能不同。'
        ],
        list: [
          '账号信息：用户名、邮箱或手机号、密码哈希、账号角色与创建时间。密码使用 bcrypt 哈希保存，我们不会保存可直接读取的明文密码。',
          '标签内容：你向我们提供的名称、简介、电话、地址、社交账号、图片、菜单、作品、宠物信息、儿童信息、紧急联系人等内容。相关页面由平台根据你提供的信息创建和更新。',
          '订单信息：收件人姓名、联系电话、收货地址、邮编、订单备注、订单状态与支付结果。支付信息由支付服务商处理，我们不保存完整银行卡信息。',
          '使用信息：访问时间、IP 地址、浏览器与设备类型、访问页面、错误日志等，用于保障服务安全和排查问题。',
          '本地存储：登录令牌、主题偏好及少量必要的界面设置，用于保持登录状态和改善使用体验。'
        ]
      },
      {
        title: '我们如何使用信息',
        paragraphs: [
          '我们仅在提供、维护和改进服务所需的范围内使用信息，主要包括以下用途。'
        ],
        list: [
          '根据你提供的信息创建、保存、展示和更新你的 NFC 标签页面。',
          '完成订单、安排发货、处理售后与客服请求。',
          '验证账号身份、保护账号安全、预防欺诈和滥用。',
          '分析服务运行情况，修复故障并改进产品体验。',
          '遵守法律法规、响应司法机关或监管部门的合法要求。',
          '在取得同意的情况下，向你发送与服务相关的通知或更新。'
        ]
      },
      {
        title: '信息共享与委托处理',
        paragraphs: [
          '我们不会出售你的个人信息。为完成服务，我们可能委托云托管、短信、支付、物流或数据分析等服务商处理必要信息，并要求其在约定范围内使用。',
          '在涉及合并、收购、资产转让或类似交易时，相关信息可能作为交易的一部分被转移；我们会要求承接方继续受本政策约束。',
          '我们也可能在法律法规要求、司法或行政机关依法提出要求，或为保护用户、公众及平台合法权益所必需时披露信息。'
        ]
      },
      {
        title: 'Cookie 与本地存储',
        paragraphs: [
          '我们使用必要的本地存储来保持登录状态、记住主题偏好和改善页面体验。禁用这些存储可能导致部分功能无法正常使用。',
          '我们目前不使用第三方广告 Cookie。若未来引入新的统计或广告服务，我们会更新本政策并说明相关选择。'
        ]
      },
      {
        title: '信息保存与安全',
        paragraphs: [
          '我们会在实现服务目的所需的期限内保存信息，并在账号注销或服务终止后按法律要求删除或匿名化处理。法律法规另有规定的，从其规定。',
          '我们采取访问控制、密码哈希、传输加密（在适用场景下）和日志审计等合理措施保护信息。但互联网环境不存在绝对安全，请妥善保管账号和密码。'
        ]
      },
      {
        title: '你的权利',
        paragraphs: [
          '在适用法律允许的范围内，你有权访问、复制、更正、补充或删除你的个人信息，也可以撤回同意、注销账号或要求限制处理。',
          '为保护账号安全，我们可能需要先验证你的身份。收到请求后，我们会在合理期限内处理；法律法规另有规定的除外。'
        ]
      },
      {
        title: '未成年人信息',
        paragraphs: [
          '未满 14 周岁的未成年人不应独立使用本服务。如确需使用，应由监护人陪同、同意，并由监护人向我们提供必要信息。',
          '儿童信息类标签应由监护人提出需求、提供并确认信息，由平台协助创建和管理。未成年人本人不应自行提交或修改相关内容。',
          '若我们发现未经监护人同意收集了未成年人信息，会尽快删除或采取其他适当措施。'
        ]
      },
      {
        title: '第三方链接与服务',
        paragraphs: [
          '你的标签页面可能包含指向第三方网站、社交平台、地图或支付服务的链接。这些第三方服务由相应运营者独立提供，我们无法控制其信息处理行为。',
          '在访问第三方服务前，建议你阅读其隐私政策，并谨慎提供个人信息。'
        ]
      },
      {
        title: '政策更新',
        paragraphs: [
          '我们可能因服务调整或法律要求更新本政策。更新后的政策会在本页面公布，并标注更新时间。',
          '若变更会实质性影响你的权利，我们会通过站内提示或其他合理方式通知你。继续使用服务即表示你接受更新后的政策。'
        ]
      },
      {
        title: '联系我们',
        paragraphs: [
          '如你对个人信息处理有疑问、投诉或请求，可通过平台内订单页或客服入口联系我们，并注明「隐私请求」。我们会尽快核实并回复。'
        ]
      }
    ],
    other: {
      to: '/disclaimer',
      label: '免责声明',
      desc: '了解 NFC 兼容性、内容责任与服务可用性的边界。'
    }
  },
  '/disclaimer': {
    eyebrow: 'LEGAL · DISCLAIMER',
    title: '免责声明',
    lead: '在使用 Tapiito 之前，请了解 NFC 兼容性、内容责任与服务可用性的边界。',
    meta: '平台使用说明',
    updated: '2026 年 10 月 21 日',
    icon: Scale,
    intro:
      '本免责声明适用于 Tapiito 平台、NFC 标签及相关页面。使用本服务前，请仔细阅读。若你不同意本声明，请停止使用相关服务。',
    sections: [
      {
        title: '服务说明',
        paragraphs: [
          'Tapiito 提供 NFC 标签与配套展示页面，帮助用户通过手机碰触打开名片、菜单、门店、作品、宠物铭牌等内容。',
          '平台根据用户提供的信息创建和配置展示页面；用户不直接创建或编辑页面内容，相关创建、修改和发布由平台完成。',
          '服务效果取决于标签状态、手机 NFC 能力、浏览器、网络环境及页面内容。我们努力保障服务稳定，但不对所有设备和使用场景作出保证。'
        ]
      },
      {
        title: 'NFC 兼容性与硬件',
        paragraphs: [
          '并非所有手机都支持 NFC，部分设备需要在系统设置中手动开启 NFC，或对标签位置、手机壳厚度较为敏感。',
          '标签可能因弯折、潮湿、高温、强磁场、磨损或粘贴不当而损坏或失效。因硬件、设备设置或标签物理损坏导致无法读取的，不属于平台可控范围。'
        ]
      },
      {
        title: '用户内容与责任',
        paragraphs: [
          '标签页面的内容由用户提供，平台根据用户提供的信息创建和维护。你应确保你提供的内容真实、准确、合法，并拥有使用相关文字、图片、商标和其他素材的权利。',
          '你不得向我们提供违法、侵权、虚假、欺诈、骚扰、恶意引流或其他损害他人权益的内容。我们有权在收到投诉或发现违规时拒绝创建、限制、删除内容或暂停服务。'
        ]
      },
      {
        title: '未成年人信息与儿童标签',
        paragraphs: [
          '未满 14 周岁的未成年人不应自行提交信息或独立使用本服务。儿童信息类标签应由监护人提出需求、提供并确认信息，由平台协助创建和管理。',
          '监护人应确保其提供的未成年人信息真实、必要，并有权同意平台按照本声明及《隐私政策》处理相关信息。未成年人本人不应自行提交或修改相关内容。',
          '若我们发现未经监护人同意收集了未成年人信息，会尽快删除或采取其他适当措施。'
        ]
      },
      {
        title: '信息准确性与时效',
        paragraphs: [
          '标签内容由用户提供，平台按用户请求更新。用户应及时告知我们内容变更；我们不对用户提供内容的准确性、完整性或时效性作保证，也不对因内容过期、错误或遗漏造成的损失承担责任。',
          '儿童信息、老人紧急联系人和宠物医疗信息等标签仅作为辅助展示，不应替代监护人、医疗机构、公安机关或其他专业机构的正式记录与判断。'
        ]
      },
      {
        title: '第三方服务与链接',
        paragraphs: [
          '平台可能接入支付、短信、地图、统计、社交平台等第三方服务，标签页面也可能包含第三方链接。',
          '第三方服务由其运营者独立提供并承担责任。我们不对第三方服务的内容、可用性、安全性或信息处理行为作保证。'
        ]
      },
      {
        title: '服务可用性',
        paragraphs: [
          '我们会尽力维持服务正常运行，但可能因系统维护、升级、网络故障、设备故障、黑客攻击、疫情、自然灾害、政府行为或其他不可抗力而中断。',
          '在法律允许的范围内，我们不对服务中断、延迟、数据丢失或无法访问承担超出法定范围的责任。建议你对自己的重要内容保留备份。'
        ]
      },
      {
        title: '知识产权',
        paragraphs: [
          'Tapiito 的页面设计、界面、代码、商标和品牌素材归平台或相应权利人所有。未经授权，不得复制、修改、传播或用于商业用途。',
          '你向我们提供的内容仍归你或原权利人所有。你授权我们在提供服务所需的范围内展示、存储、创建页面和处理这些内容。'
        ]
      },
      {
        title: '责任限制',
        paragraphs: [
          '在法律允许的最大范围内，我们不对因使用或无法使用本服务产生的间接损失、附带损失、利润损失、商誉损失或数据损失承担责任。',
          '本条不排除或限制依法不能排除或限制的责任，包括因故意或重大过失造成的人身损害责任。若相关服务为付费服务，责任上限以你为该服务实际支付的费用为限。'
        ]
      },
      {
        title: '紧急信息特别提示',
        paragraphs: [
          '老人、儿童、宠物等防走失或紧急联系标签是辅助工具，不具备 GPS 定位、实时追踪或自动报警功能，也不能替代 110、120 等紧急服务。',
          '如遇人员走失、医疗急症或其他紧急情况，请立即联系公安机关、医疗机构或相关专业机构，不要仅依赖标签页面。'
        ]
      },
      {
        title: '声明更新与联系',
        paragraphs: [
          '我们可能根据服务变化或法律要求更新本声明。更新后的内容会在本页面公布，并标注更新时间。',
          '如你对本声明有疑问，可通过平台内订单页或客服入口联系我们。'
        ]
      }
    ],
    other: {
      to: '/privacy',
      label: '隐私政策',
      desc: '了解我们如何收集、使用和保护你的个人信息。'
    }
  }
}

const page = computed(() => {
  const path = route.path.replace(/\/+$/, '') || '/'
  return PAGES[path] || PAGES['/privacy']
})

const pad = (value) => String(value).padStart(2, '0')

watch(
  page,
  (current) => {
    document.title = `${current.title} · Tapiito`
  },
  { immediate: true }
)

onBeforeUnmount(() => {
  document.title = 'Tapiito'
})
</script>

<style scoped>
.legal-page {
  --ink: #14181d;
  --cream: #fbfaf7;
  --orange: #ff6b35;
  --sun: #ffd166;
  --mint: #7bdff2;
  min-height: 100vh;
  background: var(--cream);
  color: var(--ink);
  font-family: 'Geist', ui-sans-serif, system-ui, sans-serif;
}

.legal-shell {
  width: min(1120px, calc(100% - 48px));
  margin: 0 auto;
}

.legal-nav-wrap {
  position: sticky;
  top: 0;
  z-index: 40;
  border-bottom: 1px solid rgba(20, 24, 29, 0.07);
  background: rgba(251, 250, 247, 0.86);
  backdrop-filter: blur(16px);
}

.legal-nav {
  display: flex;
  height: 68px;
  align-items: center;
  justify-content: space-between;
}

.legal-brand {
  display: flex;
  align-items: center;
  gap: 10px;
  font-size: 19px;
  font-weight: 650;
  letter-spacing: -0.03em;
  color: var(--ink);
  text-decoration: none;
}

.legal-brand-mark {
  display: inline-flex;
  height: 32px;
  width: 32px;
  align-items: center;
  justify-content: center;
  border-radius: 10px;
  background: var(--ink);
  box-shadow: 0 2px 8px rgba(20, 24, 29, 0.16);
}

.legal-brand-mark-sm {
  height: 20px;
  width: 20px;
  border-radius: 6px;

  background: #15151B;}

.legal-nav-links {
  display: flex;
  align-items: center;
  gap: 6px;
}

.legal-nav-link {
  display: none;
  padding: 9px 14px;
  border-radius: 999px;
  font-size: 12px;
  font-weight: 500;
  color: rgba(20, 24, 29, 0.62);
  text-decoration: none;
  transition: background 0.2s ease, color 0.2s ease;
}

.legal-nav-link:hover {
  background: rgba(20, 24, 29, 0.05);
  color: var(--ink);
}

.legal-nav-cta {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  padding: 9px 16px;
  border-radius: 999px;
  background: var(--ink);
  color: #fff;
  font-size: 12px;
  font-weight: 600;
  text-decoration: none;
  transition: transform 0.2s ease, background 0.2s ease;
}

.legal-nav-cta:hover {
  background: var(--orange);
  transform: translateY(-1px);
}

.legal-hero {
  position: relative;
  overflow: hidden;
  border-bottom: 1px solid rgba(20, 24, 29, 0.08);
  background:
    radial-gradient(circle at 18% 8%, rgba(255, 107, 53, 0.09), transparent 40%),
    radial-gradient(circle at 88% 20%, rgba(255, 209, 102, 0.22), transparent 42%),
    var(--cream);
}

.legal-hero-grid {
  position: absolute;
  inset: 0;
  background-image:
    linear-gradient(rgba(20, 24, 29, 0.04) 1px, transparent 1px),
    linear-gradient(90deg, rgba(20, 24, 29, 0.04) 1px, transparent 1px);
  background-size: 56px 56px;
  mask-image: linear-gradient(to bottom, rgba(0, 0, 0, 0.6), transparent 84%);
  pointer-events: none;
}

.legal-hero-glow {
  position: absolute;
  border-radius: 999px;
  filter: blur(80px);
  pointer-events: none;
}

.legal-hero-glow-a {
  top: -120px;
  left: -80px;
  height: 300px;
  width: 300px;
  background: rgba(255, 107, 53, 0.18);
}

.legal-hero-glow-b {
  top: -80px;
  right: -40px;
  height: 280px;
  width: 280px;
  background: rgba(123, 223, 242, 0.17);
}

.legal-hero-inner {
  position: relative;
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 48px;
  padding-top: 76px;
  padding-bottom: 72px;
}

.legal-kicker {
  margin: 0 0 18px;
  font-family: ui-monospace, SFMono-Regular, Menlo, Consolas, monospace;
  font-size: 10px;
  font-weight: 600;
  letter-spacing: 0.2em;
  color: var(--orange);
}

.legal-title {
  margin: 0;
  font-size: clamp(42px, 6vw, 68px);
  font-weight: 650;
  line-height: 1.04;
  letter-spacing: -0.055em;
}

.legal-title-dot {
  color: var(--orange);
}

.legal-lead {
  max-width: 620px;
  margin: 22px 0 0;
  font-size: 15px;
  line-height: 1.85;
  color: rgba(20, 24, 29, 0.58);
}

.legal-meta {
  display: flex;
  flex-wrap: wrap;
  align-items: center;
  gap: 12px;
  margin-top: 26px;
  font-family: ui-monospace, SFMono-Regular, Menlo, Consolas, monospace;
  font-size: 10px;
  letter-spacing: 0.06em;
  color: rgba(20, 24, 29, 0.4);
}

.legal-meta-pill {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  padding: 6px 10px;
  border: 1px solid rgba(20, 24, 29, 0.1);
  border-radius: 999px;
  background: rgba(255, 255, 255, 0.62);
  color: rgba(20, 24, 29, 0.62);
}

.legal-hero-seal {
  position: relative;
  display: none;
  height: 190px;
  width: 190px;
  flex: 0 0 190px;
  align-items: center;
  justify-content: center;
}

.legal-seal-ring {
  position: absolute;
  inset: 0;
  border: 1px dashed rgba(20, 24, 29, 0.16);
  border-radius: 999px;
  animation: legal-spin 28s linear infinite;
}

.legal-seal-ring-inner {
  inset: 16px;
  border-style: solid;
  border-color: rgba(255, 107, 53, 0.18);
  animation-direction: reverse;
  animation-duration: 20s;
}

.legal-seal-core {
  position: relative;
  display: flex;
  height: 122px;
  width: 122px;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  gap: 8px;
  border-radius: 999px;
  background: rgba(255, 255, 255, 0.75);
  box-shadow: 0 18px 50px rgba(20, 24, 29, 0.08);
  color: var(--orange);
}

.legal-seal-core span {
  font-family: ui-monospace, SFMono-Regular, Menlo, Consolas, monospace;
  font-size: 8px;
  line-height: 1.5;
  letter-spacing: 0.18em;
  text-align: center;
  color: rgba(20, 24, 29, 0.46);
}

.legal-layout {
  display: grid;
  gap: 56px;
  padding-top: 64px;
  padding-bottom: 96px;
}

.legal-aside {
  display: none;
}

.legal-aside-inner {
  position: sticky;
  top: 100px;
}

.legal-aside-title {
  margin: 0 0 16px;
  font-family: ui-monospace, SFMono-Regular, Menlo, Consolas, monospace;
  font-size: 9px;
  font-weight: 600;
  letter-spacing: 0.18em;
  color: rgba(20, 24, 29, 0.36);
}

.legal-toc-link {
  display: flex;
  align-items: baseline;
  gap: 9px;
  padding: 6px 0;
  font-size: 12px;
  line-height: 1.5;
  color: rgba(20, 24, 29, 0.52);
  text-decoration: none;
  transition: color 0.2s ease;
}

.legal-toc-link span {
  font-family: ui-monospace, SFMono-Regular, Menlo, Consolas, monospace;
  font-size: 9px;
  color: rgba(255, 107, 53, 0.6);
}

.legal-toc-link:hover {
  color: var(--ink);
}

.legal-aside-note {
  margin-top: 28px;
  padding: 16px;
  border: 1px solid rgba(20, 24, 29, 0.08);
  border-radius: 14px;
  background: rgba(255, 255, 255, 0.55);
}

.legal-aside-note p {
  margin: 0 0 6px;
  font-size: 11px;
  font-weight: 650;
}

.legal-aside-note span {
  display: block;
  font-size: 10px;
  line-height: 1.7;
  color: rgba(20, 24, 29, 0.46);
}

.legal-article {
  min-width: 0;
}

.legal-intro {
  margin: 0 0 42px;
  padding-bottom: 32px;
  border-bottom: 1px solid rgba(20, 24, 29, 0.1);
  font-size: 16px;
  line-height: 1.9;
  color: rgba(20, 24, 29, 0.64);
}

.legal-section {
  scroll-margin-top: 96px;
  padding: 30px 0 34px;
  border-bottom: 1px solid rgba(20, 24, 29, 0.08);
}

.legal-section-head {
  display: flex;
  align-items: baseline;
  gap: 14px;
  margin-bottom: 16px;
}

.legal-section-num {
  font-family: ui-monospace, SFMono-Regular, Menlo, Consolas, monospace;
  font-size: 10px;
  font-weight: 600;
  letter-spacing: 0.1em;
  color: var(--orange);
}

.legal-section h2 {
  margin: 0;
  font-size: 20px;
  font-weight: 650;
  letter-spacing: -0.025em;
}

.legal-p {
  margin: 0 0 12px;
  font-size: 14px;
  line-height: 1.95;
  color: rgba(20, 24, 29, 0.62);
}

.legal-p:last-child {
  margin-bottom: 0;
}

.legal-list {
  margin: 14px 0 0;
  padding: 0;
  list-style: none;
}

.legal-list li {
  position: relative;
  margin: 0 0 10px;
  padding-left: 18px;
  font-size: 14px;
  line-height: 1.85;
  color: rgba(20, 24, 29, 0.62);
}

.legal-list li::before {
  content: '';
  position: absolute;
  top: 11px;
  left: 2px;
  height: 5px;
  width: 5px;
  border-radius: 999px;
  background: var(--orange);
}

.legal-contact {
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  gap: 22px;
  margin-top: 48px;
  padding: 28px;
  border-radius: 22px;
  background:
    radial-gradient(circle at 92% 0%, rgba(255, 107, 53, 0.24), transparent 38%),
    var(--ink);
  color: #fff;
}

.legal-contact-kicker {
  margin: 0 0 8px;
  font-family: ui-monospace, SFMono-Regular, Menlo, Consolas, monospace;
  font-size: 9px;
  letter-spacing: 0.18em;
  color: rgba(255, 209, 102, 0.8);
}

.legal-contact h2 {
  margin: 0 0 10px;
  font-size: 26px;
  font-weight: 600;
  letter-spacing: -0.035em;
}

.legal-contact p {
  max-width: 520px;
  margin: 0;
  font-size: 13px;
  line-height: 1.8;
  color: rgba(255, 255, 255, 0.58);
}

.legal-contact-link {
  display: inline-flex;
  align-items: center;
  gap: 7px;
  padding: 10px 15px;
  border-radius: 999px;
  background: var(--orange);
  color: #fff;
  font-size: 12px;
  font-weight: 600;
  text-decoration: none;
  transition: transform 0.2s ease, background 0.2s ease;
}

.legal-contact-link:hover {
  background: #f25a22;
  transform: translateY(-1px);
}

.legal-related {
  margin-top: 48px;
}

.legal-related > p {
  margin: 0 0 12px;
  font-family: ui-monospace, SFMono-Regular, Menlo, Consolas, monospace;
  font-size: 9px;
  letter-spacing: 0.18em;
  color: rgba(20, 24, 29, 0.36);
}

.legal-related-card {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 18px;
  margin-bottom: 10px;
  padding: 17px 18px;
  border: 1px solid rgba(20, 24, 29, 0.09);
  border-radius: 15px;
  background: rgba(255, 255, 255, 0.48);
  color: var(--ink);
  text-decoration: none;
  transition: border-color 0.2s ease, background 0.2s ease, transform 0.2s ease;
}

.legal-related-card:hover {
  border-color: rgba(255, 107, 53, 0.35);
  background: #fff;
  transform: translateY(-1px);
}

.legal-related-card strong {
  display: block;
  margin-bottom: 4px;
  font-size: 13px;
  font-weight: 650;
}

.legal-related-card small {
  display: block;
  font-size: 11px;
  line-height: 1.6;
  color: rgba(20, 24, 29, 0.46);
}

.legal-related-card svg {
  flex: 0 0 auto;
  color: var(--orange);
}

.legal-back {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  margin-top: 22px;
  font-size: 11px;
  color: rgba(20, 24, 29, 0.46);
  text-decoration: none;
  transition: color 0.2s ease;
}

.legal-back:hover {
  color: var(--orange);
}

.legal-footer {
  border-top: 1px solid rgba(20, 24, 29, 0.08);
}

.legal-footer-inner {
  display: flex;
  flex-direction: column;
  gap: 16px;
  padding-top: 28px;
  padding-bottom: 28px;
  font-size: 11.5px;
  color: rgba(20, 24, 29, 0.4);
}

.legal-footer-brand {
  display: flex;
  align-items: center;
  gap: 8px;
  color: rgba(20, 24, 29, 0.7);
  font-weight: 600;
}

.legal-footer-tagline {
  font-weight: 400;
  color: rgba(20, 24, 29, 0.32);
}

.legal-footer-links {
  display: flex;
  flex-wrap: wrap;
  gap: 8px 18px;
}

.legal-footer-links a {
  color: rgba(20, 24, 29, 0.46);
  text-decoration: none;
  transition: color 0.2s ease;
}

.legal-footer-links a:hover {
  color: var(--orange);
}

.legal-footer-copy {
  margin: 0;
  font-family: ui-monospace, SFMono-Regular, Menlo, Consolas, monospace;
  font-size: 9.5px;
  letter-spacing: 0.08em;
  color: rgba(20, 24, 29, 0.28);
}

@keyframes legal-spin {
  to {
    transform: rotate(360deg);
  }
}

@media (min-width: 640px) {
  .legal-nav-link {
    display: inline-flex;
  }

  .legal-contact {
    flex-direction: row;
    align-items: center;
    justify-content: space-between;
    padding: 34px 38px;
  }
}

@media (min-width: 960px) {
  .legal-hero-seal {
    display: flex;
  }

  .legal-layout {
    grid-template-columns: 220px minmax(0, 1fr);
    gap: 72px;
  }

  .legal-aside {
    display: block;
  }

  .legal-footer-inner {
    flex-direction: row;
    align-items: center;
    justify-content: space-between;
  }
}

@media (max-width: 639px) {
  .legal-shell {
    width: calc(100% - 32px);
  }

  .legal-nav {
    height: 60px;
  }

  .legal-brand {
    font-size: 17px;
  }

  .legal-brand-mark {
    height: 30px;
    width: 30px;
  }

  .legal-nav-cta {
    padding: 8px 13px;
    font-size: 11px;
  }

  .legal-hero-inner {
    padding-top: 52px;
    padding-bottom: 50px;
  }

  .legal-title {
    font-size: 40px;
  }

  .legal-lead {
    font-size: 14px;
  }

  .legal-layout {
    padding-top: 42px;
    padding-bottom: 64px;
  }

  .legal-intro {
    font-size: 15px;
    margin-bottom: 30px;
    padding-bottom: 24px;
  }

  .legal-section {
    padding: 26px 0 28px;
  }

  .legal-section h2 {
    font-size: 18px;
  }

  .legal-contact {
    padding: 24px;
  }

  .legal-contact h2 {
    font-size: 23px;
  }
}

@media (prefers-reduced-motion: reduce) {
  .legal-seal-ring {
    animation: none;
  }

  .legal-nav-cta,
  .legal-contact-link,
  .legal-related-card {
    transition: none;
  }
}
</style>
