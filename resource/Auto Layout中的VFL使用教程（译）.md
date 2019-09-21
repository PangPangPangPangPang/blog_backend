# Auto Layout 中的 VFL 使用教程（译）

[date] 2015-12-11 14:12:37
[tag] AutoLayout

[原文链接](http://www.raywenderlich.com/110393/auto-layout-visual-format-language-tutorial)

Auto Layout 的可视格式化语言（以下简称 VFL）允许使用者通过 ASCII-art 格式化字符串定义约束。
用一行简单的代码，你可以定义多个水平或垂直方向的约束。对比一个一个加约束，这样可以可以节省很多代码量。
在这个教程中，你可以用 VFL 做下面这些事情哦：!

- 构建水平和垂直的约束
- 在 VFL 中使用**views**描述
- 在 VFL 中使用**metrics**
- 使用**layout options**去关联其他界面元素
- 使用**layout guides**处理视图的上下边距（译者：比如**UINavigationBar**）

> **注意：建议读者对 Auto Layout 有充分了解的情况下阅读此文，如果对于 Auto Layout 不是很熟悉，建议先阅读[Auto Layout Tutorial Part 1: Getting Started](http://www.raywenderlich.com/115440/auto-layout-tutorial-in-ios-9-part-1-getting-started-2)和[Auto Layout Tutorial Part 2: Constraints](http://www.raywenderlich.com/115444/auto-layout-tutorial-in-ios-9-part-2-constraints)**

### 准备开始吧！

首先下载[事例工程](http://cdn4.raywenderlich.com/wp-content/uploads/2015/09/Grapevine-Starter.zip)便于教程使用，该工程提供了一个初级网络社交 app-**Grapevine**的基本欢迎页面。在 Xcode 中运行工程；你将看到如下画面（在模拟器的**Hardware\Rotate Right**中旋转屏幕）:
![welcomepage](https://koenig-media.raywenderlich.com/uploads/2015/09/Grapevine-Initial-Screen.png)
好吧，这个页面真是一团乱，为什么这种情况会发生呢？面对这种情况我们应该怎么做呢？
当前界面的所有元素都是跟界面的上边缘（top）和左边缘（left）联系的，这是因为它们没有用 Auto Layout 约束。通过接下来的教程你会让视图看起来更漂亮。
打开**Main.storyboard**观察界面元素。注意到这些元素都被设置为在编译期移除 Auto Layout 约束。你不应该在真实项目中这样使用，但是这会让你节省一些元素的初始化时间。
接下来，打开**ViewController.swift**。在顶部，你可以看到在**Main.storyboard**中跟 Interface Builder（IB）视图元素联系的 outlet 和一些在 runtime 代替约束的属性。
这个时候没啥可以说，但是接下来有一大堆跟 VFL 有关的东西要学！

### VFL 语法

在你开始编写布局和约束之前，你需要有一些关于 VFL 格式化串的相关知识。
第一件要知道的事情：VFL 格式化串可以分成如下组成：
![formatString](https://koenig-media.raywenderlich.com/uploads/2015/07/VisualFormatLanguageOptionsImage.png)
接下来一个一个解释 VFL 格式化串：

1. 约束的方向，非必须。可以有以下的值：
   - H:表示水平方向。
   - V:表示垂直方向。
   - 不指定:Auto Layout 默认水平方向。
2. 和父视图的头部关联，非必须
   - 父视图的上边缘和视图的上边缘的距离（垂直方向）
   - 父视图的头部边缘和视图的头部边缘的距离（水平方向）
3. 需要布局的视图，必须。
4. 跟另一个视图关联，非必须。
5. 和父视图的尾部关联，非必须。
   - 父视图的下边缘和视图的下边缘的距离（垂直方向）
   - 父视图的尾部边缘和视图的尾部边缘的距离（水平方向）
6. 另外在上图中还有两个特殊的字符，他们的意思是：

   - **"?"**代表在格式化串中非必须。
   - **"\*"**代表允许在格式化串中出现 0 次或多次。

### 可使用的符号

VFL 使用一系列符号去描述布局

- **|** 父视图
- **-** 标准间距(通常 8pt；如果这个代表到父视图边缘的间距可以改变)
- **==** 宽度相等（可被删除）
- **-20-** 不标准间距（20pt）
- **<=** 小于等于
- **>=** 大于等于
- **@250** 约束权重；可以为 0 到 1000 得任意值

  - 250 - 低权重
  - 750 - 高权重
  - 1000 - 绝对权重

### 格式化串实例

      H:|-[icon(==iconDate)]-20-[iconLabel(120@250)]-20@750-[iconDate(>=50)]-|

接下来一步一步解释这个串：

- H: 水平方向。
- |-[icon icon 的头边缘和父视图有一个标准间距。
- ==iconDate icon 的宽度应该和 iconDate 的宽度相等。
- ]-20-[iconLabel icon 的尾边缘和 iconLabel 的头边缘有 20pt 的距离。
- [iconLabel(120@250)]iconLabel 有一个 120pt 的宽度，设置成低权重，如果出现冲突 Auto Layout 会打破这条约束。
- -20@750- iconLabel 的的尾边缘和 iconDate 的头边缘有 20pt 的距离，设置成高权重，如果出现冲突 Auto Layout 不会打破这条约束。
- [iconDate(>=50)] iconDate 的宽度应该大于等于 50pt。
- -| iconDate 的尾边缘和父视图的尾边缘有一个标准间距。

![good](https://koenig-media.raywenderlich.com/uploads/2015/08/got_it.png)

现在你对 VFL 已经有了一个基本的认识--接下来就要把这些知识用到实际应用中了。

### 创建约束

Apple 在**NSLayoutConstraint**提供了类方法**constraintsWithVisualFormat**去创建约束。你将在**Grapevine**程序化的创建约束

在 Xcode 中打开**ViewController.swift**，并且添加如下代码到**viewDidLoad()**中：

```swift
appImageView.hidden = true
welcomeLabel.hidden = true
summaryLabel.hidden = true
pageControl.hidden = true
```

这些代码会隐藏除了**iconImageView**，**appNameLabel**和**skipButton**之外的元素。运行工程；你会看到如下：
![good](https://koenig-media.raywenderlich.com/uploads/2015/09/Grapevine-Hidden-Icons.png)
棒！你现在已经清除了烦人的元素了，现在在**viewDidLoad()**添加如下代码：

```swift
// 1
let views = ["iconImageView": iconImageView,
  "appNameLabel": appNameLabel,
  "skipButton": skipButton]

// 2
var allConstraints = [NSLayoutConstraint]()

// 3
let iconVerticalConstraints = NSLayoutConstraint.constraintsWithVisualFormat(
  "V:|-20-[iconImageView(30)]",
  options: [],
  metrics: nil,
  views: views)
allConstraints += iconVerticalConstraints

// 4
let nameLabelVerticalConstraints = NSLayoutConstraint.constraintsWithVisualFormat(
  "V:|-23-[appNameLabel]",
  options: [],
  metrics: nil,
  views: views)
allConstraints += nameLabelVerticalConstraints

// 5
let skipButtonVerticalConstraints = NSLayoutConstraint.constraintsWithVisualFormat(
  "V:|-20-[skipButton]",
  options: [],
  metrics: nil,
  views: views)
allConstraints += skipButtonVerticalConstraints

// 6
let topRowHorizontalConstraints = NSLayoutConstraint.constraintsWithVisualFormat(
  "H:|-15-[iconImageView(30)]-[appNameLabel]-[skipButton]-15-|",
  options: [],
  metrics: nil,
  views: views)
allConstraints += topRowHorizontalConstraints

// 7
NSLayoutConstraint.activateConstraints(allConstraints)
```

接下来一步步解释上面的代码：

1. 创建一个字典，这个字典用字符串和 view 对应，用来在格式化串中使用。
2. 创建一个约束数组，你会在接下来的代码中向里面添加约束。
3. 创建**iconImageView**的垂直约束，距父视图的上边缘 20pt，本身高度 30pt。
4. 创建**appNameLabel**的垂直约束，距父视图的上边缘 23pt。
5. 创建**skipButton**的垂直约束，距父视图的上边缘 20pt。
6. 设置上面三个元素的水平约束，**iconImageView**的头边缘距父视图的头边缘 8pt，宽度 30pt。接下来，**iconImageView**的尾边缘距**appNameLabel**头边缘 8pt，**appNameLabel**的尾边缘距**skipButton**的头边缘 8pt，最后**skipButton**的尾边缘距离父视图的尾边缘 15pt。
7. 用**NSLayoutConstraint**的类方法**activateConstraints(\_:)**启用约束。在这个步骤你需要添加**allConstraints**数组。

> **注意：在 views 字典中的 key 必须在格式化串中得 view 串匹配。如果没有，Auto Layout 将不能找到引用并且在 runtime 崩溃。**

运行工程，元素现在看起来怎么样？
![good](https://koenig-media.raywenderlich.com/uploads/2015/09/Grapevine-Horizontal-Layout.png)
哈哈，看看是不是已经变得好看多了？
现在把它放着，这不过是个前戏（误）。你还要有一大坨代码要写呢，但是到最后这些都是值得的。
接下来，你需要给剩下的元素布局，首先，你需要把最开始加到**viewDidLoad()**的代码去掉。不要有怨言，删除下面这些：

```swift
	appImageView.hidden = true
	welcomeLabel.hidden = true
	summaryLabel.hidden = true
	pageControl.hidden = true
```

这样最开始隐藏的元素就又出现了。
接下来，把当前的**views**替换成如下的代码：

```swift
let views = ["iconImageView": iconImageView,
 "appNameLabel": appNameLabel,
 "skipButton": skipButton,
 "appImageView": appImageView,
 "welcomeLabel": welcomeLabel,
 "summaryLabel": summaryLabel,
 "pageControl": pageControl]
```

现在你已经为**appImageView**，**welcomeLabel**，**summaryLabel**和**pageControl**添加了视图定义，这些都可以在 VFL 格式化串中使用。
在**activateConstraints()**调用之前，在**viewDidLoad()**中添加如下代码：

```swift
// 1
let summaryHorizontalConstraints = NSLayoutConstraint.constraintsWithVisualFormat(
  "H:|-15-[summaryLabel]-15-|",
  options: [],
  metrics: nil,
  views: views)
allConstraints += summaryHorizontalConstraints

let welcomeHorizontalConstraints = NSLayoutConstraint.constraintsWithVisualFormat(
  "H:|-15-[welcomeLabel]-15-|",
  options: [],
  metrics: nil,
  views: views)
allConstraints += welcomeHorizontalConstraints

// 2
let iconToImageVerticalConstraints = NSLayoutConstraint.constraintsWithVisualFormat(
  "V:[iconImageView]-10-[appImageView]",
  options: [],
  metrics: nil,
  views: views)
allConstraints += iconToImageVerticalConstraints

// 3
let imageToWelcomeVerticalConstraints = NSLayoutConstraint.constraintsWithVisualFormat(
  "V:[appImageView]-10-[welcomeLabel]",
  options: [],
  metrics: nil,
  views: views)
allConstraints += imageToWelcomeVerticalConstraints

// 4
let summaryLabelVerticalConstraints = NSLayoutConstraint.constraintsWithVisualFormat(
  "V:[welcomeLabel]-4-[summaryLabel]",
  options: [],
  metrics: nil,
  views: views)
allConstraints += summaryLabelVerticalConstraints

// 5
let summaryToPageVerticalConstraints = NSLayoutConstraint.constraintsWithVisualFormat(
  "V:[summaryLabel]-15-[pageControl(9)]-15-|",
  options: [],
  metrics: nil,
  views: views)
allConstraints += summaryToPageVerticalConstraints
```

接下来一步步解释上面的代码：

1. 创建**summaryLabel**和**welcomeLabel**的水平约束，让它们的头边缘和尾边缘分别距父视图的头边缘和尾边缘 15pt。
2. 创建 icon 和 app image 的垂直约束，两者距离 10pt。
3. 创建 app image 和 welcome label 的垂直约束，两者距离 10pt。
4. 创建 welcome label 和 summary label 的垂直约束，两者距离 4pt。
5. 创建 summary label 和 page control 的垂直约束，两者相距 15pt，并且 page control 高度为 9pt，和父视图的底边缘距离 15。

运行工程；这些元素看起来怎么样？
![good](https://koenig-media.raywenderlich.com/uploads/2015/09/Grapevine-Layout-Before-Options.png)
现在看起来还不错了哦。错，其中的一些元素的布局是正确的，然后，有些并没有，image 和 page control 并没有居中！
![bad](https://koenig-media.raywenderlich.com/uploads/2015/07/No-Center-RageMakger.png)
不要害怕，下一节将会告诉你更多关于布局的工具。

### Layout Options

Layout Options 提供了一个让你在定义约束的时候对视图进行垂线方向上的约束。
使用**NSLayoutFormatOptions.AlignAllCenterY**是一个使用 Layout Options 的例子，它可以让 view 在创建水平约束的时候同时让垂直方向居中。
如果你不想让水平布局的时候垂直方向都居中，而是边对边的话，那就不应该用这个选项。
接下来，让我们看看 Layout Options 在创建约束的时候是多么有用。移除**viewDidLoad()**中如下的代码：

```swift
let nameLabelVerticalConstraints = NSLayoutConstraint.constraintsWithVisualFormat(
  "V:|-23-[appNameLabel]",
  options: [],
  metrics: nil,
  views: views)
allConstraints += nameLabelVerticalConstraints

let skipButtonVerticalConstraints = NSLayoutConstraint.constraintsWithVisualFormat(
  "V:|-20-[skipButton]",
  options: [],
  metrics: nil,
  views: views)
allConstraints += skipButtonVerticalConstraints
```

你刚刚移除了**appNameLabel**和**skipButton**的垂直布局。作为替代，你将用 Layout Options 去给它们添加垂直约束。
找到创建**topRowHorizontalConstraints**的代码并且设置**options**为**[.AlignAllCenterY]**。看起来是这个样子的：

```swift
let topRowHorizontalConstraints = NSLayoutConstraint.constraintsWithVisualFormat(
  "H:|-15-[iconImageView(30)]-[appNameLabel]-[skipButton]-15-|",
  options: [.AlignAllCenterY],
  metrics: nil,
  views: views)
```

添加**NSLayoutFormatOption .AlignAllCenterY**对上面格式化串中的所有视图都有效，并且创建了一个它们垂直方向中心的约束。如果**iconImageView**提前创建了包含高度的垂直约束也是有效的。因此，**appNameLabel**和**skipButton**同**iconImageView**一样垂直居中。
如果你现在运行，布局看起来可能没有改变，但是代码变得更棒了。移除创建**welcomeHorizontalConstraints**和将它放进数组的代码。这样就移除了**welcomeLabel**的水平约束。接下来，更新创建**summaryLabelVerticalConstraints**的 Layout Options：

```swift
summaryLabelVerticalConstraints = NSLayoutConstraint.constraintsWithVisualFormat("V:[welcomeLabel]-4-[summaryLabel]",
options: [.AlignAllLeading, .AlignAllTrailing],
metrics: nil,
views: views);
```

这个代码增加了**NSLayoutFormatOptions**的**.NSLayoutFormatOptions**和**.AlignAllTrailing**，**welcomeLabel**和**summaryLabel's**的头边缘和尾边缘会距离它们的父视图的边缘 15pt。由于提前为**summaryLabel**定义了水平约束，所以上述代码才会有效。虽然上面的代码带来的是同样的效果，但是实现起来更加优雅了。
接下来，更新你在创建**summaryToPageVerticalConstraints**时候的选项：

```swift
let pageControlVerticalConstraints = NSLayoutConstraint.constraintsWithVisualFormat(
  "V:[summaryLabel]-15-[pageControl(9)]-15-|",
  options: [.AlignAllCenterX],
  metrics: nil,
  views: views)
```

这样就添加了沿 x 轴中心对齐。同样为**imageToWelcomeVerticalConstraints**添加选项：

```swift
let imageToWelcomeVerticalConstraints = NSLayoutConstraint.constraintsWithVisualFormat(
  "V:[appImageView]-10-[welcomeLabel]",
  options: [.AlignAllCenterX],
  metrics: nil,
  views: views)
```

运行工程，看看发生了什么？
![perfect](https://koenig-media.raywenderlich.com/uploads/2015/09/Grapevine-SublayoutViewHeights.png)
感觉都居中了是吧？Layout Options 让你做出了一个更棒的交互界面。

###NSLayoutFormat 选项快速参考
下面是在**Grapevine**中使用过的属性：

- **.AlignAllCenterX** --使用**NSLayoutAttributeCenterX**的对齐元素
- **.AlignAllCenterY** --使用**NSLayoutAttributeCenterY**的对齐元素
- **.AlignAllLeading** --使用**NSLayoutAttributeLeading**的对齐元素
- **.AlignAllTrailing** --使用**NSLayoutAttributeTrailing**的对齐元素

（译者：由于有些种类的文字是从右到左书写的，所以它们的**.AlignAllLeading**等价于**.AlignAllRight**，而对于中文来说，**.AlignAllLeading**等价于**.AlignAllLeft**）

下面是剩余的一些属性：

- **.AlignAllLeft** --使用**NSLayoutAttributeLeft**的对齐元素
- **.AlignAllRight** --使用**NSLayoutAttributeRight**的对齐元素
- **.AlignAllTop** --使用**NSLayoutAttributeBottom**的对齐元素
- **.AlignAllBottom** --使用**NSLayoutAttributeCenterX**的对齐元素
- **.AlignAllBaseline** --使用**NSLayoutAttributeBaseline**的对齐元素

你同样可以在[文档](https://developer.apple.com/library/ios/documentation/AppKit/Reference/NSLayoutConstraint_Class/#//apple_ref/c/tdef/NSLayoutFormatOptions)详细查看。

> ** 注意：为了让 Layout Options 有效，至少要有一个元素定义过垂直方向的约束。看下面的例子：**

```swift
   NSLayoutConstraints.constraintsWithVisualFormat(
     "V:[topView]-[middleView]-[bottomView]",
     options: [.AlignAllLeading],
     metrics: nil,
     views: ["topView": topView, "middleView": middleView, "bottomView":"bottomView"])
```

> \*\* **topView**，**middleView**或者**bottomView**其中一个必须要有一个约束来布局它们的头缘，这样 Auto Layout 才会正确的产生正确的约束。\*\*

接下来学习新的概念！**Metrics**

### Metrics

Metrics 是一个能在 VFL 格式化串中出现的以 number 为 value 的字典。如果你需要让距离变得标准化或者有些距离需要计算所以不能直接放在格式化串中的话，Metrics 将会变得非常有用！
将如下常量声明在**ViewController.swift**的变量之上：

```swift
// MARK: - Constants
private let horizontalPadding: CGFloat = 15.0
```

现在你有了一个用于 padding 的常量，你可以创建一个 metrics 字典并且将这个常量使用进去。将如下代码添加到**views**声明的上面：

```swift
let metrics = ["hp": horizontalPadding,
  "iconImageViewWidth": 30.0]
```

上面的代码创建的字典中的 key 可以再格式化串中使用。
接下来，用如下代码代替**topRowHorizontalConstraints**和**summaryHorizontalConstraints**的定义：

```swift
let horizontalConstraints = NSLayoutConstraint.constraintsWithVisualFormat(
  "H:|-hp-[iconImageView(iconImageViewWidth)]-[appNameLabel]-[skipButton]-hp-|",
  options: [.AlignAllCenterY],
  metrics: metrics,
  views: views)

let summaryHorizontalConstraints = NSLayoutConstraint.constraintsWithVisualFormat(
  "H:|-hp-[summaryLabel]-hp-|",
  options: [],
  metrics: metrics,
  views: views)
```

现在你已经将格式化串中得硬代码用 metrics 字典中 keys 代替掉了。
Auto Layout 可以进行串替换，将 metrics 字典中的 value 替换到格式化串中的 key。所以最终，**hp**将会被替换成 15pt，**iconImageViewWidth**将会被替换成 30pt。
你将一个重复出现的莫名其妙的数字变成了一个优雅的变量。如果你想要改变 padding，现在就只需要做一件事了。这不是更好吗？metrics 字典并不仅限制于常量；如果你需要在 runtime 期间进行计算，同样可以把这种变量放到 metrics 中。
最后的一点小问题是如果你想把这些元素放进**UINavigationController**或者**UITabBarController**中，那该怎么办呢？

### Layout Guides

视图控制器有两个可用的 Layout Guides：

- topLayoutGuide
- bottomLayoutGuide

它们都指定了试图控制器的视图中顶部或者底部导航栏边缘的位置，但是在**Grapevine**中，唯一的导航栏边缘是从状态栏开始的。
更新**iconVerticalConstraints**的声明代码：

```swift
let verticalConstraints = NSLayoutConstraint.constraintsWithVisualFormat(
  "V:|-[iconImageView(30)]",
  options: [],
  metrics: nil,
  views: views)
```

这样你就把状态栏和**iconImageView**之间的 20pt 的距离移除了，运行代码：
![remove](https://koenig-media.raywenderlich.com/uploads/2015/09/Grapevine-Without-20pts.png)
现在你的状态栏覆盖掉了视图上的一些元素。在横屏模式时，iOS 为了给小屏幕设备提供更多的有效空间移除状态栏，这样**iconImageView**会紧靠在屏幕的上方。
使用**topLayoutGuide**将会解决这种问题，用如下代码代替**views**字典：

```swift
  let views: [String: AnyObject] = ["iconImageView": iconImageView,
    "appNameLabel": appNameLabel,
    "skipButton": skipButton,
    "appImageView": appImageView,
    "welcomeLabel": welcomeLabel,
    "summaryLabel": summaryLabel,
    "pageControl": pageControl,
    "topLayoutGuide": topLayoutGuide,
    "bottomLayoutGuide": bottomLayoutGuide]
```

这次增加了**topLayoutGuide**和**bottomLayoutGuide**，它们继承自**UILayoutSupport**，比不是**UIView**。
接下来，就可以使用 layout guides 去对齐界面元素了。更新**iconVerticalConstraints**的声明：

```swift
let verticalConstraints = NSLayoutConstraint.constraintsWithVisualFormat(
  "V:[topLayoutGuide]-[iconImageView(30)]",
  options: [],
  metrics: nil,
  views: views)
```

接下来运行工程，完美！  
![PERFECT](https://koenig-media.raywenderlich.com/uploads/2015/09/Grapevine-Final.png)
现在你的顶部的界面元素都依赖着**topLayoutGuide**布局并且无论在横屏或者竖屏模式下状态栏的展现都控制着布局。
在这一节，你已经学会了当界面存在状态栏的时候如何利用**topLayoutGuide**来控制界面元素的布局。如果你的视图控制器在**UINavigationController**中，**topLayoutGuide**将会包含状态栏和**UINavigationBar**的状态。同时，如果你的试图控制器在**UITabBarController**中，**bottomLayoutGuide**将会提供底部边缘的状态。

### 限制

VFL 让你用一行代码写出了多个约束，大大降低了手指的负担。但是对于当前的实现，还存在一些限制；还有一些重要的东西需要理解。

- 视图中心
- 使用约束中的 Multiplier

#### 视图中心

在**Grapevine**中，你用了**.AlignAllCenterY**和**.AlignAllCenterX**。
使用这些表示你让一些视图和其他的一些视图的垂直中心或者水平中心对齐，然而只有在这些视图中存在已经有足够约束能够确定它们的水平和垂直中心位置的时候才能变得有效。
即使现在通过 VFL 你可以用一些小把戏来处理中心视图，但是这也不保证在将来的版本中依然有效。
![CONSTRAIN](https://koenig-media.raywenderlich.com/uploads/2015/08/constraints_constraining.png)

####使用约束中的 Multiplier
通过 Multiplier，你可以通过比例来对视图进行布局，比如你可以让一个 label 的宽度是它父视图的 60%。由于 VFL 会同时创建多个没有名字的约束，所以不能通过格式化串来设置百分比系数。

> ** 注意：你可以通过**constraintsWithVisualFormat**返回的数组来遍历约束，但是你需要去确定它们的**NSLayoutAttribute**属性，这样才能正确的设定 Multiplier，但即使是这样，你依然需要替换这些约束，因为约束的 Multiplier 是不可变的。**

### 现在要干什么？

你可以下载完整的[工程](http://cdn3.raywenderlich.com/wp-content/uploads/2015/09/Grapevine-Final.zip)。

> ** 注意：如果你有多个工程使用相同的 bundle id，Xcode 可能会出现问题。所以如果你完成了这个教程并且想最后运行一下刚才下载的工程，你可以使用**shift+option+command+K**清空一下 build 目录。**

现在你已经知道 VFL 如何工作啦，你已经可以在你的界面中使用这种布局咯。
你已经知道了如何使用 layout options 来减少需要定义的约束。你也已经知道如何使用 metrics 来在 runtime 定义距离而不仅仅是编译期。最后，你也知道了 VFL 的一些限制，但是利大于弊，你应该好好的利用它。
如果你对该教程或者 Auto Layout 有什么问题或者建议的话，请留言！
