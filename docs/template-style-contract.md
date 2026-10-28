# 模板风格扩展契约（每类型 30 套独立风格）

## 目标

最终每个模板类型有 **30 套独立风格**，每套是一个独立的 `.vue` 文件，不共享布局包装组件。后续可以单独打开任意一个文件随意修改，不影响其他风格。

- 现有 7 个类型：保留已有 5 款，新增 25 款。
- 新增 3 个类型（夜店交际、二次元闪卡、宠物铭牌）：各 30 款。

## 文件位置

| 内容 | 路径 |
| --- | --- |
| 风格清单 | `frontend/src/templates/extensions/<type>.js` |
| 风格组件 | `frontend/src/pages/templates/ext/<type>/<StyleName>.vue` |
| 截图清单 | `frontend/src/templates/extensions/<type>.json` |

`<type>` 取值：`business_card`、`child_info`、`emergency_contact`、`company_website`、`product_display`、`recipe`、`custom`、`nightclub`、`anime_card`、`pet_tag`。

## 清单条目格式

`frontend/src/templates/extensions/<type>.js` 默认导出一个数组：

```js
import NeonNight from '../../pages/templates/ext/nightclub/NeonNight.vue'

export default [
  {
    value: 'neon_night',
    label: '霓虹夜色',
    preview: '深紫霓虹 + 超大姓名',
    gradient: 'linear-gradient(135deg,#1e1b4b,#a21caf)',
    component: NeonNight,
    ambient: 'amb-dark',
  },
]
```

- `value`：小写字母、数字、下划线，类型内唯一；不能与已有 5 款的 value 重复。
- `label`：2–5 个中文字。
- `preview`：8–16 个字的中文描述。
- `gradient`：CSS 渐变，用于后台表单和展厅缩略背景。
- `ambient` 只能取：`amb-dark`、`amb-cool`、`amb-plain`、`amb-warm-pink`、`amb-warm-green`、`amb-warm`、`amb-dark-red`、`amb-dark-gold`、`amb-dark-green`。
- 现有类型只放新增的 25 条；新类型放完整 30 条。

`frontend/src/templates/extensions/<type>.json` 只记录**新增**风格的 value 数组，例如：

```json
["neon_night", "velvet_lounge", "chrome_floor"]
```

## 组件契约

- props：`payload`（Object）与 `nfcId`（String / Number）。
- 根元素：`min-h-screen`；内容 `mx-auto w-full max-w-md` 居中，手机优先。
- 只读取 `payload` 里的字段；字段缺失时显示“未配置”或“—”，不能报错。
- 电话用 `<a :href="'tel:' + payload.phone">` 包住。
- 所有样式写在 `<style scoped>`，类名加唯一前缀（如 `.nc-`、`.ac-`、`.pt-`），避免互相污染。
- 图标使用 `lucide-vue-next`；装饰可用 CSS 渐变、内联 SVG，**不要引用外网图片**。
- 动效：入场动画 + 轻微 hover/active 反馈；必须写 `@media (prefers-reduced-motion: reduce)` 关闭动画。
- 不引入新依赖，不发网络请求。

## 差异化要求（每个类型 30 款）

- 至少覆盖 **8 个视觉家族**，例如：暗夜霓虹、极简白、奶油纸感、杂志编辑、玻璃拟态、赛博终端、复古印刷、金属奢华、手账拼贴、渐变流体、像素风、水墨、胶片、蓝图。
- 至少覆盖 **6 种版式骨架**，例如：居中名片、左对齐长页、分屏、卡片堆叠、全屏海报、时间线、票根、杂志封面。
- 任意两款不能同时满足“相同骨架 + 相同主色 + 相同字体气质”。
- 文案要贴合类型语境：夜店交际偏夜色与社交，二次元闪卡偏角色与收藏感，宠物铭牌偏温暖与安全。
- 30 款必须全部真实可用：禁止占位符、TODO、Lorem，禁止复制同一文件只换颜色。

## 禁止事项

- 不要修改 `frontend/src/templates/registry.js`、`ShowroomPage.vue`、`App.vue`、路由或其他类型的任何文件。
- 不要修改已有 5 款组件。
- 不要运行 build、dev server、截图命令（主流程统一执行）。
- 不要创建跨类型共用的包装组件；每套风格必须独立。

## 交付回复

完成后回复：

1. 新增风格数量；
2. value 列表；
3. 组件目录与清单文件路径；
4. 任何未完成或存疑的地方。
