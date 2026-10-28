import GrandmaNotebook from '../../pages/templates/ext/recipe/GrandmaNotebook.vue'
import FoodEditorial from '../../pages/templates/ext/recipe/FoodEditorial.vue'
import KonbiniLabel from '../../pages/templates/ext/recipe/KonbiniLabel.vue'
import InkBroadsheet from '../../pages/templates/ext/recipe/InkBroadsheet.vue'
import ButcherBlock from '../../pages/templates/ext/recipe/ButcherBlock.vue'
import MidnightDiner from '../../pages/templates/ext/recipe/MidnightDiner.vue'
import SwissMenu from '../../pages/templates/ext/recipe/SwissMenu.vue'
import Route66Diner from '../../pages/templates/ext/recipe/Route66Diner.vue'
import WagamiKaiseki from '../../pages/templates/ext/recipe/WagamiKaiseki.vue'
import NordicKitchen from '../../pages/templates/ext/recipe/NordicKitchen.vue'
import BakeryCraft from '../../pages/templates/ext/recipe/BakeryCraft.vue'
import HotpotNine from '../../pages/templates/ext/recipe/HotpotNine.vue'
import PourOverBar from '../../pages/templates/ext/recipe/PourOverBar.vue'
import CocktailTicket from '../../pages/templates/ext/recipe/CocktailTicket.vue'
import MacroLunchbox from '../../pages/templates/ext/recipe/MacroLunchbox.vue'
import BlueprintKitchen from '../../pages/templates/ext/recipe/BlueprintKitchen.vue'
import MorningPost from '../../pages/templates/ext/recipe/MorningPost.vue'
import MealTicket from '../../pages/templates/ext/recipe/MealTicket.vue'
import PolaroidScrapbook from '../../pages/templates/ext/recipe/PolaroidScrapbook.vue'
import GildedSupper from '../../pages/templates/ext/recipe/GildedSupper.vue'
import PixelCanteen from '../../pages/templates/ext/recipe/PixelCanteen.vue'
import InkWashMenu from '../../pages/templates/ext/recipe/InkWashMenu.vue'
import FilmContact from '../../pages/templates/ext/recipe/FilmContact.vue'
import CyberCanteen from '../../pages/templates/ext/recipe/CyberCanteen.vue'
import HerbGarden from '../../pages/templates/ext/recipe/HerbGarden.vue'

export default [
  {
    value: 'grandma_notebook',
    label: '外婆手札',
    preview: '横线本 + 胶带 + 手写价',
    gradient: 'linear-gradient(135deg,#f7efdd,#b98a4e)',
    component: GrandmaNotebook,
    ambient: 'amb-warm',
  },
  {
    value: 'food_editorial',
    label: '食帖大片',
    preview: '杂志大片 + 超大衬线标题',
    gradient: 'linear-gradient(135deg,#f2eee5,#a33d28)',
    component: FoodEditorial,
    ambient: 'amb-plain',
  },
  {
    value: 'konbini_label',
    label: '便利店签',
    preview: '荧光价签 + 条码货架',
    gradient: 'linear-gradient(135deg,#eef0ea,#1f8a4c)',
    component: KonbiniLabel,
    ambient: 'amb-warm-green',
  },
  {
    value: 'ink_broadsheet',
    label: '铅字晨报',
    preview: '黑白报纸 + 三栏密排',
    gradient: 'linear-gradient(135deg,#f4f1e8,#131211)',
    component: InkBroadsheet,
    ambient: 'amb-plain',
  },
  {
    value: 'butcher_block',
    label: '暖木砧板',
    preview: '木纹砧板 + 吊牌菜单',
    gradient: 'linear-gradient(135deg,#2b1a0d,#c89556)',
    component: ButcherBlock,
    ambient: 'amb-dark-gold',
  },
  {
    value: 'midnight_diner',
    label: '深夜食肆',
    preview: '靛蓝夜色 + 灯笼红光',
    gradient: 'linear-gradient(135deg,#0b1020,#b4472b)',
    component: MidnightDiner,
    ambient: 'amb-dark',
  },
  {
    value: 'swiss_menu',
    label: '瑞士清单',
    preview: '白底网格 + 红色序号',
    gradient: 'linear-gradient(135deg,#f7f7f4,#e5322d)',
    component: SwissMenu,
    ambient: 'amb-plain',
  },
  {
    value: 'route66_diner',
    label: '公路餐厅',
    preview: '复古红 + 棋盘格徽章',
    gradient: 'linear-gradient(135deg,#f5e7cb,#a52f22)',
    component: Route66Diner,
    ambient: 'amb-warm',
  },
  {
    value: 'wagami_kaiseki',
    label: '和纸献立',
    preview: '和纸竖排 + 朱红印章',
    gradient: 'linear-gradient(135deg,#f2ecdf,#9e3b32)',
    component: WagamiKaiseki,
    ambient: 'amb-plain',
  },
  {
    value: 'nordic_kitchen',
    label: '北岸厨房',
    preview: '鼠尾草绿 + 大量留白',
    gradient: 'linear-gradient(135deg,#eef0e6,#7c8b6f)',
    component: NordicKitchen,
    ambient: 'amb-warm-green',
  },
  {
    value: 'bakery_craft',
    label: '面包手记',
    preview: '牛皮纸 + 黄油色标签',
    gradient: 'linear-gradient(135deg,#cfa877,#f4d47a)',
    component: BakeryCraft,
    ambient: 'amb-warm',
  },
  {
    value: 'hotpot_nine',
    label: '九宫格锅',
    preview: '九宫格 + 沸腾红油',
    gradient: 'linear-gradient(135deg,#160b08,#e0603d)',
    component: HotpotNine,
    ambient: 'amb-dark-red',
  },
  {
    value: 'pour_over_bar',
    label: '手冲吧台',
    preview: '深可可 + 克数参数',
    gradient: 'linear-gradient(135deg,#1b120d,#c08457)',
    component: PourOverBar,
    ambient: 'amb-dark-gold',
  },
  {
    value: 'cocktail_ticket',
    label: '琥珀酒单',
    preview: '墨绿票根 + 烫金字',
    gradient: 'linear-gradient(135deg,#0c1b14,#c9a84c)',
    component: CocktailTicket,
    ambient: 'amb-dark-green',
  },
  {
    value: 'macro_lunchbox',
    label: '增肌便当',
    preview: '荧光绿 + 菜单数据条',
    gradient: 'linear-gradient(135deg,#0a0b09,#a3e635)',
    component: MacroLunchbox,
    ambient: 'amb-dark-green',
  },
  {
    value: 'blueprint_kitchen',
    label: '厨房蓝图',
    preview: '蓝图网格 + 白色标注',
    gradient: 'linear-gradient(135deg,#0a2a45,#7fb9de)',
    component: BlueprintKitchen,
    ambient: 'amb-cool',
  },
  {
    value: 'morning_post',
    label: '晨报食单',
    preview: '红刊头 + 头版头条',
    gradient: 'linear-gradient(135deg,#fbf8ef,#b3261e)',
    component: MorningPost,
    ambient: 'amb-plain',
  },
  {
    value: 'meal_ticket',
    label: '食堂饭票',
    preview: '打孔饭票 + 红色盖章',
    gradient: 'linear-gradient(135deg,#d9c9a3,#8a3a22)',
    component: MealTicket,
    ambient: 'amb-warm',
  },
  {
    value: 'polaroid_scrapbook',
    label: '拍立得食记',
    preview: '拍立得斜贴 + 手写注',
    gradient: 'linear-gradient(135deg,#efe6d3,#b98a4e)',
    component: PolaroidScrapbook,
    ambient: 'amb-warm',
  },
  {
    value: 'gilded_supper',
    label: '金箔晚宴',
    preview: '黑金对称 + 装饰艺术',
    gradient: 'linear-gradient(135deg,#08080a,#c9a85f)',
    component: GildedSupper,
    ambient: 'amb-dark-gold',
  },
  {
    value: 'pixel_canteen',
    label: '像素食堂',
    preview: '像素游戏 + 复古面板',
    gradient: 'linear-gradient(135deg,#12142a,#facc15)',
    component: PixelCanteen,
    ambient: 'amb-dark',
  },
  {
    value: 'ink_wash_menu',
    label: '水墨食单',
    preview: '宣纸留白 + 墨迹印章',
    gradient: 'linear-gradient(135deg,#f3efe3,#a8452c)',
    component: InkWashMenu,
    ambient: 'amb-plain',
  },
  {
    value: 'film_contact',
    label: '胶片食记',
    preview: '胶片齿孔 + 接触印样',
    gradient: 'linear-gradient(135deg,#0d0d0d,#c7a86a)',
    component: FilmContact,
    ambient: 'amb-dark',
  },
  {
    value: 'cyber_canteen',
    label: '赛博食堂',
    preview: '终端绿字 + 扫描线',
    gradient: 'linear-gradient(135deg,#03100a,#79e69a)',
    component: CyberCanteen,
    ambient: 'amb-dark-green',
  },
  {
    value: 'herb_garden',
    label: '香草庭园',
    preview: '植物标本 + 花园手账',
    gradient: 'linear-gradient(135deg,#f3f1e3,#6f8557)',
    component: HerbGarden,
    ambient: 'amb-warm-green',
  },
]
