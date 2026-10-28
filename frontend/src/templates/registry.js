/**
 * 模板类型注册表
 *
 * 每个类型集中描述：图标、可视化表单结构（schema）、30 种风格（含组件）、示例数据。
 * 新增类型 = 新建一个 manifest 文件 + 在 index.js 里 import 一行。
 */
import {
  text, textarea, tel, email, url, price, toggle, select, image, tags, keyValue, list, group,
  SPICY_OPTIONS,
} from './schema.js'

import {
  Briefcase, Baby, Siren, Globe, ShoppingBag, Palette, UtensilsCrossed,
  Martini, Sparkles, PawPrint, Eye,
} from 'lucide-vue-next'

// —— 页面组件 ——
import BusinessCardPage from '../pages/templates/BusinessCardPage.vue'
import BusinessCardModern from '../pages/templates/BusinessCardModern.vue'
import BusinessCardCreative from '../pages/templates/BusinessCardCreative.vue'
import BusinessCardMinimal from '../pages/templates/BusinessCardMinimal.vue'
import BusinessCardCorporate from '../pages/templates/BusinessCardCorporate.vue'

import ChildInfoPage from '../pages/templates/ChildInfoPage.vue'
import ChildInfoCartoon from '../pages/templates/ChildInfoCartoon.vue'
import ChildInfoNature from '../pages/templates/ChildInfoNature.vue'
import ChildInfoWarm from '../pages/templates/ChildInfoWarm.vue'
import ChildInfoStarry from '../pages/templates/ChildInfoStarry.vue'

import EmergencyContactPage from '../pages/templates/EmergencyContactPage.vue'
import EmergencyMedical from '../pages/templates/EmergencyMedical.vue'
import EmergencySOS from '../pages/templates/EmergencySOS.vue'
import EmergencyFamily from '../pages/templates/EmergencyFamily.vue'
import EmergencyPet from '../pages/templates/EmergencyPet.vue'

import CompanyWebsitePage from '../pages/templates/CompanyWebsitePage.vue'
import CompanyWebsiteTech from '../pages/templates/CompanyWebsiteTech.vue'
import CompanyWebsiteElegant from '../pages/templates/CompanyWebsiteElegant.vue'
import CompanyWebsiteVibrant from '../pages/templates/CompanyWebsiteVibrant.vue'
import CompanyWebsiteFullpage from '../pages/templates/CompanyWebsiteFullpage.vue'

import ProductDisplayPage from '../pages/templates/ProductDisplayPage.vue'
import ProductPremium from '../pages/templates/ProductPremium.vue'
import ProductShowcase from '../pages/templates/ProductShowcase.vue'
import ProductComparison from '../pages/templates/ProductComparison.vue'
import ProductStory from '../pages/templates/ProductStory.vue'

import RecipePage from '../pages/templates/RecipePage.vue'
import RecipeChalk from '../pages/templates/RecipeChalk.vue'
import RecipeNightMarket from '../pages/templates/RecipeNightMarket.vue'
import RecipeBreakfast from '../pages/templates/RecipeBreakfast.vue'
import RecipeFineDining from '../pages/templates/RecipeFineDining.vue'

import CustomCompanyPage from '../pages/templates/divPage/CustomCompanyPage.vue'
import CustomRestaurant from '../pages/templates/divPage/CustomRestaurant.vue'
import CustomSalon from '../pages/templates/divPage/CustomSalon.vue'
import CustomFitness from '../pages/templates/divPage/CustomFitness.vue'
import CustomEducation from '../pages/templates/divPage/CustomEducation.vue'

// —— 扩展风格模块（每个类型独立维护，向 30 款扩充） ——
import businessCardExtra from './extensions/business_card.js'
import childInfoExtra from './extensions/child_info.js'
import emergencyContactExtra from './extensions/emergency_contact.js'
import companyWebsiteExtra from './extensions/company_website.js'
import productDisplayExtra from './extensions/product_display.js'
import recipeExtra from './extensions/recipe.js'
import customExtra from './extensions/custom.js'
import nightclubStyles from './extensions/nightclub.js'
import animeCardStyles from './extensions/anime_card.js'
import petTagStyles from './extensions/pet_tag.js'
import weirdStyles from './extensions/weird.js'

/* ------------------------------------------------------------------ */
/* 复用片段                                                            */
/* ------------------------------------------------------------------ */

const contactFields = [
  text('name', '姓名', { required: true, placeholder: '例如：陈建国' }),
  text('relation', '关系', { placeholder: '例如：父亲 / 母亲 / 邻居' }),
  tel('phone', '电话', { required: true, placeholder: '138 0000 0000' }),
]

const recipeItemFields = [
  text('name', '菜名', { required: true, placeholder: '例如：焖肉面' }),
  price('price', '价格', { placeholder: '28' }),
  text('desc', '描述', { placeholder: '例如：五花肉焖足三小时，酥而不散' }),
  toggle('signature', '招牌'),
  toggle('vegetarian', '素食'),
  select('spicy', '辣度', SPICY_OPTIONS),
]

/* ------------------------------------------------------------------ */
/* 1. 名片                                                             */
/* ------------------------------------------------------------------ */

const businessCard = {
  key: 'business_card',
  label: '名片',
  icon: Briefcase,
  iconBg: '#ede9fe',
  iconColor: '#7c3aed',
  ambient: 'amb-dark',
  schema: [
    text('name', '姓名', { required: true, placeholder: '例如：林晓峰' }),
    text('title', '职位', { placeholder: '例如：产品设计师' }),
    text('company', '公司', { placeholder: '例如：云图科技' }),
    tel('phone', '手机', { placeholder: '138 0013 8000' }),
    email('email', '邮箱', { placeholder: 'name@company.com' }),
    url('website', '网站', { placeholder: 'https://example.com' }),
    text('wechat', '微信号', { placeholder: 'wechat_id' }),
  ],
  styles: [
    { value: '', label: '经典商务', preview: '深色卡片 + 联系列表', gradient: 'linear-gradient(135deg,#1e293b,#0f172a)', component: BusinessCardPage, ambient: 'amb-dark' },
    { value: 'modern', label: '现代简约', preview: '暖白底 + 卡片网格', gradient: 'linear-gradient(135deg,#fafaf9,#e7e5e4)', component: BusinessCardModern, ambient: 'amb-plain' },
    { value: 'creative', label: '创意撞色', preview: '橙黑斜切 + 大字', gradient: 'linear-gradient(135deg,#ff5a1f,#17130f)', component: BusinessCardCreative, ambient: 'amb-dark' },
    { value: 'minimal', label: '极简黑白', preview: '衬线大字 + 细线', gradient: 'linear-gradient(135deg,#ffffff,#f5f5f4)', component: BusinessCardMinimal, ambient: 'amb-plain' },
    { value: 'corporate', label: '企业品牌', preview: '深蓝 Banner + 信笺', gradient: 'linear-gradient(135deg,#123a8f,#0b2357)', component: BusinessCardCorporate, ambient: 'amb-cool' },
  ],
  samples: {
    '': { name: '林晓峰', title: '产品设计师', company: '云图科技', phone: '138 0013 8000', email: 'xiaofeng@yuntu.design', website: 'https://yuntu.design', wechat: 'xiaofeng_ux' },
  },
}

/* ------------------------------------------------------------------ */
/* 2. 求助信息卡                                                       */
/* ------------------------------------------------------------------ */

const childInfo = {
  key: 'child_info',
  label: '求助信息卡',
  icon: Baby,
  iconBg: '#e0f2fe',
  iconColor: '#0284c7',
  ambient: 'amb-cool',
  schema: [
    textarea('baseInfo', '基本情况', { required: true, rows: 2, placeholder: '例如：姓名、年龄、衣着特征' }),
    text('school', '学校 / 单位', { placeholder: '例如：青禾社区' }),
    textarea('allergies', '过敏史 / 病史', { rows: 2, placeholder: '例如：青霉素过敏；如无请留空' }),
    tel('emergencyPhone', '紧急联系电话', { required: true, placeholder: '138 0000 0000' }),
  ],
  styles: [
    { value: '', label: '天空蓝', preview: '天蓝渐变 + 圆角卡片', gradient: 'linear-gradient(135deg,#7dd3fc,#e0f2fe)', component: ChildInfoPage, ambient: 'amb-cool' },
    { value: 'cartoon', label: '卡通趣味', preview: '柔和色块 + 气泡卡', gradient: 'linear-gradient(135deg,#fbcfe8,#fef08a)', component: ChildInfoCartoon, ambient: 'amb-warm-pink' },
    { value: 'nature', label: '自然清新', preview: '森绿 + 时间轴', gradient: 'linear-gradient(135deg,#86efac,#34d399)', component: ChildInfoNature, ambient: 'amb-warm-green' },
    { value: 'warm', label: '暖阳橙', preview: '橙黄暖色 + 大字号', gradient: 'linear-gradient(135deg,#fdba74,#fbbf24)', component: ChildInfoWarm, ambient: 'amb-warm' },
    { value: 'starry', label: '星空梦幻', preview: '深蓝夜空 + 荧光点', gradient: 'linear-gradient(135deg,#1e1b4b,#4c1d95)', component: ChildInfoStarry, ambient: 'amb-dark' },
  ],
  samples: {
    '': { baseInfo: '李建国，72岁，穿深灰外套，戴黑框眼镜', school: '青禾社区', allergies: '青霉素过敏', emergencyPhone: '138 1024 7788' },
  },
}

/* ------------------------------------------------------------------ */
/* 3. 紧急联系人                                                       */
/* ------------------------------------------------------------------ */

const emergencyContact = {
  key: 'emergency_contact',
  label: '紧急联系人',
  icon: Siren,
  iconBg: '#fee2e2',
  iconColor: '#dc2626',
  ambient: 'amb-dark-red',
  schema: [
    list('contacts', '联系人', {
      itemLabel: '联系人',
      addLabel: '添加联系人',
      fields: contactFields,
      defaultItem: { name: '', relation: '', phone: '' },
    }),
    // 以下字段仅在「医疗专业」风格中展示
    group('医疗信息（仅「医疗专业」风格展示）', [
      select('bloodType', '血型', ['', 'A', 'B', 'AB', 'O', '未知']),
      textarea('allergies', '过敏史', { rows: 2 }),
      textarea('medications', '正在服用的药物', { rows: 2 }),
      textarea('conditions', '既往病史', { rows: 2 }),
    ], { onlyStyles: ['medical'] }),
    // 以下字段仅在「宠物专用」风格中展示
    group('宠物信息（仅「宠物专用」风格展示）', [
      text('petName', '宠物名字'),
      text('breed', '品种'),
      text('chipId', '芯片编号'),
      tel('vetPhone', '宠物医院电话'),
    ], { onlyStyles: ['pet'] }),
  ],
  styles: [
    { value: '', label: '红色警告', preview: '红色头部 + 联系卡', gradient: 'linear-gradient(135deg,#dc2626,#f87171)', component: EmergencyContactPage, ambient: 'amb-dark-red' },
    { value: 'medical', label: '医疗专业', preview: '白底 + 医疗信息', gradient: 'linear-gradient(135deg,#eff6ff,#93c5fd)', component: EmergencyMedical, ambient: 'amb-cool' },
    { value: 'sos', label: 'SOS 高对比', preview: '黑底红字 + 大号码', gradient: 'linear-gradient(135deg,#0a0a0a,#7f1d1d)', component: EmergencySOS, ambient: 'amb-dark-red' },
    { value: 'family', label: '家庭温馨', preview: '暖粉 + 家人卡片', gradient: 'linear-gradient(135deg,#fce7f3,#fef3c7)', component: EmergencyFamily, ambient: 'amb-warm-pink' },
    { value: 'pet', label: '宠物专用', preview: '森绿 + 宠物档案', gradient: 'linear-gradient(135deg,#86efac,#d97706)', component: EmergencyPet, ambient: 'amb-warm-green' },
  ],
  samples: {
    '': {
      contacts: [
        { name: '陈建国', relation: '父亲', phone: '138 1024 7788' },
        { name: '李静', relation: '母亲', phone: '139 2086 3344' },
        { name: '陈爷爷', relation: '祖父', phone: '010-6874 2210' },
      ],
    },
    medical: {
      contacts: [
        { name: '陈建国', relation: '父亲', phone: '138 1024 7788' },
        { name: '李静', relation: '母亲', phone: '139 2086 3344' },
      ],
      bloodType: 'A', allergies: '青霉素过敏', medications: '降压药（每日一次）', conditions: '高血压',
    },
    pet: {
      petName: '豆包', breed: '柯基 · 3 岁', chipId: 'CN-9001-2203-8871', vetPhone: '021-5466 8890',
      contacts: [{ name: '陈晓', relation: '主人', phone: '138 0013 8000' }],
    },
  },
}

/* ------------------------------------------------------------------ */
/* 4. 企业官网                                                         */
/* ------------------------------------------------------------------ */

const companyWebsite = {
  key: 'company_website',
  label: '企业官网',
  icon: Globe,
  iconBg: '#dbeafe',
  iconColor: '#2563eb',
  ambient: 'amb-dark',
  schema: [
    text('name', '企业名称', { required: true, placeholder: '例如：派若科技' }),
    textarea('description', '企业简介', { rows: 3, placeholder: '一两句话说明你们做什么' }),
    url('url', '官网地址', { placeholder: 'https://example.com' }),
    tags('features', '核心优势', { placeholder: '输入后按回车添加，例如：故障提前 72 小时预警' }),
    group('联系方式', [
      tel('contactInfo.phone', '电话'),
      email('contactInfo.email', '邮箱'),
      text('contactInfo.address', '地址'),
    ]),
  ],
  styles: [
    { value: '', label: '经典企业', preview: '冷灰纸底 + 钴蓝', gradient: 'linear-gradient(135deg,#f8fafc,#cbd5e1)', component: CompanyWebsitePage, ambient: 'amb-dark' },
    { value: 'tech', label: '科技风', preview: '终端界面 + 祖母绿', gradient: 'linear-gradient(135deg,#020617,#064e3b)', component: CompanyWebsiteTech, ambient: 'amb-dark-green' },
    { value: 'elegant', label: '典雅商务', preview: '象牙纸 + 衬线', gradient: 'linear-gradient(135deg,#fefaf0,#e7dcc4)', component: CompanyWebsiteElegant, ambient: 'amb-warm' },
    { value: 'vibrant', label: '活力多彩', preview: '奶白 + 电黄色块', gradient: 'linear-gradient(135deg,#fefce8,#facc15)', component: CompanyWebsiteVibrant, ambient: 'amb-warm' },
    { value: 'fullpage', label: '全屏叙事', preview: '分段滚动 + 大数字', gradient: 'linear-gradient(135deg,#14110e,#f5f0e6)', component: CompanyWebsiteFullpage, ambient: 'amb-dark' },
  ],
  samples: {
    '': {
      name: '派若科技',
      description: '工业设备预测性维护平台。用振动频谱与温度数据，在故障发生前 72 小时发出预警。',
      url: 'https://example.com',
      features: ['振动频谱分析', '故障提前 72 小时预警', '支持私有化部署', '兼容 Modbus / OPC-UA'],
      contactInfo: { phone: '400-820-7712', email: 'hello@piro.tech', address: '南京市江北新区研创园 5 栋 12 层' },
    },
    elegant: {
      name: '鼎盛集团',
      description: '三十年匠心品质，铸就行业标杆。我们专注于为客户提供卓越的咨询与投资服务。',
      url: 'https://example.com',
      contactInfo: { phone: '400-888-0000', email: 'info@example.com', address: '上海市浦东新区陆家嘴金融中心' },
    },
    vibrant: {
      name: '桔子设计',
      description: '品牌视觉与包装设计工作室',
      url: 'https://example.com',
      features: ['品牌视觉', '包装设计', '展陈设计', '插画定制'],
    },
    fullpage: {
      name: '星辰科技',
      description: '微小卫星制造与在轨服务',
      url: 'https://example.com',
      features: ['卫星平台', '地面站服务', '遥感数据应用'],
      contactInfo: { phone: '400-806-2200', email: 'hello@example.com' },
    },
  },
}

/* ------------------------------------------------------------------ */
/* 5. 产品展示                                                         */
/* ------------------------------------------------------------------ */

const productDisplay = {
  key: 'product_display',
  label: '产品展示',
  icon: ShoppingBag,
  iconBg: '#fce7f3',
  iconColor: '#db2777',
  ambient: 'amb-plain',
  schema: [
    text('title', '产品名称', { required: true, placeholder: '例如：AirFlow Pro 降噪耳机' }),
    text('subtitle', '副标题', { placeholder: '例如：45dB 深度降噪 · 40 小时续航' }),
    image('mainImage', '产品主图', { placeholder: '粘贴图片地址，或上传后自动填入' }),
    textarea('description', '产品描述', { rows: 3 }),
    price('price', '售价', { placeholder: '899' }),
    price('originalPrice', '原价', { placeholder: '1099（选填，会显示删除线）' }),
    tags('features', '产品特点', { placeholder: '输入后按回车添加' }),
    keyValue('specifications', '规格参数', { keyPlaceholder: '例如：续航', valuePlaceholder: '例如：40 小时' }),
    group('购买按钮', [
      text('ctaText', '按钮文案', { placeholder: '例如：立即购买' }),
      url('ctaAction', '按钮跳转地址', { placeholder: 'https://example.com/buy' }),
    ]),
  ],
  styles: [
    { value: '', label: '标准白底', preview: '干净商品卡 + 底栏', gradient: 'linear-gradient(135deg,#fbfbfa,#e7e5e4)', component: ProductDisplayPage, ambient: 'amb-plain' },
    { value: 'premium', label: '奢华暗金', preview: '暖黑 + 香槟金', gradient: 'linear-gradient(135deg,#1c1917,#b45309)', component: ProductPremium, ambient: 'amb-dark-gold' },
    { value: 'showcase', label: '画廊展示', preview: '全屏图像 + 上浮卡', gradient: 'linear-gradient(135deg,#18181b,#3f3f46)', component: ProductShowcase, ambient: 'amb-dark' },
    { value: 'comparison', label: '对比型', preview: '价格面板 + 参数表', gradient: 'linear-gradient(135deg,#f1f5f9,#cbd5e1)', component: ProductComparison, ambient: 'amb-cool' },
    { value: 'story', label: '故事叙述', preview: '暖纸 + 编辑排版', gradient: 'linear-gradient(135deg,#fffbeb,#fed7aa)', component: ProductStory, ambient: 'amb-warm' },
  ],
  samples: {
    '': {
      title: 'AirFlow Pro 降噪耳机', subtitle: '45dB 深度降噪 · 40 小时续航',
      mainImage: 'https://images.unsplash.com/photo-1505740420928-5e560c06d30e?w=600&h=600&fit=crop',
      description: '双芯混合主动降噪，通勤与飞行场景实测降噪深度 45dB。LDAC 高解析传输，续航 40 小时。',
      features: ['混合主动降噪 45dB', 'LDAC / aptX 高解析', '双设备同时连接', '6 麦克风通话降噪'],
      specifications: { 单元: '12mm 动圈', 续航: '40 小时', 重量: '48g', 充电: 'USB-C / 无线' },
      price: '899', originalPrice: '1099', ctaText: '立即购买', ctaAction: 'https://example.com/buy',
    },
    premium: {
      title: 'Orbit 机械腕表', subtitle: '瑞士机芯 · 蓝宝石镜面',
      mainImage: 'https://images.unsplash.com/photo-1523275335684-37898b6baf30?w=600&h=600&fit=crop',
      description: '316L 精钢表壳，机芯经 5 方位调校，日差 ±5 秒。',
      features: ['瑞士 ETA 机芯', '蓝宝石镜面', '鳄鱼皮表带'],
      specifications: { 机芯: 'ETA 2824', 镜面: '蓝宝石', 防水: '100 米' },
      price: '8800', originalPrice: '9800', ctaText: '预约鉴赏', ctaAction: 'https://example.com',
    },
  },
}

/* ------------------------------------------------------------------ */
/* 6. 菜谱                                                             */
/* ------------------------------------------------------------------ */

const recipe = {
  key: 'recipe',
  label: '菜谱',
  icon: UtensilsCrossed,
  iconBg: '#fff7ed',
  iconColor: '#c2410c',
  ambient: 'amb-warm',
  schema: [
    text('shopName', '店名', { required: true, placeholder: '例如：巷口面馆' }),
    text('tagline', '一句话介绍', { placeholder: '例如：苏式汤面 · 现炒浇头' }),
    text('hours', '营业时间', { placeholder: '例如：06:30 - 20:00' }),
    text('address', '地址', { placeholder: '例如：苏州市姑苏区平江路 112 号' }),
    tel('phone', '订餐电话'),
    textarea('notice', '温馨提示', { rows: 2, placeholder: '例如：浇头现点现炒，出餐约需 8 分钟' }),
    list('categories', '菜单分类', {
      itemLabel: '分类',
      addLabel: '添加分类',
      defaultItem: { name: '', items: [] },
      fields: [
        text('name', '分类名称', { required: true, placeholder: '例如：汤面 / 小菜 / 酒水' }),
        list('items', '菜品', {
          itemLabel: '菜品',
          addLabel: '添加菜品',
          defaultItem: { name: '', price: '', desc: '', signature: false, vegetarian: false, spicy: 0 },
          fields: recipeItemFields,
        }),
      ],
    }),
  ],
  styles: [
    { value: '', label: '街边小馆', preview: '红纸价目表 + 招牌标记', gradient: 'linear-gradient(135deg,#fbf7ee,#b91c1c)', component: RecipePage, ambient: 'amb-warm' },
    { value: 'chalk', label: '黑板粉笔', preview: '深绿黑板 + 粉笔字', gradient: 'linear-gradient(135deg,#16201c,#4a5a50)', component: RecipeChalk, ambient: 'amb-dark' },
    { value: 'night_market', label: '夜市烟火', preview: '黑底霓虹 + 辣度', gradient: 'linear-gradient(135deg,#0b0908,#f97316)', component: RecipeNightMarket, ambient: 'amb-dark' },
    { value: 'breakfast', label: '清晨早点', preview: '米白清爽 + 分类块', gradient: 'linear-gradient(135deg,#fdfbf5,#eab308)', component: RecipeBreakfast, ambient: 'amb-warm' },
    { value: 'fine_dining', label: '精致餐厅', preview: '衬线排版 + 细线', gradient: 'linear-gradient(135deg,#f6f4ef,#787162)', component: RecipeFineDining, ambient: 'amb-plain' },
  ],
  samples: {
    '': {
      shopName: '巷口面馆', tagline: '苏式汤面 · 现炒浇头 · 每天早上六点半开门',
      hours: '06:30 - 20:00', address: '苏州市姑苏区平江路 112 号', phone: '0512-6552 8890',
      notice: '浇头现点现炒，出餐约需 8 分钟。堂食免茶位费。',
      categories: [
        {
          name: '汤面', items: [
            { name: '焖肉面', price: '28', desc: '五花肉焖足三小时，酥而不散', signature: true },
            { name: '爆鱼面', price: '26', desc: '青鱼现炸，外脆里嫩', spicy: 0 },
            { name: '虾仁两面黄', price: '42', desc: '两面煎脆，浇上河虾仁', signature: true },
          ],
        },
        {
          name: '小笼与点心', items: [
            { name: '蟹粉小笼', price: '36', desc: '现拆蟹粉，一笼六只' },
            { name: '鲜肉小笼', price: '22', desc: '皮薄汁多，趁热吃' },
          ],
        },
        {
          name: '小菜', items: [
            { name: '油氽花生', price: '8', vegetarian: true },
            { name: '腌笃鲜', price: '18', desc: '春笋咸肉慢炖' },
          ],
        },
      ],
    },
    chalk: {
      shopName: '老陈家常菜', tagline: '今天买到什么就做什么，黑板上的都是现做的',
      hours: '11:00 - 14:00 · 17:00 - 21:00', address: '杭州市拱墅区小河直街 28 号', phone: '0571-8823 6612',
      categories: [
        {
          name: '今日热炒', items: [
            { name: '青椒炒肉丝', price: '32', signature: true, spicy: 1 },
            { name: '干锅花菜', price: '28', vegetarian: true, spicy: 2 },
            { name: '番茄炒蛋', price: '24', vegetarian: true },
          ],
        },
        {
          name: '砂锅', items: [
            { name: '砂锅鱼头', price: '68', desc: '千岛湖胖头鱼，需等 20 分钟' },
            { name: '腌菜豆腐煲', price: '26', vegetarian: true },
          ],
        },
        { name: '主食', items: [{ name: '猪油拌饭', price: '12' }, { name: '手擀面', price: '16' }] },
      ],
    },
    night_market: {
      shopName: '老周烧烤', tagline: '炭火现烤 · 啤酒冰镇 · 营业到凌晨两点',
      hours: '18:00 - 02:00', address: '成都市武侯区玉林西路 46 号', phone: '138 8090 5566',
      notice: '所有烤串现点现烤，辣度可选微辣 / 中辣 / 特辣。',
      categories: [
        {
          name: '烤串', items: [
            { name: '羊肉串', price: '6', desc: '内蒙羔羊，肥瘦相间', signature: true, spicy: 2 },
            { name: '烤五花肉', price: '5', signature: true, spicy: 1 },
            { name: '烤韭菜', price: '4', vegetarian: true, spicy: 1 },
            { name: '烤茄子', price: '12', vegetarian: true, spicy: 2 },
          ],
        },
        {
          name: '招牌', items: [
            { name: '麻辣小龙虾', price: '88', desc: '一斤半，蒜蓉或十三香', signature: true, spicy: 3 },
            { name: '烤脑花', price: '22', spicy: 3 },
          ],
        },
        { name: '酒水', items: [{ name: '冰镇扎啤', price: '15' }, { name: '酸梅汤', price: '8' }] },
      ],
    },
    breakfast: {
      shopName: '阿婆早点', tagline: '现磨豆浆 · 手工包子 · 六点出摊，卖完收摊',
      hours: '06:00 - 10:30', address: '南京市秦淮区三条巷 7 号', phone: '025-8446 1120',
      notice: '包子每日限量，售完即止。可提前电话预留。',
      categories: [
        {
          name: '包子', items: [
            { name: '鲜肉大包', price: '4', signature: true },
            { name: '青菜香菇包', price: '3.5', vegetarian: true },
            { name: '豆沙包', price: '3', vegetarian: true },
          ],
        },
        {
          name: '粥与豆浆', items: [
            { name: '现磨豆浆', price: '3', vegetarian: true, desc: '可加糖 / 无糖' },
            { name: '皮蛋瘦肉粥', price: '7' },
            { name: '小米粥', price: '4', vegetarian: true },
          ],
        },
        {
          name: '面点', items: [
            { name: '葱油饼', price: '5', vegetarian: true, signature: true },
            { name: '茶叶蛋', price: '2' },
          ],
        },
      ],
    },
    fine_dining: {
      shopName: '山月', tagline: '一间只有八张桌子的小馆，按季节换菜单。',
      hours: '周二至周日 17:30 - 22:00', address: '上海市徐汇区安福路 201 号 2 层', phone: '021-5466 8890',
      notice: '本店需提前一日预订，菜单随当季食材调整。',
      categories: [
        {
          name: '前菜', items: [
            { name: '糟卤鲜鲍', price: '68', desc: '三年花雕糟卤，冰镇四小时' },
            { name: '春笋拌马兰头', price: '42', vegetarian: true, desc: '当季春笋，芝麻油轻拌' },
          ],
        },
        {
          name: '主菜', items: [
            { name: '清蒸太湖白鱼', price: '168', desc: '仅用盐与姜，蒸八分钟', signature: true },
            { name: '红烧肉配栗子', price: '98', desc: '黑毛猪五花，慢炖两小时', signature: true },
            { name: '蟹粉豆腐', price: '88', desc: '现拆蟹粉，嫩豆腐' },
          ],
        },
        {
          name: '主食与甜品', items: [
            { name: '猪油葱油拌面', price: '38' },
            { name: '桂花酒酿圆子', price: '32', vegetarian: true },
          ],
        },
      ],
    },
  },
}

/* ------------------------------------------------------------------ */
/* 7. 自定义企业（5 个行业，字段各不相同）                              */
/* ------------------------------------------------------------------ */

const custom = {
  key: 'custom',
  label: '自定义企业',
  icon: Palette,
  iconBg: '#f0fdf4',
  iconColor: '#16a34a',
  ambient: 'amb-plain',
  // 各行业字段不同，因此 schema 挂在每个 style 上
  schema: [],
  styles: [
    {
      value: 'custom_company', label: '企业服务', preview: '企业主页 + 服务网格', gradient: 'linear-gradient(135deg,#eff6ff,#bfdbfe)', component: CustomCompanyPage, ambient: 'amb-plain',
      schema: [
        text('companyName', '公司名称', { required: true, placeholder: '例如：拾光花艺' }),
        tel('phone', '联系电话'),
        textarea('description', '业务介绍', { rows: 3 }),
        tags('images', '相册图片地址', { placeholder: '粘贴图片 URL 后按回车' }),
        list('services', '服务项目', {
          itemLabel: '服务', addLabel: '添加服务',
          defaultItem: { icon: 'sparkles', name: '' },
          fields: [
            text('name', '服务名称', { required: true, placeholder: '例如：婚礼花艺' }),
            select('icon', '图标', [
              { value: 'sparkles', label: '闪光' }, { value: 'cake', label: '蛋糕' }, { value: 'gift', label: '礼物' },
              { value: 'truck', label: '配送' }, { value: 'heart', label: '爱心' }, { value: 'camera', label: '摄影' },
              { value: 'scissors', label: '剪刀' }, { value: 'wrench', label: '维修' }, { value: 'leaf', label: '植物' },
              { value: 'coffee', label: '咖啡' }, { value: 'star', label: '星标' }, { value: 'users', label: '团队' },
            ]),
          ],
        }),
      ],
    },
    {
      value: 'restaurant', label: '餐饮美食', preview: '陶土菜单 + 营业牌', gradient: 'linear-gradient(135deg,#fdf6ec,#c2410c)', component: CustomRestaurant, ambient: 'amb-warm',
      schema: [
        text('shopName', '店名', { required: true, placeholder: '例如：巷口面馆' }),
        text('cuisine', '菜系', { placeholder: '例如：苏式面 / 本帮菜' }),
        text('hours', '营业时间', { placeholder: '例如：06:30 - 20:00' }),
        text('address', '地址'),
        tel('phone', '订餐电话'),
        list('menu', '菜单', {
          itemLabel: '菜品', addLabel: '添加菜品',
          defaultItem: { name: '', price: '', image: '' },
          fields: [
            text('name', '菜名', { required: true }),
            price('price', '价格'),
            image('image', '菜品图片（选填）'),
          ],
        }),
      ],
    },
    {
      value: 'salon', label: '美业沙龙', preview: '价目册 + 作品集', gradient: 'linear-gradient(135deg,#fdf2f8,#be185d)', component: CustomSalon, ambient: 'amb-warm-pink',
      schema: [
        text('salonName', '店名', { required: true, placeholder: '例如：拾壹造型' }),
        tel('phone', '预约电话'),
        text('address', '地址'),
        tags('gallery', '作品图片地址', { placeholder: '粘贴图片 URL 后按回车' }),
        list('services', '服务价目', {
          itemLabel: '服务', addLabel: '添加服务',
          defaultItem: { name: '', price: '', duration: '' },
          fields: [
            text('name', '服务名称', { required: true, placeholder: '例如：设计师精剪' }),
            price('price', '价格'),
            text('duration', '时长', { placeholder: '例如：60min' }),
          ],
        }),
      ],
    },
    {
      value: 'fitness', label: '健身运动', preview: '课表时间轴 + 设施', gradient: 'linear-gradient(135deg,#0b0f0c,#a3e635)', component: CustomFitness, ambient: 'amb-dark-green',
      schema: [
        text('gymName', '场馆名称', { required: true, placeholder: '例如：山丘健身工作室' }),
        tel('phone', '咨询电话'),
        text('address', '地址'),
        tags('facilities', '设施', { placeholder: '输入后按回车添加，例如：自由力量区' }),
        list('classes', '课程表', {
          itemLabel: '课程', addLabel: '添加课程',
          defaultItem: { name: '', time: '', coach: '' },
          fields: [
            text('name', '课程名称', { required: true, placeholder: '例如：燃脂搏击' }),
            text('time', '时间', { placeholder: '例如：周一/三/五 19:30' }),
            text('coach', '教练', { placeholder: '例如：阿哲教练' }),
          ],
        }),
      ],
    },
    {
      value: 'education', label: '教育培训', preview: '课程卡 + 师资名录', gradient: 'linear-gradient(135deg,#0f172a,#b45309)', component: CustomEducation, ambient: 'amb-cool',
      schema: [
        text('schoolName', '机构名称', { required: true, placeholder: '例如：木棉书院' }),
        tel('phone', '咨询电话'),
        text('address', '地址'),
        list('courses', '课程', {
          itemLabel: '课程', addLabel: '添加课程',
          defaultItem: { name: '', desc: '', price: '' },
          fields: [
            text('name', '课程名称', { required: true, placeholder: '例如：儿童创意写作' }),
            text('desc', '课程说明', { placeholder: '例如：8-12 岁 · 12 周课程' }),
            price('price', '价格'),
          ],
        }),
        list('teachers', '师资', {
          itemLabel: '教师', addLabel: '添加教师',
          defaultItem: { name: '', title: '', avatar: '' },
          fields: [
            text('name', '姓名', { required: true }),
            text('title', '头衔', { placeholder: '例如：课程主理人' }),
            image('avatar', '头像图片地址（选填）'),
          ],
        }),
      ],
    },
  ],
  samples: {
    custom_company: {
      companyName: '拾光花艺', phone: '138 0013 8000', email: 'hello@shiguang.studio',
      description: '社区花店 · 每周三束鲜花订阅 · 婚礼与活动花艺',
      images: [], services: [{ icon: 'sparkles', name: '婚礼花艺' }, { icon: 'cake', name: '生日花束' }, { icon: 'truck', name: '同城配送' }],
    },
    restaurant: {
      shopName: '巷口面馆', cuisine: '苏式面 / 本帮菜', hours: '06:30 - 20:00',
      address: '苏州市姑苏区平江路 112 号', phone: '0512-6552 8890',
      menu: [{ name: '焖肉面', price: '28', image: '' }, { name: '虾仁两面黄', price: '42', image: '' }, { name: '蟹粉小笼', price: '36', image: '' }],
    },
    salon: {
      salonName: '拾壹造型',
      services: [{ name: '设计师精剪', price: '168', duration: '60min' }, { name: '法式烫', price: '588', duration: '150min' }, { name: '头皮护理', price: '288', duration: '45min' }],
      gallery: [], phone: '138 1880 2211', address: '上海市静安区愚园路 68 号',
    },
    fitness: {
      gymName: '山丘健身工作室',
      classes: [{ name: '燃脂搏击', time: '周一/三/五 19:30', coach: '阿哲教练' }, { name: '普拉提', time: '周二/四 10:00', coach: '小鹿教练' }],
      facilities: ['自由力量区', '普拉提房', '淋浴间', '体测室'], phone: '138 2244 6677', address: '杭州市西湖区文二路 218 号',
    },
    education: {
      schoolName: '木棉书院',
      courses: [{ name: '儿童创意写作', desc: '8-12 岁 · 12 周课程 · 小班 8 人', price: '2680' }, { name: '少儿书法', desc: '6-10 岁 · 每周六上午', price: '1980' }],
      teachers: [{ name: '周老师', title: '课程主理人', avatar: '' }, { name: '陆老师', title: '书法教师', avatar: '' }],
      phone: '400-820-3311', address: '成都市锦江区红星路 99 号',
    },
  },
}

/* ------------------------------------------------------------------ */
/* 导出                                                                */
/* ------------------------------------------------------------------ */

/* ------------------------------------------------------------------ */
/* 8. 夜店交际                                                         */
/* ------------------------------------------------------------------ */

const nightclub = {
  key: 'nightclub',
  label: '夜店交际',
  icon: Martini,
  iconBg: '#fae8ff',
  iconColor: '#a21caf',
  ambient: 'amb-dark',
  schema: [
    text('name', '姓名 / 艺名', { required: true, placeholder: '例如：Nova' }),
    text('alias', '别名 / 头衔', { placeholder: '例如：NOVA' }),
    text('role', '身份', { placeholder: '例如：驻场 DJ' }),
    text('venue', '常驻场地', { placeholder: '例如：NEON ROOM' }),
    text('city', '城市', { placeholder: '例如：上海' }),
    text('wechat', '微信', { placeholder: '例如：nova_on_tour' }),
    tel('phone', '电话', { placeholder: '138 0000 0000' }),
  ],
  styles: [...nightclubStyles],
  samples: {
    '': {
      name: 'Nova',
      alias: 'NOVA',
      role: '驻场 DJ',
      venue: 'NEON ROOM',
      city: '上海',
      wechat: 'nova_on_tour',
      phone: '138 0013 8000',
    },
  },
}

/* ------------------------------------------------------------------ */
/* 9. 二次元闪卡                                                       */
/* ------------------------------------------------------------------ */

const animeCard = {
  key: 'anime_card',
  label: '二次元闪卡',
  icon: Sparkles,
  iconBg: '#fce7f3',
  iconColor: '#db2777',
  ambient: 'amb-dark-green',
  schema: [
    text('name', '角色名', { required: true, placeholder: '例如：星野澪' }),
    text('title', '称号', { placeholder: '例如：月下观测者' }),
    text('series', '出自作品', { placeholder: '例如：星海纪元' }),
    text('element', '属性', { placeholder: '例如：水 · 辅助' }),
    textarea('quote', '台词', { rows: 2, placeholder: '例如：……你也是来看星星的吗？' }),
    text('social', '社交账号', { placeholder: '例如：mio_star' }),
  ],
  styles: [...animeCardStyles],
  samples: {
    '': {
      name: '星野澪',
      title: '月下观测者',
      series: '星海纪元',
      element: '水 · 辅助',
      quote: '……你也是来看星星的吗？',
      social: 'mio_star',
    },
  },
}

/* ------------------------------------------------------------------ */
/* 10. 宠物铭牌                                                        */
/* ------------------------------------------------------------------ */

const petTag = {
  key: 'pet_tag',
  label: '宠物铭牌',
  icon: PawPrint,
  iconBg: '#ecfccb',
  iconColor: '#65a30d',
  ambient: 'amb-dark-gold',
  schema: [
    text('petName', '宠物名', { required: true, placeholder: '例如：豆包' }),
    text('breed', '品种', { placeholder: '例如：柯基' }),
    text('age', '年龄', { placeholder: '例如：3 岁' }),
    text('ownerName', '主人称呼', { placeholder: '例如：李女士' }),
    tel('phone', '联系电话', { required: true, placeholder: '138 0000 0000' }),
    text('address', '住址', { placeholder: '例如：青禾社区 12 栋' }),
    text('medical', '过敏 / 用药', { placeholder: '例如：对鸡肉过敏' }),
    textarea('note', '提示', { rows: 2, placeholder: '例如：怕生，请勿强行抱抱' }),
  ],
  styles: [...petTagStyles],
  samples: {
    '': {
      petName: '豆包',
      breed: '柯基',
      age: '3 岁',
      ownerName: '李女士',
      phone: '138 0013 8000',
      address: '青禾社区 12 栋',
      medical: '对鸡肉过敏',
      note: '怕生，请勿强行抱抱',
    },
  },
}

/* ------------------------------------------------------------------ */
/* 11. 奇怪的想法（横屏猎奇）                                          */
/* ------------------------------------------------------------------ */

/** 速度 / 强度用 select 而非滑块，保证表单控件与其它类型一致 */
const SPEED_OPTIONS = [
  { value: '1', label: '很慢' },
  { value: '2', label: '慢' },
  { value: '3', label: '正常' },
  { value: '4', label: '快' },
  { value: '5', label: '很快' },
]

const weird = {
  key: 'weird',
  label: '奇怪的想法',
  icon: Eye,
  iconBg: '#fae8ff',
  iconColor: '#c026d3',
  ambient: 'amb-dark',
  // 横屏版面：展示页要换全宽容器，不能用竖屏手机宽度
  layout: 'landscape',
  // 暂时隐藏：动画精细度未达标，先从所有选型入口移除；
  // 已创建的标签仍然可以正常渲染（componentFor 不受 hidden 影响）
  hidden: true,
  schema: [
    text('title', '主文案', { placeholder: '例如：别盯着看' }),
    text('subtitle', '副文案', { placeholder: '例如：TAP IT ALIVE' }),
    select('speed', '动效速度', SPEED_OPTIONS, { default: '3' }),
    select('intensity', '视觉强度', SPEED_OPTIONS, { default: '3' }),
    text('accent', '主题色（#hex，可留空）', { placeholder: '例如：#ef4444' }),
  ],
  /** 预留：30 款猎奇风格由扩展模块注入 */
  styles: [],
  samples: {
    '': {
      title: '别盯着看',
      subtitle: 'TAP IT ALIVE',
      speed: '3',
      intensity: '3',
      accent: '',
    },
  },
}

/* —— 把扩展风格注入各类型（追加在原有 5 款之后） —— */
businessCard.styles.push(...businessCardExtra)
childInfo.styles.push(...childInfoExtra)
emergencyContact.styles.push(...emergencyContactExtra)
companyWebsite.styles.push(...companyWebsiteExtra)
productDisplay.styles.push(...productDisplayExtra)
recipe.styles.push(...recipeExtra)
custom.styles.push(...customExtra)
weird.styles.push(...weirdStyles)

export const templateTypes = {
  business_card: businessCard,
  child_info: childInfo,
  emergency_contact: emergencyContact,
  company_website: companyWebsite,
  product_display: productDisplay,
  recipe,
  custom,
  nightclub,
  anime_card: animeCard,
  pet_tag: petTag,
  weird,
}

/** 类型顺序（表单里按这个顺序展示） */
export const templateOrder = [
  'business_card', 'child_info', 'emergency_contact',
  'company_website', 'product_display', 'recipe', 'custom',
  'nightclub', 'anime_card', 'pet_tag', 'weird',
]

/** 对外可见的类型顺序（过滤掉 hidden），选型入口一律用它 */
export const visibleTypeOrder = templateOrder.filter((key) => !templateTypes[key]?.hidden)

/** 该类型是否对用户隐藏 */
export function isHidden(typeKey) {
  return templateTypes[typeKey]?.hidden === true
}

/** 取某类型下某风格对应的 schema（风格可覆盖类型级 schema） */
export function schemaFor(typeKey, styleValue) {
  const type = templateTypes[typeKey]
  if (!type) return []
  const style = type.styles.find((s) => s.value === (styleValue || ''))
  if (style && Array.isArray(style.schema) && style.schema.length > 0) return style.schema
  return type.schema || []
}

/** 取某类型下某风格的示例数据 */
export function sampleFor(typeKey, styleValue) {
  const type = templateTypes[typeKey]
  if (!type) return {}
  const samples = type.samples || {}
  return samples[styleValue || ''] || samples[''] || Object.values(samples)[0] || {}
}

/** 取某类型下某风格对应的页面组件 */
export function componentFor(typeKey, styleValue) {
  const type = templateTypes[typeKey]
  if (!type) return null
  const style = type.styles.find((s) => s.value === (styleValue || ''))
  if (style) return style.component
  return type.styles[0]?.component || null
}

/** 取某类型下某风格的氛围背景类 */
/**
 * 版式： portrait（默认，手机竖屏） / landscape（横屏沉浸式）
 * 展示页据此切换容器宽度，横屏类型不能用 max-w-md。
 */
export function layoutFor(typeKey) {
  const type = templateTypes[typeKey]
  return type?.layout || 'portrait'
}

export function ambientFor(typeKey, styleValue) {
  const type = templateTypes[typeKey]
  if (!type) return 'amb-plain'
  const style = type.styles.find((s) => s.value === (styleValue || ''))
  return style?.ambient || type.ambient || 'amb-plain'
}
