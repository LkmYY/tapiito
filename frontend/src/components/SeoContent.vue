<template>
  <section class="seo-section">
    <div class="seo-inner">
      <!-- 1. 概念科普 -->
      <article class="seo-block">
        <h2 class="seo-h2">NFC 标签是什么？为什么碰一下就能打开内容</h2>
        <p>
          NFC 是「近场通信」（Near Field Communication）的缩写，一种工作在 13.56MHz
          的短距离无线技术。它和我们每天用的公交卡、门禁卡是同一类技术，区别在于 NFC
          标签里可以写入一个网址。
        </p>
        <p>
          Tapiito 标签内部是一颗带天线的芯片，<strong>不需要电池，也不需要充电</strong>。手机靠近时，
          手机发出的电磁场会给芯片供电，芯片随即把预先写入的网址回传给手机，手机自动打开浏览器加载页面。
          整个过程通常在半秒内完成，用户只需要把手机背面贴近标签。
        </p>
        <p>
          和二维码相比，NFC 的优势是不需要对准、不需要光线、不怕磨损污损。二维码印在纸上一旦淋湿、
          折皱或被贴纸覆盖就读不出来；NFC 标签可以封装在 PVC 卡片、亚克力牌、金属铭牌里，
          在户外、厨房、宠物项圈这些场景下依然稳定可用。缺点是手机必须支持 NFC——
          目前 iPhone 7 及以后机型、绝大多数安卓中高端机型都已支持，且从 iOS 14 起
          iPhone 可以直接在后台读取 NFC 标签，无需打开任何 App。
        </p>
      </article>

      <!-- 2. 九种类型 -->
      <article class="seo-block">
        <h2 class="seo-h2">九种标签，分别适合什么场景</h2>
        <div class="seo-types">
          <div v-for="t in typeDetails" :key="t.name" class="seo-type">
            <h3 class="seo-h3">{{ t.name }}</h3>
            <p>{{ t.desc }}</p>
          </div>
        </div>
      </article>

      <!-- 3. 使用流程 -->
      <article class="seo-block">
        <h2 class="seo-h2">从下单到用上，一共三步</h2>
        <ol class="seo-steps">
          <li>
            <strong>在线填内容。</strong>在下单页选择标签类型和视觉风格，填写要展示的信息。
            右侧会实时预览手机上看到的样子，不满意可以随时换风格，已填的内容不会丢。
          </li>
          <li>
            <strong>我们写入并发货。</strong>收到订单后把内容写入标签芯片，做成品检后寄出。
            首次写入包含在售价内，不需要你自己买读写器。
          </li>
          <li>
            <strong>贴上去就能用。</strong>收到后贴在名片夹、工位、餐桌、宠物项圈或随身物品上。
            手机碰一下即可打开，无需下载 App、无需注册账号。
          </li>
        </ol>
        <p class="seo-note">
          后续内容需要修改时，可以直接在后台改，改完立即生效，不用重新买标签。
          每枚标签含 1 次免费内容调整。
        </p>
      </article>

      <!-- 4. 隐私与安全 -->
      <article class="seo-block">
        <h2 class="seo-h2">关于隐私：谁能看到你的内容</h2>
        <p>
          Tapiito 的每一枚标签都可以单独设置访问权限，四种模式自由组合：
          <strong>完全公开</strong>（任何人碰一下都能看）、<strong>暗号保护</strong>
          （需要输入你设置的暗号）、<strong>一次性链接</strong>（链接只能被第一个打开的人使用，
          转发给别人会失效）、<strong>暗号加一次性链接</strong>（两者同时生效）。
        </p>
        <p>
          夜店交际场景通常用「暗号 + 一次性链接」：只有知道当晚暗号的人能打开，而且转发出去的链接
          在对方手机上是打不开的。求助卡、宠物铭牌这类场景则用「一次性链接 + 角标水印」，
          保证任何善意路人都能第一时间扫开看到紧急联系方式，同时水印可以防止截图外传。
          普通商务名片建议直接用公开模式，扫码即看，没有任何摩擦。
        </p>
        <p>
          暗号在服务器上以 bcrypt 哈希形式存储，我们不保存明文，即使数据库泄露也无法还原。
          同一 IP 或同一标签连续输错三次会锁定三十分钟，防止暴力尝试。
        </p>
      </article>

      <!-- 5. FAQ -->
      <article class="seo-block">
        <h2 class="seo-h2">常见问题</h2>
        <div class="seo-faq">
          <details v-for="(q, i) in faqs" :key="i" class="seo-faq-item">
            <summary>{{ q.q }}</summary>
            <p>{{ q.a }}</p>
          </details>
        </div>
      </article>
    </div>
  </section>
</template>

<script setup>
const typeDetails = [
  {
    name: '商务名片',
    desc:
      '最常见的一类。碰一下展示姓名、职位、公司、电话、邮箱、网站和微信号，访客可以直接保存到手机通讯录，不用手动输入。适合销售、设计师、咨询顾问、创业者，以及任何需要频繁交换联系方式的人。',
  },
  {
    name: '门店菜单',
    desc:
      '贴在餐桌角或门口，顾客碰一下就能看到菜品、价格、图片和推荐。相比桌贴二维码，不用对准、不怕油污水渍，翻台时也不需要反复擦拭。适合餐厅、咖啡馆、酒吧、茶室。',
  },
  {
    name: '企业官网',
    desc:
      '把公司介绍、产品线、联系方式、地图位置浓缩到一个碰触即开的页面。适合展会、招商会、门店前台，替代厚重的宣传册，且内容随时可改，不用重印。',
  },
  {
    name: '产品展示',
    desc:
      '贴在实物旁边，碰一下看到参数、使用方法、售后说明、购买链接。适合数码产品、家电、家具、手工作品，降低导购的解释成本。',
  },
  {
    name: '夜店交际',
    desc:
      '带暗号保护的一类。只有知道当晚暗号的人才能打开，且链接一次性使用，转发无效。适合酒吧、夜店、派对场景下有限度地交换联系方式。',
  },
  {
    name: '求助信息卡',
    desc:
      '给独居老人、慢病患者、语言不通的旅行者使用。路人碰一下即可看到紧急联系人、病史、用药禁忌、过敏源，并可直接拨号。支持上传现场照片和留言反馈。',
  },
  {
    name: '宠物铭牌',
    desc:
      '挂在项圈上，捡到宠物的人碰一下就能看到宠物名字、主人电话、走失酬谢说明。相比刻字吊牌，信息可随时更新，也不受字数限制。',
  },
  {
    name: '儿童信息卡',
    desc:
      '放在书包、校服内袋或手环上。走失时旁人碰一下即可联系家长，同时可以标注过敏史、接送人信息。内容仅监护人可改。',
  },
  {
    name: '菜谱',
    desc:
      '贴在厨房台面或食材包装上，碰一下看到分步做法、火候时间、食材用量。适合私厨、烘焙工作室、食品品牌做延伸内容。',
  },
]

const faqs = [
  {
    q: '我的手机支持 NFC 吗？',
    a: 'iPhone 7 及之后的机型都支持，且 iOS 14 起可以直接在后台读取，不需要打开任何 App。安卓方面，绝大多数中高端机型（华为、小米、OPPO、vivo、三星等）都支持，可在设置里搜索「NFC」确认。少数入门机型不支持，这种情况下可以用标签背面的二维码作为备选。',
  },
  {
    q: '需要装 App 吗？需要充电吗？',
    a: '都不需要。标签是被动式芯片，靠手机发出的电磁场供电，没有电池也不会没电。访客也不需要安装任何 App，碰一下直接打开浏览器。',
  },
  {
    q: '内容以后想改怎么办？',
    a: '在后台修改即可，保存后立即生效，不用重新买标签。每枚标签含 1 次免费内容调整，之后可以按次购买。',
  },
  {
    q: '标签能贴在金属表面吗？',
    a: '普通标签贴在金属上会被干扰，读取距离明显缩短。如果你的使用场景是金属表面（如冰箱、机箱、金属工牌），下单时备注一下，我们会用抗金属专用标签。',
  },
  {
    q: '能用多久？会消磁吗？',
    a: 'NFC 标签没有磁性、不会消磁，数据写入后可以保存十年以上。物理寿命主要取决于封装：PVC 卡片日常使用三到五年没问题，户外场景建议选防水封装的亚克力牌或金属牌。',
  },
  {
    q: '别人能改我的标签内容吗？',
    a: '不会。出厂时我们会把标签设为只读锁定，写入后任何设备都无法覆盖。只有你在后台修改内容时，页面才会更新（标签里存的是一个固定网址，内容由服务器下发）。',
  },
  {
    q: '一次性链接是什么原理？转发真的会失效吗？',
    a: '每次打开都会校验一次凭证，凭证首次使用时绑定到当前浏览器会话。你自己刷新页面没问题，但把链接发给别人，对方浏览器没有对应会话，就会提示链接已失效。这样既不影响本人使用，又能阻止转发。',
  },
  {
    q: '暗号忘了怎么办？',
    a: '在后台重新设置即可，旧暗号立即失效。如果你不是标签所有者，只能联系标签主人。',
  },
  {
    q: '发货要多久？',
    a: '内容确认后一般一到两个工作日内发出。定制款（刻字、异形、金属牌）需要三到五个工作日。',
  },
  {
    q: '支持哪些设计风格？',
    a: '每种标签类型都有 30 种视觉风格，从极简白、杂志排版到赛博朋克、像素游戏机、动漫角色卡都有覆盖。下单时可以逐个切换预览，选定后还可随时更换，不影响已填内容。',
  },
]
</script>

<style scoped>
.seo-section {
  border-top: 1px solid rgba(20, 24, 29, 0.07);
  background: #fbfaf7;
  padding: 72px 24px 88px;
}
.seo-inner {
  max-width: 880px;
  margin: 0 auto;
}
.seo-block + .seo-block {
  margin-top: 56px;
}
.seo-h2 {
  font-size: 24px;
  font-weight: 600;
  letter-spacing: -0.01em;
  color: #14181d;
}
.seo-block > p {
  margin-top: 14px;
  font-size: 15px;
  line-height: 1.85;
  color: rgba(20, 24, 29, 0.66);
}
.seo-types {
  margin-top: 22px;
  display: grid;
  gap: 22px;
}
.seo-type {
  padding-left: 14px;
  border-left: 2px solid rgba(255, 107, 53, 0.35);
}
.seo-h3 {
  font-size: 15.5px;
  font-weight: 600;
  color: #14181d;
}
.seo-type p {
  margin-top: 6px;
  font-size: 14px;
  line-height: 1.8;
  color: rgba(20, 24, 29, 0.6);
}
.seo-steps {
  margin-top: 20px;
  padding-left: 20px;
  display: grid;
  gap: 14px;
}
.seo-steps li {
  font-size: 15px;
  line-height: 1.85;
  color: rgba(20, 24, 29, 0.66);
}
.seo-note {
  margin-top: 16px;
  padding: 12px 16px;
  border-radius: 10px;
  background: rgba(255, 107, 53, 0.06);
  font-size: 13.5px;
  line-height: 1.75;
  color: rgba(20, 24, 29, 0.6);
}
.seo-faq {
  margin-top: 22px;
  display: grid;
  gap: 2px;
}
.seo-faq-item {
  border-bottom: 1px solid rgba(20, 24, 29, 0.07);
  padding: 14px 0;
}
.seo-faq-item summary {
  cursor: pointer;
  font-size: 15px;
  font-weight: 500;
  color: #14181d;
  list-style: none;
}
.seo-faq-item summary::-webkit-details-marker {
  display: none;
}
.seo-faq-item summary::before {
  content: '+';
  display: inline-block;
  width: 18px;
  color: rgba(255, 107, 53, 0.8);
}
.seo-faq-item[open] summary::before {
  content: '−';
}
.seo-faq-item p {
  margin-top: 10px;
  padding-left: 18px;
  font-size: 14px;
  line-height: 1.85;
  color: rgba(20, 24, 29, 0.6);
}
</style>
