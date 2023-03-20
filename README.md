# 设计模式golang版本
## 介绍
&emsp;&emsp;在开始使用go语言开发时，因为go没有显示的继承语法，想要实现一些设计模式时总是感到非常别扭，创建这个项目的初衷是希望通过
用go语言实现各种设计模式来提升对go语言的理解，同时也想提供一个文档，让大家在使用go语言实现设计模式有疑问时，通过这个项目能得到一些启发。
<br/>
&emsp;&emsp;希望在github这个平台上跟大家一起交流go语言的使用心得。
## 设计模式简介
&emsp;&emsp;设计模式（Design pattern）代表了最佳的实践，通常被有经验的面向对象的软件开发人员所采用。设计模式是软件开发人员在软件开发过程中
面临的一般问题的解决方案。这些解决方案是众多软件开发人员经过相当长的一段时间的试验和错误总结出来的。
<br/>
## 什么是 GOF（四人帮，全拼 Gang of Four）？
在 1994 年，由 Erich Gamma、Richard Helm、Ralph Johnson 和 John Vlissides 四人合著出版了一本名为 Design Patterns - Elements 
of Reusable Object-Oriented Software（中文译名：设计模式 - 可复用的面向对象软件元素） 的书，该书首次提到了软件开发中设计模式的概念。
四位作者合称 GOF（四人帮，全拼 Gang of Four）。他们所提出的设计模式主要是基于以下的面向对象设计原则。
* 对接口编程而不是对实现编程。
* 优先使用对象组合而不是继承。
<br/>
## 设计模式的六大原则
### 开闭原则（Open Close Principle）

开闭原则的意思是：对扩展开放，对修改关闭。在程序需要进行拓展的时候，不能去修改原有的代码，实现一个热插拔的效果。简言之，是为了使程序的扩展性好，易于维护和升级。想要达到这样的效果，我们需要使用接口和抽象类，后面的具体设计中我们会提到这点。

### 里氏代换原则（Liskov Substitution Principle）

里氏代换原则是面向对象设计的基本原则之一。 里氏代换原则中说，任何基类可以出现的地方，子类一定可以出现。LSP 是继承复用的基石，只有当派生类可以替换掉基类，且软件单位的功能不受到影响时，基类才能真正被复用，而派生类也能够在基类的基础上增加新的行为。里氏代换原则是对开闭原则的补充。实现开闭原则的关键步骤就是抽象化，而基类与子类的继承关系就是抽象化的具体实现，所以里氏代换原则是对实现抽象化的具体步骤的规范。

### 依赖倒转原则（Dependence Inversion Principle）

这个原则是开闭原则的基础，具体内容：针对接口编程，依赖于抽象而不依赖于具体。

### 接口隔离原则（Interface Segregation Principle）

这个原则的意思是：使用多个隔离的接口，比使用单个接口要好。它还有另外一个意思是：降低类之间的耦合度。由此可见，其实设计模式就是从大型软件架构出发、便于升级和维护的软件设计思想，它强调降低依赖，降低耦合。

### 迪米特法则，又称最少知道原则（Demeter Principle）

最少知道原则是指：一个实体应当尽量少地与其他实体之间发生相互作用，使得系统功能模块相对独立。

### 接口隔离原则（Interface Segregation Principle，ISP）

尽量将臃肿庞大的接口拆分成更小的和更具体的接口，让接口中只包含客户感兴趣的方法
要为各个类建立它们需要的专用接口，而不要试图去建立一个很庞大的接口供所有依赖它的类去调用。

接口隔离原则和单一职责都是为了提高类的内聚性、降低它们之间的耦合性，体现了封装的思想，但两者是不同的：

单一职责原则注重的是职责，而接口隔离原则注重的是对接口依赖的隔离。

单一职责原则主要是约束类，它针对的是程序中的实现和细节；接口隔离原则主要约束接口，主要针对抽象和程序整体框架的构建。

## 设计模式汇总
根据设计模式的参考书 Design Patterns - Elements of Reusable Object-Oriented Software（中文译名：设计模式 - 可复用的面向对象软件元素） 
中所提到的，总共有 23 种设计模式。这些模式可以分为三大类：创建型模式（Creational Patterns）、结构型模式（Structural Patterns）、
行为型模式（Behavioral Patterns）。

### 创建型模式
创建型设计模式提供了一种在创建对象的同时隐藏创建逻辑的方式，而不是使用 new 运算符直接实例化对象。这使得程序在判断针对某个给定实例需要创建哪些对象时更加灵活。
* <a href="https://github.com/caixunshi/go-design-patterns/tree/main/creational/factory">工厂模式（Factory Pattern）</a>
* <a href="https://github.com/caixunshi/go-design-patterns/tree/main/creational/abstract-factory">抽象工厂模式（Abstract Factory Pattern）</a>
* <a href="https://github.com/caixunshi/go-design-patterns/tree/main/creational/singleton">单例模式（Singleton Pattern）</a>
* <a href="https://github.com/caixunshi/go-design-patterns/tree/main/creational/builder">建造者模式（Builder Pattern）</a>
* <a href="https://github.com/caixunshi/go-design-patterns/tree/main/creational/prototype">原型模式（Prototype Pattern）</a>

### 结构型模式
结构型设计模式关注类和对象的组合。继承的概念被用来组合接口和定义组合对象获得新功能的方式。
* <a href="">适配器模式（Adapter Pattern）</a>
* <a href="">桥接模式（Bridge Pattern）</a>
* <a href="">过滤器模式（Filter、Criteria Pattern）</a>
* <a href="">组合模式（Composite Pattern）</a>
* <a href="">装饰器模式（Decorator Pattern）</a>
* <a href="">外观模式（Facade Pattern）</a>
* <a href="">享元模式（Flyweight Pattern）</a>
* <a href="">代理模式（Proxy Pattern）</a>

### 行为型模式
行为型模式不仅仅关注类和对象的结构，而且重点关注它们之间的相互作用。 通过行为型模式，可以更加清晰地划分类与对象的职责，并研究系统在运行时实例对象之间的交互。在系统运行时，对象并不是孤 立的，它们可以通过相互通信与协作完成某些复杂功能，一个对象在运行时也将影响到其他对象的运行。
* <a href="">责任链模式（Chain of Responsibility Pattern）</a>
* <a href="">命令模式（Command Pattern）</a>
* <a href="">解释器模式（Interpreter Pattern）</a>
* <a href="">迭代器模式（Iterator Pattern）</a>
* <a href="">中介者模式（Mediator Pattern）</a>
* <a href="">备忘录模式（Memento Pattern）</a>
* <a href="">观察者模式（Observer Pattern）</a>
* <a href="">状态模式（State Pattern）</a>
* <a href="">空对象模式（Null Object Pattern）</a>
* <a href="">策略模式（Strategy Pattern）</a>
* <a href="">模板模式（Template Pattern）</a>
* <a href="">访问者模式（Visitor Pattern）</a>
## 总结
设计原则重道，设计模式重术，设计模式是设计原则的具体实现
