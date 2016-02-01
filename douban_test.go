package bookstore

import (
	"fmt"
	"testing"
)

func TestUnmarshal(t *testing.T) {
	rawbookinfo := []byte(`{"rating":{"max":10,"numRaters":21,"average":"7.8","min":0},"subtitle":"","author":["王晓华"],"pubdate":"2015-4","tags":[{"count":121,"name":"算法","title":"算法"},{"count":43,"name":"算法\/数据结构","title":"算法\/数据结构"},{"count":41,"name":"计算机","title":"计算机"},{"count":36,"name":"编程","title":"编程"},{"count":15,"name":"计算机科学","title":"计算机科学"},{"count":12,"name":"程序员","title":"程序员"},{"count":10,"name":"计算机科学-算法与数据结构","title":"计算机科学-算法与数据结构"},{"count":5,"name":"programming","title":"programming"}],"origin_title":"","image":"https://img1.doubanio.com\/mpic\/s28033547.jpg","binding":"平装","translator":[],"catalog":"第1章 程序员与算法　　1\n1.1 什么是算法　　2\n1.2 程序员必须要会算法吗　　2\n1.2.1 一个队列引发的惨案　　3\n1.2.2 我的第一个算法　　5\n1.3 算法的乐趣在哪里　　7\n1.4 算法与代码　　8\n1.5 总结　　9\n1.6 参考资料　　9\n第2章 算法设计的基础　　10\n2.1 程序的基本结构　　10\n2.1.1 顺序执行　　10\n2.1.2 循环结构　　11\n2.1.3 分支和跳转结构　　13\n2.2 算法实现与数据结构　　16\n2.2.1 基本数据结构在算法设计中的应用　　16\n2.2.2 复杂数据结构在算法设计中的应用　　19\n2.3 数据结构和数学模型与算法的关系　　24\n2.4 总结　　25\n2.5 参考资料　　25\n第3章 算法设计的常用思想　　26\n3.1 贪婪法　　26\n3.1.1 贪婪法的基本思想　　27\n3.1.2 贪婪法的例子：0-1 背包问题　　27\n3.2 分治法　　30\n3.2.1 分治法的基本思想　　30\n3.2.2 递归和分治，一对好朋友　　31\n3.2.3 分治法的例子：大整数Karatsuba 乘法算法　　32\n3.3 动态规划　　34\n3.3.1 动态规划的基本思想　　34\n3.3.2 动态规划法的例子：字符串的编辑距离　　37\n3.4 解空间的穷举搜索　　40\n3.4.1 解空间的定义　　41\n3.4.2 穷举解空间的策略　　42\n3.4.3 穷举搜索的例子：Google 方程式　　44\n3.5 总结　　46\n3.6 参考资料　　46\n第4章 阿拉伯数字与中文数字　　47\n4.1 中文数字的特点　　47\n4.1.1 中文数字的权位和小节　　48\n4.1.2 中文数字的零　　48\n4.2 阿拉伯数字转中文数字　　49\n4.2.1 一个转换示例　　49\n4.2.2 转换算法设计　　49\n4.2.3 算法实现　　50\n4.2.4 中文大写数字　　51\n4.3 中文数字转阿拉伯数字　　52\n4.3.1 转换的基本方法　　52\n4.3.2 算法实现　　52\n4.4 数字转换的测试用例　　54\n4.5 总结　　55\n4.6 参考资料　　55\n第5章 三个水桶等分8 升水的问题　　56\n5.1 问题与求解思路　　57\n5.2 建立数学模型　　58\n5.2.1 状态的数学模型与状态树　　58\n5.2.2 倒水动作的数学模型　　59\n5.3 搜索算法　　60\n5.3.1 状态树的遍历　　60\n5.3.2 剪枝和重复状态判断　　61\n5.4 算法实现　　62\n5.5 总结　　64\n5.6 参考资料　　64\n第6章 妖怪与和尚过河问题　　65\n6.1 问题与求解思路　　66\n6.2 建立数学模型　　66\n6.2.1 状态的数学模型与状态树　　67\n6.2.2 过河动作的数学模型　　67\n6.3 搜索算法　　69\n6.3.1 状态树的遍历　　70\n6.3.2 剪枝和重复状态判断　　70\n6.4 算法实现　　71\n6.5 总结　　72\n6.6 参考资料　　73\n第7章 稳定匹配与舞伴问题　　 74\n7.1 稳定匹配问题　　74\n7.1.1 什么是稳定匹配　　74\n7.1.2 Gale-Shapley 算法原理　　75\n7.2 Gale-Shapley 算法的应用实例　　77\n7.2.1 算法实现　　77\n7.2.2 改进优化：空间换时间　　80\n7.3 有多少稳定匹配　　81\n7.3.1 穷举所有的完美匹配　　 81\n7.3.2 不稳定因素的判断算法　　 82\n7.3.3 穷举的结果　　 84\n7.4 二部图与二分匹配　　 84\n7.4.1 最大匹配与匈牙利算法　　 85\n7.4.2 带权匹配与Kuhn-Munkres算法　　 88\n7.5 总结　　 93\n7.6 参考资料　　 94\n第8章 爱因斯坦的思考题　　 95\n8.1 问题的答案　　 96\n8.2 分析问题的数学模型　　 96\n8.2.1 基本模型定义　　 96\n8.2.2 线索模型定义　　 98\n8.3 算法设计　　 99\n8.3.1 穷举所有的组合结果　　 99\n8.3.2 利用线索判定结果的正确性　　 101\n8.4 总结　　 103\n8.5 参考资料　　 104\n第9章 项目管理与图的拓扑排序　　 105\n9.1 AOV 网和AOE 网　　 107\n9.2 拓扑排序　　 108\n9.2.1 拓扑排序的基本过程　　 108\n9.2.2 按照活动开始时间排序　　 108\n9.3 关键路径算法　　 111\n9.3.1 什么是关键路径　　 112\n9.3.2 计算关键路径的算法　　 113\n9.4 总结　　 116\n9.5 参考资料　　 116\n第10章 RLE 压缩算法与PCX 图像文件格式　　 117\n10.1 RLE 压缩算法　　 117\n10.1.1 连续重复数据的处理　　 117\n10.1.2 连续非重复数据的处理　　 118\n10.1.3 算法实现　　118\n10.2 RLE 与PCX 图像文件格式　　121\n10.2.1 PCX 图像文件格式　　121\n10.2.2 PCX_RLE 算法　　122\n10.2.3 256 色PCX 文件的解码和显示　　123\n10.3 总结　　124\n10.4 参考资料　　125\n第11章 算法与历法　　126\n11.1 格里历（公历）生成算法　　126\n11.1.1 格里历的历法规则　　126\n11.1.2 今天星期几　　127\n11.1.3 生成日历的算法　　131\n11.1.4 日历变更那点事儿　　132\n11.2 二十四节气的天文学计算　　134\n11.2.1 二十四节气的起源　　134\n11.2.2 二十四节气的天文学定义　　135\n11.2.3 VSOP-82\/87 行星理论　　137\n11.2.4 误差修正——章动　　141\n11.2.5 误差修正——光行差　　143\n11.2.6 用牛顿迭代法计算二十四节气　　144\n11.3 农历朔日（新月）的天文学计算　　146\n11.3.1 日月合朔的天文学定义　　147\n11.3.2 ELP-2000\/82 月球理论　　147\n11.3.3 误差修正——地球轨道离心率修正　　149\n11.3.4 误差修正——黄经摄动　　149\n11.3.5 月球地心视黄经和最后的修正——地球章动　　150\n11.3.6 用牛顿迭代法计算日月合朔　　151\n11.4 农历的生成算法　　152\n11.4.1 中国农历的起源与历法规则　　153\n11.4.2 中国农历的推算　　157\n11.4.3 一个简单的“年历” 　　165\n11.5 总结　　166\n11.6 参考资料　　167\n第12章 实验数据与曲线拟合　　168\n12.1 曲线拟合　　168\n12.1.1 曲线拟合的定义　　168\n12.1.2 简单线性数据拟合的例子　　168\n12.2 最小二乘法曲线拟合　　169\n12.2.1 最小二乘法原理　　170\n12.2.2 高斯消元法求解方程组　　171\n12.2.3 最小二乘法解决“速度与加速度”实验　　172\n12.3 三次样条曲线拟合　　173\n12.3.1 插值函数　　174\n12.3.2 样条函数的定义　　174\n12.3.3 边界条件　　175\n12.3.4 推导三次样条函数　　176\n12.3.5 追赶法求解方程组　　179\n12.3.6 三次样条曲线拟合算法实现　　181\n12.3.7 三次样条曲线拟合的效果　　183\n12.4 总结　　184\n12.5 参考资料　　184\n第13章 非线性方程与牛顿迭代法　　185\n13.1 非线性方程求解的常用方法　　185\n13.1.1 公式法　　185\n13.1.2 二分逼近法　　186\n13.2 牛顿迭代法的数学原理　　187\n13.3 用牛顿迭代法求解非线性方程的实例　　188\n13.3.1 导函数的求解与近似公式　　188\n13.3.2 算法实现　　188\n13.4 参考资料　　189\n第14章 计算几何与计算机图形学　　190\n14.1 计算几何的基本算法　　190\n14.1.1 点与矩形的关系　　190\n14.1.2 点与圆的关系　　191\n14.1.3 矢量的基础知识　　191\n14.1.4 点与直线的关系　　194\n14.1.5 直线与直线的关系　　194\n14.1.6 点与多边形的关系　　196\n14.2 直线生成算法　　199\n14.2.1 什么是光栅图形扫描转换　　200\n14.2.2 数值微分法　　200\n14.2.3 Bresenham 算法　　202\n14.2.4 对称直线生成算法　　204\n14.2.5 两步算法　　205\n14.2.6 其他直线生成算法　　207\n14.3 圆生成算法　　207\n14.3.1 圆的八分对称性　　208\n14.3.2 中点画圆法　　209\n14.3.3 改进的中点画圆法——Bresenham 算法　　210\n14.3.4 正负判定画圆法　　211\n14.4 椭圆生成算法　　212\n14.4.1 中点画椭圆法　　213\n14.4.2 Bresenham 椭圆算法　　215\n14.5 多边形区域填充算法　　217\n14.5.1 种子填充算法　　218\n14.5.2 扫描线填充算法　　223\n14.5.3 改进的扫描线填充算法　　229\n14.5.4 边界标志填充算法　　233\n14.6 总结　　236\n14.7 参考资料　　236\n第15章 音频频谱和均衡器与傅里叶变换算法　　237\n15.1 实时频谱显示的原理　　237\n15.2 离散傅里叶变换　　238\n15.2.1 什么是傅里叶变换　　239\n15.2.2 傅里叶变换原理　　 239\n15.2.3 快速傅里叶变换算法的实现　　 243\n15.3 傅里叶变换与音频播放的实时频谱显示　　 245\n15.3.1 频域数值的特点分析　　 245\n15.3.2 从音频数据到功率频谱　　 246\n15.3.3 音频播放时实时频谱显示的例子　　 248\n15.4 破解电话号码的小把戏　　 251\n15.4.1 拨号音的频谱分析　　 251\n15.4.2 根据频谱数据反推电话号码　　 252\n15.5 离散傅里叶逆变换　　 253\n15.5.1 快速傅里叶逆变换的推导　　 254\n15.5.2 快速傅里叶逆变换的算法实现　　 254\n15.6 利用傅里叶变换实现频域均衡器　　 255\n15.6.1 频域均衡器的实现原理　　 255\n15.6.2 频域信号的增益与衰减　　 256\n15.6.3 均衡器的实现——仿Foobar的18 段均衡器　　 258\n15.7 总结　　 259\n15.8 参考资料　　 259\n第16章 全局最优解与遗传算法　　 260\n16.1 遗传算法的原理　　 260\n16.1.1 遗传算法的基本概念　　 261\n16.1.2 遗传算法的处理流程　　 262\n16.2 遗传算法求解0-1 背包问题　　 267\n16.2.1 基因编码和种群初始化　　 267\n16.2.2 适应度函数　　 268\n16.2.3 选择算子设计与轮盘赌算法　　 268\n16.2.4 交叉算子设计　　 270\n16.2.5 变异算子设计　　 271\n16.2.6 这就是遗传算法　　 272\n16.3 总结　　272\n16.4 参考资料　　273\n第17章 计算器程序与大整数计算　　274\n17.1 哦，溢出了，出洋相的计算器程序　　274\n17.2 大整数计算的原理　　275\n17.2.1 大整数加法　　276\n17.2.2 大整数减法　　278\n17.2.3 大整数乘法　　279\n17.2.4 大整数除法与模　　281\n17.2.5 大整数乘方运算　　282\n17.3 大整数类的使用　　283\n17.3.1 与Windows的计算器程序一决高下　　283\n17.3.2 最大公约数和最小公倍数　　284\n17.3.3 用扩展欧几里得算法求模的逆元　　286\n17.4 总结　　288\n17.5 参考资料　　288\n第18章 RSA 算法——加密与签名　　289\n18.1 RSA 算法的开胃菜　　289\n18.1.1 将模幂运算转化为模乘运算　　290\n18.1.2 模乘运算与蒙哥马利算法　　291\n18.1.3 模幂算法　　292\n18.1.4 素数检验与米勒—拉宾算法　　292\n18.2 RSA 算法原理　　295\n18.2.1 RSA 算法的数学理论　　295\n18.2.2 加密和解密算法　　296\n18.2.3 RSA 算法的安全性　　297\n18.3 数据块分组加密　　297\n18.3.1 字节流与大整数的转换　　298\n18.3.2 PCKS 与OAEP 加密填充模式　　298\n18.3.3 数据加密算法实现　　300\n18.3.4 数据解密算法实现　　301\n18.4 RSA 签名与身份验证　　302\n18.4.1 RSASSA-PKCS 与RSASSAPSS签名填充模式　　302\n18.4.2 签名算法实现　　304\n18.4.3 验证签名算法实现　　305\n18.5 总结　　305\n18.6 参考资料　　306\n第19章 数独游戏　　307\n19.1 数独游戏的规则与技巧　　307\n19.1.1 数独游戏的规则　　307\n19.1.2 数独游戏的常用技巧　　308\n19.2 计算机求解数独问题　　308\n19.2.1 建立问题的数学模型　　310\n19.2.2 算法实现　　311\n19.2.3 与传统穷举方法的结果对比　　312\n19.3 关于数独的趣味话题　　312\n19.3.1 数独游戏有多少终盘　　313\n19.3.2 史上最难的数独游戏　　314\n19.4 总结　　314\n19.5 参考资料　　315\n第20章 华容道游戏　　316\n20.1 华容道游戏介绍　　316\n20.2 自动求解的算法原理　　317\n20.2.1 定义棋盘的局面　　317\n20.2.2 算法思路　　319\n20.3 自动求解的算法实现　　320\n20.3.1 棋局状态与Zobrist 哈希算法　　321\n20.3.2 重复棋局和左右镜像的处理　　323\n20.3.3 正确结果的判断条件　　325\n20.3.4 武将棋子的移动　　325\n20.3.5 棋局的搜索算法　　328\n20.4 总结　　329\n20.5 参考资料　　329\n第21章 A*寻径算法　　330\n21.1 寻径算法演示程序　　330\n21.2 Dijkstra 算法　　331\n21.2.1 Dijkstra 算法原理　　332\n21.2.2 Dijkstra 算法实现　　332\n21.2.3 Dijkstra 算法演示程序　　333\n21.3 带启发的搜索算法——A*算法　　335\n21.3.1 A*算法原理　　336\n21.3.2 常用的距离评估函数　　337\n21.3.3 A*算法实现　　340\n21.4 总结　　342\n21.5 参考资料　　342\n第22章 俄罗斯方块游戏　　343\n22.1 俄罗斯方块游戏规则　　343\n22.2 俄罗斯方块游戏人工智能的算法原理　　344\n22.2.1 影响评价结果的因素　　345\n22.2.2 常用的俄罗斯方块游戏人工智能算法　　346\n22.2.3 Pierre Dellacherie 评估算法　　347\n22.3 Pierre Dellacherie 算法实现　　349\n22.3.1 基本数学模型和数据结构定义　　350\n22.3.2 算法实现　　352\n22.4 总结　　358\n22.5 参考资料　　358\n第23章 博弈树与棋类游戏　　359\n23.1 棋类游戏的AI　　359\n23.1.1 博弈与博弈树　　360\n23.1.2 极大极小值搜索算法　　361\n23.1.3 负极大极搜索算法　　 362\n23.1.4 “α-β”剪枝算法　　 363\n23.1.5 估值函数　　 365\n23.1.6 置换表与哈希函数　　 366\n23.1.7 开局库与终局库　　 368\n23.2 井字棋——最简单的博弈游戏　　 368\n23.2.1 棋盘与棋子的数学模型　　 369\n23.2.2 估值函数与估值算法　　 370\n23.2.3 如何产生走法（落子方法）　　 371\n23.3 奥赛罗棋（黑白棋） 　　 373\n23.3.1 棋盘与棋子的数学模型　　 374\n23.3.2 估值函数与估值算法　　 377\n23.3.3 搜索算法实现　　 380\n23.3.4 最终结果　　 384\n23.4 五子棋　　 385\n23.4.1 棋盘与棋子的数学模型　　 386\n23.4.2 估值函数与估值算法　　 388\n23.4.3 搜索算法实现　　 391\n23.4.4 最终结果　　 393\n23.5 总结　　 393\n23.6 参考资料　　 393\n附录A 算法设计的常用技巧　　 395\nA.1 数组下标处理　　 395\nA.2 一重循环实现两重循环的功能　　 396\nA.3 棋盘（迷宫）类算法方向遍历　　 396\nA.4 代码的一致性处理技巧　　 397\nA.5 链表和数组的配合使用　　 398\nA.6 “以空间换时间”的常用技巧　　 399\nA.7 利用表驱动避免长长的switch-case 　　 400\n附录B 一个棋类游戏的设计框架　　 401\nB.1 代码框架的整体结构　　 401\nB.2 代码框架的使用方法　　 403","ebook_url":"http:\/\/read.douban.com\/ebook\/12053861\/","pages":"420","images":{"small":"https://img1.doubanio.com\/spic\/s28033547.jpg","large":"https://img1.doubanio.com\/lpic\/s28033547.jpg","medium":"https://img1.doubanio.com\/mpic\/s28033547.jpg"},"alt":"http:\/\/book.douban.com\/subject\/26351257\/","id":"26351257","publisher":"人民邮电出版社","isbn10":"7115385378","isbn13":"9787115385376","title":"算法的乐趣","url":"http:\/\/api.douban.com\/v2\/book\/26351257","alt_title":"","author_intro":"王晓华\n2005年毕业于华中科技大学，目前在中兴通讯上海研发中心从事光纤接入网通讯设备开发，担任EPON（以太网无源光网络）业务软件开发经理，参与开发的PON设备在全球部署过亿线，为数亿家庭提供宽带接入服务。\n业余时间喜欢研究算法和写作博客（http:\/\/blog.csdn.net\/orbit），最大的乐趣就是用程序解决生活中的问题：\n为了方便使用Visual Studio 6.0开发软件，曾特意编写并开源了一个tabbar插件；\n为了文档安全，开发了一个基于layerFSD技术的透明文件加密系统；\n使用Source Insight软件觉得不习惯，于是以外挂的形式开发了TabSiPlus插件……\n算法可以做的事情还有很多，期待我们会有更多发现！","summary":"算法之大，大到可以囊括宇宙万物的运行规律；算法之小，小到寥寥数行代码即可展现一个神奇的功能。算法的应用和乐趣在生活中无处不在：\n历法和二十四节气计算使用的是霍纳法则和求解一元高次方程的牛顿迭代法；\n音频播放器跳动的实时频谱背后是离散傅立叶变换算法；\nDOS时代著名的PCX图像文件格式使用的是简单有效的RLE压缩算法；\nRSA加密算法的光环之下是朴实的欧几里德算法、蒙哥马利算法和米勒-拉宾算法；\n井字棋、黑白棋、五子棋和俄罗斯方块游戏背后是各种有趣的AI算法；\n华容道游戏求解的简单穷举算法中还蕴藏着对棋盘状态的哈希算法；\n遗传算法神秘不可测，但用遗传算法求解0-1背包问题只用了60多行代码……\n一本书带你走进色彩缤纷的算法世界，让你尽享算法的乐趣。","ebook_price":"39.99","price":"79.00元"}`)

	bookinfo, err := Unmarshal(rawbookinfo)
	if err != nil {
		t.Errorf("Unmarshal fail, %v", err)
	}
	fmt.Printf("%s %s %s %s", bookinfo.ISBN13, bookinfo.Title, bookinfo.Price, bookinfo.Publisher)

}
