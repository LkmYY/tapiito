import SignalFlare from '../../pages/templates/ext/emergency_contact/SignalFlare.vue'
import ClinicCalm from '../../pages/templates/ext/emergency_contact/ClinicCalm.vue'
import PaperEmergency from '../../pages/templates/ext/emergency_contact/PaperEmergency.vue'
import GlassGuard from '../../pages/templates/ext/emergency_contact/GlassGuard.vue'
import TerminalRescue from '../../pages/templates/ext/emergency_contact/TerminalRescue.vue'
import RisoAlert from '../../pages/templates/ext/emergency_contact/RisoAlert.vue'
import GoldShield from '../../pages/templates/ext/emergency_contact/GoldShield.vue'
import CollageKindred from '../../pages/templates/ext/emergency_contact/CollageKindred.vue'
import AuroraCall from '../../pages/templates/ext/emergency_contact/AuroraCall.vue'
import PixelGuard from '../../pages/templates/ext/emergency_contact/PixelGuard.vue'
import SumiContact from '../../pages/templates/ext/emergency_contact/SumiContact.vue'
import FilmAlert from '../../pages/templates/ext/emergency_contact/FilmAlert.vue'
import BlueprintSafe from '../../pages/templates/ext/emergency_contact/BlueprintSafe.vue'
import SwissSignal from '../../pages/templates/ext/emergency_contact/SwissSignal.vue'
import BrutalCall from '../../pages/templates/ext/emergency_contact/BrutalCall.vue'
import DecoAssist from '../../pages/templates/ext/emergency_contact/DecoAssist.vue'
import TriageCard from '../../pages/templates/ext/emergency_contact/TriageCard.vue'
import TransitSOS from '../../pages/templates/ext/emergency_contact/TransitSOS.vue'
import VhsEmergency from '../../pages/templates/ext/emergency_contact/VhsEmergency.vue'
import HerbHaven from '../../pages/templates/ext/emergency_contact/HerbHaven.vue'
import PressBulletin from '../../pages/templates/ext/emergency_contact/PressBulletin.vue'
import PassportHelp from '../../pages/templates/ext/emergency_contact/PassportHelp.vue'
import PolaroidFamily from '../../pages/templates/ext/emergency_contact/PolaroidFamily.vue'
import OpsConsole from '../../pages/templates/ext/emergency_contact/OpsConsole.vue'
import EquipmentTag from '../../pages/templates/ext/emergency_contact/EquipmentTag.vue'

export default [
  {
    value: 'signal_flare',
    label: '信号烽火',
    preview: '暗红信标 + 超大号码',
    gradient: 'linear-gradient(135deg,#1a0505,#dc2626)',
    component: SignalFlare,
    ambient: 'amb-dark-red',
  },
  {
    value: 'clinic_calm',
    label: '冷静诊室',
    preview: '青绿分栏与病历清单',
    gradient: 'linear-gradient(135deg,#f0fdfa,#0f766e)',
    component: ClinicCalm,
    ambient: 'amb-cool',
  },
  {
    value: 'paper_emergency',
    label: '纸本应急',
    preview: '米纸编辑与编号联络',
    gradient: 'linear-gradient(135deg,#f4efe3,#9f2d20)',
    component: PaperEmergency,
    ambient: 'amb-warm',
  },
  {
    value: 'glass_guard',
    label: '玻璃守护',
    preview: '冷夜玻璃与极光卡片',
    gradient: 'linear-gradient(135deg,#081120,#7c3aed)',
    component: GlassGuard,
    ambient: 'amb-cool',
  },
  {
    value: 'terminal_rescue',
    label: '终端救援',
    preview: '黑绿终端与命令列表',
    gradient: 'linear-gradient(135deg,#03110a,#4ade80)',
    component: TerminalRescue,
    ambient: 'amb-dark-green',
  },
  {
    value: 'riso_alert',
    label: '复写警报',
    preview: '双色印刷与撕票券',
    gradient: 'linear-gradient(135deg,#f4ead5,#e2452f)',
    component: RisoAlert,
    ambient: 'amb-warm',
  },
  {
    value: 'gold_shield',
    label: '金盾守护',
    preview: '黑金盾徽与衬线数字',
    gradient: 'linear-gradient(135deg,#0f0c08,#d4af37)',
    component: GoldShield,
    ambient: 'amb-dark-gold',
  },
  {
    value: 'collage_kindred',
    label: '拼贴守护',
    preview: '胶带拼贴与手写温度',
    gradient: 'linear-gradient(135deg,#f6e7dd,#c97b63)',
    component: CollageKindred,
    ambient: 'amb-warm-pink',
  },
  {
    value: 'aurora_call',
    label: '极光呼叫',
    preview: '流体极光与发光拨号',
    gradient: 'linear-gradient(135deg,#170b2e,#0891b2)',
    component: AuroraCall,
    ambient: 'amb-dark',
  },
  {
    value: 'pixel_guard',
    label: '像素守卫',
    preview: '八位像素与游戏界面',
    gradient: 'linear-gradient(135deg,#0b0b13,#a3e635)',
    component: PixelGuard,
    ambient: 'amb-dark-green',
  },
  {
    value: 'sumi_contact',
    label: '墨线联络',
    preview: '宣纸留白与朱砂印章',
    gradient: 'linear-gradient(135deg,#f3efe6,#b23a2b)',
    component: SumiContact,
    ambient: 'amb-plain',
  },
  {
    value: 'film_alert',
    label: '胶片警报',
    preview: '胶片齿孔与琥珀画面',
    gradient: 'linear-gradient(135deg,#111111,#f5b544)',
    component: FilmAlert,
    ambient: 'amb-dark-gold',
  },
  {
    value: 'blueprint_safe',
    label: '蓝图预案',
    preview: '蓝晒网格与技术标注',
    gradient: 'linear-gradient(135deg,#072a4a,#7dd3fc)',
    component: BlueprintSafe,
    ambient: 'amb-cool',
  },
  {
    value: 'swiss_signal',
    label: '瑞士信号',
    preview: '红黑网格与巨型号码',
    gradient: 'linear-gradient(135deg,#ffffff,#e30613)',
    component: SwissSignal,
    ambient: 'amb-plain',
  },
  {
    value: 'brutal_call',
    label: '粗野呼叫',
    preview: '黄黑粗框与硬影贴纸',
    gradient: 'linear-gradient(135deg,#f5d90a,#111111)',
    component: BrutalCall,
    ambient: 'amb-warm',
  },
  {
    value: 'deco_assist',
    label: '装饰援助',
    preview: '祖母绿与金色放射',
    gradient: 'linear-gradient(135deg,#06231e,#c9a24b)',
    component: DecoAssist,
    ambient: 'amb-dark-gold',
  },
  {
    value: 'triage_card',
    label: '分诊卡',
    preview: '三色分诊与打孔标签',
    gradient: 'linear-gradient(135deg,#f7f4ee,#c0392b)',
    component: TriageCard,
    ambient: 'amb-warm',
  },
  {
    value: 'transit_sos',
    label: '站点 SOS',
    preview: '站牌路线与安全黄',
    gradient: 'linear-gradient(135deg,#0c1116,#ffd400)',
    component: TransitSOS,
    ambient: 'amb-dark-gold',
  },
  {
    value: 'vhs_emergency',
    label: '录像带',
    preview: '扫描线与故障粉青',
    gradient: 'linear-gradient(135deg,#140b1c,#ff4fa3)',
    component: VhsEmergency,
    ambient: 'amb-dark',
  },
  {
    value: 'herb_haven',
    label: '草木守护',
    preview: '草本药房与森绿纸感',
    gradient: 'linear-gradient(135deg,#eef3e6,#2f4a35)',
    component: HerbHaven,
    ambient: 'amb-warm-green',
  },
  {
    value: 'press_bulletin',
    label: '新闻快报',
    preview: '号外报纸与栏线标题',
    gradient: 'linear-gradient(135deg,#f2ede2,#1a1a1a)',
    component: PressBulletin,
    ambient: 'amb-warm',
  },
  {
    value: 'passport_help',
    label: '护照援助',
    preview: '护照内页与钢印数据',
    gradient: 'linear-gradient(135deg,#12233f,#c9a24b)',
    component: PassportHelp,
    ambient: 'amb-dark-gold',
  },
  {
    value: 'polaroid_family',
    label: '拍立得',
    preview: '家庭照片墙与暖色贴纸',
    gradient: 'linear-gradient(135deg,#fff1e6,#f59e8b)',
    component: PolaroidFamily,
    ambient: 'amb-warm-pink',
  },
  {
    value: 'ops_console',
    label: '运维值班',
    preview: '深色看板与值班状态',
    gradient: 'linear-gradient(135deg,#0b0f14,#38bdf8)',
    component: OpsConsole,
    ambient: 'amb-cool',
  },
  {
    value: 'equipment_tag',
    label: '设备铭牌',
    preview: '金属铭牌与铆钉标签',
    gradient: 'linear-gradient(135deg,#d9dde1,#f26b21)',
    component: EquipmentTag,
    ambient: 'amb-warm',
  },
]
