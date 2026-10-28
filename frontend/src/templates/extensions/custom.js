import BlueprintDraft from '../../pages/templates/ext/custom/BlueprintDraft.vue'
import BoutiqueStay from '../../pages/templates/ext/custom/BoutiqueStay.vue'
import BrutalSlab from '../../pages/templates/ext/custom/BrutalSlab.vue'
import ChamberClassic from '../../pages/templates/ext/custom/ChamberClassic.vue'
import ClayCeramic from '../../pages/templates/ext/custom/ClayCeramic.vue'
import ConcreteModel from '../../pages/templates/ext/custom/ConcreteModel.vue'
import CounselFolio from '../../pages/templates/ext/custom/CounselFolio.vue'
import DailyNews from '../../pages/templates/ext/custom/DailyNews.vue'
import EditorialAtelier from '../../pages/templates/ext/custom/EditorialAtelier.vue'
import FilmContact from '../../pages/templates/ext/custom/FilmContact.vue'
import FloristPress from '../../pages/templates/ext/custom/FloristPress.vue'
import GildedMonogram from '../../pages/templates/ext/custom/GildedMonogram.vue'
import GlassConsult from '../../pages/templates/ext/custom/GlassConsult.vue'
import LaunchOrbit from '../../pages/templates/ext/custom/LaunchOrbit.vue'
import ManifestoPop from '../../pages/templates/ext/custom/ManifestoPop.vue'
import PostcardSoft from '../../pages/templates/ext/custom/PostcardSoft.vue'
import RisoPress from '../../pages/templates/ext/custom/RisoPress.vue'
import RoasteryWarm from '../../pages/templates/ext/custom/RoasteryWarm.vue'
import SwissGrid from '../../pages/templates/ext/custom/SwissGrid.vue'
import TerminalConsole from '../../pages/templates/ext/custom/TerminalConsole.vue'
import VinylSession from '../../pages/templates/ext/custom/VinylSession.vue'
import WhiteCube from '../../pages/templates/ext/custom/WhiteCube.vue'
import ZenVoid from '../../pages/templates/ext/custom/ZenVoid.vue'
import NeonStudio from '../../pages/templates/ext/custom/NeonStudio.vue'
import PaperCut from '../../pages/templates/ext/custom/PaperCut.vue'
import { text, textarea, tel, email, tags, list, select } from '../schema.js'

/** 扩展风格共用的企业档案字段 */
const companySchema = [
  text('companyName', '公司名称', { required: true, placeholder: '例如：拾光花艺' }),
  tel('phone', '联系电话'),
  email('email', '邮箱'),
  textarea('description', '业务介绍', { rows: 3 }),
  tags('images', '相册图片地址', { placeholder: '粘贴图片 URL 后按回车' }),
  list('services', '服务项目', {
    itemLabel: '服务',
    addLabel: '添加服务',
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
]

export default [
  {
    value: 'blueprint_draft',
    label: '蓝图草案',
    preview: '深蓝工程蓝图标注',
    gradient: 'linear-gradient(135deg,#0a2a46,#38bdf8)',
    component: BlueprintDraft,
    ambient: 'amb-cool',
    schema: companySchema,
  },
  {
    value: 'boutique_stay',
    label: '精品民宿',
    preview: '奶油房卡住宿卡',
    gradient: 'linear-gradient(135deg,#e9dfcc,#7d8a6d)',
    component: BoutiqueStay,
    ambient: 'amb-warm',
    schema: companySchema,
  },
  {
    value: 'brutal_slab',
    label: '粗野宣言',
    preview: '荧光黄大字海报',
    gradient: 'linear-gradient(135deg,#e9ff3f,#18181b)',
    component: BrutalSlab,
    ambient: 'amb-warm',
    schema: companySchema,
  },
  {
    value: 'chamber_classic',
    label: '古典商事',
    preview: '墨绿金箔顾问函',
    gradient: 'linear-gradient(135deg,#0d2b25,#c6a14e)',
    component: ChamberClassic,
    ambient: 'amb-dark-green',
    schema: companySchema,
  },
  {
    value: 'clay_ceramic',
    label: '陶土手作',
    preview: '米色陶土工作室',
    gradient: 'linear-gradient(135deg,#efe7d8,#b45309)',
    component: ClayCeramic,
    ambient: 'amb-warm',
    schema: companySchema,
  },
  {
    value: 'concrete_model',
    label: '混凝土模型',
    preview: '灰米建筑模型册',
    gradient: 'linear-gradient(135deg,#d9d4c9,#57534e)',
    component: ConcreteModel,
    ambient: 'amb-plain',
    schema: companySchema,
  },
  {
    value: 'counsel_folio',
    label: '顾问备忘',
    preview: '米白机密备忘',
    gradient: 'linear-gradient(135deg,#f4efe2,#a16207)',
    component: CounselFolio,
    ambient: 'amb-warm',
    schema: companySchema,
  },
  {
    value: 'daily_news',
    label: '今日企业刊',
    preview: '报纸头版排版',
    gradient: 'linear-gradient(135deg,#eee9dc,#1c1917)',
    component: DailyNews,
    ambient: 'amb-warm',
    schema: companySchema,
  },
  {
    value: 'editorial_atelier',
    label: '编辑档案',
    preview: '杂志编辑档案',
    gradient: 'linear-gradient(135deg,#f3eee3,#9f1239)',
    component: EditorialAtelier,
    ambient: 'amb-warm',
    schema: companySchema,
  },
  {
    value: 'film_contact',
    label: '胶片接触印',
    preview: '黑白胶片接触印样',
    gradient: 'linear-gradient(135deg,#080808,#a3a3a3)',
    component: FilmContact,
    ambient: 'amb-dark',
    schema: companySchema,
  },
  {
    value: 'florist_press',
    label: '植物事务所',
    preview: '森绿植物图鉴',
    gradient: 'linear-gradient(135deg,#eef1e7,#4d7c0f)',
    component: FloristPress,
    ambient: 'amb-warm-green',
    schema: companySchema,
  },
  {
    value: 'gilded_monogram',
    label: '鎏金私邸',
    preview: '黑金字母徽章',
    gradient: 'linear-gradient(135deg,#0b0a08,#d4af6a)',
    component: GildedMonogram,
    ambient: 'amb-dark-gold',
    schema: companySchema,
  },
  {
    value: 'glass_consult',
    label: '玻璃咨询',
    preview: '紫蓝玻璃顾问卡',
    gradient: 'linear-gradient(135deg,#e9e4ff,#7c3aed)',
    component: GlassConsult,
    ambient: 'amb-cool',
    schema: companySchema,
  },
  {
    value: 'launch_orbit',
    label: '发射轨道',
    preview: '深空科技产品页',
    gradient: 'linear-gradient(135deg,#060a1a,#2563eb,#22d3ee)',
    component: LaunchOrbit,
    ambient: 'amb-dark',
    schema: companySchema,
  },
  {
    value: 'manifesto_pop',
    label: '宣言波普',
    preview: '钴蓝撞色宣言',
    gradient: 'linear-gradient(135deg,#1b2df0,#facc15)',
    component: ManifestoPop,
    ambient: 'amb-cool',
    schema: companySchema,
  },
  {
    value: 'postcard_soft',
    label: '柔光明信片',
    preview: '蜜桃色问候卡',
    gradient: 'linear-gradient(135deg,#f7e7df,#fb7185)',
    component: PostcardSoft,
    ambient: 'amb-warm-pink',
    schema: companySchema,
  },
  {
    value: 'riso_press',
    label: '孔版印刷',
    preview: '双色 Riso 印刷样张',
    gradient: 'linear-gradient(135deg,#f2ead8,#ea580c,#1d4ed8)',
    component: RisoPress,
    ambient: 'amb-warm',
    schema: companySchema,
  },
  {
    value: 'roastery_warm',
    label: '烘焙暖意',
    preview: '咖啡烘焙供应单',
    gradient: 'linear-gradient(135deg,#e9dfcc,#92400e)',
    component: RoasteryWarm,
    ambient: 'amb-warm',
    schema: companySchema,
  },
  {
    value: 'swiss_grid',
    label: '瑞士网格',
    preview: '极简瑞士网格',
    gradient: 'linear-gradient(135deg,#f6f5f1,#111827)',
    component: SwissGrid,
    ambient: 'amb-plain',
    schema: companySchema,
  },
  {
    value: 'terminal_console',
    label: '终端控制台',
    preview: '绿字命令行',
    gradient: 'linear-gradient(135deg,#040705,#16a34a)',
    component: TerminalConsole,
    ambient: 'amb-dark-green',
    schema: companySchema,
  },
  {
    value: 'vinyl_session',
    label: '黑胶会话',
    preview: '米色黑胶 LP 封套',
    gradient: 'linear-gradient(135deg,#efebe1,#1c1917)',
    component: VinylSession,
    ambient: 'amb-warm',
    schema: companySchema,
  },
  {
    value: 'white_cube',
    label: '白盒子',
    preview: '美术馆墙标白盒',
    gradient: 'linear-gradient(135deg,#fbfbf9,#a8a29e)',
    component: WhiteCube,
    ambient: 'amb-plain',
    schema: companySchema,
  },
  {
    value: 'zen_void',
    label: '禅意留白',
    preview: '大字留白静谧',
    gradient: 'linear-gradient(135deg,#f7f5ef,#78716c)',
    component: ZenVoid,
    ambient: 'amb-plain',
    schema: companySchema,
  },
  {
    value: 'neon_studio',
    label: '霓虹工作室',
    preview: '暗夜霓虹工作室',
    gradient: 'linear-gradient(135deg,#08060f,#a855f7,#22d3ee)',
    component: NeonStudio,
    ambient: 'amb-dark',
    schema: companySchema,
  },
  {
    value: 'paper_cut',
    label: '纸雕拼贴',
    preview: '暖色纸雕层次',
    gradient: 'linear-gradient(135deg,#fff7ed,#f59e0b,#0f766e)',
    component: PaperCut,
    ambient: 'amb-warm',
    schema: companySchema,
  },
]
